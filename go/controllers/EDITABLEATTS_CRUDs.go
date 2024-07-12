// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __EDITABLEATTS__dummysDeclaration__ models.EDITABLEATTS
var __EDITABLEATTS_time__dummyDeclaration time.Duration

var mutexEDITABLEATTS sync.Mutex

// An EDITABLEATTSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEDITABLEATTS updateEDITABLEATTS deleteEDITABLEATTS
type EDITABLEATTSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EDITABLEATTSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEDITABLEATTS updateEDITABLEATTS
type EDITABLEATTSInput struct {
	// The EDITABLEATTS to submit or modify
	// in: body
	EDITABLEATTS *orm.EDITABLEATTSAPI
}

// GetEDITABLEATTSs
//
// swagger:route GET /editableattss editableattss getEDITABLEATTSs
//
// # Get all editableattss
//
// Responses:
// default: genericError
//
//	200: editableattsDBResponse
func (controller *Controller) GetEDITABLEATTSs(c *gin.Context) {

	// source slice
	var editableattsDBs []orm.EDITABLEATTSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEDITABLEATTSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEDITABLEATTS.GetDB()

	query := db.Find(&editableattsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	editableattsAPIs := make([]orm.EDITABLEATTSAPI, 0)

	// for each editableatts, update fields from the database nullable fields
	for idx := range editableattsDBs {
		editableattsDB := &editableattsDBs[idx]
		_ = editableattsDB
		var editableattsAPI orm.EDITABLEATTSAPI

		// insertion point for updating fields
		editableattsAPI.ID = editableattsDB.ID
		editableattsDB.CopyBasicFieldsToEDITABLEATTS_WOP(&editableattsAPI.EDITABLEATTS_WOP)
		editableattsAPI.EDITABLEATTSPointersEncoding = editableattsDB.EDITABLEATTSPointersEncoding
		editableattsAPIs = append(editableattsAPIs, editableattsAPI)
	}

	c.JSON(http.StatusOK, editableattsAPIs)
}

// PostEDITABLEATTS
//
// swagger:route POST /editableattss editableattss postEDITABLEATTS
//
// Creates a editableatts
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEDITABLEATTS(c *gin.Context) {

	mutexEDITABLEATTS.Lock()
	defer mutexEDITABLEATTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEDITABLEATTSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEDITABLEATTS.GetDB()

	// Validate input
	var input orm.EDITABLEATTSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create editableatts
	editableattsDB := orm.EDITABLEATTSDB{}
	editableattsDB.EDITABLEATTSPointersEncoding = input.EDITABLEATTSPointersEncoding
	editableattsDB.CopyBasicFieldsFromEDITABLEATTS_WOP(&input.EDITABLEATTS_WOP)

	query := db.Create(&editableattsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEDITABLEATTS.CheckoutPhaseOneInstance(&editableattsDB)
	editableatts := backRepo.BackRepoEDITABLEATTS.Map_EDITABLEATTSDBID_EDITABLEATTSPtr[editableattsDB.ID]

	if editableatts != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), editableatts)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, editableattsDB)
}

// GetEDITABLEATTS
//
// swagger:route GET /editableattss/{ID} editableattss getEDITABLEATTS
//
// Gets the details for a editableatts.
//
// Responses:
// default: genericError
//
//	200: editableattsDBResponse
func (controller *Controller) GetEDITABLEATTS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEDITABLEATTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEDITABLEATTS.GetDB()

	// Get editableattsDB in DB
	var editableattsDB orm.EDITABLEATTSDB
	if err := db.First(&editableattsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var editableattsAPI orm.EDITABLEATTSAPI
	editableattsAPI.ID = editableattsDB.ID
	editableattsAPI.EDITABLEATTSPointersEncoding = editableattsDB.EDITABLEATTSPointersEncoding
	editableattsDB.CopyBasicFieldsToEDITABLEATTS_WOP(&editableattsAPI.EDITABLEATTS_WOP)

	c.JSON(http.StatusOK, editableattsAPI)
}

// UpdateEDITABLEATTS
//
// swagger:route PATCH /editableattss/{ID} editableattss updateEDITABLEATTS
//
// # Update a editableatts
//
// Responses:
// default: genericError
//
//	200: editableattsDBResponse
func (controller *Controller) UpdateEDITABLEATTS(c *gin.Context) {

	mutexEDITABLEATTS.Lock()
	defer mutexEDITABLEATTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEDITABLEATTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEDITABLEATTS.GetDB()

	// Validate input
	var input orm.EDITABLEATTSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var editableattsDB orm.EDITABLEATTSDB

	// fetch the editableatts
	query := db.First(&editableattsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	editableattsDB.CopyBasicFieldsFromEDITABLEATTS_WOP(&input.EDITABLEATTS_WOP)
	editableattsDB.EDITABLEATTSPointersEncoding = input.EDITABLEATTSPointersEncoding

	query = db.Model(&editableattsDB).Updates(editableattsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	editableattsNew := new(models.EDITABLEATTS)
	editableattsDB.CopyBasicFieldsToEDITABLEATTS(editableattsNew)

	// redeem pointers
	editableattsDB.DecodePointers(backRepo, editableattsNew)

	// get stage instance from DB instance, and call callback function
	editableattsOld := backRepo.BackRepoEDITABLEATTS.Map_EDITABLEATTSDBID_EDITABLEATTSPtr[editableattsDB.ID]
	if editableattsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), editableattsOld, editableattsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the editableattsDB
	c.JSON(http.StatusOK, editableattsDB)
}

// DeleteEDITABLEATTS
//
// swagger:route DELETE /editableattss/{ID} editableattss deleteEDITABLEATTS
//
// # Delete a editableatts
//
// default: genericError
//
//	200: editableattsDBResponse
func (controller *Controller) DeleteEDITABLEATTS(c *gin.Context) {

	mutexEDITABLEATTS.Lock()
	defer mutexEDITABLEATTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEDITABLEATTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEDITABLEATTS.GetDB()

	// Get model if exist
	var editableattsDB orm.EDITABLEATTSDB
	if err := db.First(&editableattsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&editableattsDB)

	// get an instance (not staged) from DB instance, and call callback function
	editableattsDeleted := new(models.EDITABLEATTS)
	editableattsDB.CopyBasicFieldsToEDITABLEATTS(editableattsDeleted)

	// get stage instance from DB instance, and call callback function
	editableattsStaged := backRepo.BackRepoEDITABLEATTS.Map_EDITABLEATTSDBID_EDITABLEATTSPtr[editableattsDB.ID]
	if editableattsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), editableattsStaged, editableattsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
