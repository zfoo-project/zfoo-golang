/*
 * Copyright (C) 2020 The zfoo Authors
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */
package net

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"net"
	"sync/atomic"
)

// Session struct
type Session struct {
	sid uint64
	uid uint64

	conn        net.Conn
	sendChan    chan []byte
	messageChan chan any
	doneChan    chan error
}

var uuid uint64

// NewSession create a new session
func NewSession(conn net.Conn) *Session {
	var suuid = atomic.AddUint64(&uuid, 1)

	session := &Session{
		sid:         suuid,
		uid:         0, // 可以为用户的id
		conn:        conn,
		sendChan:    make(chan []byte, 100),
		doneChan:    make(chan error),
		messageChan: make(chan any, 100),
	}

	return session
}

// Close close connection
func (session *Session) Close() {
	session.conn.Close()
}

// SendMessage send message
func (session *Session) SendMessage(msg any) error {
	var buffer = Encode(msg)
	session.sendChan <- buffer.ToBytes()
	return nil
}

// writeCoroutine write coroutine
func (session *Session) writeCoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case packet := <-session.sendChan:
			if packet == nil {
				continue
			}

			if _, err := session.conn.Write(packet); err != nil {
				session.doneChan <- err
			}
		}
	}
}

// readCoroutine read coroutine
func (session *Session) readCoroutine(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return

		default:
			// 读取长度
			var bufferHeader = make([]byte, 4)
			_, err := io.ReadFull(session.conn, bufferHeader)
			if err != nil {
				session.doneChan <- err
				continue
			}

			reader := bytes.NewReader(bufferHeader)

			var length int32
			err = binary.Read(reader, binary.BigEndian, &length)
			if err != nil {
				session.doneChan <- err
				continue
			}

			// 读取数据
			var bufferBody = make([]byte, length)
			_, err = io.ReadFull(session.conn, bufferBody)
			if err != nil {
				session.doneChan <- err
				continue
			}

			// 解码
			var packet = Decode(bufferBody)
			session.messageChan <- packet
		}
	}
}
