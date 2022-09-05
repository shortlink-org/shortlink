import MoonIcon from "@heroicons/react/24/solid/MoonIcon";
import SunIcon from "@heroicons/react/24/solid/SunIcon";
import { useTheme as nextUseTheme } from 'next-themes'
import { useState, useContext, useEffect } from "react";
import { useTheme } from "@mui/material/styles"
import ColorModeContext from "../../theme/ColorModeContext";

const ToggleDarkMode = () => {
  // @ts-ignore
  const {systemTheme , theme, setTheme} = nextUseTheme()
  const [mounted, setMounted] = useState(false);

  useEffect(() =>{
    setMounted(true);
  },[])

  // @ts-ignore
  const { darkMode, setDarkMode } = useContext(ColorModeContext);

  // @ts-ignore
  const onClick = e => {
    setDarkMode(!darkMode)
    setTheme(e)
  }

  const renderThemeChanger= () => {
    if (!mounted) return null;

    const currentTheme = theme === "system" ? systemTheme : theme ;

    if (currentTheme ==="dark") {
      return (
        <SunIcon className="w-10 h-10 text-yellow-500 " role="button" onClick={() => onClick('light')} />
      )
    }

    else {
      return (
        <MoonIcon className="w-10 h-10 text-gray-900 " role="button" onClick={() => onClick('dark')} />
      )
    }
  }

  return renderThemeChanger()
}

export default ToggleDarkMode;
