package structsconverter

import (
	"testing"
)

func TestToMapString(t *testing.T) {
	type exampleStruct struct {
		Field1 int     `json:"field_1"`
		Field2 bool    `json:"field_2"`
		Field3 float64 `json:"field_3"`
		Field4 float32 `json:"field_4"`
		Field5 int64   `json:"field_5"`
	}

	exampleObj := exampleStruct{
		Field1: 1,
		Field2: false,
		Field3: 13.3,
		Field4: 12.2,
		Field5: 67,
	}

	_, err := ToMapString(exampleObj)
	if err != nil {
		t.Error(err)
	}

}
