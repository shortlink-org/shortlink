import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import Header from '@/components/Page/Header'

describe('Page Header', () => {
  it('renders title', () => {
    render(<Header title="Test Title" />)
    expect(screen.getByRole('heading', { name: /test title/i })).toBeInTheDocument()
  })

  it('renders description when provided', () => {
    render(<Header title="Title" description="This is a description" />)
    expect(screen.getByText(/this is a description/i)).toBeInTheDocument()
  })

  it('does not render description when not provided', () => {
    render(<Header title="Title Only" />)
    expect(screen.queryByText(/description/i)).not.toBeInTheDocument()
  })
})
