package models

import "github.com/ROBGUI09/go-mcproto/packets"

type LoginStartPacket struct {
	// *packets.UncompressedPacket
	packets.MinecraftPacket

	Name string `mc:"string"`
}

