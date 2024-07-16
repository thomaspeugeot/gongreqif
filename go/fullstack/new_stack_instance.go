// do not modify, generated file
package fullstack

import (
	"github.com/thomaspeugeot/gongreqif/go/controllers"
	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/thomaspeugeot/gongreqif/ng-github.com-thomaspeugeot-gongreqif/projects"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.ALTERNATIVEID](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONBOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONDATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONINTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONREAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONSTRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEDEFINITIONXHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEBOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEDATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEINTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEREAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUESTRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTEVALUEXHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.CHILDREN](stage)
	models.SetOrchestratorOnAfterUpdate[models.CORECONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONBOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONDATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONINTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONREAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONSTRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPEDEFINITIONXHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.DEFAULTVALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.DEFINITION](stage)
	models.SetOrchestratorOnAfterUpdate[models.EDITABLEATTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.EMBEDDEDVALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ENUMVALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.PROPERTIES](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATIONGROUP](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATIONGROUPTYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQIF](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQIFCONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQIFHEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQIFTOOLEXTENSION](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQTYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SOURCE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SOURCESPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECATTRIBUTES](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECHIERARCHY](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATIONTYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFIEDVALUES](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECOBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECOBJECTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECOBJECTTYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECRELATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECRELATIONGROUPS](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECRELATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECRELATIONTYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECTYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.TARGET](stage)
	models.SetOrchestratorOnAfterUpdate[models.TARGETSPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.THEHEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.TOOLEXTENSIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.VALUES](stage)
	models.SetOrchestratorOnAfterUpdate[models.XHTMLCONTENT](stage)

	return
}
