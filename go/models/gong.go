// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	CONTENTs           map[*CONTENT]any
	CONTENTs_mapString map[string]*CONTENT

	// insertion point for slice of pointers maps

	OnAfterCONTENTCreateCallback OnAfterCreateInterface[CONTENT]
	OnAfterCONTENTUpdateCallback OnAfterUpdateInterface[CONTENT]
	OnAfterCONTENTDeleteCallback OnAfterDeleteInterface[CONTENT]
	OnAfterCONTENTReadCallback   OnAfterReadInterface[CONTENT]

	HEADERs           map[*HEADER]any
	HEADERs_mapString map[string]*HEADER

	// insertion point for slice of pointers maps

	OnAfterHEADERCreateCallback OnAfterCreateInterface[HEADER]
	OnAfterHEADERUpdateCallback OnAfterUpdateInterface[HEADER]
	OnAfterHEADERDeleteCallback OnAfterDeleteInterface[HEADER]
	OnAfterHEADERReadCallback   OnAfterReadInterface[HEADER]

	REQIFs           map[*REQIF]any
	REQIFs_mapString map[string]*REQIF

	// insertion point for slice of pointers maps

	OnAfterREQIFCreateCallback OnAfterCreateInterface[REQIF]
	OnAfterREQIFUpdateCallback OnAfterUpdateInterface[REQIF]
	OnAfterREQIFDeleteCallback OnAfterDeleteInterface[REQIF]
	OnAfterREQIFReadCallback   OnAfterReadInterface[REQIF]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/thomaspeugeot/gongreqif/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCONTENT(content *CONTENT)
	CheckoutCONTENT(content *CONTENT)
	CommitHEADER(header *HEADER)
	CheckoutHEADER(header *HEADER)
	CommitREQIF(reqif *REQIF)
	CheckoutREQIF(reqif *REQIF)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		CONTENTs:           make(map[*CONTENT]any),
		CONTENTs_mapString: make(map[string]*CONTENT),

		HEADERs:           make(map[*HEADER]any),
		HEADERs_mapString: make(map[string]*HEADER),

		REQIFs:           make(map[*REQIF]any),
		REQIFs_mapString: make(map[string]*REQIF),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CONTENT"] = len(stage.CONTENTs)
	stage.Map_GongStructName_InstancesNb["HEADER"] = len(stage.HEADERs)
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CONTENT"] = len(stage.CONTENTs)
	stage.Map_GongStructName_InstancesNb["HEADER"] = len(stage.HEADERs)
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts content to the model stage
func (content *CONTENT) Stage(stage *StageStruct) *CONTENT {
	stage.CONTENTs[content] = __member
	stage.CONTENTs_mapString[content.Name] = content

	return content
}

// Unstage removes content off the model stage
func (content *CONTENT) Unstage(stage *StageStruct) *CONTENT {
	delete(stage.CONTENTs, content)
	delete(stage.CONTENTs_mapString, content.Name)
	return content
}

// UnstageVoid removes content off the model stage
func (content *CONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.CONTENTs, content)
	delete(stage.CONTENTs_mapString, content.Name)
}

// commit content to the back repo (if it is already staged)
func (content *CONTENT) Commit(stage *StageStruct) *CONTENT {
	if _, ok := stage.CONTENTs[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCONTENT(content)
		}
	}
	return content
}

func (content *CONTENT) CommitVoid(stage *StageStruct) {
	content.Commit(stage)
}

// Checkout content to the back repo (if it is already staged)
func (content *CONTENT) Checkout(stage *StageStruct) *CONTENT {
	if _, ok := stage.CONTENTs[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCONTENT(content)
		}
	}
	return content
}

// for satisfaction of GongStruct interface
func (content *CONTENT) GetName() (res string) {
	return content.Name
}

// Stage puts header to the model stage
func (header *HEADER) Stage(stage *StageStruct) *HEADER {
	stage.HEADERs[header] = __member
	stage.HEADERs_mapString[header.Name] = header

	return header
}

// Unstage removes header off the model stage
func (header *HEADER) Unstage(stage *StageStruct) *HEADER {
	delete(stage.HEADERs, header)
	delete(stage.HEADERs_mapString, header.Name)
	return header
}

