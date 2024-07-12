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
var __RELATIONGROUPTYPE__dummysDeclaration__ models.RELATIONGROUPTYPE
var __RELATIONGROUPTYPE_time__dummyDeclaration time.Duration

var mutexRELATIONGROUPTYPE sync.Mutex

// An RELATIONGROUPTYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRELATIONGROUPTYPE updateRELATIONGROUPTYPE deleteRELATIONGROUPTYPE
type RELATIONGROUPTYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RELATIONGROUPTYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRELATIONGROUPTYPE updateRELATIONGROUPTYPE
type RELATIONGROUPTYPEInput struct {
	// The RELATIONGROUPTYPE to submit or modify
	// in: body
	RELATIONGROUPTYPE *orm.RELATIONGROUPTYPEAPI
}

// GetRELATIONGROUPTYPEs
//
// swagger:route GET /relationgrouptypes relationgrouptypes getRELATIONGROUPTYPEs
//
// # Get all relationgrouptypes
//
// Responses:
// default: genericError
//
//	200: relationgrouptypeDBResponse
func (controller *Controller) GetRELATIONGROUPTYPEs(c *gin.Context) {

	// source slice
	var relationgrouptypeDBs []orm.RELATIONGROUPTYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATIONGROUPTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUPTYPE.GetDB()

	query := db.Find(&relationgrouptypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	relationgrouptypeAPIs := make([]orm.RELATIONGROUPTYPEAPI, 0)

	// for each relationgrouptype, update fields from the database nullable fields
	for idx := range relationgrouptypeDBs {
		relationgrouptypeDB := &relationgrouptypeDBs[idx]
		_ = relationgrouptypeDB
		var relationgrouptypeAPI orm.RELATIONGROUPTYPEAPI

		// insertion point for updating fields
		relationgrouptypeAPI.ID = relationgrouptypeDB.ID
		relationgrouptypeDB.CopyBasicFieldsToRELATIONGROUPTYPE_WOP(&relationgrouptypeAPI.RELATIONGROUPTYPE_WOP)
		relationgrouptypeAPI.RELATIONGROUPTYPEPointersEncoding = relationgrouptypeDB.RELATIONGROUPTYPEPointersEncoding
		relationgrouptypeAPIs = append(relationgrouptypeAPIs, relationgrouptypeAPI)
	}

	c.JSON(http.StatusOK, relationgrouptypeAPIs)
}

// PostRELATIONGROUPTYPE
//
// swagger:route POST /relationgrouptypes relationgrouptypes postRELATIONGROUPTYPE
//
// Creates a relationgrouptype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRELATIONGROUPTYPE(c *gin.Context) {

	mutexRELATIONGROUPTYPE.Lock()
	defer mutexRELATIONGROUPTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRELATIONGROUPTYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUPTYPE.GetDB()

	// Validate input
	var input orm.RELATIONGROUPTYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create relationgrouptype
	relationgrouptypeDB := orm.RELATIONGROUPTYPEDB{}
	relationgrouptypeDB.RELATIONGROUPTYPEPointersEncoding = input.RELATIONGROUPTYPEPointersEncoding
	relationgrouptypeDB.CopyBasicFieldsFromRELATIONGROUPTYPE_WOP(&input.RELATIONGROUPTYPE_WOP)

	query := db.Create(&relationgrouptypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRELATIONGROUPTYPE.CheckoutPhaseOneInstance(&relationgrouptypeDB)
	relationgrouptype := backRepo.BackRepoRELATIONGROUPTYPE.Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEPtr[relationgrouptypeDB.ID]

	if relationgrouptype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), relationgrouptype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, relationgrouptypeDB)
}

