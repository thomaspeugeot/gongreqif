// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *REQIF:
		ok = stage.IsStagedREQIF(target)

	case *REQIFHEADER:
		ok = stage.IsStagedREQIFHEADER(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedREQIF(reqif *REQIF) (ok bool) {

	_, ok = stage.REQIFs[reqif]

	return
}

func (stage *StageStruct) IsStagedREQIFHEADER(reqifheader *REQIFHEADER) (ok bool) {

	_, ok = stage.REQIFHEADERs[reqifheader]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *REQIF:
		stage.StageBranchREQIF(target)

	case *REQIFHEADER:
		stage.StageBranchREQIFHEADER(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if IsStaged(stage, reqif) {
		return
	}

	reqif.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.REQIFHEADER != nil {
		StageBranch(stage, reqif.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIFHEADER(reqifheader *REQIFHEADER) {

	// check if instance is already staged
	if IsStaged(stage, reqifheader) {
		return
	}

	reqifheader.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *REQIF:
		toT := CopyBranchREQIF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIFHEADER:
		toT := CopyBranchREQIFHEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchREQIF(mapOrigCopy map[any]any, reqifFrom *REQIF) (reqifTo *REQIF) {

	// reqifFrom has already been copied
	if _reqifTo, ok := mapOrigCopy[reqifFrom]; ok {
		reqifTo = _reqifTo.(*REQIF)
		return
	}

	reqifTo = new(REQIF)
	mapOrigCopy[reqifFrom] = reqifTo
	reqifFrom.CopyBasicFields(reqifTo)

	//insertion point for the staging of instances referenced by pointers
	if reqifFrom.REQIFHEADER != nil {
		reqifTo.REQIFHEADER = CopyBranchREQIFHEADER(mapOrigCopy, reqifFrom.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQIFHEADER(mapOrigCopy map[any]any, reqifheaderFrom *REQIFHEADER) (reqifheaderTo *REQIFHEADER) {

	// reqifheaderFrom has already been copied
	if _reqifheaderTo, ok := mapOrigCopy[reqifheaderFrom]; ok {
		reqifheaderTo = _reqifheaderTo.(*REQIFHEADER)
		return
	}

	reqifheaderTo = new(REQIFHEADER)
	mapOrigCopy[reqifheaderFrom] = reqifheaderTo
	reqifheaderFrom.CopyBasicFields(reqifheaderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *REQIF:
		stage.UnstageBranchREQIF(target)

	case *REQIFHEADER:
		stage.UnstageBranchREQIFHEADER(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if !IsStaged(stage, reqif) {
		return
	}

	reqif.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.REQIFHEADER != nil {
		UnstageBranch(stage, reqif.REQIFHEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIFHEADER(reqifheader *REQIFHEADER) {

	// check if instance is already staged
	if !IsStaged(stage, reqifheader) {
		return
	}

	reqifheader.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
