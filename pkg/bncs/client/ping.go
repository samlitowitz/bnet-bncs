package client

// MessageID 0x25
// Name SID_PING

// Ping is the structure of a SID_PING request
type Ping struct {
	PingValue uint32
}
