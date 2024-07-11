package models

import (
	"github.com/thomaspeugeot/gongreqif/go/schema"
)

var stage *StageStruct

type REQIF struct {
	Name string

	d *schema.REQIF

	REQIFHEADER *REQIFHEADER
}

func (reqif *REQIF) SetD(_reqif *schema.REQIF) {
	reqif.d = _reqif
}

func (reqif *REQIF) SetStage(_stage *StageStruct) {
	stage = _stage
}

func (reqif *REQIF) Walk() {
	reqif.Name = reqif.d.XMLName.Local

	reqif.REQIFHEADER = new(REQIFHEADER).Stage(stage)
	reqif.REQIFHEADER.d = reqif.d.HEADER.REQIFHEADER
	reqif.REQIFHEADER.Walk()
}

type REQIFHEADER struct {
	d *schema.REQIFHEADER

	Name           string
	IDENTIFIERAttr string
	COMMENT        string
	CREATIONTIME   string
	REPOSITORYID   string
	REQIFTOOLID    string
	REQIFVERSION   string
	SOURCETOOLID   string
	TITLE          string
}

func (_REQIFHEADER *REQIFHEADER) Walk() {

	_REQIFHEADER.Name = _REQIFHEADER.d.XMLName.Local
	_REQIFHEADER.COMMENT = _REQIFHEADER.d.COMMENT
	_REQIFHEADER.CREATIONTIME = _REQIFHEADER.d.CREATIONTIME
	_REQIFHEADER.REPOSITORYID = _REQIFHEADER.d.REPOSITORYID
	_REQIFHEADER.REQIFTOOLID = _REQIFHEADER.d.REQIFTOOLID
	_REQIFHEADER.REQIFVERSION = _REQIFHEADER.d.REQIFVERSION
	_REQIFHEADER.SOURCETOOLID = _REQIFHEADER.d.SOURCETOOLID
	_REQIFHEADER.TITLE = _REQIFHEADER.d.TITLE
}
