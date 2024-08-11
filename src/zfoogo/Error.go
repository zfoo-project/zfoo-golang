package zfoogo


type Error struct {
	Code int
	Message string
}

func (protocol Error) ProtocolId() int16 {
	return 101
}

func (protocol Error) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*Error)
	buffer.WriteInt(-1)
	buffer.WriteInt(message.Code)
	buffer.WriteString(message.Message)
}

func (protocol Error) read(buffer *ByteBuffer) any {
	var packet = new(Error)
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