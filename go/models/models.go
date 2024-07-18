package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type REQIF struct {
	Name string

	REQ_IF_HEADER  *REQ_IF_HEADER
	REQ_IF_CONTENT *REQ_IF_CONTENT
}

func (reqif *REQIF) Walk(_reqif *schema.REQ_IF, stage *StageStruct) {
	log.Println("reqif:Walk")

	reqif.REQ_IF_HEADER = new(REQ_IF_HEADER).Stage(stage)

	reqif.REQ_IF_HEADER.Name = time.Now().Format(time.DateTime)
	reqif.REQ_IF_HEADER.Walk(_reqif.THE_HEADER.REQ_IF_HEADER, stage)

	reqif.REQ_IF_CONTENT = new(REQ_IF_CONTENT).Stage(stage)
	reqif.REQ_IF_CONTENT.Walk(_reqif.CORE_CONTENT.REQ_IF_CONTENT, stage)
}
