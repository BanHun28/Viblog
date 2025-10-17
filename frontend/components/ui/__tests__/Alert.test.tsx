import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Alert } from '../Alert';

describe('Alert', () => {
  it('renders children content', () => {
    render(<Alert>Alert message</Alert>);
    expect(screen.getByText('Alert message')).toBeInTheDocument();
  });

  it('has alert role', () => {
    render(<Alert>Message</Alert>);
    expect(screen.getByRole('alert')).toBeInTheDocument();
  });

  it('renders with title', () => {
    render(<Alert title="Warning">Alert content</Alert>);
    expect(screen.getByText('Warning')).toBeInTheDocument();
  });

  it('applies info variant by default', () => {
    const { container } = render(<Alert>Info message</Alert>);
    expect(container.firstChild).toHaveClass('bg-blue-50', 'border-blue-200');
  });

  it('applies success variant styling', () => {
    const { container } = render(<Alert variant="success">Success!</Alert>);
    expect(container.firstChild).toHaveClass('bg-green-50', 'border-green-200');
  });

  it('applies warning variant styling', () => {
    const { container } = render(<Alert variant="warning">Warning!</Alert>);
    expect(container.firstChild).toHaveClass('bg-yellow-50', 'border-yellow-200');
  });

  it('applies error variant styling', () => {
    const { container } = render(<Alert variant="error">Error!</Alert>);
    expect(container.firstChild).toHaveClass('bg-red-50', 'border-red-200');
  });

  it('displays appropriate icon for each variant', () => {
    const { container: infoContainer } = render(<Alert variant="info">Info</Alert>);
    expect(infoContainer.querySelector('svg')).toBeInTheDocument();

    const { container: successContainer } = render(<Alert variant="success">Success</Alert>);
    expect(successContainer.querySelector('svg')).toBeInTheDocument();

    const { container: warningContainer } = render(<Alert variant="warning">Warning</Alert>);
    expect(warningContainer.querySelector('svg')).toBeInTheDocument();

    const { container: errorContainer } = render(<Alert variant="error">Error</Alert>);
    expect(errorContainer.querySelector('svg')).toBeInTheDocument();
  });

  it('renders close button when onClose is provided', () => {
    render(<Alert onClose={() => {}}>Closeable alert</Alert>);
    expect(screen.getByLabelText('Close')).toBeInTheDocument();
  });

  it('calls onClose when close button is clicked', () => {
    const handleClose = vi.fn();
    render(<Alert onClose={handleClose}>Alert</Alert>);

    fireEvent.click(screen.getByLabelText('Close'));
    expect(handleClose).toHaveBeenCalledTimes(1);
  });

  it('does not render close button when onClose is not provided', () => {
    render(<Alert>Alert without close</Alert>);
    expect(screen.queryByLabelText('Close')).not.toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(<Alert className="custom-alert">Alert</Alert>);
    expect(container.firstChild).toHaveClass('custom-alert');
  });

  it('renders title and content together', () => {
    render(
      <Alert title="Important" variant="warning">
        This is an important warning message.
      </Alert>
    );

    expect(screen.getByText('Important')).toBeInTheDocument();
    expect(screen.getByText('This is an important warning message.')).toBeInTheDocument();
  });
});
