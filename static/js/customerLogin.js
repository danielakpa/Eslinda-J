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

document.getElementById("login-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const response = await fetch("/customer/login", {
    method: "POST",
    body: formData,
  });

  if (response.ok) {
    document.getElementById("login-form").style.display = "none";
    document.getElementById("success-box").style.display = "block";

    setTimeout(function () {
      window.location.href = "/";
    }, 1500);
  } else {
    const message = await response.text();
    document.getElementById("login-message").innerText = message;
  }
});