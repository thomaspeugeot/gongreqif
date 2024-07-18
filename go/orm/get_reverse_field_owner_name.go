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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPES":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPECIFICATION_TYPES_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPES":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS":
				res = stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPES":
				res = stage.REQ_IF_CONTENT_SPECIFICATION_TYPES_reverseMap[inst]
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPES":
				res = stage.REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
