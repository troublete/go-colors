# go-colors
> color calculations and conversions for every-day use

## Introduction

This provides some 'every-day' color calculation and conversions that allows easier color discovery, matching and
handling. 

* convert HEX to sRGB
* convert sRGB to HEX
* convert sRGB to XYZ
* convert sRGB to HSL
* convert XYZ to L*a*b*
* calculate distance (delta E) between two L*a*b* colors
* calculate WCAG contrast ratio between to sRGB colors

**Example**
```go
a := NewRGB(120, 120, 12).ToXYZ().ToLAB()
b := NewRGB(230, 123, 40).ToXYZ().ToLAB()

deltaE := a.DistanceTo(b)
```
## CLI

### match-colors

This CLI accepts two CSV files (a 'source' collection of available colors; a 'match' which contains the colors you try
to match) – with format: collection(string),name(string),lrv(float64),r(float64),g(float64),b(
float64) – which then are compared concurrently and outputs the top 10 closest colors in respect to the match CSV
file.

**Example**
```bash
$ go run ./cmd/match -match my-colors.csv -source colors_available.csv -no-header -filter *available_name_filter_regex*
```

### delta-e

This CLI accepts two HEX color codes and returns the delta E value of both colors.

### hex-to-rgb

This CLI accepts a HEX color code and returns a ready-to-use RGB snippet for CSS.

### rgb-to-hsl

This CLI accepts three parameters (R, G, B) and convert them to a ready-to-use HSL snippet for CSS.

### rgb-to-hex

This CLI accepts three parameters (R, G, B) and converts them to a ready-to-use HEX color code.


### wcag-contrast-check

This CLI accepts two HEX color codes and returns the WCAG contrast ratio + if it succeeds the required 7:1 ratio.