// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ALTERNATIVEID:
		ok = stage.IsStagedALTERNATIVEID(target)

	case *ATTRIBUTEDEFINITIONBOOLEAN:
		ok = stage.IsStagedATTRIBUTEDEFINITIONBOOLEAN(target)

	case *ATTRIBUTEDEFINITIONDATE:
		ok = stage.IsStagedATTRIBUTEDEFINITIONDATE(target)

	case *ATTRIBUTEDEFINITIONENUMERATION:
		ok = stage.IsStagedATTRIBUTEDEFINITIONENUMERATION(target)

	case *ATTRIBUTEDEFINITIONINTEGER:
		ok = stage.IsStagedATTRIBUTEDEFINITIONINTEGER(target)

	case *ATTRIBUTEDEFINITIONREAL:
		ok = stage.IsStagedATTRIBUTEDEFINITIONREAL(target)

	case *ATTRIBUTEDEFINITIONSTRING:
		ok = stage.IsStagedATTRIBUTEDEFINITIONSTRING(target)

	case *ATTRIBUTEDEFINITIONXHTML:
		ok = stage.IsStagedATTRIBUTEDEFINITIONXHTML(target)

	case *ATTRIBUTEVALUEBOOLEAN:
		ok = stage.IsStagedATTRIBUTEVALUEBOOLEAN(target)

	case *ATTRIBUTEVALUEDATE:
		ok = stage.IsStagedATTRIBUTEVALUEDATE(target)

	case *ATTRIBUTEVALUEENUMERATION:
		ok = stage.IsStagedATTRIBUTEVALUEENUMERATION(target)

	case *ATTRIBUTEVALUEINTEGER:
		ok = stage.IsStagedATTRIBUTEVALUEINTEGER(target)

	case *ATTRIBUTEVALUEREAL:
		ok = stage.IsStagedATTRIBUTEVALUEREAL(target)

	case *ATTRIBUTEVALUESTRING:
		ok = stage.IsStagedATTRIBUTEVALUESTRING(target)

	case *ATTRIBUTEVALUEXHTML:
		ok = stage.IsStagedATTRIBUTEVALUEXHTML(target)

	case *CHILDREN:
		ok = stage.IsStagedCHILDREN(target)

	case *CORECONTENT:
		ok = stage.IsStagedCORECONTENT(target)

	case *DATATYPEDEFINITIONBOOLEAN:
		ok = stage.IsStagedDATATYPEDEFINITIONBOOLEAN(target)

	case *DATATYPEDEFINITIONDATE:
		ok = stage.IsStagedDATATYPEDEFINITIONDATE(target)

	case *DATATYPEDEFINITIONENUMERATION:
		ok = stage.IsStagedDATATYPEDEFINITIONENUMERATION(target)

	case *DATATYPEDEFINITIONINTEGER:
		ok = stage.IsStagedDATATYPEDEFINITIONINTEGER(target)

	case *DATATYPEDEFINITIONREAL:
		ok = stage.IsStagedDATATYPEDEFINITIONREAL(target)

	case *DATATYPEDEFINITIONSTRING:
		ok = stage.IsStagedDATATYPEDEFINITIONSTRING(target)

	case *DATATYPEDEFINITIONXHTML:
		ok = stage.IsStagedDATATYPEDEFINITIONXHTML(target)

	case *DATATYPES:
		ok = stage.IsStagedDATATYPES(target)

	case *DEFAULTVALUE:
		ok = stage.IsStagedDEFAULTVALUE(target)

	case *DEFINITION:
		ok = stage.IsStagedDEFINITION(target)

	case *EDITABLEATTS:
		ok = stage.IsStagedEDITABLEATTS(target)

	case *EMBEDDEDVALUE:
		ok = stage.IsStagedEMBEDDEDVALUE(target)

	case *ENUMVALUE:
		ok = stage.IsStagedENUMVALUE(target)

	case *OBJECT:
		ok = stage.IsStagedOBJECT(target)

	case *PROPERTIES:
		ok = stage.IsStagedPROPERTIES(target)

	case *RELATIONGROUP:
		ok = stage.IsStagedRELATIONGROUP(target)

	case *RELATIONGROUPTYPE:
		ok = stage.IsStagedRELATIONGROUPTYPE(target)

	case *REQIF:
		ok = stage.IsStagedREQIF(target)

	case *REQIFCONTENT:
		ok = stage.IsStagedREQIFCONTENT(target)

	case *REQIFHEADER:
		ok = stage.IsStagedREQIFHEADER(target)

	case *REQIFTOOLEXTENSION:
		ok = stage.IsStagedREQIFTOOLEXTENSION(target)

	case *REQIFTYPE:
		ok = stage.IsStagedREQIFTYPE(target)

	case *SOURCE:
		ok = stage.IsStagedSOURCE(target)

	case *SOURCESPECIFICATION:
		ok = stage.IsStagedSOURCESPECIFICATION(target)

	case *SPECATTRIBUTES:
		ok = stage.IsStagedSPECATTRIBUTES(target)

	case *SPECHIERARCHY:
		ok = stage.IsStagedSPECHIERARCHY(target)

	case *SPECIFICATION:
		ok = stage.IsStagedSPECIFICATION(target)

	case *SPECIFICATIONS:
		ok = stage.IsStagedSPECIFICATIONS(target)

	case *SPECIFICATIONTYPE:
		ok = stage.IsStagedSPECIFICATIONTYPE(target)

	case *SPECIFIEDVALUES:
		ok = stage.IsStagedSPECIFIEDVALUES(target)

	case *SPECOBJECT:
		ok = stage.IsStagedSPECOBJECT(target)

	case *SPECOBJECTS:
		ok = stage.IsStagedSPECOBJECTS(target)

	case *SPECOBJECTTYPE:
		ok = stage.IsStagedSPECOBJECTTYPE(target)

	case *SPECRELATION:
		ok = stage.IsStagedSPECRELATION(target)

	case *SPECRELATIONGROUPS:
		ok = stage.IsStagedSPECRELATIONGROUPS(target)

	case *SPECRELATIONS:
		ok = stage.IsStagedSPECRELATIONS(target)

	case *SPECRELATIONTYPE:
		ok = stage.IsStagedSPECRELATIONTYPE(target)

	case *SPECTYPES:
		ok = stage.IsStagedSPECTYPES(target)

	case *TARGET:
		ok = stage.IsStagedTARGET(target)

	case *TARGETSPECIFICATION:
		ok = stage.IsStagedTARGETSPECIFICATION(target)

	case *THEHEADER:
		ok = stage.IsStagedTHEHEADER(target)

	case *TOOLEXTENSIONS:
		ok = stage.IsStagedTOOLEXTENSIONS(target)

	case *VALUES:
		ok = stage.IsStagedVALUES(target)

	case *XHTMLCONTENT:
		ok = stage.IsStagedXHTMLCONTENT(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedALTERNATIVEID(alternativeid *ALTERNATIVEID) (ok bool) {

	_, ok = stage.ALTERNATIVEIDs[alternativeid]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONBOOLEANs[attributedefinitionboolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONDATE(attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONDATEs[attributedefinitiondate]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONENUMERATIONs[attributedefinitionenumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONINTEGERs[attributedefinitioninteger]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONREAL(attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONREALs[attributedefinitionreal]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONSTRINGs[attributedefinitionstring]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) (ok bool) {

	_, ok = stage.ATTRIBUTEDEFINITIONXHTMLs[attributedefinitionxhtml]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEBOOLEAN(attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEBOOLEANs[attributevalueboolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEDATE(attributevaluedate *ATTRIBUTEVALUEDATE) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEDATEs[attributevaluedate]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEENUMERATION(attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEENUMERATIONs[attributevalueenumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEINTEGER(attributevalueinteger *ATTRIBUTEVALUEINTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEINTEGERs[attributevalueinteger]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEREAL(attributevaluereal *ATTRIBUTEVALUEREAL) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEREALs[attributevaluereal]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUESTRING(attributevaluestring *ATTRIBUTEVALUESTRING) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUESTRINGs[attributevaluestring]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTEVALUEXHTML(attributevaluexhtml *ATTRIBUTEVALUEXHTML) (ok bool) {

	_, ok = stage.ATTRIBUTEVALUEXHTMLs[attributevaluexhtml]

	return
}

func (stage *StageStruct) IsStagedCHILDREN(children *CHILDREN) (ok bool) {

	_, ok = stage.CHILDRENs[children]

	return
}

func (stage *StageStruct) IsStagedCORECONTENT(corecontent *CORECONTENT) (ok bool) {

	_, ok = stage.CORECONTENTs[corecontent]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONBOOLEANs[datatypedefinitionboolean]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONDATE(datatypedefinitiondate *DATATYPEDEFINITIONDATE) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONDATEs[datatypedefinitiondate]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONENUMERATIONs[datatypedefinitionenumeration]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONINTEGERs[datatypedefinitioninteger]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONREAL(datatypedefinitionreal *DATATYPEDEFINITIONREAL) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONREALs[datatypedefinitionreal]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONSTRING(datatypedefinitionstring *DATATYPEDEFINITIONSTRING) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONSTRINGs[datatypedefinitionstring]

	return
}

func (stage *StageStruct) IsStagedDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) (ok bool) {

	_, ok = stage.DATATYPEDEFINITIONXHTMLs[datatypedefinitionxhtml]

	return
}

func (stage *StageStruct) IsStagedDATATYPES(datatypes *DATATYPES) (ok bool) {

	_, ok = stage.DATATYPESs[datatypes]

	return
}

func (stage *StageStruct) IsStagedDEFAULTVALUE(defaultvalue *DEFAULTVALUE) (ok bool) {

	_, ok = stage.DEFAULTVALUEs[defaultvalue]

	return
}

func (stage *StageStruct) IsStagedDEFINITION(definition *DEFINITION) (ok bool) {

	_, ok = stage.DEFINITIONs[definition]

	return
}

func (stage *StageStruct) IsStagedEDITABLEATTS(editableatts *EDITABLEATTS) (ok bool) {

	_, ok = stage.EDITABLEATTSs[editableatts]

	return
}

func (stage *StageStruct) IsStagedEMBEDDEDVALUE(embeddedvalue *EMBEDDEDVALUE) (ok bool) {

	_, ok = stage.EMBEDDEDVALUEs[embeddedvalue]

	return
}

func (stage *StageStruct) IsStagedENUMVALUE(enumvalue *ENUMVALUE) (ok bool) {

	_, ok = stage.ENUMVALUEs[enumvalue]

	return
}

func (stage *StageStruct) IsStagedOBJECT(object *OBJECT) (ok bool) {

	_, ok = stage.OBJECTs[object]

	return
}

func (stage *StageStruct) IsStagedPROPERTIES(properties *PROPERTIES) (ok bool) {

	_, ok = stage.PROPERTIESs[properties]

	return
}

func (stage *StageStruct) IsStagedRELATIONGROUP(relationgroup *RELATIONGROUP) (ok bool) {

	_, ok = stage.RELATIONGROUPs[relationgroup]

	return
}

func (stage *StageStruct) IsStagedRELATIONGROUPTYPE(relationgrouptype *RELATIONGROUPTYPE) (ok bool) {

	_, ok = stage.RELATIONGROUPTYPEs[relationgrouptype]

	return
}

func (stage *StageStruct) IsStagedREQIF(reqif *REQIF) (ok bool) {

	_, ok = stage.REQIFs[reqif]

	return
}

func (stage *StageStruct) IsStagedREQIFCONTENT(reqifcontent *REQIFCONTENT) (ok bool) {

	_, ok = stage.REQIFCONTENTs[reqifcontent]

	return
}

func (stage *StageStruct) IsStagedREQIFHEADER(reqifheader *REQIFHEADER) (ok bool) {

	_, ok = stage.REQIFHEADERs[reqifheader]

	return
}

func (stage *StageStruct) IsStagedREQIFTOOLEXTENSION(reqiftoolextension *REQIFTOOLEXTENSION) (ok bool) {

	_, ok = stage.REQIFTOOLEXTENSIONs[reqiftoolextension]

	return
}

func (stage *StageStruct) IsStagedREQIFTYPE(reqiftype *REQIFTYPE) (ok bool) {

	_, ok = stage.REQIFTYPEs[reqiftype]

	return
}

func (stage *StageStruct) IsStagedSOURCE(source *SOURCE) (ok bool) {

	_, ok = stage.SOURCEs[source]

	return
}

func (stage *StageStruct) IsStagedSOURCESPECIFICATION(sourcespecification *SOURCESPECIFICATION) (ok bool) {

	_, ok = stage.SOURCESPECIFICATIONs[sourcespecification]

	return
}

func (stage *StageStruct) IsStagedSPECATTRIBUTES(specattributes *SPECATTRIBUTES) (ok bool) {

	_, ok = stage.SPECATTRIBUTESs[specattributes]

	return
}

func (stage *StageStruct) IsStagedSPECHIERARCHY(spechierarchy *SPECHIERARCHY) (ok bool) {

	_, ok = stage.SPECHIERARCHYs[spechierarchy]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATION(specification *SPECIFICATION) (ok bool) {

	_, ok = stage.SPECIFICATIONs[specification]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATIONS(specifications *SPECIFICATIONS) (ok bool) {

	_, ok = stage.SPECIFICATIONSs[specifications]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATIONTYPE(specificationtype *SPECIFICATIONTYPE) (ok bool) {

	_, ok = stage.SPECIFICATIONTYPEs[specificationtype]

	return
}

func (stage *StageStruct) IsStagedSPECIFIEDVALUES(specifiedvalues *SPECIFIEDVALUES) (ok bool) {

	_, ok = stage.SPECIFIEDVALUESs[specifiedvalues]

	return
}

func (stage *StageStruct) IsStagedSPECOBJECT(specobject *SPECOBJECT) (ok bool) {

	_, ok = stage.SPECOBJECTs[specobject]

	return
}

func (stage *StageStruct) IsStagedSPECOBJECTS(specobjects *SPECOBJECTS) (ok bool) {

	_, ok = stage.SPECOBJECTSs[specobjects]

	return
}

func (stage *StageStruct) IsStagedSPECOBJECTTYPE(specobjecttype *SPECOBJECTTYPE) (ok bool) {

	_, ok = stage.SPECOBJECTTYPEs[specobjecttype]

	return
}

func (stage *StageStruct) IsStagedSPECRELATION(specrelation *SPECRELATION) (ok bool) {

	_, ok = stage.SPECRELATIONs[specrelation]

	return
}

func (stage *StageStruct) IsStagedSPECRELATIONGROUPS(specrelationgroups *SPECRELATIONGROUPS) (ok bool) {

	_, ok = stage.SPECRELATIONGROUPSs[specrelationgroups]

	return
}

func (stage *StageStruct) IsStagedSPECRELATIONS(specrelations *SPECRELATIONS) (ok bool) {

	_, ok = stage.SPECRELATIONSs[specrelations]

	return
}

func (stage *StageStruct) IsStagedSPECRELATIONTYPE(specrelationtype *SPECRELATIONTYPE) (ok bool) {

	_, ok = stage.SPECRELATIONTYPEs[specrelationtype]

	return
}

func (stage *StageStruct) IsStagedSPECTYPES(spectypes *SPECTYPES) (ok bool) {

	_, ok = stage.SPECTYPESs[spectypes]

	return
}

func (stage *StageStruct) IsStagedTARGET(target *TARGET) (ok bool) {

	_, ok = stage.TARGETs[target]

	return
}

func (stage *StageStruct) IsStagedTARGETSPECIFICATION(targetspecification *TARGETSPECIFICATION) (ok bool) {

	_, ok = stage.TARGETSPECIFICATIONs[targetspecification]

	return
}

func (stage *StageStruct) IsStagedTHEHEADER(theheader *THEHEADER) (ok bool) {

	_, ok = stage.THEHEADERs[theheader]

	return
}

func (stage *StageStruct) IsStagedTOOLEXTENSIONS(toolextensions *TOOLEXTENSIONS) (ok bool) {

	_, ok = stage.TOOLEXTENSIONSs[toolextensions]

	return
}

func (stage *StageStruct) IsStagedVALUES(values *VALUES) (ok bool) {

	_, ok = stage.VALUESs[values]

	return
}

func (stage *StageStruct) IsStagedXHTMLCONTENT(xhtmlcontent *XHTMLCONTENT) (ok bool) {

	_, ok = stage.XHTMLCONTENTs[xhtmlcontent]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *ALTERNATIVEID:
		stage.StageBranchALTERNATIVEID(target)

	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.StageBranchATTRIBUTEDEFINITIONBOOLEAN(target)

	case *ATTRIBUTEDEFINITIONDATE:
		stage.StageBranchATTRIBUTEDEFINITIONDATE(target)

	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.StageBranchATTRIBUTEDEFINITIONENUMERATION(target)

	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.StageBranchATTRIBUTEDEFINITIONINTEGER(target)

	case *ATTRIBUTEDEFINITIONREAL:
		stage.StageBranchATTRIBUTEDEFINITIONREAL(target)

	case *ATTRIBUTEDEFINITIONSTRING:
		stage.StageBranchATTRIBUTEDEFINITIONSTRING(target)

	case *ATTRIBUTEDEFINITIONXHTML:
		stage.StageBranchATTRIBUTEDEFINITIONXHTML(target)

	case *ATTRIBUTEVALUEBOOLEAN:
		stage.StageBranchATTRIBUTEVALUEBOOLEAN(target)

	case *ATTRIBUTEVALUEDATE:
		stage.StageBranchATTRIBUTEVALUEDATE(target)

	case *ATTRIBUTEVALUEENUMERATION:
		stage.StageBranchATTRIBUTEVALUEENUMERATION(target)

	case *ATTRIBUTEVALUEINTEGER:
		stage.StageBranchATTRIBUTEVALUEINTEGER(target)

	case *ATTRIBUTEVALUEREAL:
		stage.StageBranchATTRIBUTEVALUEREAL(target)

	case *ATTRIBUTEVALUESTRING:
		stage.StageBranchATTRIBUTEVALUESTRING(target)

	case *ATTRIBUTEVALUEXHTML:
		stage.StageBranchATTRIBUTEVALUEXHTML(target)

	case *CHILDREN:
		stage.StageBranchCHILDREN(target)

	case *CORECONTENT:
		stage.StageBranchCORECONTENT(target)

	case *DATATYPEDEFINITIONBOOLEAN:
		stage.StageBranchDATATYPEDEFINITIONBOOLEAN(target)

	case *DATATYPEDEFINITIONDATE:
		stage.StageBranchDATATYPEDEFINITIONDATE(target)

	case *DATATYPEDEFINITIONENUMERATION:
		stage.StageBranchDATATYPEDEFINITIONENUMERATION(target)

	case *DATATYPEDEFINITIONINTEGER:
		stage.StageBranchDATATYPEDEFINITIONINTEGER(target)

	case *DATATYPEDEFINITIONREAL:
		stage.StageBranchDATATYPEDEFINITIONREAL(target)

	case *DATATYPEDEFINITIONSTRING:
		stage.StageBranchDATATYPEDEFINITIONSTRING(target)

	case *DATATYPEDEFINITIONXHTML:
		stage.StageBranchDATATYPEDEFINITIONXHTML(target)

	case *DATATYPES:
		stage.StageBranchDATATYPES(target)

	case *DEFAULTVALUE:
		stage.StageBranchDEFAULTVALUE(target)

	case *DEFINITION:
		stage.StageBranchDEFINITION(target)

	case *EDITABLEATTS:
		stage.StageBranchEDITABLEATTS(target)

	case *EMBEDDEDVALUE:
		stage.StageBranchEMBEDDEDVALUE(target)

	case *ENUMVALUE:
		stage.StageBranchENUMVALUE(target)

	case *OBJECT:
		stage.StageBranchOBJECT(target)

	case *PROPERTIES:
		stage.StageBranchPROPERTIES(target)

	case *RELATIONGROUP:
		stage.StageBranchRELATIONGROUP(target)

	case *RELATIONGROUPTYPE:
		stage.StageBranchRELATIONGROUPTYPE(target)

	case *REQIF:
		stage.StageBranchREQIF(target)

	case *REQIFCONTENT:
		stage.StageBranchREQIFCONTENT(target)

	case *REQIFHEADER:
		stage.StageBranchREQIFHEADER(target)

	case *REQIFTOOLEXTENSION:
		stage.StageBranchREQIFTOOLEXTENSION(target)

	case *REQIFTYPE:
		stage.StageBranchREQIFTYPE(target)

	case *SOURCE:
		stage.StageBranchSOURCE(target)

	case *SOURCESPECIFICATION:
		stage.StageBranchSOURCESPECIFICATION(target)

	case *SPECATTRIBUTES:
		stage.StageBranchSPECATTRIBUTES(target)

	case *SPECHIERARCHY:
		stage.StageBranchSPECHIERARCHY(target)

	case *SPECIFICATION:
		stage.StageBranchSPECIFICATION(target)

	case *SPECIFICATIONS:
		stage.StageBranchSPECIFICATIONS(target)

	case *SPECIFICATIONTYPE:
		stage.StageBranchSPECIFICATIONTYPE(target)

	case *SPECIFIEDVALUES:
		stage.StageBranchSPECIFIEDVALUES(target)

	case *SPECOBJECT:
		stage.StageBranchSPECOBJECT(target)

	case *SPECOBJECTS:
		stage.StageBranchSPECOBJECTS(target)

	case *SPECOBJECTTYPE:
		stage.StageBranchSPECOBJECTTYPE(target)

	case *SPECRELATION:
		stage.StageBranchSPECRELATION(target)

	case *SPECRELATIONGROUPS:
		stage.StageBranchSPECRELATIONGROUPS(target)

	case *SPECRELATIONS:
		stage.StageBranchSPECRELATIONS(target)

	case *SPECRELATIONTYPE:
		stage.StageBranchSPECRELATIONTYPE(target)

	case *SPECTYPES:
		stage.StageBranchSPECTYPES(target)

	case *TARGET:
		stage.StageBranchTARGET(target)

	case *TARGETSPECIFICATION:
		stage.StageBranchTARGETSPECIFICATION(target)

	case *THEHEADER:
		stage.StageBranchTHEHEADER(target)

	case *TOOLEXTENSIONS:
		stage.StageBranchTOOLEXTENSIONS(target)

	case *VALUES:
		stage.StageBranchVALUES(target)

	case *XHTMLCONTENT:
		stage.StageBranchXHTMLCONTENT(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchALTERNATIVEID(alternativeid *ALTERNATIVEID) {

	// check if instance is already staged
	if IsStaged(stage, alternativeid) {
		return
	}

	alternativeid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitionboolean) {
		return
	}

	attributedefinitionboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionboolean.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitionboolean.ALTERNATIVEID)
	}
	if attributedefinitionboolean.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitionboolean.DEFAULTVALUE)
	}
	if attributedefinitionboolean.TYPE != nil {
		StageBranch(stage, attributedefinitionboolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONDATE(attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitiondate) {
		return
	}

	attributedefinitiondate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitiondate.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitiondate.ALTERNATIVEID)
	}
	if attributedefinitiondate.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitiondate.DEFAULTVALUE)
	}
	if attributedefinitiondate.TYPE != nil {
		StageBranch(stage, attributedefinitiondate.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitionenumeration) {
		return
	}

	attributedefinitionenumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionenumeration.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitionenumeration.DEFAULTVALUE)
	}
	if attributedefinitionenumeration.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitionenumeration.ALTERNATIVEID)
	}
	if attributedefinitionenumeration.TYPE != nil {
		StageBranch(stage, attributedefinitionenumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitioninteger) {
		return
	}

	attributedefinitioninteger.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitioninteger.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitioninteger.ALTERNATIVEID)
	}
	if attributedefinitioninteger.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitioninteger.DEFAULTVALUE)
	}
	if attributedefinitioninteger.TYPE != nil {
		StageBranch(stage, attributedefinitioninteger.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONREAL(attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitionreal) {
		return
	}

	attributedefinitionreal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionreal.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitionreal.ALTERNATIVEID)
	}
	if attributedefinitionreal.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitionreal.DEFAULTVALUE)
	}
	if attributedefinitionreal.TYPE != nil {
		StageBranch(stage, attributedefinitionreal.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitionstring) {
		return
	}

	attributedefinitionstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionstring.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitionstring.ALTERNATIVEID)
	}
	if attributedefinitionstring.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitionstring.DEFAULTVALUE)
	}
	if attributedefinitionstring.TYPE != nil {
		StageBranch(stage, attributedefinitionstring.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) {

	// check if instance is already staged
	if IsStaged(stage, attributedefinitionxhtml) {
		return
	}

	attributedefinitionxhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionxhtml.ALTERNATIVEID != nil {
		StageBranch(stage, attributedefinitionxhtml.ALTERNATIVEID)
	}
	if attributedefinitionxhtml.DEFAULTVALUE != nil {
		StageBranch(stage, attributedefinitionxhtml.DEFAULTVALUE)
	}
	if attributedefinitionxhtml.TYPE != nil {
		StageBranch(stage, attributedefinitionxhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEBOOLEAN(attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attributevalueboolean) {
		return
	}

	attributevalueboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueboolean.DEFINITION != nil {
		StageBranch(stage, attributevalueboolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEDATE(attributevaluedate *ATTRIBUTEVALUEDATE) {

	// check if instance is already staged
	if IsStaged(stage, attributevaluedate) {
		return
	}

	attributevaluedate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluedate.DEFINITION != nil {
		StageBranch(stage, attributevaluedate.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEENUMERATION(attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attributevalueenumeration) {
		return
	}

	attributevalueenumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueenumeration.DEFINITION != nil {
		StageBranch(stage, attributevalueenumeration.DEFINITION)
	}
	if attributevalueenumeration.VALUES != nil {
		StageBranch(stage, attributevalueenumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEINTEGER(attributevalueinteger *ATTRIBUTEVALUEINTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attributevalueinteger) {
		return
	}

	attributevalueinteger.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueinteger.DEFINITION != nil {
		StageBranch(stage, attributevalueinteger.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEREAL(attributevaluereal *ATTRIBUTEVALUEREAL) {

	// check if instance is already staged
	if IsStaged(stage, attributevaluereal) {
		return
	}

	attributevaluereal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluereal.DEFINITION != nil {
		StageBranch(stage, attributevaluereal.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUESTRING(attributevaluestring *ATTRIBUTEVALUESTRING) {

	// check if instance is already staged
	if IsStaged(stage, attributevaluestring) {
		return
	}

	attributevaluestring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluestring.DEFINITION != nil {
		StageBranch(stage, attributevaluestring.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTEVALUEXHTML(attributevaluexhtml *ATTRIBUTEVALUEXHTML) {

	// check if instance is already staged
	if IsStaged(stage, attributevaluexhtml) {
		return
	}

	attributevaluexhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluexhtml.THEVALUE != nil {
		StageBranch(stage, attributevaluexhtml.THEVALUE)
	}
	if attributevaluexhtml.THEORIGINALVALUE != nil {
		StageBranch(stage, attributevaluexhtml.THEORIGINALVALUE)
	}
	if attributevaluexhtml.DEFINITION != nil {
		StageBranch(stage, attributevaluexhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCHILDREN(children *CHILDREN) {

	// check if instance is already staged
	if IsStaged(stage, children) {
		return
	}

	children.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spechierarchy := range children.SPECHIERARCHY {
		StageBranch(stage, _spechierarchy)
	}

}

func (stage *StageStruct) StageBranchCORECONTENT(corecontent *CORECONTENT) {

	// check if instance is already staged
	if IsStaged(stage, corecontent) {
		return
	}

	corecontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if corecontent.REQIFCONTENT != nil {
		StageBranch(stage, corecontent.REQIFCONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitionboolean) {
		return
	}

	datatypedefinitionboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionboolean.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitionboolean.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONDATE(datatypedefinitiondate *DATATYPEDEFINITIONDATE) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitiondate) {
		return
	}

	datatypedefinitiondate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitiondate.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitiondate.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitionenumeration) {
		return
	}

	datatypedefinitionenumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionenumeration.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitionenumeration.ALTERNATIVEID)
	}
	if datatypedefinitionenumeration.SPECIFIEDVALUES != nil {
		StageBranch(stage, datatypedefinitionenumeration.SPECIFIEDVALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitioninteger) {
		return
	}

	datatypedefinitioninteger.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitioninteger.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitioninteger.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONREAL(datatypedefinitionreal *DATATYPEDEFINITIONREAL) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitionreal) {
		return
	}

	datatypedefinitionreal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionreal.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitionreal.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONSTRING(datatypedefinitionstring *DATATYPEDEFINITIONSTRING) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitionstring) {
		return
	}

	datatypedefinitionstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionstring.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitionstring.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) {

	// check if instance is already staged
	if IsStaged(stage, datatypedefinitionxhtml) {
		return
	}

	datatypedefinitionxhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionxhtml.ALTERNATIVEID != nil {
		StageBranch(stage, datatypedefinitionxhtml.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPES(datatypes *DATATYPES) {

	// check if instance is already staged
	if IsStaged(stage, datatypes) {
		return
	}

	datatypes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatypedefinitionboolean := range datatypes.DATATYPEDEFINITIONBOOLEAN {
		StageBranch(stage, _datatypedefinitionboolean)
	}
	for _, _datatypedefinitiondate := range datatypes.DATATYPEDEFINITIONDATE {
		StageBranch(stage, _datatypedefinitiondate)
	}
	for _, _datatypedefinitionenumeration := range datatypes.DATATYPEDEFINITIONENUMERATION {
		StageBranch(stage, _datatypedefinitionenumeration)
	}
	for _, _datatypedefinitioninteger := range datatypes.DATATYPEDEFINITIONINTEGER {
		StageBranch(stage, _datatypedefinitioninteger)
	}
	for _, _datatypedefinitionreal := range datatypes.DATATYPEDEFINITIONREAL {
		StageBranch(stage, _datatypedefinitionreal)
	}
	for _, _datatypedefinitionstring := range datatypes.DATATYPEDEFINITIONSTRING {
		StageBranch(stage, _datatypedefinitionstring)
	}
	for _, _datatypedefinitionxhtml := range datatypes.DATATYPEDEFINITIONXHTML {
		StageBranch(stage, _datatypedefinitionxhtml)
	}

}

func (stage *StageStruct) StageBranchDEFAULTVALUE(defaultvalue *DEFAULTVALUE) {

	// check if instance is already staged
	if IsStaged(stage, defaultvalue) {
		return
	}

	defaultvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaultvalue.ATTRIBUTEVALUEBOOLEAN != nil {
		StageBranch(stage, defaultvalue.ATTRIBUTEVALUEBOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDEFINITION(definition *DEFINITION) {

	// check if instance is already staged
	if IsStaged(stage, definition) {
		return
	}

	definition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEDITABLEATTS(editableatts *EDITABLEATTS) {

	// check if instance is already staged
	if IsStaged(stage, editableatts) {
		return
	}

	editableatts.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEMBEDDEDVALUE(embeddedvalue *EMBEDDEDVALUE) {

	// check if instance is already staged
	if IsStaged(stage, embeddedvalue) {
		return
	}

	embeddedvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchENUMVALUE(enumvalue *ENUMVALUE) {

	// check if instance is already staged
	if IsStaged(stage, enumvalue) {
		return
	}

	enumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumvalue.ALTERNATIVEID != nil {
		StageBranch(stage, enumvalue.ALTERNATIVEID)
	}
	if enumvalue.PROPERTIES != nil {
		StageBranch(stage, enumvalue.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOBJECT(object *OBJECT) {

	// check if instance is already staged
	if IsStaged(stage, object) {
		return
	}

	object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPROPERTIES(properties *PROPERTIES) {

	// check if instance is already staged
	if IsStaged(stage, properties) {
		return
	}

	properties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if properties.EMBEDDEDVALUE != nil {
		StageBranch(stage, properties.EMBEDDEDVALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRELATIONGROUP(relationgroup *RELATIONGROUP) {

	// check if instance is already staged
	if IsStaged(stage, relationgroup) {
		return
	}

	relationgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relationgroup.ALTERNATIVEID != nil {
		StageBranch(stage, relationgroup.ALTERNATIVEID)
	}
	if relationgroup.SOURCESPECIFICATION != nil {
		StageBranch(stage, relationgroup.SOURCESPECIFICATION)
	}
	if relationgroup.SPECRELATIONS != nil {
		StageBranch(stage, relationgroup.SPECRELATIONS)
	}
	if relationgroup.TARGETSPECIFICATION != nil {
		StageBranch(stage, relationgroup.TARGETSPECIFICATION)
	}
	if relationgroup.TYPE != nil {
		StageBranch(stage, relationgroup.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRELATIONGROUPTYPE(relationgrouptype *RELATIONGROUPTYPE) {

	// check if instance is already staged
	if IsStaged(stage, relationgrouptype) {
		return
	}

	relationgrouptype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relationgrouptype.ALTERNATIVEID != nil {
		StageBranch(stage, relationgrouptype.ALTERNATIVEID)
	}
	if relationgrouptype.SPECATTRIBUTES != nil {
		StageBranch(stage, relationgrouptype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if IsStaged(stage, reqif) {
		return
	}

	reqif.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.HEADER != nil {
		StageBranch(stage, reqif.HEADER)
	}
	if reqif.CORECONTENT != nil {
		StageBranch(stage, reqif.CORECONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIFCONTENT(reqifcontent *REQIFCONTENT) {

	// check if instance is already staged
	if IsStaged(stage, reqifcontent) {
		return
	}

	reqifcontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqifcontent.DATATYPES != nil {
		StageBranch(stage, reqifcontent.DATATYPES)
	}
	if reqifcontent.SPECTYPES != nil {
		StageBranch(stage, reqifcontent.SPECTYPES)
	}
	if reqifcontent.SPECOBJECTS != nil {
		StageBranch(stage, reqifcontent.SPECOBJECTS)
	}
	if reqifcontent.SPECRELATIONS != nil {
		StageBranch(stage, reqifcontent.SPECRELATIONS)
	}
	if reqifcontent.SPECIFICATIONS != nil {
		StageBranch(stage, reqifcontent.SPECIFICATIONS)
	}
	if reqifcontent.SPECRELATIONGROUPS != nil {
		StageBranch(stage, reqifcontent.SPECRELATIONGROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIFHEADER(reqifheader *REQIFHEADER) {

	// check if instance is already staged
	if IsStaged(stage, reqifheader) {
		return
	}

	reqifheader.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIFTOOLEXTENSION(reqiftoolextension *REQIFTOOLEXTENSION) {

	// check if instance is already staged
	if IsStaged(stage, reqiftoolextension) {
		return
	}

	reqiftoolextension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIFTYPE(reqiftype *REQIFTYPE) {

	// check if instance is already staged
	if IsStaged(stage, reqiftype) {
		return
	}

	reqiftype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSOURCE(source *SOURCE) {

	// check if instance is already staged
	if IsStaged(stage, source) {
		return
	}

	source.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSOURCESPECIFICATION(sourcespecification *SOURCESPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, sourcespecification) {
		return
	}

	sourcespecification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECATTRIBUTES(specattributes *SPECATTRIBUTES) {

	// check if instance is already staged
	if IsStaged(stage, specattributes) {
		return
	}

	specattributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributedefinitionboolean := range specattributes.ATTRIBUTEDEFINITIONBOOLEAN {
		StageBranch(stage, _attributedefinitionboolean)
	}
	for _, _attributedefinitiondate := range specattributes.ATTRIBUTEDEFINITIONDATE {
		StageBranch(stage, _attributedefinitiondate)
	}
	for _, _attributedefinitionenumeration := range specattributes.ATTRIBUTEDEFINITIONENUMERATION {
		StageBranch(stage, _attributedefinitionenumeration)
	}
	for _, _attributedefinitioninteger := range specattributes.ATTRIBUTEDEFINITIONINTEGER {
		StageBranch(stage, _attributedefinitioninteger)
	}
	for _, _attributedefinitionreal := range specattributes.ATTRIBUTEDEFINITIONREAL {
		StageBranch(stage, _attributedefinitionreal)
	}
	for _, _attributedefinitionstring := range specattributes.ATTRIBUTEDEFINITIONSTRING {
		StageBranch(stage, _attributedefinitionstring)
	}
	for _, _attributedefinitionxhtml := range specattributes.ATTRIBUTEDEFINITIONXHTML {
		StageBranch(stage, _attributedefinitionxhtml)
	}

}

func (stage *StageStruct) StageBranchSPECHIERARCHY(spechierarchy *SPECHIERARCHY) {

	// check if instance is already staged
	if IsStaged(stage, spechierarchy) {
		return
	}

	spechierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spechierarchy.ALTERNATIVEID != nil {
		StageBranch(stage, spechierarchy.ALTERNATIVEID)
	}
	if spechierarchy.CHILDREN != nil {
		StageBranch(stage, spechierarchy.CHILDREN)
	}
	if spechierarchy.EDITABLEATTS != nil {
		StageBranch(stage, spechierarchy.EDITABLEATTS)
	}
	if spechierarchy.OBJECT != nil {
		StageBranch(stage, spechierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, specification) {
		return
	}

	specification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVEID != nil {
		StageBranch(stage, specification.ALTERNATIVEID)
	}
	if specification.VALUES != nil {
		StageBranch(stage, specification.VALUES)
	}
	if specification.CHILDREN != nil {
		StageBranch(stage, specification.CHILDREN)
	}
	if specification.TYPE != nil {
		StageBranch(stage, specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFICATIONS(specifications *SPECIFICATIONS) {

	// check if instance is already staged
	if IsStaged(stage, specifications) {
		return
	}

	specifications.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range specifications.SPECIFICATION {
		StageBranch(stage, _specification)
	}

}

func (stage *StageStruct) StageBranchSPECIFICATIONTYPE(specificationtype *SPECIFICATIONTYPE) {

	// check if instance is already staged
	if IsStaged(stage, specificationtype) {
		return
	}

	specificationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specificationtype.ALTERNATIVEID != nil {
		StageBranch(stage, specificationtype.ALTERNATIVEID)
	}
	if specificationtype.SPECATTRIBUTES != nil {
		StageBranch(stage, specificationtype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFIEDVALUES(specifiedvalues *SPECIFIEDVALUES) {

	// check if instance is already staged
	if IsStaged(stage, specifiedvalues) {
		return
	}

	specifiedvalues.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumvalue := range specifiedvalues.ENUMVALUE {
		StageBranch(stage, _enumvalue)
	}

}

func (stage *StageStruct) StageBranchSPECOBJECT(specobject *SPECOBJECT) {

	// check if instance is already staged
	if IsStaged(stage, specobject) {
		return
	}

	specobject.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specobject.ALTERNATIVEID != nil {
		StageBranch(stage, specobject.ALTERNATIVEID)
	}
	if specobject.VALUES != nil {
		StageBranch(stage, specobject.VALUES)
	}
	if specobject.TYPE != nil {
		StageBranch(stage, specobject.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECOBJECTS(specobjects *SPECOBJECTS) {

	// check if instance is already staged
	if IsStaged(stage, specobjects) {
		return
	}

	specobjects.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specobject := range specobjects.SPECOBJECT {
		StageBranch(stage, _specobject)
	}

}

func (stage *StageStruct) StageBranchSPECOBJECTTYPE(specobjecttype *SPECOBJECTTYPE) {

	// check if instance is already staged
	if IsStaged(stage, specobjecttype) {
		return
	}

	specobjecttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specobjecttype.ALTERNATIVEID != nil {
		StageBranch(stage, specobjecttype.ALTERNATIVEID)
	}
	if specobjecttype.SPECATTRIBUTES != nil {
		StageBranch(stage, specobjecttype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECRELATION(specrelation *SPECRELATION) {

	// check if instance is already staged
	if IsStaged(stage, specrelation) {
		return
	}

	specrelation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specrelation.ALTERNATIVEID != nil {
		StageBranch(stage, specrelation.ALTERNATIVEID)
	}
	if specrelation.VALUES != nil {
		StageBranch(stage, specrelation.VALUES)
	}
	if specrelation.SOURCE != nil {
		StageBranch(stage, specrelation.SOURCE)
	}
	if specrelation.TARGET != nil {
		StageBranch(stage, specrelation.TARGET)
	}
	if specrelation.TYPE != nil {
		StageBranch(stage, specrelation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECRELATIONGROUPS(specrelationgroups *SPECRELATIONGROUPS) {

	// check if instance is already staged
	if IsStaged(stage, specrelationgroups) {
		return
	}

	specrelationgroups.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgroup := range specrelationgroups.RELATIONGROUP {
		StageBranch(stage, _relationgroup)
	}

}

func (stage *StageStruct) StageBranchSPECRELATIONS(specrelations *SPECRELATIONS) {

	// check if instance is already staged
	if IsStaged(stage, specrelations) {
		return
	}

	specrelations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECRELATIONTYPE(specrelationtype *SPECRELATIONTYPE) {

	// check if instance is already staged
	if IsStaged(stage, specrelationtype) {
		return
	}

	specrelationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specrelationtype.ALTERNATIVEID != nil {
		StageBranch(stage, specrelationtype.ALTERNATIVEID)
	}
	if specrelationtype.SPECATTRIBUTES != nil {
		StageBranch(stage, specrelationtype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECTYPES(spectypes *SPECTYPES) {

	// check if instance is already staged
	if IsStaged(stage, spectypes) {
		return
	}

	spectypes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgrouptype := range spectypes.RELATIONGROUPTYPE {
		StageBranch(stage, _relationgrouptype)
	}
	for _, _specobjecttype := range spectypes.SPECOBJECTTYPE {
		StageBranch(stage, _specobjecttype)
	}
	for _, _specrelationtype := range spectypes.SPECRELATIONTYPE {
		StageBranch(stage, _specrelationtype)
	}
	for _, _specificationtype := range spectypes.SPECIFICATIONTYPE {
		StageBranch(stage, _specificationtype)
	}

}

func (stage *StageStruct) StageBranchTARGET(target *TARGET) {

	// check if instance is already staged
	if IsStaged(stage, target) {
		return
	}

	target.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTARGETSPECIFICATION(targetspecification *TARGETSPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, targetspecification) {
		return
	}

	targetspecification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTHEHEADER(theheader *THEHEADER) {

	// check if instance is already staged
	if IsStaged(stage, theheader) {
		return
	}

	theheader.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if theheader.REQIFHEADER != nil {
		StageBranch(stage, theheader.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTOOLEXTENSIONS(toolextensions *TOOLEXTENSIONS) {

	// check if instance is already staged
	if IsStaged(stage, toolextensions) {
		return
	}

	toolextensions.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _reqiftoolextension := range toolextensions.REQIFTOOLEXTENSION {
		StageBranch(stage, _reqiftoolextension)
	}

}

func (stage *StageStruct) StageBranchVALUES(values *VALUES) {

	// check if instance is already staged
	if IsStaged(stage, values) {
		return
	}

	values.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXHTMLCONTENT(xhtmlcontent *XHTMLCONTENT) {

	// check if instance is already staged
	if IsStaged(stage, xhtmlcontent) {
		return
	}

	xhtmlcontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ALTERNATIVEID:
		toT := CopyBranchALTERNATIVEID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONBOOLEAN:
		toT := CopyBranchATTRIBUTEDEFINITIONBOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONDATE:
		toT := CopyBranchATTRIBUTEDEFINITIONDATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONENUMERATION:
		toT := CopyBranchATTRIBUTEDEFINITIONENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONINTEGER:
		toT := CopyBranchATTRIBUTEDEFINITIONINTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONREAL:
		toT := CopyBranchATTRIBUTEDEFINITIONREAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONSTRING:
		toT := CopyBranchATTRIBUTEDEFINITIONSTRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEDEFINITIONXHTML:
		toT := CopyBranchATTRIBUTEDEFINITIONXHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEBOOLEAN:
		toT := CopyBranchATTRIBUTEVALUEBOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEDATE:
		toT := CopyBranchATTRIBUTEVALUEDATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEENUMERATION:
		toT := CopyBranchATTRIBUTEVALUEENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEINTEGER:
		toT := CopyBranchATTRIBUTEVALUEINTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEREAL:
		toT := CopyBranchATTRIBUTEVALUEREAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUESTRING:
		toT := CopyBranchATTRIBUTEVALUESTRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTEVALUEXHTML:
		toT := CopyBranchATTRIBUTEVALUEXHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CHILDREN:
		toT := CopyBranchCHILDREN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CORECONTENT:
		toT := CopyBranchCORECONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONBOOLEAN:
		toT := CopyBranchDATATYPEDEFINITIONBOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONDATE:
		toT := CopyBranchDATATYPEDEFINITIONDATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONENUMERATION:
		toT := CopyBranchDATATYPEDEFINITIONENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONINTEGER:
		toT := CopyBranchDATATYPEDEFINITIONINTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONREAL:
		toT := CopyBranchDATATYPEDEFINITIONREAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONSTRING:
		toT := CopyBranchDATATYPEDEFINITIONSTRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPEDEFINITIONXHTML:
		toT := CopyBranchDATATYPEDEFINITIONXHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPES:
		toT := CopyBranchDATATYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DEFAULTVALUE:
		toT := CopyBranchDEFAULTVALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DEFINITION:
		toT := CopyBranchDEFINITION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EDITABLEATTS:
		toT := CopyBranchEDITABLEATTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EMBEDDEDVALUE:
		toT := CopyBranchEMBEDDEDVALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ENUMVALUE:
		toT := CopyBranchENUMVALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OBJECT:
		toT := CopyBranchOBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PROPERTIES:
		toT := CopyBranchPROPERTIES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATIONGROUP:
		toT := CopyBranchRELATIONGROUP(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATIONGROUPTYPE:
		toT := CopyBranchRELATIONGROUPTYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIF:
		toT := CopyBranchREQIF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIFCONTENT:
		toT := CopyBranchREQIFCONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIFHEADER:
		toT := CopyBranchREQIFHEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIFTOOLEXTENSION:
		toT := CopyBranchREQIFTOOLEXTENSION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIFTYPE:
		toT := CopyBranchREQIFTYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SOURCE:
		toT := CopyBranchSOURCE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SOURCESPECIFICATION:
		toT := CopyBranchSOURCESPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECATTRIBUTES:
		toT := CopyBranchSPECATTRIBUTES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECHIERARCHY:
		toT := CopyBranchSPECHIERARCHY(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION:
		toT := CopyBranchSPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATIONS:
		toT := CopyBranchSPECIFICATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATIONTYPE:
		toT := CopyBranchSPECIFICATIONTYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFIEDVALUES:
		toT := CopyBranchSPECIFIEDVALUES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECOBJECT:
		toT := CopyBranchSPECOBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECOBJECTS:
		toT := CopyBranchSPECOBJECTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECOBJECTTYPE:
		toT := CopyBranchSPECOBJECTTYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECRELATION:
		toT := CopyBranchSPECRELATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECRELATIONGROUPS:
		toT := CopyBranchSPECRELATIONGROUPS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECRELATIONS:
		toT := CopyBranchSPECRELATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECRELATIONTYPE:
		toT := CopyBranchSPECRELATIONTYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECTYPES:
		toT := CopyBranchSPECTYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TARGET:
		toT := CopyBranchTARGET(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TARGETSPECIFICATION:
		toT := CopyBranchTARGETSPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *THEHEADER:
		toT := CopyBranchTHEHEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TOOLEXTENSIONS:
		toT := CopyBranchTOOLEXTENSIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VALUES:
		toT := CopyBranchVALUES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XHTMLCONTENT:
		toT := CopyBranchXHTMLCONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchALTERNATIVEID(mapOrigCopy map[any]any, alternativeidFrom *ALTERNATIVEID) (alternativeidTo *ALTERNATIVEID) {

	// alternativeidFrom has already been copied
	if _alternativeidTo, ok := mapOrigCopy[alternativeidFrom]; ok {
		alternativeidTo = _alternativeidTo.(*ALTERNATIVEID)
		return
	}

	alternativeidTo = new(ALTERNATIVEID)
	mapOrigCopy[alternativeidFrom] = alternativeidTo
	alternativeidFrom.CopyBasicFields(alternativeidTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONBOOLEAN(mapOrigCopy map[any]any, attributedefinitionbooleanFrom *ATTRIBUTEDEFINITIONBOOLEAN) (attributedefinitionbooleanTo *ATTRIBUTEDEFINITIONBOOLEAN) {

	// attributedefinitionbooleanFrom has already been copied
	if _attributedefinitionbooleanTo, ok := mapOrigCopy[attributedefinitionbooleanFrom]; ok {
		attributedefinitionbooleanTo = _attributedefinitionbooleanTo.(*ATTRIBUTEDEFINITIONBOOLEAN)
		return
	}

	attributedefinitionbooleanTo = new(ATTRIBUTEDEFINITIONBOOLEAN)
	mapOrigCopy[attributedefinitionbooleanFrom] = attributedefinitionbooleanTo
	attributedefinitionbooleanFrom.CopyBasicFields(attributedefinitionbooleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionbooleanFrom.ALTERNATIVEID != nil {
		attributedefinitionbooleanTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionbooleanFrom.ALTERNATIVEID)
	}
	if attributedefinitionbooleanFrom.DEFAULTVALUE != nil {
		attributedefinitionbooleanTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionbooleanFrom.DEFAULTVALUE)
	}
	if attributedefinitionbooleanFrom.TYPE != nil {
		attributedefinitionbooleanTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionbooleanFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONDATE(mapOrigCopy map[any]any, attributedefinitiondateFrom *ATTRIBUTEDEFINITIONDATE) (attributedefinitiondateTo *ATTRIBUTEDEFINITIONDATE) {

	// attributedefinitiondateFrom has already been copied
	if _attributedefinitiondateTo, ok := mapOrigCopy[attributedefinitiondateFrom]; ok {
		attributedefinitiondateTo = _attributedefinitiondateTo.(*ATTRIBUTEDEFINITIONDATE)
		return
	}

	attributedefinitiondateTo = new(ATTRIBUTEDEFINITIONDATE)
	mapOrigCopy[attributedefinitiondateFrom] = attributedefinitiondateTo
	attributedefinitiondateFrom.CopyBasicFields(attributedefinitiondateTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitiondateFrom.ALTERNATIVEID != nil {
		attributedefinitiondateTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitiondateFrom.ALTERNATIVEID)
	}
	if attributedefinitiondateFrom.DEFAULTVALUE != nil {
		attributedefinitiondateTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitiondateFrom.DEFAULTVALUE)
	}
	if attributedefinitiondateFrom.TYPE != nil {
		attributedefinitiondateTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitiondateFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONENUMERATION(mapOrigCopy map[any]any, attributedefinitionenumerationFrom *ATTRIBUTEDEFINITIONENUMERATION) (attributedefinitionenumerationTo *ATTRIBUTEDEFINITIONENUMERATION) {

	// attributedefinitionenumerationFrom has already been copied
	if _attributedefinitionenumerationTo, ok := mapOrigCopy[attributedefinitionenumerationFrom]; ok {
		attributedefinitionenumerationTo = _attributedefinitionenumerationTo.(*ATTRIBUTEDEFINITIONENUMERATION)
		return
	}

	attributedefinitionenumerationTo = new(ATTRIBUTEDEFINITIONENUMERATION)
	mapOrigCopy[attributedefinitionenumerationFrom] = attributedefinitionenumerationTo
	attributedefinitionenumerationFrom.CopyBasicFields(attributedefinitionenumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionenumerationFrom.DEFAULTVALUE != nil {
		attributedefinitionenumerationTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionenumerationFrom.DEFAULTVALUE)
	}
	if attributedefinitionenumerationFrom.ALTERNATIVEID != nil {
		attributedefinitionenumerationTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionenumerationFrom.ALTERNATIVEID)
	}
	if attributedefinitionenumerationFrom.TYPE != nil {
		attributedefinitionenumerationTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionenumerationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONINTEGER(mapOrigCopy map[any]any, attributedefinitionintegerFrom *ATTRIBUTEDEFINITIONINTEGER) (attributedefinitionintegerTo *ATTRIBUTEDEFINITIONINTEGER) {

	// attributedefinitionintegerFrom has already been copied
	if _attributedefinitionintegerTo, ok := mapOrigCopy[attributedefinitionintegerFrom]; ok {
		attributedefinitionintegerTo = _attributedefinitionintegerTo.(*ATTRIBUTEDEFINITIONINTEGER)
		return
	}

	attributedefinitionintegerTo = new(ATTRIBUTEDEFINITIONINTEGER)
	mapOrigCopy[attributedefinitionintegerFrom] = attributedefinitionintegerTo
	attributedefinitionintegerFrom.CopyBasicFields(attributedefinitionintegerTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionintegerFrom.ALTERNATIVEID != nil {
		attributedefinitionintegerTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionintegerFrom.ALTERNATIVEID)
	}
	if attributedefinitionintegerFrom.DEFAULTVALUE != nil {
		attributedefinitionintegerTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionintegerFrom.DEFAULTVALUE)
	}
	if attributedefinitionintegerFrom.TYPE != nil {
		attributedefinitionintegerTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionintegerFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONREAL(mapOrigCopy map[any]any, attributedefinitionrealFrom *ATTRIBUTEDEFINITIONREAL) (attributedefinitionrealTo *ATTRIBUTEDEFINITIONREAL) {

	// attributedefinitionrealFrom has already been copied
	if _attributedefinitionrealTo, ok := mapOrigCopy[attributedefinitionrealFrom]; ok {
		attributedefinitionrealTo = _attributedefinitionrealTo.(*ATTRIBUTEDEFINITIONREAL)
		return
	}

	attributedefinitionrealTo = new(ATTRIBUTEDEFINITIONREAL)
	mapOrigCopy[attributedefinitionrealFrom] = attributedefinitionrealTo
	attributedefinitionrealFrom.CopyBasicFields(attributedefinitionrealTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionrealFrom.ALTERNATIVEID != nil {
		attributedefinitionrealTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionrealFrom.ALTERNATIVEID)
	}
	if attributedefinitionrealFrom.DEFAULTVALUE != nil {
		attributedefinitionrealTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionrealFrom.DEFAULTVALUE)
	}
	if attributedefinitionrealFrom.TYPE != nil {
		attributedefinitionrealTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionrealFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONSTRING(mapOrigCopy map[any]any, attributedefinitionstringFrom *ATTRIBUTEDEFINITIONSTRING) (attributedefinitionstringTo *ATTRIBUTEDEFINITIONSTRING) {

	// attributedefinitionstringFrom has already been copied
	if _attributedefinitionstringTo, ok := mapOrigCopy[attributedefinitionstringFrom]; ok {
		attributedefinitionstringTo = _attributedefinitionstringTo.(*ATTRIBUTEDEFINITIONSTRING)
		return
	}

	attributedefinitionstringTo = new(ATTRIBUTEDEFINITIONSTRING)
	mapOrigCopy[attributedefinitionstringFrom] = attributedefinitionstringTo
	attributedefinitionstringFrom.CopyBasicFields(attributedefinitionstringTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionstringFrom.ALTERNATIVEID != nil {
		attributedefinitionstringTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionstringFrom.ALTERNATIVEID)
	}
	if attributedefinitionstringFrom.DEFAULTVALUE != nil {
		attributedefinitionstringTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionstringFrom.DEFAULTVALUE)
	}
	if attributedefinitionstringFrom.TYPE != nil {
		attributedefinitionstringTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionstringFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEDEFINITIONXHTML(mapOrigCopy map[any]any, attributedefinitionxhtmlFrom *ATTRIBUTEDEFINITIONXHTML) (attributedefinitionxhtmlTo *ATTRIBUTEDEFINITIONXHTML) {

	// attributedefinitionxhtmlFrom has already been copied
	if _attributedefinitionxhtmlTo, ok := mapOrigCopy[attributedefinitionxhtmlFrom]; ok {
		attributedefinitionxhtmlTo = _attributedefinitionxhtmlTo.(*ATTRIBUTEDEFINITIONXHTML)
		return
	}

	attributedefinitionxhtmlTo = new(ATTRIBUTEDEFINITIONXHTML)
	mapOrigCopy[attributedefinitionxhtmlFrom] = attributedefinitionxhtmlTo
	attributedefinitionxhtmlFrom.CopyBasicFields(attributedefinitionxhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionxhtmlFrom.ALTERNATIVEID != nil {
		attributedefinitionxhtmlTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, attributedefinitionxhtmlFrom.ALTERNATIVEID)
	}
	if attributedefinitionxhtmlFrom.DEFAULTVALUE != nil {
		attributedefinitionxhtmlTo.DEFAULTVALUE = CopyBranchDEFAULTVALUE(mapOrigCopy, attributedefinitionxhtmlFrom.DEFAULTVALUE)
	}
	if attributedefinitionxhtmlFrom.TYPE != nil {
		attributedefinitionxhtmlTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, attributedefinitionxhtmlFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEBOOLEAN(mapOrigCopy map[any]any, attributevaluebooleanFrom *ATTRIBUTEVALUEBOOLEAN) (attributevaluebooleanTo *ATTRIBUTEVALUEBOOLEAN) {

	// attributevaluebooleanFrom has already been copied
	if _attributevaluebooleanTo, ok := mapOrigCopy[attributevaluebooleanFrom]; ok {
		attributevaluebooleanTo = _attributevaluebooleanTo.(*ATTRIBUTEVALUEBOOLEAN)
		return
	}

	attributevaluebooleanTo = new(ATTRIBUTEVALUEBOOLEAN)
	mapOrigCopy[attributevaluebooleanFrom] = attributevaluebooleanTo
	attributevaluebooleanFrom.CopyBasicFields(attributevaluebooleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluebooleanFrom.DEFINITION != nil {
		attributevaluebooleanTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevaluebooleanFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEDATE(mapOrigCopy map[any]any, attributevaluedateFrom *ATTRIBUTEVALUEDATE) (attributevaluedateTo *ATTRIBUTEVALUEDATE) {

	// attributevaluedateFrom has already been copied
	if _attributevaluedateTo, ok := mapOrigCopy[attributevaluedateFrom]; ok {
		attributevaluedateTo = _attributevaluedateTo.(*ATTRIBUTEVALUEDATE)
		return
	}

	attributevaluedateTo = new(ATTRIBUTEVALUEDATE)
	mapOrigCopy[attributevaluedateFrom] = attributevaluedateTo
	attributevaluedateFrom.CopyBasicFields(attributevaluedateTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluedateFrom.DEFINITION != nil {
		attributevaluedateTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevaluedateFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEENUMERATION(mapOrigCopy map[any]any, attributevalueenumerationFrom *ATTRIBUTEVALUEENUMERATION) (attributevalueenumerationTo *ATTRIBUTEVALUEENUMERATION) {

	// attributevalueenumerationFrom has already been copied
	if _attributevalueenumerationTo, ok := mapOrigCopy[attributevalueenumerationFrom]; ok {
		attributevalueenumerationTo = _attributevalueenumerationTo.(*ATTRIBUTEVALUEENUMERATION)
		return
	}

	attributevalueenumerationTo = new(ATTRIBUTEVALUEENUMERATION)
	mapOrigCopy[attributevalueenumerationFrom] = attributevalueenumerationTo
	attributevalueenumerationFrom.CopyBasicFields(attributevalueenumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueenumerationFrom.DEFINITION != nil {
		attributevalueenumerationTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevalueenumerationFrom.DEFINITION)
	}
	if attributevalueenumerationFrom.VALUES != nil {
		attributevalueenumerationTo.VALUES = CopyBranchVALUES(mapOrigCopy, attributevalueenumerationFrom.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEINTEGER(mapOrigCopy map[any]any, attributevalueintegerFrom *ATTRIBUTEVALUEINTEGER) (attributevalueintegerTo *ATTRIBUTEVALUEINTEGER) {

	// attributevalueintegerFrom has already been copied
	if _attributevalueintegerTo, ok := mapOrigCopy[attributevalueintegerFrom]; ok {
		attributevalueintegerTo = _attributevalueintegerTo.(*ATTRIBUTEVALUEINTEGER)
		return
	}

	attributevalueintegerTo = new(ATTRIBUTEVALUEINTEGER)
	mapOrigCopy[attributevalueintegerFrom] = attributevalueintegerTo
	attributevalueintegerFrom.CopyBasicFields(attributevalueintegerTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueintegerFrom.DEFINITION != nil {
		attributevalueintegerTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevalueintegerFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEREAL(mapOrigCopy map[any]any, attributevaluerealFrom *ATTRIBUTEVALUEREAL) (attributevaluerealTo *ATTRIBUTEVALUEREAL) {

	// attributevaluerealFrom has already been copied
	if _attributevaluerealTo, ok := mapOrigCopy[attributevaluerealFrom]; ok {
		attributevaluerealTo = _attributevaluerealTo.(*ATTRIBUTEVALUEREAL)
		return
	}

	attributevaluerealTo = new(ATTRIBUTEVALUEREAL)
	mapOrigCopy[attributevaluerealFrom] = attributevaluerealTo
	attributevaluerealFrom.CopyBasicFields(attributevaluerealTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluerealFrom.DEFINITION != nil {
		attributevaluerealTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevaluerealFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUESTRING(mapOrigCopy map[any]any, attributevaluestringFrom *ATTRIBUTEVALUESTRING) (attributevaluestringTo *ATTRIBUTEVALUESTRING) {

	// attributevaluestringFrom has already been copied
	if _attributevaluestringTo, ok := mapOrigCopy[attributevaluestringFrom]; ok {
		attributevaluestringTo = _attributevaluestringTo.(*ATTRIBUTEVALUESTRING)
		return
	}

	attributevaluestringTo = new(ATTRIBUTEVALUESTRING)
	mapOrigCopy[attributevaluestringFrom] = attributevaluestringTo
	attributevaluestringFrom.CopyBasicFields(attributevaluestringTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluestringFrom.DEFINITION != nil {
		attributevaluestringTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevaluestringFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTEVALUEXHTML(mapOrigCopy map[any]any, attributevaluexhtmlFrom *ATTRIBUTEVALUEXHTML) (attributevaluexhtmlTo *ATTRIBUTEVALUEXHTML) {

	// attributevaluexhtmlFrom has already been copied
	if _attributevaluexhtmlTo, ok := mapOrigCopy[attributevaluexhtmlFrom]; ok {
		attributevaluexhtmlTo = _attributevaluexhtmlTo.(*ATTRIBUTEVALUEXHTML)
		return
	}

	attributevaluexhtmlTo = new(ATTRIBUTEVALUEXHTML)
	mapOrigCopy[attributevaluexhtmlFrom] = attributevaluexhtmlTo
	attributevaluexhtmlFrom.CopyBasicFields(attributevaluexhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluexhtmlFrom.THEVALUE != nil {
		attributevaluexhtmlTo.THEVALUE = CopyBranchXHTMLCONTENT(mapOrigCopy, attributevaluexhtmlFrom.THEVALUE)
	}
	if attributevaluexhtmlFrom.THEORIGINALVALUE != nil {
		attributevaluexhtmlTo.THEORIGINALVALUE = CopyBranchXHTMLCONTENT(mapOrigCopy, attributevaluexhtmlFrom.THEORIGINALVALUE)
	}
	if attributevaluexhtmlFrom.DEFINITION != nil {
		attributevaluexhtmlTo.DEFINITION = CopyBranchDEFINITION(mapOrigCopy, attributevaluexhtmlFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCHILDREN(mapOrigCopy map[any]any, childrenFrom *CHILDREN) (childrenTo *CHILDREN) {

	// childrenFrom has already been copied
	if _childrenTo, ok := mapOrigCopy[childrenFrom]; ok {
		childrenTo = _childrenTo.(*CHILDREN)
		return
	}

	childrenTo = new(CHILDREN)
	mapOrigCopy[childrenFrom] = childrenTo
	childrenFrom.CopyBasicFields(childrenTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spechierarchy := range childrenFrom.SPECHIERARCHY {
		childrenTo.SPECHIERARCHY = append(childrenTo.SPECHIERARCHY, CopyBranchSPECHIERARCHY(mapOrigCopy, _spechierarchy))
	}

	return
}

func CopyBranchCORECONTENT(mapOrigCopy map[any]any, corecontentFrom *CORECONTENT) (corecontentTo *CORECONTENT) {

	// corecontentFrom has already been copied
	if _corecontentTo, ok := mapOrigCopy[corecontentFrom]; ok {
		corecontentTo = _corecontentTo.(*CORECONTENT)
		return
	}

	corecontentTo = new(CORECONTENT)
	mapOrigCopy[corecontentFrom] = corecontentTo
	corecontentFrom.CopyBasicFields(corecontentTo)

	//insertion point for the staging of instances referenced by pointers
	if corecontentFrom.REQIFCONTENT != nil {
		corecontentTo.REQIFCONTENT = CopyBranchREQIFCONTENT(mapOrigCopy, corecontentFrom.REQIFCONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONBOOLEAN(mapOrigCopy map[any]any, datatypedefinitionbooleanFrom *DATATYPEDEFINITIONBOOLEAN) (datatypedefinitionbooleanTo *DATATYPEDEFINITIONBOOLEAN) {

	// datatypedefinitionbooleanFrom has already been copied
	if _datatypedefinitionbooleanTo, ok := mapOrigCopy[datatypedefinitionbooleanFrom]; ok {
		datatypedefinitionbooleanTo = _datatypedefinitionbooleanTo.(*DATATYPEDEFINITIONBOOLEAN)
		return
	}

	datatypedefinitionbooleanTo = new(DATATYPEDEFINITIONBOOLEAN)
	mapOrigCopy[datatypedefinitionbooleanFrom] = datatypedefinitionbooleanTo
	datatypedefinitionbooleanFrom.CopyBasicFields(datatypedefinitionbooleanTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionbooleanFrom.ALTERNATIVEID != nil {
		datatypedefinitionbooleanTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionbooleanFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONDATE(mapOrigCopy map[any]any, datatypedefinitiondateFrom *DATATYPEDEFINITIONDATE) (datatypedefinitiondateTo *DATATYPEDEFINITIONDATE) {

	// datatypedefinitiondateFrom has already been copied
	if _datatypedefinitiondateTo, ok := mapOrigCopy[datatypedefinitiondateFrom]; ok {
		datatypedefinitiondateTo = _datatypedefinitiondateTo.(*DATATYPEDEFINITIONDATE)
		return
	}

	datatypedefinitiondateTo = new(DATATYPEDEFINITIONDATE)
	mapOrigCopy[datatypedefinitiondateFrom] = datatypedefinitiondateTo
	datatypedefinitiondateFrom.CopyBasicFields(datatypedefinitiondateTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitiondateFrom.ALTERNATIVEID != nil {
		datatypedefinitiondateTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitiondateFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONENUMERATION(mapOrigCopy map[any]any, datatypedefinitionenumerationFrom *DATATYPEDEFINITIONENUMERATION) (datatypedefinitionenumerationTo *DATATYPEDEFINITIONENUMERATION) {

	// datatypedefinitionenumerationFrom has already been copied
	if _datatypedefinitionenumerationTo, ok := mapOrigCopy[datatypedefinitionenumerationFrom]; ok {
		datatypedefinitionenumerationTo = _datatypedefinitionenumerationTo.(*DATATYPEDEFINITIONENUMERATION)
		return
	}

	datatypedefinitionenumerationTo = new(DATATYPEDEFINITIONENUMERATION)
	mapOrigCopy[datatypedefinitionenumerationFrom] = datatypedefinitionenumerationTo
	datatypedefinitionenumerationFrom.CopyBasicFields(datatypedefinitionenumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionenumerationFrom.ALTERNATIVEID != nil {
		datatypedefinitionenumerationTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionenumerationFrom.ALTERNATIVEID)
	}
	if datatypedefinitionenumerationFrom.SPECIFIEDVALUES != nil {
		datatypedefinitionenumerationTo.SPECIFIEDVALUES = CopyBranchSPECIFIEDVALUES(mapOrigCopy, datatypedefinitionenumerationFrom.SPECIFIEDVALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONINTEGER(mapOrigCopy map[any]any, datatypedefinitionintegerFrom *DATATYPEDEFINITIONINTEGER) (datatypedefinitionintegerTo *DATATYPEDEFINITIONINTEGER) {

	// datatypedefinitionintegerFrom has already been copied
	if _datatypedefinitionintegerTo, ok := mapOrigCopy[datatypedefinitionintegerFrom]; ok {
		datatypedefinitionintegerTo = _datatypedefinitionintegerTo.(*DATATYPEDEFINITIONINTEGER)
		return
	}

	datatypedefinitionintegerTo = new(DATATYPEDEFINITIONINTEGER)
	mapOrigCopy[datatypedefinitionintegerFrom] = datatypedefinitionintegerTo
	datatypedefinitionintegerFrom.CopyBasicFields(datatypedefinitionintegerTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionintegerFrom.ALTERNATIVEID != nil {
		datatypedefinitionintegerTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionintegerFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONREAL(mapOrigCopy map[any]any, datatypedefinitionrealFrom *DATATYPEDEFINITIONREAL) (datatypedefinitionrealTo *DATATYPEDEFINITIONREAL) {

	// datatypedefinitionrealFrom has already been copied
	if _datatypedefinitionrealTo, ok := mapOrigCopy[datatypedefinitionrealFrom]; ok {
		datatypedefinitionrealTo = _datatypedefinitionrealTo.(*DATATYPEDEFINITIONREAL)
		return
	}

	datatypedefinitionrealTo = new(DATATYPEDEFINITIONREAL)
	mapOrigCopy[datatypedefinitionrealFrom] = datatypedefinitionrealTo
	datatypedefinitionrealFrom.CopyBasicFields(datatypedefinitionrealTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionrealFrom.ALTERNATIVEID != nil {
		datatypedefinitionrealTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionrealFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONSTRING(mapOrigCopy map[any]any, datatypedefinitionstringFrom *DATATYPEDEFINITIONSTRING) (datatypedefinitionstringTo *DATATYPEDEFINITIONSTRING) {

	// datatypedefinitionstringFrom has already been copied
	if _datatypedefinitionstringTo, ok := mapOrigCopy[datatypedefinitionstringFrom]; ok {
		datatypedefinitionstringTo = _datatypedefinitionstringTo.(*DATATYPEDEFINITIONSTRING)
		return
	}

	datatypedefinitionstringTo = new(DATATYPEDEFINITIONSTRING)
	mapOrigCopy[datatypedefinitionstringFrom] = datatypedefinitionstringTo
	datatypedefinitionstringFrom.CopyBasicFields(datatypedefinitionstringTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionstringFrom.ALTERNATIVEID != nil {
		datatypedefinitionstringTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionstringFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPEDEFINITIONXHTML(mapOrigCopy map[any]any, datatypedefinitionxhtmlFrom *DATATYPEDEFINITIONXHTML) (datatypedefinitionxhtmlTo *DATATYPEDEFINITIONXHTML) {

	// datatypedefinitionxhtmlFrom has already been copied
	if _datatypedefinitionxhtmlTo, ok := mapOrigCopy[datatypedefinitionxhtmlFrom]; ok {
		datatypedefinitionxhtmlTo = _datatypedefinitionxhtmlTo.(*DATATYPEDEFINITIONXHTML)
		return
	}

	datatypedefinitionxhtmlTo = new(DATATYPEDEFINITIONXHTML)
	mapOrigCopy[datatypedefinitionxhtmlFrom] = datatypedefinitionxhtmlTo
	datatypedefinitionxhtmlFrom.CopyBasicFields(datatypedefinitionxhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionxhtmlFrom.ALTERNATIVEID != nil {
		datatypedefinitionxhtmlTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, datatypedefinitionxhtmlFrom.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPES(mapOrigCopy map[any]any, datatypesFrom *DATATYPES) (datatypesTo *DATATYPES) {

	// datatypesFrom has already been copied
	if _datatypesTo, ok := mapOrigCopy[datatypesFrom]; ok {
		datatypesTo = _datatypesTo.(*DATATYPES)
		return
	}

	datatypesTo = new(DATATYPES)
	mapOrigCopy[datatypesFrom] = datatypesTo
	datatypesFrom.CopyBasicFields(datatypesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatypedefinitionboolean := range datatypesFrom.DATATYPEDEFINITIONBOOLEAN {
		datatypesTo.DATATYPEDEFINITIONBOOLEAN = append(datatypesTo.DATATYPEDEFINITIONBOOLEAN, CopyBranchDATATYPEDEFINITIONBOOLEAN(mapOrigCopy, _datatypedefinitionboolean))
	}
	for _, _datatypedefinitiondate := range datatypesFrom.DATATYPEDEFINITIONDATE {
		datatypesTo.DATATYPEDEFINITIONDATE = append(datatypesTo.DATATYPEDEFINITIONDATE, CopyBranchDATATYPEDEFINITIONDATE(mapOrigCopy, _datatypedefinitiondate))
	}
	for _, _datatypedefinitionenumeration := range datatypesFrom.DATATYPEDEFINITIONENUMERATION {
		datatypesTo.DATATYPEDEFINITIONENUMERATION = append(datatypesTo.DATATYPEDEFINITIONENUMERATION, CopyBranchDATATYPEDEFINITIONENUMERATION(mapOrigCopy, _datatypedefinitionenumeration))
	}
	for _, _datatypedefinitioninteger := range datatypesFrom.DATATYPEDEFINITIONINTEGER {
		datatypesTo.DATATYPEDEFINITIONINTEGER = append(datatypesTo.DATATYPEDEFINITIONINTEGER, CopyBranchDATATYPEDEFINITIONINTEGER(mapOrigCopy, _datatypedefinitioninteger))
	}
	for _, _datatypedefinitionreal := range datatypesFrom.DATATYPEDEFINITIONREAL {
		datatypesTo.DATATYPEDEFINITIONREAL = append(datatypesTo.DATATYPEDEFINITIONREAL, CopyBranchDATATYPEDEFINITIONREAL(mapOrigCopy, _datatypedefinitionreal))
	}
	for _, _datatypedefinitionstring := range datatypesFrom.DATATYPEDEFINITIONSTRING {
		datatypesTo.DATATYPEDEFINITIONSTRING = append(datatypesTo.DATATYPEDEFINITIONSTRING, CopyBranchDATATYPEDEFINITIONSTRING(mapOrigCopy, _datatypedefinitionstring))
	}
	for _, _datatypedefinitionxhtml := range datatypesFrom.DATATYPEDEFINITIONXHTML {
		datatypesTo.DATATYPEDEFINITIONXHTML = append(datatypesTo.DATATYPEDEFINITIONXHTML, CopyBranchDATATYPEDEFINITIONXHTML(mapOrigCopy, _datatypedefinitionxhtml))
	}

	return
}

func CopyBranchDEFAULTVALUE(mapOrigCopy map[any]any, defaultvalueFrom *DEFAULTVALUE) (defaultvalueTo *DEFAULTVALUE) {

	// defaultvalueFrom has already been copied
	if _defaultvalueTo, ok := mapOrigCopy[defaultvalueFrom]; ok {
		defaultvalueTo = _defaultvalueTo.(*DEFAULTVALUE)
		return
	}

	defaultvalueTo = new(DEFAULTVALUE)
	mapOrigCopy[defaultvalueFrom] = defaultvalueTo
	defaultvalueFrom.CopyBasicFields(defaultvalueTo)

	//insertion point for the staging of instances referenced by pointers
	if defaultvalueFrom.ATTRIBUTEVALUEBOOLEAN != nil {
		defaultvalueTo.ATTRIBUTEVALUEBOOLEAN = CopyBranchATTRIBUTEVALUEBOOLEAN(mapOrigCopy, defaultvalueFrom.ATTRIBUTEVALUEBOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDEFINITION(mapOrigCopy map[any]any, definitionFrom *DEFINITION) (definitionTo *DEFINITION) {

	// definitionFrom has already been copied
	if _definitionTo, ok := mapOrigCopy[definitionFrom]; ok {
		definitionTo = _definitionTo.(*DEFINITION)
		return
	}

	definitionTo = new(DEFINITION)
	mapOrigCopy[definitionFrom] = definitionTo
	definitionFrom.CopyBasicFields(definitionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEDITABLEATTS(mapOrigCopy map[any]any, editableattsFrom *EDITABLEATTS) (editableattsTo *EDITABLEATTS) {

	// editableattsFrom has already been copied
	if _editableattsTo, ok := mapOrigCopy[editableattsFrom]; ok {
		editableattsTo = _editableattsTo.(*EDITABLEATTS)
		return
	}

	editableattsTo = new(EDITABLEATTS)
	mapOrigCopy[editableattsFrom] = editableattsTo
	editableattsFrom.CopyBasicFields(editableattsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEMBEDDEDVALUE(mapOrigCopy map[any]any, embeddedvalueFrom *EMBEDDEDVALUE) (embeddedvalueTo *EMBEDDEDVALUE) {

	// embeddedvalueFrom has already been copied
	if _embeddedvalueTo, ok := mapOrigCopy[embeddedvalueFrom]; ok {
		embeddedvalueTo = _embeddedvalueTo.(*EMBEDDEDVALUE)
		return
	}

	embeddedvalueTo = new(EMBEDDEDVALUE)
	mapOrigCopy[embeddedvalueFrom] = embeddedvalueTo
	embeddedvalueFrom.CopyBasicFields(embeddedvalueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchENUMVALUE(mapOrigCopy map[any]any, enumvalueFrom *ENUMVALUE) (enumvalueTo *ENUMVALUE) {

	// enumvalueFrom has already been copied
	if _enumvalueTo, ok := mapOrigCopy[enumvalueFrom]; ok {
		enumvalueTo = _enumvalueTo.(*ENUMVALUE)
		return
	}

	enumvalueTo = new(ENUMVALUE)
	mapOrigCopy[enumvalueFrom] = enumvalueTo
	enumvalueFrom.CopyBasicFields(enumvalueTo)

	//insertion point for the staging of instances referenced by pointers
	if enumvalueFrom.ALTERNATIVEID != nil {
		enumvalueTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, enumvalueFrom.ALTERNATIVEID)
	}
	if enumvalueFrom.PROPERTIES != nil {
		enumvalueTo.PROPERTIES = CopyBranchPROPERTIES(mapOrigCopy, enumvalueFrom.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOBJECT(mapOrigCopy map[any]any, objectFrom *OBJECT) (objectTo *OBJECT) {

	// objectFrom has already been copied
	if _objectTo, ok := mapOrigCopy[objectFrom]; ok {
		objectTo = _objectTo.(*OBJECT)
		return
	}

	objectTo = new(OBJECT)
	mapOrigCopy[objectFrom] = objectTo
	objectFrom.CopyBasicFields(objectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPROPERTIES(mapOrigCopy map[any]any, propertiesFrom *PROPERTIES) (propertiesTo *PROPERTIES) {

	// propertiesFrom has already been copied
	if _propertiesTo, ok := mapOrigCopy[propertiesFrom]; ok {
		propertiesTo = _propertiesTo.(*PROPERTIES)
		return
	}

	propertiesTo = new(PROPERTIES)
	mapOrigCopy[propertiesFrom] = propertiesTo
	propertiesFrom.CopyBasicFields(propertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if propertiesFrom.EMBEDDEDVALUE != nil {
		propertiesTo.EMBEDDEDVALUE = CopyBranchEMBEDDEDVALUE(mapOrigCopy, propertiesFrom.EMBEDDEDVALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRELATIONGROUP(mapOrigCopy map[any]any, relationgroupFrom *RELATIONGROUP) (relationgroupTo *RELATIONGROUP) {

	// relationgroupFrom has already been copied
	if _relationgroupTo, ok := mapOrigCopy[relationgroupFrom]; ok {
		relationgroupTo = _relationgroupTo.(*RELATIONGROUP)
		return
	}

	relationgroupTo = new(RELATIONGROUP)
	mapOrigCopy[relationgroupFrom] = relationgroupTo
	relationgroupFrom.CopyBasicFields(relationgroupTo)

	//insertion point for the staging of instances referenced by pointers
	if relationgroupFrom.ALTERNATIVEID != nil {
		relationgroupTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, relationgroupFrom.ALTERNATIVEID)
	}
	if relationgroupFrom.SOURCESPECIFICATION != nil {
		relationgroupTo.SOURCESPECIFICATION = CopyBranchSOURCESPECIFICATION(mapOrigCopy, relationgroupFrom.SOURCESPECIFICATION)
	}
	if relationgroupFrom.SPECRELATIONS != nil {
		relationgroupTo.SPECRELATIONS = CopyBranchSPECRELATIONS(mapOrigCopy, relationgroupFrom.SPECRELATIONS)
	}
	if relationgroupFrom.TARGETSPECIFICATION != nil {
		relationgroupTo.TARGETSPECIFICATION = CopyBranchTARGETSPECIFICATION(mapOrigCopy, relationgroupFrom.TARGETSPECIFICATION)
	}
	if relationgroupFrom.TYPE != nil {
		relationgroupTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, relationgroupFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRELATIONGROUPTYPE(mapOrigCopy map[any]any, relationgrouptypeFrom *RELATIONGROUPTYPE) (relationgrouptypeTo *RELATIONGROUPTYPE) {

	// relationgrouptypeFrom has already been copied
	if _relationgrouptypeTo, ok := mapOrigCopy[relationgrouptypeFrom]; ok {
		relationgrouptypeTo = _relationgrouptypeTo.(*RELATIONGROUPTYPE)
		return
	}

	relationgrouptypeTo = new(RELATIONGROUPTYPE)
	mapOrigCopy[relationgrouptypeFrom] = relationgrouptypeTo
	relationgrouptypeFrom.CopyBasicFields(relationgrouptypeTo)

	//insertion point for the staging of instances referenced by pointers
	if relationgrouptypeFrom.ALTERNATIVEID != nil {
		relationgrouptypeTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, relationgrouptypeFrom.ALTERNATIVEID)
	}
	if relationgrouptypeFrom.SPECATTRIBUTES != nil {
		relationgrouptypeTo.SPECATTRIBUTES = CopyBranchSPECATTRIBUTES(mapOrigCopy, relationgrouptypeFrom.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIF(mapOrigCopy map[any]any, reqifFrom *REQIF) (reqifTo *REQIF) {

	// reqifFrom has already been copied
	if _reqifTo, ok := mapOrigCopy[reqifFrom]; ok {
		reqifTo = _reqifTo.(*REQIF)
		return
	}

	reqifTo = new(REQIF)
	mapOrigCopy[reqifFrom] = reqifTo
	reqifFrom.CopyBasicFields(reqifTo)

	//insertion point for the staging of instances referenced by pointers
	if reqifFrom.HEADER != nil {
		reqifTo.HEADER = CopyBranchTHEHEADER(mapOrigCopy, reqifFrom.HEADER)
	}
	if reqifFrom.CORECONTENT != nil {
		reqifTo.CORECONTENT = CopyBranchCORECONTENT(mapOrigCopy, reqifFrom.CORECONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIFCONTENT(mapOrigCopy map[any]any, reqifcontentFrom *REQIFCONTENT) (reqifcontentTo *REQIFCONTENT) {

	// reqifcontentFrom has already been copied
	if _reqifcontentTo, ok := mapOrigCopy[reqifcontentFrom]; ok {
		reqifcontentTo = _reqifcontentTo.(*REQIFCONTENT)
		return
	}

	reqifcontentTo = new(REQIFCONTENT)
	mapOrigCopy[reqifcontentFrom] = reqifcontentTo
	reqifcontentFrom.CopyBasicFields(reqifcontentTo)

	//insertion point for the staging of instances referenced by pointers
	if reqifcontentFrom.DATATYPES != nil {
		reqifcontentTo.DATATYPES = CopyBranchDATATYPES(mapOrigCopy, reqifcontentFrom.DATATYPES)
	}
	if reqifcontentFrom.SPECTYPES != nil {
		reqifcontentTo.SPECTYPES = CopyBranchSPECTYPES(mapOrigCopy, reqifcontentFrom.SPECTYPES)
	}
	if reqifcontentFrom.SPECOBJECTS != nil {
		reqifcontentTo.SPECOBJECTS = CopyBranchSPECOBJECTS(mapOrigCopy, reqifcontentFrom.SPECOBJECTS)
	}
	if reqifcontentFrom.SPECRELATIONS != nil {
		reqifcontentTo.SPECRELATIONS = CopyBranchSPECRELATIONS(mapOrigCopy, reqifcontentFrom.SPECRELATIONS)
	}
	if reqifcontentFrom.SPECIFICATIONS != nil {
		reqifcontentTo.SPECIFICATIONS = CopyBranchSPECIFICATIONS(mapOrigCopy, reqifcontentFrom.SPECIFICATIONS)
	}
	if reqifcontentFrom.SPECRELATIONGROUPS != nil {
		reqifcontentTo.SPECRELATIONGROUPS = CopyBranchSPECRELATIONGROUPS(mapOrigCopy, reqifcontentFrom.SPECRELATIONGROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIFHEADER(mapOrigCopy map[any]any, reqifheaderFrom *REQIFHEADER) (reqifheaderTo *REQIFHEADER) {

	// reqifheaderFrom has already been copied
	if _reqifheaderTo, ok := mapOrigCopy[reqifheaderFrom]; ok {
		reqifheaderTo = _reqifheaderTo.(*REQIFHEADER)
		return
	}

	reqifheaderTo = new(REQIFHEADER)
	mapOrigCopy[reqifheaderFrom] = reqifheaderTo
	reqifheaderFrom.CopyBasicFields(reqifheaderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIFTOOLEXTENSION(mapOrigCopy map[any]any, reqiftoolextensionFrom *REQIFTOOLEXTENSION) (reqiftoolextensionTo *REQIFTOOLEXTENSION) {

	// reqiftoolextensionFrom has already been copied
	if _reqiftoolextensionTo, ok := mapOrigCopy[reqiftoolextensionFrom]; ok {
		reqiftoolextensionTo = _reqiftoolextensionTo.(*REQIFTOOLEXTENSION)
		return
	}

	reqiftoolextensionTo = new(REQIFTOOLEXTENSION)
	mapOrigCopy[reqiftoolextensionFrom] = reqiftoolextensionTo
	reqiftoolextensionFrom.CopyBasicFields(reqiftoolextensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIFTYPE(mapOrigCopy map[any]any, reqiftypeFrom *REQIFTYPE) (reqiftypeTo *REQIFTYPE) {

	// reqiftypeFrom has already been copied
	if _reqiftypeTo, ok := mapOrigCopy[reqiftypeFrom]; ok {
		reqiftypeTo = _reqiftypeTo.(*REQIFTYPE)
		return
	}

	reqiftypeTo = new(REQIFTYPE)
	mapOrigCopy[reqiftypeFrom] = reqiftypeTo
	reqiftypeFrom.CopyBasicFields(reqiftypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSOURCE(mapOrigCopy map[any]any, sourceFrom *SOURCE) (sourceTo *SOURCE) {

	// sourceFrom has already been copied
	if _sourceTo, ok := mapOrigCopy[sourceFrom]; ok {
		sourceTo = _sourceTo.(*SOURCE)
		return
	}

	sourceTo = new(SOURCE)
	mapOrigCopy[sourceFrom] = sourceTo
	sourceFrom.CopyBasicFields(sourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSOURCESPECIFICATION(mapOrigCopy map[any]any, sourcespecificationFrom *SOURCESPECIFICATION) (sourcespecificationTo *SOURCESPECIFICATION) {

	// sourcespecificationFrom has already been copied
	if _sourcespecificationTo, ok := mapOrigCopy[sourcespecificationFrom]; ok {
		sourcespecificationTo = _sourcespecificationTo.(*SOURCESPECIFICATION)
		return
	}

	sourcespecificationTo = new(SOURCESPECIFICATION)
	mapOrigCopy[sourcespecificationFrom] = sourcespecificationTo
	sourcespecificationFrom.CopyBasicFields(sourcespecificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECATTRIBUTES(mapOrigCopy map[any]any, specattributesFrom *SPECATTRIBUTES) (specattributesTo *SPECATTRIBUTES) {

	// specattributesFrom has already been copied
	if _specattributesTo, ok := mapOrigCopy[specattributesFrom]; ok {
		specattributesTo = _specattributesTo.(*SPECATTRIBUTES)
		return
	}

	specattributesTo = new(SPECATTRIBUTES)
	mapOrigCopy[specattributesFrom] = specattributesTo
	specattributesFrom.CopyBasicFields(specattributesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributedefinitionboolean := range specattributesFrom.ATTRIBUTEDEFINITIONBOOLEAN {
		specattributesTo.ATTRIBUTEDEFINITIONBOOLEAN = append(specattributesTo.ATTRIBUTEDEFINITIONBOOLEAN, CopyBranchATTRIBUTEDEFINITIONBOOLEAN(mapOrigCopy, _attributedefinitionboolean))
	}
	for _, _attributedefinitiondate := range specattributesFrom.ATTRIBUTEDEFINITIONDATE {
		specattributesTo.ATTRIBUTEDEFINITIONDATE = append(specattributesTo.ATTRIBUTEDEFINITIONDATE, CopyBranchATTRIBUTEDEFINITIONDATE(mapOrigCopy, _attributedefinitiondate))
	}
	for _, _attributedefinitionenumeration := range specattributesFrom.ATTRIBUTEDEFINITIONENUMERATION {
		specattributesTo.ATTRIBUTEDEFINITIONENUMERATION = append(specattributesTo.ATTRIBUTEDEFINITIONENUMERATION, CopyBranchATTRIBUTEDEFINITIONENUMERATION(mapOrigCopy, _attributedefinitionenumeration))
	}
	for _, _attributedefinitioninteger := range specattributesFrom.ATTRIBUTEDEFINITIONINTEGER {
		specattributesTo.ATTRIBUTEDEFINITIONINTEGER = append(specattributesTo.ATTRIBUTEDEFINITIONINTEGER, CopyBranchATTRIBUTEDEFINITIONINTEGER(mapOrigCopy, _attributedefinitioninteger))
	}
	for _, _attributedefinitionreal := range specattributesFrom.ATTRIBUTEDEFINITIONREAL {
		specattributesTo.ATTRIBUTEDEFINITIONREAL = append(specattributesTo.ATTRIBUTEDEFINITIONREAL, CopyBranchATTRIBUTEDEFINITIONREAL(mapOrigCopy, _attributedefinitionreal))
	}
	for _, _attributedefinitionstring := range specattributesFrom.ATTRIBUTEDEFINITIONSTRING {
		specattributesTo.ATTRIBUTEDEFINITIONSTRING = append(specattributesTo.ATTRIBUTEDEFINITIONSTRING, CopyBranchATTRIBUTEDEFINITIONSTRING(mapOrigCopy, _attributedefinitionstring))
	}
	for _, _attributedefinitionxhtml := range specattributesFrom.ATTRIBUTEDEFINITIONXHTML {
		specattributesTo.ATTRIBUTEDEFINITIONXHTML = append(specattributesTo.ATTRIBUTEDEFINITIONXHTML, CopyBranchATTRIBUTEDEFINITIONXHTML(mapOrigCopy, _attributedefinitionxhtml))
	}

	return
}

func CopyBranchSPECHIERARCHY(mapOrigCopy map[any]any, spechierarchyFrom *SPECHIERARCHY) (spechierarchyTo *SPECHIERARCHY) {

	// spechierarchyFrom has already been copied
	if _spechierarchyTo, ok := mapOrigCopy[spechierarchyFrom]; ok {
		spechierarchyTo = _spechierarchyTo.(*SPECHIERARCHY)
		return
	}

	spechierarchyTo = new(SPECHIERARCHY)
	mapOrigCopy[spechierarchyFrom] = spechierarchyTo
	spechierarchyFrom.CopyBasicFields(spechierarchyTo)

	//insertion point for the staging of instances referenced by pointers
	if spechierarchyFrom.ALTERNATIVEID != nil {
		spechierarchyTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, spechierarchyFrom.ALTERNATIVEID)
	}
	if spechierarchyFrom.CHILDREN != nil {
		spechierarchyTo.CHILDREN = CopyBranchCHILDREN(mapOrigCopy, spechierarchyFrom.CHILDREN)
	}
	if spechierarchyFrom.EDITABLEATTS != nil {
		spechierarchyTo.EDITABLEATTS = CopyBranchEDITABLEATTS(mapOrigCopy, spechierarchyFrom.EDITABLEATTS)
	}
	if spechierarchyFrom.OBJECT != nil {
		spechierarchyTo.OBJECT = CopyBranchOBJECT(mapOrigCopy, spechierarchyFrom.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFICATION(mapOrigCopy map[any]any, specificationFrom *SPECIFICATION) (specificationTo *SPECIFICATION) {

	// specificationFrom has already been copied
	if _specificationTo, ok := mapOrigCopy[specificationFrom]; ok {
		specificationTo = _specificationTo.(*SPECIFICATION)
		return
	}

	specificationTo = new(SPECIFICATION)
	mapOrigCopy[specificationFrom] = specificationTo
	specificationFrom.CopyBasicFields(specificationTo)

	//insertion point for the staging of instances referenced by pointers
	if specificationFrom.ALTERNATIVEID != nil {
		specificationTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specificationFrom.ALTERNATIVEID)
	}
	if specificationFrom.VALUES != nil {
		specificationTo.VALUES = CopyBranchVALUES(mapOrigCopy, specificationFrom.VALUES)
	}
	if specificationFrom.CHILDREN != nil {
		specificationTo.CHILDREN = CopyBranchCHILDREN(mapOrigCopy, specificationFrom.CHILDREN)
	}
	if specificationFrom.TYPE != nil {
		specificationTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, specificationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFICATIONS(mapOrigCopy map[any]any, specificationsFrom *SPECIFICATIONS) (specificationsTo *SPECIFICATIONS) {

	// specificationsFrom has already been copied
	if _specificationsTo, ok := mapOrigCopy[specificationsFrom]; ok {
		specificationsTo = _specificationsTo.(*SPECIFICATIONS)
		return
	}

	specificationsTo = new(SPECIFICATIONS)
	mapOrigCopy[specificationsFrom] = specificationsTo
	specificationsFrom.CopyBasicFields(specificationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range specificationsFrom.SPECIFICATION {
		specificationsTo.SPECIFICATION = append(specificationsTo.SPECIFICATION, CopyBranchSPECIFICATION(mapOrigCopy, _specification))
	}

	return
}

func CopyBranchSPECIFICATIONTYPE(mapOrigCopy map[any]any, specificationtypeFrom *SPECIFICATIONTYPE) (specificationtypeTo *SPECIFICATIONTYPE) {

	// specificationtypeFrom has already been copied
	if _specificationtypeTo, ok := mapOrigCopy[specificationtypeFrom]; ok {
		specificationtypeTo = _specificationtypeTo.(*SPECIFICATIONTYPE)
		return
	}

	specificationtypeTo = new(SPECIFICATIONTYPE)
	mapOrigCopy[specificationtypeFrom] = specificationtypeTo
	specificationtypeFrom.CopyBasicFields(specificationtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if specificationtypeFrom.ALTERNATIVEID != nil {
		specificationtypeTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specificationtypeFrom.ALTERNATIVEID)
	}
	if specificationtypeFrom.SPECATTRIBUTES != nil {
		specificationtypeTo.SPECATTRIBUTES = CopyBranchSPECATTRIBUTES(mapOrigCopy, specificationtypeFrom.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFIEDVALUES(mapOrigCopy map[any]any, specifiedvaluesFrom *SPECIFIEDVALUES) (specifiedvaluesTo *SPECIFIEDVALUES) {

	// specifiedvaluesFrom has already been copied
	if _specifiedvaluesTo, ok := mapOrigCopy[specifiedvaluesFrom]; ok {
		specifiedvaluesTo = _specifiedvaluesTo.(*SPECIFIEDVALUES)
		return
	}

	specifiedvaluesTo = new(SPECIFIEDVALUES)
	mapOrigCopy[specifiedvaluesFrom] = specifiedvaluesTo
	specifiedvaluesFrom.CopyBasicFields(specifiedvaluesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumvalue := range specifiedvaluesFrom.ENUMVALUE {
		specifiedvaluesTo.ENUMVALUE = append(specifiedvaluesTo.ENUMVALUE, CopyBranchENUMVALUE(mapOrigCopy, _enumvalue))
	}

	return
}

func CopyBranchSPECOBJECT(mapOrigCopy map[any]any, specobjectFrom *SPECOBJECT) (specobjectTo *SPECOBJECT) {

	// specobjectFrom has already been copied
	if _specobjectTo, ok := mapOrigCopy[specobjectFrom]; ok {
		specobjectTo = _specobjectTo.(*SPECOBJECT)
		return
	}

	specobjectTo = new(SPECOBJECT)
	mapOrigCopy[specobjectFrom] = specobjectTo
	specobjectFrom.CopyBasicFields(specobjectTo)

	//insertion point for the staging of instances referenced by pointers
	if specobjectFrom.ALTERNATIVEID != nil {
		specobjectTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specobjectFrom.ALTERNATIVEID)
	}
	if specobjectFrom.VALUES != nil {
		specobjectTo.VALUES = CopyBranchVALUES(mapOrigCopy, specobjectFrom.VALUES)
	}
	if specobjectFrom.TYPE != nil {
		specobjectTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, specobjectFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECOBJECTS(mapOrigCopy map[any]any, specobjectsFrom *SPECOBJECTS) (specobjectsTo *SPECOBJECTS) {

	// specobjectsFrom has already been copied
	if _specobjectsTo, ok := mapOrigCopy[specobjectsFrom]; ok {
		specobjectsTo = _specobjectsTo.(*SPECOBJECTS)
		return
	}

	specobjectsTo = new(SPECOBJECTS)
	mapOrigCopy[specobjectsFrom] = specobjectsTo
	specobjectsFrom.CopyBasicFields(specobjectsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specobject := range specobjectsFrom.SPECOBJECT {
		specobjectsTo.SPECOBJECT = append(specobjectsTo.SPECOBJECT, CopyBranchSPECOBJECT(mapOrigCopy, _specobject))
	}

	return
}

func CopyBranchSPECOBJECTTYPE(mapOrigCopy map[any]any, specobjecttypeFrom *SPECOBJECTTYPE) (specobjecttypeTo *SPECOBJECTTYPE) {

	// specobjecttypeFrom has already been copied
	if _specobjecttypeTo, ok := mapOrigCopy[specobjecttypeFrom]; ok {
		specobjecttypeTo = _specobjecttypeTo.(*SPECOBJECTTYPE)
		return
	}

	specobjecttypeTo = new(SPECOBJECTTYPE)
	mapOrigCopy[specobjecttypeFrom] = specobjecttypeTo
	specobjecttypeFrom.CopyBasicFields(specobjecttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if specobjecttypeFrom.ALTERNATIVEID != nil {
		specobjecttypeTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specobjecttypeFrom.ALTERNATIVEID)
	}
	if specobjecttypeFrom.SPECATTRIBUTES != nil {
		specobjecttypeTo.SPECATTRIBUTES = CopyBranchSPECATTRIBUTES(mapOrigCopy, specobjecttypeFrom.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECRELATION(mapOrigCopy map[any]any, specrelationFrom *SPECRELATION) (specrelationTo *SPECRELATION) {

	// specrelationFrom has already been copied
	if _specrelationTo, ok := mapOrigCopy[specrelationFrom]; ok {
		specrelationTo = _specrelationTo.(*SPECRELATION)
		return
	}

	specrelationTo = new(SPECRELATION)
	mapOrigCopy[specrelationFrom] = specrelationTo
	specrelationFrom.CopyBasicFields(specrelationTo)

	//insertion point for the staging of instances referenced by pointers
	if specrelationFrom.ALTERNATIVEID != nil {
		specrelationTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specrelationFrom.ALTERNATIVEID)
	}
	if specrelationFrom.VALUES != nil {
		specrelationTo.VALUES = CopyBranchVALUES(mapOrigCopy, specrelationFrom.VALUES)
	}
	if specrelationFrom.SOURCE != nil {
		specrelationTo.SOURCE = CopyBranchSOURCE(mapOrigCopy, specrelationFrom.SOURCE)
	}
	if specrelationFrom.TARGET != nil {
		specrelationTo.TARGET = CopyBranchTARGET(mapOrigCopy, specrelationFrom.TARGET)
	}
	if specrelationFrom.TYPE != nil {
		specrelationTo.TYPE = CopyBranchREQIFTYPE(mapOrigCopy, specrelationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECRELATIONGROUPS(mapOrigCopy map[any]any, specrelationgroupsFrom *SPECRELATIONGROUPS) (specrelationgroupsTo *SPECRELATIONGROUPS) {

	// specrelationgroupsFrom has already been copied
	if _specrelationgroupsTo, ok := mapOrigCopy[specrelationgroupsFrom]; ok {
		specrelationgroupsTo = _specrelationgroupsTo.(*SPECRELATIONGROUPS)
		return
	}

	specrelationgroupsTo = new(SPECRELATIONGROUPS)
	mapOrigCopy[specrelationgroupsFrom] = specrelationgroupsTo
	specrelationgroupsFrom.CopyBasicFields(specrelationgroupsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgroup := range specrelationgroupsFrom.RELATIONGROUP {
		specrelationgroupsTo.RELATIONGROUP = append(specrelationgroupsTo.RELATIONGROUP, CopyBranchRELATIONGROUP(mapOrigCopy, _relationgroup))
	}

	return
}

func CopyBranchSPECRELATIONS(mapOrigCopy map[any]any, specrelationsFrom *SPECRELATIONS) (specrelationsTo *SPECRELATIONS) {

	// specrelationsFrom has already been copied
	if _specrelationsTo, ok := mapOrigCopy[specrelationsFrom]; ok {
		specrelationsTo = _specrelationsTo.(*SPECRELATIONS)
		return
	}

	specrelationsTo = new(SPECRELATIONS)
	mapOrigCopy[specrelationsFrom] = specrelationsTo
	specrelationsFrom.CopyBasicFields(specrelationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECRELATIONTYPE(mapOrigCopy map[any]any, specrelationtypeFrom *SPECRELATIONTYPE) (specrelationtypeTo *SPECRELATIONTYPE) {

	// specrelationtypeFrom has already been copied
	if _specrelationtypeTo, ok := mapOrigCopy[specrelationtypeFrom]; ok {
		specrelationtypeTo = _specrelationtypeTo.(*SPECRELATIONTYPE)
		return
	}

	specrelationtypeTo = new(SPECRELATIONTYPE)
	mapOrigCopy[specrelationtypeFrom] = specrelationtypeTo
	specrelationtypeFrom.CopyBasicFields(specrelationtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if specrelationtypeFrom.ALTERNATIVEID != nil {
		specrelationtypeTo.ALTERNATIVEID = CopyBranchALTERNATIVEID(mapOrigCopy, specrelationtypeFrom.ALTERNATIVEID)
	}
	if specrelationtypeFrom.SPECATTRIBUTES != nil {
		specrelationtypeTo.SPECATTRIBUTES = CopyBranchSPECATTRIBUTES(mapOrigCopy, specrelationtypeFrom.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECTYPES(mapOrigCopy map[any]any, spectypesFrom *SPECTYPES) (spectypesTo *SPECTYPES) {

	// spectypesFrom has already been copied
	if _spectypesTo, ok := mapOrigCopy[spectypesFrom]; ok {
		spectypesTo = _spectypesTo.(*SPECTYPES)
		return
	}

	spectypesTo = new(SPECTYPES)
	mapOrigCopy[spectypesFrom] = spectypesTo
	spectypesFrom.CopyBasicFields(spectypesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgrouptype := range spectypesFrom.RELATIONGROUPTYPE {
		spectypesTo.RELATIONGROUPTYPE = append(spectypesTo.RELATIONGROUPTYPE, CopyBranchRELATIONGROUPTYPE(mapOrigCopy, _relationgrouptype))
	}
	for _, _specobjecttype := range spectypesFrom.SPECOBJECTTYPE {
		spectypesTo.SPECOBJECTTYPE = append(spectypesTo.SPECOBJECTTYPE, CopyBranchSPECOBJECTTYPE(mapOrigCopy, _specobjecttype))
	}
	for _, _specrelationtype := range spectypesFrom.SPECRELATIONTYPE {
		spectypesTo.SPECRELATIONTYPE = append(spectypesTo.SPECRELATIONTYPE, CopyBranchSPECRELATIONTYPE(mapOrigCopy, _specrelationtype))
	}
	for _, _specificationtype := range spectypesFrom.SPECIFICATIONTYPE {
		spectypesTo.SPECIFICATIONTYPE = append(spectypesTo.SPECIFICATIONTYPE, CopyBranchSPECIFICATIONTYPE(mapOrigCopy, _specificationtype))
	}

	return
}

func CopyBranchTARGET(mapOrigCopy map[any]any, targetFrom *TARGET) (targetTo *TARGET) {

	// targetFrom has already been copied
	if _targetTo, ok := mapOrigCopy[targetFrom]; ok {
		targetTo = _targetTo.(*TARGET)
		return
	}

	targetTo = new(TARGET)
	mapOrigCopy[targetFrom] = targetTo
	targetFrom.CopyBasicFields(targetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTARGETSPECIFICATION(mapOrigCopy map[any]any, targetspecificationFrom *TARGETSPECIFICATION) (targetspecificationTo *TARGETSPECIFICATION) {

	// targetspecificationFrom has already been copied
	if _targetspecificationTo, ok := mapOrigCopy[targetspecificationFrom]; ok {
		targetspecificationTo = _targetspecificationTo.(*TARGETSPECIFICATION)
		return
	}

	targetspecificationTo = new(TARGETSPECIFICATION)
	mapOrigCopy[targetspecificationFrom] = targetspecificationTo
	targetspecificationFrom.CopyBasicFields(targetspecificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTHEHEADER(mapOrigCopy map[any]any, theheaderFrom *THEHEADER) (theheaderTo *THEHEADER) {

	// theheaderFrom has already been copied
	if _theheaderTo, ok := mapOrigCopy[theheaderFrom]; ok {
		theheaderTo = _theheaderTo.(*THEHEADER)
		return
	}

	theheaderTo = new(THEHEADER)
	mapOrigCopy[theheaderFrom] = theheaderTo
	theheaderFrom.CopyBasicFields(theheaderTo)

	//insertion point for the staging of instances referenced by pointers
	if theheaderFrom.REQIFHEADER != nil {
		theheaderTo.REQIFHEADER = CopyBranchREQIFHEADER(mapOrigCopy, theheaderFrom.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTOOLEXTENSIONS(mapOrigCopy map[any]any, toolextensionsFrom *TOOLEXTENSIONS) (toolextensionsTo *TOOLEXTENSIONS) {

	// toolextensionsFrom has already been copied
	if _toolextensionsTo, ok := mapOrigCopy[toolextensionsFrom]; ok {
		toolextensionsTo = _toolextensionsTo.(*TOOLEXTENSIONS)
		return
	}

	toolextensionsTo = new(TOOLEXTENSIONS)
	mapOrigCopy[toolextensionsFrom] = toolextensionsTo
	toolextensionsFrom.CopyBasicFields(toolextensionsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _reqiftoolextension := range toolextensionsFrom.REQIFTOOLEXTENSION {
		toolextensionsTo.REQIFTOOLEXTENSION = append(toolextensionsTo.REQIFTOOLEXTENSION, CopyBranchREQIFTOOLEXTENSION(mapOrigCopy, _reqiftoolextension))
	}

	return
}

func CopyBranchVALUES(mapOrigCopy map[any]any, valuesFrom *VALUES) (valuesTo *VALUES) {

	// valuesFrom has already been copied
	if _valuesTo, ok := mapOrigCopy[valuesFrom]; ok {
		valuesTo = _valuesTo.(*VALUES)
		return
	}

	valuesTo = new(VALUES)
	mapOrigCopy[valuesFrom] = valuesTo
	valuesFrom.CopyBasicFields(valuesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXHTMLCONTENT(mapOrigCopy map[any]any, xhtmlcontentFrom *XHTMLCONTENT) (xhtmlcontentTo *XHTMLCONTENT) {

	// xhtmlcontentFrom has already been copied
	if _xhtmlcontentTo, ok := mapOrigCopy[xhtmlcontentFrom]; ok {
		xhtmlcontentTo = _xhtmlcontentTo.(*XHTMLCONTENT)
		return
	}

	xhtmlcontentTo = new(XHTMLCONTENT)
	mapOrigCopy[xhtmlcontentFrom] = xhtmlcontentTo
	xhtmlcontentFrom.CopyBasicFields(xhtmlcontentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *ALTERNATIVEID:
		stage.UnstageBranchALTERNATIVEID(target)

	case *ATTRIBUTEDEFINITIONBOOLEAN:
		stage.UnstageBranchATTRIBUTEDEFINITIONBOOLEAN(target)

	case *ATTRIBUTEDEFINITIONDATE:
		stage.UnstageBranchATTRIBUTEDEFINITIONDATE(target)

	case *ATTRIBUTEDEFINITIONENUMERATION:
		stage.UnstageBranchATTRIBUTEDEFINITIONENUMERATION(target)

	case *ATTRIBUTEDEFINITIONINTEGER:
		stage.UnstageBranchATTRIBUTEDEFINITIONINTEGER(target)

	case *ATTRIBUTEDEFINITIONREAL:
		stage.UnstageBranchATTRIBUTEDEFINITIONREAL(target)

	case *ATTRIBUTEDEFINITIONSTRING:
		stage.UnstageBranchATTRIBUTEDEFINITIONSTRING(target)

	case *ATTRIBUTEDEFINITIONXHTML:
		stage.UnstageBranchATTRIBUTEDEFINITIONXHTML(target)

	case *ATTRIBUTEVALUEBOOLEAN:
		stage.UnstageBranchATTRIBUTEVALUEBOOLEAN(target)

	case *ATTRIBUTEVALUEDATE:
		stage.UnstageBranchATTRIBUTEVALUEDATE(target)

	case *ATTRIBUTEVALUEENUMERATION:
		stage.UnstageBranchATTRIBUTEVALUEENUMERATION(target)

	case *ATTRIBUTEVALUEINTEGER:
		stage.UnstageBranchATTRIBUTEVALUEINTEGER(target)

	case *ATTRIBUTEVALUEREAL:
		stage.UnstageBranchATTRIBUTEVALUEREAL(target)

	case *ATTRIBUTEVALUESTRING:
		stage.UnstageBranchATTRIBUTEVALUESTRING(target)

	case *ATTRIBUTEVALUEXHTML:
		stage.UnstageBranchATTRIBUTEVALUEXHTML(target)

	case *CHILDREN:
		stage.UnstageBranchCHILDREN(target)

	case *CORECONTENT:
		stage.UnstageBranchCORECONTENT(target)

	case *DATATYPEDEFINITIONBOOLEAN:
		stage.UnstageBranchDATATYPEDEFINITIONBOOLEAN(target)

	case *DATATYPEDEFINITIONDATE:
		stage.UnstageBranchDATATYPEDEFINITIONDATE(target)

	case *DATATYPEDEFINITIONENUMERATION:
		stage.UnstageBranchDATATYPEDEFINITIONENUMERATION(target)

	case *DATATYPEDEFINITIONINTEGER:
		stage.UnstageBranchDATATYPEDEFINITIONINTEGER(target)

	case *DATATYPEDEFINITIONREAL:
		stage.UnstageBranchDATATYPEDEFINITIONREAL(target)

	case *DATATYPEDEFINITIONSTRING:
		stage.UnstageBranchDATATYPEDEFINITIONSTRING(target)

	case *DATATYPEDEFINITIONXHTML:
		stage.UnstageBranchDATATYPEDEFINITIONXHTML(target)

	case *DATATYPES:
		stage.UnstageBranchDATATYPES(target)

	case *DEFAULTVALUE:
		stage.UnstageBranchDEFAULTVALUE(target)

	case *DEFINITION:
		stage.UnstageBranchDEFINITION(target)

	case *EDITABLEATTS:
		stage.UnstageBranchEDITABLEATTS(target)

	case *EMBEDDEDVALUE:
		stage.UnstageBranchEMBEDDEDVALUE(target)

	case *ENUMVALUE:
		stage.UnstageBranchENUMVALUE(target)

	case *OBJECT:
		stage.UnstageBranchOBJECT(target)

	case *PROPERTIES:
		stage.UnstageBranchPROPERTIES(target)

	case *RELATIONGROUP:
		stage.UnstageBranchRELATIONGROUP(target)

	case *RELATIONGROUPTYPE:
		stage.UnstageBranchRELATIONGROUPTYPE(target)

	case *REQIF:
		stage.UnstageBranchREQIF(target)

	case *REQIFCONTENT:
		stage.UnstageBranchREQIFCONTENT(target)

	case *REQIFHEADER:
		stage.UnstageBranchREQIFHEADER(target)

	case *REQIFTOOLEXTENSION:
		stage.UnstageBranchREQIFTOOLEXTENSION(target)

	case *REQIFTYPE:
		stage.UnstageBranchREQIFTYPE(target)

	case *SOURCE:
		stage.UnstageBranchSOURCE(target)

	case *SOURCESPECIFICATION:
		stage.UnstageBranchSOURCESPECIFICATION(target)

	case *SPECATTRIBUTES:
		stage.UnstageBranchSPECATTRIBUTES(target)

	case *SPECHIERARCHY:
		stage.UnstageBranchSPECHIERARCHY(target)

	case *SPECIFICATION:
		stage.UnstageBranchSPECIFICATION(target)

	case *SPECIFICATIONS:
		stage.UnstageBranchSPECIFICATIONS(target)

	case *SPECIFICATIONTYPE:
		stage.UnstageBranchSPECIFICATIONTYPE(target)

	case *SPECIFIEDVALUES:
		stage.UnstageBranchSPECIFIEDVALUES(target)

	case *SPECOBJECT:
		stage.UnstageBranchSPECOBJECT(target)

	case *SPECOBJECTS:
		stage.UnstageBranchSPECOBJECTS(target)

	case *SPECOBJECTTYPE:
		stage.UnstageBranchSPECOBJECTTYPE(target)

	case *SPECRELATION:
		stage.UnstageBranchSPECRELATION(target)

	case *SPECRELATIONGROUPS:
		stage.UnstageBranchSPECRELATIONGROUPS(target)

	case *SPECRELATIONS:
		stage.UnstageBranchSPECRELATIONS(target)

	case *SPECRELATIONTYPE:
		stage.UnstageBranchSPECRELATIONTYPE(target)

	case *SPECTYPES:
		stage.UnstageBranchSPECTYPES(target)

	case *TARGET:
		stage.UnstageBranchTARGET(target)

	case *TARGETSPECIFICATION:
		stage.UnstageBranchTARGETSPECIFICATION(target)

	case *THEHEADER:
		stage.UnstageBranchTHEHEADER(target)

	case *TOOLEXTENSIONS:
		stage.UnstageBranchTOOLEXTENSIONS(target)

	case *VALUES:
		stage.UnstageBranchVALUES(target)

	case *XHTMLCONTENT:
		stage.UnstageBranchXHTMLCONTENT(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchALTERNATIVEID(alternativeid *ALTERNATIVEID) {

	// check if instance is already staged
	if !IsStaged(stage, alternativeid) {
		return
	}

	alternativeid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitionboolean) {
		return
	}

	attributedefinitionboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionboolean.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitionboolean.ALTERNATIVEID)
	}
	if attributedefinitionboolean.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitionboolean.DEFAULTVALUE)
	}
	if attributedefinitionboolean.TYPE != nil {
		UnstageBranch(stage, attributedefinitionboolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONDATE(attributedefinitiondate *ATTRIBUTEDEFINITIONDATE) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitiondate) {
		return
	}

	attributedefinitiondate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitiondate.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitiondate.ALTERNATIVEID)
	}
	if attributedefinitiondate.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitiondate.DEFAULTVALUE)
	}
	if attributedefinitiondate.TYPE != nil {
		UnstageBranch(stage, attributedefinitiondate.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumeration *ATTRIBUTEDEFINITIONENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitionenumeration) {
		return
	}

	attributedefinitionenumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionenumeration.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitionenumeration.DEFAULTVALUE)
	}
	if attributedefinitionenumeration.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitionenumeration.ALTERNATIVEID)
	}
	if attributedefinitionenumeration.TYPE != nil {
		UnstageBranch(stage, attributedefinitionenumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONINTEGER(attributedefinitioninteger *ATTRIBUTEDEFINITIONINTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitioninteger) {
		return
	}

	attributedefinitioninteger.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitioninteger.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitioninteger.ALTERNATIVEID)
	}
	if attributedefinitioninteger.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitioninteger.DEFAULTVALUE)
	}
	if attributedefinitioninteger.TYPE != nil {
		UnstageBranch(stage, attributedefinitioninteger.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONREAL(attributedefinitionreal *ATTRIBUTEDEFINITIONREAL) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitionreal) {
		return
	}

	attributedefinitionreal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionreal.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitionreal.ALTERNATIVEID)
	}
	if attributedefinitionreal.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitionreal.DEFAULTVALUE)
	}
	if attributedefinitionreal.TYPE != nil {
		UnstageBranch(stage, attributedefinitionreal.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRING) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitionstring) {
		return
	}

	attributedefinitionstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionstring.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitionstring.ALTERNATIVEID)
	}
	if attributedefinitionstring.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitionstring.DEFAULTVALUE)
	}
	if attributedefinitionstring.TYPE != nil {
		UnstageBranch(stage, attributedefinitionstring.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtml *ATTRIBUTEDEFINITIONXHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attributedefinitionxhtml) {
		return
	}

	attributedefinitionxhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributedefinitionxhtml.ALTERNATIVEID != nil {
		UnstageBranch(stage, attributedefinitionxhtml.ALTERNATIVEID)
	}
	if attributedefinitionxhtml.DEFAULTVALUE != nil {
		UnstageBranch(stage, attributedefinitionxhtml.DEFAULTVALUE)
	}
	if attributedefinitionxhtml.TYPE != nil {
		UnstageBranch(stage, attributedefinitionxhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEBOOLEAN(attributevalueboolean *ATTRIBUTEVALUEBOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attributevalueboolean) {
		return
	}

	attributevalueboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueboolean.DEFINITION != nil {
		UnstageBranch(stage, attributevalueboolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEDATE(attributevaluedate *ATTRIBUTEVALUEDATE) {

	// check if instance is already staged
	if !IsStaged(stage, attributevaluedate) {
		return
	}

	attributevaluedate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluedate.DEFINITION != nil {
		UnstageBranch(stage, attributevaluedate.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEENUMERATION(attributevalueenumeration *ATTRIBUTEVALUEENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attributevalueenumeration) {
		return
	}

	attributevalueenumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueenumeration.DEFINITION != nil {
		UnstageBranch(stage, attributevalueenumeration.DEFINITION)
	}
	if attributevalueenumeration.VALUES != nil {
		UnstageBranch(stage, attributevalueenumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEINTEGER(attributevalueinteger *ATTRIBUTEVALUEINTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attributevalueinteger) {
		return
	}

	attributevalueinteger.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevalueinteger.DEFINITION != nil {
		UnstageBranch(stage, attributevalueinteger.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEREAL(attributevaluereal *ATTRIBUTEVALUEREAL) {

	// check if instance is already staged
	if !IsStaged(stage, attributevaluereal) {
		return
	}

	attributevaluereal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluereal.DEFINITION != nil {
		UnstageBranch(stage, attributevaluereal.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUESTRING(attributevaluestring *ATTRIBUTEVALUESTRING) {

	// check if instance is already staged
	if !IsStaged(stage, attributevaluestring) {
		return
	}

	attributevaluestring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluestring.DEFINITION != nil {
		UnstageBranch(stage, attributevaluestring.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTEVALUEXHTML(attributevaluexhtml *ATTRIBUTEVALUEXHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attributevaluexhtml) {
		return
	}

	attributevaluexhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributevaluexhtml.THEVALUE != nil {
		UnstageBranch(stage, attributevaluexhtml.THEVALUE)
	}
	if attributevaluexhtml.THEORIGINALVALUE != nil {
		UnstageBranch(stage, attributevaluexhtml.THEORIGINALVALUE)
	}
	if attributevaluexhtml.DEFINITION != nil {
		UnstageBranch(stage, attributevaluexhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCHILDREN(children *CHILDREN) {

	// check if instance is already staged
	if !IsStaged(stage, children) {
		return
	}

	children.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spechierarchy := range children.SPECHIERARCHY {
		UnstageBranch(stage, _spechierarchy)
	}

}

func (stage *StageStruct) UnstageBranchCORECONTENT(corecontent *CORECONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, corecontent) {
		return
	}

	corecontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if corecontent.REQIFCONTENT != nil {
		UnstageBranch(stage, corecontent.REQIFCONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONBOOLEAN(datatypedefinitionboolean *DATATYPEDEFINITIONBOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitionboolean) {
		return
	}

	datatypedefinitionboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionboolean.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitionboolean.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONDATE(datatypedefinitiondate *DATATYPEDEFINITIONDATE) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitiondate) {
		return
	}

	datatypedefinitiondate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitiondate.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitiondate.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumeration *DATATYPEDEFINITIONENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitionenumeration) {
		return
	}

	datatypedefinitionenumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionenumeration.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitionenumeration.ALTERNATIVEID)
	}
	if datatypedefinitionenumeration.SPECIFIEDVALUES != nil {
		UnstageBranch(stage, datatypedefinitionenumeration.SPECIFIEDVALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONINTEGER(datatypedefinitioninteger *DATATYPEDEFINITIONINTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitioninteger) {
		return
	}

	datatypedefinitioninteger.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitioninteger.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitioninteger.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONREAL(datatypedefinitionreal *DATATYPEDEFINITIONREAL) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitionreal) {
		return
	}

	datatypedefinitionreal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionreal.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitionreal.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONSTRING(datatypedefinitionstring *DATATYPEDEFINITIONSTRING) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitionstring) {
		return
	}

	datatypedefinitionstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionstring.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitionstring.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPEDEFINITIONXHTML(datatypedefinitionxhtml *DATATYPEDEFINITIONXHTML) {

	// check if instance is already staged
	if !IsStaged(stage, datatypedefinitionxhtml) {
		return
	}

	datatypedefinitionxhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatypedefinitionxhtml.ALTERNATIVEID != nil {
		UnstageBranch(stage, datatypedefinitionxhtml.ALTERNATIVEID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPES(datatypes *DATATYPES) {

	// check if instance is already staged
	if !IsStaged(stage, datatypes) {
		return
	}

	datatypes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatypedefinitionboolean := range datatypes.DATATYPEDEFINITIONBOOLEAN {
		UnstageBranch(stage, _datatypedefinitionboolean)
	}
	for _, _datatypedefinitiondate := range datatypes.DATATYPEDEFINITIONDATE {
		UnstageBranch(stage, _datatypedefinitiondate)
	}
	for _, _datatypedefinitionenumeration := range datatypes.DATATYPEDEFINITIONENUMERATION {
		UnstageBranch(stage, _datatypedefinitionenumeration)
	}
	for _, _datatypedefinitioninteger := range datatypes.DATATYPEDEFINITIONINTEGER {
		UnstageBranch(stage, _datatypedefinitioninteger)
	}
	for _, _datatypedefinitionreal := range datatypes.DATATYPEDEFINITIONREAL {
		UnstageBranch(stage, _datatypedefinitionreal)
	}
	for _, _datatypedefinitionstring := range datatypes.DATATYPEDEFINITIONSTRING {
		UnstageBranch(stage, _datatypedefinitionstring)
	}
	for _, _datatypedefinitionxhtml := range datatypes.DATATYPEDEFINITIONXHTML {
		UnstageBranch(stage, _datatypedefinitionxhtml)
	}

}

func (stage *StageStruct) UnstageBranchDEFAULTVALUE(defaultvalue *DEFAULTVALUE) {

	// check if instance is already staged
	if !IsStaged(stage, defaultvalue) {
		return
	}

	defaultvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaultvalue.ATTRIBUTEVALUEBOOLEAN != nil {
		UnstageBranch(stage, defaultvalue.ATTRIBUTEVALUEBOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDEFINITION(definition *DEFINITION) {

	// check if instance is already staged
	if !IsStaged(stage, definition) {
		return
	}

	definition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEDITABLEATTS(editableatts *EDITABLEATTS) {

	// check if instance is already staged
	if !IsStaged(stage, editableatts) {
		return
	}

	editableatts.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEMBEDDEDVALUE(embeddedvalue *EMBEDDEDVALUE) {

	// check if instance is already staged
	if !IsStaged(stage, embeddedvalue) {
		return
	}

	embeddedvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchENUMVALUE(enumvalue *ENUMVALUE) {

	// check if instance is already staged
	if !IsStaged(stage, enumvalue) {
		return
	}

	enumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumvalue.ALTERNATIVEID != nil {
		UnstageBranch(stage, enumvalue.ALTERNATIVEID)
	}
	if enumvalue.PROPERTIES != nil {
		UnstageBranch(stage, enumvalue.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOBJECT(object *OBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, object) {
		return
	}

	object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPROPERTIES(properties *PROPERTIES) {

	// check if instance is already staged
	if !IsStaged(stage, properties) {
		return
	}

	properties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if properties.EMBEDDEDVALUE != nil {
		UnstageBranch(stage, properties.EMBEDDEDVALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRELATIONGROUP(relationgroup *RELATIONGROUP) {

	// check if instance is already staged
	if !IsStaged(stage, relationgroup) {
		return
	}

	relationgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relationgroup.ALTERNATIVEID != nil {
		UnstageBranch(stage, relationgroup.ALTERNATIVEID)
	}
	if relationgroup.SOURCESPECIFICATION != nil {
		UnstageBranch(stage, relationgroup.SOURCESPECIFICATION)
	}
	if relationgroup.SPECRELATIONS != nil {
		UnstageBranch(stage, relationgroup.SPECRELATIONS)
	}
	if relationgroup.TARGETSPECIFICATION != nil {
		UnstageBranch(stage, relationgroup.TARGETSPECIFICATION)
	}
	if relationgroup.TYPE != nil {
		UnstageBranch(stage, relationgroup.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRELATIONGROUPTYPE(relationgrouptype *RELATIONGROUPTYPE) {

	// check if instance is already staged
	if !IsStaged(stage, relationgrouptype) {
		return
	}

	relationgrouptype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relationgrouptype.ALTERNATIVEID != nil {
		UnstageBranch(stage, relationgrouptype.ALTERNATIVEID)
	}
	if relationgrouptype.SPECATTRIBUTES != nil {
		UnstageBranch(stage, relationgrouptype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if !IsStaged(stage, reqif) {
		return
	}

	reqif.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.HEADER != nil {
		UnstageBranch(stage, reqif.HEADER)
	}
	if reqif.CORECONTENT != nil {
		UnstageBranch(stage, reqif.CORECONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIFCONTENT(reqifcontent *REQIFCONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, reqifcontent) {
		return
	}

	reqifcontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqifcontent.DATATYPES != nil {
		UnstageBranch(stage, reqifcontent.DATATYPES)
	}
	if reqifcontent.SPECTYPES != nil {
		UnstageBranch(stage, reqifcontent.SPECTYPES)
	}
	if reqifcontent.SPECOBJECTS != nil {
		UnstageBranch(stage, reqifcontent.SPECOBJECTS)
	}
	if reqifcontent.SPECRELATIONS != nil {
		UnstageBranch(stage, reqifcontent.SPECRELATIONS)
	}
	if reqifcontent.SPECIFICATIONS != nil {
		UnstageBranch(stage, reqifcontent.SPECIFICATIONS)
	}
	if reqifcontent.SPECRELATIONGROUPS != nil {
		UnstageBranch(stage, reqifcontent.SPECRELATIONGROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIFHEADER(reqifheader *REQIFHEADER) {

	// check if instance is already staged
	if !IsStaged(stage, reqifheader) {
		return
	}

	reqifheader.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIFTOOLEXTENSION(reqiftoolextension *REQIFTOOLEXTENSION) {

	// check if instance is already staged
	if !IsStaged(stage, reqiftoolextension) {
		return
	}

	reqiftoolextension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIFTYPE(reqiftype *REQIFTYPE) {

	// check if instance is already staged
	if !IsStaged(stage, reqiftype) {
		return
	}

	reqiftype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSOURCE(source *SOURCE) {

	// check if instance is already staged
	if !IsStaged(stage, source) {
		return
	}

	source.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSOURCESPECIFICATION(sourcespecification *SOURCESPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, sourcespecification) {
		return
	}

	sourcespecification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECATTRIBUTES(specattributes *SPECATTRIBUTES) {

	// check if instance is already staged
	if !IsStaged(stage, specattributes) {
		return
	}

	specattributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributedefinitionboolean := range specattributes.ATTRIBUTEDEFINITIONBOOLEAN {
		UnstageBranch(stage, _attributedefinitionboolean)
	}
	for _, _attributedefinitiondate := range specattributes.ATTRIBUTEDEFINITIONDATE {
		UnstageBranch(stage, _attributedefinitiondate)
	}
	for _, _attributedefinitionenumeration := range specattributes.ATTRIBUTEDEFINITIONENUMERATION {
		UnstageBranch(stage, _attributedefinitionenumeration)
	}
	for _, _attributedefinitioninteger := range specattributes.ATTRIBUTEDEFINITIONINTEGER {
		UnstageBranch(stage, _attributedefinitioninteger)
	}
	for _, _attributedefinitionreal := range specattributes.ATTRIBUTEDEFINITIONREAL {
		UnstageBranch(stage, _attributedefinitionreal)
	}
	for _, _attributedefinitionstring := range specattributes.ATTRIBUTEDEFINITIONSTRING {
		UnstageBranch(stage, _attributedefinitionstring)
	}
	for _, _attributedefinitionxhtml := range specattributes.ATTRIBUTEDEFINITIONXHTML {
		UnstageBranch(stage, _attributedefinitionxhtml)
	}

}

func (stage *StageStruct) UnstageBranchSPECHIERARCHY(spechierarchy *SPECHIERARCHY) {

	// check if instance is already staged
	if !IsStaged(stage, spechierarchy) {
		return
	}

	spechierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spechierarchy.ALTERNATIVEID != nil {
		UnstageBranch(stage, spechierarchy.ALTERNATIVEID)
	}
	if spechierarchy.CHILDREN != nil {
		UnstageBranch(stage, spechierarchy.CHILDREN)
	}
	if spechierarchy.EDITABLEATTS != nil {
		UnstageBranch(stage, spechierarchy.EDITABLEATTS)
	}
	if spechierarchy.OBJECT != nil {
		UnstageBranch(stage, spechierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, specification) {
		return
	}

	specification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVEID != nil {
		UnstageBranch(stage, specification.ALTERNATIVEID)
	}
	if specification.VALUES != nil {
		UnstageBranch(stage, specification.VALUES)
	}
	if specification.CHILDREN != nil {
		UnstageBranch(stage, specification.CHILDREN)
	}
	if specification.TYPE != nil {
		UnstageBranch(stage, specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFICATIONS(specifications *SPECIFICATIONS) {

	// check if instance is already staged
	if !IsStaged(stage, specifications) {
		return
	}

	specifications.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range specifications.SPECIFICATION {
		UnstageBranch(stage, _specification)
	}

}

func (stage *StageStruct) UnstageBranchSPECIFICATIONTYPE(specificationtype *SPECIFICATIONTYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specificationtype) {
		return
	}

	specificationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specificationtype.ALTERNATIVEID != nil {
		UnstageBranch(stage, specificationtype.ALTERNATIVEID)
	}
	if specificationtype.SPECATTRIBUTES != nil {
		UnstageBranch(stage, specificationtype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFIEDVALUES(specifiedvalues *SPECIFIEDVALUES) {

	// check if instance is already staged
	if !IsStaged(stage, specifiedvalues) {
		return
	}

	specifiedvalues.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumvalue := range specifiedvalues.ENUMVALUE {
		UnstageBranch(stage, _enumvalue)
	}

}

func (stage *StageStruct) UnstageBranchSPECOBJECT(specobject *SPECOBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, specobject) {
		return
	}

	specobject.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specobject.ALTERNATIVEID != nil {
		UnstageBranch(stage, specobject.ALTERNATIVEID)
	}
	if specobject.VALUES != nil {
		UnstageBranch(stage, specobject.VALUES)
	}
	if specobject.TYPE != nil {
		UnstageBranch(stage, specobject.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECOBJECTS(specobjects *SPECOBJECTS) {

	// check if instance is already staged
	if !IsStaged(stage, specobjects) {
		return
	}

	specobjects.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specobject := range specobjects.SPECOBJECT {
		UnstageBranch(stage, _specobject)
	}

}

func (stage *StageStruct) UnstageBranchSPECOBJECTTYPE(specobjecttype *SPECOBJECTTYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specobjecttype) {
		return
	}

	specobjecttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specobjecttype.ALTERNATIVEID != nil {
		UnstageBranch(stage, specobjecttype.ALTERNATIVEID)
	}
	if specobjecttype.SPECATTRIBUTES != nil {
		UnstageBranch(stage, specobjecttype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECRELATION(specrelation *SPECRELATION) {

	// check if instance is already staged
	if !IsStaged(stage, specrelation) {
		return
	}

	specrelation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specrelation.ALTERNATIVEID != nil {
		UnstageBranch(stage, specrelation.ALTERNATIVEID)
	}
	if specrelation.VALUES != nil {
		UnstageBranch(stage, specrelation.VALUES)
	}
	if specrelation.SOURCE != nil {
		UnstageBranch(stage, specrelation.SOURCE)
	}
	if specrelation.TARGET != nil {
		UnstageBranch(stage, specrelation.TARGET)
	}
	if specrelation.TYPE != nil {
		UnstageBranch(stage, specrelation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECRELATIONGROUPS(specrelationgroups *SPECRELATIONGROUPS) {

	// check if instance is already staged
	if !IsStaged(stage, specrelationgroups) {
		return
	}

	specrelationgroups.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgroup := range specrelationgroups.RELATIONGROUP {
		UnstageBranch(stage, _relationgroup)
	}

}

func (stage *StageStruct) UnstageBranchSPECRELATIONS(specrelations *SPECRELATIONS) {

	// check if instance is already staged
	if !IsStaged(stage, specrelations) {
		return
	}

	specrelations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECRELATIONTYPE(specrelationtype *SPECRELATIONTYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specrelationtype) {
		return
	}

	specrelationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specrelationtype.ALTERNATIVEID != nil {
		UnstageBranch(stage, specrelationtype.ALTERNATIVEID)
	}
	if specrelationtype.SPECATTRIBUTES != nil {
		UnstageBranch(stage, specrelationtype.SPECATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECTYPES(spectypes *SPECTYPES) {

	// check if instance is already staged
	if !IsStaged(stage, spectypes) {
		return
	}

	spectypes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relationgrouptype := range spectypes.RELATIONGROUPTYPE {
		UnstageBranch(stage, _relationgrouptype)
	}
	for _, _specobjecttype := range spectypes.SPECOBJECTTYPE {
		UnstageBranch(stage, _specobjecttype)
	}
	for _, _specrelationtype := range spectypes.SPECRELATIONTYPE {
		UnstageBranch(stage, _specrelationtype)
	}
	for _, _specificationtype := range spectypes.SPECIFICATIONTYPE {
		UnstageBranch(stage, _specificationtype)
	}

}

func (stage *StageStruct) UnstageBranchTARGET(target *TARGET) {

	// check if instance is already staged
	if !IsStaged(stage, target) {
		return
	}

	target.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTARGETSPECIFICATION(targetspecification *TARGETSPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, targetspecification) {
		return
	}

	targetspecification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTHEHEADER(theheader *THEHEADER) {

	// check if instance is already staged
	if !IsStaged(stage, theheader) {
		return
	}

	theheader.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if theheader.REQIFHEADER != nil {
		UnstageBranch(stage, theheader.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTOOLEXTENSIONS(toolextensions *TOOLEXTENSIONS) {

	// check if instance is already staged
	if !IsStaged(stage, toolextensions) {
		return
	}

	toolextensions.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _reqiftoolextension := range toolextensions.REQIFTOOLEXTENSION {
		UnstageBranch(stage, _reqiftoolextension)
	}

}

func (stage *StageStruct) UnstageBranchVALUES(values *VALUES) {

	// check if instance is already staged
	if !IsStaged(stage, values) {
		return
	}

	values.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXHTMLCONTENT(xhtmlcontent *XHTMLCONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, xhtmlcontent) {
		return
	}

	xhtmlcontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
