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
var dummy_SPECOBJECTS_sql sql.NullBool
var dummy_SPECOBJECTS_time time.Duration
var dummy_SPECOBJECTS_sort sort.Float64Slice

// SPECOBJECTSAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model specobjectsAPI
type SPECOBJECTSAPI struct {
	gorm.Model

	models.SPECOBJECTS_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SPECOBJECTSPointersEncoding SPECOBJECTSPointersEncoding
}

// SPECOBJECTSPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SPECOBJECTSPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field SPECOBJECT is a slice of pointers to another Struct (optional or 0..1)
	SPECOBJECT IntSlice `gorm:"type:TEXT"`
}

// SPECOBJECTSDB describes a specobjects in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model specobjectsDB
type SPECOBJECTSDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field specobjectsDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SPECOBJECTSPointersEncoding
}

// SPECOBJECTSDBs arrays specobjectsDBs
// swagger:response specobjectsDBsResponse
type SPECOBJECTSDBs []SPECOBJECTSDB

// SPECOBJECTSDBResponse provides response
// swagger:response specobjectsDBResponse
type SPECOBJECTSDBResponse struct {
	SPECOBJECTSDB
}

// SPECOBJECTSWOP is a SPECOBJECTS without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SPECOBJECTSWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var SPECOBJECTS_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoSPECOBJECTSStruct struct {
	// stores SPECOBJECTSDB according to their gorm ID
	Map_SPECOBJECTSDBID_SPECOBJECTSDB map[uint]*SPECOBJECTSDB

	// stores SPECOBJECTSDB ID according to SPECOBJECTS address
	Map_SPECOBJECTSPtr_SPECOBJECTSDBID map[*models.SPECOBJECTS]uint

	// stores SPECOBJECTS according to their gorm ID
	Map_SPECOBJECTSDBID_SPECOBJECTSPtr map[uint]*models.SPECOBJECTS

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSPECOBJECTS.stage
	return
}

func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) GetDB() *gorm.DB {
	return backRepoSPECOBJECTS.db
}

// GetSPECOBJECTSDBFromSPECOBJECTSPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) GetSPECOBJECTSDBFromSPECOBJECTSPtr(specobjects *models.SPECOBJECTS) (specobjectsDB *SPECOBJECTSDB) {
	id := backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]
	specobjectsDB = backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[id]
	return
}

// BackRepoSPECOBJECTS.CommitPhaseOne commits all staged instances of SPECOBJECTS to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for specobjects := range stage.SPECOBJECTSs {
		backRepoSPECOBJECTS.CommitPhaseOneInstance(specobjects)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, specobjects := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr {
		if _, ok := stage.SPECOBJECTSs[specobjects]; !ok {
			backRepoSPECOBJECTS.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSPECOBJECTS.CommitDeleteInstance commits deletion of SPECOBJECTS to the BackRepo
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CommitDeleteInstance(id uint) (Error error) {

	specobjects := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[id]

	// specobjects is not staged anymore, remove specobjectsDB
	specobjectsDB := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[id]
	query := backRepoSPECOBJECTS.db.Unscoped().Delete(&specobjectsDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID, specobjects)
	delete(backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr, id)
	delete(backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB, id)

	return
}

// BackRepoSPECOBJECTS.CommitPhaseOneInstance commits specobjects staged instances of SPECOBJECTS to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CommitPhaseOneInstance(specobjects *models.SPECOBJECTS) (Error error) {

	// check if the specobjects is not commited yet
	if _, ok := backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]; ok {
		return
	}

	// initiate specobjects
	var specobjectsDB SPECOBJECTSDB
	specobjectsDB.CopyBasicFieldsFromSPECOBJECTS(specobjects)

	query := backRepoSPECOBJECTS.db.Create(&specobjectsDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects] = specobjectsDB.ID
	backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID] = specobjects
	backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[specobjectsDB.ID] = &specobjectsDB

	return
}

