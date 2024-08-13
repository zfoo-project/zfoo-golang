package net

import (
	"context"
	"net"
	"sync"
)

// TcpServer struct
type TcpServer struct {
	onReceiver   func(*Session, any)
	onConnect    func(*Session)
	onDisconnect func(*Session, error)
	sessions     *sync.Map
	address      string
	listener     net.Listener
}

// NewTcpServer create a new socket service
func NewTcpServer(address string) *TcpServer {
	var listen, _ = net.Listen("tcp", address)
	var tcpServer = new(TcpServer)
	tcpServer.address = address
	tcpServer.sessions = new(sync.Map)
	tcpServer.listener = listen
	return tcpServer
}

// Start Start socket service
func (server *TcpServer) Start() {
	var ctx, cancel = context.WithCancel(context.Background())

	defer func() {
		cancel()
		server.listener.Close()
	}()

	for {
		var conn, _ = server.listener.Accept()

		// 将conn保存到全局的map中
		var session = NewSession(conn)
		server.sessions.Store(session.sid, session)

		go server.connectHandler(ctx, session)
	}
}

func (server *TcpServer) connectHandler(ctx context.Context, session *Session) {
	var conn_ctx, cancel = context.WithCancel(ctx)

	defer func() {
		cancel()
		session.Close()
		server.sessions.Delete(session.sid)
	}()

	go session.readCoroutine(conn_ctx)
	go session.writeCoroutine(conn_ctx)

	if server.onConnect != nil {
		server.onConnect(session)
	}
	for {
		select {
		case err := <-session.doneChan:
			if server.onDisconnect != nil {
				server.onDisconnect(session, err)
			}
			return
		case packet := <-session.messageChan:
			if server.onReceiver != nil {
				server.onReceiver(session, packet)
			}
		}
	}
}
