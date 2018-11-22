package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Address   *Address `json:"address"`
}

// Rules will return the rules
func (user *User) Rules() map[string]interface{} {
	data := map[string]interface{}{
		"first_name": "firstName",
		"last_name":  "lastName",
		"age":        "age|string",
		"fullName": func(data interface{}) interface{} {
			var user User

			bytes, _ := json.Marshal(data)
			json.Unmarshal(bytes, &user)

			return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
		},
		"address": []interface{}{
			"address",
			user.Address.Rules(),
		},
	}
	return data
}

type Address struct {
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
}

func (address *Address) Rules() map[string]interface{} {
	data := map[string]interface{}{
		"city":        "City",
		"postal_code": "Code",
	}
	return data
}

func TestSum(t *testing.T) {

	address := &Address{"Singapore", "770124", "Singapore"}
	user := &User{"Nyan", "Win", 24, address}
	rules := user.Rules()

	data, err := Transform(user, rules)

	if data == nil || err != nil {
		t.Errorf("Failed")
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Failed")
	}
	fmt.Println(string(bytes))
}
