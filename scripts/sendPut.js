function sendPut(votacionNombre, title) {
    fetch('/api/putCuentaVotacion', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'fantadepina' 
        },
        body: JSON.stringify({ votacionNombre: votacionNombre, title: title })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Server response was not ok.');
        }
    })
    .catch((error) => {
      console.error('Error:', error);
    });
}