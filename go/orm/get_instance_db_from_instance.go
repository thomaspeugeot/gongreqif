// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/gongreqif/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | ALTERNATIVEIDDB | ATTRIBUTEDEFINITIONBOOLEANDB | ATTRIBUTEDEFINITIONDATEDB | ATTRIBUTEDEFINITIONENUMERATIONDB | ATTRIBUTEDEFINITIONINTEGERDB | ATTRIBUTEDEFINITIONREALDB | ATTRIBUTEDEFINITIONSTRINGDB | ATTRIBUTEDEFINITIONXHTMLDB | ATTRIBUTEVALUEBOOLEANDB | ATTRIBUTEVALUEDATEDB | ATTRIBUTEVALUEENUMERATIONDB | ATTRIBUTEVALUEINTEGERDB | ATTRIBUTEVALUEREALDB | ATTRIBUTEVALUESTRINGDB | ATTRIBUTEVALUEXHTMLDB | CHILDRENDB | CORECONTENTDB | DATATYPEDEFINITIONBOOLEANDB | DATATYPEDEFINITIONDATEDB | DATATYPEDEFINITIONENUMERATIONDB | DATATYPEDEFINITIONINTEGERDB | DATATYPEDEFINITIONREALDB | DATATYPEDEFINITIONSTRINGDB | DATATYPEDEFINITIONXHTMLDB | DATATYPESDB | DEFAULTVALUEDB | DEFINITIONDB | EDITABLEATTSDB | EMBEDDEDVALUEDB | ENUMVALUEDB | OBJECTDB | PROPERTIESDB | RELATIONGROUPDB | RELATIONGROUPTYPEDB | REQIFDB | REQIFCONTENTDB | REQIFHEADERDB | REQIFTOOLEXTENSIONDB | REQIFTYPEDB | SOURCEDB | SOURCESPECIFICATIONDB | SPECATTRIBUTESDB | SPECHIERARCHYDB | SPECIFICATIONDB | SPECIFICATIONSDB | SPECIFICATIONTYPEDB | SPECIFIEDVALUESDB | SPECOBJECTDB | SPECOBJECTSDB | SPECOBJECTTYPEDB | SPECRELATIONDB | SPECRELATIONGROUPSDB | SPECRELATIONSDB | SPECRELATIONTYPEDB | SPECTYPESDB | TARGETDB | TARGETSPECIFICATIONDB | THEHEADERDB | TOOLEXTENSIONSDB | VALUESDB | XHTMLCONTENTDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVEID:
		alternativeidInstance := any(concreteInstance).(*models.ALTERNATIVEID)
		ret2 := backRepo.BackRepoALTERNATIVEID.GetALTERNATIVEIDDBFromALTERNATIVEIDPtr(alternativeidInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		attributedefinitionbooleanInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONBOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetATTRIBUTEDEFINITIONBOOLEANDBFromATTRIBUTEDEFINITIONBOOLEANPtr(attributedefinitionbooleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONDATE:
		attributedefinitiondateInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONDATE)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetATTRIBUTEDEFINITIONDATEDBFromATTRIBUTEDEFINITIONDATEPtr(attributedefinitiondateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		attributedefinitionenumerationInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetATTRIBUTEDEFINITIONENUMERATIONDBFromATTRIBUTEDEFINITIONENUMERATIONPtr(attributedefinitionenumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		attributedefinitionintegerInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONINTEGER)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetATTRIBUTEDEFINITIONINTEGERDBFromATTRIBUTEDEFINITIONINTEGERPtr(attributedefinitionintegerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONREAL:
		attributedefinitionrealInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONREAL)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetATTRIBUTEDEFINITIONREALDBFromATTRIBUTEDEFINITIONREALPtr(attributedefinitionrealInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		attributedefinitionstringInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONSTRING)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetATTRIBUTEDEFINITIONSTRINGDBFromATTRIBUTEDEFINITIONSTRINGPtr(attributedefinitionstringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		attributedefinitionxhtmlInstance := any(concreteInstance).(*models.ATTRIBUTEDEFINITIONXHTML)
		ret2 := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetATTRIBUTEDEFINITIONXHTMLDBFromATTRIBUTEDEFINITIONXHTMLPtr(attributedefinitionxhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		attributevaluebooleanInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEBOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetATTRIBUTEVALUEBOOLEANDBFromATTRIBUTEVALUEBOOLEANPtr(attributevaluebooleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEDATE:
		attributevaluedateInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEDATE)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEDATE.GetATTRIBUTEVALUEDATEDBFromATTRIBUTEVALUEDATEPtr(attributevaluedateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEENUMERATION:
		attributevalueenumerationInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetATTRIBUTEVALUEENUMERATIONDBFromATTRIBUTEVALUEENUMERATIONPtr(attributevalueenumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEINTEGER:
		attributevalueintegerInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEINTEGER)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetATTRIBUTEVALUEINTEGERDBFromATTRIBUTEVALUEINTEGERPtr(attributevalueintegerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEREAL:
		attributevaluerealInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEREAL)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEREAL.GetATTRIBUTEVALUEREALDBFromATTRIBUTEVALUEREALPtr(attributevaluerealInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUESTRING:
		attributevaluestringInstance := any(concreteInstance).(*models.ATTRIBUTEVALUESTRING)
		ret2 := backRepo.BackRepoATTRIBUTEVALUESTRING.GetATTRIBUTEVALUESTRINGDBFromATTRIBUTEVALUESTRINGPtr(attributevaluestringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTEVALUEXHTML:
		attributevaluexhtmlInstance := any(concreteInstance).(*models.ATTRIBUTEVALUEXHTML)
		ret2 := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetATTRIBUTEVALUEXHTMLDBFromATTRIBUTEVALUEXHTMLPtr(attributevaluexhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.CHILDREN:
		childrenInstance := any(concreteInstance).(*models.CHILDREN)
		ret2 := backRepo.BackRepoCHILDREN.GetCHILDRENDBFromCHILDRENPtr(childrenInstance)
		ret = any(ret2).(*T2)
	case *models.CORECONTENT:
		corecontentInstance := any(concreteInstance).(*models.CORECONTENT)
		ret2 := backRepo.BackRepoCORECONTENT.GetCORECONTENTDBFromCORECONTENTPtr(corecontentInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		datatypedefinitionbooleanInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONBOOLEAN)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDATATYPEDEFINITIONBOOLEANDBFromDATATYPEDEFINITIONBOOLEANPtr(datatypedefinitionbooleanInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONDATE:
		datatypedefinitiondateInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONDATE)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDATATYPEDEFINITIONDATEDBFromDATATYPEDEFINITIONDATEPtr(datatypedefinitiondateInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONENUMERATION:
		datatypedefinitionenumerationInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONENUMERATION)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDATATYPEDEFINITIONENUMERATIONDBFromDATATYPEDEFINITIONENUMERATIONPtr(datatypedefinitionenumerationInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONINTEGER:
		datatypedefinitionintegerInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONINTEGER)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDATATYPEDEFINITIONINTEGERDBFromDATATYPEDEFINITIONINTEGERPtr(datatypedefinitionintegerInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONREAL:
		datatypedefinitionrealInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONREAL)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDATATYPEDEFINITIONREALDBFromDATATYPEDEFINITIONREALPtr(datatypedefinitionrealInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONSTRING:
		datatypedefinitionstringInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONSTRING)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDATATYPEDEFINITIONSTRINGDBFromDATATYPEDEFINITIONSTRINGPtr(datatypedefinitionstringInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPEDEFINITIONXHTML:
		datatypedefinitionxhtmlInstance := any(concreteInstance).(*models.DATATYPEDEFINITIONXHTML)
		ret2 := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDATATYPEDEFINITIONXHTMLDBFromDATATYPEDEFINITIONXHTMLPtr(datatypedefinitionxhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPES:
		datatypesInstance := any(concreteInstance).(*models.DATATYPES)
		ret2 := backRepo.BackRepoDATATYPES.GetDATATYPESDBFromDATATYPESPtr(datatypesInstance)
		ret = any(ret2).(*T2)
	case *models.DEFAULTVALUE:
		defaultvalueInstance := any(concreteInstance).(*models.DEFAULTVALUE)
		ret2 := backRepo.BackRepoDEFAULTVALUE.GetDEFAULTVALUEDBFromDEFAULTVALUEPtr(defaultvalueInstance)
		ret = any(ret2).(*T2)
	case *models.DEFINITION:
		definitionInstance := any(concreteInstance).(*models.DEFINITION)
		ret2 := backRepo.BackRepoDEFINITION.GetDEFINITIONDBFromDEFINITIONPtr(definitionInstance)
		ret = any(ret2).(*T2)
	case *models.EDITABLEATTS:
		editableattsInstance := any(concreteInstance).(*models.EDITABLEATTS)
		ret2 := backRepo.BackRepoEDITABLEATTS.GetEDITABLEATTSDBFromEDITABLEATTSPtr(editableattsInstance)
		ret = any(ret2).(*T2)
	case *models.EMBEDDEDVALUE:
		embeddedvalueInstance := any(concreteInstance).(*models.EMBEDDEDVALUE)
		ret2 := backRepo.BackRepoEMBEDDEDVALUE.GetEMBEDDEDVALUEDBFromEMBEDDEDVALUEPtr(embeddedvalueInstance)
		ret = any(ret2).(*T2)
	case *models.ENUMVALUE:
		enumvalueInstance := any(concreteInstance).(*models.ENUMVALUE)
		ret2 := backRepo.BackRepoENUMVALUE.GetENUMVALUEDBFromENUMVALUEPtr(enumvalueInstance)
		ret = any(ret2).(*T2)
	case *models.OBJECT:
		objectInstance := any(concreteInstance).(*models.OBJECT)
		ret2 := backRepo.BackRepoOBJECT.GetOBJECTDBFromOBJECTPtr(objectInstance)
		ret = any(ret2).(*T2)
	case *models.PROPERTIES:
		propertiesInstance := any(concreteInstance).(*models.PROPERTIES)
		ret2 := backRepo.BackRepoPROPERTIES.GetPROPERTIESDBFromPROPERTIESPtr(propertiesInstance)
		ret = any(ret2).(*T2)
	case *models.RELATIONGROUP:
		relationgroupInstance := any(concreteInstance).(*models.RELATIONGROUP)
		ret2 := backRepo.BackRepoRELATIONGROUP.GetRELATIONGROUPDBFromRELATIONGROUPPtr(relationgroupInstance)
		ret = any(ret2).(*T2)
	case *models.RELATIONGROUPTYPE:
		relationgrouptypeInstance := any(concreteInstance).(*models.RELATIONGROUPTYPE)
		ret2 := backRepo.BackRepoRELATIONGROUPTYPE.GetRELATIONGROUPTYPEDBFromRELATIONGROUPTYPEPtr(relationgrouptypeInstance)
		ret = any(ret2).(*T2)
	case *models.REQIF:
		reqifInstance := any(concreteInstance).(*models.REQIF)
		ret2 := backRepo.BackRepoREQIF.GetREQIFDBFromREQIFPtr(reqifInstance)
		ret = any(ret2).(*T2)
	case *models.REQIFCONTENT:
		reqifcontentInstance := any(concreteInstance).(*models.REQIFCONTENT)
		ret2 := backRepo.BackRepoREQIFCONTENT.GetREQIFCONTENTDBFromREQIFCONTENTPtr(reqifcontentInstance)
		ret = any(ret2).(*T2)
	case *models.REQIFHEADER:
		reqifheaderInstance := any(concreteInstance).(*models.REQIFHEADER)
		ret2 := backRepo.BackRepoREQIFHEADER.GetREQIFHEADERDBFromREQIFHEADERPtr(reqifheaderInstance)
		ret = any(ret2).(*T2)
	case *models.REQIFTOOLEXTENSION:
		reqiftoolextensionInstance := any(concreteInstance).(*models.REQIFTOOLEXTENSION)
		ret2 := backRepo.BackRepoREQIFTOOLEXTENSION.GetREQIFTOOLEXTENSIONDBFromREQIFTOOLEXTENSIONPtr(reqiftoolextensionInstance)
		ret = any(ret2).(*T2)
	case *models.REQIFTYPE:
		reqiftypeInstance := any(concreteInstance).(*models.REQIFTYPE)
		ret2 := backRepo.BackRepoREQIFTYPE.GetREQIFTYPEDBFromREQIFTYPEPtr(reqiftypeInstance)
		ret = any(ret2).(*T2)
	case *models.SOURCE:
		sourceInstance := any(concreteInstance).(*models.SOURCE)
		ret2 := backRepo.BackRepoSOURCE.GetSOURCEDBFromSOURCEPtr(sourceInstance)
		ret = any(ret2).(*T2)
	case *models.SOURCESPECIFICATION:
		sourcespecificationInstance := any(concreteInstance).(*models.SOURCESPECIFICATION)
		ret2 := backRepo.BackRepoSOURCESPECIFICATION.GetSOURCESPECIFICATIONDBFromSOURCESPECIFICATIONPtr(sourcespecificationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECATTRIBUTES:
		specattributesInstance := any(concreteInstance).(*models.SPECATTRIBUTES)
		ret2 := backRepo.BackRepoSPECATTRIBUTES.GetSPECATTRIBUTESDBFromSPECATTRIBUTESPtr(specattributesInstance)
		ret = any(ret2).(*T2)
	case *models.SPECHIERARCHY:
		spechierarchyInstance := any(concreteInstance).(*models.SPECHIERARCHY)
		ret2 := backRepo.BackRepoSPECHIERARCHY.GetSPECHIERARCHYDBFromSPECHIERARCHYPtr(spechierarchyInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATION:
		specificationInstance := any(concreteInstance).(*models.SPECIFICATION)
		ret2 := backRepo.BackRepoSPECIFICATION.GetSPECIFICATIONDBFromSPECIFICATIONPtr(specificationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATIONS:
		specificationsInstance := any(concreteInstance).(*models.SPECIFICATIONS)
		ret2 := backRepo.BackRepoSPECIFICATIONS.GetSPECIFICATIONSDBFromSPECIFICATIONSPtr(specificationsInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATIONTYPE:
		specificationtypeInstance := any(concreteInstance).(*models.SPECIFICATIONTYPE)
		ret2 := backRepo.BackRepoSPECIFICATIONTYPE.GetSPECIFICATIONTYPEDBFromSPECIFICATIONTYPEPtr(specificationtypeInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFIEDVALUES:
		specifiedvaluesInstance := any(concreteInstance).(*models.SPECIFIEDVALUES)
		ret2 := backRepo.BackRepoSPECIFIEDVALUES.GetSPECIFIEDVALUESDBFromSPECIFIEDVALUESPtr(specifiedvaluesInstance)
		ret = any(ret2).(*T2)
	case *models.SPECOBJECT:
		specobjectInstance := any(concreteInstance).(*models.SPECOBJECT)
		ret2 := backRepo.BackRepoSPECOBJECT.GetSPECOBJECTDBFromSPECOBJECTPtr(specobjectInstance)
		ret = any(ret2).(*T2)
	case *models.SPECOBJECTS:
		specobjectsInstance := any(concreteInstance).(*models.SPECOBJECTS)
		ret2 := backRepo.BackRepoSPECOBJECTS.GetSPECOBJECTSDBFromSPECOBJECTSPtr(specobjectsInstance)
		ret = any(ret2).(*T2)
	case *models.SPECOBJECTTYPE:
		specobjecttypeInstance := any(concreteInstance).(*models.SPECOBJECTTYPE)
		ret2 := backRepo.BackRepoSPECOBJECTTYPE.GetSPECOBJECTTYPEDBFromSPECOBJECTTYPEPtr(specobjecttypeInstance)
		ret = any(ret2).(*T2)
	case *models.SPECRELATION:
		specrelationInstance := any(concreteInstance).(*models.SPECRELATION)
		ret2 := backRepo.BackRepoSPECRELATION.GetSPECRELATIONDBFromSPECRELATIONPtr(specrelationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECRELATIONGROUPS:
		specrelationgroupsInstance := any(concreteInstance).(*models.SPECRELATIONGROUPS)
		ret2 := backRepo.BackRepoSPECRELATIONGROUPS.GetSPECRELATIONGROUPSDBFromSPECRELATIONGROUPSPtr(specrelationgroupsInstance)
		ret = any(ret2).(*T2)
	case *models.SPECRELATIONS:
		specrelationsInstance := any(concreteInstance).(*models.SPECRELATIONS)
		ret2 := backRepo.BackRepoSPECRELATIONS.GetSPECRELATIONSDBFromSPECRELATIONSPtr(specrelationsInstance)
		ret = any(ret2).(*T2)
	case *models.SPECRELATIONTYPE:
		specrelationtypeInstance := any(concreteInstance).(*models.SPECRELATIONTYPE)
		ret2 := backRepo.BackRepoSPECRELATIONTYPE.GetSPECRELATIONTYPEDBFromSPECRELATIONTYPEPtr(specrelationtypeInstance)
		ret = any(ret2).(*T2)
	case *models.SPECTYPES:
		spectypesInstance := any(concreteInstance).(*models.SPECTYPES)
		ret2 := backRepo.BackRepoSPECTYPES.GetSPECTYPESDBFromSPECTYPESPtr(spectypesInstance)
		ret = any(ret2).(*T2)
	case *models.TARGET:
		targetInstance := any(concreteInstance).(*models.TARGET)
		ret2 := backRepo.BackRepoTARGET.GetTARGETDBFromTARGETPtr(targetInstance)
		ret = any(ret2).(*T2)
	case *models.TARGETSPECIFICATION:
		targetspecificationInstance := any(concreteInstance).(*models.TARGETSPECIFICATION)
		ret2 := backRepo.BackRepoTARGETSPECIFICATION.GetTARGETSPECIFICATIONDBFromTARGETSPECIFICATIONPtr(targetspecificationInstance)
		ret = any(ret2).(*T2)
	case *models.THEHEADER:
		theheaderInstance := any(concreteInstance).(*models.THEHEADER)
		ret2 := backRepo.BackRepoTHEHEADER.GetTHEHEADERDBFromTHEHEADERPtr(theheaderInstance)
		ret = any(ret2).(*T2)
	case *models.TOOLEXTENSIONS:
		toolextensionsInstance := any(concreteInstance).(*models.TOOLEXTENSIONS)
		ret2 := backRepo.BackRepoTOOLEXTENSIONS.GetTOOLEXTENSIONSDBFromTOOLEXTENSIONSPtr(toolextensionsInstance)
		ret = any(ret2).(*T2)
	case *models.VALUES:
		valuesInstance := any(concreteInstance).(*models.VALUES)
		ret2 := backRepo.BackRepoVALUES.GetVALUESDBFromVALUESPtr(valuesInstance)
		ret = any(ret2).(*T2)
	case *models.XHTMLCONTENT:
		xhtmlcontentInstance := any(concreteInstance).(*models.XHTMLCONTENT)
		ret2 := backRepo.BackRepoXHTMLCONTENT.GetXHTMLCONTENTDBFromXHTMLCONTENTPtr(xhtmlcontentInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVEID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVEID, ALTERNATIVEIDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONBOOLEAN, ATTRIBUTEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONDATE, ATTRIBUTEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONENUMERATION, ATTRIBUTEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONINTEGER, ATTRIBUTEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONREAL, ATTRIBUTEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONSTRING, ATTRIBUTEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONXHTML, ATTRIBUTEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEBOOLEAN, ATTRIBUTEVALUEBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEDATE, ATTRIBUTEVALUEDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEENUMERATION, ATTRIBUTEVALUEENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEINTEGER, ATTRIBUTEVALUEINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEREAL, ATTRIBUTEVALUEREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUESTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUESTRING, ATTRIBUTEVALUESTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEXHTML, ATTRIBUTEVALUEXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CHILDREN:
		tmp := GetInstanceDBFromInstance[models.CHILDREN, CHILDRENDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CORECONTENT:
		tmp := GetInstanceDBFromInstance[models.CORECONTENT, CORECONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONBOOLEAN, DATATYPEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONDATE, DATATYPEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONENUMERATION, DATATYPEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONINTEGER, DATATYPEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONREAL, DATATYPEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONSTRING, DATATYPEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONXHTML, DATATYPEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPES:
		tmp := GetInstanceDBFromInstance[models.DATATYPES, DATATYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DEFAULTVALUE:
		tmp := GetInstanceDBFromInstance[models.DEFAULTVALUE, DEFAULTVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DEFINITION:
		tmp := GetInstanceDBFromInstance[models.DEFINITION, DEFINITIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EDITABLEATTS:
		tmp := GetInstanceDBFromInstance[models.EDITABLEATTS, EDITABLEATTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDEDVALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDEDVALUE, EMBEDDEDVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUMVALUE:
		tmp := GetInstanceDBFromInstance[models.ENUMVALUE, ENUMVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.OBJECT:
		tmp := GetInstanceDBFromInstance[models.OBJECT, OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.PROPERTIES, PROPERTIESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATIONGROUP:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUP, RELATIONGROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATIONGROUPTYPE:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUPTYPE, RELATIONGROUPTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFCONTENT:
		tmp := GetInstanceDBFromInstance[models.REQIFCONTENT, REQIFCONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFTOOLEXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQIFTOOLEXTENSION, REQIFTOOLEXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFTYPE:
		tmp := GetInstanceDBFromInstance[models.REQIFTYPE, REQIFTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SOURCE:
		tmp := GetInstanceDBFromInstance[models.SOURCE, SOURCEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SOURCESPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SOURCESPECIFICATION, SOURCESPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.SPECATTRIBUTES, SPECATTRIBUTESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECHIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPECHIERARCHY, SPECHIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONS, SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONTYPE, SPECIFICATIONTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFIEDVALUES:
		tmp := GetInstanceDBFromInstance[models.SPECIFIEDVALUES, SPECIFIEDVALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECT:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECT, SPECOBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECTS:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTS, SPECOBJECTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECTTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTTYPE, SPECOBJECTTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATION:
		tmp := GetInstanceDBFromInstance[models.SPECRELATION, SPECRELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONGROUPS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONGROUPS, SPECRELATIONGROUPSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONS, SPECRELATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONTYPE, SPECRELATIONTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECTYPES:
		tmp := GetInstanceDBFromInstance[models.SPECTYPES, SPECTYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TARGET:
		tmp := GetInstanceDBFromInstance[models.TARGET, TARGETDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TARGETSPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.TARGETSPECIFICATION, TARGETSPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.THEHEADER:
		tmp := GetInstanceDBFromInstance[models.THEHEADER, THEHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TOOLEXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.TOOLEXTENSIONS, TOOLEXTENSIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VALUES:
		tmp := GetInstanceDBFromInstance[models.VALUES, VALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTMLCONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTMLCONTENT, XHTMLCONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVEID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVEID, ALTERNATIVEIDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONBOOLEAN, ATTRIBUTEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONDATE, ATTRIBUTEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONENUMERATION, ATTRIBUTEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONINTEGER, ATTRIBUTEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONREAL, ATTRIBUTEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONSTRING, ATTRIBUTEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONXHTML, ATTRIBUTEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEBOOLEAN, ATTRIBUTEVALUEBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEDATE, ATTRIBUTEVALUEDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEENUMERATION, ATTRIBUTEVALUEENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEINTEGER, ATTRIBUTEVALUEINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEREAL, ATTRIBUTEVALUEREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUESTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUESTRING, ATTRIBUTEVALUESTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTEVALUEXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEXHTML, ATTRIBUTEVALUEXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CHILDREN:
		tmp := GetInstanceDBFromInstance[models.CHILDREN, CHILDRENDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.CORECONTENT:
		tmp := GetInstanceDBFromInstance[models.CORECONTENT, CORECONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONBOOLEAN, DATATYPEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONDATE, DATATYPEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONENUMERATION, DATATYPEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONINTEGER, DATATYPEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONREAL, DATATYPEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONSTRING, DATATYPEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONXHTML, DATATYPEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPES:
		tmp := GetInstanceDBFromInstance[models.DATATYPES, DATATYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DEFAULTVALUE:
		tmp := GetInstanceDBFromInstance[models.DEFAULTVALUE, DEFAULTVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DEFINITION:
		tmp := GetInstanceDBFromInstance[models.DEFINITION, DEFINITIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EDITABLEATTS:
		tmp := GetInstanceDBFromInstance[models.EDITABLEATTS, EDITABLEATTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDEDVALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDEDVALUE, EMBEDDEDVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUMVALUE:
		tmp := GetInstanceDBFromInstance[models.ENUMVALUE, ENUMVALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.OBJECT:
		tmp := GetInstanceDBFromInstance[models.OBJECT, OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.PROPERTIES, PROPERTIESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATIONGROUP:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUP, RELATIONGROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATIONGROUPTYPE:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUPTYPE, RELATIONGROUPTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFCONTENT:
		tmp := GetInstanceDBFromInstance[models.REQIFCONTENT, REQIFCONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFTOOLEXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQIFTOOLEXTENSION, REQIFTOOLEXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFTYPE:
		tmp := GetInstanceDBFromInstance[models.REQIFTYPE, REQIFTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SOURCE:
		tmp := GetInstanceDBFromInstance[models.SOURCE, SOURCEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SOURCESPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SOURCESPECIFICATION, SOURCESPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.SPECATTRIBUTES, SPECATTRIBUTESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECHIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPECHIERARCHY, SPECHIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONS, SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONTYPE, SPECIFICATIONTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFIEDVALUES:
		tmp := GetInstanceDBFromInstance[models.SPECIFIEDVALUES, SPECIFIEDVALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECT:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECT, SPECOBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECTS:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTS, SPECOBJECTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECOBJECTTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTTYPE, SPECOBJECTTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATION:
		tmp := GetInstanceDBFromInstance[models.SPECRELATION, SPECRELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONGROUPS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONGROUPS, SPECRELATIONGROUPSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONS, SPECRELATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECRELATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONTYPE, SPECRELATIONTYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECTYPES:
		tmp := GetInstanceDBFromInstance[models.SPECTYPES, SPECTYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TARGET:
		tmp := GetInstanceDBFromInstance[models.TARGET, TARGETDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TARGETSPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.TARGETSPECIFICATION, TARGETSPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.THEHEADER:
		tmp := GetInstanceDBFromInstance[models.THEHEADER, THEHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TOOLEXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.TOOLEXTENSIONS, TOOLEXTENSIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.VALUES:
		tmp := GetInstanceDBFromInstance[models.VALUES, VALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTMLCONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTMLCONTENT, XHTMLCONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
