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

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.REQIF": &(ref_models.REQIF{}),

	"ref_models.REQIF.Name": (ref_models.REQIF{}).Name,

	"ref_models.REQIF.REQIFHEADER": (ref_models.REQIF{}).HEADER,

	"ref_models.REQIFHEADER": &(ref_models.HEADER{}),

	"ref_models.REQIFHEADER.COMMENT": (ref_models.HEADER{}).COMMENT,

	"ref_models.REQIFHEADER.CREATIONTIME": (ref_models.HEADER{}).CREATIONTIME,

	"ref_models.REQIFHEADER.IDENTIFIERAttr": (ref_models.HEADER{}).IDENTIFIERAttr,

	"ref_models.REQIFHEADER.Name": (ref_models.HEADER{}).Name,

	"ref_models.REQIFHEADER.REPOSITORYID": (ref_models.HEADER{}).REPOSITORYID,

	"ref_models.REQIFHEADER.REQIFTOOLID": (ref_models.HEADER{}).REQIFTOOLID,

	"ref_models.REQIFHEADER.REQIFVERSION": (ref_models.HEADER{}).REQIFVERSION,

	"ref_models.REQIFHEADER.SOURCETOOLID": (ref_models.HEADER{}).SOURCETOOLID,

	"ref_models.REQIFHEADER.TITLE": (ref_models.HEADER{}).TITLE,
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__GongStructShape__000000_Default_REQIF := (&models.GongStructShape{Name: `Default-REQIF`}).Stage(stage)
	__GongStructShape__000001_Default_REQIFHEADER := (&models.GongStructShape{Name: `Default-REQIFHEADER`}).Stage(stage)

	__Link__000000_REQIFHEADER := (&models.Link{Name: `REQIFHEADER`}).Stage(stage)

	__Position__000000_Pos_Default_REQIF := (&models.Position{Name: `Pos-Default-REQIF`}).Stage(stage)
	__Position__000001_Pos_Default_REQIFHEADER := (&models.Position{Name: `Pos-Default-REQIFHEADER`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQIFHEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQIF and Default-REQIFHEADER`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__GongStructShape__000000_Default_REQIF.Name = `Default-REQIF`

	//gong:ident [ref_models.REQIF] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_REQIF.Identifier = `ref_models.REQIF`
	__GongStructShape__000000_Default_REQIF.ShowNbInstances = false
	__GongStructShape__000000_Default_REQIF.NbInstances = 0
	__GongStructShape__000000_Default_REQIF.Width = 240.000000
	__GongStructShape__000000_Default_REQIF.Height = 63.000000
	__GongStructShape__000000_Default_REQIF.IsSelected = false

	__GongStructShape__000001_Default_REQIFHEADER.Name = `Default-REQIFHEADER`

	//gong:ident [ref_models.REQIFHEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQIFHEADER.Identifier = `ref_models.REQIFHEADER`
	__GongStructShape__000001_Default_REQIFHEADER.ShowNbInstances = false
	__GongStructShape__000001_Default_REQIFHEADER.NbInstances = 0
	__GongStructShape__000001_Default_REQIFHEADER.Width = 240.000000
	__GongStructShape__000001_Default_REQIFHEADER.Height = 63.000000
	__GongStructShape__000001_Default_REQIFHEADER.IsSelected = false

	__Link__000000_REQIFHEADER.Name = `REQIFHEADER`

	//gong:ident [ref_models.REQIF.REQIFHEADER] comment added to overcome the problem with the comment map association
	__Link__000000_REQIFHEADER.Identifier = `ref_models.REQIF.REQIFHEADER`

	//gong:ident [ref_models.REQIFHEADER] comment added to overcome the problem with the comment map association
	__Link__000000_REQIFHEADER.Fieldtypename = `ref_models.REQIFHEADER`
	__Link__000000_REQIFHEADER.FieldOffsetX = -50.000000
	__Link__000000_REQIFHEADER.FieldOffsetY = -16.000000
	__Link__000000_REQIFHEADER.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_REQIFHEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_REQIFHEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_REQIFHEADER.SourceMultiplicity = models.MANY
	__Link__000000_REQIFHEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_REQIFHEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_REQIFHEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_REQIFHEADER.StartRatio = 0.500000
	__Link__000000_REQIFHEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_REQIFHEADER.EndRatio = 0.500000
	__Link__000000_REQIFHEADER.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_REQIF.X = 88.000000
	__Position__000000_Pos_Default_REQIF.Y = 90.000000
	__Position__000000_Pos_Default_REQIF.Name = `Pos-Default-REQIF`

	__Position__000001_Pos_Default_REQIFHEADER.X = 594.000000
	__Position__000001_Pos_Default_REQIFHEADER.Y = 86.000000
	__Position__000001_Pos_Default_REQIFHEADER.Name = `Pos-Default-REQIFHEADER`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQIFHEADER.X = 701.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQIFHEADER.Y = 119.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQIFHEADER.Name = `Verticle in class diagram Default in middle between Default-REQIF and Default-REQIFHEADER`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_REQIF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_REQIFHEADER)
	__GongStructShape__000000_Default_REQIF.Position = __Position__000000_Pos_Default_REQIF
	__GongStructShape__000000_Default_REQIF.Links = append(__GongStructShape__000000_Default_REQIF.Links, __Link__000000_REQIFHEADER)
	__GongStructShape__000001_Default_REQIFHEADER.Position = __Position__000001_Pos_Default_REQIFHEADER
	__Link__000000_REQIFHEADER.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQIF_and_Default_REQIFHEADER
}
