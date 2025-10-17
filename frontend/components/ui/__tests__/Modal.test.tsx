import React from 'react';
import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Modal, ModalHeader, ModalBody, ModalFooter } from '../Modal';

describe('Modal', () => {
  beforeEach(() => {
    // Reset body overflow style before each test
    document.body.style.overflow = '';
  });

  afterEach(() => {
    // Clean up after each test
    document.body.style.overflow = '';
  });

  it('does not render when isOpen is false', () => {
    render(
      <Modal isOpen={false} onClose={() => {}}>
        <div>Modal content</div>
      </Modal>
    );

    expect(screen.queryByText('Modal content')).not.toBeInTheDocument();
  });

  it('renders when isOpen is true', () => {
    render(
      <Modal isOpen={true} onClose={() => {}}>
        <div>Modal content</div>
      </Modal>
    );

    expect(screen.getByText('Modal content')).toBeInTheDocument();
  });

  it('has dialog role', () => {
    render(
      <Modal isOpen={true} onClose={() => {}}>
        <div>Content</div>
      </Modal>
    );

    expect(screen.getByRole('dialog')).toBeInTheDocument();
  });

  it('has aria-modal attribute', () => {
    render(
      <Modal isOpen={true} onClose={() => {}}>
        <div>Content</div>
      </Modal>
    );

    expect(screen.getByRole('dialog')).toHaveAttribute('aria-modal', 'true');
  });

  it('renders title when provided', () => {
    render(
      <Modal isOpen={true} onClose={() => {}} title="Modal Title">
        <div>Content</div>
      </Modal>
    );

    expect(screen.getByText('Modal Title')).toBeInTheDocument();
  });

  it('renders close button in title area', () => {
    render(
      <Modal isOpen={true} onClose={() => {}} title="Title">
        <div>Content</div>
      </Modal>
    );

    expect(screen.getByLabelText('Close')).toBeInTheDocument();
  });

  it('calls onClose when close button is clicked', () => {
    const handleClose = vi.fn();
    render(
      <Modal isOpen={true} onClose={handleClose} title="Title">
        <div>Content</div>
      </Modal>
    );

    fireEvent.click(screen.getByLabelText('Close'));
    expect(handleClose).toHaveBeenCalledTimes(1);
  });

  it('calls onClose when backdrop is clicked', () => {
    const handleClose = vi.fn();
    const { container } = render(
      <Modal isOpen={true} onClose={handleClose}>
        <div>Content</div>
      </Modal>
    );

    const backdrop = container.querySelector('.bg-black.bg-opacity-50');
    if (backdrop) {
      fireEvent.click(backdrop);
      expect(handleClose).toHaveBeenCalledTimes(1);
    }
  });

  it('calls onClose when Escape key is pressed', () => {
    const handleClose = vi.fn();
    render(
      <Modal isOpen={true} onClose={handleClose}>
        <div>Content</div>
      </Modal>
    );

    fireEvent.keyDown(document, { key: 'Escape' });
    expect(handleClose).toHaveBeenCalledTimes(1);
  });

  it('sets body overflow to hidden when open', () => {
    render(
      <Modal isOpen={true} onClose={() => {}}>
        <div>Content</div>
      </Modal>
    );

    expect(document.body.style.overflow).toBe('hidden');
  });

  it('applies medium size by default', () => {
    const { container } = render(
      <Modal isOpen={true} onClose={() => {}}>
        <div>Content</div>
      </Modal>
    );

    const modalPanel = container.querySelector('.max-w-lg');
    expect(modalPanel).toBeInTheDocument();
  });

  it('applies small size styling', () => {
    const { container } = render(
      <Modal isOpen={true} onClose={() => {}} size="sm">
        <div>Content</div>
      </Modal>
    );

    expect(container.querySelector('.max-w-md')).toBeInTheDocument();
  });

  it('applies large size styling', () => {
    const { container } = render(
      <Modal isOpen={true} onClose={() => {}} size="lg">
        <div>Content</div>
      </Modal>
    );

    expect(container.querySelector('.max-w-2xl')).toBeInTheDocument();
  });

  it('applies extra large size styling', () => {
    const { container } = render(
      <Modal isOpen={true} onClose={() => {}} size="xl">
        <div>Content</div>
      </Modal>
    );

    expect(container.querySelector('.max-w-4xl')).toBeInTheDocument();
  });
});

describe('ModalHeader', () => {
  it('renders children', () => {
    render(<ModalHeader>Header content</ModalHeader>);
    expect(screen.getByText('Header content')).toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(<ModalHeader className="custom-header">Header</ModalHeader>);
    expect(container.firstChild).toHaveClass('custom-header');
  });
});

describe('ModalBody', () => {
  it('renders children', () => {
    render(<ModalBody>Body content</ModalBody>);
    expect(screen.getByText('Body content')).toBeInTheDocument();
  });

  it('applies custom className', () => {
    const { container } = render(<ModalBody className="custom-body">Body</ModalBody>);
    expect(container.firstChild).toHaveClass('custom-body');
  });
});

describe('ModalFooter', () => {
  it('renders children', () => {
    render(<ModalFooter>Footer content</ModalFooter>);
    expect(screen.getByText('Footer content')).toBeInTheDocument();
  });

  it('has flex layout for buttons', () => {
    const { container } = render(<ModalFooter>Footer</ModalFooter>);
    expect(container.firstChild).toHaveClass('flex', 'justify-end');
  });

  it('has border top', () => {
    const { container } = render(<ModalFooter>Footer</ModalFooter>);
    expect(container.firstChild).toHaveClass('border-t', 'border-gray-200');
  });

  it('applies custom className', () => {
    const { container } = render(<ModalFooter className="custom-footer">Footer</ModalFooter>);
    expect(container.firstChild).toHaveClass('custom-footer');
  });
});
