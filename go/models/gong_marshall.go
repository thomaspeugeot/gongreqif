// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_ALTERNATIVEID_Identifiers := make(map[*ALTERNATIVEID]string)
	_ = map_ALTERNATIVEID_Identifiers

	alternativeidOrdered := []*ALTERNATIVEID{}
	for alternativeid := range stage.ALTERNATIVEIDs {
		alternativeidOrdered = append(alternativeidOrdered, alternativeid)
	}
	sort.Slice(alternativeidOrdered[:], func(i, j int) bool {
		return alternativeidOrdered[i].Name < alternativeidOrdered[j].Name
	})
	if len(alternativeidOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, alternativeid := range alternativeidOrdered {

		id = generatesIdentifier("ALTERNATIVEID", idx, alternativeid.Name)
		map_ALTERNATIVEID_Identifiers[alternativeid] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ALTERNATIVEID")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", alternativeid.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternativeid.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternativeid.IDENTIFIERAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONBOOLEAN_Identifiers := make(map[*ATTRIBUTEDEFINITIONBOOLEAN]string)
	_ = map_ATTRIBUTEDEFINITIONBOOLEAN_Identifiers

	attributedefinitionbooleanOrdered := []*ATTRIBUTEDEFINITIONBOOLEAN{}
	for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
		attributedefinitionbooleanOrdered = append(attributedefinitionbooleanOrdered, attributedefinitionboolean)
	}
	sort.Slice(attributedefinitionbooleanOrdered[:], func(i, j int) bool {
		return attributedefinitionbooleanOrdered[i].Name < attributedefinitionbooleanOrdered[j].Name
	})
	if len(attributedefinitionbooleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitionboolean := range attributedefinitionbooleanOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONBOOLEAN", idx, attributedefinitionboolean.Name)
		map_ATTRIBUTEDEFINITIONBOOLEAN_Identifiers[attributedefinitionboolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONBOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitionboolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionboolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionboolean.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionboolean.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionboolean.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionboolean.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionboolean.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONDATE_Identifiers := make(map[*ATTRIBUTEDEFINITIONDATE]string)
	_ = map_ATTRIBUTEDEFINITIONDATE_Identifiers

	attributedefinitiondateOrdered := []*ATTRIBUTEDEFINITIONDATE{}
	for attributedefinitiondate := range stage.ATTRIBUTEDEFINITIONDATEs {
		attributedefinitiondateOrdered = append(attributedefinitiondateOrdered, attributedefinitiondate)
	}
	sort.Slice(attributedefinitiondateOrdered[:], func(i, j int) bool {
		return attributedefinitiondateOrdered[i].Name < attributedefinitiondateOrdered[j].Name
	})
	if len(attributedefinitiondateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitiondate := range attributedefinitiondateOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONDATE", idx, attributedefinitiondate.Name)
		map_ATTRIBUTEDEFINITIONDATE_Identifiers[attributedefinitiondate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONDATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitiondate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitiondate.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitiondate.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitiondate.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitiondate.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitiondate.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitiondate.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONENUMERATION_Identifiers := make(map[*ATTRIBUTEDEFINITIONENUMERATION]string)
	_ = map_ATTRIBUTEDEFINITIONENUMERATION_Identifiers

	attributedefinitionenumerationOrdered := []*ATTRIBUTEDEFINITIONENUMERATION{}
	for attributedefinitionenumeration := range stage.ATTRIBUTEDEFINITIONENUMERATIONs {
		attributedefinitionenumerationOrdered = append(attributedefinitionenumerationOrdered, attributedefinitionenumeration)
	}
	sort.Slice(attributedefinitionenumerationOrdered[:], func(i, j int) bool {
		return attributedefinitionenumerationOrdered[i].Name < attributedefinitionenumerationOrdered[j].Name
	})
	if len(attributedefinitionenumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitionenumeration := range attributedefinitionenumerationOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONENUMERATION", idx, attributedefinitionenumeration.Name)
		map_ATTRIBUTEDEFINITIONENUMERATION_Identifiers[attributedefinitionenumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitionenumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionenumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionenumeration.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionenumeration.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionenumeration.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionenumeration.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionenumeration.LONGNAMEAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MULTIVALUEDAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionenumeration.MULTIVALUEDAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONINTEGER_Identifiers := make(map[*ATTRIBUTEDEFINITIONINTEGER]string)
	_ = map_ATTRIBUTEDEFINITIONINTEGER_Identifiers

	attributedefinitionintegerOrdered := []*ATTRIBUTEDEFINITIONINTEGER{}
	for attributedefinitioninteger := range stage.ATTRIBUTEDEFINITIONINTEGERs {
		attributedefinitionintegerOrdered = append(attributedefinitionintegerOrdered, attributedefinitioninteger)
	}
	sort.Slice(attributedefinitionintegerOrdered[:], func(i, j int) bool {
		return attributedefinitionintegerOrdered[i].Name < attributedefinitionintegerOrdered[j].Name
	})
	if len(attributedefinitionintegerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitioninteger := range attributedefinitionintegerOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONINTEGER", idx, attributedefinitioninteger.Name)
		map_ATTRIBUTEDEFINITIONINTEGER_Identifiers[attributedefinitioninteger] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONINTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitioninteger.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitioninteger.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitioninteger.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitioninteger.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitioninteger.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitioninteger.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitioninteger.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONREAL_Identifiers := make(map[*ATTRIBUTEDEFINITIONREAL]string)
	_ = map_ATTRIBUTEDEFINITIONREAL_Identifiers

	attributedefinitionrealOrdered := []*ATTRIBUTEDEFINITIONREAL{}
	for attributedefinitionreal := range stage.ATTRIBUTEDEFINITIONREALs {
		attributedefinitionrealOrdered = append(attributedefinitionrealOrdered, attributedefinitionreal)
	}
	sort.Slice(attributedefinitionrealOrdered[:], func(i, j int) bool {
		return attributedefinitionrealOrdered[i].Name < attributedefinitionrealOrdered[j].Name
	})
	if len(attributedefinitionrealOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitionreal := range attributedefinitionrealOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONREAL", idx, attributedefinitionreal.Name)
		map_ATTRIBUTEDEFINITIONREAL_Identifiers[attributedefinitionreal] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONREAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitionreal.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionreal.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionreal.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionreal.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionreal.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionreal.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionreal.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONSTRING_Identifiers := make(map[*ATTRIBUTEDEFINITIONSTRING]string)
	_ = map_ATTRIBUTEDEFINITIONSTRING_Identifiers

	attributedefinitionstringOrdered := []*ATTRIBUTEDEFINITIONSTRING{}
	for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
		attributedefinitionstringOrdered = append(attributedefinitionstringOrdered, attributedefinitionstring)
	}
	sort.Slice(attributedefinitionstringOrdered[:], func(i, j int) bool {
		return attributedefinitionstringOrdered[i].Name < attributedefinitionstringOrdered[j].Name
	})
	if len(attributedefinitionstringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitionstring := range attributedefinitionstringOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONSTRING", idx, attributedefinitionstring.Name)
		map_ATTRIBUTEDEFINITIONSTRING_Identifiers[attributedefinitionstring] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONSTRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitionstring.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionstring.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionstring.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionstring.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionstring.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionstring.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionstring.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEDEFINITIONXHTML_Identifiers := make(map[*ATTRIBUTEDEFINITIONXHTML]string)
	_ = map_ATTRIBUTEDEFINITIONXHTML_Identifiers

	attributedefinitionxhtmlOrdered := []*ATTRIBUTEDEFINITIONXHTML{}
	for attributedefinitionxhtml := range stage.ATTRIBUTEDEFINITIONXHTMLs {
		attributedefinitionxhtmlOrdered = append(attributedefinitionxhtmlOrdered, attributedefinitionxhtml)
	}
	sort.Slice(attributedefinitionxhtmlOrdered[:], func(i, j int) bool {
		return attributedefinitionxhtmlOrdered[i].Name < attributedefinitionxhtmlOrdered[j].Name
	})
	if len(attributedefinitionxhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributedefinitionxhtml := range attributedefinitionxhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTEDEFINITIONXHTML", idx, attributedefinitionxhtml.Name)
		map_ATTRIBUTEDEFINITIONXHTML_Identifiers[attributedefinitionxhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEDEFINITIONXHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributedefinitionxhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionxhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionxhtml.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionxhtml.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributedefinitionxhtml.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionxhtml.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributedefinitionxhtml.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEBOOLEAN_Identifiers := make(map[*ATTRIBUTEVALUEBOOLEAN]string)
	_ = map_ATTRIBUTEVALUEBOOLEAN_Identifiers

	attributevaluebooleanOrdered := []*ATTRIBUTEVALUEBOOLEAN{}
	for attributevalueboolean := range stage.ATTRIBUTEVALUEBOOLEANs {
		attributevaluebooleanOrdered = append(attributevaluebooleanOrdered, attributevalueboolean)
	}
	sort.Slice(attributevaluebooleanOrdered[:], func(i, j int) bool {
		return attributevaluebooleanOrdered[i].Name < attributevaluebooleanOrdered[j].Name
	})
	if len(attributevaluebooleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevalueboolean := range attributevaluebooleanOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEBOOLEAN", idx, attributevalueboolean.Name)
		map_ATTRIBUTEVALUEBOOLEAN_Identifiers[attributevalueboolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEBOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevalueboolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevalueboolean.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THEVALUEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributevalueboolean.THEVALUEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEDATE_Identifiers := make(map[*ATTRIBUTEVALUEDATE]string)
	_ = map_ATTRIBUTEVALUEDATE_Identifiers

	attributevaluedateOrdered := []*ATTRIBUTEVALUEDATE{}
	for attributevaluedate := range stage.ATTRIBUTEVALUEDATEs {
		attributevaluedateOrdered = append(attributevaluedateOrdered, attributevaluedate)
	}
	sort.Slice(attributevaluedateOrdered[:], func(i, j int) bool {
		return attributevaluedateOrdered[i].Name < attributevaluedateOrdered[j].Name
	})
	if len(attributevaluedateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevaluedate := range attributevaluedateOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEDATE", idx, attributevaluedate.Name)
		map_ATTRIBUTEVALUEDATE_Identifiers[attributevaluedate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEDATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevaluedate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluedate.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THEVALUEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluedate.THEVALUEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEENUMERATION_Identifiers := make(map[*ATTRIBUTEVALUEENUMERATION]string)
	_ = map_ATTRIBUTEVALUEENUMERATION_Identifiers

	attributevalueenumerationOrdered := []*ATTRIBUTEVALUEENUMERATION{}
	for attributevalueenumeration := range stage.ATTRIBUTEVALUEENUMERATIONs {
		attributevalueenumerationOrdered = append(attributevalueenumerationOrdered, attributevalueenumeration)
	}
	sort.Slice(attributevalueenumerationOrdered[:], func(i, j int) bool {
		return attributevalueenumerationOrdered[i].Name < attributevalueenumerationOrdered[j].Name
	})
	if len(attributevalueenumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevalueenumeration := range attributevalueenumerationOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEENUMERATION", idx, attributevalueenumeration.Name)
		map_ATTRIBUTEVALUEENUMERATION_Identifiers[attributevalueenumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevalueenumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevalueenumeration.Name))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEINTEGER_Identifiers := make(map[*ATTRIBUTEVALUEINTEGER]string)
	_ = map_ATTRIBUTEVALUEINTEGER_Identifiers

	attributevalueintegerOrdered := []*ATTRIBUTEVALUEINTEGER{}
	for attributevalueinteger := range stage.ATTRIBUTEVALUEINTEGERs {
		attributevalueintegerOrdered = append(attributevalueintegerOrdered, attributevalueinteger)
	}
	sort.Slice(attributevalueintegerOrdered[:], func(i, j int) bool {
		return attributevalueintegerOrdered[i].Name < attributevalueintegerOrdered[j].Name
	})
	if len(attributevalueintegerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevalueinteger := range attributevalueintegerOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEINTEGER", idx, attributevalueinteger.Name)
		map_ATTRIBUTEVALUEINTEGER_Identifiers[attributevalueinteger] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEINTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevalueinteger.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevalueinteger.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THEVALUEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attributevalueinteger.THEVALUEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEREAL_Identifiers := make(map[*ATTRIBUTEVALUEREAL]string)
	_ = map_ATTRIBUTEVALUEREAL_Identifiers

	attributevaluerealOrdered := []*ATTRIBUTEVALUEREAL{}
	for attributevaluereal := range stage.ATTRIBUTEVALUEREALs {
		attributevaluerealOrdered = append(attributevaluerealOrdered, attributevaluereal)
	}
	sort.Slice(attributevaluerealOrdered[:], func(i, j int) bool {
		return attributevaluerealOrdered[i].Name < attributevaluerealOrdered[j].Name
	})
	if len(attributevaluerealOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevaluereal := range attributevaluerealOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEREAL", idx, attributevaluereal.Name)
		map_ATTRIBUTEVALUEREAL_Identifiers[attributevaluereal] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEREAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevaluereal.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluereal.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THEVALUEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", attributevaluereal.THEVALUEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUESTRING_Identifiers := make(map[*ATTRIBUTEVALUESTRING]string)
	_ = map_ATTRIBUTEVALUESTRING_Identifiers

	attributevaluestringOrdered := []*ATTRIBUTEVALUESTRING{}
	for attributevaluestring := range stage.ATTRIBUTEVALUESTRINGs {
		attributevaluestringOrdered = append(attributevaluestringOrdered, attributevaluestring)
	}
	sort.Slice(attributevaluestringOrdered[:], func(i, j int) bool {
		return attributevaluestringOrdered[i].Name < attributevaluestringOrdered[j].Name
	})
	if len(attributevaluestringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevaluestring := range attributevaluestringOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUESTRING", idx, attributevaluestring.Name)
		map_ATTRIBUTEVALUESTRING_Identifiers[attributevaluestring] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUESTRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevaluestring.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluestring.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THEVALUEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluestring.THEVALUEAttr))
		initializerStatements += setValueField

	}

	map_ATTRIBUTEVALUEXHTML_Identifiers := make(map[*ATTRIBUTEVALUEXHTML]string)
	_ = map_ATTRIBUTEVALUEXHTML_Identifiers

	attributevaluexhtmlOrdered := []*ATTRIBUTEVALUEXHTML{}
	for attributevaluexhtml := range stage.ATTRIBUTEVALUEXHTMLs {
		attributevaluexhtmlOrdered = append(attributevaluexhtmlOrdered, attributevaluexhtml)
	}
	sort.Slice(attributevaluexhtmlOrdered[:], func(i, j int) bool {
		return attributevaluexhtmlOrdered[i].Name < attributevaluexhtmlOrdered[j].Name
	})
	if len(attributevaluexhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributevaluexhtml := range attributevaluexhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTEVALUEXHTML", idx, attributevaluexhtml.Name)
		map_ATTRIBUTEVALUEXHTML_Identifiers[attributevaluexhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTEVALUEXHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributevaluexhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributevaluexhtml.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISSIMPLIFIEDAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributevaluexhtml.ISSIMPLIFIEDAttr))
		initializerStatements += setValueField

	}

	map_CHILDREN_Identifiers := make(map[*CHILDREN]string)
	_ = map_CHILDREN_Identifiers

	childrenOrdered := []*CHILDREN{}
	for children := range stage.CHILDRENs {
		childrenOrdered = append(childrenOrdered, children)
	}
	sort.Slice(childrenOrdered[:], func(i, j int) bool {
		return childrenOrdered[i].Name < childrenOrdered[j].Name
	})
	if len(childrenOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, children := range childrenOrdered {

		id = generatesIdentifier("CHILDREN", idx, children.Name)
		map_CHILDREN_Identifiers[children] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CHILDREN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", children.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(children.Name))
		initializerStatements += setValueField

	}

	map_CORECONTENT_Identifiers := make(map[*CORECONTENT]string)
	_ = map_CORECONTENT_Identifiers

	corecontentOrdered := []*CORECONTENT{}
	for corecontent := range stage.CORECONTENTs {
		corecontentOrdered = append(corecontentOrdered, corecontent)
	}
	sort.Slice(corecontentOrdered[:], func(i, j int) bool {
		return corecontentOrdered[i].Name < corecontentOrdered[j].Name
	})
	if len(corecontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, corecontent := range corecontentOrdered {

		id = generatesIdentifier("CORECONTENT", idx, corecontent.Name)
		map_CORECONTENT_Identifiers[corecontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CORECONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", corecontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(corecontent.Name))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONBOOLEAN_Identifiers := make(map[*DATATYPEDEFINITIONBOOLEAN]string)
	_ = map_DATATYPEDEFINITIONBOOLEAN_Identifiers

	datatypedefinitionbooleanOrdered := []*DATATYPEDEFINITIONBOOLEAN{}
	for datatypedefinitionboolean := range stage.DATATYPEDEFINITIONBOOLEANs {
		datatypedefinitionbooleanOrdered = append(datatypedefinitionbooleanOrdered, datatypedefinitionboolean)
	}
	sort.Slice(datatypedefinitionbooleanOrdered[:], func(i, j int) bool {
		return datatypedefinitionbooleanOrdered[i].Name < datatypedefinitionbooleanOrdered[j].Name
	})
	if len(datatypedefinitionbooleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitionboolean := range datatypedefinitionbooleanOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONBOOLEAN", idx, datatypedefinitionboolean.Name)
		map_DATATYPEDEFINITIONBOOLEAN_Identifiers[datatypedefinitionboolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONBOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitionboolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionboolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionboolean.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionboolean.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionboolean.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionboolean.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONDATE_Identifiers := make(map[*DATATYPEDEFINITIONDATE]string)
	_ = map_DATATYPEDEFINITIONDATE_Identifiers

	datatypedefinitiondateOrdered := []*DATATYPEDEFINITIONDATE{}
	for datatypedefinitiondate := range stage.DATATYPEDEFINITIONDATEs {
		datatypedefinitiondateOrdered = append(datatypedefinitiondateOrdered, datatypedefinitiondate)
	}
	sort.Slice(datatypedefinitiondateOrdered[:], func(i, j int) bool {
		return datatypedefinitiondateOrdered[i].Name < datatypedefinitiondateOrdered[j].Name
	})
	if len(datatypedefinitiondateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitiondate := range datatypedefinitiondateOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONDATE", idx, datatypedefinitiondate.Name)
		map_DATATYPEDEFINITIONDATE_Identifiers[datatypedefinitiondate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONDATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitiondate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitiondate.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitiondate.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitiondate.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitiondate.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitiondate.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONENUMERATION_Identifiers := make(map[*DATATYPEDEFINITIONENUMERATION]string)
	_ = map_DATATYPEDEFINITIONENUMERATION_Identifiers

	datatypedefinitionenumerationOrdered := []*DATATYPEDEFINITIONENUMERATION{}
	for datatypedefinitionenumeration := range stage.DATATYPEDEFINITIONENUMERATIONs {
		datatypedefinitionenumerationOrdered = append(datatypedefinitionenumerationOrdered, datatypedefinitionenumeration)
	}
	sort.Slice(datatypedefinitionenumerationOrdered[:], func(i, j int) bool {
		return datatypedefinitionenumerationOrdered[i].Name < datatypedefinitionenumerationOrdered[j].Name
	})
	if len(datatypedefinitionenumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitionenumeration := range datatypedefinitionenumerationOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONENUMERATION", idx, datatypedefinitionenumeration.Name)
		map_DATATYPEDEFINITIONENUMERATION_Identifiers[datatypedefinitionenumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitionenumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionenumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionenumeration.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionenumeration.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionenumeration.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionenumeration.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONINTEGER_Identifiers := make(map[*DATATYPEDEFINITIONINTEGER]string)
	_ = map_DATATYPEDEFINITIONINTEGER_Identifiers

	datatypedefinitionintegerOrdered := []*DATATYPEDEFINITIONINTEGER{}
	for datatypedefinitioninteger := range stage.DATATYPEDEFINITIONINTEGERs {
		datatypedefinitionintegerOrdered = append(datatypedefinitionintegerOrdered, datatypedefinitioninteger)
	}
	sort.Slice(datatypedefinitionintegerOrdered[:], func(i, j int) bool {
		return datatypedefinitionintegerOrdered[i].Name < datatypedefinitionintegerOrdered[j].Name
	})
	if len(datatypedefinitionintegerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitioninteger := range datatypedefinitionintegerOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONINTEGER", idx, datatypedefinitioninteger.Name)
		map_DATATYPEDEFINITIONINTEGER_Identifiers[datatypedefinitioninteger] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONINTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitioninteger.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitioninteger.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitioninteger.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitioninteger.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitioninteger.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitioninteger.LONGNAMEAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAXAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatypedefinitioninteger.MAXAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MINAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatypedefinitioninteger.MINAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONREAL_Identifiers := make(map[*DATATYPEDEFINITIONREAL]string)
	_ = map_DATATYPEDEFINITIONREAL_Identifiers

	datatypedefinitionrealOrdered := []*DATATYPEDEFINITIONREAL{}
	for datatypedefinitionreal := range stage.DATATYPEDEFINITIONREALs {
		datatypedefinitionrealOrdered = append(datatypedefinitionrealOrdered, datatypedefinitionreal)
	}
	sort.Slice(datatypedefinitionrealOrdered[:], func(i, j int) bool {
		return datatypedefinitionrealOrdered[i].Name < datatypedefinitionrealOrdered[j].Name
	})
	if len(datatypedefinitionrealOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitionreal := range datatypedefinitionrealOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONREAL", idx, datatypedefinitionreal.Name)
		map_DATATYPEDEFINITIONREAL_Identifiers[datatypedefinitionreal] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONREAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitionreal.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionreal.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ACCURACYAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatypedefinitionreal.ACCURACYAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionreal.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionreal.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionreal.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionreal.LONGNAMEAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAXAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatypedefinitionreal.MAXAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MINAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatypedefinitionreal.MINAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONSTRING_Identifiers := make(map[*DATATYPEDEFINITIONSTRING]string)
	_ = map_DATATYPEDEFINITIONSTRING_Identifiers

	datatypedefinitionstringOrdered := []*DATATYPEDEFINITIONSTRING{}
	for datatypedefinitionstring := range stage.DATATYPEDEFINITIONSTRINGs {
		datatypedefinitionstringOrdered = append(datatypedefinitionstringOrdered, datatypedefinitionstring)
	}
	sort.Slice(datatypedefinitionstringOrdered[:], func(i, j int) bool {
		return datatypedefinitionstringOrdered[i].Name < datatypedefinitionstringOrdered[j].Name
	})
	if len(datatypedefinitionstringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitionstring := range datatypedefinitionstringOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONSTRING", idx, datatypedefinitionstring.Name)
		map_DATATYPEDEFINITIONSTRING_Identifiers[datatypedefinitionstring] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONSTRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitionstring.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionstring.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionstring.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionstring.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionstring.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionstring.LONGNAMEAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAXLENGTHAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatypedefinitionstring.MAXLENGTHAttr))
		initializerStatements += setValueField

	}

	map_DATATYPEDEFINITIONXHTML_Identifiers := make(map[*DATATYPEDEFINITIONXHTML]string)
	_ = map_DATATYPEDEFINITIONXHTML_Identifiers

	datatypedefinitionxhtmlOrdered := []*DATATYPEDEFINITIONXHTML{}
	for datatypedefinitionxhtml := range stage.DATATYPEDEFINITIONXHTMLs {
		datatypedefinitionxhtmlOrdered = append(datatypedefinitionxhtmlOrdered, datatypedefinitionxhtml)
	}
	sort.Slice(datatypedefinitionxhtmlOrdered[:], func(i, j int) bool {
		return datatypedefinitionxhtmlOrdered[i].Name < datatypedefinitionxhtmlOrdered[j].Name
	})
	if len(datatypedefinitionxhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypedefinitionxhtml := range datatypedefinitionxhtmlOrdered {

		id = generatesIdentifier("DATATYPEDEFINITIONXHTML", idx, datatypedefinitionxhtml.Name)
		map_DATATYPEDEFINITIONXHTML_Identifiers[datatypedefinitionxhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPEDEFINITIONXHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypedefinitionxhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionxhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionxhtml.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionxhtml.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionxhtml.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypedefinitionxhtml.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_DATATYPES_Identifiers := make(map[*DATATYPES]string)
	_ = map_DATATYPES_Identifiers

	datatypesOrdered := []*DATATYPES{}
	for datatypes := range stage.DATATYPESs {
		datatypesOrdered = append(datatypesOrdered, datatypes)
	}
	sort.Slice(datatypesOrdered[:], func(i, j int) bool {
		return datatypesOrdered[i].Name < datatypesOrdered[j].Name
	})
	if len(datatypesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatypes := range datatypesOrdered {

		id = generatesIdentifier("DATATYPES", idx, datatypes.Name)
		map_DATATYPES_Identifiers[datatypes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatypes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatypes.Name))
		initializerStatements += setValueField

	}

	map_DEFAULTVALUE_Identifiers := make(map[*DEFAULTVALUE]string)
	_ = map_DEFAULTVALUE_Identifiers

	defaultvalueOrdered := []*DEFAULTVALUE{}
	for defaultvalue := range stage.DEFAULTVALUEs {
		defaultvalueOrdered = append(defaultvalueOrdered, defaultvalue)
	}
	sort.Slice(defaultvalueOrdered[:], func(i, j int) bool {
		return defaultvalueOrdered[i].Name < defaultvalueOrdered[j].Name
	})
	if len(defaultvalueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, defaultvalue := range defaultvalueOrdered {

		id = generatesIdentifier("DEFAULTVALUE", idx, defaultvalue.Name)
		map_DEFAULTVALUE_Identifiers[defaultvalue] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DEFAULTVALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", defaultvalue.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(defaultvalue.Name))
		initializerStatements += setValueField

	}

	map_DEFINITION_Identifiers := make(map[*DEFINITION]string)
	_ = map_DEFINITION_Identifiers

	definitionOrdered := []*DEFINITION{}
	for definition := range stage.DEFINITIONs {
		definitionOrdered = append(definitionOrdered, definition)
	}
	sort.Slice(definitionOrdered[:], func(i, j int) bool {
		return definitionOrdered[i].Name < definitionOrdered[j].Name
	})
	if len(definitionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, definition := range definitionOrdered {

		id = generatesIdentifier("DEFINITION", idx, definition.Name)
		map_DEFINITION_Identifiers[definition] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DEFINITION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", definition.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(definition.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONBOOLEANREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(definition.ATTRIBUTEDEFINITIONBOOLEANREF))
		initializerStatements += setValueField

	}

	map_EDITABLEATTS_Identifiers := make(map[*EDITABLEATTS]string)
	_ = map_EDITABLEATTS_Identifiers

	editableattsOrdered := []*EDITABLEATTS{}
	for editableatts := range stage.EDITABLEATTSs {
		editableattsOrdered = append(editableattsOrdered, editableatts)
	}
	sort.Slice(editableattsOrdered[:], func(i, j int) bool {
		return editableattsOrdered[i].Name < editableattsOrdered[j].Name
	})
	if len(editableattsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, editableatts := range editableattsOrdered {

		id = generatesIdentifier("EDITABLEATTS", idx, editableatts.Name)
		map_EDITABLEATTS_Identifiers[editableatts] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EDITABLEATTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", editableatts.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(editableatts.Name))
		initializerStatements += setValueField

	}

	map_EMBEDDEDVALUE_Identifiers := make(map[*EMBEDDEDVALUE]string)
	_ = map_EMBEDDEDVALUE_Identifiers

	embeddedvalueOrdered := []*EMBEDDEDVALUE{}
	for embeddedvalue := range stage.EMBEDDEDVALUEs {
		embeddedvalueOrdered = append(embeddedvalueOrdered, embeddedvalue)
	}
	sort.Slice(embeddedvalueOrdered[:], func(i, j int) bool {
		return embeddedvalueOrdered[i].Name < embeddedvalueOrdered[j].Name
	})
	if len(embeddedvalueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, embeddedvalue := range embeddedvalueOrdered {

		id = generatesIdentifier("EMBEDDEDVALUE", idx, embeddedvalue.Name)
		map_EMBEDDEDVALUE_Identifiers[embeddedvalue] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EMBEDDEDVALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embeddedvalue.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedvalue.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "KEYAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", embeddedvalue.KEYAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OTHERCONTENTAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embeddedvalue.OTHERCONTENTAttr))
		initializerStatements += setValueField

	}

	map_ENUMVALUE_Identifiers := make(map[*ENUMVALUE]string)
	_ = map_ENUMVALUE_Identifiers

	enumvalueOrdered := []*ENUMVALUE{}
	for enumvalue := range stage.ENUMVALUEs {
		enumvalueOrdered = append(enumvalueOrdered, enumvalue)
	}
	sort.Slice(enumvalueOrdered[:], func(i, j int) bool {
		return enumvalueOrdered[i].Name < enumvalueOrdered[j].Name
	})
	if len(enumvalueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, enumvalue := range enumvalueOrdered {

		id = generatesIdentifier("ENUMVALUE", idx, enumvalue.Name)
		map_ENUMVALUE_Identifiers[enumvalue] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ENUMVALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enumvalue.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumvalue.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumvalue.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumvalue.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumvalue.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumvalue.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_OBJECT_Identifiers := make(map[*OBJECT]string)
	_ = map_OBJECT_Identifiers

	objectOrdered := []*OBJECT{}
	for object := range stage.OBJECTs {
		objectOrdered = append(objectOrdered, object)
	}
	sort.Slice(objectOrdered[:], func(i, j int) bool {
		return objectOrdered[i].Name < objectOrdered[j].Name
	})
	if len(objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, object := range objectOrdered {

		id = generatesIdentifier("OBJECT", idx, object.Name)
		map_OBJECT_Identifiers[object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "OBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(object.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECOBJECTREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(object.SPECOBJECTREF))
		initializerStatements += setValueField

	}

	map_PROPERTIES_Identifiers := make(map[*PROPERTIES]string)
	_ = map_PROPERTIES_Identifiers

	propertiesOrdered := []*PROPERTIES{}
	for properties := range stage.PROPERTIESs {
		propertiesOrdered = append(propertiesOrdered, properties)
	}
	sort.Slice(propertiesOrdered[:], func(i, j int) bool {
		return propertiesOrdered[i].Name < propertiesOrdered[j].Name
	})
	if len(propertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, properties := range propertiesOrdered {

		id = generatesIdentifier("PROPERTIES", idx, properties.Name)
		map_PROPERTIES_Identifiers[properties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PROPERTIES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", properties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(properties.Name))
		initializerStatements += setValueField

	}

	map_RELATIONGROUP_Identifiers := make(map[*RELATIONGROUP]string)
	_ = map_RELATIONGROUP_Identifiers

	relationgroupOrdered := []*RELATIONGROUP{}
	for relationgroup := range stage.RELATIONGROUPs {
		relationgroupOrdered = append(relationgroupOrdered, relationgroup)
	}
	sort.Slice(relationgroupOrdered[:], func(i, j int) bool {
		return relationgroupOrdered[i].Name < relationgroupOrdered[j].Name
	})
	if len(relationgroupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, relationgroup := range relationgroupOrdered {

		id = generatesIdentifier("RELATIONGROUP", idx, relationgroup.Name)
		map_RELATIONGROUP_Identifiers[relationgroup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATIONGROUP")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relationgroup.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgroup.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgroup.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgroup.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgroup.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgroup.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_RELATIONGROUPTYPE_Identifiers := make(map[*RELATIONGROUPTYPE]string)
	_ = map_RELATIONGROUPTYPE_Identifiers

	relationgrouptypeOrdered := []*RELATIONGROUPTYPE{}
	for relationgrouptype := range stage.RELATIONGROUPTYPEs {
		relationgrouptypeOrdered = append(relationgrouptypeOrdered, relationgrouptype)
	}
	sort.Slice(relationgrouptypeOrdered[:], func(i, j int) bool {
		return relationgrouptypeOrdered[i].Name < relationgrouptypeOrdered[j].Name
	})
	if len(relationgrouptypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, relationgrouptype := range relationgrouptypeOrdered {

		id = generatesIdentifier("RELATIONGROUPTYPE", idx, relationgrouptype.Name)
		map_RELATIONGROUPTYPE_Identifiers[relationgrouptype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATIONGROUPTYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relationgrouptype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgrouptype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgrouptype.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgrouptype.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgrouptype.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relationgrouptype.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_REQIF_Identifiers := make(map[*REQIF]string)
	_ = map_REQIF_Identifiers

	reqifOrdered := []*REQIF{}
	for reqif := range stage.REQIFs {
		reqifOrdered = append(reqifOrdered, reqif)
	}
	sort.Slice(reqifOrdered[:], func(i, j int) bool {
		return reqifOrdered[i].Name < reqifOrdered[j].Name
	})
	if len(reqifOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, reqif := range reqifOrdered {

		id = generatesIdentifier("REQIF", idx, reqif.Name)
		map_REQIF_Identifiers[reqif] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQIF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", reqif.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqif.Name))
		initializerStatements += setValueField

	}

	map_REQIFCONTENT_Identifiers := make(map[*REQIFCONTENT]string)
	_ = map_REQIFCONTENT_Identifiers

	reqifcontentOrdered := []*REQIFCONTENT{}
	for reqifcontent := range stage.REQIFCONTENTs {
		reqifcontentOrdered = append(reqifcontentOrdered, reqifcontent)
	}
	sort.Slice(reqifcontentOrdered[:], func(i, j int) bool {
		return reqifcontentOrdered[i].Name < reqifcontentOrdered[j].Name
	})
	if len(reqifcontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, reqifcontent := range reqifcontentOrdered {

		id = generatesIdentifier("REQIFCONTENT", idx, reqifcontent.Name)
		map_REQIFCONTENT_Identifiers[reqifcontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQIFCONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", reqifcontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifcontent.Name))
		initializerStatements += setValueField

	}

	map_REQIFHEADER_Identifiers := make(map[*REQIFHEADER]string)
	_ = map_REQIFHEADER_Identifiers

	reqifheaderOrdered := []*REQIFHEADER{}
	for reqifheader := range stage.REQIFHEADERs {
		reqifheaderOrdered = append(reqifheaderOrdered, reqifheader)
	}
	sort.Slice(reqifheaderOrdered[:], func(i, j int) bool {
		return reqifheaderOrdered[i].Name < reqifheaderOrdered[j].Name
	})
	if len(reqifheaderOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, reqifheader := range reqifheaderOrdered {

		id = generatesIdentifier("REQIFHEADER", idx, reqifheader.Name)
		map_REQIFHEADER_Identifiers[reqifheader] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQIFHEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", reqifheader.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "COMMENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.COMMENT))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CREATIONTIME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.CREATIONTIME))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REPOSITORYID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.REPOSITORYID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQIFTOOLID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.REQIFTOOLID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQIFVERSION")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.REQIFVERSION))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SOURCETOOLID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.SOURCETOOLID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TITLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqifheader.TITLE))
		initializerStatements += setValueField

	}

	map_REQIFTOOLEXTENSION_Identifiers := make(map[*REQIFTOOLEXTENSION]string)
	_ = map_REQIFTOOLEXTENSION_Identifiers

	reqiftoolextensionOrdered := []*REQIFTOOLEXTENSION{}
	for reqiftoolextension := range stage.REQIFTOOLEXTENSIONs {
		reqiftoolextensionOrdered = append(reqiftoolextensionOrdered, reqiftoolextension)
	}
	sort.Slice(reqiftoolextensionOrdered[:], func(i, j int) bool {
		return reqiftoolextensionOrdered[i].Name < reqiftoolextensionOrdered[j].Name
	})
	if len(reqiftoolextensionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, reqiftoolextension := range reqiftoolextensionOrdered {

		id = generatesIdentifier("REQIFTOOLEXTENSION", idx, reqiftoolextension.Name)
		map_REQIFTOOLEXTENSION_Identifiers[reqiftoolextension] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQIFTOOLEXTENSION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", reqiftoolextension.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqiftoolextension.Name))
		initializerStatements += setValueField

	}

	map_REQIFTYPE_Identifiers := make(map[*REQIFTYPE]string)
	_ = map_REQIFTYPE_Identifiers

	reqiftypeOrdered := []*REQIFTYPE{}
	for reqiftype := range stage.REQIFTYPEs {
		reqiftypeOrdered = append(reqiftypeOrdered, reqiftype)
	}
	sort.Slice(reqiftypeOrdered[:], func(i, j int) bool {
		return reqiftypeOrdered[i].Name < reqiftypeOrdered[j].Name
	})
	if len(reqiftypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, reqiftype := range reqiftypeOrdered {

		id = generatesIdentifier("REQIFTYPE", idx, reqiftype.Name)
		map_REQIFTYPE_Identifiers[reqiftype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQIFTYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", reqiftype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqiftype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONBOOLEANREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(reqiftype.DATATYPEDEFINITIONBOOLEANREF))
		initializerStatements += setValueField

	}

	map_SOURCE_Identifiers := make(map[*SOURCE]string)
	_ = map_SOURCE_Identifiers

	sourceOrdered := []*SOURCE{}
	for source := range stage.SOURCEs {
		sourceOrdered = append(sourceOrdered, source)
	}
	sort.Slice(sourceOrdered[:], func(i, j int) bool {
		return sourceOrdered[i].Name < sourceOrdered[j].Name
	})
	if len(sourceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, source := range sourceOrdered {

		id = generatesIdentifier("SOURCE", idx, source.Name)
		map_SOURCE_Identifiers[source] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SOURCE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", source.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(source.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECOBJECTREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(source.SPECOBJECTREF))
		initializerStatements += setValueField

	}

	map_SOURCESPECIFICATION_Identifiers := make(map[*SOURCESPECIFICATION]string)
	_ = map_SOURCESPECIFICATION_Identifiers

	sourcespecificationOrdered := []*SOURCESPECIFICATION{}
	for sourcespecification := range stage.SOURCESPECIFICATIONs {
		sourcespecificationOrdered = append(sourcespecificationOrdered, sourcespecification)
	}
	sort.Slice(sourcespecificationOrdered[:], func(i, j int) bool {
		return sourcespecificationOrdered[i].Name < sourcespecificationOrdered[j].Name
	})
	if len(sourcespecificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, sourcespecification := range sourcespecificationOrdered {

		id = generatesIdentifier("SOURCESPECIFICATION", idx, sourcespecification.Name)
		map_SOURCESPECIFICATION_Identifiers[sourcespecification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SOURCESPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sourcespecification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sourcespecification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATIONREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sourcespecification.SPECIFICATIONREF))
		initializerStatements += setValueField

	}

	map_SPECATTRIBUTES_Identifiers := make(map[*SPECATTRIBUTES]string)
	_ = map_SPECATTRIBUTES_Identifiers

	specattributesOrdered := []*SPECATTRIBUTES{}
	for specattributes := range stage.SPECATTRIBUTESs {
		specattributesOrdered = append(specattributesOrdered, specattributes)
	}
	sort.Slice(specattributesOrdered[:], func(i, j int) bool {
		return specattributesOrdered[i].Name < specattributesOrdered[j].Name
	})
	if len(specattributesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specattributes := range specattributesOrdered {

		id = generatesIdentifier("SPECATTRIBUTES", idx, specattributes.Name)
		map_SPECATTRIBUTES_Identifiers[specattributes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECATTRIBUTES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specattributes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specattributes.Name))
		initializerStatements += setValueField

	}

	map_SPECHIERARCHY_Identifiers := make(map[*SPECHIERARCHY]string)
	_ = map_SPECHIERARCHY_Identifiers

	spechierarchyOrdered := []*SPECHIERARCHY{}
	for spechierarchy := range stage.SPECHIERARCHYs {
		spechierarchyOrdered = append(spechierarchyOrdered, spechierarchy)
	}
	sort.Slice(spechierarchyOrdered[:], func(i, j int) bool {
		return spechierarchyOrdered[i].Name < spechierarchyOrdered[j].Name
	})
	if len(spechierarchyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spechierarchy := range spechierarchyOrdered {

		id = generatesIdentifier("SPECHIERARCHY", idx, spechierarchy.Name)
		map_SPECHIERARCHY_Identifiers[spechierarchy] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECHIERARCHY")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spechierarchy.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spechierarchy.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spechierarchy.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spechierarchy.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISEDITABLEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spechierarchy.ISEDITABLEAttr))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ISTABLEINTERNALAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spechierarchy.ISTABLEINTERNALAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spechierarchy.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spechierarchy.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_Identifiers := make(map[*SPECIFICATION]string)
	_ = map_SPECIFICATION_Identifiers

	specificationOrdered := []*SPECIFICATION{}
	for specification := range stage.SPECIFICATIONs {
		specificationOrdered = append(specificationOrdered, specification)
	}
	sort.Slice(specificationOrdered[:], func(i, j int) bool {
		return specificationOrdered[i].Name < specificationOrdered[j].Name
	})
	if len(specificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specification := range specificationOrdered {

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECIFICATIONS_Identifiers := make(map[*SPECIFICATIONS]string)
	_ = map_SPECIFICATIONS_Identifiers

	specificationsOrdered := []*SPECIFICATIONS{}
	for specifications := range stage.SPECIFICATIONSs {
		specificationsOrdered = append(specificationsOrdered, specifications)
	}
	sort.Slice(specificationsOrdered[:], func(i, j int) bool {
		return specificationsOrdered[i].Name < specificationsOrdered[j].Name
	})
	if len(specificationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specifications := range specificationsOrdered {

		id = generatesIdentifier("SPECIFICATIONS", idx, specifications.Name)
		map_SPECIFICATIONS_Identifiers[specifications] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specifications.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specifications.Name))
		initializerStatements += setValueField

	}

	map_SPECIFICATIONTYPE_Identifiers := make(map[*SPECIFICATIONTYPE]string)
	_ = map_SPECIFICATIONTYPE_Identifiers

	specificationtypeOrdered := []*SPECIFICATIONTYPE{}
	for specificationtype := range stage.SPECIFICATIONTYPEs {
		specificationtypeOrdered = append(specificationtypeOrdered, specificationtype)
	}
	sort.Slice(specificationtypeOrdered[:], func(i, j int) bool {
		return specificationtypeOrdered[i].Name < specificationtypeOrdered[j].Name
	})
	if len(specificationtypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specificationtype := range specificationtypeOrdered {

		id = generatesIdentifier("SPECIFICATIONTYPE", idx, specificationtype.Name)
		map_SPECIFICATIONTYPE_Identifiers[specificationtype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATIONTYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specificationtype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specificationtype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specificationtype.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specificationtype.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specificationtype.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specificationtype.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECIFIEDVALUES_Identifiers := make(map[*SPECIFIEDVALUES]string)
	_ = map_SPECIFIEDVALUES_Identifiers

	specifiedvaluesOrdered := []*SPECIFIEDVALUES{}
	for specifiedvalues := range stage.SPECIFIEDVALUESs {
		specifiedvaluesOrdered = append(specifiedvaluesOrdered, specifiedvalues)
	}
	sort.Slice(specifiedvaluesOrdered[:], func(i, j int) bool {
		return specifiedvaluesOrdered[i].Name < specifiedvaluesOrdered[j].Name
	})
	if len(specifiedvaluesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specifiedvalues := range specifiedvaluesOrdered {

		id = generatesIdentifier("SPECIFIEDVALUES", idx, specifiedvalues.Name)
		map_SPECIFIEDVALUES_Identifiers[specifiedvalues] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFIEDVALUES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specifiedvalues.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specifiedvalues.Name))
		initializerStatements += setValueField

	}

	map_SPECOBJECT_Identifiers := make(map[*SPECOBJECT]string)
	_ = map_SPECOBJECT_Identifiers

	specobjectOrdered := []*SPECOBJECT{}
	for specobject := range stage.SPECOBJECTs {
		specobjectOrdered = append(specobjectOrdered, specobject)
	}
	sort.Slice(specobjectOrdered[:], func(i, j int) bool {
		return specobjectOrdered[i].Name < specobjectOrdered[j].Name
	})
	if len(specobjectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specobject := range specobjectOrdered {

		id = generatesIdentifier("SPECOBJECT", idx, specobject.Name)
		map_SPECOBJECT_Identifiers[specobject] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECOBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specobject.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobject.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobject.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobject.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobject.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobject.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECOBJECTS_Identifiers := make(map[*SPECOBJECTS]string)
	_ = map_SPECOBJECTS_Identifiers

	specobjectsOrdered := []*SPECOBJECTS{}
	for specobjects := range stage.SPECOBJECTSs {
		specobjectsOrdered = append(specobjectsOrdered, specobjects)
	}
	sort.Slice(specobjectsOrdered[:], func(i, j int) bool {
		return specobjectsOrdered[i].Name < specobjectsOrdered[j].Name
	})
	if len(specobjectsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specobjects := range specobjectsOrdered {

		id = generatesIdentifier("SPECOBJECTS", idx, specobjects.Name)
		map_SPECOBJECTS_Identifiers[specobjects] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECOBJECTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specobjects.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjects.Name))
		initializerStatements += setValueField

	}

	map_SPECOBJECTTYPE_Identifiers := make(map[*SPECOBJECTTYPE]string)
	_ = map_SPECOBJECTTYPE_Identifiers

	specobjecttypeOrdered := []*SPECOBJECTTYPE{}
	for specobjecttype := range stage.SPECOBJECTTYPEs {
		specobjecttypeOrdered = append(specobjecttypeOrdered, specobjecttype)
	}
	sort.Slice(specobjecttypeOrdered[:], func(i, j int) bool {
		return specobjecttypeOrdered[i].Name < specobjecttypeOrdered[j].Name
	})
	if len(specobjecttypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specobjecttype := range specobjecttypeOrdered {

		id = generatesIdentifier("SPECOBJECTTYPE", idx, specobjecttype.Name)
		map_SPECOBJECTTYPE_Identifiers[specobjecttype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECOBJECTTYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specobjecttype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjecttype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjecttype.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjecttype.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjecttype.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specobjecttype.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECRELATION_Identifiers := make(map[*SPECRELATION]string)
	_ = map_SPECRELATION_Identifiers

	specrelationOrdered := []*SPECRELATION{}
	for specrelation := range stage.SPECRELATIONs {
		specrelationOrdered = append(specrelationOrdered, specrelation)
	}
	sort.Slice(specrelationOrdered[:], func(i, j int) bool {
		return specrelationOrdered[i].Name < specrelationOrdered[j].Name
	})
	if len(specrelationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specrelation := range specrelationOrdered {

		id = generatesIdentifier("SPECRELATION", idx, specrelation.Name)
		map_SPECRELATION_Identifiers[specrelation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECRELATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specrelation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelation.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelation.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelation.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelation.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECRELATIONGROUPS_Identifiers := make(map[*SPECRELATIONGROUPS]string)
	_ = map_SPECRELATIONGROUPS_Identifiers

	specrelationgroupsOrdered := []*SPECRELATIONGROUPS{}
	for specrelationgroups := range stage.SPECRELATIONGROUPSs {
		specrelationgroupsOrdered = append(specrelationgroupsOrdered, specrelationgroups)
	}
	sort.Slice(specrelationgroupsOrdered[:], func(i, j int) bool {
		return specrelationgroupsOrdered[i].Name < specrelationgroupsOrdered[j].Name
	})
	if len(specrelationgroupsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specrelationgroups := range specrelationgroupsOrdered {

		id = generatesIdentifier("SPECRELATIONGROUPS", idx, specrelationgroups.Name)
		map_SPECRELATIONGROUPS_Identifiers[specrelationgroups] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECRELATIONGROUPS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specrelationgroups.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationgroups.Name))
		initializerStatements += setValueField

	}

	map_SPECRELATIONS_Identifiers := make(map[*SPECRELATIONS]string)
	_ = map_SPECRELATIONS_Identifiers

	specrelationsOrdered := []*SPECRELATIONS{}
	for specrelations := range stage.SPECRELATIONSs {
		specrelationsOrdered = append(specrelationsOrdered, specrelations)
	}
	sort.Slice(specrelationsOrdered[:], func(i, j int) bool {
		return specrelationsOrdered[i].Name < specrelationsOrdered[j].Name
	})
	if len(specrelationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specrelations := range specrelationsOrdered {

		id = generatesIdentifier("SPECRELATIONS", idx, specrelations.Name)
		map_SPECRELATIONS_Identifiers[specrelations] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECRELATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specrelations.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelations.Name))
		initializerStatements += setValueField

	}

	map_SPECRELATIONTYPE_Identifiers := make(map[*SPECRELATIONTYPE]string)
	_ = map_SPECRELATIONTYPE_Identifiers

	specrelationtypeOrdered := []*SPECRELATIONTYPE{}
	for specrelationtype := range stage.SPECRELATIONTYPEs {
		specrelationtypeOrdered = append(specrelationtypeOrdered, specrelationtype)
	}
	sort.Slice(specrelationtypeOrdered[:], func(i, j int) bool {
		return specrelationtypeOrdered[i].Name < specrelationtypeOrdered[j].Name
	})
	if len(specrelationtypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specrelationtype := range specrelationtypeOrdered {

		id = generatesIdentifier("SPECRELATIONTYPE", idx, specrelationtype.Name)
		map_SPECRELATIONTYPE_Identifiers[specrelationtype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECRELATIONTYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specrelationtype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationtype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESCAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationtype.DESCAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIERAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationtype.IDENTIFIERAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LASTCHANGEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationtype.LASTCHANGEAttr))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONGNAMEAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specrelationtype.LONGNAMEAttr))
		initializerStatements += setValueField

	}

	map_SPECTYPES_Identifiers := make(map[*SPECTYPES]string)
	_ = map_SPECTYPES_Identifiers

	spectypesOrdered := []*SPECTYPES{}
	for spectypes := range stage.SPECTYPESs {
		spectypesOrdered = append(spectypesOrdered, spectypes)
	}
	sort.Slice(spectypesOrdered[:], func(i, j int) bool {
		return spectypesOrdered[i].Name < spectypesOrdered[j].Name
	})
	if len(spectypesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spectypes := range spectypesOrdered {

		id = generatesIdentifier("SPECTYPES", idx, spectypes.Name)
		map_SPECTYPES_Identifiers[spectypes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECTYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spectypes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spectypes.Name))
		initializerStatements += setValueField

	}

	map_TARGET_Identifiers := make(map[*TARGET]string)
	_ = map_TARGET_Identifiers

	targetOrdered := []*TARGET{}
	for target := range stage.TARGETs {
		targetOrdered = append(targetOrdered, target)
	}
	sort.Slice(targetOrdered[:], func(i, j int) bool {
		return targetOrdered[i].Name < targetOrdered[j].Name
	})
	if len(targetOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, target := range targetOrdered {

		id = generatesIdentifier("TARGET", idx, target.Name)
		map_TARGET_Identifiers[target] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TARGET")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", target.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(target.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECOBJECTREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(target.SPECOBJECTREF))
		initializerStatements += setValueField

	}

	map_TARGETSPECIFICATION_Identifiers := make(map[*TARGETSPECIFICATION]string)
	_ = map_TARGETSPECIFICATION_Identifiers

	targetspecificationOrdered := []*TARGETSPECIFICATION{}
	for targetspecification := range stage.TARGETSPECIFICATIONs {
		targetspecificationOrdered = append(targetspecificationOrdered, targetspecification)
	}
	sort.Slice(targetspecificationOrdered[:], func(i, j int) bool {
		return targetspecificationOrdered[i].Name < targetspecificationOrdered[j].Name
	})
	if len(targetspecificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, targetspecification := range targetspecificationOrdered {

		id = generatesIdentifier("TARGETSPECIFICATION", idx, targetspecification.Name)
		map_TARGETSPECIFICATION_Identifiers[targetspecification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TARGETSPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", targetspecification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(targetspecification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATIONREF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(targetspecification.SPECIFICATIONREF))
		initializerStatements += setValueField

	}

	map_THEHEADER_Identifiers := make(map[*THEHEADER]string)
	_ = map_THEHEADER_Identifiers

	theheaderOrdered := []*THEHEADER{}
	for theheader := range stage.THEHEADERs {
		theheaderOrdered = append(theheaderOrdered, theheader)
	}
	sort.Slice(theheaderOrdered[:], func(i, j int) bool {
		return theheaderOrdered[i].Name < theheaderOrdered[j].Name
	})
	if len(theheaderOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, theheader := range theheaderOrdered {

		id = generatesIdentifier("THEHEADER", idx, theheader.Name)
		map_THEHEADER_Identifiers[theheader] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "THEHEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", theheader.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(theheader.Name))
		initializerStatements += setValueField

	}

	map_TOOLEXTENSIONS_Identifiers := make(map[*TOOLEXTENSIONS]string)
	_ = map_TOOLEXTENSIONS_Identifiers

	toolextensionsOrdered := []*TOOLEXTENSIONS{}
	for toolextensions := range stage.TOOLEXTENSIONSs {
		toolextensionsOrdered = append(toolextensionsOrdered, toolextensions)
	}
	sort.Slice(toolextensionsOrdered[:], func(i, j int) bool {
		return toolextensionsOrdered[i].Name < toolextensionsOrdered[j].Name
	})
	if len(toolextensionsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, toolextensions := range toolextensionsOrdered {

		id = generatesIdentifier("TOOLEXTENSIONS", idx, toolextensions.Name)
		map_TOOLEXTENSIONS_Identifiers[toolextensions] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TOOLEXTENSIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", toolextensions.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(toolextensions.Name))
		initializerStatements += setValueField

	}

	map_VALUES_Identifiers := make(map[*VALUES]string)
	_ = map_VALUES_Identifiers

	valuesOrdered := []*VALUES{}
	for values := range stage.VALUESs {
		valuesOrdered = append(valuesOrdered, values)
	}
	sort.Slice(valuesOrdered[:], func(i, j int) bool {
		return valuesOrdered[i].Name < valuesOrdered[j].Name
	})
	if len(valuesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, values := range valuesOrdered {

		id = generatesIdentifier("VALUES", idx, values.Name)
		map_VALUES_Identifiers[values] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VALUES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", values.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(values.Name))
		initializerStatements += setValueField

	}

	map_XHTMLCONTENT_Identifiers := make(map[*XHTMLCONTENT]string)
	_ = map_XHTMLCONTENT_Identifiers

	xhtmlcontentOrdered := []*XHTMLCONTENT{}
	for xhtmlcontent := range stage.XHTMLCONTENTs {
		xhtmlcontentOrdered = append(xhtmlcontentOrdered, xhtmlcontent)
	}
	sort.Slice(xhtmlcontentOrdered[:], func(i, j int) bool {
		return xhtmlcontentOrdered[i].Name < xhtmlcontentOrdered[j].Name
	})
	if len(xhtmlcontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtmlcontent := range xhtmlcontentOrdered {

		id = generatesIdentifier("XHTMLCONTENT", idx, xhtmlcontent.Name)
		map_XHTMLCONTENT_Identifiers[xhtmlcontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XHTMLCONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtmlcontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtmlcontent.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, alternativeid := range alternativeidOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ALTERNATIVEID", idx, alternativeid.Name)
		map_ALTERNATIVEID_Identifiers[alternativeid] = id

		// Initialisation of values
	}

	for idx, attributedefinitionboolean := range attributedefinitionbooleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONBOOLEAN", idx, attributedefinitionboolean.Name)
		map_ATTRIBUTEDEFINITIONBOOLEAN_Identifiers[attributedefinitionboolean] = id

		// Initialisation of values
		if attributedefinitionboolean.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitionboolean.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionboolean.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitionboolean.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionboolean.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitionboolean.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitiondate := range attributedefinitiondateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONDATE", idx, attributedefinitiondate.Name)
		map_ATTRIBUTEDEFINITIONDATE_Identifiers[attributedefinitiondate] = id

		// Initialisation of values
		if attributedefinitiondate.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitiondate.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitiondate.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitiondate.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitiondate.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitiondate.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitionenumeration := range attributedefinitionenumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONENUMERATION", idx, attributedefinitionenumeration.Name)
		map_ATTRIBUTEDEFINITIONENUMERATION_Identifiers[attributedefinitionenumeration] = id

		// Initialisation of values
		if attributedefinitionenumeration.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitionenumeration.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionenumeration.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitionenumeration.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionenumeration.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitionenumeration.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitioninteger := range attributedefinitionintegerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONINTEGER", idx, attributedefinitioninteger.Name)
		map_ATTRIBUTEDEFINITIONINTEGER_Identifiers[attributedefinitioninteger] = id

		// Initialisation of values
		if attributedefinitioninteger.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitioninteger.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitioninteger.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitioninteger.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitioninteger.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitioninteger.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitionreal := range attributedefinitionrealOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONREAL", idx, attributedefinitionreal.Name)
		map_ATTRIBUTEDEFINITIONREAL_Identifiers[attributedefinitionreal] = id

		// Initialisation of values
		if attributedefinitionreal.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitionreal.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionreal.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitionreal.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionreal.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitionreal.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitionstring := range attributedefinitionstringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONSTRING", idx, attributedefinitionstring.Name)
		map_ATTRIBUTEDEFINITIONSTRING_Identifiers[attributedefinitionstring] = id

		// Initialisation of values
		if attributedefinitionstring.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitionstring.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionstring.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitionstring.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionstring.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitionstring.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributedefinitionxhtml := range attributedefinitionxhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEDEFINITIONXHTML", idx, attributedefinitionxhtml.Name)
		map_ATTRIBUTEDEFINITIONXHTML_Identifiers[attributedefinitionxhtml] = id

		// Initialisation of values
		if attributedefinitionxhtml.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[attributedefinitionxhtml.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionxhtml.DEFAULTVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULTVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFAULTVALUE_Identifiers[attributedefinitionxhtml.DEFAULTVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributedefinitionxhtml.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[attributedefinitionxhtml.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevalueboolean := range attributevaluebooleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEBOOLEAN", idx, attributevalueboolean.Name)
		map_ATTRIBUTEVALUEBOOLEAN_Identifiers[attributevalueboolean] = id

		// Initialisation of values
		if attributevalueboolean.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevalueboolean.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevaluedate := range attributevaluedateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEDATE", idx, attributevaluedate.Name)
		map_ATTRIBUTEVALUEDATE_Identifiers[attributevaluedate] = id

		// Initialisation of values
		if attributevaluedate.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevaluedate.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevalueenumeration := range attributevalueenumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEENUMERATION", idx, attributevalueenumeration.Name)
		map_ATTRIBUTEVALUEENUMERATION_Identifiers[attributevalueenumeration] = id

		// Initialisation of values
		if attributevalueenumeration.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevalueenumeration.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

		if attributevalueenumeration.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_VALUES_Identifiers[attributevalueenumeration.VALUES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevalueinteger := range attributevalueintegerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEINTEGER", idx, attributevalueinteger.Name)
		map_ATTRIBUTEVALUEINTEGER_Identifiers[attributevalueinteger] = id

		// Initialisation of values
		if attributevalueinteger.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevalueinteger.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevaluereal := range attributevaluerealOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEREAL", idx, attributevaluereal.Name)
		map_ATTRIBUTEVALUEREAL_Identifiers[attributevaluereal] = id

		// Initialisation of values
		if attributevaluereal.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevaluereal.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevaluestring := range attributevaluestringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUESTRING", idx, attributevaluestring.Name)
		map_ATTRIBUTEVALUESTRING_Identifiers[attributevaluestring] = id

		// Initialisation of values
		if attributevaluestring.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevaluestring.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributevaluexhtml := range attributevaluexhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTEVALUEXHTML", idx, attributevaluexhtml.Name)
		map_ATTRIBUTEVALUEXHTML_Identifiers[attributevaluexhtml] = id

		// Initialisation of values
		if attributevaluexhtml.THEVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THEVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTMLCONTENT_Identifiers[attributevaluexhtml.THEVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributevaluexhtml.THEORIGINALVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THEORIGINALVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTMLCONTENT_Identifiers[attributevaluexhtml.THEORIGINALVALUE])
			pointersInitializesStatements += setPointerField
		}

		if attributevaluexhtml.DEFINITION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DEFINITION_Identifiers[attributevaluexhtml.DEFINITION])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, children := range childrenOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("CHILDREN", idx, children.Name)
		map_CHILDREN_Identifiers[children] = id

		// Initialisation of values
		for _, _spechierarchy := range children.SPECHIERARCHY {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECHIERARCHY")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECHIERARCHY_Identifiers[_spechierarchy])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, corecontent := range corecontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("CORECONTENT", idx, corecontent.Name)
		map_CORECONTENT_Identifiers[corecontent] = id

		// Initialisation of values
		if corecontent.REQIFCONTENT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQIFCONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFCONTENT_Identifiers[corecontent.REQIFCONTENT])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitionboolean := range datatypedefinitionbooleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONBOOLEAN", idx, datatypedefinitionboolean.Name)
		map_DATATYPEDEFINITIONBOOLEAN_Identifiers[datatypedefinitionboolean] = id

		// Initialisation of values
		if datatypedefinitionboolean.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitionboolean.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitiondate := range datatypedefinitiondateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONDATE", idx, datatypedefinitiondate.Name)
		map_DATATYPEDEFINITIONDATE_Identifiers[datatypedefinitiondate] = id

		// Initialisation of values
		if datatypedefinitiondate.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitiondate.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitionenumeration := range datatypedefinitionenumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONENUMERATION", idx, datatypedefinitionenumeration.Name)
		map_DATATYPEDEFINITIONENUMERATION_Identifiers[datatypedefinitionenumeration] = id

		// Initialisation of values
		if datatypedefinitionenumeration.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitionenumeration.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if datatypedefinitionenumeration.SPECIFIEDVALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFIEDVALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFIEDVALUES_Identifiers[datatypedefinitionenumeration.SPECIFIEDVALUES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitioninteger := range datatypedefinitionintegerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONINTEGER", idx, datatypedefinitioninteger.Name)
		map_DATATYPEDEFINITIONINTEGER_Identifiers[datatypedefinitioninteger] = id

		// Initialisation of values
		if datatypedefinitioninteger.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitioninteger.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitionreal := range datatypedefinitionrealOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONREAL", idx, datatypedefinitionreal.Name)
		map_DATATYPEDEFINITIONREAL_Identifiers[datatypedefinitionreal] = id

		// Initialisation of values
		if datatypedefinitionreal.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitionreal.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitionstring := range datatypedefinitionstringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONSTRING", idx, datatypedefinitionstring.Name)
		map_DATATYPEDEFINITIONSTRING_Identifiers[datatypedefinitionstring] = id

		// Initialisation of values
		if datatypedefinitionstring.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitionstring.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypedefinitionxhtml := range datatypedefinitionxhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPEDEFINITIONXHTML", idx, datatypedefinitionxhtml.Name)
		map_DATATYPEDEFINITIONXHTML_Identifiers[datatypedefinitionxhtml] = id

		// Initialisation of values
		if datatypedefinitionxhtml.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[datatypedefinitionxhtml.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatypes := range datatypesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPES", idx, datatypes.Name)
		map_DATATYPES_Identifiers[datatypes] = id

		// Initialisation of values
		for _, _datatypedefinitionboolean := range datatypes.DATATYPEDEFINITIONBOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONBOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONBOOLEAN_Identifiers[_datatypedefinitionboolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitiondate := range datatypes.DATATYPEDEFINITIONDATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONDATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONDATE_Identifiers[_datatypedefinitiondate])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitionenumeration := range datatypes.DATATYPEDEFINITIONENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONENUMERATION_Identifiers[_datatypedefinitionenumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitioninteger := range datatypes.DATATYPEDEFINITIONINTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONINTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONINTEGER_Identifiers[_datatypedefinitioninteger])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitionreal := range datatypes.DATATYPEDEFINITIONREAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONREAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONREAL_Identifiers[_datatypedefinitionreal])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitionstring := range datatypes.DATATYPEDEFINITIONSTRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONSTRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONSTRING_Identifiers[_datatypedefinitionstring])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatypedefinitionxhtml := range datatypes.DATATYPEDEFINITIONXHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPEDEFINITIONXHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPEDEFINITIONXHTML_Identifiers[_datatypedefinitionxhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, defaultvalue := range defaultvalueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DEFAULTVALUE", idx, defaultvalue.Name)
		map_DEFAULTVALUE_Identifiers[defaultvalue] = id

		// Initialisation of values
		if defaultvalue.ATTRIBUTEVALUEBOOLEAN != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEVALUEBOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEVALUEBOOLEAN_Identifiers[defaultvalue.ATTRIBUTEVALUEBOOLEAN])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, definition := range definitionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DEFINITION", idx, definition.Name)
		map_DEFINITION_Identifiers[definition] = id

		// Initialisation of values
	}

	for idx, editableatts := range editableattsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EDITABLEATTS", idx, editableatts.Name)
		map_EDITABLEATTS_Identifiers[editableatts] = id

		// Initialisation of values
	}

	for idx, embeddedvalue := range embeddedvalueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EMBEDDEDVALUE", idx, embeddedvalue.Name)
		map_EMBEDDEDVALUE_Identifiers[embeddedvalue] = id

		// Initialisation of values
	}

	for idx, enumvalue := range enumvalueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ENUMVALUE", idx, enumvalue.Name)
		map_ENUMVALUE_Identifiers[enumvalue] = id

		// Initialisation of values
		if enumvalue.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[enumvalue.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if enumvalue.PROPERTIES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PROPERTIES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_PROPERTIES_Identifiers[enumvalue.PROPERTIES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, object := range objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("OBJECT", idx, object.Name)
		map_OBJECT_Identifiers[object] = id

		// Initialisation of values
	}

	for idx, properties := range propertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("PROPERTIES", idx, properties.Name)
		map_PROPERTIES_Identifiers[properties] = id

		// Initialisation of values
		if properties.EMBEDDEDVALUE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EMBEDDEDVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_EMBEDDEDVALUE_Identifiers[properties.EMBEDDEDVALUE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, relationgroup := range relationgroupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATIONGROUP", idx, relationgroup.Name)
		map_RELATIONGROUP_Identifiers[relationgroup] = id

		// Initialisation of values
		if relationgroup.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[relationgroup.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if relationgroup.SOURCESPECIFICATION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCESPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SOURCESPECIFICATION_Identifiers[relationgroup.SOURCESPECIFICATION])
			pointersInitializesStatements += setPointerField
		}

		if relationgroup.SPECRELATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECRELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECRELATIONS_Identifiers[relationgroup.SPECRELATIONS])
			pointersInitializesStatements += setPointerField
		}

		if relationgroup.TARGETSPECIFICATION != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TARGETSPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TARGETSPECIFICATION_Identifiers[relationgroup.TARGETSPECIFICATION])
			pointersInitializesStatements += setPointerField
		}

		if relationgroup.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[relationgroup.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, relationgrouptype := range relationgrouptypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATIONGROUPTYPE", idx, relationgrouptype.Name)
		map_RELATIONGROUPTYPE_Identifiers[relationgrouptype] = id

		// Initialisation of values
		if relationgrouptype.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[relationgrouptype.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if relationgrouptype.SPECATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECATTRIBUTES_Identifiers[relationgrouptype.SPECATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, reqif := range reqifOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQIF", idx, reqif.Name)
		map_REQIF_Identifiers[reqif] = id

		// Initialisation of values
		if reqif.HEADER != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_THEHEADER_Identifiers[reqif.HEADER])
			pointersInitializesStatements += setPointerField
		}

		if reqif.CORECONTENT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CORECONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CORECONTENT_Identifiers[reqif.CORECONTENT])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, reqifcontent := range reqifcontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQIFCONTENT", idx, reqifcontent.Name)
		map_REQIFCONTENT_Identifiers[reqifcontent] = id

		// Initialisation of values
		if reqifcontent.DATATYPES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPES_Identifiers[reqifcontent.DATATYPES])
			pointersInitializesStatements += setPointerField
		}

		if reqifcontent.SPECTYPES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECTYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECTYPES_Identifiers[reqifcontent.SPECTYPES])
			pointersInitializesStatements += setPointerField
		}

		if reqifcontent.SPECOBJECTS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECOBJECTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECOBJECTS_Identifiers[reqifcontent.SPECOBJECTS])
			pointersInitializesStatements += setPointerField
		}

		if reqifcontent.SPECRELATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECRELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECRELATIONS_Identifiers[reqifcontent.SPECRELATIONS])
			pointersInitializesStatements += setPointerField
		}

		if reqifcontent.SPECIFICATIONS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATIONS_Identifiers[reqifcontent.SPECIFICATIONS])
			pointersInitializesStatements += setPointerField
		}

		if reqifcontent.SPECRELATIONGROUPS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECRELATIONGROUPS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECRELATIONGROUPS_Identifiers[reqifcontent.SPECRELATIONGROUPS])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, reqifheader := range reqifheaderOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQIFHEADER", idx, reqifheader.Name)
		map_REQIFHEADER_Identifiers[reqifheader] = id

		// Initialisation of values
	}

	for idx, reqiftoolextension := range reqiftoolextensionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQIFTOOLEXTENSION", idx, reqiftoolextension.Name)
		map_REQIFTOOLEXTENSION_Identifiers[reqiftoolextension] = id

		// Initialisation of values
	}

	for idx, reqiftype := range reqiftypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQIFTYPE", idx, reqiftype.Name)
		map_REQIFTYPE_Identifiers[reqiftype] = id

		// Initialisation of values
	}

	for idx, source := range sourceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SOURCE", idx, source.Name)
		map_SOURCE_Identifiers[source] = id

		// Initialisation of values
	}

	for idx, sourcespecification := range sourcespecificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SOURCESPECIFICATION", idx, sourcespecification.Name)
		map_SOURCESPECIFICATION_Identifiers[sourcespecification] = id

		// Initialisation of values
	}

	for idx, specattributes := range specattributesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECATTRIBUTES", idx, specattributes.Name)
		map_SPECATTRIBUTES_Identifiers[specattributes] = id

		// Initialisation of values
		for _, _attributedefinitionboolean := range specattributes.ATTRIBUTEDEFINITIONBOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONBOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONBOOLEAN_Identifiers[_attributedefinitionboolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitiondate := range specattributes.ATTRIBUTEDEFINITIONDATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONDATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONDATE_Identifiers[_attributedefinitiondate])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitionenumeration := range specattributes.ATTRIBUTEDEFINITIONENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONENUMERATION_Identifiers[_attributedefinitionenumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitioninteger := range specattributes.ATTRIBUTEDEFINITIONINTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONINTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONINTEGER_Identifiers[_attributedefinitioninteger])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitionreal := range specattributes.ATTRIBUTEDEFINITIONREAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONREAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONREAL_Identifiers[_attributedefinitionreal])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitionstring := range specattributes.ATTRIBUTEDEFINITIONSTRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONSTRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONSTRING_Identifiers[_attributedefinitionstring])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributedefinitionxhtml := range specattributes.ATTRIBUTEDEFINITIONXHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTEDEFINITIONXHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTEDEFINITIONXHTML_Identifiers[_attributedefinitionxhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spechierarchy := range spechierarchyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECHIERARCHY", idx, spechierarchy.Name)
		map_SPECHIERARCHY_Identifiers[spechierarchy] = id

		// Initialisation of values
		if spechierarchy.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[spechierarchy.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if spechierarchy.CHILDREN != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CHILDREN_Identifiers[spechierarchy.CHILDREN])
			pointersInitializesStatements += setPointerField
		}

		if spechierarchy.EDITABLEATTS != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EDITABLEATTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_EDITABLEATTS_Identifiers[spechierarchy.EDITABLEATTS])
			pointersInitializesStatements += setPointerField
		}

		if spechierarchy.OBJECT != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_OBJECT_Identifiers[spechierarchy.OBJECT])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specification := range specificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		// Initialisation of values
		if specification.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specification.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specification.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_VALUES_Identifiers[specification.VALUES])
			pointersInitializesStatements += setPointerField
		}

		if specification.CHILDREN != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CHILDREN_Identifiers[specification.CHILDREN])
			pointersInitializesStatements += setPointerField
		}

		if specification.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[specification.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specifications := range specificationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATIONS", idx, specifications.Name)
		map_SPECIFICATIONS_Identifiers[specifications] = id

		// Initialisation of values
		for _, _specification := range specifications.SPECIFICATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_Identifiers[_specification])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specificationtype := range specificationtypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATIONTYPE", idx, specificationtype.Name)
		map_SPECIFICATIONTYPE_Identifiers[specificationtype] = id

		// Initialisation of values
		if specificationtype.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specificationtype.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specificationtype.SPECATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECATTRIBUTES_Identifiers[specificationtype.SPECATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specifiedvalues := range specifiedvaluesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFIEDVALUES", idx, specifiedvalues.Name)
		map_SPECIFIEDVALUES_Identifiers[specifiedvalues] = id

		// Initialisation of values
		for _, _enumvalue := range specifiedvalues.ENUMVALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ENUMVALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ENUMVALUE_Identifiers[_enumvalue])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specobject := range specobjectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECOBJECT", idx, specobject.Name)
		map_SPECOBJECT_Identifiers[specobject] = id

		// Initialisation of values
		if specobject.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specobject.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specobject.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_VALUES_Identifiers[specobject.VALUES])
			pointersInitializesStatements += setPointerField
		}

		if specobject.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[specobject.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specobjects := range specobjectsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECOBJECTS", idx, specobjects.Name)
		map_SPECOBJECTS_Identifiers[specobjects] = id

		// Initialisation of values
		for _, _specobject := range specobjects.SPECOBJECT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECOBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECOBJECT_Identifiers[_specobject])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specobjecttype := range specobjecttypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECOBJECTTYPE", idx, specobjecttype.Name)
		map_SPECOBJECTTYPE_Identifiers[specobjecttype] = id

		// Initialisation of values
		if specobjecttype.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specobjecttype.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specobjecttype.SPECATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECATTRIBUTES_Identifiers[specobjecttype.SPECATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specrelation := range specrelationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECRELATION", idx, specrelation.Name)
		map_SPECRELATION_Identifiers[specrelation] = id

		// Initialisation of values
		if specrelation.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specrelation.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specrelation.VALUES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_VALUES_Identifiers[specrelation.VALUES])
			pointersInitializesStatements += setPointerField
		}

		if specrelation.SOURCE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SOURCE_Identifiers[specrelation.SOURCE])
			pointersInitializesStatements += setPointerField
		}

		if specrelation.TARGET != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TARGET")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TARGET_Identifiers[specrelation.TARGET])
			pointersInitializesStatements += setPointerField
		}

		if specrelation.TYPE != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTYPE_Identifiers[specrelation.TYPE])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specrelationgroups := range specrelationgroupsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECRELATIONGROUPS", idx, specrelationgroups.Name)
		map_SPECRELATIONGROUPS_Identifiers[specrelationgroups] = id

		// Initialisation of values
		for _, _relationgroup := range specrelationgroups.RELATIONGROUP {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATIONGROUP")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATIONGROUP_Identifiers[_relationgroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specrelations := range specrelationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECRELATIONS", idx, specrelations.Name)
		map_SPECRELATIONS_Identifiers[specrelations] = id

		// Initialisation of values
	}

	for idx, specrelationtype := range specrelationtypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECRELATIONTYPE", idx, specrelationtype.Name)
		map_SPECRELATIONTYPE_Identifiers[specrelationtype] = id

		// Initialisation of values
		if specrelationtype.ALTERNATIVEID != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVEID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVEID_Identifiers[specrelationtype.ALTERNATIVEID])
			pointersInitializesStatements += setPointerField
		}

		if specrelationtype.SPECATTRIBUTES != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECATTRIBUTES_Identifiers[specrelationtype.SPECATTRIBUTES])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spectypes := range spectypesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECTYPES", idx, spectypes.Name)
		map_SPECTYPES_Identifiers[spectypes] = id

		// Initialisation of values
		for _, _relationgrouptype := range spectypes.RELATIONGROUPTYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATIONGROUPTYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATIONGROUPTYPE_Identifiers[_relationgrouptype])
			pointersInitializesStatements += setPointerField
		}

		for _, _specobjecttype := range spectypes.SPECOBJECTTYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECOBJECTTYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECOBJECTTYPE_Identifiers[_specobjecttype])
			pointersInitializesStatements += setPointerField
		}

		for _, _specrelationtype := range spectypes.SPECRELATIONTYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECRELATIONTYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECRELATIONTYPE_Identifiers[_specrelationtype])
			pointersInitializesStatements += setPointerField
		}

		for _, _specificationtype := range spectypes.SPECIFICATIONTYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATIONTYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATIONTYPE_Identifiers[_specificationtype])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, target := range targetOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TARGET", idx, target.Name)
		map_TARGET_Identifiers[target] = id

		// Initialisation of values
	}

	for idx, targetspecification := range targetspecificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TARGETSPECIFICATION", idx, targetspecification.Name)
		map_TARGETSPECIFICATION_Identifiers[targetspecification] = id

		// Initialisation of values
	}

	for idx, theheader := range theheaderOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("THEHEADER", idx, theheader.Name)
		map_THEHEADER_Identifiers[theheader] = id

		// Initialisation of values
		if theheader.REQIFHEADER != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQIFHEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFHEADER_Identifiers[theheader.REQIFHEADER])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, toolextensions := range toolextensionsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TOOLEXTENSIONS", idx, toolextensions.Name)
		map_TOOLEXTENSIONS_Identifiers[toolextensions] = id

		// Initialisation of values
		for _, _reqiftoolextension := range toolextensions.REQIFTOOLEXTENSION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQIFTOOLEXTENSION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQIFTOOLEXTENSION_Identifiers[_reqiftoolextension])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, values := range valuesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("VALUES", idx, values.Name)
		map_VALUES_Identifiers[values] = id

		// Initialisation of values
	}

	for idx, xhtmlcontent := range xhtmlcontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XHTMLCONTENT", idx, xhtmlcontent.Name)
		map_XHTMLCONTENT_Identifiers[xhtmlcontent] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
