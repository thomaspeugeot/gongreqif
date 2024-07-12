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
var __DEFAULTVALUE__dummysDeclaration__ models.DEFAULTVALUE
var __DEFAULTVALUE_time__dummyDeclaration time.Duration

var mutexDEFAULTVALUE sync.Mutex

// An DEFAULTVALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDEFAULTVALUE updateDEFAULTVALUE deleteDEFAULTVALUE
type DEFAULTVALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DEFAULTVALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDEFAULTVALUE updateDEFAULTVALUE
type DEFAULTVALUEInput struct {
	// The DEFAULTVALUE to submit or modify
	// in: body
	DEFAULTVALUE *orm.DEFAULTVALUEAPI
}

// GetDEFAULTVALUEs
//
// swagger:route GET /defaultvalues defaultvalues getDEFAULTVALUEs
//
// # Get all defaultvalues
//
// Responses:
// default: genericError
//
//	200: defaultvalueDBResponse
func (controller *Controller) GetDEFAULTVALUEs(c *gin.Context) {

	// source slice
	var defaultvalueDBs []orm.DEFAULTVALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDEFAULTVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFAULTVALUE.GetDB()

	query := db.Find(&defaultvalueDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	defaultvalueAPIs := make([]orm.DEFAULTVALUEAPI, 0)

	// for each defaultvalue, update fields from the database nullable fields
	for idx := range defaultvalueDBs {
		defaultvalueDB := &defaultvalueDBs[idx]
		_ = defaultvalueDB
		var defaultvalueAPI orm.DEFAULTVALUEAPI

		// insertion point for updating fields
		defaultvalueAPI.ID = defaultvalueDB.ID
		defaultvalueDB.CopyBasicFieldsToDEFAULTVALUE_WOP(&defaultvalueAPI.DEFAULTVALUE_WOP)
		defaultvalueAPI.DEFAULTVALUEPointersEncoding = defaultvalueDB.DEFAULTVALUEPointersEncoding
		defaultvalueAPIs = append(defaultvalueAPIs, defaultvalueAPI)
	}

	c.JSON(http.StatusOK, defaultvalueAPIs)
}

// PostDEFAULTVALUE
//
// swagger:route POST /defaultvalues defaultvalues postDEFAULTVALUE
//
// Creates a defaultvalue
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDEFAULTVALUE(c *gin.Context) {

	mutexDEFAULTVALUE.Lock()
	defer mutexDEFAULTVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDEFAULTVALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFAULTVALUE.GetDB()

	// Validate input
	var input orm.DEFAULTVALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create defaultvalue
	defaultvalueDB := orm.DEFAULTVALUEDB{}
	defaultvalueDB.DEFAULTVALUEPointersEncoding = input.DEFAULTVALUEPointersEncoding
	defaultvalueDB.CopyBasicFieldsFromDEFAULTVALUE_WOP(&input.DEFAULTVALUE_WOP)

	query := db.Create(&defaultvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDEFAULTVALUE.CheckoutPhaseOneInstance(&defaultvalueDB)
	defaultvalue := backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr[defaultvalueDB.ID]

	if defaultvalue != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), defaultvalue)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, defaultvalueDB)
}

// GetDEFAULTVALUE
//
// swagger:route GET /defaultvalues/{ID} defaultvalues getDEFAULTVALUE
//
// Gets the details for a defaultvalue.
//
// Responses:
// default: genericError
//
//	200: defaultvalueDBResponse
func (controller *Controller) GetDEFAULTVALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDEFAULTVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFAULTVALUE.GetDB()

	// Get defaultvalueDB in DB
	var defaultvalueDB orm.DEFAULTVALUEDB
	if err := db.First(&defaultvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var defaultvalueAPI orm.DEFAULTVALUEAPI
	defaultvalueAPI.ID = defaultvalueDB.ID
	defaultvalueAPI.DEFAULTVALUEPointersEncoding = defaultvalueDB.DEFAULTVALUEPointersEncoding
	defaultvalueDB.CopyBasicFieldsToDEFAULTVALUE_WOP(&defaultvalueAPI.DEFAULTVALUE_WOP)

	c.JSON(http.StatusOK, defaultvalueAPI)
}

// UpdateDEFAULTVALUE
//
// swagger:route PATCH /defaultvalues/{ID} defaultvalues updateDEFAULTVALUE
//
// # Update a defaultvalue
//
// Responses:
// default: genericError
//
//	200: defaultvalueDBResponse
func (controller *Controller) UpdateDEFAULTVALUE(c *gin.Context) {

	mutexDEFAULTVALUE.Lock()
	defer mutexDEFAULTVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDEFAULTVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFAULTVALUE.GetDB()

	// Validate input
	var input orm.DEFAULTVALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var defaultvalueDB orm.DEFAULTVALUEDB

	// fetch the defaultvalue
	query := db.First(&defaultvalueDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	defaultvalueDB.CopyBasicFieldsFromDEFAULTVALUE_WOP(&input.DEFAULTVALUE_WOP)
	defaultvalueDB.DEFAULTVALUEPointersEncoding = input.DEFAULTVALUEPointersEncoding

	query = db.Model(&defaultvalueDB).Updates(defaultvalueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	defaultvalueNew := new(models.DEFAULTVALUE)
	defaultvalueDB.CopyBasicFieldsToDEFAULTVALUE(defaultvalueNew)

	// redeem pointers
	defaultvalueDB.DecodePointers(backRepo, defaultvalueNew)

	// get stage instance from DB instance, and call callback function
	defaultvalueOld := backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr[defaultvalueDB.ID]
	if defaultvalueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), defaultvalueOld, defaultvalueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the defaultvalueDB
	c.JSON(http.StatusOK, defaultvalueDB)
}

// DeleteDEFAULTVALUE
//
// swagger:route DELETE /defaultvalues/{ID} defaultvalues deleteDEFAULTVALUE
//
// # Delete a defaultvalue
//
// default: genericError
//
//	200: defaultvalueDBResponse
func (controller *Controller) DeleteDEFAULTVALUE(c *gin.Context) {

	mutexDEFAULTVALUE.Lock()
	defer mutexDEFAULTVALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDEFAULTVALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDEFAULTVALUE.GetDB()

	// Get model if exist
	var defaultvalueDB orm.DEFAULTVALUEDB
	if err := db.First(&defaultvalueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&defaultvalueDB)

	// get an instance (not staged) from DB instance, and call callback function
	defaultvalueDeleted := new(models.DEFAULTVALUE)
	defaultvalueDB.CopyBasicFieldsToDEFAULTVALUE(defaultvalueDeleted)

	// get stage instance from DB instance, and call callback function
	defaultvalueStaged := backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr[defaultvalueDB.ID]
	if defaultvalueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), defaultvalueStaged, defaultvalueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
