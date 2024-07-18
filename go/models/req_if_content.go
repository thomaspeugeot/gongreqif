package models

import (
	"log"
	"time"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

type REQ_IF_CONTENT struct {
	Name string

	SPEC_OBJECT_TYPES   []*SPEC_OBJECT_TYPE
	SPECIFICATION_TYPES []*SPECIFICATION_TYPE
	SPECIFICATIONS      []*SPECIFICATION
}

func (reqif *REQ_IF_CONTENT) Walk(_reqif *schema.REQ_IF_CONTENT, stage *StageStruct) {
	log.Println("REQ_IF_CONTENT", "Walk")

	reqif.Name = time.Now().Format(time.DateTime)

	//
	// PARSE THE TYPE BEFORE THE OBJECTS
	//
	for _, r := range _reqif.SPEC_TYPES.SPEC_OBJECT_TYPE {
		spec_object_type := new(SPEC_OBJECT_TYPE).Stage(stage)
		reqif.SPEC_OBJECT_TYPES = append(reqif.SPEC_OBJECT_TYPES, spec_object_type)
		spec_object_type.Walk(r, stage)
	}
	for _, r := range _reqif.SPEC_TYPES.SPECIFICATION_TYPE {
		specification_type := new(SPECIFICATION_TYPE).Stage(stage)
		reqif.SPECIFICATION_TYPES = append(reqif.SPECIFICATION_TYPES, specification_type)
		specification_type.Walk(r, stage)
	}

	//
	// PARSE THE
	for _, r := range _reqif.SPECIFICATIONS.SPECIFICATION {
		specification := new(SPECIFICATION).Stage(stage)
		reqif.SPECIFICATIONS = append(reqif.SPECIFICATIONS, specification)
		specification.Walk(r, stage)
	}

}
