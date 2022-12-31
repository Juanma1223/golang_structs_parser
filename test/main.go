package main

import (
	"os"

	"github.com/Juanma1223/golang_structs_parser"
)

func main() {
	file, _ := os.ReadFile("../test_struct.json")

	golang_structs_parser.ParseJsonToGo(string(file))
}
