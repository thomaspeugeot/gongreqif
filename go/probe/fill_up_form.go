// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.REQIF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQ_IF_HEADER", instanceWithInferedType.REQ_IF_HEADER, formGroup, probe)
		AssociationFieldToForm("REQ_IF_CONTENT", instanceWithInferedType.REQ_IF_CONTENT, formGroup, probe)

	case *models.REQ_IF_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPEC_OBJECT_TYPES", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECT_TYPES, formGroup, probe)
		AssociationSliceToForm("SPECIFICATIONS", instanceWithInferedType, &instanceWithInferedType.SPECIFICATIONS, formGroup, probe)

	case *models.REQ_IF_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("COMMENT", instanceWithInferedType.COMMENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CREATION_TIME", instanceWithInferedType.CREATION_TIME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REPOSITORY_ID", instanceWithInferedType.REPOSITORY_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_TOOL_ID", instanceWithInferedType.REQ_IF_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_VERSION", instanceWithInferedType.REQ_IF_VERSION, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SOURCE_TOOL_ID", instanceWithInferedType.SOURCE_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TITLE", instanceWithInferedType.TITLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPECIFICATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPECIFICATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPECIFICATION](
					nil,
					"SPECIFICATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SPEC_HIERARCHY:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		BasicFieldtoForm("OBJECT", instanceWithInferedType.OBJECT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_TABLE_INTERNAL", instanceWithInferedType.IS_TABLE_INTERNAL, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SPEC_OBJECT_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_OBJECT_TYPES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_OBJECT_TYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.SPEC_OBJECT_TYPE](
					nil,
					"SPEC_OBJECT_TYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
