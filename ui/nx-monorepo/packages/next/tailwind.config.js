/* eslint-disable */

module.exports = {
  mode: 'jit',
  darkMode: 'class',
  content: {
    files: [
      './pages/**/*.{js,ts,jsx,tsx}',
      './components/**/*.{js,ts,jsx,tsx}',
      './stories/**/*.{js,ts,jsx,tsx}',
    ],
    options: {
      safelist: ['dark'], // specific classes
    },
  },
  theme: {
    fontFamily: {
      display: ['Roboto Mono', 'Menlo', 'monospace'],
      body: ['Roboto Mono', 'Menlo', 'monospace'],
    },
    extend: {
      typography: (theme) => ({
        dark: {
          css: {
            color: 'white',
          },
        },
      }),
    },
  },
  variants: {
    typography: ['light', 'dark'],
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
  ],
}
