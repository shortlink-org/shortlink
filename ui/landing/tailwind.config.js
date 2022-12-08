/* eslint-disable */

module.exports = {
  mode: 'jit',
  darkMode: 'class',
  content: {
    files: [
      './pages/**/*.{js,ts,jsx,tsx}',
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
  important: '#__next',
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
    typography: ['dark'],
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/line-clamp'),
    require('@tailwindcss/aspect-ratio'),
  ],
}
