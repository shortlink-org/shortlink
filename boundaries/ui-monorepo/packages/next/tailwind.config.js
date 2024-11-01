import { fontFamily } from 'tailwindcss/defaultTheme'
import typography from '@tailwindcss/typography'
import forms from '@tailwindcss/forms'
import aspectRatio from '@tailwindcss/aspect-ratio'
import containerQueries from '@tailwindcss/container-queries'

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
