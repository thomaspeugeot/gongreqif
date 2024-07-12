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
var __DATATYPEDEFINITIONDATE__dummysDeclaration__ models.DATATYPEDEFINITIONDATE
var __DATATYPEDEFINITIONDATE_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONDATE sync.Mutex

// An DATATYPEDEFINITIONDATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONDATE updateDATATYPEDEFINITIONDATE deleteDATATYPEDEFINITIONDATE
type DATATYPEDEFINITIONDATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONDATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONDATE updateDATATYPEDEFINITIONDATE
type DATATYPEDEFINITIONDATEInput struct {
	// The DATATYPEDEFINITIONDATE to submit or modify
	// in: body
	DATATYPEDEFINITIONDATE *orm.DATATYPEDEFINITIONDATEAPI
}

// GetDATATYPEDEFINITIONDATEs
//
// swagger:route GET /datatypedefinitiondates datatypedefinitiondates getDATATYPEDEFINITIONDATEs
//
// # Get all datatypedefinitiondates
//
// Responses:
// default: genericError
//
//	200: datatypedefinitiondateDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONDATEs(c *gin.Context) {

	// source slice
	var datatypedefinitiondateDBs []orm.DATATYPEDEFINITIONDATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDB()

	query := db.Find(&datatypedefinitiondateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitiondateAPIs := make([]orm.DATATYPEDEFINITIONDATEAPI, 0)

	// for each datatypedefinitiondate, update fields from the database nullable fields
	for idx := range datatypedefinitiondateDBs {
		datatypedefinitiondateDB := &datatypedefinitiondateDBs[idx]
		_ = datatypedefinitiondateDB
		var datatypedefinitiondateAPI orm.DATATYPEDEFINITIONDATEAPI

		// insertion point for updating fields
		datatypedefinitiondateAPI.ID = datatypedefinitiondateDB.ID
		datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE_WOP(&datatypedefinitiondateAPI.DATATYPEDEFINITIONDATE_WOP)
		datatypedefinitiondateAPI.DATATYPEDEFINITIONDATEPointersEncoding = datatypedefinitiondateDB.DATATYPEDEFINITIONDATEPointersEncoding
		datatypedefinitiondateAPIs = append(datatypedefinitiondateAPIs, datatypedefinitiondateAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitiondateAPIs)
}

// PostDATATYPEDEFINITIONDATE
//
// swagger:route POST /datatypedefinitiondates datatypedefinitiondates postDATATYPEDEFINITIONDATE
//
// Creates a datatypedefinitiondate
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONDATE(c *gin.Context) {

	mutexDATATYPEDEFINITIONDATE.Lock()
	defer mutexDATATYPEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONDATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONDATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitiondate
	datatypedefinitiondateDB := orm.DATATYPEDEFINITIONDATEDB{}
	datatypedefinitiondateDB.DATATYPEDEFINITIONDATEPointersEncoding = input.DATATYPEDEFINITIONDATEPointersEncoding
	datatypedefinitiondateDB.CopyBasicFieldsFromDATATYPEDEFINITIONDATE_WOP(&input.DATATYPEDEFINITIONDATE_WOP)

	query := db.Create(&datatypedefinitiondateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseOneInstance(&datatypedefinitiondateDB)
	datatypedefinitiondate := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]

	if datatypedefinitiondate != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitiondate)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitiondateDB)
}

// GetDATATYPEDEFINITIONDATE
//
// swagger:route GET /datatypedefinitiondates/{ID} datatypedefinitiondates getDATATYPEDEFINITIONDATE
//
// Gets the details for a datatypedefinitiondate.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitiondateDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONDATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDB()

	// Get datatypedefinitiondateDB in DB
	var datatypedefinitiondateDB orm.DATATYPEDEFINITIONDATEDB
	if err := db.First(&datatypedefinitiondateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitiondateAPI orm.DATATYPEDEFINITIONDATEAPI
	datatypedefinitiondateAPI.ID = datatypedefinitiondateDB.ID
	datatypedefinitiondateAPI.DATATYPEDEFINITIONDATEPointersEncoding = datatypedefinitiondateDB.DATATYPEDEFINITIONDATEPointersEncoding
	datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE_WOP(&datatypedefinitiondateAPI.DATATYPEDEFINITIONDATE_WOP)

	c.JSON(http.StatusOK, datatypedefinitiondateAPI)
}

// UpdateDATATYPEDEFINITIONDATE
//
// swagger:route PATCH /datatypedefinitiondates/{ID} datatypedefinitiondates updateDATATYPEDEFINITIONDATE
//
// # Update a datatypedefinitiondate
//
// Responses:
// default: genericError
//
//	200: datatypedefinitiondateDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONDATE(c *gin.Context) {

	mutexDATATYPEDEFINITIONDATE.Lock()
	defer mutexDATATYPEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONDATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitiondateDB orm.DATATYPEDEFINITIONDATEDB

	// fetch the datatypedefinitiondate
	query := db.First(&datatypedefinitiondateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitiondateDB.CopyBasicFieldsFromDATATYPEDEFINITIONDATE_WOP(&input.DATATYPEDEFINITIONDATE_WOP)
	datatypedefinitiondateDB.DATATYPEDEFINITIONDATEPointersEncoding = input.DATATYPEDEFINITIONDATEPointersEncoding

	query = db.Model(&datatypedefinitiondateDB).Updates(datatypedefinitiondateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitiondateNew := new(models.DATATYPEDEFINITIONDATE)
	datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE(datatypedefinitiondateNew)

	// redeem pointers
	datatypedefinitiondateDB.DecodePointers(backRepo, datatypedefinitiondateNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitiondateOld := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]
	if datatypedefinitiondateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitiondateOld, datatypedefinitiondateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitiondateDB
	c.JSON(http.StatusOK, datatypedefinitiondateDB)
}

// DeleteDATATYPEDEFINITIONDATE
//
// swagger:route DELETE /datatypedefinitiondates/{ID} datatypedefinitiondates deleteDATATYPEDEFINITIONDATE
//
// # Delete a datatypedefinitiondate
//
// default: genericError
//
//	200: datatypedefinitiondateDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONDATE(c *gin.Context) {

	mutexDATATYPEDEFINITIONDATE.Lock()
	defer mutexDATATYPEDEFINITIONDATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONDATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONDATE.GetDB()

	// Get model if exist
	var datatypedefinitiondateDB orm.DATATYPEDEFINITIONDATEDB
	if err := db.First(&datatypedefinitiondateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitiondateDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitiondateDeleted := new(models.DATATYPEDEFINITIONDATE)
	datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE(datatypedefinitiondateDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitiondateStaged := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]
	if datatypedefinitiondateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitiondateStaged, datatypedefinitiondateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
