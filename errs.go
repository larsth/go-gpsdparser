package gpsdparser

import "errors"

var (
	//ErrNilParseArgs is an error with the text "The ParseArgs is a nil pointer"
	ErrNilParseArgs = errors.New("The ParseArgs is a nil pointer")

	//ErrNilByteSlice is an error with the text
	//"The byte slice is nil, or the empty byte slice"
	ErrEmptyOrNilByteSlice = errors.New(
		"The byte slice is nil, or the empty byte slice")

	//ErrNilFilter is an error with the text "The *Filter is a nil pointer"
	ErrNilFilter = errors.New("The *Filter is a nil pointer")

	//ErrClassZeroLength is an error wit thhe text "Class of type string is a zero length"
	ErrClassZeroLength = errors.New("Class of type string is a zero length")

	//ErrUnknownClass is an error with the text "Unknown gpsd JSON document class type"
	ErrUnknownClass = errors.New("Unknown gpsd JSON document class type")

	//ErrNilRule is an error with the text ""
	ErrNilRule = errors.New("go-gpsdfilter.*Rule is 'nil'")
)
