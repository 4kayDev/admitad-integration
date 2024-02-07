package jsoner

import "encoding/json"

func Jsonify(obj interface{}) string {
	if obj == nil {
		return ""
	}
	str, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		return ""
	}

	return string(str)
}
