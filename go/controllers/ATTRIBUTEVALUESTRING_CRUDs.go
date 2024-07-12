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
var __ATTRIBUTEVALUESTRING__dummysDeclaration__ models.ATTRIBUTEVALUESTRING
var __ATTRIBUTEVALUESTRING_time__dummyDeclaration time.Duration

var mutexATTRIBUTEVALUESTRING sync.Mutex

// An ATTRIBUTEVALUESTRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTEVALUESTRING updateATTRIBUTEVALUESTRING deleteATTRIBUTEVALUESTRING
type ATTRIBUTEVALUESTRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTEVALUESTRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTEVALUESTRING updateATTRIBUTEVALUESTRING
type ATTRIBUTEVALUESTRINGInput struct {
	// The ATTRIBUTEVALUESTRING to submit or modify
	// in: body
	ATTRIBUTEVALUESTRING *orm.ATTRIBUTEVALUESTRINGAPI
}

// GetATTRIBUTEVALUESTRINGs
//
// swagger:route GET /attributevaluestrings attributevaluestrings getATTRIBUTEVALUESTRINGs
//
// # Get all attributevaluestrings
//
// Responses:
// default: genericError
//
//	200: attributevaluestringDBResponse
func (controller *Controller) GetATTRIBUTEVALUESTRINGs(c *gin.Context) {

	// source slice
	var attributevaluestringDBs []orm.ATTRIBUTEVALUESTRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUESTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUESTRING.GetDB()

	query := db.Find(&attributevaluestringDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributevaluestringAPIs := make([]orm.ATTRIBUTEVALUESTRINGAPI, 0)

	// for each attributevaluestring, update fields from the database nullable fields
	for idx := range attributevaluestringDBs {
		attributevaluestringDB := &attributevaluestringDBs[idx]
		_ = attributevaluestringDB
		var attributevaluestringAPI orm.ATTRIBUTEVALUESTRINGAPI

		// insertion point for updating fields
		attributevaluestringAPI.ID = attributevaluestringDB.ID
		attributevaluestringDB.CopyBasicFieldsToATTRIBUTEVALUESTRING_WOP(&attributevaluestringAPI.ATTRIBUTEVALUESTRING_WOP)
		attributevaluestringAPI.ATTRIBUTEVALUESTRINGPointersEncoding = attributevaluestringDB.ATTRIBUTEVALUESTRINGPointersEncoding
		attributevaluestringAPIs = append(attributevaluestringAPIs, attributevaluestringAPI)
	}

	c.JSON(http.StatusOK, attributevaluestringAPIs)
}

// PostATTRIBUTEVALUESTRING
//
// swagger:route POST /attributevaluestrings attributevaluestrings postATTRIBUTEVALUESTRING
//
// Creates a attributevaluestring
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTEVALUESTRING(c *gin.Context) {

	mutexATTRIBUTEVALUESTRING.Lock()
	defer mutexATTRIBUTEVALUESTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTEVALUESTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUESTRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUESTRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributevaluestring
	attributevaluestringDB := orm.ATTRIBUTEVALUESTRINGDB{}
	attributevaluestringDB.ATTRIBUTEVALUESTRINGPointersEncoding = input.ATTRIBUTEVALUESTRINGPointersEncoding
	attributevaluestringDB.CopyBasicFieldsFromATTRIBUTEVALUESTRING_WOP(&input.ATTRIBUTEVALUESTRING_WOP)

	query := db.Create(&attributevaluestringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTEVALUESTRING.CheckoutPhaseOneInstance(&attributevaluestringDB)
	attributevaluestring := backRepo.BackRepoATTRIBUTEVALUESTRING.Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGPtr[attributevaluestringDB.ID]

	if attributevaluestring != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributevaluestring)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributevaluestringDB)
}

