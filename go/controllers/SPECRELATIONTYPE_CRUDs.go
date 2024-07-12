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
var __SPECRELATIONTYPE__dummysDeclaration__ models.SPECRELATIONTYPE
var __SPECRELATIONTYPE_time__dummyDeclaration time.Duration

var mutexSPECRELATIONTYPE sync.Mutex

// An SPECRELATIONTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECRELATIONTYPE updateSPECRELATIONTYPE deleteSPECRELATIONTYPE
type SPECRELATIONTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECRELATIONTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECRELATIONTYPE updateSPECRELATIONTYPE
type SPECRELATIONTYPEInput struct {
	// The SPECRELATIONTYPE to submit or modify
	// in: body
	SPECRELATIONTYPE *orm.SPECRELATIONTYPEAPI
}

// GetSPECRELATIONTYPEs
//
// swagger:route GET /specrelationtypes specrelationtypes getSPECRELATIONTYPEs
//
// # Get all specrelationtypes
//
// Responses:
// default: genericError
//
//	200: specrelationtypeDBResponse
func (controller *Controller) GetSPECRELATIONTYPEs(c *gin.Context) {

	// source slice
	var specrelationtypeDBs []orm.SPECRELATIONTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONTYPE.GetDB()

	query := db.Find(&specrelationtypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specrelationtypeAPIs := make([]orm.SPECRELATIONTYPEAPI, 0)

	// for each specrelationtype, update fields from the database nullable fields
	for idx := range specrelationtypeDBs {
		specrelationtypeDB := &specrelationtypeDBs[idx]
		_ = specrelationtypeDB
		var specrelationtypeAPI orm.SPECRELATIONTYPEAPI

		// insertion point for updating fields
		specrelationtypeAPI.ID = specrelationtypeDB.ID
		specrelationtypeDB.CopyBasicFieldsToSPECRELATIONTYPE_WOP(&specrelationtypeAPI.SPECRELATIONTYPE_WOP)
		specrelationtypeAPI.SPECRELATIONTYPEPointersEncoding = specrelationtypeDB.SPECRELATIONTYPEPointersEncoding
		specrelationtypeAPIs = append(specrelationtypeAPIs, specrelationtypeAPI)
	}

	c.JSON(http.StatusOK, specrelationtypeAPIs)
}

// PostSPECRELATIONTYPE
//
// swagger:route POST /specrelationtypes specrelationtypes postSPECRELATIONTYPE
//
// Creates a specrelationtype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECRELATIONTYPE(c *gin.Context) {

	mutexSPECRELATIONTYPE.Lock()
	defer mutexSPECRELATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECRELATIONTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONTYPE.GetDB()

	// Validate input
	var input orm.SPECRELATIONTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specrelationtype
	specrelationtypeDB := orm.SPECRELATIONTYPEDB{}
	specrelationtypeDB.SPECRELATIONTYPEPointersEncoding = input.SPECRELATIONTYPEPointersEncoding
	specrelationtypeDB.CopyBasicFieldsFromSPECRELATIONTYPE_WOP(&input.SPECRELATIONTYPE_WOP)

	query := db.Create(&specrelationtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECRELATIONTYPE.CheckoutPhaseOneInstance(&specrelationtypeDB)
	specrelationtype := backRepo.BackRepoSPECRELATIONTYPE.Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEPtr[specrelationtypeDB.ID]

	if specrelationtype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specrelationtype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specrelationtypeDB)
}

// GetSPECRELATIONTYPE
//
// swagger:route GET /specrelationtypes/{ID} specrelationtypes getSPECRELATIONTYPE
//
// Gets the details for a specrelationtype.
//
// Responses:
// default: genericError
//
//	200: specrelationtypeDBResponse
func (controller *Controller) GetSPECRELATIONTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECRELATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONTYPE.GetDB()

	// Get specrelationtypeDB in DB
	var specrelationtypeDB orm.SPECRELATIONTYPEDB
	if err := db.First(&specrelationtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specrelationtypeAPI orm.SPECRELATIONTYPEAPI
	specrelationtypeAPI.ID = specrelationtypeDB.ID
	specrelationtypeAPI.SPECRELATIONTYPEPointersEncoding = specrelationtypeDB.SPECRELATIONTYPEPointersEncoding
	specrelationtypeDB.CopyBasicFieldsToSPECRELATIONTYPE_WOP(&specrelationtypeAPI.SPECRELATIONTYPE_WOP)

	c.JSON(http.StatusOK, specrelationtypeAPI)
}

// UpdateSPECRELATIONTYPE
//
// swagger:route PATCH /specrelationtypes/{ID} specrelationtypes updateSPECRELATIONTYPE
//
// # Update a specrelationtype
//
// Responses:
// default: genericError
//
//	200: specrelationtypeDBResponse
func (controller *Controller) UpdateSPECRELATIONTYPE(c *gin.Context) {

	mutexSPECRELATIONTYPE.Lock()
	defer mutexSPECRELATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECRELATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONTYPE.GetDB()

	// Validate input
	var input orm.SPECRELATIONTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specrelationtypeDB orm.SPECRELATIONTYPEDB

	// fetch the specrelationtype
	query := db.First(&specrelationtypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specrelationtypeDB.CopyBasicFieldsFromSPECRELATIONTYPE_WOP(&input.SPECRELATIONTYPE_WOP)
	specrelationtypeDB.SPECRELATIONTYPEPointersEncoding = input.SPECRELATIONTYPEPointersEncoding

	query = db.Model(&specrelationtypeDB).Updates(specrelationtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specrelationtypeNew := new(models.SPECRELATIONTYPE)
	specrelationtypeDB.CopyBasicFieldsToSPECRELATIONTYPE(specrelationtypeNew)

	// redeem pointers
	specrelationtypeDB.DecodePointers(backRepo, specrelationtypeNew)

	// get stage instance from DB instance, and call callback function
	specrelationtypeOld := backRepo.BackRepoSPECRELATIONTYPE.Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEPtr[specrelationtypeDB.ID]
	if specrelationtypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specrelationtypeOld, specrelationtypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specrelationtypeDB
	c.JSON(http.StatusOK, specrelationtypeDB)
}

// DeleteSPECRELATIONTYPE
//
// swagger:route DELETE /specrelationtypes/{ID} specrelationtypes deleteSPECRELATIONTYPE
//
// # Delete a specrelationtype
//
// default: genericError
//
//	200: specrelationtypeDBResponse
func (controller *Controller) DeleteSPECRELATIONTYPE(c *gin.Context) {

	mutexSPECRELATIONTYPE.Lock()
	defer mutexSPECRELATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECRELATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECRELATIONTYPE.GetDB()

	// Get model if exist
	var specrelationtypeDB orm.SPECRELATIONTYPEDB
	if err := db.First(&specrelationtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specrelationtypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	specrelationtypeDeleted := new(models.SPECRELATIONTYPE)
	specrelationtypeDB.CopyBasicFieldsToSPECRELATIONTYPE(specrelationtypeDeleted)

	// get stage instance from DB instance, and call callback function
	specrelationtypeStaged := backRepo.BackRepoSPECRELATIONTYPE.Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEPtr[specrelationtypeDB.ID]
	if specrelationtypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specrelationtypeStaged, specrelationtypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
