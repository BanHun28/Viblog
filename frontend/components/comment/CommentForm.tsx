'use client'

import { useState } from 'react'
import { useAuthStore } from '@/lib/store/authStore'
import { commentsApi } from '@/lib/api/comments'
import { Comment } from '@/types/comment'
import { Button } from '@/components/ui/Button'
import { Textarea } from '@/components/ui/Textarea'
import { Input } from '@/components/ui/Input'
import { Alert } from '@/components/ui/Alert'

interface CommentFormProps {
  postId: number | string
  parentId?: number | string
  onCommentAdded: (comment: Comment) => void
  onCancel?: () => void
}

export function CommentForm({ postId, parentId, onCommentAdded, onCancel }: CommentFormProps) {
  const { user } = useAuthStore()
  const [content, setContent] = useState('')
  const [authorName, setAuthorName] = useState('')
  const [authorEmail, setAuthorEmail] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!content.trim()) {
      setError('Comment content is required')
      return
    }

    if (!user && (!authorName.trim() || !authorEmail.trim())) {
      setError('Name and email are required for anonymous comments')
      return
    }

    try {
      setLoading(true)
      setError(null)

      const commentData = {
        content: content.trim(),
        ...(user ? {} : { authorName: authorName.trim(), authorEmail: authorEmail.trim() })
      }

      const newComment = parentId
        ? await commentsApi.createReply(parentId, commentData)
        : await commentsApi.createComment(postId, commentData)

      onCommentAdded(newComment)
      setContent('')
      setAuthorName('')
      setAuthorEmail('')
    } catch (err: any) {
      setError(err.message || 'Failed to post comment')
    } finally {
      setLoading(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      {error && <Alert variant="error">{error}</Alert>}

      {!user && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <Input
            type="text"
            placeholder="Your name"
            value={authorName}
            onChange={(e) => setAuthorName(e.target.value)}
            required
          />
          <Input
            type="email"
            placeholder="Your email"
            value={authorEmail}
            onChange={(e) => setAuthorEmail(e.target.value)}
            required
          />
        </div>
      )}

      <Textarea
        placeholder={user ? 'Write a comment...' : 'Write a comment... (Name and email required)'}
        value={content}
        onChange={(e) => setContent(e.target.value)}
        rows={4}
        required
      />

      <div className="flex gap-2">
        <Button type="submit" disabled={loading}>
          {loading ? 'Posting...' : parentId ? 'Post Reply' : 'Post Comment'}
        </Button>
        {onCancel && (
          <Button type="button" variant="outline" onClick={onCancel}>
            Cancel
          </Button>
        )}
      </div>
    </form>
  )
}
