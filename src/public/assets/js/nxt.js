// Commented out passcode related code
// function generatePasscode() {
//   let passcode = "";
//   const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
//   for (let i = 0; i < 20; i++) {
//     passcode += characters.charAt(
//       Math.floor(Math.random() * characters.length)
//     );
//     if ((i + 1) % 4 == 0 && i != 19) {
//       passcode += "-";
//     }
//   }
//   return passcode;
// }
// function getPasscode() {
//   if (localStorage.getItem("nxtup")) {
//     return localStorage.getItem("nxtup");
//   } else {
//     const defaultPasscode = generatePasscode();
//     Swal.fire({
//       title: "Enter your passcode",
//       input: "text",
//       inputValue: defaultPasscode,
//       inputPlaceholder: "Enter your passcode",
//       showCancelButton: true,
//       confirmButtonText: "I've written my passcode down",
//       showLoaderOnConfirm: true,
//       preConfirm: (passcode) => {
//         localStorage.setItem("nxtup", passcode);
//         document.getElementById("passcode").value = passcode;
//       },
//       allowOutsideClick: () => !Swal.isLoading(),
//     });
//   }
// }
function onSubmit() {
  const url = document.getElementById("url").value;
  const urlPattern = /^(http|https):\/\/[^ "]+$/;
  if (!urlPattern.test(url)) {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Please enter a valid URL.",
    });
    return;
  }
  const domain = window.location.hostname;
  if (url.includes(domain)) {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "You cannot shorten URLs that are the same as the domain we are currently on.",
    });
    return;
  }
  // const passcode = getPasscode();
  fetch("/api/link", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ url: url }),
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
        document.getElementById("card").style.display = "none";
        setTimeout(function () {
          document.getElementById("card").style.display = "block";
        }, 4600);
        document.getElementById("url").value = data.url;
        navigator.clipboard.writeText(data.url);
        Swal.fire({
          title: "Copied link to clipboard 🔗",
          showConfirmButton: false,
          timer: 4500,
        });
        fetch("/api/statistics/links")
          .then((response) => response.json())
          .then((data) => {
            document.getElementById("uploads").textContent =
              " (" + data.links + ")";
          });
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}

function checkURL() {
  const url = document.getElementById("url").value;
  const button = document.getElementById("actionButton");
  const domain = window.location.hostname;
  if (url.includes(domain)) {
    button.textContent = "Copy";
    button.onclick = function () {
      document.getElementById("card").style.display = "none";
      setTimeout(function () {
        document.getElementById("card").style.display = "block";
      }, 1600);
      navigator.clipboard.writeText(url);
      Swal.fire({
        title: "Copied link to clipboard 🔗",
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
  document.getElementById("domainTitle").textContent = domain;
  fetch("/api/statistics/links")
    .then((response) => response.json())
    .then((data) => {
      document.getElementById("uploads").textContent = " (" + data.links + ")";
    });
};
