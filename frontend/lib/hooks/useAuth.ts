import { useCallback, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { useAuthStore } from '@/lib/store/authStore'
import { authApi } from '@/lib/api/auth'
import {
  LoginRequest,
  RegisterRequest,
  UpdateProfileRequest,
} from '@/types/user'

export function useAuth() {
  const router = useRouter()
  const {
    user,
    tokens,
    isAuthenticated,
    isAdmin,
    isLoading,
    login: loginStore,
    logout: logoutStore,
    updateProfile: updateProfileStore,
    setLoading,
  } = useAuthStore()

  /**
   * Login user
   */
  const login = useCallback(
    async (credentials: LoginRequest) => {
      try {
        setLoading(true)
        const response = await authApi.login(credentials)

        loginStore(response.user, {
          access_token: response.access_token,
          refresh_token: response.refresh_token,
        })

        return { success: true, data: response }
      } catch (error: any) {
        const message = error.response?.data?.error || 'Login failed'
        return { success: false, error: message }
      } finally {
        setLoading(false)
      }
    },
    [loginStore, setLoading]
  )

  /**
   * Register new user
   */
  const register = useCallback(
    async (data: RegisterRequest) => {
      try {
        setLoading(true)
        const response = await authApi.register(data)
        return { success: true, data: response }
      } catch (error: any) {
        const message = error.response?.data?.error || 'Registration failed'
        return { success: false, error: message }
      } finally {
        setLoading(false)
      }
    },
    [setLoading]
  )

  /**
   * Logout user
   */
  const logout = useCallback(async () => {
    try {
      setLoading(true)
      await authApi.logout()
    } catch (error) {
      // Continue with logout even if API call fails
      console.error('Logout error:', error)
    } finally {
      logoutStore()
      setLoading(false)
      router.push('/login')
    }
  }, [logoutStore, router, setLoading])

  /**
   * Update user profile
   */
  const updateProfile = useCallback(
    async (data: UpdateProfileRequest) => {
      try {
        setLoading(true)
        const response = await authApi.updateProfile(data)
        updateProfileStore(response.user)
        return { success: true, data: response }
      } catch (error: any) {
        const message = error.response?.data?.error || 'Profile update failed'
        return { success: false, error: message }
      } finally {
        setLoading(false)
      }
    },
    [updateProfileStore, setLoading]
  )

  /**
   * Check and restore authentication on mount
   */
  const checkAuth = useCallback(async () => {
    if (!tokens?.access_token) return

    try {
      const response = await authApi.getProfile()
      if (response.user) {
        updateProfileStore(response.user)
      }
    } catch (error) {
      // If profile fetch fails, logout
      logoutStore()
    }
  }, [tokens, updateProfileStore, logoutStore])

  useEffect(() => {
    checkAuth()
  }, []) // Run once on mount

  return {
    user,
    tokens,
    isAuthenticated,
    isAdmin,
    isLoading,
    login,
    register,
    logout,
    updateProfile,
    checkAuth,
  }
}
