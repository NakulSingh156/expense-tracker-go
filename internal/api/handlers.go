package api

import (
	"expense-tracker/internal/algorithm"
	"net/http"
	"sync" // <-- NEW: Import sync package

	"github.com/gin-gonic/gin"
)

var (
	balances = make(map[string]float64)
	mu       sync.Mutex // <-- NEW: The bouncer for our map
)

// AddExpense Request Payload
type ExpenseRequest struct {
	Payer        string   `json:"payer" binding:"required"`
	Amount       float64  `json:"amount" binding:"required,gt=0"`
	Participants []string `json:"participants" binding:"required,min=1"`
}

// @Summary Add a new shared expense
// @Description Calculates the split and updates the in-memory balances
// @Tags expenses
// @Accept json
// @Produce json
// @Param expense body ExpenseRequest true "Expense Details"
// @Success 200 {object} map[string]interface{}
// @Router /expense [post]
func AddExpense(c *gin.Context) {
	var req ExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	splitAmount := req.Amount / float64(len(req.Participants))

	// LOCK THE MAP BEFORE WRITING
	mu.Lock()
	balances[req.Payer] += req.Amount
	for _, person := range req.Participants {
		balances[person] -= splitAmount
	}
	// UNLOCK AFTER WRITING
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "Expense added successfully!",
	})
}

// @Summary Get current balances
// @Description Returns the current outstanding net balances for all users
// @Tags balances
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /balances [get]
func GetBalances(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock() // Ensures it unlocks when the function finishes
	
	c.JSON(http.StatusOK, gin.H{
		"current_balances": balances,
	})
}

// @Summary Settle all debts
// @Description Runs the greedy algorithm to return the optimized minimum transactions
// @Tags settlement
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /settle [get]
func SettleDebts(c *gin.Context) {
	mu.Lock()
	// Pass a copy of the balances to the algorithm to prevent locking it up for too long
	balancesCopy := make(map[string]float64)
	for k, v := range balances {
		balancesCopy[k] = v
	}
	mu.Unlock()

	transactions := algorithm.MinimizeCashFlow(balancesCopy)

	c.JSON(http.StatusOK, gin.H{
		"message":                "Debts optimized successfully",
		"optimized_transactions": transactions,
	})
}
