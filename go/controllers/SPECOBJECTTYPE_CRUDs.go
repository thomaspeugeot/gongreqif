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
var __SPECOBJECTTYPE__dummysDeclaration__ models.SPECOBJECTTYPE
var __SPECOBJECTTYPE_time__dummyDeclaration time.Duration

var mutexSPECOBJECTTYPE sync.Mutex

// An SPECOBJECTTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECOBJECTTYPE updateSPECOBJECTTYPE deleteSPECOBJECTTYPE
type SPECOBJECTTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECOBJECTTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECOBJECTTYPE updateSPECOBJECTTYPE
type SPECOBJECTTYPEInput struct {
	// The SPECOBJECTTYPE to submit or modify
	// in: body
	SPECOBJECTTYPE *orm.SPECOBJECTTYPEAPI
}

// GetSPECOBJECTTYPEs
//
// swagger:route GET /specobjecttypes specobjecttypes getSPECOBJECTTYPEs
//
// # Get all specobjecttypes
//
// Responses:
// default: genericError
//
//	200: specobjecttypeDBResponse
func (controller *Controller) GetSPECOBJECTTYPEs(c *gin.Context) {

	// source slice
	var specobjecttypeDBs []orm.SPECOBJECTTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECTTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTTYPE.GetDB()

	query := db.Find(&specobjecttypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specobjecttypeAPIs := make([]orm.SPECOBJECTTYPEAPI, 0)

	// for each specobjecttype, update fields from the database nullable fields
	for idx := range specobjecttypeDBs {
		specobjecttypeDB := &specobjecttypeDBs[idx]
		_ = specobjecttypeDB
		var specobjecttypeAPI orm.SPECOBJECTTYPEAPI

		// insertion point for updating fields
		specobjecttypeAPI.ID = specobjecttypeDB.ID
		specobjecttypeDB.CopyBasicFieldsToSPECOBJECTTYPE_WOP(&specobjecttypeAPI.SPECOBJECTTYPE_WOP)
		specobjecttypeAPI.SPECOBJECTTYPEPointersEncoding = specobjecttypeDB.SPECOBJECTTYPEPointersEncoding
		specobjecttypeAPIs = append(specobjecttypeAPIs, specobjecttypeAPI)
	}

	c.JSON(http.StatusOK, specobjecttypeAPIs)
}

// PostSPECOBJECTTYPE
//
// swagger:route POST /specobjecttypes specobjecttypes postSPECOBJECTTYPE
//
// Creates a specobjecttype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECOBJECTTYPE(c *gin.Context) {

	mutexSPECOBJECTTYPE.Lock()
	defer mutexSPECOBJECTTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECOBJECTTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTTYPE.GetDB()

	// Validate input
	var input orm.SPECOBJECTTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specobjecttype
	specobjecttypeDB := orm.SPECOBJECTTYPEDB{}
	specobjecttypeDB.SPECOBJECTTYPEPointersEncoding = input.SPECOBJECTTYPEPointersEncoding
	specobjecttypeDB.CopyBasicFieldsFromSPECOBJECTTYPE_WOP(&input.SPECOBJECTTYPE_WOP)

	query := db.Create(&specobjecttypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECOBJECTTYPE.CheckoutPhaseOneInstance(&specobjecttypeDB)
	specobjecttype := backRepo.BackRepoSPECOBJECTTYPE.Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEPtr[specobjecttypeDB.ID]

	if specobjecttype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specobjecttype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specobjecttypeDB)
}

// GetSPECOBJECTTYPE
//
// swagger:route GET /specobjecttypes/{ID} specobjecttypes getSPECOBJECTTYPE
//
// Gets the details for a specobjecttype.
//
// Responses:
// default: genericError
//
//	200: specobjecttypeDBResponse
func (controller *Controller) GetSPECOBJECTTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECTTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTTYPE.GetDB()

	// Get specobjecttypeDB in DB
	var specobjecttypeDB orm.SPECOBJECTTYPEDB
	if err := db.First(&specobjecttypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specobjecttypeAPI orm.SPECOBJECTTYPEAPI
	specobjecttypeAPI.ID = specobjecttypeDB.ID
	specobjecttypeAPI.SPECOBJECTTYPEPointersEncoding = specobjecttypeDB.SPECOBJECTTYPEPointersEncoding
	specobjecttypeDB.CopyBasicFieldsToSPECOBJECTTYPE_WOP(&specobjecttypeAPI.SPECOBJECTTYPE_WOP)

	c.JSON(http.StatusOK, specobjecttypeAPI)
}

// UpdateSPECOBJECTTYPE
//
// swagger:route PATCH /specobjecttypes/{ID} specobjecttypes updateSPECOBJECTTYPE
//
// # Update a specobjecttype
//
// Responses:
// default: genericError
//
//	200: specobjecttypeDBResponse
func (controller *Controller) UpdateSPECOBJECTTYPE(c *gin.Context) {

	mutexSPECOBJECTTYPE.Lock()
	defer mutexSPECOBJECTTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECOBJECTTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTTYPE.GetDB()

	// Validate input
	var input orm.SPECOBJECTTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specobjecttypeDB orm.SPECOBJECTTYPEDB

	// fetch the specobjecttype
	query := db.First(&specobjecttypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specobjecttypeDB.CopyBasicFieldsFromSPECOBJECTTYPE_WOP(&input.SPECOBJECTTYPE_WOP)
	specobjecttypeDB.SPECOBJECTTYPEPointersEncoding = input.SPECOBJECTTYPEPointersEncoding

	query = db.Model(&specobjecttypeDB).Updates(specobjecttypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specobjecttypeNew := new(models.SPECOBJECTTYPE)
	specobjecttypeDB.CopyBasicFieldsToSPECOBJECTTYPE(specobjecttypeNew)

	// redeem pointers
	specobjecttypeDB.DecodePointers(backRepo, specobjecttypeNew)

	// get stage instance from DB instance, and call callback function
	specobjecttypeOld := backRepo.BackRepoSPECOBJECTTYPE.Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEPtr[specobjecttypeDB.ID]
	if specobjecttypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specobjecttypeOld, specobjecttypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specobjecttypeDB
	c.JSON(http.StatusOK, specobjecttypeDB)
}

// DeleteSPECOBJECTTYPE
//
// swagger:route DELETE /specobjecttypes/{ID} specobjecttypes deleteSPECOBJECTTYPE
//
// # Delete a specobjecttype
//
// default: genericError
//
//	200: specobjecttypeDBResponse
func (controller *Controller) DeleteSPECOBJECTTYPE(c *gin.Context) {

	mutexSPECOBJECTTYPE.Lock()
	defer mutexSPECOBJECTTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECOBJECTTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECTTYPE.GetDB()

	// Get model if exist
	var specobjecttypeDB orm.SPECOBJECTTYPEDB
	if err := db.First(&specobjecttypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specobjecttypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	specobjecttypeDeleted := new(models.SPECOBJECTTYPE)
	specobjecttypeDB.CopyBasicFieldsToSPECOBJECTTYPE(specobjecttypeDeleted)

	// get stage instance from DB instance, and call callback function
	specobjecttypeStaged := backRepo.BackRepoSPECOBJECTTYPE.Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEPtr[specobjecttypeDB.ID]
	if specobjecttypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specobjecttypeStaged, specobjecttypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
