# Drokkit Game Server

Drokkit is a multiplayer, turn-based game server built in Go, designed to support a variety of game genres through REST APIs and WebSockets. It leverages MariaDB for relational data storage, Redis for caching and real-time data, and NATS for messaging. This server is designed to be modular and extensible, with a focus on reliability and performance.

## Features

- Player Authentication: Secure login and registration with JWT-based authentication.
- Game Management: Create and manage game sessions, handle game events, and process moves in real time.
- Faction and Alliance Management: Set up and manage alliances, factions, and resource allocation.
- Leaderboard Tracking: Maintain player leaderboards with Redis for fast retrieval.
- Real-Time Gameplay: Uses WebSocket communication for player interactions and live game state updates.

## Requirements

### For Docker

- Docker

### To build from source

- Go 1.18+
- Nginx (for reverse proxy configuration)
- NATS
- Redis
- MariaDB

## Documentation

Comprehensive documentation is provided in the drokkumentation folder:

[Server Configuration and Setup](drokkumentation/running.md)
[Developing a Game Client in C#](drokkumentation/writingagameincsharp.md)
[Developing a Game Client in GDScript (Godot)](drokkumentation/writingagameingdscript.md)
[API Endpoints and WebSocket Guide](drokkumentation/apiendpoints.md)

## Project Structure

- **dist/** - Contains the server binary, drokkgen, and the runner script.
- **drokkumentation/** - Documentation files for setup, client development, and API usage.
- **models/** - GORM models representing database entities.
- **handlers/** - API handlers for routes.
- **config/** - Configuration files for database, Redis, and NATS setup.
- **routes/** - HTTP routes defined for the server.
- **scripts/** - Additional scripts, e.g., for building and managing Docker containers.

### Setup

**Running the Server:**

The server can be built and run in a Docker environment. Make sure you have Docker and Docker Compose installed, then follow these steps:

Clone the repository:

```bash
git clone <https://github.com/yourusername/drokkit>
cd drokkit
```

Build and Run with Docker:

```bash
make docker-build
make docker-run
```

Clean Up Docker Containers:

```bash
make docker-clean
```

**Configuration:**

The server uses environment variables for configuration. These can be set in a .env file at the root level. The drokkgen utility generates a .env file with secure defaults. Run it with:

```bash
./dist/drokkgen
```

### Contributing

If you would like to contribute, please open an issue or submit a pull request. Contributions are welcome!

### License

This project is licensed under the MIT License.
