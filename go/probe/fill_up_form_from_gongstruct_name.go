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
	case "ALTERNATIVEID":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ALTERNATIVEID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVEIDFormCallback(
			nil,
			probe,
			formGroup,
		)
		alternativeid := new(models.ALTERNATIVEID)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(alternativeid, formGroup, probe)
	case "ATTRIBUTEDEFINITIONBOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONBOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitionboolean := new(models.ATTRIBUTEDEFINITIONBOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitionboolean, formGroup, probe)
	case "ATTRIBUTEDEFINITIONDATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONDATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitiondate := new(models.ATTRIBUTEDEFINITIONDATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitiondate, formGroup, probe)
	case "ATTRIBUTEDEFINITIONENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitionenumeration := new(models.ATTRIBUTEDEFINITIONENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitionenumeration, formGroup, probe)
	case "ATTRIBUTEDEFINITIONINTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONINTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitioninteger := new(models.ATTRIBUTEDEFINITIONINTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitioninteger, formGroup, probe)
	case "ATTRIBUTEDEFINITIONREAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONREALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitionreal := new(models.ATTRIBUTEDEFINITIONREAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitionreal, formGroup, probe)
	case "ATTRIBUTEDEFINITIONSTRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONSTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONSTRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitionstring := new(models.ATTRIBUTEDEFINITIONSTRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitionstring, formGroup, probe)
	case "ATTRIBUTEDEFINITIONXHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEDEFINITIONXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEDEFINITIONXHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributedefinitionxhtml := new(models.ATTRIBUTEDEFINITIONXHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributedefinitionxhtml, formGroup, probe)
	case "ATTRIBUTEVALUEBOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEBOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevalueboolean := new(models.ATTRIBUTEVALUEBOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevalueboolean, formGroup, probe)
	case "ATTRIBUTEVALUEDATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEDATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevaluedate := new(models.ATTRIBUTEVALUEDATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevaluedate, formGroup, probe)
	case "ATTRIBUTEVALUEENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevalueenumeration := new(models.ATTRIBUTEVALUEENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevalueenumeration, formGroup, probe)
	case "ATTRIBUTEVALUEINTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEINTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevalueinteger := new(models.ATTRIBUTEVALUEINTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevalueinteger, formGroup, probe)
	case "ATTRIBUTEVALUEREAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEREALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevaluereal := new(models.ATTRIBUTEVALUEREAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevaluereal, formGroup, probe)
	case "ATTRIBUTEVALUESTRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUESTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUESTRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevaluestring := new(models.ATTRIBUTEVALUESTRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevaluestring, formGroup, probe)
	case "ATTRIBUTEVALUEXHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTEVALUEXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTEVALUEXHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributevaluexhtml := new(models.ATTRIBUTEVALUEXHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributevaluexhtml, formGroup, probe)
	case "CHILDREN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CHILDREN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CHILDRENFormCallback(
			nil,
			probe,
			formGroup,
		)
		children := new(models.CHILDREN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(children, formGroup, probe)
	case "CORECONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CORECONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CORECONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		corecontent := new(models.CORECONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(corecontent, formGroup, probe)
	case "DATATYPEDEFINITIONBOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONBOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONBOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitionboolean := new(models.DATATYPEDEFINITIONBOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitionboolean, formGroup, probe)
	case "DATATYPEDEFINITIONDATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONDATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONDATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitiondate := new(models.DATATYPEDEFINITIONDATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitiondate, formGroup, probe)
	case "DATATYPEDEFINITIONENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitionenumeration := new(models.DATATYPEDEFINITIONENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitionenumeration, formGroup, probe)
	case "DATATYPEDEFINITIONINTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONINTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONINTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitioninteger := new(models.DATATYPEDEFINITIONINTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitioninteger, formGroup, probe)
	case "DATATYPEDEFINITIONREAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONREAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONREALFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitionreal := new(models.DATATYPEDEFINITIONREAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitionreal, formGroup, probe)
	case "DATATYPEDEFINITIONSTRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONSTRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONSTRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitionstring := new(models.DATATYPEDEFINITIONSTRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitionstring, formGroup, probe)
	case "DATATYPEDEFINITIONXHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPEDEFINITIONXHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPEDEFINITIONXHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypedefinitionxhtml := new(models.DATATYPEDEFINITIONXHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypedefinitionxhtml, formGroup, probe)
	case "DATATYPES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPESFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatypes := new(models.DATATYPES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatypes, formGroup, probe)
	case "DEFAULTVALUE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DEFAULTVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DEFAULTVALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		defaultvalue := new(models.DEFAULTVALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(defaultvalue, formGroup, probe)
	case "DEFINITION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DEFINITION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DEFINITIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		definition := new(models.DEFINITION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(definition, formGroup, probe)
	case "EDITABLEATTS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "EDITABLEATTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EDITABLEATTSFormCallback(
			nil,
			probe,
			formGroup,
		)
		editableatts := new(models.EDITABLEATTS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(editableatts, formGroup, probe)
	case "EMBEDDEDVALUE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "EMBEDDEDVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDEDVALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		embeddedvalue := new(models.EMBEDDEDVALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(embeddedvalue, formGroup, probe)
	case "ENUMVALUE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ENUMVALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUMVALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		enumvalue := new(models.ENUMVALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(enumvalue, formGroup, probe)
	case "OBJECT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OBJECTFormCallback(
			nil,
			probe,
			formGroup,
		)
		object := new(models.OBJECT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(object, formGroup, probe)
	case "PROPERTIES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "PROPERTIES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PROPERTIESFormCallback(
			nil,
			probe,
			formGroup,
		)
		properties := new(models.PROPERTIES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(properties, formGroup, probe)
	case "RELATIONGROUP":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RELATIONGROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATIONGROUPFormCallback(
			nil,
			probe,
			formGroup,
		)
		relationgroup := new(models.RELATIONGROUP)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relationgroup, formGroup, probe)
	case "RELATIONGROUPTYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RELATIONGROUPTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATIONGROUPTYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		relationgrouptype := new(models.RELATIONGROUPTYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relationgrouptype, formGroup, probe)
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
	case "REQIFCONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQIFCONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFCONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		reqifcontent := new(models.REQIFCONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(reqifcontent, formGroup, probe)
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
	case "REQIFTOOLEXTENSION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQIFTOOLEXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFTOOLEXTENSIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		reqiftoolextension := new(models.REQIFTOOLEXTENSION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(reqiftoolextension, formGroup, probe)
	case "REQIFTYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQIFTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQIFTYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		reqiftype := new(models.REQIFTYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(reqiftype, formGroup, probe)
	case "SOURCE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SOURCE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SOURCEFormCallback(
			nil,
			probe,
			formGroup,
		)
		source := new(models.SOURCE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(source, formGroup, probe)
	case "SOURCESPECIFICATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SOURCESPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SOURCESPECIFICATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		sourcespecification := new(models.SOURCESPECIFICATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sourcespecification, formGroup, probe)
	case "SPECATTRIBUTES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECATTRIBUTES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECATTRIBUTESFormCallback(
			nil,
			probe,
			formGroup,
		)
		specattributes := new(models.SPECATTRIBUTES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specattributes, formGroup, probe)
	case "SPECHIERARCHY":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECHIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECHIERARCHYFormCallback(
			nil,
			probe,
			formGroup,
		)
		spechierarchy := new(models.SPECHIERARCHY)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spechierarchy, formGroup, probe)
	case "SPECIFICATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		specification := new(models.SPECIFICATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specification, formGroup, probe)
	case "SPECIFICATIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFICATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		specifications := new(models.SPECIFICATIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specifications, formGroup, probe)
	case "SPECIFICATIONTYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFICATIONTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONTYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		specificationtype := new(models.SPECIFICATIONTYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specificationtype, formGroup, probe)
	case "SPECIFIEDVALUES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFIEDVALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFIEDVALUESFormCallback(
			nil,
			probe,
			formGroup,
		)
		specifiedvalues := new(models.SPECIFIEDVALUES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specifiedvalues, formGroup, probe)
	case "SPECOBJECT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECOBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTFormCallback(
			nil,
			probe,
			formGroup,
		)
		specobject := new(models.SPECOBJECT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specobject, formGroup, probe)
	case "SPECOBJECTS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECOBJECTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTSFormCallback(
			nil,
			probe,
			formGroup,
		)
		specobjects := new(models.SPECOBJECTS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specobjects, formGroup, probe)
	case "SPECOBJECTTYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECOBJECTTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECOBJECTTYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		specobjecttype := new(models.SPECOBJECTTYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specobjecttype, formGroup, probe)
	case "SPECRELATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECRELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		specrelation := new(models.SPECRELATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specrelation, formGroup, probe)
	case "SPECRELATIONGROUPS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECRELATIONGROUPS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONGROUPSFormCallback(
			nil,
			probe,
			formGroup,
		)
		specrelationgroups := new(models.SPECRELATIONGROUPS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specrelationgroups, formGroup, probe)
	case "SPECRELATIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECRELATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		specrelations := new(models.SPECRELATIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specrelations, formGroup, probe)
	case "SPECRELATIONTYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECRELATIONTYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECRELATIONTYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		specrelationtype := new(models.SPECRELATIONTYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specrelationtype, formGroup, probe)
	case "SPECTYPES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECTYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECTYPESFormCallback(
			nil,
			probe,
			formGroup,
		)
		spectypes := new(models.SPECTYPES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spectypes, formGroup, probe)
	case "TARGET":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TARGET Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TARGETFormCallback(
			nil,
			probe,
			formGroup,
		)
		target := new(models.TARGET)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(target, formGroup, probe)
	case "TARGETSPECIFICATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TARGETSPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TARGETSPECIFICATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		targetspecification := new(models.TARGETSPECIFICATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(targetspecification, formGroup, probe)
	case "THEHEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "THEHEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__THEHEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		theheader := new(models.THEHEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(theheader, formGroup, probe)
	case "TOOLEXTENSIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TOOLEXTENSIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TOOLEXTENSIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		toolextensions := new(models.TOOLEXTENSIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(toolextensions, formGroup, probe)
	case "VALUES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "VALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VALUESFormCallback(
			nil,
			probe,
			formGroup,
		)
		values := new(models.VALUES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(values, formGroup, probe)
	case "XHTMLCONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "XHTMLCONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTMLCONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtmlcontent := new(models.XHTMLCONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtmlcontent, formGroup, probe)
	}
	formStage.Commit()
}
