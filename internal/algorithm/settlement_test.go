package algorithm

import (
	"testing"
)

func TestMinimizeCashFlow(t *testing.T) {
	// Setup a test scenario:
	// Alice has 0 net balance. Bob is owed $50. Charlie owes $50.
	balances := map[string]float64{
		"Alice":   0,
		"Bob":     50,
		"Charlie": -50,
	}

	// Run the algorithm
	transactions := MinimizeCashFlow(balances)

	// We expect exactly 1 transaction: Charlie pays Bob $50
	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(transactions))
	}

	if transactions[0].From != "Charlie" || transactions[0].To != "Bob" || transactions[0].Amount != 50 {
		t.Errorf("Algorithm failed. Got: %+v", transactions[0])
	}
}
