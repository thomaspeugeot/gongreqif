// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	REQIFAPIs []*REQIFAPI

	REQ_IF_HEADERAPIs []*REQ_IF_HEADERAPI
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

	for _, req_if_headerDB := range backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB {

		var req_if_headerAPI REQ_IF_HEADERAPI
		req_if_headerAPI.ID = req_if_headerDB.ID
		req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
		req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)

		backRepoData.REQ_IF_HEADERAPIs = append(backRepoData.REQ_IF_HEADERAPIs, &req_if_headerAPI)
	}

}
