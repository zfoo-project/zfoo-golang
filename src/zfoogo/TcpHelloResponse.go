package zfoogo

type TcpHelloResponse struct {
	Message string
}

func (protocol TcpHelloResponse) ProtocolId() int16 {
	return 1301
}

func (protocol TcpHelloResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*TcpHelloResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol TcpHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(TcpHelloResponse)
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
