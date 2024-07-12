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
var __VALUES__dummysDeclaration__ models.VALUES
var __VALUES_time__dummyDeclaration time.Duration

var mutexVALUES sync.Mutex

// An VALUESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVALUES updateVALUES deleteVALUES
type VALUESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VALUESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVALUES updateVALUES
type VALUESInput struct {
	// The VALUES to submit or modify
	// in: body
	VALUES *orm.VALUESAPI
}

// GetVALUESs
//
// swagger:route GET /valuess valuess getVALUESs
//
// # Get all valuess
//
// Responses:
// default: genericError
//
//	200: valuesDBResponse
func (controller *Controller) GetVALUESs(c *gin.Context) {

	// source slice
	var valuesDBs []orm.VALUESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVALUES.GetDB()

	query := db.Find(&valuesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	valuesAPIs := make([]orm.VALUESAPI, 0)

	// for each values, update fields from the database nullable fields
	for idx := range valuesDBs {
		valuesDB := &valuesDBs[idx]
		_ = valuesDB
		var valuesAPI orm.VALUESAPI

		// insertion point for updating fields
		valuesAPI.ID = valuesDB.ID
		valuesDB.CopyBasicFieldsToVALUES_WOP(&valuesAPI.VALUES_WOP)
		valuesAPI.VALUESPointersEncoding = valuesDB.VALUESPointersEncoding
		valuesAPIs = append(valuesAPIs, valuesAPI)
	}

	c.JSON(http.StatusOK, valuesAPIs)
}

// PostVALUES
//
// swagger:route POST /valuess valuess postVALUES
//
// Creates a values
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVALUES(c *gin.Context) {

	mutexVALUES.Lock()
	defer mutexVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVALUES.GetDB()

	// Validate input
	var input orm.VALUESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create values
	valuesDB := orm.VALUESDB{}
	valuesDB.VALUESPointersEncoding = input.VALUESPointersEncoding
	valuesDB.CopyBasicFieldsFromVALUES_WOP(&input.VALUES_WOP)

	query := db.Create(&valuesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVALUES.CheckoutPhaseOneInstance(&valuesDB)
	values := backRepo.BackRepoVALUES.Map_VALUESDBID_VALUESPtr[valuesDB.ID]

	if values != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), values)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, valuesDB)
}

// GetVALUES
//
// swagger:route GET /valuess/{ID} valuess getVALUES
//
// Gets the details for a values.
//
// Responses:
// default: genericError
//
//	200: valuesDBResponse
func (controller *Controller) GetVALUES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVALUES.GetDB()

	// Get valuesDB in DB
	var valuesDB orm.VALUESDB
	if err := db.First(&valuesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var valuesAPI orm.VALUESAPI
	valuesAPI.ID = valuesDB.ID
	valuesAPI.VALUESPointersEncoding = valuesDB.VALUESPointersEncoding
	valuesDB.CopyBasicFieldsToVALUES_WOP(&valuesAPI.VALUES_WOP)

	c.JSON(http.StatusOK, valuesAPI)
}

// UpdateVALUES
//
// swagger:route PATCH /valuess/{ID} valuess updateVALUES
//
// # Update a values
//
// Responses:
// default: genericError
//
//	200: valuesDBResponse
func (controller *Controller) UpdateVALUES(c *gin.Context) {

	mutexVALUES.Lock()
	defer mutexVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVALUES.GetDB()

	// Validate input
	var input orm.VALUESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var valuesDB orm.VALUESDB

	// fetch the values
	query := db.First(&valuesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	valuesDB.CopyBasicFieldsFromVALUES_WOP(&input.VALUES_WOP)
	valuesDB.VALUESPointersEncoding = input.VALUESPointersEncoding

	query = db.Model(&valuesDB).Updates(valuesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	valuesNew := new(models.VALUES)
	valuesDB.CopyBasicFieldsToVALUES(valuesNew)

	// redeem pointers
	valuesDB.DecodePointers(backRepo, valuesNew)

	// get stage instance from DB instance, and call callback function
	valuesOld := backRepo.BackRepoVALUES.Map_VALUESDBID_VALUESPtr[valuesDB.ID]
	if valuesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), valuesOld, valuesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the valuesDB
	c.JSON(http.StatusOK, valuesDB)
}

// DeleteVALUES
//
// swagger:route DELETE /valuess/{ID} valuess deleteVALUES
//
// # Delete a values
//
// default: genericError
//
//	200: valuesDBResponse
func (controller *Controller) DeleteVALUES(c *gin.Context) {

	mutexVALUES.Lock()
	defer mutexVALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVALUES.GetDB()

	// Get model if exist
	var valuesDB orm.VALUESDB
	if err := db.First(&valuesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&valuesDB)

	// get an instance (not staged) from DB instance, and call callback function
	valuesDeleted := new(models.VALUES)
	valuesDB.CopyBasicFieldsToVALUES(valuesDeleted)

	// get stage instance from DB instance, and call callback function
	valuesStaged := backRepo.BackRepoVALUES.Map_VALUESDBID_VALUESPtr[valuesDB.ID]
	if valuesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), valuesStaged, valuesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
