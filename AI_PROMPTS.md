# AI Transparency Log

As per the project guidelines, AI was utilized for architectural scaffolding, Go syntax assistance, and algorithm implementation. Below are the prompts used during the development process:

* **Prompt 1:** "I am building a Go REST API for an Expense Tracker (like Splitwise). Provide the standard Go project folder structure and the structs needed for balances and transactions."
* **Prompt 2:** "Write a Go function to implement the minimum cash flow algorithm (debt settlement) using a greedy approach, taking a map of net balances and returning a list of optimized transactions."
* **Prompt 3:** "Write the REST API handlers using the Gin framework for an expense tracker. I need a POST endpoint to add an expense that divides the amount among participants and updates an in-memory balance map, and a GET endpoint that calls my MinimizeCashFlow algorithm. Include `sync.Mutex` for thread safety."
* **Prompt 4:** "Add Swagger UI documentation to my Gin Go API. Provide the Swaggo comments for my expense, balances, and settle endpoints, and show how to configure `main.go` to serve the Swagger interface."
