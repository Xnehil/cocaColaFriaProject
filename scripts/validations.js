// validate.js
function validateAnuncioForm() {
    var title = document.getElementById('title').value;
    var description = document.getElementById('description').value;
    var errorMessages = document.getElementById('errorMessages');
  
    if (title == "" || description == "") {
      errorMessages.innerHTML = '<p style="color:red;">Debes ingresar título y descripción o Masha te enfría</p>';
      return false;
    }

    //Check max length
    if (title.length > 30) {
      errorMessages.innerHTML = '<p style="color:red;">El título debe tener menos de 30 caracteres</p>';
      return false;
    }
    if (description.length > 250) {
      errorMessages.innerHTML = '<p style="color:red;">La descripción debe tener menos de 250 caracteres, no te tumbes la BD por favor</p>';
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