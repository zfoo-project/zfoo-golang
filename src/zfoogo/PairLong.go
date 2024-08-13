package zfoogo

type PairLong struct {
	Key   int64
	Value int64
}

func (protocol PairLong) ProtocolId() int16 {
	return 111
}

func (protocol PairLong) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*PairLong)
	buffer.WriteInt(-1)
	buffer.WriteLong(message.Key)
	buffer.WriteLong(message.Value)
}

func (protocol PairLong) read(buffer *ByteBuffer) any {
	var packet = new(PairLong)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadLong()
	packet.Key = result0
	var result1 = buffer.ReadLong()
	packet.Value = result1
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
