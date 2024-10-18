package colors

import (
	"testing"
)

func TestRGB_PerceivedLRV(t *testing.T) {
	v := RGB{
		R: 246, G: 244, B: 241,
	}

	lrv := v.PerceivedLRVBeta()
	if lrv != 90 {
		t.Fail()
	}
}
