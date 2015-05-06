package digits

import (
	"math"
	"testing"
	"fmt"
)

func check(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		t.Error("Expected ", expected, ", got ", got)
	}
}

func TestIhex(t *testing.T) {
	// check(t, "6C65E52CB4", string(*ihex(0.423429797567303, 10)))
}

func TestMain(t *testing.T) {
	id := 1000000
	s1 := series(1, id)
	check(t, "0.181039533801509", fmt.Sprintf("%.15f",s1))
	s2 := series(4, id)
	check(t, "0.776065549807812", fmt.Sprintf("%.15f",s2))
	s3 := series(5, id)
	check(t, "0.362458564070547", fmt.Sprintf("%.15f",s3))
	s4 := series(6, id)
	check(t, "0.386138673951969", fmt.Sprintf("%.15f",s4))
	pid := 4*s1 - 2*s2 - s3 - s4
	pid = pid - math.Floor(pid)
	check(t, "0.423429797567895", fmt.Sprintf("%.15f",pid))
	chx := ihex(pid, 16)
	check(t, "6C65E52CB4", fmt.Sprintf("%10.10s",string(*chx)))
}
