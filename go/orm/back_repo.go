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
	BackRepoREQIF BackRepoREQIFStruct

	BackRepoREQ_IF_CONTENT BackRepoREQ_IF_CONTENTStruct

	BackRepoREQ_IF_HEADER BackRepoREQ_IF_HEADERStruct

	BackRepoSPECIFICATION BackRepoSPECIFICATIONStruct

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
		&REQIFDB{},
		&REQ_IF_CONTENTDB{},
		&REQ_IF_HEADERDB{},
		&SPECIFICATIONDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoREQIF = BackRepoREQIFStruct{
		Map_REQIFDBID_REQIFPtr: make(map[uint]*models.REQIF, 0),
		Map_REQIFDBID_REQIFDB:  make(map[uint]*REQIFDB, 0),
		Map_REQIFPtr_REQIFDBID: make(map[*models.REQIF]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_CONTENT = BackRepoREQ_IF_CONTENTStruct{
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr: make(map[uint]*models.REQ_IF_CONTENT, 0),
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB:  make(map[uint]*REQ_IF_CONTENTDB, 0),
		Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID: make(map[*models.REQ_IF_CONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_HEADER = BackRepoREQ_IF_HEADERStruct{
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr: make(map[uint]*models.REQ_IF_HEADER, 0),
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB:  make(map[uint]*REQ_IF_HEADERDB, 0),
		Map_REQ_IF_HEADERPtr_REQ_IF_HEADERDBID: make(map[*models.REQ_IF_HEADER]uint, 0),

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
	backRepo.BackRepoREQIF.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoREQIF.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoREQIF.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoREQIF.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoREQIF.Backup(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.Backup(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoREQIF.BackupXL(file)
	backRepo.BackRepoREQ_IF_CONTENT.BackupXL(file)
	backRepo.BackRepoREQ_IF_HEADER.BackupXL(file)
	backRepo.BackRepoSPECIFICATION.BackupXL(file)

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
	backRepo.BackRepoREQIF.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoREQIF.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION.RestorePhaseTwo()

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
	backRepo.BackRepoREQIF.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_CONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_HEADER.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION.RestoreXLPhaseOne(file)

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
