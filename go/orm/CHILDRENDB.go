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
var dummy_CHILDREN_sql sql.NullBool
var dummy_CHILDREN_time time.Duration
var dummy_CHILDREN_sort sort.Float64Slice

// CHILDRENAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model childrenAPI
type CHILDRENAPI struct {
	gorm.Model

	models.CHILDREN_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	CHILDRENPointersEncoding CHILDRENPointersEncoding
}

// CHILDRENPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type CHILDRENPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field SPECHIERARCHY is a slice of pointers to another Struct (optional or 0..1)
	SPECHIERARCHY IntSlice `gorm:"type:TEXT"`
}

// CHILDRENDB describes a children in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model childrenDB
type CHILDRENDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field childrenDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	CHILDRENPointersEncoding
}

// CHILDRENDBs arrays childrenDBs
// swagger:response childrenDBsResponse
type CHILDRENDBs []CHILDRENDB

// CHILDRENDBResponse provides response
// swagger:response childrenDBResponse
type CHILDRENDBResponse struct {
	CHILDRENDB
}

// CHILDRENWOP is a CHILDREN without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type CHILDRENWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var CHILDREN_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoCHILDRENStruct struct {
	// stores CHILDRENDB according to their gorm ID
	Map_CHILDRENDBID_CHILDRENDB map[uint]*CHILDRENDB

	// stores CHILDRENDB ID according to CHILDREN address
	Map_CHILDRENPtr_CHILDRENDBID map[*models.CHILDREN]uint

	// stores CHILDREN according to their gorm ID
	Map_CHILDRENDBID_CHILDRENPtr map[uint]*models.CHILDREN

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoCHILDREN *BackRepoCHILDRENStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoCHILDREN.stage
	return
}

func (backRepoCHILDREN *BackRepoCHILDRENStruct) GetDB() *gorm.DB {
	return backRepoCHILDREN.db
}

// GetCHILDRENDBFromCHILDRENPtr is a handy function to access the back repo instance from the stage instance
func (backRepoCHILDREN *BackRepoCHILDRENStruct) GetCHILDRENDBFromCHILDRENPtr(children *models.CHILDREN) (childrenDB *CHILDRENDB) {
	id := backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]
	childrenDB = backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[id]
	return
}

// BackRepoCHILDREN.CommitPhaseOne commits all staged instances of CHILDREN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for children := range stage.CHILDRENs {
		backRepoCHILDREN.CommitPhaseOneInstance(children)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, children := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr {
		if _, ok := stage.CHILDRENs[children]; !ok {
			backRepoCHILDREN.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoCHILDREN.CommitDeleteInstance commits deletion of CHILDREN to the BackRepo
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CommitDeleteInstance(id uint) (Error error) {

	children := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[id]

	// children is not staged anymore, remove childrenDB
	childrenDB := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[id]
	query := backRepoCHILDREN.db.Unscoped().Delete(&childrenDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID, children)
	delete(backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr, id)
	delete(backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB, id)

	return
}

// BackRepoCHILDREN.CommitPhaseOneInstance commits children staged instances of CHILDREN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CommitPhaseOneInstance(children *models.CHILDREN) (Error error) {

	// check if the children is not commited yet
	if _, ok := backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]; ok {
		return
	}

	// initiate children
	var childrenDB CHILDRENDB
	childrenDB.CopyBasicFieldsFromCHILDREN(children)

	query := backRepoCHILDREN.db.Create(&childrenDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children] = childrenDB.ID
	backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID] = children
	backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[childrenDB.ID] = &childrenDB

	return
}

// BackRepoCHILDREN.CommitPhaseTwo commits all staged instances of CHILDREN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, children := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr {
		backRepoCHILDREN.CommitPhaseTwoInstance(backRepo, idx, children)
	}

	return
}

// BackRepoCHILDREN.CommitPhaseTwoInstance commits {{structname }} of models.CHILDREN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, children *models.CHILDREN) (Error error) {

	// fetch matching childrenDB
	if childrenDB, ok := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[idx]; ok {

		childrenDB.CopyBasicFieldsFromCHILDREN(children)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		childrenDB.CHILDRENPointersEncoding.SPECHIERARCHY = make([]int, 0)
		// 2. encode
		for _, spechierarchyAssocEnd := range children.SPECHIERARCHY {
			spechierarchyAssocEnd_DB :=
				backRepo.BackRepoSPECHIERARCHY.GetSPECHIERARCHYDBFromSPECHIERARCHYPtr(spechierarchyAssocEnd)
			
			// the stage might be inconsistant, meaning that the spechierarchyAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if spechierarchyAssocEnd_DB == nil {
				continue
			}
			
			childrenDB.CHILDRENPointersEncoding.SPECHIERARCHY =
				append(childrenDB.CHILDRENPointersEncoding.SPECHIERARCHY, int(spechierarchyAssocEnd_DB.ID))
		}

		query := backRepoCHILDREN.db.Save(&childrenDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown CHILDREN intance %s", children.Name))
		return err
	}

	return
}

