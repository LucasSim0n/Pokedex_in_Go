# Pokedex_in_Go

CLI Pokedex written in Go, connected to [PokeAPI](https://pokeapi.co). Guided project from Boot.dev academy.

---

## Table of Contents

- [Description](#description)  
- [Features](#features)  
- [Requirements](#requirements)  
- [Installation](#installation)  
- [Usage](#usage)  
- [Commands](#commands)  
- [Project Structure](#project-structure)  
- [Contributing](#contributing)  
- [License](#license)  

---

## Description

Pokedex_in_Go is a command-line interface (CLI) application built with Go. It fetches data from PokeAPI to provide Pokémon information (names, details, location, map data, etc.) via terminal commands. The project is intended to practice CLI development, API consumption, and idiomatic Go.  

---

## Features

- Real-time fetch of Pokémon data from PokeAPI  
- Multiple CLI commands: explore, map, help, etc.  
- REPL (Read-Eval-Print Loop) mode for interactive use  
- Structured internal code: separate files for different commands
- Cached data for a better performance

---

## Requirements

- Go (version ≥ 1.18 recommended)  
- Internet connection to access PokeAPI  

---

## Installation

1. Clone the repository:  
   ```bash
   git clone https://github.com/LucasSim0n/Pokedex_in_Go.git
   cd Pokedex_in_Go
   ```

2. Install dependencies:  
   ```bash
   go mod download
   ```

3. Build the CLI executable:  
   ```bash
   go build -o pokedex
   ```

---

## Usage

Run the executable, or use `go run`:

```bash
./pokedex
```

or

```bash
go run main.go
```

---

## Commands

| Command             | Description                                     |
|---------------------|-------------------------------------------------|
| `inspect <name>`    | Get detailed info about a specific Pokémon     |
| `catch <name>`      | Try to capture a pokemon                       |
| `list`              | List your captured pokemons                    |
| `explore <area>`    | View Pokémon available in a specific area      |
| `map`               | Display locations to explore                   |
| `help`              | Show help / list of available commands         |
| `exit`              | Exit REPL session and quit the program         |

---

## Project Structure

```
Pokedex_in_Go/
├── internal/
│   ├── commands.go
│   ├── inspectCommand.go
│   ├── listCommand.go
│   ├── mapCommand.go
│   ├── exploreCommand.go
│   └── other command files...
├── main.go
├── repl.go
├── helpCommand.go
├── exitCommand.go
├── catchCommand.go
├── go.mod
└── go.sum
```

---

## Contributing

Contributions are welcome! If you'd like to help:

1. Fork the repository  
2. Create a new branch for your feature or bugfix  
3. Ensure code is formatted (e.g. `go fmt`)  
4. Write tests if applicable  
5. Open a pull request  

