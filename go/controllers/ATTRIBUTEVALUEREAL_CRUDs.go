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
var __ATTRIBUTEVALUEREAL__dummysDeclaration__ models.ATTRIBUTEVALUEREAL
var __ATTRIBUTEVALUEREAL_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUEREAL sync.Mutex

// An ATTRIBUTEVALUEREALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUEREAL updateATTRIBUTEVALUEREAL deleteATTRIBUTEVALUEREAL
type ATTRIBUTEVALUEREALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUEREALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUEREAL updateATTRIBUTEVALUEREAL
type ATTRIBUTEVALUEREALInput struct {
	// The ATTRIBUTEVALUEREAL to submit or modify
	// in: body
	ATTRIBUTEVALUEREAL *orm.ATTRIBUTEVALUEREALAPI
}

// GetATTRIBUTEVALUEREALs
//
// swagger:route GET /attributevaluereals attributevaluereals getATTRIBUTEVALUEREALs
//
// # Get all attributevaluereals
//
// Responses:
// default: genericError
//
//	200: attributevaluerealDBResponse
func (controller *Controller) GetATTRIBUTEVALUEREALs(c *gin.Context) {

	// source slice
	var attributevaluerealDBs []orm.ATTRIBUTEVALUEREALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEREAL.GetDB()

	query := db.Find(&attributevaluerealDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevaluerealAPIs := make([]orm.ATTRIBUTEVALUEREALAPI, 0)

	// for each attributevaluereal, update fields from the database nullable fields
	for idx := range attributevaluerealDBs {
		attributevaluerealDB := &attributevaluerealDBs[idx]
		_ = attributevaluerealDB
		var attributevaluerealAPI orm.ATTRIBUTEVALUEREALAPI

		// insertion point for updating fields
		attributevaluerealAPI.ID = attributevaluerealDB.ID
		attributevaluerealDB.CopyBasicFieldsToATTRIBUTEVALUEREAL_WOP(&attributevaluerealAPI.ATTRIBUTEVALUEREAL_WOP)
		attributevaluerealAPI.ATTRIBUTEVALUEREALPointersEncoding = attributevaluerealDB.ATTRIBUTEVALUEREALPointersEncoding
		attributevaluerealAPIs = append(attributevaluerealAPIs, attributevaluerealAPI)
	}

	c.JSON(http.StatusOK, attributevaluerealAPIs)
}

// PostATTRIBUTEVALUEREAL
//
// swagger:route POST /attributevaluereals attributevaluereals postATTRIBUTEVALUEREAL
//
// Creates a attributevaluereal
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUEREAL(c *gin.Context) {

	mutexATTRIBUTEVALUEREAL.Lock()
	defer mutexATTRIBUTEVALUEREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUEREALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEREAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEREALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevaluereal
	attributevaluerealDB := orm.ATTRIBUTEVALUEREALDB{}
	attributevaluerealDB.ATTRIBUTEVALUEREALPointersEncoding = input.ATTRIBUTEVALUEREALPointersEncoding
	attributevaluerealDB.CopyBasicFieldsFromATTRIBUTEVALUEREAL_WOP(&input.ATTRIBUTEVALUEREAL_WOP)

	query := db.Create(&attributevaluerealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUEREAL.CheckoutPhaseOneInstance(&attributevaluerealDB)
	attributevaluereal := backRepo.BackRepoATTRIBUTEVALUEREAL.Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALPtr[attributevaluerealDB.ID]

	if attributevaluereal != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevaluereal)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevaluerealDB)
}

// GetATTRIBUTEVALUEREAL
//
// swagger:route GET /attributevaluereals/{ID} attributevaluereals getATTRIBUTEVALUEREAL
//
// Gets the details for a attributevaluereal.
//
// Responses:
// default: genericError
//
//	200: attributevaluerealDBResponse
func (controller *Controller) GetATTRIBUTEVALUEREAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUEREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEREAL.GetDB()

	// Get attributevaluerealDB in DB
	var attributevaluerealDB orm.ATTRIBUTEVALUEREALDB
	if err := db.First(&attributevaluerealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevaluerealAPI orm.ATTRIBUTEVALUEREALAPI
	attributevaluerealAPI.ID = attributevaluerealDB.ID
	attributevaluerealAPI.ATTRIBUTEVALUEREALPointersEncoding = attributevaluerealDB.ATTRIBUTEVALUEREALPointersEncoding
	attributevaluerealDB.CopyBasicFieldsToATTRIBUTEVALUEREAL_WOP(&attributevaluerealAPI.ATTRIBUTEVALUEREAL_WOP)

	c.JSON(http.StatusOK, attributevaluerealAPI)
}

// UpdateATTRIBUTEVALUEREAL
//
// swagger:route PATCH /attributevaluereals/{ID} attributevaluereals updateATTRIBUTEVALUEREAL
//
// # Update a attributevaluereal
//
// Responses:
// default: genericError
//
//	200: attributevaluerealDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUEREAL(c *gin.Context) {

	mutexATTRIBUTEVALUEREAL.Lock()
	defer mutexATTRIBUTEVALUEREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUEREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEREAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUEREALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevaluerealDB orm.ATTRIBUTEVALUEREALDB

	// fetch the attributevaluereal
	query := db.First(&attributevaluerealDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevaluerealDB.CopyBasicFieldsFromATTRIBUTEVALUEREAL_WOP(&input.ATTRIBUTEVALUEREAL_WOP)
	attributevaluerealDB.ATTRIBUTEVALUEREALPointersEncoding = input.ATTRIBUTEVALUEREALPointersEncoding

	query = db.Model(&attributevaluerealDB).Updates(attributevaluerealDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluerealNew := new(models.ATTRIBUTEVALUEREAL)
	attributevaluerealDB.CopyBasicFieldsToATTRIBUTEVALUEREAL(attributevaluerealNew)

	// redeem pointers
	attributevaluerealDB.DecodePointers(backRepo, attributevaluerealNew)

	// get stage instance from DB instance, and call callback function
	attributevaluerealOld := backRepo.BackRepoATTRIBUTEVALUEREAL.Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALPtr[attributevaluerealDB.ID]
	if attributevaluerealOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevaluerealOld, attributevaluerealNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevaluerealDB
	c.JSON(http.StatusOK, attributevaluerealDB)
}

// DeleteATTRIBUTEVALUEREAL
//
// swagger:route DELETE /attributevaluereals/{ID} attributevaluereals deleteATTRIBUTEVALUEREAL
//
// # Delete a attributevaluereal
//
// default: genericError
//
//	200: attributevaluerealDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUEREAL(c *gin.Context) {

	mutexATTRIBUTEVALUEREAL.Lock()
	defer mutexATTRIBUTEVALUEREAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUEREAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUEREAL.GetDB()

	// Get model if exist
	var attributevaluerealDB orm.ATTRIBUTEVALUEREALDB
	if err := db.First(&attributevaluerealDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevaluerealDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluerealDeleted := new(models.ATTRIBUTEVALUEREAL)
	attributevaluerealDB.CopyBasicFieldsToATTRIBUTEVALUEREAL(attributevaluerealDeleted)

	// get stage instance from DB instance, and call callback function
	attributevaluerealStaged := backRepo.BackRepoATTRIBUTEVALUEREAL.Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALPtr[attributevaluerealDB.ID]
	if attributevaluerealStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevaluerealStaged, attributevaluerealDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
