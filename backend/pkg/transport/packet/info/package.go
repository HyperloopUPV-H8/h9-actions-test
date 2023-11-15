package info

import (
	"github.com/HyperloopUPV-H8/h9-backend/pkg/abstraction"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/transport/packet"
)

// infoData is the contents of an info Packet
type infoData string

// Packet is an info packet.
//
// info packets are used to send arbitrary strings for debugging purposes.
type Packet struct {
	id        abstraction.PacketId
	BoardId   abstraction.BoardId `json:"boardId"`
	Timestamp packet.Timestamp    `json:"timestamp"`
	Msg       infoData            `json:"msg"`
}

// Id returns the packet id
func (packet *Packet) Id() abstraction.PacketId {
	return packet.id
}
