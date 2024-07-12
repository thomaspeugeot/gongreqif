// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *ALTERNATIVEID:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONBOOLEAN:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONDATE:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONENUMERATION:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONINTEGER:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONREAL:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONSTRING:
		// insertion point per field

	case *ATTRIBUTEDEFINITIONXHTML:
		// insertion point per field

	case *ATTRIBUTEVALUEBOOLEAN:
		// insertion point per field

	case *ATTRIBUTEVALUEDATE:
		// insertion point per field

	case *ATTRIBUTEVALUEENUMERATION:
		// insertion point per field

	case *ATTRIBUTEVALUEINTEGER:
		// insertion point per field

	case *ATTRIBUTEVALUEREAL:
		// insertion point per field

	case *ATTRIBUTEVALUESTRING:
		// insertion point per field

	case *ATTRIBUTEVALUEXHTML:
		// insertion point per field

	case *CHILDREN:
		// insertion point per field
		if fieldName == "SPECHIERARCHY" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*CHILDREN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*CHILDREN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECHIERARCHY).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECHIERARCHY = _inferedTypeInstance.SPECHIERARCHY[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECHIERARCHY =
								append(_inferedTypeInstance.SPECHIERARCHY, any(fieldInstance).(*SPECHIERARCHY))
						}
					}
				}
			}
		}

	case *CORECONTENT:
		// insertion point per field

	case *DATATYPEDEFINITIONBOOLEAN:
		// insertion point per field

	case *DATATYPEDEFINITIONDATE:
		// insertion point per field

	case *DATATYPEDEFINITIONENUMERATION:
		// insertion point per field

	case *DATATYPEDEFINITIONINTEGER:
		// insertion point per field

	case *DATATYPEDEFINITIONREAL:
		// insertion point per field

	case *DATATYPEDEFINITIONSTRING:
		// insertion point per field

	case *DATATYPEDEFINITIONXHTML:
		// insertion point per field

	case *DATATYPES:
		// insertion point per field
		if fieldName == "DATATYPEDEFINITIONBOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONBOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONBOOLEAN = _inferedTypeInstance.DATATYPEDEFINITIONBOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONBOOLEAN =
								append(_inferedTypeInstance.DATATYPEDEFINITIONBOOLEAN, any(fieldInstance).(*DATATYPEDEFINITIONBOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONDATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONDATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONDATE = _inferedTypeInstance.DATATYPEDEFINITIONDATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONDATE =
								append(_inferedTypeInstance.DATATYPEDEFINITIONDATE, any(fieldInstance).(*DATATYPEDEFINITIONDATE))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONENUMERATION = _inferedTypeInstance.DATATYPEDEFINITIONENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONENUMERATION =
								append(_inferedTypeInstance.DATATYPEDEFINITIONENUMERATION, any(fieldInstance).(*DATATYPEDEFINITIONENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONINTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONINTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONINTEGER = _inferedTypeInstance.DATATYPEDEFINITIONINTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONINTEGER =
								append(_inferedTypeInstance.DATATYPEDEFINITIONINTEGER, any(fieldInstance).(*DATATYPEDEFINITIONINTEGER))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONREAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONREAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONREAL = _inferedTypeInstance.DATATYPEDEFINITIONREAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONREAL =
								append(_inferedTypeInstance.DATATYPEDEFINITIONREAL, any(fieldInstance).(*DATATYPEDEFINITIONREAL))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONSTRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONSTRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONSTRING = _inferedTypeInstance.DATATYPEDEFINITIONSTRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONSTRING =
								append(_inferedTypeInstance.DATATYPEDEFINITIONSTRING, any(fieldInstance).(*DATATYPEDEFINITIONSTRING))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPEDEFINITIONXHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPEDEFINITIONXHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPEDEFINITIONXHTML = _inferedTypeInstance.DATATYPEDEFINITIONXHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPEDEFINITIONXHTML =
								append(_inferedTypeInstance.DATATYPEDEFINITIONXHTML, any(fieldInstance).(*DATATYPEDEFINITIONXHTML))
						}
					}
				}
			}
		}

	case *DEFAULTVALUE:
		// insertion point per field

	case *DEFINITION:
		// insertion point per field

	case *EDITABLEATTS:
		// insertion point per field

	case *EMBEDDEDVALUE:
		// insertion point per field

	case *ENUMVALUE:
		// insertion point per field

	case *OBJECT:
		// insertion point per field

	case *PROPERTIES:
		// insertion point per field

	case *RELATIONGROUP:
		// insertion point per field

	case *RELATIONGROUPTYPE:
		// insertion point per field

	case *REQIF:
		// insertion point per field

	case *REQIFCONTENT:
		// insertion point per field

	case *REQIFHEADER:
		// insertion point per field

	case *REQIFTOOLEXTENSION:
		// insertion point per field

	case *REQIFTYPE:
		// insertion point per field

	case *SOURCE:
		// insertion point per field

	case *SOURCESPECIFICATION:
		// insertion point per field

	case *SPECATTRIBUTES:
		// insertion point per field
		if fieldName == "ATTRIBUTEDEFINITIONBOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONBOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONBOOLEAN = _inferedTypeInstance.ATTRIBUTEDEFINITIONBOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONBOOLEAN =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONBOOLEAN, any(fieldInstance).(*ATTRIBUTEDEFINITIONBOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONDATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONDATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONDATE = _inferedTypeInstance.ATTRIBUTEDEFINITIONDATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONDATE =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONDATE, any(fieldInstance).(*ATTRIBUTEDEFINITIONDATE))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONENUMERATION = _inferedTypeInstance.ATTRIBUTEDEFINITIONENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONENUMERATION =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONENUMERATION, any(fieldInstance).(*ATTRIBUTEDEFINITIONENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONINTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONINTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONINTEGER = _inferedTypeInstance.ATTRIBUTEDEFINITIONINTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONINTEGER =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONINTEGER, any(fieldInstance).(*ATTRIBUTEDEFINITIONINTEGER))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONREAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONREAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONREAL = _inferedTypeInstance.ATTRIBUTEDEFINITIONREAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONREAL =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONREAL, any(fieldInstance).(*ATTRIBUTEDEFINITIONREAL))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONSTRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONSTRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONSTRING = _inferedTypeInstance.ATTRIBUTEDEFINITIONSTRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONSTRING =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONSTRING, any(fieldInstance).(*ATTRIBUTEDEFINITIONSTRING))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTEDEFINITIONXHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTEDEFINITIONXHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTEDEFINITIONXHTML = _inferedTypeInstance.ATTRIBUTEDEFINITIONXHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTEDEFINITIONXHTML =
								append(_inferedTypeInstance.ATTRIBUTEDEFINITIONXHTML, any(fieldInstance).(*ATTRIBUTEDEFINITIONXHTML))
						}
					}
				}
			}
		}

	case *SPECHIERARCHY:
		// insertion point per field

	case *SPECIFICATION:
		// insertion point per field

	case *SPECIFICATIONS:
		// insertion point per field
		if fieldName == "SPECIFICATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATIONS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATIONS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATION = _inferedTypeInstance.SPECIFICATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATION =
								append(_inferedTypeInstance.SPECIFICATION, any(fieldInstance).(*SPECIFICATION))
						}
					}
				}
			}
		}

	case *SPECIFICATIONTYPE:
		// insertion point per field

	case *SPECIFIEDVALUES:
		// insertion point per field
		if fieldName == "ENUMVALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFIEDVALUES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFIEDVALUES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ENUMVALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ENUMVALUE = _inferedTypeInstance.ENUMVALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ENUMVALUE =
								append(_inferedTypeInstance.ENUMVALUE, any(fieldInstance).(*ENUMVALUE))
						}
					}
				}
			}
		}

	case *SPECOBJECT:
		// insertion point per field

	case *SPECOBJECTS:
		// insertion point per field
		if fieldName == "SPECOBJECT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECOBJECTS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECOBJECTS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECOBJECT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECOBJECT = _inferedTypeInstance.SPECOBJECT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECOBJECT =
								append(_inferedTypeInstance.SPECOBJECT, any(fieldInstance).(*SPECOBJECT))
						}
					}
				}
			}
		}

	case *SPECOBJECTTYPE:
		// insertion point per field

	case *SPECRELATION:
		// insertion point per field

	case *SPECRELATIONGROUPS:
		// insertion point per field
		if fieldName == "RELATIONGROUP" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECRELATIONGROUPS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECRELATIONGROUPS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RELATIONGROUP).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RELATIONGROUP = _inferedTypeInstance.RELATIONGROUP[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RELATIONGROUP =
								append(_inferedTypeInstance.RELATIONGROUP, any(fieldInstance).(*RELATIONGROUP))
						}
					}
				}
			}
		}

	case *SPECRELATIONS:
		// insertion point per field

	case *SPECRELATIONTYPE:
		// insertion point per field

	case *SPECTYPES:
		// insertion point per field
		if fieldName == "RELATIONGROUPTYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECTYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECTYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RELATIONGROUPTYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RELATIONGROUPTYPE = _inferedTypeInstance.RELATIONGROUPTYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RELATIONGROUPTYPE =
								append(_inferedTypeInstance.RELATIONGROUPTYPE, any(fieldInstance).(*RELATIONGROUPTYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPECOBJECTTYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECTYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECTYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECOBJECTTYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECOBJECTTYPE = _inferedTypeInstance.SPECOBJECTTYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECOBJECTTYPE =
								append(_inferedTypeInstance.SPECOBJECTTYPE, any(fieldInstance).(*SPECOBJECTTYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPECRELATIONTYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECTYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECTYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECRELATIONTYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECRELATIONTYPE = _inferedTypeInstance.SPECRELATIONTYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECRELATIONTYPE =
								append(_inferedTypeInstance.SPECRELATIONTYPE, any(fieldInstance).(*SPECRELATIONTYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFICATIONTYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECTYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECTYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATIONTYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATIONTYPE = _inferedTypeInstance.SPECIFICATIONTYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATIONTYPE =
								append(_inferedTypeInstance.SPECIFICATIONTYPE, any(fieldInstance).(*SPECIFICATIONTYPE))
						}
					}
				}
			}
		}

	case *TARGET:
		// insertion point per field

	case *TARGETSPECIFICATION:
		// insertion point per field

	case *THEHEADER:
		// insertion point per field

	case *TOOLEXTENSIONS:
		// insertion point per field
		if fieldName == "REQIFTOOLEXTENSION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TOOLEXTENSIONS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TOOLEXTENSIONS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.REQIFTOOLEXTENSION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.REQIFTOOLEXTENSION = _inferedTypeInstance.REQIFTOOLEXTENSION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.REQIFTOOLEXTENSION =
								append(_inferedTypeInstance.REQIFTOOLEXTENSION, any(fieldInstance).(*REQIFTOOLEXTENSION))
						}
					}
				}
			}
		}

	case *VALUES:
		// insertion point per field

	case *XHTMLCONTENT:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ALTERNATIVEID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONBOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONDATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONINTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONREAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONSTRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEDEFINITIONXHTML
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEBOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEDATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEINTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEREAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUESTRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTEVALUEXHTML
	// insertion point per field

	// Compute reverse map for named struct CHILDREN
	// insertion point per field
	clear(stage.CHILDREN_SPECHIERARCHY_reverseMap)
	stage.CHILDREN_SPECHIERARCHY_reverseMap = make(map[*SPECHIERARCHY]*CHILDREN)
	for children := range stage.CHILDRENs {
		_ = children
		for _, _spechierarchy := range children.SPECHIERARCHY {
			stage.CHILDREN_SPECHIERARCHY_reverseMap[_spechierarchy] = children
		}
	}

	// Compute reverse map for named struct CORECONTENT
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONBOOLEAN
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONDATE
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONENUMERATION
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONINTEGER
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONREAL
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONSTRING
	// insertion point per field

	// Compute reverse map for named struct DATATYPEDEFINITIONXHTML
	// insertion point per field

	// Compute reverse map for named struct DATATYPES
	// insertion point per field
	clear(stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap = make(map[*DATATYPEDEFINITIONBOOLEAN]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitionboolean := range datatypes.DATATYPEDEFINITIONBOOLEAN {
			stage.DATATYPES_DATATYPEDEFINITIONBOOLEAN_reverseMap[_datatypedefinitionboolean] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap = make(map[*DATATYPEDEFINITIONDATE]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitiondate := range datatypes.DATATYPEDEFINITIONDATE {
			stage.DATATYPES_DATATYPEDEFINITIONDATE_reverseMap[_datatypedefinitiondate] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap = make(map[*DATATYPEDEFINITIONENUMERATION]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitionenumeration := range datatypes.DATATYPEDEFINITIONENUMERATION {
			stage.DATATYPES_DATATYPEDEFINITIONENUMERATION_reverseMap[_datatypedefinitionenumeration] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap = make(map[*DATATYPEDEFINITIONINTEGER]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitioninteger := range datatypes.DATATYPEDEFINITIONINTEGER {
			stage.DATATYPES_DATATYPEDEFINITIONINTEGER_reverseMap[_datatypedefinitioninteger] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap = make(map[*DATATYPEDEFINITIONREAL]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitionreal := range datatypes.DATATYPEDEFINITIONREAL {
			stage.DATATYPES_DATATYPEDEFINITIONREAL_reverseMap[_datatypedefinitionreal] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap = make(map[*DATATYPEDEFINITIONSTRING]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitionstring := range datatypes.DATATYPEDEFINITIONSTRING {
			stage.DATATYPES_DATATYPEDEFINITIONSTRING_reverseMap[_datatypedefinitionstring] = datatypes
		}
	}
	clear(stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap)
	stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap = make(map[*DATATYPEDEFINITIONXHTML]*DATATYPES)
	for datatypes := range stage.DATATYPESs {
		_ = datatypes
		for _, _datatypedefinitionxhtml := range datatypes.DATATYPEDEFINITIONXHTML {
			stage.DATATYPES_DATATYPEDEFINITIONXHTML_reverseMap[_datatypedefinitionxhtml] = datatypes
		}
	}

	// Compute reverse map for named struct DEFAULTVALUE
	// insertion point per field

	// Compute reverse map for named struct DEFINITION
	// insertion point per field

	// Compute reverse map for named struct EDITABLEATTS
	// insertion point per field

	// Compute reverse map for named struct EMBEDDEDVALUE
	// insertion point per field

	// Compute reverse map for named struct ENUMVALUE
	// insertion point per field

	// Compute reverse map for named struct OBJECT
	// insertion point per field

	// Compute reverse map for named struct PROPERTIES
	// insertion point per field

	// Compute reverse map for named struct RELATIONGROUP
	// insertion point per field

	// Compute reverse map for named struct RELATIONGROUPTYPE
	// insertion point per field

	// Compute reverse map for named struct REQIF
	// insertion point per field

	// Compute reverse map for named struct REQIFCONTENT
	// insertion point per field

	// Compute reverse map for named struct REQIFHEADER
	// insertion point per field

	// Compute reverse map for named struct REQIFTOOLEXTENSION
	// insertion point per field

	// Compute reverse map for named struct REQIFTYPE
	// insertion point per field

	// Compute reverse map for named struct SOURCE
	// insertion point per field

	// Compute reverse map for named struct SOURCESPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECATTRIBUTES
	// insertion point per field
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap = make(map[*ATTRIBUTEDEFINITIONBOOLEAN]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitionboolean := range specattributes.ATTRIBUTEDEFINITIONBOOLEAN {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONBOOLEAN_reverseMap[_attributedefinitionboolean] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap = make(map[*ATTRIBUTEDEFINITIONDATE]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitiondate := range specattributes.ATTRIBUTEDEFINITIONDATE {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONDATE_reverseMap[_attributedefinitiondate] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap = make(map[*ATTRIBUTEDEFINITIONENUMERATION]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitionenumeration := range specattributes.ATTRIBUTEDEFINITIONENUMERATION {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONENUMERATION_reverseMap[_attributedefinitionenumeration] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap = make(map[*ATTRIBUTEDEFINITIONINTEGER]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitioninteger := range specattributes.ATTRIBUTEDEFINITIONINTEGER {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONINTEGER_reverseMap[_attributedefinitioninteger] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap = make(map[*ATTRIBUTEDEFINITIONREAL]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitionreal := range specattributes.ATTRIBUTEDEFINITIONREAL {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONREAL_reverseMap[_attributedefinitionreal] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap = make(map[*ATTRIBUTEDEFINITIONSTRING]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitionstring := range specattributes.ATTRIBUTEDEFINITIONSTRING {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONSTRING_reverseMap[_attributedefinitionstring] = specattributes
		}
	}
	clear(stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap)
	stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap = make(map[*ATTRIBUTEDEFINITIONXHTML]*SPECATTRIBUTES)
	for specattributes := range stage.SPECATTRIBUTESs {
		_ = specattributes
		for _, _attributedefinitionxhtml := range specattributes.ATTRIBUTEDEFINITIONXHTML {
			stage.SPECATTRIBUTES_ATTRIBUTEDEFINITIONXHTML_reverseMap[_attributedefinitionxhtml] = specattributes
		}
	}

	// Compute reverse map for named struct SPECHIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATIONS
	// insertion point per field
	clear(stage.SPECIFICATIONS_SPECIFICATION_reverseMap)
	stage.SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*SPECIFICATIONS)
	for specifications := range stage.SPECIFICATIONSs {
		_ = specifications
		for _, _specification := range specifications.SPECIFICATION {
			stage.SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = specifications
		}
	}

	// Compute reverse map for named struct SPECIFICATIONTYPE
	// insertion point per field

	// Compute reverse map for named struct SPECIFIEDVALUES
	// insertion point per field
	clear(stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap)
	stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap = make(map[*ENUMVALUE]*SPECIFIEDVALUES)
	for specifiedvalues := range stage.SPECIFIEDVALUESs {
		_ = specifiedvalues
		for _, _enumvalue := range specifiedvalues.ENUMVALUE {
			stage.SPECIFIEDVALUES_ENUMVALUE_reverseMap[_enumvalue] = specifiedvalues
		}
	}

	// Compute reverse map for named struct SPECOBJECT
	// insertion point per field

	// Compute reverse map for named struct SPECOBJECTS
	// insertion point per field
	clear(stage.SPECOBJECTS_SPECOBJECT_reverseMap)
	stage.SPECOBJECTS_SPECOBJECT_reverseMap = make(map[*SPECOBJECT]*SPECOBJECTS)
	for specobjects := range stage.SPECOBJECTSs {
		_ = specobjects
		for _, _specobject := range specobjects.SPECOBJECT {
			stage.SPECOBJECTS_SPECOBJECT_reverseMap[_specobject] = specobjects
		}
	}

	// Compute reverse map for named struct SPECOBJECTTYPE
	// insertion point per field

	// Compute reverse map for named struct SPECRELATION
	// insertion point per field

	// Compute reverse map for named struct SPECRELATIONGROUPS
	// insertion point per field
	clear(stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap)
	stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap = make(map[*RELATIONGROUP]*SPECRELATIONGROUPS)
	for specrelationgroups := range stage.SPECRELATIONGROUPSs {
		_ = specrelationgroups
		for _, _relationgroup := range specrelationgroups.RELATIONGROUP {
			stage.SPECRELATIONGROUPS_RELATIONGROUP_reverseMap[_relationgroup] = specrelationgroups
		}
	}

	// Compute reverse map for named struct SPECRELATIONS
	// insertion point per field

	// Compute reverse map for named struct SPECRELATIONTYPE
	// insertion point per field

	// Compute reverse map for named struct SPECTYPES
	// insertion point per field
	clear(stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap)
	stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap = make(map[*RELATIONGROUPTYPE]*SPECTYPES)
	for spectypes := range stage.SPECTYPESs {
		_ = spectypes
		for _, _relationgrouptype := range spectypes.RELATIONGROUPTYPE {
			stage.SPECTYPES_RELATIONGROUPTYPE_reverseMap[_relationgrouptype] = spectypes
		}
	}
	clear(stage.SPECTYPES_SPECOBJECTTYPE_reverseMap)
	stage.SPECTYPES_SPECOBJECTTYPE_reverseMap = make(map[*SPECOBJECTTYPE]*SPECTYPES)
	for spectypes := range stage.SPECTYPESs {
		_ = spectypes
		for _, _specobjecttype := range spectypes.SPECOBJECTTYPE {
			stage.SPECTYPES_SPECOBJECTTYPE_reverseMap[_specobjecttype] = spectypes
		}
	}
	clear(stage.SPECTYPES_SPECRELATIONTYPE_reverseMap)
	stage.SPECTYPES_SPECRELATIONTYPE_reverseMap = make(map[*SPECRELATIONTYPE]*SPECTYPES)
	for spectypes := range stage.SPECTYPESs {
		_ = spectypes
		for _, _specrelationtype := range spectypes.SPECRELATIONTYPE {
			stage.SPECTYPES_SPECRELATIONTYPE_reverseMap[_specrelationtype] = spectypes
		}
	}
	clear(stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap)
	stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap = make(map[*SPECIFICATIONTYPE]*SPECTYPES)
	for spectypes := range stage.SPECTYPESs {
		_ = spectypes
		for _, _specificationtype := range spectypes.SPECIFICATIONTYPE {
			stage.SPECTYPES_SPECIFICATIONTYPE_reverseMap[_specificationtype] = spectypes
		}
	}

	// Compute reverse map for named struct TARGET
	// insertion point per field

	// Compute reverse map for named struct TARGETSPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct THEHEADER
	// insertion point per field

	// Compute reverse map for named struct TOOLEXTENSIONS
	// insertion point per field
	clear(stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap)
	stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap = make(map[*REQIFTOOLEXTENSION]*TOOLEXTENSIONS)
	for toolextensions := range stage.TOOLEXTENSIONSs {
		_ = toolextensions
		for _, _reqiftoolextension := range toolextensions.REQIFTOOLEXTENSION {
			stage.TOOLEXTENSIONS_REQIFTOOLEXTENSION_reverseMap[_reqiftoolextension] = toolextensions
		}
	}

	// Compute reverse map for named struct VALUES
	// insertion point per field

	// Compute reverse map for named struct XHTMLCONTENT
	// insertion point per field

}
