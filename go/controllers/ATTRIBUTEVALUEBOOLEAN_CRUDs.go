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
var __ATTRIBUTEVALUEBOOLEAN__dummysDeclaration__ models.ATTRIBUTEVALUEBOOLEAN
var __ATTRIBUTEVALUEBOOLEAN_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEBOOLEAN sync.Mutex

// An ATTRIBUTEVALUEBOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEBOOLEAN updateATTRIBUTEVALUEBOOLEAN deleteATTRIBUTEVALUEBOOLEAN
type ATTRIBUTEVALUEBOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEBOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEBOOLEAN updateATTRIBUTEVALUEBOOLEAN
type ATTRIBUTEVALUEBOOLEANInput struct {
	// The ATTRIBUTEVALUEBOOLEAN to submit or modify
	// in: body
	ATTRIBUTEVALUEBOOLEAN *orm.ATTRIBUTEVALUEBOOLEANAPI
}

// GetATTRIBUTEVALUEBOOLEANs
//
// swagger:route GET /attributevaluebooleans attributevaluebooleans getATTRIBUTEVALUEBOOLEANs
//
// # Get all attributevaluebooleans
//
// Responses:
// default: genericError
//
//	200: attributevaluebooleanDBResponse
func (controller *Controller) GetATTRIBUTEVALUEBOOLEANs(c *gin.Context) {

	// source slice
	var attributevaluebooleanDBs []orm.ATTRIBUTEVALUEBOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetDB()

	query := db.Find(&attributevaluebooleanDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevaluebooleanAPIs := make([]orm.ATTRIBUTEVALUEBOOLEANAPI, 0)

	// for each attributevalueboolean, update fields from the database nullable fields
	for idx := range attributevaluebooleanDBs {
		attributevaluebooleanDB := &attributevaluebooleanDBs[idx]
		_ = attributevaluebooleanDB
		var attributevaluebooleanAPI orm.ATTRIBUTEVALUEBOOLEANAPI

		// insertion point for updating fields
		attributevaluebooleanAPI.ID = attributevaluebooleanDB.ID
		attributevaluebooleanDB.CopyBasicFieldsToATTRIBUTEVALUEBOOLEAN_WOP(&attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEAN_WOP)
		attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEANPointersEncoding = attributevaluebooleanDB.ATTRIBUTEVALUEBOOLEANPointersEncoding
		attributevaluebooleanAPIs = append(attributevaluebooleanAPIs, attributevaluebooleanAPI)
	}

	c.JSON(http.StatusOK, attributevaluebooleanAPIs)
}

// PostATTRIBUTEVALUEBOOLEAN
//
// swagger:route POST /attributevaluebooleans attributevaluebooleans postATTRIBUTEVALUEBOOLEAN
//
// Creates a attributevalueboolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEVALUEBOOLEAN.Lock()
	defer mutexATTRIBUTEVALUEBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEBOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevalueboolean
	attributevaluebooleanDB := orm.ATTRIBUTEVALUEBOOLEANDB{}
	attributevaluebooleanDB.ATTRIBUTEVALUEBOOLEANPointersEncoding = input.ATTRIBUTEVALUEBOOLEANPointersEncoding
	attributevaluebooleanDB.CopyBasicFieldsFromATTRIBUTEVALUEBOOLEAN_WOP(&input.ATTRIBUTEVALUEBOOLEAN_WOP)

	query := db.Create(&attributevaluebooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.CheckoutPhaseOneInstance(&attributevaluebooleanDB)
	attributevalueboolean := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANPtr[attributevaluebooleanDB.ID]

	if attributevalueboolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevalueboolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevaluebooleanDB)
}

// GetATTRIBUTEVALUEBOOLEAN
//
// swagger:route GET /attributevaluebooleans/{ID} attributevaluebooleans getATTRIBUTEVALUEBOOLEAN
//
// Gets the details for a attributevalueboolean.
//
// Responses:
// default: genericError
//
//	200: attributevaluebooleanDBResponse
func (controller *Controller) GetATTRIBUTEVALUEBOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetDB()

	// Get attributevaluebooleanDB in DB
	var attributevaluebooleanDB orm.ATTRIBUTEVALUEBOOLEANDB
	if err := db.First(&attributevaluebooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevaluebooleanAPI orm.ATTRIBUTEVALUEBOOLEANAPI
	attributevaluebooleanAPI.ID = attributevaluebooleanDB.ID
	attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEANPointersEncoding = attributevaluebooleanDB.ATTRIBUTEVALUEBOOLEANPointersEncoding
	attributevaluebooleanDB.CopyBasicFieldsToATTRIBUTEVALUEBOOLEAN_WOP(&attributevaluebooleanAPI.ATTRIBUTEVALUEBOOLEAN_WOP)

	c.JSON(http.StatusOK, attributevaluebooleanAPI)
}

// UpdateATTRIBUTEVALUEBOOLEAN
//
// swagger:route PATCH /attributevaluebooleans/{ID} attributevaluebooleans updateATTRIBUTEVALUEBOOLEAN
//
// # Update a attributevalueboolean
//
// Responses:
// default: genericError
//
//	200: attributevaluebooleanDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEVALUEBOOLEAN.Lock()
	defer mutexATTRIBUTEVALUEBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEBOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevaluebooleanDB orm.ATTRIBUTEVALUEBOOLEANDB

	// fetch the attributevalueboolean
	query := db.First(&attributevaluebooleanDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevaluebooleanDB.CopyBasicFieldsFromATTRIBUTEVALUEBOOLEAN_WOP(&input.ATTRIBUTEVALUEBOOLEAN_WOP)
	attributevaluebooleanDB.ATTRIBUTEVALUEBOOLEANPointersEncoding = input.ATTRIBUTEVALUEBOOLEANPointersEncoding

	query = db.Model(&attributevaluebooleanDB).Updates(attributevaluebooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluebooleanNew := new(models.ATTRIBUTEVALUEBOOLEAN)
	attributevaluebooleanDB.CopyBasicFieldsToATTRIBUTEVALUEBOOLEAN(attributevaluebooleanNew)

	// redeem pointers
	attributevaluebooleanDB.DecodePointers(backRepo, attributevaluebooleanNew)

	// get stage instance from DB instance, and call callback function
	attributevaluebooleanOld := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANPtr[attributevaluebooleanDB.ID]
	if attributevaluebooleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevaluebooleanOld, attributevaluebooleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevaluebooleanDB
	c.JSON(http.StatusOK, attributevaluebooleanDB)
}

// DeleteATTRIBUTEVALUEBOOLEAN
//
// swagger:route DELETE /attributevaluebooleans/{ID} attributevaluebooleans deleteATTRIBUTEVALUEBOOLEAN
//
// # Delete a attributevalueboolean
//
// default: genericError
//
//	200: attributevaluebooleanDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEBOOLEAN(c *gin.Context) {

	mutexATTRIBUTEVALUEBOOLEAN.Lock()
	defer mutexATTRIBUTEVALUEBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.GetDB()

	// Get model if exist
	var attributevaluebooleanDB orm.ATTRIBUTEVALUEBOOLEANDB
	if err := db.First(&attributevaluebooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevaluebooleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluebooleanDeleted := new(models.ATTRIBUTEVALUEBOOLEAN)
	attributevaluebooleanDB.CopyBasicFieldsToATTRIBUTEVALUEBOOLEAN(attributevaluebooleanDeleted)

	// get stage instance from DB instance, and call callback function
	attributevaluebooleanStaged := backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANPtr[attributevaluebooleanDB.ID]
	if attributevaluebooleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevaluebooleanStaged, attributevaluebooleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