// BackRepoSPECOBJECTS.CommitPhaseTwo commits all staged instances of SPECOBJECTS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, specobjects := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr {
		backRepoSPECOBJECTS.CommitPhaseTwoInstance(backRepo, idx, specobjects)
	}

	return
}

// BackRepoSPECOBJECTS.CommitPhaseTwoInstance commits {{structname }} of models.SPECOBJECTS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, specobjects *models.SPECOBJECTS) (Error error) {

	// fetch matching specobjectsDB
	if specobjectsDB, ok := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[idx]; ok {

		specobjectsDB.CopyBasicFieldsFromSPECOBJECTS(specobjects)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		specobjectsDB.SPECOBJECTSPointersEncoding.SPECOBJECT = make([]int, 0)
		// 2. encode
		for _, specobjectAssocEnd := range specobjects.SPECOBJECT {
			specobjectAssocEnd_DB :=
				backRepo.BackRepoSPECOBJECT.GetSPECOBJECTDBFromSPECOBJECTPtr(specobjectAssocEnd)
			
			// the stage might be inconsistant, meaning that the specobjectAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if specobjectAssocEnd_DB == nil {
				continue
			}
			
			specobjectsDB.SPECOBJECTSPointersEncoding.SPECOBJECT =
				append(specobjectsDB.SPECOBJECTSPointersEncoding.SPECOBJECT, int(specobjectAssocEnd_DB.ID))
		}

		query := backRepoSPECOBJECTS.db.Save(&specobjectsDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SPECOBJECTS intance %s", specobjects.Name))
		return err
	}

	return
}

// BackRepoSPECOBJECTS.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CheckoutPhaseOne() (Error error) {

	specobjectsDBArray := make([]SPECOBJECTSDB, 0)
	query := backRepoSPECOBJECTS.db.Find(&specobjectsDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	specobjectsInstancesToBeRemovedFromTheStage := make(map[*models.SPECOBJECTS]any)
	for key, value := range backRepoSPECOBJECTS.stage.SPECOBJECTSs {
		specobjectsInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, specobjectsDB := range specobjectsDBArray {
		backRepoSPECOBJECTS.CheckoutPhaseOneInstance(&specobjectsDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		specobjects, ok := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]
		if ok {
			delete(specobjectsInstancesToBeRemovedFromTheStage, specobjects)
		}
	}

	// remove from stage and back repo's 3 maps all specobjectss that are not in the checkout
	for specobjects := range specobjectsInstancesToBeRemovedFromTheStage {
		specobjects.Unstage(backRepoSPECOBJECTS.GetStage())

		// remove instance from the back repo 3 maps
		specobjectsID := backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]
		delete(backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID, specobjects)
		delete(backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB, specobjectsID)
		delete(backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr, specobjectsID)
	}

	return
}

// CheckoutPhaseOneInstance takes a specobjectsDB that has been found in the DB, updates the backRepo and stages the
// models version of the specobjectsDB
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CheckoutPhaseOneInstance(specobjectsDB *SPECOBJECTSDB) (Error error) {

	specobjects, ok := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]
	if !ok {
		specobjects = new(models.SPECOBJECTS)

		backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID] = specobjects
		backRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects] = specobjectsDB.ID

		// append model store with the new element
		specobjects.Name = specobjectsDB.Name_Data.String
		specobjects.Stage(backRepoSPECOBJECTS.GetStage())
	}
	specobjectsDB.CopyBasicFieldsToSPECOBJECTS(specobjects)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	specobjects.Stage(backRepoSPECOBJECTS.GetStage())

	// preserve pointer to specobjectsDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SPECOBJECTSDBID_SPECOBJECTSDB)[specobjectsDB hold variable pointers
	specobjectsDB_Data := *specobjectsDB
	preservedPtrToSPECOBJECTS := &specobjectsDB_Data
	backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[specobjectsDB.ID] = preservedPtrToSPECOBJECTS

	return
}

