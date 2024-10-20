package main

import (
	"bytes"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/troublete/color-distance-cli/colors"
)

type ColorValue struct {
	Name string
	RGB  colors.RGB
}

type ColorList []ColorValue

func parseCSVFile(content []byte, filter *regexp.Regexp, withoutHeader bool) ColorList {
	file := bytes.NewBuffer(content)
	lines := strings.Split(file.String(), "\n")

	start := 1
	if withoutHeader {
		start = 0
	}
	var cl ColorList
	for _, l := range lines[start:] {
		if strings.TrimSpace(l) == "" {
			continue
		}

		cells := strings.Split(l, ",")
		name := strings.TrimSpace(strings.Join(cells[0:2], " "))

		if filter != nil && !filter.MatchString(name) {
			continue
		}

		_ = strings.TrimSpace(cells[2])
		r, err := strconv.ParseFloat(strings.TrimSpace(cells[3]), 64)
		if err != nil {
			slog.Error("failed to parse r", "err", err, "value", cells[3], "line", l)
			os.Exit(1)
		}
		g, err := strconv.ParseFloat(strings.TrimSpace(cells[4]), 64)
		if err != nil {
			slog.Error("failed to parse g", "err", err, "value", cells[4], "line", l)
			os.Exit(1)
		}
		b, err := strconv.ParseFloat(strings.TrimSpace(cells[5]), 64)
		if err != nil {
			slog.Error("failed to parse b", "err", err, "value", cells[5], "line", l)
			os.Exit(1)
		}

		cl = append(cl, ColorValue{
			Name: name,
			RGB:  colors.NewRGB(r, g, b),
		})
	}

	return cl
}

func main() {
	i := flag.String("source", "", "input .csv file path (csv format: collection,name,lrv,r,g,b)")
	m := flag.String("match", "", "match-with .csv file path")
	f := flag.String("filter", "", "regex pattern for name to filter for (only source)")
	h := flag.Bool("no-header", false, "to start 0 indexed; and don't skip 'header line'")
	flag.Parse()

	source, err := os.ReadFile(*i)
	if err != nil {
		slog.Error("failed to read source file", "err", err)
		os.Exit(1)
	}

	var filter *regexp.Regexp
	if strings.TrimSpace(*f) != "" {
		filter = regexp.MustCompile(*f)
	}

	sourceList := parseCSVFile(source, filter, *h)

	match, err := os.ReadFile(*m)
	if err != nil {
		slog.Error("failed to read match file", "err", err)
		os.Exit(1)
	}

	matchList := parseCSVFile(match, nil, *h)

	var wg sync.WaitGroup
	wg.Add(len(matchList))

	results := make(chan chan string, len(matchList))
	go func() {
		for {
			msg := <-results
			for c := range msg {
				for m := range c {
					fmt.Println(m)
				}
			}
		}
	}()

	for _, a := range matchList {
		go func(a ColorValue) {
			defer func() {
				wg.Done()
			}()

			buf := bytes.NewBufferString("")
			buf.WriteString(fmt.Sprintf("%v %v\n", strings.Repeat("=", 5), a.Name))
			var values []struct {
				name  string
				delta float64
			}

			for _, b := range sourceList {
				bLab := b.RGB.ToXYZ().ToLAB()
				delta := bLab.DistanceTo(a.RGB.ToXYZ().ToLAB())

				values = append(values, struct {
					name  string
					delta float64
				}{name: b.Name, delta: delta})
			}
			sort.Slice(values, func(a, b int) bool {
				return values[a].delta < values[b].delta
			})
			for _, v := range values[0:10] {
				buf.WriteString(fmt.Sprintf("ΔE %v – %v\n", v.delta, v.name))
			}

			fmt.Println(buf)
		}(a)
	}

	wg.Wait()
}
