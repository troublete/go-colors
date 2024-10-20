package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/troublete/go-colors/colors"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("nothing to do")
		os.Exit(0)
	}

	r, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		slog.Error("failed to parse R", "err", err)
		os.Exit(1)
	}

	g, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		slog.Error("failed to parse G", "err", err)
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		slog.Error("failed to parse B", "err", err)
		os.Exit(1)
	}

	color := colors.NewRGB(r, g, b)
	fmt.Println(color.ToHSL().String())
}
