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
var __SPECOBJECT__dummysDeclaration__ models.SPECOBJECT
var __SPECOBJECT_time__dummyDeclaration time.Duration

var mutexSPECOBJECT sync.Mutex

// An SPECOBJECTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECOBJECT updateSPECOBJECT deleteSPECOBJECT
type SPECOBJECTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECOBJECTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECOBJECT updateSPECOBJECT
type SPECOBJECTInput struct {
	// The SPECOBJECT to submit or modify
	// in: body
	SPECOBJECT *orm.SPECOBJECTAPI
}

// GetSPECOBJECTs
//
// swagger:route GET /specobjects specobjects getSPECOBJECTs
//
// # Get all specobjects
//
// Responses:
// default: genericError
//
//	200: specobjectDBResponse
func (controller *Controller) GetSPECOBJECTs(c *gin.Context) {

	// source slice
	var specobjectDBs []orm.SPECOBJECTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECT.GetDB()

	query := db.Find(&specobjectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specobjectAPIs := make([]orm.SPECOBJECTAPI, 0)

	// for each specobject, update fields from the database nullable fields
	for idx := range specobjectDBs {
		specobjectDB := &specobjectDBs[idx]
		_ = specobjectDB
		var specobjectAPI orm.SPECOBJECTAPI

		// insertion point for updating fields
		specobjectAPI.ID = specobjectDB.ID
		specobjectDB.CopyBasicFieldsToSPECOBJECT_WOP(&specobjectAPI.SPECOBJECT_WOP)
		specobjectAPI.SPECOBJECTPointersEncoding = specobjectDB.SPECOBJECTPointersEncoding
		specobjectAPIs = append(specobjectAPIs, specobjectAPI)
	}

	c.JSON(http.StatusOK, specobjectAPIs)
}

// PostSPECOBJECT
//
// swagger:route POST /specobjects specobjects postSPECOBJECT
//
// Creates a specobject
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECOBJECT(c *gin.Context) {

	mutexSPECOBJECT.Lock()
	defer mutexSPECOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECOBJECTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECT.GetDB()

	// Validate input
	var input orm.SPECOBJECTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specobject
	specobjectDB := orm.SPECOBJECTDB{}
	specobjectDB.SPECOBJECTPointersEncoding = input.SPECOBJECTPointersEncoding
	specobjectDB.CopyBasicFieldsFromSPECOBJECT_WOP(&input.SPECOBJECT_WOP)

	query := db.Create(&specobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECOBJECT.CheckoutPhaseOneInstance(&specobjectDB)
	specobject := backRepo.BackRepoSPECOBJECT.Map_SPECOBJECTDBID_SPECOBJECTPtr[specobjectDB.ID]

	if specobject != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specobject)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specobjectDB)
}

// GetSPECOBJECT
//
// swagger:route GET /specobjects/{ID} specobjects getSPECOBJECT
//
// Gets the details for a specobject.
//
// Responses:
// default: genericError
//
//	200: specobjectDBResponse
func (controller *Controller) GetSPECOBJECT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECT.GetDB()

	// Get specobjectDB in DB
	var specobjectDB orm.SPECOBJECTDB
	if err := db.First(&specobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specobjectAPI orm.SPECOBJECTAPI
	specobjectAPI.ID = specobjectDB.ID
	specobjectAPI.SPECOBJECTPointersEncoding = specobjectDB.SPECOBJECTPointersEncoding
	specobjectDB.CopyBasicFieldsToSPECOBJECT_WOP(&specobjectAPI.SPECOBJECT_WOP)

	c.JSON(http.StatusOK, specobjectAPI)
}

// UpdateSPECOBJECT
//
// swagger:route PATCH /specobjects/{ID} specobjects updateSPECOBJECT
//
// # Update a specobject
//
// Responses:
// default: genericError
//
//	200: specobjectDBResponse
func (controller *Controller) UpdateSPECOBJECT(c *gin.Context) {

	mutexSPECOBJECT.Lock()
	defer mutexSPECOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECT.GetDB()

	// Validate input
	var input orm.SPECOBJECTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specobjectDB orm.SPECOBJECTDB

	// fetch the specobject
	query := db.First(&specobjectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specobjectDB.CopyBasicFieldsFromSPECOBJECT_WOP(&input.SPECOBJECT_WOP)
	specobjectDB.SPECOBJECTPointersEncoding = input.SPECOBJECTPointersEncoding

	query = db.Model(&specobjectDB).Updates(specobjectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specobjectNew := new(models.SPECOBJECT)
	specobjectDB.CopyBasicFieldsToSPECOBJECT(specobjectNew)

	// redeem pointers
	specobjectDB.DecodePointers(backRepo, specobjectNew)

	// get stage instance from DB instance, and call callback function
	specobjectOld := backRepo.BackRepoSPECOBJECT.Map_SPECOBJECTDBID_SPECOBJECTPtr[specobjectDB.ID]
	if specobjectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specobjectOld, specobjectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specobjectDB
	c.JSON(http.StatusOK, specobjectDB)
}

// DeleteSPECOBJECT
//
// swagger:route DELETE /specobjects/{ID} specobjects deleteSPECOBJECT
//
// # Delete a specobject
//
// default: genericError
//
//	200: specobjectDBResponse
func (controller *Controller) DeleteSPECOBJECT(c *gin.Context) {

	mutexSPECOBJECT.Lock()
	defer mutexSPECOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECOBJECT.GetDB()

	// Get model if exist
	var specobjectDB orm.SPECOBJECTDB
	if err := db.First(&specobjectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specobjectDB)

	// get an instance (not staged) from DB instance, and call callback function
	specobjectDeleted := new(models.SPECOBJECT)
	specobjectDB.CopyBasicFieldsToSPECOBJECT(specobjectDeleted)

	// get stage instance from DB instance, and call callback function
	specobjectStaged := backRepo.BackRepoSPECOBJECT.Map_SPECOBJECTDBID_SPECOBJECTPtr[specobjectDB.ID]
	if specobjectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specobjectStaged, specobjectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
