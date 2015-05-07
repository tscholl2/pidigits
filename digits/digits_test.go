package digits

import (
	"fmt"
	"math"
	"testing"
)

func check(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}
}

func TestGet(t *testing.T) {
	check(t, "243F6A8885", string(*Get(0, 10)))
	check(t, "43F6A8885A", string(*Get(1, 10)))
	check(t, "013", string(*Get(75, 3)))
}

func TestSeries(t *testing.T) {
	id := 1000000
	s1 := series(1, id)
	check(t, "0.181039533801509", fmt.Sprintf("%.15f", s1))
	s2 := series(4, id)
	check(t, "0.776065549807812", fmt.Sprintf("%.15f", s2))
	s3 := series(5, id)
	check(t, "0.362458564070547", fmt.Sprintf("%.15f", s3))
	s4 := series(6, id)
	check(t, "0.386138673951969", fmt.Sprintf("%.15f", s4))
	pid := 4*s1 - 2*s2 - s3 - s4
	pid = pid - math.Floor(pid)
	check(t, "0.423429797567895", fmt.Sprintf("%.15f", pid))
}

func TestIhex(t *testing.T) {
	check(t, "0", string(*ihex(0, 1)))
	check(t, "0000000", string(*ihex(0, 7)))
	check(t, "31B", string(*ihex(3.0/16+1.0/(16*16)+11.0/(16*16*16), 3)))
}
