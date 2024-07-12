// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *StageStruct, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelize[ALTERNATIVEID](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONBOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONDATE](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONINTEGER](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONREAL](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONSTRING](stage, f)
		SerializeExcelize[ATTRIBUTEDEFINITIONXHTML](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEBOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEDATE](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEINTEGER](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEREAL](stage, f)
		SerializeExcelize[ATTRIBUTEVALUESTRING](stage, f)
		SerializeExcelize[ATTRIBUTEVALUEXHTML](stage, f)
		SerializeExcelize[CHILDREN](stage, f)
		SerializeExcelize[CORECONTENT](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONBOOLEAN](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONDATE](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONENUMERATION](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONINTEGER](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONREAL](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONSTRING](stage, f)
		SerializeExcelize[DATATYPEDEFINITIONXHTML](stage, f)
		SerializeExcelize[DATATYPES](stage, f)
		SerializeExcelize[DEFAULTVALUE](stage, f)
		SerializeExcelize[DEFINITION](stage, f)
		SerializeExcelize[EDITABLEATTS](stage, f)
		SerializeExcelize[EMBEDDEDVALUE](stage, f)
		SerializeExcelize[ENUMVALUE](stage, f)
		SerializeExcelize[OBJECT](stage, f)
		SerializeExcelize[PROPERTIES](stage, f)
		SerializeExcelize[RELATIONGROUP](stage, f)
		SerializeExcelize[RELATIONGROUPTYPE](stage, f)
		SerializeExcelize[REQIF](stage, f)
		SerializeExcelize[REQIFCONTENT](stage, f)
		SerializeExcelize[REQIFHEADER](stage, f)
		SerializeExcelize[REQIFTOOLEXTENSION](stage, f)
		SerializeExcelize[REQIFTYPE](stage, f)
		SerializeExcelize[SOURCE](stage, f)
		SerializeExcelize[SOURCESPECIFICATION](stage, f)
		SerializeExcelize[SPECATTRIBUTES](stage, f)
		SerializeExcelize[SPECHIERARCHY](stage, f)
		SerializeExcelize[SPECIFICATION](stage, f)
		SerializeExcelize[SPECIFICATIONS](stage, f)
		SerializeExcelize[SPECIFICATIONTYPE](stage, f)
		SerializeExcelize[SPECIFIEDVALUES](stage, f)
		SerializeExcelize[SPECOBJECT](stage, f)
		SerializeExcelize[SPECOBJECTS](stage, f)
		SerializeExcelize[SPECOBJECTTYPE](stage, f)
		SerializeExcelize[SPECRELATION](stage, f)
		SerializeExcelize[SPECRELATIONGROUPS](stage, f)
		SerializeExcelize[SPECRELATIONS](stage, f)
		SerializeExcelize[SPECRELATIONTYPE](stage, f)
		SerializeExcelize[SPECTYPES](stage, f)
		SerializeExcelize[TARGET](stage, f)
		SerializeExcelize[TARGETSPECIFICATION](stage, f)
		SerializeExcelize[THEHEADER](stage, f)
		SerializeExcelize[TOOLEXTENSIONS](stage, f)
		SerializeExcelize[VALUES](stage, f)
		SerializeExcelize[XHTMLCONTENT](stage, f)
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}

}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

func Serialize[Type Gongstruct](stage *StageStruct, tab Tabulator) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	tab.AddSheet(sheetName)

	headerRowIndex := tab.AddRow(sheetName)
	for colIndex, fieldName := range GetFields[Type]() {
		tab.AddCell(sheetName, headerRowIndex, colIndex, fieldName)
		// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}

	set := *GetGongstructInstancesSet[Type](stage)
	for instance := range set {
		line := tab.AddRow(sheetName)
		for index, fieldName := range GetFields[Type]() {
			tab.AddCell(sheetName, line, index, GetFieldStringValue(
				any(*instance).(Type), fieldName))
			// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
			// 	any(*instance).(Type), fieldName))
		}
	}
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

func SerializeExcelize[Type Gongstruct](stage *StageStruct, f *excelize.File) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSet[Type](stage)
	line := 1

	for index, fieldName := range GetFields[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for instance := range set {
		line = line + 1
		for index, fieldName := range GetFields[Type]() {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
				any(*instance).(Type), fieldName))
		}
	}

	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Panicln("SerializeExcelize")
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			log.Panicln("SerializeExcelize")
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
