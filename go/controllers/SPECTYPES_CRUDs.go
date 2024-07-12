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
var __SPECTYPES__dummysDeclaration__ models.SPECTYPES
var __SPECTYPES_time__dummyDeclaration time.Duration

var mutexSPECTYPES sync.Mutex

// An SPECTYPESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECTYPES updateSPECTYPES deleteSPECTYPES
type SPECTYPESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECTYPESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECTYPES updateSPECTYPES
type SPECTYPESInput struct {
	// The SPECTYPES to submit or modify
	// in: body
	SPECTYPES *orm.SPECTYPESAPI
}

// GetSPECTYPESs
//
// swagger:route GET /spectypess spectypess getSPECTYPESs
//
// # Get all spectypess
//
// Responses:
// default: genericError
//
//	200: spectypesDBResponse
func (controller *Controller) GetSPECTYPESs(c *gin.Context) {

	// source slice
	var spectypesDBs []orm.SPECTYPESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECTYPESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECTYPES.GetDB()

	query := db.Find(&spectypesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spectypesAPIs := make([]orm.SPECTYPESAPI, 0)

	// for each spectypes, update fields from the database nullable fields
	for idx := range spectypesDBs {
		spectypesDB := &spectypesDBs[idx]
		_ = spectypesDB
		var spectypesAPI orm.SPECTYPESAPI

		// insertion point for updating fields
		spectypesAPI.ID = spectypesDB.ID
		spectypesDB.CopyBasicFieldsToSPECTYPES_WOP(&spectypesAPI.SPECTYPES_WOP)
		spectypesAPI.SPECTYPESPointersEncoding = spectypesDB.SPECTYPESPointersEncoding
		spectypesAPIs = append(spectypesAPIs, spectypesAPI)
	}

	c.JSON(http.StatusOK, spectypesAPIs)
}

// PostSPECTYPES
//
// swagger:route POST /spectypess spectypess postSPECTYPES
//
// Creates a spectypes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECTYPES(c *gin.Context) {

	mutexSPECTYPES.Lock()
	defer mutexSPECTYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECTYPESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECTYPES.GetDB()

	// Validate input
	var input orm.SPECTYPESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spectypes
	spectypesDB := orm.SPECTYPESDB{}
	spectypesDB.SPECTYPESPointersEncoding = input.SPECTYPESPointersEncoding
	spectypesDB.CopyBasicFieldsFromSPECTYPES_WOP(&input.SPECTYPES_WOP)

	query := db.Create(&spectypesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECTYPES.CheckoutPhaseOneInstance(&spectypesDB)
	spectypes := backRepo.BackRepoSPECTYPES.Map_SPECTYPESDBID_SPECTYPESPtr[spectypesDB.ID]

	if spectypes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spectypes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spectypesDB)
}

// GetSPECTYPES
//
// swagger:route GET /spectypess/{ID} spectypess getSPECTYPES
//
// Gets the details for a spectypes.
//
// Responses:
// default: genericError
//
//	200: spectypesDBResponse
func (controller *Controller) GetSPECTYPES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECTYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECTYPES.GetDB()

	// Get spectypesDB in DB
	var spectypesDB orm.SPECTYPESDB
	if err := db.First(&spectypesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spectypesAPI orm.SPECTYPESAPI
	spectypesAPI.ID = spectypesDB.ID
	spectypesAPI.SPECTYPESPointersEncoding = spectypesDB.SPECTYPESPointersEncoding
	spectypesDB.CopyBasicFieldsToSPECTYPES_WOP(&spectypesAPI.SPECTYPES_WOP)

	c.JSON(http.StatusOK, spectypesAPI)
}

// UpdateSPECTYPES
//
// swagger:route PATCH /spectypess/{ID} spectypess updateSPECTYPES
//
// # Update a spectypes
//
// Responses:
// default: genericError
//
//	200: spectypesDBResponse
func (controller *Controller) UpdateSPECTYPES(c *gin.Context) {

	mutexSPECTYPES.Lock()
	defer mutexSPECTYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECTYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECTYPES.GetDB()

	// Validate input
	var input orm.SPECTYPESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spectypesDB orm.SPECTYPESDB

	// fetch the spectypes
	query := db.First(&spectypesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spectypesDB.CopyBasicFieldsFromSPECTYPES_WOP(&input.SPECTYPES_WOP)
	spectypesDB.SPECTYPESPointersEncoding = input.SPECTYPESPointersEncoding

	query = db.Model(&spectypesDB).Updates(spectypesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spectypesNew := new(models.SPECTYPES)
	spectypesDB.CopyBasicFieldsToSPECTYPES(spectypesNew)

	// redeem pointers
	spectypesDB.DecodePointers(backRepo, spectypesNew)

	// get stage instance from DB instance, and call callback function
	spectypesOld := backRepo.BackRepoSPECTYPES.Map_SPECTYPESDBID_SPECTYPESPtr[spectypesDB.ID]
	if spectypesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spectypesOld, spectypesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spectypesDB
	c.JSON(http.StatusOK, spectypesDB)
}

// DeleteSPECTYPES
//
// swagger:route DELETE /spectypess/{ID} spectypess deleteSPECTYPES
//
// # Delete a spectypes
//
// default: genericError
//
//	200: spectypesDBResponse
func (controller *Controller) DeleteSPECTYPES(c *gin.Context) {

	mutexSPECTYPES.Lock()
	defer mutexSPECTYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECTYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECTYPES.GetDB()

	// Get model if exist
	var spectypesDB orm.SPECTYPESDB
	if err := db.First(&spectypesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spectypesDB)

	// get an instance (not staged) from DB instance, and call callback function
	spectypesDeleted := new(models.SPECTYPES)
	spectypesDB.CopyBasicFieldsToSPECTYPES(spectypesDeleted)

	// get stage instance from DB instance, and call callback function
	spectypesStaged := backRepo.BackRepoSPECTYPES.Map_SPECTYPESDBID_SPECTYPESPtr[spectypesDB.ID]
	if spectypesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spectypesStaged, spectypesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
