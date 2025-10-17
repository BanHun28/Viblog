'use client'

import Link from 'next/link'
import { useState } from 'react'
import { useAuth } from '@/lib/hooks/useAuth'
import { ThemeToggle } from '@/components/ui/ThemeToggle'

export function Header() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)
  const { isAuthenticated, user, isAdmin } = useAuth()

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-white dark:bg-gray-900 shadow-sm dark:border-gray-800">
      <div className="container mx-auto px-4">
        <div className="flex h-16 items-center justify-between">
          {/* Logo */}
          <Link href="/" className="text-2xl font-bold text-gray-900 dark:text-white hover:text-gray-700 dark:hover:text-gray-300">
            Viblog
          </Link>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <Link href="/posts" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors">
              Posts
            </Link>
            <Link href="/categories" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors">
              Categories
            </Link>
            <Link href="/tags" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors">
              Tags
            </Link>
          </nav>

          {/* Search & User Actions */}
          <div className="flex items-center space-x-4">
            {/* Search Bar */}
            <div className="hidden md:block">
              <form action="/search" method="GET" className="relative">
                <input
                  type="text"
                  name="q"
                  placeholder="Search..."
                  className="w-64 px-4 py-2 pr-10 text-sm border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100 placeholder-gray-400 dark:placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent"
                />
                <button
                  type="submit"
                  className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-400 dark:text-gray-500 hover:text-gray-600 dark:hover:text-gray-300"
                >
                  <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </button>
              </form>
            </div>

            {/* Theme Toggle */}
            <ThemeToggle />

            {/* User Menu */}
            {isAuthenticated && user ? (
              <div className="flex items-center space-x-4">
                {/* Admin Link */}
                {isAdmin && (
                  <Link
                    href="/admin"
                    className="hidden lg:flex items-center space-x-1 text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300"
                  >
                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    <span className="text-sm font-medium">Admin</span>
                  </Link>
                )}

                {/* Notifications */}
                <Link href="/notifications" className="relative text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white">
                  <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                  </svg>
                </Link>

                {/* Bookmarks */}
                <Link href="/bookmarks" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white">
                  <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                  </svg>
                </Link>

                {/* Profile */}
                <Link href="/profile" className="flex items-center space-x-2 text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white group">
                  {user.avatar_url ? (
                    <img
                      src={user.avatar_url}
                      alt={user.nickname}
                      className="w-8 h-8 rounded-full object-cover border-2 border-transparent group-hover:border-blue-500"
                    />
                  ) : (
                    <div className="w-8 h-8 bg-blue-500 dark:bg-blue-600 rounded-full flex items-center justify-center border-2 border-transparent group-hover:border-blue-600 dark:group-hover:border-blue-400">
                      <span className="text-sm font-semibold text-white">
                        {user.nickname.charAt(0).toUpperCase()}
                      </span>
                    </div>
                  )}
                  <span className="hidden lg:inline text-sm font-medium">
                    {user.nickname}
                  </span>
                </Link>
              </div>
            ) : (
              <div className="hidden md:flex items-center space-x-4">
                <Link
                  href="/login"
                  className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white transition-colors"
                >
                  Login
                </Link>
                <Link
                  href="/register"
                  className="px-4 py-2 bg-blue-600 dark:bg-blue-500 text-white rounded-lg hover:bg-blue-700 dark:hover:bg-blue-600 transition-colors"
                >
                  Sign Up
                </Link>
              </div>
            )}

            {/* Mobile Menu Button */}
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="md:hidden text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                {isMenuOpen ? (
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                ) : (
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                )}
              </svg>
            </button>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMenuOpen && (
          <div className="md:hidden py-4 border-t dark:border-gray-800">
            {/* Mobile Search */}
            <form action="/search" method="GET" className="mb-4">
              <input
                type="text"
                name="q"
                placeholder="Search..."
                className="w-full px-4 py-2 text-sm border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100 placeholder-gray-400 dark:placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400"
              />
            </form>

            {/* Mobile Navigation */}
            <nav className="flex flex-col space-y-3">
              <Link href="/posts" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                Posts
              </Link>
              <Link href="/categories" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                Categories
              </Link>
              <Link href="/tags" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                Tags
              </Link>

              {isAuthenticated && user ? (
                <>
                  {isAdmin && (
                    <Link href="/admin" className="text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300 py-2 font-semibold">
                      Admin Dashboard
                    </Link>
                  )}
                  <Link href="/notifications" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                    Notifications
                  </Link>
                  <Link href="/bookmarks" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                    Bookmarks
                  </Link>
                  <Link href="/profile" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                    Profile ({user.nickname})
                  </Link>
                </>
              ) : (
                <>
                  <Link href="/login" className="text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white py-2">
                    Login
                  </Link>
                  <Link href="/register" className="text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-500 py-2 font-semibold">
                    Sign Up
                  </Link>
                </>
              )}
            </nav>
          </div>
        )}
      </div>
    </header>
  )
}
