// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomaspeugeot/gongreqif/go/orm"

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
	v1 := r.Group("/api/github.com/thomaspeugeot/gongreqif/go")
	{ // insertion point for registrations
		v1.GET("/v1/alternativeids", GetController().GetALTERNATIVEIDs)
		v1.GET("/v1/alternativeids/:id", GetController().GetALTERNATIVEID)
		v1.POST("/v1/alternativeids", GetController().PostALTERNATIVEID)
		v1.PATCH("/v1/alternativeids/:id", GetController().UpdateALTERNATIVEID)
		v1.PUT("/v1/alternativeids/:id", GetController().UpdateALTERNATIVEID)
		v1.DELETE("/v1/alternativeids/:id", GetController().DeleteALTERNATIVEID)

		v1.GET("/v1/attributedefinitionbooleans", GetController().GetATTRIBUTEDEFINITIONBOOLEANs)
		v1.GET("/v1/attributedefinitionbooleans/:id", GetController().GetATTRIBUTEDEFINITIONBOOLEAN)
		v1.POST("/v1/attributedefinitionbooleans", GetController().PostATTRIBUTEDEFINITIONBOOLEAN)
		v1.PATCH("/v1/attributedefinitionbooleans/:id", GetController().UpdateATTRIBUTEDEFINITIONBOOLEAN)
		v1.PUT("/v1/attributedefinitionbooleans/:id", GetController().UpdateATTRIBUTEDEFINITIONBOOLEAN)
		v1.DELETE("/v1/attributedefinitionbooleans/:id", GetController().DeleteATTRIBUTEDEFINITIONBOOLEAN)

		v1.GET("/v1/attributedefinitiondates", GetController().GetATTRIBUTEDEFINITIONDATEs)
		v1.GET("/v1/attributedefinitiondates/:id", GetController().GetATTRIBUTEDEFINITIONDATE)
		v1.POST("/v1/attributedefinitiondates", GetController().PostATTRIBUTEDEFINITIONDATE)
		v1.PATCH("/v1/attributedefinitiondates/:id", GetController().UpdateATTRIBUTEDEFINITIONDATE)
		v1.PUT("/v1/attributedefinitiondates/:id", GetController().UpdateATTRIBUTEDEFINITIONDATE)
		v1.DELETE("/v1/attributedefinitiondates/:id", GetController().DeleteATTRIBUTEDEFINITIONDATE)

		v1.GET("/v1/attributedefinitionenumerations", GetController().GetATTRIBUTEDEFINITIONENUMERATIONs)
		v1.GET("/v1/attributedefinitionenumerations/:id", GetController().GetATTRIBUTEDEFINITIONENUMERATION)
		v1.POST("/v1/attributedefinitionenumerations", GetController().PostATTRIBUTEDEFINITIONENUMERATION)
		v1.PATCH("/v1/attributedefinitionenumerations/:id", GetController().UpdateATTRIBUTEDEFINITIONENUMERATION)
		v1.PUT("/v1/attributedefinitionenumerations/:id", GetController().UpdateATTRIBUTEDEFINITIONENUMERATION)
		v1.DELETE("/v1/attributedefinitionenumerations/:id", GetController().DeleteATTRIBUTEDEFINITIONENUMERATION)

		v1.GET("/v1/attributedefinitionintegers", GetController().GetATTRIBUTEDEFINITIONINTEGERs)
		v1.GET("/v1/attributedefinitionintegers/:id", GetController().GetATTRIBUTEDEFINITIONINTEGER)
		v1.POST("/v1/attributedefinitionintegers", GetController().PostATTRIBUTEDEFINITIONINTEGER)
		v1.PATCH("/v1/attributedefinitionintegers/:id", GetController().UpdateATTRIBUTEDEFINITIONINTEGER)
		v1.PUT("/v1/attributedefinitionintegers/:id", GetController().UpdateATTRIBUTEDEFINITIONINTEGER)
		v1.DELETE("/v1/attributedefinitionintegers/:id", GetController().DeleteATTRIBUTEDEFINITIONINTEGER)

		v1.GET("/v1/attributedefinitionreals", GetController().GetATTRIBUTEDEFINITIONREALs)
		v1.GET("/v1/attributedefinitionreals/:id", GetController().GetATTRIBUTEDEFINITIONREAL)
		v1.POST("/v1/attributedefinitionreals", GetController().PostATTRIBUTEDEFINITIONREAL)
		v1.PATCH("/v1/attributedefinitionreals/:id", GetController().UpdateATTRIBUTEDEFINITIONREAL)
		v1.PUT("/v1/attributedefinitionreals/:id", GetController().UpdateATTRIBUTEDEFINITIONREAL)
		v1.DELETE("/v1/attributedefinitionreals/:id", GetController().DeleteATTRIBUTEDEFINITIONREAL)

		v1.GET("/v1/attributedefinitionstrings", GetController().GetATTRIBUTEDEFINITIONSTRINGs)
		v1.GET("/v1/attributedefinitionstrings/:id", GetController().GetATTRIBUTEDEFINITIONSTRING)
		v1.POST("/v1/attributedefinitionstrings", GetController().PostATTRIBUTEDEFINITIONSTRING)
		v1.PATCH("/v1/attributedefinitionstrings/:id", GetController().UpdateATTRIBUTEDEFINITIONSTRING)
		v1.PUT("/v1/attributedefinitionstrings/:id", GetController().UpdateATTRIBUTEDEFINITIONSTRING)
		v1.DELETE("/v1/attributedefinitionstrings/:id", GetController().DeleteATTRIBUTEDEFINITIONSTRING)

		v1.GET("/v1/attributedefinitionxhtmls", GetController().GetATTRIBUTEDEFINITIONXHTMLs)
		v1.GET("/v1/attributedefinitionxhtmls/:id", GetController().GetATTRIBUTEDEFINITIONXHTML)
		v1.POST("/v1/attributedefinitionxhtmls", GetController().PostATTRIBUTEDEFINITIONXHTML)
		v1.PATCH("/v1/attributedefinitionxhtmls/:id", GetController().UpdateATTRIBUTEDEFINITIONXHTML)
		v1.PUT("/v1/attributedefinitionxhtmls/:id", GetController().UpdateATTRIBUTEDEFINITIONXHTML)
		v1.DELETE("/v1/attributedefinitionxhtmls/:id", GetController().DeleteATTRIBUTEDEFINITIONXHTML)

		v1.GET("/v1/attributevaluebooleans", GetController().GetATTRIBUTEVALUEBOOLEANs)
		v1.GET("/v1/attributevaluebooleans/:id", GetController().GetATTRIBUTEVALUEBOOLEAN)
		v1.POST("/v1/attributevaluebooleans", GetController().PostATTRIBUTEVALUEBOOLEAN)
		v1.PATCH("/v1/attributevaluebooleans/:id", GetController().UpdateATTRIBUTEVALUEBOOLEAN)
		v1.PUT("/v1/attributevaluebooleans/:id", GetController().UpdateATTRIBUTEVALUEBOOLEAN)
		v1.DELETE("/v1/attributevaluebooleans/:id", GetController().DeleteATTRIBUTEVALUEBOOLEAN)

		v1.GET("/v1/attributevaluedates", GetController().GetATTRIBUTEVALUEDATEs)
		v1.GET("/v1/attributevaluedates/:id", GetController().GetATTRIBUTEVALUEDATE)
		v1.POST("/v1/attributevaluedates", GetController().PostATTRIBUTEVALUEDATE)
		v1.PATCH("/v1/attributevaluedates/:id", GetController().UpdateATTRIBUTEVALUEDATE)
		v1.PUT("/v1/attributevaluedates/:id", GetController().UpdateATTRIBUTEVALUEDATE)
		v1.DELETE("/v1/attributevaluedates/:id", GetController().DeleteATTRIBUTEVALUEDATE)

		v1.GET("/v1/attributevalueenumerations", GetController().GetATTRIBUTEVALUEENUMERATIONs)
		v1.GET("/v1/attributevalueenumerations/:id", GetController().GetATTRIBUTEVALUEENUMERATION)
		v1.POST("/v1/attributevalueenumerations", GetController().PostATTRIBUTEVALUEENUMERATION)
		v1.PATCH("/v1/attributevalueenumerations/:id", GetController().UpdateATTRIBUTEVALUEENUMERATION)
		v1.PUT("/v1/attributevalueenumerations/:id", GetController().UpdateATTRIBUTEVALUEENUMERATION)
		v1.DELETE("/v1/attributevalueenumerations/:id", GetController().DeleteATTRIBUTEVALUEENUMERATION)

		v1.GET("/v1/attributevalueintegers", GetController().GetATTRIBUTEVALUEINTEGERs)
		v1.GET("/v1/attributevalueintegers/:id", GetController().GetATTRIBUTEVALUEINTEGER)
		v1.POST("/v1/attributevalueintegers", GetController().PostATTRIBUTEVALUEINTEGER)
		v1.PATCH("/v1/attributevalueintegers/:id", GetController().UpdateATTRIBUTEVALUEINTEGER)
		v1.PUT("/v1/attributevalueintegers/:id", GetController().UpdateATTRIBUTEVALUEINTEGER)
		v1.DELETE("/v1/attributevalueintegers/:id", GetController().DeleteATTRIBUTEVALUEINTEGER)

		v1.GET("/v1/attributevaluereals", GetController().GetATTRIBUTEVALUEREALs)
		v1.GET("/v1/attributevaluereals/:id", GetController().GetATTRIBUTEVALUEREAL)
		v1.POST("/v1/attributevaluereals", GetController().PostATTRIBUTEVALUEREAL)
		v1.PATCH("/v1/attributevaluereals/:id", GetController().UpdateATTRIBUTEVALUEREAL)
		v1.PUT("/v1/attributevaluereals/:id", GetController().UpdateATTRIBUTEVALUEREAL)
		v1.DELETE("/v1/attributevaluereals/:id", GetController().DeleteATTRIBUTEVALUEREAL)

		v1.GET("/v1/attributevaluestrings", GetController().GetATTRIBUTEVALUESTRINGs)
		v1.GET("/v1/attributevaluestrings/:id", GetController().GetATTRIBUTEVALUESTRING)
		v1.POST("/v1/attributevaluestrings", GetController().PostATTRIBUTEVALUESTRING)
		v1.PATCH("/v1/attributevaluestrings/:id", GetController().UpdateATTRIBUTEVALUESTRING)
		v1.PUT("/v1/attributevaluestrings/:id", GetController().UpdateATTRIBUTEVALUESTRING)
		v1.DELETE("/v1/attributevaluestrings/:id", GetController().DeleteATTRIBUTEVALUESTRING)

		v1.GET("/v1/attributevaluexhtmls", GetController().GetATTRIBUTEVALUEXHTMLs)
		v1.GET("/v1/attributevaluexhtmls/:id", GetController().GetATTRIBUTEVALUEXHTML)
		v1.POST("/v1/attributevaluexhtmls", GetController().PostATTRIBUTEVALUEXHTML)
		v1.PATCH("/v1/attributevaluexhtmls/:id", GetController().UpdateATTRIBUTEVALUEXHTML)
		v1.PUT("/v1/attributevaluexhtmls/:id", GetController().UpdateATTRIBUTEVALUEXHTML)
		v1.DELETE("/v1/attributevaluexhtmls/:id", GetController().DeleteATTRIBUTEVALUEXHTML)

		v1.GET("/v1/childrens", GetController().GetCHILDRENs)
		v1.GET("/v1/childrens/:id", GetController().GetCHILDREN)
		v1.POST("/v1/childrens", GetController().PostCHILDREN)
		v1.PATCH("/v1/childrens/:id", GetController().UpdateCHILDREN)
		v1.PUT("/v1/childrens/:id", GetController().UpdateCHILDREN)
		v1.DELETE("/v1/childrens/:id", GetController().DeleteCHILDREN)

		v1.GET("/v1/corecontents", GetController().GetCORECONTENTs)
		v1.GET("/v1/corecontents/:id", GetController().GetCORECONTENT)
		v1.POST("/v1/corecontents", GetController().PostCORECONTENT)
		v1.PATCH("/v1/corecontents/:id", GetController().UpdateCORECONTENT)
		v1.PUT("/v1/corecontents/:id", GetController().UpdateCORECONTENT)
		v1.DELETE("/v1/corecontents/:id", GetController().DeleteCORECONTENT)

		v1.GET("/v1/datatypedefinitionbooleans", GetController().GetDATATYPEDEFINITIONBOOLEANs)
		v1.GET("/v1/datatypedefinitionbooleans/:id", GetController().GetDATATYPEDEFINITIONBOOLEAN)
		v1.POST("/v1/datatypedefinitionbooleans", GetController().PostDATATYPEDEFINITIONBOOLEAN)
		v1.PATCH("/v1/datatypedefinitionbooleans/:id", GetController().UpdateDATATYPEDEFINITIONBOOLEAN)
		v1.PUT("/v1/datatypedefinitionbooleans/:id", GetController().UpdateDATATYPEDEFINITIONBOOLEAN)
		v1.DELETE("/v1/datatypedefinitionbooleans/:id", GetController().DeleteDATATYPEDEFINITIONBOOLEAN)

		v1.GET("/v1/datatypedefinitiondates", GetController().GetDATATYPEDEFINITIONDATEs)
		v1.GET("/v1/datatypedefinitiondates/:id", GetController().GetDATATYPEDEFINITIONDATE)
		v1.POST("/v1/datatypedefinitiondates", GetController().PostDATATYPEDEFINITIONDATE)
		v1.PATCH("/v1/datatypedefinitiondates/:id", GetController().UpdateDATATYPEDEFINITIONDATE)
		v1.PUT("/v1/datatypedefinitiondates/:id", GetController().UpdateDATATYPEDEFINITIONDATE)
		v1.DELETE("/v1/datatypedefinitiondates/:id", GetController().DeleteDATATYPEDEFINITIONDATE)

		v1.GET("/v1/datatypedefinitionenumerations", GetController().GetDATATYPEDEFINITIONENUMERATIONs)
		v1.GET("/v1/datatypedefinitionenumerations/:id", GetController().GetDATATYPEDEFINITIONENUMERATION)
		v1.POST("/v1/datatypedefinitionenumerations", GetController().PostDATATYPEDEFINITIONENUMERATION)
		v1.PATCH("/v1/datatypedefinitionenumerations/:id", GetController().UpdateDATATYPEDEFINITIONENUMERATION)
		v1.PUT("/v1/datatypedefinitionenumerations/:id", GetController().UpdateDATATYPEDEFINITIONENUMERATION)
		v1.DELETE("/v1/datatypedefinitionenumerations/:id", GetController().DeleteDATATYPEDEFINITIONENUMERATION)

		v1.GET("/v1/datatypedefinitionintegers", GetController().GetDATATYPEDEFINITIONINTEGERs)
		v1.GET("/v1/datatypedefinitionintegers/:id", GetController().GetDATATYPEDEFINITIONINTEGER)
		v1.POST("/v1/datatypedefinitionintegers", GetController().PostDATATYPEDEFINITIONINTEGER)
		v1.PATCH("/v1/datatypedefinitionintegers/:id", GetController().UpdateDATATYPEDEFINITIONINTEGER)
		v1.PUT("/v1/datatypedefinitionintegers/:id", GetController().UpdateDATATYPEDEFINITIONINTEGER)
		v1.DELETE("/v1/datatypedefinitionintegers/:id", GetController().DeleteDATATYPEDEFINITIONINTEGER)

		v1.GET("/v1/datatypedefinitionreals", GetController().GetDATATYPEDEFINITIONREALs)
		v1.GET("/v1/datatypedefinitionreals/:id", GetController().GetDATATYPEDEFINITIONREAL)
		v1.POST("/v1/datatypedefinitionreals", GetController().PostDATATYPEDEFINITIONREAL)
		v1.PATCH("/v1/datatypedefinitionreals/:id", GetController().UpdateDATATYPEDEFINITIONREAL)
		v1.PUT("/v1/datatypedefinitionreals/:id", GetController().UpdateDATATYPEDEFINITIONREAL)
		v1.DELETE("/v1/datatypedefinitionreals/:id", GetController().DeleteDATATYPEDEFINITIONREAL)

		v1.GET("/v1/datatypedefinitionstrings", GetController().GetDATATYPEDEFINITIONSTRINGs)
		v1.GET("/v1/datatypedefinitionstrings/:id", GetController().GetDATATYPEDEFINITIONSTRING)
		v1.POST("/v1/datatypedefinitionstrings", GetController().PostDATATYPEDEFINITIONSTRING)
		v1.PATCH("/v1/datatypedefinitionstrings/:id", GetController().UpdateDATATYPEDEFINITIONSTRING)
		v1.PUT("/v1/datatypedefinitionstrings/:id", GetController().UpdateDATATYPEDEFINITIONSTRING)
		v1.DELETE("/v1/datatypedefinitionstrings/:id", GetController().DeleteDATATYPEDEFINITIONSTRING)

		v1.GET("/v1/datatypedefinitionxhtmls", GetController().GetDATATYPEDEFINITIONXHTMLs)
		v1.GET("/v1/datatypedefinitionxhtmls/:id", GetController().GetDATATYPEDEFINITIONXHTML)
		v1.POST("/v1/datatypedefinitionxhtmls", GetController().PostDATATYPEDEFINITIONXHTML)
		v1.PATCH("/v1/datatypedefinitionxhtmls/:id", GetController().UpdateDATATYPEDEFINITIONXHTML)
		v1.PUT("/v1/datatypedefinitionxhtmls/:id", GetController().UpdateDATATYPEDEFINITIONXHTML)
		v1.DELETE("/v1/datatypedefinitionxhtmls/:id", GetController().DeleteDATATYPEDEFINITIONXHTML)

		v1.GET("/v1/datatypess", GetController().GetDATATYPESs)
		v1.GET("/v1/datatypess/:id", GetController().GetDATATYPES)
		v1.POST("/v1/datatypess", GetController().PostDATATYPES)
		v1.PATCH("/v1/datatypess/:id", GetController().UpdateDATATYPES)
		v1.PUT("/v1/datatypess/:id", GetController().UpdateDATATYPES)
		v1.DELETE("/v1/datatypess/:id", GetController().DeleteDATATYPES)

		v1.GET("/v1/defaultvalues", GetController().GetDEFAULTVALUEs)
		v1.GET("/v1/defaultvalues/:id", GetController().GetDEFAULTVALUE)
		v1.POST("/v1/defaultvalues", GetController().PostDEFAULTVALUE)
		v1.PATCH("/v1/defaultvalues/:id", GetController().UpdateDEFAULTVALUE)
		v1.PUT("/v1/defaultvalues/:id", GetController().UpdateDEFAULTVALUE)
		v1.DELETE("/v1/defaultvalues/:id", GetController().DeleteDEFAULTVALUE)

		v1.GET("/v1/definitions", GetController().GetDEFINITIONs)
		v1.GET("/v1/definitions/:id", GetController().GetDEFINITION)
		v1.POST("/v1/definitions", GetController().PostDEFINITION)
		v1.PATCH("/v1/definitions/:id", GetController().UpdateDEFINITION)
		v1.PUT("/v1/definitions/:id", GetController().UpdateDEFINITION)
		v1.DELETE("/v1/definitions/:id", GetController().DeleteDEFINITION)

		v1.GET("/v1/editableattss", GetController().GetEDITABLEATTSs)
		v1.GET("/v1/editableattss/:id", GetController().GetEDITABLEATTS)
		v1.POST("/v1/editableattss", GetController().PostEDITABLEATTS)
		v1.PATCH("/v1/editableattss/:id", GetController().UpdateEDITABLEATTS)
		v1.PUT("/v1/editableattss/:id", GetController().UpdateEDITABLEATTS)
		v1.DELETE("/v1/editableattss/:id", GetController().DeleteEDITABLEATTS)

		v1.GET("/v1/embeddedvalues", GetController().GetEMBEDDEDVALUEs)
		v1.GET("/v1/embeddedvalues/:id", GetController().GetEMBEDDEDVALUE)
		v1.POST("/v1/embeddedvalues", GetController().PostEMBEDDEDVALUE)
		v1.PATCH("/v1/embeddedvalues/:id", GetController().UpdateEMBEDDEDVALUE)
		v1.PUT("/v1/embeddedvalues/:id", GetController().UpdateEMBEDDEDVALUE)
		v1.DELETE("/v1/embeddedvalues/:id", GetController().DeleteEMBEDDEDVALUE)

		v1.GET("/v1/enumvalues", GetController().GetENUMVALUEs)
		v1.GET("/v1/enumvalues/:id", GetController().GetENUMVALUE)
		v1.POST("/v1/enumvalues", GetController().PostENUMVALUE)
		v1.PATCH("/v1/enumvalues/:id", GetController().UpdateENUMVALUE)
		v1.PUT("/v1/enumvalues/:id", GetController().UpdateENUMVALUE)
		v1.DELETE("/v1/enumvalues/:id", GetController().DeleteENUMVALUE)

		v1.GET("/v1/objects", GetController().GetOBJECTs)
		v1.GET("/v1/objects/:id", GetController().GetOBJECT)
		v1.POST("/v1/objects", GetController().PostOBJECT)
		v1.PATCH("/v1/objects/:id", GetController().UpdateOBJECT)
		v1.PUT("/v1/objects/:id", GetController().UpdateOBJECT)
		v1.DELETE("/v1/objects/:id", GetController().DeleteOBJECT)

		v1.GET("/v1/propertiess", GetController().GetPROPERTIESs)
		v1.GET("/v1/propertiess/:id", GetController().GetPROPERTIES)
		v1.POST("/v1/propertiess", GetController().PostPROPERTIES)
		v1.PATCH("/v1/propertiess/:id", GetController().UpdatePROPERTIES)
		v1.PUT("/v1/propertiess/:id", GetController().UpdatePROPERTIES)
		v1.DELETE("/v1/propertiess/:id", GetController().DeletePROPERTIES)

		v1.GET("/v1/relationgroups", GetController().GetRELATIONGROUPs)
		v1.GET("/v1/relationgroups/:id", GetController().GetRELATIONGROUP)
		v1.POST("/v1/relationgroups", GetController().PostRELATIONGROUP)
		v1.PATCH("/v1/relationgroups/:id", GetController().UpdateRELATIONGROUP)
		v1.PUT("/v1/relationgroups/:id", GetController().UpdateRELATIONGROUP)
		v1.DELETE("/v1/relationgroups/:id", GetController().DeleteRELATIONGROUP)

		v1.GET("/v1/relationgrouptypes", GetController().GetRELATIONGROUPTYPEs)
		v1.GET("/v1/relationgrouptypes/:id", GetController().GetRELATIONGROUPTYPE)
		v1.POST("/v1/relationgrouptypes", GetController().PostRELATIONGROUPTYPE)
		v1.PATCH("/v1/relationgrouptypes/:id", GetController().UpdateRELATIONGROUPTYPE)
		v1.PUT("/v1/relationgrouptypes/:id", GetController().UpdateRELATIONGROUPTYPE)
		v1.DELETE("/v1/relationgrouptypes/:id", GetController().DeleteRELATIONGROUPTYPE)

		v1.GET("/v1/reqifs", GetController().GetREQIFs)
		v1.GET("/v1/reqifs/:id", GetController().GetREQIF)
		v1.POST("/v1/reqifs", GetController().PostREQIF)
		v1.PATCH("/v1/reqifs/:id", GetController().UpdateREQIF)
		v1.PUT("/v1/reqifs/:id", GetController().UpdateREQIF)
		v1.DELETE("/v1/reqifs/:id", GetController().DeleteREQIF)

		v1.GET("/v1/reqifcontents", GetController().GetREQIFCONTENTs)
		v1.GET("/v1/reqifcontents/:id", GetController().GetREQIFCONTENT)
		v1.POST("/v1/reqifcontents", GetController().PostREQIFCONTENT)
		v1.PATCH("/v1/reqifcontents/:id", GetController().UpdateREQIFCONTENT)
		v1.PUT("/v1/reqifcontents/:id", GetController().UpdateREQIFCONTENT)
		v1.DELETE("/v1/reqifcontents/:id", GetController().DeleteREQIFCONTENT)

		v1.GET("/v1/reqifheaders", GetController().GetREQIFHEADERs)
		v1.GET("/v1/reqifheaders/:id", GetController().GetREQIFHEADER)
		v1.POST("/v1/reqifheaders", GetController().PostREQIFHEADER)
		v1.PATCH("/v1/reqifheaders/:id", GetController().UpdateREQIFHEADER)
		v1.PUT("/v1/reqifheaders/:id", GetController().UpdateREQIFHEADER)
		v1.DELETE("/v1/reqifheaders/:id", GetController().DeleteREQIFHEADER)

		v1.GET("/v1/reqiftoolextensions", GetController().GetREQIFTOOLEXTENSIONs)
		v1.GET("/v1/reqiftoolextensions/:id", GetController().GetREQIFTOOLEXTENSION)
		v1.POST("/v1/reqiftoolextensions", GetController().PostREQIFTOOLEXTENSION)
		v1.PATCH("/v1/reqiftoolextensions/:id", GetController().UpdateREQIFTOOLEXTENSION)
		v1.PUT("/v1/reqiftoolextensions/:id", GetController().UpdateREQIFTOOLEXTENSION)
		v1.DELETE("/v1/reqiftoolextensions/:id", GetController().DeleteREQIFTOOLEXTENSION)

		v1.GET("/v1/reqiftypes", GetController().GetREQIFTYPEs)
		v1.GET("/v1/reqiftypes/:id", GetController().GetREQIFTYPE)
		v1.POST("/v1/reqiftypes", GetController().PostREQIFTYPE)
		v1.PATCH("/v1/reqiftypes/:id", GetController().UpdateREQIFTYPE)
		v1.PUT("/v1/reqiftypes/:id", GetController().UpdateREQIFTYPE)
		v1.DELETE("/v1/reqiftypes/:id", GetController().DeleteREQIFTYPE)

		v1.GET("/v1/sources", GetController().GetSOURCEs)
		v1.GET("/v1/sources/:id", GetController().GetSOURCE)
		v1.POST("/v1/sources", GetController().PostSOURCE)
		v1.PATCH("/v1/sources/:id", GetController().UpdateSOURCE)
		v1.PUT("/v1/sources/:id", GetController().UpdateSOURCE)
		v1.DELETE("/v1/sources/:id", GetController().DeleteSOURCE)

		v1.GET("/v1/sourcespecifications", GetController().GetSOURCESPECIFICATIONs)
		v1.GET("/v1/sourcespecifications/:id", GetController().GetSOURCESPECIFICATION)
		v1.POST("/v1/sourcespecifications", GetController().PostSOURCESPECIFICATION)
		v1.PATCH("/v1/sourcespecifications/:id", GetController().UpdateSOURCESPECIFICATION)
		v1.PUT("/v1/sourcespecifications/:id", GetController().UpdateSOURCESPECIFICATION)
		v1.DELETE("/v1/sourcespecifications/:id", GetController().DeleteSOURCESPECIFICATION)

		v1.GET("/v1/specattributess", GetController().GetSPECATTRIBUTESs)
		v1.GET("/v1/specattributess/:id", GetController().GetSPECATTRIBUTES)
		v1.POST("/v1/specattributess", GetController().PostSPECATTRIBUTES)
		v1.PATCH("/v1/specattributess/:id", GetController().UpdateSPECATTRIBUTES)
		v1.PUT("/v1/specattributess/:id", GetController().UpdateSPECATTRIBUTES)
		v1.DELETE("/v1/specattributess/:id", GetController().DeleteSPECATTRIBUTES)

		v1.GET("/v1/spechierarchys", GetController().GetSPECHIERARCHYs)
		v1.GET("/v1/spechierarchys/:id", GetController().GetSPECHIERARCHY)
		v1.POST("/v1/spechierarchys", GetController().PostSPECHIERARCHY)
		v1.PATCH("/v1/spechierarchys/:id", GetController().UpdateSPECHIERARCHY)
		v1.PUT("/v1/spechierarchys/:id", GetController().UpdateSPECHIERARCHY)
		v1.DELETE("/v1/spechierarchys/:id", GetController().DeleteSPECHIERARCHY)

		v1.GET("/v1/specifications", GetController().GetSPECIFICATIONs)
		v1.GET("/v1/specifications/:id", GetController().GetSPECIFICATION)
		v1.POST("/v1/specifications", GetController().PostSPECIFICATION)
		v1.PATCH("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.PUT("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.DELETE("/v1/specifications/:id", GetController().DeleteSPECIFICATION)

		v1.GET("/v1/specificationss", GetController().GetSPECIFICATIONSs)
		v1.GET("/v1/specificationss/:id", GetController().GetSPECIFICATIONS)
		v1.POST("/v1/specificationss", GetController().PostSPECIFICATIONS)
		v1.PATCH("/v1/specificationss/:id", GetController().UpdateSPECIFICATIONS)
		v1.PUT("/v1/specificationss/:id", GetController().UpdateSPECIFICATIONS)
		v1.DELETE("/v1/specificationss/:id", GetController().DeleteSPECIFICATIONS)

		v1.GET("/v1/specificationtypes", GetController().GetSPECIFICATIONTYPEs)
		v1.GET("/v1/specificationtypes/:id", GetController().GetSPECIFICATIONTYPE)
		v1.POST("/v1/specificationtypes", GetController().PostSPECIFICATIONTYPE)
		v1.PATCH("/v1/specificationtypes/:id", GetController().UpdateSPECIFICATIONTYPE)
		v1.PUT("/v1/specificationtypes/:id", GetController().UpdateSPECIFICATIONTYPE)
		v1.DELETE("/v1/specificationtypes/:id", GetController().DeleteSPECIFICATIONTYPE)

		v1.GET("/v1/specifiedvaluess", GetController().GetSPECIFIEDVALUESs)
		v1.GET("/v1/specifiedvaluess/:id", GetController().GetSPECIFIEDVALUES)
		v1.POST("/v1/specifiedvaluess", GetController().PostSPECIFIEDVALUES)
		v1.PATCH("/v1/specifiedvaluess/:id", GetController().UpdateSPECIFIEDVALUES)
		v1.PUT("/v1/specifiedvaluess/:id", GetController().UpdateSPECIFIEDVALUES)
		v1.DELETE("/v1/specifiedvaluess/:id", GetController().DeleteSPECIFIEDVALUES)

		v1.GET("/v1/specobjects", GetController().GetSPECOBJECTs)
		v1.GET("/v1/specobjects/:id", GetController().GetSPECOBJECT)
		v1.POST("/v1/specobjects", GetController().PostSPECOBJECT)
		v1.PATCH("/v1/specobjects/:id", GetController().UpdateSPECOBJECT)
		v1.PUT("/v1/specobjects/:id", GetController().UpdateSPECOBJECT)
		v1.DELETE("/v1/specobjects/:id", GetController().DeleteSPECOBJECT)

		v1.GET("/v1/specobjectss", GetController().GetSPECOBJECTSs)
		v1.GET("/v1/specobjectss/:id", GetController().GetSPECOBJECTS)
		v1.POST("/v1/specobjectss", GetController().PostSPECOBJECTS)
		v1.PATCH("/v1/specobjectss/:id", GetController().UpdateSPECOBJECTS)
		v1.PUT("/v1/specobjectss/:id", GetController().UpdateSPECOBJECTS)
		v1.DELETE("/v1/specobjectss/:id", GetController().DeleteSPECOBJECTS)

		v1.GET("/v1/specobjecttypes", GetController().GetSPECOBJECTTYPEs)
		v1.GET("/v1/specobjecttypes/:id", GetController().GetSPECOBJECTTYPE)
		v1.POST("/v1/specobjecttypes", GetController().PostSPECOBJECTTYPE)
		v1.PATCH("/v1/specobjecttypes/:id", GetController().UpdateSPECOBJECTTYPE)
		v1.PUT("/v1/specobjecttypes/:id", GetController().UpdateSPECOBJECTTYPE)
		v1.DELETE("/v1/specobjecttypes/:id", GetController().DeleteSPECOBJECTTYPE)

		v1.GET("/v1/specrelations", GetController().GetSPECRELATIONs)
		v1.GET("/v1/specrelations/:id", GetController().GetSPECRELATION)
		v1.POST("/v1/specrelations", GetController().PostSPECRELATION)
		v1.PATCH("/v1/specrelations/:id", GetController().UpdateSPECRELATION)
		v1.PUT("/v1/specrelations/:id", GetController().UpdateSPECRELATION)
		v1.DELETE("/v1/specrelations/:id", GetController().DeleteSPECRELATION)

		v1.GET("/v1/specrelationgroupss", GetController().GetSPECRELATIONGROUPSs)
		v1.GET("/v1/specrelationgroupss/:id", GetController().GetSPECRELATIONGROUPS)
		v1.POST("/v1/specrelationgroupss", GetController().PostSPECRELATIONGROUPS)
		v1.PATCH("/v1/specrelationgroupss/:id", GetController().UpdateSPECRELATIONGROUPS)
		v1.PUT("/v1/specrelationgroupss/:id", GetController().UpdateSPECRELATIONGROUPS)
		v1.DELETE("/v1/specrelationgroupss/:id", GetController().DeleteSPECRELATIONGROUPS)

		v1.GET("/v1/specrelationss", GetController().GetSPECRELATIONSs)
		v1.GET("/v1/specrelationss/:id", GetController().GetSPECRELATIONS)
		v1.POST("/v1/specrelationss", GetController().PostSPECRELATIONS)
		v1.PATCH("/v1/specrelationss/:id", GetController().UpdateSPECRELATIONS)
		v1.PUT("/v1/specrelationss/:id", GetController().UpdateSPECRELATIONS)
		v1.DELETE("/v1/specrelationss/:id", GetController().DeleteSPECRELATIONS)

		v1.GET("/v1/specrelationtypes", GetController().GetSPECRELATIONTYPEs)
		v1.GET("/v1/specrelationtypes/:id", GetController().GetSPECRELATIONTYPE)
		v1.POST("/v1/specrelationtypes", GetController().PostSPECRELATIONTYPE)
		v1.PATCH("/v1/specrelationtypes/:id", GetController().UpdateSPECRELATIONTYPE)
		v1.PUT("/v1/specrelationtypes/:id", GetController().UpdateSPECRELATIONTYPE)
		v1.DELETE("/v1/specrelationtypes/:id", GetController().DeleteSPECRELATIONTYPE)

		v1.GET("/v1/spectypess", GetController().GetSPECTYPESs)
		v1.GET("/v1/spectypess/:id", GetController().GetSPECTYPES)
		v1.POST("/v1/spectypess", GetController().PostSPECTYPES)
		v1.PATCH("/v1/spectypess/:id", GetController().UpdateSPECTYPES)
		v1.PUT("/v1/spectypess/:id", GetController().UpdateSPECTYPES)
		v1.DELETE("/v1/spectypess/:id", GetController().DeleteSPECTYPES)

		v1.GET("/v1/targets", GetController().GetTARGETs)
		v1.GET("/v1/targets/:id", GetController().GetTARGET)
		v1.POST("/v1/targets", GetController().PostTARGET)
		v1.PATCH("/v1/targets/:id", GetController().UpdateTARGET)
		v1.PUT("/v1/targets/:id", GetController().UpdateTARGET)
		v1.DELETE("/v1/targets/:id", GetController().DeleteTARGET)

		v1.GET("/v1/targetspecifications", GetController().GetTARGETSPECIFICATIONs)
		v1.GET("/v1/targetspecifications/:id", GetController().GetTARGETSPECIFICATION)
		v1.POST("/v1/targetspecifications", GetController().PostTARGETSPECIFICATION)
		v1.PATCH("/v1/targetspecifications/:id", GetController().UpdateTARGETSPECIFICATION)
		v1.PUT("/v1/targetspecifications/:id", GetController().UpdateTARGETSPECIFICATION)
		v1.DELETE("/v1/targetspecifications/:id", GetController().DeleteTARGETSPECIFICATION)

		v1.GET("/v1/theheaders", GetController().GetTHEHEADERs)
		v1.GET("/v1/theheaders/:id", GetController().GetTHEHEADER)
		v1.POST("/v1/theheaders", GetController().PostTHEHEADER)
		v1.PATCH("/v1/theheaders/:id", GetController().UpdateTHEHEADER)
		v1.PUT("/v1/theheaders/:id", GetController().UpdateTHEHEADER)
		v1.DELETE("/v1/theheaders/:id", GetController().DeleteTHEHEADER)

		v1.GET("/v1/toolextensionss", GetController().GetTOOLEXTENSIONSs)
		v1.GET("/v1/toolextensionss/:id", GetController().GetTOOLEXTENSIONS)
		v1.POST("/v1/toolextensionss", GetController().PostTOOLEXTENSIONS)
		v1.PATCH("/v1/toolextensionss/:id", GetController().UpdateTOOLEXTENSIONS)
		v1.PUT("/v1/toolextensionss/:id", GetController().UpdateTOOLEXTENSIONS)
		v1.DELETE("/v1/toolextensionss/:id", GetController().DeleteTOOLEXTENSIONS)

		v1.GET("/v1/valuess", GetController().GetVALUESs)
		v1.GET("/v1/valuess/:id", GetController().GetVALUES)
		v1.POST("/v1/valuess", GetController().PostVALUES)
		v1.PATCH("/v1/valuess/:id", GetController().UpdateVALUES)
		v1.PUT("/v1/valuess/:id", GetController().UpdateVALUES)
		v1.DELETE("/v1/valuess/:id", GetController().DeleteVALUES)

		v1.GET("/v1/xhtmlcontents", GetController().GetXHTMLCONTENTs)
		v1.GET("/v1/xhtmlcontents/:id", GetController().GetXHTMLCONTENT)
		v1.POST("/v1/xhtmlcontents", GetController().PostXHTMLCONTENT)
		v1.PATCH("/v1/xhtmlcontents/:id", GetController().UpdateXHTMLCONTENT)
		v1.PUT("/v1/xhtmlcontents/:id", GetController().UpdateXHTMLCONTENT)
		v1.DELETE("/v1/xhtmlcontents/:id", GetController().DeleteXHTMLCONTENT)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k, _ := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/thomaspeugeot/gongreqif/go, onWebSocketRequestForCommitFromBackNb")

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
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go, Unkown stack", stackPath)
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

	// log.Println("Stack github.com/thomaspeugeot/gongreqif/go, onWebSocketRequestForBackRepoContent")

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
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/thomaspeugeot/gongreqif/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
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

		// log.Println("Stack github.com/thomaspeugeot/gongreqif/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

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
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
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
		log.Panic("Stack github.com/thomaspeugeot/gongreqif/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
