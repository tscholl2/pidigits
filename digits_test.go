package digits

import (
	"math"
	"testing"
)

func check(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}
}

func TestIhex(t *testing.T) {
	check(t, "6C65E52CB4", string(*ihex(0.423429797567303, 10)))
}

func TestMain(t *testing.T) {
	id := 1000000
	s1 := series(1, id)
	s2 := series(4, id)
	s3 := series(5, id)
	s4 := series(6, id)
	pid := 4.*s1 - 2.*s2 - s3 - s4
	pid = pid - math.Floor(pid+1)
	chx := ihex(pid, 16)
	check(t, "6C65E52CB4", string(*chx))
	check(t, 0.423429797567303, pid)
}
