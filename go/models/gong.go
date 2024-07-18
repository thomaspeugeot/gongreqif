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
	REQIFs           map[*REQIF]any
	REQIFs_mapString map[string]*REQIF

	// insertion point for slice of pointers maps

	OnAfterREQIFCreateCallback OnAfterCreateInterface[REQIF]
	OnAfterREQIFUpdateCallback OnAfterUpdateInterface[REQIF]
	OnAfterREQIFDeleteCallback OnAfterDeleteInterface[REQIF]
	OnAfterREQIFReadCallback   OnAfterReadInterface[REQIF]

	REQ_IF_CONTENTs           map[*REQ_IF_CONTENT]any
	REQ_IF_CONTENTs_mapString map[string]*REQ_IF_CONTENT

	// insertion point for slice of pointers maps
	REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT
	REQ_IF_CONTENT_SPECIFICATION_TYPES_reverseMap map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT
	REQ_IF_CONTENT_SPECIFICATIONS_reverseMap map[*SPECIFICATION]*REQ_IF_CONTENT

	OnAfterREQ_IF_CONTENTCreateCallback OnAfterCreateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTUpdateCallback OnAfterUpdateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTDeleteCallback OnAfterDeleteInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTReadCallback   OnAfterReadInterface[REQ_IF_CONTENT]

	REQ_IF_HEADERs           map[*REQ_IF_HEADER]any
	REQ_IF_HEADERs_mapString map[string]*REQ_IF_HEADER

	// insertion point for slice of pointers maps

	OnAfterREQ_IF_HEADERCreateCallback OnAfterCreateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERUpdateCallback OnAfterUpdateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERDeleteCallback OnAfterDeleteInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERReadCallback   OnAfterReadInterface[REQ_IF_HEADER]

	SPECIFICATIONs           map[*SPECIFICATION]any
	SPECIFICATIONs_mapString map[string]*SPECIFICATION

	// insertion point for slice of pointers maps

	OnAfterSPECIFICATIONCreateCallback OnAfterCreateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONUpdateCallback OnAfterUpdateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONDeleteCallback OnAfterDeleteInterface[SPECIFICATION]
	OnAfterSPECIFICATIONReadCallback   OnAfterReadInterface[SPECIFICATION]

	SPECIFICATION_TYPEs           map[*SPECIFICATION_TYPE]any
	SPECIFICATION_TYPEs_mapString map[string]*SPECIFICATION_TYPE

	// insertion point for slice of pointers maps

	OnAfterSPECIFICATION_TYPECreateCallback OnAfterCreateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEUpdateCallback OnAfterUpdateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEDeleteCallback OnAfterDeleteInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEReadCallback   OnAfterReadInterface[SPECIFICATION_TYPE]

	SPEC_HIERARCHYs           map[*SPEC_HIERARCHY]any
	SPEC_HIERARCHYs_mapString map[string]*SPEC_HIERARCHY

	// insertion point for slice of pointers maps

	OnAfterSPEC_HIERARCHYCreateCallback OnAfterCreateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYUpdateCallback OnAfterUpdateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYDeleteCallback OnAfterDeleteInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYReadCallback   OnAfterReadInterface[SPEC_HIERARCHY]

	SPEC_OBJECT_TYPEs           map[*SPEC_OBJECT_TYPE]any
	SPEC_OBJECT_TYPEs_mapString map[string]*SPEC_OBJECT_TYPE

	// insertion point for slice of pointers maps

	OnAfterSPEC_OBJECT_TYPECreateCallback OnAfterCreateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEUpdateCallback OnAfterUpdateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEDeleteCallback OnAfterDeleteInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEReadCallback   OnAfterReadInterface[SPEC_OBJECT_TYPE]

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
	CommitREQIF(reqif *REQIF)
	CheckoutREQIF(reqif *REQIF)
	CommitREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CheckoutREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CommitREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CheckoutREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CommitSPECIFICATION(specification *SPECIFICATION)
	CheckoutSPECIFICATION(specification *SPECIFICATION)
	CommitSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CheckoutSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CommitSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CheckoutSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CommitSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	CheckoutSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		REQIFs:           make(map[*REQIF]any),
		REQIFs_mapString: make(map[string]*REQIF),

		REQ_IF_CONTENTs:           make(map[*REQ_IF_CONTENT]any),
		REQ_IF_CONTENTs_mapString: make(map[string]*REQ_IF_CONTENT),

		REQ_IF_HEADERs:           make(map[*REQ_IF_HEADER]any),
		REQ_IF_HEADERs_mapString: make(map[string]*REQ_IF_HEADER),

		SPECIFICATIONs:           make(map[*SPECIFICATION]any),
		SPECIFICATIONs_mapString: make(map[string]*SPECIFICATION),

		SPECIFICATION_TYPEs:           make(map[*SPECIFICATION_TYPE]any),
		SPECIFICATION_TYPEs_mapString: make(map[string]*SPECIFICATION_TYPE),

		SPEC_HIERARCHYs:           make(map[*SPEC_HIERARCHY]any),
		SPEC_HIERARCHYs_mapString: make(map[string]*SPEC_HIERARCHY),

		SPEC_OBJECT_TYPEs:           make(map[*SPEC_OBJECT_TYPE]any),
		SPEC_OBJECT_TYPEs_mapString: make(map[string]*SPEC_OBJECT_TYPE),

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
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["REQIF"] = len(stage.REQIFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)

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

