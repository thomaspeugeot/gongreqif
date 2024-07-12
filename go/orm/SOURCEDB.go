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
var dummy_SOURCE_sql sql.NullBool
var dummy_SOURCE_time time.Duration
var dummy_SOURCE_sort sort.Float64Slice

// SOURCEAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model sourceAPI
type SOURCEAPI struct {
	gorm.Model

	models.SOURCE_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SOURCEPointersEncoding SOURCEPointersEncoding
}

// SOURCEPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SOURCEPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// SOURCEDB describes a source in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model sourceDB
type SOURCEDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field sourceDB.Name
	Name_Data sql.NullString

	// Declation for basic field sourceDB.SPECOBJECTREF
	SPECOBJECTREF_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SOURCEPointersEncoding
}

// SOURCEDBs arrays sourceDBs
// swagger:response sourceDBsResponse
type SOURCEDBs []SOURCEDB

// SOURCEDBResponse provides response
// swagger:response sourceDBResponse
type SOURCEDBResponse struct {
	SOURCEDB
}

// SOURCEWOP is a SOURCE without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SOURCEWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	SPECOBJECTREF string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var SOURCE_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"SPECOBJECTREF",
}

type BackRepoSOURCEStruct struct {
	// stores SOURCEDB according to their gorm ID
	Map_SOURCEDBID_SOURCEDB map[uint]*SOURCEDB

	// stores SOURCEDB ID according to SOURCE address
	Map_SOURCEPtr_SOURCEDBID map[*models.SOURCE]uint

	// stores SOURCE according to their gorm ID
	Map_SOURCEDBID_SOURCEPtr map[uint]*models.SOURCE

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSOURCE *BackRepoSOURCEStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSOURCE.stage
	return
}

func (backRepoSOURCE *BackRepoSOURCEStruct) GetDB() *gorm.DB {
	return backRepoSOURCE.db
}

// GetSOURCEDBFromSOURCEPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSOURCE *BackRepoSOURCEStruct) GetSOURCEDBFromSOURCEPtr(source *models.SOURCE) (sourceDB *SOURCEDB) {
	id := backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]
	sourceDB = backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[id]
	return
}

// BackRepoSOURCE.CommitPhaseOne commits all staged instances of SOURCE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSOURCE *BackRepoSOURCEStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for source := range stage.SOURCEs {
		backRepoSOURCE.CommitPhaseOneInstance(source)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, source := range backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr {
		if _, ok := stage.SOURCEs[source]; !ok {
			backRepoSOURCE.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSOURCE.CommitDeleteInstance commits deletion of SOURCE to the BackRepo
func (backRepoSOURCE *BackRepoSOURCEStruct) CommitDeleteInstance(id uint) (Error error) {

	source := backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[id]

	// source is not staged anymore, remove sourceDB
	sourceDB := backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[id]
	query := backRepoSOURCE.db.Unscoped().Delete(&sourceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID, source)
	delete(backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr, id)
	delete(backRepoSOURCE.Map_SOURCEDBID_SOURCEDB, id)

	return
}

// BackRepoSOURCE.CommitPhaseOneInstance commits source staged instances of SOURCE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSOURCE *BackRepoSOURCEStruct) CommitPhaseOneInstance(source *models.SOURCE) (Error error) {

	// check if the source is not commited yet
	if _, ok := backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]; ok {
		return
	}

	// initiate source
	var sourceDB SOURCEDB
	sourceDB.CopyBasicFieldsFromSOURCE(source)

	query := backRepoSOURCE.db.Create(&sourceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source] = sourceDB.ID
	backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID] = source
	backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[sourceDB.ID] = &sourceDB

	return
}

// BackRepoSOURCE.CommitPhaseTwo commits all staged instances of SOURCE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSOURCE *BackRepoSOURCEStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, source := range backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr {
		backRepoSOURCE.CommitPhaseTwoInstance(backRepo, idx, source)
	}

	return
}

// BackRepoSOURCE.CommitPhaseTwoInstance commits {{structname }} of models.SOURCE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSOURCE *BackRepoSOURCEStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, source *models.SOURCE) (Error error) {

	// fetch matching sourceDB
	if sourceDB, ok := backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[idx]; ok {

		sourceDB.CopyBasicFieldsFromSOURCE(source)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoSOURCE.db.Save(&sourceDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SOURCE intance %s", source.Name))
		return err
	}

	return
}

