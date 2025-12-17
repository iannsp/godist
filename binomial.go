// Frequency distribution of the possible number od sucess outcomes in a givn number
// of trials in each of which there is the same probability of sucess.

// this implementation use int to have to check for overflow.

package godist

import (
	"errors"
	"math"
)
type Binomial struct{
    trials int
    probSuccess float64
}


func NewBinomial(trials int, probSuccess float64) (Binomial, error) {

    if probSuccess < 0 || probSuccess > 1 {
        return Binomial{}, errors.New("Probability Success must be between 0 and 1")
    }

    if trials < 0 {
        return Binomial{}, errors.New("Number of Trials must be non negative")
    }

    return Binomial{ trials: trials, probSuccess: probSuccess}, nil
}
// BinomialCoefficient computes "n choose k".
//
// I use factorial method in the first implementation but
// change to the multiplicative formula to avoid calculating large factorials,
// allowing for higher n than the factorial method.
//
// parameters:
// k: Number of sucesses
func (b *Binomial)Coefficient(k int) (int, error) {
    n := b.trials
	if k > b.trials || k < 0 {
		err := errors.New("k must be between 0 and number of Trials")
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
// k: number of successes
func (b *Binomial)PMF(k int) (float64, error) {

	binomCoeff, err := b.Coefficient(k)
	if err != nil {
		return 0, err
	}

	q := 1 - b.probSuccess

	pk := math.Pow(b.probSuccess, float64(k))
	qn := math.Pow(q, float64(b.trials-k))

	return float64(binomCoeff) * pk * qn, nil
}
