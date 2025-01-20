document.addEventListener('DOMContentLoaded', function() {
    const deleteButton = document.getElementById('delete-button');
    const monitorId = deleteButton.getAttribute('data-monitor-id');

    deleteButton.addEventListener('click', function() {
        const confirmation = confirm('Are you sure you want to delete this monitor?');
        if (confirmation) {
            fetch(`/dashboard/monitors/${monitorId}/delete`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/dashboard/monitors';
                } else {
                    alert('Failed to delete the monitor.');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('An error occurred while deleting the monitor.');
            });
        }
    });

    // Fake test data for latency graph
    const latencyData = {
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
        datasets: [{
            label: 'Latency (ms)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderColor: 'rgba(75, 192, 192, 1)',
            borderWidth: 1,
            data: [65, 59, 80, 81, 56, 55, 40]
        }]
    };

    const ctx = document.getElementById('latency-graph').getContext('2d');
    new Chart(ctx, {
        type: 'line',
        data: latencyData,
        options: {
            responsive: true,
            scales: {
                x: {
                    display: true,
                    title: {
                        display: true,
                        text: 'Month'
                    }
                },
                y: {
                    display: true,
                    title: {
                        display: true,
                        text: 'Latency (ms)'
                    }
                }
            }
        }
    });
});
