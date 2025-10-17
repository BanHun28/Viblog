import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { render } from '@testing-library/react'
import { ThemeProvider } from '../ThemeProvider'
import { useUiStore } from '@/lib/store/uiStore'

// Mock zustand store
vi.mock('@/lib/store/uiStore', () => ({
  useUiStore: vi.fn(),
}))

describe('ThemeProvider', () => {
  const mockSetTheme = vi.fn()
  let localStorageMock: { [key: string]: string } = {}

  beforeEach(() => {
    vi.clearAllMocks()

    // Mock localStorage
    localStorageMock = {}
    global.localStorage = {
      getItem: vi.fn((key: string) => localStorageMock[key] || null),
      setItem: vi.fn((key: string, value: string) => {
        localStorageMock[key] = value
      }),
      removeItem: vi.fn((key: string) => {
        delete localStorageMock[key]
      }),
      clear: vi.fn(() => {
        localStorageMock = {}
      }),
      length: 0,
      key: vi.fn(),
    } as Storage

    // Mock matchMedia
    Object.defineProperty(window, 'matchMedia', {
      writable: true,
      value: vi.fn().mockImplementation((query: string) => ({
        matches: false,
        media: query,
        onchange: null,
        addListener: vi.fn(),
        removeListener: vi.fn(),
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        dispatchEvent: vi.fn(),
      })),
    })

    // Mock useUiStore default state
    vi.mocked(useUiStore).mockImplementation((selector: any) => {
      if (typeof selector === 'function') {
        return selector({
          theme: 'light',
          setTheme: mockSetTheme,
        })
      }
      return { theme: 'light', setTheme: mockSetTheme }
    })

    // Mock getState
    ;(useUiStore as any).getState = vi.fn(() => ({
      theme: 'light',
      setTheme: mockSetTheme,
    }))
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  it('renders children correctly', () => {
    const { container } = render(
      <ThemeProvider>
        <div data-testid="child">Test Child</div>
      </ThemeProvider>
    )

    expect(container.querySelector('[data-testid="child"]')).toBeInTheDocument()
  })

  it('applies light class to root element when theme is light', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'light', setTheme: mockSetTheme })
    )

    render(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    expect(document.documentElement.classList.contains('light')).toBe(true)
    expect(document.documentElement.classList.contains('dark')).toBe(false)
  })

  it('applies dark class to root element when theme is dark', () => {
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'dark', setTheme: mockSetTheme })
    )

    render(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    expect(document.documentElement.classList.contains('dark')).toBe(true)
    expect(document.documentElement.classList.contains('light')).toBe(false)
  })

  it('applies system theme class when theme is system', () => {
    // Mock system prefers dark mode
    Object.defineProperty(window, 'matchMedia', {
      writable: true,
      value: vi.fn().mockImplementation((query: string) => ({
        matches: query === '(prefers-color-scheme: dark)',
        media: query,
        onchange: null,
        addListener: vi.fn(),
        removeListener: vi.fn(),
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        dispatchEvent: vi.fn(),
      })),
    })

    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'system', setTheme: mockSetTheme })
    )

    render(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    expect(document.documentElement.classList.contains('dark')).toBe(true)
  })

  it('saves theme to localStorage when theme changes', () => {
    const { rerender } = render(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    // Change theme to dark
    vi.mocked(useUiStore).mockImplementation((selector: any) =>
      selector({ theme: 'dark', setTheme: mockSetTheme })
    )

    rerender(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    expect(localStorage.setItem).toHaveBeenCalledWith('theme', 'dark')
  })

  it('loads theme from localStorage on mount', () => {
    localStorageMock['theme'] = 'dark'

    render(
      <ThemeProvider>
        <div>Content</div>
      </ThemeProvider>
    )

    expect(localStorage.getItem).toHaveBeenCalledWith('theme')
  })
})
