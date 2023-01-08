package golang_structs_parser

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/iancoleman/strcase"
)

// Queue that holds every structure sto be parsed recursively
var structQueue []map[string]interface{}
var structNames []string

// Parsed structures
var structStrings []string

func ParseJsonToGo(structName, jsonString string) {
	// fmt.Println(jsonString)
	var parsedJson map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &parsedJson)
	if err != nil {
		log.Fatal(err)
	}
	structQueue = append(structQueue, parsedJson)
	structNames = append(structNames, structName)
	parseStruct()
}

func parseStruct() {
	for len(structQueue) > 0 {
		newStruct := structQueue[0]
		// Pop parsed struct
		structQueue = structQueue[1:]
		// Pop struct name
		structName := structNames[0]
		structNames = structNames[1:]
		structToString(structName, newStruct)
		// Look for new struct to parse inside current struct
		for key, content := range newStruct {
			// Type assertion
			parsedContent, ok := content.(map[string]interface{})
			if ok {
				structNames = append(structNames, key)
				structQueue = append(structQueue, parsedContent)
			} else {
				// Check if field is an array
				checkArray, ok := content.([]interface{})
				if ok {
					// Get array's first value and parse it
					if len(checkArray) > 0 {
						newStruct, ok = checkArray[0].(map[string]interface{})
						if !ok {
							fmt.Println("Couldn't parse struct", key)
						}
						structQueue = append(structQueue, newStruct)
						structNames = append(structNames, key)
					}
				}
			}

		}
	}
}

// Transform golang map[string]interface{} into a golang struct string
func structToString(structName string, currStruct map[string]interface{}) {
	fields := ""
	for key, content := range currStruct {
		// Type assertion
		dataType := typeAssertion(content)
		if dataType == "" {
			// Datatype is struct
			dataType = strcase.ToCamel(key)
		} else if dataType == "[]" {
			// Detected structure array
			dataType = dataType + strcase.ToCamel(key)
		}
		fields = fields + `	` + strcase.ToCamel(key) + " " + dataType + " `json:\"" + strcase.ToLowerCamel(key) + "\",omitempty` \n"
	}
	parsedStruct := "struct { \n" + fields + `}`
	structStrings = append(structStrings, parsedStruct)
	fmt.Println(parsedStruct)
}
