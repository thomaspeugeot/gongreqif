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
var __ATTRIBUTEVALUEXHTML__dummysDeclaration__ models.ATTRIBUTEVALUEXHTML
var __ATTRIBUTEVALUEXHTML_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEXHTML sync.Mutex

// An ATTRIBUTEVALUEXHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEXHTML updateATTRIBUTEVALUEXHTML deleteATTRIBUTEVALUEXHTML
type ATTRIBUTEVALUEXHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEXHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEXHTML updateATTRIBUTEVALUEXHTML
type ATTRIBUTEVALUEXHTMLInput struct {
	// The ATTRIBUTEVALUEXHTML to submit or modify
	// in: body
	ATTRIBUTEVALUEXHTML *orm.ATTRIBUTEVALUEXHTMLAPI
}

// GetATTRIBUTEVALUEXHTMLs
//
// swagger:route GET /attributevaluexhtmls attributevaluexhtmls getATTRIBUTEVALUEXHTMLs
//
// # Get all attributevaluexhtmls
//
// Responses:
// default: genericError
//
//	200: attributevaluexhtmlDBResponse
func (controller *Controller) GetATTRIBUTEVALUEXHTMLs(c *gin.Context) {

	// source slice
	var attributevaluexhtmlDBs []orm.ATTRIBUTEVALUEXHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetDB()

	query := db.Find(&attributevaluexhtmlDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevaluexhtmlAPIs := make([]orm.ATTRIBUTEVALUEXHTMLAPI, 0)

	// for each attributevaluexhtml, update fields from the database nullable fields
	for idx := range attributevaluexhtmlDBs {
		attributevaluexhtmlDB := &attributevaluexhtmlDBs[idx]
		_ = attributevaluexhtmlDB
		var attributevaluexhtmlAPI orm.ATTRIBUTEVALUEXHTMLAPI

		// insertion point for updating fields
		attributevaluexhtmlAPI.ID = attributevaluexhtmlDB.ID
		attributevaluexhtmlDB.CopyBasicFieldsToATTRIBUTEVALUEXHTML_WOP(&attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTML_WOP)
		attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTMLPointersEncoding = attributevaluexhtmlDB.ATTRIBUTEVALUEXHTMLPointersEncoding
		attributevaluexhtmlAPIs = append(attributevaluexhtmlAPIs, attributevaluexhtmlAPI)
	}

	c.JSON(http.StatusOK, attributevaluexhtmlAPIs)
}

// PostATTRIBUTEVALUEXHTML
//
// swagger:route POST /attributevaluexhtmls attributevaluexhtmls postATTRIBUTEVALUEXHTML
//
// Creates a attributevaluexhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEXHTML(c *gin.Context) {

	mutexATTRIBUTEVALUEXHTML.Lock()
	defer mutexATTRIBUTEVALUEXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEXHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevaluexhtml
	attributevaluexhtmlDB := orm.ATTRIBUTEVALUEXHTMLDB{}
	attributevaluexhtmlDB.ATTRIBUTEVALUEXHTMLPointersEncoding = input.ATTRIBUTEVALUEXHTMLPointersEncoding
	attributevaluexhtmlDB.CopyBasicFieldsFromATTRIBUTEVALUEXHTML_WOP(&input.ATTRIBUTEVALUEXHTML_WOP)

	query := db.Create(&attributevaluexhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEXHTML.CheckoutPhaseOneInstance(&attributevaluexhtmlDB)
	attributevaluexhtml := backRepo.BackRepoATTRIBUTEVALUEXHTML.Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLPtr[attributevaluexhtmlDB.ID]

	if attributevaluexhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevaluexhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevaluexhtmlDB)
}

