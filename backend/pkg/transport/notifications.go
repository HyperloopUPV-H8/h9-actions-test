package transport

import (
	"github.com/HyperloopUPV-H8/h9-backend/pkg/abstraction"
	"time"
)

// PacketNotification notifies of an incoming message
type PacketNotification struct {
	Packet    abstraction.Packet
	Raw       []byte
	From      string
	To        string
	Timestamp time.Time
}

func NewPacketNotification(packet abstraction.Packet, raw []byte, from string, to string, timestamp time.Time) PacketNotification {
	return PacketNotification{
		Packet:    packet,
		Raw:       raw,
		From:      from,
		To:        to,
		Timestamp: timestamp,
	}
}

// Event always returns PacketEvent
func (notification PacketNotification) Event() abstraction.TransportEvent {
	return PacketEvent
}
