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
var dummy_REQIF_sql sql.NullBool
var dummy_REQIF_time time.Duration
var dummy_REQIF_sort sort.Float64Slice

// REQIFAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model reqifAPI
type REQIFAPI struct {
	gorm.Model

	models.REQIF_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	REQIFPointersEncoding REQIFPointersEncoding
}

// REQIFPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type REQIFPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field REQ_IF_HEADER is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	REQ_IF_HEADERID sql.NullInt64

	// field REQ_IF_CONTENT is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	REQ_IF_CONTENTID sql.NullInt64
}

// REQIFDB describes a reqif in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model reqifDB
type REQIFDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field reqifDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	REQIFPointersEncoding
}

// REQIFDBs arrays reqifDBs
// swagger:response reqifDBsResponse
type REQIFDBs []REQIFDB

// REQIFDBResponse provides response
// swagger:response reqifDBResponse
type REQIFDBResponse struct {
	REQIFDB
}

// REQIFWOP is a REQIF without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type REQIFWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var REQIF_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoREQIFStruct struct {
	// stores REQIFDB according to their gorm ID
	Map_REQIFDBID_REQIFDB map[uint]*REQIFDB

	// stores REQIFDB ID according to REQIF address
	Map_REQIFPtr_REQIFDBID map[*models.REQIF]uint

	// stores REQIF according to their gorm ID
	Map_REQIFDBID_REQIFPtr map[uint]*models.REQIF

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoREQIF *BackRepoREQIFStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoREQIF.stage
	return
}

func (backRepoREQIF *BackRepoREQIFStruct) GetDB() *gorm.DB {
	return backRepoREQIF.db
}

// GetREQIFDBFromREQIFPtr is a handy function to access the back repo instance from the stage instance
func (backRepoREQIF *BackRepoREQIFStruct) GetREQIFDBFromREQIFPtr(reqif *models.REQIF) (reqifDB *REQIFDB) {
	id := backRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]
	reqifDB = backRepoREQIF.Map_REQIFDBID_REQIFDB[id]
	return
}

// BackRepoREQIF.CommitPhaseOne commits all staged instances of REQIF to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQIF *BackRepoREQIFStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for reqif := range stage.REQIFs {
		backRepoREQIF.CommitPhaseOneInstance(reqif)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, reqif := range backRepoREQIF.Map_REQIFDBID_REQIFPtr {
		if _, ok := stage.REQIFs[reqif]; !ok {
			backRepoREQIF.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoREQIF.CommitDeleteInstance commits deletion of REQIF to the BackRepo
func (backRepoREQIF *BackRepoREQIFStruct) CommitDeleteInstance(id uint) (Error error) {

	reqif := backRepoREQIF.Map_REQIFDBID_REQIFPtr[id]

	// reqif is not staged anymore, remove reqifDB
	reqifDB := backRepoREQIF.Map_REQIFDBID_REQIFDB[id]
	query := backRepoREQIF.db.Unscoped().Delete(&reqifDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoREQIF.Map_REQIFPtr_REQIFDBID, reqif)
	delete(backRepoREQIF.Map_REQIFDBID_REQIFPtr, id)
	delete(backRepoREQIF.Map_REQIFDBID_REQIFDB, id)

	return
}

// BackRepoREQIF.CommitPhaseOneInstance commits reqif staged instances of REQIF to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQIF *BackRepoREQIFStruct) CommitPhaseOneInstance(reqif *models.REQIF) (Error error) {

	// check if the reqif is not commited yet
	if _, ok := backRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]; ok {
		return
	}

	// initiate reqif
	var reqifDB REQIFDB
	reqifDB.CopyBasicFieldsFromREQIF(reqif)

	query := backRepoREQIF.db.Create(&reqifDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif] = reqifDB.ID
	backRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID] = reqif
	backRepoREQIF.Map_REQIFDBID_REQIFDB[reqifDB.ID] = &reqifDB

	return
}

// BackRepoREQIF.CommitPhaseTwo commits all staged instances of REQIF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQIF *BackRepoREQIFStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, reqif := range backRepoREQIF.Map_REQIFDBID_REQIFPtr {
		backRepoREQIF.CommitPhaseTwoInstance(backRepo, idx, reqif)
	}

	return
}

