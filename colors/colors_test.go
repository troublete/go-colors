package colors

import (
	"testing"
)

func Test_RGBtoXYZ(t *testing.T) {
	for _, tc := range []struct {
		rgb RGB
		xyz XYZ
	}{
		{
			NewRGB(120, 120, 12),
			XYZ{
				X: 0.14529147863690453,
				Y: 0.17453017875806678,
				Z: 0.02951184150536878,
			},
		},
		{
			NewRGB(1, 1, 1),
			XYZ{
				X: 0.0002884932920536636,
				Y: 0.0003035270139015359,
				Z: 0.00033048928549748075,
			},
		},
	} {
		t.Run(tc.rgb.String(), func(t *testing.T) {
			v := tc.rgb.ToXYZ()
			if v.X != tc.xyz.X {
				t.Errorf("X: expected %v, got %v", tc.xyz.X, v.X)
			}
			if v.Y != tc.xyz.Y {
				t.Errorf("Y: expected %v, got %v", tc.xyz.Y, v.Y)
			}
			if v.Z != tc.xyz.Z {
				t.Errorf("Z: expected %v, got %v", tc.xyz.Z, v.Z)
			}
		})
	}
}

func Test_RGBtoXYZ_50(t *testing.T) {
	for _, tc := range []struct {
		rgb RGB
		xyz XYZ
	}{
		{
			NewRGB(120, 120, 12),
			XYZ{
				X: 0.15475310997719732,
				Y: 0.17665851780501599,
				Z: 0.023480662115807804,
			},
		},
	} {
		tc.rgb.Profile = D50()
		v := tc.rgb.ToXYZ()
		if v.X != tc.xyz.X {
			t.Errorf("X: expected %v, got %v", tc.xyz.X, v.X)
		}
		if v.Y != tc.xyz.Y {
			t.Errorf("Y: expected %v, got %v", tc.xyz.Y, v.Y)
		}
		if v.Z != tc.xyz.Z {
			t.Errorf("Z: expected %v, got %v", tc.xyz.Z, v.Z)
		}
	}
}

func Test_XYZtoLAB(t *testing.T) {
	for _, tc := range []struct {
		rgb RGB
		lab LAB
	}{
		{
			NewRGB(120, 120, 12),
			LAB{
				L: 48.82584210219513,
				A: -12.077647436082472,
				B: 51.69162385420878,
			},
		},
		{
			NewRGB(1, 1, 1),
			LAB{
				L: 0.27408489355308907,
				A: -0.00000011814002554011438,
				B: 0.00000004725601021604575,
			},
		},
	} {
		t.Run(tc.rgb.String(), func(t *testing.T) {
			v := tc.rgb.ToXYZ().ToLAB()
			if v.L != tc.lab.L {
				t.Errorf("L: expected %v, got %v", tc.lab.L, v.L)
			}
			if v.A != tc.lab.A {
				t.Errorf("A: expected %v, got %v", tc.lab.A, v.A)
			}
			if v.B != tc.lab.B {
				t.Errorf("B: expected %v, got %v", tc.lab.B, v.B)
			}
		})
	}
}

func Test_XYZtoLAB_50(t *testing.T) {
	for _, tc := range []struct {
		rgb RGB
		lab LAB
	}{
		{
			NewRGB(120, 120, 12),
			LAB{
				L: 49.08828821552924,
				A: -8.830453658783654,
				B: 51.162263417787514,
			},
		},
	} {
		t.Run(tc.rgb.String(), func(t *testing.T) {
			tc.rgb.Profile = D50()
			v := tc.rgb.ToXYZ().ToLAB()
			if v.L != tc.lab.L {
				t.Errorf("L: expected %v, got %v", tc.lab.L, v.L)
			}
			if v.A != tc.lab.A {
				t.Errorf("A: expected %v, got %v", tc.lab.A, v.A)
			}
			if v.B != tc.lab.B {
				t.Errorf("B: expected %v, got %v", tc.lab.B, v.B)
			}
		})
	}
}

