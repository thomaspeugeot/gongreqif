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
	case *REQIF:
		// insertion point per field

	case *REQ_IF_CONTENT:
		// insertion point per field
		if fieldName == "SPEC_OBJECT_TYPES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_OBJECT_TYPES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_OBJECT_TYPES = _inferedTypeInstance.SPEC_OBJECT_TYPES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_OBJECT_TYPES =
								append(_inferedTypeInstance.SPEC_OBJECT_TYPES, any(fieldInstance).(*SPEC_OBJECT_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFICATIONS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATIONS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATIONS = _inferedTypeInstance.SPECIFICATIONS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATIONS =
								append(_inferedTypeInstance.SPECIFICATIONS, any(fieldInstance).(*SPECIFICATION))
						}
					}
				}
			}
		}

	case *REQ_IF_HEADER:
		// insertion point per field

	case *SPECIFICATION:
		// insertion point per field

	case *SPECIFICATION_TYPE:
		// insertion point per field

	case *SPEC_HIERARCHY:
		// insertion point per field

	case *SPEC_OBJECT_TYPE:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct REQIF
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field
	clear(stage.REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap = make(map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_object_type := range req_if_content.SPEC_OBJECT_TYPES {
			stage.REQ_IF_CONTENT_SPEC_OBJECT_TYPES_reverseMap[_spec_object_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap)
	stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap = make(map[*SPECIFICATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _specification := range req_if_content.SPECIFICATIONS {
			stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[_specification] = req_if_content
		}
	}

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

}
