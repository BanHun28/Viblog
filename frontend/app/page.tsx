'use client'

import { MainLayout } from '@/components/layout/MainLayout'
import { useAuth } from '@/lib/hooks/useAuth'
import { useEffect, useState } from 'react'
import { postsApi } from '@/lib/api/posts'
import { Post } from '@/types/post'
import { PostCard } from '@/components/post/PostCard'
import { Spinner } from '@/components/ui/Spinner'
import { Button } from '@/components/ui/Button'
import { Container } from '@/components/ui/Container'
import Link from 'next/link'
import { ArrowRight } from 'lucide-react'

export default function Home() {
  const { user, isAuthenticated } = useAuth()
  const [recentPosts, setRecentPosts] = useState<Post[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadRecentPosts()
  }, [])

  const loadRecentPosts = async () => {
    try {
      const response = await postsApi.getPosts({ page: 1, limit: 6 })
      setRecentPosts(response.data)
    } catch (err) {
      console.error('Failed to load recent posts:', err)
    } finally {
      setLoading(false)
    }
  }

  return (
    <MainLayout>
      {/* Hero Section */}
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 dark:from-blue-700 dark:to-purple-700 text-white py-16 mb-12">
        <Container>
          <div className="text-center">
            <h1 className="text-5xl font-bold mb-4">
              {isAuthenticated ? `Welcome back, ${user?.nickname}!` : 'Welcome to Viblog'}
            </h1>
            <p className="text-xl mb-8 text-blue-100">
              {isAuthenticated
                ? 'Discover amazing stories and share your thoughts'
                : 'Your personal blogging platform to share ideas and connect'}
            </p>
            <div className="flex justify-center gap-4">
              <Link href="/posts">
                <Button size="lg" variant="secondary" className="shadow-lg">
                  Browse All Posts
                  <ArrowRight size={20} className="ml-2" />
                </Button>
              </Link>
              {!isAuthenticated && (
                <Link href="/register">
                  <Button size="lg" variant="outline" className="text-white border-white hover:bg-white hover:text-blue-600 shadow-lg">
                    Get Started
                  </Button>
                </Link>
              )}
            </div>
          </div>
        </Container>
      </div>

      {/* Recent Posts Section */}
      <Container className="pb-12">
        <div className="mb-8">
          <h2 className="text-3xl font-bold text-gray-900 dark:text-white mb-2">
            Recent Posts
          </h2>
          <p className="text-gray-600 dark:text-gray-400">
            Check out the latest articles from our community
          </p>
        </div>

        {loading ? (
          <div className="flex justify-center py-12">
            <Spinner size="lg" />
          </div>
        ) : recentPosts.length > 0 ? (
          <>
            <div className="grid gap-6 mb-8">
              {recentPosts.map((post) => (
                <PostCard key={post.id} post={post} />
              ))}
            </div>
            <div className="text-center">
              <Link href="/posts">
                <Button variant="outline" size="lg">
                  View All Posts
                  <ArrowRight size={20} className="ml-2" />
                </Button>
              </Link>
            </div>
          </>
        ) : (
          <div className="text-center py-12 text-gray-500 dark:text-gray-400">
            No posts yet. Be the first to create one!
          </div>
        )}
      </Container>
    </MainLayout>
  )
}
