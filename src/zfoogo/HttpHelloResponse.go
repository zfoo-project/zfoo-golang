package zfoogo

type HttpHelloResponse struct {
	Message string
}

func (protocol HttpHelloResponse) ProtocolId() int16 {
	return 1701
}

func (protocol HttpHelloResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*HttpHelloResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol HttpHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(HttpHelloResponse)
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
