package exchange

import (
	"fmt"

	veUtils "virtual.economy/src/shared/utils"
)

type Policy struct {
	FeePercentage uint8
}

func NewPolicy(percentage uint8) (Policy, error) {
	if percentage > 100 {
		return Policy{}, fmt.Errorf("Fee percentage is over 100: %v", percentage)
	}
	return Policy{
		FeePercentage: percentage,
	}, nil
}

func (p Policy) CalculateFee(amount veUtils.Currency) CalculatedFee {
	fee := (amount.Int() / 100) * int(p.FeePercentage)
	net := amount - veUtils.Currency(fee)
	return CalculatedFee{
		Net: net,
		Fee: veUtils.Currency(fee),
	}
}

type CalculatedFee struct {
	Net, Fee veUtils.Currency
}

// Ouputs the id used to reference the exchange that the system will use to external send/recieve with the profile
func (p Policy) ProfileSystemExchangeId(profileId string) string {
	return profileId + "_SYS"
}
