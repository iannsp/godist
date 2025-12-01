package godist

import (
	"math"
	"testing"
)

func TestBinomialCoefficient(t *testing.T) {
	// Define a struct for our test cases
	tests := []struct {
		name        string
		n           int
		k           int
		expected    int
		expectError bool
	}{
		// Standard Math Cases
		{"5 choose 3", 5, 3, 10, false},
		{"4 choose 2", 4, 2, 6, false},
		{"10 choose 5", 10, 5, 252, false},

		// Boundary / Edge Cases
		{"k is 0", 10, 0, 1, false},        // C(n,0) is always 1
		{"k is n", 10, 10, 1, false},       // C(n,n) is always 1
		{"k is 1", 5, 1, 5, false},         // C(n,1) is always n
		{"n is 0, k is 0", 0, 0, 1, false}, // C(0,0) is 1

		// Large Number Check (Safety limit for int64)
		{"Max safe int64 (66 choose 23)", 66, 23, 348524321356411200, false},
		{"Max safe int64 (100 choose 10)", 100, 15, 253338471349988640, false},

		// Error Cases
		{"k > n", 5, 6, 0, true},
		{"k < 0", 5, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinomialCoefficient(tt.n, tt.k)

			// Check Error Expectation
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return // Stop test here if we expected an error
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Check Value
			if got != tt.expected {
				t.Errorf("BinomialCoefficient(%d, %d) = %d; want %d", tt.n, tt.k, got, tt.expected)
			}
		})
	}
}

func TestBinomialPMF(t *testing.T) {
	// Floating point math requires a small delta for comparison
	const epsilon = 1e-9

	tests := []struct {
		name        string
		n           int
		p           float64
		k           int
		expected    float64
		expectError bool
	}{
		// Standard Case: Fair coin, 5 tosses, 3 heads
		// C(5,3) * 0.5^3 * 0.5^2 = 10 * 0.125 * 0.25 = 0.3125
		{"Fair coin 3 heads in 5", 5, 0.5, 3, 0.3125, false},

		// Standard Case: Biased coin (p=0.8), 3 trials, 3 successes
		// 1 * 0.8^3 * 0.2^0 = 0.512
		{"Biased coin 3/3", 3, 0.8, 3, 0.512, false},

		// Edge Cases
		{"0 probability, 0 successes", 5, 0.0, 0, 1.0, false}, // If p=0, 0 successes is guaranteed
		{"0 probability, 1 success", 5, 0.0, 1, 0.0, false},   // Impossible
		{"100% probability, n successes", 5, 1.0, 5, 1.0, false},

		// Error Cases
		{"p < 0", 5, -0.1, 2, 0.0, true},
		{"p > 1", 5, 1.1, 2, 0.0, true},
		{"k > n", 5, 0.5, 6, 0.0, true}, // Should bubble up error from BinomialCoefficient
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinomialPMF(tt.n, tt.p, tt.k)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Compare float values with epsilon
			if math.Abs(got-tt.expected) > epsilon {
				t.Errorf("BinomialPMF(%d, %.2f, %d) = %.9f; want %.9f", tt.n, tt.p, tt.k, got, tt.expected)
			}
		})
	}
}
