package models

import (
	"github.com/thomaspeugeot/gongreqif/go/schema"
)

var stage *StageStruct

type REQIF struct {
	Name string

	d *schema.REQIF

	HEADER  *HEADER
	CONTENT *CONTENT
}

func (reqif *REQIF) SetD(_reqif *schema.REQIF) {
	reqif.d = _reqif
}

func (reqif *REQIF) SetStage(_stage *StageStruct) {
	stage = _stage
}

func (reqif *REQIF) Walk() {
	reqif.Name = reqif.d.XMLName.Local

	reqif.HEADER = new(HEADER).Stage(stage)
	reqif.HEADER.d = reqif.d.HEADER.REQIFHEADER
	reqif.HEADER.Walk()

	reqif.CONTENT = new(CONTENT).Stage(stage)
	reqif.CONTENT.d = reqif.d.CORECONTENT.REQIFCONTENT
	reqif.CONTENT.Walk()
}

type HEADER struct {
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

func (header *HEADER) Walk() {

	header.Name = header.d.XMLName.Local
	header.IDENTIFIERAttr = header.d.IDENTIFIERAttr
	header.COMMENT = header.d.COMMENT
	header.CREATIONTIME = header.d.CREATIONTIME
	header.REPOSITORYID = header.d.REPOSITORYID
	header.REQIFTOOLID = header.d.REQIFTOOLID
	header.REQIFVERSION = header.d.REQIFVERSION
	header.SOURCETOOLID = header.d.SOURCETOOLID
	header.TITLE = header.d.TITLE
}

type CONTENT struct {
	Name string
	d    *schema.REQIFCONTENT
}

func (content *CONTENT) Walk() {

}
