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
	case "REQIFHEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQIFHEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFHEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		reqifheader := new(models.REQIFHEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(reqifheader, formGroup, probe)
	}
	formStage.Commit()
}
