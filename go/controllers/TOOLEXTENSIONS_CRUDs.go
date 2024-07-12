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
var __TOOLEXTENSIONS__dummysDeclaration__ models.TOOLEXTENSIONS
var __TOOLEXTENSIONS_time__dummyDeclaration time.Duration

var mutexTOOLEXTENSIONS sync.Mutex

// An TOOLEXTENSIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTOOLEXTENSIONS updateTOOLEXTENSIONS deleteTOOLEXTENSIONS
type TOOLEXTENSIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TOOLEXTENSIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTOOLEXTENSIONS updateTOOLEXTENSIONS
type TOOLEXTENSIONSInput struct {
	// The TOOLEXTENSIONS to submit or modify
	// in: body
	TOOLEXTENSIONS *orm.TOOLEXTENSIONSAPI
}

// GetTOOLEXTENSIONSs
//
// swagger:route GET /toolextensionss toolextensionss getTOOLEXTENSIONSs
//
// # Get all toolextensionss
//
// Responses:
// default: genericError
//
//	200: toolextensionsDBResponse
func (controller *Controller) GetTOOLEXTENSIONSs(c *gin.Context) {

	// source slice
	var toolextensionsDBs []orm.TOOLEXTENSIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTOOLEXTENSIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTOOLEXTENSIONS.GetDB()

	query := db.Find(&toolextensionsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	toolextensionsAPIs := make([]orm.TOOLEXTENSIONSAPI, 0)

	// for each toolextensions, update fields from the database nullable fields
	for idx := range toolextensionsDBs {
		toolextensionsDB := &toolextensionsDBs[idx]
		_ = toolextensionsDB
		var toolextensionsAPI orm.TOOLEXTENSIONSAPI

		// insertion point for updating fields
		toolextensionsAPI.ID = toolextensionsDB.ID
		toolextensionsDB.CopyBasicFieldsToTOOLEXTENSIONS_WOP(&toolextensionsAPI.TOOLEXTENSIONS_WOP)
		toolextensionsAPI.TOOLEXTENSIONSPointersEncoding = toolextensionsDB.TOOLEXTENSIONSPointersEncoding
		toolextensionsAPIs = append(toolextensionsAPIs, toolextensionsAPI)
	}

	c.JSON(http.StatusOK, toolextensionsAPIs)
}

// PostTOOLEXTENSIONS
//
// swagger:route POST /toolextensionss toolextensionss postTOOLEXTENSIONS
//
// Creates a toolextensions
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTOOLEXTENSIONS(c *gin.Context) {

	mutexTOOLEXTENSIONS.Lock()
	defer mutexTOOLEXTENSIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTOOLEXTENSIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTOOLEXTENSIONS.GetDB()

	// Validate input
	var input orm.TOOLEXTENSIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create toolextensions
	toolextensionsDB := orm.TOOLEXTENSIONSDB{}
	toolextensionsDB.TOOLEXTENSIONSPointersEncoding = input.TOOLEXTENSIONSPointersEncoding
	toolextensionsDB.CopyBasicFieldsFromTOOLEXTENSIONS_WOP(&input.TOOLEXTENSIONS_WOP)

	query := db.Create(&toolextensionsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTOOLEXTENSIONS.CheckoutPhaseOneInstance(&toolextensionsDB)
	toolextensions := backRepo.BackRepoTOOLEXTENSIONS.Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSPtr[toolextensionsDB.ID]

	if toolextensions != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), toolextensions)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, toolextensionsDB)
}

// GetTOOLEXTENSIONS
//
// swagger:route GET /toolextensionss/{ID} toolextensionss getTOOLEXTENSIONS
//
// Gets the details for a toolextensions.
//
// Responses:
// default: genericError
//
//	200: toolextensionsDBResponse
func (controller *Controller) GetTOOLEXTENSIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTOOLEXTENSIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTOOLEXTENSIONS.GetDB()

	// Get toolextensionsDB in DB
	var toolextensionsDB orm.TOOLEXTENSIONSDB
	if err := db.First(&toolextensionsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var toolextensionsAPI orm.TOOLEXTENSIONSAPI
	toolextensionsAPI.ID = toolextensionsDB.ID
	toolextensionsAPI.TOOLEXTENSIONSPointersEncoding = toolextensionsDB.TOOLEXTENSIONSPointersEncoding
	toolextensionsDB.CopyBasicFieldsToTOOLEXTENSIONS_WOP(&toolextensionsAPI.TOOLEXTENSIONS_WOP)

	c.JSON(http.StatusOK, toolextensionsAPI)
}

// UpdateTOOLEXTENSIONS
//
// swagger:route PATCH /toolextensionss/{ID} toolextensionss updateTOOLEXTENSIONS
//
// # Update a toolextensions
//
// Responses:
// default: genericError
//
//	200: toolextensionsDBResponse
func (controller *Controller) UpdateTOOLEXTENSIONS(c *gin.Context) {

	mutexTOOLEXTENSIONS.Lock()
	defer mutexTOOLEXTENSIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTOOLEXTENSIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTOOLEXTENSIONS.GetDB()

	// Validate input
	var input orm.TOOLEXTENSIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var toolextensionsDB orm.TOOLEXTENSIONSDB

	// fetch the toolextensions
	query := db.First(&toolextensionsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	toolextensionsDB.CopyBasicFieldsFromTOOLEXTENSIONS_WOP(&input.TOOLEXTENSIONS_WOP)
	toolextensionsDB.TOOLEXTENSIONSPointersEncoding = input.TOOLEXTENSIONSPointersEncoding

	query = db.Model(&toolextensionsDB).Updates(toolextensionsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	toolextensionsNew := new(models.TOOLEXTENSIONS)
	toolextensionsDB.CopyBasicFieldsToTOOLEXTENSIONS(toolextensionsNew)

	// redeem pointers
	toolextensionsDB.DecodePointers(backRepo, toolextensionsNew)

	// get stage instance from DB instance, and call callback function
	toolextensionsOld := backRepo.BackRepoTOOLEXTENSIONS.Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSPtr[toolextensionsDB.ID]
	if toolextensionsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), toolextensionsOld, toolextensionsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the toolextensionsDB
	c.JSON(http.StatusOK, toolextensionsDB)
}

// DeleteTOOLEXTENSIONS
//
// swagger:route DELETE /toolextensionss/{ID} toolextensionss deleteTOOLEXTENSIONS
//
// # Delete a toolextensions
//
// default: genericError
//
//	200: toolextensionsDBResponse
func (controller *Controller) DeleteTOOLEXTENSIONS(c *gin.Context) {

	mutexTOOLEXTENSIONS.Lock()
	defer mutexTOOLEXTENSIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTOOLEXTENSIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTOOLEXTENSIONS.GetDB()

	// Get model if exist
	var toolextensionsDB orm.TOOLEXTENSIONSDB
	if err := db.First(&toolextensionsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&toolextensionsDB)

	// get an instance (not staged) from DB instance, and call callback function
	toolextensionsDeleted := new(models.TOOLEXTENSIONS)
	toolextensionsDB.CopyBasicFieldsToTOOLEXTENSIONS(toolextensionsDeleted)

	// get stage instance from DB instance, and call callback function
	toolextensionsStaged := backRepo.BackRepoTOOLEXTENSIONS.Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSPtr[toolextensionsDB.ID]
	if toolextensionsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), toolextensionsStaged, toolextensionsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
