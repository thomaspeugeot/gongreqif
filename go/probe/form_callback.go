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
func __gong__New__ALTERNATIVEIDFormCallback(
	alternativeid *models.ALTERNATIVEID,
	probe *Probe,
	formGroup *table.FormGroup,
) (alternativeidFormCallback *ALTERNATIVEIDFormCallback) {
	alternativeidFormCallback = new(ALTERNATIVEIDFormCallback)
	alternativeidFormCallback.probe = probe
	alternativeidFormCallback.alternativeid = alternativeid
	alternativeidFormCallback.formGroup = formGroup

	alternativeidFormCallback.CreationMode = (alternativeid == nil)

	return
}

type ALTERNATIVEIDFormCallback struct {
	alternativeid *models.ALTERNATIVEID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (alternativeidFormCallback *ALTERNATIVEIDFormCallback) OnSave() {

	log.Println("ALTERNATIVEIDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	alternativeidFormCallback.probe.formStage.Checkout()

	if alternativeidFormCallback.alternativeid == nil {
		alternativeidFormCallback.alternativeid = new(models.ALTERNATIVEID).Stage(alternativeidFormCallback.probe.stageOfInterest)
	}
	alternativeid_ := alternativeidFormCallback.alternativeid
	_ = alternativeid_

	for _, formDiv := range alternativeidFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(alternativeid_.Name), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(alternativeid_.IDENTIFIERAttr), formDiv)
		}
	}

	// manage the suppress operation
	if alternativeidFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternativeid_.Unstage(alternativeidFormCallback.probe.stageOfInterest)
	}

	alternativeidFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ALTERNATIVEID](
		alternativeidFormCallback.probe,
	)
	alternativeidFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if alternativeidFormCallback.CreationMode || alternativeidFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternativeidFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(alternativeidFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ALTERNATIVEIDFormCallback(
			nil,
			alternativeidFormCallback.probe,
			newFormGroup,
		)
		alternativeid := new(models.ALTERNATIVEID)
		FillUpForm(alternativeid, newFormGroup, alternativeidFormCallback.probe)
		alternativeidFormCallback.probe.formStage.Commit()
	}

	fillUpTree(alternativeidFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONBOOLEANFormCallback(
	attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionbooleanFormCallback *ATTRIBUTEDEFINITIONBOOLEANFormCallback) {
	attributedefinitionbooleanFormCallback = new(ATTRIBUTEDEFINITIONBOOLEANFormCallback)
	attributedefinitionbooleanFormCallback.probe = probe
	attributedefinitionbooleanFormCallback.attributedefinitionboolean = attributedefinitionboolean
	attributedefinitionbooleanFormCallback.formGroup = formGroup

	attributedefinitionbooleanFormCallback.CreationMode = (attributedefinitionboolean == nil)

	return
}

type ATTRIBUTEDEFINITIONBOOLEANFormCallback struct {
	attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionbooleanFormCallback *ATTRIBUTEDEFINITIONBOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONBOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionbooleanFormCallback.probe.formStage.Checkout()

	if attributedefinitionbooleanFormCallback.attributedefinitionboolean == nil {
		attributedefinitionbooleanFormCallback.attributedefinitionboolean = new(models.ATTRIBUTEDEFINITIONBOOLEAN).Stage(attributedefinitionbooleanFormCallback.probe.stageOfInterest)
	}
	attributedefinitionboolean_ := attributedefinitionbooleanFormCallback.attributedefinitionboolean
	_ = attributedefinitionboolean_

	for _, formDiv := range attributedefinitionbooleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitionboolean_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitionboolean_.ALTERNATIVEID), attributedefinitionbooleanFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitionboolean_.DEFAULTVALUE), attributedefinitionbooleanFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitionboolean_.TYPE), attributedefinitionbooleanFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONBOOLEAN":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONBOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionbooleanFormCallback.probe.stageOfInterest,
				attributedefinitionbooleanFormCallback.probe.backRepoOfInterest,
				attributedefinitionboolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, attributedefinitionboolean_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionbooleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, attributedefinitionboolean_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, attributedefinitionboolean_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONBOOLEAN, attributedefinitionboolean_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionboolean_.Unstage(attributedefinitionbooleanFormCallback.probe.stageOfInterest)
	}

	attributedefinitionbooleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONBOOLEAN](
		attributedefinitionbooleanFormCallback.probe,
	)
	attributedefinitionbooleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionbooleanFormCallback.CreationMode || attributedefinitionbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionbooleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionbooleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONBOOLEANFormCallback(
			nil,
			attributedefinitionbooleanFormCallback.probe,
			newFormGroup,
		)
		attributedefinitionboolean := new(models.ATTRIBUTEDEFINITIONBOOLEAN)
		FillUpForm(attributedefinitionboolean, newFormGroup, attributedefinitionbooleanFormCallback.probe)
		attributedefinitionbooleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionbooleanFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONDATEFormCallback(
	attributedefinitiondate *models.ATTRIBUTEDEFINITIONDATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitiondateFormCallback *ATTRIBUTEDEFINITIONDATEFormCallback) {
	attributedefinitiondateFormCallback = new(ATTRIBUTEDEFINITIONDATEFormCallback)
	attributedefinitiondateFormCallback.probe = probe
	attributedefinitiondateFormCallback.attributedefinitiondate = attributedefinitiondate
	attributedefinitiondateFormCallback.formGroup = formGroup

	attributedefinitiondateFormCallback.CreationMode = (attributedefinitiondate == nil)

	return
}

type ATTRIBUTEDEFINITIONDATEFormCallback struct {
	attributedefinitiondate *models.ATTRIBUTEDEFINITIONDATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitiondateFormCallback *ATTRIBUTEDEFINITIONDATEFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONDATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitiondateFormCallback.probe.formStage.Checkout()

	if attributedefinitiondateFormCallback.attributedefinitiondate == nil {
		attributedefinitiondateFormCallback.attributedefinitiondate = new(models.ATTRIBUTEDEFINITIONDATE).Stage(attributedefinitiondateFormCallback.probe.stageOfInterest)
	}
	attributedefinitiondate_ := attributedefinitiondateFormCallback.attributedefinitiondate
	_ = attributedefinitiondate_

	for _, formDiv := range attributedefinitiondateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitiondate_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitiondate_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitiondate_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitiondate_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitiondate_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitiondate_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitiondate_.ALTERNATIVEID), attributedefinitiondateFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitiondate_.DEFAULTVALUE), attributedefinitiondateFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitiondate_.TYPE), attributedefinitiondateFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONDATE":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONDATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitiondateFormCallback.probe.stageOfInterest,
				attributedefinitiondateFormCallback.probe.backRepoOfInterest,
				attributedefinitiondate_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, attributedefinitiondate_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitiondateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, attributedefinitiondate_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, attributedefinitiondate_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONDATE, attributedefinitiondate_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitiondateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitiondate_.Unstage(attributedefinitiondateFormCallback.probe.stageOfInterest)
	}

	attributedefinitiondateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONDATE](
		attributedefinitiondateFormCallback.probe,
	)
	attributedefinitiondateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitiondateFormCallback.CreationMode || attributedefinitiondateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitiondateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitiondateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONDATEFormCallback(
			nil,
			attributedefinitiondateFormCallback.probe,
			newFormGroup,
		)
		attributedefinitiondate := new(models.ATTRIBUTEDEFINITIONDATE)
		FillUpForm(attributedefinitiondate, newFormGroup, attributedefinitiondateFormCallback.probe)
		attributedefinitiondateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitiondateFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONENUMERATIONFormCallback(
	attributedefinitionenumeration *models.ATTRIBUTEDEFINITIONENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionenumerationFormCallback *ATTRIBUTEDEFINITIONENUMERATIONFormCallback) {
	attributedefinitionenumerationFormCallback = new(ATTRIBUTEDEFINITIONENUMERATIONFormCallback)
	attributedefinitionenumerationFormCallback.probe = probe
	attributedefinitionenumerationFormCallback.attributedefinitionenumeration = attributedefinitionenumeration
	attributedefinitionenumerationFormCallback.formGroup = formGroup

	attributedefinitionenumerationFormCallback.CreationMode = (attributedefinitionenumeration == nil)

	return
}

type ATTRIBUTEDEFINITIONENUMERATIONFormCallback struct {
	attributedefinitionenumeration *models.ATTRIBUTEDEFINITIONENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionenumerationFormCallback *ATTRIBUTEDEFINITIONENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionenumerationFormCallback.probe.formStage.Checkout()

	if attributedefinitionenumerationFormCallback.attributedefinitionenumeration == nil {
		attributedefinitionenumerationFormCallback.attributedefinitionenumeration = new(models.ATTRIBUTEDEFINITIONENUMERATION).Stage(attributedefinitionenumerationFormCallback.probe.stageOfInterest)
	}
	attributedefinitionenumeration_ := attributedefinitionenumerationFormCallback.attributedefinitionenumeration
	_ = attributedefinitionenumeration_

	for _, formDiv := range attributedefinitionenumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.LONGNAMEAttr), formDiv)
		case "MULTIVALUEDAttr":
			FormDivBasicFieldToField(&(attributedefinitionenumeration_.MULTIVALUEDAttr), formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitionenumeration_.DEFAULTVALUE), attributedefinitionenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitionenumeration_.ALTERNATIVEID), attributedefinitionenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitionenumeration_.TYPE), attributedefinitionenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONENUMERATION":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionenumerationFormCallback.probe.stageOfInterest,
				attributedefinitionenumerationFormCallback.probe.backRepoOfInterest,
				attributedefinitionenumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, attributedefinitionenumeration_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionenumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, attributedefinitionenumeration_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, attributedefinitionenumeration_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONENUMERATION, attributedefinitionenumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionenumeration_.Unstage(attributedefinitionenumerationFormCallback.probe.stageOfInterest)
	}

	attributedefinitionenumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONENUMERATION](
		attributedefinitionenumerationFormCallback.probe,
	)
	attributedefinitionenumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionenumerationFormCallback.CreationMode || attributedefinitionenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionenumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionenumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONENUMERATIONFormCallback(
			nil,
			attributedefinitionenumerationFormCallback.probe,
			newFormGroup,
		)
		attributedefinitionenumeration := new(models.ATTRIBUTEDEFINITIONENUMERATION)
		FillUpForm(attributedefinitionenumeration, newFormGroup, attributedefinitionenumerationFormCallback.probe)
		attributedefinitionenumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionenumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONINTEGERFormCallback(
	attributedefinitioninteger *models.ATTRIBUTEDEFINITIONINTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionintegerFormCallback *ATTRIBUTEDEFINITIONINTEGERFormCallback) {
	attributedefinitionintegerFormCallback = new(ATTRIBUTEDEFINITIONINTEGERFormCallback)
	attributedefinitionintegerFormCallback.probe = probe
	attributedefinitionintegerFormCallback.attributedefinitioninteger = attributedefinitioninteger
	attributedefinitionintegerFormCallback.formGroup = formGroup

	attributedefinitionintegerFormCallback.CreationMode = (attributedefinitioninteger == nil)

	return
}

type ATTRIBUTEDEFINITIONINTEGERFormCallback struct {
	attributedefinitioninteger *models.ATTRIBUTEDEFINITIONINTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionintegerFormCallback *ATTRIBUTEDEFINITIONINTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONINTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionintegerFormCallback.probe.formStage.Checkout()

	if attributedefinitionintegerFormCallback.attributedefinitioninteger == nil {
		attributedefinitionintegerFormCallback.attributedefinitioninteger = new(models.ATTRIBUTEDEFINITIONINTEGER).Stage(attributedefinitionintegerFormCallback.probe.stageOfInterest)
	}
	attributedefinitioninteger_ := attributedefinitionintegerFormCallback.attributedefinitioninteger
	_ = attributedefinitioninteger_

	for _, formDiv := range attributedefinitionintegerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitioninteger_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitioninteger_.ALTERNATIVEID), attributedefinitionintegerFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitioninteger_.DEFAULTVALUE), attributedefinitionintegerFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitioninteger_.TYPE), attributedefinitionintegerFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONINTEGER":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONINTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionintegerFormCallback.probe.stageOfInterest,
				attributedefinitionintegerFormCallback.probe.backRepoOfInterest,
				attributedefinitioninteger_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, attributedefinitioninteger_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionintegerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, attributedefinitioninteger_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, attributedefinitioninteger_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONINTEGER, attributedefinitioninteger_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitioninteger_.Unstage(attributedefinitionintegerFormCallback.probe.stageOfInterest)
	}

	attributedefinitionintegerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONINTEGER](
		attributedefinitionintegerFormCallback.probe,
	)
	attributedefinitionintegerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionintegerFormCallback.CreationMode || attributedefinitionintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionintegerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionintegerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONINTEGERFormCallback(
			nil,
			attributedefinitionintegerFormCallback.probe,
			newFormGroup,
		)
		attributedefinitioninteger := new(models.ATTRIBUTEDEFINITIONINTEGER)
		FillUpForm(attributedefinitioninteger, newFormGroup, attributedefinitionintegerFormCallback.probe)
		attributedefinitionintegerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionintegerFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONREALFormCallback(
	attributedefinitionreal *models.ATTRIBUTEDEFINITIONREAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionrealFormCallback *ATTRIBUTEDEFINITIONREALFormCallback) {
	attributedefinitionrealFormCallback = new(ATTRIBUTEDEFINITIONREALFormCallback)
	attributedefinitionrealFormCallback.probe = probe
	attributedefinitionrealFormCallback.attributedefinitionreal = attributedefinitionreal
	attributedefinitionrealFormCallback.formGroup = formGroup

	attributedefinitionrealFormCallback.CreationMode = (attributedefinitionreal == nil)

	return
}

type ATTRIBUTEDEFINITIONREALFormCallback struct {
	attributedefinitionreal *models.ATTRIBUTEDEFINITIONREAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionrealFormCallback *ATTRIBUTEDEFINITIONREALFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONREALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionrealFormCallback.probe.formStage.Checkout()

	if attributedefinitionrealFormCallback.attributedefinitionreal == nil {
		attributedefinitionrealFormCallback.attributedefinitionreal = new(models.ATTRIBUTEDEFINITIONREAL).Stage(attributedefinitionrealFormCallback.probe.stageOfInterest)
	}
	attributedefinitionreal_ := attributedefinitionrealFormCallback.attributedefinitionreal
	_ = attributedefinitionreal_

	for _, formDiv := range attributedefinitionrealFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitionreal_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitionreal_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitionreal_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitionreal_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitionreal_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitionreal_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitionreal_.ALTERNATIVEID), attributedefinitionrealFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitionreal_.DEFAULTVALUE), attributedefinitionrealFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitionreal_.TYPE), attributedefinitionrealFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONREAL":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONREAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionrealFormCallback.probe.stageOfInterest,
				attributedefinitionrealFormCallback.probe.backRepoOfInterest,
				attributedefinitionreal_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, attributedefinitionreal_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionrealFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, attributedefinitionreal_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, attributedefinitionreal_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONREAL, attributedefinitionreal_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionrealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionreal_.Unstage(attributedefinitionrealFormCallback.probe.stageOfInterest)
	}

	attributedefinitionrealFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONREAL](
		attributedefinitionrealFormCallback.probe,
	)
	attributedefinitionrealFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionrealFormCallback.CreationMode || attributedefinitionrealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionrealFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionrealFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONREALFormCallback(
			nil,
			attributedefinitionrealFormCallback.probe,
			newFormGroup,
		)
		attributedefinitionreal := new(models.ATTRIBUTEDEFINITIONREAL)
		FillUpForm(attributedefinitionreal, newFormGroup, attributedefinitionrealFormCallback.probe)
		attributedefinitionrealFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionrealFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONSTRINGFormCallback(
	attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionstringFormCallback *ATTRIBUTEDEFINITIONSTRINGFormCallback) {
	attributedefinitionstringFormCallback = new(ATTRIBUTEDEFINITIONSTRINGFormCallback)
	attributedefinitionstringFormCallback.probe = probe
	attributedefinitionstringFormCallback.attributedefinitionstring = attributedefinitionstring
	attributedefinitionstringFormCallback.formGroup = formGroup

	attributedefinitionstringFormCallback.CreationMode = (attributedefinitionstring == nil)

	return
}

type ATTRIBUTEDEFINITIONSTRINGFormCallback struct {
	attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionstringFormCallback *ATTRIBUTEDEFINITIONSTRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONSTRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionstringFormCallback.probe.formStage.Checkout()

	if attributedefinitionstringFormCallback.attributedefinitionstring == nil {
		attributedefinitionstringFormCallback.attributedefinitionstring = new(models.ATTRIBUTEDEFINITIONSTRING).Stage(attributedefinitionstringFormCallback.probe.stageOfInterest)
	}
	attributedefinitionstring_ := attributedefinitionstringFormCallback.attributedefinitionstring
	_ = attributedefinitionstring_

	for _, formDiv := range attributedefinitionstringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitionstring_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitionstring_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitionstring_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitionstring_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitionstring_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitionstring_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitionstring_.ALTERNATIVEID), attributedefinitionstringFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitionstring_.DEFAULTVALUE), attributedefinitionstringFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitionstring_.TYPE), attributedefinitionstringFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONSTRING":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONSTRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionstringFormCallback.probe.stageOfInterest,
				attributedefinitionstringFormCallback.probe.backRepoOfInterest,
				attributedefinitionstring_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, attributedefinitionstring_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionstringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, attributedefinitionstring_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, attributedefinitionstring_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONSTRING, attributedefinitionstring_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionstring_.Unstage(attributedefinitionstringFormCallback.probe.stageOfInterest)
	}

	attributedefinitionstringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONSTRING](
		attributedefinitionstringFormCallback.probe,
	)
	attributedefinitionstringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionstringFormCallback.CreationMode || attributedefinitionstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionstringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONSTRINGFormCallback(
			nil,
			attributedefinitionstringFormCallback.probe,
			newFormGroup,
		)
		attributedefinitionstring := new(models.ATTRIBUTEDEFINITIONSTRING)
		FillUpForm(attributedefinitionstring, newFormGroup, attributedefinitionstringFormCallback.probe)
		attributedefinitionstringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionstringFormCallback.probe)
}
func __gong__New__ATTRIBUTEDEFINITIONXHTMLFormCallback(
	attributedefinitionxhtml *models.ATTRIBUTEDEFINITIONXHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributedefinitionxhtmlFormCallback *ATTRIBUTEDEFINITIONXHTMLFormCallback) {
	attributedefinitionxhtmlFormCallback = new(ATTRIBUTEDEFINITIONXHTMLFormCallback)
	attributedefinitionxhtmlFormCallback.probe = probe
	attributedefinitionxhtmlFormCallback.attributedefinitionxhtml = attributedefinitionxhtml
	attributedefinitionxhtmlFormCallback.formGroup = formGroup

	attributedefinitionxhtmlFormCallback.CreationMode = (attributedefinitionxhtml == nil)

	return
}

