package client

// MessageID 0x33
// Name SID_GETFILETIME

// GetFiletime is the structure of a SID_GETFILETIME request
type GetFiletime struct {
	ID       uint32
	Unknown  uint32
	Filename string
}
