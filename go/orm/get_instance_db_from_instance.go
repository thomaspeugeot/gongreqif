// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/gongreqif/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | REQIFDB | REQIFHEADERDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.REQIF:
		reqifInstance := any(concreteInstance).(*models.REQIF)
		ret2 := backRepo.BackRepoREQIF.GetREQIFDBFromREQIFPtr(reqifInstance)
		ret = any(ret2).(*T2)
	case *models.REQIFHEADER:
		reqifheaderInstance := any(concreteInstance).(*models.REQIFHEADER)
		ret2 := backRepo.BackRepoREQIFHEADER.GetREQIFHEADERDBFromREQIFHEADERPtr(reqifheaderInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIFHEADER:
		tmp := GetInstanceDBFromInstance[models.REQIFHEADER, REQIFHEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
