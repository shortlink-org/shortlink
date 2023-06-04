/* eslint-disable */

const { fontFamily } = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} \*/
module.exports = {
  darkMode: 'class',
  content: {
    files: [
      './app/**/*.{js,ts,jsx,tsx,mdx}',
      './components/**/*.{js,ts,jsx,tsx}',
      './stories/**/*.{js,ts,jsx,tsx}',
      '../../next/components/**/*.{js,ts,jsx,tsx}',
    ],
    transform: {
      md: (content) => {
        return remark().process(content)
      },
    },
    extract: {
      md: (content) => {
        return content.match(/[^<>"'`\s]*/)
      },
    },
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
        fontFamily: {
          sans: ['var(--font-inter)', ...fontFamily.sans],
        },
      }),
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
  ],
}
