package colors

type ColorStimulus struct {
	X, Y, Z float64
}

type ChromaticAdaptation struct {
	XR, XG, XB float64
	YR, YG, YB float64
	ZR, ZG, ZB float64
}

type LightProfile struct {
	ColorStimulus       ColorStimulus
	ChromaticAdaptation ChromaticAdaptation
}

func D65() LightProfile {
	return LightProfile{
		ColorStimulus: ColorStimulus{
			X: 95.047,
			Y: 100.0,
			Z: 108.883,
		},
		ChromaticAdaptation: ChromaticAdaptation{
			XR: 0.4124564, XG: 0.3575761, XB: 0.1804375,
			YR: 0.2126729, YG: 0.7151522, YB: 0.0721750,
			ZR: 0.0193339, ZG: 0.1191920, ZB: 0.9503041,
		},
	}
}

func D50() LightProfile {
	return LightProfile{
		ColorStimulus: ColorStimulus{
			X: 96.4212,
			Y: 100.0,
			Z: 82.5188,
		},
		ChromaticAdaptation: ChromaticAdaptation{
			XR: 0.4360747, XG: 0.3850649, XB: 0.1430804,
			YR: 0.2225045, YG: 0.7168786, YB: 0.0606169,
			ZR: 0.0139322, ZG: 0.0971045, ZB: 0.7141733,
		},
	}
}
