package client

// MessageID 0x0b
// Name SID_GETCHANNELLIST

// GetChannelList is the structure of a SID_GETCHANNELLIST request
type GetChannelList struct {
	ProductID uint32 // This value is ignored and can be set to 0x00
}
