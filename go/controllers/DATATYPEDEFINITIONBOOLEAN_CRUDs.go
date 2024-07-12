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
var __DATATYPEDEFINITIONBOOLEAN__dummysDeclaration__ models.DATATYPEDEFINITIONBOOLEAN
var __DATATYPEDEFINITIONBOOLEAN_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONBOOLEAN sync.Mutex

// An DATATYPEDEFINITIONBOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONBOOLEAN updateDATATYPEDEFINITIONBOOLEAN deleteDATATYPEDEFINITIONBOOLEAN
type DATATYPEDEFINITIONBOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONBOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONBOOLEAN updateDATATYPEDEFINITIONBOOLEAN
type DATATYPEDEFINITIONBOOLEANInput struct {
	// The DATATYPEDEFINITIONBOOLEAN to submit or modify
	// in: body
	DATATYPEDEFINITIONBOOLEAN *orm.DATATYPEDEFINITIONBOOLEANAPI
}

// GetDATATYPEDEFINITIONBOOLEANs
//
// swagger:route GET /datatypedefinitionbooleans datatypedefinitionbooleans getDATATYPEDEFINITIONBOOLEANs
//
// # Get all datatypedefinitionbooleans
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionbooleanDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONBOOLEANs(c *gin.Context) {

	// source slice
	var datatypedefinitionbooleanDBs []orm.DATATYPEDEFINITIONBOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDB()

	query := db.Find(&datatypedefinitionbooleanDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionbooleanAPIs := make([]orm.DATATYPEDEFINITIONBOOLEANAPI, 0)

	// for each datatypedefinitionboolean, update fields from the database nullable fields
	for idx := range datatypedefinitionbooleanDBs {
		datatypedefinitionbooleanDB := &datatypedefinitionbooleanDBs[idx]
		_ = datatypedefinitionbooleanDB
		var datatypedefinitionbooleanAPI orm.DATATYPEDEFINITIONBOOLEANAPI

		// insertion point for updating fields
		datatypedefinitionbooleanAPI.ID = datatypedefinitionbooleanDB.ID
		datatypedefinitionbooleanDB.CopyBasicFieldsToDATATYPEDEFINITIONBOOLEAN_WOP(&datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEAN_WOP)
		datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEANPointersEncoding = datatypedefinitionbooleanDB.DATATYPEDEFINITIONBOOLEANPointersEncoding
		datatypedefinitionbooleanAPIs = append(datatypedefinitionbooleanAPIs, datatypedefinitionbooleanAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionbooleanAPIs)
}

// PostDATATYPEDEFINITIONBOOLEAN
//
// swagger:route POST /datatypedefinitionbooleans datatypedefinitionbooleans postDATATYPEDEFINITIONBOOLEAN
//
// Creates a datatypedefinitionboolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexDATATYPEDEFINITIONBOOLEAN.Lock()
	defer mutexDATATYPEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONBOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONBOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitionboolean
	datatypedefinitionbooleanDB := orm.DATATYPEDEFINITIONBOOLEANDB{}
	datatypedefinitionbooleanDB.DATATYPEDEFINITIONBOOLEANPointersEncoding = input.DATATYPEDEFINITIONBOOLEANPointersEncoding
	datatypedefinitionbooleanDB.CopyBasicFieldsFromDATATYPEDEFINITIONBOOLEAN_WOP(&input.DATATYPEDEFINITIONBOOLEAN_WOP)

	query := db.Create(&datatypedefinitionbooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.CheckoutPhaseOneInstance(&datatypedefinitionbooleanDB)
	datatypedefinitionboolean := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANPtr[datatypedefinitionbooleanDB.ID]

	if datatypedefinitionboolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitionboolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionbooleanDB)
}

// GetDATATYPEDEFINITIONBOOLEAN
//
// swagger:route GET /datatypedefinitionbooleans/{ID} datatypedefinitionbooleans getDATATYPEDEFINITIONBOOLEAN
//
// Gets the details for a datatypedefinitionboolean.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionbooleanDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONBOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDB()

	// Get datatypedefinitionbooleanDB in DB
	var datatypedefinitionbooleanDB orm.DATATYPEDEFINITIONBOOLEANDB
	if err := db.First(&datatypedefinitionbooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionbooleanAPI orm.DATATYPEDEFINITIONBOOLEANAPI
	datatypedefinitionbooleanAPI.ID = datatypedefinitionbooleanDB.ID
	datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEANPointersEncoding = datatypedefinitionbooleanDB.DATATYPEDEFINITIONBOOLEANPointersEncoding
	datatypedefinitionbooleanDB.CopyBasicFieldsToDATATYPEDEFINITIONBOOLEAN_WOP(&datatypedefinitionbooleanAPI.DATATYPEDEFINITIONBOOLEAN_WOP)

	c.JSON(http.StatusOK, datatypedefinitionbooleanAPI)
}

// UpdateDATATYPEDEFINITIONBOOLEAN
//
// swagger:route PATCH /datatypedefinitionbooleans/{ID} datatypedefinitionbooleans updateDATATYPEDEFINITIONBOOLEAN
//
// # Update a datatypedefinitionboolean
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionbooleanDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexDATATYPEDEFINITIONBOOLEAN.Lock()
	defer mutexDATATYPEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONBOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionbooleanDB orm.DATATYPEDEFINITIONBOOLEANDB

	// fetch the datatypedefinitionboolean
	query := db.First(&datatypedefinitionbooleanDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionbooleanDB.CopyBasicFieldsFromDATATYPEDEFINITIONBOOLEAN_WOP(&input.DATATYPEDEFINITIONBOOLEAN_WOP)
	datatypedefinitionbooleanDB.DATATYPEDEFINITIONBOOLEANPointersEncoding = input.DATATYPEDEFINITIONBOOLEANPointersEncoding

	query = db.Model(&datatypedefinitionbooleanDB).Updates(datatypedefinitionbooleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionbooleanNew := new(models.DATATYPEDEFINITIONBOOLEAN)
	datatypedefinitionbooleanDB.CopyBasicFieldsToDATATYPEDEFINITIONBOOLEAN(datatypedefinitionbooleanNew)

	// redeem pointers
	datatypedefinitionbooleanDB.DecodePointers(backRepo, datatypedefinitionbooleanNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionbooleanOld := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANPtr[datatypedefinitionbooleanDB.ID]
	if datatypedefinitionbooleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionbooleanOld, datatypedefinitionbooleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionbooleanDB
	c.JSON(http.StatusOK, datatypedefinitionbooleanDB)
}

// DeleteDATATYPEDEFINITIONBOOLEAN
//
// swagger:route DELETE /datatypedefinitionbooleans/{ID} datatypedefinitionbooleans deleteDATATYPEDEFINITIONBOOLEAN
//
// # Delete a datatypedefinitionboolean
//
// default: genericError
//
//	200: datatypedefinitionbooleanDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONBOOLEAN(c *gin.Context) {

	mutexDATATYPEDEFINITIONBOOLEAN.Lock()
	defer mutexDATATYPEDEFINITIONBOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONBOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.GetDB()

	// Get model if exist
	var datatypedefinitionbooleanDB orm.DATATYPEDEFINITIONBOOLEANDB
	if err := db.First(&datatypedefinitionbooleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionbooleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionbooleanDeleted := new(models.DATATYPEDEFINITIONBOOLEAN)
	datatypedefinitionbooleanDB.CopyBasicFieldsToDATATYPEDEFINITIONBOOLEAN(datatypedefinitionbooleanDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionbooleanStaged := backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANPtr[datatypedefinitionbooleanDB.ID]
	if datatypedefinitionbooleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionbooleanStaged, datatypedefinitionbooleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
