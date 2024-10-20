package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/troublete/go-colors/colors"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("nothing to do")
		os.Exit(0)
	}

	a := args[0]
	colorA, err := colors.NewRGBFromHex(a)
	if err != nil {
		slog.Error("failed to parse", "a", a)
		os.Exit(1)
	}

	b := args[1]
	colorB, err := colors.NewRGBFromHex(b)
	if err != nil {
		slog.Error("failed to parse", "b", b)
		os.Exit(1)
	}

	contrast, good := colorA.WACGContrastRatioTo(*colorB)
	fmt.Println(contrast, good)
}
