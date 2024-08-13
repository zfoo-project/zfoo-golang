package zfoogo

type GatewayToProviderResponse struct {
	Message string
}

func (protocol GatewayToProviderResponse) ProtocolId() int16 {
	return 5001
}

func (protocol GatewayToProviderResponse) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
		buffer.WriteInt(0)
		return
	}
	var message = packet.(*GatewayToProviderResponse)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol GatewayToProviderResponse) read(buffer *ByteBuffer) any {
	var packet = new(GatewayToProviderResponse)
	var length = buffer.ReadInt()
	if length == 0 {
		return packet
	}
	var beforeReadIndex = buffer.GetReadOffset()
	var result0 = buffer.ReadString()
	packet.Message = result0
	if length > 0 {
		buffer.SetReadOffset(beforeReadIndex + length)
	}
	return packet
}
