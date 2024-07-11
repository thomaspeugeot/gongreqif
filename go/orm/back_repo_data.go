// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CONTENTAPIs []*CONTENTAPI

	HEADERAPIs []*HEADERAPI

	REQIFAPIs []*REQIFAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, contentDB := range backRepo.BackRepoCONTENT.Map_CONTENTDBID_CONTENTDB {

		var contentAPI CONTENTAPI
		contentAPI.ID = contentDB.ID
		contentAPI.CONTENTPointersEncoding = contentDB.CONTENTPointersEncoding
		contentDB.CopyBasicFieldsToCONTENT_WOP(&contentAPI.CONTENT_WOP)

		backRepoData.CONTENTAPIs = append(backRepoData.CONTENTAPIs, &contentAPI)
	}

	for _, headerDB := range backRepo.BackRepoHEADER.Map_HEADERDBID_HEADERDB {

		var headerAPI HEADERAPI
		headerAPI.ID = headerDB.ID
		headerAPI.HEADERPointersEncoding = headerDB.HEADERPointersEncoding
		headerDB.CopyBasicFieldsToHEADER_WOP(&headerAPI.HEADER_WOP)

		backRepoData.HEADERAPIs = append(backRepoData.HEADERAPIs, &headerAPI)
	}

	for _, reqifDB := range backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFDB {

		var reqifAPI REQIFAPI
		reqifAPI.ID = reqifDB.ID
		reqifAPI.REQIFPointersEncoding = reqifDB.REQIFPointersEncoding
		reqifDB.CopyBasicFieldsToREQIF_WOP(&reqifAPI.REQIF_WOP)

		backRepoData.REQIFAPIs = append(backRepoData.REQIFAPIs, &reqifAPI)
	}

}
