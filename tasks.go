package tasks

import (
	"fmt"
)

// Arguments
//
// n - a natural number
//
// m - take this many digits from the back of `n`
//
//
// Special cases
//
//
// if `m` > `digits(n)` -> sum of all digits
//
// if `m` == 0 -> 0
//
//
// Returns
//
// Sum of the `m` last digits
func Task87(n, m uint) uint {
	max := pow(10, m)
	number := n % max
	var sum uint
	for i := number; i > 0; i = i / 10 {
		sum += i % 10
	}
	return sum
}
func Task87Bad(n, m uint) uint {
	str := fmt.Sprintf("%v", n)
	if m < uint(len(str)) {
		str = str[uint(len(str))-m:]
	}
	var sum uint = 0
	for i := range str {
		sum += uint(str[i] - '0')
	}
	return sum
}

// cSpell:disable
// Даны натуральные числа m, n. Получить
// их натуральные общие кратные, меньшие m*n.
// cSpell:enable
func Task226(m, n uint) []uint {
	mn := m * n
	var min uint
	if m < n {
		min = m
	} else {
		min = n
	}
	res := []uint{}
	for i := uint(1); i <= min; i++ {
		if n%i == 0 && m%i == 0 && i < mn {
			res = append(res, i)
		}
	}
	return res
}

// Returns
//
// Mersenne numbers <= n
func Task559(n uint) []uint {
	return task559Fn(n, PrimeSieve)
}
func task559Fn(n uint, fn func(uint) []uint) []uint {
	primes := fn(n)
	mersenne := make([]uint, 0, len(primes)/2)
	for _, p := range primes {
		mer := pow(2, p) - 1
		if mer > n {
			break
		}
		if contains(primes, mer) {
			mersenne = append(mersenne, mer)
		}
	}
	return mersenne
}
func contains(slice []uint, element uint) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

// Sieve of Eratosthenes
//
// Returns
//
// `[2, maxN]` primes or `[]`
func PrimeSieve(maxN uint) []uint {
	if maxN < 2 {
		return []uint{}
	}
	primes := []uint{2}
	for i := uint(3); i <= maxN; i++ {
		closure := func(p uint) bool {
			return i%p == 0
		}
		if any(primes, closure) {
			continue
		}
		// for _, p := range primes {
		// 	if i%p == 0 {
		// 		continue OUTER
		// 	}
		// }
		primes = append(primes, i)
	}
	return primes
}
func any(slice []uint, closure func(uint) bool) bool {
	for _, e := range slice {
		if closure(e) {
			return true
		}
	}
	return false
}
func pow(base, exp uint) uint {
	res := uint(1)
	for i := uint(0); i < exp; i++ {
		res *= base
	}
	return res
}
