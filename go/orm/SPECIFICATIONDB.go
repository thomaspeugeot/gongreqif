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
var dummy_SPECIFICATION_sql sql.NullBool
var dummy_SPECIFICATION_time time.Duration
var dummy_SPECIFICATION_sort sort.Float64Slice

// SPECIFICATIONAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model specificationAPI
type SPECIFICATIONAPI struct {
	gorm.Model

	models.SPECIFICATION_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SPECIFICATIONPointersEncoding SPECIFICATIONPointersEncoding
}

// SPECIFICATIONPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SPECIFICATIONPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ALTERNATIVEID is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ALTERNATIVEIDID sql.NullInt64

	// field VALUES is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	VALUESID sql.NullInt64

	// field CHILDREN is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	CHILDRENID sql.NullInt64

	// field TYPE is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	TYPEID sql.NullInt64
}

// SPECIFICATIONDB describes a specification in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model specificationDB
type SPECIFICATIONDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field specificationDB.Name
	Name_Data sql.NullString

	// Declation for basic field specificationDB.DESCAttr
	DESCAttr_Data sql.NullString

	// Declation for basic field specificationDB.IDENTIFIERAttr
	IDENTIFIERAttr_Data sql.NullString

	// Declation for basic field specificationDB.LASTCHANGEAttr
	LASTCHANGEAttr_Data sql.NullString

	// Declation for basic field specificationDB.LONGNAMEAttr
	LONGNAMEAttr_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SPECIFICATIONPointersEncoding
}

// SPECIFICATIONDBs arrays specificationDBs
// swagger:response specificationDBsResponse
type SPECIFICATIONDBs []SPECIFICATIONDB

// SPECIFICATIONDBResponse provides response
// swagger:response specificationDBResponse
type SPECIFICATIONDBResponse struct {
	SPECIFICATIONDB
}

// SPECIFICATIONWOP is a SPECIFICATION without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SPECIFICATIONWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESCAttr string `xlsx:"2"`

	IDENTIFIERAttr string `xlsx:"3"`

	LASTCHANGEAttr string `xlsx:"4"`

	LONGNAMEAttr string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var SPECIFICATION_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESCAttr",
	"IDENTIFIERAttr",
	"LASTCHANGEAttr",
	"LONGNAMEAttr",
}

type BackRepoSPECIFICATIONStruct struct {
	// stores SPECIFICATIONDB according to their gorm ID
	Map_SPECIFICATIONDBID_SPECIFICATIONDB map[uint]*SPECIFICATIONDB

	// stores SPECIFICATIONDB ID according to SPECIFICATION address
	Map_SPECIFICATIONPtr_SPECIFICATIONDBID map[*models.SPECIFICATION]uint

	// stores SPECIFICATION according to their gorm ID
	Map_SPECIFICATIONDBID_SPECIFICATIONPtr map[uint]*models.SPECIFICATION

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSPECIFICATION.stage
	return
}

func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) GetDB() *gorm.DB {
	return backRepoSPECIFICATION.db
}

// GetSPECIFICATIONDBFromSPECIFICATIONPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) GetSPECIFICATIONDBFromSPECIFICATIONPtr(specification *models.SPECIFICATION) (specificationDB *SPECIFICATIONDB) {
	id := backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]
	specificationDB = backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[id]
	return
}

// BackRepoSPECIFICATION.CommitPhaseOne commits all staged instances of SPECIFICATION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for specification := range stage.SPECIFICATIONs {
		backRepoSPECIFICATION.CommitPhaseOneInstance(specification)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, specification := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr {
		if _, ok := stage.SPECIFICATIONs[specification]; !ok {
			backRepoSPECIFICATION.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSPECIFICATION.CommitDeleteInstance commits deletion of SPECIFICATION to the BackRepo
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CommitDeleteInstance(id uint) (Error error) {

	specification := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[id]

	// specification is not staged anymore, remove specificationDB
	specificationDB := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[id]
	query := backRepoSPECIFICATION.db.Unscoped().Delete(&specificationDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID, specification)
	delete(backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr, id)
	delete(backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB, id)

	return
}

// BackRepoSPECIFICATION.CommitPhaseOneInstance commits specification staged instances of SPECIFICATION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CommitPhaseOneInstance(specification *models.SPECIFICATION) (Error error) {

	// check if the specification is not commited yet
	if _, ok := backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]; ok {
		return
	}

	// initiate specification
	var specificationDB SPECIFICATIONDB
	specificationDB.CopyBasicFieldsFromSPECIFICATION(specification)

	query := backRepoSPECIFICATION.db.Create(&specificationDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification] = specificationDB.ID
	backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID] = specification
	backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[specificationDB.ID] = &specificationDB

	return
}

