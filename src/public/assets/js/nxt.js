document.addEventListener("DOMContentLoaded", (event) => {
  function generatePasscode() {
    var passcode = "";
    var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    for (var i = 0; i < 20; i++) {
      passcode += characters.charAt(
        Math.floor(Math.random() * characters.length)
      );
      if ((i + 1) % 4 == 0 && i != 19) {
        passcode += "-";
      }
    }
    return passcode;
  }
  function getPasscode() {
    if (localStorage.getItem("nxtup")) {
      return localStorage.getItem("nxtup");
    } else {
      const defaultPasscode = generatePasscode();
      Swal.fire({
        title: "Enter your passcode",
        input: "text",
        inputValue: defaultPasscode,
        inputPlaceholder: "Enter your passcode",
        showCancelButton: true,
        confirmButtonText: "I've written my passcode down",
        showLoaderOnConfirm: true,
        preConfirm: (passcode) => {
          localStorage.setItem("nxtup", passcode);
          document.getElementById("passcode").value = passcode;
        },
        allowOutsideClick: () => !Swal.isLoading(),
      });
    }
  }
  function onSubmit() {
    var url = document.getElementById("url").value;
    if (url.includes(domain)) {
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "You cannot shorten URLs that are the same as the domain we are currently on.",
      });
      return;
    }
    var passcode = getPasscode();
    fetch("/api/link", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url: url, passcode: passcode }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          Swal.fire({
            icon: "error",
            title: "Oops...",
            text: data.error,
          });
        } else {
          document.getElementById("url").value = data.url;
          navigator.clipboard.writeText(data.url);
          Swal.fire({
            icon: "success",
            title: "URL has been copied to clipboard",
            showConfirmButton: false,
            timer: 4500,
          });
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
  function checkURL() {
    var url = document.getElementById("url").value;
    var button = document.getElementById("actionButton");
    if (url.includes(domain)) {
      button.textContent = "Copy";
      button.onclick = function () {
        navigator.clipboard.writeText(url);
        Swal.fire({
          icon: "success",
          title: "URL has been copied to clipboard",
          showConfirmButton: false,
          timer: 1500,
        });
      };
    } else {
      button.textContent = "Shorten";
      button.onclick = onSubmit;
    }
  }
  window.onload = function () {
    document.getElementById("url").addEventListener("input", checkURL);
    fetch("/api/statistics/links")
      .then((response) => response.json())
      .then((data) => {
        document.getElementById("domainTitle").textContent =
          domain + " (" + data.links + ")";
      });
  };
});
