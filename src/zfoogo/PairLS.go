package zfoogo


type PairLS struct {
	Key int64
	Value string
}

func (protocol PairLS) ProtocolId() int16 {
	return 113
}

func (protocol PairLS) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*PairLS)
	buffer.WriteInt(-1)
	buffer.WriteLong(message.Key)
	buffer.WriteString(message.Value)
}

func (protocol PairLS) read(buffer *ByteBuffer) any {
	var packet = new(PairLS)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadLong()
	packet.Key = result0
	var result1 = buffer.ReadString()
	packet.Value = result1
	if length > 0 {
        buffer.SetReadOffset(beforeReadIndex + length)
    }
	return packet
}