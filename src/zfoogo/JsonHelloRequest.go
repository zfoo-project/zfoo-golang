package zfoogo

type JsonHelloRequest struct {
	Message string
}

func (protocol JsonHelloRequest) ProtocolId() int16 {
	return 1600
}

func (protocol JsonHelloRequest) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*JsonHelloRequest)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol JsonHelloRequest) read(buffer *ByteBuffer) any {
	var packet = new(JsonHelloRequest)
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
