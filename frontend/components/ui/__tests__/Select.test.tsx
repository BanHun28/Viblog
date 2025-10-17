import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Select } from '../Select';

describe('Select', () => {
  const mockOptions = [
    { value: '1', label: 'Option 1' },
    { value: '2', label: 'Option 2' },
    { value: '3', label: 'Option 3', disabled: true },
  ];

  it('renders select element', () => {
    render(<Select options={mockOptions} />);
    expect(screen.getByRole('combobox')).toBeInTheDocument();
  });

  it('renders with label', () => {
    render(<Select label="Choose option" options={mockOptions} />);
    expect(screen.getByLabelText('Choose option')).toBeInTheDocument();
  });

  it('renders all options', () => {
    render(<Select options={mockOptions} />);

    expect(screen.getByRole('option', { name: 'Option 1' })).toBeInTheDocument();
    expect(screen.getByRole('option', { name: 'Option 2' })).toBeInTheDocument();
    expect(screen.getByRole('option', { name: 'Option 3' })).toBeInTheDocument();
  });

  it('renders placeholder option', () => {
    render(<Select options={mockOptions} placeholder="Select an option" />);

    const placeholderOption = screen.getByRole('option', { name: 'Select an option' });
    expect(placeholderOption).toBeInTheDocument();
    expect(placeholderOption).toHaveAttribute('disabled');
  });

  it('disables specific options', () => {
    render(<Select options={mockOptions} />);
    const disabledOption = screen.getByRole('option', { name: 'Option 3' });

    expect(disabledOption).toBeDisabled();
  });

  it('displays error message', () => {
    render(<Select options={mockOptions} error="Selection is required" />);
    const errorMessage = screen.getByText('Selection is required');

    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage).toHaveClass('text-red-600');
  });

  it('displays helper text', () => {
    render(<Select options={mockOptions} helperText="Choose one option" />);
    expect(screen.getByText('Choose one option')).toBeInTheDocument();
  });

  it('applies error styling', () => {
    render(<Select options={mockOptions} error="Error" />);
    const select = screen.getByRole('combobox');

    expect(select).toHaveClass('border-red-500');
    expect(select).toHaveAttribute('aria-invalid', 'true');
  });

  it('handles selection changes', () => {
    const handleChange = vi.fn();
    render(<Select options={mockOptions} onChange={handleChange} />);
    const select = screen.getByRole('combobox');

    fireEvent.change(select, { target: { value: '2' } });
    expect(handleChange).toHaveBeenCalled();
  });

  it('can be disabled', () => {
    render(<Select options={mockOptions} disabled />);
    const select = screen.getByRole('combobox');

    expect(select).toBeDisabled();
    expect(select).toHaveClass('disabled:bg-gray-100');
  });

  it('forwards ref correctly', () => {
    const ref = { current: null } as React.RefObject<HTMLSelectElement>;
    render(<Select options={mockOptions} ref={ref as any} />);

    expect(ref.current).toBeInstanceOf(HTMLSelectElement);
  });

  it('applies custom className', () => {
    render(<Select options={mockOptions} className="custom-select" />);
    expect(screen.getByRole('combobox')).toHaveClass('custom-select');
  });
});
