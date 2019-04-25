package client

// MessageID 0x51
// Name SID_AUTH_CHECK

// AuthCheckKey is the key structure of a SID_AUTH_CHECK request
type AuthCheckKey struct {
	Length       uint32
	ProductValue uint32
	PublicValue  uint32
	Unknown      uint32
	HashedData   [20]uint8
}

// AuthCheck is the structure of a SID_AUTH_CHECK request
type AuthCheck struct {
	ClientToken uint32
	ExeVersion  uint32
	ExeHash     uint32
	KeysCount   uint32 `bnet:"save-KeysCount"`
	SpawnKey    bool   `bnet:"size-uint32"`

	Keys []AuthCheckKey `bnet:"len-KeysCount"`

	ExeInfo  string
	KeyOwner string
}
