const { default: daisyui } = require("daisyui");
const { plugin } = require("postcss");

module.exports = {
  content: ["./src/views/**/*.templ"],
  mode: "jit",
  theme: {
    extend: {
      spacing: {
        "nav-height": "96px"
      },
      screens: {
        'xs': '480px',
      },
    },
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      'white': '#ffffff',
      'black': '#000000',
      "main": {
        50: "#ecfbf7",
        100: "#c5f2e6",
        200: "#9ee9d6",
        300: "#77e0c6",
        400: "#50d7b5",
        500: "#3CD3AD",
        600: "#30a98a",
        700: "#247f68",
        800: "#185445",
        900: "#061511",
      },
      "second": {
        25: "#eee8e9",
        50: "#ddd1d4",
        100: "#ccbabe",
        200: "#aa8c93",
        300: "#875d68",
        400: "#652f3d",
        500: "#541827",
        600: "#43131f",
        700: "#320e17",
        800: "#220a10",
        900: "#110508",
      },
      "grey": {
        50: "#f9fafb",
        100: "#f3f4f6",
        200: "#e5e7eb",
        300: "#d1d5db",
        400: "#9ca3af",
        500: "#6b7280",
        600: "#4b5563",
        700: "#374151",
        800: "#1f2937",
        900: "#111827",
        950: "#030712"
      }
    },
  },
  plugins: [
    require('daisyui'),
    require('@tailwindcss/container-queries'),
  ],
  corePlugins: {
    preflight: true,
  },
  daisyui: {
    themes: [],
  },
};
