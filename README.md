## Building a Pokédex CLI in Go: lessons from boot.dev

As part of the boot.dev's curriculum, I started building a CLI application that interacts with the PokéAPI.

### The project:
A command-line Pokédex that lets you explore Pokémon location areas with pagination.<br>
Simple concept, but lots of foundational Go patterns.

**What the application does**

The CLI allows users to navigate through Pokémon location areas with commands like:
- help: display available commands
- exit: close the application
- map: show next 20 location areas
- mapb: show previous 20 location areas (with pagination state management)

**Key architecture decisions**:

🔹 Clean package structure

- *main.go* handles user interaction and command routing
- *internal/pokeapi* manages all API communication
- Separation of concerns makes the codebase maintainable and testable

🔹 Configuration as state
```
Go code

    type config struct {
        pokeAPIClient       pokeapi.Client
        nextLocationURL     *string
        previousLocationURL *string
    }
```
- Centralizing application state in a config struct that gets passed to commands.
- Enables stateful pagination.
- Using pointer types (*string) as function parameters to represent optional parameter that can be *nil*.
    - *nil* explicitly means **optional**. It is impossible to pass an empty string as a valid URL.
    - Passing an string value instead would mean assuming that an empty string is never a valid input.
- Instead of creating the API client inside functions, I inject it through the config struct.

🔹Dependency injection pattern

As mentionned above, I pass the API client through the config struct, meaning that I give an object<br>
the elements it needs from the outside, rather than having to create them itself.

Simple analogy: instead of a chef going to the store to buy ingredients, someone hands the chef the ingredients they need (dependency injection).

Why use it?

- Enables testing with moch clients (no realAPI calls in tests).
- Allows different configurations for development/production.
- Follows the "Accept interfaces, return structs" guideline.

🔹Command pattern for extensibility
```
gotype cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

var commands = map[string]cliCommand{
    "map":  {name: "map", callback: commandMap},
    "mapb": {name: "mapb", callback: commandMapb},
}
```
Adding new commands means just adding entries to the map.

🔹 The constructor pattern
```
type Client struct {
	httpClient *http.Client
}

gofunc NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: &http.Client{Timeout: timeout},
    }
}
```
Following Go idioms: New* functions for initialization, encapsulating setup logic, and making future changes easier without.

🔹 Error handling

```
	resp, err := c.httpClient.Get(url)
	if err != nil {
		err := fmt.Errorf("Error making GET request to %s: %w", url, err)
		return LocationAreaResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Received a non-OK HTTP status code: %d", resp.StatusCode)
		return LocationAreaResponse{}, err
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaResponse); err != nil {
		err := fmt.Errorf("Error decoding JSON response: %w", err)
		return LocationAreaResponse{}, err
	}
```

Learning to wrap errors with %w to preserve the error chain while adding context.

### Next steps:
- Adding a caching layer.
- Comprehensive tests to practice using both real and mocked HTTP clients
- Expanding functionality to actually to:
    - explore: a user uses the map commands to find a location area, and see a list of all the Pokémon located there.
    - catch: catching a pokemon adds them to a user's Pokedex
    - inspect: allow players to see details of a Pokemon if they have seen (caught) it before.