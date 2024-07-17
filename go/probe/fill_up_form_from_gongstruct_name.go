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
	case "REQ_IF_HEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_header, formGroup, probe)
	}
	formStage.Commit()
}
