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
var __DATATYPEDEFINITIONXHTML__dummysDeclaration__ models.DATATYPEDEFINITIONXHTML
var __DATATYPEDEFINITIONXHTML_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONXHTML sync.Mutex

// An DATATYPEDEFINITIONXHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONXHTML updateDATATYPEDEFINITIONXHTML deleteDATATYPEDEFINITIONXHTML
type DATATYPEDEFINITIONXHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONXHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONXHTML updateDATATYPEDEFINITIONXHTML
type DATATYPEDEFINITIONXHTMLInput struct {
	// The DATATYPEDEFINITIONXHTML to submit or modify
	// in: body
	DATATYPEDEFINITIONXHTML *orm.DATATYPEDEFINITIONXHTMLAPI
}

// GetDATATYPEDEFINITIONXHTMLs
//
// swagger:route GET /datatypedefinitionxhtmls datatypedefinitionxhtmls getDATATYPEDEFINITIONXHTMLs
//
// # Get all datatypedefinitionxhtmls
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionxhtmlDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONXHTMLs(c *gin.Context) {

	// source slice
	var datatypedefinitionxhtmlDBs []orm.DATATYPEDEFINITIONXHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDB()

	query := db.Find(&datatypedefinitionxhtmlDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionxhtmlAPIs := make([]orm.DATATYPEDEFINITIONXHTMLAPI, 0)

	// for each datatypedefinitionxhtml, update fields from the database nullable fields
	for idx := range datatypedefinitionxhtmlDBs {
		datatypedefinitionxhtmlDB := &datatypedefinitionxhtmlDBs[idx]
		_ = datatypedefinitionxhtmlDB
		var datatypedefinitionxhtmlAPI orm.DATATYPEDEFINITIONXHTMLAPI

		// insertion point for updating fields
		datatypedefinitionxhtmlAPI.ID = datatypedefinitionxhtmlDB.ID
		datatypedefinitionxhtmlDB.CopyBasicFieldsToDATATYPEDEFINITIONXHTML_WOP(&datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTML_WOP)
		datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTMLPointersEncoding = datatypedefinitionxhtmlDB.DATATYPEDEFINITIONXHTMLPointersEncoding
		datatypedefinitionxhtmlAPIs = append(datatypedefinitionxhtmlAPIs, datatypedefinitionxhtmlAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionxhtmlAPIs)
}

// PostDATATYPEDEFINITIONXHTML
//
// swagger:route POST /datatypedefinitionxhtmls datatypedefinitionxhtmls postDATATYPEDEFINITIONXHTML
//
// Creates a datatypedefinitionxhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONXHTML(c *gin.Context) {

	mutexDATATYPEDEFINITIONXHTML.Lock()
	defer mutexDATATYPEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONXHTMLs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONXHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitionxhtml
	datatypedefinitionxhtmlDB := orm.DATATYPEDEFINITIONXHTMLDB{}
	datatypedefinitionxhtmlDB.DATATYPEDEFINITIONXHTMLPointersEncoding = input.DATATYPEDEFINITIONXHTMLPointersEncoding
	datatypedefinitionxhtmlDB.CopyBasicFieldsFromDATATYPEDEFINITIONXHTML_WOP(&input.DATATYPEDEFINITIONXHTML_WOP)

	query := db.Create(&datatypedefinitionxhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.CheckoutPhaseOneInstance(&datatypedefinitionxhtmlDB)
	datatypedefinitionxhtml := backRepo.BackRepoDATATYPEDEFINITIONXHTML.Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLPtr[datatypedefinitionxhtmlDB.ID]

	if datatypedefinitionxhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitionxhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionxhtmlDB)
}

// GetDATATYPEDEFINITIONXHTML
//
// swagger:route GET /datatypedefinitionxhtmls/{ID} datatypedefinitionxhtmls getDATATYPEDEFINITIONXHTML
//
// Gets the details for a datatypedefinitionxhtml.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionxhtmlDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONXHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDB()

	// Get datatypedefinitionxhtmlDB in DB
	var datatypedefinitionxhtmlDB orm.DATATYPEDEFINITIONXHTMLDB
	if err := db.First(&datatypedefinitionxhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionxhtmlAPI orm.DATATYPEDEFINITIONXHTMLAPI
	datatypedefinitionxhtmlAPI.ID = datatypedefinitionxhtmlDB.ID
	datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTMLPointersEncoding = datatypedefinitionxhtmlDB.DATATYPEDEFINITIONXHTMLPointersEncoding
	datatypedefinitionxhtmlDB.CopyBasicFieldsToDATATYPEDEFINITIONXHTML_WOP(&datatypedefinitionxhtmlAPI.DATATYPEDEFINITIONXHTML_WOP)

	c.JSON(http.StatusOK, datatypedefinitionxhtmlAPI)
}

// UpdateDATATYPEDEFINITIONXHTML
//
// swagger:route PATCH /datatypedefinitionxhtmls/{ID} datatypedefinitionxhtmls updateDATATYPEDEFINITIONXHTML
//
// # Update a datatypedefinitionxhtml
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionxhtmlDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONXHTML(c *gin.Context) {

	mutexDATATYPEDEFINITIONXHTML.Lock()
	defer mutexDATATYPEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONXHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionxhtmlDB orm.DATATYPEDEFINITIONXHTMLDB

	// fetch the datatypedefinitionxhtml
	query := db.First(&datatypedefinitionxhtmlDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionxhtmlDB.CopyBasicFieldsFromDATATYPEDEFINITIONXHTML_WOP(&input.DATATYPEDEFINITIONXHTML_WOP)
	datatypedefinitionxhtmlDB.DATATYPEDEFINITIONXHTMLPointersEncoding = input.DATATYPEDEFINITIONXHTMLPointersEncoding

	query = db.Model(&datatypedefinitionxhtmlDB).Updates(datatypedefinitionxhtmlDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionxhtmlNew := new(models.DATATYPEDEFINITIONXHTML)
	datatypedefinitionxhtmlDB.CopyBasicFieldsToDATATYPEDEFINITIONXHTML(datatypedefinitionxhtmlNew)

	// redeem pointers
	datatypedefinitionxhtmlDB.DecodePointers(backRepo, datatypedefinitionxhtmlNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionxhtmlOld := backRepo.BackRepoDATATYPEDEFINITIONXHTML.Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLPtr[datatypedefinitionxhtmlDB.ID]
	if datatypedefinitionxhtmlOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionxhtmlOld, datatypedefinitionxhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionxhtmlDB
	c.JSON(http.StatusOK, datatypedefinitionxhtmlDB)
}

// DeleteDATATYPEDEFINITIONXHTML
//
// swagger:route DELETE /datatypedefinitionxhtmls/{ID} datatypedefinitionxhtmls deleteDATATYPEDEFINITIONXHTML
//
// # Delete a datatypedefinitionxhtml
//
// default: genericError
//
//	200: datatypedefinitionxhtmlDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONXHTML(c *gin.Context) {

	mutexDATATYPEDEFINITIONXHTML.Lock()
	defer mutexDATATYPEDEFINITIONXHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONXHTML", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONXHTML.GetDB()

	// Get model if exist
	var datatypedefinitionxhtmlDB orm.DATATYPEDEFINITIONXHTMLDB
	if err := db.First(&datatypedefinitionxhtmlDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionxhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionxhtmlDeleted := new(models.DATATYPEDEFINITIONXHTML)
	datatypedefinitionxhtmlDB.CopyBasicFieldsToDATATYPEDEFINITIONXHTML(datatypedefinitionxhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionxhtmlStaged := backRepo.BackRepoDATATYPEDEFINITIONXHTML.Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLPtr[datatypedefinitionxhtmlDB.ID]
	if datatypedefinitionxhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionxhtmlStaged, datatypedefinitionxhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
