// validate.js
function validateAnuncioForm() {
    var title = document.getElementById('title').value;
    var description = document.getElementById('description').value;
    var errorMessages = document.getElementById('errorMessages');
  
    if (title == "" || description == "") {
      errorMessages.innerHTML = '<p style="color:red;">Debes ingresar título y descripción o Masha te enfría</p>';
      return false;
    }
  
    errorMessages.innerHTML = '';
    return true;
  }

  function clearAnuncioForm() {
    // Clear the error message
    var errorMessages = document.getElementById('errorMessages');
    errorMessages.innerHTML = '';

    // Clear the form fields
    var title = document.getElementById('title');
    var description = document.getElementById('description');
    title.value = '';
    description.value = '';
  }