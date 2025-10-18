'use client'

import { useState } from 'react'
import { Comment } from '@/types/comment'
import { commentsApi } from '@/lib/api/comments'
import { useAuthStore } from '@/lib/store/authStore'
import { CommentForm } from './CommentForm'
import { Button } from '@/components/ui/Button'
import { Heart, Reply, Trash2 } from 'lucide-react'

interface CommentItemProps {
  comment: Comment
  postId: number | string
  depth?: number
}

export function CommentItem({ comment, postId, depth = 0 }: CommentItemProps) {
  const { user } = useAuthStore()
  const [showReplyForm, setShowReplyForm] = useState(false)
  const [replies, setReplies] = useState<Comment[]>([])
  const [showReplies, setShowReplies] = useState(false)
  const [isLiked, setIsLiked] = useState(false)
  const [likeCount, setLikeCount] = useState(comment.likeCount || 0)

  const handleLike = async () => {
    if (!user) return

    try {
      if (isLiked) {
        await commentsApi.unlikeComment(comment.id)
        setLikeCount((prev) => prev - 1)
      } else {
        await commentsApi.likeComment(comment.id)
        setLikeCount((prev) => prev + 1)
      }
      setIsLiked(!isLiked)
    } catch (err) {
      console.error('Failed to toggle like:', err)
    }
  }

  const handleDelete = async () => {
    if (!confirm('Are you sure you want to delete this comment?')) return

    try {
      await commentsApi.deleteComment(comment.id)
      window.location.reload()
    } catch (err) {
      console.error('Failed to delete comment:', err)
    }
  }

  const loadReplies = async () => {
    try {
      const repliesData = await commentsApi.getReplies(comment.id)
      setReplies(repliesData)
      setShowReplies(true)
    } catch (err) {
      console.error('Failed to load replies:', err)
    }
  }

  const handleReplyAdded = (newReply: Comment) => {
    setReplies((prev) => [newReply, ...prev])
    setShowReplyForm(false)
    setShowReplies(true)
  }

  const canDelete = user && (user.id.toString() === comment.authorId || user.is_admin)

  return (
    <div className={`${depth > 0 ? 'ml-8 pl-4 border-l-2 border-gray-200 dark:border-gray-700' : ''}`}>
      <div className="bg-gray-50 dark:bg-gray-900 rounded-lg p-4">
        <div className="flex items-start justify-between mb-2">
          <div>
            <span className="font-medium text-gray-900 dark:text-white">
              {comment.author?.nickname || comment.anonymousNickname || 'Anonymous'}
            </span>
            <span className="text-sm text-gray-500 dark:text-gray-400 ml-2">
              {new Date(comment.createdAt).toLocaleDateString()}
            </span>
          </div>
          {canDelete && (
            <Button
              variant="ghost"
              size="sm"
              onClick={handleDelete}
              className="text-red-500 hover:text-red-700"
            >
              <Trash2 size={16} />
            </Button>
          )}
        </div>

        <p className="text-gray-700 dark:text-gray-300 mb-3">{comment.content}</p>

        <div className="flex items-center gap-4">
          <Button
            variant="ghost"
            size="sm"
            onClick={handleLike}
            className="flex items-center gap-1"
            disabled={!user}
          >
            <Heart size={16} fill={isLiked ? 'currentColor' : 'none'} />
            <span>{likeCount}</span>
          </Button>

          {depth < 3 && (
            <Button
              variant="ghost"
              size="sm"
              onClick={() => setShowReplyForm(!showReplyForm)}
              className="flex items-center gap-1"
            >
              <Reply size={16} />
              <span>Reply</span>
            </Button>
          )}

          {comment.replyCount > 0 && !showReplies && (
            <Button
              variant="ghost"
              size="sm"
              onClick={loadReplies}
            >
              Show {comment.replyCount} {comment.replyCount === 1 ? 'reply' : 'replies'}
            </Button>
          )}
        </div>
      </div>

      {showReplyForm && (
        <div className="mt-4">
          <CommentForm
            postId={postId}
            parentId={comment.id}
            onCommentAdded={handleReplyAdded}
            onCancel={() => setShowReplyForm(false)}
          />
        </div>
      )}

      {showReplies && replies.length > 0 && (
        <div className="mt-4 space-y-4">
          {replies.map((reply) => (
            <CommentItem
              key={reply.id}
              comment={reply}
              postId={postId}
              depth={depth + 1}
            />
          ))}
        </div>
      )}
    </div>
  )
}
