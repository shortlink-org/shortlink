import { red, grey } from '@mui/material/colors'
import { createTheme } from '@mui/material/styles'

// Create a theme instance.
export const lightTheme = createTheme({
  palette: {
    mode: 'light',
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
})

export const darkTheme = createTheme({
  palette: {
    mode: 'dark',
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
})
