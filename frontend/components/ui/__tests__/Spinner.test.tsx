import React from 'react';
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { Spinner, LoadingOverlay } from '../Spinner';

describe('Spinner', () => {
  it('renders spinner svg', () => {
    const { container } = render(<Spinner />);
    expect(container.querySelector('svg')).toBeInTheDocument();
  });

  it('has status role', () => {
    render(<Spinner />);
    expect(screen.getByRole('status')).toBeInTheDocument();
  });

  it('has loading aria-label', () => {
    render(<Spinner />);
    expect(screen.getByLabelText('Loading')).toBeInTheDocument();
  });

  it('applies medium size by default', () => {
    const { container } = render(<Spinner />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('w-6', 'h-6');
  });

  it('applies small size styling', () => {
    const { container } = render(<Spinner size="sm" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('w-4', 'h-4');
  });

  it('applies large size styling', () => {
    const { container } = render(<Spinner size="lg" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('w-8', 'h-8');
  });

  it('applies extra large size styling', () => {
    const { container } = render(<Spinner size="xl" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('w-12', 'h-12');
  });

  it('applies primary color by default', () => {
    const { container } = render(<Spinner />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('text-blue-600');
  });

  it('applies secondary color', () => {
    const { container } = render(<Spinner color="secondary" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('text-gray-600');
  });

  it('applies white color', () => {
    const { container } = render(<Spinner color="white" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('text-white');
  });

  it('has animate-spin class', () => {
    const { container } = render(<Spinner />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('animate-spin');
  });

  it('applies custom className', () => {
    const { container } = render(<Spinner className="custom-spinner" />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('custom-spinner');
  });
});

describe('LoadingOverlay', () => {
  it('renders children when not loading', () => {
    render(
      <LoadingOverlay isLoading={false}>
        <div>Content</div>
      </LoadingOverlay>
    );

    expect(screen.getByText('Content')).toBeInTheDocument();
    expect(screen.queryByRole('status')).not.toBeInTheDocument();
  });

  it('shows spinner when loading', () => {
    render(
      <LoadingOverlay isLoading={true}>
        <div>Content</div>
      </LoadingOverlay>
    );

    expect(screen.getByRole('status')).toBeInTheDocument();
  });

  it('displays default loading text', () => {
    render(<LoadingOverlay isLoading={true} />);
    expect(screen.getByText('Loading...')).toBeInTheDocument();
  });

  it('displays custom loading text', () => {
    render(<LoadingOverlay isLoading={true} text="Processing..." />);
    expect(screen.getByText('Processing...')).toBeInTheDocument();
  });

  it('makes content semi-transparent when loading', () => {
    const { container } = render(
      <LoadingOverlay isLoading={true}>
        <div>Content</div>
      </LoadingOverlay>
    );

    const contentWrapper = container.querySelector('.opacity-50');
    expect(contentWrapper).toBeInTheDocument();
    expect(contentWrapper).toHaveClass('pointer-events-none');
  });

  it('uses large spinner in overlay', () => {
    const { container } = render(<LoadingOverlay isLoading={true} />);
    const svg = container.querySelector('svg');
    expect(svg).toHaveClass('w-8', 'h-8');
  });

  it('centers loading overlay', () => {
    const { container } = render(<LoadingOverlay isLoading={true} />);
    const overlay = container.querySelector('.absolute.inset-0');

    expect(overlay).toHaveClass('flex', 'items-center', 'justify-center');
  });

  it('handles loading state toggle', () => {
    const { rerender } = render(
      <LoadingOverlay isLoading={false}>
        <div>Content</div>
      </LoadingOverlay>
    );

    expect(screen.queryByRole('status')).not.toBeInTheDocument();

    rerender(
      <LoadingOverlay isLoading={true}>
        <div>Content</div>
      </LoadingOverlay>
    );

    expect(screen.getByRole('status')).toBeInTheDocument();
  });
});
