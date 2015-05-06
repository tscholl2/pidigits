// see http://www.experimentalmath.info/bbp-codes/piqpr8.c

package digits

import "math"

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

//  This returns, in chx, the first nhx hex digits of the fraction of x.
func ihex(x float64, nhx int) *[]rune {

	y := math.Abs(x)
	chx := make([]rune, nhx)

	for i := 0; i < nhx; i++ {
		y = 16 * (y - math.Floor(y))
		chx[i] = hx[int(y)]
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
