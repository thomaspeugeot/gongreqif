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
var dummy_ATTRIBUTEDEFINITIONBOOLEAN_sql sql.NullBool
var dummy_ATTRIBUTEDEFINITIONBOOLEAN_time time.Duration
var dummy_ATTRIBUTEDEFINITIONBOOLEAN_sort sort.Float64Slice

// ATTRIBUTEDEFINITIONBOOLEANAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attributedefinitionbooleanAPI
type ATTRIBUTEDEFINITIONBOOLEANAPI struct {
	gorm.Model

	models.ATTRIBUTEDEFINITIONBOOLEAN_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTEDEFINITIONBOOLEANPointersEncoding ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
}

// ATTRIBUTEDEFINITIONBOOLEANPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTEDEFINITIONBOOLEANPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ALTERNATIVEID is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ALTERNATIVEIDID sql.NullInt64

	// field DEFAULTVALUE is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	DEFAULTVALUEID sql.NullInt64

	// field TYPE is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	TYPEID sql.NullInt64
}

// ATTRIBUTEDEFINITIONBOOLEANDB describes a attributedefinitionboolean in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attributedefinitionbooleanDB
type ATTRIBUTEDEFINITIONBOOLEANDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attributedefinitionbooleanDB.Name
	Name_Data sql.NullString

	// Declation for basic field attributedefinitionbooleanDB.DESCAttr
	DESCAttr_Data sql.NullString

	// Declation for basic field attributedefinitionbooleanDB.IDENTIFIERAttr
	IDENTIFIERAttr_Data sql.NullString

	// Declation for basic field attributedefinitionbooleanDB.ISEDITABLEAttr
	// provide the sql storage for the boolan
	ISEDITABLEAttr_Data sql.NullBool

	// Declation for basic field attributedefinitionbooleanDB.LASTCHANGEAttr
	LASTCHANGEAttr_Data sql.NullString

	// Declation for basic field attributedefinitionbooleanDB.LONGNAMEAttr
	LONGNAMEAttr_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTEDEFINITIONBOOLEANPointersEncoding
}

// ATTRIBUTEDEFINITIONBOOLEANDBs arrays attributedefinitionbooleanDBs
// swagger:response attributedefinitionbooleanDBsResponse
type ATTRIBUTEDEFINITIONBOOLEANDBs []ATTRIBUTEDEFINITIONBOOLEANDB

// ATTRIBUTEDEFINITIONBOOLEANDBResponse provides response
// swagger:response attributedefinitionbooleanDBResponse
type ATTRIBUTEDEFINITIONBOOLEANDBResponse struct {
	ATTRIBUTEDEFINITIONBOOLEANDB
}

// ATTRIBUTEDEFINITIONBOOLEANWOP is a ATTRIBUTEDEFINITIONBOOLEAN without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTEDEFINITIONBOOLEANWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESCAttr string `xlsx:"2"`

	IDENTIFIERAttr string `xlsx:"3"`

	ISEDITABLEAttr bool `xlsx:"4"`

	LASTCHANGEAttr string `xlsx:"5"`

	LONGNAMEAttr string `xlsx:"6"`
	// insertion for WOP pointer fields
}

var ATTRIBUTEDEFINITIONBOOLEAN_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESCAttr",
	"IDENTIFIERAttr",
	"ISEDITABLEAttr",
	"LASTCHANGEAttr",
	"LONGNAMEAttr",
}

type BackRepoATTRIBUTEDEFINITIONBOOLEANStruct struct {
	// stores ATTRIBUTEDEFINITIONBOOLEANDB according to their gorm ID
	Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB map[uint]*ATTRIBUTEDEFINITIONBOOLEANDB

	// stores ATTRIBUTEDEFINITIONBOOLEANDB ID according to ATTRIBUTEDEFINITIONBOOLEAN address
	Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID map[*models.ATTRIBUTEDEFINITIONBOOLEAN]uint

	// stores ATTRIBUTEDEFINITIONBOOLEAN according to their gorm ID
	Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr map[uint]*models.ATTRIBUTEDEFINITIONBOOLEAN

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoATTRIBUTEDEFINITIONBOOLEAN.stage
	return
}

