package golang_structs_parser

import (
	"fmt"
	"strconv"
)

func contains(s []string, v string) bool {
	for _, el := range s {
		if el == v {
			return true
		}
	}
	return false
}

// Return interface real type
func typeAssertion(field interface{}) string {
	// Try every golang type
	parsedField, ok := field.(float64)
	if ok {
		// Try to force int instead of float 64
		intLen := len(strconv.Itoa(int(parsedField)))
		floatLen := len(fmt.Sprintf("%v", parsedField))
		// Check that rounded number is equal to not rounded number, if it is, then it was int type
		if intLen == floatLen {
			return "int"
		} else {
			return "float64"
		}
	}
	_, ok = field.(bool)
	if ok {
		return "bool"
	}
	content, ok := field.(string)
	if ok {
		// Check if string is marked as interface{}
		if content == "interface{}" {
			return content
		}
		return "string"
	}
	arr, ok := field.([]interface{})
	if ok {
		if len(arr) > 0 {
			return "[]" + typeAssertion(arr[0])
		}
	}
	return ""
}
