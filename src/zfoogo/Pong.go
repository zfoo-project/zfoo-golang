package zfoogo


type Pong struct {
	Time int64
}

func (protocol Pong) ProtocolId() int16 {
	return 104
}

func (protocol Pong) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*Pong)
	buffer.WriteInt(-1)
	buffer.WriteLong(message.Time)
}

func (protocol Pong) read(buffer *ByteBuffer) any {
	var packet = new(Pong)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadLong()
	packet.Time = result0
	if length > 0 {
        buffer.SetReadOffset(beforeReadIndex + length)
    }
	return packet
}