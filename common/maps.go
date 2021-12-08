package common

import "encoding/json"

func ConvertStructToMap(model interface{}) (m map[string]interface{}, err error) {
	data, err := json.Marshal(model)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &m)
	return
}