// BackRepoREQIF.CommitPhaseTwoInstance commits {{structname }} of models.REQIF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQIF *BackRepoREQIFStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, reqif *models.REQIF) (Error error) {

	// fetch matching reqifDB
	if reqifDB, ok := backRepoREQIF.Map_REQIFDBID_REQIFDB[idx]; ok {

		reqifDB.CopyBasicFieldsFromREQIF(reqif)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value reqif.REQ_IF_HEADER translates to updating the reqif.REQ_IF_HEADERID
		reqifDB.REQ_IF_HEADERID.Valid = true // allow for a 0 value (nil association)
		if reqif.REQ_IF_HEADER != nil {
			if REQ_IF_HEADERId, ok := backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERPtr_REQ_IF_HEADERDBID[reqif.REQ_IF_HEADER]; ok {
				reqifDB.REQ_IF_HEADERID.Int64 = int64(REQ_IF_HEADERId)
				reqifDB.REQ_IF_HEADERID.Valid = true
			}
		} else {
			reqifDB.REQ_IF_HEADERID.Int64 = 0
			reqifDB.REQ_IF_HEADERID.Valid = true
		}

		// commit pointer value reqif.REQ_IF_CONTENT translates to updating the reqif.REQ_IF_CONTENTID
		reqifDB.REQ_IF_CONTENTID.Valid = true // allow for a 0 value (nil association)
		if reqif.REQ_IF_CONTENT != nil {
			if REQ_IF_CONTENTId, ok := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[reqif.REQ_IF_CONTENT]; ok {
				reqifDB.REQ_IF_CONTENTID.Int64 = int64(REQ_IF_CONTENTId)
				reqifDB.REQ_IF_CONTENTID.Valid = true
			}
		} else {
			reqifDB.REQ_IF_CONTENTID.Int64 = 0
			reqifDB.REQ_IF_CONTENTID.Valid = true
		}

		query := backRepoREQIF.db.Save(&reqifDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown REQIF intance %s", reqif.Name))
		return err
	}

	return
}

// BackRepoREQIF.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoREQIF *BackRepoREQIFStruct) CheckoutPhaseOne() (Error error) {

	reqifDBArray := make([]REQIFDB, 0)
	query := backRepoREQIF.db.Find(&reqifDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	reqifInstancesToBeRemovedFromTheStage := make(map[*models.REQIF]any)
	for key, value := range backRepoREQIF.stage.REQIFs {
		reqifInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, reqifDB := range reqifDBArray {
		backRepoREQIF.CheckoutPhaseOneInstance(&reqifDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		reqif, ok := backRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]
		if ok {
			delete(reqifInstancesToBeRemovedFromTheStage, reqif)
		}
	}

	// remove from stage and back repo's 3 maps all reqifs that are not in the checkout
	for reqif := range reqifInstancesToBeRemovedFromTheStage {
		reqif.Unstage(backRepoREQIF.GetStage())

		// remove instance from the back repo 3 maps
		reqifID := backRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]
		delete(backRepoREQIF.Map_REQIFPtr_REQIFDBID, reqif)
		delete(backRepoREQIF.Map_REQIFDBID_REQIFDB, reqifID)
		delete(backRepoREQIF.Map_REQIFDBID_REQIFPtr, reqifID)
	}

	return
}

// CheckoutPhaseOneInstance takes a reqifDB that has been found in the DB, updates the backRepo and stages the
// models version of the reqifDB
func (backRepoREQIF *BackRepoREQIFStruct) CheckoutPhaseOneInstance(reqifDB *REQIFDB) (Error error) {

	reqif, ok := backRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]
	if !ok {
		reqif = new(models.REQIF)

		backRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID] = reqif
		backRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif] = reqifDB.ID

		// append model store with the new element
		reqif.Name = reqifDB.Name_Data.String
		reqif.Stage(backRepoREQIF.GetStage())
	}
	reqifDB.CopyBasicFieldsToREQIF(reqif)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	reqif.Stage(backRepoREQIF.GetStage())

	// preserve pointer to reqifDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_REQIFDBID_REQIFDB)[reqifDB hold variable pointers
	reqifDB_Data := *reqifDB
	preservedPtrToREQIF := &reqifDB_Data
	backRepoREQIF.Map_REQIFDBID_REQIFDB[reqifDB.ID] = preservedPtrToREQIF

	return
}

