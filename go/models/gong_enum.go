// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GLOBALREF
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (globalref GLOBALREF) ToString() (res string) {

	// migration of former implementation of enum
	switch globalref {
	// insertion code per enum code
	}
	return
}

func (globalref *GLOBALREF) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (globalref *GLOBALREF) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (globalref *GLOBALREF) ToCodeString() (res string) {

	switch *globalref {
	// insertion code per enum code
	}
	return
}

func (globalref GLOBALREF) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (globalref GLOBALREF) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for LOCALREF
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (localref LOCALREF) ToString() (res string) {

	// migration of former implementation of enum
	switch localref {
	// insertion code per enum code
	}
	return
}

func (localref *LOCALREF) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (localref *LOCALREF) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (localref *LOCALREF) ToCodeString() (res string) {

	switch *localref {
	// insertion code per enum code
	}
	return
}

func (localref LOCALREF) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (localref LOCALREF) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | GLOBALREF | LOCALREF
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*GLOBALREF | *LOCALREF
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
