import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import Testimonials from '@/components/Testimonials/Testimonials'

describe('Testimonials', () => {
  it('renders section title', () => {
    render(<Testimonials />)
    expect(screen.getByText(/loved by developers/i)).toBeInTheDocument()
  })

  it('renders all testimonials', () => {
    render(<Testimonials />)
    expect(screen.getByText(/viktor login/i)).toBeInTheDocument()
    expect(screen.getByText(/alex chen/i)).toBeInTheDocument()
    expect(screen.getByText(/sarah miller/i)).toBeInTheDocument()
  })

  it('renders star ratings', () => {
    render(<Testimonials />)
    // Each testimonial has 5 stars, 3 testimonials = 15 stars
    const stars = document.querySelectorAll('svg')
    expect(stars.length).toBeGreaterThanOrEqual(15)
  })

  it('renders testimonial quotes', () => {
    render(<Testimonials />)
    expect(screen.getByText(/transformed how I manage/i)).toBeInTheDocument()
    expect(screen.getByText(/respects privacy/i)).toBeInTheDocument()
    expect(screen.getByText(/API is fantastic/i)).toBeInTheDocument()
  })
})