type ATTRIBUTEDEFINITIONXHTMLFormCallback struct {
	attributedefinitionxhtml *models.ATTRIBUTEDEFINITIONXHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributedefinitionxhtmlFormCallback *ATTRIBUTEDEFINITIONXHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTEDEFINITIONXHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributedefinitionxhtmlFormCallback.probe.formStage.Checkout()

	if attributedefinitionxhtmlFormCallback.attributedefinitionxhtml == nil {
		attributedefinitionxhtmlFormCallback.attributedefinitionxhtml = new(models.ATTRIBUTEDEFINITIONXHTML).Stage(attributedefinitionxhtmlFormCallback.probe.stageOfInterest)
	}
	attributedefinitionxhtml_ := attributedefinitionxhtmlFormCallback.attributedefinitionxhtml
	_ = attributedefinitionxhtml_

	for _, formDiv := range attributedefinitionxhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.ISEDITABLEAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(attributedefinitionxhtml_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(attributedefinitionxhtml_.ALTERNATIVEID), attributedefinitionxhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "DEFAULTVALUE":
			FormDivSelectFieldToField(&(attributedefinitionxhtml_.DEFAULTVALUE), attributedefinitionxhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(attributedefinitionxhtml_.TYPE), attributedefinitionxhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES:ATTRIBUTEDEFINITIONXHTML":
			// we need to retrieve the field owner before the change
			var pastSPECATTRIBUTESOwner *models.SPECATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONXHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributedefinitionxhtmlFormCallback.probe.stageOfInterest,
				attributedefinitionxhtmlFormCallback.probe.backRepoOfInterest,
				attributedefinitionxhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECATTRIBUTESOwner = reverseFieldOwner.(*models.SPECATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECATTRIBUTESOwner != nil {
					idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, attributedefinitionxhtml_)
					pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specattributes := range *models.GetGongstructInstancesSet[models.SPECATTRIBUTES](attributedefinitionxhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specattributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECATTRIBUTESOwner := _specattributes // we have a match
						if pastSPECATTRIBUTESOwner != nil {
							if newSPECATTRIBUTESOwner != pastSPECATTRIBUTESOwner {
								idx := slices.Index(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, attributedefinitionxhtml_)
								pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML = slices.Delete(pastSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, idx, idx+1)
								newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, attributedefinitionxhtml_)
							}
						} else {
							newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML = append(newSPECATTRIBUTESOwner.ATTRIBUTEDEFINITIONXHTML, attributedefinitionxhtml_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributedefinitionxhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionxhtml_.Unstage(attributedefinitionxhtmlFormCallback.probe.stageOfInterest)
	}

	attributedefinitionxhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEDEFINITIONXHTML](
		attributedefinitionxhtmlFormCallback.probe,
	)
	attributedefinitionxhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributedefinitionxhtmlFormCallback.CreationMode || attributedefinitionxhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributedefinitionxhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributedefinitionxhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONXHTMLFormCallback(
			nil,
			attributedefinitionxhtmlFormCallback.probe,
			newFormGroup,
		)
		attributedefinitionxhtml := new(models.ATTRIBUTEDEFINITIONXHTML)
		FillUpForm(attributedefinitionxhtml, newFormGroup, attributedefinitionxhtmlFormCallback.probe)
		attributedefinitionxhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributedefinitionxhtmlFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEBOOLEANFormCallback(
	attributevalueboolean *models.ATTRIBUTEVALUEBOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevaluebooleanFormCallback *ATTRIBUTEVALUEBOOLEANFormCallback) {
	attributevaluebooleanFormCallback = new(ATTRIBUTEVALUEBOOLEANFormCallback)
	attributevaluebooleanFormCallback.probe = probe
	attributevaluebooleanFormCallback.attributevalueboolean = attributevalueboolean
	attributevaluebooleanFormCallback.formGroup = formGroup

	attributevaluebooleanFormCallback.CreationMode = (attributevalueboolean == nil)

	return
}

type ATTRIBUTEVALUEBOOLEANFormCallback struct {
	attributevalueboolean *models.ATTRIBUTEVALUEBOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevaluebooleanFormCallback *ATTRIBUTEVALUEBOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEBOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevaluebooleanFormCallback.probe.formStage.Checkout()

	if attributevaluebooleanFormCallback.attributevalueboolean == nil {
		attributevaluebooleanFormCallback.attributevalueboolean = new(models.ATTRIBUTEVALUEBOOLEAN).Stage(attributevaluebooleanFormCallback.probe.stageOfInterest)
	}
	attributevalueboolean_ := attributevaluebooleanFormCallback.attributevalueboolean
	_ = attributevalueboolean_

	for _, formDiv := range attributevaluebooleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevalueboolean_.Name), formDiv)
		case "THEVALUEAttr":
			FormDivBasicFieldToField(&(attributevalueboolean_.THEVALUEAttr), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevalueboolean_.DEFINITION), attributevaluebooleanFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevaluebooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevalueboolean_.Unstage(attributevaluebooleanFormCallback.probe.stageOfInterest)
	}

	attributevaluebooleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEBOOLEAN](
		attributevaluebooleanFormCallback.probe,
	)
	attributevaluebooleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevaluebooleanFormCallback.CreationMode || attributevaluebooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluebooleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevaluebooleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEBOOLEANFormCallback(
			nil,
			attributevaluebooleanFormCallback.probe,
			newFormGroup,
		)
		attributevalueboolean := new(models.ATTRIBUTEVALUEBOOLEAN)
		FillUpForm(attributevalueboolean, newFormGroup, attributevaluebooleanFormCallback.probe)
		attributevaluebooleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevaluebooleanFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEDATEFormCallback(
	attributevaluedate *models.ATTRIBUTEVALUEDATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevaluedateFormCallback *ATTRIBUTEVALUEDATEFormCallback) {
	attributevaluedateFormCallback = new(ATTRIBUTEVALUEDATEFormCallback)
	attributevaluedateFormCallback.probe = probe
	attributevaluedateFormCallback.attributevaluedate = attributevaluedate
	attributevaluedateFormCallback.formGroup = formGroup

	attributevaluedateFormCallback.CreationMode = (attributevaluedate == nil)

	return
}

type ATTRIBUTEVALUEDATEFormCallback struct {
	attributevaluedate *models.ATTRIBUTEVALUEDATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevaluedateFormCallback *ATTRIBUTEVALUEDATEFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEDATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevaluedateFormCallback.probe.formStage.Checkout()

	if attributevaluedateFormCallback.attributevaluedate == nil {
		attributevaluedateFormCallback.attributevaluedate = new(models.ATTRIBUTEVALUEDATE).Stage(attributevaluedateFormCallback.probe.stageOfInterest)
	}
	attributevaluedate_ := attributevaluedateFormCallback.attributevaluedate
	_ = attributevaluedate_

	for _, formDiv := range attributevaluedateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevaluedate_.Name), formDiv)
		case "THEVALUEAttr":
			FormDivBasicFieldToField(&(attributevaluedate_.THEVALUEAttr), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevaluedate_.DEFINITION), attributevaluedateFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevaluedateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluedate_.Unstage(attributevaluedateFormCallback.probe.stageOfInterest)
	}

	attributevaluedateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEDATE](
		attributevaluedateFormCallback.probe,
	)
	attributevaluedateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevaluedateFormCallback.CreationMode || attributevaluedateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluedateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevaluedateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEDATEFormCallback(
			nil,
			attributevaluedateFormCallback.probe,
			newFormGroup,
		)
		attributevaluedate := new(models.ATTRIBUTEVALUEDATE)
		FillUpForm(attributevaluedate, newFormGroup, attributevaluedateFormCallback.probe)
		attributevaluedateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevaluedateFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEENUMERATIONFormCallback(
	attributevalueenumeration *models.ATTRIBUTEVALUEENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevalueenumerationFormCallback *ATTRIBUTEVALUEENUMERATIONFormCallback) {
	attributevalueenumerationFormCallback = new(ATTRIBUTEVALUEENUMERATIONFormCallback)
	attributevalueenumerationFormCallback.probe = probe
	attributevalueenumerationFormCallback.attributevalueenumeration = attributevalueenumeration
	attributevalueenumerationFormCallback.formGroup = formGroup

	attributevalueenumerationFormCallback.CreationMode = (attributevalueenumeration == nil)

	return
}

type ATTRIBUTEVALUEENUMERATIONFormCallback struct {
	attributevalueenumeration *models.ATTRIBUTEVALUEENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevalueenumerationFormCallback *ATTRIBUTEVALUEENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevalueenumerationFormCallback.probe.formStage.Checkout()

	if attributevalueenumerationFormCallback.attributevalueenumeration == nil {
		attributevalueenumerationFormCallback.attributevalueenumeration = new(models.ATTRIBUTEVALUEENUMERATION).Stage(attributevalueenumerationFormCallback.probe.stageOfInterest)
	}
	attributevalueenumeration_ := attributevalueenumerationFormCallback.attributevalueenumeration
	_ = attributevalueenumeration_

	for _, formDiv := range attributevalueenumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevalueenumeration_.Name), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevalueenumeration_.DEFINITION), attributevalueenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(attributevalueenumeration_.VALUES), attributevalueenumerationFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevalueenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevalueenumeration_.Unstage(attributevalueenumerationFormCallback.probe.stageOfInterest)
	}

	attributevalueenumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEENUMERATION](
		attributevalueenumerationFormCallback.probe,
	)
	attributevalueenumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevalueenumerationFormCallback.CreationMode || attributevalueenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevalueenumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevalueenumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEENUMERATIONFormCallback(
			nil,
			attributevalueenumerationFormCallback.probe,
			newFormGroup,
		)
		attributevalueenumeration := new(models.ATTRIBUTEVALUEENUMERATION)
		FillUpForm(attributevalueenumeration, newFormGroup, attributevalueenumerationFormCallback.probe)
		attributevalueenumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevalueenumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEINTEGERFormCallback(
	attributevalueinteger *models.ATTRIBUTEVALUEINTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevalueintegerFormCallback *ATTRIBUTEVALUEINTEGERFormCallback) {
	attributevalueintegerFormCallback = new(ATTRIBUTEVALUEINTEGERFormCallback)
	attributevalueintegerFormCallback.probe = probe
	attributevalueintegerFormCallback.attributevalueinteger = attributevalueinteger
	attributevalueintegerFormCallback.formGroup = formGroup

	attributevalueintegerFormCallback.CreationMode = (attributevalueinteger == nil)

	return
}

type ATTRIBUTEVALUEINTEGERFormCallback struct {
	attributevalueinteger *models.ATTRIBUTEVALUEINTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevalueintegerFormCallback *ATTRIBUTEVALUEINTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEINTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevalueintegerFormCallback.probe.formStage.Checkout()

	if attributevalueintegerFormCallback.attributevalueinteger == nil {
		attributevalueintegerFormCallback.attributevalueinteger = new(models.ATTRIBUTEVALUEINTEGER).Stage(attributevalueintegerFormCallback.probe.stageOfInterest)
	}
	attributevalueinteger_ := attributevalueintegerFormCallback.attributevalueinteger
	_ = attributevalueinteger_

	for _, formDiv := range attributevalueintegerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevalueinteger_.Name), formDiv)
		case "THEVALUEAttr":
			FormDivBasicFieldToField(&(attributevalueinteger_.THEVALUEAttr), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevalueinteger_.DEFINITION), attributevalueintegerFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevalueintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevalueinteger_.Unstage(attributevalueintegerFormCallback.probe.stageOfInterest)
	}

	attributevalueintegerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEINTEGER](
		attributevalueintegerFormCallback.probe,
	)
	attributevalueintegerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevalueintegerFormCallback.CreationMode || attributevalueintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevalueintegerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevalueintegerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEINTEGERFormCallback(
			nil,
			attributevalueintegerFormCallback.probe,
			newFormGroup,
		)
		attributevalueinteger := new(models.ATTRIBUTEVALUEINTEGER)
		FillUpForm(attributevalueinteger, newFormGroup, attributevalueintegerFormCallback.probe)
		attributevalueintegerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevalueintegerFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEREALFormCallback(
	attributevaluereal *models.ATTRIBUTEVALUEREAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevaluerealFormCallback *ATTRIBUTEVALUEREALFormCallback) {
	attributevaluerealFormCallback = new(ATTRIBUTEVALUEREALFormCallback)
	attributevaluerealFormCallback.probe = probe
	attributevaluerealFormCallback.attributevaluereal = attributevaluereal
	attributevaluerealFormCallback.formGroup = formGroup

	attributevaluerealFormCallback.CreationMode = (attributevaluereal == nil)

	return
}

type ATTRIBUTEVALUEREALFormCallback struct {
	attributevaluereal *models.ATTRIBUTEVALUEREAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevaluerealFormCallback *ATTRIBUTEVALUEREALFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEREALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevaluerealFormCallback.probe.formStage.Checkout()

	if attributevaluerealFormCallback.attributevaluereal == nil {
		attributevaluerealFormCallback.attributevaluereal = new(models.ATTRIBUTEVALUEREAL).Stage(attributevaluerealFormCallback.probe.stageOfInterest)
	}
	attributevaluereal_ := attributevaluerealFormCallback.attributevaluereal
	_ = attributevaluereal_

	for _, formDiv := range attributevaluerealFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevaluereal_.Name), formDiv)
		case "THEVALUEAttr":
			FormDivBasicFieldToField(&(attributevaluereal_.THEVALUEAttr), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevaluereal_.DEFINITION), attributevaluerealFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevaluerealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluereal_.Unstage(attributevaluerealFormCallback.probe.stageOfInterest)
	}

	attributevaluerealFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEREAL](
		attributevaluerealFormCallback.probe,
	)
	attributevaluerealFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevaluerealFormCallback.CreationMode || attributevaluerealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluerealFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevaluerealFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEREALFormCallback(
			nil,
			attributevaluerealFormCallback.probe,
			newFormGroup,
		)
		attributevaluereal := new(models.ATTRIBUTEVALUEREAL)
		FillUpForm(attributevaluereal, newFormGroup, attributevaluerealFormCallback.probe)
		attributevaluerealFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevaluerealFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUESTRINGFormCallback(
	attributevaluestring *models.ATTRIBUTEVALUESTRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevaluestringFormCallback *ATTRIBUTEVALUESTRINGFormCallback) {
	attributevaluestringFormCallback = new(ATTRIBUTEVALUESTRINGFormCallback)
	attributevaluestringFormCallback.probe = probe
	attributevaluestringFormCallback.attributevaluestring = attributevaluestring
	attributevaluestringFormCallback.formGroup = formGroup

	attributevaluestringFormCallback.CreationMode = (attributevaluestring == nil)

	return
}

type ATTRIBUTEVALUESTRINGFormCallback struct {
	attributevaluestring *models.ATTRIBUTEVALUESTRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevaluestringFormCallback *ATTRIBUTEVALUESTRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUESTRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevaluestringFormCallback.probe.formStage.Checkout()

	if attributevaluestringFormCallback.attributevaluestring == nil {
		attributevaluestringFormCallback.attributevaluestring = new(models.ATTRIBUTEVALUESTRING).Stage(attributevaluestringFormCallback.probe.stageOfInterest)
	}
	attributevaluestring_ := attributevaluestringFormCallback.attributevaluestring
	_ = attributevaluestring_

	for _, formDiv := range attributevaluestringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevaluestring_.Name), formDiv)
		case "THEVALUEAttr":
			FormDivBasicFieldToField(&(attributevaluestring_.THEVALUEAttr), formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevaluestring_.DEFINITION), attributevaluestringFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevaluestringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluestring_.Unstage(attributevaluestringFormCallback.probe.stageOfInterest)
	}

	attributevaluestringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUESTRING](
		attributevaluestringFormCallback.probe,
	)
	attributevaluestringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevaluestringFormCallback.CreationMode || attributevaluestringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluestringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevaluestringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUESTRINGFormCallback(
			nil,
			attributevaluestringFormCallback.probe,
			newFormGroup,
		)
		attributevaluestring := new(models.ATTRIBUTEVALUESTRING)
		FillUpForm(attributevaluestring, newFormGroup, attributevaluestringFormCallback.probe)
		attributevaluestringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevaluestringFormCallback.probe)
}
func __gong__New__ATTRIBUTEVALUEXHTMLFormCallback(
	attributevaluexhtml *models.ATTRIBUTEVALUEXHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributevaluexhtmlFormCallback *ATTRIBUTEVALUEXHTMLFormCallback) {
	attributevaluexhtmlFormCallback = new(ATTRIBUTEVALUEXHTMLFormCallback)
	attributevaluexhtmlFormCallback.probe = probe
	attributevaluexhtmlFormCallback.attributevaluexhtml = attributevaluexhtml
	attributevaluexhtmlFormCallback.formGroup = formGroup

	attributevaluexhtmlFormCallback.CreationMode = (attributevaluexhtml == nil)

	return
}

