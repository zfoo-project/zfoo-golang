package zfoogo


type TcpHelloRequest struct {
	Message string
}

func (protocol TcpHelloRequest) ProtocolId() int16 {
	return 1300
}

func (protocol TcpHelloRequest) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*TcpHelloRequest)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol TcpHelloRequest) read(buffer *ByteBuffer) any {
	var packet = new(TcpHelloRequest)
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