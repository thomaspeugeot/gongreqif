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
var __REQIFTYPE__dummysDeclaration__ models.REQIFTYPE
var __REQIFTYPE_time__dummyDeclaration time.Duration

var mutexREQIFTYPE sync.Mutex

// An REQIFTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQIFTYPE updateREQIFTYPE deleteREQIFTYPE
type REQIFTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQIFTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQIFTYPE updateREQIFTYPE
type REQIFTYPEInput struct {
	// The REQIFTYPE to submit or modify
	// in: body
	REQIFTYPE *orm.REQIFTYPEAPI
}

// GetREQIFTYPEs
//
// swagger:route GET /reqiftypes reqiftypes getREQIFTYPEs
//
// # Get all reqiftypes
//
// Responses:
// default: genericError
//
//	200: reqiftypeDBResponse
func (controller *Controller) GetREQIFTYPEs(c *gin.Context) {

	// source slice
	var reqiftypeDBs []orm.REQIFTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTYPE.GetDB()

	query := db.Find(&reqiftypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqiftypeAPIs := make([]orm.REQIFTYPEAPI, 0)

	// for each reqiftype, update fields from the database nullable fields
	for idx := range reqiftypeDBs {
		reqiftypeDB := &reqiftypeDBs[idx]
		_ = reqiftypeDB
		var reqiftypeAPI orm.REQIFTYPEAPI

		// insertion point for updating fields
		reqiftypeAPI.ID = reqiftypeDB.ID
		reqiftypeDB.CopyBasicFieldsToREQIFTYPE_WOP(&reqiftypeAPI.REQIFTYPE_WOP)
		reqiftypeAPI.REQIFTYPEPointersEncoding = reqiftypeDB.REQIFTYPEPointersEncoding
		reqiftypeAPIs = append(reqiftypeAPIs, reqiftypeAPI)
	}

	c.JSON(http.StatusOK, reqiftypeAPIs)
}

// PostREQIFTYPE
//
// swagger:route POST /reqiftypes reqiftypes postREQIFTYPE
//
// Creates a reqiftype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQIFTYPE(c *gin.Context) {

	mutexREQIFTYPE.Lock()
	defer mutexREQIFTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQIFTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTYPE.GetDB()

	// Validate input
	var input orm.REQIFTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqiftype
	reqiftypeDB := orm.REQIFTYPEDB{}
	reqiftypeDB.REQIFTYPEPointersEncoding = input.REQIFTYPEPointersEncoding
	reqiftypeDB.CopyBasicFieldsFromREQIFTYPE_WOP(&input.REQIFTYPE_WOP)

	query := db.Create(&reqiftypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQIFTYPE.CheckoutPhaseOneInstance(&reqiftypeDB)
	reqiftype := backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEDBID_REQIFTYPEPtr[reqiftypeDB.ID]

	if reqiftype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqiftype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqiftypeDB)
}

// GetREQIFTYPE
//
// swagger:route GET /reqiftypes/{ID} reqiftypes getREQIFTYPE
//
// Gets the details for a reqiftype.
//
// Responses:
// default: genericError
//
//	200: reqiftypeDBResponse
func (controller *Controller) GetREQIFTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTYPE.GetDB()

	// Get reqiftypeDB in DB
	var reqiftypeDB orm.REQIFTYPEDB
	if err := db.First(&reqiftypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqiftypeAPI orm.REQIFTYPEAPI
	reqiftypeAPI.ID = reqiftypeDB.ID
	reqiftypeAPI.REQIFTYPEPointersEncoding = reqiftypeDB.REQIFTYPEPointersEncoding
	reqiftypeDB.CopyBasicFieldsToREQIFTYPE_WOP(&reqiftypeAPI.REQIFTYPE_WOP)

	c.JSON(http.StatusOK, reqiftypeAPI)
}

// UpdateREQIFTYPE
//
// swagger:route PATCH /reqiftypes/{ID} reqiftypes updateREQIFTYPE
//
// # Update a reqiftype
//
// Responses:
// default: genericError
//
//	200: reqiftypeDBResponse
func (controller *Controller) UpdateREQIFTYPE(c *gin.Context) {

	mutexREQIFTYPE.Lock()
	defer mutexREQIFTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQIFTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTYPE.GetDB()

	// Validate input
	var input orm.REQIFTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqiftypeDB orm.REQIFTYPEDB

	// fetch the reqiftype
	query := db.First(&reqiftypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqiftypeDB.CopyBasicFieldsFromREQIFTYPE_WOP(&input.REQIFTYPE_WOP)
	reqiftypeDB.REQIFTYPEPointersEncoding = input.REQIFTYPEPointersEncoding

	query = db.Model(&reqiftypeDB).Updates(reqiftypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqiftypeNew := new(models.REQIFTYPE)
	reqiftypeDB.CopyBasicFieldsToREQIFTYPE(reqiftypeNew)

	// redeem pointers
	reqiftypeDB.DecodePointers(backRepo, reqiftypeNew)

	// get stage instance from DB instance, and call callback function
	reqiftypeOld := backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEDBID_REQIFTYPEPtr[reqiftypeDB.ID]
	if reqiftypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqiftypeOld, reqiftypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqiftypeDB
	c.JSON(http.StatusOK, reqiftypeDB)
}

// DeleteREQIFTYPE
//
// swagger:route DELETE /reqiftypes/{ID} reqiftypes deleteREQIFTYPE
//
// # Delete a reqiftype
//
// default: genericError
//
//	200: reqiftypeDBResponse
func (controller *Controller) DeleteREQIFTYPE(c *gin.Context) {

	mutexREQIFTYPE.Lock()
	defer mutexREQIFTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQIFTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTYPE.GetDB()

	// Get model if exist
	var reqiftypeDB orm.REQIFTYPEDB
	if err := db.First(&reqiftypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqiftypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqiftypeDeleted := new(models.REQIFTYPE)
	reqiftypeDB.CopyBasicFieldsToREQIFTYPE(reqiftypeDeleted)

	// get stage instance from DB instance, and call callback function
	reqiftypeStaged := backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEDBID_REQIFTYPEPtr[reqiftypeDB.ID]
	if reqiftypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqiftypeStaged, reqiftypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
