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
var __ATTRIBUTEVALUEENUMERATION__dummysDeclaration__ models.ATTRIBUTEVALUEENUMERATION
var __ATTRIBUTEVALUEENUMERATION_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEENUMERATION sync.Mutex

// An ATTRIBUTEVALUEENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEENUMERATION updateATTRIBUTEVALUEENUMERATION deleteATTRIBUTEVALUEENUMERATION
type ATTRIBUTEVALUEENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEENUMERATION updateATTRIBUTEVALUEENUMERATION
type ATTRIBUTEVALUEENUMERATIONInput struct {
	// The ATTRIBUTEVALUEENUMERATION to submit or modify
	// in: body
	ATTRIBUTEVALUEENUMERATION *orm.ATTRIBUTEVALUEENUMERATIONAPI
}

// GetATTRIBUTEVALUEENUMERATIONs
//
// swagger:route GET /attributevalueenumerations attributevalueenumerations getATTRIBUTEVALUEENUMERATIONs
//
// # Get all attributevalueenumerations
//
// Responses:
// default: genericError
//
//	200: attributevalueenumerationDBResponse
func (controller *Controller) GetATTRIBUTEVALUEENUMERATIONs(c *gin.Context) {

	// source slice
	var attributevalueenumerationDBs []orm.ATTRIBUTEVALUEENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetDB()

	query := db.Find(&attributevalueenumerationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevalueenumerationAPIs := make([]orm.ATTRIBUTEVALUEENUMERATIONAPI, 0)

	// for each attributevalueenumeration, update fields from the database nullable fields
	for idx := range attributevalueenumerationDBs {
		attributevalueenumerationDB := &attributevalueenumerationDBs[idx]
		_ = attributevalueenumerationDB
		var attributevalueenumerationAPI orm.ATTRIBUTEVALUEENUMERATIONAPI

		// insertion point for updating fields
		attributevalueenumerationAPI.ID = attributevalueenumerationDB.ID
		attributevalueenumerationDB.CopyBasicFieldsToATTRIBUTEVALUEENUMERATION_WOP(&attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATION_WOP)
		attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATIONPointersEncoding = attributevalueenumerationDB.ATTRIBUTEVALUEENUMERATIONPointersEncoding
		attributevalueenumerationAPIs = append(attributevalueenumerationAPIs, attributevalueenumerationAPI)
	}

	c.JSON(http.StatusOK, attributevalueenumerationAPIs)
}

// PostATTRIBUTEVALUEENUMERATION
//
// swagger:route POST /attributevalueenumerations attributevalueenumerations postATTRIBUTEVALUEENUMERATION
//
// Creates a attributevalueenumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEENUMERATION(c *gin.Context) {

	mutexATTRIBUTEVALUEENUMERATION.Lock()
	defer mutexATTRIBUTEVALUEENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevalueenumeration
	attributevalueenumerationDB := orm.ATTRIBUTEVALUEENUMERATIONDB{}
	attributevalueenumerationDB.ATTRIBUTEVALUEENUMERATIONPointersEncoding = input.ATTRIBUTEVALUEENUMERATIONPointersEncoding
	attributevalueenumerationDB.CopyBasicFieldsFromATTRIBUTEVALUEENUMERATION_WOP(&input.ATTRIBUTEVALUEENUMERATION_WOP)

	query := db.Create(&attributevalueenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.CheckoutPhaseOneInstance(&attributevalueenumerationDB)
	attributevalueenumeration := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONPtr[attributevalueenumerationDB.ID]

	if attributevalueenumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevalueenumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevalueenumerationDB)
}

