import { Preview } from '@storybook/react'
import { Provider } from 'react-wrap-balancer'

import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import '@fontsource/caveat'
import '@fontsource/material-icons'
import { CssVarsProvider } from '@mui/material/styles'
import InitColorSchemeScript from '@mui/material/InitColorSchemeScript'
import { ThemeProvider } from 'next-themes'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'

import '../src/theme/styles.css'
import { theme } from '../src/theme/theme'

const preview: Preview = {
  decorators: [
    (Story) => {
      return (
        // @ts-ignore
        (<ThemeProvider
          enableSystem
          attribute="class"
          defaultTheme={'light'}
        >
          <CssVarsProvider theme={theme}>
            <InitColorSchemeScript />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <Provider>
                <Story />
              </Provider>
            </LocalizationProvider>
          </CssVarsProvider>
        </ThemeProvider>)
      );
    },
  ],

  tags: ['autodocs']
}

export default preview
