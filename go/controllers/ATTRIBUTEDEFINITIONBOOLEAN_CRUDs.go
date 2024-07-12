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
var __ATTRIBUTEDEFINITIONBOOLEAN__dummysDeclaration__ models.ATTRIBUTEDEFINITIONBOOLEAN
var __ATTRIBUTEDEFINITIONBOOLEAN_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONBOOLEAN sync.Mutex

// An ATTRIBUTEDEFINITIONBOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONBOOLEAN updateATTRIBUTEDEFINITIONBOOLEAN deleteATTRIBUTEDEFINITIONBOOLEAN
type ATTRIBUTEDEFINITIONBOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONBOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONBOOLEAN updateATTRIBUTEDEFINITIONBOOLEAN
type ATTRIBUTEDEFINITIONBOOLEANInput struct {
	// The ATTRIBUTEDEFINITIONBOOLEAN to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONBOOLEAN *orm.ATTRIBUTEDEFINITIONBOOLEANAPI
}

// GetATTRIBUTEDEFINITIONBOOLEANs
//
// swagger:route GET /attributedefinitionbooleans attributedefinitionbooleans getATTRIBUTEDEFINITIONBOOLEANs
//
// # Get all attributedefinitionbooleans
//
// Responses:
// default: genericError
//
//	200: attributedefinitionbooleanDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONBOOLEANs(c *gin.Context) {

	// source slice
	var attributedefinitionbooleanDBs []orm.ATTRIBUTEDEFINITIONBOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetDB()

	query := db.Find(&attributedefinitionbooleanDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionbooleanAPIs := make([]orm.ATTRIBUTEDEFINITIONBOOLEANAPI, 0)

	// for each attributedefinitionboolean, update fields from the database nullable fields
	for idx := range attributedefinitionbooleanDBs {
		attributedefinitionbooleanDB := &attributedefinitionbooleanDBs[idx]
		_ = attributedefinitionbooleanDB
		var attributedefinitionbooleanAPI orm.ATTRIBUTEDEFINITIONBOOLEANAPI

		// insertion point for updating fields
		attributedefinitionbooleanAPI.ID = attributedefinitionbooleanDB.ID
		attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN_WOP(&attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEAN_WOP)
		attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding = attributedefinitionbooleanDB.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
		attributedefinitionbooleanAPIs = append(attributedefinitionbooleanAPIs, attributedefinitionbooleanAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionbooleanAPIs)
}

// PostATTRIBUTEDEFINITIONBOOLEAN
//
// swagger:route POST /attributedefinitionbooleans attributedefinitionbooleans postATTRIBUTEDEFINITIONBOOLEAN
//
// Creates a attributedefinitionboolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONBOOLEAN.Lock()
	defer mutexATTRIBUTEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONBOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitionboolean
	attributedefinitionbooleanDB := orm.ATTRIBUTEDEFINITIONBOOLEANDB{}
	attributedefinitionbooleanDB.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding = input.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
	attributedefinitionbooleanDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN_WOP(&input.ATTRIBUTEDEFINITIONBOOLEAN_WOP)

	query := db.Create(&attributedefinitionbooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseOneInstance(&attributedefinitionbooleanDB)
	attributedefinitionboolean := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]

	if attributedefinitionboolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitionboolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionbooleanDB)
}

// GetATTRIBUTEDEFINITIONBOOLEAN
//
// swagger:route GET /attributedefinitionbooleans/{ID} attributedefinitionbooleans getATTRIBUTEDEFINITIONBOOLEAN
//
// Gets the details for a attributedefinitionboolean.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionbooleanDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONBOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetDB()

	// Get attributedefinitionbooleanDB in DB
	var attributedefinitionbooleanDB orm.ATTRIBUTEDEFINITIONBOOLEANDB
	if err := db.First(&attributedefinitionbooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionbooleanAPI orm.ATTRIBUTEDEFINITIONBOOLEANAPI
	attributedefinitionbooleanAPI.ID = attributedefinitionbooleanDB.ID
	attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding = attributedefinitionbooleanDB.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
	attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN_WOP(&attributedefinitionbooleanAPI.ATTRIBUTEDEFINITIONBOOLEAN_WOP)

	c.JSON(http.StatusOK, attributedefinitionbooleanAPI)
}

// UpdateATTRIBUTEDEFINITIONBOOLEAN
//
// swagger:route PATCH /attributedefinitionbooleans/{ID} attributedefinitionbooleans updateATTRIBUTEDEFINITIONBOOLEAN
//
// # Update a attributedefinitionboolean
//
// Responses:
// default: genericError
//
//	200: attributedefinitionbooleanDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONBOOLEAN.Lock()
	defer mutexATTRIBUTEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONBOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionbooleanDB orm.ATTRIBUTEDEFINITIONBOOLEANDB

	// fetch the attributedefinitionboolean
	query := db.First(&attributedefinitionbooleanDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionbooleanDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN_WOP(&input.ATTRIBUTEDEFINITIONBOOLEAN_WOP)
	attributedefinitionbooleanDB.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding = input.ATTRIBUTEDEFINITIONBOOLEANPointersEncoding

	query = db.Model(&attributedefinitionbooleanDB).Updates(attributedefinitionbooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionbooleanNew := new(models.ATTRIBUTEDEFINITIONBOOLEAN)
	attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionbooleanNew)

	// redeem pointers
	attributedefinitionbooleanDB.DecodePointers(backRepo, attributedefinitionbooleanNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionbooleanOld := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]
	if attributedefinitionbooleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionbooleanOld, attributedefinitionbooleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionbooleanDB
	c.JSON(http.StatusOK, attributedefinitionbooleanDB)
}

// DeleteATTRIBUTEDEFINITIONBOOLEAN
//
// swagger:route DELETE /attributedefinitionbooleans/{ID} attributedefinitionbooleans deleteATTRIBUTEDEFINITIONBOOLEAN
//
// # Delete a attributedefinitionboolean
//
// default: genericError
//
//	200: attributedefinitionbooleanDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONBOOLEAN.Lock()
	defer mutexATTRIBUTEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.GetDB()

	// Get model if exist
	var attributedefinitionbooleanDB orm.ATTRIBUTEDEFINITIONBOOLEANDB
	if err := db.First(&attributedefinitionbooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionbooleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionbooleanDeleted := new(models.ATTRIBUTEDEFINITIONBOOLEAN)
	attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionbooleanDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionbooleanStaged := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]
	if attributedefinitionbooleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionbooleanStaged, attributedefinitionbooleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