// BackRepoCHILDREN.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CheckoutPhaseOne() (Error error) {

	childrenDBArray := make([]CHILDRENDB, 0)
	query := backRepoCHILDREN.db.Find(&childrenDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	childrenInstancesToBeRemovedFromTheStage := make(map[*models.CHILDREN]any)
	for key, value := range backRepoCHILDREN.stage.CHILDRENs {
		childrenInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, childrenDB := range childrenDBArray {
		backRepoCHILDREN.CheckoutPhaseOneInstance(&childrenDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		children, ok := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]
		if ok {
			delete(childrenInstancesToBeRemovedFromTheStage, children)
		}
	}

	// remove from stage and back repo's 3 maps all childrens that are not in the checkout
	for children := range childrenInstancesToBeRemovedFromTheStage {
		children.Unstage(backRepoCHILDREN.GetStage())

		// remove instance from the back repo 3 maps
		childrenID := backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]
		delete(backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID, children)
		delete(backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB, childrenID)
		delete(backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr, childrenID)
	}

	return
}

// CheckoutPhaseOneInstance takes a childrenDB that has been found in the DB, updates the backRepo and stages the
// models version of the childrenDB
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CheckoutPhaseOneInstance(childrenDB *CHILDRENDB) (Error error) {

	children, ok := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]
	if !ok {
		children = new(models.CHILDREN)

		backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID] = children
		backRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children] = childrenDB.ID

		// append model store with the new element
		children.Name = childrenDB.Name_Data.String
		children.Stage(backRepoCHILDREN.GetStage())
	}
	childrenDB.CopyBasicFieldsToCHILDREN(children)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	children.Stage(backRepoCHILDREN.GetStage())

	// preserve pointer to childrenDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_CHILDRENDBID_CHILDRENDB)[childrenDB hold variable pointers
	childrenDB_Data := *childrenDB
	preservedPtrToCHILDREN := &childrenDB_Data
	backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[childrenDB.ID] = preservedPtrToCHILDREN

	return
}

// BackRepoCHILDREN.CheckoutPhaseTwo Checkouts all staged instances of CHILDREN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, childrenDB := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB {
		backRepoCHILDREN.CheckoutPhaseTwoInstance(backRepo, childrenDB)
	}
	return
}

// BackRepoCHILDREN.CheckoutPhaseTwoInstance Checkouts staged instances of CHILDREN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCHILDREN *BackRepoCHILDRENStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, childrenDB *CHILDRENDB) (Error error) {

	children := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[childrenDB.ID]

	childrenDB.DecodePointers(backRepo, children)

	return
}

func (childrenDB *CHILDRENDB) DecodePointers(backRepo *BackRepoStruct, children *models.CHILDREN) {

	// insertion point for checkout of pointer encoding
	// This loop redeem children.SPECHIERARCHY in the stage from the encode in the back repo
	// It parses all SPECHIERARCHYDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	children.SPECHIERARCHY = children.SPECHIERARCHY[:0]
	for _, _SPECHIERARCHYid := range childrenDB.CHILDRENPointersEncoding.SPECHIERARCHY {
		children.SPECHIERARCHY = append(children.SPECHIERARCHY, backRepo.BackRepoSPECHIERARCHY.Map_SPECHIERARCHYDBID_SPECHIERARCHYPtr[uint(_SPECHIERARCHYid)])
	}

	return
}

