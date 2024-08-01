document.addEventListener("DOMContentLoaded", function() {
    if (typeof hljs !== 'undefined') {
        // Configure marked to use highlight.js for code blocks
        marked.setOptions({
            highlight: function(code, lang) {
                if (lang && hljs.getLanguage(lang)) {
                    return hljs.highlight(code, { language: lang }).value;
                } else {
                    return hljs.highlightAuto(code).value;
                }
            }
        });
    } else {
        console.error("highlight.js not loaded");
    }
});

var ws;
const EOT = "\u0004"; 
let markdownContent = "";

function startWebSocket() {
    if (!ws || ws.readyState === WebSocket.CLOSED) {
        ws = new WebSocket("ws://localhost:8080/ws");
        ws.onmessage = function(event) {
            var output = document.getElementById("output");
            if (event.data === EOT) {
                console.log("Process completed");
                output.innerHTML = marked.parse(markdownContent); // Convert Markdown to HTML
                markdownContent = ""; 
                document.getElementById("sendButton").disabled = false; 
                // Manually highlight code blocks after rendering
                document.querySelectorAll('pre code').forEach((block) => {
                    hljs.highlightElement(block);
                });
            } else {
                markdownContent += event.data; // Accumulate the Markdown content
                output.innerHTML += event.data.replace(/\\n/g, "<br>");
            }
        };
        ws.onclose = function() {
            console.log("WebSocket closed");
        };
    }
}

function sendPostRequest() {
    startWebSocket(); 

    document.getElementById("output").innerHTML = "";
    markdownContent = ""; 

    document.getElementById("sendButton").disabled = true;

    const model = document.getElementById("model").value;
    const prompt = document.getElementById("prompt").value;
    const data = { model, prompt };

    fetch('http://localhost:8080/post', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => response.text())
      .then(data => {
          console.log(data);
      });
}

