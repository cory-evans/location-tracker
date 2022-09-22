/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors');
const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Nunito', ...defaultTheme.fontFamily.sans],
      },
      colors: {
        primary: colors.blue,
        secondary: colors.orange,
        success: colors.green,
        danger: colors.red,
        warning: colors.yellow,
        info: colors.teal,
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}
