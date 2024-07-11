// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	REQIFAPIs []*REQIFAPI

	REQIFHEADERAPIs []*REQIFHEADERAPI
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

	for _, reqifheaderDB := range backRepo.BackRepoREQIFHEADER.Map_REQIFHEADERDBID_REQIFHEADERDB {

		var reqifheaderAPI REQIFHEADERAPI
		reqifheaderAPI.ID = reqifheaderDB.ID
		reqifheaderAPI.REQIFHEADERPointersEncoding = reqifheaderDB.REQIFHEADERPointersEncoding
		reqifheaderDB.CopyBasicFieldsToREQIFHEADER_WOP(&reqifheaderAPI.REQIFHEADER_WOP)

		backRepoData.REQIFHEADERAPIs = append(backRepoData.REQIFHEADERAPIs, &reqifheaderAPI)
	}

}
