package client

// MessageID 0x29
// Name SID_LOGONRESPONSE2

// LogonResponse2 is the structure of a SID_LOGONRESPONSE2 request
type LogonResponse2 struct {
	ClientToken  uint32
	ServerToken  uint32
	PasswordHash [20]uint8
	Username     string
}
