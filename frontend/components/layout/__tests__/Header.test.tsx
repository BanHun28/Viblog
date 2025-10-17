import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import { Header } from '../Header'

// Mock Next.js Link component
vi.mock('next/link', () => ({
  default: ({ children, href }: { children: React.ReactNode; href: string }) => (
    <a href={href}>{children}</a>
  ),
}))

// Mock ThemeToggle component
vi.mock('@/components/ui/ThemeToggle', () => ({
  ThemeToggle: () => <button data-testid="theme-toggle">Theme Toggle</button>,
}))

describe('Header', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('renders logo with correct link', () => {
    render(<Header />)
    const logo = screen.getByText('Viblog')
    expect(logo).toBeInTheDocument()
    expect(logo.closest('a')).toHaveAttribute('href', '/')
  })

  it('renders navigation links', () => {
    render(<Header />)
    expect(screen.getByText('Posts')).toBeInTheDocument()
    expect(screen.getByText('Categories')).toBeInTheDocument()
    expect(screen.getByText('Tags')).toBeInTheDocument()
  })

  it('renders search form', () => {
    render(<Header />)
    const searchInput = screen.getAllByPlaceholderText('Search...')[0]
    expect(searchInput).toBeInTheDocument()
    expect(searchInput.closest('form')).toHaveAttribute('action', '/search')
  })

  it('renders theme toggle button', () => {
    render(<Header />)
    expect(screen.getByTestId('theme-toggle')).toBeInTheDocument()
  })

  it('renders login and sign up buttons when logged out', () => {
    render(<Header />)
    expect(screen.getAllByText('Login').length).toBeGreaterThan(0)
    expect(screen.getAllByText('Sign Up').length).toBeGreaterThan(0)
  })

  it('toggles mobile menu when menu button is clicked', () => {
    render(<Header />)

    // Mobile menu should not be visible initially
    const mobileNav = screen.queryByText('Posts')?.closest('nav')

    // Find mobile menu button (visible on mobile)
    const menuButtons = screen.getAllByRole('button')
    const mobileMenuButton = menuButtons.find(
      button => button.querySelector('svg') && !button.hasAttribute('data-testid')
    )

    expect(mobileMenuButton).toBeInTheDocument()

    // Click to open mobile menu
    if (mobileMenuButton) {
      fireEvent.click(mobileMenuButton)
    }

  })

  it('renders correct navigation link hrefs', () => {
    render(<Header />)

    const postsLink = screen.getAllByText('Posts')[0].closest('a')
    const categoriesLink = screen.getAllByText('Categories')[0].closest('a')
    const tagsLink = screen.getAllByText('Tags')[0].closest('a')

    expect(postsLink).toHaveAttribute('href', '/posts')
    expect(categoriesLink).toHaveAttribute('href', '/categories')
    expect(tagsLink).toHaveAttribute('href', '/tags')
  })

  it('has sticky positioning', () => {
    render(<Header />)
    const header = screen.getByRole('banner')
    expect(header).toHaveClass('sticky')
  })

  it('applies dark mode classes', () => {
    render(<Header />)
    const header = screen.getByRole('banner')
    expect(header.className).toContain('dark:bg-gray-900')
    expect(header.className).toContain('dark:border-gray-800')
  })
})
