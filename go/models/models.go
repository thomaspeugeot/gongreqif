package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type REQIF struct {
	Name string

	REQ_IF_HEADER *REQ_IF_HEADER
}

func (reqif *REQIF) Walk(_reqif *schema.REQ_IF, stage *StageStruct) {
	log.Println("reqif:Walk")

	reqif.REQ_IF_HEADER = new(REQ_IF_HEADER).Stage(stage)
	reqif.REQ_IF_HEADER.Name = "Parsed REQIF Header"
	reqif.REQ_IF_HEADER.COMMENT = _reqif.THE_HEADER.REQ_IF_HEADER.COMMENT
	reqif.REQ_IF_HEADER.CREATION_TIME = _reqif.THE_HEADER.REQ_IF_HEADER.CREATION_TIME.ToGoTime()
	reqif.REQ_IF_HEADER.REPOSITORY_ID = _reqif.THE_HEADER.REQ_IF_HEADER.REPOSITORY_ID
	reqif.REQ_IF_HEADER.REQ_IF_TOOL_ID = _reqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_TOOL_ID
	reqif.REQ_IF_HEADER.REQ_IF_VERSION = _reqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_VERSION
	reqif.REQ_IF_HEADER.SOURCE_TOOL_ID = _reqif.THE_HEADER.REQ_IF_HEADER.SOURCE_TOOL_ID
	reqif.REQ_IF_HEADER.TITLE = _reqif.THE_HEADER.REQ_IF_HEADER.TITLE
}

type REQ_IF_HEADER struct {
	Name    string
	COMMENT string `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`

	CREATION_TIME time.Time `xml:"CREATION-TIME,omitempty" json:"CREATION-TIME,omitempty"`

	REPOSITORY_ID string `xml:"REPOSITORY-ID,omitempty" json:"REPOSITORY-ID,omitempty"`

	REQ_IF_TOOL_ID string `xml:"REQ-IF-TOOL-ID,omitempty" json:"REQ-IF-TOOL-ID,omitempty"`

	REQ_IF_VERSION string `xml:"REQ-IF-VERSION,omitempty" json:"REQ-IF-VERSION,omitempty"`

	SOURCE_TOOL_ID string `xml:"SOURCE-TOOL-ID,omitempty" json:"SOURCE-TOOL-ID,omitempty"`

	TITLE string `xml:"TITLE,omitempty" json:"TITLE,omitempty"`

	// IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`
}
