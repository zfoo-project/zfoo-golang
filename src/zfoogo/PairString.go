package zfoogo

type PairString struct {
	Key   string
	Value string
}

func (protocol PairString) ProtocolId() int16 {
	return 112
}

func (protocol PairString) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*PairString)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Key)
	buffer.WriteString(message.Value)
}

func (protocol PairString) read(buffer *ByteBuffer) any {
	var packet = new(PairString)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadString()
	packet.Key = result0
	var result1 = buffer.ReadString()
	packet.Value = result1
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