func Test_LabDistanceToLab(t *testing.T) {
	a := NewRGB(72, 96, 85)
	b := NewRGB(82, 100, 91)

	want := 3.5026084248263922
	got := a.ToXYZ().ToLAB().DistanceTo(b.ToXYZ().ToLAB())
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func Test_NewRGBFromHex(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rgb, err := NewRGBFromHex("#4f45d3")
		if err != nil {
			t.Error(err)
		}

		if rgb.R != 0.30980392156862746 {
			t.Errorf("R: unexpected value '%v'", rgb.R)
		}
		if rgb.G != 0.27058823529411763 {
			t.Errorf("G: unexpected value '%v'", rgb.G)
		}
		if rgb.B != 0.8274509803921568 {
			t.Errorf("B: unexpected value '%v'", rgb.B)
		}
	})

	t.Run("success second", func(t *testing.T) {
		rgb, err := NewRGBFromHex("#234532")
		if err != nil {
			t.Error(err)
		}

		if rgb.R != 0.13725490196078433 {
			t.Errorf("R: unexpected value '%v'", rgb.R)
		}
		if rgb.G != 0.27058823529411763 {
			t.Errorf("G: unexpected value '%v'", rgb.G)
		}
		if rgb.B != 0.19607843137254902 {
			t.Errorf("B: unexpected value '%v'", rgb.B)
		}
	})

	t.Run("invalid hex", func(t *testing.T) {
		_, err := NewRGBFromHex("23df")
		if err == nil {
			t.Errorf("expected error if hex invalid")
		}
	})

	t.Run("invalid r", func(t *testing.T) {
		_, err := NewRGBFromHex("#zz9f9f")
		if err == nil {
			t.Errorf("expected error if r is invalid")
		}
	})

	t.Run("invalid g", func(t *testing.T) {
		_, err := NewRGBFromHex("#9fzz9f")
		if err == nil {
			t.Errorf("expected error if g is invalid")
		}
	})

	t.Run("invalid b", func(t *testing.T) {
		_, err := NewRGBFromHex("#9f9fzz")
		if err == nil {
			t.Errorf("expected error if b is invalid")
		}
	})
}

func Test_RGBToHSL(t *testing.T) {
	for _, tc := range []struct {
		rgb RGB
		hsl HSL
	}{
		{
			NewRGB(234, 45, 12),
			HSL{
				H: 8.91891891891892,
				S: 0.902439024390244,
				L: 0.48235294117647054,
			},
		},
		{
			NewRGB(45, 234, 12),
			HSL{
				H: 111.08108108108108,
				S: 0.902439024390244,
				L: 0.48235294117647054,
			},
		},
		{

			NewRGB(56, 45, 234),
			HSL{
				H: 243.49206349206352,
				S: 0.818181818181818,
				L: 0.5470588235294117,
			},
		},
		{
			NewRGB(100, 100, 100),
			HSL{
				H: 0,
				S: 0,
				L: 100.0 / 255.0,
			},
		},
	} {
		t.Run(tc.rgb.String(), func(t *testing.T) {
			a := tc.rgb.ToHSL()
			if a.H != tc.hsl.H {
				t.Errorf("H: expected %v, got %v", tc.hsl.H, a.H)
			}
			if a.S != tc.hsl.S {
				t.Errorf("S: expected %v, got %v", tc.hsl.S, a.S)
			}
			if a.L != tc.hsl.L {
				t.Errorf("L: expected %v, got %v", tc.hsl.L, a.L)
			}
		})
	}
}

func Test_RGBString(t *testing.T) {
	r := NewRGB(123, 12, 24)
	want := "rgb(123, 12, 24)"
	got := r.String()
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func Test_HSLString(t *testing.T) {
	r := NewRGB(123, 12, 24).ToHSL()
	want := "hsl(354 82% 26%)"
	got := r.String()
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func Test_RGBToHex(t *testing.T) {
	rgb, err := NewRGBFromHex("#2d3f4e")
	if err != nil {
		t.Error(err)
	}

	if rgb.ToHex() != "#2d3f4e" {
		t.Errorf("expected '#2d3f4e', got '%v'", rgb.ToHex())
	}
}
