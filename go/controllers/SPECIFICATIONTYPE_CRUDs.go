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
var __SPECIFICATIONTYPE__dummysDeclaration__ models.SPECIFICATIONTYPE
var __SPECIFICATIONTYPE_time__dummyDeclaration time.Duration

var mutexSPECIFICATIONTYPE sync.Mutex

// An SPECIFICATIONTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECIFICATIONTYPE updateSPECIFICATIONTYPE deleteSPECIFICATIONTYPE
type SPECIFICATIONTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECIFICATIONTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECIFICATIONTYPE updateSPECIFICATIONTYPE
type SPECIFICATIONTYPEInput struct {
	// The SPECIFICATIONTYPE to submit or modify
	// in: body
	SPECIFICATIONTYPE *orm.SPECIFICATIONTYPEAPI
}

// GetSPECIFICATIONTYPEs
//
// swagger:route GET /specificationtypes specificationtypes getSPECIFICATIONTYPEs
//
// # Get all specificationtypes
//
// Responses:
// default: genericError
//
//	200: specificationtypeDBResponse
func (controller *Controller) GetSPECIFICATIONTYPEs(c *gin.Context) {

	// source slice
	var specificationtypeDBs []orm.SPECIFICATIONTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATIONTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONTYPE.GetDB()

	query := db.Find(&specificationtypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specificationtypeAPIs := make([]orm.SPECIFICATIONTYPEAPI, 0)

	// for each specificationtype, update fields from the database nullable fields
	for idx := range specificationtypeDBs {
		specificationtypeDB := &specificationtypeDBs[idx]
		_ = specificationtypeDB
		var specificationtypeAPI orm.SPECIFICATIONTYPEAPI

		// insertion point for updating fields
		specificationtypeAPI.ID = specificationtypeDB.ID
		specificationtypeDB.CopyBasicFieldsToSPECIFICATIONTYPE_WOP(&specificationtypeAPI.SPECIFICATIONTYPE_WOP)
		specificationtypeAPI.SPECIFICATIONTYPEPointersEncoding = specificationtypeDB.SPECIFICATIONTYPEPointersEncoding
		specificationtypeAPIs = append(specificationtypeAPIs, specificationtypeAPI)
	}

	c.JSON(http.StatusOK, specificationtypeAPIs)
}

// PostSPECIFICATIONTYPE
//
// swagger:route POST /specificationtypes specificationtypes postSPECIFICATIONTYPE
//
// Creates a specificationtype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECIFICATIONTYPE(c *gin.Context) {

	mutexSPECIFICATIONTYPE.Lock()
	defer mutexSPECIFICATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECIFICATIONTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONTYPE.GetDB()

	// Validate input
	var input orm.SPECIFICATIONTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specificationtype
	specificationtypeDB := orm.SPECIFICATIONTYPEDB{}
	specificationtypeDB.SPECIFICATIONTYPEPointersEncoding = input.SPECIFICATIONTYPEPointersEncoding
	specificationtypeDB.CopyBasicFieldsFromSPECIFICATIONTYPE_WOP(&input.SPECIFICATIONTYPE_WOP)

	query := db.Create(&specificationtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECIFICATIONTYPE.CheckoutPhaseOneInstance(&specificationtypeDB)
	specificationtype := backRepo.BackRepoSPECIFICATIONTYPE.Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEPtr[specificationtypeDB.ID]

	if specificationtype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specificationtype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specificationtypeDB)
}

// GetSPECIFICATIONTYPE
//
// swagger:route GET /specificationtypes/{ID} specificationtypes getSPECIFICATIONTYPE
//
// Gets the details for a specificationtype.
//
// Responses:
// default: genericError
//
//	200: specificationtypeDBResponse
func (controller *Controller) GetSPECIFICATIONTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONTYPE.GetDB()

	// Get specificationtypeDB in DB
	var specificationtypeDB orm.SPECIFICATIONTYPEDB
	if err := db.First(&specificationtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specificationtypeAPI orm.SPECIFICATIONTYPEAPI
	specificationtypeAPI.ID = specificationtypeDB.ID
	specificationtypeAPI.SPECIFICATIONTYPEPointersEncoding = specificationtypeDB.SPECIFICATIONTYPEPointersEncoding
	specificationtypeDB.CopyBasicFieldsToSPECIFICATIONTYPE_WOP(&specificationtypeAPI.SPECIFICATIONTYPE_WOP)

	c.JSON(http.StatusOK, specificationtypeAPI)
}

// UpdateSPECIFICATIONTYPE
//
// swagger:route PATCH /specificationtypes/{ID} specificationtypes updateSPECIFICATIONTYPE
//
// # Update a specificationtype
//
// Responses:
// default: genericError
//
//	200: specificationtypeDBResponse
func (controller *Controller) UpdateSPECIFICATIONTYPE(c *gin.Context) {

	mutexSPECIFICATIONTYPE.Lock()
	defer mutexSPECIFICATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECIFICATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONTYPE.GetDB()

	// Validate input
	var input orm.SPECIFICATIONTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specificationtypeDB orm.SPECIFICATIONTYPEDB

	// fetch the specificationtype
	query := db.First(&specificationtypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specificationtypeDB.CopyBasicFieldsFromSPECIFICATIONTYPE_WOP(&input.SPECIFICATIONTYPE_WOP)
	specificationtypeDB.SPECIFICATIONTYPEPointersEncoding = input.SPECIFICATIONTYPEPointersEncoding

	query = db.Model(&specificationtypeDB).Updates(specificationtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specificationtypeNew := new(models.SPECIFICATIONTYPE)
	specificationtypeDB.CopyBasicFieldsToSPECIFICATIONTYPE(specificationtypeNew)

	// redeem pointers
	specificationtypeDB.DecodePointers(backRepo, specificationtypeNew)

	// get stage instance from DB instance, and call callback function
	specificationtypeOld := backRepo.BackRepoSPECIFICATIONTYPE.Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEPtr[specificationtypeDB.ID]
	if specificationtypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specificationtypeOld, specificationtypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specificationtypeDB
	c.JSON(http.StatusOK, specificationtypeDB)
}

// DeleteSPECIFICATIONTYPE
//
// swagger:route DELETE /specificationtypes/{ID} specificationtypes deleteSPECIFICATIONTYPE
//
// # Delete a specificationtype
//
// default: genericError
//
//	200: specificationtypeDBResponse
func (controller *Controller) DeleteSPECIFICATIONTYPE(c *gin.Context) {

	mutexSPECIFICATIONTYPE.Lock()
	defer mutexSPECIFICATIONTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECIFICATIONTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONTYPE.GetDB()

	// Get model if exist
	var specificationtypeDB orm.SPECIFICATIONTYPEDB
	if err := db.First(&specificationtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specificationtypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	specificationtypeDeleted := new(models.SPECIFICATIONTYPE)
	specificationtypeDB.CopyBasicFieldsToSPECIFICATIONTYPE(specificationtypeDeleted)

	// get stage instance from DB instance, and call callback function
	specificationtypeStaged := backRepo.BackRepoSPECIFICATIONTYPE.Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEPtr[specificationtypeDB.ID]
	if specificationtypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specificationtypeStaged, specificationtypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
