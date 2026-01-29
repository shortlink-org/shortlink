import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import MobileApps from '@/components/Landing/mobile'

describe('MobileApps', () => {
  it('renders section title', () => {
    render(<MobileApps />)
    expect(screen.getByText(/shortlink on the go/i)).toBeInTheDocument()
  })

  it('renders App Store button', () => {
    render(<MobileApps />)
    expect(screen.getByRole('button', { name: /app store/i })).toBeInTheDocument()
  })

  it('renders Google Play button', () => {
    render(<MobileApps />)
    expect(screen.getByRole('button', { name: /google play/i })).toBeInTheDocument()
  })

  it('renders description text', () => {
    render(<MobileApps />)
    expect(screen.getByText(/access your links anywhere/i)).toBeInTheDocument()
  })

  it('renders mobile apps badge', () => {
    render(<MobileApps />)
    // Use getAllByText since "mobile apps" appears in both badge and description
    const elements = screen.getAllByText(/mobile apps/i)
    expect(elements.length).toBeGreaterThanOrEqual(1)
  })
})
