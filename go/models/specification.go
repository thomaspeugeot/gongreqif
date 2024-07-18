package models

import (
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPECIFICATION struct {
	Name string
	// ALTERNATIVE_ID struct {
	// 	ALTERNATIVE_ID *ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	// } `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	// VALUES struct {
	// 	ATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty" json:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`

	// 	ATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty" json:"ATTRIBUTE-VALUE-DATE,omitempty"`

	// 	ATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty" json:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`

	// 	ATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty" json:"ATTRIBUTE-VALUE-INTEGER,omitempty"`

	// 	ATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty" json:"ATTRIBUTE-VALUE-REAL,omitempty"`

	// 	ATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty" json:"ATTRIBUTE-VALUE-STRING,omitempty"`

	// 	ATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty" json:"ATTRIBUTE-VALUE-XHTML,omitempty"`
	// } `xml:"VALUES,omitempty" json:"VALUES,omitempty"`

	// CHILDREN struct {
	// 	SPEC_HIERARCHY *SPEC_HIERARCHY `xml:"SPEC-HIERARCHY,omitempty" json:"SPEC-HIERARCHY,omitempty"`
	// } `xml:"CHILDREN,omitempty" json:"CHILDREN,omitempty"`

	// TYPE struct {
	// 	SPECIFICATION_TYPE_REF *LOCAL_REF `xml:"SPECIFICATION-TYPE-REF,omitempty" json:"SPECIFICATION-TYPE-REF,omitempty"`
	// } `xml:"TYPE,omitempty" json:"TYPE,omitempty"`

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

func (reqif *SPECIFICATION) Walk(_reqif *schema.SPECIFICATION, stage *StageStruct) {

	reqif.DESC = _reqif.DESC
	reqif.IDENTIFIER = string(*_reqif.IDENTIFIER)
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME
}
