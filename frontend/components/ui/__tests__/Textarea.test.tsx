import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Textarea } from '../Textarea';

describe('Textarea', () => {
  it('renders textarea element', () => {
    render(<Textarea />);
    expect(screen.getByRole('textbox')).toBeInTheDocument();
  });

  it('renders with label', () => {
    render(<Textarea label="Description" />);
    expect(screen.getByLabelText('Description')).toBeInTheDocument();
  });

  it('displays error message', () => {
    render(<Textarea error="Description is required" />);
    const errorMessage = screen.getByText('Description is required');

    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage).toHaveClass('text-red-600');
  });

  it('displays helper text', () => {
    render(<Textarea helperText="Max 500 characters" />);
    expect(screen.getByText('Max 500 characters')).toBeInTheDocument();
  });

  it('applies error styling when error prop is provided', () => {
    render(<Textarea error="Error" />);
    const textarea = screen.getByRole('textbox');

    expect(textarea).toHaveClass('border-red-500');
    expect(textarea).toHaveAttribute('aria-invalid', 'true');
  });

  it('handles value changes', () => {
    const handleChange = vi.fn();
    render(<Textarea onChange={handleChange} />);
    const textarea = screen.getByRole('textbox');

    fireEvent.change(textarea, { target: { value: 'test content' } });
    expect(handleChange).toHaveBeenCalled();
  });

  it('can be disabled', () => {
    render(<Textarea disabled />);
    const textarea = screen.getByRole('textbox');

    expect(textarea).toBeDisabled();
    expect(textarea).toHaveClass('disabled:bg-gray-100');
  });

  it('has minimum height class', () => {
    render(<Textarea />);
    expect(screen.getByRole('textbox')).toHaveClass('min-h-[80px]');
  });

  it('is vertically resizable', () => {
    render(<Textarea />);
    expect(screen.getByRole('textbox')).toHaveClass('resize-y');
  });

  it('forwards ref correctly', () => {
    const ref = { current: null } as React.RefObject<HTMLTextAreaElement>;
    render(<Textarea ref={ref as any} />);

    expect(ref.current).toBeInstanceOf(HTMLTextAreaElement);
  });

  it('applies custom className', () => {
    render(<Textarea className="custom-textarea" />);
    expect(screen.getByRole('textbox')).toHaveClass('custom-textarea');
  });
});
