# Go Expense Tracker & Minimum Cash Flow API

A high-performance, concurrent REST API built in Go. This system acts as a shared expense tracker that mathematically reduces complex, overlapping group debts into the absolute minimum number of cash flow transactions using a Greedy Algorithm.

## Visual Proof & Performance

### 1. OpenAPI/Swagger UI Dashboard

<img width="1507" height="942" alt="image" src="https://github.com/user-attachments/assets/a57463c4-8665-4dc8-822f-afb5f18c0958" />

<img width="1509" height="942" alt="image" src="https://github.com/user-attachments/assets/438e8975-1cd7-405f-91d7-f11a58c734d8" />


### 2. Microsecond Execution Speeds

<img width="897" height="312" alt="Screenshot 2026-02-22 at 1 18 18 AM" src="https://github.com/user-attachments/assets/4c5fd5f8-ffda-4a6f-a094-d066c13ce540" />


---

## Methodology: From Zero to Deployment
This project was architected in four distinct engineering phases to ensure scalability, thread safety, and standard enterprise documentation.

### Phase 1: Domain Modeling & Memory Management
* **Static Typing:** Defined strict Go `structs` to represent `ExpenseRequests`, `Balances`, and `Transactions`.
* **State Management:** Implemented an in-memory Hash Map to track user ledgers dynamically without the latency overhead of a traditional disk-based database during the prototyping phase.

### Phase 2: Algorithmic Optimization (Core Engine)
* **The Problem:** Standard peer-to-peer payments create cyclical debt graphs (e.g., A owes B, B owes C, C owes A).
* **The Solution:** Implemented a **Greedy Algorithm** to solve the Minimum Cash Flow problem.
* **Execution:** 1. Calculates the net position of every entity.
  2. Partitions entities into `Debtors` (net negative) and `Creditors` (net positive).
  3. Recursively matches the maximum debtor with the maximum creditor until the ledger zeroes out.
* **Time Complexity:** `O(N log N)` due to the sorting mechanism applied during the greedy matching phase.

### Phase 3: Concurrent API Layer
* **Framework:** Utilized the **Gin HTTP framework** to expose the algorithmic engine via RESTful endpoints (`POST /expense`, `GET /balances`, `GET /settle`).
* **Thread Safety:** Because Go handles HTTP requests concurrently via Goroutines, a `sync.Mutex` was wrapped around the in-memory ledger. This prevents data race conditions and fatal concurrent map write errors under heavy traffic loads.

### Phase 4: OpenAPI Documentation Deployment
* **Implementation:** Integrated `swaggo/swag` and `swaggo/gin-swagger` to parse declarative block comments above handler functions.
* **Deployment:** Auto-generated a `swagger.json` specification file and served an interactive Swagger webpage directly from the Gin router, allowing frontend teams to test the API without external API clients like Postman.

---

## Tech Stack
* **Language:** Go (Golang)
* **Web Framework:** Gin
* **Concurrency:** Native Goroutines & `sync.Mutex`
* **API Documentation:** Swaggo (Swagger/OpenAPI 2.0)

---

## Quick Start Setup
1. Clone the repository:
   ```bash
   git clone [https://github.com/NakulSingh156/expense-tracker-go.git](https://github.com/NakulSingh156/expense-tracker-go.git)
   cd expense-tracker-go
Install dependencies:

go mod tidy
Run the server:

go run main.go
Open the Swagger UI in your browser to interact with the API:
 http://localhost:8080/swagger/index.html
