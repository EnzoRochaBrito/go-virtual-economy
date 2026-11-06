package wallet

import "time"

type WalletPolicy struct {
	MaxAge    int // Days
	MaxLedges int // Max ledges since newest snapshot
}

func NewWalletPolicy(maxAge, maxLedges int) WalletPolicy {
	return WalletPolicy{
		MaxAge:    maxAge,
		MaxLedges: maxLedges,
	}
}

func (p WalletPolicy) ShouldSnapshot(lastSnapshot WalletSnapshot, ledgersSinceSnapshot int) bool {
	if ledgersSinceSnapshot == 0 {
		return false
	}
	ageExceeded := time.Since(lastSnapshot.EffectiveAt) >= time.Duration(p.MaxAge)*24*time.Hour
	ledgersExceeded := ledgersSinceSnapshot >= p.MaxLedges
	return ageExceeded || ledgersExceeded
}