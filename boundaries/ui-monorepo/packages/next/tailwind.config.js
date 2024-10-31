const { fontFamily } = require('tailwindcss/defaultTheme')
const forms = require('@tailwindcss/forms')
const aspectRatio = require('@tailwindcss/aspect-ratio')
const containerQueries = require('@tailwindcss/container-queries')

/** @type {import('tailwindcss').Config} */
export default {
  mode: 'jit',
  darkMode: 'selector',
  corePlugins: {
    preflight: false,
  },
  content: {
    files: ['./app/**/*.{js,ts,jsx,tsx,mdx}', './components/**/*.{js,ts,jsx,tsx}'],
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
    container: {
      center: true,
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
  plugins: [typography, forms, aspectRatio, containerQueries],
}
