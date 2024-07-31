# Ollama Stream to WebSocket

<br><img src="./image/ollama_ui_chat_demo_compr.png" alt="Web UI Chat Demo" style="width:50%;"><br>

Welcome to the Ollama Stream to WebSocket project. This repository showcases a Golang server that handles long-running tasks by streaming responses to clients via WebSocket. This project is designed to have fun with Golang and WebSocket while also serving as a practical wrapper for the HTTP POST stream of the Ollama API.

## Features

- **HTTP POST with WebSocket Response**: Initiate long-running tasks with HTTP POST requests and receive real-time updates via WebSocket.
- **Separation of Concerns**: Cleanly organized architecture with distinct controller, service, configuration, lib, and model, making it easy to extend and maintain.
- **Comprehensive Unit Tests**: Thorough unit tests included for handling WebSocket scenarios.
- **Bonus Sample UI Chat**: A demo web client that showcases a chat interface with markdown and syntax color highlighting.

## More Details

For detailed instructions and more information, please refer to the README.md files in the respective folders:

[Server README](https://github.com/salcad/ollama-stream-to-websocket/blob/main/server/README.md)

[Client README](https://github.com/salcad/ollama-stream-to-websocket/blob/main/client/README.md)
