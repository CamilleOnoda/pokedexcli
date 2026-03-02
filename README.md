## Pokedex CLI in Go

This is a command-line Pokedex written in Go. It interacts with the PokeAPI and demonstrates foundational Go patterns, CLI design, and testable architecture.

---

## Prerequisites and setup

This project requires [Go](https://golang.org/dl/) (1.18+ recommended).

1. **Install Go:**
   - Download and install Go from the [official website](https://golang.org/dl/).
   - After installation, run:
     ```sh
     go version
     ```
     You should see output like `go version go1.21.0 linux/amd64`.

2. **Clone the repository:**
   ```sh
   git clone https://github.com/CamilleOnoda/pokedexcli.git
   cd pokedexcli
   ```

3. **Download dependencies:**
   ```sh
   go mod download
   ```

4. **Run the CLI:**
   ```sh
   go run .
   ```

If you see errors about missing Go or commands not found, double-check your Go installation and that your terminal recognizes the `go` command.

---

## Project overview

This CLI lets you:
- Explore Pokémon location areas with pagination
- Attempt to catch Pokémon and add them to your Pokedex
- View and inspect details about Pokémon you have caught
- List all Pokémon you have caught so far
- Clear your Pokedex

### Features and commands

The CLI allows you to:

- **help**: Display available commands
- **exit**: Close the application
- **map**: Show the next 20 location areas (pagination)
- **mapb**: Show the previous 20 location areas
- **explore <location>**: List all Pokémon in a specific location area
- **catch <pokemon>**: Attempt to catch a Pokémon and add it to your Pokedex
- **inspect <pokemon>**: View details about a Pokémon you have caught
- **pokedex**: List all Pokémon you have caught so far
- **clear**: Clear your Pokedex

Example usage:
```
Pokedex > map
Pokedex > mapb
Pokedex > explore viridian-forest
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
Pokedex > clear
Pokedex > exit
```

---

## Architecture and design

- **main.go**: Handles user interaction and command routing
- **internal/pokeapi**: Manages all API communication, caching, and data types
  - `client.go`: HTTP client for PokeAPI
  - `cached_client.go`: Adds caching to API requests
  - `cache.go`: In-memory cache with TTL
  - `interface.go`: Interface for API clients (enables mocking/testing)
  - `types_*.go`: Data types for API responses
- **mock_client.go**: Mock implementation for testing

**Key Patterns:**
- Centralized application state in a `config` struct
- Dependency injection for API client (enables easy mocking in tests)
- Command pattern for extensibility (commands in a map)
- Constructor pattern for initialization (e.g., `NewClient`)
- Error wrapping for informative error messages

---

## Testing

Unit tests cover input cleaning, cache logic, and command behaviors. Mock clients are used to test CLI logic without real API calls.

- `repl_test.go`: Tests input parsing and cleaning
- `mock_client_test.go`: Tests catching logic and command behaviors
- `internal/pokeapi/cache_test.go`: Tests cache set/get and expiration

To run all tests:
```sh
go test ./...
```

---

## Next steps

- Add more comprehensive tests (real and mocked HTTP clients)
- Expand functionality (more commands, richer Pokedex features)
- Improve CLI UX and error messages
- **pokedex**: List all Pokemon you have caught so far

