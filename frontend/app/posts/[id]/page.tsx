'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { postsApi } from '@/lib/api/posts'
import { commentsApi } from '@/lib/api/comments'
import { Post } from '@/types/post'
import { Comment } from '@/types/comment'
import { MarkdownViewer } from '@/components/post/MarkdownViewer'
import { CommentList } from '@/components/comment/CommentList'
import { CommentForm } from '@/components/comment/CommentForm'
import { Button } from '@/components/ui/Button'
import { Badge } from '@/components/ui/Badge'
import { Spinner } from '@/components/ui/Spinner'
import { Alert } from '@/components/ui/Alert'
import { useAuthStore } from '@/lib/store/authStore'
import { Heart, Bookmark, Eye, Calendar, User, Tag, Folder } from 'lucide-react'

export default function PostDetailPage({ params }: { params: { id: string } }) {
  const router = useRouter()
  const { user } = useAuthStore()
  const [post, setPost] = useState<Post | null>(null)
  const [comments, setComments] = useState<Comment[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [isLiked, setIsLiked] = useState(false)
  const [isBookmarked, setIsBookmarked] = useState(false)
  const [likeCount, setLikeCount] = useState(0)

  useEffect(() => {
    loadPostData()
  }, [params.id])

  const loadPostData = async () => {
    try {
      setLoading(true)
      setError(null)

      // Load post and increment view
      const [postData, commentsData] = await Promise.all([
        postsApi.getPostById(params.id),
        commentsApi.getCommentsByPostId(parseInt(params.id)),
      ])

      setPost(postData)
      setComments(commentsData)
      setIsLiked(postData.isLiked || false)
      setIsBookmarked(postData.isBookmarked || false)
      setLikeCount(postData.likeCount)

      // Increment view count
      await postsApi.incrementView(params.id)
    } catch (err: any) {
      setError(err.message || 'Failed to load post')
    } finally {
      setLoading(false)
    }
  }

  const handleLike = async () => {
    if (!user) {
      router.push('/login')
      return
    }

    try {
      if (isLiked) {
        await postsApi.unlikePost(params.id)
        setIsLiked(false)
        setLikeCount((prev) => prev - 1)
      } else {
        await postsApi.likePost(params.id)
        setIsLiked(true)
        setLikeCount((prev) => prev + 1)
      }
    } catch (err: any) {
      console.error('Failed to toggle like:', err)
    }
  }

  const handleBookmark = async () => {
    if (!user) {
      router.push('/login')
      return
    }

    try {
      if (isBookmarked) {
        await postsApi.unbookmarkPost(params.id)
        setIsBookmarked(false)
      } else {
        await postsApi.bookmarkPost(params.id)
        setIsBookmarked(true)
      }
    } catch (err: any) {
      console.error('Failed to toggle bookmark:', err)
    }
  }

  const handleCommentAdded = (newComment: Comment) => {
    setComments((prev) => [newComment, ...prev])
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <Spinner size="lg" />
      </div>
    )
  }

  if (error || !post) {
    return (
      <div className="container mx-auto px-4 py-8">
        <Alert variant="error">{error || 'Post not found'}</Alert>
      </div>
    )
  }

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      <article className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-8 mb-8">
        {/* Post Header */}
        <header className="mb-6">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-4">
            {post.title}
          </h1>

          {/* Meta Info */}
          <div className="flex flex-wrap items-center gap-4 text-sm text-gray-600 dark:text-gray-400 mb-4">
            <div className="flex items-center gap-2">
              <User size={16} />
              <span>{post.author?.nickname || 'Anonymous'}</span>
            </div>
            <div className="flex items-center gap-2">
              <Calendar size={16} />
              <span>{new Date(post.publishedAt || post.createdAt).toLocaleDateString()}</span>
            </div>
            <div className="flex items-center gap-2">
              <Eye size={16} />
              <span>{post.viewCount} views</span>
            </div>
          </div>

          {/* Category and Tags */}
          <div className="flex flex-wrap items-center gap-2 mb-4">
            {post.category && (
              <Badge
                variant="secondary"
                className="flex items-center gap-1 cursor-pointer"
                onClick={() => router.push(`/categories/${post.category?.slug}`)}
              >
                <Folder size={14} />
                {post.category.name}
              </Badge>
            )}
            {post.tags?.map((tag) => (
              <Badge
                key={tag.id}
                variant="outline"
                className="flex items-center gap-1 cursor-pointer"
                onClick={() => router.push(`/tags/${tag.slug}`)}
              >
                <Tag size={14} />
                {tag.name}
              </Badge>
            ))}
          </div>

          {/* Action Buttons */}
          <div className="flex items-center gap-4 pt-4 border-t border-gray-200 dark:border-gray-700">
            <Button
              variant={isLiked ? 'primary' : 'outline'}
              size="sm"
              onClick={handleLike}
              className="flex items-center gap-2"
            >
              <Heart size={18} fill={isLiked ? 'currentColor' : 'none'} />
              <span>{likeCount}</span>
            </Button>
            <Button
              variant={isBookmarked ? 'primary' : 'outline'}
              size="sm"
              onClick={handleBookmark}
              className="flex items-center gap-2"
            >
              <Bookmark size={18} fill={isBookmarked ? 'currentColor' : 'none'} />
              <span>Bookmark</span>
            </Button>
          </div>
        </header>

        {/* Post Content */}
        <div className="prose dark:prose-invert max-w-none mb-8">
          <MarkdownViewer content={post.content} />
        </div>
      </article>

      {/* Comments Section */}
      <section className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-8">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-6">
          Comments ({post.commentCount})
        </h2>

        {/* Comment Form */}
        <div className="mb-8">
          <CommentForm
            postId={parseInt(params.id)}
            onCommentAdded={handleCommentAdded}
          />
        </div>

        {/* Comments List */}
        <CommentList comments={comments} postId={parseInt(params.id)} />
      </section>
    </div>
  )
}
