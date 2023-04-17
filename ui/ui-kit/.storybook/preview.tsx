import {useState} from "react";
import { Preview } from '@storybook/react'

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import '@fontsource/material-icons';
import { ThemeProvider } from '@mui/material/styles'

import '../src/theme/styles.css'
import {ColorModeContext} from '../src/theme/ColorModeContext'
import {darkTheme, lightTheme} from '../src/theme/theme'

const preview: Preview = {
  decorators: [
    Story => {
      const [darkMode, setDarkMode] = useState(ColorModeContext)
      const theme = darkMode ? darkTheme : lightTheme

      return (
        <ThemeProvider theme={theme}>
          <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
            <Story />
          </ColorModeContext.Provider>
        </ThemeProvider>
      )
    },
  ],
};

export default preview;
