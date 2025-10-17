import { describe, it, expect, beforeEach } from 'vitest'
import { useAuthStore } from '../authStore'
import { User, AuthTokens } from '@/types/user'

describe('authStore', () => {
  beforeEach(() => {
    // Reset store before each test
    useAuthStore.setState({
      user: null,
      tokens: null,
      isAuthenticated: false,
      isAdmin: false,
      isLoading: false,
    })
    localStorage.clear()
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = useAuthStore.getState()

      expect(state.user).toBeNull()
      expect(state.tokens).toBeNull()
      expect(state.isAuthenticated).toBe(false)
      expect(state.isAdmin).toBe(false)
      expect(state.isLoading).toBe(false)
    })
  })

  describe('setUser', () => {
    it('should set user and update authentication status', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }

      useAuthStore.getState().setUser(mockUser)
      const state = useAuthStore.getState()

      expect(state.user).toEqual(mockUser)
      expect(state.isAuthenticated).toBe(true)
      expect(state.isAdmin).toBe(false)
    })

    it('should set isAdmin to true for admin user', () => {
      const mockAdmin: User = {
        id: '1',
        email: 'admin@example.com',
        nickname: 'admin',
        role: 'admin',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }

      useAuthStore.getState().setUser(mockAdmin)
      const state = useAuthStore.getState()

      expect(state.isAdmin).toBe(true)
    })

    it('should clear user when null is passed', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }

      useAuthStore.getState().setUser(mockUser)
      useAuthStore.getState().setUser(null)

      const state = useAuthStore.getState()
      expect(state.user).toBeNull()
      expect(state.isAuthenticated).toBe(false)
      expect(state.isAdmin).toBe(false)
    })
  })

  describe('setTokens', () => {
    it('should set tokens', () => {
      const mockTokens: AuthTokens = {
        accessToken: 'access-token',
        refreshToken: 'refresh-token',
      }

      useAuthStore.getState().setTokens(mockTokens)
      const state = useAuthStore.getState()

      expect(state.tokens).toEqual(mockTokens)
    })
  })

  describe('login', () => {
    it('should set user and tokens on login', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }
      const mockTokens: AuthTokens = {
        accessToken: 'access-token',
        refreshToken: 'refresh-token',
      }

      useAuthStore.getState().login(mockUser, mockTokens)
      const state = useAuthStore.getState()

      expect(state.user).toEqual(mockUser)
      expect(state.tokens).toEqual(mockTokens)
      expect(state.isAuthenticated).toBe(true)
      expect(state.isAdmin).toBe(false)
    })

    it('should set isAdmin for admin login', () => {
      const mockAdmin: User = {
        id: '1',
        email: 'admin@example.com',
        nickname: 'admin',
        role: 'admin',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }
      const mockTokens: AuthTokens = {
        accessToken: 'access-token',
        refreshToken: 'refresh-token',
      }

      useAuthStore.getState().login(mockAdmin, mockTokens)
      const state = useAuthStore.getState()

      expect(state.isAdmin).toBe(true)
    })
  })

  describe('logout', () => {
    it('should clear all authentication data', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }
      const mockTokens: AuthTokens = {
        accessToken: 'access-token',
        refreshToken: 'refresh-token',
      }

      useAuthStore.getState().login(mockUser, mockTokens)
      useAuthStore.getState().logout()

      const state = useAuthStore.getState()
      expect(state.user).toBeNull()
      expect(state.tokens).toBeNull()
      expect(state.isAuthenticated).toBe(false)
      expect(state.isAdmin).toBe(false)
    })
  })

  describe('updateProfile', () => {
    it('should update user profile', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }

      useAuthStore.getState().setUser(mockUser)
      useAuthStore.getState().updateProfile({ nickname: 'newname' })

      const state = useAuthStore.getState()
      expect(state.user?.nickname).toBe('newname')
      expect(state.user?.email).toBe('test@example.com')
    })

    it('should not update if user is null', () => {
      useAuthStore.getState().updateProfile({ nickname: 'newname' })

      const state = useAuthStore.getState()
      expect(state.user).toBeNull()
    })
  })

  describe('setLoading', () => {
    it('should set loading state', () => {
      useAuthStore.getState().setLoading(true)
      expect(useAuthStore.getState().isLoading).toBe(true)

      useAuthStore.getState().setLoading(false)
      expect(useAuthStore.getState().isLoading).toBe(false)
    })
  })

  describe('Persistence', () => {
    it('should persist user and tokens to localStorage', () => {
      const mockUser: User = {
        id: '1',
        email: 'test@example.com',
        nickname: 'testuser',
        role: 'user',
        createdAt: '2024-01-01',
        updatedAt: '2024-01-01',
      }
      const mockTokens: AuthTokens = {
        accessToken: 'access-token',
        refreshToken: 'refresh-token',
      }

      useAuthStore.getState().login(mockUser, mockTokens)

      // Check if data is in localStorage
      const stored = localStorage.getItem('auth-storage')
      expect(stored).toBeTruthy()

      if (stored) {
        const parsed = JSON.parse(stored)
        expect(parsed.state.user).toEqual(mockUser)
        expect(parsed.state.tokens).toEqual(mockTokens)
      }
    })
  })
})
