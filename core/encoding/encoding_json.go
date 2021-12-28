package encoding

import (
	"encoding/json"
)

func JsonEncode(data interface{}, pritty bool) ([]byte, error) {
	var jsonByte []byte
	var err error
	if pritty {
		jsonByte, err = json.MarshalIndent(data, "", "  ")
	} else {
		jsonByte, err = json.Marshal(data)
	}
	if err != nil {
		return []byte{}, err
	}
	return jsonByte, nil
}

func JsonDecode(data []byte, structure interface{}) (interface{}, error) {
	err := json.Unmarshal(data, structure)
	return structure, err
}
