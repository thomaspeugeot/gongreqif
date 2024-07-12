// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "ALTERNATIVEID" {
		fillUpTable[models.ALTERNATIVEID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONBOOLEAN" {
		fillUpTable[models.ATTRIBUTEDEFINITIONBOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONDATE" {
		fillUpTable[models.ATTRIBUTEDEFINITIONDATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONENUMERATION" {
		fillUpTable[models.ATTRIBUTEDEFINITIONENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONINTEGER" {
		fillUpTable[models.ATTRIBUTEDEFINITIONINTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONREAL" {
		fillUpTable[models.ATTRIBUTEDEFINITIONREAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONSTRING" {
		fillUpTable[models.ATTRIBUTEDEFINITIONSTRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEDEFINITIONXHTML" {
		fillUpTable[models.ATTRIBUTEDEFINITIONXHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEBOOLEAN" {
		fillUpTable[models.ATTRIBUTEVALUEBOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEDATE" {
		fillUpTable[models.ATTRIBUTEVALUEDATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEENUMERATION" {
		fillUpTable[models.ATTRIBUTEVALUEENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEINTEGER" {
		fillUpTable[models.ATTRIBUTEVALUEINTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEREAL" {
		fillUpTable[models.ATTRIBUTEVALUEREAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUESTRING" {
		fillUpTable[models.ATTRIBUTEVALUESTRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTEVALUEXHTML" {
		fillUpTable[models.ATTRIBUTEVALUEXHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CHILDREN" {
		fillUpTable[models.CHILDREN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "CORECONTENT" {
		fillUpTable[models.CORECONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONBOOLEAN" {
		fillUpTable[models.DATATYPEDEFINITIONBOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONDATE" {
		fillUpTable[models.DATATYPEDEFINITIONDATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONENUMERATION" {
		fillUpTable[models.DATATYPEDEFINITIONENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONINTEGER" {
		fillUpTable[models.DATATYPEDEFINITIONINTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONREAL" {
		fillUpTable[models.DATATYPEDEFINITIONREAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONSTRING" {
		fillUpTable[models.DATATYPEDEFINITIONSTRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPEDEFINITIONXHTML" {
		fillUpTable[models.DATATYPEDEFINITIONXHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPES" {
		fillUpTable[models.DATATYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DEFAULTVALUE" {
		fillUpTable[models.DEFAULTVALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DEFINITION" {
		fillUpTable[models.DEFINITION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EDITABLEATTS" {
		fillUpTable[models.EDITABLEATTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EMBEDDEDVALUE" {
		fillUpTable[models.EMBEDDEDVALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ENUMVALUE" {
		fillUpTable[models.ENUMVALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "OBJECT" {
		fillUpTable[models.OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "PROPERTIES" {
		fillUpTable[models.PROPERTIES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATIONGROUP" {
		fillUpTable[models.RELATIONGROUP](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATIONGROUPTYPE" {
		fillUpTable[models.RELATIONGROUPTYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQIF" {
		fillUpTable[models.REQIF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQIFCONTENT" {
		fillUpTable[models.REQIFCONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQIFHEADER" {
		fillUpTable[models.REQIFHEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQIFTOOLEXTENSION" {
		fillUpTable[models.REQIFTOOLEXTENSION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQIFTYPE" {
		fillUpTable[models.REQIFTYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SOURCE" {
		fillUpTable[models.SOURCE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SOURCESPECIFICATION" {
		fillUpTable[models.SOURCESPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECATTRIBUTES" {
		fillUpTable[models.SPECATTRIBUTES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECHIERARCHY" {
		fillUpTable[models.SPECHIERARCHY](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION" {
		fillUpTable[models.SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATIONS" {
		fillUpTable[models.SPECIFICATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATIONTYPE" {
		fillUpTable[models.SPECIFICATIONTYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFIEDVALUES" {
		fillUpTable[models.SPECIFIEDVALUES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECOBJECT" {
		fillUpTable[models.SPECOBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECOBJECTS" {
		fillUpTable[models.SPECOBJECTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECOBJECTTYPE" {
		fillUpTable[models.SPECOBJECTTYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECRELATION" {
		fillUpTable[models.SPECRELATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECRELATIONGROUPS" {
		fillUpTable[models.SPECRELATIONGROUPS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECRELATIONS" {
		fillUpTable[models.SPECRELATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECRELATIONTYPE" {
		fillUpTable[models.SPECRELATIONTYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECTYPES" {
		fillUpTable[models.SPECTYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TARGET" {
		fillUpTable[models.TARGET](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TARGETSPECIFICATION" {
		fillUpTable[models.TARGETSPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "THEHEADER" {
		fillUpTable[models.THEHEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TOOLEXTENSIONS" {
		fillUpTable[models.TOOLEXTENSIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "VALUES" {
		fillUpTable[models.VALUES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "XHTMLCONTENT" {
		fillUpTable[models.XHTMLCONTENT](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
