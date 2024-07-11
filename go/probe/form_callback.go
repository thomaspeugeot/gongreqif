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
