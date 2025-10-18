'use client'

import { Comment } from '@/types/comment'
import { CommentItem } from './CommentItem'

interface CommentListProps {
  comments: Comment[]
  postId: number | string
}

export function CommentList({ comments, postId }: CommentListProps) {
  if (comments.length === 0) {
    return (
      <div className="text-center py-8 text-gray-500 dark:text-gray-400">
        No comments yet. Be the first to comment!
      </div>
    )
  }

  return (
    <div className="space-y-4">
      {comments.map((comment) => (
        <CommentItem key={comment.id} comment={comment} postId={postId} />
      ))}
    </div>
  )
}
