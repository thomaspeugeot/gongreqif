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
var __DATATYPEDEFINITIONENUMERATION__dummysDeclaration__ models.DATATYPEDEFINITIONENUMERATION
var __DATATYPEDEFINITIONENUMERATION_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONENUMERATION sync.Mutex

// An DATATYPEDEFINITIONENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONENUMERATION updateDATATYPEDEFINITIONENUMERATION deleteDATATYPEDEFINITIONENUMERATION
type DATATYPEDEFINITIONENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONENUMERATION updateDATATYPEDEFINITIONENUMERATION
type DATATYPEDEFINITIONENUMERATIONInput struct {
	// The DATATYPEDEFINITIONENUMERATION to submit or modify
	// in: body
	DATATYPEDEFINITIONENUMERATION *orm.DATATYPEDEFINITIONENUMERATIONAPI
}

// GetDATATYPEDEFINITIONENUMERATIONs
//
// swagger:route GET /datatypedefinitionenumerations datatypedefinitionenumerations getDATATYPEDEFINITIONENUMERATIONs
//
// # Get all datatypedefinitionenumerations
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionenumerationDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONENUMERATIONs(c *gin.Context) {

	// source slice
	var datatypedefinitionenumerationDBs []orm.DATATYPEDEFINITIONENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDB()

	query := db.Find(&datatypedefinitionenumerationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionenumerationAPIs := make([]orm.DATATYPEDEFINITIONENUMERATIONAPI, 0)

	// for each datatypedefinitionenumeration, update fields from the database nullable fields
	for idx := range datatypedefinitionenumerationDBs {
		datatypedefinitionenumerationDB := &datatypedefinitionenumerationDBs[idx]
		_ = datatypedefinitionenumerationDB
		var datatypedefinitionenumerationAPI orm.DATATYPEDEFINITIONENUMERATIONAPI

		// insertion point for updating fields
		datatypedefinitionenumerationAPI.ID = datatypedefinitionenumerationDB.ID
		datatypedefinitionenumerationDB.CopyBasicFieldsToDATATYPEDEFINITIONENUMERATION_WOP(&datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATION_WOP)
		datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATIONPointersEncoding = datatypedefinitionenumerationDB.DATATYPEDEFINITIONENUMERATIONPointersEncoding
		datatypedefinitionenumerationAPIs = append(datatypedefinitionenumerationAPIs, datatypedefinitionenumerationAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionenumerationAPIs)
}

// PostDATATYPEDEFINITIONENUMERATION
//
// swagger:route POST /datatypedefinitionenumerations datatypedefinitionenumerations postDATATYPEDEFINITIONENUMERATION
//
// Creates a datatypedefinitionenumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONENUMERATION(c *gin.Context) {

	mutexDATATYPEDEFINITIONENUMERATION.Lock()
	defer mutexDATATYPEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitionenumeration
	datatypedefinitionenumerationDB := orm.DATATYPEDEFINITIONENUMERATIONDB{}
	datatypedefinitionenumerationDB.DATATYPEDEFINITIONENUMERATIONPointersEncoding = input.DATATYPEDEFINITIONENUMERATIONPointersEncoding
	datatypedefinitionenumerationDB.CopyBasicFieldsFromDATATYPEDEFINITIONENUMERATION_WOP(&input.DATATYPEDEFINITIONENUMERATION_WOP)

	query := db.Create(&datatypedefinitionenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.CheckoutPhaseOneInstance(&datatypedefinitionenumerationDB)
	datatypedefinitionenumeration := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONPtr[datatypedefinitionenumerationDB.ID]

	if datatypedefinitionenumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitionenumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionenumerationDB)
}

// GetDATATYPEDEFINITIONENUMERATION
//
// swagger:route GET /datatypedefinitionenumerations/{ID} datatypedefinitionenumerations getDATATYPEDEFINITIONENUMERATION
//
// Gets the details for a datatypedefinitionenumeration.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionenumerationDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDB()

	// Get datatypedefinitionenumerationDB in DB
	var datatypedefinitionenumerationDB orm.DATATYPEDEFINITIONENUMERATIONDB
	if err := db.First(&datatypedefinitionenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionenumerationAPI orm.DATATYPEDEFINITIONENUMERATIONAPI
	datatypedefinitionenumerationAPI.ID = datatypedefinitionenumerationDB.ID
	datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATIONPointersEncoding = datatypedefinitionenumerationDB.DATATYPEDEFINITIONENUMERATIONPointersEncoding
	datatypedefinitionenumerationDB.CopyBasicFieldsToDATATYPEDEFINITIONENUMERATION_WOP(&datatypedefinitionenumerationAPI.DATATYPEDEFINITIONENUMERATION_WOP)

	c.JSON(http.StatusOK, datatypedefinitionenumerationAPI)
}

// UpdateDATATYPEDEFINITIONENUMERATION
//
// swagger:route PATCH /datatypedefinitionenumerations/{ID} datatypedefinitionenumerations updateDATATYPEDEFINITIONENUMERATION
//
// # Update a datatypedefinitionenumeration
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionenumerationDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONENUMERATION(c *gin.Context) {

	mutexDATATYPEDEFINITIONENUMERATION.Lock()
	defer mutexDATATYPEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionenumerationDB orm.DATATYPEDEFINITIONENUMERATIONDB

	// fetch the datatypedefinitionenumeration
	query := db.First(&datatypedefinitionenumerationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionenumerationDB.CopyBasicFieldsFromDATATYPEDEFINITIONENUMERATION_WOP(&input.DATATYPEDEFINITIONENUMERATION_WOP)
	datatypedefinitionenumerationDB.DATATYPEDEFINITIONENUMERATIONPointersEncoding = input.DATATYPEDEFINITIONENUMERATIONPointersEncoding

	query = db.Model(&datatypedefinitionenumerationDB).Updates(datatypedefinitionenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionenumerationNew := new(models.DATATYPEDEFINITIONENUMERATION)
	datatypedefinitionenumerationDB.CopyBasicFieldsToDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumerationNew)

	// redeem pointers
	datatypedefinitionenumerationDB.DecodePointers(backRepo, datatypedefinitionenumerationNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionenumerationOld := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONPtr[datatypedefinitionenumerationDB.ID]
	if datatypedefinitionenumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionenumerationOld, datatypedefinitionenumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionenumerationDB
	c.JSON(http.StatusOK, datatypedefinitionenumerationDB)
}

// DeleteDATATYPEDEFINITIONENUMERATION
//
// swagger:route DELETE /datatypedefinitionenumerations/{ID} datatypedefinitionenumerations deleteDATATYPEDEFINITIONENUMERATION
//
// # Delete a datatypedefinitionenumeration
//
// default: genericError
//
//	200: datatypedefinitionenumerationDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONENUMERATION(c *gin.Context) {

	mutexDATATYPEDEFINITIONENUMERATION.Lock()
	defer mutexDATATYPEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.GetDB()

	// Get model if exist
	var datatypedefinitionenumerationDB orm.DATATYPEDEFINITIONENUMERATIONDB
	if err := db.First(&datatypedefinitionenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionenumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionenumerationDeleted := new(models.DATATYPEDEFINITIONENUMERATION)
	datatypedefinitionenumerationDB.CopyBasicFieldsToDATATYPEDEFINITIONENUMERATION(datatypedefinitionenumerationDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionenumerationStaged := backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONPtr[datatypedefinitionenumerationDB.ID]
	if datatypedefinitionenumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionenumerationStaged, datatypedefinitionenumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
