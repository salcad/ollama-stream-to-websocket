'use client';

import { useState, useEffect, useRef } from 'react';
import { createCustomMarked } from './createCustomMarked';

export default function Home() {
  const [model, setModel] = useState('llama3.1');
  const [prompt, setPrompt] = useState('Why is there something rather than nothing?');
  const [output, setOutput] = useState('');
  const [loading, setLoading] = useState(false);
  const EOT = '\u0004';
  const ws = useRef<WebSocket | null>(null);
  let markdownContent = '';
  const marked = createCustomMarked();

  useEffect(() => {
    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, []);

  const startWebSocket = () => {
    if (!ws.current || ws.current.readyState === WebSocket.CLOSED) {
      ws.current = new WebSocket('ws://localhost:8080/ws');
      ws.current.onmessage = function (event) {
        if (event.data === EOT) {
          markdownContent = '';
          setLoading(false);
        } else {
          markdownContent += event.data;
          console.log(markdownContent);
          setOutput(marked(markdownContent) as string);
          console.log(marked(markdownContent));
        }
      };
      ws.current.onclose = function () {
        console.log('WebSocket closed');
      };
    }
  };

  const sendPostRequest = async () => {
    startWebSocket();
    setOutput('');
    markdownContent = '';
    setLoading(true);

    const data = { model, prompt };

    try {
      const response = await fetch('http://localhost:8080/post', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
      const result = await response.text();
      console.log(result);
    } catch (error) {
      console.error('Error fetching data:', error);
    } 
  };

  return (
    <div className="container">
      <nav>
        <div className="nav-wrapper">
          <a href="#" className="brand-logo center">Ollama Chat</a>
        </div>
      </nav>

      <main className="main-content">
      <div className="container">
        <div className="input-field">
          <input
            type="text"
            id="model"
            name="model"
            value={model}
            onChange={(e) => setModel(e.target.value)}
            className="validate"
          />
          <label htmlFor="model" className="active">
            Model
          </label>
        </div>
        <div className="input-field">
          <textarea
            id="prompt"
            name="prompt"
            className="materialize-textarea"
            value={prompt}
            onChange={(e) => setPrompt(e.target.value)}
          ></textarea>
          <label htmlFor="prompt" className="active">
            Prompt
          </label>
        </div>
        <button
          id="sendButton"
          className="btn waves-effect waves-light btn-custom"
          onClick={sendPostRequest}
          disabled={loading}
        >
          Send prompt
        </button>
        <div
          id="output"
          className="card-panel grey lighten-3"
          style={{ whiteSpace: 'normal' }}
          dangerouslySetInnerHTML={{
            __html: `
              <style>
                h1, h2, h3, h4 {
                  margin: 0;
                }
                h1 { font-size: 1.5rem; }
                h2 { font-size: 1.3rem; }
                h3 { font-size: 1rem; }
                h4 { font-size: 0.8rem; }
              </style>
              ${output}
            `
          }}
        ></div>
        </div>
      </main>

      <footer className="page-footer">
        <div className="container"></div>
      </footer>
    </div>
  );
}
