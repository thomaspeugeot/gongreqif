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
var __SPECIFICATIONS__dummysDeclaration__ models.SPECIFICATIONS
var __SPECIFICATIONS_time__dummyDeclaration time.Duration

var mutexSPECIFICATIONS sync.Mutex

// An SPECIFICATIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECIFICATIONS updateSPECIFICATIONS deleteSPECIFICATIONS
type SPECIFICATIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECIFICATIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECIFICATIONS updateSPECIFICATIONS
type SPECIFICATIONSInput struct {
	// The SPECIFICATIONS to submit or modify
	// in: body
	SPECIFICATIONS *orm.SPECIFICATIONSAPI
}

// GetSPECIFICATIONSs
//
// swagger:route GET /specificationss specificationss getSPECIFICATIONSs
//
// # Get all specificationss
//
// Responses:
// default: genericError
//
//	200: specificationsDBResponse
func (controller *Controller) GetSPECIFICATIONSs(c *gin.Context) {

	// source slice
	var specificationsDBs []orm.SPECIFICATIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONS.GetDB()

	query := db.Find(&specificationsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specificationsAPIs := make([]orm.SPECIFICATIONSAPI, 0)

	// for each specifications, update fields from the database nullable fields
	for idx := range specificationsDBs {
		specificationsDB := &specificationsDBs[idx]
		_ = specificationsDB
		var specificationsAPI orm.SPECIFICATIONSAPI

		// insertion point for updating fields
		specificationsAPI.ID = specificationsDB.ID
		specificationsDB.CopyBasicFieldsToSPECIFICATIONS_WOP(&specificationsAPI.SPECIFICATIONS_WOP)
		specificationsAPI.SPECIFICATIONSPointersEncoding = specificationsDB.SPECIFICATIONSPointersEncoding
		specificationsAPIs = append(specificationsAPIs, specificationsAPI)
	}

	c.JSON(http.StatusOK, specificationsAPIs)
}

// PostSPECIFICATIONS
//
// swagger:route POST /specificationss specificationss postSPECIFICATIONS
//
// Creates a specifications
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECIFICATIONS(c *gin.Context) {

	mutexSPECIFICATIONS.Lock()
	defer mutexSPECIFICATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECIFICATIONSs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONS.GetDB()

	// Validate input
	var input orm.SPECIFICATIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specifications
	specificationsDB := orm.SPECIFICATIONSDB{}
	specificationsDB.SPECIFICATIONSPointersEncoding = input.SPECIFICATIONSPointersEncoding
	specificationsDB.CopyBasicFieldsFromSPECIFICATIONS_WOP(&input.SPECIFICATIONS_WOP)

	query := db.Create(&specificationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECIFICATIONS.CheckoutPhaseOneInstance(&specificationsDB)
	specifications := backRepo.BackRepoSPECIFICATIONS.Map_SPECIFICATIONSDBID_SPECIFICATIONSPtr[specificationsDB.ID]

	if specifications != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specifications)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specificationsDB)
}

// GetSPECIFICATIONS
//
// swagger:route GET /specificationss/{ID} specificationss getSPECIFICATIONS
//
// Gets the details for a specifications.
//
// Responses:
// default: genericError
//
//	200: specificationsDBResponse
func (controller *Controller) GetSPECIFICATIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONS.GetDB()

	// Get specificationsDB in DB
	var specificationsDB orm.SPECIFICATIONSDB
	if err := db.First(&specificationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specificationsAPI orm.SPECIFICATIONSAPI
	specificationsAPI.ID = specificationsDB.ID
	specificationsAPI.SPECIFICATIONSPointersEncoding = specificationsDB.SPECIFICATIONSPointersEncoding
	specificationsDB.CopyBasicFieldsToSPECIFICATIONS_WOP(&specificationsAPI.SPECIFICATIONS_WOP)

	c.JSON(http.StatusOK, specificationsAPI)
}

// UpdateSPECIFICATIONS
//
// swagger:route PATCH /specificationss/{ID} specificationss updateSPECIFICATIONS
//
// # Update a specifications
//
// Responses:
// default: genericError
//
//	200: specificationsDBResponse
func (controller *Controller) UpdateSPECIFICATIONS(c *gin.Context) {

	mutexSPECIFICATIONS.Lock()
	defer mutexSPECIFICATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECIFICATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONS.GetDB()

	// Validate input
	var input orm.SPECIFICATIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specificationsDB orm.SPECIFICATIONSDB

	// fetch the specifications
	query := db.First(&specificationsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specificationsDB.CopyBasicFieldsFromSPECIFICATIONS_WOP(&input.SPECIFICATIONS_WOP)
	specificationsDB.SPECIFICATIONSPointersEncoding = input.SPECIFICATIONSPointersEncoding

	query = db.Model(&specificationsDB).Updates(specificationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specificationsNew := new(models.SPECIFICATIONS)
	specificationsDB.CopyBasicFieldsToSPECIFICATIONS(specificationsNew)

	// redeem pointers
	specificationsDB.DecodePointers(backRepo, specificationsNew)

	// get stage instance from DB instance, and call callback function
	specificationsOld := backRepo.BackRepoSPECIFICATIONS.Map_SPECIFICATIONSDBID_SPECIFICATIONSPtr[specificationsDB.ID]
	if specificationsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specificationsOld, specificationsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specificationsDB
	c.JSON(http.StatusOK, specificationsDB)
}

// DeleteSPECIFICATIONS
//
// swagger:route DELETE /specificationss/{ID} specificationss deleteSPECIFICATIONS
//
// # Delete a specifications
//
// default: genericError
//
//	200: specificationsDBResponse
func (controller *Controller) DeleteSPECIFICATIONS(c *gin.Context) {

	mutexSPECIFICATIONS.Lock()
	defer mutexSPECIFICATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECIFICATIONS", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPECIFICATIONS.GetDB()

	// Get model if exist
	var specificationsDB orm.SPECIFICATIONSDB
	if err := db.First(&specificationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&specificationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	specificationsDeleted := new(models.SPECIFICATIONS)
	specificationsDB.CopyBasicFieldsToSPECIFICATIONS(specificationsDeleted)

	// get stage instance from DB instance, and call callback function
	specificationsStaged := backRepo.BackRepoSPECIFICATIONS.Map_SPECIFICATIONSDBID_SPECIFICATIONSPtr[specificationsDB.ID]
	if specificationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specificationsStaged, specificationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
