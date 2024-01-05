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
        return response.json();  // Convert the response to JSON
    })
    .then(data => {
        const winningPercentage = 100 / Object.keys(data).length;
    
        // Update the background color of the buttons, add the percentage as text, remove the onclick event, and change the cursor style
        for (const [title, percentage] of Object.entries(data)) {
            const button = document.getElementById(title);
            const borderColor = percentage > winningPercentage ? 'rgb(0, 255, 0)' : 'rgb(255, 0, 0)';
            button.style.transition = 'all 0.5s ease-out';  // Add a transition to animate the changes
            button.style.border = `6px solid rgb(38, 41, 49)`;  // Add a border to the button
            button.style.boxShadow = `0 0 0 2px ${borderColor}`;  // Add a box shadow matching the border color
            button.style.padding = '10px';  // Add 10px of padding
            button.style.background = `linear-gradient(to right, #3b82f6 ${percentage}%, #F2F3F5 ${percentage}%)`;
            button.textContent = `${button.textContent} (${percentage.toFixed(2)}%)`;  // Append the percentage to the current text
            button.removeAttribute('onclick');  // Remove the onclick event
            button.style.cursor = 'default';  // Change the cursor style
            button.style.color = 'black';  // Set the text color to black
    
            // Add a hover effect
            button.onmouseover = function() {
                this.style.transform = 'scale(1.05)';
            }
            button.onmouseout = function() {
                this.style.transform = 'scale(1)';
            }
        }
    })
    .catch((error) => {
      console.error('Error:', error);
    });
}