import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Input } from '../Input';

describe('Input', () => {
  it('renders input element', () => {
    render(<Input />);
    expect(screen.getByRole('textbox')).toBeInTheDocument();
  });

  it('renders with label', () => {
    render(<Input label="Username" />);
    expect(screen.getByLabelText('Username')).toBeInTheDocument();
  });

  it('displays error message', () => {
    render(<Input error="This field is required" />);
    const errorMessage = screen.getByText('This field is required');

    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage).toHaveClass('text-red-600');
  });

  it('displays helper text', () => {
    render(<Input helperText="Enter your email address" />);
    const helperText = screen.getByText('Enter your email address');

    expect(helperText).toBeInTheDocument();
    expect(helperText).toHaveClass('text-gray-500');
  });

  it('does not show helper text when error is present', () => {
    render(
      <Input
        error="Error message"
        helperText="Helper text should not appear"
      />
    );

    expect(screen.getByText('Error message')).toBeInTheDocument();
    expect(screen.queryByText('Helper text should not appear')).not.toBeInTheDocument();
  });

  it('applies error styling when error prop is provided', () => {
    render(<Input error="Error" />);
    const input = screen.getByRole('textbox');

    expect(input).toHaveClass('border-red-500');
    expect(input).toHaveAttribute('aria-invalid', 'true');
  });

  it('handles value changes', () => {
    const handleChange = vi.fn();
    render(<Input onChange={handleChange} />);
    const input = screen.getByRole('textbox');

    fireEvent.change(input, { target: { value: 'test value' } });
    expect(handleChange).toHaveBeenCalled();
  });

  it('can be disabled', () => {
    render(<Input disabled />);
    const input = screen.getByRole('textbox');

    expect(input).toBeDisabled();
    expect(input).toHaveClass('disabled:bg-gray-100');
  });

  it('applies custom className', () => {
    render(<Input className="custom-input" />);
    expect(screen.getByRole('textbox')).toHaveClass('custom-input');
  });

  it('forwards ref correctly', () => {
    const ref = { current: null } as React.RefObject<HTMLInputElement>;
    render(<Input ref={ref as any} />);

    expect(ref.current).toBeInstanceOf(HTMLInputElement);
  });

  it('supports different input types', () => {
    render(<Input type="email" />);
    const input = screen.getByRole('textbox');

    expect(input).toHaveAttribute('type', 'email');
  });

  it('associates label with input using htmlFor/id', () => {
    render(<Input label="Email" id="email-input" />);
    const input = screen.getByRole('textbox');
    const label = screen.getByText('Email');

    expect(label).toHaveAttribute('for', 'email-input');
    expect(input).toHaveAttribute('id', 'email-input');
  });

  it('generates unique id when not provided', () => {
    const { container: container1 } = render(<Input label="Input 1" />);
    const { container: container2 } = render(<Input label="Input 2" />);

    const input1 = container1.querySelector('input');
    const input2 = container2.querySelector('input');

    expect(input1?.id).not.toBe(input2?.id);
  });
});
