package zfoogo


type SignalAttachment struct {
	SignalId int
	TaskExecutorHash int
	// 0 for the server, 1 or 2 for the sync or async native client, 12 for the outside client such as browser, mobile
	Client int8
	Timestamp int64
}

func (protocol SignalAttachment) ProtocolId() int16 {
	return 0
}

func (protocol SignalAttachment) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*SignalAttachment)
	buffer.WriteInt(-1)
	buffer.WriteByte(message.Client)
	buffer.WriteInt(message.SignalId)
	buffer.WriteInt(message.TaskExecutorHash)
	buffer.WriteLong(message.Timestamp)
}

func (protocol SignalAttachment) read(buffer *ByteBuffer) any {
	var packet = new(SignalAttachment)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadByte()
	packet.Client = result0
	var result1 = buffer.ReadInt()
	packet.SignalId = result1
	var result2 = buffer.ReadInt()
	packet.TaskExecutorHash = result2
	var result3 = buffer.ReadLong()
	packet.Timestamp = result3
	if length > 0 {
        buffer.SetReadOffset(beforeReadIndex + length)
    }
	return packet
}