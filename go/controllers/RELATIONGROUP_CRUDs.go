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
var __RELATIONGROUP__dummysDeclaration__ models.RELATIONGROUP
var __RELATIONGROUP_time__dummyDeclaration time.Duration

var mutexRELATIONGROUP sync.Mutex

// An RELATIONGROUPID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRELATIONGROUP updateRELATIONGROUP deleteRELATIONGROUP
type RELATIONGROUPID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RELATIONGROUPInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRELATIONGROUP updateRELATIONGROUP
type RELATIONGROUPInput struct {
	// The RELATIONGROUP to submit or modify
	// in: body
	RELATIONGROUP *orm.RELATIONGROUPAPI
}

// GetRELATIONGROUPs
//
// swagger:route GET /relationgroups relationgroups getRELATIONGROUPs
//
// # Get all relationgroups
//
// Responses:
// default: genericError
//
//	200: relationgroupDBResponse
func (controller *Controller) GetRELATIONGROUPs(c *gin.Context) {

	// source slice
	var relationgroupDBs []orm.RELATIONGROUPDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATIONGROUPs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUP.GetDB()

	query := db.Find(&relationgroupDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	relationgroupAPIs := make([]orm.RELATIONGROUPAPI, 0)

	// for each relationgroup, update fields from the database nullable fields
	for idx := range relationgroupDBs {
		relationgroupDB := &relationgroupDBs[idx]
		_ = relationgroupDB
		var relationgroupAPI orm.RELATIONGROUPAPI

		// insertion point for updating fields
		relationgroupAPI.ID = relationgroupDB.ID
		relationgroupDB.CopyBasicFieldsToRELATIONGROUP_WOP(&relationgroupAPI.RELATIONGROUP_WOP)
		relationgroupAPI.RELATIONGROUPPointersEncoding = relationgroupDB.RELATIONGROUPPointersEncoding
		relationgroupAPIs = append(relationgroupAPIs, relationgroupAPI)
	}

	c.JSON(http.StatusOK, relationgroupAPIs)
}

// PostRELATIONGROUP
//
// swagger:route POST /relationgroups relationgroups postRELATIONGROUP
//
// Creates a relationgroup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRELATIONGROUP(c *gin.Context) {

	mutexRELATIONGROUP.Lock()
	defer mutexRELATIONGROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRELATIONGROUPs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUP.GetDB()

	// Validate input
	var input orm.RELATIONGROUPAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create relationgroup
	relationgroupDB := orm.RELATIONGROUPDB{}
	relationgroupDB.RELATIONGROUPPointersEncoding = input.RELATIONGROUPPointersEncoding
	relationgroupDB.CopyBasicFieldsFromRELATIONGROUP_WOP(&input.RELATIONGROUP_WOP)

	query := db.Create(&relationgroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRELATIONGROUP.CheckoutPhaseOneInstance(&relationgroupDB)
	relationgroup := backRepo.BackRepoRELATIONGROUP.Map_RELATIONGROUPDBID_RELATIONGROUPPtr[relationgroupDB.ID]

	if relationgroup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), relationgroup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, relationgroupDB)
}

// GetRELATIONGROUP
//
// swagger:route GET /relationgroups/{ID} relationgroups getRELATIONGROUP
//
// Gets the details for a relationgroup.
//
// Responses:
// default: genericError
//
//	200: relationgroupDBResponse
func (controller *Controller) GetRELATIONGROUP(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATIONGROUP", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUP.GetDB()

	// Get relationgroupDB in DB
	var relationgroupDB orm.RELATIONGROUPDB
	if err := db.First(&relationgroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var relationgroupAPI orm.RELATIONGROUPAPI
	relationgroupAPI.ID = relationgroupDB.ID
	relationgroupAPI.RELATIONGROUPPointersEncoding = relationgroupDB.RELATIONGROUPPointersEncoding
	relationgroupDB.CopyBasicFieldsToRELATIONGROUP_WOP(&relationgroupAPI.RELATIONGROUP_WOP)

	c.JSON(http.StatusOK, relationgroupAPI)
}

// UpdateRELATIONGROUP
//
// swagger:route PATCH /relationgroups/{ID} relationgroups updateRELATIONGROUP
//
// # Update a relationgroup
//
// Responses:
// default: genericError
//
//	200: relationgroupDBResponse
func (controller *Controller) UpdateRELATIONGROUP(c *gin.Context) {

	mutexRELATIONGROUP.Lock()
	defer mutexRELATIONGROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRELATIONGROUP", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUP.GetDB()

	// Validate input
	var input orm.RELATIONGROUPAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var relationgroupDB orm.RELATIONGROUPDB

	// fetch the relationgroup
	query := db.First(&relationgroupDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	relationgroupDB.CopyBasicFieldsFromRELATIONGROUP_WOP(&input.RELATIONGROUP_WOP)
	relationgroupDB.RELATIONGROUPPointersEncoding = input.RELATIONGROUPPointersEncoding

	query = db.Model(&relationgroupDB).Updates(relationgroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	relationgroupNew := new(models.RELATIONGROUP)
	relationgroupDB.CopyBasicFieldsToRELATIONGROUP(relationgroupNew)

	// redeem pointers
	relationgroupDB.DecodePointers(backRepo, relationgroupNew)

	// get stage instance from DB instance, and call callback function
	relationgroupOld := backRepo.BackRepoRELATIONGROUP.Map_RELATIONGROUPDBID_RELATIONGROUPPtr[relationgroupDB.ID]
	if relationgroupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), relationgroupOld, relationgroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the relationgroupDB
	c.JSON(http.StatusOK, relationgroupDB)
}

// DeleteRELATIONGROUP
//
// swagger:route DELETE /relationgroups/{ID} relationgroups deleteRELATIONGROUP
//
// # Delete a relationgroup
//
// default: genericError
//
//	200: relationgroupDBResponse
func (controller *Controller) DeleteRELATIONGROUP(c *gin.Context) {

	mutexRELATIONGROUP.Lock()
	defer mutexRELATIONGROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRELATIONGROUP", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRELATIONGROUP.GetDB()

	// Get model if exist
	var relationgroupDB orm.RELATIONGROUPDB
	if err := db.First(&relationgroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&relationgroupDB)

	// get an instance (not staged) from DB instance, and call callback function
	relationgroupDeleted := new(models.RELATIONGROUP)
	relationgroupDB.CopyBasicFieldsToRELATIONGROUP(relationgroupDeleted)

	// get stage instance from DB instance, and call callback function
	relationgroupStaged := backRepo.BackRepoRELATIONGROUP.Map_RELATIONGROUPDBID_RELATIONGROUPPtr[relationgroupDB.ID]
	if relationgroupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), relationgroupStaged, relationgroupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
