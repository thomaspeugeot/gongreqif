// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CONTENT:
		ok = stage.IsStagedCONTENT(target)

	case *HEADER:
		ok = stage.IsStagedHEADER(target)

	case *REQIF:
		ok = stage.IsStagedREQIF(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedCONTENT(content *CONTENT) (ok bool) {

	_, ok = stage.CONTENTs[content]

	return
}

func (stage *StageStruct) IsStagedHEADER(header *HEADER) (ok bool) {

	_, ok = stage.HEADERs[header]

	return
}

func (stage *StageStruct) IsStagedREQIF(reqif *REQIF) (ok bool) {

	_, ok = stage.REQIFs[reqif]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *CONTENT:
		stage.StageBranchCONTENT(target)

	case *HEADER:
		stage.StageBranchHEADER(target)

	case *REQIF:
		stage.StageBranchREQIF(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchCONTENT(content *CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHEADER(header *HEADER) {

	// check if instance is already staged
	if IsStaged(stage, header) {
		return
	}

	header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if IsStaged(stage, reqif) {
		return
	}

	reqif.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.HEADER != nil {
		StageBranch(stage, reqif.HEADER)
	}
	if reqif.CONTENT != nil {
		StageBranch(stage, reqif.CONTENT)
	}

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
	case *CONTENT:
		toT := CopyBranchCONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *HEADER:
		toT := CopyBranchHEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQIF:
		toT := CopyBranchREQIF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCONTENT(mapOrigCopy map[any]any, contentFrom *CONTENT) (contentTo *CONTENT) {

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*CONTENT)
		return
	}

	contentTo = new(CONTENT)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.CopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHEADER(mapOrigCopy map[any]any, headerFrom *HEADER) (headerTo *HEADER) {

	// headerFrom has already been copied
	if _headerTo, ok := mapOrigCopy[headerFrom]; ok {
		headerTo = _headerTo.(*HEADER)
		return
	}

	headerTo = new(HEADER)
	mapOrigCopy[headerFrom] = headerTo
	headerFrom.CopyBasicFields(headerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

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
	if reqifFrom.HEADER != nil {
		reqifTo.HEADER = CopyBranchHEADER(mapOrigCopy, reqifFrom.HEADER)
	}
	if reqifFrom.CONTENT != nil {
		reqifTo.CONTENT = CopyBranchCONTENT(mapOrigCopy, reqifFrom.CONTENT)
	}

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
	case *CONTENT:
		stage.UnstageBranchCONTENT(target)

	case *HEADER:
		stage.UnstageBranchHEADER(target)

	case *REQIF:
		stage.UnstageBranchREQIF(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchCONTENT(content *CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHEADER(header *HEADER) {

	// check if instance is already staged
	if !IsStaged(stage, header) {
		return
	}

	header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQIF(reqif *REQIF) {

	// check if instance is already staged
	if !IsStaged(stage, reqif) {
		return
	}

	reqif.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if reqif.HEADER != nil {
		UnstageBranch(stage, reqif.HEADER)
	}
	if reqif.CONTENT != nil {
		UnstageBranch(stage, reqif.CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
