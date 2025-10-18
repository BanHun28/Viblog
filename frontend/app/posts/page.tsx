'use client'

import { useEffect, useState } from 'react'
import { useSearchParams } from 'next/navigation'
import { postsApi } from '@/lib/api/posts'
import { Post } from '@/types/post'
import { PostCard } from '@/components/post/PostCard'
import { Pagination } from '@/components/ui/Pagination'
import { Spinner } from '@/components/ui/Spinner'
import { Alert } from '@/components/ui/Alert'
import { Container } from '@/components/ui/Container'

export default function PostsPage() {
  const searchParams = useSearchParams()
  const [posts, setPosts] = useState<Post[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [currentPage, setCurrentPage] = useState(1)
  const [totalPages, setTotalPages] = useState(1)
  const [total, setTotal] = useState(0)

  const limit = 20

  useEffect(() => {
    const page = parseInt(searchParams.get('page') || '1')
    setCurrentPage(page)
    loadPosts(page)
  }, [searchParams])

  const loadPosts = async (page: number) => {
    try {
      setLoading(true)
      setError(null)

      const response = await postsApi.getPosts({ page, limit })
      setPosts(response.data)
      setTotalPages(response.meta.totalPages)
      setTotal(response.meta.total)
    } catch (err: any) {
      setError(err.message || 'Failed to load posts')
    } finally {
      setLoading(false)
    }
  }

  const handlePageChange = (page: number) => {
    window.history.pushState({}, '', `/posts?page=${page}`)
    setCurrentPage(page)
    loadPosts(page)
  }

  if (loading && posts.length === 0) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <Spinner size="lg" />
      </div>
    )
  }

  return (
    <Container className="py-8">
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-2">
          All Posts
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          {total} posts found
        </p>
      </div>

      {error && (
        <Alert variant="error" className="mb-6">
          {error}
        </Alert>
      )}

      {posts.length === 0 && !loading ? (
        <Alert variant="info">No posts found</Alert>
      ) : (
        <>
          <div className="grid gap-6 mb-8">
            {posts.map((post) => (
              <PostCard key={post.id} post={post} />
            ))}
          </div>

          {totalPages > 1 && (
            <Pagination
              currentPage={currentPage}
              totalPages={totalPages}
              onPageChange={handlePageChange}
            />
          )}
        </>
      )}
    </Container>
  )
}
