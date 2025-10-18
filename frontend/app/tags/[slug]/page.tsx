'use client'

import { useEffect, useState } from 'react'
import { useSearchParams } from 'next/navigation'
import { postsApi } from '@/lib/api/posts'
import { Post, Tag } from '@/types/post'
import { PostCard } from '@/components/post/PostCard'
import { Pagination } from '@/components/ui/Pagination'
import { Spinner } from '@/components/ui/Spinner'
import { Alert } from '@/components/ui/Alert'
import { Container } from '@/components/ui/Container'
import { Badge } from '@/components/ui/Badge'
import { Tag as TagIcon } from 'lucide-react'

export const dynamic = 'force-dynamic'

export default function TagPage({ params }: { params: { slug: string } }) {
  const searchParams = useSearchParams()
  const [posts, setPosts] = useState<Post[]>([])
  const [tag, setTag] = useState<Tag | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [currentPage, setCurrentPage] = useState(1)
  const [totalPages, setTotalPages] = useState(1)
  const [total, setTotal] = useState(0)

  const limit = 20

  useEffect(() => {
    const page = parseInt(searchParams.get('page') || '1')
    setCurrentPage(page)
    loadTagPosts(page)
  }, [params.slug, searchParams])

  const loadTagPosts = async (page: number) => {
    try {
      setLoading(true)
      setError(null)

      const response = await postsApi.getPostsByTag(params.slug, { page, limit })
      setPosts(response.data)
      setTotalPages(response.meta.totalPages)
      setTotal(response.meta.total)

      // Get tag info from first post
      if (response.data.length > 0 && response.data[0].tags) {
        const foundTag = response.data[0].tags.find(t => t.slug === params.slug)
        if (foundTag) {
          setTag(foundTag)
        }
      }
    } catch (err: any) {
      setError(err.message || 'Failed to load tag posts')
    } finally {
      setLoading(false)
    }
  }

  const handlePageChange = (page: number) => {
    window.history.pushState({}, '', `/tags/${params.slug}?page=${page}`)
    setCurrentPage(page)
    loadTagPosts(page)
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
          <Badge variant="outline" className="flex items-center gap-2 text-lg px-4 py-2">
            <TagIcon size={20} />
            {tag?.name || params.slug}
          </Badge>
        </div>
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-2">
          Posts tagged with {tag?.name || params.slug}
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
        <Alert variant="info">No posts found with this tag</Alert>
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
