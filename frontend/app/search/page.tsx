'use client'

import { useEffect, useState } from 'react'
import { useSearchParams } from 'next/navigation'
import { postsApi } from '@/lib/api/posts'
import { Post } from '@/types/post'
import { PostCard } from '@/components/post/PostCard'
import { Input } from '@/components/ui/Input'
import { Spinner } from '@/components/ui/Spinner'
import { Alert } from '@/components/ui/Alert'
import { Container } from '@/components/ui/Container'
import { Search as SearchIcon } from 'lucide-react'
import { useRouter } from 'next/navigation'

export default function SearchPage() {
  const router = useRouter()
  const searchParams = useSearchParams()
  const [query, setQuery] = useState(searchParams.get('q') || '')
  const [posts, setPosts] = useState<Post[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [hasSearched, setHasSearched] = useState(false)

  useEffect(() => {
    const q = searchParams.get('q')
    if (q) {
      setQuery(q)
      performSearch(q)
    }
  }, [searchParams])

  const performSearch = async (searchQuery: string) => {
    if (!searchQuery.trim()) {
      return
    }

    try {
      setLoading(true)
      setError(null)
      setHasSearched(true)

      const results = await postsApi.searchPosts(searchQuery)
      setPosts(results)
    } catch (err: any) {
      setError(err.message || 'Failed to search posts')
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (query.trim()) {
      router.push(`/search?q=${encodeURIComponent(query)}`)
    }
  }

  return (
    <Container className="py-8">
      <div className="mb-8">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-6">
          Search Posts
        </h1>

        {/* Search Form */}
        <form onSubmit={handleSearch} className="max-w-2xl">
          <div className="relative">
            <Input
              type="text"
              placeholder="Search for posts..."
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              className="pl-10 pr-4 py-3 text-lg"
            />
            <SearchIcon
              size={20}
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
            />
          </div>
        </form>
      </div>

      {/* Loading State */}
      {loading && (
        <div className="flex justify-center py-12">
          <Spinner size="lg" />
        </div>
      )}

      {/* Error State */}
      {error && (
        <Alert variant="error" className="mb-6">
          {error}
        </Alert>
      )}

      {/* Results */}
      {!loading && hasSearched && (
        <>
          <div className="mb-6">
            <p className="text-gray-600 dark:text-gray-400">
              {posts.length > 0
                ? `Found ${posts.length} post${posts.length !== 1 ? 's' : ''} for "${searchParams.get('q')}"`
                : `No posts found for "${searchParams.get('q')}"`}
            </p>
          </div>

          {posts.length > 0 ? (
            <div className="grid gap-6">
              {posts.map((post) => (
                <PostCard key={post.id} post={post} />
              ))}
            </div>
          ) : (
            <Alert variant="info">
              Try different keywords or browse all posts
            </Alert>
          )}
        </>
      )}

      {/* Initial State */}
      {!hasSearched && !loading && (
        <div className="text-center py-12">
          <SearchIcon size={48} className="mx-auto text-gray-300 dark:text-gray-600 mb-4" />
          <p className="text-gray-500 dark:text-gray-400">
            Enter keywords to search for posts
          </p>
        </div>
      )}
    </Container>
  )
}