// Stage puts req_if_content to the model stage
func (req_if_content *REQ_IF_CONTENT) Stage(stage *StageStruct) *REQ_IF_CONTENT {
	stage.REQ_IF_CONTENTs[req_if_content] = __member
	stage.REQ_IF_CONTENTs_mapString[req_if_content.Name] = req_if_content

	return req_if_content
}

// Unstage removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) Unstage(stage *StageStruct) *REQ_IF_CONTENT {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
	return req_if_content
}

// UnstageVoid removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
}

// commit req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Commit(stage *StageStruct) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

func (req_if_content *REQ_IF_CONTENT) CommitVoid(stage *StageStruct) {
	req_if_content.Commit(stage)
}

// Checkout req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Checkout(stage *StageStruct) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

// for satisfaction of GongStruct interface
func (req_if_content *REQ_IF_CONTENT) GetName() (res string) {
	return req_if_content.Name
}

// Stage puts req_if_header to the model stage
func (req_if_header *REQ_IF_HEADER) Stage(stage *StageStruct) *REQ_IF_HEADER {
	stage.REQ_IF_HEADERs[req_if_header] = __member
	stage.REQ_IF_HEADERs_mapString[req_if_header.Name] = req_if_header

	return req_if_header
}

// Unstage removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) Unstage(stage *StageStruct) *REQ_IF_HEADER {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
	return req_if_header
}

// UnstageVoid removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
}

// commit req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Commit(stage *StageStruct) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

func (req_if_header *REQ_IF_HEADER) CommitVoid(stage *StageStruct) {
	req_if_header.Commit(stage)
}

// Checkout req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Checkout(stage *StageStruct) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

// for satisfaction of GongStruct interface
func (req_if_header *REQ_IF_HEADER) GetName() (res string) {
	return req_if_header.Name
}

// Stage puts specification to the model stage
func (specification *SPECIFICATION) Stage(stage *StageStruct) *SPECIFICATION {
	stage.SPECIFICATIONs[specification] = __member
	stage.SPECIFICATIONs_mapString[specification.Name] = specification

	return specification
}

// Unstage removes specification off the model stage
func (specification *SPECIFICATION) Unstage(stage *StageStruct) *SPECIFICATION {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
	return specification
}

// UnstageVoid removes specification off the model stage
func (specification *SPECIFICATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
}

// commit specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Commit(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION(specification)
		}
	}
	return specification
}

func (specification *SPECIFICATION) CommitVoid(stage *StageStruct) {
	specification.Commit(stage)
}

// Checkout specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Checkout(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION(specification)
		}
	}
	return specification
}

// for satisfaction of GongStruct interface
func (specification *SPECIFICATION) GetName() (res string) {
	return specification.Name
}

// Stage puts specification_type to the model stage
func (specification_type *SPECIFICATION_TYPE) Stage(stage *StageStruct) *SPECIFICATION_TYPE {
	stage.SPECIFICATION_TYPEs[specification_type] = __member
	stage.SPECIFICATION_TYPEs_mapString[specification_type.Name] = specification_type

	return specification_type
}

// Unstage removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) Unstage(stage *StageStruct) *SPECIFICATION_TYPE {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
	return specification_type
}

// UnstageVoid removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
}

