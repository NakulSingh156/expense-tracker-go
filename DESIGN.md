# System Design & Architecture

## 1. High-Level Architecture
The system is designed as a high-performance RESTful microservice. To prioritize speed and concurrency, the application is written in **Go (Golang)** utilizing the **Gin** framework for routing. State is temporarily managed in memory to demonstrate algorithmic logic without database latency.

## 2. Code Explanation & Data Flow
* `main.go`: The entry point. Initializes the Gin router, serves the Swagger UI, and maps the HTTP endpoints to their respective handler functions.
* `api/handlers.go`: The controller layer. It parses incoming JSON payloads, safely updates the in-memory ledger, and calls the algorithm package.
* `models/models.go`: The data layer. Enforces strict static typing for the `ExpenseRequest`, `Balance`, and `Transaction` structs.
* `algorithm/settlement.go`: The business logic layer containing the debt optimization algorithm.

## 3. Concurrency Handling (Thread Safety)
Go is highly concurrent. To prevent a race condition (e.g., `fatal error: concurrent map writes`) when multiple users attempt to log expenses at the exact same microsecond, the system implements a `sync.Mutex`. This acts as a lock on the memory map, ensuring only one Goroutine can update the balances at a time, resulting in a thread-safe environment.

## 4. Algorithmic Approach (The Greedy Algorithm)
The core challenge is the "Minimum Cash Flow" problem. If A owes B, and B owes C, the algorithm must bypass B and force A to pay C directly.

**Step-by-Step Execution:**
1. **Calculate Net Balances:** (Amount Paid - Amount Consumed = Net Balance).
2. **Divide:** Users are split into `Debtors` (negative balance) and `Creditors` (positive balance).
3. **Greedy Matching:**
   * Sort both lists.
   * Take the largest debtor and force them to pay the largest creditor.
   * The settled amount is the minimum of the two absolute balances.
   * Repeat recursively until all balances are zero.
4. **Time Complexity:** Sorting takes `O(N log N)`, and matching takes `O(N)`, resulting in a highly efficient `O(N log N)` overall complexity.
