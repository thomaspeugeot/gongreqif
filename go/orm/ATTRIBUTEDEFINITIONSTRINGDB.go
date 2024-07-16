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
var dummy_ATTRIBUTEDEFINITIONSTRING_sql sql.NullBool
var dummy_ATTRIBUTEDEFINITIONSTRING_time time.Duration
var dummy_ATTRIBUTEDEFINITIONSTRING_sort sort.Float64Slice

// ATTRIBUTEDEFINITIONSTRINGAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model attributedefinitionstringAPI
type ATTRIBUTEDEFINITIONSTRINGAPI struct {
	gorm.Model

	models.ATTRIBUTEDEFINITIONSTRING_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ATTRIBUTEDEFINITIONSTRINGPointersEncoding ATTRIBUTEDEFINITIONSTRINGPointersEncoding
}

// ATTRIBUTEDEFINITIONSTRINGPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ATTRIBUTEDEFINITIONSTRINGPointersEncoding struct {
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

// ATTRIBUTEDEFINITIONSTRINGDB describes a attributedefinitionstring in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model attributedefinitionstringDB
type ATTRIBUTEDEFINITIONSTRINGDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field attributedefinitionstringDB.Name
	Name_Data sql.NullString

	// Declation for basic field attributedefinitionstringDB.DESCAttr
	DESCAttr_Data sql.NullString

	// Declation for basic field attributedefinitionstringDB.IDENTIFIERAttr
	IDENTIFIERAttr_Data sql.NullString

	// Declation for basic field attributedefinitionstringDB.ISEDITABLEAttr
	// provide the sql storage for the boolan
	ISEDITABLEAttr_Data sql.NullBool

	// Declation for basic field attributedefinitionstringDB.LASTCHANGEAttr
	LASTCHANGEAttr_Data sql.NullString

	// Declation for basic field attributedefinitionstringDB.LONGNAMEAttr
	LONGNAMEAttr_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ATTRIBUTEDEFINITIONSTRINGPointersEncoding
}

// ATTRIBUTEDEFINITIONSTRINGDBs arrays attributedefinitionstringDBs
// swagger:response attributedefinitionstringDBsResponse
type ATTRIBUTEDEFINITIONSTRINGDBs []ATTRIBUTEDEFINITIONSTRINGDB

// ATTRIBUTEDEFINITIONSTRINGDBResponse provides response
// swagger:response attributedefinitionstringDBResponse
type ATTRIBUTEDEFINITIONSTRINGDBResponse struct {
	ATTRIBUTEDEFINITIONSTRINGDB
}

// ATTRIBUTEDEFINITIONSTRINGWOP is a ATTRIBUTEDEFINITIONSTRING without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ATTRIBUTEDEFINITIONSTRINGWOP struct {
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

var ATTRIBUTEDEFINITIONSTRING_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESCAttr",
	"IDENTIFIERAttr",
	"ISEDITABLEAttr",
	"LASTCHANGEAttr",
	"LONGNAMEAttr",
}

type BackRepoATTRIBUTEDEFINITIONSTRINGStruct struct {
	// stores ATTRIBUTEDEFINITIONSTRINGDB according to their gorm ID
	Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB map[uint]*ATTRIBUTEDEFINITIONSTRINGDB

	// stores ATTRIBUTEDEFINITIONSTRINGDB ID according to ATTRIBUTEDEFINITIONSTRING address
	Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID map[*models.ATTRIBUTEDEFINITIONSTRING]uint

	// stores ATTRIBUTEDEFINITIONSTRING according to their gorm ID
	Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr map[uint]*models.ATTRIBUTEDEFINITIONSTRING

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoATTRIBUTEDEFINITIONSTRING.stage
	return
}

func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) GetDB() *gorm.DB {
	return backRepoATTRIBUTEDEFINITIONSTRING.db
}

