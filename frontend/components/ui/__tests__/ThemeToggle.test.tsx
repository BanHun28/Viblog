import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import { ThemeToggle } from '../ThemeToggle'
import { useUiStore } from '@/lib/store/uiStore'

// Mock zustand store
vi.mock('@/lib/store/uiStore', () => ({
  useUiStore: vi.fn(),
}))

describe('ThemeToggle', () => {
  const mockSetTheme = vi.fn()

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('renders light mode icon when theme is light', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'light', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })
    expect(button).toBeInTheDocument()
    expect(button.querySelector('svg')).toBeInTheDocument()
  })

  it('renders dark mode icon when theme is dark', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'dark', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })
    expect(button).toBeInTheDocument()
    expect(button.querySelector('svg')).toBeInTheDocument()
  })

  it('renders system mode icon when theme is system', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'system', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })
    expect(button).toBeInTheDocument()
    expect(button.querySelector('svg')).toBeInTheDocument()
  })

  it('cycles from light to dark when clicked', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'light', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })

    fireEvent.click(button)
    expect(mockSetTheme).toHaveBeenCalledWith('dark')
  })

  it('cycles from dark to system when clicked', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'dark', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })

    fireEvent.click(button)
    expect(mockSetTheme).toHaveBeenCalledWith('system')
  })

  it('cycles from system to light when clicked', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'system', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })

    fireEvent.click(button)
    expect(mockSetTheme).toHaveBeenCalledWith('light')
  })

  it('has correct title attribute showing current theme', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'light', setTheme: mockSetTheme })
    )

    render(<ThemeToggle />)
    const button = screen.getByRole('button', { name: /toggle theme/i })
    expect(button).toHaveAttribute('title', 'Current theme: light')
  })
})
