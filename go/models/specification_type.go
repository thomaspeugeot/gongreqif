package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPECIFICATION_TYPE struct {
	Name string

	// ALTERNATIVE_ID struct {
	// 	ALTERNATIVE_ID *ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	// } `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	// SPEC_ATTRIBUTES struct {
	// 	ATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

	// 	ATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty" json:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

	// 	ATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

	// 	ATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

	// 	ATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty" json:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

	// 	ATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty" json:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

	// 	ATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
	// } `xml:"SPEC-ATTRIBUTES,omitempty" json:"SPEC-ATTRIBUTES,omitempty"`

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
	}
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME
}
