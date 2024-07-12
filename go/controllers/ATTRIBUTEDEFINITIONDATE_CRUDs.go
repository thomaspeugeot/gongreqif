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
var __ATTRIBUTEDEFINITIONDATE__dummysDeclaration__ models.ATTRIBUTEDEFINITIONDATE
var __ATTRIBUTEDEFINITIONDATE_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONDATE sync.Mutex

// An ATTRIBUTEDEFINITIONDATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONDATE updateATTRIBUTEDEFINITIONDATE deleteATTRIBUTEDEFINITIONDATE
type ATTRIBUTEDEFINITIONDATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONDATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONDATE updateATTRIBUTEDEFINITIONDATE
type ATTRIBUTEDEFINITIONDATEInput struct {
	// The ATTRIBUTEDEFINITIONDATE to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONDATE *orm.ATTRIBUTEDEFINITIONDATEAPI
}

// GetATTRIBUTEDEFINITIONDATEs
//
// swagger:route GET /attributedefinitiondates attributedefinitiondates getATTRIBUTEDEFINITIONDATEs
//
// # Get all attributedefinitiondates
//
// Responses:
// default: genericError
//
//	200: attributedefinitiondateDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONDATEs(c *gin.Context) {

	// source slice
	var attributedefinitiondateDBs []orm.ATTRIBUTEDEFINITIONDATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetDB()

	query := db.Find(&attributedefinitiondateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitiondateAPIs := make([]orm.ATTRIBUTEDEFINITIONDATEAPI, 0)

	// for each attributedefinitiondate, update fields from the database nullable fields
	for idx := range attributedefinitiondateDBs {
		attributedefinitiondateDB := &attributedefinitiondateDBs[idx]
		_ = attributedefinitiondateDB
		var attributedefinitiondateAPI orm.ATTRIBUTEDEFINITIONDATEAPI

		// insertion point for updating fields
		attributedefinitiondateAPI.ID = attributedefinitiondateDB.ID
		attributedefinitiondateDB.CopyBasicFieldsToATTRIBUTEDEFINITIONDATE_WOP(&attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATE_WOP)
		attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATEPointersEncoding = attributedefinitiondateDB.ATTRIBUTEDEFINITIONDATEPointersEncoding
		attributedefinitiondateAPIs = append(attributedefinitiondateAPIs, attributedefinitiondateAPI)
	}

	c.JSON(http.StatusOK, attributedefinitiondateAPIs)
}

// PostATTRIBUTEDEFINITIONDATE
//
// swagger:route POST /attributedefinitiondates attributedefinitiondates postATTRIBUTEDEFINITIONDATE
//
// Creates a attributedefinitiondate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONDATE(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONDATE.Lock()
	defer mutexATTRIBUTEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONDATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitiondate
	attributedefinitiondateDB := orm.ATTRIBUTEDEFINITIONDATEDB{}
	attributedefinitiondateDB.ATTRIBUTEDEFINITIONDATEPointersEncoding = input.ATTRIBUTEDEFINITIONDATEPointersEncoding
	attributedefinitiondateDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONDATE_WOP(&input.ATTRIBUTEDEFINITIONDATE_WOP)

	query := db.Create(&attributedefinitiondateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.CheckoutPhaseOneInstance(&attributedefinitiondateDB)
	attributedefinitiondate := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEPtr[attributedefinitiondateDB.ID]

	if attributedefinitiondate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitiondate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitiondateDB)
}

// GetATTRIBUTEDEFINITIONDATE
//
// swagger:route GET /attributedefinitiondates/{ID} attributedefinitiondates getATTRIBUTEDEFINITIONDATE
//
// Gets the details for a attributedefinitiondate.
//
// Responses:
// default: genericError
//
//	200: attributedefinitiondateDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONDATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetDB()

	// Get attributedefinitiondateDB in DB
	var attributedefinitiondateDB orm.ATTRIBUTEDEFINITIONDATEDB
	if err := db.First(&attributedefinitiondateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitiondateAPI orm.ATTRIBUTEDEFINITIONDATEAPI
	attributedefinitiondateAPI.ID = attributedefinitiondateDB.ID
	attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATEPointersEncoding = attributedefinitiondateDB.ATTRIBUTEDEFINITIONDATEPointersEncoding
	attributedefinitiondateDB.CopyBasicFieldsToATTRIBUTEDEFINITIONDATE_WOP(&attributedefinitiondateAPI.ATTRIBUTEDEFINITIONDATE_WOP)

	c.JSON(http.StatusOK, attributedefinitiondateAPI)
}

// UpdateATTRIBUTEDEFINITIONDATE
//
// swagger:route PATCH /attributedefinitiondates/{ID} attributedefinitiondates updateATTRIBUTEDEFINITIONDATE
//
// # Update a attributedefinitiondate
//
// Responses:
// default: genericError
//
//	200: attributedefinitiondateDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONDATE(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONDATE.Lock()
	defer mutexATTRIBUTEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONDATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitiondateDB orm.ATTRIBUTEDEFINITIONDATEDB

	// fetch the attributedefinitiondate
	query := db.First(&attributedefinitiondateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitiondateDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONDATE_WOP(&input.ATTRIBUTEDEFINITIONDATE_WOP)
	attributedefinitiondateDB.ATTRIBUTEDEFINITIONDATEPointersEncoding = input.ATTRIBUTEDEFINITIONDATEPointersEncoding

	query = db.Model(&attributedefinitiondateDB).Updates(attributedefinitiondateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitiondateNew := new(models.ATTRIBUTEDEFINITIONDATE)
	attributedefinitiondateDB.CopyBasicFieldsToATTRIBUTEDEFINITIONDATE(attributedefinitiondateNew)

	// redeem pointers
	attributedefinitiondateDB.DecodePointers(backRepo, attributedefinitiondateNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitiondateOld := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEPtr[attributedefinitiondateDB.ID]
	if attributedefinitiondateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitiondateOld, attributedefinitiondateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitiondateDB
	c.JSON(http.StatusOK, attributedefinitiondateDB)
}

// DeleteATTRIBUTEDEFINITIONDATE
//
// swagger:route DELETE /attributedefinitiondates/{ID} attributedefinitiondates deleteATTRIBUTEDEFINITIONDATE
//
// # Delete a attributedefinitiondate
//
// default: genericError
//
//	200: attributedefinitiondateDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONDATE(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONDATE.Lock()
	defer mutexATTRIBUTEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.GetDB()

	// Get model if exist
	var attributedefinitiondateDB orm.ATTRIBUTEDEFINITIONDATEDB
	if err := db.First(&attributedefinitiondateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitiondateDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitiondateDeleted := new(models.ATTRIBUTEDEFINITIONDATE)
	attributedefinitiondateDB.CopyBasicFieldsToATTRIBUTEDEFINITIONDATE(attributedefinitiondateDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitiondateStaged := backRepo.BackRepoATTRIBUTEDEFINITIONDATE.Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEPtr[attributedefinitiondateDB.ID]
	if attributedefinitiondateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitiondateStaged, attributedefinitiondateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
