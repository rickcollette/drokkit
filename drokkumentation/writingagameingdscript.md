# Drokkit Godot GDScript Client Development Guide

## Overview

This guide will walk you through developing a game client in Godot Engine using GDScript to interact with the Drokkit game server. This guide assumes you have a working Godot project and are familiar with basic GDScript concepts.

## Prerequisites

- Install [Godot Engine](https://godotengine.org/download)
- Basic knowledge of GDScript and the Godot Engine

## Table of Contents

- [Drokkit Godot GDScript Client Development Guide](#drokkit-godot-gdscript-client-development-guide)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Table of Contents](#table-of-contents)
  - [Setting up the Environment](#setting-up-the-environment)
  - [Configuring HTTP Requests](#configuring-http-requests)
    - [HTTPRequest Node Setup](#httprequest-node-setup)
    - [Function to Send HTTP Request](#function-to-send-http-request)
  - [Handling Player Registration \& Login](#handling-player-registration--login)
    - [Registration Example](#registration-example)
    - [Login Example](#login-example)
  - [WebSocket Game Events](#websocket-game-events)
    - [Setup WebSocket Client](#setup-websocket-client)
  - [Submitting and Receiving Game Moves](#submitting-and-receiving-game-moves)
    - [Handling Game Events](#handling-game-events)
  - [Leaderboard Retrieval](#leaderboard-retrieval)
  - [Error Handling](#error-handling)
  - [Best Practices](#best-practices)
  - [Conclusion](#conclusion)

## Setting up the Environment

To start, add the necessary server details in your `Project Settings` or within an autoload singleton to centralize your API configuration.

Example `global.gd`:

```gdscript
extends Node

# Replace with your Drokkit server URL and WebSocket endpoint
var SERVER_URL = "http://localhost:8080"
var WS_URL = "ws://localhost:8080/ws/play"
var TOKEN = ""
```

## Configuring HTTP Requests

In Godot, you'll use the `HTTPRequest` node to manage HTTP requests. Attach `HTTPRequest` to a node and use it to register players, authenticate, and retrieve data.

### HTTPRequest Node Setup

Add an HTTPRequest node to a main scene or singleton:

```gdscript
var http_request = HTTPRequest.new()
add_child(http_request)
```

### Function to Send HTTP Request

```gdscript
func send_request(endpoint: String, method: String, data: Dictionary = {}) -> void:
    var full_url = global.SERVER_URL + endpoint
    var headers = ["Content-Type: application/json"]
    if global.TOKEN:
        headers.append("Authorization: Bearer " + global.TOKEN)
    http_request.request(
        full_url,
        headers,
        true,
        method,
        JSON.print(data)
    )
```

## Handling Player Registration & Login

### Registration Example

```gdscript
func register(username: String, password: String) -> void:
    var data = {
        "username": username,
        "password": password
    }
    send_request("/register", HTTPClient.METHOD_POST, data)
```

### Login Example

On successful login, retrieve and store the JWT token for later requests:

```gdscript
func login(username: String, password: String) -> void:
    var data = {
        "username": username,
        "password": password
    }
    send_request("/login", HTTPClient.METHOD_POST, data)

func _on_request_completed(result, response_code, headers, body):
    if response_code == 200:
        var response = JSON.parse_string(body.get_string_from_utf8())
        global.TOKEN = response["token"]
        print("Login successful!")
```

## WebSocket Game Events

For real-time game events, connect to the server via WebSocket.

### Setup WebSocket Client

```gdscript
var ws = WebSocketClient.new()

func _ready():
    ws.connect_to_url(global.WS_URL + "?token=" + global.TOKEN)
    ws.connect("connection_established", self, "_on_ws_connected")
    ws.connect("data_received", self, "_on_data_received")

func _process(delta):
    ws.poll()

func _on_ws_connected():
    print("WebSocket connected to server.")

func _on_data_received(data):
    var message = JSON.parse_string(data.get_string_from_utf8())
    print("Received message: ", message)
```

## Submitting and Receiving Game Moves

Use the WebSocket client to send moves directly to the server:

```gdscript
func send_move(action: String):
    var move = {
        "action": action,
        "timestamp": Time.get_unix_time_from_system()
    }
    ws.send_text(JSON.print(move))
```

### Handling Game Events

```gdscript
func _on_data_received(data):
    var message = JSON.parse_string(data.get_string_from_utf8())
    match message["type"]:
        "move_response":
            handle_move_response(message)
        "game_state":
            update_game_state(message["state"])
        "error":
            handle_error(message["error"])

func handle_move_response(response):
    if response["valid"]:
        update_local_state(response["move"])
    else:
        print("Invalid move: ", response["reason"])
```

## Leaderboard Retrieval

You can retrieve leaderboard data via HTTP:

```gdscript
func get_leaderboard(type: String, timeframe: String = "all-time") -> void:
    var endpoint = "/leaderboard?type=" + type + "&timeframe=" + timeframe
    send_request(endpoint, HTTPClient.METHOD_GET)

func _on_leaderboard_received(result, response_code, headers, body):
    if response_code == 200:
        var leaderboard_data = JSON.parse_string(body.get_string_from_utf8())
        update_leaderboard_ui(leaderboard_data)
```

## Error Handling

Implement proper error handling for both HTTP and WebSocket connections:

```gdscript
func _on_request_error(error):
    print("HTTP Request Error: ", error)
    emit_signal("request_failed", error)

func _on_ws_error():
    print("WebSocket Error")
    # Implement reconnection logic
    yield(get_tree().create_timer(5.0), "timeout")
    ws.connect_to_url(global.WS_URL + "?token=" + global.TOKEN)
```

## Best Practices

1. **Token Management**
   - Store tokens securely
   - Implement token refresh logic
   - Handle token expiration gracefully

2. **Connection Management**
   - Implement WebSocket reconnection logic
   - Handle network interruptions
   - Maintain connection state

3. **State Synchronization**
   - Implement local state prediction
   - Handle server-client state reconciliation
   - Buffer inputs when appropriate

4. **Resource Cleanup**

```gdscript
func _exit_tree():
    if ws:
        ws.disconnect_from_host()
    if http_request:
        http_request.queue_free()
```

## Conclusion

You now have a foundational client setup in Godot Engine to interact with Drokkit. Customize and expand these examples as needed to build your game!

Remember to:

- Implement proper error handling
- Add loading states for network operations
- Consider implementing retry logic for failed requests
- Test your implementation thoroughly
- Handle edge cases and disconnections gracefully