type ATTRIBUTEVALUEXHTMLFormCallback struct {
	attributevaluexhtml *models.ATTRIBUTEVALUEXHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributevaluexhtmlFormCallback *ATTRIBUTEVALUEXHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTEVALUEXHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributevaluexhtmlFormCallback.probe.formStage.Checkout()

	if attributevaluexhtmlFormCallback.attributevaluexhtml == nil {
		attributevaluexhtmlFormCallback.attributevaluexhtml = new(models.ATTRIBUTEVALUEXHTML).Stage(attributevaluexhtmlFormCallback.probe.stageOfInterest)
	}
	attributevaluexhtml_ := attributevaluexhtmlFormCallback.attributevaluexhtml
	_ = attributevaluexhtml_

	for _, formDiv := range attributevaluexhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributevaluexhtml_.Name), formDiv)
		case "ISSIMPLIFIEDAttr":
			FormDivBasicFieldToField(&(attributevaluexhtml_.ISSIMPLIFIEDAttr), formDiv)
		case "THEVALUE":
			FormDivSelectFieldToField(&(attributevaluexhtml_.THEVALUE), attributevaluexhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "THEORIGINALVALUE":
			FormDivSelectFieldToField(&(attributevaluexhtml_.THEORIGINALVALUE), attributevaluexhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "DEFINITION":
			FormDivSelectFieldToField(&(attributevaluexhtml_.DEFINITION), attributevaluexhtmlFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if attributevaluexhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluexhtml_.Unstage(attributevaluexhtmlFormCallback.probe.stageOfInterest)
	}

	attributevaluexhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTEVALUEXHTML](
		attributevaluexhtmlFormCallback.probe,
	)
	attributevaluexhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributevaluexhtmlFormCallback.CreationMode || attributevaluexhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributevaluexhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributevaluexhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTEVALUEXHTMLFormCallback(
			nil,
			attributevaluexhtmlFormCallback.probe,
			newFormGroup,
		)
		attributevaluexhtml := new(models.ATTRIBUTEVALUEXHTML)
		FillUpForm(attributevaluexhtml, newFormGroup, attributevaluexhtmlFormCallback.probe)
		attributevaluexhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributevaluexhtmlFormCallback.probe)
}
func __gong__New__CHILDRENFormCallback(
	children *models.CHILDREN,
	probe *Probe,
	formGroup *table.FormGroup,
) (childrenFormCallback *CHILDRENFormCallback) {
	childrenFormCallback = new(CHILDRENFormCallback)
	childrenFormCallback.probe = probe
	childrenFormCallback.children = children
	childrenFormCallback.formGroup = formGroup

	childrenFormCallback.CreationMode = (children == nil)

	return
}

