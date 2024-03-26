package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	path := "../input/5.json"
	if verify(path) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func verify(input string) bool {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error while opening:", err)
		return true
	}
	defer file.Close()
	var policy Policy
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&policy)
	if err != nil {
		fmt.Println("Error while decoding:", err)
		return true
	}
	var policyDoc PolicyDoc
	policyDoc = policy.PolicyDocument

	switch stmt := policyDoc.Statement.(type) {
	case map[string]interface{}:
		res := stmt["Resource"].(string)
		if res == "*" {
			return false
		}
	case []interface{}:
		for i := range stmt {
			res := stmt[i].(map[string]interface{})
			if resource, ok := res["Resource"].(string); ok && resource == "*" {
				return false
			}
		}
	default:
		fmt.Println("Error: Statement type not recognized")
		return true
	}
	return true
}
