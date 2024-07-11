// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CONTENT:
		if stage.OnAfterCONTENTCreateCallback != nil {
			stage.OnAfterCONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *HEADER:
		if stage.OnAfterHEADERCreateCallback != nil {
			stage.OnAfterHEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQIF:
		if stage.OnAfterREQIFCreateCallback != nil {
			stage.OnAfterREQIFCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *CONTENT:
		newTarget := any(new).(*CONTENT)
		if stage.OnAfterCONTENTUpdateCallback != nil {
			stage.OnAfterCONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *HEADER:
		newTarget := any(new).(*HEADER)
		if stage.OnAfterHEADERUpdateCallback != nil {
			stage.OnAfterHEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQIF:
		newTarget := any(new).(*REQIF)
		if stage.OnAfterREQIFUpdateCallback != nil {
			stage.OnAfterREQIFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *CONTENT:
		if stage.OnAfterCONTENTDeleteCallback != nil {
			staged := any(staged).(*CONTENT)
			stage.OnAfterCONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *HEADER:
		if stage.OnAfterHEADERDeleteCallback != nil {
			staged := any(staged).(*HEADER)
			stage.OnAfterHEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQIF:
		if stage.OnAfterREQIFDeleteCallback != nil {
			staged := any(staged).(*REQIF)
			stage.OnAfterREQIFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CONTENT:
		if stage.OnAfterCONTENTReadCallback != nil {
			stage.OnAfterCONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *HEADER:
		if stage.OnAfterHEADERReadCallback != nil {
			stage.OnAfterHEADERReadCallback.OnAfterRead(stage, target)
		}
	case *REQIF:
		if stage.OnAfterREQIFReadCallback != nil {
			stage.OnAfterREQIFReadCallback.OnAfterRead(stage, target)
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
	case *CONTENT:
		stage.OnAfterCONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[CONTENT])
	
	case *HEADER:
		stage.OnAfterHEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[HEADER])
	
	case *REQIF:
		stage.OnAfterREQIFUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIF])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CONTENT:
		stage.OnAfterCONTENTCreateCallback = any(callback).(OnAfterCreateInterface[CONTENT])
	
	case *HEADER:
		stage.OnAfterHEADERCreateCallback = any(callback).(OnAfterCreateInterface[HEADER])
	
	case *REQIF:
		stage.OnAfterREQIFCreateCallback = any(callback).(OnAfterCreateInterface[REQIF])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CONTENT:
		stage.OnAfterCONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[CONTENT])
	
	case *HEADER:
		stage.OnAfterHEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[HEADER])
	
	case *REQIF:
		stage.OnAfterREQIFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIF])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CONTENT:
		stage.OnAfterCONTENTReadCallback = any(callback).(OnAfterReadInterface[CONTENT])
	
	case *HEADER:
		stage.OnAfterHEADERReadCallback = any(callback).(OnAfterReadInterface[HEADER])
	
	case *REQIF:
		stage.OnAfterREQIFReadCallback = any(callback).(OnAfterReadInterface[REQIF])
	
	}
}
