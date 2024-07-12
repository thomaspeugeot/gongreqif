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
var __REQIFTOOLEXTENSION__dummysDeclaration__ models.REQIFTOOLEXTENSION
var __REQIFTOOLEXTENSION_time__dummyDeclaration time.Duration

var mutexREQIFTOOLEXTENSION sync.Mutex

// An REQIFTOOLEXTENSIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQIFTOOLEXTENSION updateREQIFTOOLEXTENSION deleteREQIFTOOLEXTENSION
type REQIFTOOLEXTENSIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQIFTOOLEXTENSIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQIFTOOLEXTENSION updateREQIFTOOLEXTENSION
type REQIFTOOLEXTENSIONInput struct {
	// The REQIFTOOLEXTENSION to submit or modify
	// in: body
	REQIFTOOLEXTENSION *orm.REQIFTOOLEXTENSIONAPI
}

// GetREQIFTOOLEXTENSIONs
//
// swagger:route GET /reqiftoolextensions reqiftoolextensions getREQIFTOOLEXTENSIONs
//
// # Get all reqiftoolextensions
//
// Responses:
// default: genericError
//
//	200: reqiftoolextensionDBResponse
func (controller *Controller) GetREQIFTOOLEXTENSIONs(c *gin.Context) {

	// source slice
	var reqiftoolextensionDBs []orm.REQIFTOOLEXTENSIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFTOOLEXTENSIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTOOLEXTENSION.GetDB()

	query := db.Find(&reqiftoolextensionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqiftoolextensionAPIs := make([]orm.REQIFTOOLEXTENSIONAPI, 0)

	// for each reqiftoolextension, update fields from the database nullable fields
	for idx := range reqiftoolextensionDBs {
		reqiftoolextensionDB := &reqiftoolextensionDBs[idx]
		_ = reqiftoolextensionDB
		var reqiftoolextensionAPI orm.REQIFTOOLEXTENSIONAPI

		// insertion point for updating fields
		reqiftoolextensionAPI.ID = reqiftoolextensionDB.ID
		reqiftoolextensionDB.CopyBasicFieldsToREQIFTOOLEXTENSION_WOP(&reqiftoolextensionAPI.REQIFTOOLEXTENSION_WOP)
		reqiftoolextensionAPI.REQIFTOOLEXTENSIONPointersEncoding = reqiftoolextensionDB.REQIFTOOLEXTENSIONPointersEncoding
		reqiftoolextensionAPIs = append(reqiftoolextensionAPIs, reqiftoolextensionAPI)
	}

	c.JSON(http.StatusOK, reqiftoolextensionAPIs)
}

// PostREQIFTOOLEXTENSION
//
// swagger:route POST /reqiftoolextensions reqiftoolextensions postREQIFTOOLEXTENSION
//
// Creates a reqiftoolextension
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQIFTOOLEXTENSION(c *gin.Context) {

	mutexREQIFTOOLEXTENSION.Lock()
	defer mutexREQIFTOOLEXTENSION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQIFTOOLEXTENSIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTOOLEXTENSION.GetDB()

	// Validate input
	var input orm.REQIFTOOLEXTENSIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqiftoolextension
	reqiftoolextensionDB := orm.REQIFTOOLEXTENSIONDB{}
	reqiftoolextensionDB.REQIFTOOLEXTENSIONPointersEncoding = input.REQIFTOOLEXTENSIONPointersEncoding
	reqiftoolextensionDB.CopyBasicFieldsFromREQIFTOOLEXTENSION_WOP(&input.REQIFTOOLEXTENSION_WOP)

	query := db.Create(&reqiftoolextensionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQIFTOOLEXTENSION.CheckoutPhaseOneInstance(&reqiftoolextensionDB)
	reqiftoolextension := backRepo.BackRepoREQIFTOOLEXTENSION.Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONPtr[reqiftoolextensionDB.ID]

	if reqiftoolextension != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqiftoolextension)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqiftoolextensionDB)
}

