package zfoogo


type GatewayToProviderRequest struct {
	Message string
}

func (protocol GatewayToProviderRequest) ProtocolId() int16 {
	return 5000
}

func (protocol GatewayToProviderRequest) write(buffer *ByteBuffer, packet any) {
	if packet == nil {
	    buffer.WriteInt(0)
		return
	}
	var message = packet.(*GatewayToProviderRequest)
	buffer.WriteInt(-1)
	buffer.WriteString(message.Message)
}

func (protocol GatewayToProviderRequest) read(buffer *ByteBuffer) any {
	var packet = new(GatewayToProviderRequest)
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