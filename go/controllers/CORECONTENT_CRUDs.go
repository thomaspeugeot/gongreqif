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
var __CORECONTENT__dummysDeclaration__ models.CORECONTENT
var __CORECONTENT_time__dummyDeclaration time.Duration

var mutexCORECONTENT sync.Mutex

// An CORECONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCORECONTENT updateCORECONTENT deleteCORECONTENT
type CORECONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CORECONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCORECONTENT updateCORECONTENT
type CORECONTENTInput struct {
	// The CORECONTENT to submit or modify
	// in: body
	CORECONTENT *orm.CORECONTENTAPI
}

// GetCORECONTENTs
//
// swagger:route GET /corecontents corecontents getCORECONTENTs
//
// # Get all corecontents
//
// Responses:
// default: genericError
//
//	200: corecontentDBResponse
func (controller *Controller) GetCORECONTENTs(c *gin.Context) {

	// source slice
	var corecontentDBs []orm.CORECONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCORECONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCORECONTENT.GetDB()

	query := db.Find(&corecontentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	corecontentAPIs := make([]orm.CORECONTENTAPI, 0)

	// for each corecontent, update fields from the database nullable fields
	for idx := range corecontentDBs {
		corecontentDB := &corecontentDBs[idx]
		_ = corecontentDB
		var corecontentAPI orm.CORECONTENTAPI

		// insertion point for updating fields
		corecontentAPI.ID = corecontentDB.ID
		corecontentDB.CopyBasicFieldsToCORECONTENT_WOP(&corecontentAPI.CORECONTENT_WOP)
		corecontentAPI.CORECONTENTPointersEncoding = corecontentDB.CORECONTENTPointersEncoding
		corecontentAPIs = append(corecontentAPIs, corecontentAPI)
	}

	c.JSON(http.StatusOK, corecontentAPIs)
}

// PostCORECONTENT
//
// swagger:route POST /corecontents corecontents postCORECONTENT
//
// Creates a corecontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCORECONTENT(c *gin.Context) {

	mutexCORECONTENT.Lock()
	defer mutexCORECONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCORECONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCORECONTENT.GetDB()

	// Validate input
	var input orm.CORECONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create corecontent
	corecontentDB := orm.CORECONTENTDB{}
	corecontentDB.CORECONTENTPointersEncoding = input.CORECONTENTPointersEncoding
	corecontentDB.CopyBasicFieldsFromCORECONTENT_WOP(&input.CORECONTENT_WOP)

	query := db.Create(&corecontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCORECONTENT.CheckoutPhaseOneInstance(&corecontentDB)
	corecontent := backRepo.BackRepoCORECONTENT.Map_CORECONTENTDBID_CORECONTENTPtr[corecontentDB.ID]

	if corecontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), corecontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, corecontentDB)
}

// GetCORECONTENT
//
// swagger:route GET /corecontents/{ID} corecontents getCORECONTENT
//
// Gets the details for a corecontent.
//
// Responses:
// default: genericError
//
//	200: corecontentDBResponse
func (controller *Controller) GetCORECONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCORECONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCORECONTENT.GetDB()

	// Get corecontentDB in DB
	var corecontentDB orm.CORECONTENTDB
	if err := db.First(&corecontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var corecontentAPI orm.CORECONTENTAPI
	corecontentAPI.ID = corecontentDB.ID
	corecontentAPI.CORECONTENTPointersEncoding = corecontentDB.CORECONTENTPointersEncoding
	corecontentDB.CopyBasicFieldsToCORECONTENT_WOP(&corecontentAPI.CORECONTENT_WOP)

	c.JSON(http.StatusOK, corecontentAPI)
}

// UpdateCORECONTENT
//
// swagger:route PATCH /corecontents/{ID} corecontents updateCORECONTENT
//
// # Update a corecontent
//
// Responses:
// default: genericError
//
//	200: corecontentDBResponse
func (controller *Controller) UpdateCORECONTENT(c *gin.Context) {

	mutexCORECONTENT.Lock()
	defer mutexCORECONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCORECONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCORECONTENT.GetDB()

	// Validate input
	var input orm.CORECONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var corecontentDB orm.CORECONTENTDB

	// fetch the corecontent
	query := db.First(&corecontentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	corecontentDB.CopyBasicFieldsFromCORECONTENT_WOP(&input.CORECONTENT_WOP)
	corecontentDB.CORECONTENTPointersEncoding = input.CORECONTENTPointersEncoding

	query = db.Model(&corecontentDB).Updates(corecontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	corecontentNew := new(models.CORECONTENT)
	corecontentDB.CopyBasicFieldsToCORECONTENT(corecontentNew)

	// redeem pointers
	corecontentDB.DecodePointers(backRepo, corecontentNew)

	// get stage instance from DB instance, and call callback function
	corecontentOld := backRepo.BackRepoCORECONTENT.Map_CORECONTENTDBID_CORECONTENTPtr[corecontentDB.ID]
	if corecontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), corecontentOld, corecontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the corecontentDB
	c.JSON(http.StatusOK, corecontentDB)
}

// DeleteCORECONTENT
//
// swagger:route DELETE /corecontents/{ID} corecontents deleteCORECONTENT
//
// # Delete a corecontent
//
// default: genericError
//
//	200: corecontentDBResponse
func (controller *Controller) DeleteCORECONTENT(c *gin.Context) {

	mutexCORECONTENT.Lock()
	defer mutexCORECONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCORECONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCORECONTENT.GetDB()

	// Get model if exist
	var corecontentDB orm.CORECONTENTDB
	if err := db.First(&corecontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&corecontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	corecontentDeleted := new(models.CORECONTENT)
	corecontentDB.CopyBasicFieldsToCORECONTENT(corecontentDeleted)

	// get stage instance from DB instance, and call callback function
	corecontentStaged := backRepo.BackRepoCORECONTENT.Map_CORECONTENTDBID_CORECONTENTPtr[corecontentDB.ID]
	if corecontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), corecontentStaged, corecontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
