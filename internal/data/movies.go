package data

import (
	"time"

	"github.com/Fulim13/movie-api/internal/validator"
)

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
	Runtime   Runtime   `json:"runtime"`
	Genres    []string  `json:"genres,omitzero"`
	Version   int32     `json:"version,string"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "Must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contains duplicate values")
}