type CHILDRENFormCallback struct {
	children *models.CHILDREN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (childrenFormCallback *CHILDRENFormCallback) OnSave() {

	log.Println("CHILDRENFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	childrenFormCallback.probe.formStage.Checkout()

	if childrenFormCallback.children == nil {
		childrenFormCallback.children = new(models.CHILDREN).Stage(childrenFormCallback.probe.stageOfInterest)
	}
	children_ := childrenFormCallback.children
	_ = children_

	for _, formDiv := range childrenFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(children_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		children_.Unstage(childrenFormCallback.probe.stageOfInterest)
	}

	childrenFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CHILDREN](
		childrenFormCallback.probe,
	)
	childrenFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if childrenFormCallback.CreationMode || childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		childrenFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(childrenFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CHILDRENFormCallback(
			nil,
			childrenFormCallback.probe,
			newFormGroup,
		)
		children := new(models.CHILDREN)
		FillUpForm(children, newFormGroup, childrenFormCallback.probe)
		childrenFormCallback.probe.formStage.Commit()
	}

	fillUpTree(childrenFormCallback.probe)
}
func __gong__New__CORECONTENTFormCallback(
	corecontent *models.CORECONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (corecontentFormCallback *CORECONTENTFormCallback) {
	corecontentFormCallback = new(CORECONTENTFormCallback)
	corecontentFormCallback.probe = probe
	corecontentFormCallback.corecontent = corecontent
	corecontentFormCallback.formGroup = formGroup

	corecontentFormCallback.CreationMode = (corecontent == nil)

	return
}

type CORECONTENTFormCallback struct {
	corecontent *models.CORECONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (corecontentFormCallback *CORECONTENTFormCallback) OnSave() {

	log.Println("CORECONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	corecontentFormCallback.probe.formStage.Checkout()

	if corecontentFormCallback.corecontent == nil {
		corecontentFormCallback.corecontent = new(models.CORECONTENT).Stage(corecontentFormCallback.probe.stageOfInterest)
	}
	corecontent_ := corecontentFormCallback.corecontent
	_ = corecontent_

	for _, formDiv := range corecontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(corecontent_.Name), formDiv)
		case "REQIFCONTENT":
			FormDivSelectFieldToField(&(corecontent_.REQIFCONTENT), corecontentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if corecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		corecontent_.Unstage(corecontentFormCallback.probe.stageOfInterest)
	}

	corecontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CORECONTENT](
		corecontentFormCallback.probe,
	)
	corecontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if corecontentFormCallback.CreationMode || corecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		corecontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(corecontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CORECONTENTFormCallback(
			nil,
			corecontentFormCallback.probe,
			newFormGroup,
		)
		corecontent := new(models.CORECONTENT)
		FillUpForm(corecontent, newFormGroup, corecontentFormCallback.probe)
		corecontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(corecontentFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONBOOLEANFormCallback(
	datatypedefinitionboolean *models.DATATYPEDEFINITIONBOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionbooleanFormCallback *DATATYPEDEFINITIONBOOLEANFormCallback) {
	datatypedefinitionbooleanFormCallback = new(DATATYPEDEFINITIONBOOLEANFormCallback)
	datatypedefinitionbooleanFormCallback.probe = probe
	datatypedefinitionbooleanFormCallback.datatypedefinitionboolean = datatypedefinitionboolean
	datatypedefinitionbooleanFormCallback.formGroup = formGroup

	datatypedefinitionbooleanFormCallback.CreationMode = (datatypedefinitionboolean == nil)

	return
}

type DATATYPEDEFINITIONBOOLEANFormCallback struct {
	datatypedefinitionboolean *models.DATATYPEDEFINITIONBOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionbooleanFormCallback *DATATYPEDEFINITIONBOOLEANFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONBOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionbooleanFormCallback.probe.formStage.Checkout()

	if datatypedefinitionbooleanFormCallback.datatypedefinitionboolean == nil {
		datatypedefinitionbooleanFormCallback.datatypedefinitionboolean = new(models.DATATYPEDEFINITIONBOOLEAN).Stage(datatypedefinitionbooleanFormCallback.probe.stageOfInterest)
	}
	datatypedefinitionboolean_ := datatypedefinitionbooleanFormCallback.datatypedefinitionboolean
	_ = datatypedefinitionboolean_

	for _, formDiv := range datatypedefinitionbooleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitionboolean_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitionboolean_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitionboolean_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionboolean_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionboolean_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitionboolean_.ALTERNATIVEID), datatypedefinitionbooleanFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONBOOLEAN":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONBOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionbooleanFormCallback.probe.stageOfInterest,
				datatypedefinitionbooleanFormCallback.probe.backRepoOfInterest,
				datatypedefinitionboolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, datatypedefinitionboolean_)
					pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionbooleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, datatypedefinitionboolean_)
								pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN = append(newDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, datatypedefinitionboolean_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN = append(newDATATYPESOwner.DATATYPEDEFINITIONBOOLEAN, datatypedefinitionboolean_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionboolean_.Unstage(datatypedefinitionbooleanFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionbooleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONBOOLEAN](
		datatypedefinitionbooleanFormCallback.probe,
	)
	datatypedefinitionbooleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionbooleanFormCallback.CreationMode || datatypedefinitionbooleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionbooleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionbooleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONBOOLEANFormCallback(
			nil,
			datatypedefinitionbooleanFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitionboolean := new(models.DATATYPEDEFINITIONBOOLEAN)
		FillUpForm(datatypedefinitionboolean, newFormGroup, datatypedefinitionbooleanFormCallback.probe)
		datatypedefinitionbooleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionbooleanFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONDATEFormCallback(
	datatypedefinitiondate *models.DATATYPEDEFINITIONDATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitiondateFormCallback *DATATYPEDEFINITIONDATEFormCallback) {
	datatypedefinitiondateFormCallback = new(DATATYPEDEFINITIONDATEFormCallback)
	datatypedefinitiondateFormCallback.probe = probe
	datatypedefinitiondateFormCallback.datatypedefinitiondate = datatypedefinitiondate
	datatypedefinitiondateFormCallback.formGroup = formGroup

	datatypedefinitiondateFormCallback.CreationMode = (datatypedefinitiondate == nil)

	return
}

type DATATYPEDEFINITIONDATEFormCallback struct {
	datatypedefinitiondate *models.DATATYPEDEFINITIONDATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitiondateFormCallback *DATATYPEDEFINITIONDATEFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONDATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitiondateFormCallback.probe.formStage.Checkout()

	if datatypedefinitiondateFormCallback.datatypedefinitiondate == nil {
		datatypedefinitiondateFormCallback.datatypedefinitiondate = new(models.DATATYPEDEFINITIONDATE).Stage(datatypedefinitiondateFormCallback.probe.stageOfInterest)
	}
	datatypedefinitiondate_ := datatypedefinitiondateFormCallback.datatypedefinitiondate
	_ = datatypedefinitiondate_

	for _, formDiv := range datatypedefinitiondateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitiondate_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitiondate_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitiondate_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitiondate_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitiondate_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitiondate_.ALTERNATIVEID), datatypedefinitiondateFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONDATE":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONDATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitiondateFormCallback.probe.stageOfInterest,
				datatypedefinitiondateFormCallback.probe.backRepoOfInterest,
				datatypedefinitiondate_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONDATE, datatypedefinitiondate_)
					pastDATATYPESOwner.DATATYPEDEFINITIONDATE = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONDATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitiondateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONDATE, datatypedefinitiondate_)
								pastDATATYPESOwner.DATATYPEDEFINITIONDATE = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONDATE, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONDATE = append(newDATATYPESOwner.DATATYPEDEFINITIONDATE, datatypedefinitiondate_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONDATE = append(newDATATYPESOwner.DATATYPEDEFINITIONDATE, datatypedefinitiondate_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitiondateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitiondate_.Unstage(datatypedefinitiondateFormCallback.probe.stageOfInterest)
	}

	datatypedefinitiondateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONDATE](
		datatypedefinitiondateFormCallback.probe,
	)
	datatypedefinitiondateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitiondateFormCallback.CreationMode || datatypedefinitiondateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitiondateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitiondateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONDATEFormCallback(
			nil,
			datatypedefinitiondateFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitiondate := new(models.DATATYPEDEFINITIONDATE)
		FillUpForm(datatypedefinitiondate, newFormGroup, datatypedefinitiondateFormCallback.probe)
		datatypedefinitiondateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitiondateFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONENUMERATIONFormCallback(
	datatypedefinitionenumeration *models.DATATYPEDEFINITIONENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionenumerationFormCallback *DATATYPEDEFINITIONENUMERATIONFormCallback) {
	datatypedefinitionenumerationFormCallback = new(DATATYPEDEFINITIONENUMERATIONFormCallback)
	datatypedefinitionenumerationFormCallback.probe = probe
	datatypedefinitionenumerationFormCallback.datatypedefinitionenumeration = datatypedefinitionenumeration
	datatypedefinitionenumerationFormCallback.formGroup = formGroup

	datatypedefinitionenumerationFormCallback.CreationMode = (datatypedefinitionenumeration == nil)

	return
}

type DATATYPEDEFINITIONENUMERATIONFormCallback struct {
	datatypedefinitionenumeration *models.DATATYPEDEFINITIONENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionenumerationFormCallback *DATATYPEDEFINITIONENUMERATIONFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionenumerationFormCallback.probe.formStage.Checkout()

	if datatypedefinitionenumerationFormCallback.datatypedefinitionenumeration == nil {
		datatypedefinitionenumerationFormCallback.datatypedefinitionenumeration = new(models.DATATYPEDEFINITIONENUMERATION).Stage(datatypedefinitionenumerationFormCallback.probe.stageOfInterest)
	}
	datatypedefinitionenumeration_ := datatypedefinitionenumerationFormCallback.datatypedefinitionenumeration
	_ = datatypedefinitionenumeration_

	for _, formDiv := range datatypedefinitionenumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitionenumeration_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitionenumeration_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitionenumeration_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionenumeration_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionenumeration_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitionenumeration_.ALTERNATIVEID), datatypedefinitionenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFIEDVALUES":
			FormDivSelectFieldToField(&(datatypedefinitionenumeration_.SPECIFIEDVALUES), datatypedefinitionenumerationFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONENUMERATION":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionenumerationFormCallback.probe.stageOfInterest,
				datatypedefinitionenumerationFormCallback.probe.backRepoOfInterest,
				datatypedefinitionenumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, datatypedefinitionenumeration_)
					pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionenumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, datatypedefinitionenumeration_)
								pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONENUMERATION = append(newDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, datatypedefinitionenumeration_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONENUMERATION = append(newDATATYPESOwner.DATATYPEDEFINITIONENUMERATION, datatypedefinitionenumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionenumeration_.Unstage(datatypedefinitionenumerationFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionenumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONENUMERATION](
		datatypedefinitionenumerationFormCallback.probe,
	)
	datatypedefinitionenumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionenumerationFormCallback.CreationMode || datatypedefinitionenumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionenumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionenumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONENUMERATIONFormCallback(
			nil,
			datatypedefinitionenumerationFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitionenumeration := new(models.DATATYPEDEFINITIONENUMERATION)
		FillUpForm(datatypedefinitionenumeration, newFormGroup, datatypedefinitionenumerationFormCallback.probe)
		datatypedefinitionenumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionenumerationFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONINTEGERFormCallback(
	datatypedefinitioninteger *models.DATATYPEDEFINITIONINTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionintegerFormCallback *DATATYPEDEFINITIONINTEGERFormCallback) {
	datatypedefinitionintegerFormCallback = new(DATATYPEDEFINITIONINTEGERFormCallback)
	datatypedefinitionintegerFormCallback.probe = probe
	datatypedefinitionintegerFormCallback.datatypedefinitioninteger = datatypedefinitioninteger
	datatypedefinitionintegerFormCallback.formGroup = formGroup

	datatypedefinitionintegerFormCallback.CreationMode = (datatypedefinitioninteger == nil)

	return
}

type DATATYPEDEFINITIONINTEGERFormCallback struct {
	datatypedefinitioninteger *models.DATATYPEDEFINITIONINTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionintegerFormCallback *DATATYPEDEFINITIONINTEGERFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONINTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionintegerFormCallback.probe.formStage.Checkout()

	if datatypedefinitionintegerFormCallback.datatypedefinitioninteger == nil {
		datatypedefinitionintegerFormCallback.datatypedefinitioninteger = new(models.DATATYPEDEFINITIONINTEGER).Stage(datatypedefinitionintegerFormCallback.probe.stageOfInterest)
	}
	datatypedefinitioninteger_ := datatypedefinitionintegerFormCallback.datatypedefinitioninteger
	_ = datatypedefinitioninteger_

	for _, formDiv := range datatypedefinitionintegerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.LONGNAMEAttr), formDiv)
		case "MAXAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.MAXAttr), formDiv)
		case "MINAttr":
			FormDivBasicFieldToField(&(datatypedefinitioninteger_.MINAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitioninteger_.ALTERNATIVEID), datatypedefinitionintegerFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONINTEGER":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONINTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionintegerFormCallback.probe.stageOfInterest,
				datatypedefinitionintegerFormCallback.probe.backRepoOfInterest,
				datatypedefinitioninteger_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER, datatypedefinitioninteger_)
					pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionintegerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER, datatypedefinitioninteger_)
								pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONINTEGER, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONINTEGER = append(newDATATYPESOwner.DATATYPEDEFINITIONINTEGER, datatypedefinitioninteger_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONINTEGER = append(newDATATYPESOwner.DATATYPEDEFINITIONINTEGER, datatypedefinitioninteger_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitioninteger_.Unstage(datatypedefinitionintegerFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionintegerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONINTEGER](
		datatypedefinitionintegerFormCallback.probe,
	)
	datatypedefinitionintegerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionintegerFormCallback.CreationMode || datatypedefinitionintegerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionintegerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionintegerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONINTEGERFormCallback(
			nil,
			datatypedefinitionintegerFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitioninteger := new(models.DATATYPEDEFINITIONINTEGER)
		FillUpForm(datatypedefinitioninteger, newFormGroup, datatypedefinitionintegerFormCallback.probe)
		datatypedefinitionintegerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionintegerFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONREALFormCallback(
	datatypedefinitionreal *models.DATATYPEDEFINITIONREAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionrealFormCallback *DATATYPEDEFINITIONREALFormCallback) {
	datatypedefinitionrealFormCallback = new(DATATYPEDEFINITIONREALFormCallback)
	datatypedefinitionrealFormCallback.probe = probe
	datatypedefinitionrealFormCallback.datatypedefinitionreal = datatypedefinitionreal
	datatypedefinitionrealFormCallback.formGroup = formGroup

	datatypedefinitionrealFormCallback.CreationMode = (datatypedefinitionreal == nil)

	return
}

type DATATYPEDEFINITIONREALFormCallback struct {
	datatypedefinitionreal *models.DATATYPEDEFINITIONREAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionrealFormCallback *DATATYPEDEFINITIONREALFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONREALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionrealFormCallback.probe.formStage.Checkout()

	if datatypedefinitionrealFormCallback.datatypedefinitionreal == nil {
		datatypedefinitionrealFormCallback.datatypedefinitionreal = new(models.DATATYPEDEFINITIONREAL).Stage(datatypedefinitionrealFormCallback.probe.stageOfInterest)
	}
	datatypedefinitionreal_ := datatypedefinitionrealFormCallback.datatypedefinitionreal
	_ = datatypedefinitionreal_

	for _, formDiv := range datatypedefinitionrealFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.Name), formDiv)
		case "ACCURACYAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.ACCURACYAttr), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.LONGNAMEAttr), formDiv)
		case "MAXAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.MAXAttr), formDiv)
		case "MINAttr":
			FormDivBasicFieldToField(&(datatypedefinitionreal_.MINAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitionreal_.ALTERNATIVEID), datatypedefinitionrealFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONREAL":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONREAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionrealFormCallback.probe.stageOfInterest,
				datatypedefinitionrealFormCallback.probe.backRepoOfInterest,
				datatypedefinitionreal_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONREAL, datatypedefinitionreal_)
					pastDATATYPESOwner.DATATYPEDEFINITIONREAL = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONREAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionrealFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONREAL, datatypedefinitionreal_)
								pastDATATYPESOwner.DATATYPEDEFINITIONREAL = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONREAL, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONREAL = append(newDATATYPESOwner.DATATYPEDEFINITIONREAL, datatypedefinitionreal_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONREAL = append(newDATATYPESOwner.DATATYPEDEFINITIONREAL, datatypedefinitionreal_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionrealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionreal_.Unstage(datatypedefinitionrealFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionrealFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONREAL](
		datatypedefinitionrealFormCallback.probe,
	)
	datatypedefinitionrealFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionrealFormCallback.CreationMode || datatypedefinitionrealFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionrealFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionrealFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONREALFormCallback(
			nil,
			datatypedefinitionrealFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitionreal := new(models.DATATYPEDEFINITIONREAL)
		FillUpForm(datatypedefinitionreal, newFormGroup, datatypedefinitionrealFormCallback.probe)
		datatypedefinitionrealFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionrealFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONSTRINGFormCallback(
	datatypedefinitionstring *models.DATATYPEDEFINITIONSTRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionstringFormCallback *DATATYPEDEFINITIONSTRINGFormCallback) {
	datatypedefinitionstringFormCallback = new(DATATYPEDEFINITIONSTRINGFormCallback)
	datatypedefinitionstringFormCallback.probe = probe
	datatypedefinitionstringFormCallback.datatypedefinitionstring = datatypedefinitionstring
	datatypedefinitionstringFormCallback.formGroup = formGroup

	datatypedefinitionstringFormCallback.CreationMode = (datatypedefinitionstring == nil)

	return
}

type DATATYPEDEFINITIONSTRINGFormCallback struct {
	datatypedefinitionstring *models.DATATYPEDEFINITIONSTRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionstringFormCallback *DATATYPEDEFINITIONSTRINGFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONSTRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionstringFormCallback.probe.formStage.Checkout()

	if datatypedefinitionstringFormCallback.datatypedefinitionstring == nil {
		datatypedefinitionstringFormCallback.datatypedefinitionstring = new(models.DATATYPEDEFINITIONSTRING).Stage(datatypedefinitionstringFormCallback.probe.stageOfInterest)
	}
	datatypedefinitionstring_ := datatypedefinitionstringFormCallback.datatypedefinitionstring
	_ = datatypedefinitionstring_

	for _, formDiv := range datatypedefinitionstringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.LONGNAMEAttr), formDiv)
		case "MAXLENGTHAttr":
			FormDivBasicFieldToField(&(datatypedefinitionstring_.MAXLENGTHAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitionstring_.ALTERNATIVEID), datatypedefinitionstringFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONSTRING":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONSTRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionstringFormCallback.probe.stageOfInterest,
				datatypedefinitionstringFormCallback.probe.backRepoOfInterest,
				datatypedefinitionstring_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONSTRING, datatypedefinitionstring_)
					pastDATATYPESOwner.DATATYPEDEFINITIONSTRING = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONSTRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionstringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONSTRING, datatypedefinitionstring_)
								pastDATATYPESOwner.DATATYPEDEFINITIONSTRING = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONSTRING, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONSTRING = append(newDATATYPESOwner.DATATYPEDEFINITIONSTRING, datatypedefinitionstring_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONSTRING = append(newDATATYPESOwner.DATATYPEDEFINITIONSTRING, datatypedefinitionstring_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionstring_.Unstage(datatypedefinitionstringFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionstringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONSTRING](
		datatypedefinitionstringFormCallback.probe,
	)
	datatypedefinitionstringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionstringFormCallback.CreationMode || datatypedefinitionstringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionstringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONSTRINGFormCallback(
			nil,
			datatypedefinitionstringFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitionstring := new(models.DATATYPEDEFINITIONSTRING)
		FillUpForm(datatypedefinitionstring, newFormGroup, datatypedefinitionstringFormCallback.probe)
		datatypedefinitionstringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionstringFormCallback.probe)
}
func __gong__New__DATATYPEDEFINITIONXHTMLFormCallback(
	datatypedefinitionxhtml *models.DATATYPEDEFINITIONXHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypedefinitionxhtmlFormCallback *DATATYPEDEFINITIONXHTMLFormCallback) {
	datatypedefinitionxhtmlFormCallback = new(DATATYPEDEFINITIONXHTMLFormCallback)
	datatypedefinitionxhtmlFormCallback.probe = probe
	datatypedefinitionxhtmlFormCallback.datatypedefinitionxhtml = datatypedefinitionxhtml
	datatypedefinitionxhtmlFormCallback.formGroup = formGroup

	datatypedefinitionxhtmlFormCallback.CreationMode = (datatypedefinitionxhtml == nil)

	return
}

type DATATYPEDEFINITIONXHTMLFormCallback struct {
	datatypedefinitionxhtml *models.DATATYPEDEFINITIONXHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypedefinitionxhtmlFormCallback *DATATYPEDEFINITIONXHTMLFormCallback) OnSave() {

	log.Println("DATATYPEDEFINITIONXHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypedefinitionxhtmlFormCallback.probe.formStage.Checkout()

	if datatypedefinitionxhtmlFormCallback.datatypedefinitionxhtml == nil {
		datatypedefinitionxhtmlFormCallback.datatypedefinitionxhtml = new(models.DATATYPEDEFINITIONXHTML).Stage(datatypedefinitionxhtmlFormCallback.probe.stageOfInterest)
	}
	datatypedefinitionxhtml_ := datatypedefinitionxhtmlFormCallback.datatypedefinitionxhtml
	_ = datatypedefinitionxhtml_

	for _, formDiv := range datatypedefinitionxhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypedefinitionxhtml_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(datatypedefinitionxhtml_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(datatypedefinitionxhtml_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionxhtml_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(datatypedefinitionxhtml_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(datatypedefinitionxhtml_.ALTERNATIVEID), datatypedefinitionxhtmlFormCallback.probe.stageOfInterest, formDiv)
		case "DATATYPES:DATATYPEDEFINITIONXHTML":
			// we need to retrieve the field owner before the change
			var pastDATATYPESOwner *models.DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONXHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatypedefinitionxhtmlFormCallback.probe.stageOfInterest,
				datatypedefinitionxhtmlFormCallback.probe.backRepoOfInterest,
				datatypedefinitionxhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPESOwner = reverseFieldOwner.(*models.DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPESOwner != nil {
					idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONXHTML, datatypedefinitionxhtml_)
					pastDATATYPESOwner.DATATYPEDEFINITIONXHTML = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONXHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatypes := range *models.GetGongstructInstancesSet[models.DATATYPES](datatypedefinitionxhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPESOwner := _datatypes // we have a match
						if pastDATATYPESOwner != nil {
							if newDATATYPESOwner != pastDATATYPESOwner {
								idx := slices.Index(pastDATATYPESOwner.DATATYPEDEFINITIONXHTML, datatypedefinitionxhtml_)
								pastDATATYPESOwner.DATATYPEDEFINITIONXHTML = slices.Delete(pastDATATYPESOwner.DATATYPEDEFINITIONXHTML, idx, idx+1)
								newDATATYPESOwner.DATATYPEDEFINITIONXHTML = append(newDATATYPESOwner.DATATYPEDEFINITIONXHTML, datatypedefinitionxhtml_)
							}
						} else {
							newDATATYPESOwner.DATATYPEDEFINITIONXHTML = append(newDATATYPESOwner.DATATYPEDEFINITIONXHTML, datatypedefinitionxhtml_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datatypedefinitionxhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionxhtml_.Unstage(datatypedefinitionxhtmlFormCallback.probe.stageOfInterest)
	}

	datatypedefinitionxhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPEDEFINITIONXHTML](
		datatypedefinitionxhtmlFormCallback.probe,
	)
	datatypedefinitionxhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypedefinitionxhtmlFormCallback.CreationMode || datatypedefinitionxhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypedefinitionxhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypedefinitionxhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPEDEFINITIONXHTMLFormCallback(
			nil,
			datatypedefinitionxhtmlFormCallback.probe,
			newFormGroup,
		)
		datatypedefinitionxhtml := new(models.DATATYPEDEFINITIONXHTML)
		FillUpForm(datatypedefinitionxhtml, newFormGroup, datatypedefinitionxhtmlFormCallback.probe)
		datatypedefinitionxhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypedefinitionxhtmlFormCallback.probe)
}
func __gong__New__DATATYPESFormCallback(
	datatypes *models.DATATYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatypesFormCallback *DATATYPESFormCallback) {
	datatypesFormCallback = new(DATATYPESFormCallback)
	datatypesFormCallback.probe = probe
	datatypesFormCallback.datatypes = datatypes
	datatypesFormCallback.formGroup = formGroup

	datatypesFormCallback.CreationMode = (datatypes == nil)

	return
}

type DATATYPESFormCallback struct {
	datatypes *models.DATATYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatypesFormCallback *DATATYPESFormCallback) OnSave() {

	log.Println("DATATYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatypesFormCallback.probe.formStage.Checkout()

	if datatypesFormCallback.datatypes == nil {
		datatypesFormCallback.datatypes = new(models.DATATYPES).Stage(datatypesFormCallback.probe.stageOfInterest)
	}
	datatypes_ := datatypesFormCallback.datatypes
	_ = datatypes_

	for _, formDiv := range datatypesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatypes_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypes_.Unstage(datatypesFormCallback.probe.stageOfInterest)
	}

	datatypesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPES](
		datatypesFormCallback.probe,
	)
	datatypesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatypesFormCallback.CreationMode || datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatypesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatypesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPESFormCallback(
			nil,
			datatypesFormCallback.probe,
			newFormGroup,
		)
		datatypes := new(models.DATATYPES)
		FillUpForm(datatypes, newFormGroup, datatypesFormCallback.probe)
		datatypesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatypesFormCallback.probe)
}
func __gong__New__DEFAULTVALUEFormCallback(
	defaultvalue *models.DEFAULTVALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (defaultvalueFormCallback *DEFAULTVALUEFormCallback) {
	defaultvalueFormCallback = new(DEFAULTVALUEFormCallback)
	defaultvalueFormCallback.probe = probe
	defaultvalueFormCallback.defaultvalue = defaultvalue
	defaultvalueFormCallback.formGroup = formGroup

	defaultvalueFormCallback.CreationMode = (defaultvalue == nil)

	return
}

type DEFAULTVALUEFormCallback struct {
	defaultvalue *models.DEFAULTVALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (defaultvalueFormCallback *DEFAULTVALUEFormCallback) OnSave() {

	log.Println("DEFAULTVALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	defaultvalueFormCallback.probe.formStage.Checkout()

	if defaultvalueFormCallback.defaultvalue == nil {
		defaultvalueFormCallback.defaultvalue = new(models.DEFAULTVALUE).Stage(defaultvalueFormCallback.probe.stageOfInterest)
	}
	defaultvalue_ := defaultvalueFormCallback.defaultvalue
	_ = defaultvalue_

	for _, formDiv := range defaultvalueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(defaultvalue_.Name), formDiv)
		case "ATTRIBUTEVALUEBOOLEAN":
			FormDivSelectFieldToField(&(defaultvalue_.ATTRIBUTEVALUEBOOLEAN), defaultvalueFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if defaultvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		defaultvalue_.Unstage(defaultvalueFormCallback.probe.stageOfInterest)
	}

	defaultvalueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DEFAULTVALUE](
		defaultvalueFormCallback.probe,
	)
	defaultvalueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if defaultvalueFormCallback.CreationMode || defaultvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		defaultvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(defaultvalueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DEFAULTVALUEFormCallback(
			nil,
			defaultvalueFormCallback.probe,
			newFormGroup,
		)
		defaultvalue := new(models.DEFAULTVALUE)
		FillUpForm(defaultvalue, newFormGroup, defaultvalueFormCallback.probe)
		defaultvalueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(defaultvalueFormCallback.probe)
}
func __gong__New__DEFINITIONFormCallback(
	definition *models.DEFINITION,
	probe *Probe,
	formGroup *table.FormGroup,
) (definitionFormCallback *DEFINITIONFormCallback) {
	definitionFormCallback = new(DEFINITIONFormCallback)
	definitionFormCallback.probe = probe
	definitionFormCallback.definition = definition
	definitionFormCallback.formGroup = formGroup

	definitionFormCallback.CreationMode = (definition == nil)

	return
}

type DEFINITIONFormCallback struct {
	definition *models.DEFINITION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (definitionFormCallback *DEFINITIONFormCallback) OnSave() {

	log.Println("DEFINITIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	definitionFormCallback.probe.formStage.Checkout()

	if definitionFormCallback.definition == nil {
		definitionFormCallback.definition = new(models.DEFINITION).Stage(definitionFormCallback.probe.stageOfInterest)
	}
	definition_ := definitionFormCallback.definition
	_ = definition_

	for _, formDiv := range definitionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(definition_.Name), formDiv)
		case "ATTRIBUTEDEFINITIONBOOLEANREF":
			FormDivBasicFieldToField(&(definition_.ATTRIBUTEDEFINITIONBOOLEANREF), formDiv)
		}
	}

	// manage the suppress operation
	if definitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		definition_.Unstage(definitionFormCallback.probe.stageOfInterest)
	}

	definitionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DEFINITION](
		definitionFormCallback.probe,
	)
	definitionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if definitionFormCallback.CreationMode || definitionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		definitionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(definitionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DEFINITIONFormCallback(
			nil,
			definitionFormCallback.probe,
			newFormGroup,
		)
		definition := new(models.DEFINITION)
		FillUpForm(definition, newFormGroup, definitionFormCallback.probe)
		definitionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(definitionFormCallback.probe)
}
func __gong__New__EDITABLEATTSFormCallback(
	editableatts *models.EDITABLEATTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (editableattsFormCallback *EDITABLEATTSFormCallback) {
	editableattsFormCallback = new(EDITABLEATTSFormCallback)
	editableattsFormCallback.probe = probe
	editableattsFormCallback.editableatts = editableatts
	editableattsFormCallback.formGroup = formGroup

	editableattsFormCallback.CreationMode = (editableatts == nil)

	return
}

type EDITABLEATTSFormCallback struct {
	editableatts *models.EDITABLEATTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (editableattsFormCallback *EDITABLEATTSFormCallback) OnSave() {

	log.Println("EDITABLEATTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	editableattsFormCallback.probe.formStage.Checkout()

	if editableattsFormCallback.editableatts == nil {
		editableattsFormCallback.editableatts = new(models.EDITABLEATTS).Stage(editableattsFormCallback.probe.stageOfInterest)
	}
	editableatts_ := editableattsFormCallback.editableatts
	_ = editableatts_

	for _, formDiv := range editableattsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(editableatts_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if editableattsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		editableatts_.Unstage(editableattsFormCallback.probe.stageOfInterest)
	}

	editableattsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.EDITABLEATTS](
		editableattsFormCallback.probe,
	)
	editableattsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if editableattsFormCallback.CreationMode || editableattsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		editableattsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(editableattsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EDITABLEATTSFormCallback(
			nil,
			editableattsFormCallback.probe,
			newFormGroup,
		)
		editableatts := new(models.EDITABLEATTS)
		FillUpForm(editableatts, newFormGroup, editableattsFormCallback.probe)
		editableattsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(editableattsFormCallback.probe)
}
func __gong__New__EMBEDDEDVALUEFormCallback(
	embeddedvalue *models.EMBEDDEDVALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (embeddedvalueFormCallback *EMBEDDEDVALUEFormCallback) {
	embeddedvalueFormCallback = new(EMBEDDEDVALUEFormCallback)
	embeddedvalueFormCallback.probe = probe
	embeddedvalueFormCallback.embeddedvalue = embeddedvalue
	embeddedvalueFormCallback.formGroup = formGroup

	embeddedvalueFormCallback.CreationMode = (embeddedvalue == nil)

	return
}

type EMBEDDEDVALUEFormCallback struct {
	embeddedvalue *models.EMBEDDEDVALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (embeddedvalueFormCallback *EMBEDDEDVALUEFormCallback) OnSave() {

	log.Println("EMBEDDEDVALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	embeddedvalueFormCallback.probe.formStage.Checkout()

	if embeddedvalueFormCallback.embeddedvalue == nil {
		embeddedvalueFormCallback.embeddedvalue = new(models.EMBEDDEDVALUE).Stage(embeddedvalueFormCallback.probe.stageOfInterest)
	}
	embeddedvalue_ := embeddedvalueFormCallback.embeddedvalue
	_ = embeddedvalue_

	for _, formDiv := range embeddedvalueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(embeddedvalue_.Name), formDiv)
		case "KEYAttr":
			FormDivBasicFieldToField(&(embeddedvalue_.KEYAttr), formDiv)
		case "OTHERCONTENTAttr":
			FormDivBasicFieldToField(&(embeddedvalue_.OTHERCONTENTAttr), formDiv)
		}
	}

	// manage the suppress operation
	if embeddedvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embeddedvalue_.Unstage(embeddedvalueFormCallback.probe.stageOfInterest)
	}

	embeddedvalueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.EMBEDDEDVALUE](
		embeddedvalueFormCallback.probe,
	)
	embeddedvalueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if embeddedvalueFormCallback.CreationMode || embeddedvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embeddedvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(embeddedvalueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EMBEDDEDVALUEFormCallback(
			nil,
			embeddedvalueFormCallback.probe,
			newFormGroup,
		)
		embeddedvalue := new(models.EMBEDDEDVALUE)
		FillUpForm(embeddedvalue, newFormGroup, embeddedvalueFormCallback.probe)
		embeddedvalueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(embeddedvalueFormCallback.probe)
}
func __gong__New__ENUMVALUEFormCallback(
	enumvalue *models.ENUMVALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (enumvalueFormCallback *ENUMVALUEFormCallback) {
	enumvalueFormCallback = new(ENUMVALUEFormCallback)
	enumvalueFormCallback.probe = probe
	enumvalueFormCallback.enumvalue = enumvalue
	enumvalueFormCallback.formGroup = formGroup

	enumvalueFormCallback.CreationMode = (enumvalue == nil)

	return
}

type ENUMVALUEFormCallback struct {
	enumvalue *models.ENUMVALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enumvalueFormCallback *ENUMVALUEFormCallback) OnSave() {

	log.Println("ENUMVALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enumvalueFormCallback.probe.formStage.Checkout()

	if enumvalueFormCallback.enumvalue == nil {
		enumvalueFormCallback.enumvalue = new(models.ENUMVALUE).Stage(enumvalueFormCallback.probe.stageOfInterest)
	}
	enumvalue_ := enumvalueFormCallback.enumvalue
	_ = enumvalue_

	for _, formDiv := range enumvalueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enumvalue_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(enumvalue_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(enumvalue_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(enumvalue_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(enumvalue_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(enumvalue_.ALTERNATIVEID), enumvalueFormCallback.probe.stageOfInterest, formDiv)
		case "PROPERTIES":
			FormDivSelectFieldToField(&(enumvalue_.PROPERTIES), enumvalueFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFIEDVALUES:ENUMVALUE":
			// we need to retrieve the field owner before the change
			var pastSPECIFIEDVALUESOwner *models.SPECIFIEDVALUES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFIEDVALUES"
			rf.Fieldname = "ENUMVALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				enumvalueFormCallback.probe.stageOfInterest,
				enumvalueFormCallback.probe.backRepoOfInterest,
				enumvalue_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFIEDVALUESOwner = reverseFieldOwner.(*models.SPECIFIEDVALUES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFIEDVALUESOwner != nil {
					idx := slices.Index(pastSPECIFIEDVALUESOwner.ENUMVALUE, enumvalue_)
					pastSPECIFIEDVALUESOwner.ENUMVALUE = slices.Delete(pastSPECIFIEDVALUESOwner.ENUMVALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specifiedvalues := range *models.GetGongstructInstancesSet[models.SPECIFIEDVALUES](enumvalueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specifiedvalues.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFIEDVALUESOwner := _specifiedvalues // we have a match
						if pastSPECIFIEDVALUESOwner != nil {
							if newSPECIFIEDVALUESOwner != pastSPECIFIEDVALUESOwner {
								idx := slices.Index(pastSPECIFIEDVALUESOwner.ENUMVALUE, enumvalue_)
								pastSPECIFIEDVALUESOwner.ENUMVALUE = slices.Delete(pastSPECIFIEDVALUESOwner.ENUMVALUE, idx, idx+1)
								newSPECIFIEDVALUESOwner.ENUMVALUE = append(newSPECIFIEDVALUESOwner.ENUMVALUE, enumvalue_)
							}
						} else {
							newSPECIFIEDVALUESOwner.ENUMVALUE = append(newSPECIFIEDVALUESOwner.ENUMVALUE, enumvalue_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if enumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumvalue_.Unstage(enumvalueFormCallback.probe.stageOfInterest)
	}

	enumvalueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ENUMVALUE](
		enumvalueFormCallback.probe,
	)
	enumvalueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if enumvalueFormCallback.CreationMode || enumvalueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumvalueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(enumvalueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ENUMVALUEFormCallback(
			nil,
			enumvalueFormCallback.probe,
			newFormGroup,
		)
		enumvalue := new(models.ENUMVALUE)
		FillUpForm(enumvalue, newFormGroup, enumvalueFormCallback.probe)
		enumvalueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(enumvalueFormCallback.probe)
}
func __gong__New__OBJECTFormCallback(
	object *models.OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (objectFormCallback *OBJECTFormCallback) {
	objectFormCallback = new(OBJECTFormCallback)
	objectFormCallback.probe = probe
	objectFormCallback.object = object
	objectFormCallback.formGroup = formGroup

	objectFormCallback.CreationMode = (object == nil)

	return
}

type OBJECTFormCallback struct {
	object *models.OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (objectFormCallback *OBJECTFormCallback) OnSave() {

	log.Println("OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	objectFormCallback.probe.formStage.Checkout()

	if objectFormCallback.object == nil {
		objectFormCallback.object = new(models.OBJECT).Stage(objectFormCallback.probe.stageOfInterest)
	}
	object_ := objectFormCallback.object
	_ = object_

	for _, formDiv := range objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(object_.Name), formDiv)
		case "SPECOBJECTREF":
			FormDivBasicFieldToField(&(object_.SPECOBJECTREF), formDiv)
		}
	}

	// manage the suppress operation
	if objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		object_.Unstage(objectFormCallback.probe.stageOfInterest)
	}

	objectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.OBJECT](
		objectFormCallback.probe,
	)
	objectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if objectFormCallback.CreationMode || objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__OBJECTFormCallback(
			nil,
			objectFormCallback.probe,
			newFormGroup,
		)
		object := new(models.OBJECT)
		FillUpForm(object, newFormGroup, objectFormCallback.probe)
		objectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(objectFormCallback.probe)
}
func __gong__New__PROPERTIESFormCallback(
	properties *models.PROPERTIES,
	probe *Probe,
	formGroup *table.FormGroup,
) (propertiesFormCallback *PROPERTIESFormCallback) {
	propertiesFormCallback = new(PROPERTIESFormCallback)
	propertiesFormCallback.probe = probe
	propertiesFormCallback.properties = properties
	propertiesFormCallback.formGroup = formGroup

	propertiesFormCallback.CreationMode = (properties == nil)

	return
}

type PROPERTIESFormCallback struct {
	properties *models.PROPERTIES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (propertiesFormCallback *PROPERTIESFormCallback) OnSave() {

	log.Println("PROPERTIESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	propertiesFormCallback.probe.formStage.Checkout()

	if propertiesFormCallback.properties == nil {
		propertiesFormCallback.properties = new(models.PROPERTIES).Stage(propertiesFormCallback.probe.stageOfInterest)
	}
	properties_ := propertiesFormCallback.properties
	_ = properties_

	for _, formDiv := range propertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(properties_.Name), formDiv)
		case "EMBEDDEDVALUE":
			FormDivSelectFieldToField(&(properties_.EMBEDDEDVALUE), propertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		properties_.Unstage(propertiesFormCallback.probe.stageOfInterest)
	}

	propertiesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.PROPERTIES](
		propertiesFormCallback.probe,
	)
	propertiesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if propertiesFormCallback.CreationMode || propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		propertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(propertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PROPERTIESFormCallback(
			nil,
			propertiesFormCallback.probe,
			newFormGroup,
		)
		properties := new(models.PROPERTIES)
		FillUpForm(properties, newFormGroup, propertiesFormCallback.probe)
		propertiesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(propertiesFormCallback.probe)
}
func __gong__New__RELATIONGROUPFormCallback(
	relationgroup *models.RELATIONGROUP,
	probe *Probe,
	formGroup *table.FormGroup,
) (relationgroupFormCallback *RELATIONGROUPFormCallback) {
	relationgroupFormCallback = new(RELATIONGROUPFormCallback)
	relationgroupFormCallback.probe = probe
	relationgroupFormCallback.relationgroup = relationgroup
	relationgroupFormCallback.formGroup = formGroup

	relationgroupFormCallback.CreationMode = (relationgroup == nil)

	return
}

type RELATIONGROUPFormCallback struct {
	relationgroup *models.RELATIONGROUP

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relationgroupFormCallback *RELATIONGROUPFormCallback) OnSave() {

	log.Println("RELATIONGROUPFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relationgroupFormCallback.probe.formStage.Checkout()

	if relationgroupFormCallback.relationgroup == nil {
		relationgroupFormCallback.relationgroup = new(models.RELATIONGROUP).Stage(relationgroupFormCallback.probe.stageOfInterest)
	}
	relationgroup_ := relationgroupFormCallback.relationgroup
	_ = relationgroup_

	for _, formDiv := range relationgroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relationgroup_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(relationgroup_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(relationgroup_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(relationgroup_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(relationgroup_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(relationgroup_.ALTERNATIVEID), relationgroupFormCallback.probe.stageOfInterest, formDiv)
		case "SOURCESPECIFICATION":
			FormDivSelectFieldToField(&(relationgroup_.SOURCESPECIFICATION), relationgroupFormCallback.probe.stageOfInterest, formDiv)
		case "SPECRELATIONS":
			FormDivSelectFieldToField(&(relationgroup_.SPECRELATIONS), relationgroupFormCallback.probe.stageOfInterest, formDiv)
		case "TARGETSPECIFICATION":
			FormDivSelectFieldToField(&(relationgroup_.TARGETSPECIFICATION), relationgroupFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(relationgroup_.TYPE), relationgroupFormCallback.probe.stageOfInterest, formDiv)
		case "SPECRELATIONGROUPS:RELATIONGROUP":
			// we need to retrieve the field owner before the change
			var pastSPECRELATIONGROUPSOwner *models.SPECRELATIONGROUPS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECRELATIONGROUPS"
			rf.Fieldname = "RELATIONGROUP"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relationgroupFormCallback.probe.stageOfInterest,
				relationgroupFormCallback.probe.backRepoOfInterest,
				relationgroup_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECRELATIONGROUPSOwner = reverseFieldOwner.(*models.SPECRELATIONGROUPS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECRELATIONGROUPSOwner != nil {
					idx := slices.Index(pastSPECRELATIONGROUPSOwner.RELATIONGROUP, relationgroup_)
					pastSPECRELATIONGROUPSOwner.RELATIONGROUP = slices.Delete(pastSPECRELATIONGROUPSOwner.RELATIONGROUP, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specrelationgroups := range *models.GetGongstructInstancesSet[models.SPECRELATIONGROUPS](relationgroupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specrelationgroups.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECRELATIONGROUPSOwner := _specrelationgroups // we have a match
						if pastSPECRELATIONGROUPSOwner != nil {
							if newSPECRELATIONGROUPSOwner != pastSPECRELATIONGROUPSOwner {
								idx := slices.Index(pastSPECRELATIONGROUPSOwner.RELATIONGROUP, relationgroup_)
								pastSPECRELATIONGROUPSOwner.RELATIONGROUP = slices.Delete(pastSPECRELATIONGROUPSOwner.RELATIONGROUP, idx, idx+1)
								newSPECRELATIONGROUPSOwner.RELATIONGROUP = append(newSPECRELATIONGROUPSOwner.RELATIONGROUP, relationgroup_)
							}
						} else {
							newSPECRELATIONGROUPSOwner.RELATIONGROUP = append(newSPECRELATIONGROUPSOwner.RELATIONGROUP, relationgroup_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if relationgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relationgroup_.Unstage(relationgroupFormCallback.probe.stageOfInterest)
	}

	relationgroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATIONGROUP](
		relationgroupFormCallback.probe,
	)
	relationgroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relationgroupFormCallback.CreationMode || relationgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relationgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relationgroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATIONGROUPFormCallback(
			nil,
			relationgroupFormCallback.probe,
			newFormGroup,
		)
		relationgroup := new(models.RELATIONGROUP)
		FillUpForm(relationgroup, newFormGroup, relationgroupFormCallback.probe)
		relationgroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relationgroupFormCallback.probe)
}
func __gong__New__RELATIONGROUPTYPEFormCallback(
	relationgrouptype *models.RELATIONGROUPTYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (relationgrouptypeFormCallback *RELATIONGROUPTYPEFormCallback) {
	relationgrouptypeFormCallback = new(RELATIONGROUPTYPEFormCallback)
	relationgrouptypeFormCallback.probe = probe
	relationgrouptypeFormCallback.relationgrouptype = relationgrouptype
	relationgrouptypeFormCallback.formGroup = formGroup

	relationgrouptypeFormCallback.CreationMode = (relationgrouptype == nil)

	return
}

type RELATIONGROUPTYPEFormCallback struct {
	relationgrouptype *models.RELATIONGROUPTYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relationgrouptypeFormCallback *RELATIONGROUPTYPEFormCallback) OnSave() {

	log.Println("RELATIONGROUPTYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relationgrouptypeFormCallback.probe.formStage.Checkout()

	if relationgrouptypeFormCallback.relationgrouptype == nil {
		relationgrouptypeFormCallback.relationgrouptype = new(models.RELATIONGROUPTYPE).Stage(relationgrouptypeFormCallback.probe.stageOfInterest)
	}
	relationgrouptype_ := relationgrouptypeFormCallback.relationgrouptype
	_ = relationgrouptype_

	for _, formDiv := range relationgrouptypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relationgrouptype_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(relationgrouptype_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(relationgrouptype_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(relationgrouptype_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(relationgrouptype_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(relationgrouptype_.ALTERNATIVEID), relationgrouptypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES":
			FormDivSelectFieldToField(&(relationgrouptype_.SPECATTRIBUTES), relationgrouptypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECTYPES:RELATIONGROUPTYPE":
			// we need to retrieve the field owner before the change
			var pastSPECTYPESOwner *models.SPECTYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "RELATIONGROUPTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relationgrouptypeFormCallback.probe.stageOfInterest,
				relationgrouptypeFormCallback.probe.backRepoOfInterest,
				relationgrouptype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECTYPESOwner = reverseFieldOwner.(*models.SPECTYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECTYPESOwner != nil {
					idx := slices.Index(pastSPECTYPESOwner.RELATIONGROUPTYPE, relationgrouptype_)
					pastSPECTYPESOwner.RELATIONGROUPTYPE = slices.Delete(pastSPECTYPESOwner.RELATIONGROUPTYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spectypes := range *models.GetGongstructInstancesSet[models.SPECTYPES](relationgrouptypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spectypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECTYPESOwner := _spectypes // we have a match
						if pastSPECTYPESOwner != nil {
							if newSPECTYPESOwner != pastSPECTYPESOwner {
								idx := slices.Index(pastSPECTYPESOwner.RELATIONGROUPTYPE, relationgrouptype_)
								pastSPECTYPESOwner.RELATIONGROUPTYPE = slices.Delete(pastSPECTYPESOwner.RELATIONGROUPTYPE, idx, idx+1)
								newSPECTYPESOwner.RELATIONGROUPTYPE = append(newSPECTYPESOwner.RELATIONGROUPTYPE, relationgrouptype_)
							}
						} else {
							newSPECTYPESOwner.RELATIONGROUPTYPE = append(newSPECTYPESOwner.RELATIONGROUPTYPE, relationgrouptype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if relationgrouptypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relationgrouptype_.Unstage(relationgrouptypeFormCallback.probe.stageOfInterest)
	}

	relationgrouptypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATIONGROUPTYPE](
		relationgrouptypeFormCallback.probe,
	)
	relationgrouptypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relationgrouptypeFormCallback.CreationMode || relationgrouptypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relationgrouptypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relationgrouptypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATIONGROUPTYPEFormCallback(
			nil,
			relationgrouptypeFormCallback.probe,
			newFormGroup,
		)
		relationgrouptype := new(models.RELATIONGROUPTYPE)
		FillUpForm(relationgrouptype, newFormGroup, relationgrouptypeFormCallback.probe)
		relationgrouptypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relationgrouptypeFormCallback.probe)
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
		case "CORECONTENT":
			FormDivSelectFieldToField(&(reqif_.CORECONTENT), reqifFormCallback.probe.stageOfInterest, formDiv)
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
func __gong__New__REQIFCONTENTFormCallback(
	reqifcontent *models.REQIFCONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (reqifcontentFormCallback *REQIFCONTENTFormCallback) {
	reqifcontentFormCallback = new(REQIFCONTENTFormCallback)
	reqifcontentFormCallback.probe = probe
	reqifcontentFormCallback.reqifcontent = reqifcontent
	reqifcontentFormCallback.formGroup = formGroup

	reqifcontentFormCallback.CreationMode = (reqifcontent == nil)

	return
}

type REQIFCONTENTFormCallback struct {
	reqifcontent *models.REQIFCONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (reqifcontentFormCallback *REQIFCONTENTFormCallback) OnSave() {

	log.Println("REQIFCONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	reqifcontentFormCallback.probe.formStage.Checkout()

	if reqifcontentFormCallback.reqifcontent == nil {
		reqifcontentFormCallback.reqifcontent = new(models.REQIFCONTENT).Stage(reqifcontentFormCallback.probe.stageOfInterest)
	}
	reqifcontent_ := reqifcontentFormCallback.reqifcontent
	_ = reqifcontent_

	for _, formDiv := range reqifcontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(reqifcontent_.Name), formDiv)
		case "DATATYPES":
			FormDivSelectFieldToField(&(reqifcontent_.DATATYPES), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECTYPES":
			FormDivSelectFieldToField(&(reqifcontent_.SPECTYPES), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECOBJECTS":
			FormDivSelectFieldToField(&(reqifcontent_.SPECOBJECTS), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECRELATIONS":
			FormDivSelectFieldToField(&(reqifcontent_.SPECRELATIONS), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFICATIONS":
			FormDivSelectFieldToField(&(reqifcontent_.SPECIFICATIONS), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		case "SPECRELATIONGROUPS":
			FormDivSelectFieldToField(&(reqifcontent_.SPECRELATIONGROUPS), reqifcontentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if reqifcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqifcontent_.Unstage(reqifcontentFormCallback.probe.stageOfInterest)
	}

	reqifcontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQIFCONTENT](
		reqifcontentFormCallback.probe,
	)
	reqifcontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if reqifcontentFormCallback.CreationMode || reqifcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqifcontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(reqifcontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQIFCONTENTFormCallback(
			nil,
			reqifcontentFormCallback.probe,
			newFormGroup,
		)
		reqifcontent := new(models.REQIFCONTENT)
		FillUpForm(reqifcontent, newFormGroup, reqifcontentFormCallback.probe)
		reqifcontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(reqifcontentFormCallback.probe)
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
func __gong__New__REQIFTOOLEXTENSIONFormCallback(
	reqiftoolextension *models.REQIFTOOLEXTENSION,
	probe *Probe,
	formGroup *table.FormGroup,
) (reqiftoolextensionFormCallback *REQIFTOOLEXTENSIONFormCallback) {
	reqiftoolextensionFormCallback = new(REQIFTOOLEXTENSIONFormCallback)
	reqiftoolextensionFormCallback.probe = probe
	reqiftoolextensionFormCallback.reqiftoolextension = reqiftoolextension
	reqiftoolextensionFormCallback.formGroup = formGroup

	reqiftoolextensionFormCallback.CreationMode = (reqiftoolextension == nil)

	return
}

type REQIFTOOLEXTENSIONFormCallback struct {
	reqiftoolextension *models.REQIFTOOLEXTENSION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (reqiftoolextensionFormCallback *REQIFTOOLEXTENSIONFormCallback) OnSave() {

	log.Println("REQIFTOOLEXTENSIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	reqiftoolextensionFormCallback.probe.formStage.Checkout()

	if reqiftoolextensionFormCallback.reqiftoolextension == nil {
		reqiftoolextensionFormCallback.reqiftoolextension = new(models.REQIFTOOLEXTENSION).Stage(reqiftoolextensionFormCallback.probe.stageOfInterest)
	}
	reqiftoolextension_ := reqiftoolextensionFormCallback.reqiftoolextension
	_ = reqiftoolextension_

	for _, formDiv := range reqiftoolextensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(reqiftoolextension_.Name), formDiv)
		case "TOOLEXTENSIONS:REQIFTOOLEXTENSION":
			// we need to retrieve the field owner before the change
			var pastTOOLEXTENSIONSOwner *models.TOOLEXTENSIONS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TOOLEXTENSIONS"
			rf.Fieldname = "REQIFTOOLEXTENSION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				reqiftoolextensionFormCallback.probe.stageOfInterest,
				reqiftoolextensionFormCallback.probe.backRepoOfInterest,
				reqiftoolextension_,
				&rf)

			if reverseFieldOwner != nil {
				pastTOOLEXTENSIONSOwner = reverseFieldOwner.(*models.TOOLEXTENSIONS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTOOLEXTENSIONSOwner != nil {
					idx := slices.Index(pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, reqiftoolextension_)
					pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION = slices.Delete(pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _toolextensions := range *models.GetGongstructInstancesSet[models.TOOLEXTENSIONS](reqiftoolextensionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _toolextensions.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTOOLEXTENSIONSOwner := _toolextensions // we have a match
						if pastTOOLEXTENSIONSOwner != nil {
							if newTOOLEXTENSIONSOwner != pastTOOLEXTENSIONSOwner {
								idx := slices.Index(pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, reqiftoolextension_)
								pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION = slices.Delete(pastTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, idx, idx+1)
								newTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION = append(newTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, reqiftoolextension_)
							}
						} else {
							newTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION = append(newTOOLEXTENSIONSOwner.REQIFTOOLEXTENSION, reqiftoolextension_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if reqiftoolextensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqiftoolextension_.Unstage(reqiftoolextensionFormCallback.probe.stageOfInterest)
	}

	reqiftoolextensionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQIFTOOLEXTENSION](
		reqiftoolextensionFormCallback.probe,
	)
	reqiftoolextensionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if reqiftoolextensionFormCallback.CreationMode || reqiftoolextensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqiftoolextensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(reqiftoolextensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQIFTOOLEXTENSIONFormCallback(
			nil,
			reqiftoolextensionFormCallback.probe,
			newFormGroup,
		)
		reqiftoolextension := new(models.REQIFTOOLEXTENSION)
		FillUpForm(reqiftoolextension, newFormGroup, reqiftoolextensionFormCallback.probe)
		reqiftoolextensionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(reqiftoolextensionFormCallback.probe)
}
func __gong__New__REQTYPEFormCallback(
	reqtype *models.REQTYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (reqtypeFormCallback *REQTYPEFormCallback) {
	reqtypeFormCallback = new(REQTYPEFormCallback)
	reqtypeFormCallback.probe = probe
	reqtypeFormCallback.reqtype = reqtype
	reqtypeFormCallback.formGroup = formGroup

	reqtypeFormCallback.CreationMode = (reqtype == nil)

	return
}

type REQTYPEFormCallback struct {
	reqtype *models.REQTYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (reqtypeFormCallback *REQTYPEFormCallback) OnSave() {

	log.Println("REQTYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	reqtypeFormCallback.probe.formStage.Checkout()

	if reqtypeFormCallback.reqtype == nil {
		reqtypeFormCallback.reqtype = new(models.REQTYPE).Stage(reqtypeFormCallback.probe.stageOfInterest)
	}
	reqtype_ := reqtypeFormCallback.reqtype
	_ = reqtype_

	for _, formDiv := range reqtypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(reqtype_.Name), formDiv)
		case "DATATYPEDEFINITIONBOOLEANREF":
			FormDivBasicFieldToField(&(reqtype_.DATATYPEDEFINITIONBOOLEANREF), formDiv)
		}
	}

	// manage the suppress operation
	if reqtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqtype_.Unstage(reqtypeFormCallback.probe.stageOfInterest)
	}

	reqtypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQTYPE](
		reqtypeFormCallback.probe,
	)
	reqtypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if reqtypeFormCallback.CreationMode || reqtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		reqtypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(reqtypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQTYPEFormCallback(
			nil,
			reqtypeFormCallback.probe,
			newFormGroup,
		)
		reqtype := new(models.REQTYPE)
		FillUpForm(reqtype, newFormGroup, reqtypeFormCallback.probe)
		reqtypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(reqtypeFormCallback.probe)
}
func __gong__New__SOURCEFormCallback(
	source *models.SOURCE,
	probe *Probe,
	formGroup *table.FormGroup,
) (sourceFormCallback *SOURCEFormCallback) {
	sourceFormCallback = new(SOURCEFormCallback)
	sourceFormCallback.probe = probe
	sourceFormCallback.source = source
	sourceFormCallback.formGroup = formGroup

	sourceFormCallback.CreationMode = (source == nil)

	return
}

type SOURCEFormCallback struct {
	source *models.SOURCE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sourceFormCallback *SOURCEFormCallback) OnSave() {

	log.Println("SOURCEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sourceFormCallback.probe.formStage.Checkout()

	if sourceFormCallback.source == nil {
		sourceFormCallback.source = new(models.SOURCE).Stage(sourceFormCallback.probe.stageOfInterest)
	}
	source_ := sourceFormCallback.source
	_ = source_

	for _, formDiv := range sourceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(source_.Name), formDiv)
		case "SPECOBJECTREF":
			FormDivBasicFieldToField(&(source_.SPECOBJECTREF), formDiv)
		}
	}

	// manage the suppress operation
	if sourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		source_.Unstage(sourceFormCallback.probe.stageOfInterest)
	}

	sourceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SOURCE](
		sourceFormCallback.probe,
	)
	sourceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sourceFormCallback.CreationMode || sourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sourceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sourceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SOURCEFormCallback(
			nil,
			sourceFormCallback.probe,
			newFormGroup,
		)
		source := new(models.SOURCE)
		FillUpForm(source, newFormGroup, sourceFormCallback.probe)
		sourceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sourceFormCallback.probe)
}
func __gong__New__SOURCESPECIFICATIONFormCallback(
	sourcespecification *models.SOURCESPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (sourcespecificationFormCallback *SOURCESPECIFICATIONFormCallback) {
	sourcespecificationFormCallback = new(SOURCESPECIFICATIONFormCallback)
	sourcespecificationFormCallback.probe = probe
	sourcespecificationFormCallback.sourcespecification = sourcespecification
	sourcespecificationFormCallback.formGroup = formGroup

	sourcespecificationFormCallback.CreationMode = (sourcespecification == nil)

	return
}

type SOURCESPECIFICATIONFormCallback struct {
	sourcespecification *models.SOURCESPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sourcespecificationFormCallback *SOURCESPECIFICATIONFormCallback) OnSave() {

	log.Println("SOURCESPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sourcespecificationFormCallback.probe.formStage.Checkout()

	if sourcespecificationFormCallback.sourcespecification == nil {
		sourcespecificationFormCallback.sourcespecification = new(models.SOURCESPECIFICATION).Stage(sourcespecificationFormCallback.probe.stageOfInterest)
	}
	sourcespecification_ := sourcespecificationFormCallback.sourcespecification
	_ = sourcespecification_

	for _, formDiv := range sourcespecificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sourcespecification_.Name), formDiv)
		case "SPECIFICATIONREF":
			FormDivBasicFieldToField(&(sourcespecification_.SPECIFICATIONREF), formDiv)
		}
	}

	// manage the suppress operation
	if sourcespecificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sourcespecification_.Unstage(sourcespecificationFormCallback.probe.stageOfInterest)
	}

	sourcespecificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SOURCESPECIFICATION](
		sourcespecificationFormCallback.probe,
	)
	sourcespecificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sourcespecificationFormCallback.CreationMode || sourcespecificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sourcespecificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sourcespecificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SOURCESPECIFICATIONFormCallback(
			nil,
			sourcespecificationFormCallback.probe,
			newFormGroup,
		)
		sourcespecification := new(models.SOURCESPECIFICATION)
		FillUpForm(sourcespecification, newFormGroup, sourcespecificationFormCallback.probe)
		sourcespecificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sourcespecificationFormCallback.probe)
}
func __gong__New__SPECATTRIBUTESFormCallback(
	specattributes *models.SPECATTRIBUTES,
	probe *Probe,
	formGroup *table.FormGroup,
) (specattributesFormCallback *SPECATTRIBUTESFormCallback) {
	specattributesFormCallback = new(SPECATTRIBUTESFormCallback)
	specattributesFormCallback.probe = probe
	specattributesFormCallback.specattributes = specattributes
	specattributesFormCallback.formGroup = formGroup

	specattributesFormCallback.CreationMode = (specattributes == nil)

	return
}

type SPECATTRIBUTESFormCallback struct {
	specattributes *models.SPECATTRIBUTES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specattributesFormCallback *SPECATTRIBUTESFormCallback) OnSave() {

	log.Println("SPECATTRIBUTESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specattributesFormCallback.probe.formStage.Checkout()

	if specattributesFormCallback.specattributes == nil {
		specattributesFormCallback.specattributes = new(models.SPECATTRIBUTES).Stage(specattributesFormCallback.probe.stageOfInterest)
	}
	specattributes_ := specattributesFormCallback.specattributes
	_ = specattributes_

	for _, formDiv := range specattributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specattributes_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specattributes_.Unstage(specattributesFormCallback.probe.stageOfInterest)
	}

	specattributesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECATTRIBUTES](
		specattributesFormCallback.probe,
	)
	specattributesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specattributesFormCallback.CreationMode || specattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specattributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specattributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECATTRIBUTESFormCallback(
			nil,
			specattributesFormCallback.probe,
			newFormGroup,
		)
		specattributes := new(models.SPECATTRIBUTES)
		FillUpForm(specattributes, newFormGroup, specattributesFormCallback.probe)
		specattributesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specattributesFormCallback.probe)
}
func __gong__New__SPECHIERARCHYFormCallback(
	spechierarchy *models.SPECHIERARCHY,
	probe *Probe,
	formGroup *table.FormGroup,
) (spechierarchyFormCallback *SPECHIERARCHYFormCallback) {
	spechierarchyFormCallback = new(SPECHIERARCHYFormCallback)
	spechierarchyFormCallback.probe = probe
	spechierarchyFormCallback.spechierarchy = spechierarchy
	spechierarchyFormCallback.formGroup = formGroup

	spechierarchyFormCallback.CreationMode = (spechierarchy == nil)

	return
}

type SPECHIERARCHYFormCallback struct {
	spechierarchy *models.SPECHIERARCHY

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spechierarchyFormCallback *SPECHIERARCHYFormCallback) OnSave() {

	log.Println("SPECHIERARCHYFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spechierarchyFormCallback.probe.formStage.Checkout()

	if spechierarchyFormCallback.spechierarchy == nil {
		spechierarchyFormCallback.spechierarchy = new(models.SPECHIERARCHY).Stage(spechierarchyFormCallback.probe.stageOfInterest)
	}
	spechierarchy_ := spechierarchyFormCallback.spechierarchy
	_ = spechierarchy_

	for _, formDiv := range spechierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spechierarchy_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(spechierarchy_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(spechierarchy_.IDENTIFIERAttr), formDiv)
		case "ISEDITABLEAttr":
			FormDivBasicFieldToField(&(spechierarchy_.ISEDITABLEAttr), formDiv)
		case "ISTABLEINTERNALAttr":
			FormDivBasicFieldToField(&(spechierarchy_.ISTABLEINTERNALAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(spechierarchy_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(spechierarchy_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(spechierarchy_.ALTERNATIVEID), spechierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(spechierarchy_.CHILDREN), spechierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "EDITABLEATTS":
			FormDivSelectFieldToField(&(spechierarchy_.EDITABLEATTS), spechierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "OBJECT":
			FormDivSelectFieldToField(&(spechierarchy_.OBJECT), spechierarchyFormCallback.probe.stageOfInterest, formDiv)
		case "CHILDREN:SPECHIERARCHY":
			// we need to retrieve the field owner before the change
			var pastCHILDRENOwner *models.CHILDREN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "CHILDREN"
			rf.Fieldname = "SPECHIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spechierarchyFormCallback.probe.stageOfInterest,
				spechierarchyFormCallback.probe.backRepoOfInterest,
				spechierarchy_,
				&rf)

			if reverseFieldOwner != nil {
				pastCHILDRENOwner = reverseFieldOwner.(*models.CHILDREN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCHILDRENOwner != nil {
					idx := slices.Index(pastCHILDRENOwner.SPECHIERARCHY, spechierarchy_)
					pastCHILDRENOwner.SPECHIERARCHY = slices.Delete(pastCHILDRENOwner.SPECHIERARCHY, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _children := range *models.GetGongstructInstancesSet[models.CHILDREN](spechierarchyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _children.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCHILDRENOwner := _children // we have a match
						if pastCHILDRENOwner != nil {
							if newCHILDRENOwner != pastCHILDRENOwner {
								idx := slices.Index(pastCHILDRENOwner.SPECHIERARCHY, spechierarchy_)
								pastCHILDRENOwner.SPECHIERARCHY = slices.Delete(pastCHILDRENOwner.SPECHIERARCHY, idx, idx+1)
								newCHILDRENOwner.SPECHIERARCHY = append(newCHILDRENOwner.SPECHIERARCHY, spechierarchy_)
							}
						} else {
							newCHILDRENOwner.SPECHIERARCHY = append(newCHILDRENOwner.SPECHIERARCHY, spechierarchy_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if spechierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spechierarchy_.Unstage(spechierarchyFormCallback.probe.stageOfInterest)
	}

	spechierarchyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECHIERARCHY](
		spechierarchyFormCallback.probe,
	)
	spechierarchyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spechierarchyFormCallback.CreationMode || spechierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spechierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spechierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECHIERARCHYFormCallback(
			nil,
			spechierarchyFormCallback.probe,
			newFormGroup,
		)
		spechierarchy := new(models.SPECHIERARCHY)
		FillUpForm(spechierarchy, newFormGroup, spechierarchyFormCallback.probe)
		spechierarchyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spechierarchyFormCallback.probe)
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
		case "DESCAttr":
			FormDivBasicFieldToField(&(specification_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specification_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specification_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specification_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specification_.ALTERNATIVEID), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(specification_.VALUES), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "CHILDREN":
			FormDivSelectFieldToField(&(specification_.CHILDREN), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(specification_.TYPE), specificationFormCallback.probe.stageOfInterest, formDiv)
		case "SPECIFICATIONS:SPECIFICATION":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONSOwner *models.SPECIFICATIONS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATIONS"
			rf.Fieldname = "SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specificationFormCallback.probe.stageOfInterest,
				specificationFormCallback.probe.backRepoOfInterest,
				specification_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONSOwner = reverseFieldOwner.(*models.SPECIFICATIONS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONSOwner != nil {
					idx := slices.Index(pastSPECIFICATIONSOwner.SPECIFICATION, specification_)
					pastSPECIFICATIONSOwner.SPECIFICATION = slices.Delete(pastSPECIFICATIONSOwner.SPECIFICATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specifications := range *models.GetGongstructInstancesSet[models.SPECIFICATIONS](specificationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specifications.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONSOwner := _specifications // we have a match
						if pastSPECIFICATIONSOwner != nil {
							if newSPECIFICATIONSOwner != pastSPECIFICATIONSOwner {
								idx := slices.Index(pastSPECIFICATIONSOwner.SPECIFICATION, specification_)
								pastSPECIFICATIONSOwner.SPECIFICATION = slices.Delete(pastSPECIFICATIONSOwner.SPECIFICATION, idx, idx+1)
								newSPECIFICATIONSOwner.SPECIFICATION = append(newSPECIFICATIONSOwner.SPECIFICATION, specification_)
							}
						} else {
							newSPECIFICATIONSOwner.SPECIFICATION = append(newSPECIFICATIONSOwner.SPECIFICATION, specification_)
						}
					}
				}
			}
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
func __gong__New__SPECIFICATIONSFormCallback(
	specifications *models.SPECIFICATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationsFormCallback *SPECIFICATIONSFormCallback) {
	specificationsFormCallback = new(SPECIFICATIONSFormCallback)
	specificationsFormCallback.probe = probe
	specificationsFormCallback.specifications = specifications
	specificationsFormCallback.formGroup = formGroup

	specificationsFormCallback.CreationMode = (specifications == nil)

	return
}

type SPECIFICATIONSFormCallback struct {
	specifications *models.SPECIFICATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationsFormCallback *SPECIFICATIONSFormCallback) OnSave() {

	log.Println("SPECIFICATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationsFormCallback.probe.formStage.Checkout()

	if specificationsFormCallback.specifications == nil {
		specificationsFormCallback.specifications = new(models.SPECIFICATIONS).Stage(specificationsFormCallback.probe.stageOfInterest)
	}
	specifications_ := specificationsFormCallback.specifications
	_ = specifications_

	for _, formDiv := range specificationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specifications_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specifications_.Unstage(specificationsFormCallback.probe.stageOfInterest)
	}

	specificationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATIONS](
		specificationsFormCallback.probe,
	)
	specificationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specificationsFormCallback.CreationMode || specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specificationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONSFormCallback(
			nil,
			specificationsFormCallback.probe,
			newFormGroup,
		)
		specifications := new(models.SPECIFICATIONS)
		FillUpForm(specifications, newFormGroup, specificationsFormCallback.probe)
		specificationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specificationsFormCallback.probe)
}
func __gong__New__SPECIFICATIONTYPEFormCallback(
	specificationtype *models.SPECIFICATIONTYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationtypeFormCallback *SPECIFICATIONTYPEFormCallback) {
	specificationtypeFormCallback = new(SPECIFICATIONTYPEFormCallback)
	specificationtypeFormCallback.probe = probe
	specificationtypeFormCallback.specificationtype = specificationtype
	specificationtypeFormCallback.formGroup = formGroup

	specificationtypeFormCallback.CreationMode = (specificationtype == nil)

	return
}

type SPECIFICATIONTYPEFormCallback struct {
	specificationtype *models.SPECIFICATIONTYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationtypeFormCallback *SPECIFICATIONTYPEFormCallback) OnSave() {

	log.Println("SPECIFICATIONTYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationtypeFormCallback.probe.formStage.Checkout()

	if specificationtypeFormCallback.specificationtype == nil {
		specificationtypeFormCallback.specificationtype = new(models.SPECIFICATIONTYPE).Stage(specificationtypeFormCallback.probe.stageOfInterest)
	}
	specificationtype_ := specificationtypeFormCallback.specificationtype
	_ = specificationtype_

	for _, formDiv := range specificationtypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specificationtype_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(specificationtype_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specificationtype_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specificationtype_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specificationtype_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specificationtype_.ALTERNATIVEID), specificationtypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES":
			FormDivSelectFieldToField(&(specificationtype_.SPECATTRIBUTES), specificationtypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECTYPES:SPECIFICATIONTYPE":
			// we need to retrieve the field owner before the change
			var pastSPECTYPESOwner *models.SPECTYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECIFICATIONTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specificationtypeFormCallback.probe.stageOfInterest,
				specificationtypeFormCallback.probe.backRepoOfInterest,
				specificationtype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECTYPESOwner = reverseFieldOwner.(*models.SPECTYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECTYPESOwner != nil {
					idx := slices.Index(pastSPECTYPESOwner.SPECIFICATIONTYPE, specificationtype_)
					pastSPECTYPESOwner.SPECIFICATIONTYPE = slices.Delete(pastSPECTYPESOwner.SPECIFICATIONTYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spectypes := range *models.GetGongstructInstancesSet[models.SPECTYPES](specificationtypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spectypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECTYPESOwner := _spectypes // we have a match
						if pastSPECTYPESOwner != nil {
							if newSPECTYPESOwner != pastSPECTYPESOwner {
								idx := slices.Index(pastSPECTYPESOwner.SPECIFICATIONTYPE, specificationtype_)
								pastSPECTYPESOwner.SPECIFICATIONTYPE = slices.Delete(pastSPECTYPESOwner.SPECIFICATIONTYPE, idx, idx+1)
								newSPECTYPESOwner.SPECIFICATIONTYPE = append(newSPECTYPESOwner.SPECIFICATIONTYPE, specificationtype_)
							}
						} else {
							newSPECTYPESOwner.SPECIFICATIONTYPE = append(newSPECTYPESOwner.SPECIFICATIONTYPE, specificationtype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specificationtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationtype_.Unstage(specificationtypeFormCallback.probe.stageOfInterest)
	}

	specificationtypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATIONTYPE](
		specificationtypeFormCallback.probe,
	)
	specificationtypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specificationtypeFormCallback.CreationMode || specificationtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationtypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specificationtypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONTYPEFormCallback(
			nil,
			specificationtypeFormCallback.probe,
			newFormGroup,
		)
		specificationtype := new(models.SPECIFICATIONTYPE)
		FillUpForm(specificationtype, newFormGroup, specificationtypeFormCallback.probe)
		specificationtypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specificationtypeFormCallback.probe)
}
func __gong__New__SPECIFIEDVALUESFormCallback(
	specifiedvalues *models.SPECIFIEDVALUES,
	probe *Probe,
	formGroup *table.FormGroup,
) (specifiedvaluesFormCallback *SPECIFIEDVALUESFormCallback) {
	specifiedvaluesFormCallback = new(SPECIFIEDVALUESFormCallback)
	specifiedvaluesFormCallback.probe = probe
	specifiedvaluesFormCallback.specifiedvalues = specifiedvalues
	specifiedvaluesFormCallback.formGroup = formGroup

	specifiedvaluesFormCallback.CreationMode = (specifiedvalues == nil)

	return
}

type SPECIFIEDVALUESFormCallback struct {
	specifiedvalues *models.SPECIFIEDVALUES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specifiedvaluesFormCallback *SPECIFIEDVALUESFormCallback) OnSave() {

	log.Println("SPECIFIEDVALUESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specifiedvaluesFormCallback.probe.formStage.Checkout()

	if specifiedvaluesFormCallback.specifiedvalues == nil {
		specifiedvaluesFormCallback.specifiedvalues = new(models.SPECIFIEDVALUES).Stage(specifiedvaluesFormCallback.probe.stageOfInterest)
	}
	specifiedvalues_ := specifiedvaluesFormCallback.specifiedvalues
	_ = specifiedvalues_

	for _, formDiv := range specifiedvaluesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specifiedvalues_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specifiedvaluesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specifiedvalues_.Unstage(specifiedvaluesFormCallback.probe.stageOfInterest)
	}

	specifiedvaluesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFIEDVALUES](
		specifiedvaluesFormCallback.probe,
	)
	specifiedvaluesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specifiedvaluesFormCallback.CreationMode || specifiedvaluesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specifiedvaluesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specifiedvaluesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFIEDVALUESFormCallback(
			nil,
			specifiedvaluesFormCallback.probe,
			newFormGroup,
		)
		specifiedvalues := new(models.SPECIFIEDVALUES)
		FillUpForm(specifiedvalues, newFormGroup, specifiedvaluesFormCallback.probe)
		specifiedvaluesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specifiedvaluesFormCallback.probe)
}
func __gong__New__SPECOBJECTFormCallback(
	specobject *models.SPECOBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (specobjectFormCallback *SPECOBJECTFormCallback) {
	specobjectFormCallback = new(SPECOBJECTFormCallback)
	specobjectFormCallback.probe = probe
	specobjectFormCallback.specobject = specobject
	specobjectFormCallback.formGroup = formGroup

	specobjectFormCallback.CreationMode = (specobject == nil)

	return
}

type SPECOBJECTFormCallback struct {
	specobject *models.SPECOBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specobjectFormCallback *SPECOBJECTFormCallback) OnSave() {

	log.Println("SPECOBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specobjectFormCallback.probe.formStage.Checkout()

	if specobjectFormCallback.specobject == nil {
		specobjectFormCallback.specobject = new(models.SPECOBJECT).Stage(specobjectFormCallback.probe.stageOfInterest)
	}
	specobject_ := specobjectFormCallback.specobject
	_ = specobject_

	for _, formDiv := range specobjectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specobject_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(specobject_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specobject_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specobject_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specobject_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specobject_.ALTERNATIVEID), specobjectFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(specobject_.VALUES), specobjectFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(specobject_.TYPE), specobjectFormCallback.probe.stageOfInterest, formDiv)
		case "SPECOBJECTS:SPECOBJECT":
			// we need to retrieve the field owner before the change
			var pastSPECOBJECTSOwner *models.SPECOBJECTS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECOBJECTS"
			rf.Fieldname = "SPECOBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specobjectFormCallback.probe.stageOfInterest,
				specobjectFormCallback.probe.backRepoOfInterest,
				specobject_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECOBJECTSOwner = reverseFieldOwner.(*models.SPECOBJECTS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECOBJECTSOwner != nil {
					idx := slices.Index(pastSPECOBJECTSOwner.SPECOBJECT, specobject_)
					pastSPECOBJECTSOwner.SPECOBJECT = slices.Delete(pastSPECOBJECTSOwner.SPECOBJECT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specobjects := range *models.GetGongstructInstancesSet[models.SPECOBJECTS](specobjectFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specobjects.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECOBJECTSOwner := _specobjects // we have a match
						if pastSPECOBJECTSOwner != nil {
							if newSPECOBJECTSOwner != pastSPECOBJECTSOwner {
								idx := slices.Index(pastSPECOBJECTSOwner.SPECOBJECT, specobject_)
								pastSPECOBJECTSOwner.SPECOBJECT = slices.Delete(pastSPECOBJECTSOwner.SPECOBJECT, idx, idx+1)
								newSPECOBJECTSOwner.SPECOBJECT = append(newSPECOBJECTSOwner.SPECOBJECT, specobject_)
							}
						} else {
							newSPECOBJECTSOwner.SPECOBJECT = append(newSPECOBJECTSOwner.SPECOBJECT, specobject_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobject_.Unstage(specobjectFormCallback.probe.stageOfInterest)
	}

	specobjectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECOBJECT](
		specobjectFormCallback.probe,
	)
	specobjectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specobjectFormCallback.CreationMode || specobjectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobjectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specobjectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECOBJECTFormCallback(
			nil,
			specobjectFormCallback.probe,
			newFormGroup,
		)
		specobject := new(models.SPECOBJECT)
		FillUpForm(specobject, newFormGroup, specobjectFormCallback.probe)
		specobjectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specobjectFormCallback.probe)
}
func __gong__New__SPECOBJECTSFormCallback(
	specobjects *models.SPECOBJECTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (specobjectsFormCallback *SPECOBJECTSFormCallback) {
	specobjectsFormCallback = new(SPECOBJECTSFormCallback)
	specobjectsFormCallback.probe = probe
	specobjectsFormCallback.specobjects = specobjects
	specobjectsFormCallback.formGroup = formGroup

	specobjectsFormCallback.CreationMode = (specobjects == nil)

	return
}

type SPECOBJECTSFormCallback struct {
	specobjects *models.SPECOBJECTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specobjectsFormCallback *SPECOBJECTSFormCallback) OnSave() {

	log.Println("SPECOBJECTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specobjectsFormCallback.probe.formStage.Checkout()

	if specobjectsFormCallback.specobjects == nil {
		specobjectsFormCallback.specobjects = new(models.SPECOBJECTS).Stage(specobjectsFormCallback.probe.stageOfInterest)
	}
	specobjects_ := specobjectsFormCallback.specobjects
	_ = specobjects_

	for _, formDiv := range specobjectsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specobjects_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specobjectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobjects_.Unstage(specobjectsFormCallback.probe.stageOfInterest)
	}

	specobjectsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECOBJECTS](
		specobjectsFormCallback.probe,
	)
	specobjectsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specobjectsFormCallback.CreationMode || specobjectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobjectsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specobjectsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECOBJECTSFormCallback(
			nil,
			specobjectsFormCallback.probe,
			newFormGroup,
		)
		specobjects := new(models.SPECOBJECTS)
		FillUpForm(specobjects, newFormGroup, specobjectsFormCallback.probe)
		specobjectsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specobjectsFormCallback.probe)
}
func __gong__New__SPECOBJECTTYPEFormCallback(
	specobjecttype *models.SPECOBJECTTYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specobjecttypeFormCallback *SPECOBJECTTYPEFormCallback) {
	specobjecttypeFormCallback = new(SPECOBJECTTYPEFormCallback)
	specobjecttypeFormCallback.probe = probe
	specobjecttypeFormCallback.specobjecttype = specobjecttype
	specobjecttypeFormCallback.formGroup = formGroup

	specobjecttypeFormCallback.CreationMode = (specobjecttype == nil)

	return
}

type SPECOBJECTTYPEFormCallback struct {
	specobjecttype *models.SPECOBJECTTYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specobjecttypeFormCallback *SPECOBJECTTYPEFormCallback) OnSave() {

	log.Println("SPECOBJECTTYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specobjecttypeFormCallback.probe.formStage.Checkout()

	if specobjecttypeFormCallback.specobjecttype == nil {
		specobjecttypeFormCallback.specobjecttype = new(models.SPECOBJECTTYPE).Stage(specobjecttypeFormCallback.probe.stageOfInterest)
	}
	specobjecttype_ := specobjecttypeFormCallback.specobjecttype
	_ = specobjecttype_

	for _, formDiv := range specobjecttypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specobjecttype_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(specobjecttype_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specobjecttype_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specobjecttype_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specobjecttype_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specobjecttype_.ALTERNATIVEID), specobjecttypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES":
			FormDivSelectFieldToField(&(specobjecttype_.SPECATTRIBUTES), specobjecttypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECTYPES:SPECOBJECTTYPE":
			// we need to retrieve the field owner before the change
			var pastSPECTYPESOwner *models.SPECTYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECOBJECTTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specobjecttypeFormCallback.probe.stageOfInterest,
				specobjecttypeFormCallback.probe.backRepoOfInterest,
				specobjecttype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECTYPESOwner = reverseFieldOwner.(*models.SPECTYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECTYPESOwner != nil {
					idx := slices.Index(pastSPECTYPESOwner.SPECOBJECTTYPE, specobjecttype_)
					pastSPECTYPESOwner.SPECOBJECTTYPE = slices.Delete(pastSPECTYPESOwner.SPECOBJECTTYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spectypes := range *models.GetGongstructInstancesSet[models.SPECTYPES](specobjecttypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spectypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECTYPESOwner := _spectypes // we have a match
						if pastSPECTYPESOwner != nil {
							if newSPECTYPESOwner != pastSPECTYPESOwner {
								idx := slices.Index(pastSPECTYPESOwner.SPECOBJECTTYPE, specobjecttype_)
								pastSPECTYPESOwner.SPECOBJECTTYPE = slices.Delete(pastSPECTYPESOwner.SPECOBJECTTYPE, idx, idx+1)
								newSPECTYPESOwner.SPECOBJECTTYPE = append(newSPECTYPESOwner.SPECOBJECTTYPE, specobjecttype_)
							}
						} else {
							newSPECTYPESOwner.SPECOBJECTTYPE = append(newSPECTYPESOwner.SPECOBJECTTYPE, specobjecttype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specobjecttypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobjecttype_.Unstage(specobjecttypeFormCallback.probe.stageOfInterest)
	}

	specobjecttypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECOBJECTTYPE](
		specobjecttypeFormCallback.probe,
	)
	specobjecttypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specobjecttypeFormCallback.CreationMode || specobjecttypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specobjecttypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specobjecttypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECOBJECTTYPEFormCallback(
			nil,
			specobjecttypeFormCallback.probe,
			newFormGroup,
		)
		specobjecttype := new(models.SPECOBJECTTYPE)
		FillUpForm(specobjecttype, newFormGroup, specobjecttypeFormCallback.probe)
		specobjecttypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specobjecttypeFormCallback.probe)
}
func __gong__New__SPECRELATIONFormCallback(
	specrelation *models.SPECRELATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (specrelationFormCallback *SPECRELATIONFormCallback) {
	specrelationFormCallback = new(SPECRELATIONFormCallback)
	specrelationFormCallback.probe = probe
	specrelationFormCallback.specrelation = specrelation
	specrelationFormCallback.formGroup = formGroup

	specrelationFormCallback.CreationMode = (specrelation == nil)

	return
}

type SPECRELATIONFormCallback struct {
	specrelation *models.SPECRELATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specrelationFormCallback *SPECRELATIONFormCallback) OnSave() {

	log.Println("SPECRELATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specrelationFormCallback.probe.formStage.Checkout()

	if specrelationFormCallback.specrelation == nil {
		specrelationFormCallback.specrelation = new(models.SPECRELATION).Stage(specrelationFormCallback.probe.stageOfInterest)
	}
	specrelation_ := specrelationFormCallback.specrelation
	_ = specrelation_

	for _, formDiv := range specrelationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specrelation_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(specrelation_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specrelation_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specrelation_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specrelation_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specrelation_.ALTERNATIVEID), specrelationFormCallback.probe.stageOfInterest, formDiv)
		case "VALUES":
			FormDivSelectFieldToField(&(specrelation_.VALUES), specrelationFormCallback.probe.stageOfInterest, formDiv)
		case "SOURCE":
			FormDivSelectFieldToField(&(specrelation_.SOURCE), specrelationFormCallback.probe.stageOfInterest, formDiv)
		case "TARGET":
			FormDivSelectFieldToField(&(specrelation_.TARGET), specrelationFormCallback.probe.stageOfInterest, formDiv)
		case "TYPE":
			FormDivSelectFieldToField(&(specrelation_.TYPE), specrelationFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if specrelationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelation_.Unstage(specrelationFormCallback.probe.stageOfInterest)
	}

	specrelationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECRELATION](
		specrelationFormCallback.probe,
	)
	specrelationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specrelationFormCallback.CreationMode || specrelationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specrelationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECRELATIONFormCallback(
			nil,
			specrelationFormCallback.probe,
			newFormGroup,
		)
		specrelation := new(models.SPECRELATION)
		FillUpForm(specrelation, newFormGroup, specrelationFormCallback.probe)
		specrelationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specrelationFormCallback.probe)
}
func __gong__New__SPECRELATIONGROUPSFormCallback(
	specrelationgroups *models.SPECRELATIONGROUPS,
	probe *Probe,
	formGroup *table.FormGroup,
) (specrelationgroupsFormCallback *SPECRELATIONGROUPSFormCallback) {
	specrelationgroupsFormCallback = new(SPECRELATIONGROUPSFormCallback)
	specrelationgroupsFormCallback.probe = probe
	specrelationgroupsFormCallback.specrelationgroups = specrelationgroups
	specrelationgroupsFormCallback.formGroup = formGroup

	specrelationgroupsFormCallback.CreationMode = (specrelationgroups == nil)

	return
}

type SPECRELATIONGROUPSFormCallback struct {
	specrelationgroups *models.SPECRELATIONGROUPS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specrelationgroupsFormCallback *SPECRELATIONGROUPSFormCallback) OnSave() {

	log.Println("SPECRELATIONGROUPSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specrelationgroupsFormCallback.probe.formStage.Checkout()

	if specrelationgroupsFormCallback.specrelationgroups == nil {
		specrelationgroupsFormCallback.specrelationgroups = new(models.SPECRELATIONGROUPS).Stage(specrelationgroupsFormCallback.probe.stageOfInterest)
	}
	specrelationgroups_ := specrelationgroupsFormCallback.specrelationgroups
	_ = specrelationgroups_

	for _, formDiv := range specrelationgroupsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specrelationgroups_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specrelationgroupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationgroups_.Unstage(specrelationgroupsFormCallback.probe.stageOfInterest)
	}

	specrelationgroupsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECRELATIONGROUPS](
		specrelationgroupsFormCallback.probe,
	)
	specrelationgroupsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specrelationgroupsFormCallback.CreationMode || specrelationgroupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationgroupsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specrelationgroupsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECRELATIONGROUPSFormCallback(
			nil,
			specrelationgroupsFormCallback.probe,
			newFormGroup,
		)
		specrelationgroups := new(models.SPECRELATIONGROUPS)
		FillUpForm(specrelationgroups, newFormGroup, specrelationgroupsFormCallback.probe)
		specrelationgroupsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specrelationgroupsFormCallback.probe)
}
func __gong__New__SPECRELATIONSFormCallback(
	specrelations *models.SPECRELATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (specrelationsFormCallback *SPECRELATIONSFormCallback) {
	specrelationsFormCallback = new(SPECRELATIONSFormCallback)
	specrelationsFormCallback.probe = probe
	specrelationsFormCallback.specrelations = specrelations
	specrelationsFormCallback.formGroup = formGroup

	specrelationsFormCallback.CreationMode = (specrelations == nil)

	return
}

type SPECRELATIONSFormCallback struct {
	specrelations *models.SPECRELATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specrelationsFormCallback *SPECRELATIONSFormCallback) OnSave() {

	log.Println("SPECRELATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specrelationsFormCallback.probe.formStage.Checkout()

	if specrelationsFormCallback.specrelations == nil {
		specrelationsFormCallback.specrelations = new(models.SPECRELATIONS).Stage(specrelationsFormCallback.probe.stageOfInterest)
	}
	specrelations_ := specrelationsFormCallback.specrelations
	_ = specrelations_

	for _, formDiv := range specrelationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specrelations_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if specrelationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelations_.Unstage(specrelationsFormCallback.probe.stageOfInterest)
	}

	specrelationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECRELATIONS](
		specrelationsFormCallback.probe,
	)
	specrelationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specrelationsFormCallback.CreationMode || specrelationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specrelationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECRELATIONSFormCallback(
			nil,
			specrelationsFormCallback.probe,
			newFormGroup,
		)
		specrelations := new(models.SPECRELATIONS)
		FillUpForm(specrelations, newFormGroup, specrelationsFormCallback.probe)
		specrelationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specrelationsFormCallback.probe)
}
func __gong__New__SPECRELATIONTYPEFormCallback(
	specrelationtype *models.SPECRELATIONTYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specrelationtypeFormCallback *SPECRELATIONTYPEFormCallback) {
	specrelationtypeFormCallback = new(SPECRELATIONTYPEFormCallback)
	specrelationtypeFormCallback.probe = probe
	specrelationtypeFormCallback.specrelationtype = specrelationtype
	specrelationtypeFormCallback.formGroup = formGroup

	specrelationtypeFormCallback.CreationMode = (specrelationtype == nil)

	return
}

type SPECRELATIONTYPEFormCallback struct {
	specrelationtype *models.SPECRELATIONTYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specrelationtypeFormCallback *SPECRELATIONTYPEFormCallback) OnSave() {

	log.Println("SPECRELATIONTYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specrelationtypeFormCallback.probe.formStage.Checkout()

	if specrelationtypeFormCallback.specrelationtype == nil {
		specrelationtypeFormCallback.specrelationtype = new(models.SPECRELATIONTYPE).Stage(specrelationtypeFormCallback.probe.stageOfInterest)
	}
	specrelationtype_ := specrelationtypeFormCallback.specrelationtype
	_ = specrelationtype_

	for _, formDiv := range specrelationtypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specrelationtype_.Name), formDiv)
		case "DESCAttr":
			FormDivBasicFieldToField(&(specrelationtype_.DESCAttr), formDiv)
		case "IDENTIFIERAttr":
			FormDivBasicFieldToField(&(specrelationtype_.IDENTIFIERAttr), formDiv)
		case "LASTCHANGEAttr":
			FormDivBasicFieldToField(&(specrelationtype_.LASTCHANGEAttr), formDiv)
		case "LONGNAMEAttr":
			FormDivBasicFieldToField(&(specrelationtype_.LONGNAMEAttr), formDiv)
		case "ALTERNATIVEID":
			FormDivSelectFieldToField(&(specrelationtype_.ALTERNATIVEID), specrelationtypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECATTRIBUTES":
			FormDivSelectFieldToField(&(specrelationtype_.SPECATTRIBUTES), specrelationtypeFormCallback.probe.stageOfInterest, formDiv)
		case "SPECTYPES:SPECRELATIONTYPE":
			// we need to retrieve the field owner before the change
			var pastSPECTYPESOwner *models.SPECTYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECRELATIONTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specrelationtypeFormCallback.probe.stageOfInterest,
				specrelationtypeFormCallback.probe.backRepoOfInterest,
				specrelationtype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECTYPESOwner = reverseFieldOwner.(*models.SPECTYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECTYPESOwner != nil {
					idx := slices.Index(pastSPECTYPESOwner.SPECRELATIONTYPE, specrelationtype_)
					pastSPECTYPESOwner.SPECRELATIONTYPE = slices.Delete(pastSPECTYPESOwner.SPECRELATIONTYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spectypes := range *models.GetGongstructInstancesSet[models.SPECTYPES](specrelationtypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spectypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECTYPESOwner := _spectypes // we have a match
						if pastSPECTYPESOwner != nil {
							if newSPECTYPESOwner != pastSPECTYPESOwner {
								idx := slices.Index(pastSPECTYPESOwner.SPECRELATIONTYPE, specrelationtype_)
								pastSPECTYPESOwner.SPECRELATIONTYPE = slices.Delete(pastSPECTYPESOwner.SPECRELATIONTYPE, idx, idx+1)
								newSPECTYPESOwner.SPECRELATIONTYPE = append(newSPECTYPESOwner.SPECRELATIONTYPE, specrelationtype_)
							}
						} else {
							newSPECTYPESOwner.SPECRELATIONTYPE = append(newSPECTYPESOwner.SPECRELATIONTYPE, specrelationtype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if specrelationtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationtype_.Unstage(specrelationtypeFormCallback.probe.stageOfInterest)
	}

	specrelationtypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECRELATIONTYPE](
		specrelationtypeFormCallback.probe,
	)
	specrelationtypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specrelationtypeFormCallback.CreationMode || specrelationtypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specrelationtypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specrelationtypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECRELATIONTYPEFormCallback(
			nil,
			specrelationtypeFormCallback.probe,
			newFormGroup,
		)
		specrelationtype := new(models.SPECRELATIONTYPE)
		FillUpForm(specrelationtype, newFormGroup, specrelationtypeFormCallback.probe)
		specrelationtypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specrelationtypeFormCallback.probe)
}
func __gong__New__SPECTYPESFormCallback(
	spectypes *models.SPECTYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (spectypesFormCallback *SPECTYPESFormCallback) {
	spectypesFormCallback = new(SPECTYPESFormCallback)
	spectypesFormCallback.probe = probe
	spectypesFormCallback.spectypes = spectypes
	spectypesFormCallback.formGroup = formGroup

	spectypesFormCallback.CreationMode = (spectypes == nil)

	return
}

type SPECTYPESFormCallback struct {
	spectypes *models.SPECTYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spectypesFormCallback *SPECTYPESFormCallback) OnSave() {

	log.Println("SPECTYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spectypesFormCallback.probe.formStage.Checkout()

	if spectypesFormCallback.spectypes == nil {
		spectypesFormCallback.spectypes = new(models.SPECTYPES).Stage(spectypesFormCallback.probe.stageOfInterest)
	}
	spectypes_ := spectypesFormCallback.spectypes
	_ = spectypes_

	for _, formDiv := range spectypesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spectypes_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if spectypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spectypes_.Unstage(spectypesFormCallback.probe.stageOfInterest)
	}

	spectypesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECTYPES](
		spectypesFormCallback.probe,
	)
	spectypesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spectypesFormCallback.CreationMode || spectypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spectypesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spectypesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECTYPESFormCallback(
			nil,
			spectypesFormCallback.probe,
			newFormGroup,
		)
		spectypes := new(models.SPECTYPES)
		FillUpForm(spectypes, newFormGroup, spectypesFormCallback.probe)
		spectypesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spectypesFormCallback.probe)
}
func __gong__New__TARGETFormCallback(
	target *models.TARGET,
	probe *Probe,
	formGroup *table.FormGroup,
) (targetFormCallback *TARGETFormCallback) {
	targetFormCallback = new(TARGETFormCallback)
	targetFormCallback.probe = probe
	targetFormCallback.target = target
	targetFormCallback.formGroup = formGroup

	targetFormCallback.CreationMode = (target == nil)

	return
}

type TARGETFormCallback struct {
	target *models.TARGET

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (targetFormCallback *TARGETFormCallback) OnSave() {

	log.Println("TARGETFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	targetFormCallback.probe.formStage.Checkout()

	if targetFormCallback.target == nil {
		targetFormCallback.target = new(models.TARGET).Stage(targetFormCallback.probe.stageOfInterest)
	}
	target_ := targetFormCallback.target
	_ = target_

	for _, formDiv := range targetFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(target_.Name), formDiv)
		case "SPECOBJECTREF":
			FormDivBasicFieldToField(&(target_.SPECOBJECTREF), formDiv)
		}
	}

	// manage the suppress operation
	if targetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		target_.Unstage(targetFormCallback.probe.stageOfInterest)
	}

	targetFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TARGET](
		targetFormCallback.probe,
	)
	targetFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if targetFormCallback.CreationMode || targetFormCallback.formGroup.HasSuppressButtonBeenPressed {
		targetFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(targetFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TARGETFormCallback(
			nil,
			targetFormCallback.probe,
			newFormGroup,
		)
		target := new(models.TARGET)
		FillUpForm(target, newFormGroup, targetFormCallback.probe)
		targetFormCallback.probe.formStage.Commit()
	}

	fillUpTree(targetFormCallback.probe)
}
func __gong__New__TARGETSPECIFICATIONFormCallback(
	targetspecification *models.TARGETSPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (targetspecificationFormCallback *TARGETSPECIFICATIONFormCallback) {
	targetspecificationFormCallback = new(TARGETSPECIFICATIONFormCallback)
	targetspecificationFormCallback.probe = probe
	targetspecificationFormCallback.targetspecification = targetspecification
	targetspecificationFormCallback.formGroup = formGroup

	targetspecificationFormCallback.CreationMode = (targetspecification == nil)

	return
}

type TARGETSPECIFICATIONFormCallback struct {
	targetspecification *models.TARGETSPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (targetspecificationFormCallback *TARGETSPECIFICATIONFormCallback) OnSave() {

	log.Println("TARGETSPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	targetspecificationFormCallback.probe.formStage.Checkout()

	if targetspecificationFormCallback.targetspecification == nil {
		targetspecificationFormCallback.targetspecification = new(models.TARGETSPECIFICATION).Stage(targetspecificationFormCallback.probe.stageOfInterest)
	}
	targetspecification_ := targetspecificationFormCallback.targetspecification
	_ = targetspecification_

	for _, formDiv := range targetspecificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(targetspecification_.Name), formDiv)
		case "SPECIFICATIONREF":
			FormDivBasicFieldToField(&(targetspecification_.SPECIFICATIONREF), formDiv)
		}
	}

	// manage the suppress operation
	if targetspecificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		targetspecification_.Unstage(targetspecificationFormCallback.probe.stageOfInterest)
	}

	targetspecificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TARGETSPECIFICATION](
		targetspecificationFormCallback.probe,
	)
	targetspecificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if targetspecificationFormCallback.CreationMode || targetspecificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		targetspecificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(targetspecificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TARGETSPECIFICATIONFormCallback(
			nil,
			targetspecificationFormCallback.probe,
			newFormGroup,
		)
		targetspecification := new(models.TARGETSPECIFICATION)
		FillUpForm(targetspecification, newFormGroup, targetspecificationFormCallback.probe)
		targetspecificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(targetspecificationFormCallback.probe)
}
func __gong__New__THEHEADERFormCallback(
	theheader *models.THEHEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (theheaderFormCallback *THEHEADERFormCallback) {
	theheaderFormCallback = new(THEHEADERFormCallback)
	theheaderFormCallback.probe = probe
	theheaderFormCallback.theheader = theheader
	theheaderFormCallback.formGroup = formGroup

	theheaderFormCallback.CreationMode = (theheader == nil)

	return
}

type THEHEADERFormCallback struct {
	theheader *models.THEHEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (theheaderFormCallback *THEHEADERFormCallback) OnSave() {

	log.Println("THEHEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	theheaderFormCallback.probe.formStage.Checkout()

	if theheaderFormCallback.theheader == nil {
		theheaderFormCallback.theheader = new(models.THEHEADER).Stage(theheaderFormCallback.probe.stageOfInterest)
	}
	theheader_ := theheaderFormCallback.theheader
	_ = theheader_

	for _, formDiv := range theheaderFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(theheader_.Name), formDiv)
		case "REQIFHEADER":
			FormDivSelectFieldToField(&(theheader_.REQIFHEADER), theheaderFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if theheaderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		theheader_.Unstage(theheaderFormCallback.probe.stageOfInterest)
	}

	theheaderFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.THEHEADER](
		theheaderFormCallback.probe,
	)
	theheaderFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if theheaderFormCallback.CreationMode || theheaderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		theheaderFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(theheaderFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__THEHEADERFormCallback(
			nil,
			theheaderFormCallback.probe,
			newFormGroup,
		)
		theheader := new(models.THEHEADER)
		FillUpForm(theheader, newFormGroup, theheaderFormCallback.probe)
		theheaderFormCallback.probe.formStage.Commit()
	}

	fillUpTree(theheaderFormCallback.probe)
}
func __gong__New__TOOLEXTENSIONSFormCallback(
	toolextensions *models.TOOLEXTENSIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (toolextensionsFormCallback *TOOLEXTENSIONSFormCallback) {
	toolextensionsFormCallback = new(TOOLEXTENSIONSFormCallback)
	toolextensionsFormCallback.probe = probe
	toolextensionsFormCallback.toolextensions = toolextensions
	toolextensionsFormCallback.formGroup = formGroup

	toolextensionsFormCallback.CreationMode = (toolextensions == nil)

	return
}

type TOOLEXTENSIONSFormCallback struct {
	toolextensions *models.TOOLEXTENSIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (toolextensionsFormCallback *TOOLEXTENSIONSFormCallback) OnSave() {

	log.Println("TOOLEXTENSIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	toolextensionsFormCallback.probe.formStage.Checkout()

	if toolextensionsFormCallback.toolextensions == nil {
		toolextensionsFormCallback.toolextensions = new(models.TOOLEXTENSIONS).Stage(toolextensionsFormCallback.probe.stageOfInterest)
	}
	toolextensions_ := toolextensionsFormCallback.toolextensions
	_ = toolextensions_

	for _, formDiv := range toolextensionsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(toolextensions_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if toolextensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		toolextensions_.Unstage(toolextensionsFormCallback.probe.stageOfInterest)
	}

	toolextensionsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TOOLEXTENSIONS](
		toolextensionsFormCallback.probe,
	)
	toolextensionsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if toolextensionsFormCallback.CreationMode || toolextensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		toolextensionsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(toolextensionsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TOOLEXTENSIONSFormCallback(
			nil,
			toolextensionsFormCallback.probe,
			newFormGroup,
		)
		toolextensions := new(models.TOOLEXTENSIONS)
		FillUpForm(toolextensions, newFormGroup, toolextensionsFormCallback.probe)
		toolextensionsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(toolextensionsFormCallback.probe)
}
func __gong__New__VALUESFormCallback(
	values *models.VALUES,
	probe *Probe,
	formGroup *table.FormGroup,
) (valuesFormCallback *VALUESFormCallback) {
	valuesFormCallback = new(VALUESFormCallback)
	valuesFormCallback.probe = probe
	valuesFormCallback.values = values
	valuesFormCallback.formGroup = formGroup

	valuesFormCallback.CreationMode = (values == nil)

	return
}

type VALUESFormCallback struct {
	values *models.VALUES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (valuesFormCallback *VALUESFormCallback) OnSave() {

	log.Println("VALUESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	valuesFormCallback.probe.formStage.Checkout()

	if valuesFormCallback.values == nil {
		valuesFormCallback.values = new(models.VALUES).Stage(valuesFormCallback.probe.stageOfInterest)
	}
	values_ := valuesFormCallback.values
	_ = values_

	for _, formDiv := range valuesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(values_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		values_.Unstage(valuesFormCallback.probe.stageOfInterest)
	}

	valuesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.VALUES](
		valuesFormCallback.probe,
	)
	valuesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if valuesFormCallback.CreationMode || valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		valuesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(valuesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VALUESFormCallback(
			nil,
			valuesFormCallback.probe,
			newFormGroup,
		)
		values := new(models.VALUES)
		FillUpForm(values, newFormGroup, valuesFormCallback.probe)
		valuesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(valuesFormCallback.probe)
}
func __gong__New__XHTMLCONTENTFormCallback(
	xhtmlcontent *models.XHTMLCONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtmlcontentFormCallback *XHTMLCONTENTFormCallback) {
	xhtmlcontentFormCallback = new(XHTMLCONTENTFormCallback)
	xhtmlcontentFormCallback.probe = probe
	xhtmlcontentFormCallback.xhtmlcontent = xhtmlcontent
	xhtmlcontentFormCallback.formGroup = formGroup

	xhtmlcontentFormCallback.CreationMode = (xhtmlcontent == nil)

	return
}

type XHTMLCONTENTFormCallback struct {
	xhtmlcontent *models.XHTMLCONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtmlcontentFormCallback *XHTMLCONTENTFormCallback) OnSave() {

	log.Println("XHTMLCONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtmlcontentFormCallback.probe.formStage.Checkout()

	if xhtmlcontentFormCallback.xhtmlcontent == nil {
		xhtmlcontentFormCallback.xhtmlcontent = new(models.XHTMLCONTENT).Stage(xhtmlcontentFormCallback.probe.stageOfInterest)
	}
	xhtmlcontent_ := xhtmlcontentFormCallback.xhtmlcontent
	_ = xhtmlcontent_

	for _, formDiv := range xhtmlcontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtmlcontent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if xhtmlcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtmlcontent_.Unstage(xhtmlcontentFormCallback.probe.stageOfInterest)
	}

	xhtmlcontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XHTMLCONTENT](
		xhtmlcontentFormCallback.probe,
	)
	xhtmlcontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtmlcontentFormCallback.CreationMode || xhtmlcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtmlcontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtmlcontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XHTMLCONTENTFormCallback(
			nil,
			xhtmlcontentFormCallback.probe,
			newFormGroup,
		)
		xhtmlcontent := new(models.XHTMLCONTENT)
		FillUpForm(xhtmlcontent, newFormGroup, xhtmlcontentFormCallback.probe)
		xhtmlcontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtmlcontentFormCallback.probe)
}
