// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	ALTERNATIVEIDs           map[*ALTERNATIVEID]any
	ALTERNATIVEIDs_mapString map[string]*ALTERNATIVEID

	// insertion point for slice of pointers maps

	OnAfterALTERNATIVEIDCreateCallback OnAfterCreateInterface[ALTERNATIVEID]
	OnAfterALTERNATIVEIDUpdateCallback OnAfterUpdateInterface[ALTERNATIVEID]
	OnAfterALTERNATIVEIDDeleteCallback OnAfterDeleteInterface[ALTERNATIVEID]
	OnAfterALTERNATIVEIDReadCallback   OnAfterReadInterface[ALTERNATIVEID]

	ATTRIBUTEDEFINITIONBOOLEANs           map[*ATTRIBUTEDEFINITIONBOOLEAN]any
	ATTRIBUTEDEFINITIONBOOLEANs_mapString map[string]*ATTRIBUTEDEFINITIONBOOLEAN

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONBOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONBOOLEAN]
	OnAfterATTRIBUTEDEFINITIONBOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONBOOLEAN]
	OnAfterATTRIBUTEDEFINITIONBOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONBOOLEAN]
	OnAfterATTRIBUTEDEFINITIONBOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONBOOLEAN]

	ATTRIBUTEDEFINITIONDATEs           map[*ATTRIBUTEDEFINITIONDATE]any
	ATTRIBUTEDEFINITIONDATEs_mapString map[string]*ATTRIBUTEDEFINITIONDATE

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONDATECreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONDATE]
	OnAfterATTRIBUTEDEFINITIONDATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONDATE]
	OnAfterATTRIBUTEDEFINITIONDATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONDATE]
	OnAfterATTRIBUTEDEFINITIONDATEReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONDATE]

	ATTRIBUTEDEFINITIONENUMERATIONs           map[*ATTRIBUTEDEFINITIONENUMERATION]any
	ATTRIBUTEDEFINITIONENUMERATIONs_mapString map[string]*ATTRIBUTEDEFINITIONENUMERATION

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONENUMERATION]
	OnAfterATTRIBUTEDEFINITIONENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONENUMERATION]
	OnAfterATTRIBUTEDEFINITIONENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONENUMERATION]
	OnAfterATTRIBUTEDEFINITIONENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONENUMERATION]

	ATTRIBUTEDEFINITIONINTEGERs           map[*ATTRIBUTEDEFINITIONINTEGER]any
	ATTRIBUTEDEFINITIONINTEGERs_mapString map[string]*ATTRIBUTEDEFINITIONINTEGER

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONINTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONINTEGER]
	OnAfterATTRIBUTEDEFINITIONINTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONINTEGER]
	OnAfterATTRIBUTEDEFINITIONINTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONINTEGER]
	OnAfterATTRIBUTEDEFINITIONINTEGERReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONINTEGER]

	ATTRIBUTEDEFINITIONREALs           map[*ATTRIBUTEDEFINITIONREAL]any
	ATTRIBUTEDEFINITIONREALs_mapString map[string]*ATTRIBUTEDEFINITIONREAL

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONREALCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONREAL]
	OnAfterATTRIBUTEDEFINITIONREALUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONREAL]
	OnAfterATTRIBUTEDEFINITIONREALDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONREAL]
	OnAfterATTRIBUTEDEFINITIONREALReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONREAL]

	ATTRIBUTEDEFINITIONSTRINGs           map[*ATTRIBUTEDEFINITIONSTRING]any
	ATTRIBUTEDEFINITIONSTRINGs_mapString map[string]*ATTRIBUTEDEFINITIONSTRING

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONSTRINGCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONSTRING]
	OnAfterATTRIBUTEDEFINITIONSTRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONSTRING]
	OnAfterATTRIBUTEDEFINITIONSTRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONSTRING]
	OnAfterATTRIBUTEDEFINITIONSTRINGReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONSTRING]

	ATTRIBUTEDEFINITIONXHTMLs           map[*ATTRIBUTEDEFINITIONXHTML]any
	ATTRIBUTEDEFINITIONXHTMLs_mapString map[string]*ATTRIBUTEDEFINITIONXHTML

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEDEFINITIONXHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTEDEFINITIONXHTML]
	OnAfterATTRIBUTEDEFINITIONXHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTEDEFINITIONXHTML]
	OnAfterATTRIBUTEDEFINITIONXHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTEDEFINITIONXHTML]
	OnAfterATTRIBUTEDEFINITIONXHTMLReadCallback   OnAfterReadInterface[ATTRIBUTEDEFINITIONXHTML]

	ATTRIBUTEVALUEBOOLEANs           map[*ATTRIBUTEVALUEBOOLEAN]any
	ATTRIBUTEVALUEBOOLEANs_mapString map[string]*ATTRIBUTEVALUEBOOLEAN

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEBOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEBOOLEAN]
	OnAfterATTRIBUTEVALUEBOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEBOOLEAN]
	OnAfterATTRIBUTEVALUEBOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEBOOLEAN]
	OnAfterATTRIBUTEVALUEBOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEBOOLEAN]

	ATTRIBUTEVALUEDATEs           map[*ATTRIBUTEVALUEDATE]any
	ATTRIBUTEVALUEDATEs_mapString map[string]*ATTRIBUTEVALUEDATE

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEDATECreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEDATE]
	OnAfterATTRIBUTEVALUEDATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEDATE]
	OnAfterATTRIBUTEVALUEDATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEDATE]
	OnAfterATTRIBUTEVALUEDATEReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEDATE]

	ATTRIBUTEVALUEENUMERATIONs           map[*ATTRIBUTEVALUEENUMERATION]any
	ATTRIBUTEVALUEENUMERATIONs_mapString map[string]*ATTRIBUTEVALUEENUMERATION

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEENUMERATION]
	OnAfterATTRIBUTEVALUEENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEENUMERATION]
	OnAfterATTRIBUTEVALUEENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEENUMERATION]
	OnAfterATTRIBUTEVALUEENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEENUMERATION]

	ATTRIBUTEVALUEINTEGERs           map[*ATTRIBUTEVALUEINTEGER]any
	ATTRIBUTEVALUEINTEGERs_mapString map[string]*ATTRIBUTEVALUEINTEGER

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEINTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEINTEGER]
	OnAfterATTRIBUTEVALUEINTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEINTEGER]
	OnAfterATTRIBUTEVALUEINTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEINTEGER]
	OnAfterATTRIBUTEVALUEINTEGERReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEINTEGER]

	ATTRIBUTEVALUEREALs           map[*ATTRIBUTEVALUEREAL]any
	ATTRIBUTEVALUEREALs_mapString map[string]*ATTRIBUTEVALUEREAL

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEREALCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEREAL]
	OnAfterATTRIBUTEVALUEREALUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEREAL]
	OnAfterATTRIBUTEVALUEREALDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEREAL]
	OnAfterATTRIBUTEVALUEREALReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEREAL]

	ATTRIBUTEVALUESTRINGs           map[*ATTRIBUTEVALUESTRING]any
	ATTRIBUTEVALUESTRINGs_mapString map[string]*ATTRIBUTEVALUESTRING

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUESTRINGCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUESTRING]
	OnAfterATTRIBUTEVALUESTRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUESTRING]
	OnAfterATTRIBUTEVALUESTRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUESTRING]
	OnAfterATTRIBUTEVALUESTRINGReadCallback   OnAfterReadInterface[ATTRIBUTEVALUESTRING]

	ATTRIBUTEVALUEXHTMLs           map[*ATTRIBUTEVALUEXHTML]any
	ATTRIBUTEVALUEXHTMLs_mapString map[string]*ATTRIBUTEVALUEXHTML

	// insertion point for slice of pointers maps

	OnAfterATTRIBUTEVALUEXHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTEVALUEXHTML]
	OnAfterATTRIBUTEVALUEXHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTEVALUEXHTML]
	OnAfterATTRIBUTEVALUEXHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTEVALUEXHTML]
	OnAfterATTRIBUTEVALUEXHTMLReadCallback   OnAfterReadInterface[ATTRIBUTEVALUEXHTML]

	CHILDRENs           map[*CHILDREN]any
	CHILDRENs_mapString map[string]*CHILDREN

	// insertion point for slice of pointers maps
	CHILDREN_SPECHIERARCHY_reverseMap map[*SPECHIERARCHY]*CHILDREN

	OnAfterCHILDRENCreateCallback OnAfterCreateInterface[CHILDREN]
	OnAfterCHILDRENUpdateCallback OnAfterUpdateInterface[CHILDREN]
	OnAfterCHILDRENDeleteCallback OnAfterDeleteInterface[CHILDREN]
	OnAfterCHILDRENReadCallback   OnAfterReadInterface[CHILDREN]

	CORECONTENTs           map[*CORECONTENT]any
	CORECONTENTs_mapString map[string]*CORECONTENT

	// insertion point for slice of pointers maps

	OnAfterCORECONTENTCreateCallback OnAfterCreateInterface[CORECONTENT]
	OnAfterCORECONTENTUpdateCallback OnAfterUpdateInterface[CORECONTENT]
	OnAfterCORECONTENTDeleteCallback OnAfterDeleteInterface[CORECONTENT]
	OnAfterCORECONTENTReadCallback   OnAfterReadInterface[CORECONTENT]

	DATATYPEDEFINITIONBOOLEANs           map[*DATATYPEDEFINITIONBOOLEAN]any
	DATATYPEDEFINITIONBOOLEANs_mapString map[string]*DATATYPEDEFINITIONBOOLEAN

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONBOOLEANCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONBOOLEAN]
	OnAfterDATATYPEDEFINITIONBOOLEANUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONBOOLEAN]
	OnAfterDATATYPEDEFINITIONBOOLEANDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONBOOLEAN]
	OnAfterDATATYPEDEFINITIONBOOLEANReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONBOOLEAN]

	DATATYPEDEFINITIONDATEs           map[*DATATYPEDEFINITIONDATE]any
	DATATYPEDEFINITIONDATEs_mapString map[string]*DATATYPEDEFINITIONDATE

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONDATECreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONDATE]
	OnAfterDATATYPEDEFINITIONDATEUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONDATE]
	OnAfterDATATYPEDEFINITIONDATEDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONDATE]
	OnAfterDATATYPEDEFINITIONDATEReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONDATE]

	DATATYPEDEFINITIONENUMERATIONs           map[*DATATYPEDEFINITIONENUMERATION]any
	DATATYPEDEFINITIONENUMERATIONs_mapString map[string]*DATATYPEDEFINITIONENUMERATION

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONENUMERATIONCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONENUMERATION]
	OnAfterDATATYPEDEFINITIONENUMERATIONUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONENUMERATION]
	OnAfterDATATYPEDEFINITIONENUMERATIONDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONENUMERATION]
	OnAfterDATATYPEDEFINITIONENUMERATIONReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONENUMERATION]

	DATATYPEDEFINITIONINTEGERs           map[*DATATYPEDEFINITIONINTEGER]any
	DATATYPEDEFINITIONINTEGERs_mapString map[string]*DATATYPEDEFINITIONINTEGER

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONINTEGERCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONINTEGER]
	OnAfterDATATYPEDEFINITIONINTEGERUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONINTEGER]
	OnAfterDATATYPEDEFINITIONINTEGERDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONINTEGER]
	OnAfterDATATYPEDEFINITIONINTEGERReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONINTEGER]

	DATATYPEDEFINITIONREALs           map[*DATATYPEDEFINITIONREAL]any
	DATATYPEDEFINITIONREALs_mapString map[string]*DATATYPEDEFINITIONREAL

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONREALCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONREAL]
	OnAfterDATATYPEDEFINITIONREALUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONREAL]
	OnAfterDATATYPEDEFINITIONREALDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONREAL]
	OnAfterDATATYPEDEFINITIONREALReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONREAL]

	DATATYPEDEFINITIONSTRINGs           map[*DATATYPEDEFINITIONSTRING]any
	DATATYPEDEFINITIONSTRINGs_mapString map[string]*DATATYPEDEFINITIONSTRING

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONSTRINGCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONSTRING]
	OnAfterDATATYPEDEFINITIONSTRINGUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONSTRING]
	OnAfterDATATYPEDEFINITIONSTRINGDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONSTRING]
	OnAfterDATATYPEDEFINITIONSTRINGReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONSTRING]

	DATATYPEDEFINITIONXHTMLs           map[*DATATYPEDEFINITIONXHTML]any
	DATATYPEDEFINITIONXHTMLs_mapString map[string]*DATATYPEDEFINITIONXHTML

	// insertion point for slice of pointers maps

	OnAfterDATATYPEDEFINITIONXHTMLCreateCallback OnAfterCreateInterface[DATATYPEDEFINITIONXHTML]
	OnAfterDATATYPEDEFINITIONXHTMLUpdateCallback OnAfterUpdateInterface[DATATYPEDEFINITIONXHTML]
	OnAfterDATATYPEDEFINITIONXHTMLDeleteCallback OnAfterDeleteInterface[DATATYPEDEFINITIONXHTML]
	OnAfterDATATYPEDEFINITIONXHTMLReadCallback   OnAfterReadInterface[DATATYPEDEFINITIONXHTML]

	DATATYPESs           map[*DATATYPES]any
	DATATYPESs_mapString map[string]*DATATYPES

	// insertion point for slice of pointers maps
	DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap map[*DATATYPEDEFINITIONBOOLEAN]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONDATE_reverseMap map[*DATATYPEDEFINITIONDATE]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap map[*DATATYPEDEFINITIONENUMERATION]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap map[*DATATYPEDEFINITIONINTEGER]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONREAL_reverseMap map[*DATATYPEDEFINITIONREAL]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap map[*DATATYPEDEFINITIONSTRING]*DATATYPES
	DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap map[*DATATYPEDEFINITIONXHTML]*DATATYPES

	OnAfterDATATYPESCreateCallback OnAfterCreateInterface[DATATYPES]
	OnAfterDATATYPESUpdateCallback OnAfterUpdateInterface[DATATYPES]
	OnAfterDATATYPESDeleteCallback OnAfterDeleteInterface[DATATYPES]
	OnAfterDATATYPESReadCallback   OnAfterReadInterface[DATATYPES]

	DEFAULTVALUEs           map[*DEFAULTVALUE]any
	DEFAULTVALUEs_mapString map[string]*DEFAULTVALUE

	// insertion point for slice of pointers maps

	OnAfterDEFAULTVALUECreateCallback OnAfterCreateInterface[DEFAULTVALUE]
	OnAfterDEFAULTVALUEUpdateCallback OnAfterUpdateInterface[DEFAULTVALUE]
	OnAfterDEFAULTVALUEDeleteCallback OnAfterDeleteInterface[DEFAULTVALUE]
	OnAfterDEFAULTVALUEReadCallback   OnAfterReadInterface[DEFAULTVALUE]

	DEFINITIONs           map[*DEFINITION]any
	DEFINITIONs_mapString map[string]*DEFINITION

	// insertion point for slice of pointers maps

	OnAfterDEFINITIONCreateCallback OnAfterCreateInterface[DEFINITION]
	OnAfterDEFINITIONUpdateCallback OnAfterUpdateInterface[DEFINITION]
	OnAfterDEFINITIONDeleteCallback OnAfterDeleteInterface[DEFINITION]
	OnAfterDEFINITIONReadCallback   OnAfterReadInterface[DEFINITION]

	EDITABLEATTSs           map[*EDITABLEATTS]any
	EDITABLEATTSs_mapString map[string]*EDITABLEATTS

	// insertion point for slice of pointers maps

	OnAfterEDITABLEATTSCreateCallback OnAfterCreateInterface[EDITABLEATTS]
	OnAfterEDITABLEATTSUpdateCallback OnAfterUpdateInterface[EDITABLEATTS]
	OnAfterEDITABLEATTSDeleteCallback OnAfterDeleteInterface[EDITABLEATTS]
	OnAfterEDITABLEATTSReadCallback   OnAfterReadInterface[EDITABLEATTS]

	EMBEDDEDVALUEs           map[*EMBEDDEDVALUE]any
	EMBEDDEDVALUEs_mapString map[string]*EMBEDDEDVALUE

	// insertion point for slice of pointers maps

	OnAfterEMBEDDEDVALUECreateCallback OnAfterCreateInterface[EMBEDDEDVALUE]
	OnAfterEMBEDDEDVALUEUpdateCallback OnAfterUpdateInterface[EMBEDDEDVALUE]
	OnAfterEMBEDDEDVALUEDeleteCallback OnAfterDeleteInterface[EMBEDDEDVALUE]
	OnAfterEMBEDDEDVALUEReadCallback   OnAfterReadInterface[EMBEDDEDVALUE]

	ENUMVALUEs           map[*ENUMVALUE]any
	ENUMVALUEs_mapString map[string]*ENUMVALUE

	// insertion point for slice of pointers maps

	OnAfterENUMVALUECreateCallback OnAfterCreateInterface[ENUMVALUE]
	OnAfterENUMVALUEUpdateCallback OnAfterUpdateInterface[ENUMVALUE]
	OnAfterENUMVALUEDeleteCallback OnAfterDeleteInterface[ENUMVALUE]
	OnAfterENUMVALUEReadCallback   OnAfterReadInterface[ENUMVALUE]

	OBJECTs           map[*OBJECT]any
	OBJECTs_mapString map[string]*OBJECT

	// insertion point for slice of pointers maps

	OnAfterOBJECTCreateCallback OnAfterCreateInterface[OBJECT]
	OnAfterOBJECTUpdateCallback OnAfterUpdateInterface[OBJECT]
	OnAfterOBJECTDeleteCallback OnAfterDeleteInterface[OBJECT]
	OnAfterOBJECTReadCallback   OnAfterReadInterface[OBJECT]

	PROPERTIESs           map[*PROPERTIES]any
	PROPERTIESs_mapString map[string]*PROPERTIES

	// insertion point for slice of pointers maps

	OnAfterPROPERTIESCreateCallback OnAfterCreateInterface[PROPERTIES]
	OnAfterPROPERTIESUpdateCallback OnAfterUpdateInterface[PROPERTIES]
	OnAfterPROPERTIESDeleteCallback OnAfterDeleteInterface[PROPERTIES]
	OnAfterPROPERTIESReadCallback   OnAfterReadInterface[PROPERTIES]

	RELATIONGROUPs           map[*RELATIONGROUP]any
	RELATIONGROUPs_mapString map[string]*RELATIONGROUP

	// insertion point for slice of pointers maps

	OnAfterRELATIONGROUPCreateCallback OnAfterCreateInterface[RELATIONGROUP]
	OnAfterRELATIONGROUPUpdateCallback OnAfterUpdateInterface[RELATIONGROUP]
	OnAfterRELATIONGROUPDeleteCallback OnAfterDeleteInterface[RELATIONGROUP]
	OnAfterRELATIONGROUPReadCallback   OnAfterReadInterface[RELATIONGROUP]

	RELATIONGROUPTYPEs           map[*RELATIONGROUPTYPE]any
	RELATIONGROUPTYPEs_mapString map[string]*RELATIONGROUPTYPE

	// insertion point for slice of pointers maps

	OnAfterRELATIONGROUPTYPECreateCallback OnAfterCreateInterface[RELATIONGROUPTYPE]
	OnAfterRELATIONGROUPTYPEUpdateCallback OnAfterUpdateInterface[RELATIONGROUPTYPE]
	OnAfterRELATIONGROUPTYPEDeleteCallback OnAfterDeleteInterface[RELATIONGROUPTYPE]
	OnAfterRELATIONGROUPTYPEReadCallback   OnAfterReadInterface[RELATIONGROUPTYPE]

	REQIFs           map[*REQIF]any
	REQIFs_mapString map[string]*REQIF

	// insertion point for slice of pointers maps

	OnAfterREQIFCreateCallback OnAfterCreateInterface[REQIF]
	OnAfterREQIFUpdateCallback OnAfterUpdateInterface[REQIF]
	OnAfterREQIFDeleteCallback OnAfterDeleteInterface[REQIF]
	OnAfterREQIFReadCallback   OnAfterReadInterface[REQIF]

	REQIFCONTENTs           map[*REQIFCONTENT]any
	REQIFCONTENTs_mapString map[string]*REQIFCONTENT

	// insertion point for slice of pointers maps

	OnAfterREQIFCONTENTCreateCallback OnAfterCreateInterface[REQIFCONTENT]
	OnAfterREQIFCONTENTUpdateCallback OnAfterUpdateInterface[REQIFCONTENT]
	OnAfterREQIFCONTENTDeleteCallback OnAfterDeleteInterface[REQIFCONTENT]
	OnAfterREQIFCONTENTReadCallback   OnAfterReadInterface[REQIFCONTENT]

	REQIFHEADERs           map[*REQIFHEADER]any
	REQIFHEADERs_mapString map[string]*REQIFHEADER

	// insertion point for slice of pointers maps

	OnAfterREQIFHEADERCreateCallback OnAfterCreateInterface[REQIFHEADER]
	OnAfterREQIFHEADERUpdateCallback OnAfterUpdateInterface[REQIFHEADER]
	OnAfterREQIFHEADERDeleteCallback OnAfterDeleteInterface[REQIFHEADER]
	OnAfterREQIFHEADERReadCallback   OnAfterReadInterface[REQIFHEADER]

	REQIFTOOLEXTENSIONs           map[*REQIFTOOLEXTENSION]any
	REQIFTOOLEXTENSIONs_mapString map[string]*REQIFTOOLEXTENSION

	// insertion point for slice of pointers maps

	OnAfterREQIFTOOLEXTENSIONCreateCallback OnAfterCreateInterface[REQIFTOOLEXTENSION]
	OnAfterREQIFTOOLEXTENSIONUpdateCallback OnAfterUpdateInterface[REQIFTOOLEXTENSION]
	OnAfterREQIFTOOLEXTENSIONDeleteCallback OnAfterDeleteInterface[REQIFTOOLEXTENSION]
	OnAfterREQIFTOOLEXTENSIONReadCallback   OnAfterReadInterface[REQIFTOOLEXTENSION]

	REQTYPEs           map[*REQTYPE]any
	REQTYPEs_mapString map[string]*REQTYPE

	// insertion point for slice of pointers maps

	OnAfterREQTYPECreateCallback OnAfterCreateInterface[REQTYPE]
	OnAfterREQTYPEUpdateCallback OnAfterUpdateInterface[REQTYPE]
	OnAfterREQTYPEDeleteCallback OnAfterDeleteInterface[REQTYPE]
	OnAfterREQTYPEReadCallback   OnAfterReadInterface[REQTYPE]

	SOURCEs           map[*SOURCE]any
	SOURCEs_mapString map[string]*SOURCE

	// insertion point for slice of pointers maps

	OnAfterSOURCECreateCallback OnAfterCreateInterface[SOURCE]
	OnAfterSOURCEUpdateCallback OnAfterUpdateInterface[SOURCE]
	OnAfterSOURCEDeleteCallback OnAfterDeleteInterface[SOURCE]
	OnAfterSOURCEReadCallback   OnAfterReadInterface[SOURCE]

	SOURCESPECIFICATIONs           map[*SOURCESPECIFICATION]any
	SOURCESPECIFICATIONs_mapString map[string]*SOURCESPECIFICATION

	// insertion point for slice of pointers maps

	OnAfterSOURCESPECIFICATIONCreateCallback OnAfterCreateInterface[SOURCESPECIFICATION]
	OnAfterSOURCESPECIFICATIONUpdateCallback OnAfterUpdateInterface[SOURCESPECIFICATION]
	OnAfterSOURCESPECIFICATIONDeleteCallback OnAfterDeleteInterface[SOURCESPECIFICATION]
	OnAfterSOURCESPECIFICATIONReadCallback   OnAfterReadInterface[SOURCESPECIFICATION]

	SPECATTRIBUTESs           map[*SPECATTRIBUTES]any
	SPECATTRIBUTESs_mapString map[string]*SPECATTRIBUTES

	// insertion point for slice of pointers maps
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap map[*ATTRIBUTEDEFINITIONBOOLEAN]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap map[*ATTRIBUTEDEFINITIONDATE]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap map[*ATTRIBUTEDEFINITIONENUMERATION]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap map[*ATTRIBUTEDEFINITIONINTEGER]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap map[*ATTRIBUTEDEFINITIONREAL]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap map[*ATTRIBUTEDEFINITIONSTRING]*SPECATTRIBUTES
	SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap map[*ATTRIBUTEDEFINITIONXHTML]*SPECATTRIBUTES

	OnAfterSPECATTRIBUTESCreateCallback OnAfterCreateInterface[SPECATTRIBUTES]
	OnAfterSPECATTRIBUTESUpdateCallback OnAfterUpdateInterface[SPECATTRIBUTES]
	OnAfterSPECATTRIBUTESDeleteCallback OnAfterDeleteInterface[SPECATTRIBUTES]
	OnAfterSPECATTRIBUTESReadCallback   OnAfterReadInterface[SPECATTRIBUTES]

	SPECHIERARCHYs           map[*SPECHIERARCHY]any
	SPECHIERARCHYs_mapString map[string]*SPECHIERARCHY

	// insertion point for slice of pointers maps

	OnAfterSPECHIERARCHYCreateCallback OnAfterCreateInterface[SPECHIERARCHY]
	OnAfterSPECHIERARCHYUpdateCallback OnAfterUpdateInterface[SPECHIERARCHY]
	OnAfterSPECHIERARCHYDeleteCallback OnAfterDeleteInterface[SPECHIERARCHY]
	OnAfterSPECHIERARCHYReadCallback   OnAfterReadInterface[SPECHIERARCHY]

	SPECIFICATIONs           map[*SPECIFICATION]any
	SPECIFICATIONs_mapString map[string]*SPECIFICATION

	// insertion point for slice of pointers maps

	OnAfterSPECIFICATIONCreateCallback OnAfterCreateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONUpdateCallback OnAfterUpdateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONDeleteCallback OnAfterDeleteInterface[SPECIFICATION]
	OnAfterSPECIFICATIONReadCallback   OnAfterReadInterface[SPECIFICATION]

	SPECIFICATIONSs           map[*SPECIFICATIONS]any
	SPECIFICATIONSs_mapString map[string]*SPECIFICATIONS

	// insertion point for slice of pointers maps
	SPECIFICATIONS_SPECIFICATION_reverseMap map[*SPECIFICATION]*SPECIFICATIONS

	OnAfterSPECIFICATIONSCreateCallback OnAfterCreateInterface[SPECIFICATIONS]
	OnAfterSPECIFICATIONSUpdateCallback OnAfterUpdateInterface[SPECIFICATIONS]
	OnAfterSPECIFICATIONSDeleteCallback OnAfterDeleteInterface[SPECIFICATIONS]
	OnAfterSPECIFICATIONSReadCallback   OnAfterReadInterface[SPECIFICATIONS]

	SPECIFICATIONTYPEs           map[*SPECIFICATIONTYPE]any
	SPECIFICATIONTYPEs_mapString map[string]*SPECIFICATIONTYPE

	// insertion point for slice of pointers maps

	OnAfterSPECIFICATIONTYPECreateCallback OnAfterCreateInterface[SPECIFICATIONTYPE]
	OnAfterSPECIFICATIONTYPEUpdateCallback OnAfterUpdateInterface[SPECIFICATIONTYPE]
	OnAfterSPECIFICATIONTYPEDeleteCallback OnAfterDeleteInterface[SPECIFICATIONTYPE]
	OnAfterSPECIFICATIONTYPEReadCallback   OnAfterReadInterface[SPECIFICATIONTYPE]

	SPECIFIEDVALUESs           map[*SPECIFIEDVALUES]any
	SPECIFIEDVALUESs_mapString map[string]*SPECIFIEDVALUES

	// insertion point for slice of pointers maps
	SPECIFIEDVALUES_ENUMVALUE_reverseMap map[*ENUMVALUE]*SPECIFIEDVALUES

	OnAfterSPECIFIEDVALUESCreateCallback OnAfterCreateInterface[SPECIFIEDVALUES]
	OnAfterSPECIFIEDVALUESUpdateCallback OnAfterUpdateInterface[SPECIFIEDVALUES]
	OnAfterSPECIFIEDVALUESDeleteCallback OnAfterDeleteInterface[SPECIFIEDVALUES]
	OnAfterSPECIFIEDVALUESReadCallback   OnAfterReadInterface[SPECIFIEDVALUES]

	SPECOBJECTs           map[*SPECOBJECT]any
	SPECOBJECTs_mapString map[string]*SPECOBJECT

	// insertion point for slice of pointers maps

	OnAfterSPECOBJECTCreateCallback OnAfterCreateInterface[SPECOBJECT]
	OnAfterSPECOBJECTUpdateCallback OnAfterUpdateInterface[SPECOBJECT]
	OnAfterSPECOBJECTDeleteCallback OnAfterDeleteInterface[SPECOBJECT]
	OnAfterSPECOBJECTReadCallback   OnAfterReadInterface[SPECOBJECT]

	SPECOBJECTSs           map[*SPECOBJECTS]any
	SPECOBJECTSs_mapString map[string]*SPECOBJECTS

	// insertion point for slice of pointers maps
	SPECOBJECTS_SPECOBJECT_reverseMap map[*SPECOBJECT]*SPECOBJECTS

	OnAfterSPECOBJECTSCreateCallback OnAfterCreateInterface[SPECOBJECTS]
	OnAfterSPECOBJECTSUpdateCallback OnAfterUpdateInterface[SPECOBJECTS]
	OnAfterSPECOBJECTSDeleteCallback OnAfterDeleteInterface[SPECOBJECTS]
	OnAfterSPECOBJECTSReadCallback   OnAfterReadInterface[SPECOBJECTS]

	SPECOBJECTTYPEs           map[*SPECOBJECTTYPE]any
	SPECOBJECTTYPEs_mapString map[string]*SPECOBJECTTYPE

	// insertion point for slice of pointers maps

	OnAfterSPECOBJECTTYPECreateCallback OnAfterCreateInterface[SPECOBJECTTYPE]
	OnAfterSPECOBJECTTYPEUpdateCallback OnAfterUpdateInterface[SPECOBJECTTYPE]
	OnAfterSPECOBJECTTYPEDeleteCallback OnAfterDeleteInterface[SPECOBJECTTYPE]
	OnAfterSPECOBJECTTYPEReadCallback   OnAfterReadInterface[SPECOBJECTTYPE]

	SPECRELATIONs           map[*SPECRELATION]any
	SPECRELATIONs_mapString map[string]*SPECRELATION

	// insertion point for slice of pointers maps

	OnAfterSPECRELATIONCreateCallback OnAfterCreateInterface[SPECRELATION]
	OnAfterSPECRELATIONUpdateCallback OnAfterUpdateInterface[SPECRELATION]
	OnAfterSPECRELATIONDeleteCallback OnAfterDeleteInterface[SPECRELATION]
	OnAfterSPECRELATIONReadCallback   OnAfterReadInterface[SPECRELATION]

	SPECRELATIONGROUPSs           map[*SPECRELATIONGROUPS]any
	SPECRELATIONGROUPSs_mapString map[string]*SPECRELATIONGROUPS

	// insertion point for slice of pointers maps
	SPECRELATIONGROUPS_RELATIONGROUP_reverseMap map[*RELATIONGROUP]*SPECRELATIONGROUPS

	OnAfterSPECRELATIONGROUPSCreateCallback OnAfterCreateInterface[SPECRELATIONGROUPS]
	OnAfterSPECRELATIONGROUPSUpdateCallback OnAfterUpdateInterface[SPECRELATIONGROUPS]
	OnAfterSPECRELATIONGROUPSDeleteCallback OnAfterDeleteInterface[SPECRELATIONGROUPS]
	OnAfterSPECRELATIONGROUPSReadCallback   OnAfterReadInterface[SPECRELATIONGROUPS]

	SPECRELATIONSs           map[*SPECRELATIONS]any
	SPECRELATIONSs_mapString map[string]*SPECRELATIONS

	// insertion point for slice of pointers maps

	OnAfterSPECRELATIONSCreateCallback OnAfterCreateInterface[SPECRELATIONS]
	OnAfterSPECRELATIONSUpdateCallback OnAfterUpdateInterface[SPECRELATIONS]
	OnAfterSPECRELATIONSDeleteCallback OnAfterDeleteInterface[SPECRELATIONS]
	OnAfterSPECRELATIONSReadCallback   OnAfterReadInterface[SPECRELATIONS]

	SPECRELATIONTYPEs           map[*SPECRELATIONTYPE]any
	SPECRELATIONTYPEs_mapString map[string]*SPECRELATIONTYPE

	// insertion point for slice of pointers maps

	OnAfterSPECRELATIONTYPECreateCallback OnAfterCreateInterface[SPECRELATIONTYPE]
	OnAfterSPECRELATIONTYPEUpdateCallback OnAfterUpdateInterface[SPECRELATIONTYPE]
	OnAfterSPECRELATIONTYPEDeleteCallback OnAfterDeleteInterface[SPECRELATIONTYPE]
	OnAfterSPECRELATIONTYPEReadCallback   OnAfterReadInterface[SPECRELATIONTYPE]

	SPECTYPESs           map[*SPECTYPES]any
	SPECTYPESs_mapString map[string]*SPECTYPES

	// insertion point for slice of pointers maps
	SPECTYPES_RELATIONGROUPTYPE_reverseMap map[*RELATIONGROUPTYPE]*SPECTYPES
	SPECTYPES_SPECOBJECTTYPE_reverseMap map[*SPECOBJECTTYPE]*SPECTYPES
	SPECTYPES_SPECRELATIONTYPE_reverseMap map[*SPECRELATIONTYPE]*SPECTYPES
	SPECTYPES_SPECIFICATIONTYPE_reverseMap map[*SPECIFICATIONTYPE]*SPECTYPES

	OnAfterSPECTYPESCreateCallback OnAfterCreateInterface[SPECTYPES]
	OnAfterSPECTYPESUpdateCallback OnAfterUpdateInterface[SPECTYPES]
	OnAfterSPECTYPESDeleteCallback OnAfterDeleteInterface[SPECTYPES]
	OnAfterSPECTYPESReadCallback   OnAfterReadInterface[SPECTYPES]

	TARGETs           map[*TARGET]any
	TARGETs_mapString map[string]*TARGET

	// insertion point for slice of pointers maps

	OnAfterTARGETCreateCallback OnAfterCreateInterface[TARGET]
	OnAfterTARGETUpdateCallback OnAfterUpdateInterface[TARGET]
	OnAfterTARGETDeleteCallback OnAfterDeleteInterface[TARGET]
	OnAfterTARGETReadCallback   OnAfterReadInterface[TARGET]

	TARGETSPECIFICATIONs           map[*TARGETSPECIFICATION]any
	TARGETSPECIFICATIONs_mapString map[string]*TARGETSPECIFICATION

	// insertion point for slice of pointers maps

	OnAfterTARGETSPECIFICATIONCreateCallback OnAfterCreateInterface[TARGETSPECIFICATION]
	OnAfterTARGETSPECIFICATIONUpdateCallback OnAfterUpdateInterface[TARGETSPECIFICATION]
	OnAfterTARGETSPECIFICATIONDeleteCallback OnAfterDeleteInterface[TARGETSPECIFICATION]
	OnAfterTARGETSPECIFICATIONReadCallback   OnAfterReadInterface[TARGETSPECIFICATION]

	THEHEADERs           map[*THEHEADER]any
	THEHEADERs_mapString map[string]*THEHEADER

	// insertion point for slice of pointers maps

	OnAfterTHEHEADERCreateCallback OnAfterCreateInterface[THEHEADER]
	OnAfterTHEHEADERUpdateCallback OnAfterUpdateInterface[THEHEADER]
	OnAfterTHEHEADERDeleteCallback OnAfterDeleteInterface[THEHEADER]
	OnAfterTHEHEADERReadCallback   OnAfterReadInterface[THEHEADER]

	TOOLEXTENSIONSs           map[*TOOLEXTENSIONS]any
	TOOLEXTENSIONSs_mapString map[string]*TOOLEXTENSIONS

	// insertion point for slice of pointers maps
	TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap map[*REQIFTOOLEXTENSION]*TOOLEXTENSIONS

	OnAfterTOOLEXTENSIONSCreateCallback OnAfterCreateInterface[TOOLEXTENSIONS]
	OnAfterTOOLEXTENSIONSUpdateCallback OnAfterUpdateInterface[TOOLEXTENSIONS]
	OnAfterTOOLEXTENSIONSDeleteCallback OnAfterDeleteInterface[TOOLEXTENSIONS]
	OnAfterTOOLEXTENSIONSReadCallback   OnAfterReadInterface[TOOLEXTENSIONS]

	VALUESs           map[*VALUES]any
	VALUESs_mapString map[string]*VALUES

	// insertion point for slice of pointers maps

	OnAfterVALUESCreateCallback OnAfterCreateInterface[VALUES]
	OnAfterVALUESUpdateCallback OnAfterUpdateInterface[VALUES]
	OnAfterVALUESDeleteCallback OnAfterDeleteInterface[VALUES]
	OnAfterVALUESReadCallback   OnAfterReadInterface[VALUES]

	XHTMLCONTENTs           map[*XHTMLCONTENT]any
	XHTMLCONTENTs_mapString map[string]*XHTMLCONTENT

	// insertion point for slice of pointers maps

	OnAfterXHTMLCONTENTCreateCallback OnAfterCreateInterface[XHTMLCONTENT]
	OnAfterXHTMLCONTENTUpdateCallback OnAfterUpdateInterface[XHTMLCONTENT]
	OnAfterXHTMLCONTENTDeleteCallback OnAfterDeleteInterface[XHTMLCONTENT]
	OnAfterXHTMLCONTENTReadCallback   OnAfterReadInterface[XHTMLCONTENT]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/thomaspeugeot/gongreqif/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitALTERNATIVEID(alternativeid *ALTERNATIVEID)
	CheckoutALTERNATIVEID(alternativeid *ALTERNATIVEID)
	CommitATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN)
	CheckoutATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN)
	CommitATTRIBUTEDEFINITIONDATE(attributedefinitiondate *ATTRIBUTEDEFINITIONDATE)
	CheckoutATTRIBUTEDEFINITIONDATE(attributedefinitiondate *ATTRIBUTEDEFINITIONDATE)
	CommitATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION)
	CheckoutATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION)
	CommitATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER)
	CheckoutATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER)
	CommitATTRIBUTEDEFINITIONREAL(attributedefinitionreal *ATTRIBUTEDEFINITIONREAL)
	CheckoutATTRIBUTEDEFINITIONREAL(attributedefinitionreal *ATTRIBUTEDEFINITIONREAL)
	CommitATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING)
	CheckoutATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING)
	CommitATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML)
	CheckoutATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML)
	CommitATTRIBUTEVALUEBOOLEAN(attributevalueboolean *ATTRIBUTEVALUEBOOLEAN)
	CheckoutATTRIBUTEVALUEBOOLEAN(attributevalueboolean *ATTRIBUTEVALUEBOOLEAN)
	CommitATTRIBUTEVALUEDATE(attributevaluedate *ATTRIBUTEVALUEDATE)
	CheckoutATTRIBUTEVALUEDATE(attributevaluedate *ATTRIBUTEVALUEDATE)
	CommitATTRIBUTEVALUEENUMERATION(attributevalueenumeration *ATTRIBUTEVALUEENUMERATION)
	CheckoutATTRIBUTEVALUEENUMERATION(attributevalueenumeration *ATTRIBUTEVALUEENUMERATION)
	CommitATTRIBUTEVALUEINTEGER(attributevalueinteger *ATTRIBUTEVALUEINTEGER)
	CheckoutATTRIBUTEVALUEINTEGER(attributevalueinteger *ATTRIBUTEVALUEINTEGER)
	CommitATTRIBUTEVALUEREAL(attributevaluereal *ATTRIBUTEVALUEREAL)
	CheckoutATTRIBUTEVALUEREAL(attributevaluereal *ATTRIBUTEVALUEREAL)
	CommitATTRIBUTEVALUESTRING(attributevaluestring *ATTRIBUTEVALUESTRING)
	CheckoutATTRIBUTEVALUESTRING(attributevaluestring *ATTRIBUTEVALUESTRING)
	CommitATTRIBUTEVALUEXHTML(attributevaluexhtml *ATTRIBUTEVALUEXHTML)
	CheckoutATTRIBUTEVALUEXHTML(attributevaluexhtml *ATTRIBUTEVALUEXHTML)
	CommitCHILDREN(children *CHILDREN)
	CheckoutCHILDREN(children *CHILDREN)
	CommitCORECONTENT(corecontent *CORECONTENT)
	CheckoutCORECONTENT(corecontent *CORECONTENT)
	CommitDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN)
	CheckoutDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN)
	CommitDATATYPEDEFINITIONDATE(datatypedefinitiondate *DATATYPEDEFINITIONDATE)
	CheckoutDATATYPEDEFINITIONDATE(datatypedefinitiondate *DATATYPEDEFINITIONDATE)
	CommitDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION)
	CheckoutDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION)
	CommitDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER)
	CheckoutDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER)
	CommitDATATYPEDEFINITIONREAL(datatypedefinitionreal *DATATYPEDEFINITIONREAL)
	CheckoutDATATYPEDEFINITIONREAL(datatypedefinitionreal *DATATYPEDEFINITIONREAL)
	CommitDATATYPEDEFINITIONSTRING(datatypedefinitionstring *DATATYPEDEFINITIONSTRING)
	CheckoutDATATYPEDEFINITIONSTRING(datatypedefinitionstring *DATATYPEDEFINITIONSTRING)
	CommitDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML)
	CheckoutDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML)
	CommitDATATYPES(datatypes *DATATYPES)
	CheckoutDATATYPES(datatypes *DATATYPES)
	CommitDEFAULTVALUE(defaultvalue *DEFAULTVALUE)
	CheckoutDEFAULTVALUE(defaultvalue *DEFAULTVALUE)
	CommitDEFINITION(definition *DEFINITION)
	CheckoutDEFINITION(definition *DEFINITION)
	CommitEDITABLEATTS(editableatts *EDITABLEATTS)
	CheckoutEDITABLEATTS(editableatts *EDITABLEATTS)
	CommitEMBEDDEDVALUE(embeddedvalue *EMBEDDEDVALUE)
	CheckoutEMBEDDEDVALUE(embeddedvalue *EMBEDDEDVALUE)
	CommitENUMVALUE(enumvalue *ENUMVALUE)
	CheckoutENUMVALUE(enumvalue *ENUMVALUE)
	CommitOBJECT(object *OBJECT)
	CheckoutOBJECT(object *OBJECT)
	CommitPROPERTIES(properties *PROPERTIES)
	CheckoutPROPERTIES(properties *PROPERTIES)
	CommitRELATIONGROUP(relationgroup *RELATIONGROUP)
	CheckoutRELATIONGROUP(relationgroup *RELATIONGROUP)
	CommitRELATIONGROUPTYPE(relationgrouptype *RELATIONGROUPTYPE)
	CheckoutRELATIONGROUPTYPE(relationgrouptype *RELATIONGROUPTYPE)
	CommitREQIF(reqif *REQIF)
	CheckoutREQIF(reqif *REQIF)
	CommitREQIFCONTENT(reqifcontent *REQIFCONTENT)
	CheckoutREQIFCONTENT(reqifcontent *REQIFCONTENT)
	CommitREQIFHEADER(reqifheader *REQIFHEADER)
	CheckoutREQIFHEADER(reqifheader *REQIFHEADER)
	CommitREQIFTOOLEXTENSION(reqiftoolextension *REQIFTOOLEXTENSION)
	CheckoutREQIFTOOLEXTENSION(reqiftoolextension *REQIFTOOLEXTENSION)
	CommitREQTYPE(reqtype *REQTYPE)
	CheckoutREQTYPE(reqtype *REQTYPE)
	CommitSOURCE(source *SOURCE)
	CheckoutSOURCE(source *SOURCE)
	CommitSOURCESPECIFICATION(sourcespecification *SOURCESPECIFICATION)
	CheckoutSOURCESPECIFICATION(sourcespecification *SOURCESPECIFICATION)
	CommitSPECATTRIBUTES(specattributes *SPECATTRIBUTES)
	CheckoutSPECATTRIBUTES(specattributes *SPECATTRIBUTES)
	CommitSPECHIERARCHY(spechierarchy *SPECHIERARCHY)
	CheckoutSPECHIERARCHY(spechierarchy *SPECHIERARCHY)
	CommitSPECIFICATION(specification *SPECIFICATION)
	CheckoutSPECIFICATION(specification *SPECIFICATION)
	CommitSPECIFICATIONS(specifications *SPECIFICATIONS)
	CheckoutSPECIFICATIONS(specifications *SPECIFICATIONS)
	CommitSPECIFICATIONTYPE(specificationtype *SPECIFICATIONTYPE)
	CheckoutSPECIFICATIONTYPE(specificationtype *SPECIFICATIONTYPE)
	CommitSPECIFIEDVALUES(specifiedvalues *SPECIFIEDVALUES)
	CheckoutSPECIFIEDVALUES(specifiedvalues *SPECIFIEDVALUES)
	CommitSPECOBJECT(specobject *SPECOBJECT)
	CheckoutSPECOBJECT(specobject *SPECOBJECT)
	CommitSPECOBJECTS(specobjects *SPECOBJECTS)
	CheckoutSPECOBJECTS(specobjects *SPECOBJECTS)
	CommitSPECOBJECTTYPE(specobjecttype *SPECOBJECTTYPE)
	CheckoutSPECOBJECTTYPE(specobjecttype *SPECOBJECTTYPE)
	CommitSPECRELATION(specrelation *SPECRELATION)
	CheckoutSPECRELATION(specrelation *SPECRELATION)
	CommitSPECRELATIONGROUPS(specrelationgroups *SPECRELATIONGROUPS)
	CheckoutSPECRELATIONGROUPS(specrelationgroups *SPECRELATIONGROUPS)
	CommitSPECRELATIONS(specrelations *SPECRELATIONS)
	CheckoutSPECRELATIONS(specrelations *SPECRELATIONS)
	CommitSPECRELATIONTYPE(specrelationtype *SPECRELATIONTYPE)
	CheckoutSPECRELATIONTYPE(specrelationtype *SPECRELATIONTYPE)
	CommitSPECTYPES(spectypes *SPECTYPES)
	CheckoutSPECTYPES(spectypes *SPECTYPES)
	CommitTARGET(target *TARGET)
	CheckoutTARGET(target *TARGET)
	CommitTARGETSPECIFICATION(targetspecification *TARGETSPECIFICATION)
	CheckoutTARGETSPECIFICATION(targetspecification *TARGETSPECIFICATION)
	CommitTHEHEADER(theheader *THEHEADER)
	CheckoutTHEHEADER(theheader *THEHEADER)
	CommitTOOLEXTENSIONS(toolextensions *TOOLEXTENSIONS)
	CheckoutTOOLEXTENSIONS(toolextensions *TOOLEXTENSIONS)
	CommitVALUES(values *VALUES)
	CheckoutVALUES(values *VALUES)
	CommitXHTMLCONTENT(xhtmlcontent *XHTMLCONTENT)
	CheckoutXHTMLCONTENT(xhtmlcontent *XHTMLCONTENT)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		ALTERNATIVEIDs:           make(map[*ALTERNATIVEID]any),
		ALTERNATIVEIDs_mapString: make(map[string]*ALTERNATIVEID),

		ATTRIBUTEDEFINITIONBOOLEANs:           make(map[*ATTRIBUTEDEFINITIONBOOLEAN]any),
		ATTRIBUTEDEFINITIONBOOLEANs_mapString: make(map[string]*ATTRIBUTEDEFINITIONBOOLEAN),

		ATTRIBUTEDEFINITIONDATEs:           make(map[*ATTRIBUTEDEFINITIONDATE]any),
		ATTRIBUTEDEFINITIONDATEs_mapString: make(map[string]*ATTRIBUTEDEFINITIONDATE),

		ATTRIBUTEDEFINITIONENUMERATIONs:           make(map[*ATTRIBUTEDEFINITIONENUMERATION]any),
		ATTRIBUTEDEFINITIONENUMERATIONs_mapString: make(map[string]*ATTRIBUTEDEFINITIONENUMERATION),

		ATTRIBUTEDEFINITIONINTEGERs:           make(map[*ATTRIBUTEDEFINITIONINTEGER]any),
		ATTRIBUTEDEFINITIONINTEGERs_mapString: make(map[string]*ATTRIBUTEDEFINITIONINTEGER),

		ATTRIBUTEDEFINITIONREALs:           make(map[*ATTRIBUTEDEFINITIONREAL]any),
		ATTRIBUTEDEFINITIONREALs_mapString: make(map[string]*ATTRIBUTEDEFINITIONREAL),

		ATTRIBUTEDEFINITIONSTRINGs:           make(map[*ATTRIBUTEDEFINITIONSTRING]any),
		ATTRIBUTEDEFINITIONSTRINGs_mapString: make(map[string]*ATTRIBUTEDEFINITIONSTRING),

		ATTRIBUTEDEFINITIONXHTMLs:           make(map[*ATTRIBUTEDEFINITIONXHTML]any),
		ATTRIBUTEDEFINITIONXHTMLs_mapString: make(map[string]*ATTRIBUTEDEFINITIONXHTML),

		ATTRIBUTEVALUEBOOLEANs:           make(map[*ATTRIBUTEVALUEBOOLEAN]any),
		ATTRIBUTEVALUEBOOLEANs_mapString: make(map[string]*ATTRIBUTEVALUEBOOLEAN),

		ATTRIBUTEVALUEDATEs:           make(map[*ATTRIBUTEVALUEDATE]any),
		ATTRIBUTEVALUEDATEs_mapString: make(map[string]*ATTRIBUTEVALUEDATE),

		ATTRIBUTEVALUEENUMERATIONs:           make(map[*ATTRIBUTEVALUEENUMERATION]any),
		ATTRIBUTEVALUEENUMERATIONs_mapString: make(map[string]*ATTRIBUTEVALUEENUMERATION),

		ATTRIBUTEVALUEINTEGERs:           make(map[*ATTRIBUTEVALUEINTEGER]any),
		ATTRIBUTEVALUEINTEGERs_mapString: make(map[string]*ATTRIBUTEVALUEINTEGER),

		ATTRIBUTEVALUEREALs:           make(map[*ATTRIBUTEVALUEREAL]any),
		ATTRIBUTEVALUEREALs_mapString: make(map[string]*ATTRIBUTEVALUEREAL),

		ATTRIBUTEVALUESTRINGs:           make(map[*ATTRIBUTEVALUESTRING]any),
		ATTRIBUTEVALUESTRINGs_mapString: make(map[string]*ATTRIBUTEVALUESTRING),

		ATTRIBUTEVALUEXHTMLs:           make(map[*ATTRIBUTEVALUEXHTML]any),
		ATTRIBUTEVALUEXHTMLs_mapString: make(map[string]*ATTRIBUTEVALUEXHTML),

		CHILDRENs:           make(map[*CHILDREN]any),
		CHILDRENs_mapString: make(map[string]*CHILDREN),

		CORECONTENTs:           make(map[*CORECONTENT]any),
		CORECONTENTs_mapString: make(map[string]*CORECONTENT),

		DATATYPEDEFINITIONBOOLEANs:           make(map[*DATATYPEDEFINITIONBOOLEAN]any),
		DATATYPEDEFINITIONBOOLEANs_mapString: make(map[string]*DATATYPEDEFINITIONBOOLEAN),

		DATATYPEDEFINITIONDATEs:           make(map[*DATATYPEDEFINITIONDATE]any),
		DATATYPEDEFINITIONDATEs_mapString: make(map[string]*DATATYPEDEFINITIONDATE),

		DATATYPEDEFINITIONENUMERATIONs:           make(map[*DATATYPEDEFINITIONENUMERATION]any),
		DATATYPEDEFINITIONENUMERATIONs_mapString: make(map[string]*DATATYPEDEFINITIONENUMERATION),

		DATATYPEDEFINITIONINTEGERs:           make(map[*DATATYPEDEFINITIONINTEGER]any),
		DATATYPEDEFINITIONINTEGERs_mapString: make(map[string]*DATATYPEDEFINITIONINTEGER),

		DATATYPEDEFINITIONREALs:           make(map[*DATATYPEDEFINITIONREAL]any),
		DATATYPEDEFINITIONREALs_mapString: make(map[string]*DATATYPEDEFINITIONREAL),

		DATATYPEDEFINITIONSTRINGs:           make(map[*DATATYPEDEFINITIONSTRING]any),
		DATATYPEDEFINITIONSTRINGs_mapString: make(map[string]*DATATYPEDEFINITIONSTRING),

		DATATYPEDEFINITIONXHTMLs:           make(map[*DATATYPEDEFINITIONXHTML]any),
		DATATYPEDEFINITIONXHTMLs_mapString: make(map[string]*DATATYPEDEFINITIONXHTML),

		DATATYPESs:           make(map[*DATATYPES]any),
		DATATYPESs_mapString: make(map[string]*DATATYPES),

		DEFAULTVALUEs:           make(map[*DEFAULTVALUE]any),
		DEFAULTVALUEs_mapString: make(map[string]*DEFAULTVALUE),

		DEFINITIONs:           make(map[*DEFINITION]any),
		DEFINITIONs_mapString: make(map[string]*DEFINITION),

		EDITABLEATTSs:           make(map[*EDITABLEATTS]any),
		EDITABLEATTSs_mapString: make(map[string]*EDITABLEATTS),

		EMBEDDEDVALUEs:           make(map[*EMBEDDEDVALUE]any),
		EMBEDDEDVALUEs_mapString: make(map[string]*EMBEDDEDVALUE),

		ENUMVALUEs:           make(map[*ENUMVALUE]any),
		ENUMVALUEs_mapString: make(map[string]*ENUMVALUE),

		OBJECTs:           make(map[*OBJECT]any),
		OBJECTs_mapString: make(map[string]*OBJECT),

		PROPERTIESs:           make(map[*PROPERTIES]any),
		PROPERTIESs_mapString: make(map[string]*PROPERTIES),

		RELATIONGROUPs:           make(map[*RELATIONGROUP]any),
		RELATIONGROUPs_mapString: make(map[string]*RELATIONGROUP),

		RELATIONGROUPTYPEs:           make(map[*RELATIONGROUPTYPE]any),
		RELATIONGROUPTYPEs_mapString: make(map[string]*RELATIONGROUPTYPE),

		REQIFs:           make(map[*REQIF]any),
		REQIFs_mapString: make(map[string]*REQIF),

		REQIFCONTENTs:           make(map[*REQIFCONTENT]any),
		REQIFCONTENTs_mapString: make(map[string]*REQIFCONTENT),

		REQIFHEADERs:           make(map[*REQIFHEADER]any),
		REQIFHEADERs_mapString: make(map[string]*REQIFHEADER),

		REQIFTOOLEXTENSIONs:           make(map[*REQIFTOOLEXTENSION]any),
		REQIFTOOLEXTENSIONs_mapString: make(map[string]*REQIFTOOLEXTENSION),

		REQTYPEs:           make(map[*REQTYPE]any),
		REQTYPEs_mapString: make(map[string]*REQTYPE),

		SOURCEs:           make(map[*SOURCE]any),
		SOURCEs_mapString: make(map[string]*SOURCE),

		SOURCESPECIFICATIONs:           make(map[*SOURCESPECIFICATION]any),
		SOURCESPECIFICATIONs_mapString: make(map[string]*SOURCESPECIFICATION),

		SPECATTRIBUTESs:           make(map[*SPECATTRIBUTES]any),
		SPECATTRIBUTESs_mapString: make(map[string]*SPECATTRIBUTES),

		SPECHIERARCHYs:           make(map[*SPECHIERARCHY]any),
		SPECHIERARCHYs_mapString: make(map[string]*SPECHIERARCHY),

		SPECIFICATIONs:           make(map[*SPECIFICATION]any),
		SPECIFICATIONs_mapString: make(map[string]*SPECIFICATION),

		SPECIFICATIONSs:           make(map[*SPECIFICATIONS]any),
		SPECIFICATIONSs_mapString: make(map[string]*SPECIFICATIONS),

		SPECIFICATIONTYPEs:           make(map[*SPECIFICATIONTYPE]any),
		SPECIFICATIONTYPEs_mapString: make(map[string]*SPECIFICATIONTYPE),

		SPECIFIEDVALUESs:           make(map[*SPECIFIEDVALUES]any),
		SPECIFIEDVALUESs_mapString: make(map[string]*SPECIFIEDVALUES),

		SPECOBJECTs:           make(map[*SPECOBJECT]any),
		SPECOBJECTs_mapString: make(map[string]*SPECOBJECT),

		SPECOBJECTSs:           make(map[*SPECOBJECTS]any),
		SPECOBJECTSs_mapString: make(map[string]*SPECOBJECTS),

		SPECOBJECTTYPEs:           make(map[*SPECOBJECTTYPE]any),
		SPECOBJECTTYPEs_mapString: make(map[string]*SPECOBJECTTYPE),

		SPECRELATIONs:           make(map[*SPECRELATION]any),
		SPECRELATIONs_mapString: make(map[string]*SPECRELATION),

		SPECRELATIONGROUPSs:           make(map[*SPECRELATIONGROUPS]any),
		SPECRELATIONGROUPSs_mapString: make(map[string]*SPECRELATIONGROUPS),

		SPECRELATIONSs:           make(map[*SPECRELATIONS]any),
		SPECRELATIONSs_mapString: make(map[string]*SPECRELATIONS),

		SPECRELATIONTYPEs:           make(map[*SPECRELATIONTYPE]any),
		SPECRELATIONTYPEs_mapString: make(map[string]*SPECRELATIONTYPE),

		SPECTYPESs:           make(map[*SPECTYPES]any),
		SPECTYPESs_mapString: make(map[string]*SPECTYPES),

		TARGETs:           make(map[*TARGET]any),
		TARGETs_mapString: make(map[string]*TARGET),

		TARGETSPECIFICATIONs:           make(map[*TARGETSPECIFICATION]any),
		TARGETSPECIFICATIONs_mapString: make(map[string]*TARGETSPECIFICATION),

		THEHEADERs:           make(map[*THEHEADER]any),
		THEHEADERs_mapString: make(map[string]*THEHEADER),

		TOOLEXTENSIONSs:           make(map[*TOOLEXTENSIONS]any),
		TOOLEXTENSIONSs_mapString: make(map[string]*TOOLEXTENSIONS),

		VALUESs:           make(map[*VALUES]any),
		VALUESs_mapString: make(map[string]*VALUES),

		XHTMLCONTENTs:           make(map[*XHTMLCONTENT]any),
		XHTMLCONTENTs_mapString: make(map[string]*XHTMLCONTENT),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ALTERNATIVEID"] = len(stage.ALTERNATIVEIDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONBOOLEAN"] = len(stage.ATTRIBUTEDEFINITIONBOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONDATE"] = len(stage.ATTRIBUTEDEFINITIONDATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONENUMERATION"] = len(stage.ATTRIBUTEDEFINITIONENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONINTEGER"] = len(stage.ATTRIBUTEDEFINITIONINTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONREAL"] = len(stage.ATTRIBUTEDEFINITIONREALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONSTRING"] = len(stage.ATTRIBUTEDEFINITIONSTRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONXHTML"] = len(stage.ATTRIBUTEDEFINITIONXHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEBOOLEAN"] = len(stage.ATTRIBUTEVALUEBOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEDATE"] = len(stage.ATTRIBUTEVALUEDATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEENUMERATION"] = len(stage.ATTRIBUTEVALUEENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEINTEGER"] = len(stage.ATTRIBUTEVALUEINTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEREAL"] = len(stage.ATTRIBUTEVALUEREALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUESTRING"] = len(stage.ATTRIBUTEVALUESTRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEXHTML"] = len(stage.ATTRIBUTEVALUEXHTMLs)
	stage.Map_GongStructName_InstancesNb["CHILDREN"] = len(stage.CHILDRENs)
	stage.Map_GongStructName_InstancesNb["CORECONTENT"] = len(stage.CORECONTENTs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONBOOLEAN"] = len(stage.DATATYPEDEFINITIONBOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONDATE"] = len(stage.DATATYPEDEFINITIONDATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONENUMERATION"] = len(stage.DATATYPEDEFINITIONENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONINTEGER"] = len(stage.DATATYPEDEFINITIONINTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONREAL"] = len(stage.DATATYPEDEFINITIONREALs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONSTRING"] = len(stage.DATATYPEDEFINITIONSTRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONXHTML"] = len(stage.DATATYPEDEFINITIONXHTMLs)
	stage.Map_GongStructName_InstancesNb["DATATYPES"] = len(stage.DATATYPESs)
	stage.Map_GongStructName_InstancesNb["DEFAULTVALUE"] = len(stage.DEFAULTVALUEs)
	stage.Map_GongStructName_InstancesNb["DEFINITION"] = len(stage.DEFINITIONs)
	stage.Map_GongStructName_InstancesNb["EDITABLEATTS"] = len(stage.EDITABLEATTSs)
	stage.Map_GongStructName_InstancesNb["EMBEDDEDVALUE"] = len(stage.EMBEDDEDVALUEs)
	stage.Map_GongStructName_InstancesNb["ENUMVALUE"] = len(stage.ENUMVALUEs)
	stage.Map_GongStructName_InstancesNb["OBJECT"] = len(stage.OBJECTs)
	stage.Map_GongStructName_InstancesNb["PROPERTIES"] = len(stage.PROPERTIESs)
	stage.Map_GongStructName_InstancesNb["RELATIONGROUP"] = len(stage.RELATIONGROUPs)
	stage.Map_GongStructName_InstancesNb["RELATIONGROUPTYPE"] = len(stage.RELATIONGROUPTYPEs)
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)
	stage.Map_GongStructName_InstancesNb["REQIFCONTENT"] = len(stage.REQIFCONTENTs)
	stage.Map_GongStructName_InstancesNb["REQIFHEADER"] = len(stage.REQIFHEADERs)
	stage.Map_GongStructName_InstancesNb["REQIFTOOLEXTENSION"] = len(stage.REQIFTOOLEXTENSIONs)
	stage.Map_GongStructName_InstancesNb["REQTYPE"] = len(stage.REQTYPEs)
	stage.Map_GongStructName_InstancesNb["SOURCE"] = len(stage.SOURCEs)
	stage.Map_GongStructName_InstancesNb["SOURCESPECIFICATION"] = len(stage.SOURCESPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECATTRIBUTES"] = len(stage.SPECATTRIBUTESs)
	stage.Map_GongStructName_InstancesNb["SPECHIERARCHY"] = len(stage.SPECHIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATIONS"] = len(stage.SPECIFICATIONSs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATIONTYPE"] = len(stage.SPECIFICATIONTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECIFIEDVALUES"] = len(stage.SPECIFIEDVALUESs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECT"] = len(stage.SPECOBJECTs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECTS"] = len(stage.SPECOBJECTSs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECTTYPE"] = len(stage.SPECOBJECTTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECRELATION"] = len(stage.SPECRELATIONs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONGROUPS"] = len(stage.SPECRELATIONGROUPSs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONS"] = len(stage.SPECRELATIONSs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONTYPE"] = len(stage.SPECRELATIONTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECTYPES"] = len(stage.SPECTYPESs)
	stage.Map_GongStructName_InstancesNb["TARGET"] = len(stage.TARGETs)
	stage.Map_GongStructName_InstancesNb["TARGETSPECIFICATION"] = len(stage.TARGETSPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["THEHEADER"] = len(stage.THEHEADERs)
	stage.Map_GongStructName_InstancesNb["TOOLEXTENSIONS"] = len(stage.TOOLEXTENSIONSs)
	stage.Map_GongStructName_InstancesNb["VALUES"] = len(stage.VALUESs)
	stage.Map_GongStructName_InstancesNb["XHTMLCONTENT"] = len(stage.XHTMLCONTENTs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ALTERNATIVEID"] = len(stage.ALTERNATIVEIDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONBOOLEAN"] = len(stage.ATTRIBUTEDEFINITIONBOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONDATE"] = len(stage.ATTRIBUTEDEFINITIONDATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONENUMERATION"] = len(stage.ATTRIBUTEDEFINITIONENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONINTEGER"] = len(stage.ATTRIBUTEDEFINITIONINTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONREAL"] = len(stage.ATTRIBUTEDEFINITIONREALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONSTRING"] = len(stage.ATTRIBUTEDEFINITIONSTRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEDEFINITIONXHTML"] = len(stage.ATTRIBUTEDEFINITIONXHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEBOOLEAN"] = len(stage.ATTRIBUTEVALUEBOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEDATE"] = len(stage.ATTRIBUTEVALUEDATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEENUMERATION"] = len(stage.ATTRIBUTEVALUEENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEINTEGER"] = len(stage.ATTRIBUTEVALUEINTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEREAL"] = len(stage.ATTRIBUTEVALUEREALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUESTRING"] = len(stage.ATTRIBUTEVALUESTRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTEVALUEXHTML"] = len(stage.ATTRIBUTEVALUEXHTMLs)
	stage.Map_GongStructName_InstancesNb["CHILDREN"] = len(stage.CHILDRENs)
	stage.Map_GongStructName_InstancesNb["CORECONTENT"] = len(stage.CORECONTENTs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONBOOLEAN"] = len(stage.DATATYPEDEFINITIONBOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONDATE"] = len(stage.DATATYPEDEFINITIONDATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONENUMERATION"] = len(stage.DATATYPEDEFINITIONENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONINTEGER"] = len(stage.DATATYPEDEFINITIONINTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONREAL"] = len(stage.DATATYPEDEFINITIONREALs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONSTRING"] = len(stage.DATATYPEDEFINITIONSTRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPEDEFINITIONXHTML"] = len(stage.DATATYPEDEFINITIONXHTMLs)
	stage.Map_GongStructName_InstancesNb["DATATYPES"] = len(stage.DATATYPESs)
	stage.Map_GongStructName_InstancesNb["DEFAULTVALUE"] = len(stage.DEFAULTVALUEs)
	stage.Map_GongStructName_InstancesNb["DEFINITION"] = len(stage.DEFINITIONs)
	stage.Map_GongStructName_InstancesNb["EDITABLEATTS"] = len(stage.EDITABLEATTSs)
	stage.Map_GongStructName_InstancesNb["EMBEDDEDVALUE"] = len(stage.EMBEDDEDVALUEs)
	stage.Map_GongStructName_InstancesNb["ENUMVALUE"] = len(stage.ENUMVALUEs)
	stage.Map_GongStructName_InstancesNb["OBJECT"] = len(stage.OBJECTs)
	stage.Map_GongStructName_InstancesNb["PROPERTIES"] = len(stage.PROPERTIESs)
	stage.Map_GongStructName_InstancesNb["RELATIONGROUP"] = len(stage.RELATIONGROUPs)
	stage.Map_GongStructName_InstancesNb["RELATIONGROUPTYPE"] = len(stage.RELATIONGROUPTYPEs)
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)
	stage.Map_GongStructName_InstancesNb["REQIFCONTENT"] = len(stage.REQIFCONTENTs)
	stage.Map_GongStructName_InstancesNb["REQIFHEADER"] = len(stage.REQIFHEADERs)
	stage.Map_GongStructName_InstancesNb["REQIFTOOLEXTENSION"] = len(stage.REQIFTOOLEXTENSIONs)
	stage.Map_GongStructName_InstancesNb["REQTYPE"] = len(stage.REQTYPEs)
	stage.Map_GongStructName_InstancesNb["SOURCE"] = len(stage.SOURCEs)
	stage.Map_GongStructName_InstancesNb["SOURCESPECIFICATION"] = len(stage.SOURCESPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECATTRIBUTES"] = len(stage.SPECATTRIBUTESs)
	stage.Map_GongStructName_InstancesNb["SPECHIERARCHY"] = len(stage.SPECHIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATIONS"] = len(stage.SPECIFICATIONSs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATIONTYPE"] = len(stage.SPECIFICATIONTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECIFIEDVALUES"] = len(stage.SPECIFIEDVALUESs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECT"] = len(stage.SPECOBJECTs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECTS"] = len(stage.SPECOBJECTSs)
	stage.Map_GongStructName_InstancesNb["SPECOBJECTTYPE"] = len(stage.SPECOBJECTTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECRELATION"] = len(stage.SPECRELATIONs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONGROUPS"] = len(stage.SPECRELATIONGROUPSs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONS"] = len(stage.SPECRELATIONSs)
	stage.Map_GongStructName_InstancesNb["SPECRELATIONTYPE"] = len(stage.SPECRELATIONTYPEs)
	stage.Map_GongStructName_InstancesNb["SPECTYPES"] = len(stage.SPECTYPESs)
	stage.Map_GongStructName_InstancesNb["TARGET"] = len(stage.TARGETs)
	stage.Map_GongStructName_InstancesNb["TARGETSPECIFICATION"] = len(stage.TARGETSPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["THEHEADER"] = len(stage.THEHEADERs)
	stage.Map_GongStructName_InstancesNb["TOOLEXTENSIONS"] = len(stage.TOOLEXTENSIONSs)
	stage.Map_GongStructName_InstancesNb["VALUES"] = len(stage.VALUESs)
	stage.Map_GongStructName_InstancesNb["XHTMLCONTENT"] = len(stage.XHTMLCONTENTs)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts alternativeid to the model stage
func (alternativeid *ALTERNATIVEID) Stage(stage *StageStruct) *ALTERNATIVEID {
	stage.ALTERNATIVEIDs[alternativeid] = __member
	stage.ALTERNATIVEIDs_mapString[alternativeid.Name] = alternativeid

	return alternativeid
}

// Unstage removes alternativeid off the model stage
func (alternativeid *ALTERNATIVEID) Unstage(stage *StageStruct) *ALTERNATIVEID {
	delete(stage.ALTERNATIVEIDs, alternativeid)
	delete(stage.ALTERNATIVEIDs_mapString, alternativeid.Name)
	return alternativeid
}

// UnstageVoid removes alternativeid off the model stage
func (alternativeid *ALTERNATIVEID) UnstageVoid(stage *StageStruct) {
	delete(stage.ALTERNATIVEIDs, alternativeid)
	delete(stage.ALTERNATIVEIDs_mapString, alternativeid.Name)
}

// commit alternativeid to the back repo (if it is already staged)
func (alternativeid *ALTERNATIVEID) Commit(stage *StageStruct) *ALTERNATIVEID {
	if _, ok := stage.ALTERNATIVEIDs[alternativeid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitALTERNATIVEID(alternativeid)
		}
	}
	return alternativeid
}

func (alternativeid *ALTERNATIVEID) CommitVoid(stage *StageStruct) {
	alternativeid.Commit(stage)
}

// Checkout alternativeid to the back repo (if it is already staged)
func (alternativeid *ALTERNATIVEID) Checkout(stage *StageStruct) *ALTERNATIVEID {
	if _, ok := stage.ALTERNATIVEIDs[alternativeid]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutALTERNATIVEID(alternativeid)
		}
	}
	return alternativeid
}

// for satisfaction of GongStruct interface
func (alternativeid *ALTERNATIVEID) GetName() (res string) {
	return alternativeid.Name
}

// Stage puts attributedefinitionboolean to the model stage
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONBOOLEAN {
	stage.ATTRIBUTEDEFINITIONBOOLEANs[attributedefinitionboolean] = __member
	stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString[attributedefinitionboolean.Name] = attributedefinitionboolean

	return attributedefinitionboolean
}

// Unstage removes attributedefinitionboolean off the model stage
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONBOOLEAN {
	delete(stage.ATTRIBUTEDEFINITIONBOOLEANs, attributedefinitionboolean)
	delete(stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString, attributedefinitionboolean.Name)
	return attributedefinitionboolean
}

// UnstageVoid removes attributedefinitionboolean off the model stage
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONBOOLEANs, attributedefinitionboolean)
	delete(stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString, attributedefinitionboolean.Name)
}

// commit attributedefinitionboolean to the back repo (if it is already staged)
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONBOOLEAN {
	if _, ok := stage.ATTRIBUTEDEFINITIONBOOLEANs[attributedefinitionboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean)
		}
	}
	return attributedefinitionboolean
}

func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) CommitVoid(stage *StageStruct) {
	attributedefinitionboolean.Commit(stage)
}

// Checkout attributedefinitionboolean to the back repo (if it is already staged)
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONBOOLEAN {
	if _, ok := stage.ATTRIBUTEDEFINITIONBOOLEANs[attributedefinitionboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean)
		}
	}
	return attributedefinitionboolean
}

// for satisfaction of GongStruct interface
func (attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) GetName() (res string) {
	return attributedefinitionboolean.Name
}

// Stage puts attributedefinitiondate to the model stage
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONDATE {
	stage.ATTRIBUTEDEFINITIONDATEs[attributedefinitiondate] = __member
	stage.ATTRIBUTEDEFINITIONDATEs_mapString[attributedefinitiondate.Name] = attributedefinitiondate

	return attributedefinitiondate
}

// Unstage removes attributedefinitiondate off the model stage
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONDATE {
	delete(stage.ATTRIBUTEDEFINITIONDATEs, attributedefinitiondate)
	delete(stage.ATTRIBUTEDEFINITIONDATEs_mapString, attributedefinitiondate.Name)
	return attributedefinitiondate
}

// UnstageVoid removes attributedefinitiondate off the model stage
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONDATEs, attributedefinitiondate)
	delete(stage.ATTRIBUTEDEFINITIONDATEs_mapString, attributedefinitiondate.Name)
}

// commit attributedefinitiondate to the back repo (if it is already staged)
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONDATE {
	if _, ok := stage.ATTRIBUTEDEFINITIONDATEs[attributedefinitiondate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONDATE(attributedefinitiondate)
		}
	}
	return attributedefinitiondate
}

func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) CommitVoid(stage *StageStruct) {
	attributedefinitiondate.Commit(stage)
}

// Checkout attributedefinitiondate to the back repo (if it is already staged)
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONDATE {
	if _, ok := stage.ATTRIBUTEDEFINITIONDATEs[attributedefinitiondate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONDATE(attributedefinitiondate)
		}
	}
	return attributedefinitiondate
}

// for satisfaction of GongStruct interface
func (attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) GetName() (res string) {
	return attributedefinitiondate.Name
}

// Stage puts attributedefinitionenumeration to the model stage
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONENUMERATION {
	stage.ATTRIBUTEDEFINITIONENUMERATIONs[attributedefinitionenumeration] = __member
	stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString[attributedefinitionenumeration.Name] = attributedefinitionenumeration

	return attributedefinitionenumeration
}

// Unstage removes attributedefinitionenumeration off the model stage
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONENUMERATION {
	delete(stage.ATTRIBUTEDEFINITIONENUMERATIONs, attributedefinitionenumeration)
	delete(stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString, attributedefinitionenumeration.Name)
	return attributedefinitionenumeration
}

// UnstageVoid removes attributedefinitionenumeration off the model stage
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONENUMERATIONs, attributedefinitionenumeration)
	delete(stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString, attributedefinitionenumeration.Name)
}

// commit attributedefinitionenumeration to the back repo (if it is already staged)
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONENUMERATION {
	if _, ok := stage.ATTRIBUTEDEFINITIONENUMERATIONs[attributedefinitionenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration)
		}
	}
	return attributedefinitionenumeration
}

func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) CommitVoid(stage *StageStruct) {
	attributedefinitionenumeration.Commit(stage)
}

// Checkout attributedefinitionenumeration to the back repo (if it is already staged)
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONENUMERATION {
	if _, ok := stage.ATTRIBUTEDEFINITIONENUMERATIONs[attributedefinitionenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration)
		}
	}
	return attributedefinitionenumeration
}

// for satisfaction of GongStruct interface
func (attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) GetName() (res string) {
	return attributedefinitionenumeration.Name
}

// Stage puts attributedefinitioninteger to the model stage
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONINTEGER {
	stage.ATTRIBUTEDEFINITIONINTEGERs[attributedefinitioninteger] = __member
	stage.ATTRIBUTEDEFINITIONINTEGERs_mapString[attributedefinitioninteger.Name] = attributedefinitioninteger

	return attributedefinitioninteger
}

// Unstage removes attributedefinitioninteger off the model stage
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONINTEGER {
	delete(stage.ATTRIBUTEDEFINITIONINTEGERs, attributedefinitioninteger)
	delete(stage.ATTRIBUTEDEFINITIONINTEGERs_mapString, attributedefinitioninteger.Name)
	return attributedefinitioninteger
}

// UnstageVoid removes attributedefinitioninteger off the model stage
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONINTEGERs, attributedefinitioninteger)
	delete(stage.ATTRIBUTEDEFINITIONINTEGERs_mapString, attributedefinitioninteger.Name)
}

// commit attributedefinitioninteger to the back repo (if it is already staged)
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONINTEGER {
	if _, ok := stage.ATTRIBUTEDEFINITIONINTEGERs[attributedefinitioninteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger)
		}
	}
	return attributedefinitioninteger
}

func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) CommitVoid(stage *StageStruct) {
	attributedefinitioninteger.Commit(stage)
}

// Checkout attributedefinitioninteger to the back repo (if it is already staged)
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONINTEGER {
	if _, ok := stage.ATTRIBUTEDEFINITIONINTEGERs[attributedefinitioninteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger)
		}
	}
	return attributedefinitioninteger
}

// for satisfaction of GongStruct interface
func (attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) GetName() (res string) {
	return attributedefinitioninteger.Name
}

// Stage puts attributedefinitionreal to the model stage
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONREAL {
	stage.ATTRIBUTEDEFINITIONREALs[attributedefinitionreal] = __member
	stage.ATTRIBUTEDEFINITIONREALs_mapString[attributedefinitionreal.Name] = attributedefinitionreal

	return attributedefinitionreal
}

// Unstage removes attributedefinitionreal off the model stage
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONREAL {
	delete(stage.ATTRIBUTEDEFINITIONREALs, attributedefinitionreal)
	delete(stage.ATTRIBUTEDEFINITIONREALs_mapString, attributedefinitionreal.Name)
	return attributedefinitionreal
}

// UnstageVoid removes attributedefinitionreal off the model stage
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONREALs, attributedefinitionreal)
	delete(stage.ATTRIBUTEDEFINITIONREALs_mapString, attributedefinitionreal.Name)
}

// commit attributedefinitionreal to the back repo (if it is already staged)
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONREAL {
	if _, ok := stage.ATTRIBUTEDEFINITIONREALs[attributedefinitionreal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONREAL(attributedefinitionreal)
		}
	}
	return attributedefinitionreal
}

func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) CommitVoid(stage *StageStruct) {
	attributedefinitionreal.Commit(stage)
}

// Checkout attributedefinitionreal to the back repo (if it is already staged)
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONREAL {
	if _, ok := stage.ATTRIBUTEDEFINITIONREALs[attributedefinitionreal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONREAL(attributedefinitionreal)
		}
	}
	return attributedefinitionreal
}

// for satisfaction of GongStruct interface
func (attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) GetName() (res string) {
	return attributedefinitionreal.Name
}

// Stage puts attributedefinitionstring to the model stage
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONSTRING {
	stage.ATTRIBUTEDEFINITIONSTRINGs[attributedefinitionstring] = __member
	stage.ATTRIBUTEDEFINITIONSTRINGs_mapString[attributedefinitionstring.Name] = attributedefinitionstring

	return attributedefinitionstring
}

// Unstage removes attributedefinitionstring off the model stage
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONSTRING {
	delete(stage.ATTRIBUTEDEFINITIONSTRINGs, attributedefinitionstring)
	delete(stage.ATTRIBUTEDEFINITIONSTRINGs_mapString, attributedefinitionstring.Name)
	return attributedefinitionstring
}

// UnstageVoid removes attributedefinitionstring off the model stage
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONSTRINGs, attributedefinitionstring)
	delete(stage.ATTRIBUTEDEFINITIONSTRINGs_mapString, attributedefinitionstring.Name)
}

// commit attributedefinitionstring to the back repo (if it is already staged)
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONSTRING {
	if _, ok := stage.ATTRIBUTEDEFINITIONSTRINGs[attributedefinitionstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONSTRING(attributedefinitionstring)
		}
	}
	return attributedefinitionstring
}

func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) CommitVoid(stage *StageStruct) {
	attributedefinitionstring.Commit(stage)
}

// Checkout attributedefinitionstring to the back repo (if it is already staged)
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONSTRING {
	if _, ok := stage.ATTRIBUTEDEFINITIONSTRINGs[attributedefinitionstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONSTRING(attributedefinitionstring)
		}
	}
	return attributedefinitionstring
}

// for satisfaction of GongStruct interface
func (attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) GetName() (res string) {
	return attributedefinitionstring.Name
}

// Stage puts attributedefinitionxhtml to the model stage
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) Stage(stage *StageStruct) *ATTRIBUTEDEFINITIONXHTML {
	stage.ATTRIBUTEDEFINITIONXHTMLs[attributedefinitionxhtml] = __member
	stage.ATTRIBUTEDEFINITIONXHTMLs_mapString[attributedefinitionxhtml.Name] = attributedefinitionxhtml

	return attributedefinitionxhtml
}

// Unstage removes attributedefinitionxhtml off the model stage
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) Unstage(stage *StageStruct) *ATTRIBUTEDEFINITIONXHTML {
	delete(stage.ATTRIBUTEDEFINITIONXHTMLs, attributedefinitionxhtml)
	delete(stage.ATTRIBUTEDEFINITIONXHTMLs_mapString, attributedefinitionxhtml.Name)
	return attributedefinitionxhtml
}

// UnstageVoid removes attributedefinitionxhtml off the model stage
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEDEFINITIONXHTMLs, attributedefinitionxhtml)
	delete(stage.ATTRIBUTEDEFINITIONXHTMLs_mapString, attributedefinitionxhtml.Name)
}

// commit attributedefinitionxhtml to the back repo (if it is already staged)
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) Commit(stage *StageStruct) *ATTRIBUTEDEFINITIONXHTML {
	if _, ok := stage.ATTRIBUTEDEFINITIONXHTMLs[attributedefinitionxhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml)
		}
	}
	return attributedefinitionxhtml
}

func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) CommitVoid(stage *StageStruct) {
	attributedefinitionxhtml.Commit(stage)
}

// Checkout attributedefinitionxhtml to the back repo (if it is already staged)
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) Checkout(stage *StageStruct) *ATTRIBUTEDEFINITIONXHTML {
	if _, ok := stage.ATTRIBUTEDEFINITIONXHTMLs[attributedefinitionxhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml)
		}
	}
	return attributedefinitionxhtml
}

// for satisfaction of GongStruct interface
func (attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) GetName() (res string) {
	return attributedefinitionxhtml.Name
}

// Stage puts attributevalueboolean to the model stage
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) Stage(stage *StageStruct) *ATTRIBUTEVALUEBOOLEAN {
	stage.ATTRIBUTEVALUEBOOLEANs[attributevalueboolean] = __member
	stage.ATTRIBUTEVALUEBOOLEANs_mapString[attributevalueboolean.Name] = attributevalueboolean

	return attributevalueboolean
}

// Unstage removes attributevalueboolean off the model stage
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) Unstage(stage *StageStruct) *ATTRIBUTEVALUEBOOLEAN {
	delete(stage.ATTRIBUTEVALUEBOOLEANs, attributevalueboolean)
	delete(stage.ATTRIBUTEVALUEBOOLEANs_mapString, attributevalueboolean.Name)
	return attributevalueboolean
}

// UnstageVoid removes attributevalueboolean off the model stage
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEBOOLEANs, attributevalueboolean)
	delete(stage.ATTRIBUTEVALUEBOOLEANs_mapString, attributevalueboolean.Name)
}

// commit attributevalueboolean to the back repo (if it is already staged)
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) Commit(stage *StageStruct) *ATTRIBUTEVALUEBOOLEAN {
	if _, ok := stage.ATTRIBUTEVALUEBOOLEANs[attributevalueboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEBOOLEAN(attributevalueboolean)
		}
	}
	return attributevalueboolean
}

func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) CommitVoid(stage *StageStruct) {
	attributevalueboolean.Commit(stage)
}

// Checkout attributevalueboolean to the back repo (if it is already staged)
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) Checkout(stage *StageStruct) *ATTRIBUTEVALUEBOOLEAN {
	if _, ok := stage.ATTRIBUTEVALUEBOOLEANs[attributevalueboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEBOOLEAN(attributevalueboolean)
		}
	}
	return attributevalueboolean
}

// for satisfaction of GongStruct interface
func (attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) GetName() (res string) {
	return attributevalueboolean.Name
}

// Stage puts attributevaluedate to the model stage
func (attributevaluedate *ATTRIBUTEVALUEDATE) Stage(stage *StageStruct) *ATTRIBUTEVALUEDATE {
	stage.ATTRIBUTEVALUEDATEs[attributevaluedate] = __member
	stage.ATTRIBUTEVALUEDATEs_mapString[attributevaluedate.Name] = attributevaluedate

	return attributevaluedate
}

// Unstage removes attributevaluedate off the model stage
func (attributevaluedate *ATTRIBUTEVALUEDATE) Unstage(stage *StageStruct) *ATTRIBUTEVALUEDATE {
	delete(stage.ATTRIBUTEVALUEDATEs, attributevaluedate)
	delete(stage.ATTRIBUTEVALUEDATEs_mapString, attributevaluedate.Name)
	return attributevaluedate
}

// UnstageVoid removes attributevaluedate off the model stage
func (attributevaluedate *ATTRIBUTEVALUEDATE) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEDATEs, attributevaluedate)
	delete(stage.ATTRIBUTEVALUEDATEs_mapString, attributevaluedate.Name)
}

// commit attributevaluedate to the back repo (if it is already staged)
func (attributevaluedate *ATTRIBUTEVALUEDATE) Commit(stage *StageStruct) *ATTRIBUTEVALUEDATE {
	if _, ok := stage.ATTRIBUTEVALUEDATEs[attributevaluedate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEDATE(attributevaluedate)
		}
	}
	return attributevaluedate
}

func (attributevaluedate *ATTRIBUTEVALUEDATE) CommitVoid(stage *StageStruct) {
	attributevaluedate.Commit(stage)
}

// Checkout attributevaluedate to the back repo (if it is already staged)
func (attributevaluedate *ATTRIBUTEVALUEDATE) Checkout(stage *StageStruct) *ATTRIBUTEVALUEDATE {
	if _, ok := stage.ATTRIBUTEVALUEDATEs[attributevaluedate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEDATE(attributevaluedate)
		}
	}
	return attributevaluedate
}

// for satisfaction of GongStruct interface
func (attributevaluedate *ATTRIBUTEVALUEDATE) GetName() (res string) {
	return attributevaluedate.Name
}

// Stage puts attributevalueenumeration to the model stage
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) Stage(stage *StageStruct) *ATTRIBUTEVALUEENUMERATION {
	stage.ATTRIBUTEVALUEENUMERATIONs[attributevalueenumeration] = __member
	stage.ATTRIBUTEVALUEENUMERATIONs_mapString[attributevalueenumeration.Name] = attributevalueenumeration

	return attributevalueenumeration
}

// Unstage removes attributevalueenumeration off the model stage
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) Unstage(stage *StageStruct) *ATTRIBUTEVALUEENUMERATION {
	delete(stage.ATTRIBUTEVALUEENUMERATIONs, attributevalueenumeration)
	delete(stage.ATTRIBUTEVALUEENUMERATIONs_mapString, attributevalueenumeration.Name)
	return attributevalueenumeration
}

// UnstageVoid removes attributevalueenumeration off the model stage
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEENUMERATIONs, attributevalueenumeration)
	delete(stage.ATTRIBUTEVALUEENUMERATIONs_mapString, attributevalueenumeration.Name)
}

// commit attributevalueenumeration to the back repo (if it is already staged)
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) Commit(stage *StageStruct) *ATTRIBUTEVALUEENUMERATION {
	if _, ok := stage.ATTRIBUTEVALUEENUMERATIONs[attributevalueenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEENUMERATION(attributevalueenumeration)
		}
	}
	return attributevalueenumeration
}

func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) CommitVoid(stage *StageStruct) {
	attributevalueenumeration.Commit(stage)
}

// Checkout attributevalueenumeration to the back repo (if it is already staged)
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) Checkout(stage *StageStruct) *ATTRIBUTEVALUEENUMERATION {
	if _, ok := stage.ATTRIBUTEVALUEENUMERATIONs[attributevalueenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEENUMERATION(attributevalueenumeration)
		}
	}
	return attributevalueenumeration
}

// for satisfaction of GongStruct interface
func (attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) GetName() (res string) {
	return attributevalueenumeration.Name
}

// Stage puts attributevalueinteger to the model stage
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) Stage(stage *StageStruct) *ATTRIBUTEVALUEINTEGER {
	stage.ATTRIBUTEVALUEINTEGERs[attributevalueinteger] = __member
	stage.ATTRIBUTEVALUEINTEGERs_mapString[attributevalueinteger.Name] = attributevalueinteger

	return attributevalueinteger
}

// Unstage removes attributevalueinteger off the model stage
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) Unstage(stage *StageStruct) *ATTRIBUTEVALUEINTEGER {
	delete(stage.ATTRIBUTEVALUEINTEGERs, attributevalueinteger)
	delete(stage.ATTRIBUTEVALUEINTEGERs_mapString, attributevalueinteger.Name)
	return attributevalueinteger
}

// UnstageVoid removes attributevalueinteger off the model stage
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEINTEGERs, attributevalueinteger)
	delete(stage.ATTRIBUTEVALUEINTEGERs_mapString, attributevalueinteger.Name)
}

// commit attributevalueinteger to the back repo (if it is already staged)
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) Commit(stage *StageStruct) *ATTRIBUTEVALUEINTEGER {
	if _, ok := stage.ATTRIBUTEVALUEINTEGERs[attributevalueinteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEINTEGER(attributevalueinteger)
		}
	}
	return attributevalueinteger
}

func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) CommitVoid(stage *StageStruct) {
	attributevalueinteger.Commit(stage)
}

// Checkout attributevalueinteger to the back repo (if it is already staged)
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) Checkout(stage *StageStruct) *ATTRIBUTEVALUEINTEGER {
	if _, ok := stage.ATTRIBUTEVALUEINTEGERs[attributevalueinteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEINTEGER(attributevalueinteger)
		}
	}
	return attributevalueinteger
}

// for satisfaction of GongStruct interface
func (attributevalueinteger *ATTRIBUTEVALUEINTEGER) GetName() (res string) {
	return attributevalueinteger.Name
}

// Stage puts attributevaluereal to the model stage
func (attributevaluereal *ATTRIBUTEVALUEREAL) Stage(stage *StageStruct) *ATTRIBUTEVALUEREAL {
	stage.ATTRIBUTEVALUEREALs[attributevaluereal] = __member
	stage.ATTRIBUTEVALUEREALs_mapString[attributevaluereal.Name] = attributevaluereal

	return attributevaluereal
}

// Unstage removes attributevaluereal off the model stage
func (attributevaluereal *ATTRIBUTEVALUEREAL) Unstage(stage *StageStruct) *ATTRIBUTEVALUEREAL {
	delete(stage.ATTRIBUTEVALUEREALs, attributevaluereal)
	delete(stage.ATTRIBUTEVALUEREALs_mapString, attributevaluereal.Name)
	return attributevaluereal
}

// UnstageVoid removes attributevaluereal off the model stage
func (attributevaluereal *ATTRIBUTEVALUEREAL) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEREALs, attributevaluereal)
	delete(stage.ATTRIBUTEVALUEREALs_mapString, attributevaluereal.Name)
}

// commit attributevaluereal to the back repo (if it is already staged)
func (attributevaluereal *ATTRIBUTEVALUEREAL) Commit(stage *StageStruct) *ATTRIBUTEVALUEREAL {
	if _, ok := stage.ATTRIBUTEVALUEREALs[attributevaluereal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEREAL(attributevaluereal)
		}
	}
	return attributevaluereal
}

func (attributevaluereal *ATTRIBUTEVALUEREAL) CommitVoid(stage *StageStruct) {
	attributevaluereal.Commit(stage)
}

// Checkout attributevaluereal to the back repo (if it is already staged)
func (attributevaluereal *ATTRIBUTEVALUEREAL) Checkout(stage *StageStruct) *ATTRIBUTEVALUEREAL {
	if _, ok := stage.ATTRIBUTEVALUEREALs[attributevaluereal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEREAL(attributevaluereal)
		}
	}
	return attributevaluereal
}

// for satisfaction of GongStruct interface
func (attributevaluereal *ATTRIBUTEVALUEREAL) GetName() (res string) {
	return attributevaluereal.Name
}

// Stage puts attributevaluestring to the model stage
func (attributevaluestring *ATTRIBUTEVALUESTRING) Stage(stage *StageStruct) *ATTRIBUTEVALUESTRING {
	stage.ATTRIBUTEVALUESTRINGs[attributevaluestring] = __member
	stage.ATTRIBUTEVALUESTRINGs_mapString[attributevaluestring.Name] = attributevaluestring

	return attributevaluestring
}

// Unstage removes attributevaluestring off the model stage
func (attributevaluestring *ATTRIBUTEVALUESTRING) Unstage(stage *StageStruct) *ATTRIBUTEVALUESTRING {
	delete(stage.ATTRIBUTEVALUESTRINGs, attributevaluestring)
	delete(stage.ATTRIBUTEVALUESTRINGs_mapString, attributevaluestring.Name)
	return attributevaluestring
}

// UnstageVoid removes attributevaluestring off the model stage
func (attributevaluestring *ATTRIBUTEVALUESTRING) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUESTRINGs, attributevaluestring)
	delete(stage.ATTRIBUTEVALUESTRINGs_mapString, attributevaluestring.Name)
}

// commit attributevaluestring to the back repo (if it is already staged)
func (attributevaluestring *ATTRIBUTEVALUESTRING) Commit(stage *StageStruct) *ATTRIBUTEVALUESTRING {
	if _, ok := stage.ATTRIBUTEVALUESTRINGs[attributevaluestring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUESTRING(attributevaluestring)
		}
	}
	return attributevaluestring
}

func (attributevaluestring *ATTRIBUTEVALUESTRING) CommitVoid(stage *StageStruct) {
	attributevaluestring.Commit(stage)
}

// Checkout attributevaluestring to the back repo (if it is already staged)
func (attributevaluestring *ATTRIBUTEVALUESTRING) Checkout(stage *StageStruct) *ATTRIBUTEVALUESTRING {
	if _, ok := stage.ATTRIBUTEVALUESTRINGs[attributevaluestring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUESTRING(attributevaluestring)
		}
	}
	return attributevaluestring
}

// for satisfaction of GongStruct interface
func (attributevaluestring *ATTRIBUTEVALUESTRING) GetName() (res string) {
	return attributevaluestring.Name
}

// Stage puts attributevaluexhtml to the model stage
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) Stage(stage *StageStruct) *ATTRIBUTEVALUEXHTML {
	stage.ATTRIBUTEVALUEXHTMLs[attributevaluexhtml] = __member
	stage.ATTRIBUTEVALUEXHTMLs_mapString[attributevaluexhtml.Name] = attributevaluexhtml

	return attributevaluexhtml
}

// Unstage removes attributevaluexhtml off the model stage
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) Unstage(stage *StageStruct) *ATTRIBUTEVALUEXHTML {
	delete(stage.ATTRIBUTEVALUEXHTMLs, attributevaluexhtml)
	delete(stage.ATTRIBUTEVALUEXHTMLs_mapString, attributevaluexhtml.Name)
	return attributevaluexhtml
}

// UnstageVoid removes attributevaluexhtml off the model stage
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTEVALUEXHTMLs, attributevaluexhtml)
	delete(stage.ATTRIBUTEVALUEXHTMLs_mapString, attributevaluexhtml.Name)
}

// commit attributevaluexhtml to the back repo (if it is already staged)
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) Commit(stage *StageStruct) *ATTRIBUTEVALUEXHTML {
	if _, ok := stage.ATTRIBUTEVALUEXHTMLs[attributevaluexhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTEVALUEXHTML(attributevaluexhtml)
		}
	}
	return attributevaluexhtml
}

func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) CommitVoid(stage *StageStruct) {
	attributevaluexhtml.Commit(stage)
}

// Checkout attributevaluexhtml to the back repo (if it is already staged)
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) Checkout(stage *StageStruct) *ATTRIBUTEVALUEXHTML {
	if _, ok := stage.ATTRIBUTEVALUEXHTMLs[attributevaluexhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTEVALUEXHTML(attributevaluexhtml)
		}
	}
	return attributevaluexhtml
}

// for satisfaction of GongStruct interface
func (attributevaluexhtml *ATTRIBUTEVALUEXHTML) GetName() (res string) {
	return attributevaluexhtml.Name
}

// Stage puts children to the model stage
func (children *CHILDREN) Stage(stage *StageStruct) *CHILDREN {
	stage.CHILDRENs[children] = __member
	stage.CHILDRENs_mapString[children.Name] = children

	return children
}

// Unstage removes children off the model stage
func (children *CHILDREN) Unstage(stage *StageStruct) *CHILDREN {
	delete(stage.CHILDRENs, children)
	delete(stage.CHILDRENs_mapString, children.Name)
	return children
}

// UnstageVoid removes children off the model stage
func (children *CHILDREN) UnstageVoid(stage *StageStruct) {
	delete(stage.CHILDRENs, children)
	delete(stage.CHILDRENs_mapString, children.Name)
}

// commit children to the back repo (if it is already staged)
func (children *CHILDREN) Commit(stage *StageStruct) *CHILDREN {
	if _, ok := stage.CHILDRENs[children]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCHILDREN(children)
		}
	}
	return children
}

func (children *CHILDREN) CommitVoid(stage *StageStruct) {
	children.Commit(stage)
}

// Checkout children to the back repo (if it is already staged)
func (children *CHILDREN) Checkout(stage *StageStruct) *CHILDREN {
	if _, ok := stage.CHILDRENs[children]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCHILDREN(children)
		}
	}
	return children
}

// for satisfaction of GongStruct interface
func (children *CHILDREN) GetName() (res string) {
	return children.Name
}

// Stage puts corecontent to the model stage
func (corecontent *CORECONTENT) Stage(stage *StageStruct) *CORECONTENT {
	stage.CORECONTENTs[corecontent] = __member
	stage.CORECONTENTs_mapString[corecontent.Name] = corecontent

	return corecontent
}

// Unstage removes corecontent off the model stage
func (corecontent *CORECONTENT) Unstage(stage *StageStruct) *CORECONTENT {
	delete(stage.CORECONTENTs, corecontent)
	delete(stage.CORECONTENTs_mapString, corecontent.Name)
	return corecontent
}

// UnstageVoid removes corecontent off the model stage
func (corecontent *CORECONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.CORECONTENTs, corecontent)
	delete(stage.CORECONTENTs_mapString, corecontent.Name)
}

// commit corecontent to the back repo (if it is already staged)
func (corecontent *CORECONTENT) Commit(stage *StageStruct) *CORECONTENT {
	if _, ok := stage.CORECONTENTs[corecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCORECONTENT(corecontent)
		}
	}
	return corecontent
}

func (corecontent *CORECONTENT) CommitVoid(stage *StageStruct) {
	corecontent.Commit(stage)
}

// Checkout corecontent to the back repo (if it is already staged)
func (corecontent *CORECONTENT) Checkout(stage *StageStruct) *CORECONTENT {
	if _, ok := stage.CORECONTENTs[corecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCORECONTENT(corecontent)
		}
	}
	return corecontent
}

// for satisfaction of GongStruct interface
func (corecontent *CORECONTENT) GetName() (res string) {
	return corecontent.Name
}

// Stage puts datatypedefinitionboolean to the model stage
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) Stage(stage *StageStruct) *DATATYPEDEFINITIONBOOLEAN {
	stage.DATATYPEDEFINITIONBOOLEANs[datatypedefinitionboolean] = __member
	stage.DATATYPEDEFINITIONBOOLEANs_mapString[datatypedefinitionboolean.Name] = datatypedefinitionboolean

	return datatypedefinitionboolean
}

// Unstage removes datatypedefinitionboolean off the model stage
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) Unstage(stage *StageStruct) *DATATYPEDEFINITIONBOOLEAN {
	delete(stage.DATATYPEDEFINITIONBOOLEANs, datatypedefinitionboolean)
	delete(stage.DATATYPEDEFINITIONBOOLEANs_mapString, datatypedefinitionboolean.Name)
	return datatypedefinitionboolean
}

// UnstageVoid removes datatypedefinitionboolean off the model stage
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONBOOLEANs, datatypedefinitionboolean)
	delete(stage.DATATYPEDEFINITIONBOOLEANs_mapString, datatypedefinitionboolean.Name)
}

// commit datatypedefinitionboolean to the back repo (if it is already staged)
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) Commit(stage *StageStruct) *DATATYPEDEFINITIONBOOLEAN {
	if _, ok := stage.DATATYPEDEFINITIONBOOLEANs[datatypedefinitionboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean)
		}
	}
	return datatypedefinitionboolean
}

func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) CommitVoid(stage *StageStruct) {
	datatypedefinitionboolean.Commit(stage)
}

// Checkout datatypedefinitionboolean to the back repo (if it is already staged)
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) Checkout(stage *StageStruct) *DATATYPEDEFINITIONBOOLEAN {
	if _, ok := stage.DATATYPEDEFINITIONBOOLEANs[datatypedefinitionboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean)
		}
	}
	return datatypedefinitionboolean
}

// for satisfaction of GongStruct interface
func (datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) GetName() (res string) {
	return datatypedefinitionboolean.Name
}

// Stage puts datatypedefinitiondate to the model stage
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) Stage(stage *StageStruct) *DATATYPEDEFINITIONDATE {
	stage.DATATYPEDEFINITIONDATEs[datatypedefinitiondate] = __member
	stage.DATATYPEDEFINITIONDATEs_mapString[datatypedefinitiondate.Name] = datatypedefinitiondate

	return datatypedefinitiondate
}

// Unstage removes datatypedefinitiondate off the model stage
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) Unstage(stage *StageStruct) *DATATYPEDEFINITIONDATE {
	delete(stage.DATATYPEDEFINITIONDATEs, datatypedefinitiondate)
	delete(stage.DATATYPEDEFINITIONDATEs_mapString, datatypedefinitiondate.Name)
	return datatypedefinitiondate
}

// UnstageVoid removes datatypedefinitiondate off the model stage
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONDATEs, datatypedefinitiondate)
	delete(stage.DATATYPEDEFINITIONDATEs_mapString, datatypedefinitiondate.Name)
}

// commit datatypedefinitiondate to the back repo (if it is already staged)
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) Commit(stage *StageStruct) *DATATYPEDEFINITIONDATE {
	if _, ok := stage.DATATYPEDEFINITIONDATEs[datatypedefinitiondate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONDATE(datatypedefinitiondate)
		}
	}
	return datatypedefinitiondate
}

func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) CommitVoid(stage *StageStruct) {
	datatypedefinitiondate.Commit(stage)
}

// Checkout datatypedefinitiondate to the back repo (if it is already staged)
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) Checkout(stage *StageStruct) *DATATYPEDEFINITIONDATE {
	if _, ok := stage.DATATYPEDEFINITIONDATEs[datatypedefinitiondate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONDATE(datatypedefinitiondate)
		}
	}
	return datatypedefinitiondate
}

// for satisfaction of GongStruct interface
func (datatypedefinitiondate *DATATYPEDEFINITIONDATE) GetName() (res string) {
	return datatypedefinitiondate.Name
}

// Stage puts datatypedefinitionenumeration to the model stage
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) Stage(stage *StageStruct) *DATATYPEDEFINITIONENUMERATION {
	stage.DATATYPEDEFINITIONENUMERATIONs[datatypedefinitionenumeration] = __member
	stage.DATATYPEDEFINITIONENUMERATIONs_mapString[datatypedefinitionenumeration.Name] = datatypedefinitionenumeration

	return datatypedefinitionenumeration
}

// Unstage removes datatypedefinitionenumeration off the model stage
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) Unstage(stage *StageStruct) *DATATYPEDEFINITIONENUMERATION {
	delete(stage.DATATYPEDEFINITIONENUMERATIONs, datatypedefinitionenumeration)
	delete(stage.DATATYPEDEFINITIONENUMERATIONs_mapString, datatypedefinitionenumeration.Name)
	return datatypedefinitionenumeration
}

// UnstageVoid removes datatypedefinitionenumeration off the model stage
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONENUMERATIONs, datatypedefinitionenumeration)
	delete(stage.DATATYPEDEFINITIONENUMERATIONs_mapString, datatypedefinitionenumeration.Name)
}

// commit datatypedefinitionenumeration to the back repo (if it is already staged)
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) Commit(stage *StageStruct) *DATATYPEDEFINITIONENUMERATION {
	if _, ok := stage.DATATYPEDEFINITIONENUMERATIONs[datatypedefinitionenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration)
		}
	}
	return datatypedefinitionenumeration
}

func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) CommitVoid(stage *StageStruct) {
	datatypedefinitionenumeration.Commit(stage)
}

// Checkout datatypedefinitionenumeration to the back repo (if it is already staged)
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) Checkout(stage *StageStruct) *DATATYPEDEFINITIONENUMERATION {
	if _, ok := stage.DATATYPEDEFINITIONENUMERATIONs[datatypedefinitionenumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration)
		}
	}
	return datatypedefinitionenumeration
}

// for satisfaction of GongStruct interface
func (datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) GetName() (res string) {
	return datatypedefinitionenumeration.Name
}

// Stage puts datatypedefinitioninteger to the model stage
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) Stage(stage *StageStruct) *DATATYPEDEFINITIONINTEGER {
	stage.DATATYPEDEFINITIONINTEGERs[datatypedefinitioninteger] = __member
	stage.DATATYPEDEFINITIONINTEGERs_mapString[datatypedefinitioninteger.Name] = datatypedefinitioninteger

	return datatypedefinitioninteger
}

// Unstage removes datatypedefinitioninteger off the model stage
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) Unstage(stage *StageStruct) *DATATYPEDEFINITIONINTEGER {
	delete(stage.DATATYPEDEFINITIONINTEGERs, datatypedefinitioninteger)
	delete(stage.DATATYPEDEFINITIONINTEGERs_mapString, datatypedefinitioninteger.Name)
	return datatypedefinitioninteger
}

// UnstageVoid removes datatypedefinitioninteger off the model stage
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONINTEGERs, datatypedefinitioninteger)
	delete(stage.DATATYPEDEFINITIONINTEGERs_mapString, datatypedefinitioninteger.Name)
}

// commit datatypedefinitioninteger to the back repo (if it is already staged)
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) Commit(stage *StageStruct) *DATATYPEDEFINITIONINTEGER {
	if _, ok := stage.DATATYPEDEFINITIONINTEGERs[datatypedefinitioninteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger)
		}
	}
	return datatypedefinitioninteger
}

func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) CommitVoid(stage *StageStruct) {
	datatypedefinitioninteger.Commit(stage)
}

// Checkout datatypedefinitioninteger to the back repo (if it is already staged)
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) Checkout(stage *StageStruct) *DATATYPEDEFINITIONINTEGER {
	if _, ok := stage.DATATYPEDEFINITIONINTEGERs[datatypedefinitioninteger]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger)
		}
	}
	return datatypedefinitioninteger
}

// for satisfaction of GongStruct interface
func (datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) GetName() (res string) {
	return datatypedefinitioninteger.Name
}

// Stage puts datatypedefinitionreal to the model stage
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) Stage(stage *StageStruct) *DATATYPEDEFINITIONREAL {
	stage.DATATYPEDEFINITIONREALs[datatypedefinitionreal] = __member
	stage.DATATYPEDEFINITIONREALs_mapString[datatypedefinitionreal.Name] = datatypedefinitionreal

	return datatypedefinitionreal
}

// Unstage removes datatypedefinitionreal off the model stage
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) Unstage(stage *StageStruct) *DATATYPEDEFINITIONREAL {
	delete(stage.DATATYPEDEFINITIONREALs, datatypedefinitionreal)
	delete(stage.DATATYPEDEFINITIONREALs_mapString, datatypedefinitionreal.Name)
	return datatypedefinitionreal
}

// UnstageVoid removes datatypedefinitionreal off the model stage
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONREALs, datatypedefinitionreal)
	delete(stage.DATATYPEDEFINITIONREALs_mapString, datatypedefinitionreal.Name)
}

// commit datatypedefinitionreal to the back repo (if it is already staged)
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) Commit(stage *StageStruct) *DATATYPEDEFINITIONREAL {
	if _, ok := stage.DATATYPEDEFINITIONREALs[datatypedefinitionreal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONREAL(datatypedefinitionreal)
		}
	}
	return datatypedefinitionreal
}

func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) CommitVoid(stage *StageStruct) {
	datatypedefinitionreal.Commit(stage)
}

// Checkout datatypedefinitionreal to the back repo (if it is already staged)
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) Checkout(stage *StageStruct) *DATATYPEDEFINITIONREAL {
	if _, ok := stage.DATATYPEDEFINITIONREALs[datatypedefinitionreal]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONREAL(datatypedefinitionreal)
		}
	}
	return datatypedefinitionreal
}

// for satisfaction of GongStruct interface
func (datatypedefinitionreal *DATATYPEDEFINITIONREAL) GetName() (res string) {
	return datatypedefinitionreal.Name
}

// Stage puts datatypedefinitionstring to the model stage
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) Stage(stage *StageStruct) *DATATYPEDEFINITIONSTRING {
	stage.DATATYPEDEFINITIONSTRINGs[datatypedefinitionstring] = __member
	stage.DATATYPEDEFINITIONSTRINGs_mapString[datatypedefinitionstring.Name] = datatypedefinitionstring

	return datatypedefinitionstring
}

// Unstage removes datatypedefinitionstring off the model stage
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) Unstage(stage *StageStruct) *DATATYPEDEFINITIONSTRING {
	delete(stage.DATATYPEDEFINITIONSTRINGs, datatypedefinitionstring)
	delete(stage.DATATYPEDEFINITIONSTRINGs_mapString, datatypedefinitionstring.Name)
	return datatypedefinitionstring
}

// UnstageVoid removes datatypedefinitionstring off the model stage
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONSTRINGs, datatypedefinitionstring)
	delete(stage.DATATYPEDEFINITIONSTRINGs_mapString, datatypedefinitionstring.Name)
}

// commit datatypedefinitionstring to the back repo (if it is already staged)
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) Commit(stage *StageStruct) *DATATYPEDEFINITIONSTRING {
	if _, ok := stage.DATATYPEDEFINITIONSTRINGs[datatypedefinitionstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONSTRING(datatypedefinitionstring)
		}
	}
	return datatypedefinitionstring
}

func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) CommitVoid(stage *StageStruct) {
	datatypedefinitionstring.Commit(stage)
}

// Checkout datatypedefinitionstring to the back repo (if it is already staged)
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) Checkout(stage *StageStruct) *DATATYPEDEFINITIONSTRING {
	if _, ok := stage.DATATYPEDEFINITIONSTRINGs[datatypedefinitionstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONSTRING(datatypedefinitionstring)
		}
	}
	return datatypedefinitionstring
}

// for satisfaction of GongStruct interface
func (datatypedefinitionstring *DATATYPEDEFINITIONSTRING) GetName() (res string) {
	return datatypedefinitionstring.Name
}

// Stage puts datatypedefinitionxhtml to the model stage
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) Stage(stage *StageStruct) *DATATYPEDEFINITIONXHTML {
	stage.DATATYPEDEFINITIONXHTMLs[datatypedefinitionxhtml] = __member
	stage.DATATYPEDEFINITIONXHTMLs_mapString[datatypedefinitionxhtml.Name] = datatypedefinitionxhtml

	return datatypedefinitionxhtml
}

// Unstage removes datatypedefinitionxhtml off the model stage
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) Unstage(stage *StageStruct) *DATATYPEDEFINITIONXHTML {
	delete(stage.DATATYPEDEFINITIONXHTMLs, datatypedefinitionxhtml)
	delete(stage.DATATYPEDEFINITIONXHTMLs_mapString, datatypedefinitionxhtml.Name)
	return datatypedefinitionxhtml
}

// UnstageVoid removes datatypedefinitionxhtml off the model stage
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPEDEFINITIONXHTMLs, datatypedefinitionxhtml)
	delete(stage.DATATYPEDEFINITIONXHTMLs_mapString, datatypedefinitionxhtml.Name)
}

// commit datatypedefinitionxhtml to the back repo (if it is already staged)
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) Commit(stage *StageStruct) *DATATYPEDEFINITIONXHTML {
	if _, ok := stage.DATATYPEDEFINITIONXHTMLs[datatypedefinitionxhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml)
		}
	}
	return datatypedefinitionxhtml
}

func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) CommitVoid(stage *StageStruct) {
	datatypedefinitionxhtml.Commit(stage)
}

// Checkout datatypedefinitionxhtml to the back repo (if it is already staged)
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) Checkout(stage *StageStruct) *DATATYPEDEFINITIONXHTML {
	if _, ok := stage.DATATYPEDEFINITIONXHTMLs[datatypedefinitionxhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml)
		}
	}
	return datatypedefinitionxhtml
}

// for satisfaction of GongStruct interface
func (datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) GetName() (res string) {
	return datatypedefinitionxhtml.Name
}

// Stage puts datatypes to the model stage
func (datatypes *DATATYPES) Stage(stage *StageStruct) *DATATYPES {
	stage.DATATYPESs[datatypes] = __member
	stage.DATATYPESs_mapString[datatypes.Name] = datatypes

	return datatypes
}

// Unstage removes datatypes off the model stage
func (datatypes *DATATYPES) Unstage(stage *StageStruct) *DATATYPES {
	delete(stage.DATATYPESs, datatypes)
	delete(stage.DATATYPESs_mapString, datatypes.Name)
	return datatypes
}

// UnstageVoid removes datatypes off the model stage
func (datatypes *DATATYPES) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPESs, datatypes)
	delete(stage.DATATYPESs_mapString, datatypes.Name)
}

// commit datatypes to the back repo (if it is already staged)
func (datatypes *DATATYPES) Commit(stage *StageStruct) *DATATYPES {
	if _, ok := stage.DATATYPESs[datatypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPES(datatypes)
		}
	}
	return datatypes
}

func (datatypes *DATATYPES) CommitVoid(stage *StageStruct) {
	datatypes.Commit(stage)
}

// Checkout datatypes to the back repo (if it is already staged)
func (datatypes *DATATYPES) Checkout(stage *StageStruct) *DATATYPES {
	if _, ok := stage.DATATYPESs[datatypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPES(datatypes)
		}
	}
	return datatypes
}

// for satisfaction of GongStruct interface
func (datatypes *DATATYPES) GetName() (res string) {
	return datatypes.Name
}

// Stage puts defaultvalue to the model stage
func (defaultvalue *DEFAULTVALUE) Stage(stage *StageStruct) *DEFAULTVALUE {
	stage.DEFAULTVALUEs[defaultvalue] = __member
	stage.DEFAULTVALUEs_mapString[defaultvalue.Name] = defaultvalue

	return defaultvalue
}

// Unstage removes defaultvalue off the model stage
func (defaultvalue *DEFAULTVALUE) Unstage(stage *StageStruct) *DEFAULTVALUE {
	delete(stage.DEFAULTVALUEs, defaultvalue)
	delete(stage.DEFAULTVALUEs_mapString, defaultvalue.Name)
	return defaultvalue
}

// UnstageVoid removes defaultvalue off the model stage
func (defaultvalue *DEFAULTVALUE) UnstageVoid(stage *StageStruct) {
	delete(stage.DEFAULTVALUEs, defaultvalue)
	delete(stage.DEFAULTVALUEs_mapString, defaultvalue.Name)
}

// commit defaultvalue to the back repo (if it is already staged)
func (defaultvalue *DEFAULTVALUE) Commit(stage *StageStruct) *DEFAULTVALUE {
	if _, ok := stage.DEFAULTVALUEs[defaultvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDEFAULTVALUE(defaultvalue)
		}
	}
	return defaultvalue
}

func (defaultvalue *DEFAULTVALUE) CommitVoid(stage *StageStruct) {
	defaultvalue.Commit(stage)
}

// Checkout defaultvalue to the back repo (if it is already staged)
func (defaultvalue *DEFAULTVALUE) Checkout(stage *StageStruct) *DEFAULTVALUE {
	if _, ok := stage.DEFAULTVALUEs[defaultvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDEFAULTVALUE(defaultvalue)
		}
	}
	return defaultvalue
}

// for satisfaction of GongStruct interface
func (defaultvalue *DEFAULTVALUE) GetName() (res string) {
	return defaultvalue.Name
}

// Stage puts definition to the model stage
func (definition *DEFINITION) Stage(stage *StageStruct) *DEFINITION {
	stage.DEFINITIONs[definition] = __member
	stage.DEFINITIONs_mapString[definition.Name] = definition

	return definition
}

// Unstage removes definition off the model stage
func (definition *DEFINITION) Unstage(stage *StageStruct) *DEFINITION {
	delete(stage.DEFINITIONs, definition)
	delete(stage.DEFINITIONs_mapString, definition.Name)
	return definition
}

// UnstageVoid removes definition off the model stage
func (definition *DEFINITION) UnstageVoid(stage *StageStruct) {
	delete(stage.DEFINITIONs, definition)
	delete(stage.DEFINITIONs_mapString, definition.Name)
}

// commit definition to the back repo (if it is already staged)
func (definition *DEFINITION) Commit(stage *StageStruct) *DEFINITION {
	if _, ok := stage.DEFINITIONs[definition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDEFINITION(definition)
		}
	}
	return definition
}

func (definition *DEFINITION) CommitVoid(stage *StageStruct) {
	definition.Commit(stage)
}

// Checkout definition to the back repo (if it is already staged)
func (definition *DEFINITION) Checkout(stage *StageStruct) *DEFINITION {
	if _, ok := stage.DEFINITIONs[definition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDEFINITION(definition)
		}
	}
	return definition
}

// for satisfaction of GongStruct interface
func (definition *DEFINITION) GetName() (res string) {
	return definition.Name
}

// Stage puts editableatts to the model stage
func (editableatts *EDITABLEATTS) Stage(stage *StageStruct) *EDITABLEATTS {
	stage.EDITABLEATTSs[editableatts] = __member
	stage.EDITABLEATTSs_mapString[editableatts.Name] = editableatts

	return editableatts
}

// Unstage removes editableatts off the model stage
func (editableatts *EDITABLEATTS) Unstage(stage *StageStruct) *EDITABLEATTS {
	delete(stage.EDITABLEATTSs, editableatts)
	delete(stage.EDITABLEATTSs_mapString, editableatts.Name)
	return editableatts
}

// UnstageVoid removes editableatts off the model stage
func (editableatts *EDITABLEATTS) UnstageVoid(stage *StageStruct) {
	delete(stage.EDITABLEATTSs, editableatts)
	delete(stage.EDITABLEATTSs_mapString, editableatts.Name)
}

// commit editableatts to the back repo (if it is already staged)
func (editableatts *EDITABLEATTS) Commit(stage *StageStruct) *EDITABLEATTS {
	if _, ok := stage.EDITABLEATTSs[editableatts]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEDITABLEATTS(editableatts)
		}
	}
	return editableatts
}

func (editableatts *EDITABLEATTS) CommitVoid(stage *StageStruct) {
	editableatts.Commit(stage)
}

// Checkout editableatts to the back repo (if it is already staged)
func (editableatts *EDITABLEATTS) Checkout(stage *StageStruct) *EDITABLEATTS {
	if _, ok := stage.EDITABLEATTSs[editableatts]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEDITABLEATTS(editableatts)
		}
	}
	return editableatts
}

// for satisfaction of GongStruct interface
func (editableatts *EDITABLEATTS) GetName() (res string) {
	return editableatts.Name
}

// Stage puts embeddedvalue to the model stage
func (embeddedvalue *EMBEDDEDVALUE) Stage(stage *StageStruct) *EMBEDDEDVALUE {
	stage.EMBEDDEDVALUEs[embeddedvalue] = __member
	stage.EMBEDDEDVALUEs_mapString[embeddedvalue.Name] = embeddedvalue

	return embeddedvalue
}

// Unstage removes embeddedvalue off the model stage
func (embeddedvalue *EMBEDDEDVALUE) Unstage(stage *StageStruct) *EMBEDDEDVALUE {
	delete(stage.EMBEDDEDVALUEs, embeddedvalue)
	delete(stage.EMBEDDEDVALUEs_mapString, embeddedvalue.Name)
	return embeddedvalue
}

// UnstageVoid removes embeddedvalue off the model stage
func (embeddedvalue *EMBEDDEDVALUE) UnstageVoid(stage *StageStruct) {
	delete(stage.EMBEDDEDVALUEs, embeddedvalue)
	delete(stage.EMBEDDEDVALUEs_mapString, embeddedvalue.Name)
}

// commit embeddedvalue to the back repo (if it is already staged)
func (embeddedvalue *EMBEDDEDVALUE) Commit(stage *StageStruct) *EMBEDDEDVALUE {
	if _, ok := stage.EMBEDDEDVALUEs[embeddedvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEMBEDDEDVALUE(embeddedvalue)
		}
	}
	return embeddedvalue
}

func (embeddedvalue *EMBEDDEDVALUE) CommitVoid(stage *StageStruct) {
	embeddedvalue.Commit(stage)
}

// Checkout embeddedvalue to the back repo (if it is already staged)
func (embeddedvalue *EMBEDDEDVALUE) Checkout(stage *StageStruct) *EMBEDDEDVALUE {
	if _, ok := stage.EMBEDDEDVALUEs[embeddedvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEMBEDDEDVALUE(embeddedvalue)
		}
	}
	return embeddedvalue
}

// for satisfaction of GongStruct interface
func (embeddedvalue *EMBEDDEDVALUE) GetName() (res string) {
	return embeddedvalue.Name
}

// Stage puts enumvalue to the model stage
func (enumvalue *ENUMVALUE) Stage(stage *StageStruct) *ENUMVALUE {
	stage.ENUMVALUEs[enumvalue] = __member
	stage.ENUMVALUEs_mapString[enumvalue.Name] = enumvalue

	return enumvalue
}

// Unstage removes enumvalue off the model stage
func (enumvalue *ENUMVALUE) Unstage(stage *StageStruct) *ENUMVALUE {
	delete(stage.ENUMVALUEs, enumvalue)
	delete(stage.ENUMVALUEs_mapString, enumvalue.Name)
	return enumvalue
}

// UnstageVoid removes enumvalue off the model stage
func (enumvalue *ENUMVALUE) UnstageVoid(stage *StageStruct) {
	delete(stage.ENUMVALUEs, enumvalue)
	delete(stage.ENUMVALUEs_mapString, enumvalue.Name)
}

// commit enumvalue to the back repo (if it is already staged)
func (enumvalue *ENUMVALUE) Commit(stage *StageStruct) *ENUMVALUE {
	if _, ok := stage.ENUMVALUEs[enumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitENUMVALUE(enumvalue)
		}
	}
	return enumvalue
}

func (enumvalue *ENUMVALUE) CommitVoid(stage *StageStruct) {
	enumvalue.Commit(stage)
}

// Checkout enumvalue to the back repo (if it is already staged)
func (enumvalue *ENUMVALUE) Checkout(stage *StageStruct) *ENUMVALUE {
	if _, ok := stage.ENUMVALUEs[enumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutENUMVALUE(enumvalue)
		}
	}
	return enumvalue
}

// for satisfaction of GongStruct interface
func (enumvalue *ENUMVALUE) GetName() (res string) {
	return enumvalue.Name
}

// Stage puts object to the model stage
func (object *OBJECT) Stage(stage *StageStruct) *OBJECT {
	stage.OBJECTs[object] = __member
	stage.OBJECTs_mapString[object.Name] = object

	return object
}

// Unstage removes object off the model stage
func (object *OBJECT) Unstage(stage *StageStruct) *OBJECT {
	delete(stage.OBJECTs, object)
	delete(stage.OBJECTs_mapString, object.Name)
	return object
}

// UnstageVoid removes object off the model stage
func (object *OBJECT) UnstageVoid(stage *StageStruct) {
	delete(stage.OBJECTs, object)
	delete(stage.OBJECTs_mapString, object.Name)
}

// commit object to the back repo (if it is already staged)
func (object *OBJECT) Commit(stage *StageStruct) *OBJECT {
	if _, ok := stage.OBJECTs[object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitOBJECT(object)
		}
	}
	return object
}

func (object *OBJECT) CommitVoid(stage *StageStruct) {
	object.Commit(stage)
}

// Checkout object to the back repo (if it is already staged)
func (object *OBJECT) Checkout(stage *StageStruct) *OBJECT {
	if _, ok := stage.OBJECTs[object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutOBJECT(object)
		}
	}
	return object
}

// for satisfaction of GongStruct interface
func (object *OBJECT) GetName() (res string) {
	return object.Name
}

// Stage puts properties to the model stage
func (properties *PROPERTIES) Stage(stage *StageStruct) *PROPERTIES {
	stage.PROPERTIESs[properties] = __member
	stage.PROPERTIESs_mapString[properties.Name] = properties

	return properties
}

// Unstage removes properties off the model stage
func (properties *PROPERTIES) Unstage(stage *StageStruct) *PROPERTIES {
	delete(stage.PROPERTIESs, properties)
	delete(stage.PROPERTIESs_mapString, properties.Name)
	return properties
}

// UnstageVoid removes properties off the model stage
func (properties *PROPERTIES) UnstageVoid(stage *StageStruct) {
	delete(stage.PROPERTIESs, properties)
	delete(stage.PROPERTIESs_mapString, properties.Name)
}

// commit properties to the back repo (if it is already staged)
func (properties *PROPERTIES) Commit(stage *StageStruct) *PROPERTIES {
	if _, ok := stage.PROPERTIESs[properties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPROPERTIES(properties)
		}
	}
	return properties
}

func (properties *PROPERTIES) CommitVoid(stage *StageStruct) {
	properties.Commit(stage)
}

// Checkout properties to the back repo (if it is already staged)
func (properties *PROPERTIES) Checkout(stage *StageStruct) *PROPERTIES {
	if _, ok := stage.PROPERTIESs[properties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPROPERTIES(properties)
		}
	}
	return properties
}

// for satisfaction of GongStruct interface
func (properties *PROPERTIES) GetName() (res string) {
	return properties.Name
}

// Stage puts relationgroup to the model stage
func (relationgroup *RELATIONGROUP) Stage(stage *StageStruct) *RELATIONGROUP {
	stage.RELATIONGROUPs[relationgroup] = __member
	stage.RELATIONGROUPs_mapString[relationgroup.Name] = relationgroup

	return relationgroup
}

// Unstage removes relationgroup off the model stage
func (relationgroup *RELATIONGROUP) Unstage(stage *StageStruct) *RELATIONGROUP {
	delete(stage.RELATIONGROUPs, relationgroup)
	delete(stage.RELATIONGROUPs_mapString, relationgroup.Name)
	return relationgroup
}

// UnstageVoid removes relationgroup off the model stage
func (relationgroup *RELATIONGROUP) UnstageVoid(stage *StageStruct) {
	delete(stage.RELATIONGROUPs, relationgroup)
	delete(stage.RELATIONGROUPs_mapString, relationgroup.Name)
}

// commit relationgroup to the back repo (if it is already staged)
func (relationgroup *RELATIONGROUP) Commit(stage *StageStruct) *RELATIONGROUP {
	if _, ok := stage.RELATIONGROUPs[relationgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATIONGROUP(relationgroup)
		}
	}
	return relationgroup
}

func (relationgroup *RELATIONGROUP) CommitVoid(stage *StageStruct) {
	relationgroup.Commit(stage)
}

// Checkout relationgroup to the back repo (if it is already staged)
func (relationgroup *RELATIONGROUP) Checkout(stage *StageStruct) *RELATIONGROUP {
	if _, ok := stage.RELATIONGROUPs[relationgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATIONGROUP(relationgroup)
		}
	}
	return relationgroup
}

// for satisfaction of GongStruct interface
func (relationgroup *RELATIONGROUP) GetName() (res string) {
	return relationgroup.Name
}

// Stage puts relationgrouptype to the model stage
func (relationgrouptype *RELATIONGROUPTYPE) Stage(stage *StageStruct) *RELATIONGROUPTYPE {
	stage.RELATIONGROUPTYPEs[relationgrouptype] = __member
	stage.RELATIONGROUPTYPEs_mapString[relationgrouptype.Name] = relationgrouptype

	return relationgrouptype
}

// Unstage removes relationgrouptype off the model stage
func (relationgrouptype *RELATIONGROUPTYPE) Unstage(stage *StageStruct) *RELATIONGROUPTYPE {
	delete(stage.RELATIONGROUPTYPEs, relationgrouptype)
	delete(stage.RELATIONGROUPTYPEs_mapString, relationgrouptype.Name)
	return relationgrouptype
}

// UnstageVoid removes relationgrouptype off the model stage
func (relationgrouptype *RELATIONGROUPTYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.RELATIONGROUPTYPEs, relationgrouptype)
	delete(stage.RELATIONGROUPTYPEs_mapString, relationgrouptype.Name)
}

// commit relationgrouptype to the back repo (if it is already staged)
func (relationgrouptype *RELATIONGROUPTYPE) Commit(stage *StageStruct) *RELATIONGROUPTYPE {
	if _, ok := stage.RELATIONGROUPTYPEs[relationgrouptype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATIONGROUPTYPE(relationgrouptype)
		}
	}
	return relationgrouptype
}

func (relationgrouptype *RELATIONGROUPTYPE) CommitVoid(stage *StageStruct) {
	relationgrouptype.Commit(stage)
}

// Checkout relationgrouptype to the back repo (if it is already staged)
func (relationgrouptype *RELATIONGROUPTYPE) Checkout(stage *StageStruct) *RELATIONGROUPTYPE {
	if _, ok := stage.RELATIONGROUPTYPEs[relationgrouptype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATIONGROUPTYPE(relationgrouptype)
		}
	}
	return relationgrouptype
}

// for satisfaction of GongStruct interface
func (relationgrouptype *RELATIONGROUPTYPE) GetName() (res string) {
	return relationgrouptype.Name
}

// Stage puts reqif to the model stage
func (reqif *REQIF) Stage(stage *StageStruct) *REQIF {
	stage.REQIFs[reqif] = __member
	stage.REQIFs_mapString[reqif.Name] = reqif

	return reqif
}

// Unstage removes reqif off the model stage
func (reqif *REQIF) Unstage(stage *StageStruct) *REQIF {
	delete(stage.REQIFs, reqif)
	delete(stage.REQIFs_mapString, reqif.Name)
	return reqif
}

// UnstageVoid removes reqif off the model stage
func (reqif *REQIF) UnstageVoid(stage *StageStruct) {
	delete(stage.REQIFs, reqif)
	delete(stage.REQIFs_mapString, reqif.Name)
}

// commit reqif to the back repo (if it is already staged)
func (reqif *REQIF) Commit(stage *StageStruct) *REQIF {
	if _, ok := stage.REQIFs[reqif]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQIF(reqif)
		}
	}
	return reqif
}

func (reqif *REQIF) CommitVoid(stage *StageStruct) {
	reqif.Commit(stage)
}

// Checkout reqif to the back repo (if it is already staged)
func (reqif *REQIF) Checkout(stage *StageStruct) *REQIF {
	if _, ok := stage.REQIFs[reqif]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQIF(reqif)
		}
	}
	return reqif
}

// for satisfaction of GongStruct interface
func (reqif *REQIF) GetName() (res string) {
	return reqif.Name
}

// Stage puts reqifcontent to the model stage
func (reqifcontent *REQIFCONTENT) Stage(stage *StageStruct) *REQIFCONTENT {
	stage.REQIFCONTENTs[reqifcontent] = __member
	stage.REQIFCONTENTs_mapString[reqifcontent.Name] = reqifcontent

	return reqifcontent
}

// Unstage removes reqifcontent off the model stage
func (reqifcontent *REQIFCONTENT) Unstage(stage *StageStruct) *REQIFCONTENT {
	delete(stage.REQIFCONTENTs, reqifcontent)
	delete(stage.REQIFCONTENTs_mapString, reqifcontent.Name)
	return reqifcontent
}

// UnstageVoid removes reqifcontent off the model stage
func (reqifcontent *REQIFCONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.REQIFCONTENTs, reqifcontent)
	delete(stage.REQIFCONTENTs_mapString, reqifcontent.Name)
}

// commit reqifcontent to the back repo (if it is already staged)
func (reqifcontent *REQIFCONTENT) Commit(stage *StageStruct) *REQIFCONTENT {
	if _, ok := stage.REQIFCONTENTs[reqifcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQIFCONTENT(reqifcontent)
		}
	}
	return reqifcontent
}

func (reqifcontent *REQIFCONTENT) CommitVoid(stage *StageStruct) {
	reqifcontent.Commit(stage)
}

// Checkout reqifcontent to the back repo (if it is already staged)
func (reqifcontent *REQIFCONTENT) Checkout(stage *StageStruct) *REQIFCONTENT {
	if _, ok := stage.REQIFCONTENTs[reqifcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQIFCONTENT(reqifcontent)
		}
	}
	return reqifcontent
}

// for satisfaction of GongStruct interface
func (reqifcontent *REQIFCONTENT) GetName() (res string) {
	return reqifcontent.Name
}

// Stage puts reqifheader to the model stage
func (reqifheader *REQIFHEADER) Stage(stage *StageStruct) *REQIFHEADER {
	stage.REQIFHEADERs[reqifheader] = __member
	stage.REQIFHEADERs_mapString[reqifheader.Name] = reqifheader

	return reqifheader
}

// Unstage removes reqifheader off the model stage
func (reqifheader *REQIFHEADER) Unstage(stage *StageStruct) *REQIFHEADER {
	delete(stage.REQIFHEADERs, reqifheader)
	delete(stage.REQIFHEADERs_mapString, reqifheader.Name)
	return reqifheader
}

// UnstageVoid removes reqifheader off the model stage
func (reqifheader *REQIFHEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.REQIFHEADERs, reqifheader)
	delete(stage.REQIFHEADERs_mapString, reqifheader.Name)
}

// commit reqifheader to the back repo (if it is already staged)
func (reqifheader *REQIFHEADER) Commit(stage *StageStruct) *REQIFHEADER {
	if _, ok := stage.REQIFHEADERs[reqifheader]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQIFHEADER(reqifheader)
		}
	}
	return reqifheader
}

func (reqifheader *REQIFHEADER) CommitVoid(stage *StageStruct) {
	reqifheader.Commit(stage)
}

// Checkout reqifheader to the back repo (if it is already staged)
func (reqifheader *REQIFHEADER) Checkout(stage *StageStruct) *REQIFHEADER {
	if _, ok := stage.REQIFHEADERs[reqifheader]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQIFHEADER(reqifheader)
		}
	}
	return reqifheader
}

// for satisfaction of GongStruct interface
func (reqifheader *REQIFHEADER) GetName() (res string) {
	return reqifheader.Name
}

// Stage puts reqiftoolextension to the model stage
func (reqiftoolextension *REQIFTOOLEXTENSION) Stage(stage *StageStruct) *REQIFTOOLEXTENSION {
	stage.REQIFTOOLEXTENSIONs[reqiftoolextension] = __member
	stage.REQIFTOOLEXTENSIONs_mapString[reqiftoolextension.Name] = reqiftoolextension

	return reqiftoolextension
}

// Unstage removes reqiftoolextension off the model stage
func (reqiftoolextension *REQIFTOOLEXTENSION) Unstage(stage *StageStruct) *REQIFTOOLEXTENSION {
	delete(stage.REQIFTOOLEXTENSIONs, reqiftoolextension)
	delete(stage.REQIFTOOLEXTENSIONs_mapString, reqiftoolextension.Name)
	return reqiftoolextension
}

// UnstageVoid removes reqiftoolextension off the model stage
func (reqiftoolextension *REQIFTOOLEXTENSION) UnstageVoid(stage *StageStruct) {
	delete(stage.REQIFTOOLEXTENSIONs, reqiftoolextension)
	delete(stage.REQIFTOOLEXTENSIONs_mapString, reqiftoolextension.Name)
}

// commit reqiftoolextension to the back repo (if it is already staged)
func (reqiftoolextension *REQIFTOOLEXTENSION) Commit(stage *StageStruct) *REQIFTOOLEXTENSION {
	if _, ok := stage.REQIFTOOLEXTENSIONs[reqiftoolextension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQIFTOOLEXTENSION(reqiftoolextension)
		}
	}
	return reqiftoolextension
}

func (reqiftoolextension *REQIFTOOLEXTENSION) CommitVoid(stage *StageStruct) {
	reqiftoolextension.Commit(stage)
}

// Checkout reqiftoolextension to the back repo (if it is already staged)
func (reqiftoolextension *REQIFTOOLEXTENSION) Checkout(stage *StageStruct) *REQIFTOOLEXTENSION {
	if _, ok := stage.REQIFTOOLEXTENSIONs[reqiftoolextension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQIFTOOLEXTENSION(reqiftoolextension)
		}
	}
	return reqiftoolextension
}

// for satisfaction of GongStruct interface
func (reqiftoolextension *REQIFTOOLEXTENSION) GetName() (res string) {
	return reqiftoolextension.Name
}

// Stage puts reqtype to the model stage
func (reqtype *REQTYPE) Stage(stage *StageStruct) *REQTYPE {
	stage.REQTYPEs[reqtype] = __member
	stage.REQTYPEs_mapString[reqtype.Name] = reqtype

	return reqtype
}

// Unstage removes reqtype off the model stage
func (reqtype *REQTYPE) Unstage(stage *StageStruct) *REQTYPE {
	delete(stage.REQTYPEs, reqtype)
	delete(stage.REQTYPEs_mapString, reqtype.Name)
	return reqtype
}

// UnstageVoid removes reqtype off the model stage
func (reqtype *REQTYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.REQTYPEs, reqtype)
	delete(stage.REQTYPEs_mapString, reqtype.Name)
}

// commit reqtype to the back repo (if it is already staged)
func (reqtype *REQTYPE) Commit(stage *StageStruct) *REQTYPE {
	if _, ok := stage.REQTYPEs[reqtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQTYPE(reqtype)
		}
	}
	return reqtype
}

func (reqtype *REQTYPE) CommitVoid(stage *StageStruct) {
	reqtype.Commit(stage)
}

// Checkout reqtype to the back repo (if it is already staged)
func (reqtype *REQTYPE) Checkout(stage *StageStruct) *REQTYPE {
	if _, ok := stage.REQTYPEs[reqtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQTYPE(reqtype)
		}
	}
	return reqtype
}

// for satisfaction of GongStruct interface
func (reqtype *REQTYPE) GetName() (res string) {
	return reqtype.Name
}

// Stage puts source to the model stage
func (source *SOURCE) Stage(stage *StageStruct) *SOURCE {
	stage.SOURCEs[source] = __member
	stage.SOURCEs_mapString[source.Name] = source

	return source
}

// Unstage removes source off the model stage
func (source *SOURCE) Unstage(stage *StageStruct) *SOURCE {
	delete(stage.SOURCEs, source)
	delete(stage.SOURCEs_mapString, source.Name)
	return source
}

// UnstageVoid removes source off the model stage
func (source *SOURCE) UnstageVoid(stage *StageStruct) {
	delete(stage.SOURCEs, source)
	delete(stage.SOURCEs_mapString, source.Name)
}

// commit source to the back repo (if it is already staged)
func (source *SOURCE) Commit(stage *StageStruct) *SOURCE {
	if _, ok := stage.SOURCEs[source]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSOURCE(source)
		}
	}
	return source
}

func (source *SOURCE) CommitVoid(stage *StageStruct) {
	source.Commit(stage)
}

// Checkout source to the back repo (if it is already staged)
func (source *SOURCE) Checkout(stage *StageStruct) *SOURCE {
	if _, ok := stage.SOURCEs[source]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSOURCE(source)
		}
	}
	return source
}

// for satisfaction of GongStruct interface
func (source *SOURCE) GetName() (res string) {
	return source.Name
}

// Stage puts sourcespecification to the model stage
func (sourcespecification *SOURCESPECIFICATION) Stage(stage *StageStruct) *SOURCESPECIFICATION {
	stage.SOURCESPECIFICATIONs[sourcespecification] = __member
	stage.SOURCESPECIFICATIONs_mapString[sourcespecification.Name] = sourcespecification

	return sourcespecification
}

// Unstage removes sourcespecification off the model stage
func (sourcespecification *SOURCESPECIFICATION) Unstage(stage *StageStruct) *SOURCESPECIFICATION {
	delete(stage.SOURCESPECIFICATIONs, sourcespecification)
	delete(stage.SOURCESPECIFICATIONs_mapString, sourcespecification.Name)
	return sourcespecification
}

// UnstageVoid removes sourcespecification off the model stage
func (sourcespecification *SOURCESPECIFICATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SOURCESPECIFICATIONs, sourcespecification)
	delete(stage.SOURCESPECIFICATIONs_mapString, sourcespecification.Name)
}

// commit sourcespecification to the back repo (if it is already staged)
func (sourcespecification *SOURCESPECIFICATION) Commit(stage *StageStruct) *SOURCESPECIFICATION {
	if _, ok := stage.SOURCESPECIFICATIONs[sourcespecification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSOURCESPECIFICATION(sourcespecification)
		}
	}
	return sourcespecification
}

func (sourcespecification *SOURCESPECIFICATION) CommitVoid(stage *StageStruct) {
	sourcespecification.Commit(stage)
}

// Checkout sourcespecification to the back repo (if it is already staged)
func (sourcespecification *SOURCESPECIFICATION) Checkout(stage *StageStruct) *SOURCESPECIFICATION {
	if _, ok := stage.SOURCESPECIFICATIONs[sourcespecification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSOURCESPECIFICATION(sourcespecification)
		}
	}
	return sourcespecification
}

// for satisfaction of GongStruct interface
func (sourcespecification *SOURCESPECIFICATION) GetName() (res string) {
	return sourcespecification.Name
}

// Stage puts specattributes to the model stage
func (specattributes *SPECATTRIBUTES) Stage(stage *StageStruct) *SPECATTRIBUTES {
	stage.SPECATTRIBUTESs[specattributes] = __member
	stage.SPECATTRIBUTESs_mapString[specattributes.Name] = specattributes

	return specattributes
}

// Unstage removes specattributes off the model stage
func (specattributes *SPECATTRIBUTES) Unstage(stage *StageStruct) *SPECATTRIBUTES {
	delete(stage.SPECATTRIBUTESs, specattributes)
	delete(stage.SPECATTRIBUTESs_mapString, specattributes.Name)
	return specattributes
}

// UnstageVoid removes specattributes off the model stage
func (specattributes *SPECATTRIBUTES) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECATTRIBUTESs, specattributes)
	delete(stage.SPECATTRIBUTESs_mapString, specattributes.Name)
}

// commit specattributes to the back repo (if it is already staged)
func (specattributes *SPECATTRIBUTES) Commit(stage *StageStruct) *SPECATTRIBUTES {
	if _, ok := stage.SPECATTRIBUTESs[specattributes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECATTRIBUTES(specattributes)
		}
	}
	return specattributes
}

func (specattributes *SPECATTRIBUTES) CommitVoid(stage *StageStruct) {
	specattributes.Commit(stage)
}

// Checkout specattributes to the back repo (if it is already staged)
func (specattributes *SPECATTRIBUTES) Checkout(stage *StageStruct) *SPECATTRIBUTES {
	if _, ok := stage.SPECATTRIBUTESs[specattributes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECATTRIBUTES(specattributes)
		}
	}
	return specattributes
}

// for satisfaction of GongStruct interface
func (specattributes *SPECATTRIBUTES) GetName() (res string) {
	return specattributes.Name
}

// Stage puts spechierarchy to the model stage
func (spechierarchy *SPECHIERARCHY) Stage(stage *StageStruct) *SPECHIERARCHY {
	stage.SPECHIERARCHYs[spechierarchy] = __member
	stage.SPECHIERARCHYs_mapString[spechierarchy.Name] = spechierarchy

	return spechierarchy
}

// Unstage removes spechierarchy off the model stage
func (spechierarchy *SPECHIERARCHY) Unstage(stage *StageStruct) *SPECHIERARCHY {
	delete(stage.SPECHIERARCHYs, spechierarchy)
	delete(stage.SPECHIERARCHYs_mapString, spechierarchy.Name)
	return spechierarchy
}

// UnstageVoid removes spechierarchy off the model stage
func (spechierarchy *SPECHIERARCHY) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECHIERARCHYs, spechierarchy)
	delete(stage.SPECHIERARCHYs_mapString, spechierarchy.Name)
}

// commit spechierarchy to the back repo (if it is already staged)
func (spechierarchy *SPECHIERARCHY) Commit(stage *StageStruct) *SPECHIERARCHY {
	if _, ok := stage.SPECHIERARCHYs[spechierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECHIERARCHY(spechierarchy)
		}
	}
	return spechierarchy
}

func (spechierarchy *SPECHIERARCHY) CommitVoid(stage *StageStruct) {
	spechierarchy.Commit(stage)
}

// Checkout spechierarchy to the back repo (if it is already staged)
func (spechierarchy *SPECHIERARCHY) Checkout(stage *StageStruct) *SPECHIERARCHY {
	if _, ok := stage.SPECHIERARCHYs[spechierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECHIERARCHY(spechierarchy)
		}
	}
	return spechierarchy
}

// for satisfaction of GongStruct interface
func (spechierarchy *SPECHIERARCHY) GetName() (res string) {
	return spechierarchy.Name
}

// Stage puts specification to the model stage
func (specification *SPECIFICATION) Stage(stage *StageStruct) *SPECIFICATION {
	stage.SPECIFICATIONs[specification] = __member
	stage.SPECIFICATIONs_mapString[specification.Name] = specification

	return specification
}

// Unstage removes specification off the model stage
func (specification *SPECIFICATION) Unstage(stage *StageStruct) *SPECIFICATION {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
	return specification
}

// UnstageVoid removes specification off the model stage
func (specification *SPECIFICATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
}

// commit specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Commit(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION(specification)
		}
	}
	return specification
}

func (specification *SPECIFICATION) CommitVoid(stage *StageStruct) {
	specification.Commit(stage)
}

// Checkout specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Checkout(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION(specification)
		}
	}
	return specification
}

// for satisfaction of GongStruct interface
func (specification *SPECIFICATION) GetName() (res string) {
	return specification.Name
}

// Stage puts specifications to the model stage
func (specifications *SPECIFICATIONS) Stage(stage *StageStruct) *SPECIFICATIONS {
	stage.SPECIFICATIONSs[specifications] = __member
	stage.SPECIFICATIONSs_mapString[specifications.Name] = specifications

	return specifications
}

// Unstage removes specifications off the model stage
func (specifications *SPECIFICATIONS) Unstage(stage *StageStruct) *SPECIFICATIONS {
	delete(stage.SPECIFICATIONSs, specifications)
	delete(stage.SPECIFICATIONSs_mapString, specifications.Name)
	return specifications
}

// UnstageVoid removes specifications off the model stage
func (specifications *SPECIFICATIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATIONSs, specifications)
	delete(stage.SPECIFICATIONSs_mapString, specifications.Name)
}

// commit specifications to the back repo (if it is already staged)
func (specifications *SPECIFICATIONS) Commit(stage *StageStruct) *SPECIFICATIONS {
	if _, ok := stage.SPECIFICATIONSs[specifications]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATIONS(specifications)
		}
	}
	return specifications
}

func (specifications *SPECIFICATIONS) CommitVoid(stage *StageStruct) {
	specifications.Commit(stage)
}

// Checkout specifications to the back repo (if it is already staged)
func (specifications *SPECIFICATIONS) Checkout(stage *StageStruct) *SPECIFICATIONS {
	if _, ok := stage.SPECIFICATIONSs[specifications]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATIONS(specifications)
		}
	}
	return specifications
}

// for satisfaction of GongStruct interface
func (specifications *SPECIFICATIONS) GetName() (res string) {
	return specifications.Name
}

// Stage puts specificationtype to the model stage
func (specificationtype *SPECIFICATIONTYPE) Stage(stage *StageStruct) *SPECIFICATIONTYPE {
	stage.SPECIFICATIONTYPEs[specificationtype] = __member
	stage.SPECIFICATIONTYPEs_mapString[specificationtype.Name] = specificationtype

	return specificationtype
}

// Unstage removes specificationtype off the model stage
func (specificationtype *SPECIFICATIONTYPE) Unstage(stage *StageStruct) *SPECIFICATIONTYPE {
	delete(stage.SPECIFICATIONTYPEs, specificationtype)
	delete(stage.SPECIFICATIONTYPEs_mapString, specificationtype.Name)
	return specificationtype
}

// UnstageVoid removes specificationtype off the model stage
func (specificationtype *SPECIFICATIONTYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATIONTYPEs, specificationtype)
	delete(stage.SPECIFICATIONTYPEs_mapString, specificationtype.Name)
}

// commit specificationtype to the back repo (if it is already staged)
func (specificationtype *SPECIFICATIONTYPE) Commit(stage *StageStruct) *SPECIFICATIONTYPE {
	if _, ok := stage.SPECIFICATIONTYPEs[specificationtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATIONTYPE(specificationtype)
		}
	}
	return specificationtype
}

func (specificationtype *SPECIFICATIONTYPE) CommitVoid(stage *StageStruct) {
	specificationtype.Commit(stage)
}

// Checkout specificationtype to the back repo (if it is already staged)
func (specificationtype *SPECIFICATIONTYPE) Checkout(stage *StageStruct) *SPECIFICATIONTYPE {
	if _, ok := stage.SPECIFICATIONTYPEs[specificationtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATIONTYPE(specificationtype)
		}
	}
	return specificationtype
}

// for satisfaction of GongStruct interface
func (specificationtype *SPECIFICATIONTYPE) GetName() (res string) {
	return specificationtype.Name
}

// Stage puts specifiedvalues to the model stage
func (specifiedvalues *SPECIFIEDVALUES) Stage(stage *StageStruct) *SPECIFIEDVALUES {
	stage.SPECIFIEDVALUESs[specifiedvalues] = __member
	stage.SPECIFIEDVALUESs_mapString[specifiedvalues.Name] = specifiedvalues

	return specifiedvalues
}

// Unstage removes specifiedvalues off the model stage
func (specifiedvalues *SPECIFIEDVALUES) Unstage(stage *StageStruct) *SPECIFIEDVALUES {
	delete(stage.SPECIFIEDVALUESs, specifiedvalues)
	delete(stage.SPECIFIEDVALUESs_mapString, specifiedvalues.Name)
	return specifiedvalues
}

// UnstageVoid removes specifiedvalues off the model stage
func (specifiedvalues *SPECIFIEDVALUES) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFIEDVALUESs, specifiedvalues)
	delete(stage.SPECIFIEDVALUESs_mapString, specifiedvalues.Name)
}

// commit specifiedvalues to the back repo (if it is already staged)
func (specifiedvalues *SPECIFIEDVALUES) Commit(stage *StageStruct) *SPECIFIEDVALUES {
	if _, ok := stage.SPECIFIEDVALUESs[specifiedvalues]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFIEDVALUES(specifiedvalues)
		}
	}
	return specifiedvalues
}

func (specifiedvalues *SPECIFIEDVALUES) CommitVoid(stage *StageStruct) {
	specifiedvalues.Commit(stage)
}

// Checkout specifiedvalues to the back repo (if it is already staged)
func (specifiedvalues *SPECIFIEDVALUES) Checkout(stage *StageStruct) *SPECIFIEDVALUES {
	if _, ok := stage.SPECIFIEDVALUESs[specifiedvalues]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFIEDVALUES(specifiedvalues)
		}
	}
	return specifiedvalues
}

// for satisfaction of GongStruct interface
func (specifiedvalues *SPECIFIEDVALUES) GetName() (res string) {
	return specifiedvalues.Name
}

// Stage puts specobject to the model stage
func (specobject *SPECOBJECT) Stage(stage *StageStruct) *SPECOBJECT {
	stage.SPECOBJECTs[specobject] = __member
	stage.SPECOBJECTs_mapString[specobject.Name] = specobject

	return specobject
}

// Unstage removes specobject off the model stage
func (specobject *SPECOBJECT) Unstage(stage *StageStruct) *SPECOBJECT {
	delete(stage.SPECOBJECTs, specobject)
	delete(stage.SPECOBJECTs_mapString, specobject.Name)
	return specobject
}

// UnstageVoid removes specobject off the model stage
func (specobject *SPECOBJECT) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECOBJECTs, specobject)
	delete(stage.SPECOBJECTs_mapString, specobject.Name)
}

// commit specobject to the back repo (if it is already staged)
func (specobject *SPECOBJECT) Commit(stage *StageStruct) *SPECOBJECT {
	if _, ok := stage.SPECOBJECTs[specobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECOBJECT(specobject)
		}
	}
	return specobject
}

func (specobject *SPECOBJECT) CommitVoid(stage *StageStruct) {
	specobject.Commit(stage)
}

// Checkout specobject to the back repo (if it is already staged)
func (specobject *SPECOBJECT) Checkout(stage *StageStruct) *SPECOBJECT {
	if _, ok := stage.SPECOBJECTs[specobject]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECOBJECT(specobject)
		}
	}
	return specobject
}

// for satisfaction of GongStruct interface
func (specobject *SPECOBJECT) GetName() (res string) {
	return specobject.Name
}

// Stage puts specobjects to the model stage
func (specobjects *SPECOBJECTS) Stage(stage *StageStruct) *SPECOBJECTS {
	stage.SPECOBJECTSs[specobjects] = __member
	stage.SPECOBJECTSs_mapString[specobjects.Name] = specobjects

	return specobjects
}

// Unstage removes specobjects off the model stage
func (specobjects *SPECOBJECTS) Unstage(stage *StageStruct) *SPECOBJECTS {
	delete(stage.SPECOBJECTSs, specobjects)
	delete(stage.SPECOBJECTSs_mapString, specobjects.Name)
	return specobjects
}

// UnstageVoid removes specobjects off the model stage
func (specobjects *SPECOBJECTS) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECOBJECTSs, specobjects)
	delete(stage.SPECOBJECTSs_mapString, specobjects.Name)
}

// commit specobjects to the back repo (if it is already staged)
func (specobjects *SPECOBJECTS) Commit(stage *StageStruct) *SPECOBJECTS {
	if _, ok := stage.SPECOBJECTSs[specobjects]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECOBJECTS(specobjects)
		}
	}
	return specobjects
}

func (specobjects *SPECOBJECTS) CommitVoid(stage *StageStruct) {
	specobjects.Commit(stage)
}

// Checkout specobjects to the back repo (if it is already staged)
func (specobjects *SPECOBJECTS) Checkout(stage *StageStruct) *SPECOBJECTS {
	if _, ok := stage.SPECOBJECTSs[specobjects]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECOBJECTS(specobjects)
		}
	}
	return specobjects
}

// for satisfaction of GongStruct interface
func (specobjects *SPECOBJECTS) GetName() (res string) {
	return specobjects.Name
}

// Stage puts specobjecttype to the model stage
func (specobjecttype *SPECOBJECTTYPE) Stage(stage *StageStruct) *SPECOBJECTTYPE {
	stage.SPECOBJECTTYPEs[specobjecttype] = __member
	stage.SPECOBJECTTYPEs_mapString[specobjecttype.Name] = specobjecttype

	return specobjecttype
}

// Unstage removes specobjecttype off the model stage
func (specobjecttype *SPECOBJECTTYPE) Unstage(stage *StageStruct) *SPECOBJECTTYPE {
	delete(stage.SPECOBJECTTYPEs, specobjecttype)
	delete(stage.SPECOBJECTTYPEs_mapString, specobjecttype.Name)
	return specobjecttype
}

// UnstageVoid removes specobjecttype off the model stage
func (specobjecttype *SPECOBJECTTYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECOBJECTTYPEs, specobjecttype)
	delete(stage.SPECOBJECTTYPEs_mapString, specobjecttype.Name)
}

// commit specobjecttype to the back repo (if it is already staged)
func (specobjecttype *SPECOBJECTTYPE) Commit(stage *StageStruct) *SPECOBJECTTYPE {
	if _, ok := stage.SPECOBJECTTYPEs[specobjecttype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECOBJECTTYPE(specobjecttype)
		}
	}
	return specobjecttype
}

func (specobjecttype *SPECOBJECTTYPE) CommitVoid(stage *StageStruct) {
	specobjecttype.Commit(stage)
}

// Checkout specobjecttype to the back repo (if it is already staged)
func (specobjecttype *SPECOBJECTTYPE) Checkout(stage *StageStruct) *SPECOBJECTTYPE {
	if _, ok := stage.SPECOBJECTTYPEs[specobjecttype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECOBJECTTYPE(specobjecttype)
		}
	}
	return specobjecttype
}

// for satisfaction of GongStruct interface
func (specobjecttype *SPECOBJECTTYPE) GetName() (res string) {
	return specobjecttype.Name
}

// Stage puts specrelation to the model stage
func (specrelation *SPECRELATION) Stage(stage *StageStruct) *SPECRELATION {
	stage.SPECRELATIONs[specrelation] = __member
	stage.SPECRELATIONs_mapString[specrelation.Name] = specrelation

	return specrelation
}

// Unstage removes specrelation off the model stage
func (specrelation *SPECRELATION) Unstage(stage *StageStruct) *SPECRELATION {
	delete(stage.SPECRELATIONs, specrelation)
	delete(stage.SPECRELATIONs_mapString, specrelation.Name)
	return specrelation
}

// UnstageVoid removes specrelation off the model stage
func (specrelation *SPECRELATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECRELATIONs, specrelation)
	delete(stage.SPECRELATIONs_mapString, specrelation.Name)
}

// commit specrelation to the back repo (if it is already staged)
func (specrelation *SPECRELATION) Commit(stage *StageStruct) *SPECRELATION {
	if _, ok := stage.SPECRELATIONs[specrelation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECRELATION(specrelation)
		}
	}
	return specrelation
}

func (specrelation *SPECRELATION) CommitVoid(stage *StageStruct) {
	specrelation.Commit(stage)
}

// Checkout specrelation to the back repo (if it is already staged)
func (specrelation *SPECRELATION) Checkout(stage *StageStruct) *SPECRELATION {
	if _, ok := stage.SPECRELATIONs[specrelation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECRELATION(specrelation)
		}
	}
	return specrelation
}

// for satisfaction of GongStruct interface
func (specrelation *SPECRELATION) GetName() (res string) {
	return specrelation.Name
}

// Stage puts specrelationgroups to the model stage
func (specrelationgroups *SPECRELATIONGROUPS) Stage(stage *StageStruct) *SPECRELATIONGROUPS {
	stage.SPECRELATIONGROUPSs[specrelationgroups] = __member
	stage.SPECRELATIONGROUPSs_mapString[specrelationgroups.Name] = specrelationgroups

	return specrelationgroups
}

// Unstage removes specrelationgroups off the model stage
func (specrelationgroups *SPECRELATIONGROUPS) Unstage(stage *StageStruct) *SPECRELATIONGROUPS {
	delete(stage.SPECRELATIONGROUPSs, specrelationgroups)
	delete(stage.SPECRELATIONGROUPSs_mapString, specrelationgroups.Name)
	return specrelationgroups
}

// UnstageVoid removes specrelationgroups off the model stage
func (specrelationgroups *SPECRELATIONGROUPS) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECRELATIONGROUPSs, specrelationgroups)
	delete(stage.SPECRELATIONGROUPSs_mapString, specrelationgroups.Name)
}

// commit specrelationgroups to the back repo (if it is already staged)
func (specrelationgroups *SPECRELATIONGROUPS) Commit(stage *StageStruct) *SPECRELATIONGROUPS {
	if _, ok := stage.SPECRELATIONGROUPSs[specrelationgroups]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECRELATIONGROUPS(specrelationgroups)
		}
	}
	return specrelationgroups
}

func (specrelationgroups *SPECRELATIONGROUPS) CommitVoid(stage *StageStruct) {
	specrelationgroups.Commit(stage)
}

// Checkout specrelationgroups to the back repo (if it is already staged)
func (specrelationgroups *SPECRELATIONGROUPS) Checkout(stage *StageStruct) *SPECRELATIONGROUPS {
	if _, ok := stage.SPECRELATIONGROUPSs[specrelationgroups]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECRELATIONGROUPS(specrelationgroups)
		}
	}
	return specrelationgroups
}

// for satisfaction of GongStruct interface
func (specrelationgroups *SPECRELATIONGROUPS) GetName() (res string) {
	return specrelationgroups.Name
}

// Stage puts specrelations to the model stage
func (specrelations *SPECRELATIONS) Stage(stage *StageStruct) *SPECRELATIONS {
	stage.SPECRELATIONSs[specrelations] = __member
	stage.SPECRELATIONSs_mapString[specrelations.Name] = specrelations

	return specrelations
}

// Unstage removes specrelations off the model stage
func (specrelations *SPECRELATIONS) Unstage(stage *StageStruct) *SPECRELATIONS {
	delete(stage.SPECRELATIONSs, specrelations)
	delete(stage.SPECRELATIONSs_mapString, specrelations.Name)
	return specrelations
}

// UnstageVoid removes specrelations off the model stage
func (specrelations *SPECRELATIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECRELATIONSs, specrelations)
	delete(stage.SPECRELATIONSs_mapString, specrelations.Name)
}

// commit specrelations to the back repo (if it is already staged)
func (specrelations *SPECRELATIONS) Commit(stage *StageStruct) *SPECRELATIONS {
	if _, ok := stage.SPECRELATIONSs[specrelations]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECRELATIONS(specrelations)
		}
	}
	return specrelations
}

func (specrelations *SPECRELATIONS) CommitVoid(stage *StageStruct) {
	specrelations.Commit(stage)
}

// Checkout specrelations to the back repo (if it is already staged)
func (specrelations *SPECRELATIONS) Checkout(stage *StageStruct) *SPECRELATIONS {
	if _, ok := stage.SPECRELATIONSs[specrelations]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECRELATIONS(specrelations)
		}
	}
	return specrelations
}

// for satisfaction of GongStruct interface
func (specrelations *SPECRELATIONS) GetName() (res string) {
	return specrelations.Name
}

// Stage puts specrelationtype to the model stage
func (specrelationtype *SPECRELATIONTYPE) Stage(stage *StageStruct) *SPECRELATIONTYPE {
	stage.SPECRELATIONTYPEs[specrelationtype] = __member
	stage.SPECRELATIONTYPEs_mapString[specrelationtype.Name] = specrelationtype

	return specrelationtype
}

// Unstage removes specrelationtype off the model stage
func (specrelationtype *SPECRELATIONTYPE) Unstage(stage *StageStruct) *SPECRELATIONTYPE {
	delete(stage.SPECRELATIONTYPEs, specrelationtype)
	delete(stage.SPECRELATIONTYPEs_mapString, specrelationtype.Name)
	return specrelationtype
}

// UnstageVoid removes specrelationtype off the model stage
func (specrelationtype *SPECRELATIONTYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECRELATIONTYPEs, specrelationtype)
	delete(stage.SPECRELATIONTYPEs_mapString, specrelationtype.Name)
}

// commit specrelationtype to the back repo (if it is already staged)
func (specrelationtype *SPECRELATIONTYPE) Commit(stage *StageStruct) *SPECRELATIONTYPE {
	if _, ok := stage.SPECRELATIONTYPEs[specrelationtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECRELATIONTYPE(specrelationtype)
		}
	}
	return specrelationtype
}

func (specrelationtype *SPECRELATIONTYPE) CommitVoid(stage *StageStruct) {
	specrelationtype.Commit(stage)
}

// Checkout specrelationtype to the back repo (if it is already staged)
func (specrelationtype *SPECRELATIONTYPE) Checkout(stage *StageStruct) *SPECRELATIONTYPE {
	if _, ok := stage.SPECRELATIONTYPEs[specrelationtype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECRELATIONTYPE(specrelationtype)
		}
	}
	return specrelationtype
}

// for satisfaction of GongStruct interface
func (specrelationtype *SPECRELATIONTYPE) GetName() (res string) {
	return specrelationtype.Name
}

// Stage puts spectypes to the model stage
func (spectypes *SPECTYPES) Stage(stage *StageStruct) *SPECTYPES {
	stage.SPECTYPESs[spectypes] = __member
	stage.SPECTYPESs_mapString[spectypes.Name] = spectypes

	return spectypes
}

// Unstage removes spectypes off the model stage
func (spectypes *SPECTYPES) Unstage(stage *StageStruct) *SPECTYPES {
	delete(stage.SPECTYPESs, spectypes)
	delete(stage.SPECTYPESs_mapString, spectypes.Name)
	return spectypes
}

// UnstageVoid removes spectypes off the model stage
func (spectypes *SPECTYPES) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECTYPESs, spectypes)
	delete(stage.SPECTYPESs_mapString, spectypes.Name)
}

// commit spectypes to the back repo (if it is already staged)
func (spectypes *SPECTYPES) Commit(stage *StageStruct) *SPECTYPES {
	if _, ok := stage.SPECTYPESs[spectypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECTYPES(spectypes)
		}
	}
	return spectypes
}

func (spectypes *SPECTYPES) CommitVoid(stage *StageStruct) {
	spectypes.Commit(stage)
}

// Checkout spectypes to the back repo (if it is already staged)
func (spectypes *SPECTYPES) Checkout(stage *StageStruct) *SPECTYPES {
	if _, ok := stage.SPECTYPESs[spectypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECTYPES(spectypes)
		}
	}
	return spectypes
}

// for satisfaction of GongStruct interface
func (spectypes *SPECTYPES) GetName() (res string) {
	return spectypes.Name
}

// Stage puts target to the model stage
func (target *TARGET) Stage(stage *StageStruct) *TARGET {
	stage.TARGETs[target] = __member
	stage.TARGETs_mapString[target.Name] = target

	return target
}

// Unstage removes target off the model stage
func (target *TARGET) Unstage(stage *StageStruct) *TARGET {
	delete(stage.TARGETs, target)
	delete(stage.TARGETs_mapString, target.Name)
	return target
}

// UnstageVoid removes target off the model stage
func (target *TARGET) UnstageVoid(stage *StageStruct) {
	delete(stage.TARGETs, target)
	delete(stage.TARGETs_mapString, target.Name)
}

// commit target to the back repo (if it is already staged)
func (target *TARGET) Commit(stage *StageStruct) *TARGET {
	if _, ok := stage.TARGETs[target]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTARGET(target)
		}
	}
	return target
}

func (target *TARGET) CommitVoid(stage *StageStruct) {
	target.Commit(stage)
}

// Checkout target to the back repo (if it is already staged)
func (target *TARGET) Checkout(stage *StageStruct) *TARGET {
	if _, ok := stage.TARGETs[target]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTARGET(target)
		}
	}
	return target
}

// for satisfaction of GongStruct interface
func (target *TARGET) GetName() (res string) {
	return target.Name
}

// Stage puts targetspecification to the model stage
func (targetspecification *TARGETSPECIFICATION) Stage(stage *StageStruct) *TARGETSPECIFICATION {
	stage.TARGETSPECIFICATIONs[targetspecification] = __member
	stage.TARGETSPECIFICATIONs_mapString[targetspecification.Name] = targetspecification

	return targetspecification
}

// Unstage removes targetspecification off the model stage
func (targetspecification *TARGETSPECIFICATION) Unstage(stage *StageStruct) *TARGETSPECIFICATION {
	delete(stage.TARGETSPECIFICATIONs, targetspecification)
	delete(stage.TARGETSPECIFICATIONs_mapString, targetspecification.Name)
	return targetspecification
}

// UnstageVoid removes targetspecification off the model stage
func (targetspecification *TARGETSPECIFICATION) UnstageVoid(stage *StageStruct) {
	delete(stage.TARGETSPECIFICATIONs, targetspecification)
	delete(stage.TARGETSPECIFICATIONs_mapString, targetspecification.Name)
}

// commit targetspecification to the back repo (if it is already staged)
func (targetspecification *TARGETSPECIFICATION) Commit(stage *StageStruct) *TARGETSPECIFICATION {
	if _, ok := stage.TARGETSPECIFICATIONs[targetspecification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTARGETSPECIFICATION(targetspecification)
		}
	}
	return targetspecification
}

func (targetspecification *TARGETSPECIFICATION) CommitVoid(stage *StageStruct) {
	targetspecification.Commit(stage)
}

// Checkout targetspecification to the back repo (if it is already staged)
func (targetspecification *TARGETSPECIFICATION) Checkout(stage *StageStruct) *TARGETSPECIFICATION {
	if _, ok := stage.TARGETSPECIFICATIONs[targetspecification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTARGETSPECIFICATION(targetspecification)
		}
	}
	return targetspecification
}

// for satisfaction of GongStruct interface
func (targetspecification *TARGETSPECIFICATION) GetName() (res string) {
	return targetspecification.Name
}

// Stage puts theheader to the model stage
func (theheader *THEHEADER) Stage(stage *StageStruct) *THEHEADER {
	stage.THEHEADERs[theheader] = __member
	stage.THEHEADERs_mapString[theheader.Name] = theheader

	return theheader
}

// Unstage removes theheader off the model stage
func (theheader *THEHEADER) Unstage(stage *StageStruct) *THEHEADER {
	delete(stage.THEHEADERs, theheader)
	delete(stage.THEHEADERs_mapString, theheader.Name)
	return theheader
}

// UnstageVoid removes theheader off the model stage
func (theheader *THEHEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.THEHEADERs, theheader)
	delete(stage.THEHEADERs_mapString, theheader.Name)
}

// commit theheader to the back repo (if it is already staged)
func (theheader *THEHEADER) Commit(stage *StageStruct) *THEHEADER {
	if _, ok := stage.THEHEADERs[theheader]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTHEHEADER(theheader)
		}
	}
	return theheader
}

func (theheader *THEHEADER) CommitVoid(stage *StageStruct) {
	theheader.Commit(stage)
}

// Checkout theheader to the back repo (if it is already staged)
func (theheader *THEHEADER) Checkout(stage *StageStruct) *THEHEADER {
	if _, ok := stage.THEHEADERs[theheader]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTHEHEADER(theheader)
		}
	}
	return theheader
}

// for satisfaction of GongStruct interface
func (theheader *THEHEADER) GetName() (res string) {
	return theheader.Name
}

// Stage puts toolextensions to the model stage
func (toolextensions *TOOLEXTENSIONS) Stage(stage *StageStruct) *TOOLEXTENSIONS {
	stage.TOOLEXTENSIONSs[toolextensions] = __member
	stage.TOOLEXTENSIONSs_mapString[toolextensions.Name] = toolextensions

	return toolextensions
}

// Unstage removes toolextensions off the model stage
func (toolextensions *TOOLEXTENSIONS) Unstage(stage *StageStruct) *TOOLEXTENSIONS {
	delete(stage.TOOLEXTENSIONSs, toolextensions)
	delete(stage.TOOLEXTENSIONSs_mapString, toolextensions.Name)
	return toolextensions
}

// UnstageVoid removes toolextensions off the model stage
func (toolextensions *TOOLEXTENSIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.TOOLEXTENSIONSs, toolextensions)
	delete(stage.TOOLEXTENSIONSs_mapString, toolextensions.Name)
}

// commit toolextensions to the back repo (if it is already staged)
func (toolextensions *TOOLEXTENSIONS) Commit(stage *StageStruct) *TOOLEXTENSIONS {
	if _, ok := stage.TOOLEXTENSIONSs[toolextensions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTOOLEXTENSIONS(toolextensions)
		}
	}
	return toolextensions
}

func (toolextensions *TOOLEXTENSIONS) CommitVoid(stage *StageStruct) {
	toolextensions.Commit(stage)
}

// Checkout toolextensions to the back repo (if it is already staged)
func (toolextensions *TOOLEXTENSIONS) Checkout(stage *StageStruct) *TOOLEXTENSIONS {
	if _, ok := stage.TOOLEXTENSIONSs[toolextensions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTOOLEXTENSIONS(toolextensions)
		}
	}
	return toolextensions
}

// for satisfaction of GongStruct interface
func (toolextensions *TOOLEXTENSIONS) GetName() (res string) {
	return toolextensions.Name
}

// Stage puts values to the model stage
func (values *VALUES) Stage(stage *StageStruct) *VALUES {
	stage.VALUESs[values] = __member
	stage.VALUESs_mapString[values.Name] = values

	return values
}

// Unstage removes values off the model stage
func (values *VALUES) Unstage(stage *StageStruct) *VALUES {
	delete(stage.VALUESs, values)
	delete(stage.VALUESs_mapString, values.Name)
	return values
}

// UnstageVoid removes values off the model stage
func (values *VALUES) UnstageVoid(stage *StageStruct) {
	delete(stage.VALUESs, values)
	delete(stage.VALUESs_mapString, values.Name)
}

// commit values to the back repo (if it is already staged)
func (values *VALUES) Commit(stage *StageStruct) *VALUES {
	if _, ok := stage.VALUESs[values]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVALUES(values)
		}
	}
	return values
}

func (values *VALUES) CommitVoid(stage *StageStruct) {
	values.Commit(stage)
}

// Checkout values to the back repo (if it is already staged)
func (values *VALUES) Checkout(stage *StageStruct) *VALUES {
	if _, ok := stage.VALUESs[values]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVALUES(values)
		}
	}
	return values
}

// for satisfaction of GongStruct interface
func (values *VALUES) GetName() (res string) {
	return values.Name
}

// Stage puts xhtmlcontent to the model stage
func (xhtmlcontent *XHTMLCONTENT) Stage(stage *StageStruct) *XHTMLCONTENT {
	stage.XHTMLCONTENTs[xhtmlcontent] = __member
	stage.XHTMLCONTENTs_mapString[xhtmlcontent.Name] = xhtmlcontent

	return xhtmlcontent
}

// Unstage removes xhtmlcontent off the model stage
func (xhtmlcontent *XHTMLCONTENT) Unstage(stage *StageStruct) *XHTMLCONTENT {
	delete(stage.XHTMLCONTENTs, xhtmlcontent)
	delete(stage.XHTMLCONTENTs_mapString, xhtmlcontent.Name)
	return xhtmlcontent
}

// UnstageVoid removes xhtmlcontent off the model stage
func (xhtmlcontent *XHTMLCONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.XHTMLCONTENTs, xhtmlcontent)
	delete(stage.XHTMLCONTENTs_mapString, xhtmlcontent.Name)
}

// commit xhtmlcontent to the back repo (if it is already staged)
func (xhtmlcontent *XHTMLCONTENT) Commit(stage *StageStruct) *XHTMLCONTENT {
	if _, ok := stage.XHTMLCONTENTs[xhtmlcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXHTMLCONTENT(xhtmlcontent)
		}
	}
	return xhtmlcontent
}

func (xhtmlcontent *XHTMLCONTENT) CommitVoid(stage *StageStruct) {
	xhtmlcontent.Commit(stage)
}

// Checkout xhtmlcontent to the back repo (if it is already staged)
func (xhtmlcontent *XHTMLCONTENT) Checkout(stage *StageStruct) *XHTMLCONTENT {
	if _, ok := stage.XHTMLCONTENTs[xhtmlcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXHTMLCONTENT(xhtmlcontent)
		}
	}
	return xhtmlcontent
}

// for satisfaction of GongStruct interface
func (xhtmlcontent *XHTMLCONTENT) GetName() (res string) {
	return xhtmlcontent.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMALTERNATIVEID(ALTERNATIVEID *ALTERNATIVEID)
	CreateORMATTRIBUTEDEFINITIONBOOLEAN(ATTRIBUTEDEFINITIONBOOLEAN *ATTRIBUTEDEFINITIONBOOLEAN)
	CreateORMATTRIBUTEDEFINITIONDATE(ATTRIBUTEDEFINITIONDATE *ATTRIBUTEDEFINITIONDATE)
	CreateORMATTRIBUTEDEFINITIONENUMERATION(ATTRIBUTEDEFINITIONENUMERATION *ATTRIBUTEDEFINITIONENUMERATION)
	CreateORMATTRIBUTEDEFINITIONINTEGER(ATTRIBUTEDEFINITIONINTEGER *ATTRIBUTEDEFINITIONINTEGER)
	CreateORMATTRIBUTEDEFINITIONREAL(ATTRIBUTEDEFINITIONREAL *ATTRIBUTEDEFINITIONREAL)
	CreateORMATTRIBUTEDEFINITIONSTRING(ATTRIBUTEDEFINITIONSTRING *ATTRIBUTEDEFINITIONSTRING)
	CreateORMATTRIBUTEDEFINITIONXHTML(ATTRIBUTEDEFINITIONXHTML *ATTRIBUTEDEFINITIONXHTML)
	CreateORMATTRIBUTEVALUEBOOLEAN(ATTRIBUTEVALUEBOOLEAN *ATTRIBUTEVALUEBOOLEAN)
	CreateORMATTRIBUTEVALUEDATE(ATTRIBUTEVALUEDATE *ATTRIBUTEVALUEDATE)
	CreateORMATTRIBUTEVALUEENUMERATION(ATTRIBUTEVALUEENUMERATION *ATTRIBUTEVALUEENUMERATION)
	CreateORMATTRIBUTEVALUEINTEGER(ATTRIBUTEVALUEINTEGER *ATTRIBUTEVALUEINTEGER)
	CreateORMATTRIBUTEVALUEREAL(ATTRIBUTEVALUEREAL *ATTRIBUTEVALUEREAL)
	CreateORMATTRIBUTEVALUESTRING(ATTRIBUTEVALUESTRING *ATTRIBUTEVALUESTRING)
	CreateORMATTRIBUTEVALUEXHTML(ATTRIBUTEVALUEXHTML *ATTRIBUTEVALUEXHTML)
	CreateORMCHILDREN(CHILDREN *CHILDREN)
	CreateORMCORECONTENT(CORECONTENT *CORECONTENT)
	CreateORMDATATYPEDEFINITIONBOOLEAN(DATATYPEDEFINITIONBOOLEAN *DATATYPEDEFINITIONBOOLEAN)
	CreateORMDATATYPEDEFINITIONDATE(DATATYPEDEFINITIONDATE *DATATYPEDEFINITIONDATE)
	CreateORMDATATYPEDEFINITIONENUMERATION(DATATYPEDEFINITIONENUMERATION *DATATYPEDEFINITIONENUMERATION)
	CreateORMDATATYPEDEFINITIONINTEGER(DATATYPEDEFINITIONINTEGER *DATATYPEDEFINITIONINTEGER)
	CreateORMDATATYPEDEFINITIONREAL(DATATYPEDEFINITIONREAL *DATATYPEDEFINITIONREAL)
	CreateORMDATATYPEDEFINITIONSTRING(DATATYPEDEFINITIONSTRING *DATATYPEDEFINITIONSTRING)
	CreateORMDATATYPEDEFINITIONXHTML(DATATYPEDEFINITIONXHTML *DATATYPEDEFINITIONXHTML)
	CreateORMDATATYPES(DATATYPES *DATATYPES)
	CreateORMDEFAULTVALUE(DEFAULTVALUE *DEFAULTVALUE)
	CreateORMDEFINITION(DEFINITION *DEFINITION)
	CreateORMEDITABLEATTS(EDITABLEATTS *EDITABLEATTS)
	CreateORMEMBEDDEDVALUE(EMBEDDEDVALUE *EMBEDDEDVALUE)
	CreateORMENUMVALUE(ENUMVALUE *ENUMVALUE)
	CreateORMOBJECT(OBJECT *OBJECT)
	CreateORMPROPERTIES(PROPERTIES *PROPERTIES)
	CreateORMRELATIONGROUP(RELATIONGROUP *RELATIONGROUP)
	CreateORMRELATIONGROUPTYPE(RELATIONGROUPTYPE *RELATIONGROUPTYPE)
	CreateORMREQIF(REQIF *REQIF)
	CreateORMREQIFCONTENT(REQIFCONTENT *REQIFCONTENT)
	CreateORMREQIFHEADER(REQIFHEADER *REQIFHEADER)
	CreateORMREQIFTOOLEXTENSION(REQIFTOOLEXTENSION *REQIFTOOLEXTENSION)
	CreateORMREQTYPE(REQTYPE *REQTYPE)
	CreateORMSOURCE(SOURCE *SOURCE)
	CreateORMSOURCESPECIFICATION(SOURCESPECIFICATION *SOURCESPECIFICATION)
	CreateORMSPECATTRIBUTES(SPECATTRIBUTES *SPECATTRIBUTES)
	CreateORMSPECHIERARCHY(SPECHIERARCHY *SPECHIERARCHY)
	CreateORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	CreateORMSPECIFICATIONS(SPECIFICATIONS *SPECIFICATIONS)
	CreateORMSPECIFICATIONTYPE(SPECIFICATIONTYPE *SPECIFICATIONTYPE)
	CreateORMSPECIFIEDVALUES(SPECIFIEDVALUES *SPECIFIEDVALUES)
	CreateORMSPECOBJECT(SPECOBJECT *SPECOBJECT)
	CreateORMSPECOBJECTS(SPECOBJECTS *SPECOBJECTS)
	CreateORMSPECOBJECTTYPE(SPECOBJECTTYPE *SPECOBJECTTYPE)
	CreateORMSPECRELATION(SPECRELATION *SPECRELATION)
	CreateORMSPECRELATIONGROUPS(SPECRELATIONGROUPS *SPECRELATIONGROUPS)
	CreateORMSPECRELATIONS(SPECRELATIONS *SPECRELATIONS)
	CreateORMSPECRELATIONTYPE(SPECRELATIONTYPE *SPECRELATIONTYPE)
	CreateORMSPECTYPES(SPECTYPES *SPECTYPES)
	CreateORMTARGET(TARGET *TARGET)
	CreateORMTARGETSPECIFICATION(TARGETSPECIFICATION *TARGETSPECIFICATION)
	CreateORMTHEHEADER(THEHEADER *THEHEADER)
	CreateORMTOOLEXTENSIONS(TOOLEXTENSIONS *TOOLEXTENSIONS)
	CreateORMVALUES(VALUES *VALUES)
	CreateORMXHTMLCONTENT(XHTMLCONTENT *XHTMLCONTENT)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMALTERNATIVEID(ALTERNATIVEID *ALTERNATIVEID)
	DeleteORMATTRIBUTEDEFINITIONBOOLEAN(ATTRIBUTEDEFINITIONBOOLEAN *ATTRIBUTEDEFINITIONBOOLEAN)
	DeleteORMATTRIBUTEDEFINITIONDATE(ATTRIBUTEDEFINITIONDATE *ATTRIBUTEDEFINITIONDATE)
	DeleteORMATTRIBUTEDEFINITIONENUMERATION(ATTRIBUTEDEFINITIONENUMERATION *ATTRIBUTEDEFINITIONENUMERATION)
	DeleteORMATTRIBUTEDEFINITIONINTEGER(ATTRIBUTEDEFINITIONINTEGER *ATTRIBUTEDEFINITIONINTEGER)
	DeleteORMATTRIBUTEDEFINITIONREAL(ATTRIBUTEDEFINITIONREAL *ATTRIBUTEDEFINITIONREAL)
	DeleteORMATTRIBUTEDEFINITIONSTRING(ATTRIBUTEDEFINITIONSTRING *ATTRIBUTEDEFINITIONSTRING)
	DeleteORMATTRIBUTEDEFINITIONXHTML(ATTRIBUTEDEFINITIONXHTML *ATTRIBUTEDEFINITIONXHTML)
	DeleteORMATTRIBUTEVALUEBOOLEAN(ATTRIBUTEVALUEBOOLEAN *ATTRIBUTEVALUEBOOLEAN)
	DeleteORMATTRIBUTEVALUEDATE(ATTRIBUTEVALUEDATE *ATTRIBUTEVALUEDATE)
	DeleteORMATTRIBUTEVALUEENUMERATION(ATTRIBUTEVALUEENUMERATION *ATTRIBUTEVALUEENUMERATION)
	DeleteORMATTRIBUTEVALUEINTEGER(ATTRIBUTEVALUEINTEGER *ATTRIBUTEVALUEINTEGER)
	DeleteORMATTRIBUTEVALUEREAL(ATTRIBUTEVALUEREAL *ATTRIBUTEVALUEREAL)
	DeleteORMATTRIBUTEVALUESTRING(ATTRIBUTEVALUESTRING *ATTRIBUTEVALUESTRING)
	DeleteORMATTRIBUTEVALUEXHTML(ATTRIBUTEVALUEXHTML *ATTRIBUTEVALUEXHTML)
	DeleteORMCHILDREN(CHILDREN *CHILDREN)
	DeleteORMCORECONTENT(CORECONTENT *CORECONTENT)
	DeleteORMDATATYPEDEFINITIONBOOLEAN(DATATYPEDEFINITIONBOOLEAN *DATATYPEDEFINITIONBOOLEAN)
	DeleteORMDATATYPEDEFINITIONDATE(DATATYPEDEFINITIONDATE *DATATYPEDEFINITIONDATE)
	DeleteORMDATATYPEDEFINITIONENUMERATION(DATATYPEDEFINITIONENUMERATION *DATATYPEDEFINITIONENUMERATION)
	DeleteORMDATATYPEDEFINITIONINTEGER(DATATYPEDEFINITIONINTEGER *DATATYPEDEFINITIONINTEGER)
	DeleteORMDATATYPEDEFINITIONREAL(DATATYPEDEFINITIONREAL *DATATYPEDEFINITIONREAL)
	DeleteORMDATATYPEDEFINITIONSTRING(DATATYPEDEFINITIONSTRING *DATATYPEDEFINITIONSTRING)
	DeleteORMDATATYPEDEFINITIONXHTML(DATATYPEDEFINITIONXHTML *DATATYPEDEFINITIONXHTML)
	DeleteORMDATATYPES(DATATYPES *DATATYPES)
	DeleteORMDEFAULTVALUE(DEFAULTVALUE *DEFAULTVALUE)
	DeleteORMDEFINITION(DEFINITION *DEFINITION)
	DeleteORMEDITABLEATTS(EDITABLEATTS *EDITABLEATTS)
	DeleteORMEMBEDDEDVALUE(EMBEDDEDVALUE *EMBEDDEDVALUE)
	DeleteORMENUMVALUE(ENUMVALUE *ENUMVALUE)
	DeleteORMOBJECT(OBJECT *OBJECT)
	DeleteORMPROPERTIES(PROPERTIES *PROPERTIES)
	DeleteORMRELATIONGROUP(RELATIONGROUP *RELATIONGROUP)
	DeleteORMRELATIONGROUPTYPE(RELATIONGROUPTYPE *RELATIONGROUPTYPE)
	DeleteORMREQIF(REQIF *REQIF)
	DeleteORMREQIFCONTENT(REQIFCONTENT *REQIFCONTENT)
	DeleteORMREQIFHEADER(REQIFHEADER *REQIFHEADER)
	DeleteORMREQIFTOOLEXTENSION(REQIFTOOLEXTENSION *REQIFTOOLEXTENSION)
	DeleteORMREQTYPE(REQTYPE *REQTYPE)
	DeleteORMSOURCE(SOURCE *SOURCE)
	DeleteORMSOURCESPECIFICATION(SOURCESPECIFICATION *SOURCESPECIFICATION)
	DeleteORMSPECATTRIBUTES(SPECATTRIBUTES *SPECATTRIBUTES)
	DeleteORMSPECHIERARCHY(SPECHIERARCHY *SPECHIERARCHY)
	DeleteORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	DeleteORMSPECIFICATIONS(SPECIFICATIONS *SPECIFICATIONS)
	DeleteORMSPECIFICATIONTYPE(SPECIFICATIONTYPE *SPECIFICATIONTYPE)
	DeleteORMSPECIFIEDVALUES(SPECIFIEDVALUES *SPECIFIEDVALUES)
	DeleteORMSPECOBJECT(SPECOBJECT *SPECOBJECT)
	DeleteORMSPECOBJECTS(SPECOBJECTS *SPECOBJECTS)
	DeleteORMSPECOBJECTTYPE(SPECOBJECTTYPE *SPECOBJECTTYPE)
	DeleteORMSPECRELATION(SPECRELATION *SPECRELATION)
	DeleteORMSPECRELATIONGROUPS(SPECRELATIONGROUPS *SPECRELATIONGROUPS)
	DeleteORMSPECRELATIONS(SPECRELATIONS *SPECRELATIONS)
	DeleteORMSPECRELATIONTYPE(SPECRELATIONTYPE *SPECRELATIONTYPE)
	DeleteORMSPECTYPES(SPECTYPES *SPECTYPES)
	DeleteORMTARGET(TARGET *TARGET)
	DeleteORMTARGETSPECIFICATION(TARGETSPECIFICATION *TARGETSPECIFICATION)
	DeleteORMTHEHEADER(THEHEADER *THEHEADER)
	DeleteORMTOOLEXTENSIONS(TOOLEXTENSIONS *TOOLEXTENSIONS)
	DeleteORMVALUES(VALUES *VALUES)
	DeleteORMXHTMLCONTENT(XHTMLCONTENT *XHTMLCONTENT)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.ALTERNATIVEIDs = make(map[*ALTERNATIVEID]any)
	stage.ALTERNATIVEIDs_mapString = make(map[string]*ALTERNATIVEID)

	stage.ATTRIBUTEDEFINITIONBOOLEANs = make(map[*ATTRIBUTEDEFINITIONBOOLEAN]any)
	stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString = make(map[string]*ATTRIBUTEDEFINITIONBOOLEAN)

	stage.ATTRIBUTEDEFINITIONDATEs = make(map[*ATTRIBUTEDEFINITIONDATE]any)
	stage.ATTRIBUTEDEFINITIONDATEs_mapString = make(map[string]*ATTRIBUTEDEFINITIONDATE)

	stage.ATTRIBUTEDEFINITIONENUMERATIONs = make(map[*ATTRIBUTEDEFINITIONENUMERATION]any)
	stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString = make(map[string]*ATTRIBUTEDEFINITIONENUMERATION)

	stage.ATTRIBUTEDEFINITIONINTEGERs = make(map[*ATTRIBUTEDEFINITIONINTEGER]any)
	stage.ATTRIBUTEDEFINITIONINTEGERs_mapString = make(map[string]*ATTRIBUTEDEFINITIONINTEGER)

	stage.ATTRIBUTEDEFINITIONREALs = make(map[*ATTRIBUTEDEFINITIONREAL]any)
	stage.ATTRIBUTEDEFINITIONREALs_mapString = make(map[string]*ATTRIBUTEDEFINITIONREAL)

	stage.ATTRIBUTEDEFINITIONSTRINGs = make(map[*ATTRIBUTEDEFINITIONSTRING]any)
	stage.ATTRIBUTEDEFINITIONSTRINGs_mapString = make(map[string]*ATTRIBUTEDEFINITIONSTRING)

	stage.ATTRIBUTEDEFINITIONXHTMLs = make(map[*ATTRIBUTEDEFINITIONXHTML]any)
	stage.ATTRIBUTEDEFINITIONXHTMLs_mapString = make(map[string]*ATTRIBUTEDEFINITIONXHTML)

	stage.ATTRIBUTEVALUEBOOLEANs = make(map[*ATTRIBUTEVALUEBOOLEAN]any)
	stage.ATTRIBUTEVALUEBOOLEANs_mapString = make(map[string]*ATTRIBUTEVALUEBOOLEAN)

	stage.ATTRIBUTEVALUEDATEs = make(map[*ATTRIBUTEVALUEDATE]any)
	stage.ATTRIBUTEVALUEDATEs_mapString = make(map[string]*ATTRIBUTEVALUEDATE)

	stage.ATTRIBUTEVALUEENUMERATIONs = make(map[*ATTRIBUTEVALUEENUMERATION]any)
	stage.ATTRIBUTEVALUEENUMERATIONs_mapString = make(map[string]*ATTRIBUTEVALUEENUMERATION)

	stage.ATTRIBUTEVALUEINTEGERs = make(map[*ATTRIBUTEVALUEINTEGER]any)
	stage.ATTRIBUTEVALUEINTEGERs_mapString = make(map[string]*ATTRIBUTEVALUEINTEGER)

	stage.ATTRIBUTEVALUEREALs = make(map[*ATTRIBUTEVALUEREAL]any)
	stage.ATTRIBUTEVALUEREALs_mapString = make(map[string]*ATTRIBUTEVALUEREAL)

	stage.ATTRIBUTEVALUESTRINGs = make(map[*ATTRIBUTEVALUESTRING]any)
	stage.ATTRIBUTEVALUESTRINGs_mapString = make(map[string]*ATTRIBUTEVALUESTRING)

	stage.ATTRIBUTEVALUEXHTMLs = make(map[*ATTRIBUTEVALUEXHTML]any)
	stage.ATTRIBUTEVALUEXHTMLs_mapString = make(map[string]*ATTRIBUTEVALUEXHTML)

	stage.CHILDRENs = make(map[*CHILDREN]any)
	stage.CHILDRENs_mapString = make(map[string]*CHILDREN)

	stage.CORECONTENTs = make(map[*CORECONTENT]any)
	stage.CORECONTENTs_mapString = make(map[string]*CORECONTENT)

	stage.DATATYPEDEFINITIONBOOLEANs = make(map[*DATATYPEDEFINITIONBOOLEAN]any)
	stage.DATATYPEDEFINITIONBOOLEANs_mapString = make(map[string]*DATATYPEDEFINITIONBOOLEAN)

	stage.DATATYPEDEFINITIONDATEs = make(map[*DATATYPEDEFINITIONDATE]any)
	stage.DATATYPEDEFINITIONDATEs_mapString = make(map[string]*DATATYPEDEFINITIONDATE)

	stage.DATATYPEDEFINITIONENUMERATIONs = make(map[*DATATYPEDEFINITIONENUMERATION]any)
	stage.DATATYPEDEFINITIONENUMERATIONs_mapString = make(map[string]*DATATYPEDEFINITIONENUMERATION)

	stage.DATATYPEDEFINITIONINTEGERs = make(map[*DATATYPEDEFINITIONINTEGER]any)
	stage.DATATYPEDEFINITIONINTEGERs_mapString = make(map[string]*DATATYPEDEFINITIONINTEGER)

	stage.DATATYPEDEFINITIONREALs = make(map[*DATATYPEDEFINITIONREAL]any)
	stage.DATATYPEDEFINITIONREALs_mapString = make(map[string]*DATATYPEDEFINITIONREAL)

	stage.DATATYPEDEFINITIONSTRINGs = make(map[*DATATYPEDEFINITIONSTRING]any)
	stage.DATATYPEDEFINITIONSTRINGs_mapString = make(map[string]*DATATYPEDEFINITIONSTRING)

	stage.DATATYPEDEFINITIONXHTMLs = make(map[*DATATYPEDEFINITIONXHTML]any)
	stage.DATATYPEDEFINITIONXHTMLs_mapString = make(map[string]*DATATYPEDEFINITIONXHTML)

	stage.DATATYPESs = make(map[*DATATYPES]any)
	stage.DATATYPESs_mapString = make(map[string]*DATATYPES)

	stage.DEFAULTVALUEs = make(map[*DEFAULTVALUE]any)
	stage.DEFAULTVALUEs_mapString = make(map[string]*DEFAULTVALUE)

	stage.DEFINITIONs = make(map[*DEFINITION]any)
	stage.DEFINITIONs_mapString = make(map[string]*DEFINITION)

	stage.EDITABLEATTSs = make(map[*EDITABLEATTS]any)
	stage.EDITABLEATTSs_mapString = make(map[string]*EDITABLEATTS)

	stage.EMBEDDEDVALUEs = make(map[*EMBEDDEDVALUE]any)
	stage.EMBEDDEDVALUEs_mapString = make(map[string]*EMBEDDEDVALUE)

	stage.ENUMVALUEs = make(map[*ENUMVALUE]any)
	stage.ENUMVALUEs_mapString = make(map[string]*ENUMVALUE)

	stage.OBJECTs = make(map[*OBJECT]any)
	stage.OBJECTs_mapString = make(map[string]*OBJECT)

	stage.PROPERTIESs = make(map[*PROPERTIES]any)
	stage.PROPERTIESs_mapString = make(map[string]*PROPERTIES)

	stage.RELATIONGROUPs = make(map[*RELATIONGROUP]any)
	stage.RELATIONGROUPs_mapString = make(map[string]*RELATIONGROUP)

	stage.RELATIONGROUPTYPEs = make(map[*RELATIONGROUPTYPE]any)
	stage.RELATIONGROUPTYPEs_mapString = make(map[string]*RELATIONGROUPTYPE)

	stage.REQIFs = make(map[*REQIF]any)
	stage.REQIFs_mapString = make(map[string]*REQIF)

	stage.REQIFCONTENTs = make(map[*REQIFCONTENT]any)
	stage.REQIFCONTENTs_mapString = make(map[string]*REQIFCONTENT)

	stage.REQIFHEADERs = make(map[*REQIFHEADER]any)
	stage.REQIFHEADERs_mapString = make(map[string]*REQIFHEADER)

	stage.REQIFTOOLEXTENSIONs = make(map[*REQIFTOOLEXTENSION]any)
	stage.REQIFTOOLEXTENSIONs_mapString = make(map[string]*REQIFTOOLEXTENSION)

	stage.REQTYPEs = make(map[*REQTYPE]any)
	stage.REQTYPEs_mapString = make(map[string]*REQTYPE)

	stage.SOURCEs = make(map[*SOURCE]any)
	stage.SOURCEs_mapString = make(map[string]*SOURCE)

	stage.SOURCESPECIFICATIONs = make(map[*SOURCESPECIFICATION]any)
	stage.SOURCESPECIFICATIONs_mapString = make(map[string]*SOURCESPECIFICATION)

	stage.SPECATTRIBUTESs = make(map[*SPECATTRIBUTES]any)
	stage.SPECATTRIBUTESs_mapString = make(map[string]*SPECATTRIBUTES)

	stage.SPECHIERARCHYs = make(map[*SPECHIERARCHY]any)
	stage.SPECHIERARCHYs_mapString = make(map[string]*SPECHIERARCHY)

	stage.SPECIFICATIONs = make(map[*SPECIFICATION]any)
	stage.SPECIFICATIONs_mapString = make(map[string]*SPECIFICATION)

	stage.SPECIFICATIONSs = make(map[*SPECIFICATIONS]any)
	stage.SPECIFICATIONSs_mapString = make(map[string]*SPECIFICATIONS)

	stage.SPECIFICATIONTYPEs = make(map[*SPECIFICATIONTYPE]any)
	stage.SPECIFICATIONTYPEs_mapString = make(map[string]*SPECIFICATIONTYPE)

	stage.SPECIFIEDVALUESs = make(map[*SPECIFIEDVALUES]any)
	stage.SPECIFIEDVALUESs_mapString = make(map[string]*SPECIFIEDVALUES)

	stage.SPECOBJECTs = make(map[*SPECOBJECT]any)
	stage.SPECOBJECTs_mapString = make(map[string]*SPECOBJECT)

	stage.SPECOBJECTSs = make(map[*SPECOBJECTS]any)
	stage.SPECOBJECTSs_mapString = make(map[string]*SPECOBJECTS)

	stage.SPECOBJECTTYPEs = make(map[*SPECOBJECTTYPE]any)
	stage.SPECOBJECTTYPEs_mapString = make(map[string]*SPECOBJECTTYPE)

	stage.SPECRELATIONs = make(map[*SPECRELATION]any)
	stage.SPECRELATIONs_mapString = make(map[string]*SPECRELATION)

	stage.SPECRELATIONGROUPSs = make(map[*SPECRELATIONGROUPS]any)
	stage.SPECRELATIONGROUPSs_mapString = make(map[string]*SPECRELATIONGROUPS)

	stage.SPECRELATIONSs = make(map[*SPECRELATIONS]any)
	stage.SPECRELATIONSs_mapString = make(map[string]*SPECRELATIONS)

	stage.SPECRELATIONTYPEs = make(map[*SPECRELATIONTYPE]any)
	stage.SPECRELATIONTYPEs_mapString = make(map[string]*SPECRELATIONTYPE)

	stage.SPECTYPESs = make(map[*SPECTYPES]any)
	stage.SPECTYPESs_mapString = make(map[string]*SPECTYPES)

	stage.TARGETs = make(map[*TARGET]any)
	stage.TARGETs_mapString = make(map[string]*TARGET)

	stage.TARGETSPECIFICATIONs = make(map[*TARGETSPECIFICATION]any)
	stage.TARGETSPECIFICATIONs_mapString = make(map[string]*TARGETSPECIFICATION)

	stage.THEHEADERs = make(map[*THEHEADER]any)
	stage.THEHEADERs_mapString = make(map[string]*THEHEADER)

	stage.TOOLEXTENSIONSs = make(map[*TOOLEXTENSIONS]any)
	stage.TOOLEXTENSIONSs_mapString = make(map[string]*TOOLEXTENSIONS)

	stage.VALUESs = make(map[*VALUES]any)
	stage.VALUESs_mapString = make(map[string]*VALUES)

	stage.XHTMLCONTENTs = make(map[*XHTMLCONTENT]any)
	stage.XHTMLCONTENTs_mapString = make(map[string]*XHTMLCONTENT)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.ALTERNATIVEIDs = nil
	stage.ALTERNATIVEIDs_mapString = nil

	stage.ATTRIBUTEDEFINITIONBOOLEANs = nil
	stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString = nil

	stage.ATTRIBUTEDEFINITIONDATEs = nil
	stage.ATTRIBUTEDEFINITIONDATEs_mapString = nil

	stage.ATTRIBUTEDEFINITIONENUMERATIONs = nil
	stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString = nil

	stage.ATTRIBUTEDEFINITIONINTEGERs = nil
	stage.ATTRIBUTEDEFINITIONINTEGERs_mapString = nil

	stage.ATTRIBUTEDEFINITIONREALs = nil
	stage.ATTRIBUTEDEFINITIONREALs_mapString = nil

	stage.ATTRIBUTEDEFINITIONSTRINGs = nil
	stage.ATTRIBUTEDEFINITIONSTRINGs_mapString = nil

	stage.ATTRIBUTEDEFINITIONXHTMLs = nil
	stage.ATTRIBUTEDEFINITIONXHTMLs_mapString = nil

	stage.ATTRIBUTEVALUEBOOLEANs = nil
	stage.ATTRIBUTEVALUEBOOLEANs_mapString = nil

	stage.ATTRIBUTEVALUEDATEs = nil
	stage.ATTRIBUTEVALUEDATEs_mapString = nil

	stage.ATTRIBUTEVALUEENUMERATIONs = nil
	stage.ATTRIBUTEVALUEENUMERATIONs_mapString = nil

	stage.ATTRIBUTEVALUEINTEGERs = nil
	stage.ATTRIBUTEVALUEINTEGERs_mapString = nil

	stage.ATTRIBUTEVALUEREALs = nil
	stage.ATTRIBUTEVALUEREALs_mapString = nil

	stage.ATTRIBUTEVALUESTRINGs = nil
	stage.ATTRIBUTEVALUESTRINGs_mapString = nil

	stage.ATTRIBUTEVALUEXHTMLs = nil
	stage.ATTRIBUTEVALUEXHTMLs_mapString = nil

	stage.CHILDRENs = nil
	stage.CHILDRENs_mapString = nil

	stage.CORECONTENTs = nil
	stage.CORECONTENTs_mapString = nil

	stage.DATATYPEDEFINITIONBOOLEANs = nil
	stage.DATATYPEDEFINITIONBOOLEANs_mapString = nil

	stage.DATATYPEDEFINITIONDATEs = nil
	stage.DATATYPEDEFINITIONDATEs_mapString = nil

	stage.DATATYPEDEFINITIONENUMERATIONs = nil
	stage.DATATYPEDEFINITIONENUMERATIONs_mapString = nil

	stage.DATATYPEDEFINITIONINTEGERs = nil
	stage.DATATYPEDEFINITIONINTEGERs_mapString = nil

	stage.DATATYPEDEFINITIONREALs = nil
	stage.DATATYPEDEFINITIONREALs_mapString = nil

	stage.DATATYPEDEFINITIONSTRINGs = nil
	stage.DATATYPEDEFINITIONSTRINGs_mapString = nil

	stage.DATATYPEDEFINITIONXHTMLs = nil
	stage.DATATYPEDEFINITIONXHTMLs_mapString = nil

	stage.DATATYPESs = nil
	stage.DATATYPESs_mapString = nil

	stage.DEFAULTVALUEs = nil
	stage.DEFAULTVALUEs_mapString = nil

	stage.DEFINITIONs = nil
	stage.DEFINITIONs_mapString = nil

	stage.EDITABLEATTSs = nil
	stage.EDITABLEATTSs_mapString = nil

	stage.EMBEDDEDVALUEs = nil
	stage.EMBEDDEDVALUEs_mapString = nil

	stage.ENUMVALUEs = nil
	stage.ENUMVALUEs_mapString = nil

	stage.OBJECTs = nil
	stage.OBJECTs_mapString = nil

	stage.PROPERTIESs = nil
	stage.PROPERTIESs_mapString = nil

	stage.RELATIONGROUPs = nil
	stage.RELATIONGROUPs_mapString = nil

	stage.RELATIONGROUPTYPEs = nil
	stage.RELATIONGROUPTYPEs_mapString = nil

	stage.REQIFs = nil
	stage.REQIFs_mapString = nil

	stage.REQIFCONTENTs = nil
	stage.REQIFCONTENTs_mapString = nil

	stage.REQIFHEADERs = nil
	stage.REQIFHEADERs_mapString = nil

	stage.REQIFTOOLEXTENSIONs = nil
	stage.REQIFTOOLEXTENSIONs_mapString = nil

	stage.REQTYPEs = nil
	stage.REQTYPEs_mapString = nil

	stage.SOURCEs = nil
	stage.SOURCEs_mapString = nil

	stage.SOURCESPECIFICATIONs = nil
	stage.SOURCESPECIFICATIONs_mapString = nil

	stage.SPECATTRIBUTESs = nil
	stage.SPECATTRIBUTESs_mapString = nil

	stage.SPECHIERARCHYs = nil
	stage.SPECHIERARCHYs_mapString = nil

	stage.SPECIFICATIONs = nil
	stage.SPECIFICATIONs_mapString = nil

	stage.SPECIFICATIONSs = nil
	stage.SPECIFICATIONSs_mapString = nil

	stage.SPECIFICATIONTYPEs = nil
	stage.SPECIFICATIONTYPEs_mapString = nil

	stage.SPECIFIEDVALUESs = nil
	stage.SPECIFIEDVALUESs_mapString = nil

	stage.SPECOBJECTs = nil
	stage.SPECOBJECTs_mapString = nil

	stage.SPECOBJECTSs = nil
	stage.SPECOBJECTSs_mapString = nil

	stage.SPECOBJECTTYPEs = nil
	stage.SPECOBJECTTYPEs_mapString = nil

	stage.SPECRELATIONs = nil
	stage.SPECRELATIONs_mapString = nil

	stage.SPECRELATIONGROUPSs = nil
	stage.SPECRELATIONGROUPSs_mapString = nil

	stage.SPECRELATIONSs = nil
	stage.SPECRELATIONSs_mapString = nil

	stage.SPECRELATIONTYPEs = nil
	stage.SPECRELATIONTYPEs_mapString = nil

	stage.SPECTYPESs = nil
	stage.SPECTYPESs_mapString = nil

	stage.TARGETs = nil
	stage.TARGETs_mapString = nil

	stage.TARGETSPECIFICATIONs = nil
	stage.TARGETSPECIFICATIONs_mapString = nil

	stage.THEHEADERs = nil
	stage.THEHEADERs_mapString = nil

	stage.TOOLEXTENSIONSs = nil
	stage.TOOLEXTENSIONSs_mapString = nil

	stage.VALUESs = nil
	stage.VALUESs_mapString = nil

	stage.XHTMLCONTENTs = nil
	stage.XHTMLCONTENTs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for alternativeid := range stage.ALTERNATIVEIDs {
		alternativeid.Unstage(stage)
	}

	for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
		attributedefinitionboolean.Unstage(stage)
	}

	for attributedefinitiondate := range stage.ATTRIBUTEDEFINITIONDATEs {
		attributedefinitiondate.Unstage(stage)
	}

	for attributedefinitionenumeration := range stage.ATTRIBUTEDEFINITIONENUMERATIONs {
		attributedefinitionenumeration.Unstage(stage)
	}

	for attributedefinitioninteger := range stage.ATTRIBUTEDEFINITIONINTEGERs {
		attributedefinitioninteger.Unstage(stage)
	}

	for attributedefinitionreal := range stage.ATTRIBUTEDEFINITIONREALs {
		attributedefinitionreal.Unstage(stage)
	}

	for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
		attributedefinitionstring.Unstage(stage)
	}

	for attributedefinitionxhtml := range stage.ATTRIBUTEDEFINITIONXHTMLs {
		attributedefinitionxhtml.Unstage(stage)
	}

	for attributevalueboolean := range stage.ATTRIBUTEVALUEBOOLEANs {
		attributevalueboolean.Unstage(stage)
	}

	for attributevaluedate := range stage.ATTRIBUTEVALUEDATEs {
		attributevaluedate.Unstage(stage)
	}

	for attributevalueenumeration := range stage.ATTRIBUTEVALUEENUMERATIONs {
		attributevalueenumeration.Unstage(stage)
	}

	for attributevalueinteger := range stage.ATTRIBUTEVALUEINTEGERs {
		attributevalueinteger.Unstage(stage)
	}

	for attributevaluereal := range stage.ATTRIBUTEVALUEREALs {
		attributevaluereal.Unstage(stage)
	}

	for attributevaluestring := range stage.ATTRIBUTEVALUESTRINGs {
		attributevaluestring.Unstage(stage)
	}

	for attributevaluexhtml := range stage.ATTRIBUTEVALUEXHTMLs {
		attributevaluexhtml.Unstage(stage)
	}

	for children := range stage.CHILDRENs {
		children.Unstage(stage)
	}

	for corecontent := range stage.CORECONTENTs {
		corecontent.Unstage(stage)
	}

	for datatypedefinitionboolean := range stage.DATATYPEDEFINITIONBOOLEANs {
		datatypedefinitionboolean.Unstage(stage)
	}

	for datatypedefinitiondate := range stage.DATATYPEDEFINITIONDATEs {
		datatypedefinitiondate.Unstage(stage)
	}

	for datatypedefinitionenumeration := range stage.DATATYPEDEFINITIONENUMERATIONs {
		datatypedefinitionenumeration.Unstage(stage)
	}

	for datatypedefinitioninteger := range stage.DATATYPEDEFINITIONINTEGERs {
		datatypedefinitioninteger.Unstage(stage)
	}

	for datatypedefinitionreal := range stage.DATATYPEDEFINITIONREALs {
		datatypedefinitionreal.Unstage(stage)
	}

	for datatypedefinitionstring := range stage.DATATYPEDEFINITIONSTRINGs {
		datatypedefinitionstring.Unstage(stage)
	}

	for datatypedefinitionxhtml := range stage.DATATYPEDEFINITIONXHTMLs {
		datatypedefinitionxhtml.Unstage(stage)
	}

	for datatypes := range stage.DATATYPESs {
		datatypes.Unstage(stage)
	}

	for defaultvalue := range stage.DEFAULTVALUEs {
		defaultvalue.Unstage(stage)
	}

	for definition := range stage.DEFINITIONs {
		definition.Unstage(stage)
	}

	for editableatts := range stage.EDITABLEATTSs {
		editableatts.Unstage(stage)
	}

	for embeddedvalue := range stage.EMBEDDEDVALUEs {
		embeddedvalue.Unstage(stage)
	}

	for enumvalue := range stage.ENUMVALUEs {
		enumvalue.Unstage(stage)
	}

	for object := range stage.OBJECTs {
		object.Unstage(stage)
	}

	for properties := range stage.PROPERTIESs {
		properties.Unstage(stage)
	}

	for relationgroup := range stage.RELATIONGROUPs {
		relationgroup.Unstage(stage)
	}

	for relationgrouptype := range stage.RELATIONGROUPTYPEs {
		relationgrouptype.Unstage(stage)
	}

	for reqif := range stage.REQIFs {
		reqif.Unstage(stage)
	}

	for reqifcontent := range stage.REQIFCONTENTs {
		reqifcontent.Unstage(stage)
	}

	for reqifheader := range stage.REQIFHEADERs {
		reqifheader.Unstage(stage)
	}

	for reqiftoolextension := range stage.REQIFTOOLEXTENSIONs {
		reqiftoolextension.Unstage(stage)
	}

	for reqtype := range stage.REQTYPEs {
		reqtype.Unstage(stage)
	}

	for source := range stage.SOURCEs {
		source.Unstage(stage)
	}

	for sourcespecification := range stage.SOURCESPECIFICATIONs {
		sourcespecification.Unstage(stage)
	}

	for specattributes := range stage.SPECATTRIBUTESs {
		specattributes.Unstage(stage)
	}

	for spechierarchy := range stage.SPECHIERARCHYs {
		spechierarchy.Unstage(stage)
	}

	for specification := range stage.SPECIFICATIONs {
		specification.Unstage(stage)
	}

	for specifications := range stage.SPECIFICATIONSs {
		specifications.Unstage(stage)
	}

	for specificationtype := range stage.SPECIFICATIONTYPEs {
		specificationtype.Unstage(stage)
	}

	for specifiedvalues := range stage.SPECIFIEDVALUESs {
		specifiedvalues.Unstage(stage)
	}

	for specobject := range stage.SPECOBJECTs {
		specobject.Unstage(stage)
	}

	for specobjects := range stage.SPECOBJECTSs {
		specobjects.Unstage(stage)
	}

	for specobjecttype := range stage.SPECOBJECTTYPEs {
		specobjecttype.Unstage(stage)
	}

	for specrelation := range stage.SPECRELATIONs {
		specrelation.Unstage(stage)
	}

	for specrelationgroups := range stage.SPECRELATIONGROUPSs {
		specrelationgroups.Unstage(stage)
	}

	for specrelations := range stage.SPECRELATIONSs {
		specrelations.Unstage(stage)
	}

	for specrelationtype := range stage.SPECRELATIONTYPEs {
		specrelationtype.Unstage(stage)
	}

	for spectypes := range stage.SPECTYPESs {
		spectypes.Unstage(stage)
	}

	for target := range stage.TARGETs {
		target.Unstage(stage)
	}

	for targetspecification := range stage.TARGETSPECIFICATIONs {
		targetspecification.Unstage(stage)
	}

	for theheader := range stage.THEHEADERs {
		theheader.Unstage(stage)
	}

	for toolextensions := range stage.TOOLEXTENSIONSs {
		toolextensions.Unstage(stage)
	}

	for values := range stage.VALUESs {
		values.Unstage(stage)
	}

	for xhtmlcontent := range stage.XHTMLCONTENTs {
		xhtmlcontent.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	ALTERNATIVEID | ATTRIBUTEDEFINITIONBOOLEAN | ATTRIBUTEDEFINITIONDATE | ATTRIBUTEDEFINITIONENUMERATION | ATTRIBUTEDEFINITIONINTEGER | ATTRIBUTEDEFINITIONREAL | ATTRIBUTEDEFINITIONSTRING | ATTRIBUTEDEFINITIONXHTML | ATTRIBUTEVALUEBOOLEAN | ATTRIBUTEVALUEDATE | ATTRIBUTEVALUEENUMERATION | ATTRIBUTEVALUEINTEGER | ATTRIBUTEVALUEREAL | ATTRIBUTEVALUESTRING | ATTRIBUTEVALUEXHTML | CHILDREN | CORECONTENT | DATATYPEDEFINITIONBOOLEAN | DATATYPEDEFINITIONDATE | DATATYPEDEFINITIONENUMERATION | DATATYPEDEFINITIONINTEGER | DATATYPEDEFINITIONREAL | DATATYPEDEFINITIONSTRING | DATATYPEDEFINITIONXHTML | DATATYPES | DEFAULTVALUE | DEFINITION | EDITABLEATTS | EMBEDDEDVALUE | ENUMVALUE | OBJECT | PROPERTIES | RELATIONGROUP | RELATIONGROUPTYPE | REQIF | REQIFCONTENT | REQIFHEADER | REQIFTOOLEXTENSION | REQTYPE | SOURCE | SOURCESPECIFICATION | SPECATTRIBUTES | SPECHIERARCHY | SPECIFICATION | SPECIFICATIONS | SPECIFICATIONTYPE | SPECIFIEDVALUES | SPECOBJECT | SPECOBJECTS | SPECOBJECTTYPE | SPECRELATION | SPECRELATIONGROUPS | SPECRELATIONS | SPECRELATIONTYPE | SPECTYPES | TARGET | TARGETSPECIFICATION | THEHEADER | TOOLEXTENSIONS | VALUES | XHTMLCONTENT
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*ALTERNATIVEID | *ATTRIBUTEDEFINITIONBOOLEAN | *ATTRIBUTEDEFINITIONDATE | *ATTRIBUTEDEFINITIONENUMERATION | *ATTRIBUTEDEFINITIONINTEGER | *ATTRIBUTEDEFINITIONREAL | *ATTRIBUTEDEFINITIONSTRING | *ATTRIBUTEDEFINITIONXHTML | *ATTRIBUTEVALUEBOOLEAN | *ATTRIBUTEVALUEDATE | *ATTRIBUTEVALUEENUMERATION | *ATTRIBUTEVALUEINTEGER | *ATTRIBUTEVALUEREAL | *ATTRIBUTEVALUESTRING | *ATTRIBUTEVALUEXHTML | *CHILDREN | *CORECONTENT | *DATATYPEDEFINITIONBOOLEAN | *DATATYPEDEFINITIONDATE | *DATATYPEDEFINITIONENUMERATION | *DATATYPEDEFINITIONINTEGER | *DATATYPEDEFINITIONREAL | *DATATYPEDEFINITIONSTRING | *DATATYPEDEFINITIONXHTML | *DATATYPES | *DEFAULTVALUE | *DEFINITION | *EDITABLEATTS | *EMBEDDEDVALUE | *ENUMVALUE | *OBJECT | *PROPERTIES | *RELATIONGROUP | *RELATIONGROUPTYPE | *REQIF | *REQIFCONTENT | *REQIFHEADER | *REQIFTOOLEXTENSION | *REQTYPE | *SOURCE | *SOURCESPECIFICATION | *SPECATTRIBUTES | *SPECHIERARCHY | *SPECIFICATION | *SPECIFICATIONS | *SPECIFICATIONTYPE | *SPECIFIEDVALUES | *SPECOBJECT | *SPECOBJECTS | *SPECOBJECTTYPE | *SPECRELATION | *SPECRELATIONGROUPS | *SPECRELATIONS | *SPECRELATIONTYPE | *SPECTYPES | *TARGET | *TARGETSPECIFICATION | *THEHEADER | *TOOLEXTENSIONS | *VALUES | *XHTMLCONTENT
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*ALTERNATIVEID]any |
		map[*ATTRIBUTEDEFINITIONBOOLEAN]any |
		map[*ATTRIBUTEDEFINITIONDATE]any |
		map[*ATTRIBUTEDEFINITIONENUMERATION]any |
		map[*ATTRIBUTEDEFINITIONINTEGER]any |
		map[*ATTRIBUTEDEFINITIONREAL]any |
		map[*ATTRIBUTEDEFINITIONSTRING]any |
		map[*ATTRIBUTEDEFINITIONXHTML]any |
		map[*ATTRIBUTEVALUEBOOLEAN]any |
		map[*ATTRIBUTEVALUEDATE]any |
		map[*ATTRIBUTEVALUEENUMERATION]any |
		map[*ATTRIBUTEVALUEINTEGER]any |
		map[*ATTRIBUTEVALUEREAL]any |
		map[*ATTRIBUTEVALUESTRING]any |
		map[*ATTRIBUTEVALUEXHTML]any |
		map[*CHILDREN]any |
		map[*CORECONTENT]any |
		map[*DATATYPEDEFINITIONBOOLEAN]any |
		map[*DATATYPEDEFINITIONDATE]any |
		map[*DATATYPEDEFINITIONENUMERATION]any |
		map[*DATATYPEDEFINITIONINTEGER]any |
		map[*DATATYPEDEFINITIONREAL]any |
		map[*DATATYPEDEFINITIONSTRING]any |
		map[*DATATYPEDEFINITIONXHTML]any |
		map[*DATATYPES]any |
		map[*DEFAULTVALUE]any |
		map[*DEFINITION]any |
		map[*EDITABLEATTS]any |
		map[*EMBEDDEDVALUE]any |
		map[*ENUMVALUE]any |
		map[*OBJECT]any |
		map[*PROPERTIES]any |
		map[*RELATIONGROUP]any |
		map[*RELATIONGROUPTYPE]any |
		map[*REQIF]any |
		map[*REQIFCONTENT]any |
		map[*REQIFHEADER]any |
		map[*REQIFTOOLEXTENSION]any |
		map[*REQTYPE]any |
		map[*SOURCE]any |
		map[*SOURCESPECIFICATION]any |
		map[*SPECATTRIBUTES]any |
		map[*SPECHIERARCHY]any |
		map[*SPECIFICATION]any |
		map[*SPECIFICATIONS]any |
		map[*SPECIFICATIONTYPE]any |
		map[*SPECIFIEDVALUES]any |
		map[*SPECOBJECT]any |
		map[*SPECOBJECTS]any |
		map[*SPECOBJECTTYPE]any |
		map[*SPECRELATION]any |
		map[*SPECRELATIONGROUPS]any |
		map[*SPECRELATIONS]any |
		map[*SPECRELATIONTYPE]any |
		map[*SPECTYPES]any |
		map[*TARGET]any |
		map[*TARGETSPECIFICATION]any |
		map[*THEHEADER]any |
		map[*TOOLEXTENSIONS]any |
		map[*VALUES]any |
		map[*XHTMLCONTENT]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*ALTERNATIVEID |
		map[string]*ATTRIBUTEDEFINITIONBOOLEAN |
		map[string]*ATTRIBUTEDEFINITIONDATE |
		map[string]*ATTRIBUTEDEFINITIONENUMERATION |
		map[string]*ATTRIBUTEDEFINITIONINTEGER |
		map[string]*ATTRIBUTEDEFINITIONREAL |
		map[string]*ATTRIBUTEDEFINITIONSTRING |
		map[string]*ATTRIBUTEDEFINITIONXHTML |
		map[string]*ATTRIBUTEVALUEBOOLEAN |
		map[string]*ATTRIBUTEVALUEDATE |
		map[string]*ATTRIBUTEVALUEENUMERATION |
		map[string]*ATTRIBUTEVALUEINTEGER |
		map[string]*ATTRIBUTEVALUEREAL |
		map[string]*ATTRIBUTEVALUESTRING |
		map[string]*ATTRIBUTEVALUEXHTML |
		map[string]*CHILDREN |
		map[string]*CORECONTENT |
		map[string]*DATATYPEDEFINITIONBOOLEAN |
		map[string]*DATATYPEDEFINITIONDATE |
		map[string]*DATATYPEDEFINITIONENUMERATION |
		map[string]*DATATYPEDEFINITIONINTEGER |
		map[string]*DATATYPEDEFINITIONREAL |
		map[string]*DATATYPEDEFINITIONSTRING |
		map[string]*DATATYPEDEFINITIONXHTML |
		map[string]*DATATYPES |
		map[string]*DEFAULTVALUE |
		map[string]*DEFINITION |
		map[string]*EDITABLEATTS |
		map[string]*EMBEDDEDVALUE |
		map[string]*ENUMVALUE |
		map[string]*OBJECT |
		map[string]*PROPERTIES |
		map[string]*RELATIONGROUP |
		map[string]*RELATIONGROUPTYPE |
		map[string]*REQIF |
		map[string]*REQIFCONTENT |
		map[string]*REQIFHEADER |
		map[string]*REQIFTOOLEXTENSION |
		map[string]*REQTYPE |
		map[string]*SOURCE |
		map[string]*SOURCESPECIFICATION |
		map[string]*SPECATTRIBUTES |
		map[string]*SPECHIERARCHY |
		map[string]*SPECIFICATION |
		map[string]*SPECIFICATIONS |
		map[string]*SPECIFICATIONTYPE |
		map[string]*SPECIFIEDVALUES |
		map[string]*SPECOBJECT |
		map[string]*SPECOBJECTS |
		map[string]*SPECOBJECTTYPE |
		map[string]*SPECRELATION |
		map[string]*SPECRELATIONGROUPS |
		map[string]*SPECRELATIONS |
		map[string]*SPECRELATIONTYPE |
		map[string]*SPECTYPES |
		map[string]*TARGET |
		map[string]*TARGETSPECIFICATION |
		map[string]*THEHEADER |
		map[string]*TOOLEXTENSIONS |
		map[string]*VALUES |
		map[string]*XHTMLCONTENT |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*ALTERNATIVEID]any:
		return any(&stage.ALTERNATIVEIDs).(*Type)
	case map[*ATTRIBUTEDEFINITIONBOOLEAN]any:
		return any(&stage.ATTRIBUTEDEFINITIONBOOLEANs).(*Type)
	case map[*ATTRIBUTEDEFINITIONDATE]any:
		return any(&stage.ATTRIBUTEDEFINITIONDATEs).(*Type)
	case map[*ATTRIBUTEDEFINITIONENUMERATION]any:
		return any(&stage.ATTRIBUTEDEFINITIONENUMERATIONs).(*Type)
	case map[*ATTRIBUTEDEFINITIONINTEGER]any:
		return any(&stage.ATTRIBUTEDEFINITIONINTEGERs).(*Type)
	case map[*ATTRIBUTEDEFINITIONREAL]any:
		return any(&stage.ATTRIBUTEDEFINITIONREALs).(*Type)
	case map[*ATTRIBUTEDEFINITIONSTRING]any:
		return any(&stage.ATTRIBUTEDEFINITIONSTRINGs).(*Type)
	case map[*ATTRIBUTEDEFINITIONXHTML]any:
		return any(&stage.ATTRIBUTEDEFINITIONXHTMLs).(*Type)
	case map[*ATTRIBUTEVALUEBOOLEAN]any:
		return any(&stage.ATTRIBUTEVALUEBOOLEANs).(*Type)
	case map[*ATTRIBUTEVALUEDATE]any:
		return any(&stage.ATTRIBUTEVALUEDATEs).(*Type)
	case map[*ATTRIBUTEVALUEENUMERATION]any:
		return any(&stage.ATTRIBUTEVALUEENUMERATIONs).(*Type)
	case map[*ATTRIBUTEVALUEINTEGER]any:
		return any(&stage.ATTRIBUTEVALUEINTEGERs).(*Type)
	case map[*ATTRIBUTEVALUEREAL]any:
		return any(&stage.ATTRIBUTEVALUEREALs).(*Type)
	case map[*ATTRIBUTEVALUESTRING]any:
		return any(&stage.ATTRIBUTEVALUESTRINGs).(*Type)
	case map[*ATTRIBUTEVALUEXHTML]any:
		return any(&stage.ATTRIBUTEVALUEXHTMLs).(*Type)
	case map[*CHILDREN]any:
		return any(&stage.CHILDRENs).(*Type)
	case map[*CORECONTENT]any:
		return any(&stage.CORECONTENTs).(*Type)
	case map[*DATATYPEDEFINITIONBOOLEAN]any:
		return any(&stage.DATATYPEDEFINITIONBOOLEANs).(*Type)
	case map[*DATATYPEDEFINITIONDATE]any:
		return any(&stage.DATATYPEDEFINITIONDATEs).(*Type)
	case map[*DATATYPEDEFINITIONENUMERATION]any:
		return any(&stage.DATATYPEDEFINITIONENUMERATIONs).(*Type)
	case map[*DATATYPEDEFINITIONINTEGER]any:
		return any(&stage.DATATYPEDEFINITIONINTEGERs).(*Type)
	case map[*DATATYPEDEFINITIONREAL]any:
		return any(&stage.DATATYPEDEFINITIONREALs).(*Type)
	case map[*DATATYPEDEFINITIONSTRING]any:
		return any(&stage.DATATYPEDEFINITIONSTRINGs).(*Type)
	case map[*DATATYPEDEFINITIONXHTML]any:
		return any(&stage.DATATYPEDEFINITIONXHTMLs).(*Type)
	case map[*DATATYPES]any:
		return any(&stage.DATATYPESs).(*Type)
	case map[*DEFAULTVALUE]any:
		return any(&stage.DEFAULTVALUEs).(*Type)
	case map[*DEFINITION]any:
		return any(&stage.DEFINITIONs).(*Type)
	case map[*EDITABLEATTS]any:
		return any(&stage.EDITABLEATTSs).(*Type)
	case map[*EMBEDDEDVALUE]any:
		return any(&stage.EMBEDDEDVALUEs).(*Type)
	case map[*ENUMVALUE]any:
		return any(&stage.ENUMVALUEs).(*Type)
	case map[*OBJECT]any:
		return any(&stage.OBJECTs).(*Type)
	case map[*PROPERTIES]any:
		return any(&stage.PROPERTIESs).(*Type)
	case map[*RELATIONGROUP]any:
		return any(&stage.RELATIONGROUPs).(*Type)
	case map[*RELATIONGROUPTYPE]any:
		return any(&stage.RELATIONGROUPTYPEs).(*Type)
	case map[*REQIF]any:
		return any(&stage.REQIFs).(*Type)
	case map[*REQIFCONTENT]any:
		return any(&stage.REQIFCONTENTs).(*Type)
	case map[*REQIFHEADER]any:
		return any(&stage.REQIFHEADERs).(*Type)
	case map[*REQIFTOOLEXTENSION]any:
		return any(&stage.REQIFTOOLEXTENSIONs).(*Type)
	case map[*REQTYPE]any:
		return any(&stage.REQTYPEs).(*Type)
	case map[*SOURCE]any:
		return any(&stage.SOURCEs).(*Type)
	case map[*SOURCESPECIFICATION]any:
		return any(&stage.SOURCESPECIFICATIONs).(*Type)
	case map[*SPECATTRIBUTES]any:
		return any(&stage.SPECATTRIBUTESs).(*Type)
	case map[*SPECHIERARCHY]any:
		return any(&stage.SPECHIERARCHYs).(*Type)
	case map[*SPECIFICATION]any:
		return any(&stage.SPECIFICATIONs).(*Type)
	case map[*SPECIFICATIONS]any:
		return any(&stage.SPECIFICATIONSs).(*Type)
	case map[*SPECIFICATIONTYPE]any:
		return any(&stage.SPECIFICATIONTYPEs).(*Type)
	case map[*SPECIFIEDVALUES]any:
		return any(&stage.SPECIFIEDVALUESs).(*Type)
	case map[*SPECOBJECT]any:
		return any(&stage.SPECOBJECTs).(*Type)
	case map[*SPECOBJECTS]any:
		return any(&stage.SPECOBJECTSs).(*Type)
	case map[*SPECOBJECTTYPE]any:
		return any(&stage.SPECOBJECTTYPEs).(*Type)
	case map[*SPECRELATION]any:
		return any(&stage.SPECRELATIONs).(*Type)
	case map[*SPECRELATIONGROUPS]any:
		return any(&stage.SPECRELATIONGROUPSs).(*Type)
	case map[*SPECRELATIONS]any:
		return any(&stage.SPECRELATIONSs).(*Type)
	case map[*SPECRELATIONTYPE]any:
		return any(&stage.SPECRELATIONTYPEs).(*Type)
	case map[*SPECTYPES]any:
		return any(&stage.SPECTYPESs).(*Type)
	case map[*TARGET]any:
		return any(&stage.TARGETs).(*Type)
	case map[*TARGETSPECIFICATION]any:
		return any(&stage.TARGETSPECIFICATIONs).(*Type)
	case map[*THEHEADER]any:
		return any(&stage.THEHEADERs).(*Type)
	case map[*TOOLEXTENSIONS]any:
		return any(&stage.TOOLEXTENSIONSs).(*Type)
	case map[*VALUES]any:
		return any(&stage.VALUESs).(*Type)
	case map[*XHTMLCONTENT]any:
		return any(&stage.XHTMLCONTENTs).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*ALTERNATIVEID:
		return any(&stage.ALTERNATIVEIDs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONBOOLEAN:
		return any(&stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONDATE:
		return any(&stage.ATTRIBUTEDEFINITIONDATEs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONENUMERATION:
		return any(&stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONINTEGER:
		return any(&stage.ATTRIBUTEDEFINITIONINTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONREAL:
		return any(&stage.ATTRIBUTEDEFINITIONREALs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONSTRING:
		return any(&stage.ATTRIBUTEDEFINITIONSTRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTEDEFINITIONXHTML:
		return any(&stage.ATTRIBUTEDEFINITIONXHTMLs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEBOOLEAN:
		return any(&stage.ATTRIBUTEVALUEBOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEDATE:
		return any(&stage.ATTRIBUTEVALUEDATEs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEENUMERATION:
		return any(&stage.ATTRIBUTEVALUEENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEINTEGER:
		return any(&stage.ATTRIBUTEVALUEINTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEREAL:
		return any(&stage.ATTRIBUTEVALUEREALs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUESTRING:
		return any(&stage.ATTRIBUTEVALUESTRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTEVALUEXHTML:
		return any(&stage.ATTRIBUTEVALUEXHTMLs_mapString).(*Type)
	case map[string]*CHILDREN:
		return any(&stage.CHILDRENs_mapString).(*Type)
	case map[string]*CORECONTENT:
		return any(&stage.CORECONTENTs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONBOOLEAN:
		return any(&stage.DATATYPEDEFINITIONBOOLEANs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONDATE:
		return any(&stage.DATATYPEDEFINITIONDATEs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONENUMERATION:
		return any(&stage.DATATYPEDEFINITIONENUMERATIONs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONINTEGER:
		return any(&stage.DATATYPEDEFINITIONINTEGERs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONREAL:
		return any(&stage.DATATYPEDEFINITIONREALs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONSTRING:
		return any(&stage.DATATYPEDEFINITIONSTRINGs_mapString).(*Type)
	case map[string]*DATATYPEDEFINITIONXHTML:
		return any(&stage.DATATYPEDEFINITIONXHTMLs_mapString).(*Type)
	case map[string]*DATATYPES:
		return any(&stage.DATATYPESs_mapString).(*Type)
	case map[string]*DEFAULTVALUE:
		return any(&stage.DEFAULTVALUEs_mapString).(*Type)
	case map[string]*DEFINITION:
		return any(&stage.DEFINITIONs_mapString).(*Type)
	case map[string]*EDITABLEATTS:
		return any(&stage.EDITABLEATTSs_mapString).(*Type)
	case map[string]*EMBEDDEDVALUE:
		return any(&stage.EMBEDDEDVALUEs_mapString).(*Type)
	case map[string]*ENUMVALUE:
		return any(&stage.ENUMVALUEs_mapString).(*Type)
	case map[string]*OBJECT:
		return any(&stage.OBJECTs_mapString).(*Type)
	case map[string]*PROPERTIES:
		return any(&stage.PROPERTIESs_mapString).(*Type)
	case map[string]*RELATIONGROUP:
		return any(&stage.RELATIONGROUPs_mapString).(*Type)
	case map[string]*RELATIONGROUPTYPE:
		return any(&stage.RELATIONGROUPTYPEs_mapString).(*Type)
	case map[string]*REQIF:
		return any(&stage.REQIFs_mapString).(*Type)
	case map[string]*REQIFCONTENT:
		return any(&stage.REQIFCONTENTs_mapString).(*Type)
	case map[string]*REQIFHEADER:
		return any(&stage.REQIFHEADERs_mapString).(*Type)
	case map[string]*REQIFTOOLEXTENSION:
		return any(&stage.REQIFTOOLEXTENSIONs_mapString).(*Type)
	case map[string]*REQTYPE:
		return any(&stage.REQTYPEs_mapString).(*Type)
	case map[string]*SOURCE:
		return any(&stage.SOURCEs_mapString).(*Type)
	case map[string]*SOURCESPECIFICATION:
		return any(&stage.SOURCESPECIFICATIONs_mapString).(*Type)
	case map[string]*SPECATTRIBUTES:
		return any(&stage.SPECATTRIBUTESs_mapString).(*Type)
	case map[string]*SPECHIERARCHY:
		return any(&stage.SPECHIERARCHYs_mapString).(*Type)
	case map[string]*SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*Type)
	case map[string]*SPECIFICATIONS:
		return any(&stage.SPECIFICATIONSs_mapString).(*Type)
	case map[string]*SPECIFICATIONTYPE:
		return any(&stage.SPECIFICATIONTYPEs_mapString).(*Type)
	case map[string]*SPECIFIEDVALUES:
		return any(&stage.SPECIFIEDVALUESs_mapString).(*Type)
	case map[string]*SPECOBJECT:
		return any(&stage.SPECOBJECTs_mapString).(*Type)
	case map[string]*SPECOBJECTS:
		return any(&stage.SPECOBJECTSs_mapString).(*Type)
	case map[string]*SPECOBJECTTYPE:
		return any(&stage.SPECOBJECTTYPEs_mapString).(*Type)
	case map[string]*SPECRELATION:
		return any(&stage.SPECRELATIONs_mapString).(*Type)
	case map[string]*SPECRELATIONGROUPS:
		return any(&stage.SPECRELATIONGROUPSs_mapString).(*Type)
	case map[string]*SPECRELATIONS:
		return any(&stage.SPECRELATIONSs_mapString).(*Type)
	case map[string]*SPECRELATIONTYPE:
		return any(&stage.SPECRELATIONTYPEs_mapString).(*Type)
	case map[string]*SPECTYPES:
		return any(&stage.SPECTYPESs_mapString).(*Type)
	case map[string]*TARGET:
		return any(&stage.TARGETs_mapString).(*Type)
	case map[string]*TARGETSPECIFICATION:
		return any(&stage.TARGETSPECIFICATIONs_mapString).(*Type)
	case map[string]*THEHEADER:
		return any(&stage.THEHEADERs_mapString).(*Type)
	case map[string]*TOOLEXTENSIONS:
		return any(&stage.TOOLEXTENSIONSs_mapString).(*Type)
	case map[string]*VALUES:
		return any(&stage.VALUESs_mapString).(*Type)
	case map[string]*XHTMLCONTENT:
		return any(&stage.XHTMLCONTENTs_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ALTERNATIVEID:
		return any(&stage.ALTERNATIVEIDs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONBOOLEAN:
		return any(&stage.ATTRIBUTEDEFINITIONBOOLEANs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONDATE:
		return any(&stage.ATTRIBUTEDEFINITIONDATEs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONENUMERATION:
		return any(&stage.ATTRIBUTEDEFINITIONENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONINTEGER:
		return any(&stage.ATTRIBUTEDEFINITIONINTEGERs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONREAL:
		return any(&stage.ATTRIBUTEDEFINITIONREALs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONSTRING:
		return any(&stage.ATTRIBUTEDEFINITIONSTRINGs).(*map[*Type]any)
	case ATTRIBUTEDEFINITIONXHTML:
		return any(&stage.ATTRIBUTEDEFINITIONXHTMLs).(*map[*Type]any)
	case ATTRIBUTEVALUEBOOLEAN:
		return any(&stage.ATTRIBUTEVALUEBOOLEANs).(*map[*Type]any)
	case ATTRIBUTEVALUEDATE:
		return any(&stage.ATTRIBUTEVALUEDATEs).(*map[*Type]any)
	case ATTRIBUTEVALUEENUMERATION:
		return any(&stage.ATTRIBUTEVALUEENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTEVALUEINTEGER:
		return any(&stage.ATTRIBUTEVALUEINTEGERs).(*map[*Type]any)
	case ATTRIBUTEVALUEREAL:
		return any(&stage.ATTRIBUTEVALUEREALs).(*map[*Type]any)
	case ATTRIBUTEVALUESTRING:
		return any(&stage.ATTRIBUTEVALUESTRINGs).(*map[*Type]any)
	case ATTRIBUTEVALUEXHTML:
		return any(&stage.ATTRIBUTEVALUEXHTMLs).(*map[*Type]any)
	case CHILDREN:
		return any(&stage.CHILDRENs).(*map[*Type]any)
	case CORECONTENT:
		return any(&stage.CORECONTENTs).(*map[*Type]any)
	case DATATYPEDEFINITIONBOOLEAN:
		return any(&stage.DATATYPEDEFINITIONBOOLEANs).(*map[*Type]any)
	case DATATYPEDEFINITIONDATE:
		return any(&stage.DATATYPEDEFINITIONDATEs).(*map[*Type]any)
	case DATATYPEDEFINITIONENUMERATION:
		return any(&stage.DATATYPEDEFINITIONENUMERATIONs).(*map[*Type]any)
	case DATATYPEDEFINITIONINTEGER:
		return any(&stage.DATATYPEDEFINITIONINTEGERs).(*map[*Type]any)
	case DATATYPEDEFINITIONREAL:
		return any(&stage.DATATYPEDEFINITIONREALs).(*map[*Type]any)
	case DATATYPEDEFINITIONSTRING:
		return any(&stage.DATATYPEDEFINITIONSTRINGs).(*map[*Type]any)
	case DATATYPEDEFINITIONXHTML:
		return any(&stage.DATATYPEDEFINITIONXHTMLs).(*map[*Type]any)
	case DATATYPES:
		return any(&stage.DATATYPESs).(*map[*Type]any)
	case DEFAULTVALUE:
		return any(&stage.DEFAULTVALUEs).(*map[*Type]any)
	case DEFINITION:
		return any(&stage.DEFINITIONs).(*map[*Type]any)
	case EDITABLEATTS:
		return any(&stage.EDITABLEATTSs).(*map[*Type]any)
	case EMBEDDEDVALUE:
		return any(&stage.EMBEDDEDVALUEs).(*map[*Type]any)
	case ENUMVALUE:
		return any(&stage.ENUMVALUEs).(*map[*Type]any)
	case OBJECT:
		return any(&stage.OBJECTs).(*map[*Type]any)
	case PROPERTIES:
		return any(&stage.PROPERTIESs).(*map[*Type]any)
	case RELATIONGROUP:
		return any(&stage.RELATIONGROUPs).(*map[*Type]any)
	case RELATIONGROUPTYPE:
		return any(&stage.RELATIONGROUPTYPEs).(*map[*Type]any)
	case REQIF:
		return any(&stage.REQIFs).(*map[*Type]any)
	case REQIFCONTENT:
		return any(&stage.REQIFCONTENTs).(*map[*Type]any)
	case REQIFHEADER:
		return any(&stage.REQIFHEADERs).(*map[*Type]any)
	case REQIFTOOLEXTENSION:
		return any(&stage.REQIFTOOLEXTENSIONs).(*map[*Type]any)
	case REQTYPE:
		return any(&stage.REQTYPEs).(*map[*Type]any)
	case SOURCE:
		return any(&stage.SOURCEs).(*map[*Type]any)
	case SOURCESPECIFICATION:
		return any(&stage.SOURCESPECIFICATIONs).(*map[*Type]any)
	case SPECATTRIBUTES:
		return any(&stage.SPECATTRIBUTESs).(*map[*Type]any)
	case SPECHIERARCHY:
		return any(&stage.SPECHIERARCHYs).(*map[*Type]any)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[*Type]any)
	case SPECIFICATIONS:
		return any(&stage.SPECIFICATIONSs).(*map[*Type]any)
	case SPECIFICATIONTYPE:
		return any(&stage.SPECIFICATIONTYPEs).(*map[*Type]any)
	case SPECIFIEDVALUES:
		return any(&stage.SPECIFIEDVALUESs).(*map[*Type]any)
	case SPECOBJECT:
		return any(&stage.SPECOBJECTs).(*map[*Type]any)
	case SPECOBJECTS:
		return any(&stage.SPECOBJECTSs).(*map[*Type]any)
	case SPECOBJECTTYPE:
		return any(&stage.SPECOBJECTTYPEs).(*map[*Type]any)
	case SPECRELATION:
		return any(&stage.SPECRELATIONs).(*map[*Type]any)
	case SPECRELATIONGROUPS:
		return any(&stage.SPECRELATIONGROUPSs).(*map[*Type]any)
	case SPECRELATIONS:
		return any(&stage.SPECRELATIONSs).(*map[*Type]any)
	case SPECRELATIONTYPE:
		return any(&stage.SPECRELATIONTYPEs).(*map[*Type]any)
	case SPECTYPES:
		return any(&stage.SPECTYPESs).(*map[*Type]any)
	case TARGET:
		return any(&stage.TARGETs).(*map[*Type]any)
	case TARGETSPECIFICATION:
		return any(&stage.TARGETSPECIFICATIONs).(*map[*Type]any)
	case THEHEADER:
		return any(&stage.THEHEADERs).(*map[*Type]any)
	case TOOLEXTENSIONS:
		return any(&stage.TOOLEXTENSIONSs).(*map[*Type]any)
	case VALUES:
		return any(&stage.VALUESs).(*map[*Type]any)
	case XHTMLCONTENT:
		return any(&stage.XHTMLCONTENTs).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ALTERNATIVEID:
		return any(&stage.ALTERNATIVEIDs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		return any(&stage.ATTRIBUTEDEFINITIONBOOLEANs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONDATE:
		return any(&stage.ATTRIBUTEDEFINITIONDATEs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONENUMERATION:
		return any(&stage.ATTRIBUTEDEFINITIONENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONINTEGER:
		return any(&stage.ATTRIBUTEDEFINITIONINTEGERs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONREAL:
		return any(&stage.ATTRIBUTEDEFINITIONREALs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONSTRING:
		return any(&stage.ATTRIBUTEDEFINITIONSTRINGs).(*map[Type]any)
	case *ATTRIBUTEDEFINITIONXHTML:
		return any(&stage.ATTRIBUTEDEFINITIONXHTMLs).(*map[Type]any)
	case *ATTRIBUTEVALUEBOOLEAN:
		return any(&stage.ATTRIBUTEVALUEBOOLEANs).(*map[Type]any)
	case *ATTRIBUTEVALUEDATE:
		return any(&stage.ATTRIBUTEVALUEDATEs).(*map[Type]any)
	case *ATTRIBUTEVALUEENUMERATION:
		return any(&stage.ATTRIBUTEVALUEENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTEVALUEINTEGER:
		return any(&stage.ATTRIBUTEVALUEINTEGERs).(*map[Type]any)
	case *ATTRIBUTEVALUEREAL:
		return any(&stage.ATTRIBUTEVALUEREALs).(*map[Type]any)
	case *ATTRIBUTEVALUESTRING:
		return any(&stage.ATTRIBUTEVALUESTRINGs).(*map[Type]any)
	case *ATTRIBUTEVALUEXHTML:
		return any(&stage.ATTRIBUTEVALUEXHTMLs).(*map[Type]any)
	case *CHILDREN:
		return any(&stage.CHILDRENs).(*map[Type]any)
	case *CORECONTENT:
		return any(&stage.CORECONTENTs).(*map[Type]any)
	case *DATATYPEDEFINITIONBOOLEAN:
		return any(&stage.DATATYPEDEFINITIONBOOLEANs).(*map[Type]any)
	case *DATATYPEDEFINITIONDATE:
		return any(&stage.DATATYPEDEFINITIONDATEs).(*map[Type]any)
	case *DATATYPEDEFINITIONENUMERATION:
		return any(&stage.DATATYPEDEFINITIONENUMERATIONs).(*map[Type]any)
	case *DATATYPEDEFINITIONINTEGER:
		return any(&stage.DATATYPEDEFINITIONINTEGERs).(*map[Type]any)
	case *DATATYPEDEFINITIONREAL:
		return any(&stage.DATATYPEDEFINITIONREALs).(*map[Type]any)
	case *DATATYPEDEFINITIONSTRING:
		return any(&stage.DATATYPEDEFINITIONSTRINGs).(*map[Type]any)
	case *DATATYPEDEFINITIONXHTML:
		return any(&stage.DATATYPEDEFINITIONXHTMLs).(*map[Type]any)
	case *DATATYPES:
		return any(&stage.DATATYPESs).(*map[Type]any)
	case *DEFAULTVALUE:
		return any(&stage.DEFAULTVALUEs).(*map[Type]any)
	case *DEFINITION:
		return any(&stage.DEFINITIONs).(*map[Type]any)
	case *EDITABLEATTS:
		return any(&stage.EDITABLEATTSs).(*map[Type]any)
	case *EMBEDDEDVALUE:
		return any(&stage.EMBEDDEDVALUEs).(*map[Type]any)
	case *ENUMVALUE:
		return any(&stage.ENUMVALUEs).(*map[Type]any)
	case *OBJECT:
		return any(&stage.OBJECTs).(*map[Type]any)
	case *PROPERTIES:
		return any(&stage.PROPERTIESs).(*map[Type]any)
	case *RELATIONGROUP:
		return any(&stage.RELATIONGROUPs).(*map[Type]any)
	case *RELATIONGROUPTYPE:
		return any(&stage.RELATIONGROUPTYPEs).(*map[Type]any)
	case *REQIF:
		return any(&stage.REQIFs).(*map[Type]any)
	case *REQIFCONTENT:
		return any(&stage.REQIFCONTENTs).(*map[Type]any)
	case *REQIFHEADER:
		return any(&stage.REQIFHEADERs).(*map[Type]any)
	case *REQIFTOOLEXTENSION:
		return any(&stage.REQIFTOOLEXTENSIONs).(*map[Type]any)
	case *REQTYPE:
		return any(&stage.REQTYPEs).(*map[Type]any)
	case *SOURCE:
		return any(&stage.SOURCEs).(*map[Type]any)
	case *SOURCESPECIFICATION:
		return any(&stage.SOURCESPECIFICATIONs).(*map[Type]any)
	case *SPECATTRIBUTES:
		return any(&stage.SPECATTRIBUTESs).(*map[Type]any)
	case *SPECHIERARCHY:
		return any(&stage.SPECHIERARCHYs).(*map[Type]any)
	case *SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[Type]any)
	case *SPECIFICATIONS:
		return any(&stage.SPECIFICATIONSs).(*map[Type]any)
	case *SPECIFICATIONTYPE:
		return any(&stage.SPECIFICATIONTYPEs).(*map[Type]any)
	case *SPECIFIEDVALUES:
		return any(&stage.SPECIFIEDVALUESs).(*map[Type]any)
	case *SPECOBJECT:
		return any(&stage.SPECOBJECTs).(*map[Type]any)
	case *SPECOBJECTS:
		return any(&stage.SPECOBJECTSs).(*map[Type]any)
	case *SPECOBJECTTYPE:
		return any(&stage.SPECOBJECTTYPEs).(*map[Type]any)
	case *SPECRELATION:
		return any(&stage.SPECRELATIONs).(*map[Type]any)
	case *SPECRELATIONGROUPS:
		return any(&stage.SPECRELATIONGROUPSs).(*map[Type]any)
	case *SPECRELATIONS:
		return any(&stage.SPECRELATIONSs).(*map[Type]any)
	case *SPECRELATIONTYPE:
		return any(&stage.SPECRELATIONTYPEs).(*map[Type]any)
	case *SPECTYPES:
		return any(&stage.SPECTYPESs).(*map[Type]any)
	case *TARGET:
		return any(&stage.TARGETs).(*map[Type]any)
	case *TARGETSPECIFICATION:
		return any(&stage.TARGETSPECIFICATIONs).(*map[Type]any)
	case *THEHEADER:
		return any(&stage.THEHEADERs).(*map[Type]any)
	case *TOOLEXTENSIONS:
		return any(&stage.TOOLEXTENSIONSs).(*map[Type]any)
	case *VALUES:
		return any(&stage.VALUESs).(*map[Type]any)
	case *XHTMLCONTENT:
		return any(&stage.XHTMLCONTENTs).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ALTERNATIVEID:
		return any(&stage.ALTERNATIVEIDs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONBOOLEAN:
		return any(&stage.ATTRIBUTEDEFINITIONBOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONDATE:
		return any(&stage.ATTRIBUTEDEFINITIONDATEs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONENUMERATION:
		return any(&stage.ATTRIBUTEDEFINITIONENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONINTEGER:
		return any(&stage.ATTRIBUTEDEFINITIONINTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONREAL:
		return any(&stage.ATTRIBUTEDEFINITIONREALs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONSTRING:
		return any(&stage.ATTRIBUTEDEFINITIONSTRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTEDEFINITIONXHTML:
		return any(&stage.ATTRIBUTEDEFINITIONXHTMLs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEBOOLEAN:
		return any(&stage.ATTRIBUTEVALUEBOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEDATE:
		return any(&stage.ATTRIBUTEVALUEDATEs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEENUMERATION:
		return any(&stage.ATTRIBUTEVALUEENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEINTEGER:
		return any(&stage.ATTRIBUTEVALUEINTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEREAL:
		return any(&stage.ATTRIBUTEVALUEREALs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUESTRING:
		return any(&stage.ATTRIBUTEVALUESTRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTEVALUEXHTML:
		return any(&stage.ATTRIBUTEVALUEXHTMLs_mapString).(*map[string]*Type)
	case CHILDREN:
		return any(&stage.CHILDRENs_mapString).(*map[string]*Type)
	case CORECONTENT:
		return any(&stage.CORECONTENTs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONBOOLEAN:
		return any(&stage.DATATYPEDEFINITIONBOOLEANs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONDATE:
		return any(&stage.DATATYPEDEFINITIONDATEs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONENUMERATION:
		return any(&stage.DATATYPEDEFINITIONENUMERATIONs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONINTEGER:
		return any(&stage.DATATYPEDEFINITIONINTEGERs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONREAL:
		return any(&stage.DATATYPEDEFINITIONREALs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONSTRING:
		return any(&stage.DATATYPEDEFINITIONSTRINGs_mapString).(*map[string]*Type)
	case DATATYPEDEFINITIONXHTML:
		return any(&stage.DATATYPEDEFINITIONXHTMLs_mapString).(*map[string]*Type)
	case DATATYPES:
		return any(&stage.DATATYPESs_mapString).(*map[string]*Type)
	case DEFAULTVALUE:
		return any(&stage.DEFAULTVALUEs_mapString).(*map[string]*Type)
	case DEFINITION:
		return any(&stage.DEFINITIONs_mapString).(*map[string]*Type)
	case EDITABLEATTS:
		return any(&stage.EDITABLEATTSs_mapString).(*map[string]*Type)
	case EMBEDDEDVALUE:
		return any(&stage.EMBEDDEDVALUEs_mapString).(*map[string]*Type)
	case ENUMVALUE:
		return any(&stage.ENUMVALUEs_mapString).(*map[string]*Type)
	case OBJECT:
		return any(&stage.OBJECTs_mapString).(*map[string]*Type)
	case PROPERTIES:
		return any(&stage.PROPERTIESs_mapString).(*map[string]*Type)
	case RELATIONGROUP:
		return any(&stage.RELATIONGROUPs_mapString).(*map[string]*Type)
	case RELATIONGROUPTYPE:
		return any(&stage.RELATIONGROUPTYPEs_mapString).(*map[string]*Type)
	case REQIF:
		return any(&stage.REQIFs_mapString).(*map[string]*Type)
	case REQIFCONTENT:
		return any(&stage.REQIFCONTENTs_mapString).(*map[string]*Type)
	case REQIFHEADER:
		return any(&stage.REQIFHEADERs_mapString).(*map[string]*Type)
	case REQIFTOOLEXTENSION:
		return any(&stage.REQIFTOOLEXTENSIONs_mapString).(*map[string]*Type)
	case REQTYPE:
		return any(&stage.REQTYPEs_mapString).(*map[string]*Type)
	case SOURCE:
		return any(&stage.SOURCEs_mapString).(*map[string]*Type)
	case SOURCESPECIFICATION:
		return any(&stage.SOURCESPECIFICATIONs_mapString).(*map[string]*Type)
	case SPECATTRIBUTES:
		return any(&stage.SPECATTRIBUTESs_mapString).(*map[string]*Type)
	case SPECHIERARCHY:
		return any(&stage.SPECHIERARCHYs_mapString).(*map[string]*Type)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*map[string]*Type)
	case SPECIFICATIONS:
		return any(&stage.SPECIFICATIONSs_mapString).(*map[string]*Type)
	case SPECIFICATIONTYPE:
		return any(&stage.SPECIFICATIONTYPEs_mapString).(*map[string]*Type)
	case SPECIFIEDVALUES:
		return any(&stage.SPECIFIEDVALUESs_mapString).(*map[string]*Type)
	case SPECOBJECT:
		return any(&stage.SPECOBJECTs_mapString).(*map[string]*Type)
	case SPECOBJECTS:
		return any(&stage.SPECOBJECTSs_mapString).(*map[string]*Type)
	case SPECOBJECTTYPE:
		return any(&stage.SPECOBJECTTYPEs_mapString).(*map[string]*Type)
	case SPECRELATION:
		return any(&stage.SPECRELATIONs_mapString).(*map[string]*Type)
	case SPECRELATIONGROUPS:
		return any(&stage.SPECRELATIONGROUPSs_mapString).(*map[string]*Type)
	case SPECRELATIONS:
		return any(&stage.SPECRELATIONSs_mapString).(*map[string]*Type)
	case SPECRELATIONTYPE:
		return any(&stage.SPECRELATIONTYPEs_mapString).(*map[string]*Type)
	case SPECTYPES:
		return any(&stage.SPECTYPESs_mapString).(*map[string]*Type)
	case TARGET:
		return any(&stage.TARGETs_mapString).(*map[string]*Type)
	case TARGETSPECIFICATION:
		return any(&stage.TARGETSPECIFICATIONs_mapString).(*map[string]*Type)
	case THEHEADER:
		return any(&stage.THEHEADERs_mapString).(*map[string]*Type)
	case TOOLEXTENSIONS:
		return any(&stage.TOOLEXTENSIONSs_mapString).(*map[string]*Type)
	case VALUES:
		return any(&stage.VALUESs_mapString).(*map[string]*Type)
	case XHTMLCONTENT:
		return any(&stage.XHTMLCONTENTs_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case ALTERNATIVEID:
		return any(&ALTERNATIVEID{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTEDEFINITIONBOOLEAN:
		return any(&ATTRIBUTEDEFINITIONBOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONDATE:
		return any(&ATTRIBUTEDEFINITIONDATE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONENUMERATION:
		return any(&ATTRIBUTEDEFINITIONENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONINTEGER:
		return any(&ATTRIBUTEDEFINITIONINTEGER{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONREAL:
		return any(&ATTRIBUTEDEFINITIONREAL{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONSTRING:
		return any(&ATTRIBUTEDEFINITIONSTRING{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEDEFINITIONXHTML:
		return any(&ATTRIBUTEDEFINITIONXHTML{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of DEFAULTVALUE with the name of the field
			DEFAULTVALUE: &DEFAULTVALUE{Name: "DEFAULTVALUE"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTEVALUEBOOLEAN:
		return any(&ATTRIBUTEVALUEBOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTEVALUEDATE:
		return any(&ATTRIBUTEVALUEDATE{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTEVALUEENUMERATION:
		return any(&ATTRIBUTEVALUEENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
			// field is initialized with an instance of VALUES with the name of the field
			VALUES: &VALUES{Name: "VALUES"},
		}).(*Type)
	case ATTRIBUTEVALUEINTEGER:
		return any(&ATTRIBUTEVALUEINTEGER{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTEVALUEREAL:
		return any(&ATTRIBUTEVALUEREAL{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTEVALUESTRING:
		return any(&ATTRIBUTEVALUESTRING{
			// Initialisation of associations
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTEVALUEXHTML:
		return any(&ATTRIBUTEVALUEXHTML{
			// Initialisation of associations
			// field is initialized with an instance of XHTMLCONTENT with the name of the field
			THEVALUE: &XHTMLCONTENT{Name: "THEVALUE"},
			// field is initialized with an instance of XHTMLCONTENT with the name of the field
			THEORIGINALVALUE: &XHTMLCONTENT{Name: "THEORIGINALVALUE"},
			// field is initialized with an instance of DEFINITION with the name of the field
			DEFINITION: &DEFINITION{Name: "DEFINITION"},
		}).(*Type)
	case CHILDREN:
		return any(&CHILDREN{
			// Initialisation of associations
			// field is initialized with an instance of SPECHIERARCHY with the name of the field
			SPECHIERARCHY: []*SPECHIERARCHY{{Name: "SPECHIERARCHY"}},
		}).(*Type)
	case CORECONTENT:
		return any(&CORECONTENT{
			// Initialisation of associations
			// field is initialized with an instance of REQIFCONTENT with the name of the field
			REQIFCONTENT: &REQIFCONTENT{Name: "REQIFCONTENT"},
		}).(*Type)
	case DATATYPEDEFINITIONBOOLEAN:
		return any(&DATATYPEDEFINITIONBOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPEDEFINITIONDATE:
		return any(&DATATYPEDEFINITIONDATE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPEDEFINITIONENUMERATION:
		return any(&DATATYPEDEFINITIONENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SPECIFIEDVALUES with the name of the field
			SPECIFIEDVALUES: &SPECIFIEDVALUES{Name: "SPECIFIEDVALUES"},
		}).(*Type)
	case DATATYPEDEFINITIONINTEGER:
		return any(&DATATYPEDEFINITIONINTEGER{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPEDEFINITIONREAL:
		return any(&DATATYPEDEFINITIONREAL{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPEDEFINITIONSTRING:
		return any(&DATATYPEDEFINITIONSTRING{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPEDEFINITIONXHTML:
		return any(&DATATYPEDEFINITIONXHTML{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
		}).(*Type)
	case DATATYPES:
		return any(&DATATYPES{
			// Initialisation of associations
			// field is initialized with an instance of DATATYPEDEFINITIONBOOLEAN with the name of the field
			DATATYPEDEFINITIONBOOLEAN: []*DATATYPEDEFINITIONBOOLEAN{{Name: "DATATYPEDEFINITIONBOOLEAN"}},
			// field is initialized with an instance of DATATYPEDEFINITIONDATE with the name of the field
			DATATYPEDEFINITIONDATE: []*DATATYPEDEFINITIONDATE{{Name: "DATATYPEDEFINITIONDATE"}},
			// field is initialized with an instance of DATATYPEDEFINITIONENUMERATION with the name of the field
			DATATYPEDEFINITIONENUMERATION: []*DATATYPEDEFINITIONENUMERATION{{Name: "DATATYPEDEFINITIONENUMERATION"}},
			// field is initialized with an instance of DATATYPEDEFINITIONINTEGER with the name of the field
			DATATYPEDEFINITIONINTEGER: []*DATATYPEDEFINITIONINTEGER{{Name: "DATATYPEDEFINITIONINTEGER"}},
			// field is initialized with an instance of DATATYPEDEFINITIONREAL with the name of the field
			DATATYPEDEFINITIONREAL: []*DATATYPEDEFINITIONREAL{{Name: "DATATYPEDEFINITIONREAL"}},
			// field is initialized with an instance of DATATYPEDEFINITIONSTRING with the name of the field
			DATATYPEDEFINITIONSTRING: []*DATATYPEDEFINITIONSTRING{{Name: "DATATYPEDEFINITIONSTRING"}},
			// field is initialized with an instance of DATATYPEDEFINITIONXHTML with the name of the field
			DATATYPEDEFINITIONXHTML: []*DATATYPEDEFINITIONXHTML{{Name: "DATATYPEDEFINITIONXHTML"}},
		}).(*Type)
	case DEFAULTVALUE:
		return any(&DEFAULTVALUE{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTEVALUEBOOLEAN with the name of the field
			ATTRIBUTEVALUEBOOLEAN: &ATTRIBUTEVALUEBOOLEAN{Name: "ATTRIBUTEVALUEBOOLEAN"},
		}).(*Type)
	case DEFINITION:
		return any(&DEFINITION{
			// Initialisation of associations
		}).(*Type)
	case EDITABLEATTS:
		return any(&EDITABLEATTS{
			// Initialisation of associations
		}).(*Type)
	case EMBEDDEDVALUE:
		return any(&EMBEDDEDVALUE{
			// Initialisation of associations
		}).(*Type)
	case ENUMVALUE:
		return any(&ENUMVALUE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of PROPERTIES with the name of the field
			PROPERTIES: &PROPERTIES{Name: "PROPERTIES"},
		}).(*Type)
	case OBJECT:
		return any(&OBJECT{
			// Initialisation of associations
		}).(*Type)
	case PROPERTIES:
		return any(&PROPERTIES{
			// Initialisation of associations
			// field is initialized with an instance of EMBEDDEDVALUE with the name of the field
			EMBEDDEDVALUE: &EMBEDDEDVALUE{Name: "EMBEDDEDVALUE"},
		}).(*Type)
	case RELATIONGROUP:
		return any(&RELATIONGROUP{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SOURCESPECIFICATION with the name of the field
			SOURCESPECIFICATION: &SOURCESPECIFICATION{Name: "SOURCESPECIFICATION"},
			// field is initialized with an instance of SPECRELATIONS with the name of the field
			SPECRELATIONS: &SPECRELATIONS{Name: "SPECRELATIONS"},
			// field is initialized with an instance of TARGETSPECIFICATION with the name of the field
			TARGETSPECIFICATION: &TARGETSPECIFICATION{Name: "TARGETSPECIFICATION"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case RELATIONGROUPTYPE:
		return any(&RELATIONGROUPTYPE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SPECATTRIBUTES with the name of the field
			SPECATTRIBUTES: &SPECATTRIBUTES{Name: "SPECATTRIBUTES"},
		}).(*Type)
	case REQIF:
		return any(&REQIF{
			// Initialisation of associations
			// field is initialized with an instance of THEHEADER with the name of the field
			HEADER: &THEHEADER{Name: "HEADER"},
			// field is initialized with an instance of CORECONTENT with the name of the field
			CORECONTENT: &CORECONTENT{Name: "CORECONTENT"},
		}).(*Type)
	case REQIFCONTENT:
		return any(&REQIFCONTENT{
			// Initialisation of associations
			// field is initialized with an instance of DATATYPES with the name of the field
			DATATYPES: &DATATYPES{Name: "DATATYPES"},
			// field is initialized with an instance of SPECTYPES with the name of the field
			SPECTYPES: &SPECTYPES{Name: "SPECTYPES"},
			// field is initialized with an instance of SPECOBJECTS with the name of the field
			SPECOBJECTS: &SPECOBJECTS{Name: "SPECOBJECTS"},
			// field is initialized with an instance of SPECRELATIONS with the name of the field
			SPECRELATIONS: &SPECRELATIONS{Name: "SPECRELATIONS"},
			// field is initialized with an instance of SPECIFICATIONS with the name of the field
			SPECIFICATIONS: &SPECIFICATIONS{Name: "SPECIFICATIONS"},
			// field is initialized with an instance of SPECRELATIONGROUPS with the name of the field
			SPECRELATIONGROUPS: &SPECRELATIONGROUPS{Name: "SPECRELATIONGROUPS"},
		}).(*Type)
	case REQIFHEADER:
		return any(&REQIFHEADER{
			// Initialisation of associations
		}).(*Type)
	case REQIFTOOLEXTENSION:
		return any(&REQIFTOOLEXTENSION{
			// Initialisation of associations
		}).(*Type)
	case REQTYPE:
		return any(&REQTYPE{
			// Initialisation of associations
		}).(*Type)
	case SOURCE:
		return any(&SOURCE{
			// Initialisation of associations
		}).(*Type)
	case SOURCESPECIFICATION:
		return any(&SOURCESPECIFICATION{
			// Initialisation of associations
		}).(*Type)
	case SPECATTRIBUTES:
		return any(&SPECATTRIBUTES{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTEDEFINITIONBOOLEAN with the name of the field
			ATTRIBUTEDEFINITIONBOOLEAN: []*ATTRIBUTEDEFINITIONBOOLEAN{{Name: "ATTRIBUTEDEFINITIONBOOLEAN"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONDATE with the name of the field
			ATTRIBUTEDEFINITIONDATE: []*ATTRIBUTEDEFINITIONDATE{{Name: "ATTRIBUTEDEFINITIONDATE"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONENUMERATION with the name of the field
			ATTRIBUTEDEFINITIONENUMERATION: []*ATTRIBUTEDEFINITIONENUMERATION{{Name: "ATTRIBUTEDEFINITIONENUMERATION"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONINTEGER with the name of the field
			ATTRIBUTEDEFINITIONINTEGER: []*ATTRIBUTEDEFINITIONINTEGER{{Name: "ATTRIBUTEDEFINITIONINTEGER"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONREAL with the name of the field
			ATTRIBUTEDEFINITIONREAL: []*ATTRIBUTEDEFINITIONREAL{{Name: "ATTRIBUTEDEFINITIONREAL"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONSTRING with the name of the field
			ATTRIBUTEDEFINITIONSTRING: []*ATTRIBUTEDEFINITIONSTRING{{Name: "ATTRIBUTEDEFINITIONSTRING"}},
			// field is initialized with an instance of ATTRIBUTEDEFINITIONXHTML with the name of the field
			ATTRIBUTEDEFINITIONXHTML: []*ATTRIBUTEDEFINITIONXHTML{{Name: "ATTRIBUTEDEFINITIONXHTML"}},
		}).(*Type)
	case SPECHIERARCHY:
		return any(&SPECHIERARCHY{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of CHILDREN with the name of the field
			CHILDREN: &CHILDREN{Name: "CHILDREN"},
			// field is initialized with an instance of EDITABLEATTS with the name of the field
			EDITABLEATTS: &EDITABLEATTS{Name: "EDITABLEATTS"},
			// field is initialized with an instance of OBJECT with the name of the field
			OBJECT: &OBJECT{Name: "OBJECT"},
		}).(*Type)
	case SPECIFICATION:
		return any(&SPECIFICATION{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of VALUES with the name of the field
			VALUES: &VALUES{Name: "VALUES"},
			// field is initialized with an instance of CHILDREN with the name of the field
			CHILDREN: &CHILDREN{Name: "CHILDREN"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case SPECIFICATIONS:
		return any(&SPECIFICATIONS{
			// Initialisation of associations
			// field is initialized with an instance of SPECIFICATION with the name of the field
			SPECIFICATION: []*SPECIFICATION{{Name: "SPECIFICATION"}},
		}).(*Type)
	case SPECIFICATIONTYPE:
		return any(&SPECIFICATIONTYPE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SPECATTRIBUTES with the name of the field
			SPECATTRIBUTES: &SPECATTRIBUTES{Name: "SPECATTRIBUTES"},
		}).(*Type)
	case SPECIFIEDVALUES:
		return any(&SPECIFIEDVALUES{
			// Initialisation of associations
			// field is initialized with an instance of ENUMVALUE with the name of the field
			ENUMVALUE: []*ENUMVALUE{{Name: "ENUMVALUE"}},
		}).(*Type)
	case SPECOBJECT:
		return any(&SPECOBJECT{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of VALUES with the name of the field
			VALUES: &VALUES{Name: "VALUES"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case SPECOBJECTS:
		return any(&SPECOBJECTS{
			// Initialisation of associations
			// field is initialized with an instance of SPECOBJECT with the name of the field
			SPECOBJECT: []*SPECOBJECT{{Name: "SPECOBJECT"}},
		}).(*Type)
	case SPECOBJECTTYPE:
		return any(&SPECOBJECTTYPE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SPECATTRIBUTES with the name of the field
			SPECATTRIBUTES: &SPECATTRIBUTES{Name: "SPECATTRIBUTES"},
		}).(*Type)
	case SPECRELATION:
		return any(&SPECRELATION{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of VALUES with the name of the field
			VALUES: &VALUES{Name: "VALUES"},
			// field is initialized with an instance of SOURCE with the name of the field
			SOURCE: &SOURCE{Name: "SOURCE"},
			// field is initialized with an instance of TARGET with the name of the field
			TARGET: &TARGET{Name: "TARGET"},
			// field is initialized with an instance of REQTYPE with the name of the field
			TYPE: &REQTYPE{Name: "TYPE"},
		}).(*Type)
	case SPECRELATIONGROUPS:
		return any(&SPECRELATIONGROUPS{
			// Initialisation of associations
			// field is initialized with an instance of RELATIONGROUP with the name of the field
			RELATIONGROUP: []*RELATIONGROUP{{Name: "RELATIONGROUP"}},
		}).(*Type)
	case SPECRELATIONS:
		return any(&SPECRELATIONS{
			// Initialisation of associations
		}).(*Type)
	case SPECRELATIONTYPE:
		return any(&SPECRELATIONTYPE{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVEID with the name of the field
			ALTERNATIVEID: &ALTERNATIVEID{Name: "ALTERNATIVEID"},
			// field is initialized with an instance of SPECATTRIBUTES with the name of the field
			SPECATTRIBUTES: &SPECATTRIBUTES{Name: "SPECATTRIBUTES"},
		}).(*Type)
	case SPECTYPES:
		return any(&SPECTYPES{
			// Initialisation of associations
			// field is initialized with an instance of RELATIONGROUPTYPE with the name of the field
			RELATIONGROUPTYPE: []*RELATIONGROUPTYPE{{Name: "RELATIONGROUPTYPE"}},
			// field is initialized with an instance of SPECOBJECTTYPE with the name of the field
			SPECOBJECTTYPE: []*SPECOBJECTTYPE{{Name: "SPECOBJECTTYPE"}},
			// field is initialized with an instance of SPECRELATIONTYPE with the name of the field
			SPECRELATIONTYPE: []*SPECRELATIONTYPE{{Name: "SPECRELATIONTYPE"}},
			// field is initialized with an instance of SPECIFICATIONTYPE with the name of the field
			SPECIFICATIONTYPE: []*SPECIFICATIONTYPE{{Name: "SPECIFICATIONTYPE"}},
		}).(*Type)
	case TARGET:
		return any(&TARGET{
			// Initialisation of associations
		}).(*Type)
	case TARGETSPECIFICATION:
		return any(&TARGETSPECIFICATION{
			// Initialisation of associations
		}).(*Type)
	case THEHEADER:
		return any(&THEHEADER{
			// Initialisation of associations
			// field is initialized with an instance of REQIFHEADER with the name of the field
			REQIFHEADER: &REQIFHEADER{Name: "REQIFHEADER"},
		}).(*Type)
	case TOOLEXTENSIONS:
		return any(&TOOLEXTENSIONS{
			// Initialisation of associations
			// field is initialized with an instance of REQIFTOOLEXTENSION with the name of the field
			REQIFTOOLEXTENSION: []*REQIFTOOLEXTENSION{{Name: "REQIFTOOLEXTENSION"}},
		}).(*Type)
	case VALUES:
		return any(&VALUES{
			// Initialisation of associations
		}).(*Type)
	case XHTMLCONTENT:
		return any(&XHTMLCONTENT{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ALTERNATIVEID
	case ALTERNATIVEID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONBOOLEAN
	case ATTRIBUTEDEFINITIONBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONBOOLEAN)
			for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
				if attributedefinitionboolean.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitionboolean.ALTERNATIVEID
					var attributedefinitionbooleans []*ATTRIBUTEDEFINITIONBOOLEAN
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionbooleans = res[alternativeid_]
					} else {
						attributedefinitionbooleans = make([]*ATTRIBUTEDEFINITIONBOOLEAN, 0)
					}
					attributedefinitionbooleans = append(attributedefinitionbooleans, attributedefinitionboolean)
					res[alternativeid_] = attributedefinitionbooleans
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONBOOLEAN)
			for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
				if attributedefinitionboolean.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitionboolean.DEFAULTVALUE
					var attributedefinitionbooleans []*ATTRIBUTEDEFINITIONBOOLEAN
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionbooleans = res[defaultvalue_]
					} else {
						attributedefinitionbooleans = make([]*ATTRIBUTEDEFINITIONBOOLEAN, 0)
					}
					attributedefinitionbooleans = append(attributedefinitionbooleans, attributedefinitionboolean)
					res[defaultvalue_] = attributedefinitionbooleans
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONBOOLEAN)
			for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
				if attributedefinitionboolean.TYPE != nil {
					reqtype_ := attributedefinitionboolean.TYPE
					var attributedefinitionbooleans []*ATTRIBUTEDEFINITIONBOOLEAN
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionbooleans = res[reqtype_]
					} else {
						attributedefinitionbooleans = make([]*ATTRIBUTEDEFINITIONBOOLEAN, 0)
					}
					attributedefinitionbooleans = append(attributedefinitionbooleans, attributedefinitionboolean)
					res[reqtype_] = attributedefinitionbooleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONDATE
	case ATTRIBUTEDEFINITIONDATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONDATE)
			for attributedefinitiondate := range stage.ATTRIBUTEDEFINITIONDATEs {
				if attributedefinitiondate.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitiondate.ALTERNATIVEID
					var attributedefinitiondates []*ATTRIBUTEDEFINITIONDATE
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitiondates = res[alternativeid_]
					} else {
						attributedefinitiondates = make([]*ATTRIBUTEDEFINITIONDATE, 0)
					}
					attributedefinitiondates = append(attributedefinitiondates, attributedefinitiondate)
					res[alternativeid_] = attributedefinitiondates
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONDATE)
			for attributedefinitiondate := range stage.ATTRIBUTEDEFINITIONDATEs {
				if attributedefinitiondate.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitiondate.DEFAULTVALUE
					var attributedefinitiondates []*ATTRIBUTEDEFINITIONDATE
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitiondates = res[defaultvalue_]
					} else {
						attributedefinitiondates = make([]*ATTRIBUTEDEFINITIONDATE, 0)
					}
					attributedefinitiondates = append(attributedefinitiondates, attributedefinitiondate)
					res[defaultvalue_] = attributedefinitiondates
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONDATE)
			for attributedefinitiondate := range stage.ATTRIBUTEDEFINITIONDATEs {
				if attributedefinitiondate.TYPE != nil {
					reqtype_ := attributedefinitiondate.TYPE
					var attributedefinitiondates []*ATTRIBUTEDEFINITIONDATE
					_, ok := res[reqtype_]
					if ok {
						attributedefinitiondates = res[reqtype_]
					} else {
						attributedefinitiondates = make([]*ATTRIBUTEDEFINITIONDATE, 0)
					}
					attributedefinitiondates = append(attributedefinitiondates, attributedefinitiondate)
					res[reqtype_] = attributedefinitiondates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONENUMERATION
	case ATTRIBUTEDEFINITIONENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONENUMERATION)
			for attributedefinitionenumeration := range stage.ATTRIBUTEDEFINITIONENUMERATIONs {
				if attributedefinitionenumeration.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitionenumeration.DEFAULTVALUE
					var attributedefinitionenumerations []*ATTRIBUTEDEFINITIONENUMERATION
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionenumerations = res[defaultvalue_]
					} else {
						attributedefinitionenumerations = make([]*ATTRIBUTEDEFINITIONENUMERATION, 0)
					}
					attributedefinitionenumerations = append(attributedefinitionenumerations, attributedefinitionenumeration)
					res[defaultvalue_] = attributedefinitionenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONENUMERATION)
			for attributedefinitionenumeration := range stage.ATTRIBUTEDEFINITIONENUMERATIONs {
				if attributedefinitionenumeration.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitionenumeration.ALTERNATIVEID
					var attributedefinitionenumerations []*ATTRIBUTEDEFINITIONENUMERATION
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionenumerations = res[alternativeid_]
					} else {
						attributedefinitionenumerations = make([]*ATTRIBUTEDEFINITIONENUMERATION, 0)
					}
					attributedefinitionenumerations = append(attributedefinitionenumerations, attributedefinitionenumeration)
					res[alternativeid_] = attributedefinitionenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONENUMERATION)
			for attributedefinitionenumeration := range stage.ATTRIBUTEDEFINITIONENUMERATIONs {
				if attributedefinitionenumeration.TYPE != nil {
					reqtype_ := attributedefinitionenumeration.TYPE
					var attributedefinitionenumerations []*ATTRIBUTEDEFINITIONENUMERATION
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionenumerations = res[reqtype_]
					} else {
						attributedefinitionenumerations = make([]*ATTRIBUTEDEFINITIONENUMERATION, 0)
					}
					attributedefinitionenumerations = append(attributedefinitionenumerations, attributedefinitionenumeration)
					res[reqtype_] = attributedefinitionenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONINTEGER
	case ATTRIBUTEDEFINITIONINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONINTEGER)
			for attributedefinitioninteger := range stage.ATTRIBUTEDEFINITIONINTEGERs {
				if attributedefinitioninteger.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitioninteger.ALTERNATIVEID
					var attributedefinitionintegers []*ATTRIBUTEDEFINITIONINTEGER
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionintegers = res[alternativeid_]
					} else {
						attributedefinitionintegers = make([]*ATTRIBUTEDEFINITIONINTEGER, 0)
					}
					attributedefinitionintegers = append(attributedefinitionintegers, attributedefinitioninteger)
					res[alternativeid_] = attributedefinitionintegers
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONINTEGER)
			for attributedefinitioninteger := range stage.ATTRIBUTEDEFINITIONINTEGERs {
				if attributedefinitioninteger.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitioninteger.DEFAULTVALUE
					var attributedefinitionintegers []*ATTRIBUTEDEFINITIONINTEGER
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionintegers = res[defaultvalue_]
					} else {
						attributedefinitionintegers = make([]*ATTRIBUTEDEFINITIONINTEGER, 0)
					}
					attributedefinitionintegers = append(attributedefinitionintegers, attributedefinitioninteger)
					res[defaultvalue_] = attributedefinitionintegers
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONINTEGER)
			for attributedefinitioninteger := range stage.ATTRIBUTEDEFINITIONINTEGERs {
				if attributedefinitioninteger.TYPE != nil {
					reqtype_ := attributedefinitioninteger.TYPE
					var attributedefinitionintegers []*ATTRIBUTEDEFINITIONINTEGER
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionintegers = res[reqtype_]
					} else {
						attributedefinitionintegers = make([]*ATTRIBUTEDEFINITIONINTEGER, 0)
					}
					attributedefinitionintegers = append(attributedefinitionintegers, attributedefinitioninteger)
					res[reqtype_] = attributedefinitionintegers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONREAL
	case ATTRIBUTEDEFINITIONREAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONREAL)
			for attributedefinitionreal := range stage.ATTRIBUTEDEFINITIONREALs {
				if attributedefinitionreal.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitionreal.ALTERNATIVEID
					var attributedefinitionreals []*ATTRIBUTEDEFINITIONREAL
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionreals = res[alternativeid_]
					} else {
						attributedefinitionreals = make([]*ATTRIBUTEDEFINITIONREAL, 0)
					}
					attributedefinitionreals = append(attributedefinitionreals, attributedefinitionreal)
					res[alternativeid_] = attributedefinitionreals
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONREAL)
			for attributedefinitionreal := range stage.ATTRIBUTEDEFINITIONREALs {
				if attributedefinitionreal.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitionreal.DEFAULTVALUE
					var attributedefinitionreals []*ATTRIBUTEDEFINITIONREAL
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionreals = res[defaultvalue_]
					} else {
						attributedefinitionreals = make([]*ATTRIBUTEDEFINITIONREAL, 0)
					}
					attributedefinitionreals = append(attributedefinitionreals, attributedefinitionreal)
					res[defaultvalue_] = attributedefinitionreals
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONREAL)
			for attributedefinitionreal := range stage.ATTRIBUTEDEFINITIONREALs {
				if attributedefinitionreal.TYPE != nil {
					reqtype_ := attributedefinitionreal.TYPE
					var attributedefinitionreals []*ATTRIBUTEDEFINITIONREAL
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionreals = res[reqtype_]
					} else {
						attributedefinitionreals = make([]*ATTRIBUTEDEFINITIONREAL, 0)
					}
					attributedefinitionreals = append(attributedefinitionreals, attributedefinitionreal)
					res[reqtype_] = attributedefinitionreals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONSTRING
	case ATTRIBUTEDEFINITIONSTRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONSTRING)
			for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
				if attributedefinitionstring.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitionstring.ALTERNATIVEID
					var attributedefinitionstrings []*ATTRIBUTEDEFINITIONSTRING
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionstrings = res[alternativeid_]
					} else {
						attributedefinitionstrings = make([]*ATTRIBUTEDEFINITIONSTRING, 0)
					}
					attributedefinitionstrings = append(attributedefinitionstrings, attributedefinitionstring)
					res[alternativeid_] = attributedefinitionstrings
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONSTRING)
			for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
				if attributedefinitionstring.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitionstring.DEFAULTVALUE
					var attributedefinitionstrings []*ATTRIBUTEDEFINITIONSTRING
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionstrings = res[defaultvalue_]
					} else {
						attributedefinitionstrings = make([]*ATTRIBUTEDEFINITIONSTRING, 0)
					}
					attributedefinitionstrings = append(attributedefinitionstrings, attributedefinitionstring)
					res[defaultvalue_] = attributedefinitionstrings
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONSTRING)
			for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
				if attributedefinitionstring.TYPE != nil {
					reqtype_ := attributedefinitionstring.TYPE
					var attributedefinitionstrings []*ATTRIBUTEDEFINITIONSTRING
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionstrings = res[reqtype_]
					} else {
						attributedefinitionstrings = make([]*ATTRIBUTEDEFINITIONSTRING, 0)
					}
					attributedefinitionstrings = append(attributedefinitionstrings, attributedefinitionstring)
					res[reqtype_] = attributedefinitionstrings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONXHTML
	case ATTRIBUTEDEFINITIONXHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ATTRIBUTEDEFINITIONXHTML)
			for attributedefinitionxhtml := range stage.ATTRIBUTEDEFINITIONXHTMLs {
				if attributedefinitionxhtml.ALTERNATIVEID != nil {
					alternativeid_ := attributedefinitionxhtml.ALTERNATIVEID
					var attributedefinitionxhtmls []*ATTRIBUTEDEFINITIONXHTML
					_, ok := res[alternativeid_]
					if ok {
						attributedefinitionxhtmls = res[alternativeid_]
					} else {
						attributedefinitionxhtmls = make([]*ATTRIBUTEDEFINITIONXHTML, 0)
					}
					attributedefinitionxhtmls = append(attributedefinitionxhtmls, attributedefinitionxhtml)
					res[alternativeid_] = attributedefinitionxhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULTVALUE":
			res := make(map[*DEFAULTVALUE][]*ATTRIBUTEDEFINITIONXHTML)
			for attributedefinitionxhtml := range stage.ATTRIBUTEDEFINITIONXHTMLs {
				if attributedefinitionxhtml.DEFAULTVALUE != nil {
					defaultvalue_ := attributedefinitionxhtml.DEFAULTVALUE
					var attributedefinitionxhtmls []*ATTRIBUTEDEFINITIONXHTML
					_, ok := res[defaultvalue_]
					if ok {
						attributedefinitionxhtmls = res[defaultvalue_]
					} else {
						attributedefinitionxhtmls = make([]*ATTRIBUTEDEFINITIONXHTML, 0)
					}
					attributedefinitionxhtmls = append(attributedefinitionxhtmls, attributedefinitionxhtml)
					res[defaultvalue_] = attributedefinitionxhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*ATTRIBUTEDEFINITIONXHTML)
			for attributedefinitionxhtml := range stage.ATTRIBUTEDEFINITIONXHTMLs {
				if attributedefinitionxhtml.TYPE != nil {
					reqtype_ := attributedefinitionxhtml.TYPE
					var attributedefinitionxhtmls []*ATTRIBUTEDEFINITIONXHTML
					_, ok := res[reqtype_]
					if ok {
						attributedefinitionxhtmls = res[reqtype_]
					} else {
						attributedefinitionxhtmls = make([]*ATTRIBUTEDEFINITIONXHTML, 0)
					}
					attributedefinitionxhtmls = append(attributedefinitionxhtmls, attributedefinitionxhtml)
					res[reqtype_] = attributedefinitionxhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEBOOLEAN
	case ATTRIBUTEVALUEBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEBOOLEAN)
			for attributevalueboolean := range stage.ATTRIBUTEVALUEBOOLEANs {
				if attributevalueboolean.DEFINITION != nil {
					definition_ := attributevalueboolean.DEFINITION
					var attributevaluebooleans []*ATTRIBUTEVALUEBOOLEAN
					_, ok := res[definition_]
					if ok {
						attributevaluebooleans = res[definition_]
					} else {
						attributevaluebooleans = make([]*ATTRIBUTEVALUEBOOLEAN, 0)
					}
					attributevaluebooleans = append(attributevaluebooleans, attributevalueboolean)
					res[definition_] = attributevaluebooleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEDATE
	case ATTRIBUTEVALUEDATE:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEDATE)
			for attributevaluedate := range stage.ATTRIBUTEVALUEDATEs {
				if attributevaluedate.DEFINITION != nil {
					definition_ := attributevaluedate.DEFINITION
					var attributevaluedates []*ATTRIBUTEVALUEDATE
					_, ok := res[definition_]
					if ok {
						attributevaluedates = res[definition_]
					} else {
						attributevaluedates = make([]*ATTRIBUTEVALUEDATE, 0)
					}
					attributevaluedates = append(attributevaluedates, attributevaluedate)
					res[definition_] = attributevaluedates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEENUMERATION
	case ATTRIBUTEVALUEENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEENUMERATION)
			for attributevalueenumeration := range stage.ATTRIBUTEVALUEENUMERATIONs {
				if attributevalueenumeration.DEFINITION != nil {
					definition_ := attributevalueenumeration.DEFINITION
					var attributevalueenumerations []*ATTRIBUTEVALUEENUMERATION
					_, ok := res[definition_]
					if ok {
						attributevalueenumerations = res[definition_]
					} else {
						attributevalueenumerations = make([]*ATTRIBUTEVALUEENUMERATION, 0)
					}
					attributevalueenumerations = append(attributevalueenumerations, attributevalueenumeration)
					res[definition_] = attributevalueenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*VALUES][]*ATTRIBUTEVALUEENUMERATION)
			for attributevalueenumeration := range stage.ATTRIBUTEVALUEENUMERATIONs {
				if attributevalueenumeration.VALUES != nil {
					values_ := attributevalueenumeration.VALUES
					var attributevalueenumerations []*ATTRIBUTEVALUEENUMERATION
					_, ok := res[values_]
					if ok {
						attributevalueenumerations = res[values_]
					} else {
						attributevalueenumerations = make([]*ATTRIBUTEVALUEENUMERATION, 0)
					}
					attributevalueenumerations = append(attributevalueenumerations, attributevalueenumeration)
					res[values_] = attributevalueenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEINTEGER
	case ATTRIBUTEVALUEINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEINTEGER)
			for attributevalueinteger := range stage.ATTRIBUTEVALUEINTEGERs {
				if attributevalueinteger.DEFINITION != nil {
					definition_ := attributevalueinteger.DEFINITION
					var attributevalueintegers []*ATTRIBUTEVALUEINTEGER
					_, ok := res[definition_]
					if ok {
						attributevalueintegers = res[definition_]
					} else {
						attributevalueintegers = make([]*ATTRIBUTEVALUEINTEGER, 0)
					}
					attributevalueintegers = append(attributevalueintegers, attributevalueinteger)
					res[definition_] = attributevalueintegers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEREAL
	case ATTRIBUTEVALUEREAL:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEREAL)
			for attributevaluereal := range stage.ATTRIBUTEVALUEREALs {
				if attributevaluereal.DEFINITION != nil {
					definition_ := attributevaluereal.DEFINITION
					var attributevaluereals []*ATTRIBUTEVALUEREAL
					_, ok := res[definition_]
					if ok {
						attributevaluereals = res[definition_]
					} else {
						attributevaluereals = make([]*ATTRIBUTEVALUEREAL, 0)
					}
					attributevaluereals = append(attributevaluereals, attributevaluereal)
					res[definition_] = attributevaluereals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUESTRING
	case ATTRIBUTEVALUESTRING:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUESTRING)
			for attributevaluestring := range stage.ATTRIBUTEVALUESTRINGs {
				if attributevaluestring.DEFINITION != nil {
					definition_ := attributevaluestring.DEFINITION
					var attributevaluestrings []*ATTRIBUTEVALUESTRING
					_, ok := res[definition_]
					if ok {
						attributevaluestrings = res[definition_]
					} else {
						attributevaluestrings = make([]*ATTRIBUTEVALUESTRING, 0)
					}
					attributevaluestrings = append(attributevaluestrings, attributevaluestring)
					res[definition_] = attributevaluestrings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEXHTML
	case ATTRIBUTEVALUEXHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "THEVALUE":
			res := make(map[*XHTMLCONTENT][]*ATTRIBUTEVALUEXHTML)
			for attributevaluexhtml := range stage.ATTRIBUTEVALUEXHTMLs {
				if attributevaluexhtml.THEVALUE != nil {
					xhtmlcontent_ := attributevaluexhtml.THEVALUE
					var attributevaluexhtmls []*ATTRIBUTEVALUEXHTML
					_, ok := res[xhtmlcontent_]
					if ok {
						attributevaluexhtmls = res[xhtmlcontent_]
					} else {
						attributevaluexhtmls = make([]*ATTRIBUTEVALUEXHTML, 0)
					}
					attributevaluexhtmls = append(attributevaluexhtmls, attributevaluexhtml)
					res[xhtmlcontent_] = attributevaluexhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "THEORIGINALVALUE":
			res := make(map[*XHTMLCONTENT][]*ATTRIBUTEVALUEXHTML)
			for attributevaluexhtml := range stage.ATTRIBUTEVALUEXHTMLs {
				if attributevaluexhtml.THEORIGINALVALUE != nil {
					xhtmlcontent_ := attributevaluexhtml.THEORIGINALVALUE
					var attributevaluexhtmls []*ATTRIBUTEVALUEXHTML
					_, ok := res[xhtmlcontent_]
					if ok {
						attributevaluexhtmls = res[xhtmlcontent_]
					} else {
						attributevaluexhtmls = make([]*ATTRIBUTEVALUEXHTML, 0)
					}
					attributevaluexhtmls = append(attributevaluexhtmls, attributevaluexhtml)
					res[xhtmlcontent_] = attributevaluexhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFINITION":
			res := make(map[*DEFINITION][]*ATTRIBUTEVALUEXHTML)
			for attributevaluexhtml := range stage.ATTRIBUTEVALUEXHTMLs {
				if attributevaluexhtml.DEFINITION != nil {
					definition_ := attributevaluexhtml.DEFINITION
					var attributevaluexhtmls []*ATTRIBUTEVALUEXHTML
					_, ok := res[definition_]
					if ok {
						attributevaluexhtmls = res[definition_]
					} else {
						attributevaluexhtmls = make([]*ATTRIBUTEVALUEXHTML, 0)
					}
					attributevaluexhtmls = append(attributevaluexhtmls, attributevaluexhtml)
					res[definition_] = attributevaluexhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CHILDREN
	case CHILDREN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CORECONTENT
	case CORECONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "REQIFCONTENT":
			res := make(map[*REQIFCONTENT][]*CORECONTENT)
			for corecontent := range stage.CORECONTENTs {
				if corecontent.REQIFCONTENT != nil {
					reqifcontent_ := corecontent.REQIFCONTENT
					var corecontents []*CORECONTENT
					_, ok := res[reqifcontent_]
					if ok {
						corecontents = res[reqifcontent_]
					} else {
						corecontents = make([]*CORECONTENT, 0)
					}
					corecontents = append(corecontents, corecontent)
					res[reqifcontent_] = corecontents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONBOOLEAN
	case DATATYPEDEFINITIONBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONBOOLEAN)
			for datatypedefinitionboolean := range stage.DATATYPEDEFINITIONBOOLEANs {
				if datatypedefinitionboolean.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitionboolean.ALTERNATIVEID
					var datatypedefinitionbooleans []*DATATYPEDEFINITIONBOOLEAN
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionbooleans = res[alternativeid_]
					} else {
						datatypedefinitionbooleans = make([]*DATATYPEDEFINITIONBOOLEAN, 0)
					}
					datatypedefinitionbooleans = append(datatypedefinitionbooleans, datatypedefinitionboolean)
					res[alternativeid_] = datatypedefinitionbooleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONDATE
	case DATATYPEDEFINITIONDATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONDATE)
			for datatypedefinitiondate := range stage.DATATYPEDEFINITIONDATEs {
				if datatypedefinitiondate.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitiondate.ALTERNATIVEID
					var datatypedefinitiondates []*DATATYPEDEFINITIONDATE
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitiondates = res[alternativeid_]
					} else {
						datatypedefinitiondates = make([]*DATATYPEDEFINITIONDATE, 0)
					}
					datatypedefinitiondates = append(datatypedefinitiondates, datatypedefinitiondate)
					res[alternativeid_] = datatypedefinitiondates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONENUMERATION
	case DATATYPEDEFINITIONENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONENUMERATION)
			for datatypedefinitionenumeration := range stage.DATATYPEDEFINITIONENUMERATIONs {
				if datatypedefinitionenumeration.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitionenumeration.ALTERNATIVEID
					var datatypedefinitionenumerations []*DATATYPEDEFINITIONENUMERATION
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionenumerations = res[alternativeid_]
					} else {
						datatypedefinitionenumerations = make([]*DATATYPEDEFINITIONENUMERATION, 0)
					}
					datatypedefinitionenumerations = append(datatypedefinitionenumerations, datatypedefinitionenumeration)
					res[alternativeid_] = datatypedefinitionenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFIEDVALUES":
			res := make(map[*SPECIFIEDVALUES][]*DATATYPEDEFINITIONENUMERATION)
			for datatypedefinitionenumeration := range stage.DATATYPEDEFINITIONENUMERATIONs {
				if datatypedefinitionenumeration.SPECIFIEDVALUES != nil {
					specifiedvalues_ := datatypedefinitionenumeration.SPECIFIEDVALUES
					var datatypedefinitionenumerations []*DATATYPEDEFINITIONENUMERATION
					_, ok := res[specifiedvalues_]
					if ok {
						datatypedefinitionenumerations = res[specifiedvalues_]
					} else {
						datatypedefinitionenumerations = make([]*DATATYPEDEFINITIONENUMERATION, 0)
					}
					datatypedefinitionenumerations = append(datatypedefinitionenumerations, datatypedefinitionenumeration)
					res[specifiedvalues_] = datatypedefinitionenumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONINTEGER
	case DATATYPEDEFINITIONINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONINTEGER)
			for datatypedefinitioninteger := range stage.DATATYPEDEFINITIONINTEGERs {
				if datatypedefinitioninteger.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitioninteger.ALTERNATIVEID
					var datatypedefinitionintegers []*DATATYPEDEFINITIONINTEGER
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionintegers = res[alternativeid_]
					} else {
						datatypedefinitionintegers = make([]*DATATYPEDEFINITIONINTEGER, 0)
					}
					datatypedefinitionintegers = append(datatypedefinitionintegers, datatypedefinitioninteger)
					res[alternativeid_] = datatypedefinitionintegers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONREAL
	case DATATYPEDEFINITIONREAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONREAL)
			for datatypedefinitionreal := range stage.DATATYPEDEFINITIONREALs {
				if datatypedefinitionreal.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitionreal.ALTERNATIVEID
					var datatypedefinitionreals []*DATATYPEDEFINITIONREAL
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionreals = res[alternativeid_]
					} else {
						datatypedefinitionreals = make([]*DATATYPEDEFINITIONREAL, 0)
					}
					datatypedefinitionreals = append(datatypedefinitionreals, datatypedefinitionreal)
					res[alternativeid_] = datatypedefinitionreals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONSTRING
	case DATATYPEDEFINITIONSTRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONSTRING)
			for datatypedefinitionstring := range stage.DATATYPEDEFINITIONSTRINGs {
				if datatypedefinitionstring.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitionstring.ALTERNATIVEID
					var datatypedefinitionstrings []*DATATYPEDEFINITIONSTRING
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionstrings = res[alternativeid_]
					} else {
						datatypedefinitionstrings = make([]*DATATYPEDEFINITIONSTRING, 0)
					}
					datatypedefinitionstrings = append(datatypedefinitionstrings, datatypedefinitionstring)
					res[alternativeid_] = datatypedefinitionstrings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONXHTML
	case DATATYPEDEFINITIONXHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*DATATYPEDEFINITIONXHTML)
			for datatypedefinitionxhtml := range stage.DATATYPEDEFINITIONXHTMLs {
				if datatypedefinitionxhtml.ALTERNATIVEID != nil {
					alternativeid_ := datatypedefinitionxhtml.ALTERNATIVEID
					var datatypedefinitionxhtmls []*DATATYPEDEFINITIONXHTML
					_, ok := res[alternativeid_]
					if ok {
						datatypedefinitionxhtmls = res[alternativeid_]
					} else {
						datatypedefinitionxhtmls = make([]*DATATYPEDEFINITIONXHTML, 0)
					}
					datatypedefinitionxhtmls = append(datatypedefinitionxhtmls, datatypedefinitionxhtml)
					res[alternativeid_] = datatypedefinitionxhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPES
	case DATATYPES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DEFAULTVALUE
	case DEFAULTVALUE:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTEVALUEBOOLEAN":
			res := make(map[*ATTRIBUTEVALUEBOOLEAN][]*DEFAULTVALUE)
			for defaultvalue := range stage.DEFAULTVALUEs {
				if defaultvalue.ATTRIBUTEVALUEBOOLEAN != nil {
					attributevalueboolean_ := defaultvalue.ATTRIBUTEVALUEBOOLEAN
					var defaultvalues []*DEFAULTVALUE
					_, ok := res[attributevalueboolean_]
					if ok {
						defaultvalues = res[attributevalueboolean_]
					} else {
						defaultvalues = make([]*DEFAULTVALUE, 0)
					}
					defaultvalues = append(defaultvalues, defaultvalue)
					res[attributevalueboolean_] = defaultvalues
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DEFINITION
	case DEFINITION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EDITABLEATTS
	case EDITABLEATTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EMBEDDEDVALUE
	case EMBEDDEDVALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUMVALUE
	case ENUMVALUE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*ENUMVALUE)
			for enumvalue := range stage.ENUMVALUEs {
				if enumvalue.ALTERNATIVEID != nil {
					alternativeid_ := enumvalue.ALTERNATIVEID
					var enumvalues []*ENUMVALUE
					_, ok := res[alternativeid_]
					if ok {
						enumvalues = res[alternativeid_]
					} else {
						enumvalues = make([]*ENUMVALUE, 0)
					}
					enumvalues = append(enumvalues, enumvalue)
					res[alternativeid_] = enumvalues
				}
			}
			return any(res).(map[*End][]*Start)
		case "PROPERTIES":
			res := make(map[*PROPERTIES][]*ENUMVALUE)
			for enumvalue := range stage.ENUMVALUEs {
				if enumvalue.PROPERTIES != nil {
					properties_ := enumvalue.PROPERTIES
					var enumvalues []*ENUMVALUE
					_, ok := res[properties_]
					if ok {
						enumvalues = res[properties_]
					} else {
						enumvalues = make([]*ENUMVALUE, 0)
					}
					enumvalues = append(enumvalues, enumvalue)
					res[properties_] = enumvalues
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of OBJECT
	case OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PROPERTIES
	case PROPERTIES:
		switch fieldname {
		// insertion point for per direct association field
		case "EMBEDDEDVALUE":
			res := make(map[*EMBEDDEDVALUE][]*PROPERTIES)
			for properties := range stage.PROPERTIESs {
				if properties.EMBEDDEDVALUE != nil {
					embeddedvalue_ := properties.EMBEDDEDVALUE
					var propertiess []*PROPERTIES
					_, ok := res[embeddedvalue_]
					if ok {
						propertiess = res[embeddedvalue_]
					} else {
						propertiess = make([]*PROPERTIES, 0)
					}
					propertiess = append(propertiess, properties)
					res[embeddedvalue_] = propertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RELATIONGROUP
	case RELATIONGROUP:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*RELATIONGROUP)
			for relationgroup := range stage.RELATIONGROUPs {
				if relationgroup.ALTERNATIVEID != nil {
					alternativeid_ := relationgroup.ALTERNATIVEID
					var relationgroups []*RELATIONGROUP
					_, ok := res[alternativeid_]
					if ok {
						relationgroups = res[alternativeid_]
					} else {
						relationgroups = make([]*RELATIONGROUP, 0)
					}
					relationgroups = append(relationgroups, relationgroup)
					res[alternativeid_] = relationgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "SOURCESPECIFICATION":
			res := make(map[*SOURCESPECIFICATION][]*RELATIONGROUP)
			for relationgroup := range stage.RELATIONGROUPs {
				if relationgroup.SOURCESPECIFICATION != nil {
					sourcespecification_ := relationgroup.SOURCESPECIFICATION
					var relationgroups []*RELATIONGROUP
					_, ok := res[sourcespecification_]
					if ok {
						relationgroups = res[sourcespecification_]
					} else {
						relationgroups = make([]*RELATIONGROUP, 0)
					}
					relationgroups = append(relationgroups, relationgroup)
					res[sourcespecification_] = relationgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECRELATIONS":
			res := make(map[*SPECRELATIONS][]*RELATIONGROUP)
			for relationgroup := range stage.RELATIONGROUPs {
				if relationgroup.SPECRELATIONS != nil {
					specrelations_ := relationgroup.SPECRELATIONS
					var relationgroups []*RELATIONGROUP
					_, ok := res[specrelations_]
					if ok {
						relationgroups = res[specrelations_]
					} else {
						relationgroups = make([]*RELATIONGROUP, 0)
					}
					relationgroups = append(relationgroups, relationgroup)
					res[specrelations_] = relationgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "TARGETSPECIFICATION":
			res := make(map[*TARGETSPECIFICATION][]*RELATIONGROUP)
			for relationgroup := range stage.RELATIONGROUPs {
				if relationgroup.TARGETSPECIFICATION != nil {
					targetspecification_ := relationgroup.TARGETSPECIFICATION
					var relationgroups []*RELATIONGROUP
					_, ok := res[targetspecification_]
					if ok {
						relationgroups = res[targetspecification_]
					} else {
						relationgroups = make([]*RELATIONGROUP, 0)
					}
					relationgroups = append(relationgroups, relationgroup)
					res[targetspecification_] = relationgroups
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*RELATIONGROUP)
			for relationgroup := range stage.RELATIONGROUPs {
				if relationgroup.TYPE != nil {
					reqtype_ := relationgroup.TYPE
					var relationgroups []*RELATIONGROUP
					_, ok := res[reqtype_]
					if ok {
						relationgroups = res[reqtype_]
					} else {
						relationgroups = make([]*RELATIONGROUP, 0)
					}
					relationgroups = append(relationgroups, relationgroup)
					res[reqtype_] = relationgroups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RELATIONGROUPTYPE
	case RELATIONGROUPTYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*RELATIONGROUPTYPE)
			for relationgrouptype := range stage.RELATIONGROUPTYPEs {
				if relationgrouptype.ALTERNATIVEID != nil {
					alternativeid_ := relationgrouptype.ALTERNATIVEID
					var relationgrouptypes []*RELATIONGROUPTYPE
					_, ok := res[alternativeid_]
					if ok {
						relationgrouptypes = res[alternativeid_]
					} else {
						relationgrouptypes = make([]*RELATIONGROUPTYPE, 0)
					}
					relationgrouptypes = append(relationgrouptypes, relationgrouptype)
					res[alternativeid_] = relationgrouptypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECATTRIBUTES":
			res := make(map[*SPECATTRIBUTES][]*RELATIONGROUPTYPE)
			for relationgrouptype := range stage.RELATIONGROUPTYPEs {
				if relationgrouptype.SPECATTRIBUTES != nil {
					specattributes_ := relationgrouptype.SPECATTRIBUTES
					var relationgrouptypes []*RELATIONGROUPTYPE
					_, ok := res[specattributes_]
					if ok {
						relationgrouptypes = res[specattributes_]
					} else {
						relationgrouptypes = make([]*RELATIONGROUPTYPE, 0)
					}
					relationgrouptypes = append(relationgrouptypes, relationgrouptype)
					res[specattributes_] = relationgrouptypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		case "HEADER":
			res := make(map[*THEHEADER][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.HEADER != nil {
					theheader_ := reqif.HEADER
					var reqifs []*REQIF
					_, ok := res[theheader_]
					if ok {
						reqifs = res[theheader_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[theheader_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "CORECONTENT":
			res := make(map[*CORECONTENT][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.CORECONTENT != nil {
					corecontent_ := reqif.CORECONTENT
					var reqifs []*REQIF
					_, ok := res[corecontent_]
					if ok {
						reqifs = res[corecontent_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[corecontent_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQIFCONTENT
	case REQIFCONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "DATATYPES":
			res := make(map[*DATATYPES][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.DATATYPES != nil {
					datatypes_ := reqifcontent.DATATYPES
					var reqifcontents []*REQIFCONTENT
					_, ok := res[datatypes_]
					if ok {
						reqifcontents = res[datatypes_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[datatypes_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECTYPES":
			res := make(map[*SPECTYPES][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.SPECTYPES != nil {
					spectypes_ := reqifcontent.SPECTYPES
					var reqifcontents []*REQIFCONTENT
					_, ok := res[spectypes_]
					if ok {
						reqifcontents = res[spectypes_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[spectypes_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECOBJECTS":
			res := make(map[*SPECOBJECTS][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.SPECOBJECTS != nil {
					specobjects_ := reqifcontent.SPECOBJECTS
					var reqifcontents []*REQIFCONTENT
					_, ok := res[specobjects_]
					if ok {
						reqifcontents = res[specobjects_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[specobjects_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECRELATIONS":
			res := make(map[*SPECRELATIONS][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.SPECRELATIONS != nil {
					specrelations_ := reqifcontent.SPECRELATIONS
					var reqifcontents []*REQIFCONTENT
					_, ok := res[specrelations_]
					if ok {
						reqifcontents = res[specrelations_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[specrelations_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFICATIONS":
			res := make(map[*SPECIFICATIONS][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.SPECIFICATIONS != nil {
					specifications_ := reqifcontent.SPECIFICATIONS
					var reqifcontents []*REQIFCONTENT
					_, ok := res[specifications_]
					if ok {
						reqifcontents = res[specifications_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[specifications_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECRELATIONGROUPS":
			res := make(map[*SPECRELATIONGROUPS][]*REQIFCONTENT)
			for reqifcontent := range stage.REQIFCONTENTs {
				if reqifcontent.SPECRELATIONGROUPS != nil {
					specrelationgroups_ := reqifcontent.SPECRELATIONGROUPS
					var reqifcontents []*REQIFCONTENT
					_, ok := res[specrelationgroups_]
					if ok {
						reqifcontents = res[specrelationgroups_]
					} else {
						reqifcontents = make([]*REQIFCONTENT, 0)
					}
					reqifcontents = append(reqifcontents, reqifcontent)
					res[specrelationgroups_] = reqifcontents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQIFHEADER
	case REQIFHEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIFTOOLEXTENSION
	case REQIFTOOLEXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQTYPE
	case REQTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SOURCE
	case SOURCE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SOURCESPECIFICATION
	case SOURCESPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECATTRIBUTES
	case SPECATTRIBUTES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECHIERARCHY
	case SPECHIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECHIERARCHY)
			for spechierarchy := range stage.SPECHIERARCHYs {
				if spechierarchy.ALTERNATIVEID != nil {
					alternativeid_ := spechierarchy.ALTERNATIVEID
					var spechierarchys []*SPECHIERARCHY
					_, ok := res[alternativeid_]
					if ok {
						spechierarchys = res[alternativeid_]
					} else {
						spechierarchys = make([]*SPECHIERARCHY, 0)
					}
					spechierarchys = append(spechierarchys, spechierarchy)
					res[alternativeid_] = spechierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "CHILDREN":
			res := make(map[*CHILDREN][]*SPECHIERARCHY)
			for spechierarchy := range stage.SPECHIERARCHYs {
				if spechierarchy.CHILDREN != nil {
					children_ := spechierarchy.CHILDREN
					var spechierarchys []*SPECHIERARCHY
					_, ok := res[children_]
					if ok {
						spechierarchys = res[children_]
					} else {
						spechierarchys = make([]*SPECHIERARCHY, 0)
					}
					spechierarchys = append(spechierarchys, spechierarchy)
					res[children_] = spechierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "EDITABLEATTS":
			res := make(map[*EDITABLEATTS][]*SPECHIERARCHY)
			for spechierarchy := range stage.SPECHIERARCHYs {
				if spechierarchy.EDITABLEATTS != nil {
					editableatts_ := spechierarchy.EDITABLEATTS
					var spechierarchys []*SPECHIERARCHY
					_, ok := res[editableatts_]
					if ok {
						spechierarchys = res[editableatts_]
					} else {
						spechierarchys = make([]*SPECHIERARCHY, 0)
					}
					spechierarchys = append(spechierarchys, spechierarchy)
					res[editableatts_] = spechierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "OBJECT":
			res := make(map[*OBJECT][]*SPECHIERARCHY)
			for spechierarchy := range stage.SPECHIERARCHYs {
				if spechierarchy.OBJECT != nil {
					object_ := spechierarchy.OBJECT
					var spechierarchys []*SPECHIERARCHY
					_, ok := res[object_]
					if ok {
						spechierarchys = res[object_]
					} else {
						spechierarchys = make([]*SPECHIERARCHY, 0)
					}
					spechierarchys = append(spechierarchys, spechierarchy)
					res[object_] = spechierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.ALTERNATIVEID != nil {
					alternativeid_ := specification.ALTERNATIVEID
					var specifications []*SPECIFICATION
					_, ok := res[alternativeid_]
					if ok {
						specifications = res[alternativeid_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[alternativeid_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*VALUES][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.VALUES != nil {
					values_ := specification.VALUES
					var specifications []*SPECIFICATION
					_, ok := res[values_]
					if ok {
						specifications = res[values_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[values_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "CHILDREN":
			res := make(map[*CHILDREN][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.CHILDREN != nil {
					children_ := specification.CHILDREN
					var specifications []*SPECIFICATION
					_, ok := res[children_]
					if ok {
						specifications = res[children_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[children_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.TYPE != nil {
					reqtype_ := specification.TYPE
					var specifications []*SPECIFICATION
					_, ok := res[reqtype_]
					if ok {
						specifications = res[reqtype_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[reqtype_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECIFICATIONS
	case SPECIFICATIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATIONTYPE
	case SPECIFICATIONTYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECIFICATIONTYPE)
			for specificationtype := range stage.SPECIFICATIONTYPEs {
				if specificationtype.ALTERNATIVEID != nil {
					alternativeid_ := specificationtype.ALTERNATIVEID
					var specificationtypes []*SPECIFICATIONTYPE
					_, ok := res[alternativeid_]
					if ok {
						specificationtypes = res[alternativeid_]
					} else {
						specificationtypes = make([]*SPECIFICATIONTYPE, 0)
					}
					specificationtypes = append(specificationtypes, specificationtype)
					res[alternativeid_] = specificationtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECATTRIBUTES":
			res := make(map[*SPECATTRIBUTES][]*SPECIFICATIONTYPE)
			for specificationtype := range stage.SPECIFICATIONTYPEs {
				if specificationtype.SPECATTRIBUTES != nil {
					specattributes_ := specificationtype.SPECATTRIBUTES
					var specificationtypes []*SPECIFICATIONTYPE
					_, ok := res[specattributes_]
					if ok {
						specificationtypes = res[specattributes_]
					} else {
						specificationtypes = make([]*SPECIFICATIONTYPE, 0)
					}
					specificationtypes = append(specificationtypes, specificationtype)
					res[specattributes_] = specificationtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECIFIEDVALUES
	case SPECIFIEDVALUES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECOBJECT
	case SPECOBJECT:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECOBJECT)
			for specobject := range stage.SPECOBJECTs {
				if specobject.ALTERNATIVEID != nil {
					alternativeid_ := specobject.ALTERNATIVEID
					var specobjects []*SPECOBJECT
					_, ok := res[alternativeid_]
					if ok {
						specobjects = res[alternativeid_]
					} else {
						specobjects = make([]*SPECOBJECT, 0)
					}
					specobjects = append(specobjects, specobject)
					res[alternativeid_] = specobjects
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*VALUES][]*SPECOBJECT)
			for specobject := range stage.SPECOBJECTs {
				if specobject.VALUES != nil {
					values_ := specobject.VALUES
					var specobjects []*SPECOBJECT
					_, ok := res[values_]
					if ok {
						specobjects = res[values_]
					} else {
						specobjects = make([]*SPECOBJECT, 0)
					}
					specobjects = append(specobjects, specobject)
					res[values_] = specobjects
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*SPECOBJECT)
			for specobject := range stage.SPECOBJECTs {
				if specobject.TYPE != nil {
					reqtype_ := specobject.TYPE
					var specobjects []*SPECOBJECT
					_, ok := res[reqtype_]
					if ok {
						specobjects = res[reqtype_]
					} else {
						specobjects = make([]*SPECOBJECT, 0)
					}
					specobjects = append(specobjects, specobject)
					res[reqtype_] = specobjects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECOBJECTS
	case SPECOBJECTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECOBJECTTYPE
	case SPECOBJECTTYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECOBJECTTYPE)
			for specobjecttype := range stage.SPECOBJECTTYPEs {
				if specobjecttype.ALTERNATIVEID != nil {
					alternativeid_ := specobjecttype.ALTERNATIVEID
					var specobjecttypes []*SPECOBJECTTYPE
					_, ok := res[alternativeid_]
					if ok {
						specobjecttypes = res[alternativeid_]
					} else {
						specobjecttypes = make([]*SPECOBJECTTYPE, 0)
					}
					specobjecttypes = append(specobjecttypes, specobjecttype)
					res[alternativeid_] = specobjecttypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECATTRIBUTES":
			res := make(map[*SPECATTRIBUTES][]*SPECOBJECTTYPE)
			for specobjecttype := range stage.SPECOBJECTTYPEs {
				if specobjecttype.SPECATTRIBUTES != nil {
					specattributes_ := specobjecttype.SPECATTRIBUTES
					var specobjecttypes []*SPECOBJECTTYPE
					_, ok := res[specattributes_]
					if ok {
						specobjecttypes = res[specattributes_]
					} else {
						specobjecttypes = make([]*SPECOBJECTTYPE, 0)
					}
					specobjecttypes = append(specobjecttypes, specobjecttype)
					res[specattributes_] = specobjecttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECRELATION
	case SPECRELATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECRELATION)
			for specrelation := range stage.SPECRELATIONs {
				if specrelation.ALTERNATIVEID != nil {
					alternativeid_ := specrelation.ALTERNATIVEID
					var specrelations []*SPECRELATION
					_, ok := res[alternativeid_]
					if ok {
						specrelations = res[alternativeid_]
					} else {
						specrelations = make([]*SPECRELATION, 0)
					}
					specrelations = append(specrelations, specrelation)
					res[alternativeid_] = specrelations
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*VALUES][]*SPECRELATION)
			for specrelation := range stage.SPECRELATIONs {
				if specrelation.VALUES != nil {
					values_ := specrelation.VALUES
					var specrelations []*SPECRELATION
					_, ok := res[values_]
					if ok {
						specrelations = res[values_]
					} else {
						specrelations = make([]*SPECRELATION, 0)
					}
					specrelations = append(specrelations, specrelation)
					res[values_] = specrelations
				}
			}
			return any(res).(map[*End][]*Start)
		case "SOURCE":
			res := make(map[*SOURCE][]*SPECRELATION)
			for specrelation := range stage.SPECRELATIONs {
				if specrelation.SOURCE != nil {
					source_ := specrelation.SOURCE
					var specrelations []*SPECRELATION
					_, ok := res[source_]
					if ok {
						specrelations = res[source_]
					} else {
						specrelations = make([]*SPECRELATION, 0)
					}
					specrelations = append(specrelations, specrelation)
					res[source_] = specrelations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TARGET":
			res := make(map[*TARGET][]*SPECRELATION)
			for specrelation := range stage.SPECRELATIONs {
				if specrelation.TARGET != nil {
					target_ := specrelation.TARGET
					var specrelations []*SPECRELATION
					_, ok := res[target_]
					if ok {
						specrelations = res[target_]
					} else {
						specrelations = make([]*SPECRELATION, 0)
					}
					specrelations = append(specrelations, specrelation)
					res[target_] = specrelations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*REQTYPE][]*SPECRELATION)
			for specrelation := range stage.SPECRELATIONs {
				if specrelation.TYPE != nil {
					reqtype_ := specrelation.TYPE
					var specrelations []*SPECRELATION
					_, ok := res[reqtype_]
					if ok {
						specrelations = res[reqtype_]
					} else {
						specrelations = make([]*SPECRELATION, 0)
					}
					specrelations = append(specrelations, specrelation)
					res[reqtype_] = specrelations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECRELATIONGROUPS
	case SPECRELATIONGROUPS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECRELATIONS
	case SPECRELATIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECRELATIONTYPE
	case SPECRELATIONTYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVEID":
			res := make(map[*ALTERNATIVEID][]*SPECRELATIONTYPE)
			for specrelationtype := range stage.SPECRELATIONTYPEs {
				if specrelationtype.ALTERNATIVEID != nil {
					alternativeid_ := specrelationtype.ALTERNATIVEID
					var specrelationtypes []*SPECRELATIONTYPE
					_, ok := res[alternativeid_]
					if ok {
						specrelationtypes = res[alternativeid_]
					} else {
						specrelationtypes = make([]*SPECRELATIONTYPE, 0)
					}
					specrelationtypes = append(specrelationtypes, specrelationtype)
					res[alternativeid_] = specrelationtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECATTRIBUTES":
			res := make(map[*SPECATTRIBUTES][]*SPECRELATIONTYPE)
			for specrelationtype := range stage.SPECRELATIONTYPEs {
				if specrelationtype.SPECATTRIBUTES != nil {
					specattributes_ := specrelationtype.SPECATTRIBUTES
					var specrelationtypes []*SPECRELATIONTYPE
					_, ok := res[specattributes_]
					if ok {
						specrelationtypes = res[specattributes_]
					} else {
						specrelationtypes = make([]*SPECRELATIONTYPE, 0)
					}
					specrelationtypes = append(specrelationtypes, specrelationtype)
					res[specattributes_] = specrelationtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECTYPES
	case SPECTYPES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TARGET
	case TARGET:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TARGETSPECIFICATION
	case TARGETSPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of THEHEADER
	case THEHEADER:
		switch fieldname {
		// insertion point for per direct association field
		case "REQIFHEADER":
			res := make(map[*REQIFHEADER][]*THEHEADER)
			for theheader := range stage.THEHEADERs {
				if theheader.REQIFHEADER != nil {
					reqifheader_ := theheader.REQIFHEADER
					var theheaders []*THEHEADER
					_, ok := res[reqifheader_]
					if ok {
						theheaders = res[reqifheader_]
					} else {
						theheaders = make([]*THEHEADER, 0)
					}
					theheaders = append(theheaders, theheader)
					res[reqifheader_] = theheaders
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TOOLEXTENSIONS
	case TOOLEXTENSIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VALUES
	case VALUES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTMLCONTENT
	case XHTMLCONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ALTERNATIVEID
	case ALTERNATIVEID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONBOOLEAN
	case ATTRIBUTEDEFINITIONBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONDATE
	case ATTRIBUTEDEFINITIONDATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONENUMERATION
	case ATTRIBUTEDEFINITIONENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONINTEGER
	case ATTRIBUTEDEFINITIONINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONREAL
	case ATTRIBUTEDEFINITIONREAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONSTRING
	case ATTRIBUTEDEFINITIONSTRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEDEFINITIONXHTML
	case ATTRIBUTEDEFINITIONXHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEBOOLEAN
	case ATTRIBUTEVALUEBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEDATE
	case ATTRIBUTEVALUEDATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEENUMERATION
	case ATTRIBUTEVALUEENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEINTEGER
	case ATTRIBUTEVALUEINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEREAL
	case ATTRIBUTEVALUEREAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUESTRING
	case ATTRIBUTEVALUESTRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTEVALUEXHTML
	case ATTRIBUTEVALUEXHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CHILDREN
	case CHILDREN:
		switch fieldname {
		// insertion point for per direct association field
		case "SPECHIERARCHY":
			res := make(map[*SPECHIERARCHY]*CHILDREN)
			for children := range stage.CHILDRENs {
				for _, spechierarchy_ := range children.SPECHIERARCHY {
					res[spechierarchy_] = children
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of CORECONTENT
	case CORECONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONBOOLEAN
	case DATATYPEDEFINITIONBOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONDATE
	case DATATYPEDEFINITIONDATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONENUMERATION
	case DATATYPEDEFINITIONENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONINTEGER
	case DATATYPEDEFINITIONINTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONREAL
	case DATATYPEDEFINITIONREAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONSTRING
	case DATATYPEDEFINITIONSTRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPEDEFINITIONXHTML
	case DATATYPEDEFINITIONXHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPES
	case DATATYPES:
		switch fieldname {
		// insertion point for per direct association field
		case "DATATYPEDEFINITIONBOOLEAN":
			res := make(map[*DATATYPEDEFINITIONBOOLEAN]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitionboolean_ := range datatypes.DATATYPEDEFINITIONBOOLEAN {
					res[datatypedefinitionboolean_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONDATE":
			res := make(map[*DATATYPEDEFINITIONDATE]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitiondate_ := range datatypes.DATATYPEDEFINITIONDATE {
					res[datatypedefinitiondate_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONENUMERATION":
			res := make(map[*DATATYPEDEFINITIONENUMERATION]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitionenumeration_ := range datatypes.DATATYPEDEFINITIONENUMERATION {
					res[datatypedefinitionenumeration_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONINTEGER":
			res := make(map[*DATATYPEDEFINITIONINTEGER]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitioninteger_ := range datatypes.DATATYPEDEFINITIONINTEGER {
					res[datatypedefinitioninteger_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONREAL":
			res := make(map[*DATATYPEDEFINITIONREAL]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitionreal_ := range datatypes.DATATYPEDEFINITIONREAL {
					res[datatypedefinitionreal_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONSTRING":
			res := make(map[*DATATYPEDEFINITIONSTRING]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitionstring_ := range datatypes.DATATYPEDEFINITIONSTRING {
					res[datatypedefinitionstring_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		case "DATATYPEDEFINITIONXHTML":
			res := make(map[*DATATYPEDEFINITIONXHTML]*DATATYPES)
			for datatypes := range stage.DATATYPESs {
				for _, datatypedefinitionxhtml_ := range datatypes.DATATYPEDEFINITIONXHTML {
					res[datatypedefinitionxhtml_] = datatypes
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DEFAULTVALUE
	case DEFAULTVALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DEFINITION
	case DEFINITION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EDITABLEATTS
	case EDITABLEATTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EMBEDDEDVALUE
	case EMBEDDEDVALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUMVALUE
	case ENUMVALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OBJECT
	case OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PROPERTIES
	case PROPERTIES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATIONGROUP
	case RELATIONGROUP:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATIONGROUPTYPE
	case RELATIONGROUPTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIFCONTENT
	case REQIFCONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIFHEADER
	case REQIFHEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIFTOOLEXTENSION
	case REQIFTOOLEXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQTYPE
	case REQTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SOURCE
	case SOURCE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SOURCESPECIFICATION
	case SOURCESPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECATTRIBUTES
	case SPECATTRIBUTES:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			res := make(map[*ATTRIBUTEDEFINITIONBOOLEAN]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitionboolean_ := range specattributes.ATTRIBUTEDEFINITIONBOOLEAN {
					res[attributedefinitionboolean_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONDATE":
			res := make(map[*ATTRIBUTEDEFINITIONDATE]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitiondate_ := range specattributes.ATTRIBUTEDEFINITIONDATE {
					res[attributedefinitiondate_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONENUMERATION":
			res := make(map[*ATTRIBUTEDEFINITIONENUMERATION]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitionenumeration_ := range specattributes.ATTRIBUTEDEFINITIONENUMERATION {
					res[attributedefinitionenumeration_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONINTEGER":
			res := make(map[*ATTRIBUTEDEFINITIONINTEGER]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitioninteger_ := range specattributes.ATTRIBUTEDEFINITIONINTEGER {
					res[attributedefinitioninteger_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONREAL":
			res := make(map[*ATTRIBUTEDEFINITIONREAL]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitionreal_ := range specattributes.ATTRIBUTEDEFINITIONREAL {
					res[attributedefinitionreal_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONSTRING":
			res := make(map[*ATTRIBUTEDEFINITIONSTRING]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitionstring_ := range specattributes.ATTRIBUTEDEFINITIONSTRING {
					res[attributedefinitionstring_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		case "ATTRIBUTEDEFINITIONXHTML":
			res := make(map[*ATTRIBUTEDEFINITIONXHTML]*SPECATTRIBUTES)
			for specattributes := range stage.SPECATTRIBUTESs {
				for _, attributedefinitionxhtml_ := range specattributes.ATTRIBUTEDEFINITIONXHTML {
					res[attributedefinitionxhtml_] = specattributes
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECHIERARCHY
	case SPECHIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATIONS
	case SPECIFICATIONS:
		switch fieldname {
		// insertion point for per direct association field
		case "SPECIFICATION":
			res := make(map[*SPECIFICATION]*SPECIFICATIONS)
			for specifications := range stage.SPECIFICATIONSs {
				for _, specification_ := range specifications.SPECIFICATION {
					res[specification_] = specifications
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECIFICATIONTYPE
	case SPECIFICATIONTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFIEDVALUES
	case SPECIFIEDVALUES:
		switch fieldname {
		// insertion point for per direct association field
		case "ENUMVALUE":
			res := make(map[*ENUMVALUE]*SPECIFIEDVALUES)
			for specifiedvalues := range stage.SPECIFIEDVALUESs {
				for _, enumvalue_ := range specifiedvalues.ENUMVALUE {
					res[enumvalue_] = specifiedvalues
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECOBJECT
	case SPECOBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECOBJECTS
	case SPECOBJECTS:
		switch fieldname {
		// insertion point for per direct association field
		case "SPECOBJECT":
			res := make(map[*SPECOBJECT]*SPECOBJECTS)
			for specobjects := range stage.SPECOBJECTSs {
				for _, specobject_ := range specobjects.SPECOBJECT {
					res[specobject_] = specobjects
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECOBJECTTYPE
	case SPECOBJECTTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECRELATION
	case SPECRELATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECRELATIONGROUPS
	case SPECRELATIONGROUPS:
		switch fieldname {
		// insertion point for per direct association field
		case "RELATIONGROUP":
			res := make(map[*RELATIONGROUP]*SPECRELATIONGROUPS)
			for specrelationgroups := range stage.SPECRELATIONGROUPSs {
				for _, relationgroup_ := range specrelationgroups.RELATIONGROUP {
					res[relationgroup_] = specrelationgroups
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SPECRELATIONS
	case SPECRELATIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECRELATIONTYPE
	case SPECRELATIONTYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECTYPES
	case SPECTYPES:
		switch fieldname {
		// insertion point for per direct association field
		case "RELATIONGROUPTYPE":
			res := make(map[*RELATIONGROUPTYPE]*SPECTYPES)
			for spectypes := range stage.SPECTYPESs {
				for _, relationgrouptype_ := range spectypes.RELATIONGROUPTYPE {
					res[relationgrouptype_] = spectypes
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECOBJECTTYPE":
			res := make(map[*SPECOBJECTTYPE]*SPECTYPES)
			for spectypes := range stage.SPECTYPESs {
				for _, specobjecttype_ := range spectypes.SPECOBJECTTYPE {
					res[specobjecttype_] = spectypes
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECRELATIONTYPE":
			res := make(map[*SPECRELATIONTYPE]*SPECTYPES)
			for spectypes := range stage.SPECTYPESs {
				for _, specrelationtype_ := range spectypes.SPECRELATIONTYPE {
					res[specrelationtype_] = spectypes
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECIFICATIONTYPE":
			res := make(map[*SPECIFICATIONTYPE]*SPECTYPES)
			for spectypes := range stage.SPECTYPESs {
				for _, specificationtype_ := range spectypes.SPECIFICATIONTYPE {
					res[specificationtype_] = spectypes
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TARGET
	case TARGET:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TARGETSPECIFICATION
	case TARGETSPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of THEHEADER
	case THEHEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TOOLEXTENSIONS
	case TOOLEXTENSIONS:
		switch fieldname {
		// insertion point for per direct association field
		case "REQIFTOOLEXTENSION":
			res := make(map[*REQIFTOOLEXTENSION]*TOOLEXTENSIONS)
			for toolextensions := range stage.TOOLEXTENSIONSs {
				for _, reqiftoolextension_ := range toolextensions.REQIFTOOLEXTENSION {
					res[reqiftoolextension_] = toolextensions
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of VALUES
	case VALUES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTMLCONTENT
	case XHTMLCONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ALTERNATIVEID:
		res = "ALTERNATIVEID"
	case ATTRIBUTEDEFINITIONBOOLEAN:
		res = "ATTRIBUTEDEFINITIONBOOLEAN"
	case ATTRIBUTEDEFINITIONDATE:
		res = "ATTRIBUTEDEFINITIONDATE"
	case ATTRIBUTEDEFINITIONENUMERATION:
		res = "ATTRIBUTEDEFINITIONENUMERATION"
	case ATTRIBUTEDEFINITIONINTEGER:
		res = "ATTRIBUTEDEFINITIONINTEGER"
	case ATTRIBUTEDEFINITIONREAL:
		res = "ATTRIBUTEDEFINITIONREAL"
	case ATTRIBUTEDEFINITIONSTRING:
		res = "ATTRIBUTEDEFINITIONSTRING"
	case ATTRIBUTEDEFINITIONXHTML:
		res = "ATTRIBUTEDEFINITIONXHTML"
	case ATTRIBUTEVALUEBOOLEAN:
		res = "ATTRIBUTEVALUEBOOLEAN"
	case ATTRIBUTEVALUEDATE:
		res = "ATTRIBUTEVALUEDATE"
	case ATTRIBUTEVALUEENUMERATION:
		res = "ATTRIBUTEVALUEENUMERATION"
	case ATTRIBUTEVALUEINTEGER:
		res = "ATTRIBUTEVALUEINTEGER"
	case ATTRIBUTEVALUEREAL:
		res = "ATTRIBUTEVALUEREAL"
	case ATTRIBUTEVALUESTRING:
		res = "ATTRIBUTEVALUESTRING"
	case ATTRIBUTEVALUEXHTML:
		res = "ATTRIBUTEVALUEXHTML"
	case CHILDREN:
		res = "CHILDREN"
	case CORECONTENT:
		res = "CORECONTENT"
	case DATATYPEDEFINITIONBOOLEAN:
		res = "DATATYPEDEFINITIONBOOLEAN"
	case DATATYPEDEFINITIONDATE:
		res = "DATATYPEDEFINITIONDATE"
	case DATATYPEDEFINITIONENUMERATION:
		res = "DATATYPEDEFINITIONENUMERATION"
	case DATATYPEDEFINITIONINTEGER:
		res = "DATATYPEDEFINITIONINTEGER"
	case DATATYPEDEFINITIONREAL:
		res = "DATATYPEDEFINITIONREAL"
	case DATATYPEDEFINITIONSTRING:
		res = "DATATYPEDEFINITIONSTRING"
	case DATATYPEDEFINITIONXHTML:
		res = "DATATYPEDEFINITIONXHTML"
	case DATATYPES:
		res = "DATATYPES"
	case DEFAULTVALUE:
		res = "DEFAULTVALUE"
	case DEFINITION:
		res = "DEFINITION"
	case EDITABLEATTS:
		res = "EDITABLEATTS"
	case EMBEDDEDVALUE:
		res = "EMBEDDEDVALUE"
	case ENUMVALUE:
		res = "ENUMVALUE"
	case OBJECT:
		res = "OBJECT"
	case PROPERTIES:
		res = "PROPERTIES"
	case RELATIONGROUP:
		res = "RELATIONGROUP"
	case RELATIONGROUPTYPE:
		res = "RELATIONGROUPTYPE"
	case REQIF:
		res = "REQIF"
	case REQIFCONTENT:
		res = "REQIFCONTENT"
	case REQIFHEADER:
		res = "REQIFHEADER"
	case REQIFTOOLEXTENSION:
		res = "REQIFTOOLEXTENSION"
	case REQTYPE:
		res = "REQTYPE"
	case SOURCE:
		res = "SOURCE"
	case SOURCESPECIFICATION:
		res = "SOURCESPECIFICATION"
	case SPECATTRIBUTES:
		res = "SPECATTRIBUTES"
	case SPECHIERARCHY:
		res = "SPECHIERARCHY"
	case SPECIFICATION:
		res = "SPECIFICATION"
	case SPECIFICATIONS:
		res = "SPECIFICATIONS"
	case SPECIFICATIONTYPE:
		res = "SPECIFICATIONTYPE"
	case SPECIFIEDVALUES:
		res = "SPECIFIEDVALUES"
	case SPECOBJECT:
		res = "SPECOBJECT"
	case SPECOBJECTS:
		res = "SPECOBJECTS"
	case SPECOBJECTTYPE:
		res = "SPECOBJECTTYPE"
	case SPECRELATION:
		res = "SPECRELATION"
	case SPECRELATIONGROUPS:
		res = "SPECRELATIONGROUPS"
	case SPECRELATIONS:
		res = "SPECRELATIONS"
	case SPECRELATIONTYPE:
		res = "SPECRELATIONTYPE"
	case SPECTYPES:
		res = "SPECTYPES"
	case TARGET:
		res = "TARGET"
	case TARGETSPECIFICATION:
		res = "TARGETSPECIFICATION"
	case THEHEADER:
		res = "THEHEADER"
	case TOOLEXTENSIONS:
		res = "TOOLEXTENSIONS"
	case VALUES:
		res = "VALUES"
	case XHTMLCONTENT:
		res = "XHTMLCONTENT"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVEID:
		res = "ALTERNATIVEID"
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		res = "ATTRIBUTEDEFINITIONBOOLEAN"
	case *ATTRIBUTEDEFINITIONDATE:
		res = "ATTRIBUTEDEFINITIONDATE"
	case *ATTRIBUTEDEFINITIONENUMERATION:
		res = "ATTRIBUTEDEFINITIONENUMERATION"
	case *ATTRIBUTEDEFINITIONINTEGER:
		res = "ATTRIBUTEDEFINITIONINTEGER"
	case *ATTRIBUTEDEFINITIONREAL:
		res = "ATTRIBUTEDEFINITIONREAL"
	case *ATTRIBUTEDEFINITIONSTRING:
		res = "ATTRIBUTEDEFINITIONSTRING"
	case *ATTRIBUTEDEFINITIONXHTML:
		res = "ATTRIBUTEDEFINITIONXHTML"
	case *ATTRIBUTEVALUEBOOLEAN:
		res = "ATTRIBUTEVALUEBOOLEAN"
	case *ATTRIBUTEVALUEDATE:
		res = "ATTRIBUTEVALUEDATE"
	case *ATTRIBUTEVALUEENUMERATION:
		res = "ATTRIBUTEVALUEENUMERATION"
	case *ATTRIBUTEVALUEINTEGER:
		res = "ATTRIBUTEVALUEINTEGER"
	case *ATTRIBUTEVALUEREAL:
		res = "ATTRIBUTEVALUEREAL"
	case *ATTRIBUTEVALUESTRING:
		res = "ATTRIBUTEVALUESTRING"
	case *ATTRIBUTEVALUEXHTML:
		res = "ATTRIBUTEVALUEXHTML"
	case *CHILDREN:
		res = "CHILDREN"
	case *CORECONTENT:
		res = "CORECONTENT"
	case *DATATYPEDEFINITIONBOOLEAN:
		res = "DATATYPEDEFINITIONBOOLEAN"
	case *DATATYPEDEFINITIONDATE:
		res = "DATATYPEDEFINITIONDATE"
	case *DATATYPEDEFINITIONENUMERATION:
		res = "DATATYPEDEFINITIONENUMERATION"
	case *DATATYPEDEFINITIONINTEGER:
		res = "DATATYPEDEFINITIONINTEGER"
	case *DATATYPEDEFINITIONREAL:
		res = "DATATYPEDEFINITIONREAL"
	case *DATATYPEDEFINITIONSTRING:
		res = "DATATYPEDEFINITIONSTRING"
	case *DATATYPEDEFINITIONXHTML:
		res = "DATATYPEDEFINITIONXHTML"
	case *DATATYPES:
		res = "DATATYPES"
	case *DEFAULTVALUE:
		res = "DEFAULTVALUE"
	case *DEFINITION:
		res = "DEFINITION"
	case *EDITABLEATTS:
		res = "EDITABLEATTS"
	case *EMBEDDEDVALUE:
		res = "EMBEDDEDVALUE"
	case *ENUMVALUE:
		res = "ENUMVALUE"
	case *OBJECT:
		res = "OBJECT"
	case *PROPERTIES:
		res = "PROPERTIES"
	case *RELATIONGROUP:
		res = "RELATIONGROUP"
	case *RELATIONGROUPTYPE:
		res = "RELATIONGROUPTYPE"
	case *REQIF:
		res = "REQIF"
	case *REQIFCONTENT:
		res = "REQIFCONTENT"
	case *REQIFHEADER:
		res = "REQIFHEADER"
	case *REQIFTOOLEXTENSION:
		res = "REQIFTOOLEXTENSION"
	case *REQTYPE:
		res = "REQTYPE"
	case *SOURCE:
		res = "SOURCE"
	case *SOURCESPECIFICATION:
		res = "SOURCESPECIFICATION"
	case *SPECATTRIBUTES:
		res = "SPECATTRIBUTES"
	case *SPECHIERARCHY:
		res = "SPECHIERARCHY"
	case *SPECIFICATION:
		res = "SPECIFICATION"
	case *SPECIFICATIONS:
		res = "SPECIFICATIONS"
	case *SPECIFICATIONTYPE:
		res = "SPECIFICATIONTYPE"
	case *SPECIFIEDVALUES:
		res = "SPECIFIEDVALUES"
	case *SPECOBJECT:
		res = "SPECOBJECT"
	case *SPECOBJECTS:
		res = "SPECOBJECTS"
	case *SPECOBJECTTYPE:
		res = "SPECOBJECTTYPE"
	case *SPECRELATION:
		res = "SPECRELATION"
	case *SPECRELATIONGROUPS:
		res = "SPECRELATIONGROUPS"
	case *SPECRELATIONS:
		res = "SPECRELATIONS"
	case *SPECRELATIONTYPE:
		res = "SPECRELATIONTYPE"
	case *SPECTYPES:
		res = "SPECTYPES"
	case *TARGET:
		res = "TARGET"
	case *TARGETSPECIFICATION:
		res = "TARGETSPECIFICATION"
	case *THEHEADER:
		res = "THEHEADER"
	case *TOOLEXTENSIONS:
		res = "TOOLEXTENSIONS"
	case *VALUES:
		res = "VALUES"
	case *XHTMLCONTENT:
		res = "XHTMLCONTENT"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ALTERNATIVEID:
		res = []string{"Name", "IDENTIFIERAttr"}
	case ATTRIBUTEDEFINITIONBOOLEAN:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEDEFINITIONDATE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEDEFINITIONENUMERATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MULTIVALUEDAttr", "DEFAULTVALUE", "ALTERNATIVEID", "TYPE"}
	case ATTRIBUTEDEFINITIONINTEGER:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEDEFINITIONREAL:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEDEFINITIONSTRING:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEDEFINITIONXHTML:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case ATTRIBUTEVALUEBOOLEAN:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case ATTRIBUTEVALUEDATE:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case ATTRIBUTEVALUEENUMERATION:
		res = []string{"Name", "DEFINITION", "VALUES"}
	case ATTRIBUTEVALUEINTEGER:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case ATTRIBUTEVALUEREAL:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case ATTRIBUTEVALUESTRING:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case ATTRIBUTEVALUEXHTML:
		res = []string{"Name", "ISSIMPLIFIEDAttr", "THEVALUE", "THEORIGINALVALUE", "DEFINITION"}
	case CHILDREN:
		res = []string{"Name", "SPECHIERARCHY"}
	case CORECONTENT:
		res = []string{"Name", "REQIFCONTENT"}
	case DATATYPEDEFINITIONBOOLEAN:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case DATATYPEDEFINITIONDATE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case DATATYPEDEFINITIONENUMERATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECIFIEDVALUES"}
	case DATATYPEDEFINITIONINTEGER:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXAttr", "MINAttr", "ALTERNATIVEID"}
	case DATATYPEDEFINITIONREAL:
		res = []string{"Name", "ACCURACYAttr", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXAttr", "MINAttr", "ALTERNATIVEID"}
	case DATATYPEDEFINITIONSTRING:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXLENGTHAttr", "ALTERNATIVEID"}
	case DATATYPEDEFINITIONXHTML:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case DATATYPES:
		res = []string{"Name", "DATATYPEDEFINITIONBOOLEAN", "DATATYPEDEFINITIONDATE", "DATATYPEDEFINITIONENUMERATION", "DATATYPEDEFINITIONINTEGER", "DATATYPEDEFINITIONREAL", "DATATYPEDEFINITIONSTRING", "DATATYPEDEFINITIONXHTML"}
	case DEFAULTVALUE:
		res = []string{"Name", "ATTRIBUTEVALUEBOOLEAN"}
	case DEFINITION:
		res = []string{"Name", "ATTRIBUTEDEFINITIONBOOLEANREF"}
	case EDITABLEATTS:
		res = []string{"Name"}
	case EMBEDDEDVALUE:
		res = []string{"Name", "KEYAttr", "OTHERCONTENTAttr"}
	case ENUMVALUE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "PROPERTIES"}
	case OBJECT:
		res = []string{"Name", "SPECOBJECTREF"}
	case PROPERTIES:
		res = []string{"Name", "EMBEDDEDVALUE"}
	case RELATIONGROUP:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SOURCESPECIFICATION", "SPECRELATIONS", "TARGETSPECIFICATION", "TYPE"}
	case RELATIONGROUPTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case REQIF:
		res = []string{"Name", "HEADER", "CORECONTENT"}
	case REQIFCONTENT:
		res = []string{"Name", "DATATYPES", "SPECTYPES", "SPECOBJECTS", "SPECRELATIONS", "SPECIFICATIONS", "SPECRELATIONGROUPS"}
	case REQIFHEADER:
		res = []string{"Name", "IDENTIFIERAttr", "COMMENT", "CREATIONTIME", "REPOSITORYID", "REQIFTOOLID", "REQIFVERSION", "SOURCETOOLID", "TITLE"}
	case REQIFTOOLEXTENSION:
		res = []string{"Name"}
	case REQTYPE:
		res = []string{"Name", "DATATYPEDEFINITIONBOOLEANREF"}
	case SOURCE:
		res = []string{"Name", "SPECOBJECTREF"}
	case SOURCESPECIFICATION:
		res = []string{"Name", "SPECIFICATIONREF"}
	case SPECATTRIBUTES:
		res = []string{"Name", "ATTRIBUTEDEFINITIONBOOLEAN", "ATTRIBUTEDEFINITIONDATE", "ATTRIBUTEDEFINITIONENUMERATION", "ATTRIBUTEDEFINITIONINTEGER", "ATTRIBUTEDEFINITIONREAL", "ATTRIBUTEDEFINITIONSTRING", "ATTRIBUTEDEFINITIONXHTML"}
	case SPECHIERARCHY:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "ISTABLEINTERNALAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "CHILDREN", "EDITABLEATTS", "OBJECT"}
	case SPECIFICATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "CHILDREN", "TYPE"}
	case SPECIFICATIONS:
		res = []string{"Name", "SPECIFICATION"}
	case SPECIFICATIONTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case SPECIFIEDVALUES:
		res = []string{"Name", "ENUMVALUE"}
	case SPECOBJECT:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "TYPE"}
	case SPECOBJECTS:
		res = []string{"Name", "SPECOBJECT"}
	case SPECOBJECTTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case SPECRELATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "SOURCE", "TARGET", "TYPE"}
	case SPECRELATIONGROUPS:
		res = []string{"Name", "RELATIONGROUP"}
	case SPECRELATIONS:
		res = []string{"Name"}
	case SPECRELATIONTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case SPECTYPES:
		res = []string{"Name", "RELATIONGROUPTYPE", "SPECOBJECTTYPE", "SPECRELATIONTYPE", "SPECIFICATIONTYPE"}
	case TARGET:
		res = []string{"Name", "SPECOBJECTREF"}
	case TARGETSPECIFICATION:
		res = []string{"Name", "SPECIFICATIONREF"}
	case THEHEADER:
		res = []string{"Name", "REQIFHEADER"}
	case TOOLEXTENSIONS:
		res = []string{"Name", "REQIFTOOLEXTENSION"}
	case VALUES:
		res = []string{"Name"}
	case XHTMLCONTENT:
		res = []string{"Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case ALTERNATIVEID:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEDEFINITIONBOOLEAN:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONBOOLEAN"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONDATE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONDATE"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONENUMERATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONENUMERATION"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONINTEGER:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONINTEGER"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONREAL:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONREAL"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONSTRING:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONSTRING"
		res = append(res, rf)
	case ATTRIBUTEDEFINITIONXHTML:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECATTRIBUTES"
		rf.Fieldname = "ATTRIBUTEDEFINITIONXHTML"
		res = append(res, rf)
	case ATTRIBUTEVALUEBOOLEAN:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUEDATE:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUEENUMERATION:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUEINTEGER:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUEREAL:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUESTRING:
		var rf ReverseField
		_ = rf
	case ATTRIBUTEVALUEXHTML:
		var rf ReverseField
		_ = rf
	case CHILDREN:
		var rf ReverseField
		_ = rf
	case CORECONTENT:
		var rf ReverseField
		_ = rf
	case DATATYPEDEFINITIONBOOLEAN:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONBOOLEAN"
		res = append(res, rf)
	case DATATYPEDEFINITIONDATE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONDATE"
		res = append(res, rf)
	case DATATYPEDEFINITIONENUMERATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONENUMERATION"
		res = append(res, rf)
	case DATATYPEDEFINITIONINTEGER:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONINTEGER"
		res = append(res, rf)
	case DATATYPEDEFINITIONREAL:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONREAL"
		res = append(res, rf)
	case DATATYPEDEFINITIONSTRING:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONSTRING"
		res = append(res, rf)
	case DATATYPEDEFINITIONXHTML:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DATATYPES"
		rf.Fieldname = "DATATYPEDEFINITIONXHTML"
		res = append(res, rf)
	case DATATYPES:
		var rf ReverseField
		_ = rf
	case DEFAULTVALUE:
		var rf ReverseField
		_ = rf
	case DEFINITION:
		var rf ReverseField
		_ = rf
	case EDITABLEATTS:
		var rf ReverseField
		_ = rf
	case EMBEDDEDVALUE:
		var rf ReverseField
		_ = rf
	case ENUMVALUE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECIFIEDVALUES"
		rf.Fieldname = "ENUMVALUE"
		res = append(res, rf)
	case OBJECT:
		var rf ReverseField
		_ = rf
	case PROPERTIES:
		var rf ReverseField
		_ = rf
	case RELATIONGROUP:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECRELATIONGROUPS"
		rf.Fieldname = "RELATIONGROUP"
		res = append(res, rf)
	case RELATIONGROUPTYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECTYPES"
		rf.Fieldname = "RELATIONGROUPTYPE"
		res = append(res, rf)
	case REQIF:
		var rf ReverseField
		_ = rf
	case REQIFCONTENT:
		var rf ReverseField
		_ = rf
	case REQIFHEADER:
		var rf ReverseField
		_ = rf
	case REQIFTOOLEXTENSION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TOOLEXTENSIONS"
		rf.Fieldname = "REQIFTOOLEXTENSION"
		res = append(res, rf)
	case REQTYPE:
		var rf ReverseField
		_ = rf
	case SOURCE:
		var rf ReverseField
		_ = rf
	case SOURCESPECIFICATION:
		var rf ReverseField
		_ = rf
	case SPECATTRIBUTES:
		var rf ReverseField
		_ = rf
	case SPECHIERARCHY:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "CHILDREN"
		rf.Fieldname = "SPECHIERARCHY"
		res = append(res, rf)
	case SPECIFICATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECIFICATIONS"
		rf.Fieldname = "SPECIFICATION"
		res = append(res, rf)
	case SPECIFICATIONS:
		var rf ReverseField
		_ = rf
	case SPECIFICATIONTYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECTYPES"
		rf.Fieldname = "SPECIFICATIONTYPE"
		res = append(res, rf)
	case SPECIFIEDVALUES:
		var rf ReverseField
		_ = rf
	case SPECOBJECT:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECOBJECTS"
		rf.Fieldname = "SPECOBJECT"
		res = append(res, rf)
	case SPECOBJECTS:
		var rf ReverseField
		_ = rf
	case SPECOBJECTTYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECTYPES"
		rf.Fieldname = "SPECOBJECTTYPE"
		res = append(res, rf)
	case SPECRELATION:
		var rf ReverseField
		_ = rf
	case SPECRELATIONGROUPS:
		var rf ReverseField
		_ = rf
	case SPECRELATIONS:
		var rf ReverseField
		_ = rf
	case SPECRELATIONTYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "SPECTYPES"
		rf.Fieldname = "SPECRELATIONTYPE"
		res = append(res, rf)
	case SPECTYPES:
		var rf ReverseField
		_ = rf
	case TARGET:
		var rf ReverseField
		_ = rf
	case TARGETSPECIFICATION:
		var rf ReverseField
		_ = rf
	case THEHEADER:
		var rf ReverseField
		_ = rf
	case TOOLEXTENSIONS:
		var rf ReverseField
		_ = rf
	case VALUES:
		var rf ReverseField
		_ = rf
	case XHTMLCONTENT:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVEID:
		res = []string{"Name", "IDENTIFIERAttr"}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEDEFINITIONDATE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MULTIVALUEDAttr", "DEFAULTVALUE", "ALTERNATIVEID", "TYPE"}
	case *ATTRIBUTEDEFINITIONINTEGER:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEDEFINITIONREAL:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEDEFINITIONSTRING:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEDEFINITIONXHTML:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "DEFAULTVALUE", "TYPE"}
	case *ATTRIBUTEVALUEBOOLEAN:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case *ATTRIBUTEVALUEDATE:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case *ATTRIBUTEVALUEENUMERATION:
		res = []string{"Name", "DEFINITION", "VALUES"}
	case *ATTRIBUTEVALUEINTEGER:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case *ATTRIBUTEVALUEREAL:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case *ATTRIBUTEVALUESTRING:
		res = []string{"Name", "THEVALUEAttr", "DEFINITION"}
	case *ATTRIBUTEVALUEXHTML:
		res = []string{"Name", "ISSIMPLIFIEDAttr", "THEVALUE", "THEORIGINALVALUE", "DEFINITION"}
	case *CHILDREN:
		res = []string{"Name", "SPECHIERARCHY"}
	case *CORECONTENT:
		res = []string{"Name", "REQIFCONTENT"}
	case *DATATYPEDEFINITIONBOOLEAN:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case *DATATYPEDEFINITIONDATE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case *DATATYPEDEFINITIONENUMERATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECIFIEDVALUES"}
	case *DATATYPEDEFINITIONINTEGER:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXAttr", "MINAttr", "ALTERNATIVEID"}
	case *DATATYPEDEFINITIONREAL:
		res = []string{"Name", "ACCURACYAttr", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXAttr", "MINAttr", "ALTERNATIVEID"}
	case *DATATYPEDEFINITIONSTRING:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "MAXLENGTHAttr", "ALTERNATIVEID"}
	case *DATATYPEDEFINITIONXHTML:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID"}
	case *DATATYPES:
		res = []string{"Name", "DATATYPEDEFINITIONBOOLEAN", "DATATYPEDEFINITIONDATE", "DATATYPEDEFINITIONENUMERATION", "DATATYPEDEFINITIONINTEGER", "DATATYPEDEFINITIONREAL", "DATATYPEDEFINITIONSTRING", "DATATYPEDEFINITIONXHTML"}
	case *DEFAULTVALUE:
		res = []string{"Name", "ATTRIBUTEVALUEBOOLEAN"}
	case *DEFINITION:
		res = []string{"Name", "ATTRIBUTEDEFINITIONBOOLEANREF"}
	case *EDITABLEATTS:
		res = []string{"Name"}
	case *EMBEDDEDVALUE:
		res = []string{"Name", "KEYAttr", "OTHERCONTENTAttr"}
	case *ENUMVALUE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "PROPERTIES"}
	case *OBJECT:
		res = []string{"Name", "SPECOBJECTREF"}
	case *PROPERTIES:
		res = []string{"Name", "EMBEDDEDVALUE"}
	case *RELATIONGROUP:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SOURCESPECIFICATION", "SPECRELATIONS", "TARGETSPECIFICATION", "TYPE"}
	case *RELATIONGROUPTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case *REQIF:
		res = []string{"Name", "HEADER", "CORECONTENT"}
	case *REQIFCONTENT:
		res = []string{"Name", "DATATYPES", "SPECTYPES", "SPECOBJECTS", "SPECRELATIONS", "SPECIFICATIONS", "SPECRELATIONGROUPS"}
	case *REQIFHEADER:
		res = []string{"Name", "IDENTIFIERAttr", "COMMENT", "CREATIONTIME", "REPOSITORYID", "REQIFTOOLID", "REQIFVERSION", "SOURCETOOLID", "TITLE"}
	case *REQIFTOOLEXTENSION:
		res = []string{"Name"}
	case *REQTYPE:
		res = []string{"Name", "DATATYPEDEFINITIONBOOLEANREF"}
	case *SOURCE:
		res = []string{"Name", "SPECOBJECTREF"}
	case *SOURCESPECIFICATION:
		res = []string{"Name", "SPECIFICATIONREF"}
	case *SPECATTRIBUTES:
		res = []string{"Name", "ATTRIBUTEDEFINITIONBOOLEAN", "ATTRIBUTEDEFINITIONDATE", "ATTRIBUTEDEFINITIONENUMERATION", "ATTRIBUTEDEFINITIONINTEGER", "ATTRIBUTEDEFINITIONREAL", "ATTRIBUTEDEFINITIONSTRING", "ATTRIBUTEDEFINITIONXHTML"}
	case *SPECHIERARCHY:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "ISEDITABLEAttr", "ISTABLEINTERNALAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "CHILDREN", "EDITABLEATTS", "OBJECT"}
	case *SPECIFICATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "CHILDREN", "TYPE"}
	case *SPECIFICATIONS:
		res = []string{"Name", "SPECIFICATION"}
	case *SPECIFICATIONTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case *SPECIFIEDVALUES:
		res = []string{"Name", "ENUMVALUE"}
	case *SPECOBJECT:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "TYPE"}
	case *SPECOBJECTS:
		res = []string{"Name", "SPECOBJECT"}
	case *SPECOBJECTTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case *SPECRELATION:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "VALUES", "SOURCE", "TARGET", "TYPE"}
	case *SPECRELATIONGROUPS:
		res = []string{"Name", "RELATIONGROUP"}
	case *SPECRELATIONS:
		res = []string{"Name"}
	case *SPECRELATIONTYPE:
		res = []string{"Name", "DESCAttr", "IDENTIFIERAttr", "LASTCHANGEAttr", "LONGNAMEAttr", "ALTERNATIVEID", "SPECATTRIBUTES"}
	case *SPECTYPES:
		res = []string{"Name", "RELATIONGROUPTYPE", "SPECOBJECTTYPE", "SPECRELATIONTYPE", "SPECIFICATIONTYPE"}
	case *TARGET:
		res = []string{"Name", "SPECOBJECTREF"}
	case *TARGETSPECIFICATION:
		res = []string{"Name", "SPECIFICATIONREF"}
	case *THEHEADER:
		res = []string{"Name", "REQIFHEADER"}
	case *TOOLEXTENSIONS:
		res = []string{"Name", "REQIFTOOLEXTENSION"}
	case *VALUES:
		res = []string{"Name"}
	case *XHTMLCONTENT:
		res = []string{"Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *ALTERNATIVEID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MULTIVALUEDAttr":
			res = fmt.Sprintf("%t", inferedInstance.MULTIVALUEDAttr)
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONSTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEDEFINITIONXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTEVALUEBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%t", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTEVALUEDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = inferedInstance.THEVALUEAttr
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTEVALUEENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		}
	case *ATTRIBUTEVALUEINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%d", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTEVALUEREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%f", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTEVALUESTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = inferedInstance.THEVALUEAttr
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTEVALUEXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ISSIMPLIFIEDAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISSIMPLIFIEDAttr)
		case "THEVALUE":
			if inferedInstance.THEVALUE != nil {
				res = inferedInstance.THEVALUE.Name
			}
		case "THEORIGINALVALUE":
			if inferedInstance.THEORIGINALVALUE != nil {
				res = inferedInstance.THEORIGINALVALUE.Name
			}
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *CHILDREN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECHIERARCHY":
			for idx, __instance__ := range inferedInstance.SPECHIERARCHY {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *CORECONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFCONTENT":
			if inferedInstance.REQIFCONTENT != nil {
				res = inferedInstance.REQIFCONTENT.Name
			}
		}
	case *DATATYPEDEFINITIONBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPEDEFINITIONDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPEDEFINITIONENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECIFIEDVALUES":
			if inferedInstance.SPECIFIEDVALUES != nil {
				res = inferedInstance.SPECIFIEDVALUES.Name
			}
		}
	case *DATATYPEDEFINITIONINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXAttr":
			res = fmt.Sprintf("%d", inferedInstance.MAXAttr)
		case "MINAttr":
			res = fmt.Sprintf("%d", inferedInstance.MINAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPEDEFINITIONREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ACCURACYAttr":
			res = fmt.Sprintf("%d", inferedInstance.ACCURACYAttr)
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXAttr":
			res = fmt.Sprintf("%f", inferedInstance.MAXAttr)
		case "MINAttr":
			res = fmt.Sprintf("%f", inferedInstance.MINAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPEDEFINITIONSTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXLENGTHAttr":
			res = fmt.Sprintf("%d", inferedInstance.MAXLENGTHAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPEDEFINITIONXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case *DATATYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPEDEFINITIONBOOLEAN":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONBOOLEAN {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONDATE":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONDATE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONENUMERATION":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONENUMERATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONINTEGER":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONINTEGER {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONREAL":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONREAL {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONSTRING":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONSTRING {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONXHTML":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONXHTML {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *DEFAULTVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEVALUEBOOLEAN":
			if inferedInstance.ATTRIBUTEVALUEBOOLEAN != nil {
				res = inferedInstance.ATTRIBUTEVALUEBOOLEAN.Name
			}
		}
	case *DEFINITION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEDEFINITIONBOOLEANREF":
			res = inferedInstance.ATTRIBUTEDEFINITIONBOOLEANREF
		}
	case *EDITABLEATTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *EMBEDDEDVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "KEYAttr":
			res = fmt.Sprintf("%d", inferedInstance.KEYAttr)
		case "OTHERCONTENTAttr":
			res = inferedInstance.OTHERCONTENTAttr
		}
	case *ENUMVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "PROPERTIES":
			if inferedInstance.PROPERTIES != nil {
				res = inferedInstance.PROPERTIES.Name
			}
		}
	case *OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case *PROPERTIES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EMBEDDEDVALUE":
			if inferedInstance.EMBEDDEDVALUE != nil {
				res = inferedInstance.EMBEDDEDVALUE.Name
			}
		}
	case *RELATIONGROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SOURCESPECIFICATION":
			if inferedInstance.SOURCESPECIFICATION != nil {
				res = inferedInstance.SOURCESPECIFICATION.Name
			}
		case "SPECRELATIONS":
			if inferedInstance.SPECRELATIONS != nil {
				res = inferedInstance.SPECRELATIONS.Name
			}
		case "TARGETSPECIFICATION":
			if inferedInstance.TARGETSPECIFICATION != nil {
				res = inferedInstance.TARGETSPECIFICATION.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *RELATIONGROUPTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case *REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "HEADER":
			if inferedInstance.HEADER != nil {
				res = inferedInstance.HEADER.Name
			}
		case "CORECONTENT":
			if inferedInstance.CORECONTENT != nil {
				res = inferedInstance.CORECONTENT.Name
			}
		}
	case *REQIFCONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPES":
			if inferedInstance.DATATYPES != nil {
				res = inferedInstance.DATATYPES.Name
			}
		case "SPECTYPES":
			if inferedInstance.SPECTYPES != nil {
				res = inferedInstance.SPECTYPES.Name
			}
		case "SPECOBJECTS":
			if inferedInstance.SPECOBJECTS != nil {
				res = inferedInstance.SPECOBJECTS.Name
			}
		case "SPECRELATIONS":
			if inferedInstance.SPECRELATIONS != nil {
				res = inferedInstance.SPECRELATIONS.Name
			}
		case "SPECIFICATIONS":
			if inferedInstance.SPECIFICATIONS != nil {
				res = inferedInstance.SPECIFICATIONS.Name
			}
		case "SPECRELATIONGROUPS":
			if inferedInstance.SPECRELATIONGROUPS != nil {
				res = inferedInstance.SPECRELATIONGROUPS.Name
			}
		}
	case *REQIFHEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATIONTIME":
			res = inferedInstance.CREATIONTIME
		case "REPOSITORYID":
			res = inferedInstance.REPOSITORYID
		case "REQIFTOOLID":
			res = inferedInstance.REQIFTOOLID
		case "REQIFVERSION":
			res = inferedInstance.REQIFVERSION
		case "SOURCETOOLID":
			res = inferedInstance.SOURCETOOLID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case *REQIFTOOLEXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *REQTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPEDEFINITIONBOOLEANREF":
			res = inferedInstance.DATATYPEDEFINITIONBOOLEANREF
		}
	case *SOURCE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case *SOURCESPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATIONREF":
			res = inferedInstance.SPECIFICATIONREF
		}
	case *SPECATTRIBUTES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONBOOLEAN {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONDATE":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONDATE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONENUMERATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONINTEGER {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONREAL":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONREAL {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONSTRING {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONXHTML {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SPECHIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "ISTABLEINTERNALAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISTABLEINTERNALAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "EDITABLEATTS":
			if inferedInstance.EDITABLEATTS != nil {
				res = inferedInstance.EDITABLEATTS.Name
			}
		case "OBJECT":
			if inferedInstance.OBJECT != nil {
				res = inferedInstance.OBJECT.Name
			}
		}
	case *SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *SPECIFICATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION":
			for idx, __instance__ := range inferedInstance.SPECIFICATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SPECIFICATIONTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case *SPECIFIEDVALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUMVALUE":
			for idx, __instance__ := range inferedInstance.ENUMVALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SPECOBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *SPECOBJECTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECT":
			for idx, __instance__ := range inferedInstance.SPECOBJECT {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SPECOBJECTTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case *SPECRELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "SOURCE":
			if inferedInstance.SOURCE != nil {
				res = inferedInstance.SOURCE.Name
			}
		case "TARGET":
			if inferedInstance.TARGET != nil {
				res = inferedInstance.TARGET.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *SPECRELATIONGROUPS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATIONGROUP":
			for idx, __instance__ := range inferedInstance.RELATIONGROUP {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SPECRELATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *SPECRELATIONTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case *SPECTYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATIONGROUPTYPE":
			for idx, __instance__ := range inferedInstance.RELATIONGROUPTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECOBJECTTYPE":
			for idx, __instance__ := range inferedInstance.SPECOBJECTTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECRELATIONTYPE":
			for idx, __instance__ := range inferedInstance.SPECRELATIONTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATIONTYPE":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *TARGET:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case *TARGETSPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATIONREF":
			res = inferedInstance.SPECIFICATIONREF
		}
	case *THEHEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFHEADER":
			if inferedInstance.REQIFHEADER != nil {
				res = inferedInstance.REQIFHEADER.Name
			}
		}
	case *TOOLEXTENSIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFTOOLEXTENSION":
			for idx, __instance__ := range inferedInstance.REQIFTOOLEXTENSION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *VALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *XHTMLCONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case ALTERNATIVEID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		}
	case ATTRIBUTEDEFINITIONBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MULTIVALUEDAttr":
			res = fmt.Sprintf("%t", inferedInstance.MULTIVALUEDAttr)
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONSTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEDEFINITIONXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "DEFAULTVALUE":
			if inferedInstance.DEFAULTVALUE != nil {
				res = inferedInstance.DEFAULTVALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTEVALUEBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%t", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTEVALUEDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = inferedInstance.THEVALUEAttr
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTEVALUEENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		}
	case ATTRIBUTEVALUEINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%d", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTEVALUEREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = fmt.Sprintf("%f", inferedInstance.THEVALUEAttr)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTEVALUESTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THEVALUEAttr":
			res = inferedInstance.THEVALUEAttr
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTEVALUEXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ISSIMPLIFIEDAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISSIMPLIFIEDAttr)
		case "THEVALUE":
			if inferedInstance.THEVALUE != nil {
				res = inferedInstance.THEVALUE.Name
			}
		case "THEORIGINALVALUE":
			if inferedInstance.THEORIGINALVALUE != nil {
				res = inferedInstance.THEORIGINALVALUE.Name
			}
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case CHILDREN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECHIERARCHY":
			for idx, __instance__ := range inferedInstance.SPECHIERARCHY {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case CORECONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFCONTENT":
			if inferedInstance.REQIFCONTENT != nil {
				res = inferedInstance.REQIFCONTENT.Name
			}
		}
	case DATATYPEDEFINITIONBOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPEDEFINITIONDATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPEDEFINITIONENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECIFIEDVALUES":
			if inferedInstance.SPECIFIEDVALUES != nil {
				res = inferedInstance.SPECIFIEDVALUES.Name
			}
		}
	case DATATYPEDEFINITIONINTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXAttr":
			res = fmt.Sprintf("%d", inferedInstance.MAXAttr)
		case "MINAttr":
			res = fmt.Sprintf("%d", inferedInstance.MINAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPEDEFINITIONREAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ACCURACYAttr":
			res = fmt.Sprintf("%d", inferedInstance.ACCURACYAttr)
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXAttr":
			res = fmt.Sprintf("%f", inferedInstance.MAXAttr)
		case "MINAttr":
			res = fmt.Sprintf("%f", inferedInstance.MINAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPEDEFINITIONSTRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "MAXLENGTHAttr":
			res = fmt.Sprintf("%d", inferedInstance.MAXLENGTHAttr)
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPEDEFINITIONXHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		}
	case DATATYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPEDEFINITIONBOOLEAN":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONBOOLEAN {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONDATE":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONDATE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONENUMERATION":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONENUMERATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONINTEGER":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONINTEGER {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONREAL":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONREAL {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONSTRING":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONSTRING {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "DATATYPEDEFINITIONXHTML":
			for idx, __instance__ := range inferedInstance.DATATYPEDEFINITIONXHTML {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case DEFAULTVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEVALUEBOOLEAN":
			if inferedInstance.ATTRIBUTEVALUEBOOLEAN != nil {
				res = inferedInstance.ATTRIBUTEVALUEBOOLEAN.Name
			}
		}
	case DEFINITION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEDEFINITIONBOOLEANREF":
			res = inferedInstance.ATTRIBUTEDEFINITIONBOOLEANREF
		}
	case EDITABLEATTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case EMBEDDEDVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "KEYAttr":
			res = fmt.Sprintf("%d", inferedInstance.KEYAttr)
		case "OTHERCONTENTAttr":
			res = inferedInstance.OTHERCONTENTAttr
		}
	case ENUMVALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "PROPERTIES":
			if inferedInstance.PROPERTIES != nil {
				res = inferedInstance.PROPERTIES.Name
			}
		}
	case OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case PROPERTIES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EMBEDDEDVALUE":
			if inferedInstance.EMBEDDEDVALUE != nil {
				res = inferedInstance.EMBEDDEDVALUE.Name
			}
		}
	case RELATIONGROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SOURCESPECIFICATION":
			if inferedInstance.SOURCESPECIFICATION != nil {
				res = inferedInstance.SOURCESPECIFICATION.Name
			}
		case "SPECRELATIONS":
			if inferedInstance.SPECRELATIONS != nil {
				res = inferedInstance.SPECRELATIONS.Name
			}
		case "TARGETSPECIFICATION":
			if inferedInstance.TARGETSPECIFICATION != nil {
				res = inferedInstance.TARGETSPECIFICATION.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case RELATIONGROUPTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "HEADER":
			if inferedInstance.HEADER != nil {
				res = inferedInstance.HEADER.Name
			}
		case "CORECONTENT":
			if inferedInstance.CORECONTENT != nil {
				res = inferedInstance.CORECONTENT.Name
			}
		}
	case REQIFCONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPES":
			if inferedInstance.DATATYPES != nil {
				res = inferedInstance.DATATYPES.Name
			}
		case "SPECTYPES":
			if inferedInstance.SPECTYPES != nil {
				res = inferedInstance.SPECTYPES.Name
			}
		case "SPECOBJECTS":
			if inferedInstance.SPECOBJECTS != nil {
				res = inferedInstance.SPECOBJECTS.Name
			}
		case "SPECRELATIONS":
			if inferedInstance.SPECRELATIONS != nil {
				res = inferedInstance.SPECRELATIONS.Name
			}
		case "SPECIFICATIONS":
			if inferedInstance.SPECIFICATIONS != nil {
				res = inferedInstance.SPECIFICATIONS.Name
			}
		case "SPECRELATIONGROUPS":
			if inferedInstance.SPECRELATIONGROUPS != nil {
				res = inferedInstance.SPECRELATIONGROUPS.Name
			}
		}
	case REQIFHEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATIONTIME":
			res = inferedInstance.CREATIONTIME
		case "REPOSITORYID":
			res = inferedInstance.REPOSITORYID
		case "REQIFTOOLID":
			res = inferedInstance.REQIFTOOLID
		case "REQIFVERSION":
			res = inferedInstance.REQIFVERSION
		case "SOURCETOOLID":
			res = inferedInstance.SOURCETOOLID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case REQIFTOOLEXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case REQTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPEDEFINITIONBOOLEANREF":
			res = inferedInstance.DATATYPEDEFINITIONBOOLEANREF
		}
	case SOURCE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case SOURCESPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATIONREF":
			res = inferedInstance.SPECIFICATIONREF
		}
	case SPECATTRIBUTES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONBOOLEAN {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONDATE":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONDATE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONENUMERATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONINTEGER {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONREAL":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONREAL {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONSTRING {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			for idx, __instance__ := range inferedInstance.ATTRIBUTEDEFINITIONXHTML {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SPECHIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "ISEDITABLEAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISEDITABLEAttr)
		case "ISTABLEINTERNALAttr":
			res = fmt.Sprintf("%t", inferedInstance.ISTABLEINTERNALAttr)
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "EDITABLEATTS":
			if inferedInstance.EDITABLEATTS != nil {
				res = inferedInstance.EDITABLEATTS.Name
			}
		case "OBJECT":
			if inferedInstance.OBJECT != nil {
				res = inferedInstance.OBJECT.Name
			}
		}
	case SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case SPECIFICATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION":
			for idx, __instance__ := range inferedInstance.SPECIFICATION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SPECIFICATIONTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case SPECIFIEDVALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUMVALUE":
			for idx, __instance__ := range inferedInstance.ENUMVALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SPECOBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case SPECOBJECTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECT":
			for idx, __instance__ := range inferedInstance.SPECOBJECT {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SPECOBJECTTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case SPECRELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "SOURCE":
			if inferedInstance.SOURCE != nil {
				res = inferedInstance.SOURCE.Name
			}
		case "TARGET":
			if inferedInstance.TARGET != nil {
				res = inferedInstance.TARGET.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case SPECRELATIONGROUPS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATIONGROUP":
			for idx, __instance__ := range inferedInstance.RELATIONGROUP {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SPECRELATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case SPECRELATIONTYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESCAttr":
			res = inferedInstance.DESCAttr
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "LASTCHANGEAttr":
			res = inferedInstance.LASTCHANGEAttr
		case "LONGNAMEAttr":
			res = inferedInstance.LONGNAMEAttr
		case "ALTERNATIVEID":
			if inferedInstance.ALTERNATIVEID != nil {
				res = inferedInstance.ALTERNATIVEID.Name
			}
		case "SPECATTRIBUTES":
			if inferedInstance.SPECATTRIBUTES != nil {
				res = inferedInstance.SPECATTRIBUTES.Name
			}
		}
	case SPECTYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATIONGROUPTYPE":
			for idx, __instance__ := range inferedInstance.RELATIONGROUPTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECOBJECTTYPE":
			for idx, __instance__ := range inferedInstance.SPECOBJECTTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECRELATIONTYPE":
			for idx, __instance__ := range inferedInstance.SPECRELATIONTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATIONTYPE":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONTYPE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case TARGET:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECOBJECTREF":
			res = inferedInstance.SPECOBJECTREF
		}
	case TARGETSPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATIONREF":
			res = inferedInstance.SPECIFICATIONREF
		}
	case THEHEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFHEADER":
			if inferedInstance.REQIFHEADER != nil {
				res = inferedInstance.REQIFHEADER.Name
			}
		}
	case TOOLEXTENSIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQIFTOOLEXTENSION":
			for idx, __instance__ := range inferedInstance.REQIFTOOLEXTENSION {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case VALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case XHTMLCONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
