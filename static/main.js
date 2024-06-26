const themeToggler = document.getElementById("theme-toggler");
const body = document.querySelector("body");

themeToggler.addEventListener("click", () => {
  const currentTheme = body.getAttribute("data-theme");

  if (currentTheme === "light") {
    body.setAttribute("data-theme", "dark");
  } else {
    body.setAttribute("data-theme", "light");
  }
});
