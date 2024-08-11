package zfoogo


type TripleLSS struct {
	Left int64
	Middle string
	Right string
}

func (protocol TripleLSS) ProtocolId() int16 {
	return 116
}

func (protocol TripleLSS) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*TripleLSS)
	buffer.WriteInt(-1)
	buffer.WriteLong(message.Left)
	buffer.WriteString(message.Middle)
	buffer.WriteString(message.Right)
}

func (protocol TripleLSS) read(buffer *ByteBuffer) any {
	var packet = new(TripleLSS)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadLong()
	packet.Left = result0
	var result1 = buffer.ReadString()
	packet.Middle = result1
	var result2 = buffer.ReadString()
	packet.Right = result2
	if length > 0 {
        buffer.SetReadOffset(beforeReadIndex + length)
    }
	return packet
}