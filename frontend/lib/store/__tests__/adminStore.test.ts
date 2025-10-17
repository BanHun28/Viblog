import { describe, it, expect, beforeEach } from 'vitest'
import { useAdminStore } from '../adminStore'
import {
  DashboardStats,
  TrendData,
  PopularPost,
  PopularTag,
  RecentComment,
} from '@/types/admin'

describe('adminStore', () => {
  const mockStats: DashboardStats = {
    totalPosts: 100,
    totalComments: 500,
    totalUsers: 50,
    totalViews: 10000,
    totalLikes: 1000,
    recentPosts: 10,
    recentComments: 25,
    recentUsers: 5,
  }

  const mockTrends: TrendData[] = [
    { date: '2024-01-01', views: 100, likes: 10 },
    { date: '2024-01-02', views: 150, likes: 15 },
    { date: '2024-01-03', views: 200, likes: 20 },
  ]

  const mockPopularPosts: PopularPost[] = [
    {
      id: '1',
      title: 'Popular Post 1',
      viewCount: 1000,
      likeCount: 100,
      commentCount: 50,
      publishedAt: '2024-01-01',
    },
    {
      id: '2',
      title: 'Popular Post 2',
      viewCount: 800,
      likeCount: 80,
      commentCount: 40,
      publishedAt: '2024-01-02',
    },
  ]

  const mockPopularTags: PopularTag[] = [
    { id: '1', name: 'JavaScript', postCount: 50 },
    { id: '2', name: 'React', postCount: 40 },
  ]

  const mockRecentComments: RecentComment[] = [
    {
      id: '1',
      content: 'Great post!',
      postId: 'post1',
      postTitle: 'Test Post',
      authorNickname: 'user1',
      isAnonymous: false,
      createdAt: '2024-01-01',
    },
    {
      id: '2',
      content: 'Nice article',
      postId: 'post2',
      postTitle: 'Another Post',
      authorNickname: 'Anonymous',
      isAnonymous: true,
      createdAt: '2024-01-02',
    },
  ]

  beforeEach(() => {
    useAdminStore.setState({
      stats: null,
      trends: [],
      popularPosts: [],
      popularTags: [],
      recentComments: [],
      isLoading: false,
      error: null,
    })
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = useAdminStore.getState()

      expect(state.stats).toBeNull()
      expect(state.trends).toEqual([])
      expect(state.popularPosts).toEqual([])
      expect(state.popularTags).toEqual([])
      expect(state.recentComments).toEqual([])
      expect(state.isLoading).toBe(false)
      expect(state.error).toBeNull()
    })
  })

  describe('Stats Management', () => {
    it('should set stats', () => {
      useAdminStore.getState().setStats(mockStats)

      const state = useAdminStore.getState()
      expect(state.stats).toEqual(mockStats)
    })

    it('should update stats', () => {
      useAdminStore.getState().setStats(mockStats)

      const updatedStats: DashboardStats = {
        ...mockStats,
        totalPosts: 150,
      }

      useAdminStore.getState().setStats(updatedStats)

      const state = useAdminStore.getState()
      expect(state.stats?.totalPosts).toBe(150)
      expect(state.stats?.totalComments).toBe(500)
    })
  })

  describe('Trends Management', () => {
    it('should set trends', () => {
      useAdminStore.getState().setTrends(mockTrends)

      const state = useAdminStore.getState()
      expect(state.trends).toEqual(mockTrends)
      expect(state.trends).toHaveLength(3)
    })

    it('should replace existing trends', () => {
      useAdminStore.getState().setTrends(mockTrends)

      const newTrends: TrendData[] = [
        { date: '2024-01-04', views: 250, likes: 25 },
      ]

      useAdminStore.getState().setTrends(newTrends)

      const state = useAdminStore.getState()
      expect(state.trends).toEqual(newTrends)
      expect(state.trends).toHaveLength(1)
    })
  })

  describe('Popular Posts Management', () => {
    it('should set popular posts', () => {
      useAdminStore.getState().setPopularPosts(mockPopularPosts)

      const state = useAdminStore.getState()
      expect(state.popularPosts).toEqual(mockPopularPosts)
      expect(state.popularPosts).toHaveLength(2)
    })

    it('should replace existing popular posts', () => {
      useAdminStore.getState().setPopularPosts(mockPopularPosts)

      const newPosts: PopularPost[] = [
        {
          id: '3',
          title: 'New Popular Post',
          viewCount: 2000,
          likeCount: 200,
          commentCount: 100,
          publishedAt: '2024-01-03',
        },
      ]

      useAdminStore.getState().setPopularPosts(newPosts)

      const state = useAdminStore.getState()
      expect(state.popularPosts).toEqual(newPosts)
      expect(state.popularPosts).toHaveLength(1)
    })
  })

  describe('Popular Tags Management', () => {
    it('should set popular tags', () => {
      useAdminStore.getState().setPopularTags(mockPopularTags)

      const state = useAdminStore.getState()
      expect(state.popularTags).toEqual(mockPopularTags)
      expect(state.popularTags).toHaveLength(2)
    })

    it('should replace existing popular tags', () => {
      useAdminStore.getState().setPopularTags(mockPopularTags)

      const newTags: PopularTag[] = [
        { id: '3', name: 'TypeScript', postCount: 60 },
      ]

      useAdminStore.getState().setPopularTags(newTags)

      const state = useAdminStore.getState()
      expect(state.popularTags).toEqual(newTags)
      expect(state.popularTags).toHaveLength(1)
    })
  })

  describe('Recent Comments Management', () => {
    it('should set recent comments', () => {
      useAdminStore.getState().setRecentComments(mockRecentComments)

      const state = useAdminStore.getState()
      expect(state.recentComments).toEqual(mockRecentComments)
      expect(state.recentComments).toHaveLength(2)
    })

    it('should handle anonymous comments', () => {
      useAdminStore.getState().setRecentComments(mockRecentComments)

      const state = useAdminStore.getState()
      expect(state.recentComments[1].isAnonymous).toBe(true)
      expect(state.recentComments[1].authorNickname).toBe('Anonymous')
    })
  })

  describe('Dashboard Data Management', () => {
    it('should set all dashboard data at once', () => {
      useAdminStore.getState().setDashboardData({
        stats: mockStats,
        trends: mockTrends,
        popularPosts: mockPopularPosts,
        popularTags: mockPopularTags,
        recentComments: mockRecentComments,
      })

      const state = useAdminStore.getState()
      expect(state.stats).toEqual(mockStats)
      expect(state.trends).toEqual(mockTrends)
      expect(state.popularPosts).toEqual(mockPopularPosts)
      expect(state.popularTags).toEqual(mockPopularTags)
      expect(state.recentComments).toEqual(mockRecentComments)
    })

    it('should replace all existing data', () => {
      // Set initial data
      useAdminStore.getState().setDashboardData({
        stats: mockStats,
        trends: mockTrends,
        popularPosts: mockPopularPosts,
        popularTags: mockPopularTags,
        recentComments: mockRecentComments,
      })

      // Set new data
      const newStats: DashboardStats = {
        ...mockStats,
        totalPosts: 200,
      }

      useAdminStore.getState().setDashboardData({
        stats: newStats,
        trends: [],
        popularPosts: [],
        popularTags: [],
        recentComments: [],
      })

      const state = useAdminStore.getState()
      expect(state.stats?.totalPosts).toBe(200)
      expect(state.trends).toEqual([])
      expect(state.popularPosts).toEqual([])
    })
  })

  describe('Clear Dashboard', () => {
    it('should clear all dashboard data', () => {
      // Set data first
      useAdminStore.getState().setDashboardData({
        stats: mockStats,
        trends: mockTrends,
        popularPosts: mockPopularPosts,
        popularTags: mockPopularTags,
        recentComments: mockRecentComments,
      })

      // Clear
      useAdminStore.getState().clearDashboard()

      const state = useAdminStore.getState()
      expect(state.stats).toBeNull()
      expect(state.trends).toEqual([])
      expect(state.popularPosts).toEqual([])
      expect(state.popularTags).toEqual([])
      expect(state.recentComments).toEqual([])
    })

    it('should not affect loading and error states', () => {
      useAdminStore.getState().setLoading(true)
      useAdminStore.getState().setError('Test error')

      useAdminStore.getState().clearDashboard()

      const state = useAdminStore.getState()
      expect(state.isLoading).toBe(true)
      expect(state.error).toBe('Test error')
    })
  })

  describe('Loading and Error States', () => {
    it('should set loading state', () => {
      useAdminStore.getState().setLoading(true)
      expect(useAdminStore.getState().isLoading).toBe(true)

      useAdminStore.getState().setLoading(false)
      expect(useAdminStore.getState().isLoading).toBe(false)
    })

    it('should set error state', () => {
      useAdminStore.getState().setError('Test error')
      expect(useAdminStore.getState().error).toBe('Test error')

      useAdminStore.getState().setError(null)
      expect(useAdminStore.getState().error).toBeNull()
    })

    it('should handle error while loading', () => {
      useAdminStore.getState().setLoading(true)
      useAdminStore.getState().setError('Load failed')

      const state = useAdminStore.getState()
      expect(state.isLoading).toBe(true)
      expect(state.error).toBe('Load failed')
    })
  })

  describe('Complex Scenarios', () => {
    it('should handle full dashboard workflow', () => {
      // Start loading
      useAdminStore.getState().setLoading(true)

      // Load data
      useAdminStore.getState().setDashboardData({
        stats: mockStats,
        trends: mockTrends,
        popularPosts: mockPopularPosts,
        popularTags: mockPopularTags,
        recentComments: mockRecentComments,
      })

      // Stop loading
      useAdminStore.getState().setLoading(false)

      const state = useAdminStore.getState()
      expect(state.stats).toEqual(mockStats)
      expect(state.isLoading).toBe(false)
      expect(state.error).toBeNull()
    })

    it('should handle error workflow', () => {
      // Start loading
      useAdminStore.getState().setLoading(true)

      // Error occurs
      useAdminStore.getState().setError('Failed to load')
      useAdminStore.getState().setLoading(false)

      const state = useAdminStore.getState()
      expect(state.stats).toBeNull()
      expect(state.isLoading).toBe(false)
      expect(state.error).toBe('Failed to load')
    })

    it('should handle retry after error', () => {
      // First attempt fails
      useAdminStore.getState().setLoading(true)
      useAdminStore.getState().setError('Failed')
      useAdminStore.getState().setLoading(false)

      // Clear error for retry
      useAdminStore.getState().setError(null)
      useAdminStore.getState().setLoading(true)

      // Success
      useAdminStore.getState().setDashboardData({
        stats: mockStats,
        trends: mockTrends,
        popularPosts: mockPopularPosts,
        popularTags: mockPopularTags,
        recentComments: mockRecentComments,
      })
      useAdminStore.getState().setLoading(false)

      const state = useAdminStore.getState()
      expect(state.stats).toEqual(mockStats)
      expect(state.error).toBeNull()
    })
  })
})
