package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// Transform will transform Data using declartive rules
func Transform(original interface{}, rules map[string]interface{}) (map[string]interface{}, error) {

	jsonBytes, err := json.Marshal(original)
	json := string(jsonBytes)

	if err != nil {
		return nil, err
	}

	data := rules
	returnData := make(map[string]interface{})

	for k, v := range data {

		value := gjson.Get(json, k)

		switch typ := v.(type) {
		case string:

			newValues := strings.Split(typ, "|")

			if len(newValues) == 2 {
				switch newValues[1] {
				case "string":
					returnData[newValues[0]] = value.String()

				case "integer":
					returnData[newValues[0]] = value.Int()

				case "float":
					returnData[newValues[0]] = value.Float()

				case "bool":
					returnData[newValues[0]] = value.Bool()

				case "time":
					returnData[newValues[0]] = value.Time()

				case "uint":
					returnData[newValues[0]] = value.Uint()

				default:
					returnData[typ] = value.Value()
				}
			} else {
				returnData[typ] = value.Value()
			}

		case []interface{}:

			values := value.Value()
			newValues, ok := v.([]interface{})

			if !ok {
				return nil, fmt.Errorf("Wrong values provided for struct transformation")
			}

			key, ok := newValues[0].(string)

			if !ok {
				return nil, fmt.Errorf("Wrong key provided for struct transformation")
			}
			rules, ok := newValues[1].(map[string]interface{})

			if !ok {
				return nil, fmt.Errorf("Wrong Rules provided for struct transformation")
			}

			data, err := Transform(values, rules)

			if err != nil {
				return nil, err
			}

			returnData[key] = data

		case func(interface{}) interface{}:

			value := typ(original)
			returnData[k] = value
		}
	}

	return returnData, nil
}
