// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/thomaspeugeot/gongreqif/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoALTERNATIVEID BackRepoALTERNATIVEIDStruct

	BackRepoATTRIBUTEDEFINITIONBOOLEAN BackRepoATTRIBUTEDEFINITIONBOOLEANStruct

	BackRepoATTRIBUTEDEFINITIONDATE BackRepoATTRIBUTEDEFINITIONDATEStruct

	BackRepoATTRIBUTEDEFINITIONENUMERATION BackRepoATTRIBUTEDEFINITIONENUMERATIONStruct

	BackRepoATTRIBUTEDEFINITIONINTEGER BackRepoATTRIBUTEDEFINITIONINTEGERStruct

	BackRepoATTRIBUTEDEFINITIONREAL BackRepoATTRIBUTEDEFINITIONREALStruct

	BackRepoATTRIBUTEDEFINITIONSTRING BackRepoATTRIBUTEDEFINITIONSTRINGStruct

	BackRepoATTRIBUTEDEFINITIONXHTML BackRepoATTRIBUTEDEFINITIONXHTMLStruct

	BackRepoATTRIBUTEVALUEBOOLEAN BackRepoATTRIBUTEVALUEBOOLEANStruct

	BackRepoATTRIBUTEVALUEDATE BackRepoATTRIBUTEVALUEDATEStruct

	BackRepoATTRIBUTEVALUEENUMERATION BackRepoATTRIBUTEVALUEENUMERATIONStruct

	BackRepoATTRIBUTEVALUEINTEGER BackRepoATTRIBUTEVALUEINTEGERStruct

	BackRepoATTRIBUTEVALUEREAL BackRepoATTRIBUTEVALUEREALStruct

	BackRepoATTRIBUTEVALUESTRING BackRepoATTRIBUTEVALUESTRINGStruct

	BackRepoATTRIBUTEVALUEXHTML BackRepoATTRIBUTEVALUEXHTMLStruct

	BackRepoCHILDREN BackRepoCHILDRENStruct

	BackRepoCORECONTENT BackRepoCORECONTENTStruct

	BackRepoDATATYPEDEFINITIONBOOLEAN BackRepoDATATYPEDEFINITIONBOOLEANStruct

	BackRepoDATATYPEDEFINITIONDATE BackRepoDATATYPEDEFINITIONDATEStruct

	BackRepoDATATYPEDEFINITIONENUMERATION BackRepoDATATYPEDEFINITIONENUMERATIONStruct

	BackRepoDATATYPEDEFINITIONINTEGER BackRepoDATATYPEDEFINITIONINTEGERStruct

	BackRepoDATATYPEDEFINITIONREAL BackRepoDATATYPEDEFINITIONREALStruct

	BackRepoDATATYPEDEFINITIONSTRING BackRepoDATATYPEDEFINITIONSTRINGStruct

	BackRepoDATATYPEDEFINITIONXHTML BackRepoDATATYPEDEFINITIONXHTMLStruct

	BackRepoDATATYPES BackRepoDATATYPESStruct

	BackRepoDEFAULTVALUE BackRepoDEFAULTVALUEStruct

	BackRepoDEFINITION BackRepoDEFINITIONStruct

	BackRepoEDITABLEATTS BackRepoEDITABLEATTSStruct

	BackRepoEMBEDDEDVALUE BackRepoEMBEDDEDVALUEStruct

	BackRepoENUMVALUE BackRepoENUMVALUEStruct

	BackRepoOBJECT BackRepoOBJECTStruct

	BackRepoPROPERTIES BackRepoPROPERTIESStruct

	BackRepoRELATIONGROUP BackRepoRELATIONGROUPStruct

	BackRepoRELATIONGROUPTYPE BackRepoRELATIONGROUPTYPEStruct

	BackRepoREQIF BackRepoREQIFStruct

	BackRepoREQIFCONTENT BackRepoREQIFCONTENTStruct

	BackRepoREQIFHEADER BackRepoREQIFHEADERStruct

	BackRepoREQIFTOOLEXTENSION BackRepoREQIFTOOLEXTENSIONStruct

	BackRepoREQIFTYPE BackRepoREQIFTYPEStruct

	BackRepoSOURCE BackRepoSOURCEStruct

	BackRepoSOURCESPECIFICATION BackRepoSOURCESPECIFICATIONStruct

	BackRepoSPECATTRIBUTES BackRepoSPECATTRIBUTESStruct

	BackRepoSPECHIERARCHY BackRepoSPECHIERARCHYStruct

	BackRepoSPECIFICATION BackRepoSPECIFICATIONStruct

	BackRepoSPECIFICATIONS BackRepoSPECIFICATIONSStruct

	BackRepoSPECIFICATIONTYPE BackRepoSPECIFICATIONTYPEStruct

	BackRepoSPECIFIEDVALUES BackRepoSPECIFIEDVALUESStruct

	BackRepoSPECOBJECT BackRepoSPECOBJECTStruct

	BackRepoSPECOBJECTS BackRepoSPECOBJECTSStruct

	BackRepoSPECOBJECTTYPE BackRepoSPECOBJECTTYPEStruct

	BackRepoSPECRELATION BackRepoSPECRELATIONStruct

	BackRepoSPECRELATIONGROUPS BackRepoSPECRELATIONGROUPSStruct

	BackRepoSPECRELATIONS BackRepoSPECRELATIONSStruct

	BackRepoSPECRELATIONTYPE BackRepoSPECRELATIONTYPEStruct

	BackRepoSPECTYPES BackRepoSPECTYPESStruct

	BackRepoTARGET BackRepoTARGETStruct

	BackRepoTARGETSPECIFICATION BackRepoTARGETSPECIFICATIONStruct

	BackRepoTHEHEADER BackRepoTHEHEADERStruct

	BackRepoTOOLEXTENSIONS BackRepoTOOLEXTENSIONSStruct

	BackRepoVALUES BackRepoVALUESStruct

	BackRepoXHTMLCONTENT BackRepoXHTMLCONTENTStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&ALTERNATIVEIDDB{},
		&ATTRIBUTEDEFINITIONBOOLEANDB{},
		&ATTRIBUTEDEFINITIONDATEDB{},
		&ATTRIBUTEDEFINITIONENUMERATIONDB{},
		&ATTRIBUTEDEFINITIONINTEGERDB{},
		&ATTRIBUTEDEFINITIONREALDB{},
		&ATTRIBUTEDEFINITIONSTRINGDB{},
		&ATTRIBUTEDEFINITIONXHTMLDB{},
		&ATTRIBUTEVALUEBOOLEANDB{},
		&ATTRIBUTEVALUEDATEDB{},
		&ATTRIBUTEVALUEENUMERATIONDB{},
		&ATTRIBUTEVALUEINTEGERDB{},
		&ATTRIBUTEVALUEREALDB{},
		&ATTRIBUTEVALUESTRINGDB{},
		&ATTRIBUTEVALUEXHTMLDB{},
		&CHILDRENDB{},
		&CORECONTENTDB{},
		&DATATYPEDEFINITIONBOOLEANDB{},
		&DATATYPEDEFINITIONDATEDB{},
		&DATATYPEDEFINITIONENUMERATIONDB{},
		&DATATYPEDEFINITIONINTEGERDB{},
		&DATATYPEDEFINITIONREALDB{},
		&DATATYPEDEFINITIONSTRINGDB{},
		&DATATYPEDEFINITIONXHTMLDB{},
		&DATATYPESDB{},
		&DEFAULTVALUEDB{},
		&DEFINITIONDB{},
		&EDITABLEATTSDB{},
		&EMBEDDEDVALUEDB{},
		&ENUMVALUEDB{},
		&OBJECTDB{},
		&PROPERTIESDB{},
		&RELATIONGROUPDB{},
		&RELATIONGROUPTYPEDB{},
		&REQIFDB{},
		&REQIFCONTENTDB{},
		&REQIFHEADERDB{},
		&REQIFTOOLEXTENSIONDB{},
		&REQIFTYPEDB{},
		&SOURCEDB{},
		&SOURCESPECIFICATIONDB{},
		&SPECATTRIBUTESDB{},
		&SPECHIERARCHYDB{},
		&SPECIFICATIONDB{},
		&SPECIFICATIONSDB{},
		&SPECIFICATIONTYPEDB{},
		&SPECIFIEDVALUESDB{},
		&SPECOBJECTDB{},
		&SPECOBJECTSDB{},
		&SPECOBJECTTYPEDB{},
		&SPECRELATIONDB{},
		&SPECRELATIONGROUPSDB{},
		&SPECRELATIONSDB{},
		&SPECRELATIONTYPEDB{},
		&SPECTYPESDB{},
		&TARGETDB{},
		&TARGETSPECIFICATIONDB{},
		&THEHEADERDB{},
		&TOOLEXTENSIONSDB{},
		&VALUESDB{},
		&XHTMLCONTENTDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoALTERNATIVEID = BackRepoALTERNATIVEIDStruct{
		Map_ALTERNATIVEIDDBID_ALTERNATIVEIDPtr: make(map[uint]*models.ALTERNATIVEID, 0),
		Map_ALTERNATIVEIDDBID_ALTERNATIVEIDDB:  make(map[uint]*ALTERNATIVEIDDB, 0),
		Map_ALTERNATIVEIDPtr_ALTERNATIVEIDDBID: make(map[*models.ALTERNATIVEID]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN = BackRepoATTRIBUTEDEFINITIONBOOLEANStruct{
		Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONBOOLEAN, 0),
		Map_ATTRIBUTEDEFINITIONBOOLEANDBID_ATTRIBUTEDEFINITIONBOOLEANDB:  make(map[uint]*ATTRIBUTEDEFINITIONBOOLEANDB, 0),
		Map_ATTRIBUTEDEFINITIONBOOLEANPtr_ATTRIBUTEDEFINITIONBOOLEANDBID: make(map[*models.ATTRIBUTEDEFINITIONBOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE = BackRepoATTRIBUTEDEFINITIONDATEStruct{
		Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONDATE, 0),
		Map_ATTRIBUTEDEFINITIONDATEDBID_ATTRIBUTEDEFINITIONDATEDB:  make(map[uint]*ATTRIBUTEDEFINITIONDATEDB, 0),
		Map_ATTRIBUTEDEFINITIONDATEPtr_ATTRIBUTEDEFINITIONDATEDBID: make(map[*models.ATTRIBUTEDEFINITIONDATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION = BackRepoATTRIBUTEDEFINITIONENUMERATIONStruct{
		Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONENUMERATION, 0),
		Map_ATTRIBUTEDEFINITIONENUMERATIONDBID_ATTRIBUTEDEFINITIONENUMERATIONDB:  make(map[uint]*ATTRIBUTEDEFINITIONENUMERATIONDB, 0),
		Map_ATTRIBUTEDEFINITIONENUMERATIONPtr_ATTRIBUTEDEFINITIONENUMERATIONDBID: make(map[*models.ATTRIBUTEDEFINITIONENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER = BackRepoATTRIBUTEDEFINITIONINTEGERStruct{
		Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONINTEGER, 0),
		Map_ATTRIBUTEDEFINITIONINTEGERDBID_ATTRIBUTEDEFINITIONINTEGERDB:  make(map[uint]*ATTRIBUTEDEFINITIONINTEGERDB, 0),
		Map_ATTRIBUTEDEFINITIONINTEGERPtr_ATTRIBUTEDEFINITIONINTEGERDBID: make(map[*models.ATTRIBUTEDEFINITIONINTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL = BackRepoATTRIBUTEDEFINITIONREALStruct{
		Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONREAL, 0),
		Map_ATTRIBUTEDEFINITIONREALDBID_ATTRIBUTEDEFINITIONREALDB:  make(map[uint]*ATTRIBUTEDEFINITIONREALDB, 0),
		Map_ATTRIBUTEDEFINITIONREALPtr_ATTRIBUTEDEFINITIONREALDBID: make(map[*models.ATTRIBUTEDEFINITIONREAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING = BackRepoATTRIBUTEDEFINITIONSTRINGStruct{
		Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONSTRING, 0),
		Map_ATTRIBUTEDEFINITIONSTRINGDBID_ATTRIBUTEDEFINITIONSTRINGDB:  make(map[uint]*ATTRIBUTEDEFINITIONSTRINGDB, 0),
		Map_ATTRIBUTEDEFINITIONSTRINGPtr_ATTRIBUTEDEFINITIONSTRINGDBID: make(map[*models.ATTRIBUTEDEFINITIONSTRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML = BackRepoATTRIBUTEDEFINITIONXHTMLStruct{
		Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLPtr: make(map[uint]*models.ATTRIBUTEDEFINITIONXHTML, 0),
		Map_ATTRIBUTEDEFINITIONXHTMLDBID_ATTRIBUTEDEFINITIONXHTMLDB:  make(map[uint]*ATTRIBUTEDEFINITIONXHTMLDB, 0),
		Map_ATTRIBUTEDEFINITIONXHTMLPtr_ATTRIBUTEDEFINITIONXHTMLDBID: make(map[*models.ATTRIBUTEDEFINITIONXHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN = BackRepoATTRIBUTEVALUEBOOLEANStruct{
		Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANPtr: make(map[uint]*models.ATTRIBUTEVALUEBOOLEAN, 0),
		Map_ATTRIBUTEVALUEBOOLEANDBID_ATTRIBUTEVALUEBOOLEANDB:  make(map[uint]*ATTRIBUTEVALUEBOOLEANDB, 0),
		Map_ATTRIBUTEVALUEBOOLEANPtr_ATTRIBUTEVALUEBOOLEANDBID: make(map[*models.ATTRIBUTEVALUEBOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEDATE = BackRepoATTRIBUTEVALUEDATEStruct{
		Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEPtr: make(map[uint]*models.ATTRIBUTEVALUEDATE, 0),
		Map_ATTRIBUTEVALUEDATEDBID_ATTRIBUTEVALUEDATEDB:  make(map[uint]*ATTRIBUTEVALUEDATEDB, 0),
		Map_ATTRIBUTEVALUEDATEPtr_ATTRIBUTEVALUEDATEDBID: make(map[*models.ATTRIBUTEVALUEDATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION = BackRepoATTRIBUTEVALUEENUMERATIONStruct{
		Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONPtr: make(map[uint]*models.ATTRIBUTEVALUEENUMERATION, 0),
		Map_ATTRIBUTEVALUEENUMERATIONDBID_ATTRIBUTEVALUEENUMERATIONDB:  make(map[uint]*ATTRIBUTEVALUEENUMERATIONDB, 0),
		Map_ATTRIBUTEVALUEENUMERATIONPtr_ATTRIBUTEVALUEENUMERATIONDBID: make(map[*models.ATTRIBUTEVALUEENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEINTEGER = BackRepoATTRIBUTEVALUEINTEGERStruct{
		Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERPtr: make(map[uint]*models.ATTRIBUTEVALUEINTEGER, 0),
		Map_ATTRIBUTEVALUEINTEGERDBID_ATTRIBUTEVALUEINTEGERDB:  make(map[uint]*ATTRIBUTEVALUEINTEGERDB, 0),
		Map_ATTRIBUTEVALUEINTEGERPtr_ATTRIBUTEVALUEINTEGERDBID: make(map[*models.ATTRIBUTEVALUEINTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEREAL = BackRepoATTRIBUTEVALUEREALStruct{
		Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALPtr: make(map[uint]*models.ATTRIBUTEVALUEREAL, 0),
		Map_ATTRIBUTEVALUEREALDBID_ATTRIBUTEVALUEREALDB:  make(map[uint]*ATTRIBUTEVALUEREALDB, 0),
		Map_ATTRIBUTEVALUEREALPtr_ATTRIBUTEVALUEREALDBID: make(map[*models.ATTRIBUTEVALUEREAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUESTRING = BackRepoATTRIBUTEVALUESTRINGStruct{
		Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGPtr: make(map[uint]*models.ATTRIBUTEVALUESTRING, 0),
		Map_ATTRIBUTEVALUESTRINGDBID_ATTRIBUTEVALUESTRINGDB:  make(map[uint]*ATTRIBUTEVALUESTRINGDB, 0),
		Map_ATTRIBUTEVALUESTRINGPtr_ATTRIBUTEVALUESTRINGDBID: make(map[*models.ATTRIBUTEVALUESTRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTEVALUEXHTML = BackRepoATTRIBUTEVALUEXHTMLStruct{
		Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLPtr: make(map[uint]*models.ATTRIBUTEVALUEXHTML, 0),
		Map_ATTRIBUTEVALUEXHTMLDBID_ATTRIBUTEVALUEXHTMLDB:  make(map[uint]*ATTRIBUTEVALUEXHTMLDB, 0),
		Map_ATTRIBUTEVALUEXHTMLPtr_ATTRIBUTEVALUEXHTMLDBID: make(map[*models.ATTRIBUTEVALUEXHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCHILDREN = BackRepoCHILDRENStruct{
		Map_CHILDRENDBID_CHILDRENPtr: make(map[uint]*models.CHILDREN, 0),
		Map_CHILDRENDBID_CHILDRENDB:  make(map[uint]*CHILDRENDB, 0),
		Map_CHILDRENPtr_CHILDRENDBID: make(map[*models.CHILDREN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCORECONTENT = BackRepoCORECONTENTStruct{
		Map_CORECONTENTDBID_CORECONTENTPtr: make(map[uint]*models.CORECONTENT, 0),
		Map_CORECONTENTDBID_CORECONTENTDB:  make(map[uint]*CORECONTENTDB, 0),
		Map_CORECONTENTPtr_CORECONTENTDBID: make(map[*models.CORECONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN = BackRepoDATATYPEDEFINITIONBOOLEANStruct{
		Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANPtr: make(map[uint]*models.DATATYPEDEFINITIONBOOLEAN, 0),
		Map_DATATYPEDEFINITIONBOOLEANDBID_DATATYPEDEFINITIONBOOLEANDB:  make(map[uint]*DATATYPEDEFINITIONBOOLEANDB, 0),
		Map_DATATYPEDEFINITIONBOOLEANPtr_DATATYPEDEFINITIONBOOLEANDBID: make(map[*models.DATATYPEDEFINITIONBOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONDATE = BackRepoDATATYPEDEFINITIONDATEStruct{
		Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEPtr: make(map[uint]*models.DATATYPEDEFINITIONDATE, 0),
		Map_DATATYPEDEFINITIONDATEDBID_DATATYPEDEFINITIONDATEDB:  make(map[uint]*DATATYPEDEFINITIONDATEDB, 0),
		Map_DATATYPEDEFINITIONDATEPtr_DATATYPEDEFINITIONDATEDBID: make(map[*models.DATATYPEDEFINITIONDATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION = BackRepoDATATYPEDEFINITIONENUMERATIONStruct{
		Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONPtr: make(map[uint]*models.DATATYPEDEFINITIONENUMERATION, 0),
		Map_DATATYPEDEFINITIONENUMERATIONDBID_DATATYPEDEFINITIONENUMERATIONDB:  make(map[uint]*DATATYPEDEFINITIONENUMERATIONDB, 0),
		Map_DATATYPEDEFINITIONENUMERATIONPtr_DATATYPEDEFINITIONENUMERATIONDBID: make(map[*models.DATATYPEDEFINITIONENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER = BackRepoDATATYPEDEFINITIONINTEGERStruct{
		Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERPtr: make(map[uint]*models.DATATYPEDEFINITIONINTEGER, 0),
		Map_DATATYPEDEFINITIONINTEGERDBID_DATATYPEDEFINITIONINTEGERDB:  make(map[uint]*DATATYPEDEFINITIONINTEGERDB, 0),
		Map_DATATYPEDEFINITIONINTEGERPtr_DATATYPEDEFINITIONINTEGERDBID: make(map[*models.DATATYPEDEFINITIONINTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONREAL = BackRepoDATATYPEDEFINITIONREALStruct{
		Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALPtr: make(map[uint]*models.DATATYPEDEFINITIONREAL, 0),
		Map_DATATYPEDEFINITIONREALDBID_DATATYPEDEFINITIONREALDB:  make(map[uint]*DATATYPEDEFINITIONREALDB, 0),
		Map_DATATYPEDEFINITIONREALPtr_DATATYPEDEFINITIONREALDBID: make(map[*models.DATATYPEDEFINITIONREAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONSTRING = BackRepoDATATYPEDEFINITIONSTRINGStruct{
		Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGPtr: make(map[uint]*models.DATATYPEDEFINITIONSTRING, 0),
		Map_DATATYPEDEFINITIONSTRINGDBID_DATATYPEDEFINITIONSTRINGDB:  make(map[uint]*DATATYPEDEFINITIONSTRINGDB, 0),
		Map_DATATYPEDEFINITIONSTRINGPtr_DATATYPEDEFINITIONSTRINGDBID: make(map[*models.DATATYPEDEFINITIONSTRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPEDEFINITIONXHTML = BackRepoDATATYPEDEFINITIONXHTMLStruct{
		Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLPtr: make(map[uint]*models.DATATYPEDEFINITIONXHTML, 0),
		Map_DATATYPEDEFINITIONXHTMLDBID_DATATYPEDEFINITIONXHTMLDB:  make(map[uint]*DATATYPEDEFINITIONXHTMLDB, 0),
		Map_DATATYPEDEFINITIONXHTMLPtr_DATATYPEDEFINITIONXHTMLDBID: make(map[*models.DATATYPEDEFINITIONXHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPES = BackRepoDATATYPESStruct{
		Map_DATATYPESDBID_DATATYPESPtr: make(map[uint]*models.DATATYPES, 0),
		Map_DATATYPESDBID_DATATYPESDB:  make(map[uint]*DATATYPESDB, 0),
		Map_DATATYPESPtr_DATATYPESDBID: make(map[*models.DATATYPES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDEFAULTVALUE = BackRepoDEFAULTVALUEStruct{
		Map_DEFAULTVALUEDBID_DEFAULTVALUEPtr: make(map[uint]*models.DEFAULTVALUE, 0),
		Map_DEFAULTVALUEDBID_DEFAULTVALUEDB:  make(map[uint]*DEFAULTVALUEDB, 0),
		Map_DEFAULTVALUEPtr_DEFAULTVALUEDBID: make(map[*models.DEFAULTVALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDEFINITION = BackRepoDEFINITIONStruct{
		Map_DEFINITIONDBID_DEFINITIONPtr: make(map[uint]*models.DEFINITION, 0),
		Map_DEFINITIONDBID_DEFINITIONDB:  make(map[uint]*DEFINITIONDB, 0),
		Map_DEFINITIONPtr_DEFINITIONDBID: make(map[*models.DEFINITION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEDITABLEATTS = BackRepoEDITABLEATTSStruct{
		Map_EDITABLEATTSDBID_EDITABLEATTSPtr: make(map[uint]*models.EDITABLEATTS, 0),
		Map_EDITABLEATTSDBID_EDITABLEATTSDB:  make(map[uint]*EDITABLEATTSDB, 0),
		Map_EDITABLEATTSPtr_EDITABLEATTSDBID: make(map[*models.EDITABLEATTS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEMBEDDEDVALUE = BackRepoEMBEDDEDVALUEStruct{
		Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEPtr: make(map[uint]*models.EMBEDDEDVALUE, 0),
		Map_EMBEDDEDVALUEDBID_EMBEDDEDVALUEDB:  make(map[uint]*EMBEDDEDVALUEDB, 0),
		Map_EMBEDDEDVALUEPtr_EMBEDDEDVALUEDBID: make(map[*models.EMBEDDEDVALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoENUMVALUE = BackRepoENUMVALUEStruct{
		Map_ENUMVALUEDBID_ENUMVALUEPtr: make(map[uint]*models.ENUMVALUE, 0),
		Map_ENUMVALUEDBID_ENUMVALUEDB:  make(map[uint]*ENUMVALUEDB, 0),
		Map_ENUMVALUEPtr_ENUMVALUEDBID: make(map[*models.ENUMVALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOBJECT = BackRepoOBJECTStruct{
		Map_OBJECTDBID_OBJECTPtr: make(map[uint]*models.OBJECT, 0),
		Map_OBJECTDBID_OBJECTDB:  make(map[uint]*OBJECTDB, 0),
		Map_OBJECTPtr_OBJECTDBID: make(map[*models.OBJECT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPROPERTIES = BackRepoPROPERTIESStruct{
		Map_PROPERTIESDBID_PROPERTIESPtr: make(map[uint]*models.PROPERTIES, 0),
		Map_PROPERTIESDBID_PROPERTIESDB:  make(map[uint]*PROPERTIESDB, 0),
		Map_PROPERTIESPtr_PROPERTIESDBID: make(map[*models.PROPERTIES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATIONGROUP = BackRepoRELATIONGROUPStruct{
		Map_RELATIONGROUPDBID_RELATIONGROUPPtr: make(map[uint]*models.RELATIONGROUP, 0),
		Map_RELATIONGROUPDBID_RELATIONGROUPDB:  make(map[uint]*RELATIONGROUPDB, 0),
		Map_RELATIONGROUPPtr_RELATIONGROUPDBID: make(map[*models.RELATIONGROUP]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATIONGROUPTYPE = BackRepoRELATIONGROUPTYPEStruct{
		Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEPtr: make(map[uint]*models.RELATIONGROUPTYPE, 0),
		Map_RELATIONGROUPTYPEDBID_RELATIONGROUPTYPEDB:  make(map[uint]*RELATIONGROUPTYPEDB, 0),
		Map_RELATIONGROUPTYPEPtr_RELATIONGROUPTYPEDBID: make(map[*models.RELATIONGROUPTYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQIF = BackRepoREQIFStruct{
		Map_REQIFDBID_REQIFPtr: make(map[uint]*models.REQIF, 0),
		Map_REQIFDBID_REQIFDB:  make(map[uint]*REQIFDB, 0),
		Map_REQIFPtr_REQIFDBID: make(map[*models.REQIF]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQIFCONTENT = BackRepoREQIFCONTENTStruct{
		Map_REQIFCONTENTDBID_REQIFCONTENTPtr: make(map[uint]*models.REQIFCONTENT, 0),
		Map_REQIFCONTENTDBID_REQIFCONTENTDB:  make(map[uint]*REQIFCONTENTDB, 0),
		Map_REQIFCONTENTPtr_REQIFCONTENTDBID: make(map[*models.REQIFCONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQIFHEADER = BackRepoREQIFHEADERStruct{
		Map_REQIFHEADERDBID_REQIFHEADERPtr: make(map[uint]*models.REQIFHEADER, 0),
		Map_REQIFHEADERDBID_REQIFHEADERDB:  make(map[uint]*REQIFHEADERDB, 0),
		Map_REQIFHEADERPtr_REQIFHEADERDBID: make(map[*models.REQIFHEADER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQIFTOOLEXTENSION = BackRepoREQIFTOOLEXTENSIONStruct{
		Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONPtr: make(map[uint]*models.REQIFTOOLEXTENSION, 0),
		Map_REQIFTOOLEXTENSIONDBID_REQIFTOOLEXTENSIONDB:  make(map[uint]*REQIFTOOLEXTENSIONDB, 0),
		Map_REQIFTOOLEXTENSIONPtr_REQIFTOOLEXTENSIONDBID: make(map[*models.REQIFTOOLEXTENSION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQIFTYPE = BackRepoREQIFTYPEStruct{
		Map_REQIFTYPEDBID_REQIFTYPEPtr: make(map[uint]*models.REQIFTYPE, 0),
		Map_REQIFTYPEDBID_REQIFTYPEDB:  make(map[uint]*REQIFTYPEDB, 0),
		Map_REQIFTYPEPtr_REQIFTYPEDBID: make(map[*models.REQIFTYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSOURCE = BackRepoSOURCEStruct{
		Map_SOURCEDBID_SOURCEPtr: make(map[uint]*models.SOURCE, 0),
		Map_SOURCEDBID_SOURCEDB:  make(map[uint]*SOURCEDB, 0),
		Map_SOURCEPtr_SOURCEDBID: make(map[*models.SOURCE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSOURCESPECIFICATION = BackRepoSOURCESPECIFICATIONStruct{
		Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONPtr: make(map[uint]*models.SOURCESPECIFICATION, 0),
		Map_SOURCESPECIFICATIONDBID_SOURCESPECIFICATIONDB:  make(map[uint]*SOURCESPECIFICATIONDB, 0),
		Map_SOURCESPECIFICATIONPtr_SOURCESPECIFICATIONDBID: make(map[*models.SOURCESPECIFICATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECATTRIBUTES = BackRepoSPECATTRIBUTESStruct{
		Map_SPECATTRIBUTESDBID_SPECATTRIBUTESPtr: make(map[uint]*models.SPECATTRIBUTES, 0),
		Map_SPECATTRIBUTESDBID_SPECATTRIBUTESDB:  make(map[uint]*SPECATTRIBUTESDB, 0),
		Map_SPECATTRIBUTESPtr_SPECATTRIBUTESDBID: make(map[*models.SPECATTRIBUTES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECHIERARCHY = BackRepoSPECHIERARCHYStruct{
		Map_SPECHIERARCHYDBID_SPECHIERARCHYPtr: make(map[uint]*models.SPECHIERARCHY, 0),
		Map_SPECHIERARCHYDBID_SPECHIERARCHYDB:  make(map[uint]*SPECHIERARCHYDB, 0),
		Map_SPECHIERARCHYPtr_SPECHIERARCHYDBID: make(map[*models.SPECHIERARCHY]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATION = BackRepoSPECIFICATIONStruct{
		Map_SPECIFICATIONDBID_SPECIFICATIONPtr: make(map[uint]*models.SPECIFICATION, 0),
		Map_SPECIFICATIONDBID_SPECIFICATIONDB:  make(map[uint]*SPECIFICATIONDB, 0),
		Map_SPECIFICATIONPtr_SPECIFICATIONDBID: make(map[*models.SPECIFICATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATIONS = BackRepoSPECIFICATIONSStruct{
		Map_SPECIFICATIONSDBID_SPECIFICATIONSPtr: make(map[uint]*models.SPECIFICATIONS, 0),
		Map_SPECIFICATIONSDBID_SPECIFICATIONSDB:  make(map[uint]*SPECIFICATIONSDB, 0),
		Map_SPECIFICATIONSPtr_SPECIFICATIONSDBID: make(map[*models.SPECIFICATIONS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATIONTYPE = BackRepoSPECIFICATIONTYPEStruct{
		Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEPtr: make(map[uint]*models.SPECIFICATIONTYPE, 0),
		Map_SPECIFICATIONTYPEDBID_SPECIFICATIONTYPEDB:  make(map[uint]*SPECIFICATIONTYPEDB, 0),
		Map_SPECIFICATIONTYPEPtr_SPECIFICATIONTYPEDBID: make(map[*models.SPECIFICATIONTYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFIEDVALUES = BackRepoSPECIFIEDVALUESStruct{
		Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESPtr: make(map[uint]*models.SPECIFIEDVALUES, 0),
		Map_SPECIFIEDVALUESDBID_SPECIFIEDVALUESDB:  make(map[uint]*SPECIFIEDVALUESDB, 0),
		Map_SPECIFIEDVALUESPtr_SPECIFIEDVALUESDBID: make(map[*models.SPECIFIEDVALUES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECOBJECT = BackRepoSPECOBJECTStruct{
		Map_SPECOBJECTDBID_SPECOBJECTPtr: make(map[uint]*models.SPECOBJECT, 0),
		Map_SPECOBJECTDBID_SPECOBJECTDB:  make(map[uint]*SPECOBJECTDB, 0),
		Map_SPECOBJECTPtr_SPECOBJECTDBID: make(map[*models.SPECOBJECT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECOBJECTS = BackRepoSPECOBJECTSStruct{
		Map_SPECOBJECTSDBID_SPECOBJECTSPtr: make(map[uint]*models.SPECOBJECTS, 0),
		Map_SPECOBJECTSDBID_SPECOBJECTSDB:  make(map[uint]*SPECOBJECTSDB, 0),
		Map_SPECOBJECTSPtr_SPECOBJECTSDBID: make(map[*models.SPECOBJECTS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECOBJECTTYPE = BackRepoSPECOBJECTTYPEStruct{
		Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEPtr: make(map[uint]*models.SPECOBJECTTYPE, 0),
		Map_SPECOBJECTTYPEDBID_SPECOBJECTTYPEDB:  make(map[uint]*SPECOBJECTTYPEDB, 0),
		Map_SPECOBJECTTYPEPtr_SPECOBJECTTYPEDBID: make(map[*models.SPECOBJECTTYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECRELATION = BackRepoSPECRELATIONStruct{
		Map_SPECRELATIONDBID_SPECRELATIONPtr: make(map[uint]*models.SPECRELATION, 0),
		Map_SPECRELATIONDBID_SPECRELATIONDB:  make(map[uint]*SPECRELATIONDB, 0),
		Map_SPECRELATIONPtr_SPECRELATIONDBID: make(map[*models.SPECRELATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECRELATIONGROUPS = BackRepoSPECRELATIONGROUPSStruct{
		Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSPtr: make(map[uint]*models.SPECRELATIONGROUPS, 0),
		Map_SPECRELATIONGROUPSDBID_SPECRELATIONGROUPSDB:  make(map[uint]*SPECRELATIONGROUPSDB, 0),
		Map_SPECRELATIONGROUPSPtr_SPECRELATIONGROUPSDBID: make(map[*models.SPECRELATIONGROUPS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECRELATIONS = BackRepoSPECRELATIONSStruct{
		Map_SPECRELATIONSDBID_SPECRELATIONSPtr: make(map[uint]*models.SPECRELATIONS, 0),
		Map_SPECRELATIONSDBID_SPECRELATIONSDB:  make(map[uint]*SPECRELATIONSDB, 0),
		Map_SPECRELATIONSPtr_SPECRELATIONSDBID: make(map[*models.SPECRELATIONS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECRELATIONTYPE = BackRepoSPECRELATIONTYPEStruct{
		Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEPtr: make(map[uint]*models.SPECRELATIONTYPE, 0),
		Map_SPECRELATIONTYPEDBID_SPECRELATIONTYPEDB:  make(map[uint]*SPECRELATIONTYPEDB, 0),
		Map_SPECRELATIONTYPEPtr_SPECRELATIONTYPEDBID: make(map[*models.SPECRELATIONTYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECTYPES = BackRepoSPECTYPESStruct{
		Map_SPECTYPESDBID_SPECTYPESPtr: make(map[uint]*models.SPECTYPES, 0),
		Map_SPECTYPESDBID_SPECTYPESDB:  make(map[uint]*SPECTYPESDB, 0),
		Map_SPECTYPESPtr_SPECTYPESDBID: make(map[*models.SPECTYPES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTARGET = BackRepoTARGETStruct{
		Map_TARGETDBID_TARGETPtr: make(map[uint]*models.TARGET, 0),
		Map_TARGETDBID_TARGETDB:  make(map[uint]*TARGETDB, 0),
		Map_TARGETPtr_TARGETDBID: make(map[*models.TARGET]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTARGETSPECIFICATION = BackRepoTARGETSPECIFICATIONStruct{
		Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONPtr: make(map[uint]*models.TARGETSPECIFICATION, 0),
		Map_TARGETSPECIFICATIONDBID_TARGETSPECIFICATIONDB:  make(map[uint]*TARGETSPECIFICATIONDB, 0),
		Map_TARGETSPECIFICATIONPtr_TARGETSPECIFICATIONDBID: make(map[*models.TARGETSPECIFICATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTHEHEADER = BackRepoTHEHEADERStruct{
		Map_THEHEADERDBID_THEHEADERPtr: make(map[uint]*models.THEHEADER, 0),
		Map_THEHEADERDBID_THEHEADERDB:  make(map[uint]*THEHEADERDB, 0),
		Map_THEHEADERPtr_THEHEADERDBID: make(map[*models.THEHEADER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTOOLEXTENSIONS = BackRepoTOOLEXTENSIONSStruct{
		Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSPtr: make(map[uint]*models.TOOLEXTENSIONS, 0),
		Map_TOOLEXTENSIONSDBID_TOOLEXTENSIONSDB:  make(map[uint]*TOOLEXTENSIONSDB, 0),
		Map_TOOLEXTENSIONSPtr_TOOLEXTENSIONSDBID: make(map[*models.TOOLEXTENSIONS]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVALUES = BackRepoVALUESStruct{
		Map_VALUESDBID_VALUESPtr: make(map[uint]*models.VALUES, 0),
		Map_VALUESDBID_VALUESDB:  make(map[uint]*VALUESDB, 0),
		Map_VALUESPtr_VALUESDBID: make(map[*models.VALUES]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXHTMLCONTENT = BackRepoXHTMLCONTENTStruct{
		Map_XHTMLCONTENTDBID_XHTMLCONTENTPtr: make(map[uint]*models.XHTMLCONTENT, 0),
		Map_XHTMLCONTENTDBID_XHTMLCONTENTDB:  make(map[uint]*XHTMLCONTENTDB, 0),
		Map_XHTMLCONTENTPtr_XHTMLCONTENTDBID: make(map[*models.XHTMLCONTENT]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVEID.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEDATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEREAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUESTRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.CommitPhaseOne(stage)
	backRepo.BackRepoCHILDREN.CommitPhaseOne(stage)
	backRepo.BackRepoCORECONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPES.CommitPhaseOne(stage)
	backRepo.BackRepoDEFAULTVALUE.CommitPhaseOne(stage)
	backRepo.BackRepoDEFINITION.CommitPhaseOne(stage)
	backRepo.BackRepoEDITABLEATTS.CommitPhaseOne(stage)
	backRepo.BackRepoEMBEDDEDVALUE.CommitPhaseOne(stage)
	backRepo.BackRepoENUMVALUE.CommitPhaseOne(stage)
	backRepo.BackRepoOBJECT.CommitPhaseOne(stage)
	backRepo.BackRepoPROPERTIES.CommitPhaseOne(stage)
	backRepo.BackRepoRELATIONGROUP.CommitPhaseOne(stage)
	backRepo.BackRepoRELATIONGROUPTYPE.CommitPhaseOne(stage)
	backRepo.BackRepoREQIF.CommitPhaseOne(stage)
	backRepo.BackRepoREQIFCONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoREQIFHEADER.CommitPhaseOne(stage)
	backRepo.BackRepoREQIFTOOLEXTENSION.CommitPhaseOne(stage)
	backRepo.BackRepoREQIFTYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSOURCE.CommitPhaseOne(stage)
	backRepo.BackRepoSOURCESPECIFICATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECATTRIBUTES.CommitPhaseOne(stage)
	backRepo.BackRepoSPECHIERARCHY.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATIONS.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATIONTYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFIEDVALUES.CommitPhaseOne(stage)
	backRepo.BackRepoSPECOBJECT.CommitPhaseOne(stage)
	backRepo.BackRepoSPECOBJECTS.CommitPhaseOne(stage)
	backRepo.BackRepoSPECOBJECTTYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPECRELATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECRELATIONGROUPS.CommitPhaseOne(stage)
	backRepo.BackRepoSPECRELATIONS.CommitPhaseOne(stage)
	backRepo.BackRepoSPECRELATIONTYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPECTYPES.CommitPhaseOne(stage)
	backRepo.BackRepoTARGET.CommitPhaseOne(stage)
	backRepo.BackRepoTARGETSPECIFICATION.CommitPhaseOne(stage)
	backRepo.BackRepoTHEHEADER.CommitPhaseOne(stage)
	backRepo.BackRepoTOOLEXTENSIONS.CommitPhaseOne(stage)
	backRepo.BackRepoVALUES.CommitPhaseOne(stage)
	backRepo.BackRepoXHTMLCONTENT.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVEID.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEDATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEREAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUESTRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCHILDREN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCORECONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDEFAULTVALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDEFINITION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEDITABLEATTS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDEDVALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoENUMVALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOBJECT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPROPERTIES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATIONGROUP.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATIONGROUPTYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQIF.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQIFCONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQIFHEADER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQIFTOOLEXTENSION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQIFTYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSOURCE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSOURCESPECIFICATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECATTRIBUTES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECHIERARCHY.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATIONS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATIONTYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFIEDVALUES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECTS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECTTYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONGROUPS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONTYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECTYPES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTARGET.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTARGETSPECIFICATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTHEHEADER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTOOLEXTENSIONS.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVALUES.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXHTMLCONTENT.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVEID.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEDATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEREAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUESTRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTEVALUEXHTML.CheckoutPhaseOne()
	backRepo.BackRepoCHILDREN.CheckoutPhaseOne()
	backRepo.BackRepoCORECONTENT.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONREAL.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPES.CheckoutPhaseOne()
	backRepo.BackRepoDEFAULTVALUE.CheckoutPhaseOne()
	backRepo.BackRepoDEFINITION.CheckoutPhaseOne()
	backRepo.BackRepoEDITABLEATTS.CheckoutPhaseOne()
	backRepo.BackRepoEMBEDDEDVALUE.CheckoutPhaseOne()
	backRepo.BackRepoENUMVALUE.CheckoutPhaseOne()
	backRepo.BackRepoOBJECT.CheckoutPhaseOne()
	backRepo.BackRepoPROPERTIES.CheckoutPhaseOne()
	backRepo.BackRepoRELATIONGROUP.CheckoutPhaseOne()
	backRepo.BackRepoRELATIONGROUPTYPE.CheckoutPhaseOne()
	backRepo.BackRepoREQIF.CheckoutPhaseOne()
	backRepo.BackRepoREQIFCONTENT.CheckoutPhaseOne()
	backRepo.BackRepoREQIFHEADER.CheckoutPhaseOne()
	backRepo.BackRepoREQIFTOOLEXTENSION.CheckoutPhaseOne()
	backRepo.BackRepoREQIFTYPE.CheckoutPhaseOne()
	backRepo.BackRepoSOURCE.CheckoutPhaseOne()
	backRepo.BackRepoSOURCESPECIFICATION.CheckoutPhaseOne()
	backRepo.BackRepoSPECATTRIBUTES.CheckoutPhaseOne()
	backRepo.BackRepoSPECHIERARCHY.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATIONS.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATIONTYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFIEDVALUES.CheckoutPhaseOne()
	backRepo.BackRepoSPECOBJECT.CheckoutPhaseOne()
	backRepo.BackRepoSPECOBJECTS.CheckoutPhaseOne()
	backRepo.BackRepoSPECOBJECTTYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPECRELATION.CheckoutPhaseOne()
	backRepo.BackRepoSPECRELATIONGROUPS.CheckoutPhaseOne()
	backRepo.BackRepoSPECRELATIONS.CheckoutPhaseOne()
	backRepo.BackRepoSPECRELATIONTYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPECTYPES.CheckoutPhaseOne()
	backRepo.BackRepoTARGET.CheckoutPhaseOne()
	backRepo.BackRepoTARGETSPECIFICATION.CheckoutPhaseOne()
	backRepo.BackRepoTHEHEADER.CheckoutPhaseOne()
	backRepo.BackRepoTOOLEXTENSIONS.CheckoutPhaseOne()
	backRepo.BackRepoVALUES.CheckoutPhaseOne()
	backRepo.BackRepoXHTMLCONTENT.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVEID.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEDATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEREAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUESTRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCHILDREN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCORECONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDEFAULTVALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDEFINITION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEDITABLEATTS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDEDVALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoENUMVALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOBJECT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPROPERTIES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATIONGROUP.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATIONGROUPTYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQIF.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQIFCONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQIFHEADER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQIFTOOLEXTENSION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQIFTYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSOURCE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSOURCESPECIFICATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECATTRIBUTES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECHIERARCHY.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATIONS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATIONTYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFIEDVALUES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECTS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECOBJECTTYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONGROUPS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECRELATIONTYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECTYPES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTARGET.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTARGETSPECIFICATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTHEHEADER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTOOLEXTENSIONS.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVALUES.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXHTMLCONTENT.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVEID.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEDATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEREAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUESTRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.Backup(dirPath)
	backRepo.BackRepoCHILDREN.Backup(dirPath)
	backRepo.BackRepoCORECONTENT.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.Backup(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.Backup(dirPath)
	backRepo.BackRepoDATATYPES.Backup(dirPath)
	backRepo.BackRepoDEFAULTVALUE.Backup(dirPath)
	backRepo.BackRepoDEFINITION.Backup(dirPath)
	backRepo.BackRepoEDITABLEATTS.Backup(dirPath)
	backRepo.BackRepoEMBEDDEDVALUE.Backup(dirPath)
	backRepo.BackRepoENUMVALUE.Backup(dirPath)
	backRepo.BackRepoOBJECT.Backup(dirPath)
	backRepo.BackRepoPROPERTIES.Backup(dirPath)
	backRepo.BackRepoRELATIONGROUP.Backup(dirPath)
	backRepo.BackRepoRELATIONGROUPTYPE.Backup(dirPath)
	backRepo.BackRepoREQIF.Backup(dirPath)
	backRepo.BackRepoREQIFCONTENT.Backup(dirPath)
	backRepo.BackRepoREQIFHEADER.Backup(dirPath)
	backRepo.BackRepoREQIFTOOLEXTENSION.Backup(dirPath)
	backRepo.BackRepoREQIFTYPE.Backup(dirPath)
	backRepo.BackRepoSOURCE.Backup(dirPath)
	backRepo.BackRepoSOURCESPECIFICATION.Backup(dirPath)
	backRepo.BackRepoSPECATTRIBUTES.Backup(dirPath)
	backRepo.BackRepoSPECHIERARCHY.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION.Backup(dirPath)
	backRepo.BackRepoSPECIFICATIONS.Backup(dirPath)
	backRepo.BackRepoSPECIFICATIONTYPE.Backup(dirPath)
	backRepo.BackRepoSPECIFIEDVALUES.Backup(dirPath)
	backRepo.BackRepoSPECOBJECT.Backup(dirPath)
	backRepo.BackRepoSPECOBJECTS.Backup(dirPath)
	backRepo.BackRepoSPECOBJECTTYPE.Backup(dirPath)
	backRepo.BackRepoSPECRELATION.Backup(dirPath)
	backRepo.BackRepoSPECRELATIONGROUPS.Backup(dirPath)
	backRepo.BackRepoSPECRELATIONS.Backup(dirPath)
	backRepo.BackRepoSPECRELATIONTYPE.Backup(dirPath)
	backRepo.BackRepoSPECTYPES.Backup(dirPath)
	backRepo.BackRepoTARGET.Backup(dirPath)
	backRepo.BackRepoTARGETSPECIFICATION.Backup(dirPath)
	backRepo.BackRepoTHEHEADER.Backup(dirPath)
	backRepo.BackRepoTOOLEXTENSIONS.Backup(dirPath)
	backRepo.BackRepoVALUES.Backup(dirPath)
	backRepo.BackRepoXHTMLCONTENT.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVEID.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEDATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEREAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUESTRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.BackupXL(file)
	backRepo.BackRepoCHILDREN.BackupXL(file)
	backRepo.BackRepoCORECONTENT.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.BackupXL(file)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.BackupXL(file)
	backRepo.BackRepoDATATYPES.BackupXL(file)
	backRepo.BackRepoDEFAULTVALUE.BackupXL(file)
	backRepo.BackRepoDEFINITION.BackupXL(file)
	backRepo.BackRepoEDITABLEATTS.BackupXL(file)
	backRepo.BackRepoEMBEDDEDVALUE.BackupXL(file)
	backRepo.BackRepoENUMVALUE.BackupXL(file)
	backRepo.BackRepoOBJECT.BackupXL(file)
	backRepo.BackRepoPROPERTIES.BackupXL(file)
	backRepo.BackRepoRELATIONGROUP.BackupXL(file)
	backRepo.BackRepoRELATIONGROUPTYPE.BackupXL(file)
	backRepo.BackRepoREQIF.BackupXL(file)
	backRepo.BackRepoREQIFCONTENT.BackupXL(file)
	backRepo.BackRepoREQIFHEADER.BackupXL(file)
	backRepo.BackRepoREQIFTOOLEXTENSION.BackupXL(file)
	backRepo.BackRepoREQIFTYPE.BackupXL(file)
	backRepo.BackRepoSOURCE.BackupXL(file)
	backRepo.BackRepoSOURCESPECIFICATION.BackupXL(file)
	backRepo.BackRepoSPECATTRIBUTES.BackupXL(file)
	backRepo.BackRepoSPECHIERARCHY.BackupXL(file)
	backRepo.BackRepoSPECIFICATION.BackupXL(file)
	backRepo.BackRepoSPECIFICATIONS.BackupXL(file)
	backRepo.BackRepoSPECIFICATIONTYPE.BackupXL(file)
	backRepo.BackRepoSPECIFIEDVALUES.BackupXL(file)
	backRepo.BackRepoSPECOBJECT.BackupXL(file)
	backRepo.BackRepoSPECOBJECTS.BackupXL(file)
	backRepo.BackRepoSPECOBJECTTYPE.BackupXL(file)
	backRepo.BackRepoSPECRELATION.BackupXL(file)
	backRepo.BackRepoSPECRELATIONGROUPS.BackupXL(file)
	backRepo.BackRepoSPECRELATIONS.BackupXL(file)
	backRepo.BackRepoSPECRELATIONTYPE.BackupXL(file)
	backRepo.BackRepoSPECTYPES.BackupXL(file)
	backRepo.BackRepoTARGET.BackupXL(file)
	backRepo.BackRepoTARGETSPECIFICATION.BackupXL(file)
	backRepo.BackRepoTHEHEADER.BackupXL(file)
	backRepo.BackRepoTOOLEXTENSIONS.BackupXL(file)
	backRepo.BackRepoVALUES.BackupXL(file)
	backRepo.BackRepoXHTMLCONTENT.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVEID.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEDATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEREAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUESTRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoCHILDREN.RestorePhaseOne(dirPath)
	backRepo.BackRepoCORECONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPES.RestorePhaseOne(dirPath)
	backRepo.BackRepoDEFAULTVALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoDEFINITION.RestorePhaseOne(dirPath)
	backRepo.BackRepoEDITABLEATTS.RestorePhaseOne(dirPath)
	backRepo.BackRepoEMBEDDEDVALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoENUMVALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoOBJECT.RestorePhaseOne(dirPath)
	backRepo.BackRepoPROPERTIES.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATIONGROUP.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATIONGROUPTYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQIF.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQIFCONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQIFHEADER.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQIFTOOLEXTENSION.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQIFTYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSOURCE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSOURCESPECIFICATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECATTRIBUTES.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECHIERARCHY.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATIONS.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATIONTYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFIEDVALUES.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECOBJECT.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECOBJECTS.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECOBJECTTYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECRELATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECRELATIONGROUPS.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECRELATIONS.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECRELATIONTYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECTYPES.RestorePhaseOne(dirPath)
	backRepo.BackRepoTARGET.RestorePhaseOne(dirPath)
	backRepo.BackRepoTARGETSPECIFICATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoTHEHEADER.RestorePhaseOne(dirPath)
	backRepo.BackRepoTOOLEXTENSIONS.RestorePhaseOne(dirPath)
	backRepo.BackRepoVALUES.RestorePhaseOne(dirPath)
	backRepo.BackRepoXHTMLCONTENT.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVEID.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEDATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEREAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUESTRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTEVALUEXHTML.RestorePhaseTwo()
	backRepo.BackRepoCHILDREN.RestorePhaseTwo()
	backRepo.BackRepoCORECONTENT.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONDATE.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONREAL.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.RestorePhaseTwo()
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.RestorePhaseTwo()
	backRepo.BackRepoDATATYPES.RestorePhaseTwo()
	backRepo.BackRepoDEFAULTVALUE.RestorePhaseTwo()
	backRepo.BackRepoDEFINITION.RestorePhaseTwo()
	backRepo.BackRepoEDITABLEATTS.RestorePhaseTwo()
	backRepo.BackRepoEMBEDDEDVALUE.RestorePhaseTwo()
	backRepo.BackRepoENUMVALUE.RestorePhaseTwo()
	backRepo.BackRepoOBJECT.RestorePhaseTwo()
	backRepo.BackRepoPROPERTIES.RestorePhaseTwo()
	backRepo.BackRepoRELATIONGROUP.RestorePhaseTwo()
	backRepo.BackRepoRELATIONGROUPTYPE.RestorePhaseTwo()
	backRepo.BackRepoREQIF.RestorePhaseTwo()
	backRepo.BackRepoREQIFCONTENT.RestorePhaseTwo()
	backRepo.BackRepoREQIFHEADER.RestorePhaseTwo()
	backRepo.BackRepoREQIFTOOLEXTENSION.RestorePhaseTwo()
	backRepo.BackRepoREQIFTYPE.RestorePhaseTwo()
	backRepo.BackRepoSOURCE.RestorePhaseTwo()
	backRepo.BackRepoSOURCESPECIFICATION.RestorePhaseTwo()
	backRepo.BackRepoSPECATTRIBUTES.RestorePhaseTwo()
	backRepo.BackRepoSPECHIERARCHY.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATIONS.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATIONTYPE.RestorePhaseTwo()
	backRepo.BackRepoSPECIFIEDVALUES.RestorePhaseTwo()
	backRepo.BackRepoSPECOBJECT.RestorePhaseTwo()
	backRepo.BackRepoSPECOBJECTS.RestorePhaseTwo()
	backRepo.BackRepoSPECOBJECTTYPE.RestorePhaseTwo()
	backRepo.BackRepoSPECRELATION.RestorePhaseTwo()
	backRepo.BackRepoSPECRELATIONGROUPS.RestorePhaseTwo()
	backRepo.BackRepoSPECRELATIONS.RestorePhaseTwo()
	backRepo.BackRepoSPECRELATIONTYPE.RestorePhaseTwo()
	backRepo.BackRepoSPECTYPES.RestorePhaseTwo()
	backRepo.BackRepoTARGET.RestorePhaseTwo()
	backRepo.BackRepoTARGETSPECIFICATION.RestorePhaseTwo()
	backRepo.BackRepoTHEHEADER.RestorePhaseTwo()
	backRepo.BackRepoTOOLEXTENSIONS.RestorePhaseTwo()
	backRepo.BackRepoVALUES.RestorePhaseTwo()
	backRepo.BackRepoXHTMLCONTENT.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVEID.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONBOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONDATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONINTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONREAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONSTRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEDEFINITIONXHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEBOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEDATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEINTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEREAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUESTRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTEVALUEXHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoCHILDREN.RestoreXLPhaseOne(file)
	backRepo.BackRepoCORECONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONBOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONDATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONINTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONREAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONSTRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPEDEFINITIONXHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPES.RestoreXLPhaseOne(file)
	backRepo.BackRepoDEFAULTVALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoDEFINITION.RestoreXLPhaseOne(file)
	backRepo.BackRepoEDITABLEATTS.RestoreXLPhaseOne(file)
	backRepo.BackRepoEMBEDDEDVALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoENUMVALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoOBJECT.RestoreXLPhaseOne(file)
	backRepo.BackRepoPROPERTIES.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATIONGROUP.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATIONGROUPTYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQIF.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQIFCONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQIFHEADER.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQIFTOOLEXTENSION.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQIFTYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSOURCE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSOURCESPECIFICATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECATTRIBUTES.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECHIERARCHY.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATIONS.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATIONTYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFIEDVALUES.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECOBJECT.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECOBJECTS.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECOBJECTTYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECRELATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECRELATIONGROUPS.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECRELATIONS.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECRELATIONTYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECTYPES.RestoreXLPhaseOne(file)
	backRepo.BackRepoTARGET.RestoreXLPhaseOne(file)
	backRepo.BackRepoTARGETSPECIFICATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoTHEHEADER.RestoreXLPhaseOne(file)
	backRepo.BackRepoTOOLEXTENSIONS.RestoreXLPhaseOne(file)
	backRepo.BackRepoVALUES.RestoreXLPhaseOne(file)
	backRepo.BackRepoXHTMLCONTENT.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
