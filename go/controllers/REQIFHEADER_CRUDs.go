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
var __REQIFHEADER__dummysDeclaration__ models.REQIFHEADER
var __REQIFHEADER_time__dummyDeclaration time.Duration

var mutexREQIFHEADER sync.Mutex

// An REQIFHEADERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQIFHEADER updateREQIFHEADER deleteREQIFHEADER
type REQIFHEADERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQIFHEADERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQIFHEADER updateREQIFHEADER
type REQIFHEADERInput struct {
	// The REQIFHEADER to submit or modify
	// in: body
	REQIFHEADER *orm.REQIFHEADERAPI
}

// GetREQIFHEADERs
//
// swagger:route GET /reqifheaders reqifheaders getREQIFHEADERs
//
// # Get all reqifheaders
//
// Responses:
// default: genericError
//
//	200: reqifheaderDBResponse
func (controller *Controller) GetREQIFHEADERs(c *gin.Context) {

	// source slice
	var reqifheaderDBs []orm.REQIFHEADERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFHEADER.GetDB()

	query := db.Find(&reqifheaderDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqifheaderAPIs := make([]orm.REQIFHEADERAPI, 0)

	// for each reqifheader, update fields from the database nullable fields
	for idx := range reqifheaderDBs {
		reqifheaderDB := &reqifheaderDBs[idx]
		_ = reqifheaderDB
		var reqifheaderAPI orm.REQIFHEADERAPI

		// insertion point for updating fields
		reqifheaderAPI.ID = reqifheaderDB.ID
		reqifheaderDB.CopyBasicFieldsToREQIFHEADER_WOP(&reqifheaderAPI.REQIFHEADER_WOP)
		reqifheaderAPI.REQIFHEADERPointersEncoding = reqifheaderDB.REQIFHEADERPointersEncoding
		reqifheaderAPIs = append(reqifheaderAPIs, reqifheaderAPI)
	}

	c.JSON(http.StatusOK, reqifheaderAPIs)
}

// PostREQIFHEADER
//
// swagger:route POST /reqifheaders reqifheaders postREQIFHEADER
//
// Creates a reqifheader
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQIFHEADER(c *gin.Context) {

	mutexREQIFHEADER.Lock()
	defer mutexREQIFHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQIFHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFHEADER.GetDB()

	// Validate input
	var input orm.REQIFHEADERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqifheader
	reqifheaderDB := orm.REQIFHEADERDB{}
	reqifheaderDB.REQIFHEADERPointersEncoding = input.REQIFHEADERPointersEncoding
	reqifheaderDB.CopyBasicFieldsFromREQIFHEADER_WOP(&input.REQIFHEADER_WOP)

	query := db.Create(&reqifheaderDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQIFHEADER.CheckoutPhaseOneInstance(&reqifheaderDB)
	reqifheader := backRepo.BackRepoREQIFHEADER.Map_REQIFHEADERDBID_REQIFHEADERPtr[reqifheaderDB.ID]

	if reqifheader != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqifheader)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqifheaderDB)
}

// GetREQIFHEADER
//
// swagger:route GET /reqifheaders/{ID} reqifheaders getREQIFHEADER
//
// Gets the details for a reqifheader.
//
// Responses:
// default: genericError
//
//	200: reqifheaderDBResponse
func (controller *Controller) GetREQIFHEADER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFHEADER.GetDB()

	// Get reqifheaderDB in DB
	var reqifheaderDB orm.REQIFHEADERDB
	if err := db.First(&reqifheaderDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqifheaderAPI orm.REQIFHEADERAPI
	reqifheaderAPI.ID = reqifheaderDB.ID
	reqifheaderAPI.REQIFHEADERPointersEncoding = reqifheaderDB.REQIFHEADERPointersEncoding
	reqifheaderDB.CopyBasicFieldsToREQIFHEADER_WOP(&reqifheaderAPI.REQIFHEADER_WOP)

	c.JSON(http.StatusOK, reqifheaderAPI)
}

// UpdateREQIFHEADER
//
// swagger:route PATCH /reqifheaders/{ID} reqifheaders updateREQIFHEADER
//
// # Update a reqifheader
//
// Responses:
// default: genericError
//
//	200: reqifheaderDBResponse
func (controller *Controller) UpdateREQIFHEADER(c *gin.Context) {

	mutexREQIFHEADER.Lock()
	defer mutexREQIFHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQIFHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFHEADER.GetDB()

	// Validate input
	var input orm.REQIFHEADERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqifheaderDB orm.REQIFHEADERDB

	// fetch the reqifheader
	query := db.First(&reqifheaderDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqifheaderDB.CopyBasicFieldsFromREQIFHEADER_WOP(&input.REQIFHEADER_WOP)
	reqifheaderDB.REQIFHEADERPointersEncoding = input.REQIFHEADERPointersEncoding

	query = db.Model(&reqifheaderDB).Updates(reqifheaderDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqifheaderNew := new(models.REQIFHEADER)
	reqifheaderDB.CopyBasicFieldsToREQIFHEADER(reqifheaderNew)

	// redeem pointers
	reqifheaderDB.DecodePointers(backRepo, reqifheaderNew)

	// get stage instance from DB instance, and call callback function
	reqifheaderOld := backRepo.BackRepoREQIFHEADER.Map_REQIFHEADERDBID_REQIFHEADERPtr[reqifheaderDB.ID]
	if reqifheaderOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqifheaderOld, reqifheaderNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqifheaderDB
	c.JSON(http.StatusOK, reqifheaderDB)
}

// DeleteREQIFHEADER
//
// swagger:route DELETE /reqifheaders/{ID} reqifheaders deleteREQIFHEADER
//
// # Delete a reqifheader
//
// default: genericError
//
//	200: reqifheaderDBResponse
func (controller *Controller) DeleteREQIFHEADER(c *gin.Context) {

	mutexREQIFHEADER.Lock()
	defer mutexREQIFHEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQIFHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFHEADER.GetDB()

	// Get model if exist
	var reqifheaderDB orm.REQIFHEADERDB
	if err := db.First(&reqifheaderDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqifheaderDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqifheaderDeleted := new(models.REQIFHEADER)
	reqifheaderDB.CopyBasicFieldsToREQIFHEADER(reqifheaderDeleted)

	// get stage instance from DB instance, and call callback function
	reqifheaderStaged := backRepo.BackRepoREQIFHEADER.Map_REQIFHEADERDBID_REQIFHEADERPtr[reqifheaderDB.ID]
	if reqifheaderStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqifheaderStaged, reqifheaderDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
