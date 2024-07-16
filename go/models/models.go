package models

import (
	"encoding/xml"
)

// REQIF ...
type REQIF struct {
	Name        string
	XMLName     xml.Name     `xml:"REQ-IF"`
	HEADER      *THEHEADER   `xml:"THE-HEADER"`
	CORECONTENT *CORECONTENT `xml:"CORE-CONTENT"`
}

// LOCALREF ...
type LOCALREF string

// GLOBALREF ...
type GLOBALREF string

// ALTERNATIVEID ...
type ALTERNATIVEID struct {
	Name           string
	XMLName        xml.Name `xml:"ALTERNATIVE-ID"`
	IDENTIFIERAttr string   `xml:"IDENTIFIER,attr"`
}

// DEFAULTVALUE ...
type DEFAULTVALUE struct {
	Name                  string
	XMLName               xml.Name               `xml:"DEFAULT-VALUE"`
	ATTRIBUTEVALUEBOOLEAN *ATTRIBUTEVALUEBOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN"`
}

// REQTYPE ...
type REQTYPE struct {
	Name                         string
	DATATYPEDEFINITIONBOOLEANREF string `xml:"DATATYPE-DEFINITION-BOOLEAN-REF"`
}

// ATTRIBUTEDEFINITIONBOOLEAN ...
type ATTRIBUTEDEFINITIONBOOLEAN struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-BOOLEAN"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONDATE ...
type ATTRIBUTEDEFINITIONDATE struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-DATE"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONENUMERATION ...
type ATTRIBUTEDEFINITIONENUMERATION struct {
	Name            string
	XMLName         xml.Name       `xml:"ATTRIBUTE-DEFINITION-ENUMERATION"`
	DESCAttr        string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr  string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr  bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr  string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr    string         `xml:"LONG-NAME,attr,omitempty"`
	MULTIVALUEDAttr bool           `xml:"MULTI-VALUED,attr"`
	DEFAULTVALUE    *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	ALTERNATIVEID   *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	TYPE            *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONINTEGER ...
type ATTRIBUTEDEFINITIONINTEGER struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-INTEGER"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONREAL ...
type ATTRIBUTEDEFINITIONREAL struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-REAL"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONSTRING ...
type ATTRIBUTEDEFINITIONSTRING struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-STRING"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// ATTRIBUTEDEFINITIONXHTML ...
type ATTRIBUTEDEFINITIONXHTML struct {
	Name           string
	XMLName        xml.Name       `xml:"ATTRIBUTE-DEFINITION-XHTML"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr bool           `xml:"IS-EDITABLE,attr,omitempty"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	DEFAULTVALUE   *DEFAULTVALUE  `xml:"DEFAULT-VALUE"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// DEFINITION ...
type DEFINITION struct {
	Name                          string
	ATTRIBUTEDEFINITIONBOOLEANREF string `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF"`
}

// ATTRIBUTEVALUEBOOLEAN ...
type ATTRIBUTEVALUEBOOLEAN struct {
	Name         string
	XMLName      xml.Name    `xml:"ATTRIBUTE-VALUE-BOOLEAN"`
	THEVALUEAttr bool        `xml:"THE-VALUE,attr"`
	DEFINITION   *DEFINITION `xml:"DEFINITION"`
}

// ATTRIBUTEVALUEDATE ...
type ATTRIBUTEVALUEDATE struct {
	Name         string
	XMLName      xml.Name    `xml:"ATTRIBUTE-VALUE-DATE"`
	THEVALUEAttr string      `xml:"THE-VALUE,attr"`
	DEFINITION   *DEFINITION `xml:"DEFINITION"`
}

// VALUES ...
type VALUES struct {
	Name         string
	ENUMVALUEREF []string `xml:"ENUM-VALUE-REF"`
}

// ATTRIBUTEVALUEENUMERATION ...
type ATTRIBUTEVALUEENUMERATION struct {
	Name       string
	XMLName    xml.Name    `xml:"ATTRIBUTE-VALUE-ENUMERATION"`
	DEFINITION *DEFINITION `xml:"DEFINITION"`
	VALUES     *VALUES     `xml:"VALUES"`
}

// ATTRIBUTEVALUEINTEGER ...
type ATTRIBUTEVALUEINTEGER struct {
	Name         string
	XMLName      xml.Name    `xml:"ATTRIBUTE-VALUE-INTEGER"`
	THEVALUEAttr int         `xml:"THE-VALUE,attr"`
	DEFINITION   *DEFINITION `xml:"DEFINITION"`
}

// ATTRIBUTEVALUEREAL ...
type ATTRIBUTEVALUEREAL struct {
	Name         string
	XMLName      xml.Name    `xml:"ATTRIBUTE-VALUE-REAL"`
	THEVALUEAttr float64     `xml:"THE-VALUE,attr"`
	DEFINITION   *DEFINITION `xml:"DEFINITION"`
}

// ATTRIBUTEVALUESTRING ...
type ATTRIBUTEVALUESTRING struct {
	Name         string
	XMLName      xml.Name    `xml:"ATTRIBUTE-VALUE-STRING"`
	THEVALUEAttr string      `xml:"THE-VALUE,attr"`
	DEFINITION   *DEFINITION `xml:"DEFINITION"`
}

// ATTRIBUTEVALUEXHTML ...
type ATTRIBUTEVALUEXHTML struct {
	Name             string
	XMLName          xml.Name      `xml:"ATTRIBUTE-VALUE-XHTML"`
	ISSIMPLIFIEDAttr bool          `xml:"IS-SIMPLIFIED,attr,omitempty"`
	THEVALUE         *XHTMLCONTENT `xml:"THE-VALUE"`
	THEORIGINALVALUE *XHTMLCONTENT `xml:"THE-ORIGINAL-VALUE"`
	DEFINITION       *DEFINITION   `xml:"DEFINITION"`
}

// DATATYPEDEFINITIONBOOLEAN ...
type DATATYPEDEFINITIONBOOLEAN struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-BOOLEAN"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// DATATYPEDEFINITIONDATE ...
type DATATYPEDEFINITIONDATE struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-DATE"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// SPECIFIEDVALUES ...
type SPECIFIEDVALUES struct {
	Name      string
	XMLName   xml.Name     `xml:"SPECIFIED-VALUES"`
	ENUMVALUE []*ENUMVALUE `xml:"ENUM-VALUE"`
}

// DATATYPEDEFINITIONENUMERATION ...
type DATATYPEDEFINITIONENUMERATION struct {
	Name            string
	XMLName         xml.Name         `xml:"DATATYPE-DEFINITION-ENUMERATION"`
	DESCAttr        string           `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr  string           `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr  string           `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr    string           `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID   *ALTERNATIVEID   `xml:"ALTERNATIVE-ID"`
	SPECIFIEDVALUES *SPECIFIEDVALUES `xml:"SPECIFIED-VALUES"`
}

// DATATYPEDEFINITIONINTEGER ...
type DATATYPEDEFINITIONINTEGER struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-INTEGER"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	MAXAttr        int            `xml:"MAX,attr"`
	MINAttr        int            `xml:"MIN,attr"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// DATATYPEDEFINITIONREAL ...
type DATATYPEDEFINITIONREAL struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-REAL"`
	ACCURACYAttr   int            `xml:"ACCURACY,attr"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	MAXAttr        float64        `xml:"MAX,attr"`
	MINAttr        float64        `xml:"MIN,attr"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// DATATYPEDEFINITIONSTRING ...
type DATATYPEDEFINITIONSTRING struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-STRING"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	MAXLENGTHAttr  int            `xml:"MAX-LENGTH,attr"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// DATATYPEDEFINITIONXHTML ...
type DATATYPEDEFINITIONXHTML struct {
	Name           string
	XMLName        xml.Name       `xml:"DATATYPE-DEFINITION-XHTML"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
}

// EMBEDDEDVALUE ...
type EMBEDDEDVALUE struct {
	Name             string
	XMLName          xml.Name `xml:"EMBEDDED-VALUE"`
	KEYAttr          int      `xml:"KEY,attr"`
	OTHERCONTENTAttr string   `xml:"OTHER-CONTENT,attr"`
}

// PROPERTIES ...
type PROPERTIES struct {
	Name          string
	EMBEDDEDVALUE *EMBEDDEDVALUE `xml:"EMBEDDED-VALUE"`
}

// ENUMVALUE ...
type ENUMVALUE struct {
	Name           string
	XMLName        xml.Name       `xml:"ENUM-VALUE"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	PROPERTIES     *PROPERTIES    `xml:"PROPERTIES"`
}

// SOURCESPECIFICATION ...
type SOURCESPECIFICATION struct {
	Name             string
	XMLName          xml.Name `xml:"SOURCE-SPECIFICATION"`
	SPECIFICATIONREF string   `xml:"SPECIFICATION-REF"`
}

// SPECRELATIONS ...
type SPECRELATIONS struct {
	Name            string
	XMLName         xml.Name `xml:"SPEC-RELATIONS"`
	SPECRELATIONREF []string `xml:"SPEC-RELATION-REF"`
}

// TARGETSPECIFICATION ...
type TARGETSPECIFICATION struct {
	Name             string
	XMLName          xml.Name `xml:"TARGET-SPECIFICATION"`
	SPECIFICATIONREF string   `xml:"SPECIFICATION-REF"`
}

// RELATIONGROUP ...
type RELATIONGROUP struct {
	Name                string
	XMLName             xml.Name             `xml:"RELATION-GROUP"`
	DESCAttr            string               `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr      string               `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr      string               `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr        string               `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID       *ALTERNATIVEID       `xml:"ALTERNATIVE-ID"`
	SOURCESPECIFICATION *SOURCESPECIFICATION `xml:"SOURCE-SPECIFICATION"`
	SPECRELATIONS       *SPECRELATIONS       `xml:"SPEC-RELATIONS"`
	TARGETSPECIFICATION *TARGETSPECIFICATION `xml:"TARGET-SPECIFICATION"`
	TYPE                *REQTYPE             `xml:"TYPE"`
}

// SPECATTRIBUTES ...
type SPECATTRIBUTES struct {
	Name                           string
	XMLName                        xml.Name                          `xml:"SPEC-ATTRIBUTES"`
	ATTRIBUTEDEFINITIONBOOLEAN     []*ATTRIBUTEDEFINITIONBOOLEAN     `xml:"ATTRIBUTE-DEFINITION-BOOLEAN"`
	ATTRIBUTEDEFINITIONDATE        []*ATTRIBUTEDEFINITIONDATE        `xml:"ATTRIBUTE-DEFINITION-DATE"`
	ATTRIBUTEDEFINITIONENUMERATION []*ATTRIBUTEDEFINITIONENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION"`
	ATTRIBUTEDEFINITIONINTEGER     []*ATTRIBUTEDEFINITIONINTEGER     `xml:"ATTRIBUTE-DEFINITION-INTEGER"`
	ATTRIBUTEDEFINITIONREAL        []*ATTRIBUTEDEFINITIONREAL        `xml:"ATTRIBUTE-DEFINITION-REAL"`
	ATTRIBUTEDEFINITIONSTRING      []*ATTRIBUTEDEFINITIONSTRING      `xml:"ATTRIBUTE-DEFINITION-STRING"`
	ATTRIBUTEDEFINITIONXHTML       []*ATTRIBUTEDEFINITIONXHTML       `xml:"ATTRIBUTE-DEFINITION-XHTML"`
}

// RELATIONGROUPTYPE ...
type RELATIONGROUPTYPE struct {
	Name           string
	XMLName        xml.Name        `xml:"RELATION-GROUP-TYPE"`
	DESCAttr       string          `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string          `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string          `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string          `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID  `xml:"ALTERNATIVE-ID"`
	SPECATTRIBUTES *SPECATTRIBUTES `xml:"SPEC-ATTRIBUTES"`
}

// THEHEADER ...
type THEHEADER struct {
	Name        string
	XMLName     xml.Name     `xml:"THE-HEADER"`
	REQIFHEADER *REQIFHEADER `xml:"REQ-IF-HEADER"`
}

// CORECONTENT ...
type CORECONTENT struct {
	Name         string
	XMLName      xml.Name      `xml:"CORE-CONTENT"`
	REQIFCONTENT *REQIFCONTENT `xml:"REQ-IF-CONTENT"`
}

// TOOLEXTENSIONS ...
type TOOLEXTENSIONS struct {
	Name               string
	XMLName            xml.Name              `xml:"TOOL-EXTENSIONS"`
	REQIFTOOLEXTENSION []*REQIFTOOLEXTENSION `xml:"REQ-IF-TOOL-EXTENSION"`
}

// DATATYPES ...
type DATATYPES struct {
	Name                          string
	DATATYPEDEFINITIONBOOLEAN     []*DATATYPEDEFINITIONBOOLEAN     `xml:"DATATYPE-DEFINITION-BOOLEAN"`
	DATATYPEDEFINITIONDATE        []*DATATYPEDEFINITIONDATE        `xml:"DATATYPE-DEFINITION-DATE"`
	DATATYPEDEFINITIONENUMERATION []*DATATYPEDEFINITIONENUMERATION `xml:"DATATYPE-DEFINITION-ENUMERATION"`
	DATATYPEDEFINITIONINTEGER     []*DATATYPEDEFINITIONINTEGER     `xml:"DATATYPE-DEFINITION-INTEGER"`
	DATATYPEDEFINITIONREAL        []*DATATYPEDEFINITIONREAL        `xml:"DATATYPE-DEFINITION-REAL"`
	DATATYPEDEFINITIONSTRING      []*DATATYPEDEFINITIONSTRING      `xml:"DATATYPE-DEFINITION-STRING"`
	DATATYPEDEFINITIONXHTML       []*DATATYPEDEFINITIONXHTML       `xml:"DATATYPE-DEFINITION-XHTML"`
}

// SPECTYPES ...
type SPECTYPES struct {
	Name              string
	XMLName           xml.Name             `xml:"SPEC-TYPES"`
	RELATIONGROUPTYPE []*RELATIONGROUPTYPE `xml:"RELATION-GROUP-TYPE"`
	SPECOBJECTTYPE    []*SPECOBJECTTYPE    `xml:"SPEC-OBJECT-TYPE"`
	SPECRELATIONTYPE  []*SPECRELATIONTYPE  `xml:"SPEC-RELATION-TYPE"`
	SPECIFICATIONTYPE []*SPECIFICATIONTYPE `xml:"SPECIFICATION-TYPE"`
}

// SPECOBJECTS ...
type SPECOBJECTS struct {
	Name       string
	XMLName    xml.Name      `xml:"SPEC-OBJECTS"`
	SPECOBJECT []*SPECOBJECT `xml:"SPEC-OBJECT"`
}

// SPECIFICATIONS ...
type SPECIFICATIONS struct {
	Name          string
	SPECIFICATION []*SPECIFICATION `xml:"SPECIFICATION"`
}

// SPECRELATIONGROUPS ...
type SPECRELATIONGROUPS struct {
	Name          string
	XMLName       xml.Name         `xml:"SPEC-RELATION-GROUPS"`
	RELATIONGROUP []*RELATIONGROUP `xml:"RELATION-GROUP"`
}

// REQIFCONTENT ...
type REQIFCONTENT struct {
	Name               string
	XMLName            xml.Name            `xml:"REQ-IF-CONTENT"`
	DATATYPES          *DATATYPES          `xml:"DATATYPES"`
	SPECTYPES          *SPECTYPES          `xml:"SPEC-TYPES"`
	SPECOBJECTS        *SPECOBJECTS        `xml:"SPEC-OBJECTS"`
	SPECRELATIONS      *SPECRELATIONS      `xml:"SPEC-RELATIONS"`
	SPECIFICATIONS     *SPECIFICATIONS     `xml:"SPECIFICATIONS"`
	SPECRELATIONGROUPS *SPECRELATIONGROUPS `xml:"SPEC-RELATION-GROUPS"`
}

// REQIFHEADER ...
type REQIFHEADER struct {
	Name           string
	XMLName        xml.Name `xml:"REQ-IF-HEADER"`
	IDENTIFIERAttr string   `xml:"IDENTIFIER,attr"`
	COMMENT        string   `xml:"COMMENT"`
	CREATIONTIME   string   `xml:"CREATION-TIME"`
	REPOSITORYID   string   `xml:"REPOSITORY-ID"`
	REQIFTOOLID    string   `xml:"REQ-IF-TOOL-ID"`
	REQIFVERSION   string   `xml:"REQ-IF-VERSION"`
	SOURCETOOLID   string   `xml:"SOURCE-TOOL-ID"`
	TITLE          string   `xml:"TITLE"`
}

// CHILDREN ...
type CHILDREN struct {
	Name          string
	SPECHIERARCHY []*SPECHIERARCHY `xml:"SPEC-HIERARCHY"`
}

// EDITABLEATTS ...
type EDITABLEATTS struct {
	Name                              string
	XMLName                           xml.Name `xml:"EDITABLE-ATTS"`
	ATTRIBUTEDEFINITIONBOOLEANREF     []string `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF"`
	ATTRIBUTEDEFINITIONDATEREF        []string `xml:"ATTRIBUTE-DEFINITION-DATE-REF"`
	ATTRIBUTEDEFINITIONENUMERATIONREF []string `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF"`
	ATTRIBUTEDEFINITIONINTEGERREF     []string `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF"`
	ATTRIBUTEDEFINITIONREALREF        []string `xml:"ATTRIBUTE-DEFINITION-REAL-REF"`
	ATTRIBUTEDEFINITIONSTRINGREF      []string `xml:"ATTRIBUTE-DEFINITION-STRING-REF"`
	ATTRIBUTEDEFINITIONXHTMLREF       []string `xml:"ATTRIBUTE-DEFINITION-XHTML-REF"`
}

// OBJECT ...
type OBJECT struct {
	Name          string
	SPECOBJECTREF string `xml:"SPEC-OBJECT-REF"`
}

// SPECHIERARCHY ...
type SPECHIERARCHY struct {
	Name                string
	XMLName             xml.Name       `xml:"SPEC-HIERARCHY"`
	DESCAttr            string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr      string         `xml:"IDENTIFIER,attr"`
	ISEDITABLEAttr      bool           `xml:"IS-EDITABLE,attr,omitempty"`
	ISTABLEINTERNALAttr bool           `xml:"IS-TABLE-INTERNAL,attr,omitempty"`
	LASTCHANGEAttr      string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr        string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID       *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	CHILDREN            *CHILDREN      `xml:"CHILDREN"`
	EDITABLEATTS        *EDITABLEATTS  `xml:"EDITABLE-ATTS"`
	OBJECT              *OBJECT        `xml:"OBJECT"`
}

// SPECOBJECT ...
type SPECOBJECT struct {
	Name           string
	XMLName        xml.Name       `xml:"SPEC-OBJECT"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	VALUES         *VALUES        `xml:"VALUES"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// SPECOBJECTTYPE ...
type SPECOBJECTTYPE struct {
	Name           string
	XMLName        xml.Name        `xml:"SPEC-OBJECT-TYPE"`
	DESCAttr       string          `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string          `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string          `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string          `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID  `xml:"ALTERNATIVE-ID"`
	SPECATTRIBUTES *SPECATTRIBUTES `xml:"SPEC-ATTRIBUTES"`
}

// SOURCE ...
type SOURCE struct {
	Name          string
	SPECOBJECTREF string `xml:"SPEC-OBJECT-REF"`
}

// TARGET ...
type TARGET struct {
	Name          string
	SPECOBJECTREF string `xml:"SPEC-OBJECT-REF"`
}

// SPECRELATION ...
type SPECRELATION struct {
	Name           string
	XMLName        xml.Name       `xml:"SPEC-RELATION"`
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	VALUES         *VALUES        `xml:"VALUES"`
	SOURCE         *SOURCE        `xml:"SOURCE"`
	TARGET         *TARGET        `xml:"TARGET"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// SPECRELATIONTYPE ...
type SPECRELATIONTYPE struct {
	Name           string
	XMLName        xml.Name        `xml:"SPEC-RELATION-TYPE"`
	DESCAttr       string          `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string          `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string          `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string          `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID  `xml:"ALTERNATIVE-ID"`
	SPECATTRIBUTES *SPECATTRIBUTES `xml:"SPEC-ATTRIBUTES"`
}

// SPECIFICATION ...
type SPECIFICATION struct {
	Name           string
	DESCAttr       string         `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string         `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string         `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string         `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID `xml:"ALTERNATIVE-ID"`
	VALUES         *VALUES        `xml:"VALUES"`
	CHILDREN       *CHILDREN      `xml:"CHILDREN"`
	TYPE           *REQTYPE       `xml:"TYPE"`
}

// SPECIFICATIONTYPE ...
type SPECIFICATIONTYPE struct {
	Name           string
	XMLName        xml.Name        `xml:"SPECIFICATION-TYPE"`
	DESCAttr       string          `xml:"DESC,attr,omitempty"`
	IDENTIFIERAttr string          `xml:"IDENTIFIER,attr"`
	LASTCHANGEAttr string          `xml:"LAST-CHANGE,attr"`
	LONGNAMEAttr   string          `xml:"LONG-NAME,attr,omitempty"`
	ALTERNATIVEID  *ALTERNATIVEID  `xml:"ALTERNATIVE-ID"`
	SPECATTRIBUTES *SPECATTRIBUTES `xml:"SPEC-ATTRIBUTES"`
}

// REQIFTOOLEXTENSION ...
type REQIFTOOLEXTENSION struct {
	Name    string
	XMLName xml.Name `xml:"REQ-IF-TOOL-EXTENSION"`
}

// XHTMLCONTENT ...
type XHTMLCONTENT struct {
	Name    string
	XMLName xml.Name `xml:"XHTML-CONTENT"`
}
