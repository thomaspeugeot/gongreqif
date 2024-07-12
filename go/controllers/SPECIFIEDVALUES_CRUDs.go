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
var __SPECIFIEDVALUES__dummysDeclaration__ models.SPECIFIEDVALUES
var __SPECIFIEDVALUES_time__dummyDeclaration time.Duration

var mutexSPECIFIEDVALUES sync.Mutex

// An SPECIFIEDVALUESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECIFIEDVALUES updateSPECIFIEDVALUES deleteSPECIFIEDVALUES
type SPECIFIEDVALUESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECIFIEDVALUESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECIFIEDVALUES updateSPECIFIEDVALUES
type SPECIFIEDVALUESInput struct {
	// The SPECIFIEDVALUES to submit or modify
	// in: body
	SPECIFIEDVALUES *orm.SPECIFIEDVALUESAPI
}

// GetSPECIFIEDVALUESs
//
// swagger:route GET /specifiedvaluess specifiedvaluess getSPECIFIEDVALUESs
//
// # Get all specifiedvaluess
//
// Responses:
// default: genericError
//
//	200: specifiedvaluesDBResponse
func (controller *Controller) GetSPECIFIEDVALUESs(c *gin.Context) {

	// source slice
	var specifiedvaluesDBs []orm.SPECIFIEDVALUESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFIEDVALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFIEDVALUES.GetDB()

	query := db.Find(&specifiedvaluesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specifiedvaluesAPIs := make([]orm.SPECIFIEDVALUESAPI, 0)

	// for each specifiedvalues, update fields from the database nullable fields
	for idx := range specifiedvaluesDBs {
		specifiedvaluesDB := &specifiedvaluesDBs[idx]
		_ = specifiedvaluesDB
		var specifiedvaluesAPI orm.SPECIFIEDVALUESAPI

		// insertion point for updating fields
		specifiedvaluesAPI.ID = specifiedvaluesDB.ID
		specifiedvaluesDB.CopyBasicFieldsToSPECIFIEDVALUES_WOP(&specifiedvaluesAPI.SPECIFIEDVALUES_WOP)
		specifiedvaluesAPI.SPECIFIEDVALUESPointersEncoding = specifiedvaluesDB.SPECIFIEDVALUESPointersEncoding
		specifiedvaluesAPIs = append(specifiedvaluesAPIs, specifiedvaluesAPI)
	}

	c.JSON(http.StatusOK, specifiedvaluesAPIs)
}

// PostSPECIFIEDVALUES
//
// swagger:route POST /specifiedvaluess specifiedvaluess postSPECIFIEDVALUES
//
// Creates a specifiedvalues
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECIFIEDVALUES(c *gin.Context) {

	mutexSPECIFIEDVALUES.Lock()
	defer mutexSPECIFIEDVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECIFIEDVALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFIEDVALUES.GetDB()

	// Validate input
	var input orm.SPECIFIEDVALUESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specifiedvalues
	specifiedvaluesDB := orm.SPECIFIEDVALUESDB{}
	specifiedvaluesDB.SPECIFIEDVALUESPointersEncoding = input.SPECIFIEDVALUESPointersEncoding
	specifiedvaluesDB.CopyBasicFieldsFromSPECIFIEDVALUES_WOP(&input.SPECIFIEDVALUES_WOP)

	query := db.Create(&specifiedvaluesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECIFIEDVALUES.CheckoutPhaseOneInstance(&specifiedvaluesDB)
	specifiedvalues := backRepo.BackRepoSPECIFIEDVALUES.Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESPtr[specifiedvaluesDB.ID]

	if specifiedvalues != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specifiedvalues)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specifiedvaluesDB)
}

