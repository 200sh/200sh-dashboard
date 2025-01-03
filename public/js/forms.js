document.getElementById('cancel-button').addEventListener('click', function () {
    window.location.href = this.getAttribute('data-link');
});
