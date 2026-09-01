package data

import "time"

// - is used to now shown the field in json output

// omitzero is omit the field when the value of the field is zero value
// Go type			Zero value
// bool				false
// string			""
// int*, uint*, float*, rune	0
// slices, maps, pointers	nil
// structs			each field in the struct has its zero value
// arrays			each element in the array is set to the zero value for its type

// omitempty is when your array is [], it will omit the whole field

// string just convert to string in json output
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`
	Runtime   int32     `json:"runtime"`
	Genres    []string  `json:"genres,omitzero"`
	Version   int32     `json:"version,string"`
}
