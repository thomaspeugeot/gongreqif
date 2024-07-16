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
	case *models.ALTERNATIVEID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONDATE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONDATE":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONENUMERATION":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONINTEGER":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONREAL:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONREAL":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONSTRING:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONSTRING":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEDEFINITIONXHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONXHTML":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		}

	case *models.ATTRIBUTEVALUEBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEDATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEREAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUESTRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEXHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CORECONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPEDEFINITIONBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONBOOLEAN":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONDATE:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONDATE":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONENUMERATION":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONINTEGER":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONREAL:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONREAL":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONSTRING:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONSTRING":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPEDEFINITIONXHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONXHTML":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		}

	case *models.DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DEFAULTVALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DEFINITION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EDITABLEATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EMBEDDEDVALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUMVALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			case "ENUMVALUE":
				if _specifiedvalues, ok := stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap[inst]; ok {
					res = _specifiedvalues.Name
				}
			}
		}

	case *models.OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATIONGROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			case "RELATIONGROUP":
				if _specrelationgroups, ok := stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap[inst]; ok {
					res = _specrelationgroups.Name
				}
			}
		}

	case *models.RELATIONGROUPTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "RELATIONGROUPTYPE":
				if _spectypes, ok := stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		}

	case *models.REQIF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFCONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFHEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFTOOLEXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			case "REQIFTOOLEXTENSION":
				if _toolextensions, ok := stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap[inst]; ok {
					res = _toolextensions.Name
				}
			}
		}

	case *models.REQTYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SOURCE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SOURCESPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECHIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "CHILDREN":
			switch reverseField.Fieldname {
			case "SPECHIERARCHY":
				if _children, ok := stage.CHILDREN_SPECHIERARCHY_reverseMap[inst]; ok {
					res = _children.Name
				}
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				if _specifications, ok := stage.SPECIFICATIONS_SPECIFICATION_reverseMap[inst]; ok {
					res = _specifications.Name
				}
			}
		}

	case *models.SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATIONTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATIONTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		}

	case *models.SPECIFIEDVALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECOBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			case "SPECOBJECT":
				if _specobjects, ok := stage.SPECOBJECTS_SPECOBJECT_reverseMap[inst]; ok {
					res = _specobjects.Name
				}
			}
		}

	case *models.SPECOBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECOBJECTTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECOBJECTTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECOBJECTTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		}

	case *models.SPECRELATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONGROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECRELATIONTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECRELATIONTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		}

	case *models.SPECTYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TARGET:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TARGETSPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.THEHEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TOOLEXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XHTMLCONTENT:
		switch reverseField.GongstructName {
		// insertion point
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
	case *models.ALTERNATIVEID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONDATE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONDATE":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONENUMERATION":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONINTEGER":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONREAL:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONREAL":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONSTRING:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONSTRING":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEDEFINITIONXHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONXHTML":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTEVALUEBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEDATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEREAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUESTRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTEVALUEXHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CORECONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPEDEFINITIONBOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONBOOLEAN":
				res = stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONDATE:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONDATE":
				res = stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONENUMERATION":
				res = stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONINTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONINTEGER":
				res = stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONREAL:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONREAL":
				res = stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONSTRING:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONSTRING":
				res = stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap[inst]
			}
		}

	case *models.DATATYPEDEFINITIONXHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONXHTML":
				res = stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap[inst]
			}
		}

	case *models.DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DEFAULTVALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DEFINITION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EDITABLEATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EMBEDDEDVALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUMVALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			case "ENUMVALUE":
				res = stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap[inst]
			}
		}

	case *models.OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATIONGROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			case "RELATIONGROUP":
				res = stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap[inst]
			}
		}

	case *models.RELATIONGROUPTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "RELATIONGROUPTYPE":
				res = stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap[inst]
			}
		}

	case *models.REQIF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFCONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFHEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQIFTOOLEXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			case "REQIFTOOLEXTENSION":
				res = stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap[inst]
			}
		}

	case *models.REQTYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SOURCE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SOURCESPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECHIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "CHILDREN":
			switch reverseField.Fieldname {
			case "SPECHIERARCHY":
				res = stage.CHILDREN_SPECHIERARCHY_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				res = stage.SPECIFICATIONS_SPECIFICATION_reverseMap[inst]
			}
		}

	case *models.SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATIONTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATIONTYPE":
				res = stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap[inst]
			}
		}

	case *models.SPECIFIEDVALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECOBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			case "SPECOBJECT":
				res = stage.SPECOBJECTS_SPECOBJECT_reverseMap[inst]
			}
		}

	case *models.SPECOBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECOBJECTTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECOBJECTTYPE":
				res = stage.SPECTYPES_SPECOBJECTTYPE_reverseMap[inst]
			}
		}

	case *models.SPECRELATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONGROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECRELATIONTYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECRELATIONTYPE":
				res = stage.SPECTYPES_SPECRELATIONTYPE_reverseMap[inst]
			}
		}

	case *models.SPECTYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TARGET:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TARGETSPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.THEHEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TOOLEXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XHTMLCONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
