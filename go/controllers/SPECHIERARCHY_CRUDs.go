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
var __SPECHIERARCHY__dummysDeclaration__ models.SPECHIERARCHY
var __SPECHIERARCHY_time__dummyDeclaration time.Duration

var mutexSPECHIERARCHY sync.Mutex

// An SPECHIERARCHYID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECHIERARCHY updateSPECHIERARCHY deleteSPECHIERARCHY
type SPECHIERARCHYID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECHIERARCHYInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECHIERARCHY updateSPECHIERARCHY
type SPECHIERARCHYInput struct {
	// The SPECHIERARCHY to submit or modify
	// in: body
	SPECHIERARCHY *orm.SPECHIERARCHYAPI
}

// GetSPECHIERARCHYs
//
// swagger:route GET /spechierarchys spechierarchys getSPECHIERARCHYs
//
// # Get all spechierarchys
//
// Responses:
// default: genericError
//
//	200: spechierarchyDBResponse
func (controller *Controller) GetSPECHIERARCHYs(c *gin.Context) {

	// source slice
	var spechierarchyDBs []orm.SPECHIERARCHYDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECHIERARCHYs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECHIERARCHY.GetDB()

	query := db.Find(&spechierarchyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spechierarchyAPIs := make([]orm.SPECHIERARCHYAPI, 0)

	// for each spechierarchy, update fields from the database nullable fields
	for idx := range spechierarchyDBs {
		spechierarchyDB := &spechierarchyDBs[idx]
		_ = spechierarchyDB
		var spechierarchyAPI orm.SPECHIERARCHYAPI

		// insertion point for updating fields
		spechierarchyAPI.ID = spechierarchyDB.ID
		spechierarchyDB.CopyBasicFieldsToSPECHIERARCHY_WOP(&spechierarchyAPI.SPECHIERARCHY_WOP)
		spechierarchyAPI.SPECHIERARCHYPointersEncoding = spechierarchyDB.SPECHIERARCHYPointersEncoding
		spechierarchyAPIs = append(spechierarchyAPIs, spechierarchyAPI)
	}

	c.JSON(http.StatusOK, spechierarchyAPIs)
}

// PostSPECHIERARCHY
//
// swagger:route POST /spechierarchys spechierarchys postSPECHIERARCHY
//
// Creates a spechierarchy
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECHIERARCHY(c *gin.Context) {

	mutexSPECHIERARCHY.Lock()
	defer mutexSPECHIERARCHY.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECHIERARCHYs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECHIERARCHY.GetDB()

	// Validate input
	var input orm.SPECHIERARCHYAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spechierarchy
	spechierarchyDB := orm.SPECHIERARCHYDB{}
	spechierarchyDB.SPECHIERARCHYPointersEncoding = input.SPECHIERARCHYPointersEncoding
	spechierarchyDB.CopyBasicFieldsFromSPECHIERARCHY_WOP(&input.SPECHIERARCHY_WOP)

	query := db.Create(&spechierarchyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECHIERARCHY.CheckoutPhaseOneInstance(&spechierarchyDB)
	spechierarchy := backRepo.BackRepoSPECHIERARCHY.Map_SPECHIERARCHYDBID_SPECHIERARCHYPtr[spechierarchyDB.ID]

	if spechierarchy != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spechierarchy)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spechierarchyDB)
}

// GetSPECHIERARCHY
//
// swagger:route GET /spechierarchys/{ID} spechierarchys getSPECHIERARCHY
//
// Gets the details for a spechierarchy.
//
// Responses:
// default: genericError
//
//	200: spechierarchyDBResponse
func (controller *Controller) GetSPECHIERARCHY(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECHIERARCHY", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECHIERARCHY.GetDB()

	// Get spechierarchyDB in DB
	var spechierarchyDB orm.SPECHIERARCHYDB
	if err := db.First(&spechierarchyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spechierarchyAPI orm.SPECHIERARCHYAPI
	spechierarchyAPI.ID = spechierarchyDB.ID
	spechierarchyAPI.SPECHIERARCHYPointersEncoding = spechierarchyDB.SPECHIERARCHYPointersEncoding
	spechierarchyDB.CopyBasicFieldsToSPECHIERARCHY_WOP(&spechierarchyAPI.SPECHIERARCHY_WOP)

	c.JSON(http.StatusOK, spechierarchyAPI)
}

// UpdateSPECHIERARCHY
//
// swagger:route PATCH /spechierarchys/{ID} spechierarchys updateSPECHIERARCHY
//
// # Update a spechierarchy
//
// Responses:
// default: genericError
//
//	200: spechierarchyDBResponse
func (controller *Controller) UpdateSPECHIERARCHY(c *gin.Context) {

	mutexSPECHIERARCHY.Lock()
	defer mutexSPECHIERARCHY.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECHIERARCHY", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECHIERARCHY.GetDB()

	// Validate input
	var input orm.SPECHIERARCHYAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spechierarchyDB orm.SPECHIERARCHYDB

	// fetch the spechierarchy
	query := db.First(&spechierarchyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spechierarchyDB.CopyBasicFieldsFromSPECHIERARCHY_WOP(&input.SPECHIERARCHY_WOP)
	spechierarchyDB.SPECHIERARCHYPointersEncoding = input.SPECHIERARCHYPointersEncoding

	query = db.Model(&spechierarchyDB).Updates(spechierarchyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spechierarchyNew := new(models.SPECHIERARCHY)
	spechierarchyDB.CopyBasicFieldsToSPECHIERARCHY(spechierarchyNew)

	// redeem pointers
	spechierarchyDB.DecodePointers(backRepo, spechierarchyNew)

	// get stage instance from DB instance, and call callback function
	spechierarchyOld := backRepo.BackRepoSPECHIERARCHY.Map_SPECHIERARCHYDBID_SPECHIERARCHYPtr[spechierarchyDB.ID]
	if spechierarchyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spechierarchyOld, spechierarchyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spechierarchyDB
	c.JSON(http.StatusOK, spechierarchyDB)
}

// DeleteSPECHIERARCHY
//
// swagger:route DELETE /spechierarchys/{ID} spechierarchys deleteSPECHIERARCHY
//
// # Delete a spechierarchy
//
// default: genericError
//
//	200: spechierarchyDBResponse
func (controller *Controller) DeleteSPECHIERARCHY(c *gin.Context) {

	mutexSPECHIERARCHY.Lock()
	defer mutexSPECHIERARCHY.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECHIERARCHY", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECHIERARCHY.GetDB()

	// Get model if exist
	var spechierarchyDB orm.SPECHIERARCHYDB
	if err := db.First(&spechierarchyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spechierarchyDB)

	// get an instance (not staged) from DB instance, and call callback function
	spechierarchyDeleted := new(models.SPECHIERARCHY)
	spechierarchyDB.CopyBasicFieldsToSPECHIERARCHY(spechierarchyDeleted)

	// get stage instance from DB instance, and call callback function
	spechierarchyStaged := backRepo.BackRepoSPECHIERARCHY.Map_SPECHIERARCHYDBID_SPECHIERARCHYPtr[spechierarchyDB.ID]
	if spechierarchyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spechierarchyStaged, spechierarchyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
