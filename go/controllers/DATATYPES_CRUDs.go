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
var __DATATYPES__dummysDeclaration__ models.DATATYPES
var __DATATYPES_time__dummyDeclaration time.Duration

var mutexDATATYPES sync.Mutex

// An DATATYPESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPES updateDATATYPES deleteDATATYPES
type DATATYPESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPES updateDATATYPES
type DATATYPESInput struct {
	// The DATATYPES to submit or modify
	// in: body
	DATATYPES *orm.DATATYPESAPI
}

// GetDATATYPESs
//
// swagger:route GET /datatypess datatypess getDATATYPESs
//
// # Get all datatypess
//
// Responses:
// default: genericError
//
//	200: datatypesDBResponse
func (controller *Controller) GetDATATYPESs(c *gin.Context) {

	// source slice
	var datatypesDBs []orm.DATATYPESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPES.GetDB()

	query := db.Find(&datatypesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypesAPIs := make([]orm.DATATYPESAPI, 0)

	// for each datatypes, update fields from the database nullable fields
	for idx := range datatypesDBs {
		datatypesDB := &datatypesDBs[idx]
		_ = datatypesDB
		var datatypesAPI orm.DATATYPESAPI

		// insertion point for updating fields
		datatypesAPI.ID = datatypesDB.ID
		datatypesDB.CopyBasicFieldsToDATATYPES_WOP(&datatypesAPI.DATATYPES_WOP)
		datatypesAPI.DATATYPESPointersEncoding = datatypesDB.DATATYPESPointersEncoding
		datatypesAPIs = append(datatypesAPIs, datatypesAPI)
	}

	c.JSON(http.StatusOK, datatypesAPIs)
}

// PostDATATYPES
//
// swagger:route POST /datatypess datatypess postDATATYPES
//
// Creates a datatypes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPES(c *gin.Context) {

	mutexDATATYPES.Lock()
	defer mutexDATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPES.GetDB()

	// Validate input
	var input orm.DATATYPESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypes
	datatypesDB := orm.DATATYPESDB{}
	datatypesDB.DATATYPESPointersEncoding = input.DATATYPESPointersEncoding
	datatypesDB.CopyBasicFieldsFromDATATYPES_WOP(&input.DATATYPES_WOP)

	query := db.Create(&datatypesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPES.CheckoutPhaseOneInstance(&datatypesDB)
	datatypes := backRepo.BackRepoDATATYPES.Map_DATATYPESDBID_DATATYPESPtr[datatypesDB.ID]

	if datatypes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypesDB)
}

// GetDATATYPES
//
// swagger:route GET /datatypess/{ID} datatypess getDATATYPES
//
// Gets the details for a datatypes.
//
// Responses:
// default: genericError
//
//	200: datatypesDBResponse
func (controller *Controller) GetDATATYPES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPES.GetDB()

	// Get datatypesDB in DB
	var datatypesDB orm.DATATYPESDB
	if err := db.First(&datatypesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypesAPI orm.DATATYPESAPI
	datatypesAPI.ID = datatypesDB.ID
	datatypesAPI.DATATYPESPointersEncoding = datatypesDB.DATATYPESPointersEncoding
	datatypesDB.CopyBasicFieldsToDATATYPES_WOP(&datatypesAPI.DATATYPES_WOP)

	c.JSON(http.StatusOK, datatypesAPI)
}

// UpdateDATATYPES
//
// swagger:route PATCH /datatypess/{ID} datatypess updateDATATYPES
//
// # Update a datatypes
//
// Responses:
// default: genericError
//
//	200: datatypesDBResponse
func (controller *Controller) UpdateDATATYPES(c *gin.Context) {

	mutexDATATYPES.Lock()
	defer mutexDATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPES.GetDB()

	// Validate input
	var input orm.DATATYPESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypesDB orm.DATATYPESDB

	// fetch the datatypes
	query := db.First(&datatypesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypesDB.CopyBasicFieldsFromDATATYPES_WOP(&input.DATATYPES_WOP)
	datatypesDB.DATATYPESPointersEncoding = input.DATATYPESPointersEncoding

	query = db.Model(&datatypesDB).Updates(datatypesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypesNew := new(models.DATATYPES)
	datatypesDB.CopyBasicFieldsToDATATYPES(datatypesNew)

	// redeem pointers
	datatypesDB.DecodePointers(backRepo, datatypesNew)

	// get stage instance from DB instance, and call callback function
	datatypesOld := backRepo.BackRepoDATATYPES.Map_DATATYPESDBID_DATATYPESPtr[datatypesDB.ID]
	if datatypesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypesOld, datatypesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypesDB
	c.JSON(http.StatusOK, datatypesDB)
}

// DeleteDATATYPES
//
// swagger:route DELETE /datatypess/{ID} datatypess deleteDATATYPES
//
// # Delete a datatypes
//
// default: genericError
//
//	200: datatypesDBResponse
func (controller *Controller) DeleteDATATYPES(c *gin.Context) {

	mutexDATATYPES.Lock()
	defer mutexDATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPES.GetDB()

	// Get model if exist
	var datatypesDB orm.DATATYPESDB
	if err := db.First(&datatypesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypesDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypesDeleted := new(models.DATATYPES)
	datatypesDB.CopyBasicFieldsToDATATYPES(datatypesDeleted)

	// get stage instance from DB instance, and call callback function
	datatypesStaged := backRepo.BackRepoDATATYPES.Map_DATATYPESDBID_DATATYPESPtr[datatypesDB.ID]
	if datatypesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypesStaged, datatypesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
