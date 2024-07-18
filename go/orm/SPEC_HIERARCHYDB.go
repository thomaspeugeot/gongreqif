// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/thomaspeugeot/gongreqif/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SPEC_HIERARCHY_sql sql.NullBool
var dummy_SPEC_HIERARCHY_time time.Duration
var dummy_SPEC_HIERARCHY_sort sort.Float64Slice

// SPEC_HIERARCHYAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spec_hierarchyAPI
type SPEC_HIERARCHYAPI struct {
	gorm.Model

	models.SPEC_HIERARCHY_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SPEC_HIERARCHYPointersEncoding SPEC_HIERARCHYPointersEncoding
}

// SPEC_HIERARCHYPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SPEC_HIERARCHYPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// SPEC_HIERARCHYDB describes a spec_hierarchy in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spec_hierarchyDB
type SPEC_HIERARCHYDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spec_hierarchyDB.Name
	Name_Data sql.NullString

	// Declation for basic field spec_hierarchyDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field spec_hierarchyDB.IDENTIFIER
	IDENTIFIER_Data sql.NullString

	// Declation for basic field spec_hierarchyDB.IS_EDITABLE
	// provide the sql storage for the boolan
	IS_EDITABLE_Data sql.NullBool

	// Declation for basic field spec_hierarchyDB.IS_TABLE_INTERNAL
	// provide the sql storage for the boolan
	IS_TABLE_INTERNAL_Data sql.NullBool

	// Declation for basic field spec_hierarchyDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullTime

	// Declation for basic field spec_hierarchyDB.LONG_NAME
	LONG_NAME_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SPEC_HIERARCHYPointersEncoding
}

// SPEC_HIERARCHYDBs arrays spec_hierarchyDBs
// swagger:response spec_hierarchyDBsResponse
type SPEC_HIERARCHYDBs []SPEC_HIERARCHYDB

// SPEC_HIERARCHYDBResponse provides response
// swagger:response spec_hierarchyDBResponse
type SPEC_HIERARCHYDBResponse struct {
	SPEC_HIERARCHYDB
}

// SPEC_HIERARCHYWOP is a SPEC_HIERARCHY without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SPEC_HIERARCHYWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IDENTIFIER string `xlsx:"3"`

	IS_EDITABLE bool `xlsx:"4"`

	IS_TABLE_INTERNAL bool `xlsx:"5"`

	LAST_CHANGE time.Time `xlsx:"6"`

	LONG_NAME string `xlsx:"7"`
	// insertion for WOP pointer fields
}

var SPEC_HIERARCHY_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IDENTIFIER",
	"IS_EDITABLE",
	"IS_TABLE_INTERNAL",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoSPEC_HIERARCHYStruct struct {
	// stores SPEC_HIERARCHYDB according to their gorm ID
	Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB map[uint]*SPEC_HIERARCHYDB

	// stores SPEC_HIERARCHYDB ID according to SPEC_HIERARCHY address
	Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID map[*models.SPEC_HIERARCHY]uint

	// stores SPEC_HIERARCHY according to their gorm ID
	Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr map[uint]*models.SPEC_HIERARCHY

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSPEC_HIERARCHY.stage
	return
}

func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) GetDB() *gorm.DB {
	return backRepoSPEC_HIERARCHY.db
}

// GetSPEC_HIERARCHYDBFromSPEC_HIERARCHYPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) GetSPEC_HIERARCHYDBFromSPEC_HIERARCHYPtr(spec_hierarchy *models.SPEC_HIERARCHY) (spec_hierarchyDB *SPEC_HIERARCHYDB) {
	id := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]
	spec_hierarchyDB = backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[id]
	return
}

// BackRepoSPEC_HIERARCHY.CommitPhaseOne commits all staged instances of SPEC_HIERARCHY to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		backRepoSPEC_HIERARCHY.CommitPhaseOneInstance(spec_hierarchy)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spec_hierarchy := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr {
		if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; !ok {
			backRepoSPEC_HIERARCHY.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSPEC_HIERARCHY.CommitDeleteInstance commits deletion of SPEC_HIERARCHY to the BackRepo
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CommitDeleteInstance(id uint) (Error error) {

	spec_hierarchy := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[id]

	// spec_hierarchy is not staged anymore, remove spec_hierarchyDB
	spec_hierarchyDB := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[id]
	query := backRepoSPEC_HIERARCHY.db.Unscoped().Delete(&spec_hierarchyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID, spec_hierarchy)
	delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr, id)
	delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB, id)

	return
}