// BackRepoSPECOBJECTS.CheckoutPhaseTwo Checkouts all staged instances of SPECOBJECTS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, specobjectsDB := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB {
		backRepoSPECOBJECTS.CheckoutPhaseTwoInstance(backRepo, specobjectsDB)
	}
	return
}

// BackRepoSPECOBJECTS.CheckoutPhaseTwoInstance Checkouts staged instances of SPECOBJECTS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, specobjectsDB *SPECOBJECTSDB) (Error error) {

	specobjects := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr[specobjectsDB.ID]

	specobjectsDB.DecodePointers(backRepo, specobjects)

	return
}

func (specobjectsDB *SPECOBJECTSDB) DecodePointers(backRepo *BackRepoStruct, specobjects *models.SPECOBJECTS) {

	// insertion point for checkout of pointer encoding
	// This loop redeem specobjects.SPECOBJECT in the stage from the encode in the back repo
	// It parses all SPECOBJECTDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	specobjects.SPECOBJECT = specobjects.SPECOBJECT[:0]
	for _, _SPECOBJECTid := range specobjectsDB.SPECOBJECTSPointersEncoding.SPECOBJECT {
		specobjects.SPECOBJECT = append(specobjects.SPECOBJECT, backRepo.BackRepoSPECOBJECT.Map_SPECOBJECTDBID_SPECOBJECTPtr[uint(_SPECOBJECTid)])
	}

	return
}

