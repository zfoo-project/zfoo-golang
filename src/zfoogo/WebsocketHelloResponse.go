package zfoogo

type WebsocketHelloResponse struct {
	Message string
}

func (protocol WebsocketHelloResponse) ProtocolId() int16 {
	return 1401
}

func (protocol WebsocketHelloResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*WebsocketHelloResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol WebsocketHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(WebsocketHelloResponse)
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
