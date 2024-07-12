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
var __ENUMVALUE__dummysDeclaration__ models.ENUMVALUE
var __ENUMVALUE_time__dummyDeclaration time.Duration

var mutexENUMVALUE sync.Mutex

// An ENUMVALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getENUMVALUE updateENUMVALUE deleteENUMVALUE
type ENUMVALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ENUMVALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postENUMVALUE updateENUMVALUE
type ENUMVALUEInput struct {
	// The ENUMVALUE to submit or modify
	// in: body
	ENUMVALUE *orm.ENUMVALUEAPI
}

// GetENUMVALUEs
//
// swagger:route GET /enumvalues enumvalues getENUMVALUEs
//
// # Get all enumvalues
//
// Responses:
// default: genericError
//
//	200: enumvalueDBResponse
func (controller *Controller) GetENUMVALUEs(c *gin.Context) {

	// source slice
	var enumvalueDBs []orm.ENUMVALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetENUMVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoENUMVALUE.GetDB()

	query := db.Find(&enumvalueDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	enumvalueAPIs := make([]orm.ENUMVALUEAPI, 0)

	// for each enumvalue, update fields from the database nullable fields
	for idx := range enumvalueDBs {
		enumvalueDB := &enumvalueDBs[idx]
		_ = enumvalueDB
		var enumvalueAPI orm.ENUMVALUEAPI

		// insertion point for updating fields
		enumvalueAPI.ID = enumvalueDB.ID
		enumvalueDB.CopyBasicFieldsToENUMVALUE_WOP(&enumvalueAPI.ENUMVALUE_WOP)
		enumvalueAPI.ENUMVALUEPointersEncoding = enumvalueDB.ENUMVALUEPointersEncoding
		enumvalueAPIs = append(enumvalueAPIs, enumvalueAPI)
	}

	c.JSON(http.StatusOK, enumvalueAPIs)
}

// PostENUMVALUE
//
// swagger:route POST /enumvalues enumvalues postENUMVALUE
//
// Creates a enumvalue
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostENUMVALUE(c *gin.Context) {

	mutexENUMVALUE.Lock()
	defer mutexENUMVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostENUMVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoENUMVALUE.GetDB()

	// Validate input
	var input orm.ENUMVALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create enumvalue
	enumvalueDB := orm.ENUMVALUEDB{}
	enumvalueDB.ENUMVALUEPointersEncoding = input.ENUMVALUEPointersEncoding
	enumvalueDB.CopyBasicFieldsFromENUMVALUE_WOP(&input.ENUMVALUE_WOP)

	query := db.Create(&enumvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoENUMVALUE.CheckoutPhaseOneInstance(&enumvalueDB)
	enumvalue := backRepo.BackRepoENUMVALUE.Map_ENUMVALUEDBID_ENUMVALUEPtr[enumvalueDB.ID]

	if enumvalue != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), enumvalue)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, enumvalueDB)
}

// GetENUMVALUE
//
// swagger:route GET /enumvalues/{ID} enumvalues getENUMVALUE
//
// Gets the details for a enumvalue.
//
// Responses:
// default: genericError
//
//	200: enumvalueDBResponse
func (controller *Controller) GetENUMVALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetENUMVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoENUMVALUE.GetDB()

	// Get enumvalueDB in DB
	var enumvalueDB orm.ENUMVALUEDB
	if err := db.First(&enumvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var enumvalueAPI orm.ENUMVALUEAPI
	enumvalueAPI.ID = enumvalueDB.ID
	enumvalueAPI.ENUMVALUEPointersEncoding = enumvalueDB.ENUMVALUEPointersEncoding
	enumvalueDB.CopyBasicFieldsToENUMVALUE_WOP(&enumvalueAPI.ENUMVALUE_WOP)

	c.JSON(http.StatusOK, enumvalueAPI)
}

// UpdateENUMVALUE
//
// swagger:route PATCH /enumvalues/{ID} enumvalues updateENUMVALUE
//
// # Update a enumvalue
//
// Responses:
// default: genericError
//
//	200: enumvalueDBResponse
func (controller *Controller) UpdateENUMVALUE(c *gin.Context) {

	mutexENUMVALUE.Lock()
	defer mutexENUMVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateENUMVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoENUMVALUE.GetDB()

	// Validate input
	var input orm.ENUMVALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var enumvalueDB orm.ENUMVALUEDB

	// fetch the enumvalue
	query := db.First(&enumvalueDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	enumvalueDB.CopyBasicFieldsFromENUMVALUE_WOP(&input.ENUMVALUE_WOP)
	enumvalueDB.ENUMVALUEPointersEncoding = input.ENUMVALUEPointersEncoding

	query = db.Model(&enumvalueDB).Updates(enumvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	enumvalueNew := new(models.ENUMVALUE)
	enumvalueDB.CopyBasicFieldsToENUMVALUE(enumvalueNew)

	// redeem pointers
	enumvalueDB.DecodePointers(backRepo, enumvalueNew)

	// get stage instance from DB instance, and call callback function
	enumvalueOld := backRepo.BackRepoENUMVALUE.Map_ENUMVALUEDBID_ENUMVALUEPtr[enumvalueDB.ID]
	if enumvalueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), enumvalueOld, enumvalueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the enumvalueDB
	c.JSON(http.StatusOK, enumvalueDB)
}

// DeleteENUMVALUE
//
// swagger:route DELETE /enumvalues/{ID} enumvalues deleteENUMVALUE
//
// # Delete a enumvalue
//
// default: genericError
//
//	200: enumvalueDBResponse
func (controller *Controller) DeleteENUMVALUE(c *gin.Context) {

	mutexENUMVALUE.Lock()
	defer mutexENUMVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteENUMVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoENUMVALUE.GetDB()

	// Get model if exist
	var enumvalueDB orm.ENUMVALUEDB
	if err := db.First(&enumvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&enumvalueDB)

	// get an instance (not staged) from DB instance, and call callback function
	enumvalueDeleted := new(models.ENUMVALUE)
	enumvalueDB.CopyBasicFieldsToENUMVALUE(enumvalueDeleted)

	// get stage instance from DB instance, and call callback function
	enumvalueStaged := backRepo.BackRepoENUMVALUE.Map_ENUMVALUEDBID_ENUMVALUEPtr[enumvalueDB.ID]
	if enumvalueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), enumvalueStaged, enumvalueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
