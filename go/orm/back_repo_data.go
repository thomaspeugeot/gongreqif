// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	REQIFAPIs []*REQIFAPI

	REQ_IF_CONTENTAPIs []*REQ_IF_CONTENTAPI

	REQ_IF_HEADERAPIs []*REQ_IF_HEADERAPI

	SPECIFICATIONAPIs []*SPECIFICATIONAPI
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

}
