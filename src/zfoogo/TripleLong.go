package zfoogo

type TripleLong struct {
	Left   int64
	Middle int64
	Right  int64
}

func (protocol TripleLong) ProtocolId() int16 {
	return 114
}

func (protocol TripleLong) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*TripleLong)
	buffer.WriteInt(-1)
	buffer.WriteLong(message.Left)
	buffer.WriteLong(message.Middle)
	buffer.WriteLong(message.Right)
}

func (protocol TripleLong) read(buffer *ByteBuffer) any {
	var packet = new(TripleLong)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadLong()
	packet.Left = result0
	var result1 = buffer.ReadLong()
	packet.Middle = result1
	var result2 = buffer.ReadLong()
	packet.Right = result2
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
