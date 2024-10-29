
# Drokkit C# Game Client Development Guide

This guide provides instructions on how to develop a C# client that interacts with the Drokkit game server. This setup assumes the Drokkit server is running and accessible through the configured endpoints. The Drokkit server uses JWT for authentication, REST for game actions, and WebSockets for real-time interactions.

---

## Table of Contents

- [Drokkit C# Game Client Development Guide](#drokkit-c-game-client-development-guide)
  - [Table of Contents](#table-of-contents)
  - [Setting Up the C# Project](#setting-up-the-c-project)
  - [Installing Required Libraries](#installing-required-libraries)
  - [Authenticating with the Server](#authenticating-with-the-server)
  - [Interacting with the REST API](#interacting-with-the-rest-api)
  - [Setting Up WebSocket for Real-Time Play](#setting-up-websocket-for-real-time-play)
  - [Handling Game Events](#handling-game-events)
  - [Example Usage](#example-usage)
  - [Conclusion](#conclusion)

---

## Setting Up the C# Project

1. **Create a New C# Console Application**  
   Open Visual Studio or your preferred C# editor and create a new .NET Console Application project. Name it `DrokkitClient`.

2. **Configure Environment Variables**  
   Create a `.env` file or add environment variables directly to your project configuration. You’ll need the following variables:
   - `API_BASE_URL`: Base URL for REST endpoints (e.g., `http://localhost:8080/api`).
   - `WS_BASE_URL`: WebSocket URL for real-time interaction (e.g., `ws://localhost:8080/ws/play`).

---

## Installing Required Libraries

Install the following libraries for HTTP requests, JWT handling, and WebSocket support:

```bash
dotnet add package System.Net.Http
dotnet add package JWT
dotnet add package WebSocketSharp
```

---

## Authenticating with the Server

To interact with the game server, your client must first authenticate and obtain a JWT token.

```csharp
using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

public class Authentication
{
    private static readonly HttpClient client = new HttpClient();

    public static async Task<string> LoginAsync(string username, string password)
    {
        var loginData = new { username, password };
        var content = new StringContent(JsonSerializer.Serialize(loginData), Encoding.UTF8, "application/json");

        HttpResponseMessage response = await client.PostAsync($"{Environment.GetEnvironmentVariable("API_BASE_URL")}/login", content);
        
        if (response.IsSuccessStatusCode)
        {
            var responseBody = await response.Content.ReadAsStringAsync();
            var jsonDoc = JsonDocument.Parse(responseBody);
            return jsonDoc.RootElement.GetProperty("token").GetString();
        }
        else
        {
            throw new Exception("Failed to authenticate");
        }
    }
}
```

---

## Interacting with the REST API

With the JWT token, your client can make authenticated requests to the server. Here’s an example of creating a new match:

```csharp
public static async Task CreateMatch(string token)
{
    client.DefaultRequestHeaders.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", token);
    var content = new StringContent("{}", Encoding.UTF8, "application/json");

    HttpResponseMessage response = await client.PostAsync($"{Environment.GetEnvironmentVariable("API_BASE_URL")}/match", content);

    if (response.IsSuccessStatusCode)
    {
        Console.WriteLine("Match created successfully");
    }
    else
    {
        Console.WriteLine("Failed to create match");
    }
}
```

---

## Setting Up WebSocket for Real-Time Play

For real-time interactions (e.g., moves), establish a WebSocket connection.

```csharp
using WebSocketSharp;

public class GameWebSocket
{
    private WebSocket ws;

    public GameWebSocket(string token)
    {
        ws = new WebSocket($"{Environment.GetEnvironmentVariable("WS_BASE_URL")}?token={token}");

        ws.OnMessage += (sender, e) =>
        {
            Console.WriteLine($"Message from server: {e.Data}");
        };

        ws.OnOpen += (sender, e) =>
        {
            Console.WriteLine("Connected to WebSocket!");
        };

        ws.OnClose += (sender, e) =>
        {
            Console.WriteLine("Disconnected from WebSocket.");
        };
    }

    public void Connect() => ws.Connect();
    public void Disconnect() => ws.Close();
    public void SendMessage(string message) => ws.Send(message);
}
```

---

## Handling Game Events

Using the WebSocket, handle events such as player moves or game updates:

```csharp
public void SendMove(int matchId, int playerId, string action)
{
    var moveData = new { matchId, playerId, action };
    string moveJson = JsonSerializer.Serialize(moveData);
    ws.Send(moveJson);
}
```

---

## Example Usage

```csharp
public class Program
{
    public static async Task Main(string[] args)
    {
        string token = await Authentication.LoginAsync("player1", "password123");

        var gameSocket = new GameWebSocket(token);
        gameSocket.Connect();

        Console.WriteLine("Press any key to send a move...");
        Console.ReadKey();
        gameSocket.SendMove(1, 123, "move-forward");

        Console.ReadKey();
        gameSocket.Disconnect();
    }
}
```

This example demonstrates logging in, connecting to the WebSocket, and sending a move. Customize it as needed for other game actions.

---

## Conclusion

This guide provides a C# client setup for interacting with Drokkit’s game server. Follow the steps for secure authentication, REST API calls, and real-time WebSocket communication for a smooth player experience.
