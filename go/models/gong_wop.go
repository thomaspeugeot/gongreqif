// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type REQIF_WOP struct {
	// insertion point
	Name string
}

func (from *REQIF) CopyBasicFields(to *REQIF) {
	// insertion point
	to.Name = from.Name
}

type REQ_IF_CONTENT_WOP struct {
	// insertion point
	Name string
}

func (from *REQ_IF_CONTENT) CopyBasicFields(to *REQ_IF_CONTENT) {
	// insertion point
	to.Name = from.Name
}

type REQ_IF_HEADER_WOP struct {
	// insertion point
	Name string
	COMMENT string
	CREATION_TIME time.Time
	REPOSITORY_ID string
	REQ_IF_TOOL_ID string
	REQ_IF_VERSION string
	SOURCE_TOOL_ID string
	TITLE string
}

func (from *REQ_IF_HEADER) CopyBasicFields(to *REQ_IF_HEADER) {
	// insertion point
	to.Name = from.Name
	to.COMMENT = from.COMMENT
	to.CREATION_TIME = from.CREATION_TIME
	to.REPOSITORY_ID = from.REPOSITORY_ID
	to.REQ_IF_TOOL_ID = from.REQ_IF_TOOL_ID
	to.REQ_IF_VERSION = from.REQ_IF_VERSION
	to.SOURCE_TOOL_ID = from.SOURCE_TOOL_ID
	to.TITLE = from.TITLE
}

