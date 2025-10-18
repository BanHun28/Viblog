'use client'

import { useRouter } from 'next/navigation'
import { Post } from '@/types/post'
import { Badge } from '@/components/ui/Badge'
import { Button } from '@/components/ui/Button'
import { useAuthStore } from '@/lib/store/authStore'
import { postsApi } from '@/lib/api/posts'
import { Heart, Bookmark, Eye, MessageSquare, Calendar, User, Tag, Folder } from 'lucide-react'
import { useState } from 'react'

interface PostCardProps {
  post: Post
  onLikeChange?: (postId: string, isLiked: boolean) => void
  onBookmarkChange?: (postId: string, isBookmarked: boolean) => void
}

export function PostCard({ post, onLikeChange, onBookmarkChange }: PostCardProps) {
  const router = useRouter()
  const { user } = useAuthStore()
  const [isLiked, setIsLiked] = useState(post.isLiked || false)
  const [isBookmarked, setIsBookmarked] = useState(post.isBookmarked || false)
  const [likeCount, setLikeCount] = useState(post.likeCount)

  const handleLike = async (e: React.MouseEvent) => {
    e.stopPropagation()

    if (!user) {
      router.push('/login')
      return
    }

    try {
      const newLikedState = !isLiked
      if (newLikedState) {
        await postsApi.likePost(post.id)
        setLikeCount((prev) => prev + 1)
      } else {
        await postsApi.unlikePost(post.id)
        setLikeCount((prev) => prev - 1)
      }
      setIsLiked(newLikedState)
      onLikeChange?.(post.id, newLikedState)
    } catch (err) {
      console.error('Failed to toggle like:', err)
    }
  }

  const handleBookmark = async (e: React.MouseEvent) => {
    e.stopPropagation()

    if (!user) {
      router.push('/login')
      return
    }

    try {
      const newBookmarkedState = !isBookmarked
      if (newBookmarkedState) {
        await postsApi.bookmarkPost(post.id)
      } else {
        await postsApi.unbookmarkPost(post.id)
      }
      setIsBookmarked(newBookmarkedState)
      onBookmarkChange?.(post.id, newBookmarkedState)
    } catch (err) {
      console.error('Failed to toggle bookmark:', err)
    }
  }

  const handleCardClick = () => {
    router.push(`/posts/${post.id}`)
  }

  const handleCategoryClick = (e: React.MouseEvent) => {
    e.stopPropagation()
    if (post.category) {
      router.push(`/categories/${post.category.slug}`)
    }
  }

  const handleTagClick = (e: React.MouseEvent, slug: string) => {
    e.stopPropagation()
    router.push(`/tags/${slug}`)
  }

  return (
    <article
      onClick={handleCardClick}
      className="bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-lg transition-shadow cursor-pointer p-6"
    >
      {/* Header with category and date */}
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center gap-2">
          {post.category && (
            <Badge
              variant="secondary"
              className="flex items-center gap-1"
              onClick={handleCategoryClick}
            >
              <Folder size={14} />
              {post.category.name}
            </Badge>
          )}
        </div>
        <div className="flex items-center gap-2 text-sm text-gray-500 dark:text-gray-400">
          <Calendar size={14} />
          <span>{new Date(post.publishedAt || post.createdAt).toLocaleDateString()}</span>
        </div>
      </div>

      {/* Title and excerpt */}
      <div className="mb-4">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-2 hover:text-blue-600 dark:hover:text-blue-400 transition-colors">
          {post.title}
        </h2>
        {post.excerpt && (
          <p className="text-gray-600 dark:text-gray-300 line-clamp-2">
            {post.excerpt}
          </p>
        )}
      </div>

      {/* Tags */}
      {post.tags && post.tags.length > 0 && (
        <div className="flex flex-wrap gap-2 mb-4">
          {post.tags.map((tag) => (
            <Badge
              key={tag.id}
              variant="outline"
              className="flex items-center gap-1"
              onClick={(e) => handleTagClick(e, tag.slug)}
            >
              <Tag size={12} />
              {tag.name}
            </Badge>
          ))}
        </div>
      )}

      {/* Author and stats */}
      <div className="flex items-center justify-between pt-4 border-t border-gray-200 dark:border-gray-700">
        <div className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
          <User size={14} />
          <span>{post.author?.nickname || 'Anonymous'}</span>
        </div>

        <div className="flex items-center gap-4">
          {/* View count */}
          <div className="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400">
            <Eye size={16} />
            <span>{post.viewCount}</span>
          </div>

          {/* Comment count */}
          <div className="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400">
            <MessageSquare size={16} />
            <span>{post.commentCount}</span>
          </div>

          {/* Like button */}
          <Button
            variant={isLiked ? 'primary' : 'ghost'}
            size="sm"
            onClick={handleLike}
            className="flex items-center gap-1"
          >
            <Heart size={16} fill={isLiked ? 'currentColor' : 'none'} />
            <span>{likeCount}</span>
          </Button>

          {/* Bookmark button */}
          <Button
            variant={isBookmarked ? 'primary' : 'ghost'}
            size="sm"
            onClick={handleBookmark}
          >
            <Bookmark size={16} fill={isBookmarked ? 'currentColor' : 'none'} />
          </Button>
        </div>
      </div>
    </article>
  )
}