// BackRepoSPECIFICATION.CommitPhaseTwo commits all staged instances of SPECIFICATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, specification := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr {
		backRepoSPECIFICATION.CommitPhaseTwoInstance(backRepo, idx, specification)
	}

	return
}

// BackRepoSPECIFICATION.CommitPhaseTwoInstance commits {{structname }} of models.SPECIFICATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, specification *models.SPECIFICATION) (Error error) {

	// fetch matching specificationDB
	if specificationDB, ok := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[idx]; ok {

		specificationDB.CopyBasicFieldsFromSPECIFICATION(specification)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value specification.ALTERNATIVEID translates to updating the specification.ALTERNATIVEIDID
		specificationDB.ALTERNATIVEIDID.Valid = true // allow for a 0 value (nil association)
		if specification.ALTERNATIVEID != nil {
			if ALTERNATIVEIDId, ok := backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDPtr_ALTERNATIVEIDDBID[specification.ALTERNATIVEID]; ok {
				specificationDB.ALTERNATIVEIDID.Int64 = int64(ALTERNATIVEIDId)
				specificationDB.ALTERNATIVEIDID.Valid = true
			}
		} else {
			specificationDB.ALTERNATIVEIDID.Int64 = 0
			specificationDB.ALTERNATIVEIDID.Valid = true
		}

		// commit pointer value specification.VALUES translates to updating the specification.VALUESID
		specificationDB.VALUESID.Valid = true // allow for a 0 value (nil association)
		if specification.VALUES != nil {
			if VALUESId, ok := backRepo.BackRepoVALUES.Map_VALUESPtr_VALUESDBID[specification.VALUES]; ok {
				specificationDB.VALUESID.Int64 = int64(VALUESId)
				specificationDB.VALUESID.Valid = true
			}
		} else {
			specificationDB.VALUESID.Int64 = 0
			specificationDB.VALUESID.Valid = true
		}

		// commit pointer value specification.CHILDREN translates to updating the specification.CHILDRENID
		specificationDB.CHILDRENID.Valid = true // allow for a 0 value (nil association)
		if specification.CHILDREN != nil {
			if CHILDRENId, ok := backRepo.BackRepoCHILDREN.Map_CHILDRENPtr_CHILDRENDBID[specification.CHILDREN]; ok {
				specificationDB.CHILDRENID.Int64 = int64(CHILDRENId)
				specificationDB.CHILDRENID.Valid = true
			}
		} else {
			specificationDB.CHILDRENID.Int64 = 0
			specificationDB.CHILDRENID.Valid = true
		}

		// commit pointer value specification.TYPE translates to updating the specification.TYPEID
		specificationDB.TYPEID.Valid = true // allow for a 0 value (nil association)
		if specification.TYPE != nil {
			if TYPEId, ok := backRepo.BackRepoREQTYPE.Map_REQTYPEPtr_REQTYPEDBID[specification.TYPE]; ok {
				specificationDB.TYPEID.Int64 = int64(TYPEId)
				specificationDB.TYPEID.Valid = true
			}
		} else {
			specificationDB.TYPEID.Int64 = 0
			specificationDB.TYPEID.Valid = true
		}

		query := backRepoSPECIFICATION.db.Save(&specificationDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SPECIFICATION intance %s", specification.Name))
		return err
	}

	return
}

// BackRepoSPECIFICATION.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CheckoutPhaseOne() (Error error) {

	specificationDBArray := make([]SPECIFICATIONDB, 0)
	query := backRepoSPECIFICATION.db.Find(&specificationDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	specificationInstancesToBeRemovedFromTheStage := make(map[*models.SPECIFICATION]any)
	for key, value := range backRepoSPECIFICATION.stage.SPECIFICATIONs {
		specificationInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, specificationDB := range specificationDBArray {
		backRepoSPECIFICATION.CheckoutPhaseOneInstance(&specificationDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		specification, ok := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]
		if ok {
			delete(specificationInstancesToBeRemovedFromTheStage, specification)
		}
	}

	// remove from stage and back repo's 3 maps all specifications that are not in the checkout
	for specification := range specificationInstancesToBeRemovedFromTheStage {
		specification.Unstage(backRepoSPECIFICATION.GetStage())

		// remove instance from the back repo 3 maps
		specificationID := backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]
		delete(backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID, specification)
		delete(backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB, specificationID)
		delete(backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr, specificationID)
	}

	return
}

