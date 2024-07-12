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
var __DATATYPEDEFINITIONINTEGER__dummysDeclaration__ models.DATATYPEDEFINITIONINTEGER
var __DATATYPEDEFINITIONINTEGER_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONINTEGER sync.Mutex

// An DATATYPEDEFINITIONINTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONINTEGER updateDATATYPEDEFINITIONINTEGER deleteDATATYPEDEFINITIONINTEGER
type DATATYPEDEFINITIONINTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONINTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONINTEGER updateDATATYPEDEFINITIONINTEGER
type DATATYPEDEFINITIONINTEGERInput struct {
	// The DATATYPEDEFINITIONINTEGER to submit or modify
	// in: body
	DATATYPEDEFINITIONINTEGER *orm.DATATYPEDEFINITIONINTEGERAPI
}

// GetDATATYPEDEFINITIONINTEGERs
//
// swagger:route GET /datatypedefinitionintegers datatypedefinitionintegers getDATATYPEDEFINITIONINTEGERs
//
// # Get all datatypedefinitionintegers
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionintegerDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONINTEGERs(c *gin.Context) {

	// source slice
	var datatypedefinitionintegerDBs []orm.DATATYPEDEFINITIONINTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDB()

	query := db.Find(&datatypedefinitionintegerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionintegerAPIs := make([]orm.DATATYPEDEFINITIONINTEGERAPI, 0)

	// for each datatypedefinitioninteger, update fields from the database nullable fields
	for idx := range datatypedefinitionintegerDBs {
		datatypedefinitionintegerDB := &datatypedefinitionintegerDBs[idx]
		_ = datatypedefinitionintegerDB
		var datatypedefinitionintegerAPI orm.DATATYPEDEFINITIONINTEGERAPI

		// insertion point for updating fields
		datatypedefinitionintegerAPI.ID = datatypedefinitionintegerDB.ID
		datatypedefinitionintegerDB.CopyBasicFieldsToDATATYPEDEFINITIONINTEGER_WOP(&datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGER_WOP)
		datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGERPointersEncoding = datatypedefinitionintegerDB.DATATYPEDEFINITIONINTEGERPointersEncoding
		datatypedefinitionintegerAPIs = append(datatypedefinitionintegerAPIs, datatypedefinitionintegerAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionintegerAPIs)
}

// PostDATATYPEDEFINITIONINTEGER
//
// swagger:route POST /datatypedefinitionintegers datatypedefinitionintegers postDATATYPEDEFINITIONINTEGER
//
// Creates a datatypedefinitioninteger
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONINTEGER(c *gin.Context) {

	mutexDATATYPEDEFINITIONINTEGER.Lock()
	defer mutexDATATYPEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONINTEGERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONINTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitioninteger
	datatypedefinitionintegerDB := orm.DATATYPEDEFINITIONINTEGERDB{}
	datatypedefinitionintegerDB.DATATYPEDEFINITIONINTEGERPointersEncoding = input.DATATYPEDEFINITIONINTEGERPointersEncoding
	datatypedefinitionintegerDB.CopyBasicFieldsFromDATATYPEDEFINITIONINTEGER_WOP(&input.DATATYPEDEFINITIONINTEGER_WOP)

	query := db.Create(&datatypedefinitionintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.CheckoutPhaseOneInstance(&datatypedefinitionintegerDB)
	datatypedefinitioninteger := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERPtr[datatypedefinitionintegerDB.ID]

	if datatypedefinitioninteger != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitioninteger)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionintegerDB)
}

// GetDATATYPEDEFINITIONINTEGER
//
// swagger:route GET /datatypedefinitionintegers/{ID} datatypedefinitionintegers getDATATYPEDEFINITIONINTEGER
//
// Gets the details for a datatypedefinitioninteger.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionintegerDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONINTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDB()

	// Get datatypedefinitionintegerDB in DB
	var datatypedefinitionintegerDB orm.DATATYPEDEFINITIONINTEGERDB
	if err := db.First(&datatypedefinitionintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionintegerAPI orm.DATATYPEDEFINITIONINTEGERAPI
	datatypedefinitionintegerAPI.ID = datatypedefinitionintegerDB.ID
	datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGERPointersEncoding = datatypedefinitionintegerDB.DATATYPEDEFINITIONINTEGERPointersEncoding
	datatypedefinitionintegerDB.CopyBasicFieldsToDATATYPEDEFINITIONINTEGER_WOP(&datatypedefinitionintegerAPI.DATATYPEDEFINITIONINTEGER_WOP)

	c.JSON(http.StatusOK, datatypedefinitionintegerAPI)
}

// UpdateDATATYPEDEFINITIONINTEGER
//
// swagger:route PATCH /datatypedefinitionintegers/{ID} datatypedefinitionintegers updateDATATYPEDEFINITIONINTEGER
//
// # Update a datatypedefinitioninteger
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionintegerDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONINTEGER(c *gin.Context) {

	mutexDATATYPEDEFINITIONINTEGER.Lock()
	defer mutexDATATYPEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONINTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionintegerDB orm.DATATYPEDEFINITIONINTEGERDB

	// fetch the datatypedefinitioninteger
	query := db.First(&datatypedefinitionintegerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionintegerDB.CopyBasicFieldsFromDATATYPEDEFINITIONINTEGER_WOP(&input.DATATYPEDEFINITIONINTEGER_WOP)
	datatypedefinitionintegerDB.DATATYPEDEFINITIONINTEGERPointersEncoding = input.DATATYPEDEFINITIONINTEGERPointersEncoding

	query = db.Model(&datatypedefinitionintegerDB).Updates(datatypedefinitionintegerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionintegerNew := new(models.DATATYPEDEFINITIONINTEGER)
	datatypedefinitionintegerDB.CopyBasicFieldsToDATATYPEDEFINITIONINTEGER(datatypedefinitionintegerNew)

	// redeem pointers
	datatypedefinitionintegerDB.DecodePointers(backRepo, datatypedefinitionintegerNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionintegerOld := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERPtr[datatypedefinitionintegerDB.ID]
	if datatypedefinitionintegerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionintegerOld, datatypedefinitionintegerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionintegerDB
	c.JSON(http.StatusOK, datatypedefinitionintegerDB)
}

// DeleteDATATYPEDEFINITIONINTEGER
//
// swagger:route DELETE /datatypedefinitionintegers/{ID} datatypedefinitionintegers deleteDATATYPEDEFINITIONINTEGER
//
// # Delete a datatypedefinitioninteger
//
// default: genericError
//
//	200: datatypedefinitionintegerDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONINTEGER(c *gin.Context) {

	mutexDATATYPEDEFINITIONINTEGER.Lock()
	defer mutexDATATYPEDEFINITIONINTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONINTEGER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.GetDB()

	// Get model if exist
	var datatypedefinitionintegerDB orm.DATATYPEDEFINITIONINTEGERDB
	if err := db.First(&datatypedefinitionintegerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionintegerDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionintegerDeleted := new(models.DATATYPEDEFINITIONINTEGER)
	datatypedefinitionintegerDB.CopyBasicFieldsToDATATYPEDEFINITIONINTEGER(datatypedefinitionintegerDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionintegerStaged := backRepo.BackRepoDATATYPEDEFINITIONINTEGER.Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERPtr[datatypedefinitionintegerDB.ID]
	if datatypedefinitionintegerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionintegerStaged, datatypedefinitionintegerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
