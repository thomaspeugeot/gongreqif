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
var __THEHEADER__dummysDeclaration__ models.THEHEADER
var __THEHEADER_time__dummyDeclaration time.Duration

var mutexTHEHEADER sync.Mutex

// An THEHEADERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTHEHEADER updateTHEHEADER deleteTHEHEADER
type THEHEADERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// THEHEADERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTHEHEADER updateTHEHEADER
type THEHEADERInput struct {
	// The THEHEADER to submit or modify
	// in: body
	THEHEADER *orm.THEHEADERAPI
}

// GetTHEHEADERs
//
// swagger:route GET /theheaders theheaders getTHEHEADERs
//
// # Get all theheaders
//
// Responses:
// default: genericError
//
//	200: theheaderDBResponse
func (controller *Controller) GetTHEHEADERs(c *gin.Context) {

	// source slice
	var theheaderDBs []orm.THEHEADERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTHEHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTHEHEADER.GetDB()

	query := db.Find(&theheaderDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	theheaderAPIs := make([]orm.THEHEADERAPI, 0)

	// for each theheader, update fields from the database nullable fields
	for idx := range theheaderDBs {
		theheaderDB := &theheaderDBs[idx]
		_ = theheaderDB
		var theheaderAPI orm.THEHEADERAPI

		// insertion point for updating fields
		theheaderAPI.ID = theheaderDB.ID
		theheaderDB.CopyBasicFieldsToTHEHEADER_WOP(&theheaderAPI.THEHEADER_WOP)
		theheaderAPI.THEHEADERPointersEncoding = theheaderDB.THEHEADERPointersEncoding
		theheaderAPIs = append(theheaderAPIs, theheaderAPI)
	}

	c.JSON(http.StatusOK, theheaderAPIs)
}

// PostTHEHEADER
//
// swagger:route POST /theheaders theheaders postTHEHEADER
//
// Creates a theheader
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTHEHEADER(c *gin.Context) {

	mutexTHEHEADER.Lock()
	defer mutexTHEHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTHEHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTHEHEADER.GetDB()

	// Validate input
	var input orm.THEHEADERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create theheader
	theheaderDB := orm.THEHEADERDB{}
	theheaderDB.THEHEADERPointersEncoding = input.THEHEADERPointersEncoding
	theheaderDB.CopyBasicFieldsFromTHEHEADER_WOP(&input.THEHEADER_WOP)

	query := db.Create(&theheaderDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTHEHEADER.CheckoutPhaseOneInstance(&theheaderDB)
	theheader := backRepo.BackRepoTHEHEADER.Map_THEHEADERDBID_THEHEADERPtr[theheaderDB.ID]

	if theheader != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), theheader)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, theheaderDB)
}

// GetTHEHEADER
//
// swagger:route GET /theheaders/{ID} theheaders getTHEHEADER
//
// Gets the details for a theheader.
//
// Responses:
// default: genericError
//
//	200: theheaderDBResponse
func (controller *Controller) GetTHEHEADER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTHEHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTHEHEADER.GetDB()

	// Get theheaderDB in DB
	var theheaderDB orm.THEHEADERDB
	if err := db.First(&theheaderDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var theheaderAPI orm.THEHEADERAPI
	theheaderAPI.ID = theheaderDB.ID
	theheaderAPI.THEHEADERPointersEncoding = theheaderDB.THEHEADERPointersEncoding
	theheaderDB.CopyBasicFieldsToTHEHEADER_WOP(&theheaderAPI.THEHEADER_WOP)

	c.JSON(http.StatusOK, theheaderAPI)
}

// UpdateTHEHEADER
//
// swagger:route PATCH /theheaders/{ID} theheaders updateTHEHEADER
//
// # Update a theheader
//
// Responses:
// default: genericError
//
//	200: theheaderDBResponse
func (controller *Controller) UpdateTHEHEADER(c *gin.Context) {

	mutexTHEHEADER.Lock()
	defer mutexTHEHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTHEHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTHEHEADER.GetDB()

	// Validate input
	var input orm.THEHEADERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var theheaderDB orm.THEHEADERDB

	// fetch the theheader
	query := db.First(&theheaderDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	theheaderDB.CopyBasicFieldsFromTHEHEADER_WOP(&input.THEHEADER_WOP)
	theheaderDB.THEHEADERPointersEncoding = input.THEHEADERPointersEncoding

	query = db.Model(&theheaderDB).Updates(theheaderDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	theheaderNew := new(models.THEHEADER)
	theheaderDB.CopyBasicFieldsToTHEHEADER(theheaderNew)

	// redeem pointers
	theheaderDB.DecodePointers(backRepo, theheaderNew)

	// get stage instance from DB instance, and call callback function
	theheaderOld := backRepo.BackRepoTHEHEADER.Map_THEHEADERDBID_THEHEADERPtr[theheaderDB.ID]
	if theheaderOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), theheaderOld, theheaderNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the theheaderDB
	c.JSON(http.StatusOK, theheaderDB)
}

// DeleteTHEHEADER
//
// swagger:route DELETE /theheaders/{ID} theheaders deleteTHEHEADER
//
// # Delete a theheader
//
// default: genericError
//
//	200: theheaderDBResponse
func (controller *Controller) DeleteTHEHEADER(c *gin.Context) {

	mutexTHEHEADER.Lock()
	defer mutexTHEHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTHEHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTHEHEADER.GetDB()

	// Get model if exist
	var theheaderDB orm.THEHEADERDB
	if err := db.First(&theheaderDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&theheaderDB)

	// get an instance (not staged) from DB instance, and call callback function
	theheaderDeleted := new(models.THEHEADER)
	theheaderDB.CopyBasicFieldsToTHEHEADER(theheaderDeleted)

	// get stage instance from DB instance, and call callback function
	theheaderStaged := backRepo.BackRepoTHEHEADER.Map_THEHEADERDBID_THEHEADERPtr[theheaderDB.ID]
	if theheaderStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), theheaderStaged, theheaderDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
