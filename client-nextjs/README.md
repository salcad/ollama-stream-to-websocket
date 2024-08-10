# Demo UI Chat to LLM

This is a simple web application that allows users to interact with a Large Language Model (LLM) via a chat interface. The application uses Next.js, WebSocket for real-time communication, and marked.js with highlight.js for rendering and syntax highlighting of Markdown content.

## Features

- Real-time chat interface with LLM
- Markdown support with syntax highlighting for code blocks
- Responsive design using Materialize CSS

## Installation

## Prerequisites

- Node.js (v20.11 or higher)
- npm (v10.2.4 or higher)

### Clone the repository:

```bash
git clone https://github.com/salcad/ollama-stream-to-websocket.git
cd ollama-stream-to-websocket/client-nextjs
```

## [Next.js](https://nextjs.org/) 
#### Project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

1. If you want you can create a `.env.local` file in the root of the project and update the variables to match your needs.

    ```bash
    NEXT_PUBLIC_WEBSOCKET_URL=ws://localhost:8080/ws
    NEXT_PUBLIC_POST_URL=http://localhost:8080/post
    ```

2. Run the development server

    ```bash
    npm run dev
    ```

3. Run test

    ```bash
    npm test
    ```

4. To build and run the project in production mode:

    ```bash
    npm run build
    npm start
    ```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.