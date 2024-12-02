import { Preview } from '@storybook/react'
import { Provider } from 'react-wrap-balancer'

import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import '@fontsource/caveat'
import '@fontsource/material-icons'
import { ThemeProvider } from '@mui/material/styles'
import InitColorSchemeScript from '@mui/material/InitColorSchemeScript'
import { ThemeProvider as ThemeProviderNext } from 'next-themes'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'

import '../src/theme/styles.css'
import { theme } from '../src/theme/theme'

const preview: Preview = {
  decorators: [
    (Story) => {
      return (
        // @ts-ignore
        (<ThemeProviderNext
          enableSystem
          attribute="class"
          defaultTheme={'light'}
        >
          <ThemeProvider theme={theme}>
            <InitColorSchemeScript />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <Provider>
                <Story />
              </Provider>
            </LocalizationProvider>
          </ThemeProvider>
        </ThemeProviderNext>)
      );
    },
  ],

  tags: []
}

export default preview
