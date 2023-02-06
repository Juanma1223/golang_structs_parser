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

var savedStructs map[string](map[string]interface{}) = make(map[string]map[string]interface{})

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
	addOrMergeStruct(structName, parsedJson)
	// Parse structures recursively
	parseStruct()
	// Parse structs to string
	for key, structure := range savedStructs {
		structToString(key, structure)
	}
}

func parseStruct() {
	for len(structQueue) > 0 {
		newStruct := structQueue[0]
		// Pop parsed struct
		structQueue = structQueue[1:]
		// Pop struct name
		// structName := structNames[0]
		structNames = structNames[1:]
		// Look for new struct to parse inside current struct
		for key, content := range newStruct {
			// Type assertion
			parsedContent, ok := content.(map[string]interface{})
			if ok {
				structNames = append(structNames, key)
				structQueue = append(structQueue, parsedContent)
				addOrMergeStruct(key, parsedContent)

			} else {
				// Check if field is an array
				checkArray, ok := content.([]interface{})
				if ok {
					for i := range checkArray {
						newStruct, ok = checkArray[i].(map[string]interface{})
						if ok {
							structQueue = append(structQueue, newStruct)
							structNames = append(structNames, key)
							addOrMergeStruct(key, newStruct)
						} else {
							fmt.Println("Couldn't parse struct", key)
						}
					}
				}
			}

		}
	}
}

// This function checks if a structure with the same name is already defined, then merge it's attributes
// If not, it adds the struct to the saved structures
func addOrMergeStruct(structName string, newStruct map[string]interface{}) {
	structure, ok := savedStructs[structName]
	// Struct found, then merge attributes
	if ok {
		for key, att := range newStruct {
			_, ok := structure[key]
			// This attribute is not defined in struct with the same name
			if !ok {
				structure[key] = att
			} else {
				// Check if types are equal, if not use interface{}
				if typeAssertion(att) != typeAssertion(structure[key]) {
					// Mark key as interface{}
					structure[key] = "interface{}"
				}
			}
		}
	} else {
		savedStructs[structName] = newStruct
	}
}

// Transform golang map[string]interface{} into a golang struct string
func structToString(structName string, currStruct map[string]interface{}) {
	fields := ""
	for key, content := range currStruct {
		// Type assertion
		dataType := typeAssertion(content)
		switch dataType {
		case "":
			// Datatype is struct
			dataType = strcase.ToCamel(key)
		case "[]":
			// Detected structure array
			dataType = dataType + strcase.ToCamel(key)
		default:
		}
		fields = fields + `	` + strcase.ToCamel(key) + " " + dataType + " `json:\"" + strcase.ToLowerCamel(key) + ",omitempty\"` \n"
	}
	parsedStruct := `type ` + strcase.ToCamel(structName) + " struct { \n" + fields + `}`
	structStrings = append(structStrings, parsedStruct)
	fmt.Println(parsedStruct)
}
