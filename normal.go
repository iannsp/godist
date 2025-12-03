// Normal Distribution 

package godist

import (
_	"errors"
	"math"
)

type Normal struct{
    mean float64
    stddev float64 // standard Deviation
}

func NewNormal( mean, stdDev float64) Normal{
    return Normal{mean: mean, stddev: stdDev}
}

// PDF (Probability Dentisty Function)
// Calculate how plausible a specific value is compared to others in a continuous distribution
func (n *Normal) PDF( v float64) float64{
    // z score
    z := n.SingleDataZ(v)
    // The bell shape.
    exponentialDecay := math.Exp(-0.5 * z * z)
    // the total Area under the curve must Equal 1.0
    heightAdjuster := (1.0 / (n.stddev * math.Sqrt(2*math.Pi)))
    return  heightAdjuster * exponentialDecay
}

// CDF (Cumulative Distribution Function)
// Calculate the probability that a random variable is Less than or Equal to v
// using math.Erf
func (n *Normal) CDF( v float64) float64{
    return 0.5 * (1 + n.SampleMeanZ( v ))
}

func (n *Normal) Less( v float64) float64 {
    return n.CDF(v)
}
func (n *Normal) Grater (v float64) float64 {
    return 1 - n.CDF(v)
}

func (n *Normal) Between (a,b float64) float64{
    return math.Abs(n.CDF(b) - n.CDF(a))
}
// Z Standardizes the input value v.
// The exponential part of the bell curve relies on distance from the center regardless of wheter you are measuring so Z convert your specific unit to a "universal"distance.
//
// How many standard deviations away from the mean the z value is.
// If z is 0, v is exacly the average.
// If z is 2.0, v is far to the right.
// if z is -1.0, v is to the left.
func (n *Normal) SingleDataZ( v float64) float64{
    return (v - n.mean) / n.stddev
}
func (n *Normal) SampleMeanZ( v float64) float64{
    return math.Erf( (v - n.mean) / (n.stddev * math.Sqrt(2)) )
}
