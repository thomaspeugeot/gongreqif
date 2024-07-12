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
var __XHTMLCONTENT__dummysDeclaration__ models.XHTMLCONTENT
var __XHTMLCONTENT_time__dummyDeclaration time.Duration

var mutexXHTMLCONTENT sync.Mutex

// An XHTMLCONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXHTMLCONTENT updateXHTMLCONTENT deleteXHTMLCONTENT
type XHTMLCONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XHTMLCONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXHTMLCONTENT updateXHTMLCONTENT
type XHTMLCONTENTInput struct {
	// The XHTMLCONTENT to submit or modify
	// in: body
	XHTMLCONTENT *orm.XHTMLCONTENTAPI
}

// GetXHTMLCONTENTs
//
// swagger:route GET /xhtmlcontents xhtmlcontents getXHTMLCONTENTs
//
// # Get all xhtmlcontents
//
// Responses:
// default: genericError
//
//	200: xhtmlcontentDBResponse
func (controller *Controller) GetXHTMLCONTENTs(c *gin.Context) {

	// source slice
	var xhtmlcontentDBs []orm.XHTMLCONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXHTMLCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXHTMLCONTENT.GetDB()

	query := db.Find(&xhtmlcontentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtmlcontentAPIs := make([]orm.XHTMLCONTENTAPI, 0)

	// for each xhtmlcontent, update fields from the database nullable fields
	for idx := range xhtmlcontentDBs {
		xhtmlcontentDB := &xhtmlcontentDBs[idx]
		_ = xhtmlcontentDB
		var xhtmlcontentAPI orm.XHTMLCONTENTAPI

		// insertion point for updating fields
		xhtmlcontentAPI.ID = xhtmlcontentDB.ID
		xhtmlcontentDB.CopyBasicFieldsToXHTMLCONTENT_WOP(&xhtmlcontentAPI.XHTMLCONTENT_WOP)
		xhtmlcontentAPI.XHTMLCONTENTPointersEncoding = xhtmlcontentDB.XHTMLCONTENTPointersEncoding
		xhtmlcontentAPIs = append(xhtmlcontentAPIs, xhtmlcontentAPI)
	}

	c.JSON(http.StatusOK, xhtmlcontentAPIs)
}

// PostXHTMLCONTENT
//
// swagger:route POST /xhtmlcontents xhtmlcontents postXHTMLCONTENT
//
// Creates a xhtmlcontent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXHTMLCONTENT(c *gin.Context) {

	mutexXHTMLCONTENT.Lock()
	defer mutexXHTMLCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXHTMLCONTENTs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXHTMLCONTENT.GetDB()

	// Validate input
	var input orm.XHTMLCONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtmlcontent
	xhtmlcontentDB := orm.XHTMLCONTENTDB{}
	xhtmlcontentDB.XHTMLCONTENTPointersEncoding = input.XHTMLCONTENTPointersEncoding
	xhtmlcontentDB.CopyBasicFieldsFromXHTMLCONTENT_WOP(&input.XHTMLCONTENT_WOP)

	query := db.Create(&xhtmlcontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXHTMLCONTENT.CheckoutPhaseOneInstance(&xhtmlcontentDB)
	xhtmlcontent := backRepo.BackRepoXHTMLCONTENT.Map_XHTMLCONTENTDBID_XHTMLCONTENTPtr[xhtmlcontentDB.ID]

	if xhtmlcontent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtmlcontent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtmlcontentDB)
}

// GetXHTMLCONTENT
//
// swagger:route GET /xhtmlcontents/{ID} xhtmlcontents getXHTMLCONTENT
//
// Gets the details for a xhtmlcontent.
//
// Responses:
// default: genericError
//
//	200: xhtmlcontentDBResponse
func (controller *Controller) GetXHTMLCONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXHTMLCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXHTMLCONTENT.GetDB()

	// Get xhtmlcontentDB in DB
	var xhtmlcontentDB orm.XHTMLCONTENTDB
	if err := db.First(&xhtmlcontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtmlcontentAPI orm.XHTMLCONTENTAPI
	xhtmlcontentAPI.ID = xhtmlcontentDB.ID
	xhtmlcontentAPI.XHTMLCONTENTPointersEncoding = xhtmlcontentDB.XHTMLCONTENTPointersEncoding
	xhtmlcontentDB.CopyBasicFieldsToXHTMLCONTENT_WOP(&xhtmlcontentAPI.XHTMLCONTENT_WOP)

	c.JSON(http.StatusOK, xhtmlcontentAPI)
}

// UpdateXHTMLCONTENT
//
// swagger:route PATCH /xhtmlcontents/{ID} xhtmlcontents updateXHTMLCONTENT
//
// # Update a xhtmlcontent
//
// Responses:
// default: genericError
//
//	200: xhtmlcontentDBResponse
func (controller *Controller) UpdateXHTMLCONTENT(c *gin.Context) {

	mutexXHTMLCONTENT.Lock()
	defer mutexXHTMLCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXHTMLCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXHTMLCONTENT.GetDB()

	// Validate input
	var input orm.XHTMLCONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtmlcontentDB orm.XHTMLCONTENTDB

	// fetch the xhtmlcontent
	query := db.First(&xhtmlcontentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtmlcontentDB.CopyBasicFieldsFromXHTMLCONTENT_WOP(&input.XHTMLCONTENT_WOP)
	xhtmlcontentDB.XHTMLCONTENTPointersEncoding = input.XHTMLCONTENTPointersEncoding

	query = db.Model(&xhtmlcontentDB).Updates(xhtmlcontentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtmlcontentNew := new(models.XHTMLCONTENT)
	xhtmlcontentDB.CopyBasicFieldsToXHTMLCONTENT(xhtmlcontentNew)

	// redeem pointers
	xhtmlcontentDB.DecodePointers(backRepo, xhtmlcontentNew)

	// get stage instance from DB instance, and call callback function
	xhtmlcontentOld := backRepo.BackRepoXHTMLCONTENT.Map_XHTMLCONTENTDBID_XHTMLCONTENTPtr[xhtmlcontentDB.ID]
	if xhtmlcontentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtmlcontentOld, xhtmlcontentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtmlcontentDB
	c.JSON(http.StatusOK, xhtmlcontentDB)
}

// DeleteXHTMLCONTENT
//
// swagger:route DELETE /xhtmlcontents/{ID} xhtmlcontents deleteXHTMLCONTENT
//
// # Delete a xhtmlcontent
//
// default: genericError
//
//	200: xhtmlcontentDBResponse
func (controller *Controller) DeleteXHTMLCONTENT(c *gin.Context) {

	mutexXHTMLCONTENT.Lock()
	defer mutexXHTMLCONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXHTMLCONTENT", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXHTMLCONTENT.GetDB()

	// Get model if exist
	var xhtmlcontentDB orm.XHTMLCONTENTDB
	if err := db.First(&xhtmlcontentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtmlcontentDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtmlcontentDeleted := new(models.XHTMLCONTENT)
	xhtmlcontentDB.CopyBasicFieldsToXHTMLCONTENT(xhtmlcontentDeleted)

	// get stage instance from DB instance, and call callback function
	xhtmlcontentStaged := backRepo.BackRepoXHTMLCONTENT.Map_XHTMLCONTENTDBID_XHTMLCONTENTPtr[xhtmlcontentDB.ID]
	if xhtmlcontentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtmlcontentStaged, xhtmlcontentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
