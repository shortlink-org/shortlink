import { red, grey } from '@mui/material/colors'
import { experimental_extendTheme as extendTheme } from '@mui/material/styles'

export const theme = extendTheme({
  colorSchemes: {
    light: {
      palette: {
        primary: {
          main: '#556cd6',
        },
        secondary: {
          main: red.A400,
        },
        error: {
          main: red.A400,
        },
        text: {
          primary: grey.A700,
          secondary: grey.A400,
        },
      },
    },
    dark: {
      palette: {
        primary: {
          main: '#556cd6',
          light: '#f1f5f9',
        },
        secondary: {
          main: '#94a3b8',
        },
        error: {
          main: red.A400,
        },
        background: {
          default: grey.A100,
        },
        text: {
          primary: grey[100],
          secondary: grey[200],
        },
      },
    },
  },
})

export default theme
