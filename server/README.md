# Ollama Stream to WebSocket

This project serves as a wrapper for the HTTP POST stream of the Ollama API, showcasing a Golang server that handles long-running tasks by streaming responses to clients via WebSocket. The use case? It's a bit of a joke because, let's be honest, do we really need another chat app LLM? Or maybe it's just my excuse to have some fun with Golang and WebSocket.

## Installation

### 1. **Clone the repository**:

```bash
git clone https://github.com/salcad/ollama-stream-to-websocket.git
cd ollama-stream-to-websocket/server/src/main
```

### 2. **Install dependencies**:

```bash
go mod tidy
```

### 3. **Install Ollama**:

Ollama allows us to run open-source Large language models (LLMs) locally on our system. If you don't have Ollama installed on your system and don't know how to use it, I suggest you go through the [README.md on Ollama's GitHub](https://github.com/ollama/ollama).

#### On macOS you can install Ollama using Homebrew

If you don't have Homebrew installed, please install it first by following the instructions at [Homebrew Installation Guide](https://docs.brew.sh/Installation)

#### Install Ollama:

```sh
brew install ollama
```
##### Verify Installation

Open Terminal and run:

```sh
ollama
```

This command should display available commands and verify that Ollama is installed correctly.

##### Download and Run a Model

To list available models:

```sh
ollama list
```

To download a model:

```bash
ollama pull llama3.1:latest
```

To run the model:

```bash
ollama run llama3.1:latest
```

## Configuration

Update the `config.json` file with your desired configuration:

```go
{
    "ollamaAPIUrl": "http://localhost:11434/api/generate",
    "allowedOrigin": "http://localhost:3000"
}
```

## Usage

### Run the server:

```bash
go run main.go
```

### Run the unit tests:

```bash
go test -v -count=1 ../test
```

### Client interaction:

- Setup and run node.js client
- Open your web browser and navigate to http://localhost:3000.
- Enter the model and prompt values.
- Click "SEND PROMPT" to start chat with Ollama LLM.
