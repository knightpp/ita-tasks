package tasks_test

import (
	"reflect"
	"testing"

	tasks "github.com/knightpp"
)

func BenchmarkTask87(t *testing.B) {
	for i := 0; i < 100000; i++ {
		tasks.Task87(123456789, 5)
	}
}
func BenchmarkTask87Bad(t *testing.B) {
	for i := 0; i < 100000; i++ {
		tasks.Task87Bad(123456789, 5)
	}
}
func TestPrimeSieve(t *testing.T) {
	cases := []struct {
		MaxN   uint
		Answer []uint
	}{
		{1000, []uint{2, 3, 5, 7, 11, 13, 17,
			19, 23, 29, 31, 37, 41, 43, 47, 53,
			59, 61, 67, 71, 73, 79, 83, 89, 97,
			101, 103, 107, 109, 113, 127, 131,
			137, 139, 149, 151, 157, 163, 167,
			173, 179, 181, 191, 193, 197, 199,
			211, 223, 227, 229, 233, 239, 241,
			251, 257, 263, 269, 271, 277, 281,
			283, 293, 307, 311, 313, 317, 331,
			337, 347, 349, 353, 359, 367, 373,
			379, 383, 389, 397, 401, 409, 419,
			421, 431, 433, 439, 443, 449, 457,
			461, 463, 467, 479, 487, 491, 499,
			503, 509, 521, 523, 541, 547, 557,
			563, 569, 571, 577, 587, 593, 599,
			601, 607, 613, 617, 619, 631, 641,
			643, 647, 653, 659, 661, 673, 677,
			683, 691, 701, 709, 719, 727, 733,
			739, 743, 751, 757, 761, 769, 773,
			787, 797, 809, 811, 821, 823, 827,
			829, 839, 853, 857, 859, 863, 877,
			881, 883, 887, 907, 911, 919, 929,
			937, 941, 947, 953, 967, 971, 977,
			983, 991, 997}},
		{0, []uint{}},
		{1, []uint{}},
		{2, []uint{2}},
	}
	for _, c := range cases {
		primes := tasks.PrimeSieve(c.MaxN)
		if !reflect.DeepEqual(primes, c.Answer) {
			t.Errorf("maxN = %v\nresult = %v\nexpected = %v",
				c.MaxN, primes, c.Answer)
		}
	}
}
func TestTask87(t *testing.T) {
	cases := []struct{ N, M, Right uint }{
		{
			123, 1, 3,
		},
		{
			1234, 15, 1 + 2 + 3 + 4,
		},
		{
			1, 1, 1,
		},
		{
			123456789, 9, 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9,
		},
		{
			123456789, 8, 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9,
		},
		{
			/// ???
			9999, 0, 0,
		},
	}
	for _, c := range cases {
		eval := tasks.Task87(c.N, c.M)
		if eval != c.Right {
			t.Errorf("N = %v, M = %v, Result = %v but expected %v",
				c.N, c.M, eval, c.Right)
		}
	}
}

func TestTask226(t *testing.T) {
	cases := []struct {
		N, M  uint
		Right []uint
	}{
		{
			2, 4, []uint{1, 2},
		},
		{
			4, 8, []uint{1, 2, 4},
		},
		{
			1, 999, []uint{1},
		},
		{
			0, 0, []uint{},
		},
		{
			3 * 3, 3 * 3 * 3, []uint{1, 3, 9},
		},
		{
			9 * 9 * 9, 1000, []uint{1},
		},
	}
	for _, c := range cases {
		eval := tasks.Task226(c.M, c.N)
		if !reflect.DeepEqual(c.Right, eval) {
			t.Errorf("N = %v, M = %v, Result = %v but expected %v",
				c.N, c.M, eval, c.Right)
		}
	}
}

func TestTask559(t *testing.T) {
	cases := []struct {
		N      uint
		Answer []uint
	}{
		{1, []uint{}},
		{3, []uint{3}},
		{511, []uint{3, 7, 31, 127}},
		{131_071, []uint{3, 7, 31, 127, 8191, 131_071}},
	}
	for _, c := range cases {
		mer := tasks.Task559(c.N)
		if !reflect.DeepEqual(mer, c.Answer) {
			t.Errorf("N = %v, result = %v but expected %v",
				c.N, mer, c.Answer)
		}
	}
}
