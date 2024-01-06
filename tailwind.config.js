/** @type {import('tailwindcss').Config} */

const defaultTheme = require("tailwindcss/defaultTheme");
module.exports = {
  content: ["view/**/*.templ", "./node_modules/flowbite/**/*.js"],
  darkMode: "class", // Doc: https://tailwindcss.com/docs/dark-mode#supporting-system-preference-and-manual-selection or https://flowbite.com/docs/customize/dark-mode/
  theme: {
    extend: {
      fontFamily: {
        sans: ["InterVariable", "Inter", ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("flowbite/plugin"),
  ],
};