// UnstageVoid removes header off the model stage
func (header *HEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.HEADERs, header)
	delete(stage.HEADERs_mapString, header.Name)
}

// commit header to the back repo (if it is already staged)
func (header *HEADER) Commit(stage *StageStruct) *HEADER {
	if _, ok := stage.HEADERs[header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitHEADER(header)
		}
	}
	return header
}

func (header *HEADER) CommitVoid(stage *StageStruct) {
	header.Commit(stage)
}

// Checkout header to the back repo (if it is already staged)
func (header *HEADER) Checkout(stage *StageStruct) *HEADER {
	if _, ok := stage.HEADERs[header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutHEADER(header)
		}
	}
	return header
}

// for satisfaction of GongStruct interface
func (header *HEADER) GetName() (res string) {
	return header.Name
}

// Stage puts reqif to the model stage
func (reqif *REQIF) Stage(stage *StageStruct) *REQIF {
	stage.REQIFs[reqif] = __member
	stage.REQIFs_mapString[reqif.Name] = reqif

	return reqif
}

// Unstage removes reqif off the model stage
func (reqif *REQIF) Unstage(stage *StageStruct) *REQIF {
	delete(stage.REQIFs, reqif)
	delete(stage.REQIFs_mapString, reqif.Name)
	return reqif
}

// UnstageVoid removes reqif off the model stage
func (reqif *REQIF) UnstageVoid(stage *StageStruct) {
	delete(stage.REQIFs, reqif)
	delete(stage.REQIFs_mapString, reqif.Name)
}

// commit reqif to the back repo (if it is already staged)
func (reqif *REQIF) Commit(stage *StageStruct) *REQIF {
	if _, ok := stage.REQIFs[reqif]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQIF(reqif)
		}
	}
	return reqif
}

func (reqif *REQIF) CommitVoid(stage *StageStruct) {
	reqif.Commit(stage)
}

// Checkout reqif to the back repo (if it is already staged)
func (reqif *REQIF) Checkout(stage *StageStruct) *REQIF {
	if _, ok := stage.REQIFs[reqif]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQIF(reqif)
		}
	}
	return reqif
}

