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
var __ATTRIBUTEVALUEINTEGER__dummysDeclaration__ models.ATTRIBUTEVALUEINTEGER
var __ATTRIBUTEVALUEINTEGER_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEINTEGER sync.Mutex

// An ATTRIBUTEVALUEINTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEINTEGER updateATTRIBUTEVALUEINTEGER deleteATTRIBUTEVALUEINTEGER
type ATTRIBUTEVALUEINTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEINTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEINTEGER updateATTRIBUTEVALUEINTEGER
type ATTRIBUTEVALUEINTEGERInput struct {
	// The ATTRIBUTEVALUEINTEGER to submit or modify
	// in: body
	ATTRIBUTEVALUEINTEGER *orm.ATTRIBUTEVALUEINTEGERAPI
}

// GetATTRIBUTEVALUEINTEGERs
//
// swagger:route GET /attributevalueintegers attributevalueintegers getATTRIBUTEVALUEINTEGERs
//
// # Get all attributevalueintegers
//
// Responses:
// default: genericError
//
//	200: attributevalueintegerDBResponse
func (controller *Controller) GetATTRIBUTEVALUEINTEGERs(c *gin.Context) {

	// source slice
	var attributevalueintegerDBs []orm.ATTRIBUTEVALUEINTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetDB()

	query := db.Find(&attributevalueintegerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevalueintegerAPIs := make([]orm.ATTRIBUTEVALUEINTEGERAPI, 0)

	// for each attributevalueinteger, update fields from the database nullable fields
	for idx := range attributevalueintegerDBs {
		attributevalueintegerDB := &attributevalueintegerDBs[idx]
		_ = attributevalueintegerDB
		var attributevalueintegerAPI orm.ATTRIBUTEVALUEINTEGERAPI

		// insertion point for updating fields
		attributevalueintegerAPI.ID = attributevalueintegerDB.ID
		attributevalueintegerDB.CopyBasicFieldsToATTRIBUTEVALUEINTEGER_WOP(&attributevalueintegerAPI.ATTRIBUTEVALUEINTEGER_WOP)
		attributevalueintegerAPI.ATTRIBUTEVALUEINTEGERPointersEncoding = attributevalueintegerDB.ATTRIBUTEVALUEINTEGERPointersEncoding
		attributevalueintegerAPIs = append(attributevalueintegerAPIs, attributevalueintegerAPI)
	}

	c.JSON(http.StatusOK, attributevalueintegerAPIs)
}

// PostATTRIBUTEVALUEINTEGER
//
// swagger:route POST /attributevalueintegers attributevalueintegers postATTRIBUTEVALUEINTEGER
//
// Creates a attributevalueinteger
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEINTEGER(c *gin.Context) {

	mutexATTRIBUTEVALUEINTEGER.Lock()
	defer mutexATTRIBUTEVALUEINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEINTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevalueinteger
	attributevalueintegerDB := orm.ATTRIBUTEVALUEINTEGERDB{}
	attributevalueintegerDB.ATTRIBUTEVALUEINTEGERPointersEncoding = input.ATTRIBUTEVALUEINTEGERPointersEncoding
	attributevalueintegerDB.CopyBasicFieldsFromATTRIBUTEVALUEINTEGER_WOP(&input.ATTRIBUTEVALUEINTEGER_WOP)

	query := db.Create(&attributevalueintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.CheckoutPhaseOneInstance(&attributevalueintegerDB)
	attributevalueinteger := backRepo.BackRepoATTRIBUTEVALUEINTEGER.Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERPtr[attributevalueintegerDB.ID]

	if attributevalueinteger != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevalueinteger)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevalueintegerDB)
}

