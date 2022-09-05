import MoonIcon from "@heroicons/react/24/solid/MoonIcon";
import SunIcon from "@heroicons/react/24/solid/SunIcon";
import { useTheme } from 'next-themes'
import {useState, useEffect} from "react";

const ToggleDarkMode = () => {
  // @ts-ignore
  const {systemTheme , theme, setTheme} = useTheme ()
  const [mounted, setMounted] = useState(false);

  useEffect(() =>{
    setMounted(true);
  },[])

  const renderThemeChanger= () => {
    if (!mounted) return null;

    const currentTheme = theme === "system" ? systemTheme : theme ;

    if (currentTheme ==="dark") {
      return (
        <SunIcon className="w-10 h-10 text-yellow-500 " role="button" onClick={() => setTheme('light')} />
      )
    }

    else {
      return (
        <MoonIcon className="w-10 h-10 text-gray-900 " role="button" onClick={() => setTheme('dark')} />
      )
    }
  }

  return renderThemeChanger()
}

export default ToggleDarkMode;
