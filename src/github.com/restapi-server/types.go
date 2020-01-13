package main

// Json will only converted the filed which has a capital letter starting
type Book struct {
	Title       string `json:"title"`
	Page        int    `json:"page"`
	description string `json:"description"`   // never exported in json
	Tag         string `json:"tag,omitempty"` // if the value is 0/"" will be ignore, that is empty omitted.
	Unknown     string `json:"-"`             // will never be converted
}
