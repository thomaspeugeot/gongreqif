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
var dummy_DATATYPEDEFINITIONDATE_sql sql.NullBool
var dummy_DATATYPEDEFINITIONDATE_time time.Duration
var dummy_DATATYPEDEFINITIONDATE_sort sort.Float64Slice

// DATATYPEDEFINITIONDATEAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model datatypedefinitiondateAPI
type DATATYPEDEFINITIONDATEAPI struct {
	gorm.Model

	models.DATATYPEDEFINITIONDATE_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	DATATYPEDEFINITIONDATEPointersEncoding DATATYPEDEFINITIONDATEPointersEncoding
}

// DATATYPEDEFINITIONDATEPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DATATYPEDEFINITIONDATEPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ALTERNATIVEID is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ALTERNATIVEIDID sql.NullInt64
}

// DATATYPEDEFINITIONDATEDB describes a datatypedefinitiondate in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model datatypedefinitiondateDB
type DATATYPEDEFINITIONDATEDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field datatypedefinitiondateDB.Name
	Name_Data sql.NullString

	// Declation for basic field datatypedefinitiondateDB.DESCAttr
	DESCAttr_Data sql.NullString

	// Declation for basic field datatypedefinitiondateDB.IDENTIFIERAttr
	IDENTIFIERAttr_Data sql.NullString

	// Declation for basic field datatypedefinitiondateDB.LASTCHANGEAttr
	LASTCHANGEAttr_Data sql.NullString

	// Declation for basic field datatypedefinitiondateDB.LONGNAMEAttr
	LONGNAMEAttr_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	DATATYPEDEFINITIONDATEPointersEncoding
}

// DATATYPEDEFINITIONDATEDBs arrays datatypedefinitiondateDBs
// swagger:response datatypedefinitiondateDBsResponse
type DATATYPEDEFINITIONDATEDBs []DATATYPEDEFINITIONDATEDB

// DATATYPEDEFINITIONDATEDBResponse provides response
// swagger:response datatypedefinitiondateDBResponse
type DATATYPEDEFINITIONDATEDBResponse struct {
	DATATYPEDEFINITIONDATEDB
}

// DATATYPEDEFINITIONDATEWOP is a DATATYPEDEFINITIONDATE without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DATATYPEDEFINITIONDATEWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESCAttr string `xlsx:"2"`

	IDENTIFIERAttr string `xlsx:"3"`

	LASTCHANGEAttr string `xlsx:"4"`

	LONGNAMEAttr string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var DATATYPEDEFINITIONDATE_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESCAttr",
	"IDENTIFIERAttr",
	"LASTCHANGEAttr",
	"LONGNAMEAttr",
}

type BackRepoDATATYPEDEFINITIONDATEStruct struct {
	// stores DATATYPEDEFINITIONDATEDB according to their gorm ID
	Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB map[uint]*DATATYPEDEFINITIONDATEDB

	// stores DATATYPEDEFINITIONDATEDB ID according to DATATYPEDEFINITIONDATE address
	Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID map[*models.DATATYPEDEFINITIONDATE]uint

	// stores DATATYPEDEFINITIONDATE according to their gorm ID
	Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr map[uint]*models.DATATYPEDEFINITIONDATE

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDATATYPEDEFINITIONDATE.stage
	return
}

func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) GetDB() *gorm.DB {
	return backRepoDATATYPEDEFINITIONDATE.db
}

// GetDATATYPEDEFINITIONDATEDBFromDATATYPEDEFINITIONDATEPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) GetDATATYPEDEFINITIONDATEDBFromDATATYPEDEFINITIONDATEPtr(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) {
	id := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]
	datatypedefinitiondateDB = backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[id]
	return
}

