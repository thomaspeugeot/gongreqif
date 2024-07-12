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
var __ATTRIBUTEDEFINITIONENUMERATION__dummysDeclaration__ models.ATTRIBUTEDEFINITIONENUMERATION
var __ATTRIBUTEDEFINITIONENUMERATION_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONENUMERATION sync.Mutex

// An ATTRIBUTEDEFINITIONENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONENUMERATION updateATTRIBUTEDEFINITIONENUMERATION deleteATTRIBUTEDEFINITIONENUMERATION
type ATTRIBUTEDEFINITIONENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONENUMERATION updateATTRIBUTEDEFINITIONENUMERATION
type ATTRIBUTEDEFINITIONENUMERATIONInput struct {
	// The ATTRIBUTEDEFINITIONENUMERATION to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONENUMERATION *orm.ATTRIBUTEDEFINITIONENUMERATIONAPI
}

// GetATTRIBUTEDEFINITIONENUMERATIONs
//
// swagger:route GET /attributedefinitionenumerations attributedefinitionenumerations getATTRIBUTEDEFINITIONENUMERATIONs
//
// # Get all attributedefinitionenumerations
//
// Responses:
// default: genericError
//
//	200: attributedefinitionenumerationDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONENUMERATIONs(c *gin.Context) {

	// source slice
	var attributedefinitionenumerationDBs []orm.ATTRIBUTEDEFINITIONENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetDB()

	query := db.Find(&attributedefinitionenumerationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionenumerationAPIs := make([]orm.ATTRIBUTEDEFINITIONENUMERATIONAPI, 0)

	// for each attributedefinitionenumeration, update fields from the database nullable fields
	for idx := range attributedefinitionenumerationDBs {
		attributedefinitionenumerationDB := &attributedefinitionenumerationDBs[idx]
		_ = attributedefinitionenumerationDB
		var attributedefinitionenumerationAPI orm.ATTRIBUTEDEFINITIONENUMERATIONAPI

		// insertion point for updating fields
		attributedefinitionenumerationAPI.ID = attributedefinitionenumerationDB.ID
		attributedefinitionenumerationDB.CopyBasicFieldsToATTRIBUTEDEFINITIONENUMERATION_WOP(&attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATION_WOP)
		attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding = attributedefinitionenumerationDB.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding
		attributedefinitionenumerationAPIs = append(attributedefinitionenumerationAPIs, attributedefinitionenumerationAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionenumerationAPIs)
}

// PostATTRIBUTEDEFINITIONENUMERATION
//
// swagger:route POST /attributedefinitionenumerations attributedefinitionenumerations postATTRIBUTEDEFINITIONENUMERATION
//
// Creates a attributedefinitionenumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONENUMERATION(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONENUMERATION.Lock()
	defer mutexATTRIBUTEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitionenumeration
	attributedefinitionenumerationDB := orm.ATTRIBUTEDEFINITIONENUMERATIONDB{}
	attributedefinitionenumerationDB.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding = input.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding
	attributedefinitionenumerationDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONENUMERATION_WOP(&input.ATTRIBUTEDEFINITIONENUMERATION_WOP)

	query := db.Create(&attributedefinitionenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.CheckoutPhaseOneInstance(&attributedefinitionenumerationDB)
	attributedefinitionenumeration := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONPtr[attributedefinitionenumerationDB.ID]

	if attributedefinitionenumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitionenumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionenumerationDB)
}

// GetATTRIBUTEDEFINITIONENUMERATION
//
// swagger:route GET /attributedefinitionenumerations/{ID} attributedefinitionenumerations getATTRIBUTEDEFINITIONENUMERATION
//
// Gets the details for a attributedefinitionenumeration.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionenumerationDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetDB()

	// Get attributedefinitionenumerationDB in DB
	var attributedefinitionenumerationDB orm.ATTRIBUTEDEFINITIONENUMERATIONDB
	if err := db.First(&attributedefinitionenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionenumerationAPI orm.ATTRIBUTEDEFINITIONENUMERATIONAPI
	attributedefinitionenumerationAPI.ID = attributedefinitionenumerationDB.ID
	attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding = attributedefinitionenumerationDB.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding
	attributedefinitionenumerationDB.CopyBasicFieldsToATTRIBUTEDEFINITIONENUMERATION_WOP(&attributedefinitionenumerationAPI.ATTRIBUTEDEFINITIONENUMERATION_WOP)

	c.JSON(http.StatusOK, attributedefinitionenumerationAPI)
}

// UpdateATTRIBUTEDEFINITIONENUMERATION
//
// swagger:route PATCH /attributedefinitionenumerations/{ID} attributedefinitionenumerations updateATTRIBUTEDEFINITIONENUMERATION
//
// # Update a attributedefinitionenumeration
//
// Responses:
// default: genericError
//
//	200: attributedefinitionenumerationDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONENUMERATION(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONENUMERATION.Lock()
	defer mutexATTRIBUTEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionenumerationDB orm.ATTRIBUTEDEFINITIONENUMERATIONDB

	// fetch the attributedefinitionenumeration
	query := db.First(&attributedefinitionenumerationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionenumerationDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONENUMERATION_WOP(&input.ATTRIBUTEDEFINITIONENUMERATION_WOP)
	attributedefinitionenumerationDB.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding = input.ATTRIBUTEDEFINITIONENUMERATIONPointersEncoding

	query = db.Model(&attributedefinitionenumerationDB).Updates(attributedefinitionenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionenumerationNew := new(models.ATTRIBUTEDEFINITIONENUMERATION)
	attributedefinitionenumerationDB.CopyBasicFieldsToATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumerationNew)

	// redeem pointers
	attributedefinitionenumerationDB.DecodePointers(backRepo, attributedefinitionenumerationNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionenumerationOld := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONPtr[attributedefinitionenumerationDB.ID]
	if attributedefinitionenumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionenumerationOld, attributedefinitionenumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionenumerationDB
	c.JSON(http.StatusOK, attributedefinitionenumerationDB)
}

// DeleteATTRIBUTEDEFINITIONENUMERATION
//
// swagger:route DELETE /attributedefinitionenumerations/{ID} attributedefinitionenumerations deleteATTRIBUTEDEFINITIONENUMERATION
//
// # Delete a attributedefinitionenumeration
//
// default: genericError
//
//	200: attributedefinitionenumerationDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONENUMERATION(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONENUMERATION.Lock()
	defer mutexATTRIBUTEDEFINITIONENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.GetDB()

	// Get model if exist
	var attributedefinitionenumerationDB orm.ATTRIBUTEDEFINITIONENUMERATIONDB
	if err := db.First(&attributedefinitionenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionenumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionenumerationDeleted := new(models.ATTRIBUTEDEFINITIONENUMERATION)
	attributedefinitionenumerationDB.CopyBasicFieldsToATTRIBUTEDEFINITIONENUMERATION(attributedefinitionenumerationDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionenumerationStaged := backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONPtr[attributedefinitionenumerationDB.ID]
	if attributedefinitionenumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionenumerationStaged, attributedefinitionenumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
