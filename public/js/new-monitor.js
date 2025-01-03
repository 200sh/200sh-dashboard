window.addEventListener('DOMContentLoaded', function() {
    document.getElementById('monitor-url').addEventListener('paste', function(event) {
        event.preventDefault();

        let content = this.value;
        const pastedText = (event.originalEvent || event).clipboardData.getData('text/plain');

        // Check if both content and pastedText contain https or http. Keep the one from the pasted text
        if (content.includes('https://') && pastedText.includes('http://')) {
            content = pastedText;
        } else if (content.includes('http://') && pastedText.includes('https://')) {
            content = pastedText;
        } else {
            content += pastedText;
        }

        document.getElementById('monitor-url').value = content;
    });
});