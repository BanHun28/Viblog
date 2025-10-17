import React from 'react';
import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { Pagination } from '../Pagination';

describe('Pagination', () => {
  const defaultProps = {
    currentPage: 1,
    totalPages: 10,
    onPageChange: vi.fn(),
  };

  it('renders pagination navigation', () => {
    render(<Pagination {...defaultProps} />);
    expect(screen.getByRole('navigation', { name: 'Pagination' })).toBeInTheDocument();
  });

  it('renders Previous and Next buttons', () => {
    render(<Pagination {...defaultProps} />);

    expect(screen.getByLabelText('Previous page')).toBeInTheDocument();
    expect(screen.getByLabelText('Next page')).toBeInTheDocument();
  });

  it('renders First and Last buttons by default', () => {
    render(<Pagination {...defaultProps} />);

    expect(screen.getByLabelText('First page')).toBeInTheDocument();
    expect(screen.getByLabelText('Last page')).toBeInTheDocument();
  });

  it('hides First and Last buttons when showFirstLast is false', () => {
    render(<Pagination {...defaultProps} showFirstLast={false} />);

    expect(screen.queryByLabelText('First page')).not.toBeInTheDocument();
    expect(screen.queryByLabelText('Last page')).not.toBeInTheDocument();
  });

  it('disables Previous button on first page', () => {
    render(<Pagination {...defaultProps} currentPage={1} />);
    expect(screen.getByLabelText('Previous page')).toBeDisabled();
  });

  it('disables Next button on last page', () => {
    render(<Pagination {...defaultProps} currentPage={10} totalPages={10} />);
    expect(screen.getByLabelText('Next page')).toBeDisabled();
  });

  it('disables First button on first page', () => {
    render(<Pagination {...defaultProps} currentPage={1} />);
    expect(screen.getByLabelText('First page')).toBeDisabled();
  });

  it('disables Last button on last page', () => {
    render(<Pagination {...defaultProps} currentPage={10} totalPages={10} />);
    expect(screen.getByLabelText('Last page')).toBeDisabled();
  });

  it('calls onPageChange when Previous is clicked', () => {
    const onPageChange = vi.fn();
    render(<Pagination {...defaultProps} currentPage={5} onPageChange={onPageChange} />);

    fireEvent.click(screen.getByLabelText('Previous page'));
    expect(onPageChange).toHaveBeenCalledWith(4);
  });

  it('calls onPageChange when Next is clicked', () => {
    const onPageChange = vi.fn();
    render(<Pagination {...defaultProps} currentPage={5} onPageChange={onPageChange} />);

    fireEvent.click(screen.getByLabelText('Next page'));
    expect(onPageChange).toHaveBeenCalledWith(6);
  });

  it('calls onPageChange when First is clicked', () => {
    const onPageChange = vi.fn();
    render(<Pagination {...defaultProps} currentPage={5} onPageChange={onPageChange} />);

    fireEvent.click(screen.getByLabelText('First page'));
    expect(onPageChange).toHaveBeenCalledWith(1);
  });

  it('calls onPageChange when Last is clicked', () => {
    const onPageChange = vi.fn();
    render(<Pagination {...defaultProps} currentPage={5} onPageChange={onPageChange} />);

    fireEvent.click(screen.getByLabelText('Last page'));
    expect(onPageChange).toHaveBeenCalledWith(10);
  });

  it('calls onPageChange when page number is clicked', () => {
    const onPageChange = vi.fn();
    render(<Pagination {...defaultProps} onPageChange={onPageChange} />);

    fireEvent.click(screen.getByLabelText('Page 3'));
    expect(onPageChange).toHaveBeenCalledWith(3);
  });

  it('highlights current page', () => {
    render(<Pagination {...defaultProps} currentPage={5} />);
    const currentPageButton = screen.getByLabelText('Page 5');

    expect(currentPageButton).toHaveClass('bg-blue-600', 'text-white');
    expect(currentPageButton).toHaveAttribute('aria-current', 'page');
  });

  it('renders all pages when totalPages <= maxVisible', () => {
    render(<Pagination {...defaultProps} totalPages={5} maxVisible={5} />);

    expect(screen.getByLabelText('Page 1')).toBeInTheDocument();
    expect(screen.getByLabelText('Page 2')).toBeInTheDocument();
    expect(screen.getByLabelText('Page 3')).toBeInTheDocument();
    expect(screen.getByLabelText('Page 4')).toBeInTheDocument();
    expect(screen.getByLabelText('Page 5')).toBeInTheDocument();
  });

  it('shows ellipsis for truncated pages', () => {
    render(<Pagination {...defaultProps} currentPage={1} totalPages={20} maxVisible={5} />);

    const ellipsis = screen.getAllByText('...');
    expect(ellipsis.length).toBeGreaterThan(0);
  });

  it('always shows first and last page numbers', () => {
    render(<Pagination {...defaultProps} currentPage={10} totalPages={20} maxVisible={5} />);

    expect(screen.getByLabelText('Page 1')).toBeInTheDocument();
    expect(screen.getByLabelText('Page 20')).toBeInTheDocument();
  });

  it('handles single page correctly', () => {
    render(<Pagination {...defaultProps} currentPage={1} totalPages={1} />);

    expect(screen.getByLabelText('Previous page')).toBeDisabled();
    expect(screen.getByLabelText('Next page')).toBeDisabled();
    expect(screen.getByLabelText('Page 1')).toBeInTheDocument();
  });

  it('adjusts page range around current page', () => {
    render(<Pagination {...defaultProps} currentPage={10} totalPages={20} maxVisible={5} />);

    // Should show pages around current page (10)
    expect(screen.getByLabelText('Page 10')).toBeInTheDocument();
  });
});
