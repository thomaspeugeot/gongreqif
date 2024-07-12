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
var __DATATYPEDEFINITIONSTRING__dummysDeclaration__ models.DATATYPEDEFINITIONSTRING
var __DATATYPEDEFINITIONSTRING_time__dummyDeclaration time.Duration

var mutexDATATYPEDEFINITIONSTRING sync.Mutex

// An DATATYPEDEFINITIONSTRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPEDEFINITIONSTRING updateDATATYPEDEFINITIONSTRING deleteDATATYPEDEFINITIONSTRING
type DATATYPEDEFINITIONSTRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPEDEFINITIONSTRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPEDEFINITIONSTRING updateDATATYPEDEFINITIONSTRING
type DATATYPEDEFINITIONSTRINGInput struct {
	// The DATATYPEDEFINITIONSTRING to submit or modify
	// in: body
	DATATYPEDEFINITIONSTRING *orm.DATATYPEDEFINITIONSTRINGAPI
}

// GetDATATYPEDEFINITIONSTRINGs
//
// swagger:route GET /datatypedefinitionstrings datatypedefinitionstrings getDATATYPEDEFINITIONSTRINGs
//
// # Get all datatypedefinitionstrings
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionstringDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONSTRINGs(c *gin.Context) {

	// source slice
	var datatypedefinitionstringDBs []orm.DATATYPEDEFINITIONSTRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONSTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDB()

	query := db.Find(&datatypedefinitionstringDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatypedefinitionstringAPIs := make([]orm.DATATYPEDEFINITIONSTRINGAPI, 0)

	// for each datatypedefinitionstring, update fields from the database nullable fields
	for idx := range datatypedefinitionstringDBs {
		datatypedefinitionstringDB := &datatypedefinitionstringDBs[idx]
		_ = datatypedefinitionstringDB
		var datatypedefinitionstringAPI orm.DATATYPEDEFINITIONSTRINGAPI

		// insertion point for updating fields
		datatypedefinitionstringAPI.ID = datatypedefinitionstringDB.ID
		datatypedefinitionstringDB.CopyBasicFieldsToDATATYPEDEFINITIONSTRING_WOP(&datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRING_WOP)
		datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRINGPointersEncoding = datatypedefinitionstringDB.DATATYPEDEFINITIONSTRINGPointersEncoding
		datatypedefinitionstringAPIs = append(datatypedefinitionstringAPIs, datatypedefinitionstringAPI)
	}

	c.JSON(http.StatusOK, datatypedefinitionstringAPIs)
}

// PostDATATYPEDEFINITIONSTRING
//
// swagger:route POST /datatypedefinitionstrings datatypedefinitionstrings postDATATYPEDEFINITIONSTRING
//
// Creates a datatypedefinitionstring
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPEDEFINITIONSTRING(c *gin.Context) {

	mutexDATATYPEDEFINITIONSTRING.Lock()
	defer mutexDATATYPEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPEDEFINITIONSTRINGs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONSTRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatypedefinitionstring
	datatypedefinitionstringDB := orm.DATATYPEDEFINITIONSTRINGDB{}
	datatypedefinitionstringDB.DATATYPEDEFINITIONSTRINGPointersEncoding = input.DATATYPEDEFINITIONSTRINGPointersEncoding
	datatypedefinitionstringDB.CopyBasicFieldsFromDATATYPEDEFINITIONSTRING_WOP(&input.DATATYPEDEFINITIONSTRING_WOP)

	query := db.Create(&datatypedefinitionstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.CheckoutPhaseOneInstance(&datatypedefinitionstringDB)
	datatypedefinitionstring := backRepo.BackRepoDATATYPEDEFINITIONSTRING.Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGPtr[datatypedefinitionstringDB.ID]

	if datatypedefinitionstring != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatypedefinitionstring)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatypedefinitionstringDB)
}

