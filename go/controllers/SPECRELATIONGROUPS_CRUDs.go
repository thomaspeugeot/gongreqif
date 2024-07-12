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
var __SPECRELATIONGROUPS__dummysDeclaration__ models.SPECRELATIONGROUPS
var __SPECRELATIONGROUPS_time__dummyDeclaration time.Duration

var mutexSPECRELATIONGROUPS sync.Mutex

// An SPECRELATIONGROUPSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECRELATIONGROUPS updateSPECRELATIONGROUPS deleteSPECRELATIONGROUPS
type SPECRELATIONGROUPSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECRELATIONGROUPSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECRELATIONGROUPS updateSPECRELATIONGROUPS
type SPECRELATIONGROUPSInput struct {
	// The SPECRELATIONGROUPS to submit or modify
	// in: body
	SPECRELATIONGROUPS *orm.SPECRELATIONGROUPSAPI
}

// GetSPECRELATIONGROUPSs
//
// swagger:route GET /specrelationgroupss specrelationgroupss getSPECRELATIONGROUPSs
//
// # Get all specrelationgroupss
//
// Responses:
// default: genericError
//
//	200: specrelationgroupsDBResponse
func (controller *Controller) GetSPECRELATIONGROUPSs(c *gin.Context) {

	// source slice
	var specrelationgroupsDBs []orm.SPECRELATIONGROUPSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONGROUPSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONGROUPS.GetDB()

	query := db.Find(&specrelationgroupsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specrelationgroupsAPIs := make([]orm.SPECRELATIONGROUPSAPI, 0)

	// for each specrelationgroups, update fields from the database nullable fields
	for idx := range specrelationgroupsDBs {
		specrelationgroupsDB := &specrelationgroupsDBs[idx]
		_ = specrelationgroupsDB
		var specrelationgroupsAPI orm.SPECRELATIONGROUPSAPI

		// insertion point for updating fields
		specrelationgroupsAPI.ID = specrelationgroupsDB.ID
		specrelationgroupsDB.CopyBasicFieldsToSPECRELATIONGROUPS_WOP(&specrelationgroupsAPI.SPECRELATIONGROUPS_WOP)
		specrelationgroupsAPI.SPECRELATIONGROUPSPointersEncoding = specrelationgroupsDB.SPECRELATIONGROUPSPointersEncoding
		specrelationgroupsAPIs = append(specrelationgroupsAPIs, specrelationgroupsAPI)
	}

	c.JSON(http.StatusOK, specrelationgroupsAPIs)
}

// PostSPECRELATIONGROUPS
//
// swagger:route POST /specrelationgroupss specrelationgroupss postSPECRELATIONGROUPS
//
// Creates a specrelationgroups
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECRELATIONGROUPS(c *gin.Context) {

	mutexSPECRELATIONGROUPS.Lock()
	defer mutexSPECRELATIONGROUPS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECRELATIONGROUPSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONGROUPS.GetDB()

	// Validate input
	var input orm.SPECRELATIONGROUPSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specrelationgroups
	specrelationgroupsDB := orm.SPECRELATIONGROUPSDB{}
	specrelationgroupsDB.SPECRELATIONGROUPSPointersEncoding = input.SPECRELATIONGROUPSPointersEncoding
	specrelationgroupsDB.CopyBasicFieldsFromSPECRELATIONGROUPS_WOP(&input.SPECRELATIONGROUPS_WOP)

	query := db.Create(&specrelationgroupsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECRELATIONGROUPS.CheckoutPhaseOneInstance(&specrelationgroupsDB)
	specrelationgroups := backRepo.BackRepoSPECRELATIONGROUPS.Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSPtr[specrelationgroupsDB.ID]

	if specrelationgroups != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specrelationgroups)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specrelationgroupsDB)
}