// GetSPECIFIEDVALUES
//
// swagger:route GET /specifiedvaluess/{ID} specifiedvaluess getSPECIFIEDVALUES
//
// Gets the details for a specifiedvalues.
//
// Responses:
// default: genericError
//
//	200: specifiedvaluesDBResponse
func (controller *Controller) GetSPECIFIEDVALUES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFIEDVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFIEDVALUES.GetDB()

	// Get specifiedvaluesDB in DB
	var specifiedvaluesDB orm.SPECIFIEDVALUESDB
	if err := db.First(&specifiedvaluesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specifiedvaluesAPI orm.SPECIFIEDVALUESAPI
	specifiedvaluesAPI.ID = specifiedvaluesDB.ID
	specifiedvaluesAPI.SPECIFIEDVALUESPointersEncoding = specifiedvaluesDB.SPECIFIEDVALUESPointersEncoding
	specifiedvaluesDB.CopyBasicFieldsToSPECIFIEDVALUES_WOP(&specifiedvaluesAPI.SPECIFIEDVALUES_WOP)

	c.JSON(http.StatusOK, specifiedvaluesAPI)
}

// UpdateSPECIFIEDVALUES
//
// swagger:route PATCH /specifiedvaluess/{ID} specifiedvaluess updateSPECIFIEDVALUES
//
// # Update a specifiedvalues
//
// Responses:
// default: genericError
//
//	200: specifiedvaluesDBResponse
func (controller *Controller) UpdateSPECIFIEDVALUES(c *gin.Context) {

	mutexSPECIFIEDVALUES.Lock()
	defer mutexSPECIFIEDVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECIFIEDVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFIEDVALUES.GetDB()

	// Validate input
	var input orm.SPECIFIEDVALUESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specifiedvaluesDB orm.SPECIFIEDVALUESDB

	// fetch the specifiedvalues
	query := db.First(&specifiedvaluesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specifiedvaluesDB.CopyBasicFieldsFromSPECIFIEDVALUES_WOP(&input.SPECIFIEDVALUES_WOP)
	specifiedvaluesDB.SPECIFIEDVALUESPointersEncoding = input.SPECIFIEDVALUESPointersEncoding

	query = db.Model(&specifiedvaluesDB).Updates(specifiedvaluesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specifiedvaluesNew := new(models.SPECIFIEDVALUES)
	specifiedvaluesDB.CopyBasicFieldsToSPECIFIEDVALUES(specifiedvaluesNew)

	// redeem pointers
	specifiedvaluesDB.DecodePointers(backRepo, specifiedvaluesNew)

	// get stage instance from DB instance, and call callback function
	specifiedvaluesOld := backRepo.BackRepoSPECIFIEDVALUES.Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESPtr[specifiedvaluesDB.ID]
	if specifiedvaluesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specifiedvaluesOld, specifiedvaluesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specifiedvaluesDB
	c.JSON(http.StatusOK, specifiedvaluesDB)
}

// DeleteSPECIFIEDVALUES
//
// swagger:route DELETE /specifiedvaluess/{ID} specifiedvaluess deleteSPECIFIEDVALUES
//
// # Delete a specifiedvalues
//
// default: genericError
//
//	200: specifiedvaluesDBResponse
func (controller *Controller) DeleteSPECIFIEDVALUES(c *gin.Context) {

	mutexSPECIFIEDVALUES.Lock()
	defer mutexSPECIFIEDVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECIFIEDVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFIEDVALUES.GetDB()

	// Get model if exist
	var specifiedvaluesDB orm.SPECIFIEDVALUESDB
	if err := db.First(&specifiedvaluesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specifiedvaluesDB)

	// get an instance (not staged) from DB instance, and call callback function
	specifiedvaluesDeleted := new(models.SPECIFIEDVALUES)
	specifiedvaluesDB.CopyBasicFieldsToSPECIFIEDVALUES(specifiedvaluesDeleted)

	// get stage instance from DB instance, and call callback function
	specifiedvaluesStaged := backRepo.BackRepoSPECIFIEDVALUES.Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESPtr[specifiedvaluesDB.ID]
	if specifiedvaluesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specifiedvaluesStaged, specifiedvaluesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
