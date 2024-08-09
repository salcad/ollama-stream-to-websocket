import { marked, Renderer } from 'marked';
import hljs from 'highlight.js';

export function createCustomMarked() {
  const customRenderer = new Renderer();

  customRenderer.code = function ({ text, lang }) {
    const language = lang && hljs.getLanguage(lang) ? lang : 'plaintext';
    return `<pre><code class="hljs ${language}">${hljs.highlight(text, { language }).value}</code></pre>`;
  };

  marked.setOptions({
    renderer: customRenderer,
  });

  return marked;
}
