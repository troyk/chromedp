package log

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

// EventEntryAdded issued when new message was logged.
type EventEntryAdded struct {
	Entry *LogEntry `json:"entry,omitempty"` // The entry.
}

// EventTypes is all event types in the domain.
var EventTypes = []MethodType{
	EventLogEntryAdded,
}