// for satisfaction of GongStruct interface
func (reqif *REQIF) GetName() (res string) {
	return reqif.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCONTENT(CONTENT *CONTENT)
	CreateORMHEADER(HEADER *HEADER)
	CreateORMREQIF(REQIF *REQIF)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCONTENT(CONTENT *CONTENT)
	DeleteORMHEADER(HEADER *HEADER)
	DeleteORMREQIF(REQIF *REQIF)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.CONTENTs = make(map[*CONTENT]any)
	stage.CONTENTs_mapString = make(map[string]*CONTENT)

	stage.HEADERs = make(map[*HEADER]any)
	stage.HEADERs_mapString = make(map[string]*HEADER)

	stage.REQIFs = make(map[*REQIF]any)
	stage.REQIFs_mapString = make(map[string]*REQIF)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.CONTENTs = nil
	stage.CONTENTs_mapString = nil

	stage.HEADERs = nil
	stage.HEADERs_mapString = nil

	stage.REQIFs = nil
	stage.REQIFs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for content := range stage.CONTENTs {
		content.Unstage(stage)
	}

	for header := range stage.HEADERs {
		header.Unstage(stage)
	}

	for reqif := range stage.REQIFs {
		reqif.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	CONTENT | HEADER | REQIF
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*CONTENT | *HEADER | *REQIF
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*CONTENT]any |
		map[*HEADER]any |
		map[*REQIF]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*CONTENT |
		map[string]*HEADER |
		map[string]*REQIF |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*CONTENT]any:
		return any(&stage.CONTENTs).(*Type)
	case map[*HEADER]any:
		return any(&stage.HEADERs).(*Type)
	case map[*REQIF]any:
		return any(&stage.REQIFs).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*CONTENT:
		return any(&stage.CONTENTs_mapString).(*Type)
	case map[string]*HEADER:
		return any(&stage.HEADERs_mapString).(*Type)
	case map[string]*REQIF:
		return any(&stage.REQIFs_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CONTENT:
		return any(&stage.CONTENTs).(*map[*Type]any)
	case HEADER:
		return any(&stage.HEADERs).(*map[*Type]any)
	case REQIF:
		return any(&stage.REQIFs).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *CONTENT:
		return any(&stage.CONTENTs).(*map[Type]any)
	case *HEADER:
		return any(&stage.HEADERs).(*map[Type]any)
	case *REQIF:
		return any(&stage.REQIFs).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CONTENT:
		return any(&stage.CONTENTs_mapString).(*map[string]*Type)
	case HEADER:
		return any(&stage.HEADERs_mapString).(*map[string]*Type)
	case REQIF:
		return any(&stage.REQIFs_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case CONTENT:
		return any(&CONTENT{
			// Initialisation of associations
		}).(*Type)
	case HEADER:
		return any(&HEADER{
			// Initialisation of associations
		}).(*Type)
	case REQIF:
		return any(&REQIF{
			// Initialisation of associations
			// field is initialized with an instance of HEADER with the name of the field
			HEADER: &HEADER{Name: "HEADER"},
			// field is initialized with an instance of CONTENT with the name of the field
			CONTENT: &CONTENT{Name: "CONTENT"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CONTENT
	case CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HEADER
	case HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		case "HEADER":
			res := make(map[*HEADER][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.HEADER != nil {
					header_ := reqif.HEADER
					var reqifs []*REQIF
					_, ok := res[header_]
					if ok {
						reqifs = res[header_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[header_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "CONTENT":
			res := make(map[*CONTENT][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.CONTENT != nil {
					content_ := reqif.CONTENT
					var reqifs []*REQIF
					_, ok := res[content_]
					if ok {
						reqifs = res[content_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[content_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CONTENT
	case CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HEADER
	case HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CONTENT:
		res = "CONTENT"
	case HEADER:
		res = "HEADER"
	case REQIF:
		res = "REQIF"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *CONTENT:
		res = "CONTENT"
	case *HEADER:
		res = "HEADER"
	case *REQIF:
		res = "REQIF"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CONTENT:
		res = []string{"Name"}
	case HEADER:
		res = []string{"Name", "IDENTIFIERAttr", "COMMENT", "CREATIONTIME", "REPOSITORYID", "REQIFTOOLID", "REQIFVERSION", "SOURCETOOLID", "TITLE"}
	case REQIF:
		res = []string{"Name", "HEADER", "CONTENT"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case CONTENT:
		var rf ReverseField
		_ = rf
	case HEADER:
		var rf ReverseField
		_ = rf
	case REQIF:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *CONTENT:
		res = []string{"Name"}
	case *HEADER:
		res = []string{"Name", "IDENTIFIERAttr", "COMMENT", "CREATIONTIME", "REPOSITORYID", "REQIFTOOLID", "REQIFVERSION", "SOURCETOOLID", "TITLE"}
	case *REQIF:
		res = []string{"Name", "HEADER", "CONTENT"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATIONTIME":
			res = inferedInstance.CREATIONTIME
		case "REPOSITORYID":
			res = inferedInstance.REPOSITORYID
		case "REQIFTOOLID":
			res = inferedInstance.REQIFTOOLID
		case "REQIFVERSION":
			res = inferedInstance.REQIFVERSION
		case "SOURCETOOLID":
			res = inferedInstance.SOURCETOOLID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case *REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "HEADER":
			if inferedInstance.HEADER != nil {
				res = inferedInstance.HEADER.Name
			}
		case "CONTENT":
			if inferedInstance.CONTENT != nil {
				res = inferedInstance.CONTENT.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIERAttr":
			res = inferedInstance.IDENTIFIERAttr
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATIONTIME":
			res = inferedInstance.CREATIONTIME
		case "REPOSITORYID":
			res = inferedInstance.REPOSITORYID
		case "REQIFTOOLID":
			res = inferedInstance.REQIFTOOLID
		case "REQIFVERSION":
			res = inferedInstance.REQIFVERSION
		case "SOURCETOOLID":
			res = inferedInstance.SOURCETOOLID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "HEADER":
			if inferedInstance.HEADER != nil {
				res = inferedInstance.HEADER.Name
			}
		case "CONTENT":
			if inferedInstance.CONTENT != nil {
				res = inferedInstance.CONTENT.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
