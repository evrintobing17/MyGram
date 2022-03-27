package structsconverter

import (
	"encoding/json"
	"fmt"
)

func ToMap(fromStruct interface{}) (map[string]interface{}, error) {

	var inInterface map[string]interface{}

	inrec, err := json.Marshal(&fromStruct)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}

	return inInterface, nil
}

func ToMapString(fromStruct interface{}) (map[string]string, error) {
	mappedStruct, err := ToMap(fromStruct)
	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]string)

	for key, val := range mappedStruct {
		resultMap[key] = fmt.Sprintf("%v", val)
	}

	return resultMap, nil
}
