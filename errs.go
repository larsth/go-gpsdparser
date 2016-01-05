package gpsdparser

import "errors"

var (
	//ErrNilParseArgs is an error withe text "The ParseArgs is a nil pointer"
	ErrNilParseArgs = errors.New("The ParseArgs is a nil pointer")
	//ErrNilByteSlice is an error withe text "The byte slice is nil"
	ErrNilByteSlice = errors.New("The byte slice is nil")
	//ErrNilFilter is an error withe text "The *Filter is a nil pointer"
	ErrNilFilter = errors.New("The *Filter is a nil pointer")
	//ErrClassZeroLength is an error withe text "Class of type string is a zero length"
	ErrClassZeroLength = errors.New("Class of type string is a zero length")
	//ErrUnknownClass is an error withe text "Unknown gpsd JSON document class type"
	ErrUnknownClass = errors.New("Unknown gpsd JSON document class type")
	//ErrNilLogger is an error withe text "The Logger is a nil pointer"
	ErrNilLogger = errors.New("The Logger is a nil pointer")
)
