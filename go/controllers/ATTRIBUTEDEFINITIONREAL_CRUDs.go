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
var __ATTRIBUTEDEFINITIONREAL__dummysDeclaration__ models.ATTRIBUTEDEFINITIONREAL
var __ATTRIBUTEDEFINITIONREAL_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONREAL sync.Mutex

// An ATTRIBUTEDEFINITIONREALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONREAL updateATTRIBUTEDEFINITIONREAL deleteATTRIBUTEDEFINITIONREAL
type ATTRIBUTEDEFINITIONREALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONREALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONREAL updateATTRIBUTEDEFINITIONREAL
type ATTRIBUTEDEFINITIONREALInput struct {
	// The ATTRIBUTEDEFINITIONREAL to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONREAL *orm.ATTRIBUTEDEFINITIONREALAPI
}

// GetATTRIBUTEDEFINITIONREALs
//
// swagger:route GET /attributedefinitionreals attributedefinitionreals getATTRIBUTEDEFINITIONREALs
//
// # Get all attributedefinitionreals
//
// Responses:
// default: genericError
//
//	200: attributedefinitionrealDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONREALs(c *gin.Context) {

	// source slice
	var attributedefinitionrealDBs []orm.ATTRIBUTEDEFINITIONREALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetDB()

	query := db.Find(&attributedefinitionrealDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionrealAPIs := make([]orm.ATTRIBUTEDEFINITIONREALAPI, 0)

	// for each attributedefinitionreal, update fields from the database nullable fields
	for idx := range attributedefinitionrealDBs {
		attributedefinitionrealDB := &attributedefinitionrealDBs[idx]
		_ = attributedefinitionrealDB
		var attributedefinitionrealAPI orm.ATTRIBUTEDEFINITIONREALAPI

		// insertion point for updating fields
		attributedefinitionrealAPI.ID = attributedefinitionrealDB.ID
		attributedefinitionrealDB.CopyBasicFieldsToATTRIBUTEDEFINITIONREAL_WOP(&attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREAL_WOP)
		attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREALPointersEncoding = attributedefinitionrealDB.ATTRIBUTEDEFINITIONREALPointersEncoding
		attributedefinitionrealAPIs = append(attributedefinitionrealAPIs, attributedefinitionrealAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionrealAPIs)
}

// PostATTRIBUTEDEFINITIONREAL
//
// swagger:route POST /attributedefinitionreals attributedefinitionreals postATTRIBUTEDEFINITIONREAL
//
// Creates a attributedefinitionreal
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONREAL(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONREAL.Lock()
	defer mutexATTRIBUTEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONREALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitionreal
	attributedefinitionrealDB := orm.ATTRIBUTEDEFINITIONREALDB{}
	attributedefinitionrealDB.ATTRIBUTEDEFINITIONREALPointersEncoding = input.ATTRIBUTEDEFINITIONREALPointersEncoding
	attributedefinitionrealDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONREAL_WOP(&input.ATTRIBUTEDEFINITIONREAL_WOP)

	query := db.Create(&attributedefinitionrealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.CheckoutPhaseOneInstance(&attributedefinitionrealDB)
	attributedefinitionreal := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALPtr[attributedefinitionrealDB.ID]

	if attributedefinitionreal != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitionreal)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionrealDB)
}

// GetATTRIBUTEDEFINITIONREAL
//
// swagger:route GET /attributedefinitionreals/{ID} attributedefinitionreals getATTRIBUTEDEFINITIONREAL
//
// Gets the details for a attributedefinitionreal.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionrealDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONREAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetDB()

	// Get attributedefinitionrealDB in DB
	var attributedefinitionrealDB orm.ATTRIBUTEDEFINITIONREALDB
	if err := db.First(&attributedefinitionrealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionrealAPI orm.ATTRIBUTEDEFINITIONREALAPI
	attributedefinitionrealAPI.ID = attributedefinitionrealDB.ID
	attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREALPointersEncoding = attributedefinitionrealDB.ATTRIBUTEDEFINITIONREALPointersEncoding
	attributedefinitionrealDB.CopyBasicFieldsToATTRIBUTEDEFINITIONREAL_WOP(&attributedefinitionrealAPI.ATTRIBUTEDEFINITIONREAL_WOP)

	c.JSON(http.StatusOK, attributedefinitionrealAPI)
}

// UpdateATTRIBUTEDEFINITIONREAL
//
// swagger:route PATCH /attributedefinitionreals/{ID} attributedefinitionreals updateATTRIBUTEDEFINITIONREAL
//
// # Update a attributedefinitionreal
//
// Responses:
// default: genericError
//
//	200: attributedefinitionrealDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONREAL(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONREAL.Lock()
	defer mutexATTRIBUTEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONREALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionrealDB orm.ATTRIBUTEDEFINITIONREALDB

	// fetch the attributedefinitionreal
	query := db.First(&attributedefinitionrealDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionrealDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONREAL_WOP(&input.ATTRIBUTEDEFINITIONREAL_WOP)
	attributedefinitionrealDB.ATTRIBUTEDEFINITIONREALPointersEncoding = input.ATTRIBUTEDEFINITIONREALPointersEncoding

	query = db.Model(&attributedefinitionrealDB).Updates(attributedefinitionrealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionrealNew := new(models.ATTRIBUTEDEFINITIONREAL)
	attributedefinitionrealDB.CopyBasicFieldsToATTRIBUTEDEFINITIONREAL(attributedefinitionrealNew)

	// redeem pointers
	attributedefinitionrealDB.DecodePointers(backRepo, attributedefinitionrealNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionrealOld := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALPtr[attributedefinitionrealDB.ID]
	if attributedefinitionrealOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionrealOld, attributedefinitionrealNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionrealDB
	c.JSON(http.StatusOK, attributedefinitionrealDB)
}

// DeleteATTRIBUTEDEFINITIONREAL
//
// swagger:route DELETE /attributedefinitionreals/{ID} attributedefinitionreals deleteATTRIBUTEDEFINITIONREAL
//
// # Delete a attributedefinitionreal
//
// default: genericError
//
//	200: attributedefinitionrealDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONREAL(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONREAL.Lock()
	defer mutexATTRIBUTEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.GetDB()

	// Get model if exist
	var attributedefinitionrealDB orm.ATTRIBUTEDEFINITIONREALDB
	if err := db.First(&attributedefinitionrealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionrealDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionrealDeleted := new(models.ATTRIBUTEDEFINITIONREAL)
	attributedefinitionrealDB.CopyBasicFieldsToATTRIBUTEDEFINITIONREAL(attributedefinitionrealDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionrealStaged := backRepo.BackRepoATTRIBUTEDEFINITIONREAL.Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALPtr[attributedefinitionrealDB.ID]
	if attributedefinitionrealStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionrealStaged, attributedefinitionrealDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
