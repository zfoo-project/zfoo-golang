package zfoogo

type Ping struct {
}

func (protocol Ping) ProtocolId() int16 {
	return 103
}

func (protocol Ping) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	buffer.WriteInt(-1)
}

func (protocol Ping) read(buffer *ByteBuffer) any {
	var packet = new(Ping)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
