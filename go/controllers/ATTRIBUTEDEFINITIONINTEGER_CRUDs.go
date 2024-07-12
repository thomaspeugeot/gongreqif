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
var __ATTRIBUTEDEFINITIONINTEGER__dummysDeclaration__ models.ATTRIBUTEDEFINITIONINTEGER
var __ATTRIBUTEDEFINITIONINTEGER_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONINTEGER sync.Mutex

// An ATTRIBUTEDEFINITIONINTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONINTEGER updateATTRIBUTEDEFINITIONINTEGER deleteATTRIBUTEDEFINITIONINTEGER
type ATTRIBUTEDEFINITIONINTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONINTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONINTEGER updateATTRIBUTEDEFINITIONINTEGER
type ATTRIBUTEDEFINITIONINTEGERInput struct {
	// The ATTRIBUTEDEFINITIONINTEGER to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONINTEGER *orm.ATTRIBUTEDEFINITIONINTEGERAPI
}

// GetATTRIBUTEDEFINITIONINTEGERs
//
// swagger:route GET /attributedefinitionintegers attributedefinitionintegers getATTRIBUTEDEFINITIONINTEGERs
//
// # Get all attributedefinitionintegers
//
// Responses:
// default: genericError
//
//	200: attributedefinitionintegerDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONINTEGERs(c *gin.Context) {

	// source slice
	var attributedefinitionintegerDBs []orm.ATTRIBUTEDEFINITIONINTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetDB()

	query := db.Find(&attributedefinitionintegerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionintegerAPIs := make([]orm.ATTRIBUTEDEFINITIONINTEGERAPI, 0)

	// for each attributedefinitioninteger, update fields from the database nullable fields
	for idx := range attributedefinitionintegerDBs {
		attributedefinitionintegerDB := &attributedefinitionintegerDBs[idx]
		_ = attributedefinitionintegerDB
		var attributedefinitionintegerAPI orm.ATTRIBUTEDEFINITIONINTEGERAPI

		// insertion point for updating fields
		attributedefinitionintegerAPI.ID = attributedefinitionintegerDB.ID
		attributedefinitionintegerDB.CopyBasicFieldsToATTRIBUTEDEFINITIONINTEGER_WOP(&attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGER_WOP)
		attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGERPointersEncoding = attributedefinitionintegerDB.ATTRIBUTEDEFINITIONINTEGERPointersEncoding
		attributedefinitionintegerAPIs = append(attributedefinitionintegerAPIs, attributedefinitionintegerAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionintegerAPIs)
}

// PostATTRIBUTEDEFINITIONINTEGER
//
// swagger:route POST /attributedefinitionintegers attributedefinitionintegers postATTRIBUTEDEFINITIONINTEGER
//
// Creates a attributedefinitioninteger
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONINTEGER(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONINTEGER.Lock()
	defer mutexATTRIBUTEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONINTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitioninteger
	attributedefinitionintegerDB := orm.ATTRIBUTEDEFINITIONINTEGERDB{}
	attributedefinitionintegerDB.ATTRIBUTEDEFINITIONINTEGERPointersEncoding = input.ATTRIBUTEDEFINITIONINTEGERPointersEncoding
	attributedefinitionintegerDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONINTEGER_WOP(&input.ATTRIBUTEDEFINITIONINTEGER_WOP)

	query := db.Create(&attributedefinitionintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.CheckoutPhaseOneInstance(&attributedefinitionintegerDB)
	attributedefinitioninteger := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERPtr[attributedefinitionintegerDB.ID]

	if attributedefinitioninteger != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitioninteger)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionintegerDB)
}

// GetATTRIBUTEDEFINITIONINTEGER
//
// swagger:route GET /attributedefinitionintegers/{ID} attributedefinitionintegers getATTRIBUTEDEFINITIONINTEGER
//
// Gets the details for a attributedefinitioninteger.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionintegerDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONINTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetDB()

	// Get attributedefinitionintegerDB in DB
	var attributedefinitionintegerDB orm.ATTRIBUTEDEFINITIONINTEGERDB
	if err := db.First(&attributedefinitionintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionintegerAPI orm.ATTRIBUTEDEFINITIONINTEGERAPI
	attributedefinitionintegerAPI.ID = attributedefinitionintegerDB.ID
	attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGERPointersEncoding = attributedefinitionintegerDB.ATTRIBUTEDEFINITIONINTEGERPointersEncoding
	attributedefinitionintegerDB.CopyBasicFieldsToATTRIBUTEDEFINITIONINTEGER_WOP(&attributedefinitionintegerAPI.ATTRIBUTEDEFINITIONINTEGER_WOP)

	c.JSON(http.StatusOK, attributedefinitionintegerAPI)
}

// UpdateATTRIBUTEDEFINITIONINTEGER
//
// swagger:route PATCH /attributedefinitionintegers/{ID} attributedefinitionintegers updateATTRIBUTEDEFINITIONINTEGER
//
// # Update a attributedefinitioninteger
//
// Responses:
// default: genericError
//
//	200: attributedefinitionintegerDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONINTEGER(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONINTEGER.Lock()
	defer mutexATTRIBUTEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONINTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionintegerDB orm.ATTRIBUTEDEFINITIONINTEGERDB

	// fetch the attributedefinitioninteger
	query := db.First(&attributedefinitionintegerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionintegerDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONINTEGER_WOP(&input.ATTRIBUTEDEFINITIONINTEGER_WOP)
	attributedefinitionintegerDB.ATTRIBUTEDEFINITIONINTEGERPointersEncoding = input.ATTRIBUTEDEFINITIONINTEGERPointersEncoding

	query = db.Model(&attributedefinitionintegerDB).Updates(attributedefinitionintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionintegerNew := new(models.ATTRIBUTEDEFINITIONINTEGER)
	attributedefinitionintegerDB.CopyBasicFieldsToATTRIBUTEDEFINITIONINTEGER(attributedefinitionintegerNew)

	// redeem pointers
	attributedefinitionintegerDB.DecodePointers(backRepo, attributedefinitionintegerNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionintegerOld := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERPtr[attributedefinitionintegerDB.ID]
	if attributedefinitionintegerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionintegerOld, attributedefinitionintegerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionintegerDB
	c.JSON(http.StatusOK, attributedefinitionintegerDB)
}

// DeleteATTRIBUTEDEFINITIONINTEGER
//
// swagger:route DELETE /attributedefinitionintegers/{ID} attributedefinitionintegers deleteATTRIBUTEDEFINITIONINTEGER
//
// # Delete a attributedefinitioninteger
//
// default: genericError
//
//	200: attributedefinitionintegerDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONINTEGER(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONINTEGER.Lock()
	defer mutexATTRIBUTEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.GetDB()

	// Get model if exist
	var attributedefinitionintegerDB orm.ATTRIBUTEDEFINITIONINTEGERDB
	if err := db.First(&attributedefinitionintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionintegerDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionintegerDeleted := new(models.ATTRIBUTEDEFINITIONINTEGER)
	attributedefinitionintegerDB.CopyBasicFieldsToATTRIBUTEDEFINITIONINTEGER(attributedefinitionintegerDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionintegerStaged := backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERPtr[attributedefinitionintegerDB.ID]
	if attributedefinitionintegerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionintegerStaged, attributedefinitionintegerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