// GetSPECRELATIONGROUPS
//
// swagger:route GET /specrelationgroupss/{ID} specrelationgroupss getSPECRELATIONGROUPS
//
// Gets the details for a specrelationgroups.
//
// Responses:
// default: genericError
//
//	200: specrelationgroupsDBResponse
func (controller *Controller) GetSPECRELATIONGROUPS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONGROUPS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONGROUPS.GetDB()

	// Get specrelationgroupsDB in DB
	var specrelationgroupsDB orm.SPECRELATIONGROUPSDB
	if err := db.First(&specrelationgroupsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specrelationgroupsAPI orm.SPECRELATIONGROUPSAPI
	specrelationgroupsAPI.ID = specrelationgroupsDB.ID
	specrelationgroupsAPI.SPECRELATIONGROUPSPointersEncoding = specrelationgroupsDB.SPECRELATIONGROUPSPointersEncoding
	specrelationgroupsDB.CopyBasicFieldsToSPECRELATIONGROUPS_WOP(&specrelationgroupsAPI.SPECRELATIONGROUPS_WOP)

	c.JSON(http.StatusOK, specrelationgroupsAPI)
}

// UpdateSPECRELATIONGROUPS
//
// swagger:route PATCH /specrelationgroupss/{ID} specrelationgroupss updateSPECRELATIONGROUPS
//
// # Update a specrelationgroups
//
// Responses:
// default: genericError
//
//	200: specrelationgroupsDBResponse
func (controller *Controller) UpdateSPECRELATIONGROUPS(c *gin.Context) {

	mutexSPECRELATIONGROUPS.Lock()
	defer mutexSPECRELATIONGROUPS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECRELATIONGROUPS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONGROUPS.GetDB()

	// Validate input
	var input orm.SPECRELATIONGROUPSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specrelationgroupsDB orm.SPECRELATIONGROUPSDB

	// fetch the specrelationgroups
	query := db.First(&specrelationgroupsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specrelationgroupsDB.CopyBasicFieldsFromSPECRELATIONGROUPS_WOP(&input.SPECRELATIONGROUPS_WOP)
	specrelationgroupsDB.SPECRELATIONGROUPSPointersEncoding = input.SPECRELATIONGROUPSPointersEncoding

	query = db.Model(&specrelationgroupsDB).Updates(specrelationgroupsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specrelationgroupsNew := new(models.SPECRELATIONGROUPS)
	specrelationgroupsDB.CopyBasicFieldsToSPECRELATIONGROUPS(specrelationgroupsNew)

	// redeem pointers
	specrelationgroupsDB.DecodePointers(backRepo, specrelationgroupsNew)

	// get stage instance from DB instance, and call callback function
	specrelationgroupsOld := backRepo.BackRepoSPECRELATIONGROUPS.Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSPtr[specrelationgroupsDB.ID]
	if specrelationgroupsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specrelationgroupsOld, specrelationgroupsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specrelationgroupsDB
	c.JSON(http.StatusOK, specrelationgroupsDB)
}

// DeleteSPECRELATIONGROUPS
//
// swagger:route DELETE /specrelationgroupss/{ID} specrelationgroupss deleteSPECRELATIONGROUPS
//
// # Delete a specrelationgroups
//
// default: genericError
//
//	200: specrelationgroupsDBResponse
func (controller *Controller) DeleteSPECRELATIONGROUPS(c *gin.Context) {

	mutexSPECRELATIONGROUPS.Lock()
	defer mutexSPECRELATIONGROUPS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECRELATIONGROUPS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONGROUPS.GetDB()

	// Get model if exist
	var specrelationgroupsDB orm.SPECRELATIONGROUPSDB
	if err := db.First(&specrelationgroupsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specrelationgroupsDB)

	// get an instance (not staged) from DB instance, and call callback function
	specrelationgroupsDeleted := new(models.SPECRELATIONGROUPS)
	specrelationgroupsDB.CopyBasicFieldsToSPECRELATIONGROUPS(specrelationgroupsDeleted)

	// get stage instance from DB instance, and call callback function
	specrelationgroupsStaged := backRepo.BackRepoSPECRELATIONGROUPS.Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSPtr[specrelationgroupsDB.ID]
	if specrelationgroupsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specrelationgroupsStaged, specrelationgroupsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