// BackRepoSPEC_HIERARCHY.CommitPhaseOneInstance commits spec_hierarchy staged instances of SPEC_HIERARCHY to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CommitPhaseOneInstance(spec_hierarchy *models.SPEC_HIERARCHY) (Error error) {

	// check if the spec_hierarchy is not commited yet
	if _, ok := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]; ok {
		return
	}

	// initiate spec_hierarchy
	var spec_hierarchyDB SPEC_HIERARCHYDB
	spec_hierarchyDB.CopyBasicFieldsFromSPEC_HIERARCHY(spec_hierarchy)

	query := backRepoSPEC_HIERARCHY.db.Create(&spec_hierarchyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy] = spec_hierarchyDB.ID
	backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID] = spec_hierarchy
	backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[spec_hierarchyDB.ID] = &spec_hierarchyDB

	return
}

// BackRepoSPEC_HIERARCHY.CommitPhaseTwo commits all staged instances of SPEC_HIERARCHY to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spec_hierarchy := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr {
		backRepoSPEC_HIERARCHY.CommitPhaseTwoInstance(backRepo, idx, spec_hierarchy)
	}

	return
}

// BackRepoSPEC_HIERARCHY.CommitPhaseTwoInstance commits {{structname }} of models.SPEC_HIERARCHY to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spec_hierarchy *models.SPEC_HIERARCHY) (Error error) {

	// fetch matching spec_hierarchyDB
	if spec_hierarchyDB, ok := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[idx]; ok {

		spec_hierarchyDB.CopyBasicFieldsFromSPEC_HIERARCHY(spec_hierarchy)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoSPEC_HIERARCHY.db.Save(&spec_hierarchyDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SPEC_HIERARCHY intance %s", spec_hierarchy.Name))
		return err
	}

	return
}

// BackRepoSPEC_HIERARCHY.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CheckoutPhaseOne() (Error error) {

	spec_hierarchyDBArray := make([]SPEC_HIERARCHYDB, 0)
	query := backRepoSPEC_HIERARCHY.db.Find(&spec_hierarchyDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spec_hierarchyInstancesToBeRemovedFromTheStage := make(map[*models.SPEC_HIERARCHY]any)
	for key, value := range backRepoSPEC_HIERARCHY.stage.SPEC_HIERARCHYs {
		spec_hierarchyInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spec_hierarchyDB := range spec_hierarchyDBArray {
		backRepoSPEC_HIERARCHY.CheckoutPhaseOneInstance(&spec_hierarchyDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spec_hierarchy, ok := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]
		if ok {
			delete(spec_hierarchyInstancesToBeRemovedFromTheStage, spec_hierarchy)
		}
	}

	// remove from stage and back repo's 3 maps all spec_hierarchys that are not in the checkout
	for spec_hierarchy := range spec_hierarchyInstancesToBeRemovedFromTheStage {
		spec_hierarchy.Unstage(backRepoSPEC_HIERARCHY.GetStage())

		// remove instance from the back repo 3 maps
		spec_hierarchyID := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]
		delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID, spec_hierarchy)
		delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB, spec_hierarchyID)
		delete(backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr, spec_hierarchyID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spec_hierarchyDB that has been found in the DB, updates the backRepo and stages the
// models version of the spec_hierarchyDB
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CheckoutPhaseOneInstance(spec_hierarchyDB *SPEC_HIERARCHYDB) (Error error) {

	spec_hierarchy, ok := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]
	if !ok {
		spec_hierarchy = new(models.SPEC_HIERARCHY)

		backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID] = spec_hierarchy
		backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy] = spec_hierarchyDB.ID

		// append model store with the new element
		spec_hierarchy.Name = spec_hierarchyDB.Name_Data.String
		spec_hierarchy.Stage(backRepoSPEC_HIERARCHY.GetStage())
	}
	spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY(spec_hierarchy)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spec_hierarchy.Stage(backRepoSPEC_HIERARCHY.GetStage())

	// preserve pointer to spec_hierarchyDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB)[spec_hierarchyDB hold variable pointers
	spec_hierarchyDB_Data := *spec_hierarchyDB
	preservedPtrToSPEC_HIERARCHY := &spec_hierarchyDB_Data
	backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[spec_hierarchyDB.ID] = preservedPtrToSPEC_HIERARCHY

	return
}

