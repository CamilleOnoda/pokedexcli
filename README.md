## Building a Pokedex CLI in Go: lessons from boot.dev


This is a command-line Pokedex written in Go, built as part of the boot.dev curriculum. It interacts with the PokeAPI and demonstrates foundational Go patterns and CLI design.

---

## Prerequisites and setup

This project requires [Go](https://golang.org/dl/) (version 1.18 or newer recommended).


1. **Install Go:**
   - Download and install Go from the [official website](https://golang.org/dl/).
   - Follow the instructions for your operating system.
   - After installation, open a new terminal and run:
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


## The project
A command-line Pokedex that lets you:
- Explore Pokemon location areas with pagination
- Attempt to catch Pokemon and add them to your Pokedex
- View details about Pokemon you have caught


### Features and commands

The CLI allows you to:

- **help**: Display available commands
- **exit**: Close the application
- **map**: Show the next 20 location areas (pagination)
- **mapb**: Show the previous 20 location areas
- **explore <location>**: List all Pokemon in a specific location area
- **catch <pokemon>**: Attempt to catch a Pokemon and add it to your Pokedex
- **inspect <pokemon>**: View details about a Pokemon you have caught
- **pokedex**: List all Pokemon you have caught so far

Example usage:
```
Pokedex > map
Pokedex > mapb
Pokedex > explore viridian-forest
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
Pokedex > exit
```

---


## Key architecture decisions


**Clean package structure**

- `main.go`: Handles user interaction and command routing
- `internal/pokeapi`: Manages all API communication
- `internal/pokecache`: In-memory caching layer for API responses
- Clear separation ensures each component has a distinct responsibility, making the codebase easier to manage and test.

**Configuration as state**

Centralized application state is managed in a `config` struct, passed to all commands. This enables stateful pagination and supports dependency injection. Optional parameters are represented using pointer types (such as `*string`).

**Dependency injection pattern**

The API client is injected via the `config` struct, making the code more testable and flexible. This allows for easy mocking in unit tests and supports different configurations for development and production.

**Command pattern for extensibility**

Commands are defined in a map, making it easy to add new commands by simply adding entries. This approach supports scalable growth of the application's functionality.

**Constructor pattern**

Go idioms are followed for initialization (e.g., `NewClient`), encapsulating setup logic and making future changes easier.

**Error handling**

Errors are wrapped using Go’s `%w` error wrapping, preserving the error chain and adding context. This ensures errors are informative and traceable throughout the application.

---

## Next steps

- Add more comprehensive tests (real and mocked HTTP clients)
- Expand functionality (more commands, richer Pokedex features)
- Improve CLI UX and error messages