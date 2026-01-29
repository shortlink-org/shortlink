import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import Statistic from '@/components/Dashboard/stats'

describe('Statistic', () => {
  it('renders count of 0', () => {
    render(<Statistic count={0} />)
    expect(screen.getByText('0')).toBeInTheDocument()
    expect(screen.getByText(/don't have any saved links/i)).toBeInTheDocument()
  })

  it('renders count of 1 with singular text', () => {
    render(<Statistic count={1} />)
    expect(screen.getByText('1')).toBeInTheDocument()
    expect(screen.getByText(/one saved link/i)).toBeInTheDocument()
  })

  it('renders count of multiple links', () => {
    render(<Statistic count={42} />)
    expect(screen.getByText('42')).toBeInTheDocument()
    expect(screen.getByText(/42 saved links/i)).toBeInTheDocument()
  })

  it('renders link icon', () => {
    render(<Statistic count={5} />)
    const svg = document.querySelector('svg')
    expect(svg).toBeInTheDocument()
  })
})
