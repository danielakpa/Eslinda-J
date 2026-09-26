// Toggles the password field between hidden and visible text
document.getElementById("toggle-password").addEventListener("click", function () {
  const passwordInput = document.getElementById("password");
  if (passwordInput.type === "password") {
    passwordInput.type = "text";
  } else {
    passwordInput.type = "password";
    this.textContent = "👁";
  }
});

document.getElementById("signup-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const response = await fetch("/customer/signup", {
    method: "POST",
    body: formData,
  });

  const message = await response.text();
  document.getElementById("signup-message").innerText = message;

  if (response.ok) {
    window.location.href = "/customer-login";
  }
});