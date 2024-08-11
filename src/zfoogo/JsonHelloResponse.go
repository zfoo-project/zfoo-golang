package zfoogo


type JsonHelloResponse struct {
	Message string
}

func (protocol JsonHelloResponse) ProtocolId() int16 {
	return 1601
}

func (protocol JsonHelloResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*JsonHelloResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol JsonHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(JsonHelloResponse)
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