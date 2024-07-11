// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gong/go")
	{ // insertion point for registrations
		v1.GET("/v1/gongbasicfields", GetController().GetGongBasicFields)
		v1.GET("/v1/gongbasicfields/:id", GetController().GetGongBasicField)
		v1.POST("/v1/gongbasicfields", GetController().PostGongBasicField)
		v1.PATCH("/v1/gongbasicfields/:id", GetController().UpdateGongBasicField)
		v1.PUT("/v1/gongbasicfields/:id", GetController().UpdateGongBasicField)
		v1.DELETE("/v1/gongbasicfields/:id", GetController().DeleteGongBasicField)

		v1.GET("/v1/gongenums", GetController().GetGongEnums)
		v1.GET("/v1/gongenums/:id", GetController().GetGongEnum)
		v1.POST("/v1/gongenums", GetController().PostGongEnum)
		v1.PATCH("/v1/gongenums/:id", GetController().UpdateGongEnum)
		v1.PUT("/v1/gongenums/:id", GetController().UpdateGongEnum)
		v1.DELETE("/v1/gongenums/:id", GetController().DeleteGongEnum)

		v1.GET("/v1/gongenumvalues", GetController().GetGongEnumValues)
		v1.GET("/v1/gongenumvalues/:id", GetController().GetGongEnumValue)
		v1.POST("/v1/gongenumvalues", GetController().PostGongEnumValue)
		v1.PATCH("/v1/gongenumvalues/:id", GetController().UpdateGongEnumValue)
		v1.PUT("/v1/gongenumvalues/:id", GetController().UpdateGongEnumValue)
		v1.DELETE("/v1/gongenumvalues/:id", GetController().DeleteGongEnumValue)

		v1.GET("/v1/gonglinks", GetController().GetGongLinks)
		v1.GET("/v1/gonglinks/:id", GetController().GetGongLink)
		v1.POST("/v1/gonglinks", GetController().PostGongLink)
		v1.PATCH("/v1/gonglinks/:id", GetController().UpdateGongLink)
		v1.PUT("/v1/gonglinks/:id", GetController().UpdateGongLink)
		v1.DELETE("/v1/gonglinks/:id", GetController().DeleteGongLink)

		v1.GET("/v1/gongnotes", GetController().GetGongNotes)
		v1.GET("/v1/gongnotes/:id", GetController().GetGongNote)
		v1.POST("/v1/gongnotes", GetController().PostGongNote)
		v1.PATCH("/v1/gongnotes/:id", GetController().UpdateGongNote)
		v1.PUT("/v1/gongnotes/:id", GetController().UpdateGongNote)
		v1.DELETE("/v1/gongnotes/:id", GetController().DeleteGongNote)

		v1.GET("/v1/gongstructs", GetController().GetGongStructs)
		v1.GET("/v1/gongstructs/:id", GetController().GetGongStruct)
		v1.POST("/v1/gongstructs", GetController().PostGongStruct)
		v1.PATCH("/v1/gongstructs/:id", GetController().UpdateGongStruct)
		v1.PUT("/v1/gongstructs/:id", GetController().UpdateGongStruct)
		v1.DELETE("/v1/gongstructs/:id", GetController().DeleteGongStruct)

		v1.GET("/v1/gongtimefields", GetController().GetGongTimeFields)
		v1.GET("/v1/gongtimefields/:id", GetController().GetGongTimeField)
		v1.POST("/v1/gongtimefields", GetController().PostGongTimeField)
		v1.PATCH("/v1/gongtimefields/:id", GetController().UpdateGongTimeField)
		v1.PUT("/v1/gongtimefields/:id", GetController().UpdateGongTimeField)
		v1.DELETE("/v1/gongtimefields/:id", GetController().DeleteGongTimeField)

		v1.GET("/v1/metas", GetController().GetMetas)
		v1.GET("/v1/metas/:id", GetController().GetMeta)
		v1.POST("/v1/metas", GetController().PostMeta)
		v1.PATCH("/v1/metas/:id", GetController().UpdateMeta)
		v1.PUT("/v1/metas/:id", GetController().UpdateMeta)
		v1.DELETE("/v1/metas/:id", GetController().DeleteMeta)

		v1.GET("/v1/metareferences", GetController().GetMetaReferences)
		v1.GET("/v1/metareferences/:id", GetController().GetMetaReference)
		v1.POST("/v1/metareferences", GetController().PostMetaReference)
		v1.PATCH("/v1/metareferences/:id", GetController().UpdateMetaReference)
		v1.PUT("/v1/metareferences/:id", GetController().UpdateMetaReference)
		v1.DELETE("/v1/metareferences/:id", GetController().DeleteMetaReference)

		v1.GET("/v1/modelpkgs", GetController().GetModelPkgs)
		v1.GET("/v1/modelpkgs/:id", GetController().GetModelPkg)
		v1.POST("/v1/modelpkgs", GetController().PostModelPkg)
		v1.PATCH("/v1/modelpkgs/:id", GetController().UpdateModelPkg)
		v1.PUT("/v1/modelpkgs/:id", GetController().UpdateModelPkg)
		v1.DELETE("/v1/modelpkgs/:id", GetController().DeleteModelPkg)

		v1.GET("/v1/pointertogongstructfields", GetController().GetPointerToGongStructFields)
		v1.GET("/v1/pointertogongstructfields/:id", GetController().GetPointerToGongStructField)
		v1.POST("/v1/pointertogongstructfields", GetController().PostPointerToGongStructField)
		v1.PATCH("/v1/pointertogongstructfields/:id", GetController().UpdatePointerToGongStructField)
		v1.PUT("/v1/pointertogongstructfields/:id", GetController().UpdatePointerToGongStructField)
		v1.DELETE("/v1/pointertogongstructfields/:id", GetController().DeletePointerToGongStructField)

		v1.GET("/v1/sliceofpointertogongstructfields", GetController().GetSliceOfPointerToGongStructFields)
		v1.GET("/v1/sliceofpointertogongstructfields/:id", GetController().GetSliceOfPointerToGongStructField)
		v1.POST("/v1/sliceofpointertogongstructfields", GetController().PostSliceOfPointerToGongStructField)
		v1.PATCH("/v1/sliceofpointertogongstructfields/:id", GetController().UpdateSliceOfPointerToGongStructField)
		v1.PUT("/v1/sliceofpointertogongstructfields/:id", GetController().UpdateSliceOfPointerToGongStructField)
		v1.DELETE("/v1/sliceofpointertogongstructfields/:id", GetController().DeleteSliceOfPointerToGongStructField)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)
	}
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gong/go, onWebSocketRequestForCommitFromBackNb")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gong/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/fullstack-lang/gong/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
	if err != nil {
		log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	}

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
		_ = nbCommitBackRepo

		backRepoData := new(orm.BackRepoData)
		orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

		// Send backRepo data
		err = wsConnection.WriteJSON(backRepoData)

		// log.Println("Stack github.com/fullstack-lang/gong/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gong/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
