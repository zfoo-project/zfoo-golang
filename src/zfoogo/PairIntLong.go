package zfoogo

type PairIntLong struct {
	Key   int
	Value int64
}

func (protocol PairIntLong) ProtocolId() int16 {
	return 110
}

func (protocol PairIntLong) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*PairIntLong)
	buffer.WriteInt(-1)
	buffer.WriteInt(message.Key)
	buffer.WriteLong(message.Value)
}

func (protocol PairIntLong) read(buffer *ByteBuffer) any {
	var packet = new(PairIntLong)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadInt()
	packet.Key = result0
	var result1 = buffer.ReadLong()
	packet.Value = result1
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
