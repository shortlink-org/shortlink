const { fontFamily } = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  darkMode: 'selector',
  content: {
    files: [
      './app/**/*.{js,ts,jsx,tsx,mdx}',
      './components/**/*.{js,ts,jsx,tsx}',
    ],
    options: {
      safelist: ['dark'], // specific classes
    },
  },
  theme: {
    fontFamily: {
      display: ['Roboto Mono', 'Menlo', 'monospace'],
      body: ['Roboto Mono', 'Menlo', 'monospace'],
      inter: ['Inter', 'sans-serif'],
      caveat: ['Caveat', 'cursive'],
    },
    extend: {
      typography: () => ({
        dark: {
          css: {
            color: 'white',
          },
        },
        fontFamily: {
          sans: ['var(--font-inter)', ...fontFamily.sans],
        },
      }),
    },
  },
  variants: {
    typography: ['light', 'dark'],
  },
  plugins: [
    // eslint-disable-next-line global-require
    require('@tailwindcss/typography'),
    // eslint-disable-next-line global-require
    require('@tailwindcss/forms'),
    // eslint-disable-next-line global-require
    require('@tailwindcss/aspect-ratio'),
    // eslint-disable-next-line global-require
    require('@tailwindcss/container-queries'),
  ],
}
