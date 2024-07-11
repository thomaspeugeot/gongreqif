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
var __CONTENT__dummysDeclaration__ models.CONTENT
var __CONTENT_time__dummyDeclaration time.Duration

var mutexCONTENT sync.Mutex

// An CONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCONTENT updateCONTENT deleteCONTENT
type CONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCONTENT updateCONTENT
type CONTENTInput struct {
	// The CONTENT to submit or modify
	// in: body
	CONTENT *orm.CONTENTAPI
}

// GetCONTENTs
//
// swagger:route GET /contents contents getCONTENTs
//
// # Get all contents
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) GetCONTENTs(c *gin.Context) {

	// source slice
	var contentDBs []orm.CONTENTDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCONTENT.GetDB()

	query := db.Find(&contentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	contentAPIs := make([]orm.CONTENTAPI, 0)

	// for each content, update fields from the database nullable fields
	for idx := range contentDBs {
		contentDB := &contentDBs[idx]
		_ = contentDB
		var contentAPI orm.CONTENTAPI

		// insertion point for updating fields
		contentAPI.ID = contentDB.ID
		contentDB.CopyBasicFieldsToCONTENT_WOP(&contentAPI.CONTENT_WOP)
		contentAPI.CONTENTPointersEncoding = contentDB.CONTENTPointersEncoding
		contentAPIs = append(contentAPIs, contentAPI)
	}

	c.JSON(http.StatusOK, contentAPIs)
}

// PostCONTENT
//
// swagger:route POST /contents contents postCONTENT
//
// Creates a content
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCONTENT(c *gin.Context) {

	mutexCONTENT.Lock()
	defer mutexCONTENT.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCONTENT.GetDB()

	// Validate input
	var input orm.CONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create content
	contentDB := orm.CONTENTDB{}
	contentDB.CONTENTPointersEncoding = input.CONTENTPointersEncoding
	contentDB.CopyBasicFieldsFromCONTENT_WOP(&input.CONTENT_WOP)

	query := db.Create(&contentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCONTENT.CheckoutPhaseOneInstance(&contentDB)
	content := backRepo.BackRepoCONTENT.Map_CONTENTDBID_CONTENTPtr[contentDB.ID]

	if content != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), content)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, contentDB)
}

// GetCONTENT
//
// swagger:route GET /contents/{ID} contents getCONTENT
//
// Gets the details for a content.
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) GetCONTENT(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCONTENT.GetDB()

	// Get contentDB in DB
	var contentDB orm.CONTENTDB
	if err := db.First(&contentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var contentAPI orm.CONTENTAPI
	contentAPI.ID = contentDB.ID
	contentAPI.CONTENTPointersEncoding = contentDB.CONTENTPointersEncoding
	contentDB.CopyBasicFieldsToCONTENT_WOP(&contentAPI.CONTENT_WOP)

	c.JSON(http.StatusOK, contentAPI)
}

// UpdateCONTENT
//
// swagger:route PATCH /contents/{ID} contents updateCONTENT
//
// # Update a content
//
// Responses:
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) UpdateCONTENT(c *gin.Context) {

	mutexCONTENT.Lock()
	defer mutexCONTENT.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCONTENT.GetDB()

	// Validate input
	var input orm.CONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var contentDB orm.CONTENTDB

	// fetch the content
	query := db.First(&contentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	contentDB.CopyBasicFieldsFromCONTENT_WOP(&input.CONTENT_WOP)
	contentDB.CONTENTPointersEncoding = input.CONTENTPointersEncoding

	query = db.Model(&contentDB).Updates(contentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	contentNew := new(models.CONTENT)
	contentDB.CopyBasicFieldsToCONTENT(contentNew)

	// redeem pointers
	contentDB.DecodePointers(backRepo, contentNew)

	// get stage instance from DB instance, and call callback function
	contentOld := backRepo.BackRepoCONTENT.Map_CONTENTDBID_CONTENTPtr[contentDB.ID]
	if contentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), contentOld, contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the contentDB
	c.JSON(http.StatusOK, contentDB)
}

// DeleteCONTENT
//
// swagger:route DELETE /contents/{ID} contents deleteCONTENT
//
// # Delete a content
//
// default: genericError
//
//	200: contentDBResponse
func (controller *Controller) DeleteCONTENT(c *gin.Context) {

	mutexCONTENT.Lock()
	defer mutexCONTENT.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCONTENT.GetDB()

	// Get model if exist
	var contentDB orm.CONTENTDB
	if err := db.First(&contentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&contentDB)

	// get an instance (not staged) from DB instance, and call callback function
	contentDeleted := new(models.CONTENT)
	contentDB.CopyBasicFieldsToCONTENT(contentDeleted)

	// get stage instance from DB instance, and call callback function
	contentStaged := backRepo.BackRepoCONTENT.Map_CONTENTDBID_CONTENTPtr[contentDB.ID]
	if contentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), contentStaged, contentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
