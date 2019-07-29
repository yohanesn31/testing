package getaverage

import "testing"

func TestGetAverage(t *testing.T) {
	var v float64
	v = GetAverage(2, 2)
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