// BackRepoREQIF.CheckoutPhaseTwo Checkouts all staged instances of REQIF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQIF *BackRepoREQIFStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, reqifDB := range backRepoREQIF.Map_REQIFDBID_REQIFDB {
		backRepoREQIF.CheckoutPhaseTwoInstance(backRepo, reqifDB)
	}
	return
}

// BackRepoREQIF.CheckoutPhaseTwoInstance Checkouts staged instances of REQIF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQIF *BackRepoREQIFStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, reqifDB *REQIFDB) (Error error) {

	reqif := backRepoREQIF.Map_REQIFDBID_REQIFPtr[reqifDB.ID]

	reqifDB.DecodePointers(backRepo, reqif)

	return
}

func (reqifDB *REQIFDB) DecodePointers(backRepo *BackRepoStruct, reqif *models.REQIF) {

	// insertion point for checkout of pointer encoding
	// REQ_IF_HEADER field
	reqif.REQ_IF_HEADER = nil
	if reqifDB.REQ_IF_HEADERID.Int64 != 0 {
		reqif.REQ_IF_HEADER = backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr[uint(reqifDB.REQ_IF_HEADERID.Int64)]
	}
	// REQ_IF_CONTENT field
	reqif.REQ_IF_CONTENT = nil
	if reqifDB.REQ_IF_CONTENTID.Int64 != 0 {
		reqif.REQ_IF_CONTENT = backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[uint(reqifDB.REQ_IF_CONTENTID.Int64)]
	}
	return
}

