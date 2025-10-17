import { create } from 'zustand'

type Theme = 'light' | 'dark' | 'system'
type Language = 'ko' | 'en'

interface UiState {
  theme: Theme
  language: Language
  sidebarOpen: boolean
  searchOpen: boolean
  mobileMenuOpen: boolean
  isLoading: boolean
  toast: {
    message: string
    type: 'success' | 'error' | 'info' | 'warning'
    visible: boolean
  } | null
}

interface UiActions {
  setTheme: (theme: Theme) => void
  setLanguage: (language: Language) => void
  toggleSidebar: () => void
  setSidebarOpen: (open: boolean) => void
  toggleSearch: () => void
  setSearchOpen: (open: boolean) => void
  toggleMobileMenu: () => void
  setMobileMenuOpen: (open: boolean) => void
  setLoading: (loading: boolean) => void
  showToast: (message: string, type?: 'success' | 'error' | 'info' | 'warning') => void
  hideToast: () => void
}

type UiStore = UiState & UiActions

export const useUiStore = create<UiStore>((set) => ({
  // State
  theme: 'system',
  language: 'ko',
  sidebarOpen: true,
  searchOpen: false,
  mobileMenuOpen: false,
  isLoading: false,
  toast: null,

  // Actions
  setTheme: (theme) => set({ theme }),
  setLanguage: (language) => set({ language }),

  toggleSidebar: () => set((state) => ({ sidebarOpen: !state.sidebarOpen })),
  setSidebarOpen: (open) => set({ sidebarOpen: open }),

  toggleSearch: () => set((state) => ({ searchOpen: !state.searchOpen })),
  setSearchOpen: (open) => set({ searchOpen: open }),

  toggleMobileMenu: () => set((state) => ({ mobileMenuOpen: !state.mobileMenuOpen })),
  setMobileMenuOpen: (open) => set({ mobileMenuOpen: open }),

  setLoading: (loading) => set({ isLoading: loading }),

  showToast: (message, type = 'info') =>
    set({ toast: { message, type, visible: true } }),

  hideToast: () =>
    set((state) => (state.toast ? { toast: { ...state.toast, visible: false } } : {})),
}))
