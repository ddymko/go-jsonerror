package jsonError

import (
	"encoding/json"
)

// An ErrorJson contains an array of Errors.
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

type Source struct {
	Pointer string `json:"pointer,omitempty"`
}

func (e *ErrorJSON) Error() string {
	out, _ := json.Marshal(e)
	return string(out)
}

func (e *ErrorJSON) ErrorByte() []byte {
	out, _ := json.Marshal(e)
	return out
}

func (e *ErrorJSON) AddError(comp ErrorComp) {
	e.Errors = append(e.Errors, comp)
}
