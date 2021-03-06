// see http://www.experimentalmath.info/bbp-codes/piqpr8.c

package digits

import (
	"math"
)

const (
	ntp = 25
)

var hx []rune
var tp []float64

func init() {
	hx = []rune("0123456789ABCDEF")
	tp = make([]float64, ntp)
	tp[0] = 1
	for i := 1; i < ntp; i++ {
		tp[i] = 2 * tp[i-1]
	}
}

// Get returns `length` hex digits of π
// starting at position `start`
// Get(0,3) returns the first 3 hex digits of pi
// that is, it takes 0.14159... and turns it into hex
func Get(start int, length int) *[]rune {
	chx := make([]rune, length)
	for pos := 0; pos < length; pos += 9 {
		next := getNext9(start + pos)
		for i := 0; i < int(math.Min(float64(length-pos), 9.0)); i++ {
			chx[pos+i] = (*next)[i]
		}
	}
	return &chx
}

// This returns 9 digits which seems to be how accurate
// the arithmetic is with float64
func getNext9(start int) *[]rune {
	s1 := series(1, start)
	s2 := series(4, start)
	s3 := series(5, start)
	s4 := series(6, start)
	pid := 4*s1 - 2*s2 - s3 - s4
	pid = pid - math.Floor(pid)
	return ihex(pid, 9)
}

// This returns, in chx, the first nhx hex digits of the fraction of x.
func ihex(x float64, nhx int) *[]rune {
	chx := make([]rune, nhx)

	y := math.Abs(x)

	for i := 0; i < nhx; i++ {
		y = 16 * (y - math.Floor(y))
		chx[i] = hx[int(math.Floor(y))]
	}

	return &chx
}

// This routine evaluates the series  sum_k 16^(id-k)/(8*k+m)
// using the modular exponentiation technique.
func series(m int, id int) float64 {
	fm := float64(m)
	fid := float64(id)

	var ak, p, t, k, s, eps float64
	eps = 0.00000000000000001 //1e-17
	s = 0

	/*  Sum the series up to id. */

	for k = 0; k < fid; k++ {
		ak = 8*k + fm
		p = fid - k
		t = expm(p, ak)
		s = s + float64(t)/float64(ak)
		s = s - math.Floor(s)
	}

	/*  Compute a few terms where k >= id. */

	for k = fid; k <= fid+100; k++ {
		ak = 8*k + fm
		t = math.Pow(16, fid-k) / ak
		if t < eps {
			break
		}
		s = s + t
		s = s - math.Floor(s)
	}
	return s
}

// expm = 16^p mod ak.  This routine uses the left-to-right binary
// exponentiation scheme.
func expm(p float64, ak float64) (r float64) {

	var i, j int
	var p1, pt float64

	if ak == 1 {
		return 0
	}

	/*  Find the greatest power of two less than or equal to p. */

	for i = 0; i < ntp; i++ {
		if tp[i] > p {
			break
		}
	}

	pt = tp[i-1]
	p1 = p
	r = 1.

	/*  Perform binary exponentiation algorithm modulo ak. */

	for j = 1; j <= i; j++ {
		if p1 >= pt {
			r = 16. * r
			r = r - math.Floor(r/ak)*ak
			p1 = p1 - pt
		}
		pt = 0.5 * pt
		if pt >= 1. {
			r = r * r
			r = r - math.Floor(r/ak)*ak
		}
	}

	return
}
