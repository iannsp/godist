// Frequency distribution of the possible number od sucess outcomes in a givn number
// of trials in each of which there is the same probability of sucess.

// this implementation use int to have to check for overflow.

package godist

import (
	"errors"
	"math"
)

// BinomialCoefficient computes "n choose k".
//
// I use factorial method in the first implementation but
// change to the multiplicative formula to avoid calculating large factorials,
// allowing for higher n than the factorial method.
//
// parameters:
// n: Number of trials
// k: Number of sucesses
func BinomialCoefficient(n, k int) (int, error) {

	if k > n || k < 0 {
		err := errors.New("k must be between 0 and n")
		return 0, err
	}

	// use identity property to optimize. C(n, k) == C(n, n-k)
	if k > n/2 {
		k = n - k
	}

	result := 1

	for i := 1; i <= k; i++ {
		if result > math.MaxInt64/(n-i+1) {
			return 0, errors.New("integer overflow")
		}

		result = result * (n - i + 1)
		result = result / i
	}
	return result, nil
}

// BinomialPMF calculates the probability of exactly k successes in n independent trials.
// P(X=k) = C(n,k) * p^k * (1-p)^(n-k)
//
// Parameters:
// n: number of trials
// p: probability of success in a single trial (0 <= p <= 1)
// k: number of successes
func BinomialPMF(n int, p float64, k int) (float64, error) {
	if p < 0 || p > 1 {
		return 0, errors.New("probability p must be between 0 and 1")
	}

	binomCoeff, err := BinomialCoefficient(n, k)
	if err != nil {
		return 0, err
	}

	q := 1 - p

	pk := math.Pow(p, float64(k))
	qn := math.Pow(q, float64(n-k))

	return float64(binomCoeff) * pk * qn, nil
}
