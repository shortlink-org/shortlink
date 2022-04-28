/* eslint-disable */

module.exports = {
  content: {
    files: [
      './pages/**/*.{js,ts,jsx,tsx}',
      './components/**/*.{js,ts,jsx,tsx}',
    ],
    transform: {
      md: (content) => {
        return remark().process(content)
      },
    },
    extract: {
      md: (content) => {
        return content.match(/[^<>"'`\s]*/);
      },
    },
  },
  // important: '#__next',
  theme: {
    fontFamily: {
      display: ['Roboto Mono', 'Menlo', 'monospace'],
      body: ['Roboto Mono', 'Menlo', 'monospace'],
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    require("@tailwindcss/forms")({
      strategy: 'class', // only generate classes
    }),
    require('@tailwindcss/line-clamp'),
    require('@tailwindcss/aspect-ratio'),
  ],
}
