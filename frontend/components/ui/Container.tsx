import React from 'react';

export interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
  size?: 'sm' | 'md' | 'lg' | 'xl' | 'full';
  padding?: boolean;
  children: React.ReactNode;
}

export function Container({
  size = 'lg',
  padding = true,
  className = '',
  children,
  ...props
}: ContainerProps) {
  const sizeStyles = {
    sm: 'max-w-2xl',
    md: 'max-w-4xl',
    lg: 'max-w-6xl',
    xl: 'max-w-7xl',
    full: 'max-w-full',
  };

  const paddingStyles = padding ? 'px-4 sm:px-6 lg:px-8' : '';

  return (
    <div
      className={`mx-auto ${sizeStyles[size]} ${paddingStyles} ${className}`}
      {...props}
    >
      {children}
    </div>
  );
}
