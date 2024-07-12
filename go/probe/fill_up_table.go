// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.ALTERNATIVEID:
		fillUpTable[models.ALTERNATIVEID](probe)
	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		fillUpTable[models.ATTRIBUTEDEFINITIONBOOLEAN](probe)
	case *models.ATTRIBUTEDEFINITIONDATE:
		fillUpTable[models.ATTRIBUTEDEFINITIONDATE](probe)
	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		fillUpTable[models.ATTRIBUTEDEFINITIONENUMERATION](probe)
	case *models.ATTRIBUTEDEFINITIONINTEGER:
		fillUpTable[models.ATTRIBUTEDEFINITIONINTEGER](probe)
	case *models.ATTRIBUTEDEFINITIONREAL:
		fillUpTable[models.ATTRIBUTEDEFINITIONREAL](probe)
	case *models.ATTRIBUTEDEFINITIONSTRING:
		fillUpTable[models.ATTRIBUTEDEFINITIONSTRING](probe)
	case *models.ATTRIBUTEDEFINITIONXHTML:
		fillUpTable[models.ATTRIBUTEDEFINITIONXHTML](probe)
	case *models.ATTRIBUTEVALUEBOOLEAN:
		fillUpTable[models.ATTRIBUTEVALUEBOOLEAN](probe)
	case *models.ATTRIBUTEVALUEDATE:
		fillUpTable[models.ATTRIBUTEVALUEDATE](probe)
	case *models.ATTRIBUTEVALUEENUMERATION:
		fillUpTable[models.ATTRIBUTEVALUEENUMERATION](probe)
	case *models.ATTRIBUTEVALUEINTEGER:
		fillUpTable[models.ATTRIBUTEVALUEINTEGER](probe)
	case *models.ATTRIBUTEVALUEREAL:
		fillUpTable[models.ATTRIBUTEVALUEREAL](probe)
	case *models.ATTRIBUTEVALUESTRING:
		fillUpTable[models.ATTRIBUTEVALUESTRING](probe)
	case *models.ATTRIBUTEVALUEXHTML:
		fillUpTable[models.ATTRIBUTEVALUEXHTML](probe)
	case *models.CHILDREN:
		fillUpTable[models.CHILDREN](probe)
	case *models.CORECONTENT:
		fillUpTable[models.CORECONTENT](probe)
	case *models.DATATYPEDEFINITIONBOOLEAN:
		fillUpTable[models.DATATYPEDEFINITIONBOOLEAN](probe)
	case *models.DATATYPEDEFINITIONDATE:
		fillUpTable[models.DATATYPEDEFINITIONDATE](probe)
	case *models.DATATYPEDEFINITIONENUMERATION:
		fillUpTable[models.DATATYPEDEFINITIONENUMERATION](probe)
	case *models.DATATYPEDEFINITIONINTEGER:
		fillUpTable[models.DATATYPEDEFINITIONINTEGER](probe)
	case *models.DATATYPEDEFINITIONREAL:
		fillUpTable[models.DATATYPEDEFINITIONREAL](probe)
	case *models.DATATYPEDEFINITIONSTRING:
		fillUpTable[models.DATATYPEDEFINITIONSTRING](probe)
	case *models.DATATYPEDEFINITIONXHTML:
		fillUpTable[models.DATATYPEDEFINITIONXHTML](probe)
	case *models.DATATYPES:
		fillUpTable[models.DATATYPES](probe)
	case *models.DEFAULTVALUE:
		fillUpTable[models.DEFAULTVALUE](probe)
	case *models.DEFINITION:
		fillUpTable[models.DEFINITION](probe)
	case *models.EDITABLEATTS:
		fillUpTable[models.EDITABLEATTS](probe)
	case *models.EMBEDDEDVALUE:
		fillUpTable[models.EMBEDDEDVALUE](probe)
	case *models.ENUMVALUE:
		fillUpTable[models.ENUMVALUE](probe)
	case *models.OBJECT:
		fillUpTable[models.OBJECT](probe)
	case *models.PROPERTIES:
		fillUpTable[models.PROPERTIES](probe)
	case *models.RELATIONGROUP:
		fillUpTable[models.RELATIONGROUP](probe)
	case *models.RELATIONGROUPTYPE:
		fillUpTable[models.RELATIONGROUPTYPE](probe)
	case *models.REQIF:
		fillUpTable[models.REQIF](probe)
	case *models.REQIFCONTENT:
		fillUpTable[models.REQIFCONTENT](probe)
	case *models.REQIFHEADER:
		fillUpTable[models.REQIFHEADER](probe)
	case *models.REQIFTOOLEXTENSION:
		fillUpTable[models.REQIFTOOLEXTENSION](probe)
	case *models.REQIFTYPE:
		fillUpTable[models.REQIFTYPE](probe)
	case *models.SOURCE:
		fillUpTable[models.SOURCE](probe)
	case *models.SOURCESPECIFICATION:
		fillUpTable[models.SOURCESPECIFICATION](probe)
	case *models.SPECATTRIBUTES:
		fillUpTable[models.SPECATTRIBUTES](probe)
	case *models.SPECHIERARCHY:
		fillUpTable[models.SPECHIERARCHY](probe)
	case *models.SPECIFICATION:
		fillUpTable[models.SPECIFICATION](probe)
	case *models.SPECIFICATIONS:
		fillUpTable[models.SPECIFICATIONS](probe)
	case *models.SPECIFICATIONTYPE:
		fillUpTable[models.SPECIFICATIONTYPE](probe)
	case *models.SPECIFIEDVALUES:
		fillUpTable[models.SPECIFIEDVALUES](probe)
	case *models.SPECOBJECT:
		fillUpTable[models.SPECOBJECT](probe)
	case *models.SPECOBJECTS:
		fillUpTable[models.SPECOBJECTS](probe)
	case *models.SPECOBJECTTYPE:
		fillUpTable[models.SPECOBJECTTYPE](probe)
	case *models.SPECRELATION:
		fillUpTable[models.SPECRELATION](probe)
	case *models.SPECRELATIONGROUPS:
		fillUpTable[models.SPECRELATIONGROUPS](probe)
	case *models.SPECRELATIONS:
		fillUpTable[models.SPECRELATIONS](probe)
	case *models.SPECRELATIONTYPE:
		fillUpTable[models.SPECRELATIONTYPE](probe)
	case *models.SPECTYPES:
		fillUpTable[models.SPECTYPES](probe)
	case *models.TARGET:
		fillUpTable[models.TARGET](probe)
	case *models.TARGETSPECIFICATION:
		fillUpTable[models.TARGETSPECIFICATION](probe)
	case *models.THEHEADER:
		fillUpTable[models.THEHEADER](probe)
	case *models.TOOLEXTENSIONS:
		fillUpTable[models.TOOLEXTENSIONS](probe)
	case *models.VALUES:
		fillUpTable[models.VALUES](probe)
	case *models.XHTMLCONTENT:
		fillUpTable[models.XHTMLCONTENT](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()
	probe.tableStage.Commit()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			probe.stageOfInterest,
			probe.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
