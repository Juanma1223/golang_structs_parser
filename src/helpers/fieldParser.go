package helpers

import (
	"fmt"
	"log"
	"strings"
)

var currState = 0

// Parse golang struct from string, return field names and it's types
// Golang struct should be correctly indented
func ParseGoStructFields(goStruct string) ([]string, []string) {
	lines := strings.Split(goStruct, "\n")

	// Clean comments and line jumps
	lines = cleanLines(lines)

	// Skip package line
	lines = lines[1:]

	// Check valid struct syntax
	if !parseStructInit(lines[0]) {
		log.Fatal("Golang struct incorrect syntax at:", lines[0])
	}

	// Pop struct declaring line from slice
	lines = lines[1:]
	var fieldNames []string
	var fieldTypes []string
	for _, line := range lines {
		if line == "}" {
			break
		}
		columnName, dataType := parseColumn(line)
		fieldNames = append(fieldNames, columnName)
		fieldTypes = append(fieldTypes, dataType)
	}
	fmt.Println(fieldNames, fieldTypes)
	return fieldNames, fieldTypes
}

// Removes white spaces and line jumps
func cleanLines(lines []string) []string {
	cleanLines := []string{}
	for _, line := range lines {
		isComment := strings.HasPrefix(line, "//")
		if line != "" && !isComment {
			cleanLines = append(cleanLines, line)
		}
	}
	return cleanLines
}

// Remove white spaces and comments
func lineCleaner(line []string) []string {
	cleanLine := []string{}
	for _, word := range line {
		// Ignore white spaces and line jumps
		if word != "" {
			cleanLine = append(cleanLine, word)
		}
	}
	return cleanLine
}

// Parse golang struct beginning syntax
func parseStructInit(stringLine string) bool {
	line := strings.Split(stringLine, " ")
	line = lineCleaner(line)

	if len(line) < 4 {
		fmt.Println("Error: Bad syntax on struct initialization")
		return false
	}

	if line[0] == "type" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"type\", found:", line[0])
		return false
	}

	if line[2] == "struct" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"struct\", found:", line[2])
		return false
	}

	if line[3] == "{" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"}\", found:", line[3])
		return false
	}
	return true
}

func parseColumn(stringLine string) (string, string) {
	line := strings.Split(stringLine, " ")
	// Clean white spaces and comments
	line = lineCleaner(line)

	if len(line) < 3 {
		log.Fatal("Error: Column", stringLine, "syntax is incorrect")
	}
	columnName := line[0]
	// Remove tabs from name
	if columnName[:1] == "\t" {
		columnName = columnName[1:]
	}
	goDataType := line[1]

	return columnName, goDataType
}
