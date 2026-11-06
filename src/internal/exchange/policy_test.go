package exchange

import "testing"

func TestPolicyFee(t *testing.T) {
	p := Policy{
		FeePercentage: 25,
	}
	calc := p.CalculateFee(100)
	if calc.Fee.Int() != 25 || calc.Net != 75 {
		t.Errorf("expected fee 25, and got: %d, expected net 75, and got %d", calc.Fee, calc.Net)
	}
}