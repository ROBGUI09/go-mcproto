package models

import "github.com/ROBGUI09/go-mcproto/packets"

type KeepAlivePacket struct {
	// *packets.CompressedPacket
	packets.MinecraftPacket

	KeepAliveID int64 `mc:"inherit"`
}

