import { describe, it, expect, beforeEach } from 'vitest'
import { useUiStore } from '../uiStore'

describe('uiStore', () => {
  beforeEach(() => {
    // Reset store before each test
    useUiStore.setState({
      theme: 'system',
      language: 'ko',
      sidebarOpen: true,
      searchOpen: false,
      mobileMenuOpen: false,
      isLoading: false,
      toast: null,
    })
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = useUiStore.getState()

      expect(state.theme).toBe('system')
      expect(state.language).toBe('ko')
      expect(state.sidebarOpen).toBe(true)
      expect(state.searchOpen).toBe(false)
      expect(state.mobileMenuOpen).toBe(false)
      expect(state.isLoading).toBe(false)
      expect(state.toast).toBeNull()
    })
  })

  describe('Theme Management', () => {
    it('should set theme', () => {
      useUiStore.getState().setTheme('dark')
      expect(useUiStore.getState().theme).toBe('dark')

      useUiStore.getState().setTheme('light')
      expect(useUiStore.getState().theme).toBe('light')
    })
  })

  describe('Language Management', () => {
    it('should set language', () => {
      useUiStore.getState().setLanguage('en')
      expect(useUiStore.getState().language).toBe('en')

      useUiStore.getState().setLanguage('ko')
      expect(useUiStore.getState().language).toBe('ko')
    })
  })

  describe('Sidebar Management', () => {
    it('should toggle sidebar', () => {
      const initialState = useUiStore.getState().sidebarOpen

      useUiStore.getState().toggleSidebar()
      expect(useUiStore.getState().sidebarOpen).toBe(!initialState)

      useUiStore.getState().toggleSidebar()
      expect(useUiStore.getState().sidebarOpen).toBe(initialState)
    })

    it('should set sidebar open state', () => {
      useUiStore.getState().setSidebarOpen(false)
      expect(useUiStore.getState().sidebarOpen).toBe(false)

      useUiStore.getState().setSidebarOpen(true)
      expect(useUiStore.getState().sidebarOpen).toBe(true)
    })
  })

  describe('Search Management', () => {
    it('should toggle search', () => {
      const initialState = useUiStore.getState().searchOpen

      useUiStore.getState().toggleSearch()
      expect(useUiStore.getState().searchOpen).toBe(!initialState)

      useUiStore.getState().toggleSearch()
      expect(useUiStore.getState().searchOpen).toBe(initialState)
    })

    it('should set search open state', () => {
      useUiStore.getState().setSearchOpen(true)
      expect(useUiStore.getState().searchOpen).toBe(true)

      useUiStore.getState().setSearchOpen(false)
      expect(useUiStore.getState().searchOpen).toBe(false)
    })
  })

  describe('Mobile Menu Management', () => {
    it('should toggle mobile menu', () => {
      const initialState = useUiStore.getState().mobileMenuOpen

      useUiStore.getState().toggleMobileMenu()
      expect(useUiStore.getState().mobileMenuOpen).toBe(!initialState)

      useUiStore.getState().toggleMobileMenu()
      expect(useUiStore.getState().mobileMenuOpen).toBe(initialState)
    })

    it('should set mobile menu open state', () => {
      useUiStore.getState().setMobileMenuOpen(true)
      expect(useUiStore.getState().mobileMenuOpen).toBe(true)

      useUiStore.getState().setMobileMenuOpen(false)
      expect(useUiStore.getState().mobileMenuOpen).toBe(false)
    })
  })

  describe('Loading State', () => {
    it('should set loading state', () => {
      useUiStore.getState().setLoading(true)
      expect(useUiStore.getState().isLoading).toBe(true)

      useUiStore.getState().setLoading(false)
      expect(useUiStore.getState().isLoading).toBe(false)
    })
  })

  describe('Toast Management', () => {
    it('should show toast with default type', () => {
      useUiStore.getState().showToast('Test message')

      const state = useUiStore.getState()
      expect(state.toast).toEqual({
        message: 'Test message',
        type: 'info',
        visible: true,
      })
    })

    it('should show toast with specific type', () => {
      useUiStore.getState().showToast('Success message', 'success')

      const state = useUiStore.getState()
      expect(state.toast).toEqual({
        message: 'Success message',
        type: 'success',
        visible: true,
      })
    })

    it('should show error toast', () => {
      useUiStore.getState().showToast('Error message', 'error')

      const state = useUiStore.getState()
      expect(state.toast?.type).toBe('error')
      expect(state.toast?.message).toBe('Error message')
    })

    it('should show warning toast', () => {
      useUiStore.getState().showToast('Warning message', 'warning')

      const state = useUiStore.getState()
      expect(state.toast?.type).toBe('warning')
      expect(state.toast?.message).toBe('Warning message')
    })

    it('should hide toast', () => {
      useUiStore.getState().showToast('Test message')
      useUiStore.getState().hideToast()

      const state = useUiStore.getState()
      expect(state.toast?.visible).toBe(false)
    })

    it('should preserve toast data when hiding', () => {
      useUiStore.getState().showToast('Test message', 'success')
      useUiStore.getState().hideToast()

      const state = useUiStore.getState()
      expect(state.toast?.message).toBe('Test message')
      expect(state.toast?.type).toBe('success')
      expect(state.toast?.visible).toBe(false)
    })

    it('should handle hiding toast when no toast exists', () => {
      useUiStore.getState().hideToast()

      const state = useUiStore.getState()
      expect(state.toast).toBeNull()
    })
  })

  describe('Multiple State Changes', () => {
    it('should handle multiple state changes', () => {
      useUiStore.getState().setTheme('dark')
      useUiStore.getState().setLanguage('en')
      useUiStore.getState().setSidebarOpen(false)
      useUiStore.getState().setSearchOpen(true)
      useUiStore.getState().showToast('Test', 'success')

      const state = useUiStore.getState()
      expect(state.theme).toBe('dark')
      expect(state.language).toBe('en')
      expect(state.sidebarOpen).toBe(false)
      expect(state.searchOpen).toBe(true)
      expect(state.toast?.message).toBe('Test')
    })
  })
})
