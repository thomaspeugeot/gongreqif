package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type REQ_IF_CONTENT struct {
	Name string
	// DATATYPES struct {
	// 	DATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN `xml:"DATATYPE-DEFINITION-BOOLEAN,omitempty" json:"DATATYPE-DEFINITION-BOOLEAN,omitempty"`

	// 	DATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE `xml:"DATATYPE-DEFINITION-DATE,omitempty" json:"DATATYPE-DEFINITION-DATE,omitempty"`

	// 	DATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION `xml:"DATATYPE-DEFINITION-ENUMERATION,omitempty" json:"DATATYPE-DEFINITION-ENUMERATION,omitempty"`

	// 	DATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER `xml:"DATATYPE-DEFINITION-INTEGER,omitempty" json:"DATATYPE-DEFINITION-INTEGER,omitempty"`

	// 	DATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL `xml:"DATATYPE-DEFINITION-REAL,omitempty" json:"DATATYPE-DEFINITION-REAL,omitempty"`

	// 	DATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING `xml:"DATATYPE-DEFINITION-STRING,omitempty" json:"DATATYPE-DEFINITION-STRING,omitempty"`

	// 	DATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML `xml:"DATATYPE-DEFINITION-XHTML,omitempty" json:"DATATYPE-DEFINITION-XHTML,omitempty"`
	// } `xml:"DATATYPES,omitempty" json:"DATATYPES,omitempty"`

	// SPEC_TYPES struct {
	// 	RELATION_GROUP_TYPE *RELATION_GROUP_TYPE `xml:"RELATION-GROUP-TYPE,omitempty" json:"RELATION-GROUP-TYPE,omitempty"`

	// 	SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE `xml:"SPEC-OBJECT-TYPE,omitempty" json:"SPEC-OBJECT-TYPE,omitempty"`

	// 	SPEC_RELATION_TYPE *SPEC_RELATION_TYPE `xml:"SPEC-RELATION-TYPE,omitempty" json:"SPEC-RELATION-TYPE,omitempty"`

	// 	SPECIFICATION_TYPE *SPECIFICATION_TYPE `xml:"SPECIFICATION-TYPE,omitempty" json:"SPECIFICATION-TYPE,omitempty"`
	// } `xml:"SPEC-TYPES,omitempty" json:"SPEC-TYPES,omitempty"`

	// SPEC_OBJECTS struct {
	// 	SPEC_OBJECT *SPEC_OBJECT `xml:"SPEC-OBJECT,omitempty" json:"SPEC-OBJECT,omitempty"`
	// } `xml:"SPEC-OBJECTS,omitempty" json:"SPEC-OBJECTS,omitempty"`

	// SPEC_RELATIONS struct {
	// 	SPEC_RELATION *SPEC_RELATION `xml:"SPEC-RELATION,omitempty" json:"SPEC-RELATION,omitempty"`
	// } `xml:"SPEC-RELATIONS,omitempty" json:"SPEC-RELATIONS,omitempty"`

	// SPECIFICATIONS struct {
	// 	SPECIFICATION *SPECIFICATION `xml:"SPECIFICATION,omitempty" json:"SPECIFICATION,omitempty"`
	// } `xml:"SPECIFICATIONS,omitempty" json:"SPECIFICATIONS,omitempty"`
	SPECIFICATION *SPECIFICATION

	// SPEC_RELATION_GROUPS struct {
	// 	RELATION_GROUP *RELATION_GROUP `xml:"RELATION-GROUP,omitempty" json:"RELATION-GROUP,omitempty"`
	// } `xml:"SPEC-RELATION-GROUPS,omitempty" json:"SPEC-RELATION-GROUPS,omitempty"`
}

func (reqifContent *REQ_IF_CONTENT) Walk(_reqifContent *schema.REQ_IF_CONTENT, stage *StageStruct) {
	log.Println("REQ_IF_CONTENT", "Walk")

	reqifContent.Name = time.Now().Format(time.DateTime)

	specification := new(SPECIFICATION).Stage(stage)
	reqifContent.SPECIFICATION = specification
	specification.Walk(_reqifContent.SPECIFICATIONS.SPECIFICATION, stage)

}
