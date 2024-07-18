// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.REQIF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQIF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_HEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_HIERARCHY:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_HIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_OBJECT_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_OBJECT_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
