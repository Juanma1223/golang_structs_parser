package golang_structs_parser

import (
	"encoding/json"
	"fmt"
	"log"
)

// Queue that holds every structure sto be parsed recursively
var structQueue []map[string]interface{}

func ParseJsonToGo(jsonString string) {
	fmt.Println(jsonString)
	var parsedJson map[string]any
	err := json.Unmarshal([]byte(jsonString), &parsedJson)
	if err != nil {
		log.Fatal(err)
	}
	structQueue = append(structQueue, parsedJson)
	parseStruct()
}

func parseStruct() {
	for len(structQueue) > 0 {
		newStruct := structQueue[0]
		// Pop parsed struct
		structQueue = structQueue[1:]
		for key, content := range newStruct {
			// Type assertion
			fmt.Println("parsing:", key)
			parsedContent, ok := content.(map[string]interface{})
			if ok {
				structQueue = append(structQueue, parsedContent)
			}
		}
	}
}
