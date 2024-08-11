package zfoogo


type UdpHelloRequest struct {
	Message string
}

func (protocol UdpHelloRequest) ProtocolId() int16 {
	return 1200
}

func (protocol UdpHelloRequest) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*UdpHelloRequest)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol UdpHelloRequest) read(buffer *ByteBuffer) any {
	var packet = new(UdpHelloRequest)
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