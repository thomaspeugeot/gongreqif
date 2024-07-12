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
var __CHILDREN__dummysDeclaration__ models.CHILDREN
var __CHILDREN_time__dummyDeclaration time.Duration

var mutexCHILDREN sync.Mutex

// An CHILDRENID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCHILDREN updateCHILDREN deleteCHILDREN
type CHILDRENID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CHILDRENInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCHILDREN updateCHILDREN
type CHILDRENInput struct {
	// The CHILDREN to submit or modify
	// in: body
	CHILDREN *orm.CHILDRENAPI
}

// GetCHILDRENs
//
// swagger:route GET /childrens childrens getCHILDRENs
//
// # Get all childrens
//
// Responses:
// default: genericError
//
//	200: childrenDBResponse
func (controller *Controller) GetCHILDRENs(c *gin.Context) {

	// source slice
	var childrenDBs []orm.CHILDRENDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCHILDRENs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCHILDREN.GetDB()

	query := db.Find(&childrenDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	childrenAPIs := make([]orm.CHILDRENAPI, 0)

	// for each children, update fields from the database nullable fields
	for idx := range childrenDBs {
		childrenDB := &childrenDBs[idx]
		_ = childrenDB
		var childrenAPI orm.CHILDRENAPI

		// insertion point for updating fields
		childrenAPI.ID = childrenDB.ID
		childrenDB.CopyBasicFieldsToCHILDREN_WOP(&childrenAPI.CHILDREN_WOP)
		childrenAPI.CHILDRENPointersEncoding = childrenDB.CHILDRENPointersEncoding
		childrenAPIs = append(childrenAPIs, childrenAPI)
	}

	c.JSON(http.StatusOK, childrenAPIs)
}

// PostCHILDREN
//
// swagger:route POST /childrens childrens postCHILDREN
//
// Creates a children
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCHILDREN(c *gin.Context) {

	mutexCHILDREN.Lock()
	defer mutexCHILDREN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCHILDRENs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCHILDREN.GetDB()

	// Validate input
	var input orm.CHILDRENAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create children
	childrenDB := orm.CHILDRENDB{}
	childrenDB.CHILDRENPointersEncoding = input.CHILDRENPointersEncoding
	childrenDB.CopyBasicFieldsFromCHILDREN_WOP(&input.CHILDREN_WOP)

	query := db.Create(&childrenDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCHILDREN.CheckoutPhaseOneInstance(&childrenDB)
	children := backRepo.BackRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]

	if children != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), children)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, childrenDB)
}

// GetCHILDREN
//
// swagger:route GET /childrens/{ID} childrens getCHILDREN
//
// Gets the details for a children.
//
// Responses:
// default: genericError
//
//	200: childrenDBResponse
func (controller *Controller) GetCHILDREN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCHILDREN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCHILDREN.GetDB()

	// Get childrenDB in DB
	var childrenDB orm.CHILDRENDB
	if err := db.First(&childrenDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var childrenAPI orm.CHILDRENAPI
	childrenAPI.ID = childrenDB.ID
	childrenAPI.CHILDRENPointersEncoding = childrenDB.CHILDRENPointersEncoding
	childrenDB.CopyBasicFieldsToCHILDREN_WOP(&childrenAPI.CHILDREN_WOP)

	c.JSON(http.StatusOK, childrenAPI)
}

// UpdateCHILDREN
//
// swagger:route PATCH /childrens/{ID} childrens updateCHILDREN
//
// # Update a children
//
// Responses:
// default: genericError
//
//	200: childrenDBResponse
func (controller *Controller) UpdateCHILDREN(c *gin.Context) {

	mutexCHILDREN.Lock()
	defer mutexCHILDREN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCHILDREN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCHILDREN.GetDB()

	// Validate input
	var input orm.CHILDRENAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var childrenDB orm.CHILDRENDB

	// fetch the children
	query := db.First(&childrenDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	childrenDB.CopyBasicFieldsFromCHILDREN_WOP(&input.CHILDREN_WOP)
	childrenDB.CHILDRENPointersEncoding = input.CHILDRENPointersEncoding

	query = db.Model(&childrenDB).Updates(childrenDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	childrenNew := new(models.CHILDREN)
	childrenDB.CopyBasicFieldsToCHILDREN(childrenNew)

	// redeem pointers
	childrenDB.DecodePointers(backRepo, childrenNew)

	// get stage instance from DB instance, and call callback function
	childrenOld := backRepo.BackRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]
	if childrenOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), childrenOld, childrenNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the childrenDB
	c.JSON(http.StatusOK, childrenDB)
}

// DeleteCHILDREN
//
// swagger:route DELETE /childrens/{ID} childrens deleteCHILDREN
//
// # Delete a children
//
// default: genericError
//
//	200: childrenDBResponse
func (controller *Controller) DeleteCHILDREN(c *gin.Context) {

	mutexCHILDREN.Lock()
	defer mutexCHILDREN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCHILDREN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCHILDREN.GetDB()

	// Get model if exist
	var childrenDB orm.CHILDRENDB
	if err := db.First(&childrenDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&childrenDB)

	// get an instance (not staged) from DB instance, and call callback function
	childrenDeleted := new(models.CHILDREN)
	childrenDB.CopyBasicFieldsToCHILDREN(childrenDeleted)

	// get stage instance from DB instance, and call callback function
	childrenStaged := backRepo.BackRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]
	if childrenStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), childrenStaged, childrenDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
