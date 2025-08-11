import typography from '@tailwindcss/typography'
import forms from '@tailwindcss/forms'
import aspectRatio from '@tailwindcss/aspect-ratio'
import containerQueries from '@tailwindcss/container-queries'

/** @type {import('tailwindcss').Config} */
export default {
  // Use 'class' or a custom selector:
  // darkMode: 'class',
  // or, if you toggle data-theme:
  darkMode: ['class', '[data-theme="dark"]'],

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
        // set global defaults you'll actually use
        sans: ['var(--font-inter)', 'system-ui', 'sans-serif'],
      },

      // Tailwind Typography customization
      typography: (theme) => ({
        // dark mode styles via `prose-invert`
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