// BackRepoSPEC_HIERARCHY.CheckoutPhaseTwo Checkouts all staged instances of SPEC_HIERARCHY to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spec_hierarchyDB := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {
		backRepoSPEC_HIERARCHY.CheckoutPhaseTwoInstance(backRepo, spec_hierarchyDB)
	}
	return
}

// BackRepoSPEC_HIERARCHY.CheckoutPhaseTwoInstance Checkouts staged instances of SPEC_HIERARCHY to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spec_hierarchyDB *SPEC_HIERARCHYDB) (Error error) {

	spec_hierarchy := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]

	spec_hierarchyDB.DecodePointers(backRepo, spec_hierarchy)

	return
}

func (spec_hierarchyDB *SPEC_HIERARCHYDB) DecodePointers(backRepo *BackRepoStruct, spec_hierarchy *models.SPEC_HIERARCHY) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitSPEC_HIERARCHY allows commit of a single spec_hierarchy (if already staged)
func (backRepo *BackRepoStruct) CommitSPEC_HIERARCHY(spec_hierarchy *models.SPEC_HIERARCHY) {
	backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseOneInstance(spec_hierarchy)
	if id, ok := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]; ok {
		backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseTwoInstance(backRepo, id, spec_hierarchy)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSPEC_HIERARCHY allows checkout of a single spec_hierarchy (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSPEC_HIERARCHY(spec_hierarchy *models.SPEC_HIERARCHY) {
	// check if the spec_hierarchy is staged
	if _, ok := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]; ok {

		if id, ok := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID[spec_hierarchy]; ok {
			var spec_hierarchyDB SPEC_HIERARCHYDB
			spec_hierarchyDB.ID = id

			if err := backRepo.BackRepoSPEC_HIERARCHY.db.First(&spec_hierarchyDB, id).Error; err != nil {
				log.Fatalln("CheckoutSPEC_HIERARCHY : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseOneInstance(&spec_hierarchyDB)
			backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseTwoInstance(backRepo, &spec_hierarchyDB)
		}
	}
}

// CopyBasicFieldsFromSPEC_HIERARCHY
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsFromSPEC_HIERARCHY(spec_hierarchy *models.SPEC_HIERARCHY) {
	// insertion point for fields commit

	spec_hierarchyDB.Name_Data.String = spec_hierarchy.Name
	spec_hierarchyDB.Name_Data.Valid = true

	spec_hierarchyDB.DESC_Data.String = spec_hierarchy.DESC
	spec_hierarchyDB.DESC_Data.Valid = true

	spec_hierarchyDB.IDENTIFIER_Data.String = spec_hierarchy.IDENTIFIER
	spec_hierarchyDB.IDENTIFIER_Data.Valid = true

	spec_hierarchyDB.IS_EDITABLE_Data.Bool = spec_hierarchy.IS_EDITABLE
	spec_hierarchyDB.IS_EDITABLE_Data.Valid = true

	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool = spec_hierarchy.IS_TABLE_INTERNAL
	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Valid = true

	spec_hierarchyDB.LAST_CHANGE_Data.Time = spec_hierarchy.LAST_CHANGE
	spec_hierarchyDB.LAST_CHANGE_Data.Valid = true

	spec_hierarchyDB.LONG_NAME_Data.String = spec_hierarchy.LONG_NAME
	spec_hierarchyDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromSPEC_HIERARCHY_WOP
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsFromSPEC_HIERARCHY_WOP(spec_hierarchy *models.SPEC_HIERARCHY_WOP) {
	// insertion point for fields commit

	spec_hierarchyDB.Name_Data.String = spec_hierarchy.Name
	spec_hierarchyDB.Name_Data.Valid = true

	spec_hierarchyDB.DESC_Data.String = spec_hierarchy.DESC
	spec_hierarchyDB.DESC_Data.Valid = true

	spec_hierarchyDB.IDENTIFIER_Data.String = spec_hierarchy.IDENTIFIER
	spec_hierarchyDB.IDENTIFIER_Data.Valid = true

	spec_hierarchyDB.IS_EDITABLE_Data.Bool = spec_hierarchy.IS_EDITABLE
	spec_hierarchyDB.IS_EDITABLE_Data.Valid = true

	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool = spec_hierarchy.IS_TABLE_INTERNAL
	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Valid = true

	spec_hierarchyDB.LAST_CHANGE_Data.Time = spec_hierarchy.LAST_CHANGE
	spec_hierarchyDB.LAST_CHANGE_Data.Valid = true

	spec_hierarchyDB.LONG_NAME_Data.String = spec_hierarchy.LONG_NAME
	spec_hierarchyDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromSPEC_HIERARCHYWOP
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsFromSPEC_HIERARCHYWOP(spec_hierarchy *SPEC_HIERARCHYWOP) {
	// insertion point for fields commit

	spec_hierarchyDB.Name_Data.String = spec_hierarchy.Name
	spec_hierarchyDB.Name_Data.Valid = true

	spec_hierarchyDB.DESC_Data.String = spec_hierarchy.DESC
	spec_hierarchyDB.DESC_Data.Valid = true

	spec_hierarchyDB.IDENTIFIER_Data.String = spec_hierarchy.IDENTIFIER
	spec_hierarchyDB.IDENTIFIER_Data.Valid = true

	spec_hierarchyDB.IS_EDITABLE_Data.Bool = spec_hierarchy.IS_EDITABLE
	spec_hierarchyDB.IS_EDITABLE_Data.Valid = true

	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool = spec_hierarchy.IS_TABLE_INTERNAL
	spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Valid = true

	spec_hierarchyDB.LAST_CHANGE_Data.Time = spec_hierarchy.LAST_CHANGE
	spec_hierarchyDB.LAST_CHANGE_Data.Valid = true

	spec_hierarchyDB.LONG_NAME_Data.String = spec_hierarchy.LONG_NAME
	spec_hierarchyDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToSPEC_HIERARCHY
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsToSPEC_HIERARCHY(spec_hierarchy *models.SPEC_HIERARCHY) {
	// insertion point for checkout of basic fields (back repo to stage)
	spec_hierarchy.Name = spec_hierarchyDB.Name_Data.String
	spec_hierarchy.DESC = spec_hierarchyDB.DESC_Data.String
	spec_hierarchy.IDENTIFIER = spec_hierarchyDB.IDENTIFIER_Data.String
	spec_hierarchy.IS_EDITABLE = spec_hierarchyDB.IS_EDITABLE_Data.Bool
	spec_hierarchy.IS_TABLE_INTERNAL = spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool
	spec_hierarchy.LAST_CHANGE = spec_hierarchyDB.LAST_CHANGE_Data.Time
	spec_hierarchy.LONG_NAME = spec_hierarchyDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToSPEC_HIERARCHY_WOP
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsToSPEC_HIERARCHY_WOP(spec_hierarchy *models.SPEC_HIERARCHY_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spec_hierarchy.Name = spec_hierarchyDB.Name_Data.String
	spec_hierarchy.DESC = spec_hierarchyDB.DESC_Data.String
	spec_hierarchy.IDENTIFIER = spec_hierarchyDB.IDENTIFIER_Data.String
	spec_hierarchy.IS_EDITABLE = spec_hierarchyDB.IS_EDITABLE_Data.Bool
	spec_hierarchy.IS_TABLE_INTERNAL = spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool
	spec_hierarchy.LAST_CHANGE = spec_hierarchyDB.LAST_CHANGE_Data.Time
	spec_hierarchy.LONG_NAME = spec_hierarchyDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToSPEC_HIERARCHYWOP
func (spec_hierarchyDB *SPEC_HIERARCHYDB) CopyBasicFieldsToSPEC_HIERARCHYWOP(spec_hierarchy *SPEC_HIERARCHYWOP) {
	spec_hierarchy.ID = int(spec_hierarchyDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spec_hierarchy.Name = spec_hierarchyDB.Name_Data.String
	spec_hierarchy.DESC = spec_hierarchyDB.DESC_Data.String
	spec_hierarchy.IDENTIFIER = spec_hierarchyDB.IDENTIFIER_Data.String
	spec_hierarchy.IS_EDITABLE = spec_hierarchyDB.IS_EDITABLE_Data.Bool
	spec_hierarchy.IS_TABLE_INTERNAL = spec_hierarchyDB.IS_TABLE_INTERNAL_Data.Bool
	spec_hierarchy.LAST_CHANGE = spec_hierarchyDB.LAST_CHANGE_Data.Time
	spec_hierarchy.LONG_NAME = spec_hierarchyDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all SPEC_HIERARCHYDB instances in the backrepo
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SPEC_HIERARCHYDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPEC_HIERARCHYDB, 0)
	for _, spec_hierarchyDB := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {
		forBackup = append(forBackup, spec_hierarchyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SPEC_HIERARCHY ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SPEC_HIERARCHY file", err.Error())
	}
}

// Backup generates a json file from a slice of all SPEC_HIERARCHYDB instances in the backrepo
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPEC_HIERARCHYDB, 0)
	for _, spec_hierarchyDB := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {
		forBackup = append(forBackup, spec_hierarchyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SPEC_HIERARCHY")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SPEC_HIERARCHY_Fields, -1)
	for _, spec_hierarchyDB := range forBackup {

		var spec_hierarchyWOP SPEC_HIERARCHYWOP
		spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHYWOP(&spec_hierarchyWOP)

		row := sh.AddRow()
		row.WriteStruct(&spec_hierarchyWOP, -1)
	}
}

// RestoreXL from the "SPEC_HIERARCHY" sheet all SPEC_HIERARCHYDB instances
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSPEC_HIERARCHYid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SPEC_HIERARCHY"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSPEC_HIERARCHY.rowVisitorSPEC_HIERARCHY)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) rowVisitorSPEC_HIERARCHY(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spec_hierarchyWOP SPEC_HIERARCHYWOP
		row.ReadStruct(&spec_hierarchyWOP)

		// add the unmarshalled struct to the stage
		spec_hierarchyDB := new(SPEC_HIERARCHYDB)
		spec_hierarchyDB.CopyBasicFieldsFromSPEC_HIERARCHYWOP(&spec_hierarchyWOP)

		spec_hierarchyDB_ID_atBackupTime := spec_hierarchyDB.ID
		spec_hierarchyDB.ID = 0
		query := backRepoSPEC_HIERARCHY.db.Create(spec_hierarchyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[spec_hierarchyDB.ID] = spec_hierarchyDB
		BackRepoSPEC_HIERARCHYid_atBckpTime_newID[spec_hierarchyDB_ID_atBackupTime] = spec_hierarchyDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SPEC_HIERARCHYDB.json" in dirPath that stores an array
// of SPEC_HIERARCHYDB and stores it in the database
// the map BackRepoSPEC_HIERARCHYid_atBckpTime_newID is updated accordingly
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSPEC_HIERARCHYid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SPEC_HIERARCHYDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SPEC_HIERARCHY file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SPEC_HIERARCHYDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB
	for _, spec_hierarchyDB := range forRestore {

		spec_hierarchyDB_ID_atBackupTime := spec_hierarchyDB.ID
		spec_hierarchyDB.ID = 0
		query := backRepoSPEC_HIERARCHY.db.Create(spec_hierarchyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[spec_hierarchyDB.ID] = spec_hierarchyDB
		BackRepoSPEC_HIERARCHYid_atBckpTime_newID[spec_hierarchyDB_ID_atBackupTime] = spec_hierarchyDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SPEC_HIERARCHY file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SPEC_HIERARCHY>id_atBckpTime_newID
// to compute new index
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) RestorePhaseTwo() {

	for _, spec_hierarchyDB := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {

		// next line of code is to avert unused variable compilation error
		_ = spec_hierarchyDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSPEC_HIERARCHY.db.Model(spec_hierarchyDB).Updates(*spec_hierarchyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSPEC_HIERARCHY.ResetReversePointers commits all staged instances of SPEC_HIERARCHY to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spec_hierarchy := range backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr {
		backRepoSPEC_HIERARCHY.ResetReversePointersInstance(backRepo, idx, spec_hierarchy)
	}

	return
}

func (backRepoSPEC_HIERARCHY *BackRepoSPEC_HIERARCHYStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spec_hierarchy *models.SPEC_HIERARCHY) (Error error) {

	// fetch matching spec_hierarchyDB
	if spec_hierarchyDB, ok := backRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB[idx]; ok {
		_ = spec_hierarchyDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSPEC_HIERARCHYid_atBckpTime_newID map[uint]uint