// GetATTRIBUTEVALUEINTEGER
//
// swagger:route GET /attributevalueintegers/{ID} attributevalueintegers getATTRIBUTEVALUEINTEGER
//
// Gets the details for a attributevalueinteger.
//
// Responses:
// default: genericError
//
//	200: attributevalueintegerDBResponse
func (controller *Controller) GetATTRIBUTEVALUEINTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetDB()

	// Get attributevalueintegerDB in DB
	var attributevalueintegerDB orm.ATTRIBUTEVALUEINTEGERDB
	if err := db.First(&attributevalueintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevalueintegerAPI orm.ATTRIBUTEVALUEINTEGERAPI
	attributevalueintegerAPI.ID = attributevalueintegerDB.ID
	attributevalueintegerAPI.ATTRIBUTEVALUEINTEGERPointersEncoding = attributevalueintegerDB.ATTRIBUTEVALUEINTEGERPointersEncoding
	attributevalueintegerDB.CopyBasicFieldsToATTRIBUTEVALUEINTEGER_WOP(&attributevalueintegerAPI.ATTRIBUTEVALUEINTEGER_WOP)

	c.JSON(http.StatusOK, attributevalueintegerAPI)
}

// UpdateATTRIBUTEVALUEINTEGER
//
// swagger:route PATCH /attributevalueintegers/{ID} attributevalueintegers updateATTRIBUTEVALUEINTEGER
//
// # Update a attributevalueinteger
//
// Responses:
// default: genericError
//
//	200: attributevalueintegerDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEINTEGER(c *gin.Context) {

	mutexATTRIBUTEVALUEINTEGER.Lock()
	defer mutexATTRIBUTEVALUEINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEINTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevalueintegerDB orm.ATTRIBUTEVALUEINTEGERDB

	// fetch the attributevalueinteger
	query := db.First(&attributevalueintegerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevalueintegerDB.CopyBasicFieldsFromATTRIBUTEVALUEINTEGER_WOP(&input.ATTRIBUTEVALUEINTEGER_WOP)
	attributevalueintegerDB.ATTRIBUTEVALUEINTEGERPointersEncoding = input.ATTRIBUTEVALUEINTEGERPointersEncoding

	query = db.Model(&attributevalueintegerDB).Updates(attributevalueintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevalueintegerNew := new(models.ATTRIBUTEVALUEINTEGER)
	attributevalueintegerDB.CopyBasicFieldsToATTRIBUTEVALUEINTEGER(attributevalueintegerNew)

	// redeem pointers
	attributevalueintegerDB.DecodePointers(backRepo, attributevalueintegerNew)

	// get stage instance from DB instance, and call callback function
	attributevalueintegerOld := backRepo.BackRepoATTRIBUTEVALUEINTEGER.Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERPtr[attributevalueintegerDB.ID]
	if attributevalueintegerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevalueintegerOld, attributevalueintegerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevalueintegerDB
	c.JSON(http.StatusOK, attributevalueintegerDB)
}

// DeleteATTRIBUTEVALUEINTEGER
//
// swagger:route DELETE /attributevalueintegers/{ID} attributevalueintegers deleteATTRIBUTEVALUEINTEGER
//
// # Delete a attributevalueinteger
//
// default: genericError
//
//	200: attributevalueintegerDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEINTEGER(c *gin.Context) {

	mutexATTRIBUTEVALUEINTEGER.Lock()
	defer mutexATTRIBUTEVALUEINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEINTEGER.GetDB()

	// Get model if exist
	var attributevalueintegerDB orm.ATTRIBUTEVALUEINTEGERDB
	if err := db.First(&attributevalueintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevalueintegerDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevalueintegerDeleted := new(models.ATTRIBUTEVALUEINTEGER)
	attributevalueintegerDB.CopyBasicFieldsToATTRIBUTEVALUEINTEGER(attributevalueintegerDeleted)

	// get stage instance from DB instance, and call callback function
	attributevalueintegerStaged := backRepo.BackRepoATTRIBUTEVALUEINTEGER.Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERPtr[attributevalueintegerDB.ID]
	if attributevalueintegerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevalueintegerStaged, attributevalueintegerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
