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
var __TARGET__dummysDeclaration__ models.TARGET
var __TARGET_time__dummyDeclaration time.Duration

var mutexTARGET sync.Mutex

// An TARGETID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTARGET updateTARGET deleteTARGET
type TARGETID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TARGETInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTARGET updateTARGET
type TARGETInput struct {
	// The TARGET to submit or modify
	// in: body
	TARGET *orm.TARGETAPI
}

// GetTARGETs
//
// swagger:route GET /targets targets getTARGETs
//
// # Get all targets
//
// Responses:
// default: genericError
//
//	200: targetDBResponse
func (controller *Controller) GetTARGETs(c *gin.Context) {

	// source slice
	var targetDBs []orm.TARGETDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTARGETs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGET.GetDB()

	query := db.Find(&targetDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	targetAPIs := make([]orm.TARGETAPI, 0)

	// for each target, update fields from the database nullable fields
	for idx := range targetDBs {
		targetDB := &targetDBs[idx]
		_ = targetDB
		var targetAPI orm.TARGETAPI

		// insertion point for updating fields
		targetAPI.ID = targetDB.ID
		targetDB.CopyBasicFieldsToTARGET_WOP(&targetAPI.TARGET_WOP)
		targetAPI.TARGETPointersEncoding = targetDB.TARGETPointersEncoding
		targetAPIs = append(targetAPIs, targetAPI)
	}

	c.JSON(http.StatusOK, targetAPIs)
}

// PostTARGET
//
// swagger:route POST /targets targets postTARGET
//
// Creates a target
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTARGET(c *gin.Context) {

	mutexTARGET.Lock()
	defer mutexTARGET.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTARGETs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGET.GetDB()

	// Validate input
	var input orm.TARGETAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create target
	targetDB := orm.TARGETDB{}
	targetDB.TARGETPointersEncoding = input.TARGETPointersEncoding
	targetDB.CopyBasicFieldsFromTARGET_WOP(&input.TARGET_WOP)

	query := db.Create(&targetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTARGET.CheckoutPhaseOneInstance(&targetDB)
	target := backRepo.BackRepoTARGET.Map_TARGETDBID_TARGETPtr[targetDB.ID]

	if target != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), target)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, targetDB)
}

// GetTARGET
//
// swagger:route GET /targets/{ID} targets getTARGET
//
// Gets the details for a target.
//
// Responses:
// default: genericError
//
//	200: targetDBResponse
func (controller *Controller) GetTARGET(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTARGET", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGET.GetDB()

	// Get targetDB in DB
	var targetDB orm.TARGETDB
	if err := db.First(&targetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var targetAPI orm.TARGETAPI
	targetAPI.ID = targetDB.ID
	targetAPI.TARGETPointersEncoding = targetDB.TARGETPointersEncoding
	targetDB.CopyBasicFieldsToTARGET_WOP(&targetAPI.TARGET_WOP)

	c.JSON(http.StatusOK, targetAPI)
}

// UpdateTARGET
//
// swagger:route PATCH /targets/{ID} targets updateTARGET
//
// # Update a target
//
// Responses:
// default: genericError
//
//	200: targetDBResponse
func (controller *Controller) UpdateTARGET(c *gin.Context) {

	mutexTARGET.Lock()
	defer mutexTARGET.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTARGET", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGET.GetDB()

	// Validate input
	var input orm.TARGETAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var targetDB orm.TARGETDB

	// fetch the target
	query := db.First(&targetDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	targetDB.CopyBasicFieldsFromTARGET_WOP(&input.TARGET_WOP)
	targetDB.TARGETPointersEncoding = input.TARGETPointersEncoding

	query = db.Model(&targetDB).Updates(targetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	targetNew := new(models.TARGET)
	targetDB.CopyBasicFieldsToTARGET(targetNew)

	// redeem pointers
	targetDB.DecodePointers(backRepo, targetNew)

	// get stage instance from DB instance, and call callback function
	targetOld := backRepo.BackRepoTARGET.Map_TARGETDBID_TARGETPtr[targetDB.ID]
	if targetOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), targetOld, targetNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the targetDB
	c.JSON(http.StatusOK, targetDB)
}

// DeleteTARGET
//
// swagger:route DELETE /targets/{ID} targets deleteTARGET
//
// # Delete a target
//
// default: genericError
//
//	200: targetDBResponse
func (controller *Controller) DeleteTARGET(c *gin.Context) {

	mutexTARGET.Lock()
	defer mutexTARGET.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTARGET", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTARGET.GetDB()

	// Get model if exist
	var targetDB orm.TARGETDB
	if err := db.First(&targetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&targetDB)

	// get an instance (not staged) from DB instance, and call callback function
	targetDeleted := new(models.TARGET)
	targetDB.CopyBasicFieldsToTARGET(targetDeleted)

	// get stage instance from DB instance, and call callback function
	targetStaged := backRepo.BackRepoTARGET.Map_TARGETDBID_TARGETPtr[targetDB.ID]
	if targetStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), targetStaged, targetDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
