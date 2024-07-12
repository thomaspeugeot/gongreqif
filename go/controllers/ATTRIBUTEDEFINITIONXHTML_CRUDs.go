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
var __ATTRIBUTEDEFINITIONXHTML__dummysDeclaration__ models.ATTRIBUTEDEFINITIONXHTML
var __ATTRIBUTEDEFINITIONXHTML_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONXHTML sync.Mutex

// An ATTRIBUTEDEFINITIONXHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONXHTML updateATTRIBUTEDEFINITIONXHTML deleteATTRIBUTEDEFINITIONXHTML
type ATTRIBUTEDEFINITIONXHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONXHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONXHTML updateATTRIBUTEDEFINITIONXHTML
type ATTRIBUTEDEFINITIONXHTMLInput struct {
	// The ATTRIBUTEDEFINITIONXHTML to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONXHTML *orm.ATTRIBUTEDEFINITIONXHTMLAPI
}

// GetATTRIBUTEDEFINITIONXHTMLs
//
// swagger:route GET /attributedefinitionxhtmls attributedefinitionxhtmls getATTRIBUTEDEFINITIONXHTMLs
//
// # Get all attributedefinitionxhtmls
//
// Responses:
// default: genericError
//
//	200: attributedefinitionxhtmlDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONXHTMLs(c *gin.Context) {

	// source slice
	var attributedefinitionxhtmlDBs []orm.ATTRIBUTEDEFINITIONXHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetDB()

	query := db.Find(&attributedefinitionxhtmlDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionxhtmlAPIs := make([]orm.ATTRIBUTEDEFINITIONXHTMLAPI, 0)

	// for each attributedefinitionxhtml, update fields from the database nullable fields
	for idx := range attributedefinitionxhtmlDBs {
		attributedefinitionxhtmlDB := &attributedefinitionxhtmlDBs[idx]
		_ = attributedefinitionxhtmlDB
		var attributedefinitionxhtmlAPI orm.ATTRIBUTEDEFINITIONXHTMLAPI

		// insertion point for updating fields
		attributedefinitionxhtmlAPI.ID = attributedefinitionxhtmlDB.ID
		attributedefinitionxhtmlDB.CopyBasicFieldsToATTRIBUTEDEFINITIONXHTML_WOP(&attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTML_WOP)
		attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTMLPointersEncoding = attributedefinitionxhtmlDB.ATTRIBUTEDEFINITIONXHTMLPointersEncoding
		attributedefinitionxhtmlAPIs = append(attributedefinitionxhtmlAPIs, attributedefinitionxhtmlAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionxhtmlAPIs)
}

// PostATTRIBUTEDEFINITIONXHTML
//
// swagger:route POST /attributedefinitionxhtmls attributedefinitionxhtmls postATTRIBUTEDEFINITIONXHTML
//
// Creates a attributedefinitionxhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONXHTML(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONXHTML.Lock()
	defer mutexATTRIBUTEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONXHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitionxhtml
	attributedefinitionxhtmlDB := orm.ATTRIBUTEDEFINITIONXHTMLDB{}
	attributedefinitionxhtmlDB.ATTRIBUTEDEFINITIONXHTMLPointersEncoding = input.ATTRIBUTEDEFINITIONXHTMLPointersEncoding
	attributedefinitionxhtmlDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONXHTML_WOP(&input.ATTRIBUTEDEFINITIONXHTML_WOP)

	query := db.Create(&attributedefinitionxhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.CheckoutPhaseOneInstance(&attributedefinitionxhtmlDB)
	attributedefinitionxhtml := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLPtr[attributedefinitionxhtmlDB.ID]

	if attributedefinitionxhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitionxhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionxhtmlDB)
}

// GetATTRIBUTEDEFINITIONXHTML
//
// swagger:route GET /attributedefinitionxhtmls/{ID} attributedefinitionxhtmls getATTRIBUTEDEFINITIONXHTML
//
// Gets the details for a attributedefinitionxhtml.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionxhtmlDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONXHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetDB()

	// Get attributedefinitionxhtmlDB in DB
	var attributedefinitionxhtmlDB orm.ATTRIBUTEDEFINITIONXHTMLDB
	if err := db.First(&attributedefinitionxhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionxhtmlAPI orm.ATTRIBUTEDEFINITIONXHTMLAPI
	attributedefinitionxhtmlAPI.ID = attributedefinitionxhtmlDB.ID
	attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTMLPointersEncoding = attributedefinitionxhtmlDB.ATTRIBUTEDEFINITIONXHTMLPointersEncoding
	attributedefinitionxhtmlDB.CopyBasicFieldsToATTRIBUTEDEFINITIONXHTML_WOP(&attributedefinitionxhtmlAPI.ATTRIBUTEDEFINITIONXHTML_WOP)

	c.JSON(http.StatusOK, attributedefinitionxhtmlAPI)
}

// UpdateATTRIBUTEDEFINITIONXHTML
//
// swagger:route PATCH /attributedefinitionxhtmls/{ID} attributedefinitionxhtmls updateATTRIBUTEDEFINITIONXHTML
//
// # Update a attributedefinitionxhtml
//
// Responses:
// default: genericError
//
//	200: attributedefinitionxhtmlDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONXHTML(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONXHTML.Lock()
	defer mutexATTRIBUTEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONXHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionxhtmlDB orm.ATTRIBUTEDEFINITIONXHTMLDB

	// fetch the attributedefinitionxhtml
	query := db.First(&attributedefinitionxhtmlDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionxhtmlDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONXHTML_WOP(&input.ATTRIBUTEDEFINITIONXHTML_WOP)
	attributedefinitionxhtmlDB.ATTRIBUTEDEFINITIONXHTMLPointersEncoding = input.ATTRIBUTEDEFINITIONXHTMLPointersEncoding

	query = db.Model(&attributedefinitionxhtmlDB).Updates(attributedefinitionxhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionxhtmlNew := new(models.ATTRIBUTEDEFINITIONXHTML)
	attributedefinitionxhtmlDB.CopyBasicFieldsToATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtmlNew)

	// redeem pointers
	attributedefinitionxhtmlDB.DecodePointers(backRepo, attributedefinitionxhtmlNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionxhtmlOld := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLPtr[attributedefinitionxhtmlDB.ID]
	if attributedefinitionxhtmlOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionxhtmlOld, attributedefinitionxhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionxhtmlDB
	c.JSON(http.StatusOK, attributedefinitionxhtmlDB)
}

// DeleteATTRIBUTEDEFINITIONXHTML
//
// swagger:route DELETE /attributedefinitionxhtmls/{ID} attributedefinitionxhtmls deleteATTRIBUTEDEFINITIONXHTML
//
// # Delete a attributedefinitionxhtml
//
// default: genericError
//
//	200: attributedefinitionxhtmlDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONXHTML(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONXHTML.Lock()
	defer mutexATTRIBUTEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.GetDB()

	// Get model if exist
	var attributedefinitionxhtmlDB orm.ATTRIBUTEDEFINITIONXHTMLDB
	if err := db.First(&attributedefinitionxhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionxhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionxhtmlDeleted := new(models.ATTRIBUTEDEFINITIONXHTML)
	attributedefinitionxhtmlDB.CopyBasicFieldsToATTRIBUTEDEFINITIONXHTML(attributedefinitionxhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionxhtmlStaged := backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLPtr[attributedefinitionxhtmlDB.ID]
	if attributedefinitionxhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionxhtmlStaged, attributedefinitionxhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
