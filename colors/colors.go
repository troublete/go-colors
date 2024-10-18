package colors

import "math"

type RGB struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
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
	// D65
	// http://www.brucelindbloom.com/index.html?Eqn_RGB_to_XYZ.html
	return XYZ{
		X: 0.4124564*linR + 0.3575761*linG + 0.1804375*linB,
		Y: 0.2126729*linR + 0.7151522*linG + 0.0721750*linB,
		Z: 0.0193339*linR + 0.1191920*linG + 0.9503041*linB,
	}
}

func NewRGB(r, g, b float64) RGB {
	return RGB{
		R: r / 255.0,
		G: g / 255.0,
		B: b / 255.0,
	}
}

type XYZ struct {
	X, Y, Z float64
}

func (xyz XYZ) ToLAB() LAB {
	// make values light relative to D65
	oX := (xyz.X / 95.047) * 100.0
	oY := (xyz.Y / 100.0) * 100.0
	oZ := (xyz.Z / 108.883) * 100.0

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