// commit specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Commit(stage *StageStruct) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

func (specification_type *SPECIFICATION_TYPE) CommitVoid(stage *StageStruct) {
	specification_type.Commit(stage)
}

// Checkout specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Checkout(stage *StageStruct) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

// for satisfaction of GongStruct interface
func (specification_type *SPECIFICATION_TYPE) GetName() (res string) {
	return specification_type.Name
}

// Stage puts spec_hierarchy to the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Stage(stage *StageStruct) *SPEC_HIERARCHY {
	stage.SPEC_HIERARCHYs[spec_hierarchy] = __member
	stage.SPEC_HIERARCHYs_mapString[spec_hierarchy.Name] = spec_hierarchy

	return spec_hierarchy
}

// Unstage removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Unstage(stage *StageStruct) *SPEC_HIERARCHY {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
	return spec_hierarchy
}

// UnstageVoid removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
}

// commit spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Commit(stage *StageStruct) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

func (spec_hierarchy *SPEC_HIERARCHY) CommitVoid(stage *StageStruct) {
	spec_hierarchy.Commit(stage)
}

// Checkout spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Checkout(stage *StageStruct) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

// for satisfaction of GongStruct interface
func (spec_hierarchy *SPEC_HIERARCHY) GetName() (res string) {
	return spec_hierarchy.Name
}

// Stage puts spec_object_type to the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Stage(stage *StageStruct) *SPEC_OBJECT_TYPE {
	stage.SPEC_OBJECT_TYPEs[spec_object_type] = __member
	stage.SPEC_OBJECT_TYPEs_mapString[spec_object_type.Name] = spec_object_type

	return spec_object_type
}

// Unstage removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Unstage(stage *StageStruct) *SPEC_OBJECT_TYPE {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
	return spec_object_type
}

// UnstageVoid removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
}

// commit spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Commit(stage *StageStruct) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

func (spec_object_type *SPEC_OBJECT_TYPE) CommitVoid(stage *StageStruct) {
	spec_object_type.Commit(stage)
}

// Checkout spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Checkout(stage *StageStruct) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

