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
var __REQIF__dummysDeclaration__ models.REQIF
var __REQIF_time__dummyDeclaration time.Duration

var mutexREQIF sync.Mutex

// An REQIFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQIF updateREQIF deleteREQIF
type REQIFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQIFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQIF updateREQIF
type REQIFInput struct {
	// The REQIF to submit or modify
	// in: body
	REQIF *orm.REQIFAPI
}

// GetREQIFs
//
// swagger:route GET /reqifs reqifs getREQIFs
//
// # Get all reqifs
//
// Responses:
// default: genericError
//
//	200: reqifDBResponse
func (controller *Controller) GetREQIFs(c *gin.Context) {

	// source slice
	var reqifDBs []orm.REQIFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIFs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIF.GetDB()

	query := db.Find(&reqifDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	reqifAPIs := make([]orm.REQIFAPI, 0)

	// for each reqif, update fields from the database nullable fields
	for idx := range reqifDBs {
		reqifDB := &reqifDBs[idx]
		_ = reqifDB
		var reqifAPI orm.REQIFAPI

		// insertion point for updating fields
		reqifAPI.ID = reqifDB.ID
		reqifDB.CopyBasicFieldsToREQIF_WOP(&reqifAPI.REQIF_WOP)
		reqifAPI.REQIFPointersEncoding = reqifDB.REQIFPointersEncoding
		reqifAPIs = append(reqifAPIs, reqifAPI)
	}

	c.JSON(http.StatusOK, reqifAPIs)
}

// PostREQIF
//
// swagger:route POST /reqifs reqifs postREQIF
//
// Creates a reqif
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQIF(c *gin.Context) {

	mutexREQIF.Lock()
	defer mutexREQIF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQIFs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIF.GetDB()

	// Validate input
	var input orm.REQIFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create reqif
	reqifDB := orm.REQIFDB{}
	reqifDB.REQIFPointersEncoding = input.REQIFPointersEncoding
	reqifDB.CopyBasicFieldsFromREQIF_WOP(&input.REQIF_WOP)

	query := db.Create(&reqifDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQIF.CheckoutPhaseOneInstance(&reqifDB)
	reqif := backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]

	if reqif != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), reqif)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, reqifDB)
}

// GetREQIF
//
// swagger:route GET /reqifs/{ID} reqifs getREQIF
//
// Gets the details for a reqif.
//
// Responses:
// default: genericError
//
//	200: reqifDBResponse
func (controller *Controller) GetREQIF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQIF", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIF.GetDB()

	// Get reqifDB in DB
	var reqifDB orm.REQIFDB
	if err := db.First(&reqifDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var reqifAPI orm.REQIFAPI
	reqifAPI.ID = reqifDB.ID
	reqifAPI.REQIFPointersEncoding = reqifDB.REQIFPointersEncoding
	reqifDB.CopyBasicFieldsToREQIF_WOP(&reqifAPI.REQIF_WOP)

	c.JSON(http.StatusOK, reqifAPI)
}

// UpdateREQIF
//
// swagger:route PATCH /reqifs/{ID} reqifs updateREQIF
//
// # Update a reqif
//
// Responses:
// default: genericError
//
//	200: reqifDBResponse
func (controller *Controller) UpdateREQIF(c *gin.Context) {

	mutexREQIF.Lock()
	defer mutexREQIF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQIF", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIF.GetDB()

	// Validate input
	var input orm.REQIFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var reqifDB orm.REQIFDB

	// fetch the reqif
	query := db.First(&reqifDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	reqifDB.CopyBasicFieldsFromREQIF_WOP(&input.REQIF_WOP)
	reqifDB.REQIFPointersEncoding = input.REQIFPointersEncoding

	query = db.Model(&reqifDB).Updates(reqifDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	reqifNew := new(models.REQIF)
	reqifDB.CopyBasicFieldsToREQIF(reqifNew)

	// redeem pointers
	reqifDB.DecodePointers(backRepo, reqifNew)

	// get stage instance from DB instance, and call callback function
	reqifOld := backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]
	if reqifOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), reqifOld, reqifNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the reqifDB
	c.JSON(http.StatusOK, reqifDB)
}

// DeleteREQIF
//
// swagger:route DELETE /reqifs/{ID} reqifs deleteREQIF
//
// # Delete a reqif
//
// default: genericError
//
//	200: reqifDBResponse
func (controller *Controller) DeleteREQIF(c *gin.Context) {

	mutexREQIF.Lock()
	defer mutexREQIF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQIF", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoREQIF.GetDB()

	// Get model if exist
	var reqifDB orm.REQIFDB
	if err := db.First(&reqifDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&reqifDB)

	// get an instance (not staged) from DB instance, and call callback function
	reqifDeleted := new(models.REQIF)
	reqifDB.CopyBasicFieldsToREQIF(reqifDeleted)

	// get stage instance from DB instance, and call callback function
	reqifStaged := backRepo.BackRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]
	if reqifStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), reqifStaged, reqifDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
