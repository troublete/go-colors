package main

import (
	"encoding/json"
	"flag"
	"github.com/troublete/color-distance-cli/colors"
	"os"
)

type ColorValue struct {
	Name    string
	R, G, B float64
}

type ColorList []ColorValue

func main() {
	i := flag.String("input", "", "input file path")
	r := flag.String("required", "", "requirement file path")
	flag.Parse()

	input, err := os.ReadFile(*i)
	if err != nil {
		panic(err)
	}

	requirement, err := os.ReadFile(*r)
	if err != nil {
		panic(err)
	}

	var inputList colors.ColorList
	err = json.Unmarshal(input, &inputList)
	if err != nil {
		panic(err)
	}

	var requiredList colors.ColorList
	err = json.Unmarshal(requirement, &requiredList)
	if err != nil {
		panic(err)
	}
}
