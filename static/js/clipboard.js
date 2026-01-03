document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.code-block').forEach(block => {
        const btn = document.createElement('button');
        btn.className = 'clipboard-btn';
        btn.innerHTML = '<img src="/icons/clipboard-color/clipboard-24.png" alt="Copy">';
        btn.onclick = function() {
            navigator.clipboard.writeText(block.querySelector('code').innerText);
            btn.innerHTML = '✓';
            setTimeout(() => btn.innerHTML = '<img src="/icons/clipboard-color/clipboard-24.png" alt="Copy">', 2000);
        };
        block.appendChild(btn);
    });
});