func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) GetDB() *gorm.DB {
	return backRepoATTRIBUTEDEFINITIONBOOLEAN.db
}

// GetATTRIBUTEDEFINITIONBOOLEANDBFromATTRIBUTEDEFINITIONBOOLEANPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) GetATTRIBUTEDEFINITIONBOOLEANDBFromATTRIBUTEDEFINITIONBOOLEANPtr(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) {
	id := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]
	attributedefinitionbooleanDB = backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[id]
	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseOne commits all staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for attributedefinitionboolean := range stage.ATTRIBUTEDEFINITIONBOOLEANs {
		backRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseOneInstance(attributedefinitionboolean)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attributedefinitionboolean := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr {
		if _, ok := stage.ATTRIBUTEDEFINITIONBOOLEANs[attributedefinitionboolean]; !ok {
			backRepoATTRIBUTEDEFINITIONBOOLEAN.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitDeleteInstance commits deletion of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CommitDeleteInstance(id uint) (Error error) {

	attributedefinitionboolean := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[id]

	// attributedefinitionboolean is not staged anymore, remove attributedefinitionbooleanDB
	attributedefinitionbooleanDB := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[id]
	query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Unscoped().Delete(&attributedefinitionbooleanDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID, attributedefinitionboolean)
	delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr, id)
	delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB, id)

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseOneInstance commits attributedefinitionboolean staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CommitPhaseOneInstance(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) (Error error) {

	// check if the attributedefinitionboolean is not commited yet
	if _, ok := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]; ok {
		return
	}

	// initiate attributedefinitionboolean
	var attributedefinitionbooleanDB ATTRIBUTEDEFINITIONBOOLEANDB
	attributedefinitionbooleanDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean)

	query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Create(&attributedefinitionbooleanDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean] = attributedefinitionbooleanDB.ID
	backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID] = attributedefinitionboolean
	backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[attributedefinitionbooleanDB.ID] = &attributedefinitionbooleanDB

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseTwo commits all staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attributedefinitionboolean := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr {
		backRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseTwoInstance(backRepo, idx, attributedefinitionboolean)
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) (Error error) {

	// fetch matching attributedefinitionbooleanDB
	if attributedefinitionbooleanDB, ok := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[idx]; ok {

		attributedefinitionbooleanDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value attributedefinitionboolean.ALTERNATIVEID translates to updating the attributedefinitionboolean.ALTERNATIVEIDID
		attributedefinitionbooleanDB.ALTERNATIVEIDID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionboolean.ALTERNATIVEID != nil {
			if ALTERNATIVEIDId, ok := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDPtr_ALTERNATIVEIDDBID[attributedefinitionboolean.ALTERNATIVEID]; ok {
				attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64 = int64(ALTERNATIVEIDId)
				attributedefinitionbooleanDB.ALTERNATIVEIDID.Valid = true
			}
		} else {
			attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64 = 0
			attributedefinitionbooleanDB.ALTERNATIVEIDID.Valid = true
		}

		// commit pointer value attributedefinitionboolean.DEFAULTVALUE translates to updating the attributedefinitionboolean.DEFAULTVALUEID
		attributedefinitionbooleanDB.DEFAULTVALUEID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionboolean.DEFAULTVALUE != nil {
			if DEFAULTVALUEId, ok := backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEPtr_DEFAULTVALUEDBID[attributedefinitionboolean.DEFAULTVALUE]; ok {
				attributedefinitionbooleanDB.DEFAULTVALUEID.Int64 = int64(DEFAULTVALUEId)
				attributedefinitionbooleanDB.DEFAULTVALUEID.Valid = true
			}
		} else {
			attributedefinitionbooleanDB.DEFAULTVALUEID.Int64 = 0
			attributedefinitionbooleanDB.DEFAULTVALUEID.Valid = true
		}

		// commit pointer value attributedefinitionboolean.TYPE translates to updating the attributedefinitionboolean.TYPEID
		attributedefinitionbooleanDB.TYPEID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionboolean.TYPE != nil {
			if TYPEId, ok := backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEPtr_REQIFTYPEDBID[attributedefinitionboolean.TYPE]; ok {
				attributedefinitionbooleanDB.TYPEID.Int64 = int64(TYPEId)
				attributedefinitionbooleanDB.TYPEID.Valid = true
			}
		} else {
			attributedefinitionbooleanDB.TYPEID.Int64 = 0
			attributedefinitionbooleanDB.TYPEID.Valid = true
		}

		query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Save(&attributedefinitionbooleanDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTEDEFINITIONBOOLEAN intance %s", attributedefinitionboolean.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CheckoutPhaseOne() (Error error) {

	attributedefinitionbooleanDBArray := make([]ATTRIBUTEDEFINITIONBOOLEANDB, 0)
	query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Find(&attributedefinitionbooleanDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attributedefinitionbooleanInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTEDEFINITIONBOOLEAN]any)
	for key, value := range backRepoATTRIBUTEDEFINITIONBOOLEAN.stage.ATTRIBUTEDEFINITIONBOOLEANs {
		attributedefinitionbooleanInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attributedefinitionbooleanDB := range attributedefinitionbooleanDBArray {
		backRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseOneInstance(&attributedefinitionbooleanDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attributedefinitionboolean, ok := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]
		if ok {
			delete(attributedefinitionbooleanInstancesToBeRemovedFromTheStage, attributedefinitionboolean)
		}
	}

	// remove from stage and back repo's 3 maps all attributedefinitionbooleans that are not in the checkout
	for attributedefinitionboolean := range attributedefinitionbooleanInstancesToBeRemovedFromTheStage {
		attributedefinitionboolean.Unstage(backRepoATTRIBUTEDEFINITIONBOOLEAN.GetStage())

		// remove instance from the back repo 3 maps
		attributedefinitionbooleanID := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]
		delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID, attributedefinitionboolean)
		delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB, attributedefinitionbooleanID)
		delete(backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr, attributedefinitionbooleanID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attributedefinitionbooleanDB that has been found in the DB, updates the backRepo and stages the
// models version of the attributedefinitionbooleanDB
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CheckoutPhaseOneInstance(attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) (Error error) {

	attributedefinitionboolean, ok := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]
	if !ok {
		attributedefinitionboolean = new(models.ATTRIBUTEDEFINITIONBOOLEAN)

		backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID] = attributedefinitionboolean
		backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean] = attributedefinitionbooleanDB.ID

		// append model store with the new element
		attributedefinitionboolean.Name = attributedefinitionbooleanDB.Name_Data.String
		attributedefinitionboolean.Stage(backRepoATTRIBUTEDEFINITIONBOOLEAN.GetStage())
	}
	attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attributedefinitionboolean.Stage(backRepoATTRIBUTEDEFINITIONBOOLEAN.GetStage())

	// preserve pointer to attributedefinitionbooleanDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB)[attributedefinitionbooleanDB hold variable pointers
	attributedefinitionbooleanDB_Data := *attributedefinitionbooleanDB
	preservedPtrToATTRIBUTEDEFINITIONBOOLEAN := &attributedefinitionbooleanDB_Data
	backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[attributedefinitionbooleanDB.ID] = preservedPtrToATTRIBUTEDEFINITIONBOOLEAN

	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attributedefinitionbooleanDB := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB {
		backRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseTwoInstance(backRepo, attributedefinitionbooleanDB)
	}
	return
}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) (Error error) {

	attributedefinitionboolean := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr[attributedefinitionbooleanDB.ID]

	attributedefinitionbooleanDB.DecodePointers(backRepo, attributedefinitionboolean)

	return
}

