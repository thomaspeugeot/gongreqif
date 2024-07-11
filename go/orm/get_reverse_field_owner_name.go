// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/gongreqif/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
