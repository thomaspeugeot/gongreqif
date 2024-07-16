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
	case *models.ALTERNATIVEID:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ALTERNATIVEID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVEIDFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONBOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONDATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONDATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONINTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONREAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONREALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONSTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONSTRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEDEFINITIONXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONXHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEBOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEDATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEDATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEINTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEINTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEREAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEREALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUESTRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUESTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUESTRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTEVALUEXHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTEVALUEXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEXHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CHILDREN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CHILDREN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CHILDRENFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.CORECONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CORECONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CORECONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONBOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONDATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONDATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONINTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONINTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONREAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONREALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONSTRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONSTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONSTRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPEDEFINITIONXHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPEDEFINITIONXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONXHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DEFAULTVALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DEFAULTVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DEFAULTVALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DEFINITION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DEFINITION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DEFINITIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EDITABLEATTS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "EDITABLEATTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EDITABLEATTSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EMBEDDEDVALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "EMBEDDEDVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDEDVALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ENUMVALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ENUMVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUMVALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.OBJECT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OBJECTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PROPERTIES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "PROPERTIES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PROPERTIESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATIONGROUP:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATIONGROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATIONGROUPFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATIONGROUPTYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATIONGROUPTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATIONGROUPTYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
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
	case *models.REQIFCONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQIFCONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFCONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQIFHEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQIFHEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFHEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQIFTOOLEXTENSION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQIFTOOLEXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFTOOLEXTENSIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQTYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQTYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SOURCE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SOURCE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SOURCEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SOURCESPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SOURCESPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SOURCESPECIFICATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECATTRIBUTES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECATTRIBUTES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECATTRIBUTESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECHIERARCHY:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECHIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECHIERARCHYFormCallback(
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
	case *models.SPECIFICATIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATIONTYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATIONTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONTYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFIEDVALUES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFIEDVALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFIEDVALUESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECOBJECT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECOBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECOBJECTS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECOBJECTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECOBJECTTYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECOBJECTTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTTYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECRELATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECRELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECRELATIONGROUPS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECRELATIONGROUPS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONGROUPSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECRELATIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECRELATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECRELATIONTYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECRELATIONTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONTYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECTYPES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECTYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECTYPESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TARGET:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TARGET Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TARGETFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TARGETSPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TARGETSPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TARGETSPECIFICATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.THEHEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "THEHEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__THEHEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TOOLEXTENSIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TOOLEXTENSIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TOOLEXTENSIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.VALUES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "VALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VALUESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XHTMLCONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XHTMLCONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTMLCONTENTFormCallback(
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
