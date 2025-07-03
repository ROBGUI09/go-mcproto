package models

import "github.com/ROBGUI09/go-mcproto/packets"

type DisconnectPacket struct {
	packets.MinecraftPacket

	Reason string `mc:"string"`
}

