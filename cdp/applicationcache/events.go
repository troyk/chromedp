package applicationcache

// AUTOGENERATED. DO NOT EDIT.

import (
	. "github.com/knq/chromedp/cdp"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

type EventApplicationCacheStatusUpdated struct {
	FrameID     FrameID `json:"frameId,omitempty"`     // Identifier of the frame containing document whose application cache updated status.
	ManifestURL string  `json:"manifestURL,omitempty"` // Manifest URL.
	Status      int64   `json:"status,omitempty"`      // Updated application cache status.
}

type EventNetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline,omitempty"`
}

// EventTypes is all event types in the domain.
var EventTypes = []MethodType{
	EventApplicationCacheApplicationCacheStatusUpdated,
	EventApplicationCacheNetworkStateUpdated,
}