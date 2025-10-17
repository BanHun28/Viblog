import React from 'react'
import { describe, it, expect, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import { MainLayout } from '../MainLayout'

// Mock Header component
vi.mock('../Header', () => ({
  Header: () => <header data-testid="mock-header">Header</header>,
}))

// Mock Footer component
vi.mock('../Footer', () => ({
  Footer: () => <footer data-testid="mock-footer">Footer</footer>,
}))

describe('MainLayout', () => {
  it('renders children correctly', () => {
    render(
      <MainLayout>
        <div data-testid="test-content">Test Content</div>
      </MainLayout>
    )

    expect(screen.getByTestId('test-content')).toBeInTheDocument()
    expect(screen.getByText('Test Content')).toBeInTheDocument()
  })

  it('renders Header component', () => {
    render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    expect(screen.getByTestId('mock-header')).toBeInTheDocument()
  })

  it('renders Footer component', () => {
    render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    expect(screen.getByTestId('mock-footer')).toBeInTheDocument()
  })

  it('has correct layout structure (Header, Main, Footer)', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const layoutDiv = container.firstChild as HTMLElement
    expect(layoutDiv.children.length).toBe(3)

    // Check order: Header -> Main -> Footer
    expect(layoutDiv.children[0]).toHaveAttribute('data-testid', 'mock-header')
    expect(layoutDiv.children[1].tagName).toBe('MAIN')
    expect(layoutDiv.children[2]).toHaveAttribute('data-testid', 'mock-footer')
  })

  it('applies min-h-screen for full viewport height', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const layoutDiv = container.firstChild as HTMLElement
    expect(layoutDiv.className).toContain('min-h-screen')
  })

  it('applies flexbox layout', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const layoutDiv = container.firstChild as HTMLElement
    expect(layoutDiv.className).toContain('flex')
    expect(layoutDiv.className).toContain('flex-col')
  })

  it('applies dark mode background classes', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const layoutDiv = container.firstChild as HTMLElement
    expect(layoutDiv.className).toContain('bg-white')
    expect(layoutDiv.className).toContain('dark:bg-gray-950')
  })

  it('applies dark mode text classes', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const layoutDiv = container.firstChild as HTMLElement
    expect(layoutDiv.className).toContain('text-gray-900')
    expect(layoutDiv.className).toContain('dark:text-gray-100')
  })

  it('main content has flex-1 for flexible height', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const main = container.querySelector('main')
    expect(main?.className).toContain('flex-1')
  })

  it('main content has container and padding classes', () => {
    const { container } = render(
      <MainLayout>
        <div>Content</div>
      </MainLayout>
    )

    const main = container.querySelector('main')
    expect(main?.className).toContain('container')
    expect(main?.className).toContain('mx-auto')
    expect(main?.className).toContain('px-4')
    expect(main?.className).toContain('py-8')
  })

  it('renders multiple children correctly', () => {
    render(
      <MainLayout>
        <div data-testid="child-1">Child 1</div>
        <div data-testid="child-2">Child 2</div>
        <div data-testid="child-3">Child 3</div>
      </MainLayout>
    )

    expect(screen.getByTestId('child-1')).toBeInTheDocument()
    expect(screen.getByTestId('child-2')).toBeInTheDocument()
    expect(screen.getByTestId('child-3')).toBeInTheDocument()
  })
})
