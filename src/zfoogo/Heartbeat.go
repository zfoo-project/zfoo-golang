package zfoogo

type Heartbeat struct {
}

func (protocol Heartbeat) ProtocolId() int16 {
	return 102
}

func (protocol Heartbeat) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	buffer.WriteInt(-1)
}

func (protocol Heartbeat) read(buffer *ByteBuffer) any {
	var packet = new(Heartbeat)
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
