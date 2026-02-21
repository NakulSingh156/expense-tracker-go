package algorithm

import (
	"expense-tracker/models"
	"math"
	"sort"
)

func MinimizeCashFlow(balances map[string]float64) []models.Transaction {
	var debtors []models.Balance
	var creditors []models.Balance

	// 1. Separate into debtors (negative) and creditors (positive)
	for user, amount := range balances {
		if amount < 0 {
			debtors = append(debtors, models.Balance{UserID: user, Amount: amount})
		} else if amount > 0 {
			creditors = append(creditors, models.Balance{UserID: user, Amount: amount})
		}
	}

	var transactions []models.Transaction

	// 2. Process until debts are settled
	for len(debtors) > 0 && len(creditors) > 0 {
		// Sort to always match the highest debtor with highest creditor
		sort.Slice(debtors, func(i, j int) bool { return debtors[i].Amount < debtors[j].Amount })
		sort.Slice(creditors, func(i, j int) bool { return creditors[i].Amount > creditors[j].Amount })

		debtor := &debtors[0]
		creditor := &creditors[0]

		// Find the minimum of what the debtor owes and creditor is owed
		settleAmount := math.Min(-debtor.Amount, creditor.Amount)

		// Record the transaction
		transactions = append(transactions, models.Transaction{
			From:   debtor.UserID,
			To:     creditor.UserID,
			Amount: settleAmount,
		})

		// Update balances
		debtor.Amount += settleAmount
		creditor.Amount -= settleAmount

		// Remove settled users from the list
		if debtor.Amount == 0 {
			debtors = debtors[1:]
		}
		if creditor.Amount == 0 {
			creditors = creditors[1:]
		}
	}

	return transactions
}