// BackRepoSOURCE.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSOURCE *BackRepoSOURCEStruct) CheckoutPhaseOne() (Error error) {

	sourceDBArray := make([]SOURCEDB, 0)
	query := backRepoSOURCE.db.Find(&sourceDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	sourceInstancesToBeRemovedFromTheStage := make(map[*models.SOURCE]any)
	for key, value := range backRepoSOURCE.stage.SOURCEs {
		sourceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, sourceDB := range sourceDBArray {
		backRepoSOURCE.CheckoutPhaseOneInstance(&sourceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		source, ok := backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]
		if ok {
			delete(sourceInstancesToBeRemovedFromTheStage, source)
		}
	}

	// remove from stage and back repo's 3 maps all sources that are not in the checkout
	for source := range sourceInstancesToBeRemovedFromTheStage {
		source.Unstage(backRepoSOURCE.GetStage())

		// remove instance from the back repo 3 maps
		sourceID := backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]
		delete(backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID, source)
		delete(backRepoSOURCE.Map_SOURCEDBID_SOURCEDB, sourceID)
		delete(backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr, sourceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a sourceDB that has been found in the DB, updates the backRepo and stages the
// models version of the sourceDB
func (backRepoSOURCE *BackRepoSOURCEStruct) CheckoutPhaseOneInstance(sourceDB *SOURCEDB) (Error error) {

	source, ok := backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]
	if !ok {
		source = new(models.SOURCE)

		backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID] = source
		backRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source] = sourceDB.ID

		// append model store with the new element
		source.Name = sourceDB.Name_Data.String
		source.Stage(backRepoSOURCE.GetStage())
	}
	sourceDB.CopyBasicFieldsToSOURCE(source)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	source.Stage(backRepoSOURCE.GetStage())

	// preserve pointer to sourceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SOURCEDBID_SOURCEDB)[sourceDB hold variable pointers
	sourceDB_Data := *sourceDB
	preservedPtrToSOURCE := &sourceDB_Data
	backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[sourceDB.ID] = preservedPtrToSOURCE

	return
}

// BackRepoSOURCE.CheckoutPhaseTwo Checkouts all staged instances of SOURCE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSOURCE *BackRepoSOURCEStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, sourceDB := range backRepoSOURCE.Map_SOURCEDBID_SOURCEDB {
		backRepoSOURCE.CheckoutPhaseTwoInstance(backRepo, sourceDB)
	}
	return
}

// BackRepoSOURCE.CheckoutPhaseTwoInstance Checkouts staged instances of SOURCE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSOURCE *BackRepoSOURCEStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, sourceDB *SOURCEDB) (Error error) {

	source := backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr[sourceDB.ID]

	sourceDB.DecodePointers(backRepo, source)

	return
}

