package client

// MessageID 0x0c
// Name SID_JOINCHANNEL

// JoinChannel is the structure of a SID_JOINCHANNEL request
type JoinChannel struct {
	Flags   uint32
	Channel string
}
