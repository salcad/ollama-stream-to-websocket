# Demo UI Chat to LLM

This is a simple web application that allows users to interact with a Large Language Model (LLM) via a chat interface. The application uses `express` for the backend server, `WebSocket` for real-time communication, and `marked.js` with `highlight.js` for rendering and syntax highlighting of Markdown content.

## Features

- Real-time chat interface with LLM
- Markdown support with syntax highlighting for code blocks
- Responsive design using Materialize CSS

## Prerequisites

- Node.js (v20.11 or higher)
- npm (v10.2.4 or higher)

## Installation

### Clone the repository:

```bash
git clone https://github.com/salcad/ollama-stream-to-websocket.git
cd cd ollama-stream-to-websocket/client
```

### Install the dependencies:

```bash
npm install
```

## Usage

### Start the server:

```bash
npm start
```

Open your browser and navigate to http://localhost:3000.