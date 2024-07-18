package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

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

func (reqif *REQ_IF_HEADER) Walk(_reqif *schema.REQ_IF_HEADER, stage *StageStruct) {
	log.Println("REQ_IF_HEADER", "Walk")

	reqif.Name = time.Now().Format(time.DateTime)
	reqif.COMMENT = _reqif.COMMENT
	reqif.CREATION_TIME = _reqif.CREATION_TIME.ToGoTime()
	reqif.REPOSITORY_ID = _reqif.REPOSITORY_ID
	reqif.REQ_IF_TOOL_ID = _reqif.REQ_IF_TOOL_ID
	reqif.REQ_IF_VERSION = _reqif.REQ_IF_VERSION
	reqif.SOURCE_TOOL_ID = _reqif.SOURCE_TOOL_ID
	reqif.TITLE = _reqif.TITLE
}
