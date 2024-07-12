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
var __ATTRIBUTEVALUEDATE__dummysDeclaration__ models.ATTRIBUTEVALUEDATE
var __ATTRIBUTEVALUEDATE_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEDATE sync.Mutex

// An ATTRIBUTEVALUEDATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEDATE updateATTRIBUTEVALUEDATE deleteATTRIBUTEVALUEDATE
type ATTRIBUTEVALUEDATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEDATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEDATE updateATTRIBUTEVALUEDATE
type ATTRIBUTEVALUEDATEInput struct {
	// The ATTRIBUTEVALUEDATE to submit or modify
	// in: body
	ATTRIBUTEVALUEDATE *orm.ATTRIBUTEVALUEDATEAPI
}

// GetATTRIBUTEVALUEDATEs
//
// swagger:route GET /attributevaluedates attributevaluedates getATTRIBUTEVALUEDATEs
//
// # Get all attributevaluedates
//
// Responses:
// default: genericError
//
//	200: attributevaluedateDBResponse
func (controller *Controller) GetATTRIBUTEVALUEDATEs(c *gin.Context) {

	// source slice
	var attributevaluedateDBs []orm.ATTRIBUTEVALUEDATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEDATE.GetDB()

	query := db.Find(&attributevaluedateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevaluedateAPIs := make([]orm.ATTRIBUTEVALUEDATEAPI, 0)

	// for each attributevaluedate, update fields from the database nullable fields
	for idx := range attributevaluedateDBs {
		attributevaluedateDB := &attributevaluedateDBs[idx]
		_ = attributevaluedateDB
		var attributevaluedateAPI orm.ATTRIBUTEVALUEDATEAPI

		// insertion point for updating fields
		attributevaluedateAPI.ID = attributevaluedateDB.ID
		attributevaluedateDB.CopyBasicFieldsToATTRIBUTEVALUEDATE_WOP(&attributevaluedateAPI.ATTRIBUTEVALUEDATE_WOP)
		attributevaluedateAPI.ATTRIBUTEVALUEDATEPointersEncoding = attributevaluedateDB.ATTRIBUTEVALUEDATEPointersEncoding
		attributevaluedateAPIs = append(attributevaluedateAPIs, attributevaluedateAPI)
	}

	c.JSON(http.StatusOK, attributevaluedateAPIs)
}

// PostATTRIBUTEVALUEDATE
//
// swagger:route POST /attributevaluedates attributevaluedates postATTRIBUTEVALUEDATE
//
// Creates a attributevaluedate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEDATE(c *gin.Context) {

	mutexATTRIBUTEVALUEDATE.Lock()
	defer mutexATTRIBUTEVALUEDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEDATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEDATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevaluedate
	attributevaluedateDB := orm.ATTRIBUTEVALUEDATEDB{}
	attributevaluedateDB.ATTRIBUTEVALUEDATEPointersEncoding = input.ATTRIBUTEVALUEDATEPointersEncoding
	attributevaluedateDB.CopyBasicFieldsFromATTRIBUTEVALUEDATE_WOP(&input.ATTRIBUTEVALUEDATE_WOP)

	query := db.Create(&attributevaluedateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEDATE.CheckoutPhaseOneInstance(&attributevaluedateDB)
	attributevaluedate := backRepo.BackRepoATTRIBUTEVALUEDATE.Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEPtr[attributevaluedateDB.ID]

	if attributevaluedate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevaluedate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevaluedateDB)
}

// GetATTRIBUTEVALUEDATE
//
// swagger:route GET /attributevaluedates/{ID} attributevaluedates getATTRIBUTEVALUEDATE
//
// Gets the details for a attributevaluedate.
//
// Responses:
// default: genericError
//
//	200: attributevaluedateDBResponse
func (controller *Controller) GetATTRIBUTEVALUEDATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEDATE.GetDB()

	// Get attributevaluedateDB in DB
	var attributevaluedateDB orm.ATTRIBUTEVALUEDATEDB
	if err := db.First(&attributevaluedateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevaluedateAPI orm.ATTRIBUTEVALUEDATEAPI
	attributevaluedateAPI.ID = attributevaluedateDB.ID
	attributevaluedateAPI.ATTRIBUTEVALUEDATEPointersEncoding = attributevaluedateDB.ATTRIBUTEVALUEDATEPointersEncoding
	attributevaluedateDB.CopyBasicFieldsToATTRIBUTEVALUEDATE_WOP(&attributevaluedateAPI.ATTRIBUTEVALUEDATE_WOP)

	c.JSON(http.StatusOK, attributevaluedateAPI)
}

// UpdateATTRIBUTEVALUEDATE
//
// swagger:route PATCH /attributevaluedates/{ID} attributevaluedates updateATTRIBUTEVALUEDATE
//
// # Update a attributevaluedate
//
// Responses:
// default: genericError
//
//	200: attributevaluedateDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEDATE(c *gin.Context) {

	mutexATTRIBUTEVALUEDATE.Lock()
	defer mutexATTRIBUTEVALUEDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEDATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEDATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevaluedateDB orm.ATTRIBUTEVALUEDATEDB

	// fetch the attributevaluedate
	query := db.First(&attributevaluedateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevaluedateDB.CopyBasicFieldsFromATTRIBUTEVALUEDATE_WOP(&input.ATTRIBUTEVALUEDATE_WOP)
	attributevaluedateDB.ATTRIBUTEVALUEDATEPointersEncoding = input.ATTRIBUTEVALUEDATEPointersEncoding

	query = db.Model(&attributevaluedateDB).Updates(attributevaluedateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluedateNew := new(models.ATTRIBUTEVALUEDATE)
	attributevaluedateDB.CopyBasicFieldsToATTRIBUTEVALUEDATE(attributevaluedateNew)

	// redeem pointers
	attributevaluedateDB.DecodePointers(backRepo, attributevaluedateNew)

	// get stage instance from DB instance, and call callback function
	attributevaluedateOld := backRepo.BackRepoATTRIBUTEVALUEDATE.Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEPtr[attributevaluedateDB.ID]
	if attributevaluedateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevaluedateOld, attributevaluedateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevaluedateDB
	c.JSON(http.StatusOK, attributevaluedateDB)
}

// DeleteATTRIBUTEVALUEDATE
//
// swagger:route DELETE /attributevaluedates/{ID} attributevaluedates deleteATTRIBUTEVALUEDATE
//
// # Delete a attributevaluedate
//
// default: genericError
//
//	200: attributevaluedateDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEDATE(c *gin.Context) {

	mutexATTRIBUTEVALUEDATE.Lock()
	defer mutexATTRIBUTEVALUEDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEDATE.GetDB()

	// Get model if exist
	var attributevaluedateDB orm.ATTRIBUTEVALUEDATEDB
	if err := db.First(&attributevaluedateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevaluedateDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluedateDeleted := new(models.ATTRIBUTEVALUEDATE)
	attributevaluedateDB.CopyBasicFieldsToATTRIBUTEVALUEDATE(attributevaluedateDeleted)

	// get stage instance from DB instance, and call callback function
	attributevaluedateStaged := backRepo.BackRepoATTRIBUTEVALUEDATE.Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEPtr[attributevaluedateDB.ID]
	if attributevaluedateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevaluedateStaged, attributevaluedateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
