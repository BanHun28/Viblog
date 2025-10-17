import { describe, it, expect, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import { Footer } from '../Footer'

// Mock Next.js Link component
vi.mock('next/link', () => ({
  default: ({ children, href }: { children: React.ReactNode; href: string }) => (
    <a href={href}>{children}</a>
  ),
}))

describe('Footer', () => {
  it('renders brand name with link to home', () => {
    render(<Footer />)
    const brandLink = screen.getByText('Viblog')
    expect(brandLink).toBeInTheDocument()
    expect(brandLink.closest('a')).toHaveAttribute('href', '/')
  })

  it('displays current year in copyright', () => {
    render(<Footer />)
    const currentYear = new Date().getFullYear()
    expect(screen.getByText(new RegExp(`© ${currentYear}`))).toBeInTheDocument()
    expect(screen.getByText(/All rights reserved/i)).toBeInTheDocument()
  })

  it('renders RSS feed link', () => {
    render(<Footer />)
    const rssLink = screen.getByText('RSS').closest('a')
    expect(rssLink).toBeInTheDocument()
    expect(rssLink).toHaveAttribute('href', '/rss')
  })

  it('renders social media links with correct attributes', () => {
    render(<Footer />)

    // Find links by aria-label
    const githubLink = screen.getByLabelText('GitHub')
    const twitterLink = screen.getByLabelText('Twitter')

    expect(githubLink).toBeInTheDocument()
    expect(githubLink).toHaveAttribute('href', 'https://github.com')
    expect(githubLink).toHaveAttribute('target', '_blank')
    expect(githubLink).toHaveAttribute('rel', 'noopener noreferrer')

    expect(twitterLink).toBeInTheDocument()
    expect(twitterLink).toHaveAttribute('href', 'https://twitter.com')
    expect(twitterLink).toHaveAttribute('target', '_blank')
    expect(twitterLink).toHaveAttribute('rel', 'noopener noreferrer')
  })

  it('has proper footer role', () => {
    render(<Footer />)
    const footer = screen.getByRole('contentinfo')
    expect(footer).toBeInTheDocument()
  })

  it('applies dark mode classes', () => {
    render(<Footer />)
    const footer = screen.getByRole('contentinfo')
    expect(footer.className).toContain('dark:bg-gray-900')
    expect(footer.className).toContain('dark:border-gray-800')
  })

  it('renders social icons as SVG elements', () => {
    const { container } = render(<Footer />)
    const svgElements = container.querySelectorAll('svg')

    // Should have at least 3 SVGs (RSS icon + 2 social icons)
    expect(svgElements.length).toBeGreaterThanOrEqual(3)
  })

  it('has compact minimal layout', () => {
    render(<Footer />)
    const footer = screen.getByRole('contentinfo')

    // Check for minimal padding class in container div
    const containerDiv = footer.querySelector(".container")
    expect(containerDiv?.className).toContain("py-4")
  })

  it('uses flexbox layout for responsive design', () => {
    render(<Footer />)
    const footer = screen.getByRole('contentinfo')
    const container = footer.querySelector('.container')

    expect(container?.firstElementChild?.className).toContain('flex')
    expect(container?.firstElementChild?.className).toContain('md:flex-row')
  })
})
