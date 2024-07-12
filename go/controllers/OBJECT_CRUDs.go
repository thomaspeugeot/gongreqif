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
var __OBJECT__dummysDeclaration__ models.OBJECT
var __OBJECT_time__dummyDeclaration time.Duration

var mutexOBJECT sync.Mutex

// An OBJECTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOBJECT updateOBJECT deleteOBJECT
type OBJECTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OBJECTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOBJECT updateOBJECT
type OBJECTInput struct {
	// The OBJECT to submit or modify
	// in: body
	OBJECT *orm.OBJECTAPI
}

// GetOBJECTs
//
// swagger:route GET /objects objects getOBJECTs
//
// # Get all objects
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) GetOBJECTs(c *gin.Context) {

	// source slice
	var objectDBs []orm.OBJECTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOBJECTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOBJECT.GetDB()

	query := db.Find(&objectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	objectAPIs := make([]orm.OBJECTAPI, 0)

	// for each object, update fields from the database nullable fields
	for idx := range objectDBs {
		objectDB := &objectDBs[idx]
		_ = objectDB
		var objectAPI orm.OBJECTAPI

		// insertion point for updating fields
		objectAPI.ID = objectDB.ID
		objectDB.CopyBasicFieldsToOBJECT_WOP(&objectAPI.OBJECT_WOP)
		objectAPI.OBJECTPointersEncoding = objectDB.OBJECTPointersEncoding
		objectAPIs = append(objectAPIs, objectAPI)
	}

	c.JSON(http.StatusOK, objectAPIs)
}

// PostOBJECT
//
// swagger:route POST /objects objects postOBJECT
//
// Creates a object
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOBJECT(c *gin.Context) {

	mutexOBJECT.Lock()
	defer mutexOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOBJECTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOBJECT.GetDB()

	// Validate input
	var input orm.OBJECTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create object
	objectDB := orm.OBJECTDB{}
	objectDB.OBJECTPointersEncoding = input.OBJECTPointersEncoding
	objectDB.CopyBasicFieldsFromOBJECT_WOP(&input.OBJECT_WOP)

	query := db.Create(&objectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOBJECT.CheckoutPhaseOneInstance(&objectDB)
	object := backRepo.BackRepoOBJECT.Map_OBJECTDBID_OBJECTPtr[objectDB.ID]

	if object != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), object)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, objectDB)
}

// GetOBJECT
//
// swagger:route GET /objects/{ID} objects getOBJECT
//
// Gets the details for a object.
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) GetOBJECT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOBJECT.GetDB()

	// Get objectDB in DB
	var objectDB orm.OBJECTDB
	if err := db.First(&objectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var objectAPI orm.OBJECTAPI
	objectAPI.ID = objectDB.ID
	objectAPI.OBJECTPointersEncoding = objectDB.OBJECTPointersEncoding
	objectDB.CopyBasicFieldsToOBJECT_WOP(&objectAPI.OBJECT_WOP)

	c.JSON(http.StatusOK, objectAPI)
}

// UpdateOBJECT
//
// swagger:route PATCH /objects/{ID} objects updateOBJECT
//
// # Update a object
//
// Responses:
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) UpdateOBJECT(c *gin.Context) {

	mutexOBJECT.Lock()
	defer mutexOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOBJECT.GetDB()

	// Validate input
	var input orm.OBJECTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var objectDB orm.OBJECTDB

	// fetch the object
	query := db.First(&objectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	objectDB.CopyBasicFieldsFromOBJECT_WOP(&input.OBJECT_WOP)
	objectDB.OBJECTPointersEncoding = input.OBJECTPointersEncoding

	query = db.Model(&objectDB).Updates(objectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	objectNew := new(models.OBJECT)
	objectDB.CopyBasicFieldsToOBJECT(objectNew)

	// redeem pointers
	objectDB.DecodePointers(backRepo, objectNew)

	// get stage instance from DB instance, and call callback function
	objectOld := backRepo.BackRepoOBJECT.Map_OBJECTDBID_OBJECTPtr[objectDB.ID]
	if objectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), objectOld, objectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the objectDB
	c.JSON(http.StatusOK, objectDB)
}

// DeleteOBJECT
//
// swagger:route DELETE /objects/{ID} objects deleteOBJECT
//
// # Delete a object
//
// default: genericError
//
//	200: objectDBResponse
func (controller *Controller) DeleteOBJECT(c *gin.Context) {

	mutexOBJECT.Lock()
	defer mutexOBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOBJECT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOBJECT.GetDB()

	// Get model if exist
	var objectDB orm.OBJECTDB
	if err := db.First(&objectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&objectDB)

	// get an instance (not staged) from DB instance, and call callback function
	objectDeleted := new(models.OBJECT)
	objectDB.CopyBasicFieldsToOBJECT(objectDeleted)

	// get stage instance from DB instance, and call callback function
	objectStaged := backRepo.BackRepoOBJECT.Map_OBJECTDBID_OBJECTPtr[objectDB.ID]
	if objectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), objectStaged, objectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
