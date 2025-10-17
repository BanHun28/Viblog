export type PostStatus = 'draft' | 'published' | 'scheduled'

export interface Post {
  id: string
  title: string
  content: string
  excerpt?: string
  status: PostStatus
  authorId: string
  author?: {
    id: string
    nickname: string
    profileImage?: string
  }
  categoryId?: string
  category?: Category
  tags?: Tag[]
  viewCount: number
  likeCount: number
  commentCount: number
  isLiked?: boolean // 현재 사용자가 좋아요 했는지
  isBookmarked?: boolean // 현재 사용자가 북마크 했는지
  // SEO 관련
  metaTitle?: string
  metaDescription?: string
  metaKeywords?: string
  // 검색 하이라이팅
  highlightedTitle?: string
  highlightedContent?: string
  publishedAt?: string
  scheduledAt?: string
  createdAt: string
  updatedAt: string
}

export interface Category {
  id: string
  name: string
  slug: string
  parentId?: string
  children?: Category[]
  postCount: number
  createdAt: string
  updatedAt: string
}

export interface Tag {
  id: string
  name: string
  slug: string
  postCount: number
  createdAt: string
  updatedAt: string
}

export interface CreatePostRequest {
  title: string
  content: string
  excerpt?: string
  status: PostStatus
  categoryId?: string
  tagIds?: string[]
  scheduledAt?: string
}

export interface UpdatePostRequest {
  title?: string
  content?: string
  excerpt?: string
  status?: PostStatus
  categoryId?: string
  tagIds?: string[]
  scheduledAt?: string
}

export interface PostListParams {
  page?: number
  limit?: number
  categoryId?: string
  tagId?: string
  status?: PostStatus
  search?: string
  cursor?: string
}

export interface Like {
  id: string
  postId: string
  userId: string
  createdAt: string
}

export interface Bookmark {
  id: string
  postId: string
  userId: string
  createdAt: string
}