// for satisfaction of GongStruct interface
func (spec_object_type *SPEC_OBJECT_TYPE) GetName() (res string) {
	return spec_object_type.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMREQIF(REQIF *REQIF)
	CreateORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	CreateORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	CreateORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	CreateORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	CreateORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	CreateORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMREQIF(REQIF *REQIF)
	DeleteORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	DeleteORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	DeleteORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	DeleteORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	DeleteORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	DeleteORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.REQIFs = make(map[*REQIF]any)
	stage.REQIFs_mapString = make(map[string]*REQIF)

	stage.REQ_IF_CONTENTs = make(map[*REQ_IF_CONTENT]any)
	stage.REQ_IF_CONTENTs_mapString = make(map[string]*REQ_IF_CONTENT)

	stage.REQ_IF_HEADERs = make(map[*REQ_IF_HEADER]any)
	stage.REQ_IF_HEADERs_mapString = make(map[string]*REQ_IF_HEADER)

	stage.SPECIFICATIONs = make(map[*SPECIFICATION]any)
	stage.SPECIFICATIONs_mapString = make(map[string]*SPECIFICATION)

	stage.SPECIFICATION_TYPEs = make(map[*SPECIFICATION_TYPE]any)
	stage.SPECIFICATION_TYPEs_mapString = make(map[string]*SPECIFICATION_TYPE)

	stage.SPEC_HIERARCHYs = make(map[*SPEC_HIERARCHY]any)
	stage.SPEC_HIERARCHYs_mapString = make(map[string]*SPEC_HIERARCHY)

	stage.SPEC_OBJECT_TYPEs = make(map[*SPEC_OBJECT_TYPE]any)
	stage.SPEC_OBJECT_TYPEs_mapString = make(map[string]*SPEC_OBJECT_TYPE)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.REQIFs = nil
	stage.REQIFs_mapString = nil

	stage.REQ_IF_CONTENTs = nil
	stage.REQ_IF_CONTENTs_mapString = nil

	stage.REQ_IF_HEADERs = nil
	stage.REQ_IF_HEADERs_mapString = nil

	stage.SPECIFICATIONs = nil
	stage.SPECIFICATIONs_mapString = nil

	stage.SPECIFICATION_TYPEs = nil
	stage.SPECIFICATION_TYPEs_mapString = nil

	stage.SPEC_HIERARCHYs = nil
	stage.SPEC_HIERARCHYs_mapString = nil

	stage.SPEC_OBJECT_TYPEs = nil
	stage.SPEC_OBJECT_TYPEs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for reqif := range stage.REQIFs {
		reqif.Unstage(stage)
	}

	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_content.Unstage(stage)
	}

	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_header.Unstage(stage)
	}

	for specification := range stage.SPECIFICATIONs {
		specification.Unstage(stage)
	}

	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_type.Unstage(stage)
	}

	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchy.Unstage(stage)
	}

	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_type.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	REQIF | REQ_IF_CONTENT | REQ_IF_HEADER | SPECIFICATION | SPECIFICATION_TYPE | SPEC_HIERARCHY | SPEC_OBJECT_TYPE
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
	*REQIF | *REQ_IF_CONTENT | *REQ_IF_HEADER | *SPECIFICATION | *SPECIFICATION_TYPE | *SPEC_HIERARCHY | *SPEC_OBJECT_TYPE
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
		map[*REQIF]any |
		map[*REQ_IF_CONTENT]any |
		map[*REQ_IF_HEADER]any |
		map[*SPECIFICATION]any |
		map[*SPECIFICATION_TYPE]any |
		map[*SPEC_HIERARCHY]any |
		map[*SPEC_OBJECT_TYPE]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*REQIF |
		map[string]*REQ_IF_CONTENT |
		map[string]*REQ_IF_HEADER |
		map[string]*SPECIFICATION |
		map[string]*SPECIFICATION_TYPE |
		map[string]*SPEC_HIERARCHY |
		map[string]*SPEC_OBJECT_TYPE |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*REQIF]any:
		return any(&stage.REQIFs).(*Type)
	case map[*REQ_IF_CONTENT]any:
		return any(&stage.REQ_IF_CONTENTs).(*Type)
	case map[*REQ_IF_HEADER]any:
		return any(&stage.REQ_IF_HEADERs).(*Type)
	case map[*SPECIFICATION]any:
		return any(&stage.SPECIFICATIONs).(*Type)
	case map[*SPECIFICATION_TYPE]any:
		return any(&stage.SPECIFICATION_TYPEs).(*Type)
	case map[*SPEC_HIERARCHY]any:
		return any(&stage.SPEC_HIERARCHYs).(*Type)
	case map[*SPEC_OBJECT_TYPE]any:
		return any(&stage.SPEC_OBJECT_TYPEs).(*Type)
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
	case map[string]*REQIF:
		return any(&stage.REQIFs_mapString).(*Type)
	case map[string]*REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*Type)
	case map[string]*REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*Type)
	case map[string]*SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*Type)
	case map[string]*SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*Type)
	case map[string]*SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*Type)
	case map[string]*SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*Type)
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
	case REQIF:
		return any(&stage.REQIFs).(*map[*Type]any)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[*Type]any)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[*Type]any)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[*Type]any)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[*Type]any)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[*Type]any)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[*Type]any)
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
	case *REQIF:
		return any(&stage.REQIFs).(*map[Type]any)
	case *REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[Type]any)
	case *REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[Type]any)
	case *SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[Type]any)
	case *SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[Type]any)
	case *SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[Type]any)
	case *SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[Type]any)
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
	case REQIF:
		return any(&stage.REQIFs_mapString).(*map[string]*Type)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*map[string]*Type)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*map[string]*Type)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*map[string]*Type)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*map[string]*Type)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*map[string]*Type)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*map[string]*Type)
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
	case REQIF:
		return any(&REQIF{
			// Initialisation of associations
			// field is initialized with an instance of REQ_IF_HEADER with the name of the field
			REQ_IF_HEADER: &REQ_IF_HEADER{Name: "REQ_IF_HEADER"},
			// field is initialized with an instance of REQ_IF_CONTENT with the name of the field
			REQ_IF_CONTENT: &REQ_IF_CONTENT{Name: "REQ_IF_CONTENT"},
		}).(*Type)
	case REQ_IF_CONTENT:
		return any(&REQ_IF_CONTENT{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_OBJECT_TYPE with the name of the field
			SPEC_OBJECT_TYPES: []*SPEC_OBJECT_TYPE{{Name: "SPEC_OBJECT_TYPES"}},
			// field is initialized with an instance of SPECIFICATION_TYPE with the name of the field
			SPECIFICATION_TYPES: []*SPECIFICATION_TYPE{{Name: "SPECIFICATION_TYPES"}},
			// field is initialized with an instance of SPECIFICATION with the name of the field
			SPECIFICATIONS: []*SPECIFICATION{{Name: "SPECIFICATIONS"}},
		}).(*Type)
	case REQ_IF_HEADER:
		return any(&REQ_IF_HEADER{
			// Initialisation of associations
		}).(*Type)
	case SPECIFICATION:
		return any(&SPECIFICATION{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_HIERARCHY with the name of the field
			CHILDREN: &SPEC_HIERARCHY{Name: "CHILDREN"},
			// field is initialized with an instance of SPECIFICATION_TYPE with the name of the field
			SPECIFICATION_TYPE: &SPECIFICATION_TYPE{Name: "SPECIFICATION_TYPE"},
		}).(*Type)
	case SPECIFICATION_TYPE:
		return any(&SPECIFICATION_TYPE{
			// Initialisation of associations
		}).(*Type)
	case SPEC_HIERARCHY:
		return any(&SPEC_HIERARCHY{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_HIERARCHY with the name of the field
			CHILDREN: &SPEC_HIERARCHY{Name: "CHILDREN"},
		}).(*Type)
	case SPEC_OBJECT_TYPE:
		return any(&SPEC_OBJECT_TYPE{
			// Initialisation of associations
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
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		case "REQ_IF_HEADER":
			res := make(map[*REQ_IF_HEADER][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.REQ_IF_HEADER != nil {
					req_if_header_ := reqif.REQ_IF_HEADER
					var reqifs []*REQIF
					_, ok := res[req_if_header_]
					if ok {
						reqifs = res[req_if_header_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[req_if_header_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "REQ_IF_CONTENT":
			res := make(map[*REQ_IF_CONTENT][]*REQIF)
			for reqif := range stage.REQIFs {
				if reqif.REQ_IF_CONTENT != nil {
					req_if_content_ := reqif.REQ_IF_CONTENT
					var reqifs []*REQIF
					_, ok := res[req_if_content_]
					if ok {
						reqifs = res[req_if_content_]
					} else {
						reqifs = make([]*REQIF, 0)
					}
					reqifs = append(reqifs, reqif)
					res[req_if_content_] = reqifs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		case "CHILDREN":
			res := make(map[*SPEC_HIERARCHY][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.CHILDREN != nil {
					spec_hierarchy_ := specification.CHILDREN
					var specifications []*SPECIFICATION
					_, ok := res[spec_hierarchy_]
					if ok {
						specifications = res[spec_hierarchy_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[spec_hierarchy_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFICATION_TYPE":
			res := make(map[*SPECIFICATION_TYPE][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.SPECIFICATION_TYPE != nil {
					specification_type_ := specification.SPECIFICATION_TYPE
					var specifications []*SPECIFICATION
					_, ok := res[specification_type_]
					if ok {
						specifications = res[specification_type_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[specification_type_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		case "CHILDREN":
			res := make(map[*SPEC_HIERARCHY][]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				if spec_hierarchy.CHILDREN != nil {
					spec_hierarchy_ := spec_hierarchy.CHILDREN
					var spec_hierarchys []*SPEC_HIERARCHY
					_, ok := res[spec_hierarchy_]
					if ok {
						spec_hierarchys = res[spec_hierarchy_]
					} else {
						spec_hierarchys = make([]*SPEC_HIERARCHY, 0)
					}
					spec_hierarchys = append(spec_hierarchys, spec_hierarchy)
					res[spec_hierarchy_] = spec_hierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of REQIF
	case REQIF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "SPEC_OBJECT_TYPES":
			res := make(map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, spec_object_type_ := range req_if_content.SPEC_OBJECT_TYPES {
					res[spec_object_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECIFICATION_TYPES":
			res := make(map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, specification_type_ := range req_if_content.SPECIFICATION_TYPES {
					res[specification_type_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		case "SPECIFICATIONS":
			res := make(map[*SPECIFICATION]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				for _, specification_ := range req_if_content.SPECIFICATIONS {
					res[specification_] = req_if_content
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
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
	case REQIF:
		res = "REQIF"
	case REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case SPECIFICATION:
		res = "SPECIFICATION"
	case SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *REQIF:
		res = "REQIF"
	case *REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case *REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case *SPECIFICATION:
		res = "SPECIFICATION"
	case *SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case *SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case *SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case REQIF:
		res = []string{"Name", "REQ_IF_HEADER", "REQ_IF_CONTENT"}
	case REQ_IF_CONTENT:
		res = []string{"Name", "SPEC_OBJECT_TYPES", "SPECIFICATION_TYPES", "SPECIFICATIONS"}
	case REQ_IF_HEADER:
		res = []string{"Name", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case SPECIFICATION:
		res = []string{"Name", "CHILDREN", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "TYPE", "SPECIFICATION_TYPE"}
	case SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_HIERARCHY:
		res = []string{"Name", "CHILDREN", "OBJECT", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
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
	case REQIF:
		var rf ReverseField
		_ = rf
	case REQ_IF_CONTENT:
		var rf ReverseField
		_ = rf
	case REQ_IF_HEADER:
		var rf ReverseField
		_ = rf
	case SPECIFICATION:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPECIFICATIONS"
		res = append(res, rf)
	case SPECIFICATION_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPECIFICATION_TYPES"
		res = append(res, rf)
	case SPEC_HIERARCHY:
		var rf ReverseField
		_ = rf
	case SPEC_OBJECT_TYPE:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "REQ_IF_CONTENT"
		rf.Fieldname = "SPEC_OBJECT_TYPES"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *REQIF:
		res = []string{"Name", "REQ_IF_HEADER", "REQ_IF_CONTENT"}
	case *REQ_IF_CONTENT:
		res = []string{"Name", "SPEC_OBJECT_TYPES", "SPECIFICATION_TYPES", "SPECIFICATIONS"}
	case *REQ_IF_HEADER:
		res = []string{"Name", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case *SPECIFICATION:
		res = []string{"Name", "CHILDREN", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "TYPE", "SPECIFICATION_TYPE"}
	case *SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_HIERARCHY:
		res = []string{"Name", "CHILDREN", "OBJECT", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_HEADER":
			if inferedInstance.REQ_IF_HEADER != nil {
				res = inferedInstance.REQ_IF_HEADER.Name
			}
		case "REQ_IF_CONTENT":
			if inferedInstance.REQ_IF_CONTENT != nil {
				res = inferedInstance.REQ_IF_CONTENT.Name
			}
		}
	case *REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_TYPES":
			for idx, __instance__ := range inferedInstance.SPEC_OBJECT_TYPES {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATION_TYPES":
			for idx, __instance__ := range inferedInstance.SPECIFICATION_TYPES {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATIONS":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONS {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATION_TIME":
			res = inferedInstance.CREATION_TIME.String()
		case "REPOSITORY_ID":
			res = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case *SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "TYPE":
			res = inferedInstance.TYPE
		case "SPECIFICATION_TYPE":
			if inferedInstance.SPECIFICATION_TYPE != nil {
				res = inferedInstance.SPECIFICATION_TYPE.Name
			}
		}
	case *SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "OBJECT":
			res = inferedInstance.OBJECT
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "IS_TABLE_INTERNAL":
			res = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case REQIF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_HEADER":
			if inferedInstance.REQ_IF_HEADER != nil {
				res = inferedInstance.REQ_IF_HEADER.Name
			}
		case "REQ_IF_CONTENT":
			if inferedInstance.REQ_IF_CONTENT != nil {
				res = inferedInstance.REQ_IF_CONTENT.Name
			}
		}
	case REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_TYPES":
			for idx, __instance__ := range inferedInstance.SPEC_OBJECT_TYPES {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATION_TYPES":
			for idx, __instance__ := range inferedInstance.SPECIFICATION_TYPES {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SPECIFICATIONS":
			for idx, __instance__ := range inferedInstance.SPECIFICATIONS {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATION_TIME":
			res = inferedInstance.CREATION_TIME.String()
		case "REPOSITORY_ID":
			res = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "TYPE":
			res = inferedInstance.TYPE
		case "SPECIFICATION_TYPE":
			if inferedInstance.SPECIFICATION_TYPE != nil {
				res = inferedInstance.SPECIFICATION_TYPE.Name
			}
		}
	case SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "OBJECT":
			res = inferedInstance.OBJECT
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "IS_TABLE_INTERNAL":
			res = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE.String()
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
