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
var __PROPERTIES__dummysDeclaration__ models.PROPERTIES
var __PROPERTIES_time__dummyDeclaration time.Duration

var mutexPROPERTIES sync.Mutex

// An PROPERTIESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPROPERTIES updatePROPERTIES deletePROPERTIES
type PROPERTIESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PROPERTIESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPROPERTIES updatePROPERTIES
type PROPERTIESInput struct {
	// The PROPERTIES to submit or modify
	// in: body
	PROPERTIES *orm.PROPERTIESAPI
}

// GetPROPERTIESs
//
// swagger:route GET /propertiess propertiess getPROPERTIESs
//
// # Get all propertiess
//
// Responses:
// default: genericError
//
//	200: propertiesDBResponse
func (controller *Controller) GetPROPERTIESs(c *gin.Context) {

	// source slice
	var propertiesDBs []orm.PROPERTIESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPROPERTIESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPROPERTIES.GetDB()

	query := db.Find(&propertiesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	propertiesAPIs := make([]orm.PROPERTIESAPI, 0)

	// for each properties, update fields from the database nullable fields
	for idx := range propertiesDBs {
		propertiesDB := &propertiesDBs[idx]
		_ = propertiesDB
		var propertiesAPI orm.PROPERTIESAPI

		// insertion point for updating fields
		propertiesAPI.ID = propertiesDB.ID
		propertiesDB.CopyBasicFieldsToPROPERTIES_WOP(&propertiesAPI.PROPERTIES_WOP)
		propertiesAPI.PROPERTIESPointersEncoding = propertiesDB.PROPERTIESPointersEncoding
		propertiesAPIs = append(propertiesAPIs, propertiesAPI)
	}

	c.JSON(http.StatusOK, propertiesAPIs)
}

// PostPROPERTIES
//
// swagger:route POST /propertiess propertiess postPROPERTIES
//
// Creates a properties
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPROPERTIES(c *gin.Context) {

	mutexPROPERTIES.Lock()
	defer mutexPROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPROPERTIESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPROPERTIES.GetDB()

	// Validate input
	var input orm.PROPERTIESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create properties
	propertiesDB := orm.PROPERTIESDB{}
	propertiesDB.PROPERTIESPointersEncoding = input.PROPERTIESPointersEncoding
	propertiesDB.CopyBasicFieldsFromPROPERTIES_WOP(&input.PROPERTIES_WOP)

	query := db.Create(&propertiesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPROPERTIES.CheckoutPhaseOneInstance(&propertiesDB)
	properties := backRepo.BackRepoPROPERTIES.Map_PROPERTIESDBID_PROPERTIESPtr[propertiesDB.ID]

	if properties != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), properties)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, propertiesDB)
}

// GetPROPERTIES
//
// swagger:route GET /propertiess/{ID} propertiess getPROPERTIES
//
// Gets the details for a properties.
//
// Responses:
// default: genericError
//
//	200: propertiesDBResponse
func (controller *Controller) GetPROPERTIES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPROPERTIES.GetDB()

	// Get propertiesDB in DB
	var propertiesDB orm.PROPERTIESDB
	if err := db.First(&propertiesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var propertiesAPI orm.PROPERTIESAPI
	propertiesAPI.ID = propertiesDB.ID
	propertiesAPI.PROPERTIESPointersEncoding = propertiesDB.PROPERTIESPointersEncoding
	propertiesDB.CopyBasicFieldsToPROPERTIES_WOP(&propertiesAPI.PROPERTIES_WOP)

	c.JSON(http.StatusOK, propertiesAPI)
}

// UpdatePROPERTIES
//
// swagger:route PATCH /propertiess/{ID} propertiess updatePROPERTIES
//
// # Update a properties
//
// Responses:
// default: genericError
//
//	200: propertiesDBResponse
func (controller *Controller) UpdatePROPERTIES(c *gin.Context) {

	mutexPROPERTIES.Lock()
	defer mutexPROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPROPERTIES.GetDB()

	// Validate input
	var input orm.PROPERTIESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var propertiesDB orm.PROPERTIESDB

	// fetch the properties
	query := db.First(&propertiesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	propertiesDB.CopyBasicFieldsFromPROPERTIES_WOP(&input.PROPERTIES_WOP)
	propertiesDB.PROPERTIESPointersEncoding = input.PROPERTIESPointersEncoding

	query = db.Model(&propertiesDB).Updates(propertiesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	propertiesNew := new(models.PROPERTIES)
	propertiesDB.CopyBasicFieldsToPROPERTIES(propertiesNew)

	// redeem pointers
	propertiesDB.DecodePointers(backRepo, propertiesNew)

	// get stage instance from DB instance, and call callback function
	propertiesOld := backRepo.BackRepoPROPERTIES.Map_PROPERTIESDBID_PROPERTIESPtr[propertiesDB.ID]
	if propertiesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), propertiesOld, propertiesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the propertiesDB
	c.JSON(http.StatusOK, propertiesDB)
}

// DeletePROPERTIES
//
// swagger:route DELETE /propertiess/{ID} propertiess deletePROPERTIES
//
// # Delete a properties
//
// default: genericError
//
//	200: propertiesDBResponse
func (controller *Controller) DeletePROPERTIES(c *gin.Context) {

	mutexPROPERTIES.Lock()
	defer mutexPROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPROPERTIES.GetDB()

	// Get model if exist
	var propertiesDB orm.PROPERTIESDB
	if err := db.First(&propertiesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&propertiesDB)

	// get an instance (not staged) from DB instance, and call callback function
	propertiesDeleted := new(models.PROPERTIES)
	propertiesDB.CopyBasicFieldsToPROPERTIES(propertiesDeleted)

	// get stage instance from DB instance, and call callback function
	propertiesStaged := backRepo.BackRepoPROPERTIES.Map_PROPERTIESDBID_PROPERTIESPtr[propertiesDB.ID]
	if propertiesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), propertiesStaged, propertiesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
