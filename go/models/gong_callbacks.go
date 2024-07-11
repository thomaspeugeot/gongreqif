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
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERCreateCallback != nil {
			stage.OnAfterREQIFHEADERCreateCallback.OnAfterCreate(stage, target)
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
	case *REQIFHEADER:
		newTarget := any(new).(*REQIFHEADER)
		if stage.OnAfterREQIFHEADERUpdateCallback != nil {
			stage.OnAfterREQIFHEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERDeleteCallback != nil {
			staged := any(staged).(*REQIFHEADER)
			stage.OnAfterREQIFHEADERDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *REQIFHEADER:
		if stage.OnAfterREQIFHEADERReadCallback != nil {
			stage.OnAfterREQIFHEADERReadCallback.OnAfterRead(stage, target)
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
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[REQIFHEADER])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFCreateCallback = any(callback).(OnAfterCreateInterface[REQIF])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERCreateCallback = any(callback).(OnAfterCreateInterface[REQIFHEADER])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIF])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[REQIFHEADER])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *REQIF:
		stage.OnAfterREQIFReadCallback = any(callback).(OnAfterReadInterface[REQIF])
	
	case *REQIFHEADER:
		stage.OnAfterREQIFHEADERReadCallback = any(callback).(OnAfterReadInterface[REQIFHEADER])
	
	}
}
