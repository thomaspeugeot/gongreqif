package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPECIFICATION_TYPE struct {
	Name string

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

func (reqif *SPECIFICATION_TYPE) Walk(_reqif *schema.SPECIFICATION_TYPE, stage *StageStruct) {
	log.Println("SPECIFICATION_TYPE", "Walk")

	reqif.DESC = _reqif.DESC
	if _reqif.IDENTIFIER != nil {
		reqif.IDENTIFIER = string(*_reqif.IDENTIFIER)
		reqif.Name = reqif.IDENTIFIER
	}
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME
}