// CommitCHILDREN allows commit of a single children (if already staged)
func (backRepo *BackRepoStruct) CommitCHILDREN(children *models.CHILDREN) {
	backRepo.BackRepoCHILDREN.CommitPhaseOneInstance(children)
	if id, ok := backRepo.BackRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]; ok {
		backRepo.BackRepoCHILDREN.CommitPhaseTwoInstance(backRepo, id, children)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitCHILDREN allows checkout of a single children (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutCHILDREN(children *models.CHILDREN) {
	// check if the children is staged
	if _, ok := backRepo.BackRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]; ok {

		if id, ok := backRepo.BackRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[children]; ok {
			var childrenDB CHILDRENDB
			childrenDB.ID = id

			if err := backRepo.BackRepoCHILDREN.db.First(&childrenDB, id).Error; err != nil {
				log.Fatalln("CheckoutCHILDREN : Problem with getting object with id:", id)
			}
			backRepo.BackRepoCHILDREN.CheckoutPhaseOneInstance(&childrenDB)
			backRepo.BackRepoCHILDREN.CheckoutPhaseTwoInstance(backRepo, &childrenDB)
		}
	}
}

// CopyBasicFieldsFromCHILDREN
func (childrenDB *CHILDRENDB) CopyBasicFieldsFromCHILDREN(children *models.CHILDREN) {
	// insertion point for fields commit

	childrenDB.Name_Data.String = children.Name
	childrenDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromCHILDREN_WOP
func (childrenDB *CHILDRENDB) CopyBasicFieldsFromCHILDREN_WOP(children *models.CHILDREN_WOP) {
	// insertion point for fields commit

	childrenDB.Name_Data.String = children.Name
	childrenDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromCHILDRENWOP
func (childrenDB *CHILDRENDB) CopyBasicFieldsFromCHILDRENWOP(children *CHILDRENWOP) {
	// insertion point for fields commit

	childrenDB.Name_Data.String = children.Name
	childrenDB.Name_Data.Valid = true
}

// CopyBasicFieldsToCHILDREN
func (childrenDB *CHILDRENDB) CopyBasicFieldsToCHILDREN(children *models.CHILDREN) {
	// insertion point for checkout of basic fields (back repo to stage)
	children.Name = childrenDB.Name_Data.String
}

// CopyBasicFieldsToCHILDREN_WOP
func (childrenDB *CHILDRENDB) CopyBasicFieldsToCHILDREN_WOP(children *models.CHILDREN_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	children.Name = childrenDB.Name_Data.String
}

// CopyBasicFieldsToCHILDRENWOP
func (childrenDB *CHILDRENDB) CopyBasicFieldsToCHILDRENWOP(children *CHILDRENWOP) {
	children.ID = int(childrenDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	children.Name = childrenDB.Name_Data.String
}

// Backup generates a json file from a slice of all CHILDRENDB instances in the backrepo
func (backRepoCHILDREN *BackRepoCHILDRENStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "CHILDRENDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CHILDRENDB, 0)
	for _, childrenDB := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB {
		forBackup = append(forBackup, childrenDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json CHILDREN ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json CHILDREN file", err.Error())
	}
}

// Backup generates a json file from a slice of all CHILDRENDB instances in the backrepo
func (backRepoCHILDREN *BackRepoCHILDRENStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CHILDRENDB, 0)
	for _, childrenDB := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB {
		forBackup = append(forBackup, childrenDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("CHILDREN")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&CHILDREN_Fields, -1)
	for _, childrenDB := range forBackup {

		var childrenWOP CHILDRENWOP
		childrenDB.CopyBasicFieldsToCHILDRENWOP(&childrenWOP)

		row := sh.AddRow()
		row.WriteStruct(&childrenWOP, -1)
	}
}

// RestoreXL from the "CHILDREN" sheet all CHILDRENDB instances
func (backRepoCHILDREN *BackRepoCHILDRENStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoCHILDRENid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["CHILDREN"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoCHILDREN.rowVisitorCHILDREN)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoCHILDREN *BackRepoCHILDRENStruct) rowVisitorCHILDREN(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var childrenWOP CHILDRENWOP
		row.ReadStruct(&childrenWOP)

		// add the unmarshalled struct to the stage
		childrenDB := new(CHILDRENDB)
		childrenDB.CopyBasicFieldsFromCHILDRENWOP(&childrenWOP)

		childrenDB_ID_atBackupTime := childrenDB.ID
		childrenDB.ID = 0
		query := backRepoCHILDREN.db.Create(childrenDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[childrenDB.ID] = childrenDB
		BackRepoCHILDRENid_atBckpTime_newID[childrenDB_ID_atBackupTime] = childrenDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "CHILDRENDB.json" in dirPath that stores an array
// of CHILDRENDB and stores it in the database
// the map BackRepoCHILDRENid_atBckpTime_newID is updated accordingly
func (backRepoCHILDREN *BackRepoCHILDRENStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoCHILDRENid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "CHILDRENDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json CHILDREN file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*CHILDRENDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_CHILDRENDBID_CHILDRENDB
	for _, childrenDB := range forRestore {

		childrenDB_ID_atBackupTime := childrenDB.ID
		childrenDB.ID = 0
		query := backRepoCHILDREN.db.Create(childrenDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[childrenDB.ID] = childrenDB
		BackRepoCHILDRENid_atBckpTime_newID[childrenDB_ID_atBackupTime] = childrenDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json CHILDREN file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<CHILDREN>id_atBckpTime_newID
// to compute new index
func (backRepoCHILDREN *BackRepoCHILDRENStruct) RestorePhaseTwo() {

	for _, childrenDB := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB {

		// next line of code is to avert unused variable compilation error
		_ = childrenDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoCHILDREN.db.Model(childrenDB).Updates(*childrenDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoCHILDREN.ResetReversePointers commits all staged instances of CHILDREN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCHILDREN *BackRepoCHILDRENStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, children := range backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr {
		backRepoCHILDREN.ResetReversePointersInstance(backRepo, idx, children)
	}

	return
}

func (backRepoCHILDREN *BackRepoCHILDRENStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, children *models.CHILDREN) (Error error) {

	// fetch matching childrenDB
	if childrenDB, ok := backRepoCHILDREN.Map_CHILDRENDBID_CHILDRENDB[idx]; ok {
		_ = childrenDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoCHILDRENid_atBckpTime_newID map[uint]uint
