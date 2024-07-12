// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVEID:
		if stage.OnAfterALTERNATIVEIDCreateCallback != nil {
			stage.OnAfterALTERNATIVEIDCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		if stage.OnAfterATTRIBUTEDEFINITIONBOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONBOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONDATE:
		if stage.OnAfterATTRIBUTEDEFINITIONDATECreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONDATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		if stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONINTEGER:
		if stage.OnAfterATTRIBUTEDEFINITIONINTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONINTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONREAL:
		if stage.OnAfterATTRIBUTEDEFINITIONREALCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONREALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONSTRING:
		if stage.OnAfterATTRIBUTEDEFINITIONSTRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONSTRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEDEFINITIONXHTML:
		if stage.OnAfterATTRIBUTEDEFINITIONXHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONXHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEBOOLEAN:
		if stage.OnAfterATTRIBUTEVALUEBOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEBOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEDATE:
		if stage.OnAfterATTRIBUTEVALUEDATECreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEDATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEENUMERATION:
		if stage.OnAfterATTRIBUTEVALUEENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEINTEGER:
		if stage.OnAfterATTRIBUTEVALUEINTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEINTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEREAL:
		if stage.OnAfterATTRIBUTEVALUEREALCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEREALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUESTRING:
		if stage.OnAfterATTRIBUTEVALUESTRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUESTRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTEVALUEXHTML:
		if stage.OnAfterATTRIBUTEVALUEXHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEXHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *CHILDREN:
		if stage.OnAfterCHILDRENCreateCallback != nil {
			stage.OnAfterCHILDRENCreateCallback.OnAfterCreate(stage, target)
		}
	case *CORECONTENT:
		if stage.OnAfterCORECONTENTCreateCallback != nil {
			stage.OnAfterCORECONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONBOOLEAN:
		if stage.OnAfterDATATYPEDEFINITIONBOOLEANCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONBOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONDATE:
		if stage.OnAfterDATATYPEDEFINITIONDATECreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONDATECreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONENUMERATION:
		if stage.OnAfterDATATYPEDEFINITIONENUMERATIONCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONINTEGER:
		if stage.OnAfterDATATYPEDEFINITIONINTEGERCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONINTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONREAL:
		if stage.OnAfterDATATYPEDEFINITIONREALCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONREALCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONSTRING:
		if stage.OnAfterDATATYPEDEFINITIONSTRINGCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONSTRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPEDEFINITIONXHTML:
		if stage.OnAfterDATATYPEDEFINITIONXHTMLCreateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONXHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPES:
		if stage.OnAfterDATATYPESCreateCallback != nil {
			stage.OnAfterDATATYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *DEFAULTVALUE:
		if stage.OnAfterDEFAULTVALUECreateCallback != nil {
			stage.OnAfterDEFAULTVALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *DEFINITION:
		if stage.OnAfterDEFINITIONCreateCallback != nil {
			stage.OnAfterDEFINITIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *EDITABLEATTS:
		if stage.OnAfterEDITABLEATTSCreateCallback != nil {
			stage.OnAfterEDITABLEATTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *EMBEDDEDVALUE:
		if stage.OnAfterEMBEDDEDVALUECreateCallback != nil {
			stage.OnAfterEMBEDDEDVALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *ENUMVALUE:
		if stage.OnAfterENUMVALUECreateCallback != nil {
			stage.OnAfterENUMVALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *OBJECT:
		if stage.OnAfterOBJECTCreateCallback != nil {
			stage.OnAfterOBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *PROPERTIES:
		if stage.OnAfterPROPERTIESCreateCallback != nil {
			stage.OnAfterPROPERTIESCreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATIONGROUP:
		if stage.OnAfterRELATIONGROUPCreateCallback != nil {
			stage.OnAfterRELATIONGROUPCreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATIONGROUPTYPE:
		if stage.OnAfterRELATIONGROUPTYPECreateCallback != nil {
			stage.OnAfterRELATIONGROUPTYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIF:
		if stage.OnAfterREQIFCreateCallback != nil {
			stage.OnAfterREQIFCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIFCONTENT:
		if stage.OnAfterREQIFCONTENTCreateCallback != nil {
			stage.OnAfterREQIFCONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERCreateCallback != nil {
			stage.OnAfterREQIFHEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIFTOOLEXTENSION:
		if stage.OnAfterREQIFTOOLEXTENSIONCreateCallback != nil {
			stage.OnAfterREQIFTOOLEXTENSIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIFTYPE:
		if stage.OnAfterREQIFTYPECreateCallback != nil {
			stage.OnAfterREQIFTYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SOURCE:
		if stage.OnAfterSOURCECreateCallback != nil {
			stage.OnAfterSOURCECreateCallback.OnAfterCreate(stage, target)
		}
	case *SOURCESPECIFICATION:
		if stage.OnAfterSOURCESPECIFICATIONCreateCallback != nil {
			stage.OnAfterSOURCESPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECATTRIBUTES:
		if stage.OnAfterSPECATTRIBUTESCreateCallback != nil {
			stage.OnAfterSPECATTRIBUTESCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECHIERARCHY:
		if stage.OnAfterSPECHIERARCHYCreateCallback != nil {
			stage.OnAfterSPECHIERARCHYCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONCreateCallback != nil {
			stage.OnAfterSPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATIONS:
		if stage.OnAfterSPECIFICATIONSCreateCallback != nil {
			stage.OnAfterSPECIFICATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATIONTYPE:
		if stage.OnAfterSPECIFICATIONTYPECreateCallback != nil {
			stage.OnAfterSPECIFICATIONTYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFIEDVALUES:
		if stage.OnAfterSPECIFIEDVALUESCreateCallback != nil {
			stage.OnAfterSPECIFIEDVALUESCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECOBJECT:
		if stage.OnAfterSPECOBJECTCreateCallback != nil {
			stage.OnAfterSPECOBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECOBJECTS:
		if stage.OnAfterSPECOBJECTSCreateCallback != nil {
			stage.OnAfterSPECOBJECTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECOBJECTTYPE:
		if stage.OnAfterSPECOBJECTTYPECreateCallback != nil {
			stage.OnAfterSPECOBJECTTYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECRELATION:
		if stage.OnAfterSPECRELATIONCreateCallback != nil {
			stage.OnAfterSPECRELATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECRELATIONGROUPS:
		if stage.OnAfterSPECRELATIONGROUPSCreateCallback != nil {
			stage.OnAfterSPECRELATIONGROUPSCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECRELATIONS:
		if stage.OnAfterSPECRELATIONSCreateCallback != nil {
			stage.OnAfterSPECRELATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECRELATIONTYPE:
		if stage.OnAfterSPECRELATIONTYPECreateCallback != nil {
			stage.OnAfterSPECRELATIONTYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECTYPES:
		if stage.OnAfterSPECTYPESCreateCallback != nil {
			stage.OnAfterSPECTYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *TARGET:
		if stage.OnAfterTARGETCreateCallback != nil {
			stage.OnAfterTARGETCreateCallback.OnAfterCreate(stage, target)
		}
	case *TARGETSPECIFICATION:
		if stage.OnAfterTARGETSPECIFICATIONCreateCallback != nil {
			stage.OnAfterTARGETSPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *THEHEADER:
		if stage.OnAfterTHEHEADERCreateCallback != nil {
			stage.OnAfterTHEHEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *TOOLEXTENSIONS:
		if stage.OnAfterTOOLEXTENSIONSCreateCallback != nil {
			stage.OnAfterTOOLEXTENSIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *VALUES:
		if stage.OnAfterVALUESCreateCallback != nil {
			stage.OnAfterVALUESCreateCallback.OnAfterCreate(stage, target)
		}
	case *XHTMLCONTENT:
		if stage.OnAfterXHTMLCONTENTCreateCallback != nil {
			stage.OnAfterXHTMLCONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ALTERNATIVEID:
		newTarget := any(new).(*ALTERNATIVEID)
		if stage.OnAfterALTERNATIVEIDUpdateCallback != nil {
			stage.OnAfterALTERNATIVEIDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONBOOLEAN)
		if stage.OnAfterATTRIBUTEDEFINITIONBOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONBOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONDATE:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONDATE)
		if stage.OnAfterATTRIBUTEDEFINITIONDATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONDATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONENUMERATION)
		if stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONINTEGER:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONINTEGER)
		if stage.OnAfterATTRIBUTEDEFINITIONINTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONINTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONREAL:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONREAL)
		if stage.OnAfterATTRIBUTEDEFINITIONREALUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONREALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONSTRING:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONSTRING)
		if stage.OnAfterATTRIBUTEDEFINITIONSTRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONSTRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEDEFINITIONXHTML:
		newTarget := any(new).(*ATTRIBUTEDEFINITIONXHTML)
		if stage.OnAfterATTRIBUTEDEFINITIONXHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONXHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEBOOLEAN:
		newTarget := any(new).(*ATTRIBUTEVALUEBOOLEAN)
		if stage.OnAfterATTRIBUTEVALUEBOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEBOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEDATE:
		newTarget := any(new).(*ATTRIBUTEVALUEDATE)
		if stage.OnAfterATTRIBUTEVALUEDATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEDATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEENUMERATION:
		newTarget := any(new).(*ATTRIBUTEVALUEENUMERATION)
		if stage.OnAfterATTRIBUTEVALUEENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEINTEGER:
		newTarget := any(new).(*ATTRIBUTEVALUEINTEGER)
		if stage.OnAfterATTRIBUTEVALUEINTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEINTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEREAL:
		newTarget := any(new).(*ATTRIBUTEVALUEREAL)
		if stage.OnAfterATTRIBUTEVALUEREALUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEREALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUESTRING:
		newTarget := any(new).(*ATTRIBUTEVALUESTRING)
		if stage.OnAfterATTRIBUTEVALUESTRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUESTRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTEVALUEXHTML:
		newTarget := any(new).(*ATTRIBUTEVALUEXHTML)
		if stage.OnAfterATTRIBUTEVALUEXHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTEVALUEXHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CHILDREN:
		newTarget := any(new).(*CHILDREN)
		if stage.OnAfterCHILDRENUpdateCallback != nil {
			stage.OnAfterCHILDRENUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CORECONTENT:
		newTarget := any(new).(*CORECONTENT)
		if stage.OnAfterCORECONTENTUpdateCallback != nil {
			stage.OnAfterCORECONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONBOOLEAN:
		newTarget := any(new).(*DATATYPEDEFINITIONBOOLEAN)
		if stage.OnAfterDATATYPEDEFINITIONBOOLEANUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONBOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONDATE:
		newTarget := any(new).(*DATATYPEDEFINITIONDATE)
		if stage.OnAfterDATATYPEDEFINITIONDATEUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONDATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONENUMERATION:
		newTarget := any(new).(*DATATYPEDEFINITIONENUMERATION)
		if stage.OnAfterDATATYPEDEFINITIONENUMERATIONUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONINTEGER:
		newTarget := any(new).(*DATATYPEDEFINITIONINTEGER)
		if stage.OnAfterDATATYPEDEFINITIONINTEGERUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONINTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONREAL:
		newTarget := any(new).(*DATATYPEDEFINITIONREAL)
		if stage.OnAfterDATATYPEDEFINITIONREALUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONREALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONSTRING:
		newTarget := any(new).(*DATATYPEDEFINITIONSTRING)
		if stage.OnAfterDATATYPEDEFINITIONSTRINGUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONSTRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPEDEFINITIONXHTML:
		newTarget := any(new).(*DATATYPEDEFINITIONXHTML)
		if stage.OnAfterDATATYPEDEFINITIONXHTMLUpdateCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONXHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPES:
		newTarget := any(new).(*DATATYPES)
		if stage.OnAfterDATATYPESUpdateCallback != nil {
			stage.OnAfterDATATYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DEFAULTVALUE:
		newTarget := any(new).(*DEFAULTVALUE)
		if stage.OnAfterDEFAULTVALUEUpdateCallback != nil {
			stage.OnAfterDEFAULTVALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DEFINITION:
		newTarget := any(new).(*DEFINITION)
		if stage.OnAfterDEFINITIONUpdateCallback != nil {
			stage.OnAfterDEFINITIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EDITABLEATTS:
		newTarget := any(new).(*EDITABLEATTS)
		if stage.OnAfterEDITABLEATTSUpdateCallback != nil {
			stage.OnAfterEDITABLEATTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EMBEDDEDVALUE:
		newTarget := any(new).(*EMBEDDEDVALUE)
		if stage.OnAfterEMBEDDEDVALUEUpdateCallback != nil {
			stage.OnAfterEMBEDDEDVALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ENUMVALUE:
		newTarget := any(new).(*ENUMVALUE)
		if stage.OnAfterENUMVALUEUpdateCallback != nil {
			stage.OnAfterENUMVALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *OBJECT:
		newTarget := any(new).(*OBJECT)
		if stage.OnAfterOBJECTUpdateCallback != nil {
			stage.OnAfterOBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PROPERTIES:
		newTarget := any(new).(*PROPERTIES)
		if stage.OnAfterPROPERTIESUpdateCallback != nil {
			stage.OnAfterPROPERTIESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATIONGROUP:
		newTarget := any(new).(*RELATIONGROUP)
		if stage.OnAfterRELATIONGROUPUpdateCallback != nil {
			stage.OnAfterRELATIONGROUPUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATIONGROUPTYPE:
		newTarget := any(new).(*RELATIONGROUPTYPE)
		if stage.OnAfterRELATIONGROUPTYPEUpdateCallback != nil {
			stage.OnAfterRELATIONGROUPTYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIF:
		newTarget := any(new).(*REQIF)
		if stage.OnAfterREQIFUpdateCallback != nil {
			stage.OnAfterREQIFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIFCONTENT:
		newTarget := any(new).(*REQIFCONTENT)
		if stage.OnAfterREQIFCONTENTUpdateCallback != nil {
			stage.OnAfterREQIFCONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIFHEADER:
		newTarget := any(new).(*REQIFHEADER)
		if stage.OnAfterREQIFHEADERUpdateCallback != nil {
			stage.OnAfterREQIFHEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIFTOOLEXTENSION:
		newTarget := any(new).(*REQIFTOOLEXTENSION)
		if stage.OnAfterREQIFTOOLEXTENSIONUpdateCallback != nil {
			stage.OnAfterREQIFTOOLEXTENSIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIFTYPE:
		newTarget := any(new).(*REQIFTYPE)
		if stage.OnAfterREQIFTYPEUpdateCallback != nil {
			stage.OnAfterREQIFTYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SOURCE:
		newTarget := any(new).(*SOURCE)
		if stage.OnAfterSOURCEUpdateCallback != nil {
			stage.OnAfterSOURCEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SOURCESPECIFICATION:
		newTarget := any(new).(*SOURCESPECIFICATION)
		if stage.OnAfterSOURCESPECIFICATIONUpdateCallback != nil {
			stage.OnAfterSOURCESPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECATTRIBUTES:
		newTarget := any(new).(*SPECATTRIBUTES)
		if stage.OnAfterSPECATTRIBUTESUpdateCallback != nil {
			stage.OnAfterSPECATTRIBUTESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECHIERARCHY:
		newTarget := any(new).(*SPECHIERARCHY)
		if stage.OnAfterSPECHIERARCHYUpdateCallback != nil {
			stage.OnAfterSPECHIERARCHYUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION:
		newTarget := any(new).(*SPECIFICATION)
		if stage.OnAfterSPECIFICATIONUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATIONS:
		newTarget := any(new).(*SPECIFICATIONS)
		if stage.OnAfterSPECIFICATIONSUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATIONTYPE:
		newTarget := any(new).(*SPECIFICATIONTYPE)
		if stage.OnAfterSPECIFICATIONTYPEUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONTYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFIEDVALUES:
		newTarget := any(new).(*SPECIFIEDVALUES)
		if stage.OnAfterSPECIFIEDVALUESUpdateCallback != nil {
			stage.OnAfterSPECIFIEDVALUESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECOBJECT:
		newTarget := any(new).(*SPECOBJECT)
		if stage.OnAfterSPECOBJECTUpdateCallback != nil {
			stage.OnAfterSPECOBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECOBJECTS:
		newTarget := any(new).(*SPECOBJECTS)
		if stage.OnAfterSPECOBJECTSUpdateCallback != nil {
			stage.OnAfterSPECOBJECTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECOBJECTTYPE:
		newTarget := any(new).(*SPECOBJECTTYPE)
		if stage.OnAfterSPECOBJECTTYPEUpdateCallback != nil {
			stage.OnAfterSPECOBJECTTYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECRELATION:
		newTarget := any(new).(*SPECRELATION)
		if stage.OnAfterSPECRELATIONUpdateCallback != nil {
			stage.OnAfterSPECRELATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECRELATIONGROUPS:
		newTarget := any(new).(*SPECRELATIONGROUPS)
		if stage.OnAfterSPECRELATIONGROUPSUpdateCallback != nil {
			stage.OnAfterSPECRELATIONGROUPSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECRELATIONS:
		newTarget := any(new).(*SPECRELATIONS)
		if stage.OnAfterSPECRELATIONSUpdateCallback != nil {
			stage.OnAfterSPECRELATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECRELATIONTYPE:
		newTarget := any(new).(*SPECRELATIONTYPE)
		if stage.OnAfterSPECRELATIONTYPEUpdateCallback != nil {
			stage.OnAfterSPECRELATIONTYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECTYPES:
		newTarget := any(new).(*SPECTYPES)
		if stage.OnAfterSPECTYPESUpdateCallback != nil {
			stage.OnAfterSPECTYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TARGET:
		newTarget := any(new).(*TARGET)
		if stage.OnAfterTARGETUpdateCallback != nil {
			stage.OnAfterTARGETUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TARGETSPECIFICATION:
		newTarget := any(new).(*TARGETSPECIFICATION)
		if stage.OnAfterTARGETSPECIFICATIONUpdateCallback != nil {
			stage.OnAfterTARGETSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *THEHEADER:
		newTarget := any(new).(*THEHEADER)
		if stage.OnAfterTHEHEADERUpdateCallback != nil {
			stage.OnAfterTHEHEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TOOLEXTENSIONS:
		newTarget := any(new).(*TOOLEXTENSIONS)
		if stage.OnAfterTOOLEXTENSIONSUpdateCallback != nil {
			stage.OnAfterTOOLEXTENSIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VALUES:
		newTarget := any(new).(*VALUES)
		if stage.OnAfterVALUESUpdateCallback != nil {
			stage.OnAfterVALUESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XHTMLCONTENT:
		newTarget := any(new).(*XHTMLCONTENT)
		if stage.OnAfterXHTMLCONTENTUpdateCallback != nil {
			stage.OnAfterXHTMLCONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ALTERNATIVEID:
		if stage.OnAfterALTERNATIVEIDDeleteCallback != nil {
			staged := any(staged).(*ALTERNATIVEID)
			stage.OnAfterALTERNATIVEIDDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		if stage.OnAfterATTRIBUTEDEFINITIONBOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONBOOLEAN)
			stage.OnAfterATTRIBUTEDEFINITIONBOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONDATE:
		if stage.OnAfterATTRIBUTEDEFINITIONDATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONDATE)
			stage.OnAfterATTRIBUTEDEFINITIONDATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		if stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONENUMERATION)
			stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONINTEGER:
		if stage.OnAfterATTRIBUTEDEFINITIONINTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONINTEGER)
			stage.OnAfterATTRIBUTEDEFINITIONINTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONREAL:
		if stage.OnAfterATTRIBUTEDEFINITIONREALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONREAL)
			stage.OnAfterATTRIBUTEDEFINITIONREALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONSTRING:
		if stage.OnAfterATTRIBUTEDEFINITIONSTRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONSTRING)
			stage.OnAfterATTRIBUTEDEFINITIONSTRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEDEFINITIONXHTML:
		if stage.OnAfterATTRIBUTEDEFINITIONXHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEDEFINITIONXHTML)
			stage.OnAfterATTRIBUTEDEFINITIONXHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEBOOLEAN:
		if stage.OnAfterATTRIBUTEVALUEBOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEBOOLEAN)
			stage.OnAfterATTRIBUTEVALUEBOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEDATE:
		if stage.OnAfterATTRIBUTEVALUEDATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEDATE)
			stage.OnAfterATTRIBUTEVALUEDATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEENUMERATION:
		if stage.OnAfterATTRIBUTEVALUEENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEENUMERATION)
			stage.OnAfterATTRIBUTEVALUEENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEINTEGER:
		if stage.OnAfterATTRIBUTEVALUEINTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEINTEGER)
			stage.OnAfterATTRIBUTEVALUEINTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEREAL:
		if stage.OnAfterATTRIBUTEVALUEREALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEREAL)
			stage.OnAfterATTRIBUTEVALUEREALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUESTRING:
		if stage.OnAfterATTRIBUTEVALUESTRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUESTRING)
			stage.OnAfterATTRIBUTEVALUESTRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTEVALUEXHTML:
		if stage.OnAfterATTRIBUTEVALUEXHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTEVALUEXHTML)
			stage.OnAfterATTRIBUTEVALUEXHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CHILDREN:
		if stage.OnAfterCHILDRENDeleteCallback != nil {
			staged := any(staged).(*CHILDREN)
			stage.OnAfterCHILDRENDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CORECONTENT:
		if stage.OnAfterCORECONTENTDeleteCallback != nil {
			staged := any(staged).(*CORECONTENT)
			stage.OnAfterCORECONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONBOOLEAN:
		if stage.OnAfterDATATYPEDEFINITIONBOOLEANDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONBOOLEAN)
			stage.OnAfterDATATYPEDEFINITIONBOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONDATE:
		if stage.OnAfterDATATYPEDEFINITIONDATEDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONDATE)
			stage.OnAfterDATATYPEDEFINITIONDATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONENUMERATION:
		if stage.OnAfterDATATYPEDEFINITIONENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONENUMERATION)
			stage.OnAfterDATATYPEDEFINITIONENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONINTEGER:
		if stage.OnAfterDATATYPEDEFINITIONINTEGERDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONINTEGER)
			stage.OnAfterDATATYPEDEFINITIONINTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONREAL:
		if stage.OnAfterDATATYPEDEFINITIONREALDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONREAL)
			stage.OnAfterDATATYPEDEFINITIONREALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONSTRING:
		if stage.OnAfterDATATYPEDEFINITIONSTRINGDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONSTRING)
			stage.OnAfterDATATYPEDEFINITIONSTRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPEDEFINITIONXHTML:
		if stage.OnAfterDATATYPEDEFINITIONXHTMLDeleteCallback != nil {
			staged := any(staged).(*DATATYPEDEFINITIONXHTML)
			stage.OnAfterDATATYPEDEFINITIONXHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPES:
		if stage.OnAfterDATATYPESDeleteCallback != nil {
			staged := any(staged).(*DATATYPES)
			stage.OnAfterDATATYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DEFAULTVALUE:
		if stage.OnAfterDEFAULTVALUEDeleteCallback != nil {
			staged := any(staged).(*DEFAULTVALUE)
			stage.OnAfterDEFAULTVALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DEFINITION:
		if stage.OnAfterDEFINITIONDeleteCallback != nil {
			staged := any(staged).(*DEFINITION)
			stage.OnAfterDEFINITIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EDITABLEATTS:
		if stage.OnAfterEDITABLEATTSDeleteCallback != nil {
			staged := any(staged).(*EDITABLEATTS)
			stage.OnAfterEDITABLEATTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EMBEDDEDVALUE:
		if stage.OnAfterEMBEDDEDVALUEDeleteCallback != nil {
			staged := any(staged).(*EMBEDDEDVALUE)
			stage.OnAfterEMBEDDEDVALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ENUMVALUE:
		if stage.OnAfterENUMVALUEDeleteCallback != nil {
			staged := any(staged).(*ENUMVALUE)
			stage.OnAfterENUMVALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *OBJECT:
		if stage.OnAfterOBJECTDeleteCallback != nil {
			staged := any(staged).(*OBJECT)
			stage.OnAfterOBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PROPERTIES:
		if stage.OnAfterPROPERTIESDeleteCallback != nil {
			staged := any(staged).(*PROPERTIES)
			stage.OnAfterPROPERTIESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATIONGROUP:
		if stage.OnAfterRELATIONGROUPDeleteCallback != nil {
			staged := any(staged).(*RELATIONGROUP)
			stage.OnAfterRELATIONGROUPDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATIONGROUPTYPE:
		if stage.OnAfterRELATIONGROUPTYPEDeleteCallback != nil {
			staged := any(staged).(*RELATIONGROUPTYPE)
			stage.OnAfterRELATIONGROUPTYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIF:
		if stage.OnAfterREQIFDeleteCallback != nil {
			staged := any(staged).(*REQIF)
			stage.OnAfterREQIFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIFCONTENT:
		if stage.OnAfterREQIFCONTENTDeleteCallback != nil {
			staged := any(staged).(*REQIFCONTENT)
			stage.OnAfterREQIFCONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERDeleteCallback != nil {
			staged := any(staged).(*REQIFHEADER)
			stage.OnAfterREQIFHEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIFTOOLEXTENSION:
		if stage.OnAfterREQIFTOOLEXTENSIONDeleteCallback != nil {
			staged := any(staged).(*REQIFTOOLEXTENSION)
			stage.OnAfterREQIFTOOLEXTENSIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIFTYPE:
		if stage.OnAfterREQIFTYPEDeleteCallback != nil {
			staged := any(staged).(*REQIFTYPE)
			stage.OnAfterREQIFTYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SOURCE:
		if stage.OnAfterSOURCEDeleteCallback != nil {
			staged := any(staged).(*SOURCE)
			stage.OnAfterSOURCEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SOURCESPECIFICATION:
		if stage.OnAfterSOURCESPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*SOURCESPECIFICATION)
			stage.OnAfterSOURCESPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECATTRIBUTES:
		if stage.OnAfterSPECATTRIBUTESDeleteCallback != nil {
			staged := any(staged).(*SPECATTRIBUTES)
			stage.OnAfterSPECATTRIBUTESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECHIERARCHY:
		if stage.OnAfterSPECHIERARCHYDeleteCallback != nil {
			staged := any(staged).(*SPECHIERARCHY)
			stage.OnAfterSPECHIERARCHYDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION)
			stage.OnAfterSPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATIONS:
		if stage.OnAfterSPECIFICATIONSDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATIONS)
			stage.OnAfterSPECIFICATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATIONTYPE:
		if stage.OnAfterSPECIFICATIONTYPEDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATIONTYPE)
			stage.OnAfterSPECIFICATIONTYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFIEDVALUES:
		if stage.OnAfterSPECIFIEDVALUESDeleteCallback != nil {
			staged := any(staged).(*SPECIFIEDVALUES)
			stage.OnAfterSPECIFIEDVALUESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECOBJECT:
		if stage.OnAfterSPECOBJECTDeleteCallback != nil {
			staged := any(staged).(*SPECOBJECT)
			stage.OnAfterSPECOBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECOBJECTS:
		if stage.OnAfterSPECOBJECTSDeleteCallback != nil {
			staged := any(staged).(*SPECOBJECTS)
			stage.OnAfterSPECOBJECTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECOBJECTTYPE:
		if stage.OnAfterSPECOBJECTTYPEDeleteCallback != nil {
			staged := any(staged).(*SPECOBJECTTYPE)
			stage.OnAfterSPECOBJECTTYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECRELATION:
		if stage.OnAfterSPECRELATIONDeleteCallback != nil {
			staged := any(staged).(*SPECRELATION)
			stage.OnAfterSPECRELATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECRELATIONGROUPS:
		if stage.OnAfterSPECRELATIONGROUPSDeleteCallback != nil {
			staged := any(staged).(*SPECRELATIONGROUPS)
			stage.OnAfterSPECRELATIONGROUPSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECRELATIONS:
		if stage.OnAfterSPECRELATIONSDeleteCallback != nil {
			staged := any(staged).(*SPECRELATIONS)
			stage.OnAfterSPECRELATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECRELATIONTYPE:
		if stage.OnAfterSPECRELATIONTYPEDeleteCallback != nil {
			staged := any(staged).(*SPECRELATIONTYPE)
			stage.OnAfterSPECRELATIONTYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECTYPES:
		if stage.OnAfterSPECTYPESDeleteCallback != nil {
			staged := any(staged).(*SPECTYPES)
			stage.OnAfterSPECTYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TARGET:
		if stage.OnAfterTARGETDeleteCallback != nil {
			staged := any(staged).(*TARGET)
			stage.OnAfterTARGETDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TARGETSPECIFICATION:
		if stage.OnAfterTARGETSPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*TARGETSPECIFICATION)
			stage.OnAfterTARGETSPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *THEHEADER:
		if stage.OnAfterTHEHEADERDeleteCallback != nil {
			staged := any(staged).(*THEHEADER)
			stage.OnAfterTHEHEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TOOLEXTENSIONS:
		if stage.OnAfterTOOLEXTENSIONSDeleteCallback != nil {
			staged := any(staged).(*TOOLEXTENSIONS)
			stage.OnAfterTOOLEXTENSIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VALUES:
		if stage.OnAfterVALUESDeleteCallback != nil {
			staged := any(staged).(*VALUES)
			stage.OnAfterVALUESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XHTMLCONTENT:
		if stage.OnAfterXHTMLCONTENTDeleteCallback != nil {
			staged := any(staged).(*XHTMLCONTENT)
			stage.OnAfterXHTMLCONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVEID:
		if stage.OnAfterALTERNATIVEIDReadCallback != nil {
			stage.OnAfterALTERNATIVEIDReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		if stage.OnAfterATTRIBUTEDEFINITIONBOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONBOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONDATE:
		if stage.OnAfterATTRIBUTEDEFINITIONDATEReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONDATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONENUMERATION:
		if stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONINTEGER:
		if stage.OnAfterATTRIBUTEDEFINITIONINTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONINTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONREAL:
		if stage.OnAfterATTRIBUTEDEFINITIONREALReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONREALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONSTRING:
		if stage.OnAfterATTRIBUTEDEFINITIONSTRINGReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONSTRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEDEFINITIONXHTML:
		if stage.OnAfterATTRIBUTEDEFINITIONXHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTEDEFINITIONXHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEBOOLEAN:
		if stage.OnAfterATTRIBUTEVALUEBOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEBOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEDATE:
		if stage.OnAfterATTRIBUTEVALUEDATEReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEDATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEENUMERATION:
		if stage.OnAfterATTRIBUTEVALUEENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEINTEGER:
		if stage.OnAfterATTRIBUTEVALUEINTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEINTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEREAL:
		if stage.OnAfterATTRIBUTEVALUEREALReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEREALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUESTRING:
		if stage.OnAfterATTRIBUTEVALUESTRINGReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUESTRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTEVALUEXHTML:
		if stage.OnAfterATTRIBUTEVALUEXHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTEVALUEXHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *CHILDREN:
		if stage.OnAfterCHILDRENReadCallback != nil {
			stage.OnAfterCHILDRENReadCallback.OnAfterRead(stage, target)
		}
	case *CORECONTENT:
		if stage.OnAfterCORECONTENTReadCallback != nil {
			stage.OnAfterCORECONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONBOOLEAN:
		if stage.OnAfterDATATYPEDEFINITIONBOOLEANReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONBOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONDATE:
		if stage.OnAfterDATATYPEDEFINITIONDATEReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONDATEReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONENUMERATION:
		if stage.OnAfterDATATYPEDEFINITIONENUMERATIONReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONINTEGER:
		if stage.OnAfterDATATYPEDEFINITIONINTEGERReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONINTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONREAL:
		if stage.OnAfterDATATYPEDEFINITIONREALReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONREALReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONSTRING:
		if stage.OnAfterDATATYPEDEFINITIONSTRINGReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONSTRINGReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPEDEFINITIONXHTML:
		if stage.OnAfterDATATYPEDEFINITIONXHTMLReadCallback != nil {
			stage.OnAfterDATATYPEDEFINITIONXHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPES:
		if stage.OnAfterDATATYPESReadCallback != nil {
			stage.OnAfterDATATYPESReadCallback.OnAfterRead(stage, target)
		}
	case *DEFAULTVALUE:
		if stage.OnAfterDEFAULTVALUEReadCallback != nil {
			stage.OnAfterDEFAULTVALUEReadCallback.OnAfterRead(stage, target)
		}
	case *DEFINITION:
		if stage.OnAfterDEFINITIONReadCallback != nil {
			stage.OnAfterDEFINITIONReadCallback.OnAfterRead(stage, target)
		}
	case *EDITABLEATTS:
		if stage.OnAfterEDITABLEATTSReadCallback != nil {
			stage.OnAfterEDITABLEATTSReadCallback.OnAfterRead(stage, target)
		}
	case *EMBEDDEDVALUE:
		if stage.OnAfterEMBEDDEDVALUEReadCallback != nil {
			stage.OnAfterEMBEDDEDVALUEReadCallback.OnAfterRead(stage, target)
		}
	case *ENUMVALUE:
		if stage.OnAfterENUMVALUEReadCallback != nil {
			stage.OnAfterENUMVALUEReadCallback.OnAfterRead(stage, target)
		}
	case *OBJECT:
		if stage.OnAfterOBJECTReadCallback != nil {
			stage.OnAfterOBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *PROPERTIES:
		if stage.OnAfterPROPERTIESReadCallback != nil {
			stage.OnAfterPROPERTIESReadCallback.OnAfterRead(stage, target)
		}
	case *RELATIONGROUP:
		if stage.OnAfterRELATIONGROUPReadCallback != nil {
			stage.OnAfterRELATIONGROUPReadCallback.OnAfterRead(stage, target)
		}
	case *RELATIONGROUPTYPE:
		if stage.OnAfterRELATIONGROUPTYPEReadCallback != nil {
			stage.OnAfterRELATIONGROUPTYPEReadCallback.OnAfterRead(stage, target)
		}
	case *REQIF:
		if stage.OnAfterREQIFReadCallback != nil {
			stage.OnAfterREQIFReadCallback.OnAfterRead(stage, target)
		}
	case *REQIFCONTENT:
		if stage.OnAfterREQIFCONTENTReadCallback != nil {
			stage.OnAfterREQIFCONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERReadCallback != nil {
			stage.OnAfterREQIFHEADERReadCallback.OnAfterRead(stage, target)
		}
	case *REQIFTOOLEXTENSION:
		if stage.OnAfterREQIFTOOLEXTENSIONReadCallback != nil {
			stage.OnAfterREQIFTOOLEXTENSIONReadCallback.OnAfterRead(stage, target)
		}
	case *REQIFTYPE:
		if stage.OnAfterREQIFTYPEReadCallback != nil {
			stage.OnAfterREQIFTYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SOURCE:
		if stage.OnAfterSOURCEReadCallback != nil {
			stage.OnAfterSOURCEReadCallback.OnAfterRead(stage, target)
		}
	case *SOURCESPECIFICATION:
		if stage.OnAfterSOURCESPECIFICATIONReadCallback != nil {
			stage.OnAfterSOURCESPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECATTRIBUTES:
		if stage.OnAfterSPECATTRIBUTESReadCallback != nil {
			stage.OnAfterSPECATTRIBUTESReadCallback.OnAfterRead(stage, target)
		}
	case *SPECHIERARCHY:
		if stage.OnAfterSPECHIERARCHYReadCallback != nil {
			stage.OnAfterSPECHIERARCHYReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONReadCallback != nil {
			stage.OnAfterSPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATIONS:
		if stage.OnAfterSPECIFICATIONSReadCallback != nil {
			stage.OnAfterSPECIFICATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATIONTYPE:
		if stage.OnAfterSPECIFICATIONTYPEReadCallback != nil {
			stage.OnAfterSPECIFICATIONTYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFIEDVALUES:
		if stage.OnAfterSPECIFIEDVALUESReadCallback != nil {
			stage.OnAfterSPECIFIEDVALUESReadCallback.OnAfterRead(stage, target)
		}
	case *SPECOBJECT:
		if stage.OnAfterSPECOBJECTReadCallback != nil {
			stage.OnAfterSPECOBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *SPECOBJECTS:
		if stage.OnAfterSPECOBJECTSReadCallback != nil {
			stage.OnAfterSPECOBJECTSReadCallback.OnAfterRead(stage, target)
		}
	case *SPECOBJECTTYPE:
		if stage.OnAfterSPECOBJECTTYPEReadCallback != nil {
			stage.OnAfterSPECOBJECTTYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPECRELATION:
		if stage.OnAfterSPECRELATIONReadCallback != nil {
			stage.OnAfterSPECRELATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECRELATIONGROUPS:
		if stage.OnAfterSPECRELATIONGROUPSReadCallback != nil {
			stage.OnAfterSPECRELATIONGROUPSReadCallback.OnAfterRead(stage, target)
		}
	case *SPECRELATIONS:
		if stage.OnAfterSPECRELATIONSReadCallback != nil {
			stage.OnAfterSPECRELATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *SPECRELATIONTYPE:
		if stage.OnAfterSPECRELATIONTYPEReadCallback != nil {
			stage.OnAfterSPECRELATIONTYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPECTYPES:
		if stage.OnAfterSPECTYPESReadCallback != nil {
			stage.OnAfterSPECTYPESReadCallback.OnAfterRead(stage, target)
		}
	case *TARGET:
		if stage.OnAfterTARGETReadCallback != nil {
			stage.OnAfterTARGETReadCallback.OnAfterRead(stage, target)
		}
	case *TARGETSPECIFICATION:
		if stage.OnAfterTARGETSPECIFICATIONReadCallback != nil {
			stage.OnAfterTARGETSPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *THEHEADER:
		if stage.OnAfterTHEHEADERReadCallback != nil {
			stage.OnAfterTHEHEADERReadCallback.OnAfterRead(stage, target)
		}
	case *TOOLEXTENSIONS:
		if stage.OnAfterTOOLEXTENSIONSReadCallback != nil {
			stage.OnAfterTOOLEXTENSIONSReadCallback.OnAfterRead(stage, target)
		}
	case *VALUES:
		if stage.OnAfterVALUESReadCallback != nil {
			stage.OnAfterVALUESReadCallback.OnAfterRead(stage, target)
		}
	case *XHTMLCONTENT:
		if stage.OnAfterXHTMLCONTENTReadCallback != nil {
			stage.OnAfterXHTMLCONTENTReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVEID:
		stage.OnAfterALTERNATIVEIDUpdateCallback = any(callback).(OnAfterUpdateInterface[ALTERNATIVEID])
	
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.OnAfterATTRIBUTEDEFINITIONBOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONBOOLEAN])
	
	case *ATTRIBUTEDEFINITIONDATE:
		stage.OnAfterATTRIBUTEDEFINITIONDATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONDATE])
	
	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONENUMERATION])
	
	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.OnAfterATTRIBUTEDEFINITIONINTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONINTEGER])
	
	case *ATTRIBUTEDEFINITIONREAL:
		stage.OnAfterATTRIBUTEDEFINITIONREALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONREAL])
	
	case *ATTRIBUTEDEFINITIONSTRING:
		stage.OnAfterATTRIBUTEDEFINITIONSTRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONSTRING])
	
	case *ATTRIBUTEDEFINITIONXHTML:
		stage.OnAfterATTRIBUTEDEFINITIONXHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEDEFINITIONXHTML])
	
	case *ATTRIBUTEVALUEBOOLEAN:
		stage.OnAfterATTRIBUTEVALUEBOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEBOOLEAN])
	
	case *ATTRIBUTEVALUEDATE:
		stage.OnAfterATTRIBUTEVALUEDATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEDATE])
	
	case *ATTRIBUTEVALUEENUMERATION:
		stage.OnAfterATTRIBUTEVALUEENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEENUMERATION])
	
	case *ATTRIBUTEVALUEINTEGER:
		stage.OnAfterATTRIBUTEVALUEINTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEINTEGER])
	
	case *ATTRIBUTEVALUEREAL:
		stage.OnAfterATTRIBUTEVALUEREALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEREAL])
	
	case *ATTRIBUTEVALUESTRING:
		stage.OnAfterATTRIBUTEVALUESTRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUESTRING])
	
	case *ATTRIBUTEVALUEXHTML:
		stage.OnAfterATTRIBUTEVALUEXHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTEVALUEXHTML])
	
	case *CHILDREN:
		stage.OnAfterCHILDRENUpdateCallback = any(callback).(OnAfterUpdateInterface[CHILDREN])
	
	case *CORECONTENT:
		stage.OnAfterCORECONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[CORECONTENT])
	
	case *DATATYPEDEFINITIONBOOLEAN:
		stage.OnAfterDATATYPEDEFINITIONBOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONBOOLEAN])
	
	case *DATATYPEDEFINITIONDATE:
		stage.OnAfterDATATYPEDEFINITIONDATEUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONDATE])
	
	case *DATATYPEDEFINITIONENUMERATION:
		stage.OnAfterDATATYPEDEFINITIONENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONENUMERATION])
	
	case *DATATYPEDEFINITIONINTEGER:
		stage.OnAfterDATATYPEDEFINITIONINTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONINTEGER])
	
	case *DATATYPEDEFINITIONREAL:
		stage.OnAfterDATATYPEDEFINITIONREALUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONREAL])
	
	case *DATATYPEDEFINITIONSTRING:
		stage.OnAfterDATATYPEDEFINITIONSTRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONSTRING])
	
	case *DATATYPEDEFINITIONXHTML:
		stage.OnAfterDATATYPEDEFINITIONXHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPEDEFINITIONXHTML])
	
	case *DATATYPES:
		stage.OnAfterDATATYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPES])
	
	case *DEFAULTVALUE:
		stage.OnAfterDEFAULTVALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[DEFAULTVALUE])
	
	case *DEFINITION:
		stage.OnAfterDEFINITIONUpdateCallback = any(callback).(OnAfterUpdateInterface[DEFINITION])
	
	case *EDITABLEATTS:
		stage.OnAfterEDITABLEATTSUpdateCallback = any(callback).(OnAfterUpdateInterface[EDITABLEATTS])
	
	case *EMBEDDEDVALUE:
		stage.OnAfterEMBEDDEDVALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[EMBEDDEDVALUE])
	
	case *ENUMVALUE:
		stage.OnAfterENUMVALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[ENUMVALUE])
	
	case *OBJECT:
		stage.OnAfterOBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[OBJECT])
	
	case *PROPERTIES:
		stage.OnAfterPROPERTIESUpdateCallback = any(callback).(OnAfterUpdateInterface[PROPERTIES])
	
	case *RELATIONGROUP:
		stage.OnAfterRELATIONGROUPUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATIONGROUP])
	
	case *RELATIONGROUPTYPE:
		stage.OnAfterRELATIONGROUPTYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATIONGROUPTYPE])
	
	case *REQIF:
		stage.OnAfterREQIFUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIF])
	
	case *REQIFCONTENT:
		stage.OnAfterREQIFCONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIFCONTENT])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIFHEADER])
	
	case *REQIFTOOLEXTENSION:
		stage.OnAfterREQIFTOOLEXTENSIONUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIFTOOLEXTENSION])
	
	case *REQIFTYPE:
		stage.OnAfterREQIFTYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIFTYPE])
	
	case *SOURCE:
		stage.OnAfterSOURCEUpdateCallback = any(callback).(OnAfterUpdateInterface[SOURCE])
	
	case *SOURCESPECIFICATION:
		stage.OnAfterSOURCESPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SOURCESPECIFICATION])
	
	case *SPECATTRIBUTES:
		stage.OnAfterSPECATTRIBUTESUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECATTRIBUTES])
	
	case *SPECHIERARCHY:
		stage.OnAfterSPECHIERARCHYUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECHIERARCHY])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION])
	
	case *SPECIFICATIONS:
		stage.OnAfterSPECIFICATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATIONS])
	
	case *SPECIFICATIONTYPE:
		stage.OnAfterSPECIFICATIONTYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATIONTYPE])
	
	case *SPECIFIEDVALUES:
		stage.OnAfterSPECIFIEDVALUESUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFIEDVALUES])
	
	case *SPECOBJECT:
		stage.OnAfterSPECOBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECOBJECT])
	
	case *SPECOBJECTS:
		stage.OnAfterSPECOBJECTSUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECOBJECTS])
	
	case *SPECOBJECTTYPE:
		stage.OnAfterSPECOBJECTTYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECOBJECTTYPE])
	
	case *SPECRELATION:
		stage.OnAfterSPECRELATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECRELATION])
	
	case *SPECRELATIONGROUPS:
		stage.OnAfterSPECRELATIONGROUPSUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECRELATIONGROUPS])
	
	case *SPECRELATIONS:
		stage.OnAfterSPECRELATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECRELATIONS])
	
	case *SPECRELATIONTYPE:
		stage.OnAfterSPECRELATIONTYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECRELATIONTYPE])
	
	case *SPECTYPES:
		stage.OnAfterSPECTYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECTYPES])
	
	case *TARGET:
		stage.OnAfterTARGETUpdateCallback = any(callback).(OnAfterUpdateInterface[TARGET])
	
	case *TARGETSPECIFICATION:
		stage.OnAfterTARGETSPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[TARGETSPECIFICATION])
	
	case *THEHEADER:
		stage.OnAfterTHEHEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[THEHEADER])
	
	case *TOOLEXTENSIONS:
		stage.OnAfterTOOLEXTENSIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[TOOLEXTENSIONS])
	
	case *VALUES:
		stage.OnAfterVALUESUpdateCallback = any(callback).(OnAfterUpdateInterface[VALUES])
	
	case *XHTMLCONTENT:
		stage.OnAfterXHTMLCONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[XHTMLCONTENT])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVEID:
		stage.OnAfterALTERNATIVEIDCreateCallback = any(callback).(OnAfterCreateInterface[ALTERNATIVEID])
	
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.OnAfterATTRIBUTEDEFINITIONBOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONBOOLEAN])
	
	case *ATTRIBUTEDEFINITIONDATE:
		stage.OnAfterATTRIBUTEDEFINITIONDATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONDATE])
	
	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONENUMERATION])
	
	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.OnAfterATTRIBUTEDEFINITIONINTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONINTEGER])
	
	case *ATTRIBUTEDEFINITIONREAL:
		stage.OnAfterATTRIBUTEDEFINITIONREALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONREAL])
	
	case *ATTRIBUTEDEFINITIONSTRING:
		stage.OnAfterATTRIBUTEDEFINITIONSTRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONSTRING])
	
	case *ATTRIBUTEDEFINITIONXHTML:
		stage.OnAfterATTRIBUTEDEFINITIONXHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEDEFINITIONXHTML])
	
	case *ATTRIBUTEVALUEBOOLEAN:
		stage.OnAfterATTRIBUTEVALUEBOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEBOOLEAN])
	
	case *ATTRIBUTEVALUEDATE:
		stage.OnAfterATTRIBUTEVALUEDATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEDATE])
	
	case *ATTRIBUTEVALUEENUMERATION:
		stage.OnAfterATTRIBUTEVALUEENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEENUMERATION])
	
	case *ATTRIBUTEVALUEINTEGER:
		stage.OnAfterATTRIBUTEVALUEINTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEINTEGER])
	
	case *ATTRIBUTEVALUEREAL:
		stage.OnAfterATTRIBUTEVALUEREALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEREAL])
	
	case *ATTRIBUTEVALUESTRING:
		stage.OnAfterATTRIBUTEVALUESTRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUESTRING])
	
	case *ATTRIBUTEVALUEXHTML:
		stage.OnAfterATTRIBUTEVALUEXHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTEVALUEXHTML])
	
	case *CHILDREN:
		stage.OnAfterCHILDRENCreateCallback = any(callback).(OnAfterCreateInterface[CHILDREN])
	
	case *CORECONTENT:
		stage.OnAfterCORECONTENTCreateCallback = any(callback).(OnAfterCreateInterface[CORECONTENT])
	
	case *DATATYPEDEFINITIONBOOLEAN:
		stage.OnAfterDATATYPEDEFINITIONBOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONBOOLEAN])
	
	case *DATATYPEDEFINITIONDATE:
		stage.OnAfterDATATYPEDEFINITIONDATECreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONDATE])
	
	case *DATATYPEDEFINITIONENUMERATION:
		stage.OnAfterDATATYPEDEFINITIONENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONENUMERATION])
	
	case *DATATYPEDEFINITIONINTEGER:
		stage.OnAfterDATATYPEDEFINITIONINTEGERCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONINTEGER])
	
	case *DATATYPEDEFINITIONREAL:
		stage.OnAfterDATATYPEDEFINITIONREALCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONREAL])
	
	case *DATATYPEDEFINITIONSTRING:
		stage.OnAfterDATATYPEDEFINITIONSTRINGCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONSTRING])
	
	case *DATATYPEDEFINITIONXHTML:
		stage.OnAfterDATATYPEDEFINITIONXHTMLCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPEDEFINITIONXHTML])
	
	case *DATATYPES:
		stage.OnAfterDATATYPESCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPES])
	
	case *DEFAULTVALUE:
		stage.OnAfterDEFAULTVALUECreateCallback = any(callback).(OnAfterCreateInterface[DEFAULTVALUE])
	
	case *DEFINITION:
		stage.OnAfterDEFINITIONCreateCallback = any(callback).(OnAfterCreateInterface[DEFINITION])
	
	case *EDITABLEATTS:
		stage.OnAfterEDITABLEATTSCreateCallback = any(callback).(OnAfterCreateInterface[EDITABLEATTS])
	
	case *EMBEDDEDVALUE:
		stage.OnAfterEMBEDDEDVALUECreateCallback = any(callback).(OnAfterCreateInterface[EMBEDDEDVALUE])
	
	case *ENUMVALUE:
		stage.OnAfterENUMVALUECreateCallback = any(callback).(OnAfterCreateInterface[ENUMVALUE])
	
	case *OBJECT:
		stage.OnAfterOBJECTCreateCallback = any(callback).(OnAfterCreateInterface[OBJECT])
	
	case *PROPERTIES:
		stage.OnAfterPROPERTIESCreateCallback = any(callback).(OnAfterCreateInterface[PROPERTIES])
	
	case *RELATIONGROUP:
		stage.OnAfterRELATIONGROUPCreateCallback = any(callback).(OnAfterCreateInterface[RELATIONGROUP])
	
	case *RELATIONGROUPTYPE:
		stage.OnAfterRELATIONGROUPTYPECreateCallback = any(callback).(OnAfterCreateInterface[RELATIONGROUPTYPE])
	
	case *REQIF:
		stage.OnAfterREQIFCreateCallback = any(callback).(OnAfterCreateInterface[REQIF])
	
	case *REQIFCONTENT:
		stage.OnAfterREQIFCONTENTCreateCallback = any(callback).(OnAfterCreateInterface[REQIFCONTENT])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERCreateCallback = any(callback).(OnAfterCreateInterface[REQIFHEADER])
	
	case *REQIFTOOLEXTENSION:
		stage.OnAfterREQIFTOOLEXTENSIONCreateCallback = any(callback).(OnAfterCreateInterface[REQIFTOOLEXTENSION])
	
	case *REQIFTYPE:
		stage.OnAfterREQIFTYPECreateCallback = any(callback).(OnAfterCreateInterface[REQIFTYPE])
	
	case *SOURCE:
		stage.OnAfterSOURCECreateCallback = any(callback).(OnAfterCreateInterface[SOURCE])
	
	case *SOURCESPECIFICATION:
		stage.OnAfterSOURCESPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[SOURCESPECIFICATION])
	
	case *SPECATTRIBUTES:
		stage.OnAfterSPECATTRIBUTESCreateCallback = any(callback).(OnAfterCreateInterface[SPECATTRIBUTES])
	
	case *SPECHIERARCHY:
		stage.OnAfterSPECHIERARCHYCreateCallback = any(callback).(OnAfterCreateInterface[SPECHIERARCHY])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION])
	
	case *SPECIFICATIONS:
		stage.OnAfterSPECIFICATIONSCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATIONS])
	
	case *SPECIFICATIONTYPE:
		stage.OnAfterSPECIFICATIONTYPECreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATIONTYPE])
	
	case *SPECIFIEDVALUES:
		stage.OnAfterSPECIFIEDVALUESCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFIEDVALUES])
	
	case *SPECOBJECT:
		stage.OnAfterSPECOBJECTCreateCallback = any(callback).(OnAfterCreateInterface[SPECOBJECT])
	
	case *SPECOBJECTS:
		stage.OnAfterSPECOBJECTSCreateCallback = any(callback).(OnAfterCreateInterface[SPECOBJECTS])
	
	case *SPECOBJECTTYPE:
		stage.OnAfterSPECOBJECTTYPECreateCallback = any(callback).(OnAfterCreateInterface[SPECOBJECTTYPE])
	
	case *SPECRELATION:
		stage.OnAfterSPECRELATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPECRELATION])
	
	case *SPECRELATIONGROUPS:
		stage.OnAfterSPECRELATIONGROUPSCreateCallback = any(callback).(OnAfterCreateInterface[SPECRELATIONGROUPS])
	
	case *SPECRELATIONS:
		stage.OnAfterSPECRELATIONSCreateCallback = any(callback).(OnAfterCreateInterface[SPECRELATIONS])
	
	case *SPECRELATIONTYPE:
		stage.OnAfterSPECRELATIONTYPECreateCallback = any(callback).(OnAfterCreateInterface[SPECRELATIONTYPE])
	
	case *SPECTYPES:
		stage.OnAfterSPECTYPESCreateCallback = any(callback).(OnAfterCreateInterface[SPECTYPES])
	
	case *TARGET:
		stage.OnAfterTARGETCreateCallback = any(callback).(OnAfterCreateInterface[TARGET])
	
	case *TARGETSPECIFICATION:
		stage.OnAfterTARGETSPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[TARGETSPECIFICATION])
	
	case *THEHEADER:
		stage.OnAfterTHEHEADERCreateCallback = any(callback).(OnAfterCreateInterface[THEHEADER])
	
	case *TOOLEXTENSIONS:
		stage.OnAfterTOOLEXTENSIONSCreateCallback = any(callback).(OnAfterCreateInterface[TOOLEXTENSIONS])
	
	case *VALUES:
		stage.OnAfterVALUESCreateCallback = any(callback).(OnAfterCreateInterface[VALUES])
	
	case *XHTMLCONTENT:
		stage.OnAfterXHTMLCONTENTCreateCallback = any(callback).(OnAfterCreateInterface[XHTMLCONTENT])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVEID:
		stage.OnAfterALTERNATIVEIDDeleteCallback = any(callback).(OnAfterDeleteInterface[ALTERNATIVEID])
	
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.OnAfterATTRIBUTEDEFINITIONBOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONBOOLEAN])
	
	case *ATTRIBUTEDEFINITIONDATE:
		stage.OnAfterATTRIBUTEDEFINITIONDATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONDATE])
	
	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONENUMERATION])
	
	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.OnAfterATTRIBUTEDEFINITIONINTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONINTEGER])
	
	case *ATTRIBUTEDEFINITIONREAL:
		stage.OnAfterATTRIBUTEDEFINITIONREALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONREAL])
	
	case *ATTRIBUTEDEFINITIONSTRING:
		stage.OnAfterATTRIBUTEDEFINITIONSTRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONSTRING])
	
	case *ATTRIBUTEDEFINITIONXHTML:
		stage.OnAfterATTRIBUTEDEFINITIONXHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEDEFINITIONXHTML])
	
	case *ATTRIBUTEVALUEBOOLEAN:
		stage.OnAfterATTRIBUTEVALUEBOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEBOOLEAN])
	
	case *ATTRIBUTEVALUEDATE:
		stage.OnAfterATTRIBUTEVALUEDATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEDATE])
	
	case *ATTRIBUTEVALUEENUMERATION:
		stage.OnAfterATTRIBUTEVALUEENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEENUMERATION])
	
	case *ATTRIBUTEVALUEINTEGER:
		stage.OnAfterATTRIBUTEVALUEINTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEINTEGER])
	
	case *ATTRIBUTEVALUEREAL:
		stage.OnAfterATTRIBUTEVALUEREALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEREAL])
	
	case *ATTRIBUTEVALUESTRING:
		stage.OnAfterATTRIBUTEVALUESTRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUESTRING])
	
	case *ATTRIBUTEVALUEXHTML:
		stage.OnAfterATTRIBUTEVALUEXHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTEVALUEXHTML])
	
	case *CHILDREN:
		stage.OnAfterCHILDRENDeleteCallback = any(callback).(OnAfterDeleteInterface[CHILDREN])
	
	case *CORECONTENT:
		stage.OnAfterCORECONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[CORECONTENT])
	
	case *DATATYPEDEFINITIONBOOLEAN:
		stage.OnAfterDATATYPEDEFINITIONBOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONBOOLEAN])
	
	case *DATATYPEDEFINITIONDATE:
		stage.OnAfterDATATYPEDEFINITIONDATEDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONDATE])
	
	case *DATATYPEDEFINITIONENUMERATION:
		stage.OnAfterDATATYPEDEFINITIONENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONENUMERATION])
	
	case *DATATYPEDEFINITIONINTEGER:
		stage.OnAfterDATATYPEDEFINITIONINTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONINTEGER])
	
	case *DATATYPEDEFINITIONREAL:
		stage.OnAfterDATATYPEDEFINITIONREALDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONREAL])
	
	case *DATATYPEDEFINITIONSTRING:
		stage.OnAfterDATATYPEDEFINITIONSTRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONSTRING])
	
	case *DATATYPEDEFINITIONXHTML:
		stage.OnAfterDATATYPEDEFINITIONXHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPEDEFINITIONXHTML])
	
	case *DATATYPES:
		stage.OnAfterDATATYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPES])
	
	case *DEFAULTVALUE:
		stage.OnAfterDEFAULTVALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[DEFAULTVALUE])
	
	case *DEFINITION:
		stage.OnAfterDEFINITIONDeleteCallback = any(callback).(OnAfterDeleteInterface[DEFINITION])
	
	case *EDITABLEATTS:
		stage.OnAfterEDITABLEATTSDeleteCallback = any(callback).(OnAfterDeleteInterface[EDITABLEATTS])
	
	case *EMBEDDEDVALUE:
		stage.OnAfterEMBEDDEDVALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[EMBEDDEDVALUE])
	
	case *ENUMVALUE:
		stage.OnAfterENUMVALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[ENUMVALUE])
	
	case *OBJECT:
		stage.OnAfterOBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[OBJECT])
	
	case *PROPERTIES:
		stage.OnAfterPROPERTIESDeleteCallback = any(callback).(OnAfterDeleteInterface[PROPERTIES])
	
	case *RELATIONGROUP:
		stage.OnAfterRELATIONGROUPDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATIONGROUP])
	
	case *RELATIONGROUPTYPE:
		stage.OnAfterRELATIONGROUPTYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATIONGROUPTYPE])
	
	case *REQIF:
		stage.OnAfterREQIFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIF])
	
	case *REQIFCONTENT:
		stage.OnAfterREQIFCONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIFCONTENT])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIFHEADER])
	
	case *REQIFTOOLEXTENSION:
		stage.OnAfterREQIFTOOLEXTENSIONDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIFTOOLEXTENSION])
	
	case *REQIFTYPE:
		stage.OnAfterREQIFTYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIFTYPE])
	
	case *SOURCE:
		stage.OnAfterSOURCEDeleteCallback = any(callback).(OnAfterDeleteInterface[SOURCE])
	
	case *SOURCESPECIFICATION:
		stage.OnAfterSOURCESPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SOURCESPECIFICATION])
	
	case *SPECATTRIBUTES:
		stage.OnAfterSPECATTRIBUTESDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECATTRIBUTES])
	
	case *SPECHIERARCHY:
		stage.OnAfterSPECHIERARCHYDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECHIERARCHY])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION])
	
	case *SPECIFICATIONS:
		stage.OnAfterSPECIFICATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATIONS])
	
	case *SPECIFICATIONTYPE:
		stage.OnAfterSPECIFICATIONTYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATIONTYPE])
	
	case *SPECIFIEDVALUES:
		stage.OnAfterSPECIFIEDVALUESDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFIEDVALUES])
	
	case *SPECOBJECT:
		stage.OnAfterSPECOBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECOBJECT])
	
	case *SPECOBJECTS:
		stage.OnAfterSPECOBJECTSDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECOBJECTS])
	
	case *SPECOBJECTTYPE:
		stage.OnAfterSPECOBJECTTYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECOBJECTTYPE])
	
	case *SPECRELATION:
		stage.OnAfterSPECRELATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECRELATION])
	
	case *SPECRELATIONGROUPS:
		stage.OnAfterSPECRELATIONGROUPSDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECRELATIONGROUPS])
	
	case *SPECRELATIONS:
		stage.OnAfterSPECRELATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECRELATIONS])
	
	case *SPECRELATIONTYPE:
		stage.OnAfterSPECRELATIONTYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECRELATIONTYPE])
	
	case *SPECTYPES:
		stage.OnAfterSPECTYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECTYPES])
	
	case *TARGET:
		stage.OnAfterTARGETDeleteCallback = any(callback).(OnAfterDeleteInterface[TARGET])
	
	case *TARGETSPECIFICATION:
		stage.OnAfterTARGETSPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[TARGETSPECIFICATION])
	
	case *THEHEADER:
		stage.OnAfterTHEHEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[THEHEADER])
	
	case *TOOLEXTENSIONS:
		stage.OnAfterTOOLEXTENSIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[TOOLEXTENSIONS])
	
	case *VALUES:
		stage.OnAfterVALUESDeleteCallback = any(callback).(OnAfterDeleteInterface[VALUES])
	
	case *XHTMLCONTENT:
		stage.OnAfterXHTMLCONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[XHTMLCONTENT])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVEID:
		stage.OnAfterALTERNATIVEIDReadCallback = any(callback).(OnAfterReadInterface[ALTERNATIVEID])
	
	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.OnAfterATTRIBUTEDEFINITIONBOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONBOOLEAN])
	
	case *ATTRIBUTEDEFINITIONDATE:
		stage.OnAfterATTRIBUTEDEFINITIONDATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONDATE])
	
	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.OnAfterATTRIBUTEDEFINITIONENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONENUMERATION])
	
	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.OnAfterATTRIBUTEDEFINITIONINTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONINTEGER])
	
	case *ATTRIBUTEDEFINITIONREAL:
		stage.OnAfterATTRIBUTEDEFINITIONREALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONREAL])
	
	case *ATTRIBUTEDEFINITIONSTRING:
		stage.OnAfterATTRIBUTEDEFINITIONSTRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONSTRING])
	
	case *ATTRIBUTEDEFINITIONXHTML:
		stage.OnAfterATTRIBUTEDEFINITIONXHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEDEFINITIONXHTML])
	
	case *ATTRIBUTEVALUEBOOLEAN:
		stage.OnAfterATTRIBUTEVALUEBOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEBOOLEAN])
	
	case *ATTRIBUTEVALUEDATE:
		stage.OnAfterATTRIBUTEVALUEDATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEDATE])
	
	case *ATTRIBUTEVALUEENUMERATION:
		stage.OnAfterATTRIBUTEVALUEENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEENUMERATION])
	
	case *ATTRIBUTEVALUEINTEGER:
		stage.OnAfterATTRIBUTEVALUEINTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEINTEGER])
	
	case *ATTRIBUTEVALUEREAL:
		stage.OnAfterATTRIBUTEVALUEREALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEREAL])
	
	case *ATTRIBUTEVALUESTRING:
		stage.OnAfterATTRIBUTEVALUESTRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUESTRING])
	
	case *ATTRIBUTEVALUEXHTML:
		stage.OnAfterATTRIBUTEVALUEXHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTEVALUEXHTML])
	
	case *CHILDREN:
		stage.OnAfterCHILDRENReadCallback = any(callback).(OnAfterReadInterface[CHILDREN])
	
	case *CORECONTENT:
		stage.OnAfterCORECONTENTReadCallback = any(callback).(OnAfterReadInterface[CORECONTENT])
	
	case *DATATYPEDEFINITIONBOOLEAN:
		stage.OnAfterDATATYPEDEFINITIONBOOLEANReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONBOOLEAN])
	
	case *DATATYPEDEFINITIONDATE:
		stage.OnAfterDATATYPEDEFINITIONDATEReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONDATE])
	
	case *DATATYPEDEFINITIONENUMERATION:
		stage.OnAfterDATATYPEDEFINITIONENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONENUMERATION])
	
	case *DATATYPEDEFINITIONINTEGER:
		stage.OnAfterDATATYPEDEFINITIONINTEGERReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONINTEGER])
	
	case *DATATYPEDEFINITIONREAL:
		stage.OnAfterDATATYPEDEFINITIONREALReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONREAL])
	
	case *DATATYPEDEFINITIONSTRING:
		stage.OnAfterDATATYPEDEFINITIONSTRINGReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONSTRING])
	
	case *DATATYPEDEFINITIONXHTML:
		stage.OnAfterDATATYPEDEFINITIONXHTMLReadCallback = any(callback).(OnAfterReadInterface[DATATYPEDEFINITIONXHTML])
	
	case *DATATYPES:
		stage.OnAfterDATATYPESReadCallback = any(callback).(OnAfterReadInterface[DATATYPES])
	
	case *DEFAULTVALUE:
		stage.OnAfterDEFAULTVALUEReadCallback = any(callback).(OnAfterReadInterface[DEFAULTVALUE])
	
	case *DEFINITION:
		stage.OnAfterDEFINITIONReadCallback = any(callback).(OnAfterReadInterface[DEFINITION])
	
	case *EDITABLEATTS:
		stage.OnAfterEDITABLEATTSReadCallback = any(callback).(OnAfterReadInterface[EDITABLEATTS])
	
	case *EMBEDDEDVALUE:
		stage.OnAfterEMBEDDEDVALUEReadCallback = any(callback).(OnAfterReadInterface[EMBEDDEDVALUE])
	
	case *ENUMVALUE:
		stage.OnAfterENUMVALUEReadCallback = any(callback).(OnAfterReadInterface[ENUMVALUE])
	
	case *OBJECT:
		stage.OnAfterOBJECTReadCallback = any(callback).(OnAfterReadInterface[OBJECT])
	
	case *PROPERTIES:
		stage.OnAfterPROPERTIESReadCallback = any(callback).(OnAfterReadInterface[PROPERTIES])
	
	case *RELATIONGROUP:
		stage.OnAfterRELATIONGROUPReadCallback = any(callback).(OnAfterReadInterface[RELATIONGROUP])
	
	case *RELATIONGROUPTYPE:
		stage.OnAfterRELATIONGROUPTYPEReadCallback = any(callback).(OnAfterReadInterface[RELATIONGROUPTYPE])
	
	case *REQIF:
		stage.OnAfterREQIFReadCallback = any(callback).(OnAfterReadInterface[REQIF])
	
	case *REQIFCONTENT:
		stage.OnAfterREQIFCONTENTReadCallback = any(callback).(OnAfterReadInterface[REQIFCONTENT])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERReadCallback = any(callback).(OnAfterReadInterface[REQIFHEADER])
	
	case *REQIFTOOLEXTENSION:
		stage.OnAfterREQIFTOOLEXTENSIONReadCallback = any(callback).(OnAfterReadInterface[REQIFTOOLEXTENSION])
	
	case *REQIFTYPE:
		stage.OnAfterREQIFTYPEReadCallback = any(callback).(OnAfterReadInterface[REQIFTYPE])
	
	case *SOURCE:
		stage.OnAfterSOURCEReadCallback = any(callback).(OnAfterReadInterface[SOURCE])
	
	case *SOURCESPECIFICATION:
		stage.OnAfterSOURCESPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[SOURCESPECIFICATION])
	
	case *SPECATTRIBUTES:
		stage.OnAfterSPECATTRIBUTESReadCallback = any(callback).(OnAfterReadInterface[SPECATTRIBUTES])
	
	case *SPECHIERARCHY:
		stage.OnAfterSPECHIERARCHYReadCallback = any(callback).(OnAfterReadInterface[SPECHIERARCHY])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION])
	
	case *SPECIFICATIONS:
		stage.OnAfterSPECIFICATIONSReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATIONS])
	
	case *SPECIFICATIONTYPE:
		stage.OnAfterSPECIFICATIONTYPEReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATIONTYPE])
	
	case *SPECIFIEDVALUES:
		stage.OnAfterSPECIFIEDVALUESReadCallback = any(callback).(OnAfterReadInterface[SPECIFIEDVALUES])
	
	case *SPECOBJECT:
		stage.OnAfterSPECOBJECTReadCallback = any(callback).(OnAfterReadInterface[SPECOBJECT])
	
	case *SPECOBJECTS:
		stage.OnAfterSPECOBJECTSReadCallback = any(callback).(OnAfterReadInterface[SPECOBJECTS])
	
	case *SPECOBJECTTYPE:
		stage.OnAfterSPECOBJECTTYPEReadCallback = any(callback).(OnAfterReadInterface[SPECOBJECTTYPE])
	
	case *SPECRELATION:
		stage.OnAfterSPECRELATIONReadCallback = any(callback).(OnAfterReadInterface[SPECRELATION])
	
	case *SPECRELATIONGROUPS:
		stage.OnAfterSPECRELATIONGROUPSReadCallback = any(callback).(OnAfterReadInterface[SPECRELATIONGROUPS])
	
	case *SPECRELATIONS:
		stage.OnAfterSPECRELATIONSReadCallback = any(callback).(OnAfterReadInterface[SPECRELATIONS])
	
	case *SPECRELATIONTYPE:
		stage.OnAfterSPECRELATIONTYPEReadCallback = any(callback).(OnAfterReadInterface[SPECRELATIONTYPE])
	
	case *SPECTYPES:
		stage.OnAfterSPECTYPESReadCallback = any(callback).(OnAfterReadInterface[SPECTYPES])
	
	case *TARGET:
		stage.OnAfterTARGETReadCallback = any(callback).(OnAfterReadInterface[TARGET])
	
	case *TARGETSPECIFICATION:
		stage.OnAfterTARGETSPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[TARGETSPECIFICATION])
	
	case *THEHEADER:
		stage.OnAfterTHEHEADERReadCallback = any(callback).(OnAfterReadInterface[THEHEADER])
	
	case *TOOLEXTENSIONS:
		stage.OnAfterTOOLEXTENSIONSReadCallback = any(callback).(OnAfterReadInterface[TOOLEXTENSIONS])
	
	case *VALUES:
		stage.OnAfterVALUESReadCallback = any(callback).(OnAfterReadInterface[VALUES])
	
	case *XHTMLCONTENT:
		stage.OnAfterXHTMLCONTENTReadCallback = any(callback).(OnAfterReadInterface[XHTMLCONTENT])
	
	}
}
