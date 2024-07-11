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
var __HEADER__dummysDeclaration__ models.HEADER
var __HEADER_time__dummyDeclaration time.Duration

var mutexHEADER sync.Mutex

// An HEADERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHEADER updateHEADER deleteHEADER
type HEADERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HEADERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHEADER updateHEADER
type HEADERInput struct {
	// The HEADER to submit or modify
	// in: body
	HEADER *orm.HEADERAPI
}

// GetHEADERs
//
// swagger:route GET /headers headers getHEADERs
//
// # Get all headers
//
// Responses:
// default: genericError
//
//	200: headerDBResponse
func (controller *Controller) GetHEADERs(c *gin.Context) {

	// source slice
	var headerDBs []orm.HEADERDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHEADER.GetDB()

	query := db.Find(&headerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	headerAPIs := make([]orm.HEADERAPI, 0)

	// for each header, update fields from the database nullable fields
	for idx := range headerDBs {
		headerDB := &headerDBs[idx]
		_ = headerDB
		var headerAPI orm.HEADERAPI

		// insertion point for updating fields
		headerAPI.ID = headerDB.ID
		headerDB.CopyBasicFieldsToHEADER_WOP(&headerAPI.HEADER_WOP)
		headerAPI.HEADERPointersEncoding = headerDB.HEADERPointersEncoding
		headerAPIs = append(headerAPIs, headerAPI)
	}

	c.JSON(http.StatusOK, headerAPIs)
}

// PostHEADER
//
// swagger:route POST /headers headers postHEADER
//
// Creates a header
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHEADER(c *gin.Context) {

	mutexHEADER.Lock()
	defer mutexHEADER.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHEADERs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHEADER.GetDB()

	// Validate input
	var input orm.HEADERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create header
	headerDB := orm.HEADERDB{}
	headerDB.HEADERPointersEncoding = input.HEADERPointersEncoding
	headerDB.CopyBasicFieldsFromHEADER_WOP(&input.HEADER_WOP)

	query := db.Create(&headerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHEADER.CheckoutPhaseOneInstance(&headerDB)
	header := backRepo.BackRepoHEADER.Map_HEADERDBID_HEADERPtr[headerDB.ID]

	if header != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), header)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, headerDB)
}

// GetHEADER
//
// swagger:route GET /headers/{ID} headers getHEADER
//
// Gets the details for a header.
//
// Responses:
// default: genericError
//
//	200: headerDBResponse
func (controller *Controller) GetHEADER(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHEADER.GetDB()

	// Get headerDB in DB
	var headerDB orm.HEADERDB
	if err := db.First(&headerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var headerAPI orm.HEADERAPI
	headerAPI.ID = headerDB.ID
	headerAPI.HEADERPointersEncoding = headerDB.HEADERPointersEncoding
	headerDB.CopyBasicFieldsToHEADER_WOP(&headerAPI.HEADER_WOP)

	c.JSON(http.StatusOK, headerAPI)
}

// UpdateHEADER
//
// swagger:route PATCH /headers/{ID} headers updateHEADER
//
// # Update a header
//
// Responses:
// default: genericError
//
//	200: headerDBResponse
func (controller *Controller) UpdateHEADER(c *gin.Context) {

	mutexHEADER.Lock()
	defer mutexHEADER.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHEADER.GetDB()

	// Validate input
	var input orm.HEADERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var headerDB orm.HEADERDB

	// fetch the header
	query := db.First(&headerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	headerDB.CopyBasicFieldsFromHEADER_WOP(&input.HEADER_WOP)
	headerDB.HEADERPointersEncoding = input.HEADERPointersEncoding

	query = db.Model(&headerDB).Updates(headerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	headerNew := new(models.HEADER)
	headerDB.CopyBasicFieldsToHEADER(headerNew)

	// redeem pointers
	headerDB.DecodePointers(backRepo, headerNew)

	// get stage instance from DB instance, and call callback function
	headerOld := backRepo.BackRepoHEADER.Map_HEADERDBID_HEADERPtr[headerDB.ID]
	if headerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), headerOld, headerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the headerDB
	c.JSON(http.StatusOK, headerDB)
}

// DeleteHEADER
//
// swagger:route DELETE /headers/{ID} headers deleteHEADER
//
// # Delete a header
//
// default: genericError
//
//	200: headerDBResponse
func (controller *Controller) DeleteHEADER(c *gin.Context) {

	mutexHEADER.Lock()
	defer mutexHEADER.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHEADER", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHEADER.GetDB()

	// Get model if exist
	var headerDB orm.HEADERDB
	if err := db.First(&headerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&headerDB)

	// get an instance (not staged) from DB instance, and call callback function
	headerDeleted := new(models.HEADER)
	headerDB.CopyBasicFieldsToHEADER(headerDeleted)

	// get stage instance from DB instance, and call callback function
	headerStaged := backRepo.BackRepoHEADER.Map_HEADERDBID_HEADERPtr[headerDB.ID]
	if headerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), headerStaged, headerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
