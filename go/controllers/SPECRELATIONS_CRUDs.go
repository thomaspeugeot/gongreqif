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
var __SPECRELATIONS__dummysDeclaration__ models.SPECRELATIONS
var __SPECRELATIONS_time__dummyDeclaration time.Duration

var mutexSPECRELATIONS sync.Mutex

// An SPECRELATIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECRELATIONS updateSPECRELATIONS deleteSPECRELATIONS
type SPECRELATIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECRELATIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECRELATIONS updateSPECRELATIONS
type SPECRELATIONSInput struct {
	// The SPECRELATIONS to submit or modify
	// in: body
	SPECRELATIONS *orm.SPECRELATIONSAPI
}

// GetSPECRELATIONSs
//
// swagger:route GET /specrelationss specrelationss getSPECRELATIONSs
//
// # Get all specrelationss
//
// Responses:
// default: genericError
//
//	200: specrelationsDBResponse
func (controller *Controller) GetSPECRELATIONSs(c *gin.Context) {

	// source slice
	var specrelationsDBs []orm.SPECRELATIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONS.GetDB()

	query := db.Find(&specrelationsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specrelationsAPIs := make([]orm.SPECRELATIONSAPI, 0)

	// for each specrelations, update fields from the database nullable fields
	for idx := range specrelationsDBs {
		specrelationsDB := &specrelationsDBs[idx]
		_ = specrelationsDB
		var specrelationsAPI orm.SPECRELATIONSAPI

		// insertion point for updating fields
		specrelationsAPI.ID = specrelationsDB.ID
		specrelationsDB.CopyBasicFieldsToSPECRELATIONS_WOP(&specrelationsAPI.SPECRELATIONS_WOP)
		specrelationsAPI.SPECRELATIONSPointersEncoding = specrelationsDB.SPECRELATIONSPointersEncoding
		specrelationsAPIs = append(specrelationsAPIs, specrelationsAPI)
	}

	c.JSON(http.StatusOK, specrelationsAPIs)
}

// PostSPECRELATIONS
//
// swagger:route POST /specrelationss specrelationss postSPECRELATIONS
//
// Creates a specrelations
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECRELATIONS(c *gin.Context) {

	mutexSPECRELATIONS.Lock()
	defer mutexSPECRELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECRELATIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONS.GetDB()

	// Validate input
	var input orm.SPECRELATIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specrelations
	specrelationsDB := orm.SPECRELATIONSDB{}
	specrelationsDB.SPECRELATIONSPointersEncoding = input.SPECRELATIONSPointersEncoding
	specrelationsDB.CopyBasicFieldsFromSPECRELATIONS_WOP(&input.SPECRELATIONS_WOP)

	query := db.Create(&specrelationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECRELATIONS.CheckoutPhaseOneInstance(&specrelationsDB)
	specrelations := backRepo.BackRepoSPECRELATIONS.Map_SPECRELATIONSDBID_SPECRELATIONSPtr[specrelationsDB.ID]

	if specrelations != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specrelations)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specrelationsDB)
}

// GetSPECRELATIONS
//
// swagger:route GET /specrelationss/{ID} specrelationss getSPECRELATIONS
//
// Gets the details for a specrelations.
//
// Responses:
// default: genericError
//
//	200: specrelationsDBResponse
func (controller *Controller) GetSPECRELATIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONS.GetDB()

	// Get specrelationsDB in DB
	var specrelationsDB orm.SPECRELATIONSDB
	if err := db.First(&specrelationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specrelationsAPI orm.SPECRELATIONSAPI
	specrelationsAPI.ID = specrelationsDB.ID
	specrelationsAPI.SPECRELATIONSPointersEncoding = specrelationsDB.SPECRELATIONSPointersEncoding
	specrelationsDB.CopyBasicFieldsToSPECRELATIONS_WOP(&specrelationsAPI.SPECRELATIONS_WOP)

	c.JSON(http.StatusOK, specrelationsAPI)
}

// UpdateSPECRELATIONS
//
// swagger:route PATCH /specrelationss/{ID} specrelationss updateSPECRELATIONS
//
// # Update a specrelations
//
// Responses:
// default: genericError
//
//	200: specrelationsDBResponse
func (controller *Controller) UpdateSPECRELATIONS(c *gin.Context) {

	mutexSPECRELATIONS.Lock()
	defer mutexSPECRELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECRELATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONS.GetDB()

	// Validate input
	var input orm.SPECRELATIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specrelationsDB orm.SPECRELATIONSDB

	// fetch the specrelations
	query := db.First(&specrelationsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specrelationsDB.CopyBasicFieldsFromSPECRELATIONS_WOP(&input.SPECRELATIONS_WOP)
	specrelationsDB.SPECRELATIONSPointersEncoding = input.SPECRELATIONSPointersEncoding

	query = db.Model(&specrelationsDB).Updates(specrelationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specrelationsNew := new(models.SPECRELATIONS)
	specrelationsDB.CopyBasicFieldsToSPECRELATIONS(specrelationsNew)

	// redeem pointers
	specrelationsDB.DecodePointers(backRepo, specrelationsNew)

	// get stage instance from DB instance, and call callback function
	specrelationsOld := backRepo.BackRepoSPECRELATIONS.Map_SPECRELATIONSDBID_SPECRELATIONSPtr[specrelationsDB.ID]
	if specrelationsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specrelationsOld, specrelationsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specrelationsDB
	c.JSON(http.StatusOK, specrelationsDB)
}

// DeleteSPECRELATIONS
//
// swagger:route DELETE /specrelationss/{ID} specrelationss deleteSPECRELATIONS
//
// # Delete a specrelations
//
// default: genericError
//
//	200: specrelationsDBResponse
func (controller *Controller) DeleteSPECRELATIONS(c *gin.Context) {

	mutexSPECRELATIONS.Lock()
	defer mutexSPECRELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECRELATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONS.GetDB()

	// Get model if exist
	var specrelationsDB orm.SPECRELATIONSDB
	if err := db.First(&specrelationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specrelationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	specrelationsDeleted := new(models.SPECRELATIONS)
	specrelationsDB.CopyBasicFieldsToSPECRELATIONS(specrelationsDeleted)

	// get stage instance from DB instance, and call callback function
	specrelationsStaged := backRepo.BackRepoSPECRELATIONS.Map_SPECRELATIONSDBID_SPECRELATIONSPtr[specrelationsDB.ID]
	if specrelationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specrelationsStaged, specrelationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
