import { Preview } from '@storybook/react'
import { Provider } from 'react-wrap-balancer'

import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'
import '@fontsource/caveat'
import '@fontsource/material-icons'
import { Experimental_CssVarsProvider as CssVarsProvider } from '@mui/material/styles'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'

import '../src/theme/styles.css'
import { theme } from '../src'

const preview: Preview = {
  decorators: [
    (Story) => {
      return (
        <CssVarsProvider theme={theme} defaultMode="light">
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Provider>
              <Story />
            </Provider>
          </LocalizationProvider>
        </CssVarsProvider>
      )
    },
  ],
}

export default preview
