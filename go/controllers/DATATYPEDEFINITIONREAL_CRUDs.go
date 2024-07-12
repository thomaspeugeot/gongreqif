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
var __DATATYPEDEFINITIONREAL__dummysDeclaration__ models.DATATYPEDEFINITIONREAL
var __DATATYPEDEFINITIONREAL_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONREAL sync.Mutex

// An DATATYPEDEFINITIONREALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONREAL updateDATATYPEDEFINITIONREAL deleteDATATYPEDEFINITIONREAL
type DATATYPEDEFINITIONREALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONREALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONREAL updateDATATYPEDEFINITIONREAL
type DATATYPEDEFINITIONREALInput struct {
	// The DATATYPEDEFINITIONREAL to submit or modify
	// in: body
	DATATYPEDEFINITIONREAL *orm.DATATYPEDEFINITIONREALAPI
}

// GetDATATYPEDEFINITIONREALs
//
// swagger:route GET /datatypedefinitionreals datatypedefinitionreals getDATATYPEDEFINITIONREALs
//
// # Get all datatypedefinitionreals
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionrealDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONREALs(c *gin.Context) {

	// source slice
	var datatypedefinitionrealDBs []orm.DATATYPEDEFINITIONREALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDB()

	query := db.Find(&datatypedefinitionrealDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionrealAPIs := make([]orm.DATATYPEDEFINITIONREALAPI, 0)

	// for each datatypedefinitionreal, update fields from the database nullable fields
	for idx := range datatypedefinitionrealDBs {
		datatypedefinitionrealDB := &datatypedefinitionrealDBs[idx]
		_ = datatypedefinitionrealDB
		var datatypedefinitionrealAPI orm.DATATYPEDEFINITIONREALAPI

		// insertion point for updating fields
		datatypedefinitionrealAPI.ID = datatypedefinitionrealDB.ID
		datatypedefinitionrealDB.CopyBasicFieldsToDATATYPEDEFINITIONREAL_WOP(&datatypedefinitionrealAPI.DATATYPEDEFINITIONREAL_WOP)
		datatypedefinitionrealAPI.DATATYPEDEFINITIONREALPointersEncoding = datatypedefinitionrealDB.DATATYPEDEFINITIONREALPointersEncoding
		datatypedefinitionrealAPIs = append(datatypedefinitionrealAPIs, datatypedefinitionrealAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionrealAPIs)
}

// PostDATATYPEDEFINITIONREAL
//
// swagger:route POST /datatypedefinitionreals datatypedefinitionreals postDATATYPEDEFINITIONREAL
//
// Creates a datatypedefinitionreal
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONREAL(c *gin.Context) {

	mutexDATATYPEDEFINITIONREAL.Lock()
	defer mutexDATATYPEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONREALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitionreal
	datatypedefinitionrealDB := orm.DATATYPEDEFINITIONREALDB{}
	datatypedefinitionrealDB.DATATYPEDEFINITIONREALPointersEncoding = input.DATATYPEDEFINITIONREALPointersEncoding
	datatypedefinitionrealDB.CopyBasicFieldsFromDATATYPEDEFINITIONREAL_WOP(&input.DATATYPEDEFINITIONREAL_WOP)

	query := db.Create(&datatypedefinitionrealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONREAL.CheckoutPhaseOneInstance(&datatypedefinitionrealDB)
	datatypedefinitionreal := backRepo.BackRepoDATATYPEDEFINITIONREAL.Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALPtr[datatypedefinitionrealDB.ID]

	if datatypedefinitionreal != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitionreal)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionrealDB)
}

// GetDATATYPEDEFINITIONREAL
//
// swagger:route GET /datatypedefinitionreals/{ID} datatypedefinitionreals getDATATYPEDEFINITIONREAL
//
// Gets the details for a datatypedefinitionreal.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionrealDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONREAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDB()

	// Get datatypedefinitionrealDB in DB
	var datatypedefinitionrealDB orm.DATATYPEDEFINITIONREALDB
	if err := db.First(&datatypedefinitionrealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionrealAPI orm.DATATYPEDEFINITIONREALAPI
	datatypedefinitionrealAPI.ID = datatypedefinitionrealDB.ID
	datatypedefinitionrealAPI.DATATYPEDEFINITIONREALPointersEncoding = datatypedefinitionrealDB.DATATYPEDEFINITIONREALPointersEncoding
	datatypedefinitionrealDB.CopyBasicFieldsToDATATYPEDEFINITIONREAL_WOP(&datatypedefinitionrealAPI.DATATYPEDEFINITIONREAL_WOP)

	c.JSON(http.StatusOK, datatypedefinitionrealAPI)
}

// UpdateDATATYPEDEFINITIONREAL
//
// swagger:route PATCH /datatypedefinitionreals/{ID} datatypedefinitionreals updateDATATYPEDEFINITIONREAL
//
// # Update a datatypedefinitionreal
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionrealDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONREAL(c *gin.Context) {

	mutexDATATYPEDEFINITIONREAL.Lock()
	defer mutexDATATYPEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONREALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionrealDB orm.DATATYPEDEFINITIONREALDB

	// fetch the datatypedefinitionreal
	query := db.First(&datatypedefinitionrealDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionrealDB.CopyBasicFieldsFromDATATYPEDEFINITIONREAL_WOP(&input.DATATYPEDEFINITIONREAL_WOP)
	datatypedefinitionrealDB.DATATYPEDEFINITIONREALPointersEncoding = input.DATATYPEDEFINITIONREALPointersEncoding

	query = db.Model(&datatypedefinitionrealDB).Updates(datatypedefinitionrealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionrealNew := new(models.DATATYPEDEFINITIONREAL)
	datatypedefinitionrealDB.CopyBasicFieldsToDATATYPEDEFINITIONREAL(datatypedefinitionrealNew)

	// redeem pointers
	datatypedefinitionrealDB.DecodePointers(backRepo, datatypedefinitionrealNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionrealOld := backRepo.BackRepoDATATYPEDEFINITIONREAL.Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALPtr[datatypedefinitionrealDB.ID]
	if datatypedefinitionrealOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionrealOld, datatypedefinitionrealNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionrealDB
	c.JSON(http.StatusOK, datatypedefinitionrealDB)
}

// DeleteDATATYPEDEFINITIONREAL
//
// swagger:route DELETE /datatypedefinitionreals/{ID} datatypedefinitionreals deleteDATATYPEDEFINITIONREAL
//
// # Delete a datatypedefinitionreal
//
// default: genericError
//
//	200: datatypedefinitionrealDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONREAL(c *gin.Context) {

	mutexDATATYPEDEFINITIONREAL.Lock()
	defer mutexDATATYPEDEFINITIONREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONREAL.GetDB()

	// Get model if exist
	var datatypedefinitionrealDB orm.DATATYPEDEFINITIONREALDB
	if err := db.First(&datatypedefinitionrealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionrealDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionrealDeleted := new(models.DATATYPEDEFINITIONREAL)
	datatypedefinitionrealDB.CopyBasicFieldsToDATATYPEDEFINITIONREAL(datatypedefinitionrealDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionrealStaged := backRepo.BackRepoDATATYPEDEFINITIONREAL.Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALPtr[datatypedefinitionrealDB.ID]
	if datatypedefinitionrealStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionrealStaged, datatypedefinitionrealDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