// CommitSPECOBJECTS allows commit of a single specobjects (if already staged)
func (backRepo *BackRepoStruct) CommitSPECOBJECTS(specobjects *models.SPECOBJECTS) {
	backRepo.BackRepoSPECOBJECTS.CommitPhaseOneInstance(specobjects)
	if id, ok := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]; ok {
		backRepo.BackRepoSPECOBJECTS.CommitPhaseTwoInstance(backRepo, id, specobjects)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSPECOBJECTS allows checkout of a single specobjects (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSPECOBJECTS(specobjects *models.SPECOBJECTS) {
	// check if the specobjects is staged
	if _, ok := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]; ok {

		if id, ok := backRepo.BackRepoSPECOBJECTS.Map_SPECOBJECTSPtr_SPECOBJECTSDBID[specobjects]; ok {
			var specobjectsDB SPECOBJECTSDB
			specobjectsDB.ID = id

			if err := backRepo.BackRepoSPECOBJECTS.db.First(&specobjectsDB, id).Error; err != nil {
				log.Fatalln("CheckoutSPECOBJECTS : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSPECOBJECTS.CheckoutPhaseOneInstance(&specobjectsDB)
			backRepo.BackRepoSPECOBJECTS.CheckoutPhaseTwoInstance(backRepo, &specobjectsDB)
		}
	}
}

// CopyBasicFieldsFromSPECOBJECTS
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsFromSPECOBJECTS(specobjects *models.SPECOBJECTS) {
	// insertion point for fields commit

	specobjectsDB.Name_Data.String = specobjects.Name
	specobjectsDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromSPECOBJECTS_WOP
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsFromSPECOBJECTS_WOP(specobjects *models.SPECOBJECTS_WOP) {
	// insertion point for fields commit

	specobjectsDB.Name_Data.String = specobjects.Name
	specobjectsDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromSPECOBJECTSWOP
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsFromSPECOBJECTSWOP(specobjects *SPECOBJECTSWOP) {
	// insertion point for fields commit

	specobjectsDB.Name_Data.String = specobjects.Name
	specobjectsDB.Name_Data.Valid = true
}

// CopyBasicFieldsToSPECOBJECTS
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsToSPECOBJECTS(specobjects *models.SPECOBJECTS) {
	// insertion point for checkout of basic fields (back repo to stage)
	specobjects.Name = specobjectsDB.Name_Data.String
}

// CopyBasicFieldsToSPECOBJECTS_WOP
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsToSPECOBJECTS_WOP(specobjects *models.SPECOBJECTS_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	specobjects.Name = specobjectsDB.Name_Data.String
}

// CopyBasicFieldsToSPECOBJECTSWOP
func (specobjectsDB *SPECOBJECTSDB) CopyBasicFieldsToSPECOBJECTSWOP(specobjects *SPECOBJECTSWOP) {
	specobjects.ID = int(specobjectsDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	specobjects.Name = specobjectsDB.Name_Data.String
}

// Backup generates a json file from a slice of all SPECOBJECTSDB instances in the backrepo
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SPECOBJECTSDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPECOBJECTSDB, 0)
	for _, specobjectsDB := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB {
		forBackup = append(forBackup, specobjectsDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SPECOBJECTS ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SPECOBJECTS file", err.Error())
	}
}

// Backup generates a json file from a slice of all SPECOBJECTSDB instances in the backrepo
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPECOBJECTSDB, 0)
	for _, specobjectsDB := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB {
		forBackup = append(forBackup, specobjectsDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SPECOBJECTS")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SPECOBJECTS_Fields, -1)
	for _, specobjectsDB := range forBackup {

		var specobjectsWOP SPECOBJECTSWOP
		specobjectsDB.CopyBasicFieldsToSPECOBJECTSWOP(&specobjectsWOP)

		row := sh.AddRow()
		row.WriteStruct(&specobjectsWOP, -1)
	}
}

// RestoreXL from the "SPECOBJECTS" sheet all SPECOBJECTSDB instances
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSPECOBJECTSid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SPECOBJECTS"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSPECOBJECTS.rowVisitorSPECOBJECTS)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) rowVisitorSPECOBJECTS(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var specobjectsWOP SPECOBJECTSWOP
		row.ReadStruct(&specobjectsWOP)

		// add the unmarshalled struct to the stage
		specobjectsDB := new(SPECOBJECTSDB)
		specobjectsDB.CopyBasicFieldsFromSPECOBJECTSWOP(&specobjectsWOP)

		specobjectsDB_ID_atBackupTime := specobjectsDB.ID
		specobjectsDB.ID = 0
		query := backRepoSPECOBJECTS.db.Create(specobjectsDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[specobjectsDB.ID] = specobjectsDB
		BackRepoSPECOBJECTSid_atBckpTime_newID[specobjectsDB_ID_atBackupTime] = specobjectsDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SPECOBJECTSDB.json" in dirPath that stores an array
// of SPECOBJECTSDB and stores it in the database
// the map BackRepoSPECOBJECTSid_atBckpTime_newID is updated accordingly
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSPECOBJECTSid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SPECOBJECTSDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SPECOBJECTS file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SPECOBJECTSDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SPECOBJECTSDBID_SPECOBJECTSDB
	for _, specobjectsDB := range forRestore {

		specobjectsDB_ID_atBackupTime := specobjectsDB.ID
		specobjectsDB.ID = 0
		query := backRepoSPECOBJECTS.db.Create(specobjectsDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[specobjectsDB.ID] = specobjectsDB
		BackRepoSPECOBJECTSid_atBckpTime_newID[specobjectsDB_ID_atBackupTime] = specobjectsDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SPECOBJECTS file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SPECOBJECTS>id_atBckpTime_newID
// to compute new index
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) RestorePhaseTwo() {

	for _, specobjectsDB := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB {

		// next line of code is to avert unused variable compilation error
		_ = specobjectsDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSPECOBJECTS.db.Model(specobjectsDB).Updates(*specobjectsDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSPECOBJECTS.ResetReversePointers commits all staged instances of SPECOBJECTS to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, specobjects := range backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSPtr {
		backRepoSPECOBJECTS.ResetReversePointersInstance(backRepo, idx, specobjects)
	}

	return
}

func (backRepoSPECOBJECTS *BackRepoSPECOBJECTSStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, specobjects *models.SPECOBJECTS) (Error error) {

	// fetch matching specobjectsDB
	if specobjectsDB, ok := backRepoSPECOBJECTS.Map_SPECOBJECTSDBID_SPECOBJECTSDB[idx]; ok {
		_ = specobjectsDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSPECOBJECTSid_atBckpTime_newID map[uint]uint
