package models

import (
	"log"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type SPEC_OBJECT_TYPE struct {
	Name string

	StandardAttributes
}

func (reqif *SPEC_OBJECT_TYPE) Walk(_reqif *schema.SPEC_OBJECT_TYPE, stage *StageStruct) {
	log.Println("SPEC_OBJECT_TYPE", "Walk")

	reqif.DESC = _reqif.DESC
	if _reqif.IDENTIFIER != nil {
		reqif.IDENTIFIER = string(*_reqif.IDENTIFIER)
	}
	reqif.LAST_CHANGE = _reqif.LAST_CHANGE.ToGoTime()
	reqif.LONG_NAME = _reqif.LONG_NAME
}
