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
var __ALTERNATIVEID__dummysDeclaration__ models.ALTERNATIVEID
var __ALTERNATIVEID_time__dummyDeclaration time.Duration

var mutexALTERNATIVEID sync.Mutex

// An ALTERNATIVEIDID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getALTERNATIVEID updateALTERNATIVEID deleteALTERNATIVEID
type ALTERNATIVEIDID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ALTERNATIVEIDInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postALTERNATIVEID updateALTERNATIVEID
type ALTERNATIVEIDInput struct {
	// The ALTERNATIVEID to submit or modify
	// in: body
	ALTERNATIVEID *orm.ALTERNATIVEIDAPI
}

// GetALTERNATIVEIDs
//
// swagger:route GET /alternativeids alternativeids getALTERNATIVEIDs
//
// # Get all alternativeids
//
// Responses:
// default: genericError
//
//	200: alternativeidDBResponse
func (controller *Controller) GetALTERNATIVEIDs(c *gin.Context) {

	// source slice
	var alternativeidDBs []orm.ALTERNATIVEIDDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetALTERNATIVEIDs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVEID.GetDB()

	query := db.Find(&alternativeidDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	alternativeidAPIs := make([]orm.ALTERNATIVEIDAPI, 0)

	// for each alternativeid, update fields from the database nullable fields
	for idx := range alternativeidDBs {
		alternativeidDB := &alternativeidDBs[idx]
		_ = alternativeidDB
		var alternativeidAPI orm.ALTERNATIVEIDAPI

		// insertion point for updating fields
		alternativeidAPI.ID = alternativeidDB.ID
		alternativeidDB.CopyBasicFieldsToALTERNATIVEID_WOP(&alternativeidAPI.ALTERNATIVEID_WOP)
		alternativeidAPI.ALTERNATIVEIDPointersEncoding = alternativeidDB.ALTERNATIVEIDPointersEncoding
		alternativeidAPIs = append(alternativeidAPIs, alternativeidAPI)
	}

	c.JSON(http.StatusOK, alternativeidAPIs)
}

// PostALTERNATIVEID
//
// swagger:route POST /alternativeids alternativeids postALTERNATIVEID
//
// Creates a alternativeid
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostALTERNATIVEID(c *gin.Context) {

	mutexALTERNATIVEID.Lock()
	defer mutexALTERNATIVEID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostALTERNATIVEIDs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVEID.GetDB()

	// Validate input
	var input orm.ALTERNATIVEIDAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create alternativeid
	alternativeidDB := orm.ALTERNATIVEIDDB{}
	alternativeidDB.ALTERNATIVEIDPointersEncoding = input.ALTERNATIVEIDPointersEncoding
	alternativeidDB.CopyBasicFieldsFromALTERNATIVEID_WOP(&input.ALTERNATIVEID_WOP)

	query := db.Create(&alternativeidDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoALTERNATIVEID.CheckoutPhaseOneInstance(&alternativeidDB)
	alternativeid := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[alternativeidDB.ID]

	if alternativeid != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), alternativeid)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, alternativeidDB)
}

// GetALTERNATIVEID
//
// swagger:route GET /alternativeids/{ID} alternativeids getALTERNATIVEID
//
// Gets the details for a alternativeid.
//
// Responses:
// default: genericError
//
//	200: alternativeidDBResponse
func (controller *Controller) GetALTERNATIVEID(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetALTERNATIVEID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVEID.GetDB()

	// Get alternativeidDB in DB
	var alternativeidDB orm.ALTERNATIVEIDDB
	if err := db.First(&alternativeidDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var alternativeidAPI orm.ALTERNATIVEIDAPI
	alternativeidAPI.ID = alternativeidDB.ID
	alternativeidAPI.ALTERNATIVEIDPointersEncoding = alternativeidDB.ALTERNATIVEIDPointersEncoding
	alternativeidDB.CopyBasicFieldsToALTERNATIVEID_WOP(&alternativeidAPI.ALTERNATIVEID_WOP)

	c.JSON(http.StatusOK, alternativeidAPI)
}

// UpdateALTERNATIVEID
//
// swagger:route PATCH /alternativeids/{ID} alternativeids updateALTERNATIVEID
//
// # Update a alternativeid
//
// Responses:
// default: genericError
//
//	200: alternativeidDBResponse
func (controller *Controller) UpdateALTERNATIVEID(c *gin.Context) {

	mutexALTERNATIVEID.Lock()
	defer mutexALTERNATIVEID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateALTERNATIVEID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVEID.GetDB()

	// Validate input
	var input orm.ALTERNATIVEIDAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var alternativeidDB orm.ALTERNATIVEIDDB

	// fetch the alternativeid
	query := db.First(&alternativeidDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	alternativeidDB.CopyBasicFieldsFromALTERNATIVEID_WOP(&input.ALTERNATIVEID_WOP)
	alternativeidDB.ALTERNATIVEIDPointersEncoding = input.ALTERNATIVEIDPointersEncoding

	query = db.Model(&alternativeidDB).Updates(alternativeidDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	alternativeidNew := new(models.ALTERNATIVEID)
	alternativeidDB.CopyBasicFieldsToALTERNATIVEID(alternativeidNew)

	// redeem pointers
	alternativeidDB.DecodePointers(backRepo, alternativeidNew)

	// get stage instance from DB instance, and call callback function
	alternativeidOld := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[alternativeidDB.ID]
	if alternativeidOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), alternativeidOld, alternativeidNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the alternativeidDB
	c.JSON(http.StatusOK, alternativeidDB)
}

// DeleteALTERNATIVEID
//
// swagger:route DELETE /alternativeids/{ID} alternativeids deleteALTERNATIVEID
//
// # Delete a alternativeid
//
// default: genericError
//
//	200: alternativeidDBResponse
func (controller *Controller) DeleteALTERNATIVEID(c *gin.Context) {

	mutexALTERNATIVEID.Lock()
	defer mutexALTERNATIVEID.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteALTERNATIVEID", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoALTERNATIVEID.GetDB()

	// Get model if exist
	var alternativeidDB orm.ALTERNATIVEIDDB
	if err := db.First(&alternativeidDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&alternativeidDB)

	// get an instance (not staged) from DB instance, and call callback function
	alternativeidDeleted := new(models.ALTERNATIVEID)
	alternativeidDB.CopyBasicFieldsToALTERNATIVEID(alternativeidDeleted)

	// get stage instance from DB instance, and call callback function
	alternativeidStaged := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[alternativeidDB.ID]
	if alternativeidStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), alternativeidStaged, alternativeidDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
