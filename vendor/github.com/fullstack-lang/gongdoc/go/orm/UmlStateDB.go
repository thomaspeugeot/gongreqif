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

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_UmlState_sql sql.NullBool
var dummy_UmlState_time time.Duration
var dummy_UmlState_sort sort.Float64Slice

// UmlStateAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model umlstateAPI
type UmlStateAPI struct {
	gorm.Model

	models.UmlState_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	UmlStatePointersEncoding UmlStatePointersEncoding
}

// UmlStatePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type UmlStatePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// UmlStateDB describes a umlstate in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model umlstateDB
type UmlStateDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field umlstateDB.Name
	Name_Data sql.NullString

	// Declation for basic field umlstateDB.X
	X_Data sql.NullFloat64

	// Declation for basic field umlstateDB.Y
	Y_Data sql.NullFloat64
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	UmlStatePointersEncoding
}

// UmlStateDBs arrays umlstateDBs
// swagger:response umlstateDBsResponse
type UmlStateDBs []UmlStateDB

// UmlStateDBResponse provides response
// swagger:response umlstateDBResponse
type UmlStateDBResponse struct {
	UmlStateDB
}

// UmlStateWOP is a UmlState without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type UmlStateWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	X float64 `xlsx:"2"`

	Y float64 `xlsx:"3"`
	// insertion for WOP pointer fields
}

var UmlState_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"X",
	"Y",
}

type BackRepoUmlStateStruct struct {
	// stores UmlStateDB according to their gorm ID
	Map_UmlStateDBID_UmlStateDB map[uint]*UmlStateDB

	// stores UmlStateDB ID according to UmlState address
	Map_UmlStatePtr_UmlStateDBID map[*models.UmlState]uint

	// stores UmlState according to their gorm ID
	Map_UmlStateDBID_UmlStatePtr map[uint]*models.UmlState

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoUmlState *BackRepoUmlStateStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoUmlState.stage
	return
}

func (backRepoUmlState *BackRepoUmlStateStruct) GetDB() *gorm.DB {
	return backRepoUmlState.db
}

// GetUmlStateDBFromUmlStatePtr is a handy function to access the back repo instance from the stage instance
func (backRepoUmlState *BackRepoUmlStateStruct) GetUmlStateDBFromUmlStatePtr(umlstate *models.UmlState) (umlstateDB *UmlStateDB) {
	id := backRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]
	umlstateDB = backRepoUmlState.Map_UmlStateDBID_UmlStateDB[id]
	return
}

// BackRepoUmlState.CommitPhaseOne commits all staged instances of UmlState to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlState *BackRepoUmlStateStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for umlstate := range stage.UmlStates {
		backRepoUmlState.CommitPhaseOneInstance(umlstate)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, umlstate := range backRepoUmlState.Map_UmlStateDBID_UmlStatePtr {
		if _, ok := stage.UmlStates[umlstate]; !ok {
			backRepoUmlState.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUmlState.CommitDeleteInstance commits deletion of UmlState to the BackRepo
func (backRepoUmlState *BackRepoUmlStateStruct) CommitDeleteInstance(id uint) (Error error) {

	umlstate := backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[id]

	// umlstate is not staged anymore, remove umlstateDB
	umlstateDB := backRepoUmlState.Map_UmlStateDBID_UmlStateDB[id]
	query := backRepoUmlState.db.Unscoped().Delete(&umlstateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoUmlState.Map_UmlStatePtr_UmlStateDBID, umlstate)
	delete(backRepoUmlState.Map_UmlStateDBID_UmlStatePtr, id)
	delete(backRepoUmlState.Map_UmlStateDBID_UmlStateDB, id)

	return
}

// BackRepoUmlState.CommitPhaseOneInstance commits umlstate staged instances of UmlState to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUmlState *BackRepoUmlStateStruct) CommitPhaseOneInstance(umlstate *models.UmlState) (Error error) {

	// check if the umlstate is not commited yet
	if _, ok := backRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]; ok {
		return
	}

	// initiate umlstate
	var umlstateDB UmlStateDB
	umlstateDB.CopyBasicFieldsFromUmlState(umlstate)

	query := backRepoUmlState.db.Create(&umlstateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate] = umlstateDB.ID
	backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID] = umlstate
	backRepoUmlState.Map_UmlStateDBID_UmlStateDB[umlstateDB.ID] = &umlstateDB

	return
}

