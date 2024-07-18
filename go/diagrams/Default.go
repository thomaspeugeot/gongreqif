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

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Setup of pointers
}
