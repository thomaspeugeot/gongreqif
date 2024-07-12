// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/gongreqif/go/models"
	"github.com/thomaspeugeot/gongreqif/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVEID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.ATTRIBUTEDEFINITIONBOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONBOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONBOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONBOOLEAN](
					nil,
					"ATTRIBUTEDEFINITIONBOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONDATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONDATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONDATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONDATE](
					nil,
					"ATTRIBUTEDEFINITIONDATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MULTIVALUEDAttr", instanceWithInferedType.MULTIVALUEDAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONENUMERATION](
					nil,
					"ATTRIBUTEDEFINITIONENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONINTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONINTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONINTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONINTEGER](
					nil,
					"ATTRIBUTEDEFINITIONINTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONREAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONREAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONREAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONREAL](
					nil,
					"ATTRIBUTEDEFINITIONREAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONSTRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONSTRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONSTRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONSTRING](
					nil,
					"ATTRIBUTEDEFINITIONSTRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEDEFINITIONXHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("DEFAULTVALUE", instanceWithInferedType.DEFAULTVALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECATTRIBUTES"
			rf.Fieldname = "ATTRIBUTEDEFINITIONXHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECATTRIBUTES),
					"ATTRIBUTEDEFINITIONXHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECATTRIBUTES, *models.ATTRIBUTEDEFINITIONXHTML](
					nil,
					"ATTRIBUTEDEFINITIONXHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTEVALUEBOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THEVALUEAttr", instanceWithInferedType.THEVALUEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTEVALUEDATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THEVALUEAttr", instanceWithInferedType.THEVALUEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTEVALUEENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)

	case *models.ATTRIBUTEVALUEINTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THEVALUEAttr", instanceWithInferedType.THEVALUEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTEVALUEREAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THEVALUEAttr", instanceWithInferedType.THEVALUEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTEVALUESTRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THEVALUEAttr", instanceWithInferedType.THEVALUEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTEVALUEXHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISSIMPLIFIEDAttr", instanceWithInferedType.ISSIMPLIFIEDAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("THEVALUE", instanceWithInferedType.THEVALUE, formGroup, probe)
		AssociationFieldToForm("THEORIGINALVALUE", instanceWithInferedType.THEORIGINALVALUE, formGroup, probe)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.CHILDREN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPECHIERARCHY", instanceWithInferedType, &instanceWithInferedType.SPECHIERARCHY, formGroup, probe)

	case *models.CORECONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQIFCONTENT", instanceWithInferedType.REQIFCONTENT, formGroup, probe)

	case *models.DATATYPEDEFINITIONBOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONBOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONBOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONBOOLEAN](
					nil,
					"DATATYPEDEFINITIONBOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONDATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONDATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONDATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONDATE](
					nil,
					"DATATYPEDEFINITIONDATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SPECIFIEDVALUES", instanceWithInferedType.SPECIFIEDVALUES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONENUMERATION](
					nil,
					"DATATYPEDEFINITIONENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONINTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAXAttr", instanceWithInferedType.MAXAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MINAttr", instanceWithInferedType.MINAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONINTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONINTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONINTEGER](
					nil,
					"DATATYPEDEFINITIONINTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONREAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ACCURACYAttr", instanceWithInferedType.ACCURACYAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAXAttr", instanceWithInferedType.MAXAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MINAttr", instanceWithInferedType.MINAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONREAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONREAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONREAL](
					nil,
					"DATATYPEDEFINITIONREAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONSTRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAXLENGTHAttr", instanceWithInferedType.MAXLENGTHAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONSTRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONSTRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONSTRING](
					nil,
					"DATATYPEDEFINITIONSTRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPEDEFINITIONXHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPES"
			rf.Fieldname = "DATATYPEDEFINITIONXHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPES),
					"DATATYPEDEFINITIONXHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPES, *models.DATATYPEDEFINITIONXHTML](
					nil,
					"DATATYPEDEFINITIONXHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DATATYPEDEFINITIONBOOLEAN", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONBOOLEAN, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONDATE", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONDATE, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONENUMERATION", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONENUMERATION, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONINTEGER", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONINTEGER, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONREAL", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONREAL, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONSTRING", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONSTRING, formGroup, probe)
		AssociationSliceToForm("DATATYPEDEFINITIONXHTML", instanceWithInferedType, &instanceWithInferedType.DATATYPEDEFINITIONXHTML, formGroup, probe)

	case *models.DEFAULTVALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTEVALUEBOOLEAN", instanceWithInferedType.ATTRIBUTEVALUEBOOLEAN, formGroup, probe)

	case *models.DEFINITION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTEDEFINITIONBOOLEANREF", instanceWithInferedType.ATTRIBUTEDEFINITIONBOOLEANREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.EDITABLEATTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.EMBEDDEDVALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("KEYAttr", instanceWithInferedType.KEYAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OTHERCONTENTAttr", instanceWithInferedType.OTHERCONTENTAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.ENUMVALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("PROPERTIES", instanceWithInferedType.PROPERTIES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFIEDVALUES"
			rf.Fieldname = "ENUMVALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFIEDVALUES),
					"ENUMVALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFIEDVALUES, *models.ENUMVALUE](
					nil,
					"ENUMVALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECOBJECTREF", instanceWithInferedType.SPECOBJECTREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.PROPERTIES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("EMBEDDEDVALUE", instanceWithInferedType.EMBEDDEDVALUE, formGroup, probe)

	case *models.RELATIONGROUP:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SOURCESPECIFICATION", instanceWithInferedType.SOURCESPECIFICATION, formGroup, probe)
		AssociationFieldToForm("SPECRELATIONS", instanceWithInferedType.SPECRELATIONS, formGroup, probe)
		AssociationFieldToForm("TARGETSPECIFICATION", instanceWithInferedType.TARGETSPECIFICATION, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECRELATIONGROUPS"
			rf.Fieldname = "RELATIONGROUP"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECRELATIONGROUPS),
					"RELATIONGROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECRELATIONGROUPS, *models.RELATIONGROUP](
					nil,
					"RELATIONGROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RELATIONGROUPTYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SPECATTRIBUTES", instanceWithInferedType.SPECATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "RELATIONGROUPTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECTYPES),
					"RELATIONGROUPTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECTYPES, *models.RELATIONGROUPTYPE](
					nil,
					"RELATIONGROUPTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQIF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("HEADER", instanceWithInferedType.HEADER, formGroup, probe)
		AssociationFieldToForm("CORECONTENT", instanceWithInferedType.CORECONTENT, formGroup, probe)

	case *models.REQIFCONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DATATYPES", instanceWithInferedType.DATATYPES, formGroup, probe)
		AssociationFieldToForm("SPECTYPES", instanceWithInferedType.SPECTYPES, formGroup, probe)
		AssociationFieldToForm("SPECOBJECTS", instanceWithInferedType.SPECOBJECTS, formGroup, probe)
		AssociationFieldToForm("SPECRELATIONS", instanceWithInferedType.SPECRELATIONS, formGroup, probe)
		AssociationFieldToForm("SPECIFICATIONS", instanceWithInferedType.SPECIFICATIONS, formGroup, probe)
		AssociationFieldToForm("SPECRELATIONGROUPS", instanceWithInferedType.SPECRELATIONGROUPS, formGroup, probe)

	case *models.REQIFHEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("COMMENT", instanceWithInferedType.COMMENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CREATIONTIME", instanceWithInferedType.CREATIONTIME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REPOSITORYID", instanceWithInferedType.REPOSITORYID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQIFTOOLID", instanceWithInferedType.REQIFTOOLID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQIFVERSION", instanceWithInferedType.REQIFVERSION, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SOURCETOOLID", instanceWithInferedType.SOURCETOOLID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TITLE", instanceWithInferedType.TITLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.REQIFTOOLEXTENSION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TOOLEXTENSIONS"
			rf.Fieldname = "REQIFTOOLEXTENSION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TOOLEXTENSIONS),
					"REQIFTOOLEXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TOOLEXTENSIONS, *models.REQIFTOOLEXTENSION](
					nil,
					"REQIFTOOLEXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQIFTYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPEDEFINITIONBOOLEANREF", instanceWithInferedType.DATATYPEDEFINITIONBOOLEANREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SOURCE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECOBJECTREF", instanceWithInferedType.SPECOBJECTREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SOURCESPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATIONREF", instanceWithInferedType.SPECIFICATIONREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SPECATTRIBUTES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONBOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONBOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONDATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONDATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONINTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONINTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONREAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONREAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONSTRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONSTRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTEDEFINITIONXHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTEDEFINITIONXHTML, formGroup, probe)

	case *models.SPECHIERARCHY:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISEDITABLEAttr", instanceWithInferedType.ISEDITABLEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ISTABLEINTERNALAttr", instanceWithInferedType.ISTABLEINTERNALAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("EDITABLEATTS", instanceWithInferedType.EDITABLEATTS, formGroup, probe)
		AssociationFieldToForm("OBJECT", instanceWithInferedType.OBJECT, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "CHILDREN"
			rf.Fieldname = "SPECHIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.CHILDREN),
					"SPECHIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.CHILDREN, *models.SPECHIERARCHY](
					nil,
					"SPECHIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATIONS"
			rf.Fieldname = "SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATIONS),
					"SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATIONS, *models.SPECIFICATION](
					nil,
					"SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.SPECIFICATION, formGroup, probe)

	case *models.SPECIFICATIONTYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SPECATTRIBUTES", instanceWithInferedType.SPECATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECIFICATIONTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECTYPES),
					"SPECIFICATIONTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECTYPES, *models.SPECIFICATIONTYPE](
					nil,
					"SPECIFICATIONTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFIEDVALUES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ENUMVALUE", instanceWithInferedType, &instanceWithInferedType.ENUMVALUE, formGroup, probe)

	case *models.SPECOBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECOBJECTS"
			rf.Fieldname = "SPECOBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECOBJECTS),
					"SPECOBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECOBJECTS, *models.SPECOBJECT](
					nil,
					"SPECOBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECOBJECTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPECOBJECT", instanceWithInferedType, &instanceWithInferedType.SPECOBJECT, formGroup, probe)

	case *models.SPECOBJECTTYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SPECATTRIBUTES", instanceWithInferedType.SPECATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECOBJECTTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECTYPES),
					"SPECOBJECTTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECTYPES, *models.SPECOBJECTTYPE](
					nil,
					"SPECOBJECTTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECRELATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("SOURCE", instanceWithInferedType.SOURCE, formGroup, probe)
		AssociationFieldToForm("TARGET", instanceWithInferedType.TARGET, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

	case *models.SPECRELATIONGROUPS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RELATIONGROUP", instanceWithInferedType, &instanceWithInferedType.RELATIONGROUP, formGroup, probe)

	case *models.SPECRELATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.SPECRELATIONTYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESCAttr", instanceWithInferedType.DESCAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIERAttr", instanceWithInferedType.IDENTIFIERAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LASTCHANGEAttr", instanceWithInferedType.LASTCHANGEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONGNAMEAttr", instanceWithInferedType.LONGNAMEAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVEID", instanceWithInferedType.ALTERNATIVEID, formGroup, probe)
		AssociationFieldToForm("SPECATTRIBUTES", instanceWithInferedType.SPECATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECTYPES"
			rf.Fieldname = "SPECRELATIONTYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECTYPES),
					"SPECRELATIONTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECTYPES, *models.SPECRELATIONTYPE](
					nil,
					"SPECRELATIONTYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECTYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RELATIONGROUPTYPE", instanceWithInferedType, &instanceWithInferedType.RELATIONGROUPTYPE, formGroup, probe)
		AssociationSliceToForm("SPECOBJECTTYPE", instanceWithInferedType, &instanceWithInferedType.SPECOBJECTTYPE, formGroup, probe)
		AssociationSliceToForm("SPECRELATIONTYPE", instanceWithInferedType, &instanceWithInferedType.SPECRELATIONTYPE, formGroup, probe)
		AssociationSliceToForm("SPECIFICATIONTYPE", instanceWithInferedType, &instanceWithInferedType.SPECIFICATIONTYPE, formGroup, probe)

	case *models.TARGET:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECOBJECTREF", instanceWithInferedType.SPECOBJECTREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.TARGETSPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATIONREF", instanceWithInferedType.SPECIFICATIONREF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.THEHEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQIFHEADER", instanceWithInferedType.REQIFHEADER, formGroup, probe)

	case *models.TOOLEXTENSIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("REQIFTOOLEXTENSION", instanceWithInferedType, &instanceWithInferedType.REQIFTOOLEXTENSION, formGroup, probe)

	case *models.VALUES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.XHTMLCONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
