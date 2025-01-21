const plugin = require('tailwindcss/plugin')


module.exports = {
  content: ["./src/views/**/*.templ"],
  theme: {
    extend: {
      textShadow: {
        DEFAULT: '2px 8px 6px rgba(0,0,0,0.2),0px -5px 35px rgba(255,255,255,0.3)',
      },
    },
  },
  plugins: [plugin(function ({ matchUtilities, theme }) {
    matchUtilities(
      {
        'text-shadow': (value) => ({
          textShadow: value,
        }),
      },
      { values: theme('textShadow') }
    )
  }),],
  corePlugins: {
    preflight: true,
  },
  mode: "jit",
};