// CheckoutPhaseOneInstance takes a specificationDB that has been found in the DB, updates the backRepo and stages the
// models version of the specificationDB
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CheckoutPhaseOneInstance(specificationDB *SPECIFICATIONDB) (Error error) {

	specification, ok := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]
	if !ok {
		specification = new(models.SPECIFICATION)

		backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID] = specification
		backRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification] = specificationDB.ID

		// append model store with the new element
		specification.Name = specificationDB.Name_Data.String
		specification.Stage(backRepoSPECIFICATION.GetStage())
	}
	specificationDB.CopyBasicFieldsToSPECIFICATION(specification)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	specification.Stage(backRepoSPECIFICATION.GetStage())

	// preserve pointer to specificationDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SPECIFICATIONDBID_SPECIFICATIONDB)[specificationDB hold variable pointers
	specificationDB_Data := *specificationDB
	preservedPtrToSPECIFICATION := &specificationDB_Data
	backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[specificationDB.ID] = preservedPtrToSPECIFICATION

	return
}

// BackRepoSPECIFICATION.CheckoutPhaseTwo Checkouts all staged instances of SPECIFICATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, specificationDB := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {
		backRepoSPECIFICATION.CheckoutPhaseTwoInstance(backRepo, specificationDB)
	}
	return
}

// BackRepoSPECIFICATION.CheckoutPhaseTwoInstance Checkouts staged instances of SPECIFICATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, specificationDB *SPECIFICATIONDB) (Error error) {

	specification := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]

	specificationDB.DecodePointers(backRepo, specification)

	return
}

func (specificationDB *SPECIFICATIONDB) DecodePointers(backRepo *BackRepoStruct, specification *models.SPECIFICATION) {

	// insertion point for checkout of pointer encoding
	// ALTERNATIVEID field
	specification.ALTERNATIVEID = nil
	if specificationDB.ALTERNATIVEIDID.Int64 != 0 {
		specification.ALTERNATIVEID = backRepo.BackRepoALTERNATIVEID.Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr[uint(specificationDB.ALTERNATIVEIDID.Int64)]
	}
	// VALUES field
	specification.VALUES = nil
	if specificationDB.VALUESID.Int64 != 0 {
		specification.VALUES = backRepo.BackRepoVALUES.Map_VALUESDBID_VALUESPtr[uint(specificationDB.VALUESID.Int64)]
	}
	// CHILDREN field
	specification.CHILDREN = nil
	if specificationDB.CHILDRENID.Int64 != 0 {
		specification.CHILDREN = backRepo.BackRepoCHILDREN.Map_CHILDRENDBID_CHILDRENPtr[uint(specificationDB.CHILDRENID.Int64)]
	}
	// TYPE field
	specification.TYPE = nil
	if specificationDB.TYPEID.Int64 != 0 {
		specification.TYPE = backRepo.BackRepoREQTYPE.Map_REQTYPEDBID_REQTYPEPtr[uint(specificationDB.TYPEID.Int64)]
	}
	return
}

