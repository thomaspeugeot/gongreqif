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
var __SOURCESPECIFICATION__dummysDeclaration__ models.SOURCESPECIFICATION
var __SOURCESPECIFICATION_time__dummyDeclaration time.Duration

var mutexSOURCESPECIFICATION sync.Mutex

// An SOURCESPECIFICATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSOURCESPECIFICATION updateSOURCESPECIFICATION deleteSOURCESPECIFICATION
type SOURCESPECIFICATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SOURCESPECIFICATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSOURCESPECIFICATION updateSOURCESPECIFICATION
type SOURCESPECIFICATIONInput struct {
	// The SOURCESPECIFICATION to submit or modify
	// in: body
	SOURCESPECIFICATION *orm.SOURCESPECIFICATIONAPI
}

// GetSOURCESPECIFICATIONs
//
// swagger:route GET /sourcespecifications sourcespecifications getSOURCESPECIFICATIONs
//
// # Get all sourcespecifications
//
// Responses:
// default: genericError
//
//	200: sourcespecificationDBResponse
func (controller *Controller) GetSOURCESPECIFICATIONs(c *gin.Context) {

	// source slice
	var sourcespecificationDBs []orm.SOURCESPECIFICATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSOURCESPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCESPECIFICATION.GetDB()

	query := db.Find(&sourcespecificationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sourcespecificationAPIs := make([]orm.SOURCESPECIFICATIONAPI, 0)

	// for each sourcespecification, update fields from the database nullable fields
	for idx := range sourcespecificationDBs {
		sourcespecificationDB := &sourcespecificationDBs[idx]
		_ = sourcespecificationDB
		var sourcespecificationAPI orm.SOURCESPECIFICATIONAPI

		// insertion point for updating fields
		sourcespecificationAPI.ID = sourcespecificationDB.ID
		sourcespecificationDB.CopyBasicFieldsToSOURCESPECIFICATION_WOP(&sourcespecificationAPI.SOURCESPECIFICATION_WOP)
		sourcespecificationAPI.SOURCESPECIFICATIONPointersEncoding = sourcespecificationDB.SOURCESPECIFICATIONPointersEncoding
		sourcespecificationAPIs = append(sourcespecificationAPIs, sourcespecificationAPI)
	}

	c.JSON(http.StatusOK, sourcespecificationAPIs)
}

// PostSOURCESPECIFICATION
//
// swagger:route POST /sourcespecifications sourcespecifications postSOURCESPECIFICATION
//
// Creates a sourcespecification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSOURCESPECIFICATION(c *gin.Context) {

	mutexSOURCESPECIFICATION.Lock()
	defer mutexSOURCESPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSOURCESPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCESPECIFICATION.GetDB()

	// Validate input
	var input orm.SOURCESPECIFICATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sourcespecification
	sourcespecificationDB := orm.SOURCESPECIFICATIONDB{}
	sourcespecificationDB.SOURCESPECIFICATIONPointersEncoding = input.SOURCESPECIFICATIONPointersEncoding
	sourcespecificationDB.CopyBasicFieldsFromSOURCESPECIFICATION_WOP(&input.SOURCESPECIFICATION_WOP)

	query := db.Create(&sourcespecificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSOURCESPECIFICATION.CheckoutPhaseOneInstance(&sourcespecificationDB)
	sourcespecification := backRepo.BackRepoSOURCESPECIFICATION.Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONPtr[sourcespecificationDB.ID]

	if sourcespecification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sourcespecification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sourcespecificationDB)
}

// GetSOURCESPECIFICATION
//
// swagger:route GET /sourcespecifications/{ID} sourcespecifications getSOURCESPECIFICATION
//
// Gets the details for a sourcespecification.
//
// Responses:
// default: genericError
//
//	200: sourcespecificationDBResponse
func (controller *Controller) GetSOURCESPECIFICATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSOURCESPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCESPECIFICATION.GetDB()

	// Get sourcespecificationDB in DB
	var sourcespecificationDB orm.SOURCESPECIFICATIONDB
	if err := db.First(&sourcespecificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sourcespecificationAPI orm.SOURCESPECIFICATIONAPI
	sourcespecificationAPI.ID = sourcespecificationDB.ID
	sourcespecificationAPI.SOURCESPECIFICATIONPointersEncoding = sourcespecificationDB.SOURCESPECIFICATIONPointersEncoding
	sourcespecificationDB.CopyBasicFieldsToSOURCESPECIFICATION_WOP(&sourcespecificationAPI.SOURCESPECIFICATION_WOP)

	c.JSON(http.StatusOK, sourcespecificationAPI)
}

// UpdateSOURCESPECIFICATION
//
// swagger:route PATCH /sourcespecifications/{ID} sourcespecifications updateSOURCESPECIFICATION
//
// # Update a sourcespecification
//
// Responses:
// default: genericError
//
//	200: sourcespecificationDBResponse
func (controller *Controller) UpdateSOURCESPECIFICATION(c *gin.Context) {

	mutexSOURCESPECIFICATION.Lock()
	defer mutexSOURCESPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSOURCESPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCESPECIFICATION.GetDB()

	// Validate input
	var input orm.SOURCESPECIFICATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sourcespecificationDB orm.SOURCESPECIFICATIONDB

	// fetch the sourcespecification
	query := db.First(&sourcespecificationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sourcespecificationDB.CopyBasicFieldsFromSOURCESPECIFICATION_WOP(&input.SOURCESPECIFICATION_WOP)
	sourcespecificationDB.SOURCESPECIFICATIONPointersEncoding = input.SOURCESPECIFICATIONPointersEncoding

	query = db.Model(&sourcespecificationDB).Updates(sourcespecificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sourcespecificationNew := new(models.SOURCESPECIFICATION)
	sourcespecificationDB.CopyBasicFieldsToSOURCESPECIFICATION(sourcespecificationNew)

	// redeem pointers
	sourcespecificationDB.DecodePointers(backRepo, sourcespecificationNew)

	// get stage instance from DB instance, and call callback function
	sourcespecificationOld := backRepo.BackRepoSOURCESPECIFICATION.Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONPtr[sourcespecificationDB.ID]
	if sourcespecificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sourcespecificationOld, sourcespecificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sourcespecificationDB
	c.JSON(http.StatusOK, sourcespecificationDB)
}

// DeleteSOURCESPECIFICATION
//
// swagger:route DELETE /sourcespecifications/{ID} sourcespecifications deleteSOURCESPECIFICATION
//
// # Delete a sourcespecification
//
// default: genericError
//
//	200: sourcespecificationDBResponse
func (controller *Controller) DeleteSOURCESPECIFICATION(c *gin.Context) {

	mutexSOURCESPECIFICATION.Lock()
	defer mutexSOURCESPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSOURCESPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCESPECIFICATION.GetDB()

	// Get model if exist
	var sourcespecificationDB orm.SOURCESPECIFICATIONDB
	if err := db.First(&sourcespecificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&sourcespecificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	sourcespecificationDeleted := new(models.SOURCESPECIFICATION)
	sourcespecificationDB.CopyBasicFieldsToSOURCESPECIFICATION(sourcespecificationDeleted)

	// get stage instance from DB instance, and call callback function
	sourcespecificationStaged := backRepo.BackRepoSOURCESPECIFICATION.Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONPtr[sourcespecificationDB.ID]
	if sourcespecificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sourcespecificationStaged, sourcespecificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