// GetRELATIONGROUPTYPE
//
// swagger:route GET /relationgrouptypes/{ID} relationgrouptypes getRELATIONGROUPTYPE
//
// Gets the details for a relationgrouptype.
//
// Responses:
// default: genericError
//
//	200: relationgrouptypeDBResponse
func (controller *Controller) GetRELATIONGROUPTYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATIONGROUPTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUPTYPE.GetDB()

	// Get relationgrouptypeDB in DB
	var relationgrouptypeDB orm.RELATIONGROUPTYPEDB
	if err := db.First(&relationgrouptypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var relationgrouptypeAPI orm.RELATIONGROUPTYPEAPI
	relationgrouptypeAPI.ID = relationgrouptypeDB.ID
	relationgrouptypeAPI.RELATIONGROUPTYPEPointersEncoding = relationgrouptypeDB.RELATIONGROUPTYPEPointersEncoding
	relationgrouptypeDB.CopyBasicFieldsToRELATIONGROUPTYPE_WOP(&relationgrouptypeAPI.RELATIONGROUPTYPE_WOP)

	c.JSON(http.StatusOK, relationgrouptypeAPI)
}

// UpdateRELATIONGROUPTYPE
//
// swagger:route PATCH /relationgrouptypes/{ID} relationgrouptypes updateRELATIONGROUPTYPE
//
// # Update a relationgrouptype
//
// Responses:
// default: genericError
//
//	200: relationgrouptypeDBResponse
func (controller *Controller) UpdateRELATIONGROUPTYPE(c *gin.Context) {

	mutexRELATIONGROUPTYPE.Lock()
	defer mutexRELATIONGROUPTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRELATIONGROUPTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUPTYPE.GetDB()

	// Validate input
	var input orm.RELATIONGROUPTYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var relationgrouptypeDB orm.RELATIONGROUPTYPEDB

	// fetch the relationgrouptype
	query := db.First(&relationgrouptypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	relationgrouptypeDB.CopyBasicFieldsFromRELATIONGROUPTYPE_WOP(&input.RELATIONGROUPTYPE_WOP)
	relationgrouptypeDB.RELATIONGROUPTYPEPointersEncoding = input.RELATIONGROUPTYPEPointersEncoding

	query = db.Model(&relationgrouptypeDB).Updates(relationgrouptypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	relationgrouptypeNew := new(models.RELATIONGROUPTYPE)
	relationgrouptypeDB.CopyBasicFieldsToRELATIONGROUPTYPE(relationgrouptypeNew)

	// redeem pointers
	relationgrouptypeDB.DecodePointers(backRepo, relationgrouptypeNew)

	// get stage instance from DB instance, and call callback function
	relationgrouptypeOld := backRepo.BackRepoRELATIONGROUPTYPE.Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEPtr[relationgrouptypeDB.ID]
	if relationgrouptypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), relationgrouptypeOld, relationgrouptypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the relationgrouptypeDB
	c.JSON(http.StatusOK, relationgrouptypeDB)
}

// DeleteRELATIONGROUPTYPE
//
// swagger:route DELETE /relationgrouptypes/{ID} relationgrouptypes deleteRELATIONGROUPTYPE
//
// # Delete a relationgrouptype
//
// default: genericError
//
//	200: relationgrouptypeDBResponse
func (controller *Controller) DeleteRELATIONGROUPTYPE(c *gin.Context) {

	mutexRELATIONGROUPTYPE.Lock()
	defer mutexRELATIONGROUPTYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRELATIONGROUPTYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUPTYPE.GetDB()

	// Get model if exist
	var relationgrouptypeDB orm.RELATIONGROUPTYPEDB
	if err := db.First(&relationgrouptypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&relationgrouptypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	relationgrouptypeDeleted := new(models.RELATIONGROUPTYPE)
	relationgrouptypeDB.CopyBasicFieldsToRELATIONGROUPTYPE(relationgrouptypeDeleted)

	// get stage instance from DB instance, and call callback function
	relationgrouptypeStaged := backRepo.BackRepoRELATIONGROUPTYPE.Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEPtr[relationgrouptypeDB.ID]
	if relationgrouptypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), relationgrouptypeStaged, relationgrouptypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
