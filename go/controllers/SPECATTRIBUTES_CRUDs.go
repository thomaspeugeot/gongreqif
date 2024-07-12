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
var __SPECATTRIBUTES__dummysDeclaration__ models.SPECATTRIBUTES
var __SPECATTRIBUTES_time__dummyDeclaration time.Duration

var mutexSPECATTRIBUTES sync.Mutex

// An SPECATTRIBUTESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECATTRIBUTES updateSPECATTRIBUTES deleteSPECATTRIBUTES
type SPECATTRIBUTESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECATTRIBUTESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECATTRIBUTES updateSPECATTRIBUTES
type SPECATTRIBUTESInput struct {
	// The SPECATTRIBUTES to submit or modify
	// in: body
	SPECATTRIBUTES *orm.SPECATTRIBUTESAPI
}

// GetSPECATTRIBUTESs
//
// swagger:route GET /specattributess specattributess getSPECATTRIBUTESs
//
// # Get all specattributess
//
// Responses:
// default: genericError
//
//	200: specattributesDBResponse
func (controller *Controller) GetSPECATTRIBUTESs(c *gin.Context) {

	// source slice
	var specattributesDBs []orm.SPECATTRIBUTESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECATTRIBUTESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECATTRIBUTES.GetDB()

	query := db.Find(&specattributesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specattributesAPIs := make([]orm.SPECATTRIBUTESAPI, 0)

	// for each specattributes, update fields from the database nullable fields
	for idx := range specattributesDBs {
		specattributesDB := &specattributesDBs[idx]
		_ = specattributesDB
		var specattributesAPI orm.SPECATTRIBUTESAPI

		// insertion point for updating fields
		specattributesAPI.ID = specattributesDB.ID
		specattributesDB.CopyBasicFieldsToSPECATTRIBUTES_WOP(&specattributesAPI.SPECATTRIBUTES_WOP)
		specattributesAPI.SPECATTRIBUTESPointersEncoding = specattributesDB.SPECATTRIBUTESPointersEncoding
		specattributesAPIs = append(specattributesAPIs, specattributesAPI)
	}

	c.JSON(http.StatusOK, specattributesAPIs)
}

// PostSPECATTRIBUTES
//
// swagger:route POST /specattributess specattributess postSPECATTRIBUTES
//
// Creates a specattributes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECATTRIBUTES(c *gin.Context) {

	mutexSPECATTRIBUTES.Lock()
	defer mutexSPECATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECATTRIBUTESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECATTRIBUTES.GetDB()

	// Validate input
	var input orm.SPECATTRIBUTESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specattributes
	specattributesDB := orm.SPECATTRIBUTESDB{}
	specattributesDB.SPECATTRIBUTESPointersEncoding = input.SPECATTRIBUTESPointersEncoding
	specattributesDB.CopyBasicFieldsFromSPECATTRIBUTES_WOP(&input.SPECATTRIBUTES_WOP)

	query := db.Create(&specattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECATTRIBUTES.CheckoutPhaseOneInstance(&specattributesDB)
	specattributes := backRepo.BackRepoSPECATTRIBUTES.Map_SPECATTRIBUTESDBID_SPECATTRIBUTESPtr[specattributesDB.ID]

	if specattributes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specattributes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specattributesDB)
}

// GetSPECATTRIBUTES
//
// swagger:route GET /specattributess/{ID} specattributess getSPECATTRIBUTES
//
// Gets the details for a specattributes.
//
// Responses:
// default: genericError
//
//	200: specattributesDBResponse
func (controller *Controller) GetSPECATTRIBUTES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECATTRIBUTES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECATTRIBUTES.GetDB()

	// Get specattributesDB in DB
	var specattributesDB orm.SPECATTRIBUTESDB
	if err := db.First(&specattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specattributesAPI orm.SPECATTRIBUTESAPI
	specattributesAPI.ID = specattributesDB.ID
	specattributesAPI.SPECATTRIBUTESPointersEncoding = specattributesDB.SPECATTRIBUTESPointersEncoding
	specattributesDB.CopyBasicFieldsToSPECATTRIBUTES_WOP(&specattributesAPI.SPECATTRIBUTES_WOP)

	c.JSON(http.StatusOK, specattributesAPI)
}

// UpdateSPECATTRIBUTES
//
// swagger:route PATCH /specattributess/{ID} specattributess updateSPECATTRIBUTES
//
// # Update a specattributes
//
// Responses:
// default: genericError
//
//	200: specattributesDBResponse
func (controller *Controller) UpdateSPECATTRIBUTES(c *gin.Context) {

	mutexSPECATTRIBUTES.Lock()
	defer mutexSPECATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECATTRIBUTES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECATTRIBUTES.GetDB()

	// Validate input
	var input orm.SPECATTRIBUTESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specattributesDB orm.SPECATTRIBUTESDB

	// fetch the specattributes
	query := db.First(&specattributesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specattributesDB.CopyBasicFieldsFromSPECATTRIBUTES_WOP(&input.SPECATTRIBUTES_WOP)
	specattributesDB.SPECATTRIBUTESPointersEncoding = input.SPECATTRIBUTESPointersEncoding

	query = db.Model(&specattributesDB).Updates(specattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specattributesNew := new(models.SPECATTRIBUTES)
	specattributesDB.CopyBasicFieldsToSPECATTRIBUTES(specattributesNew)

	// redeem pointers
	specattributesDB.DecodePointers(backRepo, specattributesNew)

	// get stage instance from DB instance, and call callback function
	specattributesOld := backRepo.BackRepoSPECATTRIBUTES.Map_SPECATTRIBUTESDBID_SPECATTRIBUTESPtr[specattributesDB.ID]
	if specattributesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specattributesOld, specattributesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specattributesDB
	c.JSON(http.StatusOK, specattributesDB)
}

// DeleteSPECATTRIBUTES
//
// swagger:route DELETE /specattributess/{ID} specattributess deleteSPECATTRIBUTES
//
// # Delete a specattributes
//
// default: genericError
//
//	200: specattributesDBResponse
func (controller *Controller) DeleteSPECATTRIBUTES(c *gin.Context) {

	mutexSPECATTRIBUTES.Lock()
	defer mutexSPECATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECATTRIBUTES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECATTRIBUTES.GetDB()

	// Get model if exist
	var specattributesDB orm.SPECATTRIBUTESDB
	if err := db.First(&specattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specattributesDB)

	// get an instance (not staged) from DB instance, and call callback function
	specattributesDeleted := new(models.SPECATTRIBUTES)
	specattributesDB.CopyBasicFieldsToSPECATTRIBUTES(specattributesDeleted)

	// get stage instance from DB instance, and call callback function
	specattributesStaged := backRepo.BackRepoSPECATTRIBUTES.Map_SPECATTRIBUTESDBID_SPECATTRIBUTESPtr[specattributesDB.ID]
	if specattributesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specattributesStaged, specattributesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
