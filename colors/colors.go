package colors

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RGB struct {
	R, G, B float64 // must be [0..1]; v/255

	Profile LightProfile
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%.0f, %.0f, %.0f)", rgb.R*255, rgb.G*255, rgb.B*255)
}

func (rgb RGB) ToHex() string {
	return fmt.Sprintf("#%x%x%x", int(rgb.R*255), int(rgb.G*255), int(rgb.B*255))
}

func (rgb RGB) linearize() (r, g, b float64) {
	v := func(x float64) float64 {
		if x > 0.04045 {
			return math.Pow((x+0.055)/1.055, 2.4)
		} else {
			return x / 12.92
		}
	}

	return v(rgb.R), v(rgb.G), v(rgb.B)
}

func (rgb RGB) ToXYZ() XYZ {
	linR, linG, linB := rgb.linearize()

	return XYZ{
		X: rgb.Profile.ChromaticAdaptation.XR*linR + rgb.Profile.ChromaticAdaptation.XG*linG + rgb.Profile.ChromaticAdaptation.XB*linB,
		Y: rgb.Profile.ChromaticAdaptation.YR*linR + rgb.Profile.ChromaticAdaptation.YG*linG + rgb.Profile.ChromaticAdaptation.YB*linB,
		Z: rgb.Profile.ChromaticAdaptation.ZR*linR + rgb.Profile.ChromaticAdaptation.ZG*linG + rgb.Profile.ChromaticAdaptation.ZB*linB,

		rgb: rgb,
	}
}

func (rgb RGB) ToHSL() HSL {
	maxv := math.Max(math.Max(rgb.R, rgb.G), rgb.B)
	minv := math.Min(math.Min(rgb.R, rgb.G), rgb.B)

	h := 0.0
	s := 0.0
	l := (maxv + minv) / 2

	if maxv != minv {
		switch {
		case maxv == rgb.R:
			h = 60.0 * ((rgb.G - rgb.B) / (maxv - minv))
		case maxv == rgb.G:
			h = 60.0 * (2 + ((rgb.B - rgb.R) / (maxv - minv)))
		case maxv == rgb.B:
			h = 60.0 * (4 + ((rgb.R - rgb.G) / (maxv - minv)))
		}
		if h < 0 {
			h = h + 360.0
		}
		s = (maxv - minv) / (1 - math.Abs(maxv+minv-1))
	}

	return HSL{
		H: h,
		S: s,
		L: l,
	}
}

func NewRGB(r, g, b float64) RGB {
	return RGB{
		R: r / 255.0,
		G: g / 255.0,
		B: b / 255.0,

		Profile: D65(),
	}
}

func NewRGBFromHex(hex string) (*RGB, error) {
	if len(hex) < 7 {
		return nil, fmt.Errorf("invalid hex '%v'", hex)
	}

	h := bytes.NewBufferString(hex).Bytes()
	var r strings.Builder
	r.Write(h[1:3])
	var g strings.Builder
	g.Write(h[3:5])
	var b strings.Builder
	b.Write(h[5:7])

	rval, err := strconv.ParseInt(r.String(), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse red '%v'", r.String())
	}

	gval, err := strconv.ParseInt(g.String(), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse red '%v'", g.String())
	}

	bval, err := strconv.ParseInt(b.String(), 16, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse red '%v'", b.String())
	}

	col := NewRGB(float64(rval), float64(gval), float64(bval))
	return &col, nil
}

type HSL struct {
	H, S, L float64
}

func (hsl HSL) String() string {
	return fmt.Sprintf("hsl(%.0f %.0f%% %.0f%%)", hsl.H, hsl.S*100, hsl.L*100)
}

type XYZ struct {
	X, Y, Z float64

	rgb RGB
}

func (xyz XYZ) ToLAB() LAB {
	// make values light relative to D65
	oX := (xyz.X / xyz.rgb.Profile.ColorStimulus.X) * 100.0
	oY := (xyz.Y / xyz.rgb.Profile.ColorStimulus.Y) * 100.0
	oZ := (xyz.Z / xyz.rgb.Profile.ColorStimulus.Z) * 100.0

	f := func(x float64) float64 {
		if x < (216.0 / 24389.0) {
			return (7.787 * x) + (16 / 116)
			//return 1.0 / 116.0 * ((24389.0/27.0)*x + 16.0)
		} else {
			return math.Cbrt(x)
		}
	}

	fX := f(oX)
	fY := f(oY)
	fZ := f(oZ)

	lab := LAB{
		L: (116.0 * fY) - 16.0,
		A: 500.0 * (fX - fY),
		B: 200.0 * (fY - fZ),
	}
	return lab
}

type LAB struct {
	L, A, B float64
}

func (a LAB) DistanceTo(b LAB) float64 {
	return math.Sqrt(math.Pow(b.L-a.L, 2) + math.Pow(b.A-a.A, 2) + math.Pow(b.B-a.B, 2))
}
