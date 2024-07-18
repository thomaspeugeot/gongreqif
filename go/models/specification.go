package models

import (
	"log"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPECIFICATION struct {
	Name string

	CHILDREN *SPEC_HIERARCHY

	StandardAttributes
}

func (reqif *SPECIFICATION) Walk(_reqif *schema.SPECIFICATION, stage *StageStruct) {
	log.Println("SPECIFICATION", "Walk")

	reqif.DESC = _reqif.DESC
	if _reqif.IDENTIFIER != nil {
		reqif.IDENTIFIER = string(*_reqif.IDENTIFIER)
	}
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME

	for _, _sh := range _reqif.CHILDREN.SPEC_HIERARCHY {
		sh := new(SPEC_HIERARCHY).Stage(stage)
		sh.Walk(_sh, stage)
		reqif.CHILDREN = sh
	}
}
