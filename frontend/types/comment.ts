export interface Comment {
  id: string
  postId: string
  content: string
  // 회원 댓글
  authorId?: string // null for anonymous comments
  author?: {
    id: string
    nickname: string
    profileImage?: string
  }
  // 익명 댓글
  isAnonymous: boolean
  anonymousNickname?: string // 익명 댓글인 경우
  // 기타
  parentId?: string // for replies
  likeCount: number
  replyCount: number
  isLiked?: boolean // 현재 사용자가 좋아요 했는지
  replies?: Comment[]
  createdAt: string
  updatedAt: string
}

export interface CreateCommentRequest {
  postId: string
  content: string
  parentId?: string
  // For anonymous comments
  isAnonymous: boolean
  anonymousNickname?: string
  anonymousPassword?: string
}

export interface UpdateCommentRequest {
  content: string
  // For anonymous comments
  anonymousPassword?: string
}

export interface DeleteCommentRequest {
  // For anonymous comments
  anonymousPassword?: string
}

export interface CommentLike {
  id: string
  commentId: string
  userId: string
  createdAt: string
}
