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
var __ATTRIBUTEDEFINITIONSTRING__dummysDeclaration__ models.ATTRIBUTEDEFINITIONSTRING
var __ATTRIBUTEDEFINITIONSTRING_time__dummyDeclaration time.Duration

var mutexATTRIBUTEDEFINITIONSTRING sync.Mutex

// An ATTRIBUTEDEFINITIONSTRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEDEFINITIONSTRING updateATTRIBUTEDEFINITIONSTRING deleteATTRIBUTEDEFINITIONSTRING
type ATTRIBUTEDEFINITIONSTRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEDEFINITIONSTRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEDEFINITIONSTRING updateATTRIBUTEDEFINITIONSTRING
type ATTRIBUTEDEFINITIONSTRINGInput struct {
	// The ATTRIBUTEDEFINITIONSTRING to submit or modify
	// in: body
	ATTRIBUTEDEFINITIONSTRING *orm.ATTRIBUTEDEFINITIONSTRINGAPI
}

// GetATTRIBUTEDEFINITIONSTRINGs
//
// swagger:route GET /attributedefinitionstrings attributedefinitionstrings getATTRIBUTEDEFINITIONSTRINGs
//
// # Get all attributedefinitionstrings
//
// Responses:
// default: genericError
//
//	200: attributedefinitionstringDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONSTRINGs(c *gin.Context) {

	// source slice
	var attributedefinitionstringDBs []orm.ATTRIBUTEDEFINITIONSTRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONSTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetDB()

	query := db.Find(&attributedefinitionstringDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributedefinitionstringAPIs := make([]orm.ATTRIBUTEDEFINITIONSTRINGAPI, 0)

	// for each attributedefinitionstring, update fields from the database nullable fields
	for idx := range attributedefinitionstringDBs {
		attributedefinitionstringDB := &attributedefinitionstringDBs[idx]
		_ = attributedefinitionstringDB
		var attributedefinitionstringAPI orm.ATTRIBUTEDEFINITIONSTRINGAPI

		// insertion point for updating fields
		attributedefinitionstringAPI.ID = attributedefinitionstringDB.ID
		attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING_WOP(&attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRING_WOP)
		attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRINGPointersEncoding = attributedefinitionstringDB.ATTRIBUTEDEFINITIONSTRINGPointersEncoding
		attributedefinitionstringAPIs = append(attributedefinitionstringAPIs, attributedefinitionstringAPI)
	}

	c.JSON(http.StatusOK, attributedefinitionstringAPIs)
}

// PostATTRIBUTEDEFINITIONSTRING
//
// swagger:route POST /attributedefinitionstrings attributedefinitionstrings postATTRIBUTEDEFINITIONSTRING
//
// Creates a attributedefinitionstring
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEDEFINITIONSTRING(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONSTRING.Lock()
	defer mutexATTRIBUTEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEDEFINITIONSTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONSTRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributedefinitionstring
	attributedefinitionstringDB := orm.ATTRIBUTEDEFINITIONSTRINGDB{}
	attributedefinitionstringDB.ATTRIBUTEDEFINITIONSTRINGPointersEncoding = input.ATTRIBUTEDEFINITIONSTRINGPointersEncoding
	attributedefinitionstringDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING_WOP(&input.ATTRIBUTEDEFINITIONSTRING_WOP)

	query := db.Create(&attributedefinitionstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseOneInstance(&attributedefinitionstringDB)
	attributedefinitionstring := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]

	if attributedefinitionstring != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributedefinitionstring)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributedefinitionstringDB)
}

// GetATTRIBUTEDEFINITIONSTRING
//
// swagger:route GET /attributedefinitionstrings/{ID} attributedefinitionstrings getATTRIBUTEDEFINITIONSTRING
//
// Gets the details for a attributedefinitionstring.
//
// Responses:
// default: genericError
//
//	200: attributedefinitionstringDBResponse
func (controller *Controller) GetATTRIBUTEDEFINITIONSTRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetDB()

	// Get attributedefinitionstringDB in DB
	var attributedefinitionstringDB orm.ATTRIBUTEDEFINITIONSTRINGDB
	if err := db.First(&attributedefinitionstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributedefinitionstringAPI orm.ATTRIBUTEDEFINITIONSTRINGAPI
	attributedefinitionstringAPI.ID = attributedefinitionstringDB.ID
	attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRINGPointersEncoding = attributedefinitionstringDB.ATTRIBUTEDEFINITIONSTRINGPointersEncoding
	attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING_WOP(&attributedefinitionstringAPI.ATTRIBUTEDEFINITIONSTRING_WOP)

	c.JSON(http.StatusOK, attributedefinitionstringAPI)
}

// UpdateATTRIBUTEDEFINITIONSTRING
//
// swagger:route PATCH /attributedefinitionstrings/{ID} attributedefinitionstrings updateATTRIBUTEDEFINITIONSTRING
//
// # Update a attributedefinitionstring
//
// Responses:
// default: genericError
//
//	200: attributedefinitionstringDBResponse
func (controller *Controller) UpdateATTRIBUTEDEFINITIONSTRING(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONSTRING.Lock()
	defer mutexATTRIBUTEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTEDEFINITIONSTRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributedefinitionstringDB orm.ATTRIBUTEDEFINITIONSTRINGDB

	// fetch the attributedefinitionstring
	query := db.First(&attributedefinitionstringDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributedefinitionstringDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING_WOP(&input.ATTRIBUTEDEFINITIONSTRING_WOP)
	attributedefinitionstringDB.ATTRIBUTEDEFINITIONSTRINGPointersEncoding = input.ATTRIBUTEDEFINITIONSTRINGPointersEncoding

	query = db.Model(&attributedefinitionstringDB).Updates(attributedefinitionstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionstringNew := new(models.ATTRIBUTEDEFINITIONSTRING)
	attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING(attributedefinitionstringNew)

	// redeem pointers
	attributedefinitionstringDB.DecodePointers(backRepo, attributedefinitionstringNew)

	// get stage instance from DB instance, and call callback function
	attributedefinitionstringOld := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]
	if attributedefinitionstringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributedefinitionstringOld, attributedefinitionstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributedefinitionstringDB
	c.JSON(http.StatusOK, attributedefinitionstringDB)
}

// DeleteATTRIBUTEDEFINITIONSTRING
//
// swagger:route DELETE /attributedefinitionstrings/{ID} attributedefinitionstrings deleteATTRIBUTEDEFINITIONSTRING
//
// # Delete a attributedefinitionstring
//
// default: genericError
//
//	200: attributedefinitionstringDBResponse
func (controller *Controller) DeleteATTRIBUTEDEFINITIONSTRING(c *gin.Context) {

	mutexATTRIBUTEDEFINITIONSTRING.Lock()
	defer mutexATTRIBUTEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.GetDB()

	// Get model if exist
	var attributedefinitionstringDB orm.ATTRIBUTEDEFINITIONSTRINGDB
	if err := db.First(&attributedefinitionstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributedefinitionstringDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributedefinitionstringDeleted := new(models.ATTRIBUTEDEFINITIONSTRING)
	attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING(attributedefinitionstringDeleted)

	// get stage instance from DB instance, and call callback function
	attributedefinitionstringStaged := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]
	if attributedefinitionstringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributedefinitionstringStaged, attributedefinitionstringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
