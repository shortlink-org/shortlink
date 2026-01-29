'use client'

import { useRef, useState, useEffect } from 'react'
import { motion, useMotionValue, useSpring, useTransform } from 'framer-motion'

interface MagneticProps {
  children: React.ReactNode
  className?: string
  strength?: number
  radius?: number
}

export function Magnetic({ children, className = '', strength = 0.35, radius = 150 }: MagneticProps) {
  const ref = useRef<HTMLDivElement>(null)

  const x = useMotionValue(0)
  const y = useMotionValue(0)

  const springConfig = { damping: 15, stiffness: 150, mass: 0.1 }
  const springX = useSpring(x, springConfig)
  const springY = useSpring(y, springConfig)

  const handleMouseMove = (e: React.MouseEvent) => {
    if (!ref.current) return

    const rect = ref.current.getBoundingClientRect()
    const centerX = rect.left + rect.width / 2
    const centerY = rect.top + rect.height / 2

    const distanceX = e.clientX - centerX
    const distanceY = e.clientY - centerY
    const distance = Math.sqrt(distanceX * distanceX + distanceY * distanceY)

    if (distance < radius) {
      const factor = 1 - distance / radius
      x.set(distanceX * strength * factor)
      y.set(distanceY * strength * factor)
    }
  }

  const handleMouseLeave = () => {
    x.set(0)
    y.set(0)
  }

  return (
    <motion.div
      ref={ref}
      className={className}
      style={{ x: springX, y: springY }}
      onMouseMove={handleMouseMove}
      onMouseLeave={handleMouseLeave}
    >
      {children}
    </motion.div>
  )
}

// Custom animated cursor
export function CursorDot() {
  const cursorX = useMotionValue(0)
  const cursorY = useMotionValue(0)

  const springConfig = { damping: 25, stiffness: 400 }
  const cursorXSpring = useSpring(cursorX, springConfig)
  const cursorYSpring = useSpring(cursorY, springConfig)

  // Slower spring for outer ring
  const ringConfig = { damping: 30, stiffness: 200 }
  const ringXSpring = useSpring(cursorX, ringConfig)
  const ringYSpring = useSpring(cursorY, ringConfig)

  const [isVisible, setIsVisible] = useState(false)
  const [isHovering, setIsHovering] = useState(false)
  const [isPressed, setIsPressed] = useState(false)

  useEffect(() => {
    // Check for touch device
    if ('ontouchstart' in window) return

    const handleMouseMove = (e: MouseEvent) => {
      cursorX.set(e.clientX)
      cursorY.set(e.clientY)
      setIsVisible(true)
    }

    const handleMouseLeave = () => setIsVisible(false)
    const handleMouseDown = () => setIsPressed(true)
    const handleMouseUp = () => setIsPressed(false)

    const handleMouseEnterInteractive = () => setIsHovering(true)
    const handleMouseLeaveInteractive = () => setIsHovering(false)

    window.addEventListener('mousemove', handleMouseMove)
    document.addEventListener('mouseleave', handleMouseLeave)
    document.addEventListener('mousedown', handleMouseDown)
    document.addEventListener('mouseup', handleMouseUp)

    // Watch for interactive elements
    const observer = new MutationObserver(() => {
      const interactiveElements = document.querySelectorAll('a, button, [role="button"], input, textarea, select')
      interactiveElements.forEach((el) => {
        el.removeEventListener('mouseenter', handleMouseEnterInteractive)
        el.removeEventListener('mouseleave', handleMouseLeaveInteractive)
        el.addEventListener('mouseenter', handleMouseEnterInteractive)
        el.addEventListener('mouseleave', handleMouseLeaveInteractive)
      })
    })

    observer.observe(document.body, { childList: true, subtree: true })

    // Initial setup
    const interactiveElements = document.querySelectorAll('a, button, [role="button"], input, textarea, select')
    interactiveElements.forEach((el) => {
      el.addEventListener('mouseenter', handleMouseEnterInteractive)
      el.addEventListener('mouseleave', handleMouseLeaveInteractive)
    })

    return () => {
      window.removeEventListener('mousemove', handleMouseMove)
      document.removeEventListener('mouseleave', handleMouseLeave)
      document.removeEventListener('mousedown', handleMouseDown)
      document.removeEventListener('mouseup', handleMouseUp)
      observer.disconnect()
    }
  }, [cursorX, cursorY])

  // Don't render on server or touch devices
  if (typeof window === 'undefined') return null

  return (
    <>
      {/* Main cursor dot */}
      <motion.div
        className="fixed top-0 left-0 pointer-events-none z-[9999] mix-blend-difference"
        style={{
          x: cursorXSpring,
          y: cursorYSpring,
        }}
        animate={{
          scale: isPressed ? 0.8 : isHovering ? 0.5 : 1,
          opacity: isVisible ? 1 : 0,
        }}
        transition={{ duration: 0.15 }}
      >
        <motion.div
          className="w-3 h-3 bg-white rounded-full -translate-x-1/2 -translate-y-1/2"
          animate={{
            scale: isHovering ? 0 : 1,
          }}
          transition={{ duration: 0.2 }}
        />
      </motion.div>

      {/* Outer ring */}
      <motion.div
        className="fixed top-0 left-0 pointer-events-none z-[9998] mix-blend-difference"
        style={{
          x: ringXSpring,
          y: ringYSpring,
        }}
        animate={{
          opacity: isVisible ? 1 : 0,
        }}
        transition={{ duration: 0.2 }}
      >
        <motion.div
          className="border-2 border-white rounded-full -translate-x-1/2 -translate-y-1/2"
          animate={{
            width: isPressed ? 30 : isHovering ? 60 : 40,
            height: isPressed ? 30 : isHovering ? 60 : 40,
            borderWidth: isHovering ? 3 : 2,
          }}
          transition={{
            type: 'spring',
            damping: 20,
            stiffness: 300,
          }}
        />
      </motion.div>
    </>
  )
}

// Magnetic button wrapper with scale effect
export function MagneticButton({ children, className = '', ...props }: MagneticProps & React.HTMLAttributes<HTMLDivElement>) {
  const [isHovered, setIsHovered] = useState(false)

  return (
    <Magnetic className={className} {...props}>
      <motion.div
        onHoverStart={() => setIsHovered(true)}
        onHoverEnd={() => setIsHovered(false)}
        animate={{
          scale: isHovered ? 1.05 : 1,
        }}
        transition={{
          type: 'spring',
          damping: 15,
          stiffness: 300,
        }}
      >
        {children}
      </motion.div>
    </Magnetic>
  )
}

export default Magnetic
