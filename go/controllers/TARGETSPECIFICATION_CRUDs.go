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
var __TARGETSPECIFICATION__dummysDeclaration__ models.TARGETSPECIFICATION
var __TARGETSPECIFICATION_time__dummyDeclaration time.Duration

var mutexTARGETSPECIFICATION sync.Mutex

// An TARGETSPECIFICATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTARGETSPECIFICATION updateTARGETSPECIFICATION deleteTARGETSPECIFICATION
type TARGETSPECIFICATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TARGETSPECIFICATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTARGETSPECIFICATION updateTARGETSPECIFICATION
type TARGETSPECIFICATIONInput struct {
	// The TARGETSPECIFICATION to submit or modify
	// in: body
	TARGETSPECIFICATION *orm.TARGETSPECIFICATIONAPI
}

// GetTARGETSPECIFICATIONs
//
// swagger:route GET /targetspecifications targetspecifications getTARGETSPECIFICATIONs
//
// # Get all targetspecifications
//
// Responses:
// default: genericError
//
//	200: targetspecificationDBResponse
func (controller *Controller) GetTARGETSPECIFICATIONs(c *gin.Context) {

	// source slice
	var targetspecificationDBs []orm.TARGETSPECIFICATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTARGETSPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGETSPECIFICATION.GetDB()

	query := db.Find(&targetspecificationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	targetspecificationAPIs := make([]orm.TARGETSPECIFICATIONAPI, 0)

	// for each targetspecification, update fields from the database nullable fields
	for idx := range targetspecificationDBs {
		targetspecificationDB := &targetspecificationDBs[idx]
		_ = targetspecificationDB
		var targetspecificationAPI orm.TARGETSPECIFICATIONAPI

		// insertion point for updating fields
		targetspecificationAPI.ID = targetspecificationDB.ID
		targetspecificationDB.CopyBasicFieldsToTARGETSPECIFICATION_WOP(&targetspecificationAPI.TARGETSPECIFICATION_WOP)
		targetspecificationAPI.TARGETSPECIFICATIONPointersEncoding = targetspecificationDB.TARGETSPECIFICATIONPointersEncoding
		targetspecificationAPIs = append(targetspecificationAPIs, targetspecificationAPI)
	}

	c.JSON(http.StatusOK, targetspecificationAPIs)
}

// PostTARGETSPECIFICATION
//
// swagger:route POST /targetspecifications targetspecifications postTARGETSPECIFICATION
//
// Creates a targetspecification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTARGETSPECIFICATION(c *gin.Context) {

	mutexTARGETSPECIFICATION.Lock()
	defer mutexTARGETSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTARGETSPECIFICATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGETSPECIFICATION.GetDB()

	// Validate input
	var input orm.TARGETSPECIFICATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create targetspecification
	targetspecificationDB := orm.TARGETSPECIFICATIONDB{}
	targetspecificationDB.TARGETSPECIFICATIONPointersEncoding = input.TARGETSPECIFICATIONPointersEncoding
	targetspecificationDB.CopyBasicFieldsFromTARGETSPECIFICATION_WOP(&input.TARGETSPECIFICATION_WOP)

	query := db.Create(&targetspecificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTARGETSPECIFICATION.CheckoutPhaseOneInstance(&targetspecificationDB)
	targetspecification := backRepo.BackRepoTARGETSPECIFICATION.Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONPtr[targetspecificationDB.ID]

	if targetspecification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), targetspecification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, targetspecificationDB)
}

// GetTARGETSPECIFICATION
//
// swagger:route GET /targetspecifications/{ID} targetspecifications getTARGETSPECIFICATION
//
// Gets the details for a targetspecification.
//
// Responses:
// default: genericError
//
//	200: targetspecificationDBResponse
func (controller *Controller) GetTARGETSPECIFICATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTARGETSPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGETSPECIFICATION.GetDB()

	// Get targetspecificationDB in DB
	var targetspecificationDB orm.TARGETSPECIFICATIONDB
	if err := db.First(&targetspecificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var targetspecificationAPI orm.TARGETSPECIFICATIONAPI
	targetspecificationAPI.ID = targetspecificationDB.ID
	targetspecificationAPI.TARGETSPECIFICATIONPointersEncoding = targetspecificationDB.TARGETSPECIFICATIONPointersEncoding
	targetspecificationDB.CopyBasicFieldsToTARGETSPECIFICATION_WOP(&targetspecificationAPI.TARGETSPECIFICATION_WOP)

	c.JSON(http.StatusOK, targetspecificationAPI)
}

// UpdateTARGETSPECIFICATION
//
// swagger:route PATCH /targetspecifications/{ID} targetspecifications updateTARGETSPECIFICATION
//
// # Update a targetspecification
//
// Responses:
// default: genericError
//
//	200: targetspecificationDBResponse
func (controller *Controller) UpdateTARGETSPECIFICATION(c *gin.Context) {

	mutexTARGETSPECIFICATION.Lock()
	defer mutexTARGETSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTARGETSPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGETSPECIFICATION.GetDB()

	// Validate input
	var input orm.TARGETSPECIFICATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var targetspecificationDB orm.TARGETSPECIFICATIONDB

	// fetch the targetspecification
	query := db.First(&targetspecificationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	targetspecificationDB.CopyBasicFieldsFromTARGETSPECIFICATION_WOP(&input.TARGETSPECIFICATION_WOP)
	targetspecificationDB.TARGETSPECIFICATIONPointersEncoding = input.TARGETSPECIFICATIONPointersEncoding

	query = db.Model(&targetspecificationDB).Updates(targetspecificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	targetspecificationNew := new(models.TARGETSPECIFICATION)
	targetspecificationDB.CopyBasicFieldsToTARGETSPECIFICATION(targetspecificationNew)

	// redeem pointers
	targetspecificationDB.DecodePointers(backRepo, targetspecificationNew)

	// get stage instance from DB instance, and call callback function
	targetspecificationOld := backRepo.BackRepoTARGETSPECIFICATION.Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONPtr[targetspecificationDB.ID]
	if targetspecificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), targetspecificationOld, targetspecificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the targetspecificationDB
	c.JSON(http.StatusOK, targetspecificationDB)
}

// DeleteTARGETSPECIFICATION
//
// swagger:route DELETE /targetspecifications/{ID} targetspecifications deleteTARGETSPECIFICATION
//
// # Delete a targetspecification
//
// default: genericError
//
//	200: targetspecificationDBResponse
func (controller *Controller) DeleteTARGETSPECIFICATION(c *gin.Context) {

	mutexTARGETSPECIFICATION.Lock()
	defer mutexTARGETSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTARGETSPECIFICATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGETSPECIFICATION.GetDB()

	// Get model if exist
	var targetspecificationDB orm.TARGETSPECIFICATIONDB
	if err := db.First(&targetspecificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&targetspecificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	targetspecificationDeleted := new(models.TARGETSPECIFICATION)
	targetspecificationDB.CopyBasicFieldsToTARGETSPECIFICATION(targetspecificationDeleted)

	// get stage instance from DB instance, and call callback function
	targetspecificationStaged := backRepo.BackRepoTARGETSPECIFICATION.Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONPtr[targetspecificationDB.ID]
	if targetspecificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), targetspecificationStaged, targetspecificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
