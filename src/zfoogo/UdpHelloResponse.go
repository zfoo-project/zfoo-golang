package zfoogo

type UdpHelloResponse struct {
	Message string
}

func (protocol UdpHelloResponse) ProtocolId() int16 {
	return 1201
}

func (protocol UdpHelloResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*UdpHelloResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol UdpHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(UdpHelloResponse)
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
