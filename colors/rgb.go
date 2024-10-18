package colors

import "math"

func (x RGB) PerceivedLRVBeta() float64 {
	return math.Sqrt(+math.Pow(0.587*x.G, 2) + math.Pow(0.114*x.B, 2))
}
