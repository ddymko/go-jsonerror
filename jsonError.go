package jsonError

import (
	"encoding/json"
)

// An ErrorJSON contains an array of Errors.
type ErrorJSON struct {
	Errors []ErrorComp
}

// An ErrorComp is a error structure that follows the json spec.
type ErrorComp struct {
	ID     string `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Source Source `json:"source,omitempty"`
}

// A Source is represents
type Source struct {
	Pointer string `json:"pointer,omitempty"`
}

// Returns the error in string format
func (e *ErrorJSON) Error() string {
	out, _ := json.Marshal(e)
	return string(out)
}

// ErrorByte returns the error in an []byte
func (e *ErrorJSON) ErrorByte() []byte {
	out, _ := json.Marshal(e)
	return out
}

// AddError allows adding fields to the error
func (e *ErrorJSON) AddError(comp ErrorComp) {
	e.Errors = append(e.Errors, comp)
}