// GetATTRIBUTEDEFINITIONSTRINGDBFromATTRIBUTEDEFINITIONSTRINGPtr is a handy function to access the back repo instance from the stage instance
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) GetATTRIBUTEDEFINITIONSTRINGDBFromATTRIBUTEDEFINITIONSTRINGPtr(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) {
	id := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]
	attributedefinitionstringDB = backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[id]
	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseOne commits all staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for attributedefinitionstring := range stage.ATTRIBUTEDEFINITIONSTRINGs {
		backRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseOneInstance(attributedefinitionstring)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, attributedefinitionstring := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr {
		if _, ok := stage.ATTRIBUTEDEFINITIONSTRINGs[attributedefinitionstring]; !ok {
			backRepoATTRIBUTEDEFINITIONSTRING.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CommitDeleteInstance commits deletion of ATTRIBUTEDEFINITIONSTRING to the BackRepo
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CommitDeleteInstance(id uint) (Error error) {

	attributedefinitionstring := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[id]

	// attributedefinitionstring is not staged anymore, remove attributedefinitionstringDB
	attributedefinitionstringDB := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[id]
	query := backRepoATTRIBUTEDEFINITIONSTRING.db.Unscoped().Delete(&attributedefinitionstringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID, attributedefinitionstring)
	delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr, id)
	delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB, id)

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseOneInstance commits attributedefinitionstring staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CommitPhaseOneInstance(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) (Error error) {

	// check if the attributedefinitionstring is not commited yet
	if _, ok := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]; ok {
		return
	}

	// initiate attributedefinitionstring
	var attributedefinitionstringDB ATTRIBUTEDEFINITIONSTRINGDB
	attributedefinitionstringDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING(attributedefinitionstring)

	query := backRepoATTRIBUTEDEFINITIONSTRING.db.Create(&attributedefinitionstringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring] = attributedefinitionstringDB.ID
	backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID] = attributedefinitionstring
	backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[attributedefinitionstringDB.ID] = &attributedefinitionstringDB

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseTwo commits all staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, attributedefinitionstring := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr {
		backRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseTwoInstance(backRepo, idx, attributedefinitionstring)
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseTwoInstance commits {{structname }} of models.ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) (Error error) {

	// fetch matching attributedefinitionstringDB
	if attributedefinitionstringDB, ok := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[idx]; ok {

		attributedefinitionstringDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING(attributedefinitionstring)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value attributedefinitionstring.ALTERNATIVEID translates to updating the attributedefinitionstring.ALTERNATIVEIDID
		attributedefinitionstringDB.ALTERNATIVEIDID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionstring.ALTERNATIVEID != nil {
			if ALTERNATIVEIDId, ok := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDPtr_ALTERNATIVEIDDBID[attributedefinitionstring.ALTERNATIVEID]; ok {
				attributedefinitionstringDB.ALTERNATIVEIDID.Int64 = int64(ALTERNATIVEIDId)
				attributedefinitionstringDB.ALTERNATIVEIDID.Valid = true
			}
		} else {
			attributedefinitionstringDB.ALTERNATIVEIDID.Int64 = 0
			attributedefinitionstringDB.ALTERNATIVEIDID.Valid = true
		}

		// commit pointer value attributedefinitionstring.DEFAULTVALUE translates to updating the attributedefinitionstring.DEFAULTVALUEID
		attributedefinitionstringDB.DEFAULTVALUEID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionstring.DEFAULTVALUE != nil {
			if DEFAULTVALUEId, ok := backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEPtr_DEFAULTVALUEDBID[attributedefinitionstring.DEFAULTVALUE]; ok {
				attributedefinitionstringDB.DEFAULTVALUEID.Int64 = int64(DEFAULTVALUEId)
				attributedefinitionstringDB.DEFAULTVALUEID.Valid = true
			}
		} else {
			attributedefinitionstringDB.DEFAULTVALUEID.Int64 = 0
			attributedefinitionstringDB.DEFAULTVALUEID.Valid = true
		}

		// commit pointer value attributedefinitionstring.TYPE translates to updating the attributedefinitionstring.TYPEID
		attributedefinitionstringDB.TYPEID.Valid = true // allow for a 0 value (nil association)
		if attributedefinitionstring.TYPE != nil {
			if TYPEId, ok := backRepo.BackRepoREQTYPE.Map_REQTYPEPtr_REQTYPEDBID[attributedefinitionstring.TYPE]; ok {
				attributedefinitionstringDB.TYPEID.Int64 = int64(TYPEId)
				attributedefinitionstringDB.TYPEID.Valid = true
			}
		} else {
			attributedefinitionstringDB.TYPEID.Int64 = 0
			attributedefinitionstringDB.TYPEID.Valid = true
		}

		query := backRepoATTRIBUTEDEFINITIONSTRING.db.Save(&attributedefinitionstringDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ATTRIBUTEDEFINITIONSTRING intance %s", attributedefinitionstring.Name))
		return err
	}

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CheckoutPhaseOne() (Error error) {

	attributedefinitionstringDBArray := make([]ATTRIBUTEDEFINITIONSTRINGDB, 0)
	query := backRepoATTRIBUTEDEFINITIONSTRING.db.Find(&attributedefinitionstringDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	attributedefinitionstringInstancesToBeRemovedFromTheStage := make(map[*models.ATTRIBUTEDEFINITIONSTRING]any)
	for key, value := range backRepoATTRIBUTEDEFINITIONSTRING.stage.ATTRIBUTEDEFINITIONSTRINGs {
		attributedefinitionstringInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, attributedefinitionstringDB := range attributedefinitionstringDBArray {
		backRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseOneInstance(&attributedefinitionstringDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		attributedefinitionstring, ok := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]
		if ok {
			delete(attributedefinitionstringInstancesToBeRemovedFromTheStage, attributedefinitionstring)
		}
	}

	// remove from stage and back repo's 3 maps all attributedefinitionstrings that are not in the checkout
	for attributedefinitionstring := range attributedefinitionstringInstancesToBeRemovedFromTheStage {
		attributedefinitionstring.Unstage(backRepoATTRIBUTEDEFINITIONSTRING.GetStage())

		// remove instance from the back repo 3 maps
		attributedefinitionstringID := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]
		delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID, attributedefinitionstring)
		delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB, attributedefinitionstringID)
		delete(backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr, attributedefinitionstringID)
	}

	return
}

// CheckoutPhaseOneInstance takes a attributedefinitionstringDB that has been found in the DB, updates the backRepo and stages the
// models version of the attributedefinitionstringDB
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CheckoutPhaseOneInstance(attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) (Error error) {

	attributedefinitionstring, ok := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]
	if !ok {
		attributedefinitionstring = new(models.ATTRIBUTEDEFINITIONSTRING)

		backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID] = attributedefinitionstring
		backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring] = attributedefinitionstringDB.ID

		// append model store with the new element
		attributedefinitionstring.Name = attributedefinitionstringDB.Name_Data.String
		attributedefinitionstring.Stage(backRepoATTRIBUTEDEFINITIONSTRING.GetStage())
	}
	attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING(attributedefinitionstring)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	attributedefinitionstring.Stage(backRepoATTRIBUTEDEFINITIONSTRING.GetStage())

	// preserve pointer to attributedefinitionstringDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB)[attributedefinitionstringDB hold variable pointers
	attributedefinitionstringDB_Data := *attributedefinitionstringDB
	preservedPtrToATTRIBUTEDEFINITIONSTRING := &attributedefinitionstringDB_Data
	backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[attributedefinitionstringDB.ID] = preservedPtrToATTRIBUTEDEFINITIONSTRING

	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseTwo Checkouts all staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, attributedefinitionstringDB := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB {
		backRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseTwoInstance(backRepo, attributedefinitionstringDB)
	}
	return
}

// BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseTwoInstance Checkouts staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) (Error error) {

	attributedefinitionstring := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr[attributedefinitionstringDB.ID]

	attributedefinitionstringDB.DecodePointers(backRepo, attributedefinitionstring)

	return
}

func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) DecodePointers(backRepo *BackRepoStruct, attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) {

	// insertion point for checkout of pointer encoding
	// ALTERNATIVEID field
	attributedefinitionstring.ALTERNATIVEID = nil
	if attributedefinitionstringDB.ALTERNATIVEIDID.Int64 != 0 {
		attributedefinitionstring.ALTERNATIVEID = backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[uint(attributedefinitionstringDB.ALTERNATIVEIDID.Int64)]
	}
	// DEFAULTVALUE field
	attributedefinitionstring.DEFAULTVALUE = nil
	if attributedefinitionstringDB.DEFAULTVALUEID.Int64 != 0 {
		attributedefinitionstring.DEFAULTVALUE = backRepo.BackRepoDEFAULTVALUE.Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr[uint(attributedefinitionstringDB.DEFAULTVALUEID.Int64)]
	}
	// TYPE field
	attributedefinitionstring.TYPE = nil
	if attributedefinitionstringDB.TYPEID.Int64 != 0 {
		attributedefinitionstring.TYPE = backRepo.BackRepoREQTYPE.Map_REQTYPEDBID_REQTYPEPtr[uint(attributedefinitionstringDB.TYPEID.Int64)]
	}
	return
}

