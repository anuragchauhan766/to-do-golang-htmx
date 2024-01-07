/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        primary: "#FFFBF5",
        secondary: "#F7EFE5",
        ctc: "#7743DB",
        "ctc-light": "#BC6FF1",
      },
      fontFamily: {
        montserrat: ["Montserrat", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
