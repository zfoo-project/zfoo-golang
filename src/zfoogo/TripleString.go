package zfoogo


type TripleString struct {
	Left string
	Middle string
	Right string
}

func (protocol TripleString) ProtocolId() int16 {
	return 115
}

func (protocol TripleString) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*TripleString)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Left)
	buffer.WriteString(message.Middle)
	buffer.WriteString(message.Right)
}

func (protocol TripleString) read(buffer *ByteBuffer) any {
	var packet = new(TripleString)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadString()
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