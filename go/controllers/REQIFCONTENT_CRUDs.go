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
var __REQIFCONTENT__dummysDeclaration__ models.REQIFCONTENT
var __REQIFCONTENT_time__dummyDeclaration time.Duration

var mutexREQIFCONTENT sync.Mutex

// An REQIFCONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQIFCONTENT updateREQIFCONTENT deleteREQIFCONTENT
type REQIFCONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQIFCONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQIFCONTENT updateREQIFCONTENT
type REQIFCONTENTInput struct {
	// The REQIFCONTENT to submit or modify
	// in: body
	REQIFCONTENT *orm.REQIFCONTENTAPI
}

// GetREQIFCONTENTs
//
// swagger:route GET /reqifcontents reqifcontents getREQIFCONTENTs
//
// # Get all reqifcontents
//
// Responses:
// default: genericError
//
//	200: reqifcontentDBResponse
func (controller *Controller) GetREQIFCONTENTs(c *gin.Context) {

	// source slice
	var reqifcontentDBs []orm.REQIFCONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFCONTENT.GetDB()

	query := db.Find(&reqifcontentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqifcontentAPIs := make([]orm.REQIFCONTENTAPI, 0)

	// for each reqifcontent, update fields from the database nullable fields
	for idx := range reqifcontentDBs {
		reqifcontentDB := &reqifcontentDBs[idx]
		_ = reqifcontentDB
		var reqifcontentAPI orm.REQIFCONTENTAPI

		// insertion point for updating fields
		reqifcontentAPI.ID = reqifcontentDB.ID
		reqifcontentDB.CopyBasicFieldsToREQIFCONTENT_WOP(&reqifcontentAPI.REQIFCONTENT_WOP)
		reqifcontentAPI.REQIFCONTENTPointersEncoding = reqifcontentDB.REQIFCONTENTPointersEncoding
		reqifcontentAPIs = append(reqifcontentAPIs, reqifcontentAPI)
	}

	c.JSON(http.StatusOK, reqifcontentAPIs)
}

// PostREQIFCONTENT
//
// swagger:route POST /reqifcontents reqifcontents postREQIFCONTENT
//
// Creates a reqifcontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQIFCONTENT(c *gin.Context) {

	mutexREQIFCONTENT.Lock()
	defer mutexREQIFCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQIFCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFCONTENT.GetDB()

	// Validate input
	var input orm.REQIFCONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqifcontent
	reqifcontentDB := orm.REQIFCONTENTDB{}
	reqifcontentDB.REQIFCONTENTPointersEncoding = input.REQIFCONTENTPointersEncoding
	reqifcontentDB.CopyBasicFieldsFromREQIFCONTENT_WOP(&input.REQIFCONTENT_WOP)

	query := db.Create(&reqifcontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQIFCONTENT.CheckoutPhaseOneInstance(&reqifcontentDB)
	reqifcontent := backRepo.BackRepoREQIFCONTENT.Map_REQIFCONTENTDBID_REQIFCONTENTPtr[reqifcontentDB.ID]

	if reqifcontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqifcontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqifcontentDB)
}

// GetREQIFCONTENT
//
// swagger:route GET /reqifcontents/{ID} reqifcontents getREQIFCONTENT
//
// Gets the details for a reqifcontent.
//
// Responses:
// default: genericError
//
//	200: reqifcontentDBResponse
func (controller *Controller) GetREQIFCONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFCONTENT.GetDB()

	// Get reqifcontentDB in DB
	var reqifcontentDB orm.REQIFCONTENTDB
	if err := db.First(&reqifcontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqifcontentAPI orm.REQIFCONTENTAPI
	reqifcontentAPI.ID = reqifcontentDB.ID
	reqifcontentAPI.REQIFCONTENTPointersEncoding = reqifcontentDB.REQIFCONTENTPointersEncoding
	reqifcontentDB.CopyBasicFieldsToREQIFCONTENT_WOP(&reqifcontentAPI.REQIFCONTENT_WOP)

	c.JSON(http.StatusOK, reqifcontentAPI)
}

// UpdateREQIFCONTENT
//
// swagger:route PATCH /reqifcontents/{ID} reqifcontents updateREQIFCONTENT
//
// # Update a reqifcontent
//
// Responses:
// default: genericError
//
//	200: reqifcontentDBResponse
func (controller *Controller) UpdateREQIFCONTENT(c *gin.Context) {

	mutexREQIFCONTENT.Lock()
	defer mutexREQIFCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQIFCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFCONTENT.GetDB()

	// Validate input
	var input orm.REQIFCONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqifcontentDB orm.REQIFCONTENTDB

	// fetch the reqifcontent
	query := db.First(&reqifcontentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqifcontentDB.CopyBasicFieldsFromREQIFCONTENT_WOP(&input.REQIFCONTENT_WOP)
	reqifcontentDB.REQIFCONTENTPointersEncoding = input.REQIFCONTENTPointersEncoding

	query = db.Model(&reqifcontentDB).Updates(reqifcontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqifcontentNew := new(models.REQIFCONTENT)
	reqifcontentDB.CopyBasicFieldsToREQIFCONTENT(reqifcontentNew)

	// redeem pointers
	reqifcontentDB.DecodePointers(backRepo, reqifcontentNew)

	// get stage instance from DB instance, and call callback function
	reqifcontentOld := backRepo.BackRepoREQIFCONTENT.Map_REQIFCONTENTDBID_REQIFCONTENTPtr[reqifcontentDB.ID]
	if reqifcontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqifcontentOld, reqifcontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqifcontentDB
	c.JSON(http.StatusOK, reqifcontentDB)
}

// DeleteREQIFCONTENT
//
// swagger:route DELETE /reqifcontents/{ID} reqifcontents deleteREQIFCONTENT
//
// # Delete a reqifcontent
//
// default: genericError
//
//	200: reqifcontentDBResponse
func (controller *Controller) DeleteREQIFCONTENT(c *gin.Context) {

	mutexREQIFCONTENT.Lock()
	defer mutexREQIFCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQIFCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIFCONTENT.GetDB()

	// Get model if exist
	var reqifcontentDB orm.REQIFCONTENTDB
	if err := db.First(&reqifcontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqifcontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqifcontentDeleted := new(models.REQIFCONTENT)
	reqifcontentDB.CopyBasicFieldsToREQIFCONTENT(reqifcontentDeleted)

	// get stage instance from DB instance, and call callback function
	reqifcontentStaged := backRepo.BackRepoREQIFCONTENT.Map_REQIFCONTENTDBID_REQIFCONTENTPtr[reqifcontentDB.ID]
	if reqifcontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqifcontentStaged, reqifcontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
