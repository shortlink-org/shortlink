import { useState } from 'react'
import { Preview } from '@storybook/react'
import { Provider } from 'react-wrap-balancer'

import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import '@fontsource/caveat'
import '@fontsource/material-icons'
import { ThemeProvider } from '@mui/material/styles'
import { ThemeProvider as TailWindProvider } from 'next-themes'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'

import '../src/theme/styles.css'
import { ColorModeContext, darkTheme, lightTheme } from '../src'

const preview: Preview = {
  decorators: [
    (Story) => {
      const [darkMode, setDarkMode] = useState(false)
      const theme = darkMode ? darkTheme : lightTheme

      return (
        <TailWindProvider enableSystem attribute="class">
          <ThemeProvider theme={theme}>
            <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <Provider>
                  <Story />
                </Provider>
              </LocalizationProvider>
            </ColorModeContext.Provider>
          </ThemeProvider>
        </TailWindProvider>
      )
    },
  ],
}

export default preview
