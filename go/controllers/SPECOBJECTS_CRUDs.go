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
var __SPECOBJECTS__dummysDeclaration__ models.SPECOBJECTS
var __SPECOBJECTS_time__dummyDeclaration time.Duration

var mutexSPECOBJECTS sync.Mutex

// An SPECOBJECTSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECOBJECTS updateSPECOBJECTS deleteSPECOBJECTS
type SPECOBJECTSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECOBJECTSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECOBJECTS updateSPECOBJECTS
type SPECOBJECTSInput struct {
	// The SPECOBJECTS to submit or modify
	// in: body
	SPECOBJECTS *orm.SPECOBJECTSAPI
}

// GetSPECOBJECTSs
//
// swagger:route GET /specobjectss specobjectss getSPECOBJECTSs
//
// # Get all specobjectss
//
// Responses:
// default: genericError
//
//	200: specobjectsDBResponse
func (controller *Controller) GetSPECOBJECTSs(c *gin.Context) {

	// source slice
	var specobjectsDBs []orm.SPECOBJECTSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECTSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTS.GetDB()

	query := db.Find(&specobjectsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specobjectsAPIs := make([]orm.SPECOBJECTSAPI, 0)

	// for each specobjects, update fields from the database nullable fields
	for idx := range specobjectsDBs {
		specobjectsDB := &specobjectsDBs[idx]
		_ = specobjectsDB
		var specobjectsAPI orm.SPECOBJECTSAPI

		// insertion point for updating fields
		specobjectsAPI.ID = specobjectsDB.ID
		specobjectsDB.CopyBasicFieldsToSPECOBJECTS_WOP(&specobjectsAPI.SPECOBJECTS_WOP)
		specobjectsAPI.SPECOBJECTSPointersEncoding = specobjectsDB.SPECOBJECTSPointersEncoding
		specobjectsAPIs = append(specobjectsAPIs, specobjectsAPI)
	}

	c.JSON(http.StatusOK, specobjectsAPIs)
}

// PostSPECOBJECTS
//
// swagger:route POST /specobjectss specobjectss postSPECOBJECTS
//
// Creates a specobjects
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECOBJECTS(c *gin.Context) {

	mutexSPECOBJECTS.Lock()
	defer mutexSPECOBJECTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECOBJECTSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTS.GetDB()

	// Validate input
	var input orm.SPECOBJECTSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specobjects
	specobjectsDB := orm.SPECOBJECTSDB{}
	specobjectsDB.SPECOBJECTSPointersEncoding = input.SPECOBJECTSPointersEncoding
	specobjectsDB.CopyBasicFieldsFromSPECOBJECTS_WOP(&input.SPECOBJECTS_WOP)

	query := db.Create(&specobjectsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECOBJECTS.CheckoutPhaseOneInstance(&specobjectsDB)
	specobjects := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]

	if specobjects != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specobjects)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specobjectsDB)
}

// GetSPECOBJECTS
//
// swagger:route GET /specobjectss/{ID} specobjectss getSPECOBJECTS
//
// Gets the details for a specobjects.
//
// Responses:
// default: genericError
//
//	200: specobjectsDBResponse
func (controller *Controller) GetSPECOBJECTS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTS.GetDB()

	// Get specobjectsDB in DB
	var specobjectsDB orm.SPECOBJECTSDB
	if err := db.First(&specobjectsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specobjectsAPI orm.SPECOBJECTSAPI
	specobjectsAPI.ID = specobjectsDB.ID
	specobjectsAPI.SPECOBJECTSPointersEncoding = specobjectsDB.SPECOBJECTSPointersEncoding
	specobjectsDB.CopyBasicFieldsToSPECOBJECTS_WOP(&specobjectsAPI.SPECOBJECTS_WOP)

	c.JSON(http.StatusOK, specobjectsAPI)
}

// UpdateSPECOBJECTS
//
// swagger:route PATCH /specobjectss/{ID} specobjectss updateSPECOBJECTS
//
// # Update a specobjects
//
// Responses:
// default: genericError
//
//	200: specobjectsDBResponse
func (controller *Controller) UpdateSPECOBJECTS(c *gin.Context) {

	mutexSPECOBJECTS.Lock()
	defer mutexSPECOBJECTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECOBJECTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTS.GetDB()

	// Validate input
	var input orm.SPECOBJECTSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specobjectsDB orm.SPECOBJECTSDB

	// fetch the specobjects
	query := db.First(&specobjectsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specobjectsDB.CopyBasicFieldsFromSPECOBJECTS_WOP(&input.SPECOBJECTS_WOP)
	specobjectsDB.SPECOBJECTSPointersEncoding = input.SPECOBJECTSPointersEncoding

	query = db.Model(&specobjectsDB).Updates(specobjectsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specobjectsNew := new(models.SPECOBJECTS)
	specobjectsDB.CopyBasicFieldsToSPECOBJECTS(specobjectsNew)

	// redeem pointers
	specobjectsDB.DecodePointers(backRepo, specobjectsNew)

	// get stage instance from DB instance, and call callback function
	specobjectsOld := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]
	if specobjectsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specobjectsOld, specobjectsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specobjectsDB
	c.JSON(http.StatusOK, specobjectsDB)
}

// DeleteSPECOBJECTS
//
// swagger:route DELETE /specobjectss/{ID} specobjectss deleteSPECOBJECTS
//
// # Delete a specobjects
//
// default: genericError
//
//	200: specobjectsDBResponse
func (controller *Controller) DeleteSPECOBJECTS(c *gin.Context) {

	mutexSPECOBJECTS.Lock()
	defer mutexSPECOBJECTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECOBJECTS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTS.GetDB()

	// Get model if exist
	var specobjectsDB orm.SPECOBJECTSDB
	if err := db.First(&specobjectsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specobjectsDB)

	// get an instance (not staged) from DB instance, and call callback function
	specobjectsDeleted := new(models.SPECOBJECTS)
	specobjectsDB.CopyBasicFieldsToSPECOBJECTS(specobjectsDeleted)

	// get stage instance from DB instance, and call callback function
	specobjectsStaged := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]
	if specobjectsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specobjectsStaged, specobjectsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