// BackRepoUmlState.CommitPhaseTwo commits all staged instances of UmlState to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlState *BackRepoUmlStateStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, umlstate := range backRepoUmlState.Map_UmlStateDBID_UmlStatePtr {
		backRepoUmlState.CommitPhaseTwoInstance(backRepo, idx, umlstate)
	}

	return
}

// BackRepoUmlState.CommitPhaseTwoInstance commits {{structname }} of models.UmlState to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlState *BackRepoUmlStateStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, umlstate *models.UmlState) (Error error) {

	// fetch matching umlstateDB
	if umlstateDB, ok := backRepoUmlState.Map_UmlStateDBID_UmlStateDB[idx]; ok {

		umlstateDB.CopyBasicFieldsFromUmlState(umlstate)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoUmlState.db.Save(&umlstateDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown UmlState intance %s", umlstate.Name))
		return err
	}

	return
}

// BackRepoUmlState.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoUmlState *BackRepoUmlStateStruct) CheckoutPhaseOne() (Error error) {

	umlstateDBArray := make([]UmlStateDB, 0)
	query := backRepoUmlState.db.Find(&umlstateDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	umlstateInstancesToBeRemovedFromTheStage := make(map[*models.UmlState]any)
	for key, value := range backRepoUmlState.stage.UmlStates {
		umlstateInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, umlstateDB := range umlstateDBArray {
		backRepoUmlState.CheckoutPhaseOneInstance(&umlstateDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		umlstate, ok := backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]
		if ok {
			delete(umlstateInstancesToBeRemovedFromTheStage, umlstate)
		}
	}

	// remove from stage and back repo's 3 maps all umlstates that are not in the checkout
	for umlstate := range umlstateInstancesToBeRemovedFromTheStage {
		umlstate.Unstage(backRepoUmlState.GetStage())

		// remove instance from the back repo 3 maps
		umlstateID := backRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]
		delete(backRepoUmlState.Map_UmlStatePtr_UmlStateDBID, umlstate)
		delete(backRepoUmlState.Map_UmlStateDBID_UmlStateDB, umlstateID)
		delete(backRepoUmlState.Map_UmlStateDBID_UmlStatePtr, umlstateID)
	}

	return
}

// CheckoutPhaseOneInstance takes a umlstateDB that has been found in the DB, updates the backRepo and stages the
// models version of the umlstateDB
func (backRepoUmlState *BackRepoUmlStateStruct) CheckoutPhaseOneInstance(umlstateDB *UmlStateDB) (Error error) {

	umlstate, ok := backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]
	if !ok {
		umlstate = new(models.UmlState)

		backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID] = umlstate
		backRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate] = umlstateDB.ID

		// append model store with the new element
		umlstate.Name = umlstateDB.Name_Data.String
		umlstate.Stage(backRepoUmlState.GetStage())
	}
	umlstateDB.CopyBasicFieldsToUmlState(umlstate)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	umlstate.Stage(backRepoUmlState.GetStage())

	// preserve pointer to umlstateDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_UmlStateDBID_UmlStateDB)[umlstateDB hold variable pointers
	umlstateDB_Data := *umlstateDB
	preservedPtrToUmlState := &umlstateDB_Data
	backRepoUmlState.Map_UmlStateDBID_UmlStateDB[umlstateDB.ID] = preservedPtrToUmlState

	return
}

// BackRepoUmlState.CheckoutPhaseTwo Checkouts all staged instances of UmlState to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlState *BackRepoUmlStateStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, umlstateDB := range backRepoUmlState.Map_UmlStateDBID_UmlStateDB {
		backRepoUmlState.CheckoutPhaseTwoInstance(backRepo, umlstateDB)
	}
	return
}

// BackRepoUmlState.CheckoutPhaseTwoInstance Checkouts staged instances of UmlState to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlState *BackRepoUmlStateStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, umlstateDB *UmlStateDB) (Error error) {

	umlstate := backRepoUmlState.Map_UmlStateDBID_UmlStatePtr[umlstateDB.ID]

	umlstateDB.DecodePointers(backRepo, umlstate)

	return
}

