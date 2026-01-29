import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, it, expect } from 'vitest'
import Subscribe from '@/components/Landing/subscribe'

describe('Subscribe', () => {
  it('renders newsletter title', () => {
    render(<Subscribe />)
    expect(screen.getByText(/subscribe to our/i)).toBeInTheDocument()
    expect(screen.getByText(/newsletter/i)).toBeInTheDocument()
  })

  it('renders email input', () => {
    render(<Subscribe />)
    expect(screen.getByPlaceholderText(/enter your email/i)).toBeInTheDocument()
  })

  it('renders subscribe button', () => {
    render(<Subscribe />)
    expect(screen.getByRole('button', { name: /subscribe/i })).toBeInTheDocument()
  })

  it('shows success message after submission', async () => {
    const user = userEvent.setup()
    render(<Subscribe />)

    const input = screen.getByPlaceholderText(/enter your email/i)
    const button = screen.getByRole('button', { name: /subscribe/i })

    await user.type(input, 'test@example.com')
    await user.click(button)

    await waitFor(
      () => {
        expect(screen.getByText(/you're subscribed/i)).toBeInTheDocument()
      },
      { timeout: 3000 },
    )
  })

  it('disables button while loading', async () => {
    const user = userEvent.setup()
    render(<Subscribe />)

    const input = screen.getByPlaceholderText(/enter your email/i)
    const button = screen.getByRole('button', { name: /subscribe/i })

    await user.type(input, 'test@example.com')
    await user.click(button)

    // Button should show loading state
    expect(screen.getByRole('button')).toBeDisabled()
  })
})
