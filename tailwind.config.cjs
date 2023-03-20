/** @type {import('tailwindcss').Config} */
module.exports = {
  // support toggling dark mode manually instead of 
  // relying on the operating system preference
  darkMode: "class",
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}