// CommitREQIF allows commit of a single reqif (if already staged)
func (backRepo *BackRepoStruct) CommitREQIF(reqif *models.REQIF) {
	backRepo.BackRepoREQIF.CommitPhaseOneInstance(reqif)
	if id, ok := backRepo.BackRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]; ok {
		backRepo.BackRepoREQIF.CommitPhaseTwoInstance(backRepo, id, reqif)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitREQIF allows checkout of a single reqif (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutREQIF(reqif *models.REQIF) {
	// check if the reqif is staged
	if _, ok := backRepo.BackRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]; ok {

		if id, ok := backRepo.BackRepoREQIF.Map_REQIFPtr_REQIFDBID[reqif]; ok {
			var reqifDB REQIFDB
			reqifDB.ID = id

			if err := backRepo.BackRepoREQIF.db.First(&reqifDB, id).Error; err != nil {
				log.Fatalln("CheckoutREQIF : Problem with getting object with id:", id)
			}
			backRepo.BackRepoREQIF.CheckoutPhaseOneInstance(&reqifDB)
			backRepo.BackRepoREQIF.CheckoutPhaseTwoInstance(backRepo, &reqifDB)
		}
	}
}

// CopyBasicFieldsFromREQIF
func (reqifDB *REQIFDB) CopyBasicFieldsFromREQIF(reqif *models.REQIF) {
	// insertion point for fields commit

	reqifDB.Name_Data.String = reqif.Name
	reqifDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQIF_WOP
func (reqifDB *REQIFDB) CopyBasicFieldsFromREQIF_WOP(reqif *models.REQIF_WOP) {
	// insertion point for fields commit

	reqifDB.Name_Data.String = reqif.Name
	reqifDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQIFWOP
func (reqifDB *REQIFDB) CopyBasicFieldsFromREQIFWOP(reqif *REQIFWOP) {
	// insertion point for fields commit

	reqifDB.Name_Data.String = reqif.Name
	reqifDB.Name_Data.Valid = true
}

// CopyBasicFieldsToREQIF
func (reqifDB *REQIFDB) CopyBasicFieldsToREQIF(reqif *models.REQIF) {
	// insertion point for checkout of basic fields (back repo to stage)
	reqif.Name = reqifDB.Name_Data.String
}

// CopyBasicFieldsToREQIF_WOP
func (reqifDB *REQIFDB) CopyBasicFieldsToREQIF_WOP(reqif *models.REQIF_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	reqif.Name = reqifDB.Name_Data.String
}

// CopyBasicFieldsToREQIFWOP
func (reqifDB *REQIFDB) CopyBasicFieldsToREQIFWOP(reqif *REQIFWOP) {
	reqif.ID = int(reqifDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	reqif.Name = reqifDB.Name_Data.String
}

// Backup generates a json file from a slice of all REQIFDB instances in the backrepo
func (backRepoREQIF *BackRepoREQIFStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "REQIFDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQIFDB, 0)
	for _, reqifDB := range backRepoREQIF.Map_REQIFDBID_REQIFDB {
		forBackup = append(forBackup, reqifDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json REQIF ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json REQIF file", err.Error())
	}
}

// Backup generates a json file from a slice of all REQIFDB instances in the backrepo
func (backRepoREQIF *BackRepoREQIFStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQIFDB, 0)
	for _, reqifDB := range backRepoREQIF.Map_REQIFDBID_REQIFDB {
		forBackup = append(forBackup, reqifDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("REQIF")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&REQIF_Fields, -1)
	for _, reqifDB := range forBackup {

		var reqifWOP REQIFWOP
		reqifDB.CopyBasicFieldsToREQIFWOP(&reqifWOP)

		row := sh.AddRow()
		row.WriteStruct(&reqifWOP, -1)
	}
}

// RestoreXL from the "REQIF" sheet all REQIFDB instances
func (backRepoREQIF *BackRepoREQIFStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoREQIFid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["REQIF"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoREQIF.rowVisitorREQIF)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoREQIF *BackRepoREQIFStruct) rowVisitorREQIF(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var reqifWOP REQIFWOP
		row.ReadStruct(&reqifWOP)

		// add the unmarshalled struct to the stage
		reqifDB := new(REQIFDB)
		reqifDB.CopyBasicFieldsFromREQIFWOP(&reqifWOP)

		reqifDB_ID_atBackupTime := reqifDB.ID
		reqifDB.ID = 0
		query := backRepoREQIF.db.Create(reqifDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQIF.Map_REQIFDBID_REQIFDB[reqifDB.ID] = reqifDB
		BackRepoREQIFid_atBckpTime_newID[reqifDB_ID_atBackupTime] = reqifDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "REQIFDB.json" in dirPath that stores an array
// of REQIFDB and stores it in the database
// the map BackRepoREQIFid_atBckpTime_newID is updated accordingly
func (backRepoREQIF *BackRepoREQIFStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoREQIFid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "REQIFDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json REQIF file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*REQIFDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_REQIFDBID_REQIFDB
	for _, reqifDB := range forRestore {

		reqifDB_ID_atBackupTime := reqifDB.ID
		reqifDB.ID = 0
		query := backRepoREQIF.db.Create(reqifDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQIF.Map_REQIFDBID_REQIFDB[reqifDB.ID] = reqifDB
		BackRepoREQIFid_atBckpTime_newID[reqifDB_ID_atBackupTime] = reqifDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json REQIF file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<REQIF>id_atBckpTime_newID
// to compute new index
func (backRepoREQIF *BackRepoREQIFStruct) RestorePhaseTwo() {

	for _, reqifDB := range backRepoREQIF.Map_REQIFDBID_REQIFDB {

		// next line of code is to avert unused variable compilation error
		_ = reqifDB

		// insertion point for reindexing pointers encoding
		// reindexing REQ_IF_HEADER field
		if reqifDB.REQ_IF_HEADERID.Int64 != 0 {
			reqifDB.REQ_IF_HEADERID.Int64 = int64(BackRepoREQ_IF_HEADERid_atBckpTime_newID[uint(reqifDB.REQ_IF_HEADERID.Int64)])
			reqifDB.REQ_IF_HEADERID.Valid = true
		}

		// reindexing REQ_IF_CONTENT field
		if reqifDB.REQ_IF_CONTENTID.Int64 != 0 {
			reqifDB.REQ_IF_CONTENTID.Int64 = int64(BackRepoREQ_IF_CONTENTid_atBckpTime_newID[uint(reqifDB.REQ_IF_CONTENTID.Int64)])
			reqifDB.REQ_IF_CONTENTID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoREQIF.db.Model(reqifDB).Updates(*reqifDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoREQIF.ResetReversePointers commits all staged instances of REQIF to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQIF *BackRepoREQIFStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, reqif := range backRepoREQIF.Map_REQIFDBID_REQIFPtr {
		backRepoREQIF.ResetReversePointersInstance(backRepo, idx, reqif)
	}

	return
}

func (backRepoREQIF *BackRepoREQIFStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, reqif *models.REQIF) (Error error) {

	// fetch matching reqifDB
	if reqifDB, ok := backRepoREQIF.Map_REQIFDBID_REQIFDB[idx]; ok {
		_ = reqifDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoREQIFid_atBckpTime_newID map[uint]uint
