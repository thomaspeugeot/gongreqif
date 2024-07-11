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

type REQIFHEADER_WOP struct {
	// insertion point
	Name string
	IDENTIFIERAttr string
	COMMENT string
	CREATIONTIME string
	REPOSITORYID string
	REQIFTOOLID string
	REQIFVERSION string
	SOURCETOOLID string
	TITLE string
}

func (from *REQIFHEADER) CopyBasicFields(to *REQIFHEADER) {
	// insertion point
	to.Name = from.Name
	to.IDENTIFIERAttr = from.IDENTIFIERAttr
	to.COMMENT = from.COMMENT
	to.CREATIONTIME = from.CREATIONTIME
	to.REPOSITORYID = from.REPOSITORYID
	to.REQIFTOOLID = from.REQIFTOOLID
	to.REQIFVERSION = from.REQIFVERSION
	to.SOURCETOOLID = from.SOURCETOOLID
	to.TITLE = from.TITLE
}

