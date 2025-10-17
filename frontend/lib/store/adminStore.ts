import { create } from 'zustand'
import {
  DashboardStats,
  TrendData,
  PopularPost,
  PopularTag,
  RecentComment,
} from '@/types/admin'

interface AdminState {
  stats: DashboardStats | null
  trends: TrendData[]
  popularPosts: PopularPost[]
  popularTags: PopularTag[]
  recentComments: RecentComment[]
  isLoading: boolean
  error: string | null
}

interface AdminActions {
  setStats: (stats: DashboardStats) => void
  setTrends: (trends: TrendData[]) => void
  setPopularPosts: (posts: PopularPost[]) => void
  setPopularTags: (tags: PopularTag[]) => void
  setRecentComments: (comments: RecentComment[]) => void
  setDashboardData: (data: {
    stats: DashboardStats
    trends: TrendData[]
    popularPosts: PopularPost[]
    popularTags: PopularTag[]
    recentComments: RecentComment[]
  }) => void
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
  clearDashboard: () => void
}

type AdminStore = AdminState & AdminActions

export const useAdminStore = create<AdminStore>((set) => ({
  // State
  stats: null,
  trends: [],
  popularPosts: [],
  popularTags: [],
  recentComments: [],
  isLoading: false,
  error: null,

  // Actions
  setStats: (stats) => set({ stats }),

  setTrends: (trends) => set({ trends }),

  setPopularPosts: (posts) => set({ popularPosts: posts }),

  setPopularTags: (tags) => set({ popularTags: tags }),

  setRecentComments: (comments) => set({ recentComments: comments }),

  setDashboardData: (data) =>
    set({
      stats: data.stats,
      trends: data.trends,
      popularPosts: data.popularPosts,
      popularTags: data.popularTags,
      recentComments: data.recentComments,
    }),

  setLoading: (loading) => set({ isLoading: loading }),

  setError: (error) => set({ error }),

  clearDashboard: () =>
    set({
      stats: null,
      trends: [],
      popularPosts: [],
      popularTags: [],
      recentComments: [],
    }),
}))
