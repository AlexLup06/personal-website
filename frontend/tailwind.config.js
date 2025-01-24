const { default: daisyui } = require("daisyui");
const { plugin } = require("postcss");

module.exports = {
  content: ["./src/views/**/*.templ"],
  theme: {
    extend: {
    },
  },
  plugins: [
    require('daisyui'),
  ],
  corePlugins: {
    preflight: true,
  },
  mode: "jit",
  daisyui: {
    themes: [],
  },
};
