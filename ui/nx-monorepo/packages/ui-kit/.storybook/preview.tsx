import { useState } from 'react'
import { Preview } from '@storybook/react'
import { Provider } from 'react-wrap-balancer'

import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import '@fontsource/material-icons'
import { ThemeProvider } from '@mui/material/styles'
import { ThemeProvider as TailWindProvider } from 'next-themes'

import '../src/theme/styles.css'
import { ColorModeContext, darkTheme, lightTheme } from '../src'

const preview: Preview = {
  decorators: [
    (Story) => {
      const [darkMode, setDarkMode] = useState(ColorModeContext)
      const theme = darkMode ? darkTheme : lightTheme

      return (
        <ThemeProvider theme={theme}>
          <TailWindProvider enableSystem attribute="class">
            <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
              <Provider><Story /></Provider>
            </ColorModeContext.Provider>
          </TailWindProvider>
        </ThemeProvider>
      )
    },
  ],
}

export default preview
