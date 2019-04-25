package client

// MessageID 0x50
// Name SID_AUTH_INFO

// AuthInfo is the structure of a SID_AUTH_INFO request
type AuthInfo struct {
	ProtocolID     uint32
	PlatformCode   uint32
	ProductCode    uint32
	Version        uint32
	LanguageCode   uint32
	LocalIP        uint32
	TimeZoneBias   uint32
	MPQLocaleID    uint32
	UserLanguageID uint32
	CountryAb      string
	Country        string
}