// CommitATTRIBUTEDEFINITIONSTRING allows commit of a single attributedefinitionstring (if already staged)
func (backRepo *BackRepoStruct) CommitATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) {
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseOneInstance(attributedefinitionstring)
	if id, ok := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]; ok {
		backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseTwoInstance(backRepo, id, attributedefinitionstring)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitATTRIBUTEDEFINITIONSTRING allows checkout of a single attributedefinitionstring (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) {
	// check if the attributedefinitionstring is staged
	if _, ok := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]; ok {

		if id, ok := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID[attributedefinitionstring]; ok {
			var attributedefinitionstringDB ATTRIBUTEDEFINITIONSTRINGDB
			attributedefinitionstringDB.ID = id

			if err := backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.db.First(&attributedefinitionstringDB, id).Error; err != nil {
				log.Fatalln("CheckoutATTRIBUTEDEFINITIONSTRING : Problem with getting object with id:", id)
			}
			backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseOneInstance(&attributedefinitionstringDB)
			backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseTwoInstance(backRepo, &attributedefinitionstringDB)
		}
	}
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) {
	// insertion point for fields commit

	attributedefinitionstringDB.Name_Data.String = attributedefinitionstring.Name
	attributedefinitionstringDB.Name_Data.Valid = true

	attributedefinitionstringDB.DESCAttr_Data.String = attributedefinitionstring.DESCAttr
	attributedefinitionstringDB.DESCAttr_Data.Valid = true

	attributedefinitionstringDB.IDENTIFIERAttr_Data.String = attributedefinitionstring.IDENTIFIERAttr
	attributedefinitionstringDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool = attributedefinitionstring.ISEDITABLEAttr
	attributedefinitionstringDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionstringDB.LASTCHANGEAttr_Data.String = attributedefinitionstring.LASTCHANGEAttr
	attributedefinitionstringDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionstringDB.LONGNAMEAttr_Data.String = attributedefinitionstring.LONGNAMEAttr
	attributedefinitionstringDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING_WOP
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRING_WOP(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING_WOP) {
	// insertion point for fields commit

	attributedefinitionstringDB.Name_Data.String = attributedefinitionstring.Name
	attributedefinitionstringDB.Name_Data.Valid = true

	attributedefinitionstringDB.DESCAttr_Data.String = attributedefinitionstring.DESCAttr
	attributedefinitionstringDB.DESCAttr_Data.Valid = true

	attributedefinitionstringDB.IDENTIFIERAttr_Data.String = attributedefinitionstring.IDENTIFIERAttr
	attributedefinitionstringDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool = attributedefinitionstring.ISEDITABLEAttr
	attributedefinitionstringDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionstringDB.LASTCHANGEAttr_Data.String = attributedefinitionstring.LASTCHANGEAttr
	attributedefinitionstringDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionstringDB.LONGNAMEAttr_Data.String = attributedefinitionstring.LONGNAMEAttr
	attributedefinitionstringDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRINGWOP
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRINGWOP(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRINGWOP) {
	// insertion point for fields commit

	attributedefinitionstringDB.Name_Data.String = attributedefinitionstring.Name
	attributedefinitionstringDB.Name_Data.Valid = true

	attributedefinitionstringDB.DESCAttr_Data.String = attributedefinitionstring.DESCAttr
	attributedefinitionstringDB.DESCAttr_Data.Valid = true

	attributedefinitionstringDB.IDENTIFIERAttr_Data.String = attributedefinitionstring.IDENTIFIERAttr
	attributedefinitionstringDB.IDENTIFIERAttr_Data.Valid = true

	attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool = attributedefinitionstring.ISEDITABLEAttr
	attributedefinitionstringDB.ISEDITABLEAttr_Data.Valid = true

	attributedefinitionstringDB.LASTCHANGEAttr_Data.String = attributedefinitionstring.LASTCHANGEAttr
	attributedefinitionstringDB.LASTCHANGEAttr_Data.Valid = true

	attributedefinitionstringDB.LONGNAMEAttr_Data.String = attributedefinitionstring.LONGNAMEAttr
	attributedefinitionstringDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) {
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionstring.Name = attributedefinitionstringDB.Name_Data.String
	attributedefinitionstring.DESCAttr = attributedefinitionstringDB.DESCAttr_Data.String
	attributedefinitionstring.IDENTIFIERAttr = attributedefinitionstringDB.IDENTIFIERAttr_Data.String
	attributedefinitionstring.ISEDITABLEAttr = attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionstring.LASTCHANGEAttr = attributedefinitionstringDB.LASTCHANGEAttr_Data.String
	attributedefinitionstring.LONGNAMEAttr = attributedefinitionstringDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING_WOP
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsToATTRIBUTEDEFINITIONSTRING_WOP(attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionstring.Name = attributedefinitionstringDB.Name_Data.String
	attributedefinitionstring.DESCAttr = attributedefinitionstringDB.DESCAttr_Data.String
	attributedefinitionstring.IDENTIFIERAttr = attributedefinitionstringDB.IDENTIFIERAttr_Data.String
	attributedefinitionstring.ISEDITABLEAttr = attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionstring.LASTCHANGEAttr = attributedefinitionstringDB.LASTCHANGEAttr_Data.String
	attributedefinitionstring.LONGNAMEAttr = attributedefinitionstringDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToATTRIBUTEDEFINITIONSTRINGWOP
func (attributedefinitionstringDB *ATTRIBUTEDEFINITIONSTRINGDB) CopyBasicFieldsToATTRIBUTEDEFINITIONSTRINGWOP(attributedefinitionstring *ATTRIBUTEDEFINITIONSTRINGWOP) {
	attributedefinitionstring.ID = int(attributedefinitionstringDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	attributedefinitionstring.Name = attributedefinitionstringDB.Name_Data.String
	attributedefinitionstring.DESCAttr = attributedefinitionstringDB.DESCAttr_Data.String
	attributedefinitionstring.IDENTIFIERAttr = attributedefinitionstringDB.IDENTIFIERAttr_Data.String
	attributedefinitionstring.ISEDITABLEAttr = attributedefinitionstringDB.ISEDITABLEAttr_Data.Bool
	attributedefinitionstring.LASTCHANGEAttr = attributedefinitionstringDB.LASTCHANGEAttr_Data.String
	attributedefinitionstring.LONGNAMEAttr = attributedefinitionstringDB.LONGNAMEAttr_Data.String
}

// Backup generates a json file from a slice of all ATTRIBUTEDEFINITIONSTRINGDB instances in the backrepo
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ATTRIBUTEDEFINITIONSTRINGDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTEDEFINITIONSTRINGDB, 0)
	for _, attributedefinitionstringDB := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB {
		forBackup = append(forBackup, attributedefinitionstringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json ATTRIBUTEDEFINITIONSTRING ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json ATTRIBUTEDEFINITIONSTRING file", err.Error())
	}
}

// Backup generates a json file from a slice of all ATTRIBUTEDEFINITIONSTRINGDB instances in the backrepo
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ATTRIBUTEDEFINITIONSTRINGDB, 0)
	for _, attributedefinitionstringDB := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB {
		forBackup = append(forBackup, attributedefinitionstringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ATTRIBUTEDEFINITIONSTRING")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ATTRIBUTEDEFINITIONSTRING_Fields, -1)
	for _, attributedefinitionstringDB := range forBackup {

		var attributedefinitionstringWOP ATTRIBUTEDEFINITIONSTRINGWOP
		attributedefinitionstringDB.CopyBasicFieldsToATTRIBUTEDEFINITIONSTRINGWOP(&attributedefinitionstringWOP)

		row := sh.AddRow()
		row.WriteStruct(&attributedefinitionstringWOP, -1)
	}
}

// RestoreXL from the "ATTRIBUTEDEFINITIONSTRING" sheet all ATTRIBUTEDEFINITIONSTRINGDB instances
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ATTRIBUTEDEFINITIONSTRING"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoATTRIBUTEDEFINITIONSTRING.rowVisitorATTRIBUTEDEFINITIONSTRING)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) rowVisitorATTRIBUTEDEFINITIONSTRING(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var attributedefinitionstringWOP ATTRIBUTEDEFINITIONSTRINGWOP
		row.ReadStruct(&attributedefinitionstringWOP)

		// add the unmarshalled struct to the stage
		attributedefinitionstringDB := new(ATTRIBUTEDEFINITIONSTRINGDB)
		attributedefinitionstringDB.CopyBasicFieldsFromATTRIBUTEDEFINITIONSTRINGWOP(&attributedefinitionstringWOP)

		attributedefinitionstringDB_ID_atBackupTime := attributedefinitionstringDB.ID
		attributedefinitionstringDB.ID = 0
		query := backRepoATTRIBUTEDEFINITIONSTRING.db.Create(attributedefinitionstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[attributedefinitionstringDB.ID] = attributedefinitionstringDB
		BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID[attributedefinitionstringDB_ID_atBackupTime] = attributedefinitionstringDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ATTRIBUTEDEFINITIONSTRINGDB.json" in dirPath that stores an array
// of ATTRIBUTEDEFINITIONSTRINGDB and stores it in the database
// the map BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID is updated accordingly
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ATTRIBUTEDEFINITIONSTRINGDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json ATTRIBUTEDEFINITIONSTRING file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ATTRIBUTEDEFINITIONSTRINGDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB
	for _, attributedefinitionstringDB := range forRestore {

		attributedefinitionstringDB_ID_atBackupTime := attributedefinitionstringDB.ID
		attributedefinitionstringDB.ID = 0
		query := backRepoATTRIBUTEDEFINITIONSTRING.db.Create(attributedefinitionstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[attributedefinitionstringDB.ID] = attributedefinitionstringDB
		BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID[attributedefinitionstringDB_ID_atBackupTime] = attributedefinitionstringDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json ATTRIBUTEDEFINITIONSTRING file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ATTRIBUTEDEFINITIONSTRING>id_atBckpTime_newID
// to compute new index
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) RestorePhaseTwo() {

	for _, attributedefinitionstringDB := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB {

		// next line of code is to avert unused variable compilation error
		_ = attributedefinitionstringDB

		// insertion point for reindexing pointers encoding
		// reindexing ALTERNATIVEID field
		if attributedefinitionstringDB.ALTERNATIVEIDID.Int64 != 0 {
			attributedefinitionstringDB.ALTERNATIVEIDID.Int64 = int64(BackRepoALTERNATIVEIDid_atBckpTime_newID[uint(attributedefinitionstringDB.ALTERNATIVEIDID.Int64)])
			attributedefinitionstringDB.ALTERNATIVEIDID.Valid = true
		}

		// reindexing DEFAULTVALUE field
		if attributedefinitionstringDB.DEFAULTVALUEID.Int64 != 0 {
			attributedefinitionstringDB.DEFAULTVALUEID.Int64 = int64(BackRepoDEFAULTVALUEid_atBckpTime_newID[uint(attributedefinitionstringDB.DEFAULTVALUEID.Int64)])
			attributedefinitionstringDB.DEFAULTVALUEID.Valid = true
		}

		// reindexing TYPE field
		if attributedefinitionstringDB.TYPEID.Int64 != 0 {
			attributedefinitionstringDB.TYPEID.Int64 = int64(BackRepoREQTYPEid_atBckpTime_newID[uint(attributedefinitionstringDB.TYPEID.Int64)])
			attributedefinitionstringDB.TYPEID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoATTRIBUTEDEFINITIONSTRING.db.Model(attributedefinitionstringDB).Updates(*attributedefinitionstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoATTRIBUTEDEFINITIONSTRING.ResetReversePointers commits all staged instances of ATTRIBUTEDEFINITIONSTRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, attributedefinitionstring := range backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr {
		backRepoATTRIBUTEDEFINITIONSTRING.ResetReversePointersInstance(backRepo, idx, attributedefinitionstring)
	}

	return
}

func (backRepoATTRIBUTEDEFINITIONSTRING *BackRepoATTRIBUTEDEFINITIONSTRINGStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, attributedefinitionstring *models.ATTRIBUTEDEFINITIONSTRING) (Error error) {

	// fetch matching attributedefinitionstringDB
	if attributedefinitionstringDB, ok := backRepoATTRIBUTEDEFINITIONSTRING.Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB[idx]; ok {
		_ = attributedefinitionstringDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoATTRIBUTEDEFINITIONSTRINGid_atBckpTime_newID map[uint]uint
