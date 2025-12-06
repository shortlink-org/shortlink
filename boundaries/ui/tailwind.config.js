import typography from '@tailwindcss/typography'
import forms from '@tailwindcss/forms'
import aspectRatio from '@tailwindcss/aspect-ratio'
import containerQueries from '@tailwindcss/container-queries'

/** @type {import('tailwindcss').Config} */
export default {
  // Enable dark mode via class strategy (next-themes uses class="dark" on <html>)
  darkMode: 'class',

  corePlugins: {
    preflight: false, // keep if you intentionally disabled base resets
  },

  content: [
    './app/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './pages/**/*.{js,ts,jsx,tsx,mdx}',   // add if you still use Pages Router anywhere
    './src/**/*.{js,ts,jsx,tsx,mdx}',     // add if you keep code in /src
    './node_modules/@shortlink-org/ui-kit/**/*.{js,ts,jsx,tsx}', // if you render this package
  ],

  theme: {
    container: { center: true },

    fontFamily: {
      display: ['Roboto Mono', 'Menlo', 'monospace'],
      body: ['Roboto Mono', 'Menlo', 'monospace'],
      inter: ['Inter', 'sans-serif'],
      caveat: ['Caveat', 'cursive'],
    },

    extend: {
      fontFamily: {
        sans: ['var(--font-inter)', 'system-ui', 'sans-serif'],
      },

      // Modern CSS Variables - override default colors
      colors: {
        gray: {
          50: 'rgb(var(--color-gray-50) / <alpha-value>)',
          100: 'rgb(var(--color-gray-100) / <alpha-value>)',
          200: 'rgb(var(--color-gray-200) / <alpha-value>)',
          400: 'rgb(var(--color-text-secondary) / <alpha-value>)',
          500: 'rgb(var(--color-text-secondary) / <alpha-value>)',
          800: 'rgb(var(--color-gray-800) / <alpha-value>)',
          900: 'rgb(var(--color-gray-900) / <alpha-value>)',
        },
        indigo: {
          100: 'rgb(var(--color-indigo-100) / <alpha-value>)',
          200: 'rgb(var(--color-indigo-200) / <alpha-value>)',
          300: 'rgb(var(--color-indigo-300) / <alpha-value>)',
          500: 'rgb(var(--color-indigo-500) / <alpha-value>)',
          600: 'rgb(var(--color-indigo-600) / <alpha-value>)',
          700: 'rgb(var(--color-indigo-700) / <alpha-value>)',
        },
        white: 'rgb(255 255 255 / <alpha-value>)',
        black: 'rgb(0 0 0 / <alpha-value>)',
      },

      // Tailwind Typography customization
      typography: (theme) => ({
        invert: {
          css: {
            color: theme('colors.slate.200'),
            a: { color: theme('colors.indigo.300') },
            strong: { color: theme('colors.slate.100') },
            h1: { color: theme('colors.slate.100') },
            h2: { color: theme('colors.slate.100') },
            h3: { color: theme('colors.slate.100') },
            hr: { borderColor: theme('colors.slate.700') },
            code: { color: theme('colors.slate.100') },
            'blockquote p': { color: theme('colors.slate.200') },
          },
        },
      }),
    },
  },

  plugins: [typography, forms, aspectRatio, containerQueries],
}
