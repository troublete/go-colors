package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/troublete/go-colors/colors"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("nothing to do")
		os.Exit(0)
	}

	from := args[0]
	color, err := colors.NewRGBFromHex(from)
	if err != nil {
		slog.Error("failed to parse", "input", from)
		os.Exit(1)
	}

	fmt.Println(color)
}
