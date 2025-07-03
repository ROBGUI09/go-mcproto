package models

import "github.com/ROBGUI09/go-mcproto/packets"

type LoginSuccessPacket struct {
	// *packets.CompressedPacket
	packets.MinecraftPacket

	UUID     []byte `mc:"bytes" len:"16"`
	Username string `mc:"string"`
}

