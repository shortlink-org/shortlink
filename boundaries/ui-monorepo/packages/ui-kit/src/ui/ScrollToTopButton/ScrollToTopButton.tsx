import React, { useEffect, useState, useCallback } from 'react'
import { ArrowUpward } from '@mui/icons-material'

interface ScrollToTopButtonProps {
  visibilityThreshold?: number // in pixels
}

const ScrollToTopButton: React.FC<ScrollToTopButtonProps> = ({
  visibilityThreshold = 500,
}) => {
  const [isVisible, setIsVisible] = useState<boolean>(false)

  const toggleVisibility = useCallback(() => {
    setIsVisible(window.scrollY > visibilityThreshold)
  }, [visibilityThreshold])

  useEffect(() => {
    window.addEventListener('scroll', toggleVisibility)
    return () => window.removeEventListener('scroll', toggleVisibility)
  }, [toggleVisibility])

  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth',
    })
  }

  return (
    <button
      onClick={scrollToTop}
      type="button"
      style={{ right: '3em', bottom: '3em' }}
      className={`fixed bottom-5 right-5 p-2 bg-indigo-500 dark:bg-gray-800 text-white rounded-full shadow-lg 
                  transition-opacity duration-500 ease-in-out 
                  ${isVisible ? 'opacity-100' : 'opacity-0'}
                  hover:bg-indigo-400 focus:outline-none focus:ring-2 focus:ring-blue-300 focus:bg-indigo-400`}
      aria-label="Scroll to top"
    >
      <ArrowUpward />
    </button>
  )
}

export default ScrollToTopButton
