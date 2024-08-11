package zfoogo


type WebSocketObjectA struct {
	A int
	ObjectB WebSocketObjectB
}

func (protocol WebSocketObjectA) ProtocolId() int16 {
	return 2071
}

func (protocol WebSocketObjectA) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*WebSocketObjectA)
	buffer.WriteInt(-1)
	buffer.WriteInt(message.A)
	buffer.WritePacket(&message.ObjectB, 2072)
}

func (protocol WebSocketObjectA) read(buffer *ByteBuffer) any {
	var packet = new(WebSocketObjectA)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadInt()
	packet.A = result0
	var result1 = *buffer.ReadPacket(2072).(*WebSocketObjectB)
	packet.ObjectB = result1
	if length > 0 {
        buffer.SetReadOffset(beforeReadIndex + length)
    }
	return packet
}