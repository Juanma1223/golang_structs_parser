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
	_, ok = field.(string)
	if ok {
		return "string"
	}
	_, ok = field.([]interface{})
	if ok {
		return "[]"
	}
	_, ok = field.(map[string]interface{})
	if ok {
		return "struct"
	}
	return ""
}
