# Go Expense Tracker & Settlement API

A high-performance REST API built in Go that acts as a shared expense tracker. It handles concurrent requests and uses a **Greedy Algorithm** to mathematically reduce complex group debts into the absolute minimum number of cash flow transactions.

## Key Features
* **Lightning Fast:** Sub-millisecond response times (`< 500µs`) powered by Go.
* **Algorithmic Optimization:** Implements the Minimum Cash Flow graph algorithm.
* **Thread-Safe:** Uses `sync.Mutex` to prevent concurrent memory access during high-volume requests.
* **Interactive UI:** Auto-generated OpenAPI/Swagger documentation.

## Project Showcases

### 1. Interactive Swagger UI
*(Drag and drop the screenshot of your Swagger UI webpage here)*

### 2. High-Speed Terminal Execution
*(Drag and drop the screenshot of your terminal showing the microsecond response times here)*

## Tech Stack
* **Language:** Go (Golang)
* **Framework:** Gin (HTTP web framework)
* **Documentation:** Swaggo (Swagger/OpenAPI 2.0)

## Quick Start Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/NakulSingh156/expense-tracker-go.git
   cd expense-tracker-go
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

Open the Swagger UI in your browser to interact with the API:
 http://localhost:8080/swagger/index.html

## Documentation
See `DESIGN.md` for a full breakdown of the architecture, concurrency handling, and algorithm time complexity.

See `AI_PROMPTS.md` for the transparency log of AI tools used during development.
