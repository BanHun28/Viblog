'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { useAuth } from './useAuth'

interface UseAuthGuardOptions {
  requireAuth?: boolean
  requireAdmin?: boolean
  redirectTo?: string
}

export function useAuthGuard(options: UseAuthGuardOptions = {}) {
  const {
    requireAuth = true,
    requireAdmin = false,
    redirectTo = '/login',
  } = options

  const { isAuthenticated, isAdmin, isLoading } = useAuth()
  const router = useRouter()

  useEffect(() => {
    if (isLoading) return

    if (requireAuth && !isAuthenticated) {
      router.push(redirectTo)
    } else if (requireAdmin && !isAdmin) {
      router.push('/')
    }
  }, [isAuthenticated, isAdmin, isLoading, requireAuth, requireAdmin, redirectTo, router])

  return {
    isLoading,
    isAuthorized: requireAdmin ? isAdmin : isAuthenticated,
  }
}