func (sourceDB *SOURCEDB) DecodePointers(backRepo *BackRepoStruct, source *models.SOURCE) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitSOURCE allows commit of a single source (if already staged)
func (backRepo *BackRepoStruct) CommitSOURCE(source *models.SOURCE) {
	backRepo.BackRepoSOURCE.CommitPhaseOneInstance(source)
	if id, ok := backRepo.BackRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]; ok {
		backRepo.BackRepoSOURCE.CommitPhaseTwoInstance(backRepo, id, source)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSOURCE allows checkout of a single source (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSOURCE(source *models.SOURCE) {
	// check if the source is staged
	if _, ok := backRepo.BackRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]; ok {

		if id, ok := backRepo.BackRepoSOURCE.Map_SOURCEPtr_SOURCEDBID[source]; ok {
			var sourceDB SOURCEDB
			sourceDB.ID = id

			if err := backRepo.BackRepoSOURCE.db.First(&sourceDB, id).Error; err != nil {
				log.Fatalln("CheckoutSOURCE : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSOURCE.CheckoutPhaseOneInstance(&sourceDB)
			backRepo.BackRepoSOURCE.CheckoutPhaseTwoInstance(backRepo, &sourceDB)
		}
	}
}

// CopyBasicFieldsFromSOURCE
func (sourceDB *SOURCEDB) CopyBasicFieldsFromSOURCE(source *models.SOURCE) {
	// insertion point for fields commit

	sourceDB.Name_Data.String = source.Name
	sourceDB.Name_Data.Valid = true

	sourceDB.SPECOBJECTREF_Data.String = source.SPECOBJECTREF
	sourceDB.SPECOBJECTREF_Data.Valid = true
}

// CopyBasicFieldsFromSOURCE_WOP
func (sourceDB *SOURCEDB) CopyBasicFieldsFromSOURCE_WOP(source *models.SOURCE_WOP) {
	// insertion point for fields commit

	sourceDB.Name_Data.String = source.Name
	sourceDB.Name_Data.Valid = true

	sourceDB.SPECOBJECTREF_Data.String = source.SPECOBJECTREF
	sourceDB.SPECOBJECTREF_Data.Valid = true
}

// CopyBasicFieldsFromSOURCEWOP
func (sourceDB *SOURCEDB) CopyBasicFieldsFromSOURCEWOP(source *SOURCEWOP) {
	// insertion point for fields commit

	sourceDB.Name_Data.String = source.Name
	sourceDB.Name_Data.Valid = true

	sourceDB.SPECOBJECTREF_Data.String = source.SPECOBJECTREF
	sourceDB.SPECOBJECTREF_Data.Valid = true
}

// CopyBasicFieldsToSOURCE
func (sourceDB *SOURCEDB) CopyBasicFieldsToSOURCE(source *models.SOURCE) {
	// insertion point for checkout of basic fields (back repo to stage)
	source.Name = sourceDB.Name_Data.String
	source.SPECOBJECTREF = sourceDB.SPECOBJECTREF_Data.String
}

// CopyBasicFieldsToSOURCE_WOP
func (sourceDB *SOURCEDB) CopyBasicFieldsToSOURCE_WOP(source *models.SOURCE_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	source.Name = sourceDB.Name_Data.String
	source.SPECOBJECTREF = sourceDB.SPECOBJECTREF_Data.String
}

// CopyBasicFieldsToSOURCEWOP
func (sourceDB *SOURCEDB) CopyBasicFieldsToSOURCEWOP(source *SOURCEWOP) {
	source.ID = int(sourceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	source.Name = sourceDB.Name_Data.String
	source.SPECOBJECTREF = sourceDB.SPECOBJECTREF_Data.String
}

// Backup generates a json file from a slice of all SOURCEDB instances in the backrepo
func (backRepoSOURCE *BackRepoSOURCEStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SOURCEDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SOURCEDB, 0)
	for _, sourceDB := range backRepoSOURCE.Map_SOURCEDBID_SOURCEDB {
		forBackup = append(forBackup, sourceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SOURCE ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SOURCE file", err.Error())
	}
}

// Backup generates a json file from a slice of all SOURCEDB instances in the backrepo
func (backRepoSOURCE *BackRepoSOURCEStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SOURCEDB, 0)
	for _, sourceDB := range backRepoSOURCE.Map_SOURCEDBID_SOURCEDB {
		forBackup = append(forBackup, sourceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SOURCE")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SOURCE_Fields, -1)
	for _, sourceDB := range forBackup {

		var sourceWOP SOURCEWOP
		sourceDB.CopyBasicFieldsToSOURCEWOP(&sourceWOP)

		row := sh.AddRow()
		row.WriteStruct(&sourceWOP, -1)
	}
}

// RestoreXL from the "SOURCE" sheet all SOURCEDB instances
func (backRepoSOURCE *BackRepoSOURCEStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSOURCEid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SOURCE"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSOURCE.rowVisitorSOURCE)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSOURCE *BackRepoSOURCEStruct) rowVisitorSOURCE(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var sourceWOP SOURCEWOP
		row.ReadStruct(&sourceWOP)

		// add the unmarshalled struct to the stage
		sourceDB := new(SOURCEDB)
		sourceDB.CopyBasicFieldsFromSOURCEWOP(&sourceWOP)

		sourceDB_ID_atBackupTime := sourceDB.ID
		sourceDB.ID = 0
		query := backRepoSOURCE.db.Create(sourceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[sourceDB.ID] = sourceDB
		BackRepoSOURCEid_atBckpTime_newID[sourceDB_ID_atBackupTime] = sourceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SOURCEDB.json" in dirPath that stores an array
// of SOURCEDB and stores it in the database
// the map BackRepoSOURCEid_atBckpTime_newID is updated accordingly
func (backRepoSOURCE *BackRepoSOURCEStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSOURCEid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SOURCEDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SOURCE file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SOURCEDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SOURCEDBID_SOURCEDB
	for _, sourceDB := range forRestore {

		sourceDB_ID_atBackupTime := sourceDB.ID
		sourceDB.ID = 0
		query := backRepoSOURCE.db.Create(sourceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[sourceDB.ID] = sourceDB
		BackRepoSOURCEid_atBckpTime_newID[sourceDB_ID_atBackupTime] = sourceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SOURCE file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SOURCE>id_atBckpTime_newID
// to compute new index
func (backRepoSOURCE *BackRepoSOURCEStruct) RestorePhaseTwo() {

	for _, sourceDB := range backRepoSOURCE.Map_SOURCEDBID_SOURCEDB {

		// next line of code is to avert unused variable compilation error
		_ = sourceDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSOURCE.db.Model(sourceDB).Updates(*sourceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSOURCE.ResetReversePointers commits all staged instances of SOURCE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSOURCE *BackRepoSOURCEStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, source := range backRepoSOURCE.Map_SOURCEDBID_SOURCEPtr {
		backRepoSOURCE.ResetReversePointersInstance(backRepo, idx, source)
	}

	return
}

func (backRepoSOURCE *BackRepoSOURCEStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, source *models.SOURCE) (Error error) {

	// fetch matching sourceDB
	if sourceDB, ok := backRepoSOURCE.Map_SOURCEDBID_SOURCEDB[idx]; ok {
		_ = sourceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSOURCEid_atBckpTime_newID map[uint]uint