// BackRepoDATATYPEDEFINITIONDATE.CommitPhaseOne commits all staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for datatypedefinitiondate := range stage.DATATYPEDEFINITIONDATEs {
		backRepoDATATYPEDEFINITIONDATE.CommitPhaseOneInstance(datatypedefinitiondate)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, datatypedefinitiondate := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr {
		if _, ok := stage.DATATYPEDEFINITIONDATEs[datatypedefinitiondate]; !ok {
			backRepoDATATYPEDEFINITIONDATE.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CommitDeleteInstance commits deletion of DATATYPEDEFINITIONDATE to the BackRepo
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CommitDeleteInstance(id uint) (Error error) {

	datatypedefinitiondate := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[id]

	// datatypedefinitiondate is not staged anymore, remove datatypedefinitiondateDB
	datatypedefinitiondateDB := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[id]
	query := backRepoDATATYPEDEFINITIONDATE.db.Unscoped().Delete(&datatypedefinitiondateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID, datatypedefinitiondate)
	delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr, id)
	delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB, id)

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CommitPhaseOneInstance commits datatypedefinitiondate staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CommitPhaseOneInstance(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) (Error error) {

	// check if the datatypedefinitiondate is not commited yet
	if _, ok := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]; ok {
		return
	}

	// initiate datatypedefinitiondate
	var datatypedefinitiondateDB DATATYPEDEFINITIONDATEDB
	datatypedefinitiondateDB.CopyBasicFieldsFromDATATYPEDEFINITIONDATE(datatypedefinitiondate)

	query := backRepoDATATYPEDEFINITIONDATE.db.Create(&datatypedefinitiondateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate] = datatypedefinitiondateDB.ID
	backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID] = datatypedefinitiondate
	backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[datatypedefinitiondateDB.ID] = &datatypedefinitiondateDB

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CommitPhaseTwo commits all staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, datatypedefinitiondate := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr {
		backRepoDATATYPEDEFINITIONDATE.CommitPhaseTwoInstance(backRepo, idx, datatypedefinitiondate)
	}

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CommitPhaseTwoInstance commits {{structname }} of models.DATATYPEDEFINITIONDATE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) (Error error) {

	// fetch matching datatypedefinitiondateDB
	if datatypedefinitiondateDB, ok := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[idx]; ok {

		datatypedefinitiondateDB.CopyBasicFieldsFromDATATYPEDEFINITIONDATE(datatypedefinitiondate)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value datatypedefinitiondate.ALTERNATIVEID translates to updating the datatypedefinitiondate.ALTERNATIVEIDID
		datatypedefinitiondateDB.ALTERNATIVEIDID.Valid = true // allow for a 0 value (nil association)
		if datatypedefinitiondate.ALTERNATIVEID != nil {
			if ALTERNATIVEIDId, ok := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDPtr_ALTERNATIVEIDDBID[datatypedefinitiondate.ALTERNATIVEID]; ok {
				datatypedefinitiondateDB.ALTERNATIVEIDID.Int64 = int64(ALTERNATIVEIDId)
				datatypedefinitiondateDB.ALTERNATIVEIDID.Valid = true
			}
		} else {
			datatypedefinitiondateDB.ALTERNATIVEIDID.Int64 = 0
			datatypedefinitiondateDB.ALTERNATIVEIDID.Valid = true
		}

		query := backRepoDATATYPEDEFINITIONDATE.db.Save(&datatypedefinitiondateDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DATATYPEDEFINITIONDATE intance %s", datatypedefinitiondate.Name))
		return err
	}

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CheckoutPhaseOne() (Error error) {

	datatypedefinitiondateDBArray := make([]DATATYPEDEFINITIONDATEDB, 0)
	query := backRepoDATATYPEDEFINITIONDATE.db.Find(&datatypedefinitiondateDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	datatypedefinitiondateInstancesToBeRemovedFromTheStage := make(map[*models.DATATYPEDEFINITIONDATE]any)
	for key, value := range backRepoDATATYPEDEFINITIONDATE.stage.DATATYPEDEFINITIONDATEs {
		datatypedefinitiondateInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, datatypedefinitiondateDB := range datatypedefinitiondateDBArray {
		backRepoDATATYPEDEFINITIONDATE.CheckoutPhaseOneInstance(&datatypedefinitiondateDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		datatypedefinitiondate, ok := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]
		if ok {
			delete(datatypedefinitiondateInstancesToBeRemovedFromTheStage, datatypedefinitiondate)
		}
	}

	// remove from stage and back repo's 3 maps all datatypedefinitiondates that are not in the checkout
	for datatypedefinitiondate := range datatypedefinitiondateInstancesToBeRemovedFromTheStage {
		datatypedefinitiondate.Unstage(backRepoDATATYPEDEFINITIONDATE.GetStage())

		// remove instance from the back repo 3 maps
		datatypedefinitiondateID := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]
		delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID, datatypedefinitiondate)
		delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB, datatypedefinitiondateID)
		delete(backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr, datatypedefinitiondateID)
	}

	return
}

// CheckoutPhaseOneInstance takes a datatypedefinitiondateDB that has been found in the DB, updates the backRepo and stages the
// models version of the datatypedefinitiondateDB
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CheckoutPhaseOneInstance(datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) (Error error) {

	datatypedefinitiondate, ok := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]
	if !ok {
		datatypedefinitiondate = new(models.DATATYPEDEFINITIONDATE)

		backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID] = datatypedefinitiondate
		backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate] = datatypedefinitiondateDB.ID

		// append model store with the new element
		datatypedefinitiondate.Name = datatypedefinitiondateDB.Name_Data.String
		datatypedefinitiondate.Stage(backRepoDATATYPEDEFINITIONDATE.GetStage())
	}
	datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATE(datatypedefinitiondate)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	datatypedefinitiondate.Stage(backRepoDATATYPEDEFINITIONDATE.GetStage())

	// preserve pointer to datatypedefinitiondateDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB)[datatypedefinitiondateDB hold variable pointers
	datatypedefinitiondateDB_Data := *datatypedefinitiondateDB
	preservedPtrToDATATYPEDEFINITIONDATE := &datatypedefinitiondateDB_Data
	backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[datatypedefinitiondateDB.ID] = preservedPtrToDATATYPEDEFINITIONDATE

	return
}

// BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseTwo Checkouts all staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, datatypedefinitiondateDB := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB {
		backRepoDATATYPEDEFINITIONDATE.CheckoutPhaseTwoInstance(backRepo, datatypedefinitiondateDB)
	}
	return
}

// BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseTwoInstance Checkouts staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) (Error error) {

	datatypedefinitiondate := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr[datatypedefinitiondateDB.ID]

	datatypedefinitiondateDB.DecodePointers(backRepo, datatypedefinitiondate)

	return
}

func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) DecodePointers(backRepo *BackRepoStruct, datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) {

	// insertion point for checkout of pointer encoding
	// ALTERNATIVEID field
	datatypedefinitiondate.ALTERNATIVEID = nil
	if datatypedefinitiondateDB.ALTERNATIVEIDID.Int64 != 0 {
		datatypedefinitiondate.ALTERNATIVEID = backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[uint(datatypedefinitiondateDB.ALTERNATIVEIDID.Int64)]
	}
	return
}

// CommitDATATYPEDEFINITIONDATE allows commit of a single datatypedefinitiondate (if already staged)
func (backRepo *BackRepoStruct) CommitDATATYPEDEFINITIONDATE(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) {
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CommitPhaseOneInstance(datatypedefinitiondate)
	if id, ok := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]; ok {
		backRepo.BackRepoDATATYPEDEFINITIONDATE.CommitPhaseTwoInstance(backRepo, id, datatypedefinitiondate)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDATATYPEDEFINITIONDATE allows checkout of a single datatypedefinitiondate (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDATATYPEDEFINITIONDATE(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) {
	// check if the datatypedefinitiondate is staged
	if _, ok := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]; ok {

		if id, ok := backRepo.BackRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID[datatypedefinitiondate]; ok {
			var datatypedefinitiondateDB DATATYPEDEFINITIONDATEDB
			datatypedefinitiondateDB.ID = id

			if err := backRepo.BackRepoDATATYPEDEFINITIONDATE.db.First(&datatypedefinitiondateDB, id).Error; err != nil {
				log.Fatalln("CheckoutDATATYPEDEFINITIONDATE : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseOneInstance(&datatypedefinitiondateDB)
			backRepo.BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseTwoInstance(backRepo, &datatypedefinitiondateDB)
		}
	}
}

// CopyBasicFieldsFromDATATYPEDEFINITIONDATE
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsFromDATATYPEDEFINITIONDATE(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) {
	// insertion point for fields commit

	datatypedefinitiondateDB.Name_Data.String = datatypedefinitiondate.Name
	datatypedefinitiondateDB.Name_Data.Valid = true

	datatypedefinitiondateDB.DESCAttr_Data.String = datatypedefinitiondate.DESCAttr
	datatypedefinitiondateDB.DESCAttr_Data.Valid = true

	datatypedefinitiondateDB.IDENTIFIERAttr_Data.String = datatypedefinitiondate.IDENTIFIERAttr
	datatypedefinitiondateDB.IDENTIFIERAttr_Data.Valid = true

	datatypedefinitiondateDB.LASTCHANGEAttr_Data.String = datatypedefinitiondate.LASTCHANGEAttr
	datatypedefinitiondateDB.LASTCHANGEAttr_Data.Valid = true

	datatypedefinitiondateDB.LONGNAMEAttr_Data.String = datatypedefinitiondate.LONGNAMEAttr
	datatypedefinitiondateDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPEDEFINITIONDATE_WOP
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsFromDATATYPEDEFINITIONDATE_WOP(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE_WOP) {
	// insertion point for fields commit

	datatypedefinitiondateDB.Name_Data.String = datatypedefinitiondate.Name
	datatypedefinitiondateDB.Name_Data.Valid = true

	datatypedefinitiondateDB.DESCAttr_Data.String = datatypedefinitiondate.DESCAttr
	datatypedefinitiondateDB.DESCAttr_Data.Valid = true

	datatypedefinitiondateDB.IDENTIFIERAttr_Data.String = datatypedefinitiondate.IDENTIFIERAttr
	datatypedefinitiondateDB.IDENTIFIERAttr_Data.Valid = true

	datatypedefinitiondateDB.LASTCHANGEAttr_Data.String = datatypedefinitiondate.LASTCHANGEAttr
	datatypedefinitiondateDB.LASTCHANGEAttr_Data.Valid = true

	datatypedefinitiondateDB.LONGNAMEAttr_Data.String = datatypedefinitiondate.LONGNAMEAttr
	datatypedefinitiondateDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPEDEFINITIONDATEWOP
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsFromDATATYPEDEFINITIONDATEWOP(datatypedefinitiondate *DATATYPEDEFINITIONDATEWOP) {
	// insertion point for fields commit

	datatypedefinitiondateDB.Name_Data.String = datatypedefinitiondate.Name
	datatypedefinitiondateDB.Name_Data.Valid = true

	datatypedefinitiondateDB.DESCAttr_Data.String = datatypedefinitiondate.DESCAttr
	datatypedefinitiondateDB.DESCAttr_Data.Valid = true

	datatypedefinitiondateDB.IDENTIFIERAttr_Data.String = datatypedefinitiondate.IDENTIFIERAttr
	datatypedefinitiondateDB.IDENTIFIERAttr_Data.Valid = true

	datatypedefinitiondateDB.LASTCHANGEAttr_Data.String = datatypedefinitiondate.LASTCHANGEAttr
	datatypedefinitiondateDB.LASTCHANGEAttr_Data.Valid = true

	datatypedefinitiondateDB.LONGNAMEAttr_Data.String = datatypedefinitiondate.LONGNAMEAttr
	datatypedefinitiondateDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsToDATATYPEDEFINITIONDATE
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsToDATATYPEDEFINITIONDATE(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatypedefinitiondate.Name = datatypedefinitiondateDB.Name_Data.String
	datatypedefinitiondate.DESCAttr = datatypedefinitiondateDB.DESCAttr_Data.String
	datatypedefinitiondate.IDENTIFIERAttr = datatypedefinitiondateDB.IDENTIFIERAttr_Data.String
	datatypedefinitiondate.LASTCHANGEAttr = datatypedefinitiondateDB.LASTCHANGEAttr_Data.String
	datatypedefinitiondate.LONGNAMEAttr = datatypedefinitiondateDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToDATATYPEDEFINITIONDATE_WOP
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsToDATATYPEDEFINITIONDATE_WOP(datatypedefinitiondate *models.DATATYPEDEFINITIONDATE_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatypedefinitiondate.Name = datatypedefinitiondateDB.Name_Data.String
	datatypedefinitiondate.DESCAttr = datatypedefinitiondateDB.DESCAttr_Data.String
	datatypedefinitiondate.IDENTIFIERAttr = datatypedefinitiondateDB.IDENTIFIERAttr_Data.String
	datatypedefinitiondate.LASTCHANGEAttr = datatypedefinitiondateDB.LASTCHANGEAttr_Data.String
	datatypedefinitiondate.LONGNAMEAttr = datatypedefinitiondateDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToDATATYPEDEFINITIONDATEWOP
func (datatypedefinitiondateDB *DATATYPEDEFINITIONDATEDB) CopyBasicFieldsToDATATYPEDEFINITIONDATEWOP(datatypedefinitiondate *DATATYPEDEFINITIONDATEWOP) {
	datatypedefinitiondate.ID = int(datatypedefinitiondateDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	datatypedefinitiondate.Name = datatypedefinitiondateDB.Name_Data.String
	datatypedefinitiondate.DESCAttr = datatypedefinitiondateDB.DESCAttr_Data.String
	datatypedefinitiondate.IDENTIFIERAttr = datatypedefinitiondateDB.IDENTIFIERAttr_Data.String
	datatypedefinitiondate.LASTCHANGEAttr = datatypedefinitiondateDB.LASTCHANGEAttr_Data.String
	datatypedefinitiondate.LONGNAMEAttr = datatypedefinitiondateDB.LONGNAMEAttr_Data.String
}

// Backup generates a json file from a slice of all DATATYPEDEFINITIONDATEDB instances in the backrepo
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DATATYPEDEFINITIONDATEDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPEDEFINITIONDATEDB, 0)
	for _, datatypedefinitiondateDB := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB {
		forBackup = append(forBackup, datatypedefinitiondateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json DATATYPEDEFINITIONDATE ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json DATATYPEDEFINITIONDATE file", err.Error())
	}
}

// Backup generates a json file from a slice of all DATATYPEDEFINITIONDATEDB instances in the backrepo
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPEDEFINITIONDATEDB, 0)
	for _, datatypedefinitiondateDB := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB {
		forBackup = append(forBackup, datatypedefinitiondateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DATATYPEDEFINITIONDATE")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DATATYPEDEFINITIONDATE_Fields, -1)
	for _, datatypedefinitiondateDB := range forBackup {

		var datatypedefinitiondateWOP DATATYPEDEFINITIONDATEWOP
		datatypedefinitiondateDB.CopyBasicFieldsToDATATYPEDEFINITIONDATEWOP(&datatypedefinitiondateWOP)

		row := sh.AddRow()
		row.WriteStruct(&datatypedefinitiondateWOP, -1)
	}
}

// RestoreXL from the "DATATYPEDEFINITIONDATE" sheet all DATATYPEDEFINITIONDATEDB instances
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DATATYPEDEFINITIONDATE"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDATATYPEDEFINITIONDATE.rowVisitorDATATYPEDEFINITIONDATE)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) rowVisitorDATATYPEDEFINITIONDATE(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var datatypedefinitiondateWOP DATATYPEDEFINITIONDATEWOP
		row.ReadStruct(&datatypedefinitiondateWOP)

		// add the unmarshalled struct to the stage
		datatypedefinitiondateDB := new(DATATYPEDEFINITIONDATEDB)
		datatypedefinitiondateDB.CopyBasicFieldsFromDATATYPEDEFINITIONDATEWOP(&datatypedefinitiondateWOP)

		datatypedefinitiondateDB_ID_atBackupTime := datatypedefinitiondateDB.ID
		datatypedefinitiondateDB.ID = 0
		query := backRepoDATATYPEDEFINITIONDATE.db.Create(datatypedefinitiondateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[datatypedefinitiondateDB.ID] = datatypedefinitiondateDB
		BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID[datatypedefinitiondateDB_ID_atBackupTime] = datatypedefinitiondateDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DATATYPEDEFINITIONDATEDB.json" in dirPath that stores an array
// of DATATYPEDEFINITIONDATEDB and stores it in the database
// the map BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID is updated accordingly
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DATATYPEDEFINITIONDATEDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json DATATYPEDEFINITIONDATE file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DATATYPEDEFINITIONDATEDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB
	for _, datatypedefinitiondateDB := range forRestore {

		datatypedefinitiondateDB_ID_atBackupTime := datatypedefinitiondateDB.ID
		datatypedefinitiondateDB.ID = 0
		query := backRepoDATATYPEDEFINITIONDATE.db.Create(datatypedefinitiondateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[datatypedefinitiondateDB.ID] = datatypedefinitiondateDB
		BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID[datatypedefinitiondateDB_ID_atBackupTime] = datatypedefinitiondateDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json DATATYPEDEFINITIONDATE file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DATATYPEDEFINITIONDATE>id_atBckpTime_newID
// to compute new index
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) RestorePhaseTwo() {

	for _, datatypedefinitiondateDB := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB {

		// next line of code is to avert unused variable compilation error
		_ = datatypedefinitiondateDB

		// insertion point for reindexing pointers encoding
		// reindexing ALTERNATIVEID field
		if datatypedefinitiondateDB.ALTERNATIVEIDID.Int64 != 0 {
			datatypedefinitiondateDB.ALTERNATIVEIDID.Int64 = int64(BackRepoALTERNATIVEIDid_atBckpTime_newID[uint(datatypedefinitiondateDB.ALTERNATIVEIDID.Int64)])
			datatypedefinitiondateDB.ALTERNATIVEIDID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoDATATYPEDEFINITIONDATE.db.Model(datatypedefinitiondateDB).Updates(*datatypedefinitiondateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDATATYPEDEFINITIONDATE.ResetReversePointers commits all staged instances of DATATYPEDEFINITIONDATE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, datatypedefinitiondate := range backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr {
		backRepoDATATYPEDEFINITIONDATE.ResetReversePointersInstance(backRepo, idx, datatypedefinitiondate)
	}

	return
}

func (backRepoDATATYPEDEFINITIONDATE *BackRepoDATATYPEDEFINITIONDATEStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, datatypedefinitiondate *models.DATATYPEDEFINITIONDATE) (Error error) {

	// fetch matching datatypedefinitiondateDB
	if datatypedefinitiondateDB, ok := backRepoDATATYPEDEFINITIONDATE.Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB[idx]; ok {
		_ = datatypedefinitiondateDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDATATYPEDEFINITIONDATEid_atBckpTime_newID map[uint]uint