// GetATTRIBUTEVALUEENUMERATION
//
// swagger:route GET /attributevalueenumerations/{ID} attributevalueenumerations getATTRIBUTEVALUEENUMERATION
//
// Gets the details for a attributevalueenumeration.
//
// Responses:
// default: genericError
//
//	200: attributevalueenumerationDBResponse
func (controller *Controller) GetATTRIBUTEVALUEENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetDB()

	// Get attributevalueenumerationDB in DB
	var attributevalueenumerationDB orm.ATTRIBUTEVALUEENUMERATIONDB
	if err := db.First(&attributevalueenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevalueenumerationAPI orm.ATTRIBUTEVALUEENUMERATIONAPI
	attributevalueenumerationAPI.ID = attributevalueenumerationDB.ID
	attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATIONPointersEncoding = attributevalueenumerationDB.ATTRIBUTEVALUEENUMERATIONPointersEncoding
	attributevalueenumerationDB.CopyBasicFieldsToATTRIBUTEVALUEENUMERATION_WOP(&attributevalueenumerationAPI.ATTRIBUTEVALUEENUMERATION_WOP)

	c.JSON(http.StatusOK, attributevalueenumerationAPI)
}

// UpdateATTRIBUTEVALUEENUMERATION
//
// swagger:route PATCH /attributevalueenumerations/{ID} attributevalueenumerations updateATTRIBUTEVALUEENUMERATION
//
// # Update a attributevalueenumeration
//
// Responses:
// default: genericError
//
//	200: attributevalueenumerationDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEENUMERATION(c *gin.Context) {

	mutexATTRIBUTEVALUEENUMERATION.Lock()
	defer mutexATTRIBUTEVALUEENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevalueenumerationDB orm.ATTRIBUTEVALUEENUMERATIONDB

	// fetch the attributevalueenumeration
	query := db.First(&attributevalueenumerationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevalueenumerationDB.CopyBasicFieldsFromATTRIBUTEVALUEENUMERATION_WOP(&input.ATTRIBUTEVALUEENUMERATION_WOP)
	attributevalueenumerationDB.ATTRIBUTEVALUEENUMERATIONPointersEncoding = input.ATTRIBUTEVALUEENUMERATIONPointersEncoding

	query = db.Model(&attributevalueenumerationDB).Updates(attributevalueenumerationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevalueenumerationNew := new(models.ATTRIBUTEVALUEENUMERATION)
	attributevalueenumerationDB.CopyBasicFieldsToATTRIBUTEVALUEENUMERATION(attributevalueenumerationNew)

	// redeem pointers
	attributevalueenumerationDB.DecodePointers(backRepo, attributevalueenumerationNew)

	// get stage instance from DB instance, and call callback function
	attributevalueenumerationOld := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONPtr[attributevalueenumerationDB.ID]
	if attributevalueenumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevalueenumerationOld, attributevalueenumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevalueenumerationDB
	c.JSON(http.StatusOK, attributevalueenumerationDB)
}

// DeleteATTRIBUTEVALUEENUMERATION
//
// swagger:route DELETE /attributevalueenumerations/{ID} attributevalueenumerations deleteATTRIBUTEVALUEENUMERATION
//
// # Delete a attributevalueenumeration
//
// default: genericError
//
//	200: attributevalueenumerationDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEENUMERATION(c *gin.Context) {

	mutexATTRIBUTEVALUEENUMERATION.Lock()
	defer mutexATTRIBUTEVALUEENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.GetDB()

	// Get model if exist
	var attributevalueenumerationDB orm.ATTRIBUTEVALUEENUMERATIONDB
	if err := db.First(&attributevalueenumerationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevalueenumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevalueenumerationDeleted := new(models.ATTRIBUTEVALUEENUMERATION)
	attributevalueenumerationDB.CopyBasicFieldsToATTRIBUTEVALUEENUMERATION(attributevalueenumerationDeleted)

	// get stage instance from DB instance, and call callback function
	attributevalueenumerationStaged := backRepo.BackRepoATTRIBUTEVALUEENUMERATION.Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONPtr[attributevalueenumerationDB.ID]
	if attributevalueenumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevalueenumerationStaged, attributevalueenumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