func (umlstateDB *UmlStateDB) DecodePointers(backRepo *BackRepoStruct, umlstate *models.UmlState) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitUmlState allows commit of a single umlstate (if already staged)
func (backRepo *BackRepoStruct) CommitUmlState(umlstate *models.UmlState) {
	backRepo.BackRepoUmlState.CommitPhaseOneInstance(umlstate)
	if id, ok := backRepo.BackRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]; ok {
		backRepo.BackRepoUmlState.CommitPhaseTwoInstance(backRepo, id, umlstate)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitUmlState allows checkout of a single umlstate (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUmlState(umlstate *models.UmlState) {
	// check if the umlstate is staged
	if _, ok := backRepo.BackRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]; ok {

		if id, ok := backRepo.BackRepoUmlState.Map_UmlStatePtr_UmlStateDBID[umlstate]; ok {
			var umlstateDB UmlStateDB
			umlstateDB.ID = id

			if err := backRepo.BackRepoUmlState.db.First(&umlstateDB, id).Error; err != nil {
				log.Fatalln("CheckoutUmlState : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUmlState.CheckoutPhaseOneInstance(&umlstateDB)
			backRepo.BackRepoUmlState.CheckoutPhaseTwoInstance(backRepo, &umlstateDB)
		}
	}
}

// CopyBasicFieldsFromUmlState
func (umlstateDB *UmlStateDB) CopyBasicFieldsFromUmlState(umlstate *models.UmlState) {
	// insertion point for fields commit

	umlstateDB.Name_Data.String = umlstate.Name
	umlstateDB.Name_Data.Valid = true

	umlstateDB.X_Data.Float64 = umlstate.X
	umlstateDB.X_Data.Valid = true

	umlstateDB.Y_Data.Float64 = umlstate.Y
	umlstateDB.Y_Data.Valid = true
}

// CopyBasicFieldsFromUmlState_WOP
func (umlstateDB *UmlStateDB) CopyBasicFieldsFromUmlState_WOP(umlstate *models.UmlState_WOP) {
	// insertion point for fields commit

	umlstateDB.Name_Data.String = umlstate.Name
	umlstateDB.Name_Data.Valid = true

	umlstateDB.X_Data.Float64 = umlstate.X
	umlstateDB.X_Data.Valid = true

	umlstateDB.Y_Data.Float64 = umlstate.Y
	umlstateDB.Y_Data.Valid = true
}

// CopyBasicFieldsFromUmlStateWOP
func (umlstateDB *UmlStateDB) CopyBasicFieldsFromUmlStateWOP(umlstate *UmlStateWOP) {
	// insertion point for fields commit

	umlstateDB.Name_Data.String = umlstate.Name
	umlstateDB.Name_Data.Valid = true

	umlstateDB.X_Data.Float64 = umlstate.X
	umlstateDB.X_Data.Valid = true

	umlstateDB.Y_Data.Float64 = umlstate.Y
	umlstateDB.Y_Data.Valid = true
}

// CopyBasicFieldsToUmlState
func (umlstateDB *UmlStateDB) CopyBasicFieldsToUmlState(umlstate *models.UmlState) {
	// insertion point for checkout of basic fields (back repo to stage)
	umlstate.Name = umlstateDB.Name_Data.String
	umlstate.X = umlstateDB.X_Data.Float64
	umlstate.Y = umlstateDB.Y_Data.Float64
}

// CopyBasicFieldsToUmlState_WOP
func (umlstateDB *UmlStateDB) CopyBasicFieldsToUmlState_WOP(umlstate *models.UmlState_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	umlstate.Name = umlstateDB.Name_Data.String
	umlstate.X = umlstateDB.X_Data.Float64
	umlstate.Y = umlstateDB.Y_Data.Float64
}

// CopyBasicFieldsToUmlStateWOP
func (umlstateDB *UmlStateDB) CopyBasicFieldsToUmlStateWOP(umlstate *UmlStateWOP) {
	umlstate.ID = int(umlstateDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	umlstate.Name = umlstateDB.Name_Data.String
	umlstate.X = umlstateDB.X_Data.Float64
	umlstate.Y = umlstateDB.Y_Data.Float64
}

// Backup generates a json file from a slice of all UmlStateDB instances in the backrepo
func (backRepoUmlState *BackRepoUmlStateStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "UmlStateDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UmlStateDB, 0)
	for _, umlstateDB := range backRepoUmlState.Map_UmlStateDBID_UmlStateDB {
		forBackup = append(forBackup, umlstateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json UmlState ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json UmlState file", err.Error())
	}
}

// Backup generates a json file from a slice of all UmlStateDB instances in the backrepo
func (backRepoUmlState *BackRepoUmlStateStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UmlStateDB, 0)
	for _, umlstateDB := range backRepoUmlState.Map_UmlStateDBID_UmlStateDB {
		forBackup = append(forBackup, umlstateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("UmlState")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&UmlState_Fields, -1)
	for _, umlstateDB := range forBackup {

		var umlstateWOP UmlStateWOP
		umlstateDB.CopyBasicFieldsToUmlStateWOP(&umlstateWOP)

		row := sh.AddRow()
		row.WriteStruct(&umlstateWOP, -1)
	}
}

// RestoreXL from the "UmlState" sheet all UmlStateDB instances
func (backRepoUmlState *BackRepoUmlStateStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoUmlStateid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["UmlState"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoUmlState.rowVisitorUmlState)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoUmlState *BackRepoUmlStateStruct) rowVisitorUmlState(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var umlstateWOP UmlStateWOP
		row.ReadStruct(&umlstateWOP)

		// add the unmarshalled struct to the stage
		umlstateDB := new(UmlStateDB)
		umlstateDB.CopyBasicFieldsFromUmlStateWOP(&umlstateWOP)

		umlstateDB_ID_atBackupTime := umlstateDB.ID
		umlstateDB.ID = 0
		query := backRepoUmlState.db.Create(umlstateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoUmlState.Map_UmlStateDBID_UmlStateDB[umlstateDB.ID] = umlstateDB
		BackRepoUmlStateid_atBckpTime_newID[umlstateDB_ID_atBackupTime] = umlstateDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "UmlStateDB.json" in dirPath that stores an array
// of UmlStateDB and stores it in the database
// the map BackRepoUmlStateid_atBckpTime_newID is updated accordingly
func (backRepoUmlState *BackRepoUmlStateStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoUmlStateid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "UmlStateDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json UmlState file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*UmlStateDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_UmlStateDBID_UmlStateDB
	for _, umlstateDB := range forRestore {

		umlstateDB_ID_atBackupTime := umlstateDB.ID
		umlstateDB.ID = 0
		query := backRepoUmlState.db.Create(umlstateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoUmlState.Map_UmlStateDBID_UmlStateDB[umlstateDB.ID] = umlstateDB
		BackRepoUmlStateid_atBckpTime_newID[umlstateDB_ID_atBackupTime] = umlstateDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json UmlState file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<UmlState>id_atBckpTime_newID
// to compute new index
func (backRepoUmlState *BackRepoUmlStateStruct) RestorePhaseTwo() {

	for _, umlstateDB := range backRepoUmlState.Map_UmlStateDBID_UmlStateDB {

		// next line of code is to avert unused variable compilation error
		_ = umlstateDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoUmlState.db.Model(umlstateDB).Updates(*umlstateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoUmlState.ResetReversePointers commits all staged instances of UmlState to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUmlState *BackRepoUmlStateStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, umlstate := range backRepoUmlState.Map_UmlStateDBID_UmlStatePtr {
		backRepoUmlState.ResetReversePointersInstance(backRepo, idx, umlstate)
	}

	return
}

func (backRepoUmlState *BackRepoUmlStateStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, umlstate *models.UmlState) (Error error) {

	// fetch matching umlstateDB
	if umlstateDB, ok := backRepoUmlState.Map_UmlStateDBID_UmlStateDB[idx]; ok {
		_ = umlstateDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoUmlStateid_atBckpTime_newID map[uint]uint
