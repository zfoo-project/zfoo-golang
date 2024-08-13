package zfoogo

type WebsocketHelloRequest struct {
	Message string
}

func (protocol WebsocketHelloRequest) ProtocolId() int16 {
	return 1400
}

func (protocol WebsocketHelloRequest) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*WebsocketHelloRequest)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol WebsocketHelloRequest) read(buffer *ByteBuffer) any {
	var packet = new(WebsocketHelloRequest)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadString()
	packet.Message = result0
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