// CommitSPECIFICATION allows commit of a single specification (if already staged)
func (backRepo *BackRepoStruct) CommitSPECIFICATION(specification *models.SPECIFICATION) {
	backRepo.BackRepoSPECIFICATION.CommitPhaseOneInstance(specification)
	if id, ok := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]; ok {
		backRepo.BackRepoSPECIFICATION.CommitPhaseTwoInstance(backRepo, id, specification)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSPECIFICATION allows checkout of a single specification (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSPECIFICATION(specification *models.SPECIFICATION) {
	// check if the specification is staged
	if _, ok := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]; ok {

		if id, ok := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONPtr_SPECIFICATIONDBID[specification]; ok {
			var specificationDB SPECIFICATIONDB
			specificationDB.ID = id

			if err := backRepo.BackRepoSPECIFICATION.db.First(&specificationDB, id).Error; err != nil {
				log.Fatalln("CheckoutSPECIFICATION : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSPECIFICATION.CheckoutPhaseOneInstance(&specificationDB)
			backRepo.BackRepoSPECIFICATION.CheckoutPhaseTwoInstance(backRepo, &specificationDB)
		}
	}
}

// CopyBasicFieldsFromSPECIFICATION
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsFromSPECIFICATION(specification *models.SPECIFICATION) {
	// insertion point for fields commit

	specificationDB.Name_Data.String = specification.Name
	specificationDB.Name_Data.Valid = true

	specificationDB.DESCAttr_Data.String = specification.DESCAttr
	specificationDB.DESCAttr_Data.Valid = true

	specificationDB.IDENTIFIERAttr_Data.String = specification.IDENTIFIERAttr
	specificationDB.IDENTIFIERAttr_Data.Valid = true

	specificationDB.LASTCHANGEAttr_Data.String = specification.LASTCHANGEAttr
	specificationDB.LASTCHANGEAttr_Data.Valid = true

	specificationDB.LONGNAMEAttr_Data.String = specification.LONGNAMEAttr
	specificationDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromSPECIFICATION_WOP
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsFromSPECIFICATION_WOP(specification *models.SPECIFICATION_WOP) {
	// insertion point for fields commit

	specificationDB.Name_Data.String = specification.Name
	specificationDB.Name_Data.Valid = true

	specificationDB.DESCAttr_Data.String = specification.DESCAttr
	specificationDB.DESCAttr_Data.Valid = true

	specificationDB.IDENTIFIERAttr_Data.String = specification.IDENTIFIERAttr
	specificationDB.IDENTIFIERAttr_Data.Valid = true

	specificationDB.LASTCHANGEAttr_Data.String = specification.LASTCHANGEAttr
	specificationDB.LASTCHANGEAttr_Data.Valid = true

	specificationDB.LONGNAMEAttr_Data.String = specification.LONGNAMEAttr
	specificationDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsFromSPECIFICATIONWOP
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsFromSPECIFICATIONWOP(specification *SPECIFICATIONWOP) {
	// insertion point for fields commit

	specificationDB.Name_Data.String = specification.Name
	specificationDB.Name_Data.Valid = true

	specificationDB.DESCAttr_Data.String = specification.DESCAttr
	specificationDB.DESCAttr_Data.Valid = true

	specificationDB.IDENTIFIERAttr_Data.String = specification.IDENTIFIERAttr
	specificationDB.IDENTIFIERAttr_Data.Valid = true

	specificationDB.LASTCHANGEAttr_Data.String = specification.LASTCHANGEAttr
	specificationDB.LASTCHANGEAttr_Data.Valid = true

	specificationDB.LONGNAMEAttr_Data.String = specification.LONGNAMEAttr
	specificationDB.LONGNAMEAttr_Data.Valid = true
}

// CopyBasicFieldsToSPECIFICATION
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsToSPECIFICATION(specification *models.SPECIFICATION) {
	// insertion point for checkout of basic fields (back repo to stage)
	specification.Name = specificationDB.Name_Data.String
	specification.DESCAttr = specificationDB.DESCAttr_Data.String
	specification.IDENTIFIERAttr = specificationDB.IDENTIFIERAttr_Data.String
	specification.LASTCHANGEAttr = specificationDB.LASTCHANGEAttr_Data.String
	specification.LONGNAMEAttr = specificationDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToSPECIFICATION_WOP
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsToSPECIFICATION_WOP(specification *models.SPECIFICATION_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	specification.Name = specificationDB.Name_Data.String
	specification.DESCAttr = specificationDB.DESCAttr_Data.String
	specification.IDENTIFIERAttr = specificationDB.IDENTIFIERAttr_Data.String
	specification.LASTCHANGEAttr = specificationDB.LASTCHANGEAttr_Data.String
	specification.LONGNAMEAttr = specificationDB.LONGNAMEAttr_Data.String
}

// CopyBasicFieldsToSPECIFICATIONWOP
func (specificationDB *SPECIFICATIONDB) CopyBasicFieldsToSPECIFICATIONWOP(specification *SPECIFICATIONWOP) {
	specification.ID = int(specificationDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	specification.Name = specificationDB.Name_Data.String
	specification.DESCAttr = specificationDB.DESCAttr_Data.String
	specification.IDENTIFIERAttr = specificationDB.IDENTIFIERAttr_Data.String
	specification.LASTCHANGEAttr = specificationDB.LASTCHANGEAttr_Data.String
	specification.LONGNAMEAttr = specificationDB.LONGNAMEAttr_Data.String
}

// Backup generates a json file from a slice of all SPECIFICATIONDB instances in the backrepo
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SPECIFICATIONDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPECIFICATIONDB, 0)
	for _, specificationDB := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {
		forBackup = append(forBackup, specificationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SPECIFICATION ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SPECIFICATION file", err.Error())
	}
}

// Backup generates a json file from a slice of all SPECIFICATIONDB instances in the backrepo
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPECIFICATIONDB, 0)
	for _, specificationDB := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {
		forBackup = append(forBackup, specificationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SPECIFICATION")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SPECIFICATION_Fields, -1)
	for _, specificationDB := range forBackup {

		var specificationWOP SPECIFICATIONWOP
		specificationDB.CopyBasicFieldsToSPECIFICATIONWOP(&specificationWOP)

		row := sh.AddRow()
		row.WriteStruct(&specificationWOP, -1)
	}
}

// RestoreXL from the "SPECIFICATION" sheet all SPECIFICATIONDB instances
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSPECIFICATIONid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SPECIFICATION"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSPECIFICATION.rowVisitorSPECIFICATION)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) rowVisitorSPECIFICATION(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var specificationWOP SPECIFICATIONWOP
		row.ReadStruct(&specificationWOP)

		// add the unmarshalled struct to the stage
		specificationDB := new(SPECIFICATIONDB)
		specificationDB.CopyBasicFieldsFromSPECIFICATIONWOP(&specificationWOP)

		specificationDB_ID_atBackupTime := specificationDB.ID
		specificationDB.ID = 0
		query := backRepoSPECIFICATION.db.Create(specificationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[specificationDB.ID] = specificationDB
		BackRepoSPECIFICATIONid_atBckpTime_newID[specificationDB_ID_atBackupTime] = specificationDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SPECIFICATIONDB.json" in dirPath that stores an array
// of SPECIFICATIONDB and stores it in the database
// the map BackRepoSPECIFICATIONid_atBckpTime_newID is updated accordingly
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSPECIFICATIONid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SPECIFICATIONDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SPECIFICATION file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SPECIFICATIONDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SPECIFICATIONDBID_SPECIFICATIONDB
	for _, specificationDB := range forRestore {

		specificationDB_ID_atBackupTime := specificationDB.ID
		specificationDB.ID = 0
		query := backRepoSPECIFICATION.db.Create(specificationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[specificationDB.ID] = specificationDB
		BackRepoSPECIFICATIONid_atBckpTime_newID[specificationDB_ID_atBackupTime] = specificationDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SPECIFICATION file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SPECIFICATION>id_atBckpTime_newID
// to compute new index
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) RestorePhaseTwo() {

	for _, specificationDB := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {

		// next line of code is to avert unused variable compilation error
		_ = specificationDB

		// insertion point for reindexing pointers encoding
		// reindexing ALTERNATIVEID field
		if specificationDB.ALTERNATIVEIDID.Int64 != 0 {
			specificationDB.ALTERNATIVEIDID.Int64 = int64(BackRepoALTERNATIVEIDid_atBckpTime_newID[uint(specificationDB.ALTERNATIVEIDID.Int64)])
			specificationDB.ALTERNATIVEIDID.Valid = true
		}

		// reindexing VALUES field
		if specificationDB.VALUESID.Int64 != 0 {
			specificationDB.VALUESID.Int64 = int64(BackRepoVALUESid_atBckpTime_newID[uint(specificationDB.VALUESID.Int64)])
			specificationDB.VALUESID.Valid = true
		}

		// reindexing CHILDREN field
		if specificationDB.CHILDRENID.Int64 != 0 {
			specificationDB.CHILDRENID.Int64 = int64(BackRepoCHILDRENid_atBckpTime_newID[uint(specificationDB.CHILDRENID.Int64)])
			specificationDB.CHILDRENID.Valid = true
		}

		// reindexing TYPE field
		if specificationDB.TYPEID.Int64 != 0 {
			specificationDB.TYPEID.Int64 = int64(BackRepoREQTYPEid_atBckpTime_newID[uint(specificationDB.TYPEID.Int64)])
			specificationDB.TYPEID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoSPECIFICATION.db.Model(specificationDB).Updates(*specificationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSPECIFICATION.ResetReversePointers commits all staged instances of SPECIFICATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, specification := range backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr {
		backRepoSPECIFICATION.ResetReversePointersInstance(backRepo, idx, specification)
	}

	return
}

func (backRepoSPECIFICATION *BackRepoSPECIFICATIONStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, specification *models.SPECIFICATION) (Error error) {

	// fetch matching specificationDB
	if specificationDB, ok := backRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB[idx]; ok {
		_ = specificationDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSPECIFICATIONid_atBckpTime_newID map[uint]uint
