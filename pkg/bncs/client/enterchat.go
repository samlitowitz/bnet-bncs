package client

// MessageID 0x0a
// Name SID_ENTERCHAT

// EnterChat is the structure of a SID_ENTERCHAT request
type EnterChat struct {
	Username   string
	Statstring string
}
