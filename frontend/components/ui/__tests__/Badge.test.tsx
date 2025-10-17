import React from 'react';
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { Badge } from '../Badge';

describe('Badge', () => {
  it('renders children content', () => {
    render(<Badge>New</Badge>);
    expect(screen.getByText('New')).toBeInTheDocument();
  });

  it('applies default variant by default', () => {
    const { container } = render(<Badge>Badge</Badge>);
    expect(container.firstChild).toHaveClass('bg-gray-100', 'text-gray-800');
  });

  it('applies primary variant styling', () => {
    const { container } = render(<Badge variant="primary">Primary</Badge>);
    expect(container.firstChild).toHaveClass('bg-blue-100', 'text-blue-800');
  });

  it('applies success variant styling', () => {
    const { container } = render(<Badge variant="success">Success</Badge>);
    expect(container.firstChild).toHaveClass('bg-green-100', 'text-green-800');
  });

  it('applies warning variant styling', () => {
    const { container } = render(<Badge variant="warning">Warning</Badge>);
    expect(container.firstChild).toHaveClass('bg-yellow-100', 'text-yellow-800');
  });

  it('applies error variant styling', () => {
    const { container } = render(<Badge variant="error">Error</Badge>);
    expect(container.firstChild).toHaveClass('bg-red-100', 'text-red-800');
  });

  it('applies outline variant styling', () => {
    const { container } = render(<Badge variant="outline">Outline</Badge>);
    expect(container.firstChild).toHaveClass('border', 'border-gray-300');
  });

  it('applies medium size by default', () => {
    const { container } = render(<Badge>Badge</Badge>);
    expect(container.firstChild).toHaveClass('text-sm', 'px-2.5');
  });

  it('applies small size styling', () => {
    const { container } = render(<Badge size="sm">Small</Badge>);
    expect(container.firstChild).toHaveClass('text-xs', 'px-2');
  });

  it('applies large size styling', () => {
    const { container } = render(<Badge size="lg">Large</Badge>);
    expect(container.firstChild).toHaveClass('text-base', 'px-3');
  });

  it('has rounded-full class', () => {
    const { container } = render(<Badge>Badge</Badge>);
    expect(container.firstChild).toHaveClass('rounded-full');
  });

  it('applies custom className', () => {
    const { container } = render(<Badge className="custom-badge">Badge</Badge>);
    expect(container.firstChild).toHaveClass('custom-badge');
  });

  it('renders as span element', () => {
    const { container } = render(<Badge>Badge</Badge>);
    expect(container.firstChild?.nodeName).toBe('SPAN');
  });

  it('supports additional props', () => {
    render(<Badge data-testid="test-badge">Badge</Badge>);
    expect(screen.getByTestId('test-badge')).toBeInTheDocument();
  });
});
