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
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVEID, ALTERNATIVEIDDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONBOOLEAN, ATTRIBUTEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONDATE, ATTRIBUTEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONDATE":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONENUMERATION, ATTRIBUTEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONENUMERATION":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONINTEGER, ATTRIBUTEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONINTEGER":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONREAL, ATTRIBUTEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONREAL":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONSTRING, ATTRIBUTEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONSTRING":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONXHTML, ATTRIBUTEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONXHTML":
				if _specattributes, ok := stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap[inst]; ok {
					res = _specattributes.Name
				}
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEBOOLEAN, ATTRIBUTEVALUEBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEDATE, ATTRIBUTEVALUEDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEENUMERATION, ATTRIBUTEVALUEENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEINTEGER, ATTRIBUTEVALUEINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEREAL, ATTRIBUTEVALUEREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUESTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUESTRING, ATTRIBUTEVALUESTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEXHTML, ATTRIBUTEVALUEXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.CHILDREN:
		tmp := GetInstanceDBFromInstance[models.CHILDREN, CHILDRENDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.CORECONTENT:
		tmp := GetInstanceDBFromInstance[models.CORECONTENT, CORECONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONBOOLEAN, DATATYPEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONBOOLEAN":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONDATE, DATATYPEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONDATE":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONENUMERATION, DATATYPEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONENUMERATION":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONINTEGER, DATATYPEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONINTEGER":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONREAL, DATATYPEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONREAL":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONSTRING, DATATYPEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONSTRING":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONXHTML, DATATYPEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONXHTML":
				if _datatypes, ok := stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap[inst]; ok {
					res = _datatypes.Name
				}
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPES:
		tmp := GetInstanceDBFromInstance[models.DATATYPES, DATATYPESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DEFAULTVALUE:
		tmp := GetInstanceDBFromInstance[models.DEFAULTVALUE, DEFAULTVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DEFINITION:
		tmp := GetInstanceDBFromInstance[models.DEFINITION, DEFINITIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.EDITABLEATTS:
		tmp := GetInstanceDBFromInstance[models.EDITABLEATTS, EDITABLEATTSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.EMBEDDEDVALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDEDVALUE, EMBEDDEDVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ENUMVALUE:
		tmp := GetInstanceDBFromInstance[models.ENUMVALUE, ENUMVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			case "ENUMVALUE":
				if _specifiedvalues, ok := stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap[inst]; ok {
					res = _specifiedvalues.Name
				}
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.OBJECT:
		tmp := GetInstanceDBFromInstance[models.OBJECT, OBJECTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.PROPERTIES, PROPERTIESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.RELATIONGROUP:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUP, RELATIONGROUPDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			case "RELATIONGROUP":
				if _specrelationgroups, ok := stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap[inst]; ok {
					res = _specrelationgroups.Name
				}
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.RELATIONGROUPTYPE:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUPTYPE, RELATIONGROUPTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "RELATIONGROUPTYPE":
				if _spectypes, ok := stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFCONTENT:
		tmp := GetInstanceDBFromInstance[models.REQIFCONTENT, REQIFCONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
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
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFTOOLEXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQIFTOOLEXTENSION, REQIFTOOLEXTENSIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			case "REQIFTOOLEXTENSION":
				if _toolextensions, ok := stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap[inst]; ok {
					res = _toolextensions.Name
				}
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFTYPE:
		tmp := GetInstanceDBFromInstance[models.REQIFTYPE, REQIFTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SOURCE:
		tmp := GetInstanceDBFromInstance[models.SOURCE, SOURCEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SOURCESPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SOURCESPECIFICATION, SOURCESPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.SPECATTRIBUTES, SPECATTRIBUTESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECHIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPECHIERARCHY, SPECHIERARCHYDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			case "SPECHIERARCHY":
				if _children, ok := stage.CHILDREN_SPECHIERARCHY_reverseMap[inst]; ok {
					res = _children.Name
				}
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				if _specifications, ok := stage.SPECIFICATIONS_SPECIFICATION_reverseMap[inst]; ok {
					res = _specifications.Name
				}
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONS, SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONTYPE, SPECIFICATIONTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATIONTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFIEDVALUES:
		tmp := GetInstanceDBFromInstance[models.SPECIFIEDVALUES, SPECIFIEDVALUESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECT:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECT, SPECOBJECTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			case "SPECOBJECT":
				if _specobjects, ok := stage.SPECOBJECTS_SPECOBJECT_reverseMap[inst]; ok {
					res = _specobjects.Name
				}
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECTS:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTS, SPECOBJECTSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECTTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTTYPE, SPECOBJECTTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECOBJECTTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECOBJECTTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATION:
		tmp := GetInstanceDBFromInstance[models.SPECRELATION, SPECRELATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONGROUPS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONGROUPS, SPECRELATIONGROUPSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONS, SPECRELATIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONTYPE, SPECRELATIONTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECRELATIONTYPE":
				if _spectypes, ok := stage.SPECTYPES_SPECRELATIONTYPE_reverseMap[inst]; ok {
					res = _spectypes.Name
				}
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECTYPES:
		tmp := GetInstanceDBFromInstance[models.SPECTYPES, SPECTYPESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TARGET:
		tmp := GetInstanceDBFromInstance[models.TARGET, TARGETDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TARGETSPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.TARGETSPECIFICATION, TARGETSPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.THEHEADER:
		tmp := GetInstanceDBFromInstance[models.THEHEADER, THEHEADERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TOOLEXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.TOOLEXTENSIONS, TOOLEXTENSIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.VALUES:
		tmp := GetInstanceDBFromInstance[models.VALUES, VALUESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.XHTMLCONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTMLCONTENT, XHTMLCONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
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
	case *models.ALTERNATIVEID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVEID, ALTERNATIVEIDDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONBOOLEAN, ATTRIBUTEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONBOOLEAN":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONDATE, ATTRIBUTEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONDATE":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONENUMERATION, ATTRIBUTEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONENUMERATION":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONINTEGER, ATTRIBUTEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONINTEGER":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONREAL, ATTRIBUTEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONREAL":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONSTRING, ATTRIBUTEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONSTRING":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEDEFINITIONXHTML, ATTRIBUTEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTEDEFINITIONXHTML":
				res = stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap[inst]
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEBOOLEAN, ATTRIBUTEVALUEBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEDATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEDATE, ATTRIBUTEVALUEDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEENUMERATION, ATTRIBUTEVALUEENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEINTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEINTEGER, ATTRIBUTEVALUEINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEREAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEREAL, ATTRIBUTEVALUEREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUESTRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUESTRING, ATTRIBUTEVALUESTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ATTRIBUTEVALUEXHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTEVALUEXHTML, ATTRIBUTEVALUEXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.CHILDREN:
		tmp := GetInstanceDBFromInstance[models.CHILDREN, CHILDRENDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.CORECONTENT:
		tmp := GetInstanceDBFromInstance[models.CORECONTENT, CORECONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONBOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONBOOLEAN, DATATYPEDEFINITIONBOOLEANDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONBOOLEAN":
				res = stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONDATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONDATE, DATATYPEDEFINITIONDATEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONDATE":
				res = stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONENUMERATION, DATATYPEDEFINITIONENUMERATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONENUMERATION":
				res = stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONINTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONINTEGER, DATATYPEDEFINITIONINTEGERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONINTEGER":
				res = stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONREAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONREAL, DATATYPEDEFINITIONREALDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONREAL":
				res = stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONSTRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONSTRING, DATATYPEDEFINITIONSTRINGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONSTRING":
				res = stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPEDEFINITIONXHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPEDEFINITIONXHTML, DATATYPEDEFINITIONXHTMLDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPEDEFINITIONXHTML":
				res = stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap[inst]
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DATATYPES:
		tmp := GetInstanceDBFromInstance[models.DATATYPES, DATATYPESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DEFAULTVALUE:
		tmp := GetInstanceDBFromInstance[models.DEFAULTVALUE, DEFAULTVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.DEFINITION:
		tmp := GetInstanceDBFromInstance[models.DEFINITION, DEFINITIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.EDITABLEATTS:
		tmp := GetInstanceDBFromInstance[models.EDITABLEATTS, EDITABLEATTSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.EMBEDDEDVALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDEDVALUE, EMBEDDEDVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.ENUMVALUE:
		tmp := GetInstanceDBFromInstance[models.ENUMVALUE, ENUMVALUEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			case "ENUMVALUE":
				res = stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap[inst]
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.OBJECT:
		tmp := GetInstanceDBFromInstance[models.OBJECT, OBJECTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.PROPERTIES, PROPERTIESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.RELATIONGROUP:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUP, RELATIONGROUPDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			case "RELATIONGROUP":
				res = stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap[inst]
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.RELATIONGROUPTYPE:
		tmp := GetInstanceDBFromInstance[models.RELATIONGROUPTYPE, RELATIONGROUPTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "RELATIONGROUPTYPE":
				res = stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap[inst]
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIF:
		tmp := GetInstanceDBFromInstance[models.REQIF, REQIFDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFCONTENT:
		tmp := GetInstanceDBFromInstance[models.REQIFCONTENT, REQIFCONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
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
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFTOOLEXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQIFTOOLEXTENSION, REQIFTOOLEXTENSIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			case "REQIFTOOLEXTENSION":
				res = stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap[inst]
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.REQIFTYPE:
		tmp := GetInstanceDBFromInstance[models.REQIFTYPE, REQIFTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SOURCE:
		tmp := GetInstanceDBFromInstance[models.SOURCE, SOURCEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SOURCESPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SOURCESPECIFICATION, SOURCESPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.SPECATTRIBUTES, SPECATTRIBUTESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECHIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPECHIERARCHY, SPECHIERARCHYDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			case "SPECHIERARCHY":
				res = stage.CHILDREN_SPECHIERARCHY_reverseMap[inst]
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				res = stage.SPECIFICATIONS_SPECIFICATION_reverseMap[inst]
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONS, SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFICATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATIONTYPE, SPECIFICATIONTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATIONTYPE":
				res = stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap[inst]
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECIFIEDVALUES:
		tmp := GetInstanceDBFromInstance[models.SPECIFIEDVALUES, SPECIFIEDVALUESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECT:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECT, SPECOBJECTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			case "SPECOBJECT":
				res = stage.SPECOBJECTS_SPECOBJECT_reverseMap[inst]
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECTS:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTS, SPECOBJECTSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECOBJECTTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECOBJECTTYPE, SPECOBJECTTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECOBJECTTYPE":
				res = stage.SPECTYPES_SPECOBJECTTYPE_reverseMap[inst]
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATION:
		tmp := GetInstanceDBFromInstance[models.SPECRELATION, SPECRELATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONGROUPS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONGROUPS, SPECRELATIONGROUPSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONS:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONS, SPECRELATIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECRELATIONTYPE:
		tmp := GetInstanceDBFromInstance[models.SPECRELATIONTYPE, SPECRELATIONTYPEDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			case "SPECRELATIONTYPE":
				res = stage.SPECTYPES_SPECRELATIONTYPE_reverseMap[inst]
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.SPECTYPES:
		tmp := GetInstanceDBFromInstance[models.SPECTYPES, SPECTYPESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TARGET:
		tmp := GetInstanceDBFromInstance[models.TARGET, TARGETDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TARGETSPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.TARGETSPECIFICATION, TARGETSPECIFICATIONDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.THEHEADER:
		tmp := GetInstanceDBFromInstance[models.THEHEADER, THEHEADERDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.TOOLEXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.TOOLEXTENSIONS, TOOLEXTENSIONSDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.VALUES:
		tmp := GetInstanceDBFromInstance[models.VALUES, VALUESDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	case *models.XHTMLCONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTMLCONTENT, XHTMLCONTENTDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "ALTERNATIVEID":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEDATE":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEENUMERATION":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEINTEGER":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEREAL":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUESTRING":
			switch reverseField.Fieldname {
			}
		case "ATTRIBUTEVALUEXHTML":
			switch reverseField.Fieldname {
			}
		case "CHILDREN":
			switch reverseField.Fieldname {
			}
		case "CORECONTENT":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONBOOLEAN":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONDATE":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONENUMERATION":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONINTEGER":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONREAL":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONSTRING":
			switch reverseField.Fieldname {
			}
		case "DATATYPEDEFINITIONXHTML":
			switch reverseField.Fieldname {
			}
		case "DATATYPES":
			switch reverseField.Fieldname {
			}
		case "DEFAULTVALUE":
			switch reverseField.Fieldname {
			}
		case "DEFINITION":
			switch reverseField.Fieldname {
			}
		case "EDITABLEATTS":
			switch reverseField.Fieldname {
			}
		case "EMBEDDEDVALUE":
			switch reverseField.Fieldname {
			}
		case "ENUMVALUE":
			switch reverseField.Fieldname {
			}
		case "OBJECT":
			switch reverseField.Fieldname {
			}
		case "PROPERTIES":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUP":
			switch reverseField.Fieldname {
			}
		case "RELATIONGROUPTYPE":
			switch reverseField.Fieldname {
			}
		case "REQIF":
			switch reverseField.Fieldname {
			}
		case "REQIFCONTENT":
			switch reverseField.Fieldname {
			}
		case "REQIFHEADER":
			switch reverseField.Fieldname {
			}
		case "REQIFTOOLEXTENSION":
			switch reverseField.Fieldname {
			}
		case "REQIFTYPE":
			switch reverseField.Fieldname {
			}
		case "SOURCE":
			switch reverseField.Fieldname {
			}
		case "SOURCESPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECATTRIBUTES":
			switch reverseField.Fieldname {
			}
		case "SPECHIERARCHY":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECIFICATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECIFIEDVALUES":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECT":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTS":
			switch reverseField.Fieldname {
			}
		case "SPECOBJECTTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECRELATION":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONGROUPS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONS":
			switch reverseField.Fieldname {
			}
		case "SPECRELATIONTYPE":
			switch reverseField.Fieldname {
			}
		case "SPECTYPES":
			switch reverseField.Fieldname {
			}
		case "TARGET":
			switch reverseField.Fieldname {
			}
		case "TARGETSPECIFICATION":
			switch reverseField.Fieldname {
			}
		case "THEHEADER":
			switch reverseField.Fieldname {
			}
		case "TOOLEXTENSIONS":
			switch reverseField.Fieldname {
			}
		case "VALUES":
			switch reverseField.Fieldname {
			}
		case "XHTMLCONTENT":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
