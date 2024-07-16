// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.ALTERNATIVEID:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONDATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONREAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEDATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEINTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEREAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUESTRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTEVALUEXHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CHILDREN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CORECONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONDATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONINTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONREAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONSTRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPEDEFINITIONXHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DEFAULTVALUE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DEFINITION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.EDITABLEATTS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.EMBEDDEDVALUE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ENUMVALUE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.OBJECT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.PROPERTIES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RELATIONGROUP:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RELATIONGROUPTYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQIF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQIFCONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQIFHEADER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQIFTOOLEXTENSION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQTYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SOURCE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SOURCESPECIFICATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECATTRIBUTES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECHIERARCHY:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFICATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFICATIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFICATIONTYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFIEDVALUES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECOBJECT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECOBJECTS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECOBJECTTYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECRELATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECRELATIONGROUPS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECRELATIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECRELATIONTYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECTYPES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TARGET:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TARGETSPECIFICATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.THEHEADER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TOOLEXTENSIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.VALUES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.XHTMLCONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