func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) DecodePointers(backRepo *BackRepoStruct, attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) {

	// insertion point for checkout of pointer encoding
	// ALTERNATIVEID field
	attributedefinitionboolean.ALTERNATIVEID = nil
	if attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64 != 0 {
		attributedefinitionboolean.ALTERNATIVEID = backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[uint(attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64)]
	}
	// DEFAULTVALUE field
	attributedefinitionboolean.DEFAULTVALUE = nil
	if attributedefinitionbooleanDB.DEFAULTVALUEID.Int64 != 0 {
		attributedefinitionboolean.DEFAULTVALUE = backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr[uint(attributedefinitionbooleanDB.DEFAULTVALUEID.Int64)]
	}
	// TYPE field
	attributedefinitionboolean.TYPE = nil
	if attributedefinitionbooleanDB.TYPEID.Int64 != 0 {
		attributedefinitionboolean.TYPE = backRepo.BackRepoREQIFTYPE.Map_REQIFTYPEDBID_REQIFTYPEPtr[uint(attributedefinitionbooleanDB.TYPEID.Int64)]
	}
	return
}

// CommitATTRIBUTEDEFINITIONBOOLEAN allows commit of a single attributedefinitionboolean (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) {
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseOneInstance(attributedefinitionboolean)
	if id, ok := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]; ok {
		backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseTwoInstance(backRepo, id, attributedefinitionboolean)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTEDEFINITIONBOOLEAN allows checkout of a single attributedefinitionboolean (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) {
	// check if the attributedefinitionboolean is staged
	if _, ok := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID[attributedefinitionboolean]; ok {
			var attributedefinitionbooleanDB ATTRIBUTEDEFINITIONBOOLEANDB
			attributedefinitionbooleanDB.ID = id

			if err := backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.db.First(&attributedefinitionbooleanDB, id).Error; err != nil {
				log.Fatalln("CheckoutATTRIBUTEDEFINITIONBOOLEAN : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseOneInstance(&attributedefinitionbooleanDB)
			backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseTwoInstance(backRepo, &attributedefinitionbooleanDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) {
	// insertion point for fields commit

	attributedefinitionbooleanDB.Name_Data.String = attributedefinitionboolean.Name
	attributedefinitionbooleanDB.Name_Data.Valid = true

	attributedefinitionbooleanDB.DESCAttr_Data.String = attributedefinitionboolean.DESCAttr
	attributedefinitionbooleanDB.DESCAttr_Data.Valid = true

	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String = attributedefinitionboolean.IDENTIFIERAttr
	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool = attributedefinitionboolean.ISEDITABLEAttr
	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String = attributedefinitionboolean.LASTCHANGEAttr
	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LONGNAMEAttr_Data.String = attributedefinitionboolean.LONGNAMEAttr
	attributedefinitionbooleanDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN_WOP
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEAN_WOP(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN_WOP) {
	// insertion point for fields commit

	attributedefinitionbooleanDB.Name_Data.String = attributedefinitionboolean.Name
	attributedefinitionbooleanDB.Name_Data.Valid = true

	attributedefinitionbooleanDB.DESCAttr_Data.String = attributedefinitionboolean.DESCAttr
	attributedefinitionbooleanDB.DESCAttr_Data.Valid = true

	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String = attributedefinitionboolean.IDENTIFIERAttr
	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool = attributedefinitionboolean.ISEDITABLEAttr
	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String = attributedefinitionboolean.LASTCHANGEAttr
	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LONGNAMEAttr_Data.String = attributedefinitionboolean.LONGNAMEAttr
	attributedefinitionbooleanDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEANWOP
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEANWOP(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEANWOP) {
	// insertion point for fields commit

	attributedefinitionbooleanDB.Name_Data.String = attributedefinitionboolean.Name
	attributedefinitionbooleanDB.Name_Data.Valid = true

	attributedefinitionbooleanDB.DESCAttr_Data.String = attributedefinitionboolean.DESCAttr
	attributedefinitionbooleanDB.DESCAttr_Data.Valid = true

	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String = attributedefinitionboolean.IDENTIFIERAttr
	attributedefinitionbooleanDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool = attributedefinitionboolean.ISEDITABLEAttr
	attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String = attributedefinitionboolean.LASTCHANGEAttr
	attributedefinitionbooleanDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionbooleanDB.LONGNAMEAttr_Data.String = attributedefinitionboolean.LONGNAMEAttr
	attributedefinitionbooleanDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) {
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionboolean.Name = attributedefinitionbooleanDB.Name_Data.String
	attributedefinitionboolean.DESCAttr = attributedefinitionbooleanDB.DESCAttr_Data.String
	attributedefinitionboolean.IDENTIFIERAttr = attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String
	attributedefinitionboolean.ISEDITABLEAttr = attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionboolean.LASTCHANGEAttr = attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String
	attributedefinitionboolean.LONGNAMEAttr = attributedefinitionbooleanDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN_WOP
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEAN_WOP(attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionboolean.Name = attributedefinitionbooleanDB.Name_Data.String
	attributedefinitionboolean.DESCAttr = attributedefinitionbooleanDB.DESCAttr_Data.String
	attributedefinitionboolean.IDENTIFIERAttr = attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String
	attributedefinitionboolean.ISEDITABLEAttr = attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionboolean.LASTCHANGEAttr = attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String
	attributedefinitionboolean.LONGNAMEAttr = attributedefinitionbooleanDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEANWOP
func (attributedefinitionbooleanDB *ATTRIBUTEDEFINITIONBOOLEANDB) CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEANWOP(attributedefinitionboolean *ATTRIBUTEDEFINITIONBOOLEANWOP) {
	attributedefinitionboolean.ID = int(attributedefinitionbooleanDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionboolean.Name = attributedefinitionbooleanDB.Name_Data.String
	attributedefinitionboolean.DESCAttr = attributedefinitionbooleanDB.DESCAttr_Data.String
	attributedefinitionboolean.IDENTIFIERAttr = attributedefinitionbooleanDB.IDENTIFIERAttr_Data.String
	attributedefinitionboolean.ISEDITABLEAttr = attributedefinitionbooleanDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionboolean.LASTCHANGEAttr = attributedefinitionbooleanDB.LASTCHANGEAttr_Data.String
	attributedefinitionboolean.LONGNAMEAttr = attributedefinitionbooleanDB.LONGNAMEAttr_Data.String
}

// Backup generates a json file from a slice of all ATTRIBUTEDEFINITIONBOOLEANDB instances in the backrepo
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTEDEFINITIONBOOLEANDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTEDEFINITIONBOOLEANDB, 0)
	for _, attributedefinitionbooleanDB := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB {
		forBackup = append(forBackup, attributedefinitionbooleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTEDEFINITIONBOOLEAN ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTEDEFINITIONBOOLEAN file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTEDEFINITIONBOOLEANDB instances in the backrepo
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTEDEFINITIONBOOLEANDB, 0)
	for _, attributedefinitionbooleanDB := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB {
		forBackup = append(forBackup, attributedefinitionbooleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTEDEFINITIONBOOLEAN")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTEDEFINITIONBOOLEAN_Fields, -1)
	for _, attributedefinitionbooleanDB := range forBackup {

		var attributedefinitionbooleanWOP ATTRIBUTEDEFINITIONBOOLEANWOP
		attributedefinitionbooleanDB.CopyBasicFieldsToATTRIBUTEDEFINITIONBOOLEANWOP(&attributedefinitionbooleanWOP)

		row := sh.AddRow()
		row.WriteStruct(&attributedefinitionbooleanWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTEDEFINITIONBOOLEAN" sheet all ATTRIBUTEDEFINITIONBOOLEANDB instances
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTEDEFINITIONBOOLEAN"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTEDEFINITIONBOOLEAN.rowVisitorATTRIBUTEDEFINITIONBOOLEAN)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) rowVisitorATTRIBUTEDEFINITIONBOOLEAN(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attributedefinitionbooleanWOP ATTRIBUTEDEFINITIONBOOLEANWOP
		row.ReadStruct(&attributedefinitionbooleanWOP)

		// add the unmarshalled struct to the stage
		attributedefinitionbooleanDB := new(ATTRIBUTEDEFINITIONBOOLEANDB)
		attributedefinitionbooleanDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONBOOLEANWOP(&attributedefinitionbooleanWOP)

		attributedefinitionbooleanDB_ID_atBackupTime := attributedefinitionbooleanDB.ID
		attributedefinitionbooleanDB.ID = 0
		query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Create(attributedefinitionbooleanDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[attributedefinitionbooleanDB.ID] = attributedefinitionbooleanDB
		BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID[attributedefinitionbooleanDB_ID_atBackupTime] = attributedefinitionbooleanDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTEDEFINITIONBOOLEANDB.json" in dirPath that stores an array
// of ATTRIBUTEDEFINITIONBOOLEANDB and stores it in the database
// the map BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTEDEFINITIONBOOLEANDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTEDEFINITIONBOOLEAN file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTEDEFINITIONBOOLEANDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB
	for _, attributedefinitionbooleanDB := range forRestore {

		attributedefinitionbooleanDB_ID_atBackupTime := attributedefinitionbooleanDB.ID
		attributedefinitionbooleanDB.ID = 0
		query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Create(attributedefinitionbooleanDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[attributedefinitionbooleanDB.ID] = attributedefinitionbooleanDB
		BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID[attributedefinitionbooleanDB_ID_atBackupTime] = attributedefinitionbooleanDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTEDEFINITIONBOOLEAN file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTEDEFINITIONBOOLEAN>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) RestorePhaseTwo() {

	for _, attributedefinitionbooleanDB := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB {

		// next line of code is to avert unused variable compilation error
		_ = attributedefinitionbooleanDB

		// insertion point for reindexing pointers encoding
		// reindexing ALTERNATIVEID field
		if attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64 != 0 {
			attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64 = int64(BackRepoALTERNATIVEIDid_atBckpTime_newID[uint(attributedefinitionbooleanDB.ALTERNATIVEIDID.Int64)])
			attributedefinitionbooleanDB.ALTERNATIVEIDID.Valid = true
		}

		// reindexing DEFAULTVALUE field
		if attributedefinitionbooleanDB.DEFAULTVALUEID.Int64 != 0 {
			attributedefinitionbooleanDB.DEFAULTVALUEID.Int64 = int64(BackRepoDEFAULTVALUEid_atBckpTime_newID[uint(attributedefinitionbooleanDB.DEFAULTVALUEID.Int64)])
			attributedefinitionbooleanDB.DEFAULTVALUEID.Valid = true
		}

		// reindexing TYPE field
		if attributedefinitionbooleanDB.TYPEID.Int64 != 0 {
			attributedefinitionbooleanDB.TYPEID.Int64 = int64(BackRepoREQIFTYPEid_atBckpTime_newID[uint(attributedefinitionbooleanDB.TYPEID.Int64)])
			attributedefinitionbooleanDB.TYPEID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoATTRIBUTEDEFINITIONBOOLEAN.db.Model(attributedefinitionbooleanDB).Updates(*attributedefinitionbooleanDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoATTRIBUTEDEFINITIONBOOLEAN.ResetReversePointers commits all staged instances of ATTRIBUTEDEFINITIONBOOLEAN to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attributedefinitionboolean := range backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr {
		backRepoATTRIBUTEDEFINITIONBOOLEAN.ResetReversePointersInstance(backRepo, idx, attributedefinitionboolean)
	}

	return
}

func (backRepoATTRIBUTEDEFINITIONBOOLEAN *BackRepoATTRIBUTEDEFINITIONBOOLEANStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attributedefinitionboolean *models.ATTRIBUTEDEFINITIONBOOLEAN) (Error error) {

	// fetch matching attributedefinitionbooleanDB
	if attributedefinitionbooleanDB, ok := backRepoATTRIBUTEDEFINITIONBOOLEAN.Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB[idx]; ok {
		_ = attributedefinitionbooleanDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTEDEFINITIONBOOLEANid_atBckpTime_newID map[uint]uint
