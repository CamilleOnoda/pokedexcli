## Building a Pokedex CLI in Go: lessons from boot.dev


This is a command-line Pokedex written in Go, built as part of the boot.dev curriculum. It interacts with the PokeAPI and demonstrates foundational Go patterns and CLI design.

---

## Quickstart

1. **Clone the repository:**
    ```sh
    git clone https://github.com/CamilleOnoda/pokedexcli.git
    cd pokedexcli
    ```
2. **Run the CLI:**
    ```sh
    go run .
    ```

---

## The project:
A command-line Pokedex that lets you explore Pokemon location areas with pagination, attempt to catch a Pokemon and add it to your Pokedex,<br> and view details about Pokemon you have caught.<br>


### Features and commands

The CLI allows you to:

- **help**: Display available commands
- **exit**: Close the application
- **map**: Show the next 20 location areas (pagination)
- **mapb**: Show the previous 20 location areas
- **explore <location>**: List all Pokemon in a specific location area
- **catch <pokemon>**: Attempt to catch a Pokemon and add it to your Pokedex
- **inspect <pokemon>**: View details of a Pokemon you have caught
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

- `main.go`: handles user interaction and command routing
- `internal/pokeapi`: manages all API communication
- `internal/pokecache`: in-memory caching layer for API responses
- This clear separation ensures that each component has a distinct responsibility, making the codebase easier to manage and test.


**Configuration as state**

Centralized application state is managed in a `config` struct, which is passed to all commands. This enables stateful pagination and supports dependency injection.<br>
Optional parameters are represented using pointer types (such as `*string`), allowing for flexible configuration and clear handling of optional values.

**Dependency injection pattern**

The dependency injection pattern is used by injecting the API client via the `config` struct. This design makes the code more testable and flexible, as it allows for easy mocking in unit tests and supports different configurations for development and production environments.


**Command pattern for extensibility**

Commands are defined in a map, making it easy to add new commands by simply adding entries. This approach supports scalable growth of the application's functionality.

**Constructor pattern**

Go idioms are followed for initialization (e.g., `NewClient`), encapsulating setup logic and making future changes easier.


**Error handling**

Error handling is performed using Go’s `%w` error wrapping, which preserves the error chain and adds context. This practice aligns with Go best practices, ensuring that errors are informative and traceable throughout the application.


---

## Next steps

- Add more comprehensive tests (real and mocked HTTP clients)
- Expand functionality (more commands, richer Pokedex features)
- Improve CLI UX and error messages