// Highlight.js initialization for code blocks
// Make sure to include Highlight.js library in your HTML file

document.addEventListener('DOMContentLoaded', (event) => {
  document.querySelectorAll('pre code').forEach((el) => {
    hljs.highlightElement(el);
  });
});

