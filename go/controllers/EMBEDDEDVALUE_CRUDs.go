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
var __EMBEDDEDVALUE__dummysDeclaration__ models.EMBEDDEDVALUE
var __EMBEDDEDVALUE_time__dummyDeclaration time.Duration

var mutexEMBEDDEDVALUE sync.Mutex

// An EMBEDDEDVALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEMBEDDEDVALUE updateEMBEDDEDVALUE deleteEMBEDDEDVALUE
type EMBEDDEDVALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EMBEDDEDVALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEMBEDDEDVALUE updateEMBEDDEDVALUE
type EMBEDDEDVALUEInput struct {
	// The EMBEDDEDVALUE to submit or modify
	// in: body
	EMBEDDEDVALUE *orm.EMBEDDEDVALUEAPI
}

// GetEMBEDDEDVALUEs
//
// swagger:route GET /embeddedvalues embeddedvalues getEMBEDDEDVALUEs
//
// # Get all embeddedvalues
//
// Responses:
// default: genericError
//
//	200: embeddedvalueDBResponse
func (controller *Controller) GetEMBEDDEDVALUEs(c *gin.Context) {

	// source slice
	var embeddedvalueDBs []orm.EMBEDDEDVALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEMBEDDEDVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEMBEDDEDVALUE.GetDB()

	query := db.Find(&embeddedvalueDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	embeddedvalueAPIs := make([]orm.EMBEDDEDVALUEAPI, 0)

	// for each embeddedvalue, update fields from the database nullable fields
	for idx := range embeddedvalueDBs {
		embeddedvalueDB := &embeddedvalueDBs[idx]
		_ = embeddedvalueDB
		var embeddedvalueAPI orm.EMBEDDEDVALUEAPI

		// insertion point for updating fields
		embeddedvalueAPI.ID = embeddedvalueDB.ID
		embeddedvalueDB.CopyBasicFieldsToEMBEDDEDVALUE_WOP(&embeddedvalueAPI.EMBEDDEDVALUE_WOP)
		embeddedvalueAPI.EMBEDDEDVALUEPointersEncoding = embeddedvalueDB.EMBEDDEDVALUEPointersEncoding
		embeddedvalueAPIs = append(embeddedvalueAPIs, embeddedvalueAPI)
	}

	c.JSON(http.StatusOK, embeddedvalueAPIs)
}

// PostEMBEDDEDVALUE
//
// swagger:route POST /embeddedvalues embeddedvalues postEMBEDDEDVALUE
//
// Creates a embeddedvalue
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEMBEDDEDVALUE(c *gin.Context) {

	mutexEMBEDDEDVALUE.Lock()
	defer mutexEMBEDDEDVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEMBEDDEDVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEMBEDDEDVALUE.GetDB()

	// Validate input
	var input orm.EMBEDDEDVALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create embeddedvalue
	embeddedvalueDB := orm.EMBEDDEDVALUEDB{}
	embeddedvalueDB.EMBEDDEDVALUEPointersEncoding = input.EMBEDDEDVALUEPointersEncoding
	embeddedvalueDB.CopyBasicFieldsFromEMBEDDEDVALUE_WOP(&input.EMBEDDEDVALUE_WOP)

	query := db.Create(&embeddedvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEMBEDDEDVALUE.CheckoutPhaseOneInstance(&embeddedvalueDB)
	embeddedvalue := backRepo.BackRepoEMBEDDEDVALUE.Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEPtr[embeddedvalueDB.ID]

	if embeddedvalue != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), embeddedvalue)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, embeddedvalueDB)
}

// GetEMBEDDEDVALUE
//
// swagger:route GET /embeddedvalues/{ID} embeddedvalues getEMBEDDEDVALUE
//
// Gets the details for a embeddedvalue.
//
// Responses:
// default: genericError
//
//	200: embeddedvalueDBResponse
func (controller *Controller) GetEMBEDDEDVALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEMBEDDEDVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEMBEDDEDVALUE.GetDB()

	// Get embeddedvalueDB in DB
	var embeddedvalueDB orm.EMBEDDEDVALUEDB
	if err := db.First(&embeddedvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var embeddedvalueAPI orm.EMBEDDEDVALUEAPI
	embeddedvalueAPI.ID = embeddedvalueDB.ID
	embeddedvalueAPI.EMBEDDEDVALUEPointersEncoding = embeddedvalueDB.EMBEDDEDVALUEPointersEncoding
	embeddedvalueDB.CopyBasicFieldsToEMBEDDEDVALUE_WOP(&embeddedvalueAPI.EMBEDDEDVALUE_WOP)

	c.JSON(http.StatusOK, embeddedvalueAPI)
}

// UpdateEMBEDDEDVALUE
//
// swagger:route PATCH /embeddedvalues/{ID} embeddedvalues updateEMBEDDEDVALUE
//
// # Update a embeddedvalue
//
// Responses:
// default: genericError
//
//	200: embeddedvalueDBResponse
func (controller *Controller) UpdateEMBEDDEDVALUE(c *gin.Context) {

	mutexEMBEDDEDVALUE.Lock()
	defer mutexEMBEDDEDVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEMBEDDEDVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEMBEDDEDVALUE.GetDB()

	// Validate input
	var input orm.EMBEDDEDVALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var embeddedvalueDB orm.EMBEDDEDVALUEDB

	// fetch the embeddedvalue
	query := db.First(&embeddedvalueDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	embeddedvalueDB.CopyBasicFieldsFromEMBEDDEDVALUE_WOP(&input.EMBEDDEDVALUE_WOP)
	embeddedvalueDB.EMBEDDEDVALUEPointersEncoding = input.EMBEDDEDVALUEPointersEncoding

	query = db.Model(&embeddedvalueDB).Updates(embeddedvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	embeddedvalueNew := new(models.EMBEDDEDVALUE)
	embeddedvalueDB.CopyBasicFieldsToEMBEDDEDVALUE(embeddedvalueNew)

	// redeem pointers
	embeddedvalueDB.DecodePointers(backRepo, embeddedvalueNew)

	// get stage instance from DB instance, and call callback function
	embeddedvalueOld := backRepo.BackRepoEMBEDDEDVALUE.Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEPtr[embeddedvalueDB.ID]
	if embeddedvalueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), embeddedvalueOld, embeddedvalueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the embeddedvalueDB
	c.JSON(http.StatusOK, embeddedvalueDB)
}

// DeleteEMBEDDEDVALUE
//
// swagger:route DELETE /embeddedvalues/{ID} embeddedvalues deleteEMBEDDEDVALUE
//
// # Delete a embeddedvalue
//
// default: genericError
//
//	200: embeddedvalueDBResponse
func (controller *Controller) DeleteEMBEDDEDVALUE(c *gin.Context) {

	mutexEMBEDDEDVALUE.Lock()
	defer mutexEMBEDDEDVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEMBEDDEDVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEMBEDDEDVALUE.GetDB()

	// Get model if exist
	var embeddedvalueDB orm.EMBEDDEDVALUEDB
	if err := db.First(&embeddedvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&embeddedvalueDB)

	// get an instance (not staged) from DB instance, and call callback function
	embeddedvalueDeleted := new(models.EMBEDDEDVALUE)
	embeddedvalueDB.CopyBasicFieldsToEMBEDDEDVALUE(embeddedvalueDeleted)

	// get stage instance from DB instance, and call callback function
	embeddedvalueStaged := backRepo.BackRepoEMBEDDEDVALUE.Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEPtr[embeddedvalueDB.ID]
	if embeddedvalueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), embeddedvalueStaged, embeddedvalueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
