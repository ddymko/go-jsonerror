package jsonerror

import (
	"reflect"
	"testing"
)

func TestErrorJSON_AddError(t *testing.T) {
	var err ErrorJSON

	errorComposition := ErrorComp{
		Detail: "this is a error message",
		Code:   "This is the code",
		Status: 200,
	}

	err.AddError(errorComposition)

	errorManually := ErrorJSON{Errors: []ErrorComp{{Detail: "this is a error message", Code: "This is the code", Status: 200}}}

	reflect.DeepEqual(err, errorManually)

	if !reflect.DeepEqual(err, errorManually) {
		t.Fatalf("%#v != %#v", err, errorManually)
	}
}

func TestErrorJSON_Error(t *testing.T) {
	var err ErrorJSON

	errorComposition := ErrorComp{
		Detail: "this is a error message",
		Code:   "This is the code",
		Source: Source{
			Pointer: "/unit/tests",
		},
		Title:  "Title Test",
		Status: 200,
	}

	err.AddError(errorComposition)

	errString := "{\"errors\":[{\"status\":200,\"code\":\"This is the code\",\"title\":\"Title Test\",\"detail\":\"this is a error message\",\"source\":{\"pointer\":\"/unit/tests\"}}]}"

	if errString != err.Error() {
		t.Fatal("The err string produced did not match up")
	}
}

func TestErrorJSON_ErrorByte(t *testing.T) {
	var err ErrorJSON

	errorComposition := ErrorComp{
		Detail: "this is a error message",
		Code:   "This is the code",
	}

	err.AddError(errorComposition)

	typ := reflect.TypeOf(err.ErrorByte()).Kind()

	if typ != reflect.Slice {
		t.Fatal("Error was not returned as []byte")
	}
}