// GetATTRIBUTEVALUEXHTML
//
// swagger:route GET /attributevaluexhtmls/{ID} attributevaluexhtmls getATTRIBUTEVALUEXHTML
//
// Gets the details for a attributevaluexhtml.
//
// Responses:
// default: genericError
//
//	200: attributevaluexhtmlDBResponse
func (controller *Controller) GetATTRIBUTEVALUEXHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetDB()

	// Get attributevaluexhtmlDB in DB
	var attributevaluexhtmlDB orm.ATTRIBUTEVALUEXHTMLDB
	if err := db.First(&attributevaluexhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevaluexhtmlAPI orm.ATTRIBUTEVALUEXHTMLAPI
	attributevaluexhtmlAPI.ID = attributevaluexhtmlDB.ID
	attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTMLPointersEncoding = attributevaluexhtmlDB.ATTRIBUTEVALUEXHTMLPointersEncoding
	attributevaluexhtmlDB.CopyBasicFieldsToATTRIBUTEVALUEXHTML_WOP(&attributevaluexhtmlAPI.ATTRIBUTEVALUEXHTML_WOP)

	c.JSON(http.StatusOK, attributevaluexhtmlAPI)
}

// UpdateATTRIBUTEVALUEXHTML
//
// swagger:route PATCH /attributevaluexhtmls/{ID} attributevaluexhtmls updateATTRIBUTEVALUEXHTML
//
// # Update a attributevaluexhtml
//
// Responses:
// default: genericError
//
//	200: attributevaluexhtmlDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEXHTML(c *gin.Context) {

	mutexATTRIBUTEVALUEXHTML.Lock()
	defer mutexATTRIBUTEVALUEXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEXHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevaluexhtmlDB orm.ATTRIBUTEVALUEXHTMLDB

	// fetch the attributevaluexhtml
	query := db.First(&attributevaluexhtmlDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevaluexhtmlDB.CopyBasicFieldsFromATTRIBUTEVALUEXHTML_WOP(&input.ATTRIBUTEVALUEXHTML_WOP)
	attributevaluexhtmlDB.ATTRIBUTEVALUEXHTMLPointersEncoding = input.ATTRIBUTEVALUEXHTMLPointersEncoding

	query = db.Model(&attributevaluexhtmlDB).Updates(attributevaluexhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluexhtmlNew := new(models.ATTRIBUTEVALUEXHTML)
	attributevaluexhtmlDB.CopyBasicFieldsToATTRIBUTEVALUEXHTML(attributevaluexhtmlNew)

	// redeem pointers
	attributevaluexhtmlDB.DecodePointers(backRepo, attributevaluexhtmlNew)

	// get stage instance from DB instance, and call callback function
	attributevaluexhtmlOld := backRepo.BackRepoATTRIBUTEVALUEXHTML.Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLPtr[attributevaluexhtmlDB.ID]
	if attributevaluexhtmlOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevaluexhtmlOld, attributevaluexhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevaluexhtmlDB
	c.JSON(http.StatusOK, attributevaluexhtmlDB)
}

// DeleteATTRIBUTEVALUEXHTML
//
// swagger:route DELETE /attributevaluexhtmls/{ID} attributevaluexhtmls deleteATTRIBUTEVALUEXHTML
//
// # Delete a attributevaluexhtml
//
// default: genericError
//
//	200: attributevaluexhtmlDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEXHTML(c *gin.Context) {

	mutexATTRIBUTEVALUEXHTML.Lock()
	defer mutexATTRIBUTEVALUEXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEXHTML.GetDB()

	// Get model if exist
	var attributevaluexhtmlDB orm.ATTRIBUTEVALUEXHTMLDB
	if err := db.First(&attributevaluexhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevaluexhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluexhtmlDeleted := new(models.ATTRIBUTEVALUEXHTML)
	attributevaluexhtmlDB.CopyBasicFieldsToATTRIBUTEVALUEXHTML(attributevaluexhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	attributevaluexhtmlStaged := backRepo.BackRepoATTRIBUTEVALUEXHTML.Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLPtr[attributevaluexhtmlDB.ID]
	if attributevaluexhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevaluexhtmlStaged, attributevaluexhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
