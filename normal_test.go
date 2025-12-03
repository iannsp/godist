package godist

import (
_	"math"
	"testing"
    "log/slog"
)

func TestNormalPDF(t *testing.T) {
/*
when manufacturing Light Buls with a Mean Life = 1000hours and Standard deviation
of 50hours. what is the PDF at the peak of the bell(1000)?

When you say the Probability Density Function (PDF) at 1000 hours is 0.0079, it means the relative likelihood of a light bulb lasting around 1000 hours is described by that value. It does not mean there is a 0.0079 probability that a bulb will last exactly 1000 hours. 

Here is a breakdown of what the value represents: 
Probability vs. Probability Density 

For a continuous random variable (like the exact lifespan of a light bulb), the probability of it taking on any single, specific value (e.g., exactly 1000.000... hours) is essentially zero because there are infinitely many possible values. 

Instead, the PDF value at a point (the y-axis value on the bell curve) describes the density of the probability distribution at that point. 

Interpretation of 0.0079 

The value 0.0079 can be interpreted in these ways: 

Relative Likelihood: It indicates that values very close to the mean (1000 hours) are more "dense" or likely to occur than values further away, which would have a lower PDF value.

Probability over an Interval: The PDF is used to find probabilities over a range of values by calculating the area under the curve (integration). The value 0.0079 at 1000 hours is the height of the curve at that specific point.

Comparison of Ranges: If you consider a very small interval, say between 1000 hours and 1000.001 hours, the probability of the lifespan falling within this range would be approximately \(0.0079\times 0.001\) (density times the small interval length). This is a very small probability, as expected for a narrow range. 

In summary, the number 0.0079 is a measure of the curve's height at the mean, illustrating that the mean lifespan is the most likely region for the light bulbs' lifespans to fall within, but it is not a direct probability for that single hour value.

*/
    n := NewNormal(1000, 50)
    pdf:= n.PDF(1000)
    slog.Info("a","pdf",  pdf, "n", n) //0.007978
}

func TestNormalCDF(t *testing.T) {
/*
finding the probability of a student scoring below 90 on a test with a mean of 75  
and a standard deviation of 10  
*/
    n := NewNormal(75, 10)
    cdf:= n.CDF(90.0)
    slog.Info("a","cdf",  cdf, "n", n)
}




