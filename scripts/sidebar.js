function sidebar() {
  fetch("/templates/components/sidebar.html")
    .then((response) => response.text())
    .then((data) => {
      document.getElementById("sidebar").innerHTML = data;
    });
}
