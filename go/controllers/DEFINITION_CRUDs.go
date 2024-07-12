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
var __DEFINITION__dummysDeclaration__ models.DEFINITION
var __DEFINITION_time__dummyDeclaration time.Duration

var mutexDEFINITION sync.Mutex

// An DEFINITIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDEFINITION updateDEFINITION deleteDEFINITION
type DEFINITIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DEFINITIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDEFINITION updateDEFINITION
type DEFINITIONInput struct {
	// The DEFINITION to submit or modify
	// in: body
	DEFINITION *orm.DEFINITIONAPI
}

// GetDEFINITIONs
//
// swagger:route GET /definitions definitions getDEFINITIONs
//
// # Get all definitions
//
// Responses:
// default: genericError
//
//	200: definitionDBResponse
func (controller *Controller) GetDEFINITIONs(c *gin.Context) {

	// source slice
	var definitionDBs []orm.DEFINITIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDEFINITIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFINITION.GetDB()

	query := db.Find(&definitionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	definitionAPIs := make([]orm.DEFINITIONAPI, 0)

	// for each definition, update fields from the database nullable fields
	for idx := range definitionDBs {
		definitionDB := &definitionDBs[idx]
		_ = definitionDB
		var definitionAPI orm.DEFINITIONAPI

		// insertion point for updating fields
		definitionAPI.ID = definitionDB.ID
		definitionDB.CopyBasicFieldsToDEFINITION_WOP(&definitionAPI.DEFINITION_WOP)
		definitionAPI.DEFINITIONPointersEncoding = definitionDB.DEFINITIONPointersEncoding
		definitionAPIs = append(definitionAPIs, definitionAPI)
	}

	c.JSON(http.StatusOK, definitionAPIs)
}

// PostDEFINITION
//
// swagger:route POST /definitions definitions postDEFINITION
//
// Creates a definition
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDEFINITION(c *gin.Context) {

	mutexDEFINITION.Lock()
	defer mutexDEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDEFINITIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFINITION.GetDB()

	// Validate input
	var input orm.DEFINITIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create definition
	definitionDB := orm.DEFINITIONDB{}
	definitionDB.DEFINITIONPointersEncoding = input.DEFINITIONPointersEncoding
	definitionDB.CopyBasicFieldsFromDEFINITION_WOP(&input.DEFINITION_WOP)

	query := db.Create(&definitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDEFINITION.CheckoutPhaseOneInstance(&definitionDB)
	definition := backRepo.BackRepoDEFINITION.Map_DEFINITIONDBID_DEFINITIONPtr[definitionDB.ID]

	if definition != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), definition)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, definitionDB)
}

// GetDEFINITION
//
// swagger:route GET /definitions/{ID} definitions getDEFINITION
//
// Gets the details for a definition.
//
// Responses:
// default: genericError
//
//	200: definitionDBResponse
func (controller *Controller) GetDEFINITION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFINITION.GetDB()

	// Get definitionDB in DB
	var definitionDB orm.DEFINITIONDB
	if err := db.First(&definitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var definitionAPI orm.DEFINITIONAPI
	definitionAPI.ID = definitionDB.ID
	definitionAPI.DEFINITIONPointersEncoding = definitionDB.DEFINITIONPointersEncoding
	definitionDB.CopyBasicFieldsToDEFINITION_WOP(&definitionAPI.DEFINITION_WOP)

	c.JSON(http.StatusOK, definitionAPI)
}

// UpdateDEFINITION
//
// swagger:route PATCH /definitions/{ID} definitions updateDEFINITION
//
// # Update a definition
//
// Responses:
// default: genericError
//
//	200: definitionDBResponse
func (controller *Controller) UpdateDEFINITION(c *gin.Context) {

	mutexDEFINITION.Lock()
	defer mutexDEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFINITION.GetDB()

	// Validate input
	var input orm.DEFINITIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var definitionDB orm.DEFINITIONDB

	// fetch the definition
	query := db.First(&definitionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	definitionDB.CopyBasicFieldsFromDEFINITION_WOP(&input.DEFINITION_WOP)
	definitionDB.DEFINITIONPointersEncoding = input.DEFINITIONPointersEncoding

	query = db.Model(&definitionDB).Updates(definitionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	definitionNew := new(models.DEFINITION)
	definitionDB.CopyBasicFieldsToDEFINITION(definitionNew)

	// redeem pointers
	definitionDB.DecodePointers(backRepo, definitionNew)

	// get stage instance from DB instance, and call callback function
	definitionOld := backRepo.BackRepoDEFINITION.Map_DEFINITIONDBID_DEFINITIONPtr[definitionDB.ID]
	if definitionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), definitionOld, definitionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the definitionDB
	c.JSON(http.StatusOK, definitionDB)
}

// DeleteDEFINITION
//
// swagger:route DELETE /definitions/{ID} definitions deleteDEFINITION
//
// # Delete a definition
//
// default: genericError
//
//	200: definitionDBResponse
func (controller *Controller) DeleteDEFINITION(c *gin.Context) {

	mutexDEFINITION.Lock()
	defer mutexDEFINITION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDEFINITION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFINITION.GetDB()

	// Get model if exist
	var definitionDB orm.DEFINITIONDB
	if err := db.First(&definitionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&definitionDB)

	// get an instance (not staged) from DB instance, and call callback function
	definitionDeleted := new(models.DEFINITION)
	definitionDB.CopyBasicFieldsToDEFINITION(definitionDeleted)

	// get stage instance from DB instance, and call callback function
	definitionStaged := backRepo.BackRepoDEFINITION.Map_DEFINITIONDBID_DEFINITIONPtr[definitionDB.ID]
	if definitionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), definitionStaged, definitionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
