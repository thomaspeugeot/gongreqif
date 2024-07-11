// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__REQIFFormCallback(
	reqif *models.REQIF,
	probe *Probe,
	formGroup *table.FormGroup,
) (reqifFormCallback *REQIFFormCallback) {
	reqifFormCallback = new(REQIFFormCallback)
	reqifFormCallback.probe = probe
	reqifFormCallback.reqif = reqif
	reqifFormCallback.formGroup = formGroup

	reqifFormCallback.CreationMode = (reqif == nil)

	return
}

type REQIFFormCallback struct {
	reqif *models.REQIF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (reqifFormCallback *REQIFFormCallback) OnSave() {

	log.Println("REQIFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	reqifFormCallback.probe.formStage.Checkout()

	if reqifFormCallback.reqif == nil {
		reqifFormCallback.reqif = new(models.REQIF).Stage(reqifFormCallback.probe.stageOfInterest)
	}
	reqif_ := reqifFormCallback.reqif
	_ = reqif_

	for _, formDiv := range reqifFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(reqif_.Name), formDiv)
		case "REQIFHEADER":
			FormDivSelectFieldToField(&(reqif_.REQIFHEADER), reqifFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if reqifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqif_.Unstage(reqifFormCallback.probe.stageOfInterest)
	}

	reqifFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQIF](
		reqifFormCallback.probe,
	)
	reqifFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if reqifFormCallback.CreationMode || reqifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqifFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(reqifFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQIFFormCallback(
			nil,
			reqifFormCallback.probe,
			newFormGroup,
		)
		reqif := new(models.REQIF)
		FillUpForm(reqif, newFormGroup, reqifFormCallback.probe)
		reqifFormCallback.probe.formStage.Commit()
	}

	fillUpTree(reqifFormCallback.probe)
}
func __gong__New__REQIFHEADERFormCallback(
	reqifheader *models.REQIFHEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (reqifheaderFormCallback *REQIFHEADERFormCallback) {
	reqifheaderFormCallback = new(REQIFHEADERFormCallback)
	reqifheaderFormCallback.probe = probe
	reqifheaderFormCallback.reqifheader = reqifheader
	reqifheaderFormCallback.formGroup = formGroup

	reqifheaderFormCallback.CreationMode = (reqifheader == nil)

	return
}

type REQIFHEADERFormCallback struct {
	reqifheader *models.REQIFHEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (reqifheaderFormCallback *REQIFHEADERFormCallback) OnSave() {

	log.Println("REQIFHEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	reqifheaderFormCallback.probe.formStage.Checkout()

	if reqifheaderFormCallback.reqifheader == nil {
		reqifheaderFormCallback.reqifheader = new(models.REQIFHEADER).Stage(reqifheaderFormCallback.probe.stageOfInterest)
	}
	reqifheader_ := reqifheaderFormCallback.reqifheader
	_ = reqifheader_

	for _, formDiv := range reqifheaderFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(reqifheader_.Name), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(reqifheader_.IDENTIFIERAttr), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(reqifheader_.COMMENT), formDiv)
		case "CREATIONTIME":
			FormDivBasicFieldToField(&(reqifheader_.CREATIONTIME), formDiv)
		case "REPOSITORYID":
			FormDivBasicFieldToField(&(reqifheader_.REPOSITORYID), formDiv)
		case "REQIFTOOLID":
			FormDivBasicFieldToField(&(reqifheader_.REQIFTOOLID), formDiv)
		case "REQIFVERSION":
			FormDivBasicFieldToField(&(reqifheader_.REQIFVERSION), formDiv)
		case "SOURCETOOLID":
			FormDivBasicFieldToField(&(reqifheader_.SOURCETOOLID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(reqifheader_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if reqifheaderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqifheader_.Unstage(reqifheaderFormCallback.probe.stageOfInterest)
	}

	reqifheaderFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQIFHEADER](
		reqifheaderFormCallback.probe,
	)
	reqifheaderFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if reqifheaderFormCallback.CreationMode || reqifheaderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqifheaderFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(reqifheaderFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQIFHEADERFormCallback(
			nil,
			reqifheaderFormCallback.probe,
			newFormGroup,
		)
		reqifheader := new(models.REQIFHEADER)
		FillUpForm(reqifheader, newFormGroup, reqifheaderFormCallback.probe)
		reqifheaderFormCallback.probe.formStage.Commit()
	}

	fillUpTree(reqifheaderFormCallback.probe)
}
