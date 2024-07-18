package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPEC_HIERARCHY struct {
	Name string
	// ALTERNATIVE_ID struct {
	// 	ALTERNATIVE_ID *ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`
	// } `xml:"ALTERNATIVE-ID,omitempty" json:"ALTERNATIVE-ID,omitempty"`

	// CHILDREN struct {
	// 	SPEC_HIERARCHY *SPEC_HIERARCHY `xml:"SPEC-HIERARCHY,omitempty" json:"SPEC-HIERARCHY,omitempty"`
	// } `xml:"CHILDREN,omitempty" json:"CHILDREN,omitempty"`
	CHILDREN *SPEC_HIERARCHY

	// EDITABLE_ATTS struct {
	// 	ATTRIBUTE_DEFINITION_BOOLEAN_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty" json:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_DATE_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty" json:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_ENUMERATION_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty" json:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_INTEGER_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty" json:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_REAL_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty" json:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_STRING_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty" json:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty"`

	// 	ATTRIBUTE_DEFINITION_XHTML_REF *LOCAL_REF `xml:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty" json:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty"`
	// } `xml:"EDITABLE-ATTS,omitempty" json:"EDITABLE-ATTS,omitempty"`

	// OBJECT struct {
	// 	SPEC_OBJECT_REF *LOCAL_REF `xml:"SPEC-OBJECT-REF,omitempty" json:"SPEC-OBJECT-REF,omitempty"`
	// } `xml:"OBJECT,omitempty" json:"OBJECT,omitempty"`
	OBJECT string

	DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	IDENTIFIER string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	IS_EDITABLE bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-EDITABLE,attr,omitempty" json:"IS-EDITABLE,omitempty"`

	IS_TABLE_INTERNAL bool `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IS-TABLE-INTERNAL,attr,omitempty" json:"IS-TABLE-INTERNAL,omitempty"`

	LAST_CHANGE time.Time `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}

func (reqif *SPEC_HIERARCHY) Walk(_reqif *schema.SPEC_HIERARCHY, stage *StageStruct) {
	log.Println("SPEC_HIERARCHY", "Walk")

	reqif.DESC = _reqif.DESC
	if _reqif.IDENTIFIER != nil {
		reqif.IDENTIFIER = string(*_reqif.IDENTIFIER)
	}
	reqif.IS_EDITABLE = _reqif.IS_EDITABLE
	reqif.IS_TABLE_INTERNAL = _reqif.IS_TABLE_INTERNAL
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME
	// if _reqif.OBJECT.SPEC_OBJECT_REF != nil {
	// 	reqif.OBJECT = string(*_reqif.OBJECT.SPEC_OBJECT_REF)
	// 	reqif.Name = reqif.OBJECT
	// }

	// if _reqif.CHILDREN.SPEC_HIERARCHY != nil {
	// 	sh := new(SPEC_HIERARCHY).Stage(stage)
	// 	sh.Walk(_reqif.CHILDREN.SPEC_HIERARCHY, stage)
	// 	reqif.CHILDREN = sh
	// }
}
