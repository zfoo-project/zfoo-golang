package zfoogo

type WebSocketObjectB struct {
	Flag bool
}

func (protocol WebSocketObjectB) ProtocolId() int16 {
	return 2072
}

func (protocol WebSocketObjectB) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*WebSocketObjectB)
	buffer.WriteInt(-1)
	buffer.WriteBool(message.Flag)
}

func (protocol WebSocketObjectB) read(buffer *ByteBuffer) any {
	var packet = new(WebSocketObjectB)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadBool()
	packet.Flag = result0
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