// GetREQIFTOOLEXTENSION
//
// swagger:route GET /reqiftoolextensions/{ID} reqiftoolextensions getREQIFTOOLEXTENSION
//
// Gets the details for a reqiftoolextension.
//
// Responses:
// default: genericError
//
//	200: reqiftoolextensionDBResponse
func (controller *Controller) GetREQIFTOOLEXTENSION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFTOOLEXTENSION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTOOLEXTENSION.GetDB()

	// Get reqiftoolextensionDB in DB
	var reqiftoolextensionDB orm.REQIFTOOLEXTENSIONDB
	if err := db.First(&reqiftoolextensionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqiftoolextensionAPI orm.REQIFTOOLEXTENSIONAPI
	reqiftoolextensionAPI.ID = reqiftoolextensionDB.ID
	reqiftoolextensionAPI.REQIFTOOLEXTENSIONPointersEncoding = reqiftoolextensionDB.REQIFTOOLEXTENSIONPointersEncoding
	reqiftoolextensionDB.CopyBasicFieldsToREQIFTOOLEXTENSION_WOP(&reqiftoolextensionAPI.REQIFTOOLEXTENSION_WOP)

	c.JSON(http.StatusOK, reqiftoolextensionAPI)
}

// UpdateREQIFTOOLEXTENSION
//
// swagger:route PATCH /reqiftoolextensions/{ID} reqiftoolextensions updateREQIFTOOLEXTENSION
//
// # Update a reqiftoolextension
//
// Responses:
// default: genericError
//
//	200: reqiftoolextensionDBResponse
func (controller *Controller) UpdateREQIFTOOLEXTENSION(c *gin.Context) {

	mutexREQIFTOOLEXTENSION.Lock()
	defer mutexREQIFTOOLEXTENSION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQIFTOOLEXTENSION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTOOLEXTENSION.GetDB()

	// Validate input
	var input orm.REQIFTOOLEXTENSIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqiftoolextensionDB orm.REQIFTOOLEXTENSIONDB

	// fetch the reqiftoolextension
	query := db.First(&reqiftoolextensionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqiftoolextensionDB.CopyBasicFieldsFromREQIFTOOLEXTENSION_WOP(&input.REQIFTOOLEXTENSION_WOP)
	reqiftoolextensionDB.REQIFTOOLEXTENSIONPointersEncoding = input.REQIFTOOLEXTENSIONPointersEncoding

	query = db.Model(&reqiftoolextensionDB).Updates(reqiftoolextensionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqiftoolextensionNew := new(models.REQIFTOOLEXTENSION)
	reqiftoolextensionDB.CopyBasicFieldsToREQIFTOOLEXTENSION(reqiftoolextensionNew)

	// redeem pointers
	reqiftoolextensionDB.DecodePointers(backRepo, reqiftoolextensionNew)

	// get stage instance from DB instance, and call callback function
	reqiftoolextensionOld := backRepo.BackRepoREQIFTOOLEXTENSION.Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONPtr[reqiftoolextensionDB.ID]
	if reqiftoolextensionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqiftoolextensionOld, reqiftoolextensionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqiftoolextensionDB
	c.JSON(http.StatusOK, reqiftoolextensionDB)
}

// DeleteREQIFTOOLEXTENSION
//
// swagger:route DELETE /reqiftoolextensions/{ID} reqiftoolextensions deleteREQIFTOOLEXTENSION
//
// # Delete a reqiftoolextension
//
// default: genericError
//
//	200: reqiftoolextensionDBResponse
func (controller *Controller) DeleteREQIFTOOLEXTENSION(c *gin.Context) {

	mutexREQIFTOOLEXTENSION.Lock()
	defer mutexREQIFTOOLEXTENSION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQIFTOOLEXTENSION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFTOOLEXTENSION.GetDB()

	// Get model if exist
	var reqiftoolextensionDB orm.REQIFTOOLEXTENSIONDB
	if err := db.First(&reqiftoolextensionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqiftoolextensionDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqiftoolextensionDeleted := new(models.REQIFTOOLEXTENSION)
	reqiftoolextensionDB.CopyBasicFieldsToREQIFTOOLEXTENSION(reqiftoolextensionDeleted)

	// get stage instance from DB instance, and call callback function
	reqiftoolextensionStaged := backRepo.BackRepoREQIFTOOLEXTENSION.Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONPtr[reqiftoolextensionDB.ID]
	if reqiftoolextensionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqiftoolextensionStaged, reqiftoolextensionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
