import { Preview } from '@storybook/react'

import '../src/theme/styles.css'
import {ColorModeContext} from '../src/theme/ColorModeContext'
import {useState} from "react";

const preview: Preview = {
  decorators: [
    Story => {
      const [darkMode, setDarkMode] = useState(ColorModeContext)

      return (
        <ColorModeContext.Provider value={{ darkMode, setDarkMode }}>
          <Story />
        </ColorModeContext.Provider>
      )
    },
  ],
};

export default preview;
