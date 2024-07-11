// generated code - do not edit
package orm

import (
	"github.com/thomaspeugeot/gongreqif/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | CONTENTDB | HEADERDB | REQIFDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CONTENT:
		contentInstance := any(concreteInstance).(*models.CONTENT)
		ret2 := backRepo.BackRepoCONTENT.GetCONTENTDBFromCONTENTPtr(contentInstance)
		ret = any(ret2).(*T2)
	case *models.HEADER:
		headerInstance := any(concreteInstance).(*models.HEADER)
		ret2 := backRepo.BackRepoHEADER.GetHEADERDBFromHEADERPtr(headerInstance)
		ret = any(ret2).(*T2)
	case *models.REQIF:
		reqifInstance := any(concreteInstance).(*models.REQIF)
		ret2 := backRepo.BackRepoREQIF.GetREQIFDBFromREQIFPtr(reqifInstance)
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
	case *models.CONTENT:
		tmp := GetInstanceDBFromInstance[models.CONTENT, CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.HEADER:
		tmp := GetInstanceDBFromInstance[models.HEADER, HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
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
	case *models.CONTENT:
		tmp := GetInstanceDBFromInstance[models.CONTENT, CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.HEADER:
		tmp := GetInstanceDBFromInstance[models.HEADER, HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
