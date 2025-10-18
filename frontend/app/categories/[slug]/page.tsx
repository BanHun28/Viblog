'use client'

import { useEffect, useState } from 'react'
import { useSearchParams } from 'next/navigation'
import { postsApi } from '@/lib/api/posts'
import { Post, Category } from '@/types/post'
import { PostCard } from '@/components/post/PostCard'
import { Pagination } from '@/components/ui/Pagination'
import { Spinner } from '@/components/ui/Spinner'
import { Alert } from '@/components/ui/Alert'
import { Container } from '@/components/ui/Container'
import { Badge } from '@/components/ui/Badge'
import { Folder } from 'lucide-react'

export const dynamic = 'force-dynamic'

export default function CategoryPage({ params }: { params: { slug: string } }) {
  const searchParams = useSearchParams()
  const [posts, setPosts] = useState<Post[]>([])
  const [category, setCategory] = useState<Category | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [currentPage, setCurrentPage] = useState(1)
  const [totalPages, setTotalPages] = useState(1)
  const [total, setTotal] = useState(0)

  const limit = 20

  useEffect(() => {
    const page = parseInt(searchParams.get('page') || '1')
    setCurrentPage(page)
    loadCategoryPosts(page)
  }, [params.slug, searchParams])

  const loadCategoryPosts = async (page: number) => {
    try {
      setLoading(true)
      setError(null)

      const response = await postsApi.getPostsByCategory(params.slug, { page, limit })
      setPosts(response.data)
      setTotalPages(response.meta.totalPages)
      setTotal(response.meta.total)

      // Get category info from first post
      if (response.data.length > 0 && response.data[0].category) {
        setCategory(response.data[0].category)
      }
    } catch (err: any) {
      setError(err.message || 'Failed to load category posts')
    } finally {
      setLoading(false)
    }
  }

  const handlePageChange = (page: number) => {
    window.history.pushState({}, '', `/categories/${params.slug}?page=${page}`)
    setCurrentPage(page)
    loadCategoryPosts(page)
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
        <div className="flex items-center gap-3 mb-4">
          <Badge variant="secondary" className="flex items-center gap-2 text-lg px-4 py-2">
            <Folder size={20} />
            {category?.name || params.slug}
          </Badge>
        </div>
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-2">
          Posts in {category?.name || params.slug}
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          {total} post{total !== 1 ? 's' : ''} found
        </p>
      </div>

      {error && (
        <Alert variant="error" className="mb-6">
          {error}
        </Alert>
      )}

      {posts.length === 0 && !loading ? (
        <Alert variant="info">No posts found in this category</Alert>
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
