package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type REQIF struct {
	Name string

	REQ_IF_HEADER  *REQ_IF_HEADER
	REQ_IF_CONTENT *REQ_IF_CONTENT
}

func (reqif *REQIF) Walk(_reqif *schema.REQ_IF, stage *StageStruct) {
	log.Println("reqif:Walk")

	reqif.REQ_IF_HEADER = new(REQ_IF_HEADER).Stage(stage)
	reqif.REQ_IF_HEADER.Name = time.Now().Format(time.DateTime)
	reqif.REQ_IF_HEADER.COMMENT = _reqif.THE_HEADER.REQ_IF_HEADER.COMMENT
	// reqif.REQ_IF_HEADER.CREATION_TIME = time.Time(_reqif.THE_HEADER.REQ_IF_HEADER.CREATION_TIME)
	reqif.REQ_IF_HEADER.CREATION_TIME = _reqif.THE_HEADER.REQ_IF_HEADER.CREATION_TIME.ToGoTime()

	const myFormat = "2006-01-02T15:04:05.000-07:00"
	log.Println("Date time", reqif.REQ_IF_HEADER.CREATION_TIME.Local().Format(myFormat))

	reqif.REQ_IF_HEADER.REPOSITORY_ID = _reqif.THE_HEADER.REQ_IF_HEADER.REPOSITORY_ID
	reqif.REQ_IF_HEADER.REQ_IF_TOOL_ID = _reqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_TOOL_ID
	reqif.REQ_IF_HEADER.REQ_IF_VERSION = _reqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_VERSION
	reqif.REQ_IF_HEADER.SOURCE_TOOL_ID = _reqif.THE_HEADER.REQ_IF_HEADER.SOURCE_TOOL_ID
	reqif.REQ_IF_HEADER.TITLE = _reqif.THE_HEADER.REQ_IF_HEADER.TITLE

	reqif.REQ_IF_CONTENT = new(REQ_IF_CONTENT).Stage(stage)
	reqif.REQ_IF_CONTENT.Walk(_reqif.CORE_CONTENT.REQ_IF_CONTENT, stage)
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

	// SPEC_RELATION_GROUPS struct {
	// 	RELATION_GROUP *RELATION_GROUP `xml:"RELATION-GROUP,omitempty" json:"RELATION-GROUP,omitempty"`
	// } `xml:"SPEC-RELATION-GROUPS,omitempty" json:"SPEC-RELATION-GROUPS,omitempty"`
}

func (reqifContent *REQ_IF_CONTENT) Walk(_reqifContent *schema.REQ_IF_CONTENT, stage *StageStruct) {
	log.Println("REQ_IF_CONTENT", "Walk")

	reqifContent.Name = time.Now().Format(time.DateTime)
}

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

	// DESC string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd DESC,attr,omitempty" json:"DESC,omitempty"`

	// IDENTIFIER *ID `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd IDENTIFIER,attr,omitempty" json:"IDENTIFIER,omitempty"`

	// LAST_CHANGE soap.XSDDateTime `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LAST-CHANGE,attr,omitempty" json:"LAST-CHANGE,omitempty"`

	// LONG_NAME string `xml:"http://www.omg.org/spec/ReqIF/20110401/reqif.xsd LONG-NAME,attr,omitempty" json:"LONG-NAME,omitempty"`
}
