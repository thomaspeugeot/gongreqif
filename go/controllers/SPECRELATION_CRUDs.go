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
var __SPECRELATION__dummysDeclaration__ models.SPECRELATION
var __SPECRELATION_time__dummyDeclaration time.Duration

var mutexSPECRELATION sync.Mutex

// An SPECRELATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECRELATION updateSPECRELATION deleteSPECRELATION
type SPECRELATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECRELATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECRELATION updateSPECRELATION
type SPECRELATIONInput struct {
	// The SPECRELATION to submit or modify
	// in: body
	SPECRELATION *orm.SPECRELATIONAPI
}

// GetSPECRELATIONs
//
// swagger:route GET /specrelations specrelations getSPECRELATIONs
//
// # Get all specrelations
//
// Responses:
// default: genericError
//
//	200: specrelationDBResponse
func (controller *Controller) GetSPECRELATIONs(c *gin.Context) {

	// source slice
	var specrelationDBs []orm.SPECRELATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATION.GetDB()

	query := db.Find(&specrelationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specrelationAPIs := make([]orm.SPECRELATIONAPI, 0)

	// for each specrelation, update fields from the database nullable fields
	for idx := range specrelationDBs {
		specrelationDB := &specrelationDBs[idx]
		_ = specrelationDB
		var specrelationAPI orm.SPECRELATIONAPI

		// insertion point for updating fields
		specrelationAPI.ID = specrelationDB.ID
		specrelationDB.CopyBasicFieldsToSPECRELATION_WOP(&specrelationAPI.SPECRELATION_WOP)
		specrelationAPI.SPECRELATIONPointersEncoding = specrelationDB.SPECRELATIONPointersEncoding
		specrelationAPIs = append(specrelationAPIs, specrelationAPI)
	}

	c.JSON(http.StatusOK, specrelationAPIs)
}

// PostSPECRELATION
//
// swagger:route POST /specrelations specrelations postSPECRELATION
//
// Creates a specrelation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECRELATION(c *gin.Context) {

	mutexSPECRELATION.Lock()
	defer mutexSPECRELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECRELATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATION.GetDB()

	// Validate input
	var input orm.SPECRELATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specrelation
	specrelationDB := orm.SPECRELATIONDB{}
	specrelationDB.SPECRELATIONPointersEncoding = input.SPECRELATIONPointersEncoding
	specrelationDB.CopyBasicFieldsFromSPECRELATION_WOP(&input.SPECRELATION_WOP)

	query := db.Create(&specrelationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECRELATION.CheckoutPhaseOneInstance(&specrelationDB)
	specrelation := backRepo.BackRepoSPECRELATION.Map_SPECRELATIONDBID_SPECRELATIONPtr[specrelationDB.ID]

	if specrelation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specrelation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specrelationDB)
}

// GetSPECRELATION
//
// swagger:route GET /specrelations/{ID} specrelations getSPECRELATION
//
// Gets the details for a specrelation.
//
// Responses:
// default: genericError
//
//	200: specrelationDBResponse
func (controller *Controller) GetSPECRELATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATION.GetDB()

	// Get specrelationDB in DB
	var specrelationDB orm.SPECRELATIONDB
	if err := db.First(&specrelationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specrelationAPI orm.SPECRELATIONAPI
	specrelationAPI.ID = specrelationDB.ID
	specrelationAPI.SPECRELATIONPointersEncoding = specrelationDB.SPECRELATIONPointersEncoding
	specrelationDB.CopyBasicFieldsToSPECRELATION_WOP(&specrelationAPI.SPECRELATION_WOP)

	c.JSON(http.StatusOK, specrelationAPI)
}

// UpdateSPECRELATION
//
// swagger:route PATCH /specrelations/{ID} specrelations updateSPECRELATION
//
// # Update a specrelation
//
// Responses:
// default: genericError
//
//	200: specrelationDBResponse
func (controller *Controller) UpdateSPECRELATION(c *gin.Context) {

	mutexSPECRELATION.Lock()
	defer mutexSPECRELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECRELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATION.GetDB()

	// Validate input
	var input orm.SPECRELATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specrelationDB orm.SPECRELATIONDB

	// fetch the specrelation
	query := db.First(&specrelationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specrelationDB.CopyBasicFieldsFromSPECRELATION_WOP(&input.SPECRELATION_WOP)
	specrelationDB.SPECRELATIONPointersEncoding = input.SPECRELATIONPointersEncoding

	query = db.Model(&specrelationDB).Updates(specrelationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specrelationNew := new(models.SPECRELATION)
	specrelationDB.CopyBasicFieldsToSPECRELATION(specrelationNew)

	// redeem pointers
	specrelationDB.DecodePointers(backRepo, specrelationNew)

	// get stage instance from DB instance, and call callback function
	specrelationOld := backRepo.BackRepoSPECRELATION.Map_SPECRELATIONDBID_SPECRELATIONPtr[specrelationDB.ID]
	if specrelationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specrelationOld, specrelationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specrelationDB
	c.JSON(http.StatusOK, specrelationDB)
}

// DeleteSPECRELATION
//
// swagger:route DELETE /specrelations/{ID} specrelations deleteSPECRELATION
//
// # Delete a specrelation
//
// default: genericError
//
//	200: specrelationDBResponse
func (controller *Controller) DeleteSPECRELATION(c *gin.Context) {

	mutexSPECRELATION.Lock()
	defer mutexSPECRELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECRELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATION.GetDB()

	// Get model if exist
	var specrelationDB orm.SPECRELATIONDB
	if err := db.First(&specrelationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specrelationDB)

	// get an instance (not staged) from DB instance, and call callback function
	specrelationDeleted := new(models.SPECRELATION)
	specrelationDB.CopyBasicFieldsToSPECRELATION(specrelationDeleted)

	// get stage instance from DB instance, and call callback function
	specrelationStaged := backRepo.BackRepoSPECRELATION.Map_SPECRELATIONDBID_SPECRELATIONPtr[specrelationDB.ID]
	if specrelationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specrelationStaged, specrelationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
