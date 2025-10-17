import React from 'react';
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter
} from '../Card';

describe('Card', () => {
  it('renders children', () => {
    render(<Card>Card content</Card>);
    expect(screen.getByText('Card content')).toBeInTheDocument();
  });

  it('applies default variant styling', () => {
    const { container } = render(<Card>Content</Card>);
    const card = container.firstChild as HTMLElement;

    expect(card).toHaveClass('bg-white', 'border');
  });

  it('applies outlined variant', () => {
    const { container } = render(<Card variant="outlined">Content</Card>);
    const card = container.firstChild as HTMLElement;

    expect(card).toHaveClass('bg-transparent', 'border');
  });

  it('applies elevated variant', () => {
    const { container } = render(<Card variant="elevated">Content</Card>);
    const card = container.firstChild as HTMLElement;

    expect(card).toHaveClass('shadow-lg');
  });

  it('applies different padding sizes', () => {
    const { rerender, container } = render(<Card padding="sm">Content</Card>);
    expect(container.firstChild).toHaveClass('p-3');

    rerender(<Card padding="md">Content</Card>);
    expect(container.firstChild).toHaveClass('p-4');

    rerender(<Card padding="lg">Content</Card>);
    expect(container.firstChild).toHaveClass('p-6');

    rerender(<Card padding="none">Content</Card>);
    expect(container.firstChild).not.toHaveClass('p-3', 'p-4', 'p-6');
  });

  it('applies custom className', () => {
    const { container } = render(<Card className="custom-card">Content</Card>);
    expect(container.firstChild).toHaveClass('custom-card');
  });

  it('renders complete card structure', () => {
    render(
      <Card>
        <CardHeader>
          <CardTitle>Card Title</CardTitle>
          <CardDescription>Card description text</CardDescription>
        </CardHeader>
        <CardContent>
          Main content goes here
        </CardContent>
        <CardFooter>
          Footer content
        </CardFooter>
      </Card>
    );

    expect(screen.getByText('Card Title')).toBeInTheDocument();
    expect(screen.getByText('Card description text')).toBeInTheDocument();
    expect(screen.getByText('Main content goes here')).toBeInTheDocument();
    expect(screen.getByText('Footer content')).toBeInTheDocument();
  });
});

describe('CardHeader', () => {
  it('renders children', () => {
    render(<CardHeader>Header content</CardHeader>);
    expect(screen.getByText('Header content')).toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(<CardHeader className="custom-header">Header</CardHeader>);
    expect(container.firstChild).toHaveClass('custom-header');
  });
});

describe('CardTitle', () => {
  it('renders as h3 element', () => {
    render(<CardTitle>Title</CardTitle>);
    const title = screen.getByText('Title');

    expect(title.tagName).toBe('H3');
    expect(title).toHaveClass('text-xl', 'font-semibold');
  });

  it('applies custom className', () => {
    render(<CardTitle className="custom-title">Title</CardTitle>);
    expect(screen.getByText('Title')).toHaveClass('custom-title');
  });
});

describe('CardDescription', () => {
  it('renders as paragraph element', () => {
    render(<CardDescription>Description</CardDescription>);
    const description = screen.getByText('Description');

    expect(description.tagName).toBe('P');
    expect(description).toHaveClass('text-sm', 'text-gray-600');
  });

  it('applies custom className', () => {
    render(<CardDescription className="custom-desc">Desc</CardDescription>);
    expect(screen.getByText('Desc')).toHaveClass('custom-desc');
  });
});

describe('CardContent', () => {
  it('renders children', () => {
    render(<CardContent>Content</CardContent>);
    expect(screen.getByText('Content')).toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(<CardContent className="custom-content">Content</CardContent>);
    expect(container.firstChild).toHaveClass('custom-content');
  });
});

describe('CardFooter', () => {
  it('renders children', () => {
    render(<CardFooter>Footer</CardFooter>);
    expect(screen.getByText('Footer')).toBeInTheDocument();
  });

  it('has border top styling', () => {
    const { container } = render(<CardFooter>Footer</CardFooter>);
    expect(container.firstChild).toHaveClass('border-t', 'border-gray-200');
  });

  it('applies custom className', () => {
    const { container } = render(<CardFooter className="custom-footer">Footer</CardFooter>);
    expect(container.firstChild).toHaveClass('custom-footer');
  });
});
