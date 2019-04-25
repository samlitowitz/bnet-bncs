package client

// MessageID 0x3e
// Name SID_LOGONREALMEX

// LogonRealmEx is the structure of a SID_LOGONREALMEX request
type LogonRealmEx struct {
	ClientToken   uint32
	RealmPassword [20]uint8 // Hashed realm password
	Title         string
}