// GetDATATYPEDEFINITIONSTRING
//
// swagger:route GET /datatypedefinitionstrings/{ID} datatypedefinitionstrings getDATATYPEDEFINITIONSTRING
//
// Gets the details for a datatypedefinitionstring.
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionstringDBResponse
func (controller *Controller) GetDATATYPEDEFINITIONSTRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDB()

	// Get datatypedefinitionstringDB in DB
	var datatypedefinitionstringDB orm.DATATYPEDEFINITIONSTRINGDB
	if err := db.First(&datatypedefinitionstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatypedefinitionstringAPI orm.DATATYPEDEFINITIONSTRINGAPI
	datatypedefinitionstringAPI.ID = datatypedefinitionstringDB.ID
	datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRINGPointersEncoding = datatypedefinitionstringDB.DATATYPEDEFINITIONSTRINGPointersEncoding
	datatypedefinitionstringDB.CopyBasicFieldsToDATATYPEDEFINITIONSTRING_WOP(&datatypedefinitionstringAPI.DATATYPEDEFINITIONSTRING_WOP)

	c.JSON(http.StatusOK, datatypedefinitionstringAPI)
}

// UpdateDATATYPEDEFINITIONSTRING
//
// swagger:route PATCH /datatypedefinitionstrings/{ID} datatypedefinitionstrings updateDATATYPEDEFINITIONSTRING
//
// # Update a datatypedefinitionstring
//
// Responses:
// default: genericError
//
//	200: datatypedefinitionstringDBResponse
func (controller *Controller) UpdateDATATYPEDEFINITIONSTRING(c *gin.Context) {

	mutexDATATYPEDEFINITIONSTRING.Lock()
	defer mutexDATATYPEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDATATYPEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDB()

	// Validate input
	var input orm.DATATYPEDEFINITIONSTRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatypedefinitionstringDB orm.DATATYPEDEFINITIONSTRINGDB

	// fetch the datatypedefinitionstring
	query := db.First(&datatypedefinitionstringDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatypedefinitionstringDB.CopyBasicFieldsFromDATATYPEDEFINITIONSTRING_WOP(&input.DATATYPEDEFINITIONSTRING_WOP)
	datatypedefinitionstringDB.DATATYPEDEFINITIONSTRINGPointersEncoding = input.DATATYPEDEFINITIONSTRINGPointersEncoding

	query = db.Model(&datatypedefinitionstringDB).Updates(datatypedefinitionstringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionstringNew := new(models.DATATYPEDEFINITIONSTRING)
	datatypedefinitionstringDB.CopyBasicFieldsToDATATYPEDEFINITIONSTRING(datatypedefinitionstringNew)

	// redeem pointers
	datatypedefinitionstringDB.DecodePointers(backRepo, datatypedefinitionstringNew)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionstringOld := backRepo.BackRepoDATATYPEDEFINITIONSTRING.Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGPtr[datatypedefinitionstringDB.ID]
	if datatypedefinitionstringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), datatypedefinitionstringOld, datatypedefinitionstringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatypedefinitionstringDB
	c.JSON(http.StatusOK, datatypedefinitionstringDB)
}

// DeleteDATATYPEDEFINITIONSTRING
//
// swagger:route DELETE /datatypedefinitionstrings/{ID} datatypedefinitionstrings deleteDATATYPEDEFINITIONSTRING
//
// # Delete a datatypedefinitionstring
//
// default: genericError
//
//	200: datatypedefinitionstringDBResponse
func (controller *Controller) DeleteDATATYPEDEFINITIONSTRING(c *gin.Context) {

	mutexDATATYPEDEFINITIONSTRING.Lock()
	defer mutexDATATYPEDEFINITIONSTRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPEDEFINITIONSTRING", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDATATYPEDEFINITIONSTRING.GetDB()

	// Get model if exist
	var datatypedefinitionstringDB orm.DATATYPEDEFINITIONSTRINGDB
	if err := db.First(&datatypedefinitionstringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&datatypedefinitionstringDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatypedefinitionstringDeleted := new(models.DATATYPEDEFINITIONSTRING)
	datatypedefinitionstringDB.CopyBasicFieldsToDATATYPEDEFINITIONSTRING(datatypedefinitionstringDeleted)

	// get stage instance from DB instance, and call callback function
	datatypedefinitionstringStaged := backRepo.BackRepoDATATYPEDEFINITIONSTRING.Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGPtr[datatypedefinitionstringDB.ID]
	if datatypedefinitionstringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatypedefinitionstringStaged, datatypedefinitionstringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