// GetATTRIBUTEVALUESTRING
//
// swagger:route GET /attributevaluestrings/{ID} attributevaluestrings getATTRIBUTEVALUESTRING
//
// Gets the details for a attributevaluestring.
//
// Responses:
// default: genericError
//
//	200: attributevaluestringDBResponse
func (controller *Controller) GetATTRIBUTEVALUESTRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTEVALUESTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUESTRING.GetDB()

	// Get attributevaluestringDB in DB
	var attributevaluestringDB orm.ATTRIBUTEVALUESTRINGDB
	if err := db.First(&attributevaluestringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributevaluestringAPI orm.ATTRIBUTEVALUESTRINGAPI
	attributevaluestringAPI.ID = attributevaluestringDB.ID
	attributevaluestringAPI.ATTRIBUTEVALUESTRINGPointersEncoding = attributevaluestringDB.ATTRIBUTEVALUESTRINGPointersEncoding
	attributevaluestringDB.CopyBasicFieldsToATTRIBUTEVALUESTRING_WOP(&attributevaluestringAPI.ATTRIBUTEVALUESTRING_WOP)

	c.JSON(http.StatusOK, attributevaluestringAPI)
}

// UpdateATTRIBUTEVALUESTRING
//
// swagger:route PATCH /attributevaluestrings/{ID} attributevaluestrings updateATTRIBUTEVALUESTRING
//
// # Update a attributevaluestring
//
// Responses:
// default: genericError
//
//	200: attributevaluestringDBResponse
func (controller *Controller) UpdateATTRIBUTEVALUESTRING(c *gin.Context) {

	mutexATTRIBUTEVALUESTRING.Lock()
	defer mutexATTRIBUTEVALUESTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTEVALUESTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUESTRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTEVALUESTRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributevaluestringDB orm.ATTRIBUTEVALUESTRINGDB

	// fetch the attributevaluestring
	query := db.First(&attributevaluestringDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributevaluestringDB.CopyBasicFieldsFromATTRIBUTEVALUESTRING_WOP(&input.ATTRIBUTEVALUESTRING_WOP)
	attributevaluestringDB.ATTRIBUTEVALUESTRINGPointersEncoding = input.ATTRIBUTEVALUESTRINGPointersEncoding

	query = db.Model(&attributevaluestringDB).Updates(attributevaluestringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluestringNew := new(models.ATTRIBUTEVALUESTRING)
	attributevaluestringDB.CopyBasicFieldsToATTRIBUTEVALUESTRING(attributevaluestringNew)

	// redeem pointers
	attributevaluestringDB.DecodePointers(backRepo, attributevaluestringNew)

	// get stage instance from DB instance, and call callback function
	attributevaluestringOld := backRepo.BackRepoATTRIBUTEVALUESTRING.Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGPtr[attributevaluestringDB.ID]
	if attributevaluestringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributevaluestringOld, attributevaluestringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributevaluestringDB
	c.JSON(http.StatusOK, attributevaluestringDB)
}

// DeleteATTRIBUTEVALUESTRING
//
// swagger:route DELETE /attributevaluestrings/{ID} attributevaluestrings deleteATTRIBUTEVALUESTRING
//
// # Delete a attributevaluestring
//
// default: genericError
//
//	200: attributevaluestringDBResponse
func (controller *Controller) DeleteATTRIBUTEVALUESTRING(c *gin.Context) {

	mutexATTRIBUTEVALUESTRING.Lock()
	defer mutexATTRIBUTEVALUESTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTEVALUESTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTEVALUESTRING.GetDB()

	// Get model if exist
	var attributevaluestringDB orm.ATTRIBUTEVALUESTRINGDB
	if err := db.First(&attributevaluestringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributevaluestringDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributevaluestringDeleted := new(models.ATTRIBUTEVALUESTRING)
	attributevaluestringDB.CopyBasicFieldsToATTRIBUTEVALUESTRING(attributevaluestringDeleted)

	// get stage instance from DB instance, and call callback function
	attributevaluestringStaged := backRepo.BackRepoATTRIBUTEVALUESTRING.Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGPtr[attributevaluestringDB.ID]
	if attributevaluestringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributevaluestringStaged, attributevaluestringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
