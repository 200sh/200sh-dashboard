document.addEventListener('DOMContentLoaded', function() {
    const deleteButton = document.getElementById('delete-button');
    const monitorId = deleteButton.getAttribute('data-monitor-id');
    
    document.getElementById('confirm-delete-button').addEventListener('click', function() {
        fetch(`/dashboard/monitors/${monitorId}`, {
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
    });

    document.getElementById('cancel-delete-button').addEventListener('click', function() {
        // No action needed as the modal will be closed by the data-dialog-close attribute
    });

    // Fake test data for latency graph
    const latencyData = {
        series: [{
            name: 'Latency (ms)',
            data: [65, 59, 80, 81, 56, 55, 40]
        }],
        chart: {
            height: 350,
            type: 'line'
        },
        xaxis: {
            categories: ['January', 'February', 'March', 'April', 'May', 'June', 'July']
        },
        yaxis: {
            title: {
                text: 'Latency (ms)'
            }
        }
    };

    const chart = new ApexCharts(document.querySelector("#latency-graph"), latencyData);
    chart.render();
});
