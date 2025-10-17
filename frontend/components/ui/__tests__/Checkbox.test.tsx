import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Checkbox } from '../Checkbox';

describe('Checkbox', () => {
  it('renders checkbox input', () => {
    render(<Checkbox />);
    expect(screen.getByRole('checkbox')).toBeInTheDocument();
  });

  it('renders with label', () => {
    render(<Checkbox label="Accept terms" />);
    expect(screen.getByLabelText('Accept terms')).toBeInTheDocument();
  });

  it('displays error message when label is present', () => {
    render(<Checkbox label="Terms" error="You must accept the terms" />);
    expect(screen.getByText('You must accept the terms')).toBeInTheDocument();
  });

  it('applies error styling', () => {
    render(<Checkbox error="Error" />);
    const checkbox = screen.getByRole('checkbox');

    expect(checkbox).toHaveClass('border-red-500');
    expect(checkbox).toHaveAttribute('aria-invalid', 'true');
  });

  it('handles check/uncheck', () => {
    const handleChange = vi.fn();
    render(<Checkbox onChange={handleChange} />);
    const checkbox = screen.getByRole('checkbox');

    fireEvent.click(checkbox);
    expect(handleChange).toHaveBeenCalled();
  });

  it('can be checked by default', () => {
    render(<Checkbox defaultChecked />);
    expect(screen.getByRole('checkbox')).toBeChecked();
  });

  it('can be controlled', () => {
    const { rerender } = render(<Checkbox checked={false} onChange={() => {}} />);
    expect(screen.getByRole('checkbox')).not.toBeChecked();

    rerender(<Checkbox checked={true} onChange={() => {}} />);
    expect(screen.getByRole('checkbox')).toBeChecked();
  });

  it('can be disabled', () => {
    render(<Checkbox disabled />);
    const checkbox = screen.getByRole('checkbox');

    expect(checkbox).toBeDisabled();
    expect(checkbox).toHaveClass('disabled:cursor-not-allowed');
  });

  it('forwards ref correctly', () => {
    const ref = { current: null } as React.RefObject<HTMLInputElement>;
    render(<Checkbox ref={ref as any} />);

    expect(ref.current).toBeInstanceOf(HTMLInputElement);
  });

  it('applies custom className', () => {
    render(<Checkbox className="custom-checkbox" />);
    expect(screen.getByRole('checkbox')).toHaveClass('custom-checkbox');
  });

  it('associates label with checkbox correctly', () => {
    render(<Checkbox label="Remember me" id="remember" />);
    const checkbox = screen.getByRole('checkbox');
    const label = screen.getByText('Remember me');

    expect(label).toHaveAttribute('for', 'remember');
    expect(checkbox).toHaveAttribute('id', 'remember');
  });
});
