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
var __REQTYPE__dummysDeclaration__ models.REQTYPE
var __REQTYPE_time__dummyDeclaration time.Duration

var mutexREQTYPE sync.Mutex

// An REQTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQTYPE updateREQTYPE deleteREQTYPE
type REQTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQTYPE updateREQTYPE
type REQTYPEInput struct {
	// The REQTYPE to submit or modify
	// in: body
	REQTYPE *orm.REQTYPEAPI
}

// GetREQTYPEs
//
// swagger:route GET /reqtypes reqtypes getREQTYPEs
//
// # Get all reqtypes
//
// Responses:
// default: genericError
//
//	200: reqtypeDBResponse
func (controller *Controller) GetREQTYPEs(c *gin.Context) {

	// source slice
	var reqtypeDBs []orm.REQTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQTYPE.GetDB()

	query := db.Find(&reqtypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqtypeAPIs := make([]orm.REQTYPEAPI, 0)

	// for each reqtype, update fields from the database nullable fields
	for idx := range reqtypeDBs {
		reqtypeDB := &reqtypeDBs[idx]
		_ = reqtypeDB
		var reqtypeAPI orm.REQTYPEAPI

		// insertion point for updating fields
		reqtypeAPI.ID = reqtypeDB.ID
		reqtypeDB.CopyBasicFieldsToREQTYPE_WOP(&reqtypeAPI.REQTYPE_WOP)
		reqtypeAPI.REQTYPEPointersEncoding = reqtypeDB.REQTYPEPointersEncoding
		reqtypeAPIs = append(reqtypeAPIs, reqtypeAPI)
	}

	c.JSON(http.StatusOK, reqtypeAPIs)
}

// PostREQTYPE
//
// swagger:route POST /reqtypes reqtypes postREQTYPE
//
// Creates a reqtype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQTYPE(c *gin.Context) {

	mutexREQTYPE.Lock()
	defer mutexREQTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQTYPE.GetDB()

	// Validate input
	var input orm.REQTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqtype
	reqtypeDB := orm.REQTYPEDB{}
	reqtypeDB.REQTYPEPointersEncoding = input.REQTYPEPointersEncoding
	reqtypeDB.CopyBasicFieldsFromREQTYPE_WOP(&input.REQTYPE_WOP)

	query := db.Create(&reqtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQTYPE.CheckoutPhaseOneInstance(&reqtypeDB)
	reqtype := backRepo.BackRepoREQTYPE.Map_REQTYPEDBID_REQTYPEPtr[reqtypeDB.ID]

	if reqtype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqtype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqtypeDB)
}

// GetREQTYPE
//
// swagger:route GET /reqtypes/{ID} reqtypes getREQTYPE
//
// Gets the details for a reqtype.
//
// Responses:
// default: genericError
//
//	200: reqtypeDBResponse
func (controller *Controller) GetREQTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQTYPE.GetDB()

	// Get reqtypeDB in DB
	var reqtypeDB orm.REQTYPEDB
	if err := db.First(&reqtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqtypeAPI orm.REQTYPEAPI
	reqtypeAPI.ID = reqtypeDB.ID
	reqtypeAPI.REQTYPEPointersEncoding = reqtypeDB.REQTYPEPointersEncoding
	reqtypeDB.CopyBasicFieldsToREQTYPE_WOP(&reqtypeAPI.REQTYPE_WOP)

	c.JSON(http.StatusOK, reqtypeAPI)
}

// UpdateREQTYPE
//
// swagger:route PATCH /reqtypes/{ID} reqtypes updateREQTYPE
//
// # Update a reqtype
//
// Responses:
// default: genericError
//
//	200: reqtypeDBResponse
func (controller *Controller) UpdateREQTYPE(c *gin.Context) {

	mutexREQTYPE.Lock()
	defer mutexREQTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQTYPE.GetDB()

	// Validate input
	var input orm.REQTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqtypeDB orm.REQTYPEDB

	// fetch the reqtype
	query := db.First(&reqtypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqtypeDB.CopyBasicFieldsFromREQTYPE_WOP(&input.REQTYPE_WOP)
	reqtypeDB.REQTYPEPointersEncoding = input.REQTYPEPointersEncoding

	query = db.Model(&reqtypeDB).Updates(reqtypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqtypeNew := new(models.REQTYPE)
	reqtypeDB.CopyBasicFieldsToREQTYPE(reqtypeNew)

	// redeem pointers
	reqtypeDB.DecodePointers(backRepo, reqtypeNew)

	// get stage instance from DB instance, and call callback function
	reqtypeOld := backRepo.BackRepoREQTYPE.Map_REQTYPEDBID_REQTYPEPtr[reqtypeDB.ID]
	if reqtypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqtypeOld, reqtypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqtypeDB
	c.JSON(http.StatusOK, reqtypeDB)
}

// DeleteREQTYPE
//
// swagger:route DELETE /reqtypes/{ID} reqtypes deleteREQTYPE
//
// # Delete a reqtype
//
// default: genericError
//
//	200: reqtypeDBResponse
func (controller *Controller) DeleteREQTYPE(c *gin.Context) {

	mutexREQTYPE.Lock()
	defer mutexREQTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQTYPE.GetDB()

	// Get model if exist
	var reqtypeDB orm.REQTYPEDB
	if err := db.First(&reqtypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqtypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqtypeDeleted := new(models.REQTYPE)
	reqtypeDB.CopyBasicFieldsToREQTYPE(reqtypeDeleted)

	// get stage instance from DB instance, and call callback function
	reqtypeStaged := backRepo.BackRepoREQTYPE.Map_REQTYPEDBID_REQTYPEPtr[reqtypeDB.ID]
	if reqtypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqtypeStaged, reqtypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
