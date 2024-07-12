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
var __SOURCE__dummysDeclaration__ models.SOURCE
var __SOURCE_time__dummyDeclaration time.Duration

var mutexSOURCE sync.Mutex

// An SOURCEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSOURCE updateSOURCE deleteSOURCE
type SOURCEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SOURCEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSOURCE updateSOURCE
type SOURCEInput struct {
	// The SOURCE to submit or modify
	// in: body
	SOURCE *orm.SOURCEAPI
}

// GetSOURCEs
//
// swagger:route GET /sources sources getSOURCEs
//
// # Get all sources
//
// Responses:
// default: genericError
//
//	200: sourceDBResponse
func (controller *Controller) GetSOURCEs(c *gin.Context) {

	// source slice
	var sourceDBs []orm.SOURCEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSOURCEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCE.GetDB()

	query := db.Find(&sourceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sourceAPIs := make([]orm.SOURCEAPI, 0)

	// for each source, update fields from the database nullable fields
	for idx := range sourceDBs {
		sourceDB := &sourceDBs[idx]
		_ = sourceDB
		var sourceAPI orm.SOURCEAPI

		// insertion point for updating fields
		sourceAPI.ID = sourceDB.ID
		sourceDB.CopyBasicFieldsToSOURCE_WOP(&sourceAPI.SOURCE_WOP)
		sourceAPI.SOURCEPointersEncoding = sourceDB.SOURCEPointersEncoding
		sourceAPIs = append(sourceAPIs, sourceAPI)
	}

	c.JSON(http.StatusOK, sourceAPIs)
}

// PostSOURCE
//
// swagger:route POST /sources sources postSOURCE
//
// Creates a source
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSOURCE(c *gin.Context) {

	mutexSOURCE.Lock()
	defer mutexSOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSOURCEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCE.GetDB()

	// Validate input
	var input orm.SOURCEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create source
	sourceDB := orm.SOURCEDB{}
	sourceDB.SOURCEPointersEncoding = input.SOURCEPointersEncoding
	sourceDB.CopyBasicFieldsFromSOURCE_WOP(&input.SOURCE_WOP)

	query := db.Create(&sourceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSOURCE.CheckoutPhaseOneInstance(&sourceDB)
	source := backRepo.BackRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]

	if source != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), source)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sourceDB)
}

// GetSOURCE
//
// swagger:route GET /sources/{ID} sources getSOURCE
//
// Gets the details for a source.
//
// Responses:
// default: genericError
//
//	200: sourceDBResponse
func (controller *Controller) GetSOURCE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCE.GetDB()

	// Get sourceDB in DB
	var sourceDB orm.SOURCEDB
	if err := db.First(&sourceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sourceAPI orm.SOURCEAPI
	sourceAPI.ID = sourceDB.ID
	sourceAPI.SOURCEPointersEncoding = sourceDB.SOURCEPointersEncoding
	sourceDB.CopyBasicFieldsToSOURCE_WOP(&sourceAPI.SOURCE_WOP)

	c.JSON(http.StatusOK, sourceAPI)
}

// UpdateSOURCE
//
// swagger:route PATCH /sources/{ID} sources updateSOURCE
//
// # Update a source
//
// Responses:
// default: genericError
//
//	200: sourceDBResponse
func (controller *Controller) UpdateSOURCE(c *gin.Context) {

	mutexSOURCE.Lock()
	defer mutexSOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCE.GetDB()

	// Validate input
	var input orm.SOURCEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sourceDB orm.SOURCEDB

	// fetch the source
	query := db.First(&sourceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sourceDB.CopyBasicFieldsFromSOURCE_WOP(&input.SOURCE_WOP)
	sourceDB.SOURCEPointersEncoding = input.SOURCEPointersEncoding

	query = db.Model(&sourceDB).Updates(sourceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sourceNew := new(models.SOURCE)
	sourceDB.CopyBasicFieldsToSOURCE(sourceNew)

	// redeem pointers
	sourceDB.DecodePointers(backRepo, sourceNew)

	// get stage instance from DB instance, and call callback function
	sourceOld := backRepo.BackRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]
	if sourceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sourceOld, sourceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sourceDB
	c.JSON(http.StatusOK, sourceDB)
}

// DeleteSOURCE
//
// swagger:route DELETE /sources/{ID} sources deleteSOURCE
//
// # Delete a source
//
// default: genericError
//
//	200: sourceDBResponse
func (controller *Controller) DeleteSOURCE(c *gin.Context) {

	mutexSOURCE.Lock()
	defer mutexSOURCE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSOURCE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSOURCE.GetDB()

	// Get model if exist
	var sourceDB orm.SOURCEDB
	if err := db.First(&sourceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&sourceDB)

	// get an instance (not staged) from DB instance, and call callback function
	sourceDeleted := new(models.SOURCE)
	sourceDB.CopyBasicFieldsToSOURCE(sourceDeleted)

	// get stage instance from DB instance, and call callback function
	sourceStaged := backRepo.BackRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]
	if sourceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sourceStaged, sourceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
