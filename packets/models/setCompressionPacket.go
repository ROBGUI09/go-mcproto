package models

import "github.com/ROBGUI09/go-mcproto/packets"

type SetCompressionPacket struct {
	// *packets.UncompressedPacket
	packets.MinecraftPacket

	Treshold int32 `mc:"varint"`
}

