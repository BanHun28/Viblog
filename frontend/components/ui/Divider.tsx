import React from 'react';

export interface DividerProps extends React.HTMLAttributes<HTMLDivElement> {
  orientation?: 'horizontal' | 'vertical';
  text?: string;
}

export function Divider({
  orientation = 'horizontal',
  text,
  className = '',
  ...props
}: DividerProps) {
  if (orientation === 'vertical') {
    return (
      <div
        className={`inline-block h-full w-px bg-gray-300 ${className}`}
        role="separator"
        aria-orientation="vertical"
        {...props}
      />
    );
  }

  if (text) {
    return (
      <div className={`relative ${className}`} role="separator" aria-orientation="horizontal" {...props}>
        <div className="absolute inset-0 flex items-center">
          <div className="w-full border-t border-gray-300" />
        </div>
        <div className="relative flex justify-center text-sm">
          <span className="px-2 bg-white text-gray-500">{text}</span>
        </div>
      </div>
    );
  }

  return (
    <hr
      className={`border-t border-gray-300 ${className}`}
      role="separator"
      aria-orientation="horizontal"
      {...props}
    />
  );
}
