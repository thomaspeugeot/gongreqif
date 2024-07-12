// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ALTERNATIVEIDAPIs []*ALTERNATIVEIDAPI

	ATTRIBUTEDEFINITIONBOOLEANAPIs []*ATTRIBUTEDEFINITIONBOOLEANAPI

	ATTRIBUTEDEFINITIONDATEAPIs []*ATTRIBUTEDEFINITIONDATEAPI

	ATTRIBUTEDEFINITIONENUMERATIONAPIs []*ATTRIBUTEDEFINITIONENUMERATIONAPI

	ATTRIBUTEDEFINITIONINTEGERAPIs []*ATTRIBUTEDEFINITIONINTEGERAPI

	ATTRIBUTEDEFINITIONREALAPIs []*ATTRIBUTEDEFINITIONREALAPI

	ATTRIBUTEDEFINITIONSTRINGAPIs []*ATTRIBUTEDEFINITIONSTRINGAPI

	ATTRIBUTEDEFINITIONXHTMLAPIs []*ATTRIBUTEDEFINITIONXHTMLAPI

	ATTRIBUTEVALUEBOOLEANAPIs []*ATTRIBUTEVALUEBOOLEANAPI

	ATTRIBUTEVALUEDATEAPIs []*ATTRIBUTEVALUEDATEAPI

	ATTRIBUTEVALUEENUMERATIONAPIs []*ATTRIBUTEVALUEENUMERATIONAPI

	ATTRIBUTEVALUEINTEGERAPIs []*ATTRIBUTEVALUEINTEGERAPI

	ATTRIBUTEVALUEREALAPIs []*ATTRIBUTEVALUEREALAPI

	ATTRIBUTEVALUESTRINGAPIs []*ATTRIBUTEVALUESTRINGAPI

	ATTRIBUTEVALUEXHTMLAPIs []*ATTRIBUTEVALUEXHTMLAPI

	CHILDRENAPIs []*CHILDRENAPI

	CORECONTENTAPIs []*CORECONTENTAPI

	DATATYPEDEFINITIONBOOLEANAPIs []*DATATYPEDEFINITIONBOOLEANAPI

	DATATYPEDEFINITIONDATEAPIs []*DATATYPEDEFINITIONDATEAPI

	DATATYPEDEFINITIONENUMERATIONAPIs []*DATATYPEDEFINITIONENUMERATIONAPI

	DATATYPEDEFINITIONINTEGERAPIs []*DATATYPEDEFINITIONINTEGERAPI

	DATATYPEDEFINITIONREALAPIs []*DATATYPEDEFINITIONREALAPI

	DATATYPEDEFINITIONSTRINGAPIs []*DATATYPEDEFINITIONSTRINGAPI

	DATATYPEDEFINITIONXHTMLAPIs []*DATATYPEDEFINITIONXHTMLAPI

	DATATYPESAPIs []*DATATYPESAPI

	DEFAULTVALUEAPIs []*DEFAULTVALUEAPI

	DEFINITIONAPIs []*DEFINITIONAPI

	EDITABLEATTSAPIs []*EDITABLEATTSAPI

	EMBEDDEDVALUEAPIs []*EMBEDDEDVALUEAPI

	ENUMVALUEAPIs []*ENUMVALUEAPI

	OBJECTAPIs []*OBJECTAPI

	PROPERTIESAPIs []*PROPERTIESAPI

	RELATIONGROUPAPIs []*RELATIONGROUPAPI

	RELATIONGROUPTYPEAPIs []*RELATIONGROUPTYPEAPI

	REQIFAPIs []*REQIFAPI

	REQIFCONTENTAPIs []*REQIFCONTENTAPI

	REQIFHEADERAPIs []*REQIFHEADERAPI

	REQIFTOOLEXTENSIONAPIs []*REQIFTOOLEXTENSIONAPI

	REQIFTYPEAPIs []*REQIFTYPEAPI

	SOURCEAPIs []*SOURCEAPI

	SOURCESPECIFICATIONAPIs []*SOURCESPECIFICATIONAPI

	SPECATTRIBUTESAPIs []*SPECATTRIBUTESAPI

	SPECHIERARCHYAPIs []*SPECHIERARCHYAPI

	SPECIFICATIONAPIs []*SPECIFICATIONAPI

	SPECIFICATIONSAPIs []*SPECIFICATIONSAPI

	SPECIFICATIONTYPEAPIs []*SPECIFICATIONTYPEAPI

	SPECIFIEDVALUESAPIs []*SPECIFIEDVALUESAPI

	SPECOBJECTAPIs []*SPECOBJECTAPI

	SPECOBJECTSAPIs []*SPECOBJECTSAPI

	SPECOBJECTTYPEAPIs []*SPECOBJECTTYPEAPI

	SPECRELATIONAPIs []*SPECRELATIONAPI

	SPECRELATIONGROUPSAPIs []*SPECRELATIONGROUPSAPI

	SPECRELATIONSAPIs []*SPECRELATIONSAPI

	SPECRELATIONTYPEAPIs []*SPECRELATIONTYPEAPI

	SPECTYPESAPIs []*SPECTYPESAPI

	TARGETAPIs []*TARGETAPI

	TARGETSPECIFICATIONAPIs []*TARGETSPECIFICATIONAPI

	THEHEADERAPIs []*THEHEADERAPI

	TOOLEXTENSIONSAPIs []*TOOLEXTENSIONSAPI

	VALUESAPIs []*VALUESAPI

	XHTMLCONTENTAPIs []*XHTMLCONTENTAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, alternativeidDB := range backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDDB {

		var alternativeidAPI ALTERNATIVEIDAPI
		alternativeidAPI.ID = alternativeidDB.ID
		alternativeidAPI.ALTERNATIVEIDPointersEncoding = alternativeidDB.ALTERNATIVEIDPointersEncoding
		alternativeidDB.CopyBasicFieldsToALTERNATIVEID_WOP(&alternativeidAPI.ALTERNATIVEID_WOP)

		backRepoData.ALTERNATIVEIDAPIs = append(backRepoData.ALTERNATIVEIDAPIs, &alternativeidAPI)
	}

	for _, attributedefinitionbooleanDB := range backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB {

		var attributedefinitionbooleanAPI ATTRIBUTEDEFINITIONBOOLEANAPI
		attributedefinitionbooleanAPI.ID = attributedefinitionbooleanDB.ID
		attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding = attributedefinitionbooleanDB.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
		attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN_WOP(&attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEAN_WOP)

		backRepoData.ATTRIBUTEDEFINITIONBOOLEANAPIs = append(backRepoData.ATTRIBUTEDEFINITIONBOOLEANAPIs, &attributedefinitionbooleanAPI)
	}

	for _, attributedefinitiondateDB := range backRepo.BackRepoATTRIBUTEDEFINITIONDATE.Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEDB {

		var attributedefinitiondateAPI ATTRIBUTEDEFINITIONDATEAPI
		attributedefinitiondateAPI.ID = attributedefinitiondateDB.ID
		attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATEPointersEncoding = attributedefinitiondateDB.ATTRIBUTEDEFINITIONDATEPointersEncoding
		attributedefinitiondateDB.CopyBasicFieldsToATTRIBUTEDEFINITIONDATE_WOP(&attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATE_WOP)

		backRepoData.ATTRIBUTEDEFINITIONDATEAPIs = append(backRepoData.ATTRIBUTEDEFINITIONDATEAPIs, &attributedefinitiondateAPI)
	}

	for _, attributedefinitionenumerationDB := range backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONDB {

		var attributedefinitionenumerationAPI ATTRIBUTEDEFINITIONENUMERATIONAPI
		attributedefinitionenumerationAPI.ID = attributedefinitionenumerationDB.ID
		attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding = attributedefinitionenumerationDB.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding
		attributedefinitionenumerationDB.CopyBasicFieldsToATTRIBUTEDEFINITIONENUMERATION_WOP(&attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATION_WOP)

		backRepoData.ATTRIBUTEDEFINITIONENUMERATIONAPIs = append(backRepoData.ATTRIBUTEDEFINITIONENUMERATIONAPIs, &attributedefinitionenumerationAPI)
	}

	for _, attributedefinitionintegerDB := range backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERDB {

		var attributedefinitionintegerAPI ATTRIBUTEDEFINITIONINTEGERAPI
		attributedefinitionintegerAPI.ID = attributedefinitionintegerDB.ID
		attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGERPointersEncoding = attributedefinitionintegerDB.ATTRIBUTEDEFINITIONINTEGERPointersEncoding
		attributedefinitionintegerDB.CopyBasicFieldsToATTRIBUTEDEFINITIONINTEGER_WOP(&attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGER_WOP)

		backRepoData.ATTRIBUTEDEFINITIONINTEGERAPIs = append(backRepoData.ATTRIBUTEDEFINITIONINTEGERAPIs, &attributedefinitionintegerAPI)
	}

	for _, attributedefinitionrealDB := range backRepo.BackRepoATTRIBUTEDEFINITIONREAL.Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALDB {

		var attributedefinitionrealAPI ATTRIBUTEDEFINITIONREALAPI
		attributedefinitionrealAPI.ID = attributedefinitionrealDB.ID
		attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREALPointersEncoding = attributedefinitionrealDB.ATTRIBUTEDEFINITIONREALPointersEncoding
		attributedefinitionrealDB.CopyBasicFieldsToATTRIBUTEDEFINITIONREAL_WOP(&attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREAL_WOP)

		backRepoData.ATTRIBUTEDEFINITIONREALAPIs = append(backRepoData.ATTRIBUTEDEFINITIONREALAPIs, &attributedefinitionrealAPI)
	}

	for _, attributedefinitionstringDB := range backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB {

		var attributedefinitionstringAPI ATTRIBUTEDEFINITIONSTRINGAPI
		attributedefinitionstringAPI.ID = attributedefinitionstringDB.ID
		attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRINGPointersEncoding = attributedefinitionstringDB.ATTRIBUTEDEFINITIONSTRINGPointersEncoding
		attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING_WOP(&attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRING_WOP)

		backRepoData.ATTRIBUTEDEFINITIONSTRINGAPIs = append(backRepoData.ATTRIBUTEDEFINITIONSTRINGAPIs, &attributedefinitionstringAPI)
	}

	for _, attributedefinitionxhtmlDB := range backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLDB {

		var attributedefinitionxhtmlAPI ATTRIBUTEDEFINITIONXHTMLAPI
		attributedefinitionxhtmlAPI.ID = attributedefinitionxhtmlDB.ID
		attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTMLPointersEncoding = attributedefinitionxhtmlDB.ATTRIBUTEDEFINITIONXHTMLPointersEncoding
		attributedefinitionxhtmlDB.CopyBasicFieldsToATTRIBUTEDEFINITIONXHTML_WOP(&attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTML_WOP)

		backRepoData.ATTRIBUTEDEFINITIONXHTMLAPIs = append(backRepoData.ATTRIBUTEDEFINITIONXHTMLAPIs, &attributedefinitionxhtmlAPI)
	}

	for _, attributevaluebooleanDB := range backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANDB {

		var attributevaluebooleanAPI ATTRIBUTEVALUEBOOLEANAPI
		attributevaluebooleanAPI.ID = attributevaluebooleanDB.ID
		attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEANPointersEncoding = attributevaluebooleanDB.ATTRIBUTEVALUEBOOLEANPointersEncoding
		attributevaluebooleanDB.CopyBasicFieldsToATTRIBUTEVALUEBOOLEAN_WOP(&attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEAN_WOP)

		backRepoData.ATTRIBUTEVALUEBOOLEANAPIs = append(backRepoData.ATTRIBUTEVALUEBOOLEANAPIs, &attributevaluebooleanAPI)
	}

	for _, attributevaluedateDB := range backRepo.BackRepoATTRIBUTEVALUEDATE.Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEDB {

		var attributevaluedateAPI ATTRIBUTEVALUEDATEAPI
		attributevaluedateAPI.ID = attributevaluedateDB.ID
		attributevaluedateAPI.ATTRIBUTEVALUEDATEPointersEncoding = attributevaluedateDB.ATTRIBUTEVALUEDATEPointersEncoding
		attributevaluedateDB.CopyBasicFieldsToATTRIBUTEVALUEDATE_WOP(&attributevaluedateAPI.ATTRIBUTEVALUEDATE_WOP)

		backRepoData.ATTRIBUTEVALUEDATEAPIs = append(backRepoData.ATTRIBUTEVALUEDATEAPIs, &attributevaluedateAPI)
	}

	for _, attributevalueenumerationDB := range backRepo.BackRepoATTRIBUTEVALUEENUMERATION.Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONDB {

		var attributevalueenumerationAPI ATTRIBUTEVALUEENUMERATIONAPI
		attributevalueenumerationAPI.ID = attributevalueenumerationDB.ID
		attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATIONPointersEncoding = attributevalueenumerationDB.ATTRIBUTEVALUEENUMERATIONPointersEncoding
		attributevalueenumerationDB.CopyBasicFieldsToATTRIBUTEVALUEENUMERATION_WOP(&attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATION_WOP)

		backRepoData.ATTRIBUTEVALUEENUMERATIONAPIs = append(backRepoData.ATTRIBUTEVALUEENUMERATIONAPIs, &attributevalueenumerationAPI)
	}

	for _, attributevalueintegerDB := range backRepo.BackRepoATTRIBUTEVALUEINTEGER.Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERDB {

		var attributevalueintegerAPI ATTRIBUTEVALUEINTEGERAPI
		attributevalueintegerAPI.ID = attributevalueintegerDB.ID
		attributevalueintegerAPI.ATTRIBUTEVALUEINTEGERPointersEncoding = attributevalueintegerDB.ATTRIBUTEVALUEINTEGERPointersEncoding
		attributevalueintegerDB.CopyBasicFieldsToATTRIBUTEVALUEINTEGER_WOP(&attributevalueintegerAPI.ATTRIBUTEVALUEINTEGER_WOP)

		backRepoData.ATTRIBUTEVALUEINTEGERAPIs = append(backRepoData.ATTRIBUTEVALUEINTEGERAPIs, &attributevalueintegerAPI)
	}

	for _, attributevaluerealDB := range backRepo.BackRepoATTRIBUTEVALUEREAL.Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALDB {

		var attributevaluerealAPI ATTRIBUTEVALUEREALAPI
		attributevaluerealAPI.ID = attributevaluerealDB.ID
		attributevaluerealAPI.ATTRIBUTEVALUEREALPointersEncoding = attributevaluerealDB.ATTRIBUTEVALUEREALPointersEncoding
		attributevaluerealDB.CopyBasicFieldsToATTRIBUTEVALUEREAL_WOP(&attributevaluerealAPI.ATTRIBUTEVALUEREAL_WOP)

		backRepoData.ATTRIBUTEVALUEREALAPIs = append(backRepoData.ATTRIBUTEVALUEREALAPIs, &attributevaluerealAPI)
	}

	for _, attributevaluestringDB := range backRepo.BackRepoATTRIBUTEVALUESTRING.Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGDB {

		var attributevaluestringAPI ATTRIBUTEVALUESTRINGAPI
		attributevaluestringAPI.ID = attributevaluestringDB.ID
		attributevaluestringAPI.ATTRIBUTEVALUESTRINGPointersEncoding = attributevaluestringDB.ATTRIBUTEVALUESTRINGPointersEncoding
		attributevaluestringDB.CopyBasicFieldsToATTRIBUTEVALUESTRING_WOP(&attributevaluestringAPI.ATTRIBUTEVALUESTRING_WOP)

		backRepoData.ATTRIBUTEVALUESTRINGAPIs = append(backRepoData.ATTRIBUTEVALUESTRINGAPIs, &attributevaluestringAPI)
	}

	for _, attributevaluexhtmlDB := range backRepo.BackRepoATTRIBUTEVALUEXHTML.Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLDB {

		var attributevaluexhtmlAPI ATTRIBUTEVALUEXHTMLAPI
		attributevaluexhtmlAPI.ID = attributevaluexhtmlDB.ID
		attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTMLPointersEncoding = attributevaluexhtmlDB.ATTRIBUTEVALUEXHTMLPointersEncoding
		attributevaluexhtmlDB.CopyBasicFieldsToATTRIBUTEVALUEXHTML_WOP(&attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTML_WOP)

		backRepoData.ATTRIBUTEVALUEXHTMLAPIs = append(backRepoData.ATTRIBUTEVALUEXHTMLAPIs, &attributevaluexhtmlAPI)
	}

	for _, childrenDB := range backRepo.BackRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB {

		var childrenAPI CHILDRENAPI
		childrenAPI.ID = childrenDB.ID
		childrenAPI.CHILDRENPointersEncoding = childrenDB.CHILDRENPointersEncoding
		childrenDB.CopyBasicFieldsToCHILDREN_WOP(&childrenAPI.CHILDREN_WOP)

		backRepoData.CHILDRENAPIs = append(backRepoData.CHILDRENAPIs, &childrenAPI)
	}

	for _, corecontentDB := range backRepo.BackRepoCORECONTENT.Map_CORECONTENTDBID_CORECONTENTDB {

		var corecontentAPI CORECONTENTAPI
		corecontentAPI.ID = corecontentDB.ID
		corecontentAPI.CORECONTENTPointersEncoding = corecontentDB.CORECONTENTPointersEncoding
		corecontentDB.CopyBasicFieldsToCORECONTENT_WOP(&corecontentAPI.CORECONTENT_WOP)

		backRepoData.CORECONTENTAPIs = append(backRepoData.CORECONTENTAPIs, &corecontentAPI)
	}

	for _, datatypedefinitionbooleanDB := range backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANDB {

		var datatypedefinitionbooleanAPI DATATYPEDEFINITIONBOOLEANAPI
		datatypedefinitionbooleanAPI.ID = datatypedefinitionbooleanDB.ID
		datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEANPointersEncoding = datatypedefinitionbooleanDB.DATATYPEDEFINITIONBOOLEANPointersEncoding
		datatypedefinitionbooleanDB.CopyBasicFieldsToDATATYPEDEFINITIONBOOLEAN_WOP(&datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEAN_WOP)

		backRepoData.DATATYPEDEFINITIONBOOLEANAPIs = append(backRepoData.DATATYPEDEFINITIONBOOLEANAPIs, &datatypedefinitionbooleanAPI)
	}

	for _, datatypedefinitiondateDB := range backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB {

		var datatypedefinitiondateAPI DATATYPEDEFINITIONDATEAPI
		datatypedefinitiondateAPI.ID = datatypedefinitiondateDB.ID
		datatypedefinitiondateAPI.DATATYPEDEFINITIONDATEPointersEncoding = datatypedefinitiondateDB.DATATYPEDEFINITIONDATEPointersEncoding
		datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE_WOP(&datatypedefinitiondateAPI.DATATYPEDEFINITIONDATE_WOP)

		backRepoData.DATATYPEDEFINITIONDATEAPIs = append(backRepoData.DATATYPEDEFINITIONDATEAPIs, &datatypedefinitiondateAPI)
	}

	for _, datatypedefinitionenumerationDB := range backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONDB {

		var datatypedefinitionenumerationAPI DATATYPEDEFINITIONENUMERATIONAPI
		datatypedefinitionenumerationAPI.ID = datatypedefinitionenumerationDB.ID
		datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATIONPointersEncoding = datatypedefinitionenumerationDB.DATATYPEDEFINITIONENUMERATIONPointersEncoding
		datatypedefinitionenumerationDB.CopyBasicFieldsToDATATYPEDEFINITIONENUMERATION_WOP(&datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATION_WOP)

		backRepoData.DATATYPEDEFINITIONENUMERATIONAPIs = append(backRepoData.DATATYPEDEFINITIONENUMERATIONAPIs, &datatypedefinitionenumerationAPI)
	}

	for _, datatypedefinitionintegerDB := range backRepo.BackRepoDATATYPEDEFINITIONINTEGER.Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERDB {

		var datatypedefinitionintegerAPI DATATYPEDEFINITIONINTEGERAPI
		datatypedefinitionintegerAPI.ID = datatypedefinitionintegerDB.ID
		datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGERPointersEncoding = datatypedefinitionintegerDB.DATATYPEDEFINITIONINTEGERPointersEncoding
		datatypedefinitionintegerDB.CopyBasicFieldsToDATATYPEDEFINITIONINTEGER_WOP(&datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGER_WOP)

		backRepoData.DATATYPEDEFINITIONINTEGERAPIs = append(backRepoData.DATATYPEDEFINITIONINTEGERAPIs, &datatypedefinitionintegerAPI)
	}

	for _, datatypedefinitionrealDB := range backRepo.BackRepoDATATYPEDEFINITIONREAL.Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALDB {

		var datatypedefinitionrealAPI DATATYPEDEFINITIONREALAPI
		datatypedefinitionrealAPI.ID = datatypedefinitionrealDB.ID
		datatypedefinitionrealAPI.DATATYPEDEFINITIONREALPointersEncoding = datatypedefinitionrealDB.DATATYPEDEFINITIONREALPointersEncoding
		datatypedefinitionrealDB.CopyBasicFieldsToDATATYPEDEFINITIONREAL_WOP(&datatypedefinitionrealAPI.DATATYPEDEFINITIONREAL_WOP)

		backRepoData.DATATYPEDEFINITIONREALAPIs = append(backRepoData.DATATYPEDEFINITIONREALAPIs, &datatypedefinitionrealAPI)
	}

	for _, datatypedefinitionstringDB := range backRepo.BackRepoDATATYPEDEFINITIONSTRING.Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGDB {

		var datatypedefinitionstringAPI DATATYPEDEFINITIONSTRINGAPI
		datatypedefinitionstringAPI.ID = datatypedefinitionstringDB.ID
		datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRINGPointersEncoding = datatypedefinitionstringDB.DATATYPEDEFINITIONSTRINGPointersEncoding
		datatypedefinitionstringDB.CopyBasicFieldsToDATATYPEDEFINITIONSTRING_WOP(&datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRING_WOP)

		backRepoData.DATATYPEDEFINITIONSTRINGAPIs = append(backRepoData.DATATYPEDEFINITIONSTRINGAPIs, &datatypedefinitionstringAPI)
	}

	for _, datatypedefinitionxhtmlDB := range backRepo.BackRepoDATATYPEDEFINITIONXHTML.Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLDB {

		var datatypedefinitionxhtmlAPI DATATYPEDEFINITIONXHTMLAPI
		datatypedefinitionxhtmlAPI.ID = datatypedefinitionxhtmlDB.ID
		datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTMLPointersEncoding = datatypedefinitionxhtmlDB.DATATYPEDEFINITIONXHTMLPointersEncoding
		datatypedefinitionxhtmlDB.CopyBasicFieldsToDATATYPEDEFINITIONXHTML_WOP(&datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTML_WOP)

		backRepoData.DATATYPEDEFINITIONXHTMLAPIs = append(backRepoData.DATATYPEDEFINITIONXHTMLAPIs, &datatypedefinitionxhtmlAPI)
	}

	for _, datatypesDB := range backRepo.BackRepoDATATYPES.Map_DATATYPESDBID_DATATYPESDB {

		var datatypesAPI DATATYPESAPI
		datatypesAPI.ID = datatypesDB.ID
		datatypesAPI.DATATYPESPointersEncoding = datatypesDB.DATATYPESPointersEncoding
		datatypesDB.CopyBasicFieldsToDATATYPES_WOP(&datatypesAPI.DATATYPES_WOP)

		backRepoData.DATATYPESAPIs = append(backRepoData.DATATYPESAPIs, &datatypesAPI)
	}

	for _, defaultvalueDB := range backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEDB {

		var defaultvalueAPI DEFAULTVALUEAPI
		defaultvalueAPI.ID = defaultvalueDB.ID
		defaultvalueAPI.DEFAULTVALUEPointersEncoding = defaultvalueDB.DEFAULTVALUEPointersEncoding
		defaultvalueDB.CopyBasicFieldsToDEFAULTVALUE_WOP(&defaultvalueAPI.DEFAULTVALUE_WOP)

		backRepoData.DEFAULTVALUEAPIs = append(backRepoData.DEFAULTVALUEAPIs, &defaultvalueAPI)
	}

	for _, definitionDB := range backRepo.BackRepoDEFINITION.Map_DEFINITIONDBID_DEFINITIONDB {

		var definitionAPI DEFINITIONAPI
		definitionAPI.ID = definitionDB.ID
		definitionAPI.DEFINITIONPointersEncoding = definitionDB.DEFINITIONPointersEncoding
		definitionDB.CopyBasicFieldsToDEFINITION_WOP(&definitionAPI.DEFINITION_WOP)

		backRepoData.DEFINITIONAPIs = append(backRepoData.DEFINITIONAPIs, &definitionAPI)
	}

	for _, editableattsDB := range backRepo.BackRepoEDITABLEATTS.Map_EDITABLEATTSDBID_EDITABLEATTSDB {

		var editableattsAPI EDITABLEATTSAPI
		editableattsAPI.ID = editableattsDB.ID
		editableattsAPI.EDITABLEATTSPointersEncoding = editableattsDB.EDITABLEATTSPointersEncoding
		editableattsDB.CopyBasicFieldsToEDITABLEATTS_WOP(&editableattsAPI.EDITABLEATTS_WOP)

		backRepoData.EDITABLEATTSAPIs = append(backRepoData.EDITABLEATTSAPIs, &editableattsAPI)
	}

	for _, embeddedvalueDB := range backRepo.BackRepoEMBEDDEDVALUE.Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEDB {

		var embeddedvalueAPI EMBEDDEDVALUEAPI
		embeddedvalueAPI.ID = embeddedvalueDB.ID
		embeddedvalueAPI.EMBEDDEDVALUEPointersEncoding = embeddedvalueDB.EMBEDDEDVALUEPointersEncoding
		embeddedvalueDB.CopyBasicFieldsToEMBEDDEDVALUE_WOP(&embeddedvalueAPI.EMBEDDEDVALUE_WOP)

		backRepoData.EMBEDDEDVALUEAPIs = append(backRepoData.EMBEDDEDVALUEAPIs, &embeddedvalueAPI)
	}

	for _, enumvalueDB := range backRepo.BackRepoENUMVALUE.Map_ENUMVALUEDBID_ENUMVALUEDB {

		var enumvalueAPI ENUMVALUEAPI
		enumvalueAPI.ID = enumvalueDB.ID
		enumvalueAPI.ENUMVALUEPointersEncoding = enumvalueDB.ENUMVALUEPointersEncoding
		enumvalueDB.CopyBasicFieldsToENUMVALUE_WOP(&enumvalueAPI.ENUMVALUE_WOP)

		backRepoData.ENUMVALUEAPIs = append(backRepoData.ENUMVALUEAPIs, &enumvalueAPI)
	}

	for _, objectDB := range backRepo.BackRepoOBJECT.Map_OBJECTDBID_OBJECTDB {

		var objectAPI OBJECTAPI
		objectAPI.ID = objectDB.ID
		objectAPI.OBJECTPointersEncoding = objectDB.OBJECTPointersEncoding
		objectDB.CopyBasicFieldsToOBJECT_WOP(&objectAPI.OBJECT_WOP)

		backRepoData.OBJECTAPIs = append(backRepoData.OBJECTAPIs, &objectAPI)
	}

	for _, propertiesDB := range backRepo.BackRepoPROPERTIES.Map_PROPERTIESDBID_PROPERTIESDB {

		var propertiesAPI PROPERTIESAPI
		propertiesAPI.ID = propertiesDB.ID
		propertiesAPI.PROPERTIESPointersEncoding = propertiesDB.PROPERTIESPointersEncoding
		propertiesDB.CopyBasicFieldsToPROPERTIES_WOP(&propertiesAPI.PROPERTIES_WOP)

		backRepoData.PROPERTIESAPIs = append(backRepoData.PROPERTIESAPIs, &propertiesAPI)
	}

	for _, relationgroupDB := range backRepo.BackRepoRELATIONGROUP.Map_RELATIONGROUPDBID_RELATIONGROUPDB {

		var relationgroupAPI RELATIONGROUPAPI
		relationgroupAPI.ID = relationgroupDB.ID
		relationgroupAPI.RELATIONGROUPPointersEncoding = relationgroupDB.RELATIONGROUPPointersEncoding
		relationgroupDB.CopyBasicFieldsToRELATIONGROUP_WOP(&relationgroupAPI.RELATIONGROUP_WOP)

		backRepoData.RELATIONGROUPAPIs = append(backRepoData.RELATIONGROUPAPIs, &relationgroupAPI)
	}

	for _, relationgrouptypeDB := range backRepo.BackRepoRELATIONGROUPTYPE.Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEDB {

		var relationgrouptypeAPI RELATIONGROUPTYPEAPI
		relationgrouptypeAPI.ID = relationgrouptypeDB.ID
		relationgrouptypeAPI.RELATIONGROUPTYPEPointersEncoding = relationgrouptypeDB.RELATIONGROUPTYPEPointersEncoding
		relationgrouptypeDB.CopyBasicFieldsToRELATIONGROUPTYPE_WOP(&relationgrouptypeAPI.RELATIONGROUPTYPE_WOP)

		backRepoData.RELATIONGROUPTYPEAPIs = append(backRepoData.RELATIONGROUPTYPEAPIs, &relationgrouptypeAPI)
	}

	for _, reqifDB := range backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFDB {

		var reqifAPI REQIFAPI
		reqifAPI.ID = reqifDB.ID
		reqifAPI.REQIFPointersEncoding = reqifDB.REQIFPointersEncoding
		reqifDB.CopyBasicFieldsToREQIF_WOP(&reqifAPI.REQIF_WOP)

		backRepoData.REQIFAPIs = append(backRepoData.REQIFAPIs, &reqifAPI)
	}

	for _, reqifcontentDB := range backRepo.BackRepoREQIFCONTENT.Map_REQIFCONTENTDBID_REQIFCONTENTDB {

		var reqifcontentAPI REQIFCONTENTAPI
		reqifcontentAPI.ID = reqifcontentDB.ID
		reqifcontentAPI.REQIFCONTENTPointersEncoding = reqifcontentDB.REQIFCONTENTPointersEncoding
		reqifcontentDB.CopyBasicFieldsToREQIFCONTENT_WOP(&reqifcontentAPI.REQIFCONTENT_WOP)

		backRepoData.REQIFCONTENTAPIs = append(backRepoData.REQIFCONTENTAPIs, &reqifcontentAPI)
	}

	for _, reqifheaderDB := range backRepo.BackRepoREQIFHEADER.Map_REQIFHEADERDBID_REQIFHEADERDB {

		var reqifheaderAPI REQIFHEADERAPI
		reqifheaderAPI.ID = reqifheaderDB.ID
		reqifheaderAPI.REQIFHEADERPointersEncoding = reqifheaderDB.REQIFHEADERPointersEncoding
		reqifheaderDB.CopyBasicFieldsToREQIFHEADER_WOP(&reqifheaderAPI.REQIFHEADER_WOP)

		backRepoData.REQIFHEADERAPIs = append(backRepoData.REQIFHEADERAPIs, &reqifheaderAPI)
	}

	for _, reqiftoolextensionDB := range backRepo.BackRepoREQIFTOOLEXTENSION.Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONDB {

		var reqiftoolextensionAPI REQIFTOOLEXTENSIONAPI
		reqiftoolextensionAPI.ID = reqiftoolextensionDB.ID
		reqiftoolextensionAPI.REQIFTOOLEXTENSIONPointersEncoding = reqiftoolextensionDB.REQIFTOOLEXTENSIONPointersEncoding
		reqiftoolextensionDB.CopyBasicFieldsToREQIFTOOLEXTENSION_WOP(&reqiftoolextensionAPI.REQIFTOOLEXTENSION_WOP)

		backRepoData.REQIFTOOLEXTENSIONAPIs = append(backRepoData.REQIFTOOLEXTENSIONAPIs, &reqiftoolextensionAPI)
	}

	for _, reqiftypeDB := range backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEDBID_REQIFTYPEDB {

		var reqiftypeAPI REQIFTYPEAPI
		reqiftypeAPI.ID = reqiftypeDB.ID
		reqiftypeAPI.REQIFTYPEPointersEncoding = reqiftypeDB.REQIFTYPEPointersEncoding
		reqiftypeDB.CopyBasicFieldsToREQIFTYPE_WOP(&reqiftypeAPI.REQIFTYPE_WOP)

		backRepoData.REQIFTYPEAPIs = append(backRepoData.REQIFTYPEAPIs, &reqiftypeAPI)
	}

	for _, sourceDB := range backRepo.BackRepoSOURCE.Map_SOURCEDBID_SOURCEDB {

		var sourceAPI SOURCEAPI
		sourceAPI.ID = sourceDB.ID
		sourceAPI.SOURCEPointersEncoding = sourceDB.SOURCEPointersEncoding
		sourceDB.CopyBasicFieldsToSOURCE_WOP(&sourceAPI.SOURCE_WOP)

		backRepoData.SOURCEAPIs = append(backRepoData.SOURCEAPIs, &sourceAPI)
	}

	for _, sourcespecificationDB := range backRepo.BackRepoSOURCESPECIFICATION.Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONDB {

		var sourcespecificationAPI SOURCESPECIFICATIONAPI
		sourcespecificationAPI.ID = sourcespecificationDB.ID
		sourcespecificationAPI.SOURCESPECIFICATIONPointersEncoding = sourcespecificationDB.SOURCESPECIFICATIONPointersEncoding
		sourcespecificationDB.CopyBasicFieldsToSOURCESPECIFICATION_WOP(&sourcespecificationAPI.SOURCESPECIFICATION_WOP)

		backRepoData.SOURCESPECIFICATIONAPIs = append(backRepoData.SOURCESPECIFICATIONAPIs, &sourcespecificationAPI)
	}

	for _, specattributesDB := range backRepo.BackRepoSPECATTRIBUTES.Map_SPECATTRIBUTESDBID_SPECATTRIBUTESDB {

		var specattributesAPI SPECATTRIBUTESAPI
		specattributesAPI.ID = specattributesDB.ID
		specattributesAPI.SPECATTRIBUTESPointersEncoding = specattributesDB.SPECATTRIBUTESPointersEncoding
		specattributesDB.CopyBasicFieldsToSPECATTRIBUTES_WOP(&specattributesAPI.SPECATTRIBUTES_WOP)

		backRepoData.SPECATTRIBUTESAPIs = append(backRepoData.SPECATTRIBUTESAPIs, &specattributesAPI)
	}

	for _, spechierarchyDB := range backRepo.BackRepoSPECHIERARCHY.Map_SPECHIERARCHYDBID_SPECHIERARCHYDB {

		var spechierarchyAPI SPECHIERARCHYAPI
		spechierarchyAPI.ID = spechierarchyDB.ID
		spechierarchyAPI.SPECHIERARCHYPointersEncoding = spechierarchyDB.SPECHIERARCHYPointersEncoding
		spechierarchyDB.CopyBasicFieldsToSPECHIERARCHY_WOP(&spechierarchyAPI.SPECHIERARCHY_WOP)

		backRepoData.SPECHIERARCHYAPIs = append(backRepoData.SPECHIERARCHYAPIs, &spechierarchyAPI)
	}

	for _, specificationDB := range backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {

		var specificationAPI SPECIFICATIONAPI
		specificationAPI.ID = specificationDB.ID
		specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
		specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)

		backRepoData.SPECIFICATIONAPIs = append(backRepoData.SPECIFICATIONAPIs, &specificationAPI)
	}

	for _, specificationsDB := range backRepo.BackRepoSPECIFICATIONS.Map_SPECIFICATIONSDBID_SPECIFICATIONSDB {

		var specificationsAPI SPECIFICATIONSAPI
		specificationsAPI.ID = specificationsDB.ID
		specificationsAPI.SPECIFICATIONSPointersEncoding = specificationsDB.SPECIFICATIONSPointersEncoding
		specificationsDB.CopyBasicFieldsToSPECIFICATIONS_WOP(&specificationsAPI.SPECIFICATIONS_WOP)

		backRepoData.SPECIFICATIONSAPIs = append(backRepoData.SPECIFICATIONSAPIs, &specificationsAPI)
	}

	for _, specificationtypeDB := range backRepo.BackRepoSPECIFICATIONTYPE.Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEDB {

		var specificationtypeAPI SPECIFICATIONTYPEAPI
		specificationtypeAPI.ID = specificationtypeDB.ID
		specificationtypeAPI.SPECIFICATIONTYPEPointersEncoding = specificationtypeDB.SPECIFICATIONTYPEPointersEncoding
		specificationtypeDB.CopyBasicFieldsToSPECIFICATIONTYPE_WOP(&specificationtypeAPI.SPECIFICATIONTYPE_WOP)

		backRepoData.SPECIFICATIONTYPEAPIs = append(backRepoData.SPECIFICATIONTYPEAPIs, &specificationtypeAPI)
	}

	for _, specifiedvaluesDB := range backRepo.BackRepoSPECIFIEDVALUES.Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESDB {

		var specifiedvaluesAPI SPECIFIEDVALUESAPI
		specifiedvaluesAPI.ID = specifiedvaluesDB.ID
		specifiedvaluesAPI.SPECIFIEDVALUESPointersEncoding = specifiedvaluesDB.SPECIFIEDVALUESPointersEncoding
		specifiedvaluesDB.CopyBasicFieldsToSPECIFIEDVALUES_WOP(&specifiedvaluesAPI.SPECIFIEDVALUES_WOP)

		backRepoData.SPECIFIEDVALUESAPIs = append(backRepoData.SPECIFIEDVALUESAPIs, &specifiedvaluesAPI)
	}

	for _, specobjectDB := range backRepo.BackRepoSPECOBJECT.Map_SPECOBJECTDBID_SPECOBJECTDB {

		var specobjectAPI SPECOBJECTAPI
		specobjectAPI.ID = specobjectDB.ID
		specobjectAPI.SPECOBJECTPointersEncoding = specobjectDB.SPECOBJECTPointersEncoding
		specobjectDB.CopyBasicFieldsToSPECOBJECT_WOP(&specobjectAPI.SPECOBJECT_WOP)

		backRepoData.SPECOBJECTAPIs = append(backRepoData.SPECOBJECTAPIs, &specobjectAPI)
	}

	for _, specobjectsDB := range backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB {

		var specobjectsAPI SPECOBJECTSAPI
		specobjectsAPI.ID = specobjectsDB.ID
		specobjectsAPI.SPECOBJECTSPointersEncoding = specobjectsDB.SPECOBJECTSPointersEncoding
		specobjectsDB.CopyBasicFieldsToSPECOBJECTS_WOP(&specobjectsAPI.SPECOBJECTS_WOP)

		backRepoData.SPECOBJECTSAPIs = append(backRepoData.SPECOBJECTSAPIs, &specobjectsAPI)
	}

	for _, specobjecttypeDB := range backRepo.BackRepoSPECOBJECTTYPE.Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEDB {

		var specobjecttypeAPI SPECOBJECTTYPEAPI
		specobjecttypeAPI.ID = specobjecttypeDB.ID
		specobjecttypeAPI.SPECOBJECTTYPEPointersEncoding = specobjecttypeDB.SPECOBJECTTYPEPointersEncoding
		specobjecttypeDB.CopyBasicFieldsToSPECOBJECTTYPE_WOP(&specobjecttypeAPI.SPECOBJECTTYPE_WOP)

		backRepoData.SPECOBJECTTYPEAPIs = append(backRepoData.SPECOBJECTTYPEAPIs, &specobjecttypeAPI)
	}

	for _, specrelationDB := range backRepo.BackRepoSPECRELATION.Map_SPECRELATIONDBID_SPECRELATIONDB {

		var specrelationAPI SPECRELATIONAPI
		specrelationAPI.ID = specrelationDB.ID
		specrelationAPI.SPECRELATIONPointersEncoding = specrelationDB.SPECRELATIONPointersEncoding
		specrelationDB.CopyBasicFieldsToSPECRELATION_WOP(&specrelationAPI.SPECRELATION_WOP)

		backRepoData.SPECRELATIONAPIs = append(backRepoData.SPECRELATIONAPIs, &specrelationAPI)
	}

	for _, specrelationgroupsDB := range backRepo.BackRepoSPECRELATIONGROUPS.Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSDB {

		var specrelationgroupsAPI SPECRELATIONGROUPSAPI
		specrelationgroupsAPI.ID = specrelationgroupsDB.ID
		specrelationgroupsAPI.SPECRELATIONGROUPSPointersEncoding = specrelationgroupsDB.SPECRELATIONGROUPSPointersEncoding
		specrelationgroupsDB.CopyBasicFieldsToSPECRELATIONGROUPS_WOP(&specrelationgroupsAPI.SPECRELATIONGROUPS_WOP)

		backRepoData.SPECRELATIONGROUPSAPIs = append(backRepoData.SPECRELATIONGROUPSAPIs, &specrelationgroupsAPI)
	}

	for _, specrelationsDB := range backRepo.BackRepoSPECRELATIONS.Map_SPECRELATIONSDBID_SPECRELATIONSDB {

		var specrelationsAPI SPECRELATIONSAPI
		specrelationsAPI.ID = specrelationsDB.ID
		specrelationsAPI.SPECRELATIONSPointersEncoding = specrelationsDB.SPECRELATIONSPointersEncoding
		specrelationsDB.CopyBasicFieldsToSPECRELATIONS_WOP(&specrelationsAPI.SPECRELATIONS_WOP)

		backRepoData.SPECRELATIONSAPIs = append(backRepoData.SPECRELATIONSAPIs, &specrelationsAPI)
	}

	for _, specrelationtypeDB := range backRepo.BackRepoSPECRELATIONTYPE.Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEDB {

		var specrelationtypeAPI SPECRELATIONTYPEAPI
		specrelationtypeAPI.ID = specrelationtypeDB.ID
		specrelationtypeAPI.SPECRELATIONTYPEPointersEncoding = specrelationtypeDB.SPECRELATIONTYPEPointersEncoding
		specrelationtypeDB.CopyBasicFieldsToSPECRELATIONTYPE_WOP(&specrelationtypeAPI.SPECRELATIONTYPE_WOP)

		backRepoData.SPECRELATIONTYPEAPIs = append(backRepoData.SPECRELATIONTYPEAPIs, &specrelationtypeAPI)
	}

	for _, spectypesDB := range backRepo.BackRepoSPECTYPES.Map_SPECTYPESDBID_SPECTYPESDB {

		var spectypesAPI SPECTYPESAPI
		spectypesAPI.ID = spectypesDB.ID
		spectypesAPI.SPECTYPESPointersEncoding = spectypesDB.SPECTYPESPointersEncoding
		spectypesDB.CopyBasicFieldsToSPECTYPES_WOP(&spectypesAPI.SPECTYPES_WOP)

		backRepoData.SPECTYPESAPIs = append(backRepoData.SPECTYPESAPIs, &spectypesAPI)
	}

	for _, targetDB := range backRepo.BackRepoTARGET.Map_TARGETDBID_TARGETDB {

		var targetAPI TARGETAPI
		targetAPI.ID = targetDB.ID
		targetAPI.TARGETPointersEncoding = targetDB.TARGETPointersEncoding
		targetDB.CopyBasicFieldsToTARGET_WOP(&targetAPI.TARGET_WOP)

		backRepoData.TARGETAPIs = append(backRepoData.TARGETAPIs, &targetAPI)
	}

	for _, targetspecificationDB := range backRepo.BackRepoTARGETSPECIFICATION.Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONDB {

		var targetspecificationAPI TARGETSPECIFICATIONAPI
		targetspecificationAPI.ID = targetspecificationDB.ID
		targetspecificationAPI.TARGETSPECIFICATIONPointersEncoding = targetspecificationDB.TARGETSPECIFICATIONPointersEncoding
		targetspecificationDB.CopyBasicFieldsToTARGETSPECIFICATION_WOP(&targetspecificationAPI.TARGETSPECIFICATION_WOP)

		backRepoData.TARGETSPECIFICATIONAPIs = append(backRepoData.TARGETSPECIFICATIONAPIs, &targetspecificationAPI)
	}

	for _, theheaderDB := range backRepo.BackRepoTHEHEADER.Map_THEHEADERDBID_THEHEADERDB {

		var theheaderAPI THEHEADERAPI
		theheaderAPI.ID = theheaderDB.ID
		theheaderAPI.THEHEADERPointersEncoding = theheaderDB.THEHEADERPointersEncoding
		theheaderDB.CopyBasicFieldsToTHEHEADER_WOP(&theheaderAPI.THEHEADER_WOP)

		backRepoData.THEHEADERAPIs = append(backRepoData.THEHEADERAPIs, &theheaderAPI)
	}

	for _, toolextensionsDB := range backRepo.BackRepoTOOLEXTENSIONS.Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSDB {

		var toolextensionsAPI TOOLEXTENSIONSAPI
		toolextensionsAPI.ID = toolextensionsDB.ID
		toolextensionsAPI.TOOLEXTENSIONSPointersEncoding = toolextensionsDB.TOOLEXTENSIONSPointersEncoding
		toolextensionsDB.CopyBasicFieldsToTOOLEXTENSIONS_WOP(&toolextensionsAPI.TOOLEXTENSIONS_WOP)

		backRepoData.TOOLEXTENSIONSAPIs = append(backRepoData.TOOLEXTENSIONSAPIs, &toolextensionsAPI)
	}

	for _, valuesDB := range backRepo.BackRepoVALUES.Map_VALUESDBID_VALUESDB {

		var valuesAPI VALUESAPI
		valuesAPI.ID = valuesDB.ID
		valuesAPI.VALUESPointersEncoding = valuesDB.VALUESPointersEncoding
		valuesDB.CopyBasicFieldsToVALUES_WOP(&valuesAPI.VALUES_WOP)

		backRepoData.VALUESAPIs = append(backRepoData.VALUESAPIs, &valuesAPI)
	}

	for _, xhtmlcontentDB := range backRepo.BackRepoXHTMLCONTENT.Map_XHTMLCONTENTDBID_XHTMLCONTENTDB {

		var xhtmlcontentAPI XHTMLCONTENTAPI
		xhtmlcontentAPI.ID = xhtmlcontentDB.ID
		xhtmlcontentAPI.XHTMLCONTENTPointersEncoding = xhtmlcontentDB.XHTMLCONTENTPointersEncoding
		xhtmlcontentDB.CopyBasicFieldsToXHTMLCONTENT_WOP(&xhtmlcontentAPI.XHTMLCONTENT_WOP)

		backRepoData.XHTMLCONTENTAPIs = append(backRepoData.XHTMLCONTENTAPIs, &xhtmlcontentAPI)
	}

}
