// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	REQIFAPIs []*REQIFAPI

	REQ_IF_CONTENTAPIs []*REQ_IF_CONTENTAPI

	REQ_IF_HEADERAPIs []*REQ_IF_HEADERAPI

	SPECIFICATIONAPIs []*SPECIFICATIONAPI

	SPECIFICATION_TYPEAPIs []*SPECIFICATION_TYPEAPI

	SPEC_HIERARCHYAPIs []*SPEC_HIERARCHYAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, reqifDB := range backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFDB {

		var reqifAPI REQIFAPI
		reqifAPI.ID = reqifDB.ID
		reqifAPI.REQIFPointersEncoding = reqifDB.REQIFPointersEncoding
		reqifDB.CopyBasicFieldsToREQIF_WOP(&reqifAPI.REQIF_WOP)

		backRepoData.REQIFAPIs = append(backRepoData.REQIFAPIs, &reqifAPI)
	}

	for _, req_if_contentDB := range backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {

		var req_if_contentAPI REQ_IF_CONTENTAPI
		req_if_contentAPI.ID = req_if_contentDB.ID
		req_if_contentAPI.REQ_IF_CONTENTPointersEncoding = req_if_contentDB.REQ_IF_CONTENTPointersEncoding
		req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT_WOP(&req_if_contentAPI.REQ_IF_CONTENT_WOP)

		backRepoData.REQ_IF_CONTENTAPIs = append(backRepoData.REQ_IF_CONTENTAPIs, &req_if_contentAPI)
	}

	for _, req_if_headerDB := range backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB {

		var req_if_headerAPI REQ_IF_HEADERAPI
		req_if_headerAPI.ID = req_if_headerDB.ID
		req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
		req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)

		backRepoData.REQ_IF_HEADERAPIs = append(backRepoData.REQ_IF_HEADERAPIs, &req_if_headerAPI)
	}

	for _, specificationDB := range backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {

		var specificationAPI SPECIFICATIONAPI
		specificationAPI.ID = specificationDB.ID
		specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
		specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)

		backRepoData.SPECIFICATIONAPIs = append(backRepoData.SPECIFICATIONAPIs, &specificationAPI)
	}

	for _, specification_typeDB := range backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEDB {

		var specification_typeAPI SPECIFICATION_TYPEAPI
		specification_typeAPI.ID = specification_typeDB.ID
		specification_typeAPI.SPECIFICATION_TYPEPointersEncoding = specification_typeDB.SPECIFICATION_TYPEPointersEncoding
		specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE_WOP(&specification_typeAPI.SPECIFICATION_TYPE_WOP)

		backRepoData.SPECIFICATION_TYPEAPIs = append(backRepoData.SPECIFICATION_TYPEAPIs, &specification_typeAPI)
	}

	for _, spec_hierarchyDB := range backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {

		var spec_hierarchyAPI SPEC_HIERARCHYAPI
		spec_hierarchyAPI.ID = spec_hierarchyDB.ID
		spec_hierarchyAPI.SPEC_HIERARCHYPointersEncoding = spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding
		spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY_WOP(&spec_hierarchyAPI.SPEC_HIERARCHY_WOP)

		backRepoData.SPEC_HIERARCHYAPIs = append(backRepoData.SPEC_HIERARCHYAPIs, &spec_hierarchyAPI)
	}

}
