package json

import (
	"encoding/json"
)

//ToString ToString
func ToString(data interface{}) (string, error) {
	if data == nil {
		return "", nil
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//FromString FromString
func FromString(str string, data interface{}) error {
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return err
	}
	return nil
}

//ToMap ToMap
func ToMap(data interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var mapObj map[string]interface{}
	err = json.Unmarshal(bytes, &mapObj)
	if err != nil {
		return nil, err
	}
	return mapObj, nil
}

//Convert Convert
func Convert(s interface{}, d interface{}) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &d)
	if err != nil {
		return err
	}
	return nil
}
