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
		case "REQ_IF_HEADER":
			FormDivSelectFieldToField(&(reqif_.REQ_IF_HEADER), reqifFormCallback.probe.stageOfInterest, formDiv)
		case "REQ_IF_CONTENT":
			FormDivSelectFieldToField(&(reqif_.REQ_IF_CONTENT), reqifFormCallback.probe.stageOfInterest, formDiv)
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
func __gong__New__REQ_IF_CONTENTFormCallback(
	req_if_content *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) {
	req_if_contentFormCallback = new(REQ_IF_CONTENTFormCallback)
	req_if_contentFormCallback.probe = probe
	req_if_contentFormCallback.req_if_content = req_if_content
	req_if_contentFormCallback.formGroup = formGroup

	req_if_contentFormCallback.CreationMode = (req_if_content == nil)

	return
}

type REQ_IF_CONTENTFormCallback struct {
	req_if_content *models.REQ_IF_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) OnSave() {

	log.Println("REQ_IF_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_contentFormCallback.probe.formStage.Checkout()

	if req_if_contentFormCallback.req_if_content == nil {
		req_if_contentFormCallback.req_if_content = new(models.REQ_IF_CONTENT).Stage(req_if_contentFormCallback.probe.stageOfInterest)
	}
	req_if_content_ := req_if_contentFormCallback.req_if_content
	_ = req_if_content_

	for _, formDiv := range req_if_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_content_.Name), formDiv)
		case "SPECIFICATION":
			FormDivSelectFieldToField(&(req_if_content_.SPECIFICATION), req_if_contentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_content_.Unstage(req_if_contentFormCallback.probe.stageOfInterest)
	}

	req_if_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_CONTENT](
		req_if_contentFormCallback.probe,
	)
	req_if_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_contentFormCallback.CreationMode || req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			req_if_contentFormCallback.probe,
			newFormGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		FillUpForm(req_if_content, newFormGroup, req_if_contentFormCallback.probe)
		req_if_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_contentFormCallback.probe)
}
func __gong__New__REQ_IF_HEADERFormCallback(
	req_if_header *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) {
	req_if_headerFormCallback = new(REQ_IF_HEADERFormCallback)
	req_if_headerFormCallback.probe = probe
	req_if_headerFormCallback.req_if_header = req_if_header
	req_if_headerFormCallback.formGroup = formGroup

	req_if_headerFormCallback.CreationMode = (req_if_header == nil)

	return
}

type REQ_IF_HEADERFormCallback struct {
	req_if_header *models.REQ_IF_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) OnSave() {

	log.Println("REQ_IF_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_headerFormCallback.probe.formStage.Checkout()

	if req_if_headerFormCallback.req_if_header == nil {
		req_if_headerFormCallback.req_if_header = new(models.REQ_IF_HEADER).Stage(req_if_headerFormCallback.probe.stageOfInterest)
	}
	req_if_header_ := req_if_headerFormCallback.req_if_header
	_ = req_if_header_

	for _, formDiv := range req_if_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_header_.Name), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(req_if_header_.COMMENT), formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&(req_if_header_.CREATION_TIME), formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&(req_if_header_.REPOSITORY_ID), formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_TOOL_ID), formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_VERSION), formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.SOURCE_TOOL_ID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(req_if_header_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_header_.Unstage(req_if_headerFormCallback.probe.stageOfInterest)
	}

	req_if_headerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_HEADER](
		req_if_headerFormCallback.probe,
	)
	req_if_headerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_headerFormCallback.CreationMode || req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			req_if_headerFormCallback.probe,
			newFormGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		FillUpForm(req_if_header, newFormGroup, req_if_headerFormCallback.probe)
		req_if_headerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_headerFormCallback.probe)
}
func __gong__New__SPECIFICATIONFormCallback(
	specification *models.SPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationFormCallback *SPECIFICATIONFormCallback) {
	specificationFormCallback = new(SPECIFICATIONFormCallback)
	specificationFormCallback.probe = probe
	specificationFormCallback.specification = specification
	specificationFormCallback.formGroup = formGroup

	specificationFormCallback.CreationMode = (specification == nil)

	return
}

type SPECIFICATIONFormCallback struct {
	specification *models.SPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationFormCallback *SPECIFICATIONFormCallback) OnSave() {

	log.Println("SPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationFormCallback.probe.formStage.Checkout()

	if specificationFormCallback.specification == nil {
		specificationFormCallback.specification = new(models.SPECIFICATION).Stage(specificationFormCallback.probe.stageOfInterest)
	}
	specification_ := specificationFormCallback.specification
	_ = specification_

	for _, formDiv := range specificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_.Name), formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(specification_.CHILDREN), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(specification_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_.Unstage(specificationFormCallback.probe.stageOfInterest)
	}

	specificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATION](
		specificationFormCallback.probe,
	)
	specificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specificationFormCallback.CreationMode || specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			specificationFormCallback.probe,
			newFormGroup,
		)
		specification := new(models.SPECIFICATION)
		FillUpForm(specification, newFormGroup, specificationFormCallback.probe)
		specificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specificationFormCallback.probe)
}
func __gong__New__SPEC_HIERARCHYFormCallback(
	spec_hierarchy *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) {
	spec_hierarchyFormCallback = new(SPEC_HIERARCHYFormCallback)
	spec_hierarchyFormCallback.probe = probe
	spec_hierarchyFormCallback.spec_hierarchy = spec_hierarchy
	spec_hierarchyFormCallback.formGroup = formGroup

	spec_hierarchyFormCallback.CreationMode = (spec_hierarchy == nil)

	return
}

type SPEC_HIERARCHYFormCallback struct {
	spec_hierarchy *models.SPEC_HIERARCHY

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) OnSave() {

	log.Println("SPEC_HIERARCHYFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_hierarchyFormCallback.probe.formStage.Checkout()

	if spec_hierarchyFormCallback.spec_hierarchy == nil {
		spec_hierarchyFormCallback.spec_hierarchy = new(models.SPEC_HIERARCHY).Stage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}
	spec_hierarchy_ := spec_hierarchyFormCallback.spec_hierarchy
	_ = spec_hierarchy_

	for _, formDiv := range spec_hierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_hierarchy_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_hierarchy_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_hierarchy_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_EDITABLE), formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_TABLE_INTERNAL), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_hierarchy_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_hierarchy_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchy_.Unstage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}

	spec_hierarchyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_HIERARCHY](
		spec_hierarchyFormCallback.probe,
	)
	spec_hierarchyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_hierarchyFormCallback.CreationMode || spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_hierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			spec_hierarchyFormCallback.probe,
			newFormGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		FillUpForm(spec_hierarchy, newFormGroup, spec_hierarchyFormCallback.probe)
		spec_hierarchyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_hierarchyFormCallback.probe)
}
