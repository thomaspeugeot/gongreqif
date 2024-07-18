package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/thomaspeugeot/gongreqif/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.REQIF": &(ref_models.REQIF{}),

	"ref_models.REQIF.Name": (ref_models.REQIF{}).Name,

	"ref_models.REQIF.REQ_IF_CONTENT": (ref_models.REQIF{}).REQ_IF_CONTENT,

	"ref_models.REQIF.REQ_IF_HEADER": (ref_models.REQIF{}).REQ_IF_HEADER,

	"ref_models.REQ_IF_CONTENT": &(ref_models.REQ_IF_CONTENT{}),

	"ref_models.REQ_IF_CONTENT.Name": (ref_models.REQ_IF_CONTENT{}).Name,

	"ref_models.REQ_IF_CONTENT.SPECIFICATIONS": (ref_models.REQ_IF_CONTENT{}).SPECIFICATIONS,

	"ref_models.REQ_IF_CONTENT.SPECIFICATION_TYPES": (ref_models.REQ_IF_CONTENT{}).SPECIFICATION_TYPES,

	"ref_models.REQ_IF_CONTENT.SPEC_OBJECT_TYPES": (ref_models.REQ_IF_CONTENT{}).SPEC_OBJECT_TYPES,

	"ref_models.REQ_IF_HEADER": &(ref_models.REQ_IF_HEADER{}),

	"ref_models.REQ_IF_HEADER.COMMENT": (ref_models.REQ_IF_HEADER{}).COMMENT,

	"ref_models.REQ_IF_HEADER.CREATION_TIME": (ref_models.REQ_IF_HEADER{}).CREATION_TIME,

	"ref_models.REQ_IF_HEADER.Name": (ref_models.REQ_IF_HEADER{}).Name,

	"ref_models.REQ_IF_HEADER.REPOSITORY_ID": (ref_models.REQ_IF_HEADER{}).REPOSITORY_ID,

	"ref_models.REQ_IF_HEADER.REQ_IF_TOOL_ID": (ref_models.REQ_IF_HEADER{}).REQ_IF_TOOL_ID,

	"ref_models.REQ_IF_HEADER.REQ_IF_VERSION": (ref_models.REQ_IF_HEADER{}).REQ_IF_VERSION,

	"ref_models.REQ_IF_HEADER.SOURCE_TOOL_ID": (ref_models.REQ_IF_HEADER{}).SOURCE_TOOL_ID,

	"ref_models.REQ_IF_HEADER.TITLE": (ref_models.REQ_IF_HEADER{}).TITLE,

	"ref_models.SPECIFICATION": &(ref_models.SPECIFICATION{}),

	"ref_models.SPECIFICATION.CHILDREN": (ref_models.SPECIFICATION{}).CHILDREN,

	"ref_models.SPECIFICATION.DESC": (ref_models.SPECIFICATION{}).DESC,

	"ref_models.SPECIFICATION.IDENTIFIER": (ref_models.SPECIFICATION{}).IDENTIFIER,

	"ref_models.SPECIFICATION.LAST_CHANGE": (ref_models.SPECIFICATION{}).LAST_CHANGE,

	"ref_models.SPECIFICATION.LONG_NAME": (ref_models.SPECIFICATION{}).LONG_NAME,

	"ref_models.SPECIFICATION.Name": (ref_models.SPECIFICATION{}).Name,

	"ref_models.SPECIFICATION.SPECIFICATION_TYPE": (ref_models.SPECIFICATION{}).SPECIFICATION_TYPE,

	"ref_models.SPECIFICATION.TYPE": (ref_models.SPECIFICATION{}).TYPE,

	"ref_models.SPECIFICATION_TYPE": &(ref_models.SPECIFICATION_TYPE{}),

	"ref_models.SPECIFICATION_TYPE.DESC": (ref_models.SPECIFICATION_TYPE{}).DESC,

	"ref_models.SPECIFICATION_TYPE.IDENTIFIER": (ref_models.SPECIFICATION_TYPE{}).IDENTIFIER,

	"ref_models.SPECIFICATION_TYPE.LAST_CHANGE": (ref_models.SPECIFICATION_TYPE{}).LAST_CHANGE,

	"ref_models.SPECIFICATION_TYPE.LONG_NAME": (ref_models.SPECIFICATION_TYPE{}).LONG_NAME,

	"ref_models.SPECIFICATION_TYPE.Name": (ref_models.SPECIFICATION_TYPE{}).Name,

	"ref_models.SPEC_HIERARCHY": &(ref_models.SPEC_HIERARCHY{}),

	"ref_models.SPEC_HIERARCHY.CHILDREN": (ref_models.SPEC_HIERARCHY{}).CHILDREN,

	"ref_models.SPEC_HIERARCHY.DESC": (ref_models.SPEC_HIERARCHY{}).DESC,

	"ref_models.SPEC_HIERARCHY.IDENTIFIER": (ref_models.SPEC_HIERARCHY{}).IDENTIFIER,

	"ref_models.SPEC_HIERARCHY.IS_EDITABLE": (ref_models.SPEC_HIERARCHY{}).IS_EDITABLE,

	"ref_models.SPEC_HIERARCHY.IS_TABLE_INTERNAL": (ref_models.SPEC_HIERARCHY{}).IS_TABLE_INTERNAL,

	"ref_models.SPEC_HIERARCHY.LAST_CHANGE": (ref_models.SPEC_HIERARCHY{}).LAST_CHANGE,

	"ref_models.SPEC_HIERARCHY.LONG_NAME": (ref_models.SPEC_HIERARCHY{}).LONG_NAME,

	"ref_models.SPEC_HIERARCHY.Name": (ref_models.SPEC_HIERARCHY{}).Name,

	"ref_models.SPEC_HIERARCHY.OBJECT": (ref_models.SPEC_HIERARCHY{}).OBJECT,

	"ref_models.SPEC_OBJECT_TYPE": &(ref_models.SPEC_OBJECT_TYPE{}),

	"ref_models.SPEC_OBJECT_TYPE.DESC": (ref_models.SPEC_OBJECT_TYPE{}).DESC,

	"ref_models.SPEC_OBJECT_TYPE.IDENTIFIER": (ref_models.SPEC_OBJECT_TYPE{}).IDENTIFIER,

	"ref_models.SPEC_OBJECT_TYPE.LAST_CHANGE": (ref_models.SPEC_OBJECT_TYPE{}).LAST_CHANGE,

	"ref_models.SPEC_OBJECT_TYPE.LONG_NAME": (ref_models.SPEC_OBJECT_TYPE{}).LONG_NAME,

	"ref_models.SPEC_OBJECT_TYPE.Name": (ref_models.SPEC_OBJECT_TYPE{}).Name,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_TYPE := (&models.Field{Name: `TYPE`}).Stage(stage)

	__GongStructShape__000000_Default_REQIF := (&models.GongStructShape{Name: `Default-REQIF`}).Stage(stage)
	__GongStructShape__000001_Default_REQ_IF_CONTENT := (&models.GongStructShape{Name: `Default-REQ_IF_CONTENT`}).Stage(stage)
	__GongStructShape__000002_Default_REQ_IF_HEADER := (&models.GongStructShape{Name: `Default-REQ_IF_HEADER`}).Stage(stage)
	__GongStructShape__000003_Default_SPECIFICATION := (&models.GongStructShape{Name: `Default-SPECIFICATION`}).Stage(stage)
	__GongStructShape__000004_Default_SPECIFICATION_TYPE := (&models.GongStructShape{Name: `Default-SPECIFICATION_TYPE`}).Stage(stage)
	__GongStructShape__000005_Default_SPEC_HIERARCHY := (&models.GongStructShape{Name: `Default-SPEC_HIERARCHY`}).Stage(stage)
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE := (&models.GongStructShape{Name: `Default-SPEC_OBJECT_TYPE`}).Stage(stage)

	__Link__000000_CHILDREN := (&models.Link{Name: `CHILDREN`}).Stage(stage)
	__Link__000001_CHILDREN := (&models.Link{Name: `CHILDREN`}).Stage(stage)
	__Link__000002_REQ_IF_CONTENT := (&models.Link{Name: `REQ_IF_CONTENT`}).Stage(stage)
	__Link__000003_REQ_IF_HEADER := (&models.Link{Name: `REQ_IF_HEADER`}).Stage(stage)
	__Link__000004_SPECIFICATIONS := (&models.Link{Name: `SPECIFICATIONS`}).Stage(stage)
	__Link__000005_SPECIFICATION_TYPE := (&models.Link{Name: `SPECIFICATION_TYPE`}).Stage(stage)
	__Link__000006_SPECIFICATION_TYPES := (&models.Link{Name: `SPECIFICATION_TYPES`}).Stage(stage)
	__Link__000007_SPEC_OBJECT_TYPES := (&models.Link{Name: `SPEC_OBJECT_TYPES`}).Stage(stage)

	__Position__000000_Pos_Default_REQIF := (&models.Position{Name: `Pos-Default-REQIF`}).Stage(stage)
	__Position__000001_Pos_Default_REQ_IF_CONTENT := (&models.Position{Name: `Pos-Default-REQ_IF_CONTENT`}).Stage(stage)
	__Position__000002_Pos_Default_REQ_IF_HEADER := (&models.Position{Name: `Pos-Default-REQ_IF_HEADER`}).Stage(stage)
	__Position__000003_Pos_Default_SPECIFICATION := (&models.Position{Name: `Pos-Default-SPECIFICATION`}).Stage(stage)
	__Position__000004_Pos_Default_SPECIFICATION_TYPE := (&models.Position{Name: `Pos-Default-SPECIFICATION_TYPE`}).Stage(stage)
	__Position__000005_Pos_Default_SPEC_HIERARCHY := (&models.Position{Name: `Pos-Default-SPEC_HIERARCHY`}).Stage(stage)
	__Position__000006_Pos_Default_SPEC_OBJECT_TYPE := (&models.Position{Name: `Pos-Default-SPEC_OBJECT_TYPE`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQIF and Default-REQ_IF_CONTENT`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQIF and Default-REQ_IF_HEADER`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPECIFICATION`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION_TYPE := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPECIFICATION_TYPE`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPEC_OBJECT_TYPE := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPEC_OBJECT_TYPE`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPECIFICATION_TYPE := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SPECIFICATION and Default-SPECIFICATION_TYPE`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPEC_HIERARCHY := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SPECIFICATION and Default-SPEC_HIERARCHY`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_HIERARCHY_and_Default_SPEC_HIERARCHY := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SPEC_HIERARCHY and Default-SPEC_HIERARCHY`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_TYPE.Name = `TYPE`

	//gong:ident [ref_models.SPECIFICATION.TYPE] comment added to overcome the problem with the comment map association
	__Field__000000_TYPE.Identifier = `ref_models.SPECIFICATION.TYPE`
	__Field__000000_TYPE.FieldTypeAsString = ``
	__Field__000000_TYPE.Structname = `SPECIFICATION`
	__Field__000000_TYPE.Fieldtypename = `string`

	__GongStructShape__000000_Default_REQIF.Name = `Default-REQIF`

	//gong:ident [ref_models.REQIF] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_REQIF.Identifier = `ref_models.REQIF`
	__GongStructShape__000000_Default_REQIF.ShowNbInstances = false
	__GongStructShape__000000_Default_REQIF.NbInstances = 0
	__GongStructShape__000000_Default_REQIF.Width = 240.000000
	__GongStructShape__000000_Default_REQIF.Height = 63.000000
	__GongStructShape__000000_Default_REQIF.IsSelected = false

	__GongStructShape__000001_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000001_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000002_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000002_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000002_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Height = 63.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.IsSelected = false

	__GongStructShape__000003_Default_SPECIFICATION.Name = `Default-SPECIFICATION`

	//gong:ident [ref_models.SPECIFICATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_SPECIFICATION.Identifier = `ref_models.SPECIFICATION`
	__GongStructShape__000003_Default_SPECIFICATION.ShowNbInstances = false
	__GongStructShape__000003_Default_SPECIFICATION.NbInstances = 0
	__GongStructShape__000003_Default_SPECIFICATION.Width = 240.000000
	__GongStructShape__000003_Default_SPECIFICATION.Height = 151.000000
	__GongStructShape__000003_Default_SPECIFICATION.IsSelected = false

	__GongStructShape__000004_Default_SPECIFICATION_TYPE.Name = `Default-SPECIFICATION_TYPE`

	//gong:ident [ref_models.SPECIFICATION_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.Identifier = `ref_models.SPECIFICATION_TYPE`
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.ShowNbInstances = false
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.NbInstances = 0
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.Width = 240.000000
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.Height = 63.000000
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.IsSelected = false

	__GongStructShape__000005_Default_SPEC_HIERARCHY.Name = `Default-SPEC_HIERARCHY`

	//gong:ident [ref_models.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_SPEC_HIERARCHY.Identifier = `ref_models.SPEC_HIERARCHY`
	__GongStructShape__000005_Default_SPEC_HIERARCHY.ShowNbInstances = false
	__GongStructShape__000005_Default_SPEC_HIERARCHY.NbInstances = 0
	__GongStructShape__000005_Default_SPEC_HIERARCHY.Width = 240.000000
	__GongStructShape__000005_Default_SPEC_HIERARCHY.Height = 63.000000
	__GongStructShape__000005_Default_SPEC_HIERARCHY.IsSelected = false

	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.Name = `Default-SPEC_OBJECT_TYPE`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.Identifier = `ref_models.SPEC_OBJECT_TYPE`
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.ShowNbInstances = false
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.NbInstances = 0
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.Width = 240.000000
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.Height = 63.000000
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.IsSelected = false

	__Link__000000_CHILDREN.Name = `CHILDREN`

	//gong:ident [ref_models.SPEC_HIERARCHY.CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000000_CHILDREN.Identifier = `ref_models.SPEC_HIERARCHY.CHILDREN`

	//gong:ident [ref_models.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__Link__000000_CHILDREN.Fieldtypename = `ref_models.SPEC_HIERARCHY`
	__Link__000000_CHILDREN.FieldOffsetX = -50.000000
	__Link__000000_CHILDREN.FieldOffsetY = -16.000000
	__Link__000000_CHILDREN.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_CHILDREN.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_CHILDREN.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_CHILDREN.SourceMultiplicity = models.MANY
	__Link__000000_CHILDREN.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_CHILDREN.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_CHILDREN.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_CHILDREN.StartRatio = 0.130139
	__Link__000000_CHILDREN.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_CHILDREN.EndRatio = 0.888472
	__Link__000000_CHILDREN.CornerOffsetRatio = 3.053172

	__Link__000001_CHILDREN.Name = `CHILDREN`

	//gong:ident [ref_models.SPECIFICATION.CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000001_CHILDREN.Identifier = `ref_models.SPECIFICATION.CHILDREN`

	//gong:ident [ref_models.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__Link__000001_CHILDREN.Fieldtypename = `ref_models.SPEC_HIERARCHY`
	__Link__000001_CHILDREN.FieldOffsetX = -50.000000
	__Link__000001_CHILDREN.FieldOffsetY = -16.000000
	__Link__000001_CHILDREN.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_CHILDREN.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_CHILDREN.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_CHILDREN.SourceMultiplicity = models.MANY
	__Link__000001_CHILDREN.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_CHILDREN.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_CHILDREN.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CHILDREN.StartRatio = 0.862696
	__Link__000001_CHILDREN.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CHILDREN.EndRatio = 0.500000
	__Link__000001_CHILDREN.CornerOffsetRatio = 1.646806

	__Link__000002_REQ_IF_CONTENT.Name = `REQ_IF_CONTENT`

	//gong:ident [ref_models.REQIF.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000002_REQ_IF_CONTENT.Identifier = `ref_models.REQIF.REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000002_REQ_IF_CONTENT.Fieldtypename = `ref_models.REQ_IF_CONTENT`
	__Link__000002_REQ_IF_CONTENT.FieldOffsetX = -50.000000
	__Link__000002_REQ_IF_CONTENT.FieldOffsetY = -16.000000
	__Link__000002_REQ_IF_CONTENT.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_REQ_IF_CONTENT.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_REQ_IF_CONTENT.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_REQ_IF_CONTENT.SourceMultiplicity = models.MANY
	__Link__000002_REQ_IF_CONTENT.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_REQ_IF_CONTENT.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_REQ_IF_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_REQ_IF_CONTENT.StartRatio = 0.500000
	__Link__000002_REQ_IF_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_REQ_IF_CONTENT.EndRatio = 0.500000
	__Link__000002_REQ_IF_CONTENT.CornerOffsetRatio = 1.380000

	__Link__000003_REQ_IF_HEADER.Name = `REQ_IF_HEADER`

	//gong:ident [ref_models.REQIF.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000003_REQ_IF_HEADER.Identifier = `ref_models.REQIF.REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000003_REQ_IF_HEADER.Fieldtypename = `ref_models.REQ_IF_HEADER`
	__Link__000003_REQ_IF_HEADER.FieldOffsetX = -50.000000
	__Link__000003_REQ_IF_HEADER.FieldOffsetY = -16.000000
	__Link__000003_REQ_IF_HEADER.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_REQ_IF_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_REQ_IF_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_REQ_IF_HEADER.SourceMultiplicity = models.MANY
	__Link__000003_REQ_IF_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_REQ_IF_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_REQ_IF_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_REQ_IF_HEADER.StartRatio = 0.500000
	__Link__000003_REQ_IF_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_REQ_IF_HEADER.EndRatio = 0.500000
	__Link__000003_REQ_IF_HEADER.CornerOffsetRatio = 1.380000

	__Link__000004_SPECIFICATIONS.Name = `SPECIFICATIONS`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPECIFICATIONS] comment added to overcome the problem with the comment map association
	__Link__000004_SPECIFICATIONS.Identifier = `ref_models.REQ_IF_CONTENT.SPECIFICATIONS`

	//gong:ident [ref_models.SPECIFICATION] comment added to overcome the problem with the comment map association
	__Link__000004_SPECIFICATIONS.Fieldtypename = `ref_models.SPECIFICATION`
	__Link__000004_SPECIFICATIONS.FieldOffsetX = -50.000000
	__Link__000004_SPECIFICATIONS.FieldOffsetY = -16.000000
	__Link__000004_SPECIFICATIONS.TargetMultiplicity = models.MANY
	__Link__000004_SPECIFICATIONS.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_SPECIFICATIONS.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_SPECIFICATIONS.SourceMultiplicity = models.MANY
	__Link__000004_SPECIFICATIONS.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_SPECIFICATIONS.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_SPECIFICATIONS.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_SPECIFICATIONS.StartRatio = 0.221806
	__Link__000004_SPECIFICATIONS.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_SPECIFICATIONS.EndRatio = 0.234105
	__Link__000004_SPECIFICATIONS.CornerOffsetRatio = 1.053172

	__Link__000005_SPECIFICATION_TYPE.Name = `SPECIFICATION_TYPE`

	//gong:ident [ref_models.SPECIFICATION.SPECIFICATION_TYPE] comment added to overcome the problem with the comment map association
	__Link__000005_SPECIFICATION_TYPE.Identifier = `ref_models.SPECIFICATION.SPECIFICATION_TYPE`

	//gong:ident [ref_models.SPECIFICATION_TYPE] comment added to overcome the problem with the comment map association
	__Link__000005_SPECIFICATION_TYPE.Fieldtypename = `ref_models.SPECIFICATION_TYPE`
	__Link__000005_SPECIFICATION_TYPE.FieldOffsetX = -50.000000
	__Link__000005_SPECIFICATION_TYPE.FieldOffsetY = -16.000000
	__Link__000005_SPECIFICATION_TYPE.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_SPECIFICATION_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_SPECIFICATION_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_SPECIFICATION_TYPE.SourceMultiplicity = models.MANY
	__Link__000005_SPECIFICATION_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_SPECIFICATION_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_SPECIFICATION_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_SPECIFICATION_TYPE.StartRatio = 0.500000
	__Link__000005_SPECIFICATION_TYPE.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_SPECIFICATION_TYPE.EndRatio = 0.188472
	__Link__000005_SPECIFICATION_TYPE.CornerOffsetRatio = -1.063908

	__Link__000006_SPECIFICATION_TYPES.Name = `SPECIFICATION_TYPES`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPECIFICATION_TYPES] comment added to overcome the problem with the comment map association
	__Link__000006_SPECIFICATION_TYPES.Identifier = `ref_models.REQ_IF_CONTENT.SPECIFICATION_TYPES`

	//gong:ident [ref_models.SPECIFICATION_TYPE] comment added to overcome the problem with the comment map association
	__Link__000006_SPECIFICATION_TYPES.Fieldtypename = `ref_models.SPECIFICATION_TYPE`
	__Link__000006_SPECIFICATION_TYPES.FieldOffsetX = -50.000000
	__Link__000006_SPECIFICATION_TYPES.FieldOffsetY = -16.000000
	__Link__000006_SPECIFICATION_TYPES.TargetMultiplicity = models.MANY
	__Link__000006_SPECIFICATION_TYPES.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_SPECIFICATION_TYPES.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_SPECIFICATION_TYPES.SourceMultiplicity = models.MANY
	__Link__000006_SPECIFICATION_TYPES.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_SPECIFICATION_TYPES.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_SPECIFICATION_TYPES.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000006_SPECIFICATION_TYPES.StartRatio = 0.834306
	__Link__000006_SPECIFICATION_TYPES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_SPECIFICATION_TYPES.EndRatio = 0.500000
	__Link__000006_SPECIFICATION_TYPES.CornerOffsetRatio = 1.005553

	__Link__000007_SPEC_OBJECT_TYPES.Name = `SPEC_OBJECT_TYPES`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPEC_OBJECT_TYPES] comment added to overcome the problem with the comment map association
	__Link__000007_SPEC_OBJECT_TYPES.Identifier = `ref_models.REQ_IF_CONTENT.SPEC_OBJECT_TYPES`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__Link__000007_SPEC_OBJECT_TYPES.Fieldtypename = `ref_models.SPEC_OBJECT_TYPE`
	__Link__000007_SPEC_OBJECT_TYPES.FieldOffsetX = -50.000000
	__Link__000007_SPEC_OBJECT_TYPES.FieldOffsetY = -16.000000
	__Link__000007_SPEC_OBJECT_TYPES.TargetMultiplicity = models.MANY
	__Link__000007_SPEC_OBJECT_TYPES.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_SPEC_OBJECT_TYPES.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_SPEC_OBJECT_TYPES.SourceMultiplicity = models.MANY
	__Link__000007_SPEC_OBJECT_TYPES.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_SPEC_OBJECT_TYPES.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_SPEC_OBJECT_TYPES.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_SPEC_OBJECT_TYPES.StartRatio = 0.055139
	__Link__000007_SPEC_OBJECT_TYPES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_SPEC_OBJECT_TYPES.EndRatio = 0.500000
	__Link__000007_SPEC_OBJECT_TYPES.CornerOffsetRatio = 1.069045

	__Position__000000_Pos_Default_REQIF.X = 44.000000
	__Position__000000_Pos_Default_REQIF.Y = 28.000000
	__Position__000000_Pos_Default_REQIF.Name = `Pos-Default-REQIF`

	__Position__000001_Pos_Default_REQ_IF_CONTENT.X = 673.000000
	__Position__000001_Pos_Default_REQ_IF_CONTENT.Y = 31.000000
	__Position__000001_Pos_Default_REQ_IF_CONTENT.Name = `Pos-Default-REQ_IF_CONTENT`

	__Position__000002_Pos_Default_REQ_IF_HEADER.X = 16.000000
	__Position__000002_Pos_Default_REQ_IF_HEADER.Y = 115.000000
	__Position__000002_Pos_Default_REQ_IF_HEADER.Name = `Pos-Default-REQ_IF_HEADER`

	__Position__000003_Pos_Default_SPECIFICATION.X = 38.000000
	__Position__000003_Pos_Default_SPECIFICATION.Y = 450.000000
	__Position__000003_Pos_Default_SPECIFICATION.Name = `Pos-Default-SPECIFICATION`

	__Position__000004_Pos_Default_SPECIFICATION_TYPE.X = 1111.000061
	__Position__000004_Pos_Default_SPECIFICATION_TYPE.Y = 243.000000
	__Position__000004_Pos_Default_SPECIFICATION_TYPE.Name = `Pos-Default-SPECIFICATION_TYPE`

	__Position__000005_Pos_Default_SPEC_HIERARCHY.X = 37.000000
	__Position__000005_Pos_Default_SPEC_HIERARCHY.Y = 682.000000
	__Position__000005_Pos_Default_SPEC_HIERARCHY.Name = `Pos-Default-SPEC_HIERARCHY`

	__Position__000006_Pos_Default_SPEC_OBJECT_TYPE.X = 37.000000
	__Position__000006_Pos_Default_SPEC_OBJECT_TYPE.Y = 338.000000
	__Position__000006_Pos_Default_SPEC_OBJECT_TYPE.Name = `Pos-Default-SPEC_OBJECT_TYPE`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_CONTENT.X = 648.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_CONTENT.Y = 74.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_CONTENT.Name = `Verticle in class diagram Default in middle between Default-REQIF and Default-REQ_IF_CONTENT`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_HEADER.X = 420.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_HEADER.Y = 127.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_HEADER.Name = `Verticle in class diagram Default in middle between Default-REQIF and Default-REQ_IF_HEADER`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION.X = 859.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION.Y = 230.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPECIFICATION`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION_TYPE.X = 715.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION_TYPE.Y = 170.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION_TYPE.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPECIFICATION_TYPE`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPEC_OBJECT_TYPE.X = 1002.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPEC_OBJECT_TYPE.Y = 191.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPEC_OBJECT_TYPE.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-SPEC_OBJECT_TYPE`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPECIFICATION_TYPE.X = 974.000030
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPECIFICATION_TYPE.Y = 413.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPECIFICATION_TYPE.Name = `Verticle in class diagram Default in middle between Default-SPECIFICATION and Default-SPECIFICATION_TYPE`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPEC_HIERARCHY.X = 402.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPEC_HIERARCHY.Y = 562.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPEC_HIERARCHY.Name = `Verticle in class diagram Default in middle between Default-SPECIFICATION and Default-SPEC_HIERARCHY`

	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_HIERARCHY_and_Default_SPEC_HIERARCHY.X = 406.000000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_HIERARCHY_and_Default_SPEC_HIERARCHY.Y = 642.500000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_HIERARCHY_and_Default_SPEC_HIERARCHY.Name = `Verticle in class diagram Default in middle between Default-SPEC_HIERARCHY and Default-SPEC_HIERARCHY`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_REQIF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_REQ_IF_HEADER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_SPECIFICATION)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_SPECIFICATION_TYPE)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_SPEC_HIERARCHY)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_SPEC_OBJECT_TYPE)
	__GongStructShape__000000_Default_REQIF.Position = __Position__000000_Pos_Default_REQIF
	__GongStructShape__000000_Default_REQIF.Links = append(__GongStructShape__000000_Default_REQIF.Links, __Link__000003_REQ_IF_HEADER)
	__GongStructShape__000000_Default_REQIF.Links = append(__GongStructShape__000000_Default_REQIF.Links, __Link__000002_REQ_IF_CONTENT)
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Position = __Position__000001_Pos_Default_REQ_IF_CONTENT
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000001_Default_REQ_IF_CONTENT.Links, __Link__000007_SPEC_OBJECT_TYPES)
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000001_Default_REQ_IF_CONTENT.Links, __Link__000006_SPECIFICATION_TYPES)
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000001_Default_REQ_IF_CONTENT.Links, __Link__000004_SPECIFICATIONS)
	__GongStructShape__000002_Default_REQ_IF_HEADER.Position = __Position__000002_Pos_Default_REQ_IF_HEADER
	__GongStructShape__000003_Default_SPECIFICATION.Position = __Position__000003_Pos_Default_SPECIFICATION
	__GongStructShape__000003_Default_SPECIFICATION.Fields = append(__GongStructShape__000003_Default_SPECIFICATION.Fields, __Field__000000_TYPE)
	__GongStructShape__000003_Default_SPECIFICATION.Links = append(__GongStructShape__000003_Default_SPECIFICATION.Links, __Link__000001_CHILDREN)
	__GongStructShape__000003_Default_SPECIFICATION.Links = append(__GongStructShape__000003_Default_SPECIFICATION.Links, __Link__000005_SPECIFICATION_TYPE)
	__GongStructShape__000004_Default_SPECIFICATION_TYPE.Position = __Position__000004_Pos_Default_SPECIFICATION_TYPE
	__GongStructShape__000005_Default_SPEC_HIERARCHY.Position = __Position__000005_Pos_Default_SPEC_HIERARCHY
	__GongStructShape__000005_Default_SPEC_HIERARCHY.Links = append(__GongStructShape__000005_Default_SPEC_HIERARCHY.Links, __Link__000000_CHILDREN)
	__GongStructShape__000006_Default_SPEC_OBJECT_TYPE.Position = __Position__000006_Pos_Default_SPEC_OBJECT_TYPE
	__Link__000000_CHILDREN.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_HIERARCHY_and_Default_SPEC_HIERARCHY
	__Link__000001_CHILDREN.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPEC_HIERARCHY
	__Link__000002_REQ_IF_CONTENT.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_CONTENT
	__Link__000003_REQ_IF_HEADER.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQ_IF_HEADER
	__Link__000004_SPECIFICATIONS.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION
	__Link__000005_SPECIFICATION_TYPE.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_SPECIFICATION_and_Default_SPECIFICATION_TYPE
	__Link__000006_SPECIFICATION_TYPES.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPECIFICATION_TYPE
	__Link__000007_SPEC_OBJECT_TYPES.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_SPEC_OBJECT_TYPE
}
