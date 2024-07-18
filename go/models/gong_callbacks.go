// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *REQIF:
		if stage.OnAfterREQIFCreateCallback != nil {
			stage.OnAfterREQIFCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTCreateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERCreateCallback != nil {
			stage.OnAfterREQ_IF_HEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONCreateCallback != nil {
			stage.OnAfterSPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYCreateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *REQIF:
		newTarget := any(new).(*REQIF)
		if stage.OnAfterREQIFUpdateCallback != nil {
			stage.OnAfterREQIFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_CONTENT:
		newTarget := any(new).(*REQ_IF_CONTENT)
		if stage.OnAfterREQ_IF_CONTENTUpdateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_HEADER:
		newTarget := any(new).(*REQ_IF_HEADER)
		if stage.OnAfterREQ_IF_HEADERUpdateCallback != nil {
			stage.OnAfterREQ_IF_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION:
		newTarget := any(new).(*SPECIFICATION)
		if stage.OnAfterSPECIFICATIONUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_HIERARCHY:
		newTarget := any(new).(*SPEC_HIERARCHY)
		if stage.OnAfterSPEC_HIERARCHYUpdateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *REQIF:
		if stage.OnAfterREQIFDeleteCallback != nil {
			staged := any(staged).(*REQIF)
			stage.OnAfterREQIFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_CONTENT)
			stage.OnAfterREQ_IF_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_HEADER)
			stage.OnAfterREQ_IF_HEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION)
			stage.OnAfterSPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYDeleteCallback != nil {
			staged := any(staged).(*SPEC_HIERARCHY)
			stage.OnAfterSPEC_HIERARCHYDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *REQIF:
		if stage.OnAfterREQIFReadCallback != nil {
			stage.OnAfterREQIFReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTReadCallback != nil {
			stage.OnAfterREQ_IF_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERReadCallback != nil {
			stage.OnAfterREQ_IF_HEADERReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONReadCallback != nil {
			stage.OnAfterSPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYReadCallback != nil {
			stage.OnAfterSPEC_HIERARCHYReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_HEADER])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_HIERARCHY])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFCreateCallback = any(callback).(OnAfterCreateInterface[REQIF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_HEADER])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_HIERARCHY])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_HEADER])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_HIERARCHY])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFReadCallback = any(callback).(OnAfterReadInterface[REQIF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_HEADER])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYReadCallback = any(callback).(OnAfterReadInterface[SPEC_HIERARCHY])
	
	}
}
