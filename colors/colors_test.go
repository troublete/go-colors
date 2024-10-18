package colors

import "testing"

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
	} {
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
		xyz XYZ
		lab LAB
	}{
		{
			NewRGB(120, 120, 12).ToXYZ(),
			LAB{
				L: 48.82584210219513,
				A: -12.077647436082472,
				B: 51.69162385420878,
			},
		},
	} {
		v := tc.xyz.ToLAB()
		if v.L != tc.lab.L {
			t.Errorf("L: expected %v, got %v", tc.lab.L, v.L)
		}
		if v.A != tc.lab.A {
			t.Errorf("A: expected %v, got %v", tc.lab.A, v.A)
		}
		if v.B != tc.lab.B {
			t.Errorf("B: expected %v, got %v", tc.lab.B, v.B)
		}
	}
}
