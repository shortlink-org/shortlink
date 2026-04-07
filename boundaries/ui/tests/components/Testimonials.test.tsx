import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect } from 'vitest'
import Testimonials from '@/components/Testimonials/Testimonials'

describe('Testimonials', () => {
  it('renders section title', () => {
    render(<Testimonials />)
    expect(screen.getByText(/loved by developers/i)).toBeInTheDocument()
  })

  it('renders first testimonial by default', () => {
    render(<Testimonials />)
    expect(screen.getByText(/viktor login/i)).toBeInTheDocument()
    expect(screen.getByText(/transformed how I manage/i)).toBeInTheDocument()
  })

  it('switches testimonial when a dot is selected', async () => {
    const user = userEvent.setup()
    render(<Testimonials />)

    await user.click(screen.getByRole('tab', { name: /testimonial 2/i }))
    expect(screen.getByText(/alex chen/i)).toBeInTheDocument()
    expect(screen.getByText(/respects privacy/i)).toBeInTheDocument()

    await user.click(screen.getByRole('tab', { name: /testimonial 3/i }))
    expect(screen.getByText(/sarah miller/i)).toBeInTheDocument()
    expect(screen.getByText(/API is fantastic/i)).toBeInTheDocument()
  })
})
