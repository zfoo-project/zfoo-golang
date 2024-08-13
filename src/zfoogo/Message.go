package zfoogo

type Message struct {
	Code    int
	Message string
}

func (protocol Message) ProtocolId() int16 {
	return 100
}

func (protocol Message) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*Message)
	buffer.WriteInt(-1)
	buffer.WriteInt(message.Code)
	buffer.WriteString(message.Message)
}

func (protocol Message) read(buffer *ByteBuffer) any {
	var packet = new(Message)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadInt()
	packet.Code = result0
	var result1 = buffer.ReadString()
	packet.Message = result1
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
