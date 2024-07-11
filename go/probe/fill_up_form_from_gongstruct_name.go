// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "CONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		content := new(models.CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(content, formGroup, probe)
	case "HEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		header := new(models.HEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(header, formGroup, probe)
	case "REQIF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQIF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFFormCallback(
			nil,
			probe,
			formGroup,
		)
		reqif := new(models.REQIF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(reqif, formGroup, probe)
	}
	formStage.Commit()
}
