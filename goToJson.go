package golang_structs_parser

import "github.com/iancoleman/strcase"

// Parse golang struct from file string, return equivalent json object
func ParseGoToJson(goStruct string) string {
	fields, types := ParseGoStructFields(goStruct)
	jsonResult := "{\n"
	for i, field := range fields {
		jsonResult = jsonResult + "\"" + strcase.ToLowerCamel(field) + "\":"
		// Parse number or string
		if types[i] == "int" || types[i] == "float32" || types[i] == "float64" {
			jsonResult = jsonResult + "" + "1"
		} else {
			jsonResult = jsonResult + "\"" + field + "\""
		}
		// Check if it's last element
		if i < len(fields)-1 {
			jsonResult = jsonResult + ",\n"
		} else {
			jsonResult = jsonResult + "\n"
		}
	}
	jsonResult = jsonResult + "}"
	return jsonResult
}
