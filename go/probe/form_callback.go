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
func __gong__New__CONTENTFormCallback(
	content *models.CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (contentFormCallback *CONTENTFormCallback) {
	contentFormCallback = new(CONTENTFormCallback)
	contentFormCallback.probe = probe
	contentFormCallback.content = content
	contentFormCallback.formGroup = formGroup

	contentFormCallback.CreationMode = (content == nil)

	return
}

type CONTENTFormCallback struct {
	content *models.CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (contentFormCallback *CONTENTFormCallback) OnSave() {

	log.Println("CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	contentFormCallback.probe.formStage.Checkout()

	if contentFormCallback.content == nil {
		contentFormCallback.content = new(models.CONTENT).Stage(contentFormCallback.probe.stageOfInterest)
	}
	content_ := contentFormCallback.content
	_ = content_

	for _, formDiv := range contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(content_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		content_.Unstage(contentFormCallback.probe.stageOfInterest)
	}

	contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CONTENT](
		contentFormCallback.probe,
	)
	contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if contentFormCallback.CreationMode || contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CONTENTFormCallback(
			nil,
			contentFormCallback.probe,
			newFormGroup,
		)
		content := new(models.CONTENT)
		FillUpForm(content, newFormGroup, contentFormCallback.probe)
		contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(contentFormCallback.probe)
}
func __gong__New__HEADERFormCallback(
	header *models.HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (headerFormCallback *HEADERFormCallback) {
	headerFormCallback = new(HEADERFormCallback)
	headerFormCallback.probe = probe
	headerFormCallback.header = header
	headerFormCallback.formGroup = formGroup

	headerFormCallback.CreationMode = (header == nil)

	return
}

type HEADERFormCallback struct {
	header *models.HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (headerFormCallback *HEADERFormCallback) OnSave() {

	log.Println("HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	headerFormCallback.probe.formStage.Checkout()

	if headerFormCallback.header == nil {
		headerFormCallback.header = new(models.HEADER).Stage(headerFormCallback.probe.stageOfInterest)
	}
	header_ := headerFormCallback.header
	_ = header_

	for _, formDiv := range headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(header_.Name), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(header_.IDENTIFIERAttr), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(header_.COMMENT), formDiv)
		case "CREATIONTIME":
			FormDivBasicFieldToField(&(header_.CREATIONTIME), formDiv)
		case "REPOSITORYID":
			FormDivBasicFieldToField(&(header_.REPOSITORYID), formDiv)
		case "REQIFTOOLID":
			FormDivBasicFieldToField(&(header_.REQIFTOOLID), formDiv)
		case "REQIFVERSION":
			FormDivBasicFieldToField(&(header_.REQIFVERSION), formDiv)
		case "SOURCETOOLID":
			FormDivBasicFieldToField(&(header_.SOURCETOOLID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(header_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		header_.Unstage(headerFormCallback.probe.stageOfInterest)
	}

	headerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.HEADER](
		headerFormCallback.probe,
	)
	headerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if headerFormCallback.CreationMode || headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__HEADERFormCallback(
			nil,
			headerFormCallback.probe,
			newFormGroup,
		)
		header := new(models.HEADER)
		FillUpForm(header, newFormGroup, headerFormCallback.probe)
		headerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(headerFormCallback.probe)
}
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
		case "HEADER":
			FormDivSelectFieldToField(&(reqif_.HEADER), reqifFormCallback.probe.stageOfInterest, formDiv)
		case "CONTENT":
			FormDivSelectFieldToField(&(reqif_.CONTENT), reqifFormCallback.probe.stageOfInterest, formDiv)
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
