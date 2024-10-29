
# Drokkit API Endpoints and WebSocket Guide

This document provides an overview of the REST API endpoints and WebSocket connections available in the Drokkit Game Server. The server enables secure authentication, resource updates, real-time game state management, and administrative actions.

## Table of Contents

- Authentication
- Player Endpoints
- Match and Game Endpoints
- Resource Management
- Faction and Alliance Management
- Leaderboard
- WebSocket Connections
- Admin Endpoints

## Authentication

- `POST /register`: Registers a new player account.
  - **Request Body**: `{"username": "<username>", "password": "<password>"}`
  - **Response**: Created player information.
- `POST /login`: Logs in a player, generating a JWT token.
  - **Request Body**: `{"username": "<username>", "password": "<password>"}`
  - **Response**: JWT token for session authentication.

The JWT token is stored as a cookie named `token` and is required for all subsequent authenticated requests.

## Player Endpoints

- `GET /api/player/<id>`: Retrieves player information and stats.
  - **Headers**: `Authorization: Bearer <JWT Token>`
  - **Response**: `{"id": 1, "username": "player1", "stats": {...}}`

## Match and Game Endpoints

- `POST /api/match`: Creates a new game match between two players.
  - **Request Body**: `{"player_one": <PlayerID>, "player_two": <PlayerID>}`
  - **Response**: Match data, including game state.
- `POST /api/match/<id>/turn`: Updates the game state with a new player move.
  - **Request Body**: `{"player_id": <PlayerID>, "action": "<move description>"}`
  - **Response**: Updated match data with the new turn and game state.

## Resource Management

- `POST /api/resource`: Updates or adds a resource for a player in a specific game.
  - **Request Body**: `{"game_instance_id": <GameID>, "player_id": <PlayerID>, "type": "<resource type>", "amount": <amount>}`
  - **Response**: Resource data after update.

## Faction and Alliance Management

- `POST /api/faction`: Creates a faction within a game instance.
  - **Request Body**: `{"game_instance_id": <GameID>, "faction_type": "<type>", "leader_id": <PlayerID>}`
  - **Response**: Created faction data.
- `POST /api/alliance`: Forms an alliance between two factions.
  - **Request Body**: `{"game_instance_id": <GameID>, "name": "<AllianceName>", "faction_ids": [<FactionID1>, <FactionID2>]}`
  - **Response**: Alliance information with member data.

## Leaderboard

- `GET /leaderboard`: Retrieves the leaderboard.
  - **Query Parameters**:
    - `type`: individual or team (optional, defaults to individual).
    - `timeframe`: all-time, monthly, or weekly (optional, defaults to all-time).
  - **Response**: JSON array of player/team leaderboard positions.

## WebSocket Connections

- `GET /ws/play`: Establishes a WebSocket connection for real-time gameplay and turn management.
  - **Query Parameters**: `token=<JWT Token>`
  - **Usage**: Used by clients to send moves and receive opponent moves in real time.

### WebSocket Messages

- **Move Submission**: Client sends a JSON object `{ "player_id": <PlayerID>, "action": "<move>" }`.
- **Broadcast**: All other connected players in the same game instance receive this message.

## Admin Endpoints

- `POST /admin/create`: Registers a new admin user.
  - **Request Body**: `{"user_id": <UserID>, "permissions": "<permission level>"}`
- `DELETE /admin/delete-player`: Deletes a player account.
  - **Request Body**: `{"player_id": <PlayerID>}`
  - **Response**: Confirmation of player deletion.

> Note: For secure access, always use the JWT token issued upon login for any API calls that require authorization.

This guide serves as a reference for developers building clients or administrative tools for the Drokkit Game Server.
