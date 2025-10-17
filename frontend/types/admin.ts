// 대시보드 통계
export interface DashboardStats {
  totalPosts: number
  totalComments: number
  totalUsers: number
  totalViews: number
  totalLikes: number
  recentPosts: number // 최근 7일 게시글
  recentComments: number // 최근 7일 댓글
  recentUsers: number // 최근 7일 가입자
}

// 조회수/좋아요 추이 (차트 데이터)
export interface TrendData {
  date: string
  views: number
  likes: number
}

// 인기 글
export interface PopularPost {
  id: string
  title: string
  viewCount: number
  likeCount: number
  commentCount: number
  publishedAt: string
}

// 인기 태그
export interface PopularTag {
  id: string
  name: string
  postCount: number
}

// 최근 댓글/답글
export interface RecentComment {
  id: string
  content: string
  postId: string
  postTitle: string
  authorNickname: string
  isAnonymous: boolean
  createdAt: string
}

// 관리자 대시보드 응답
export interface AdminDashboardResponse {
  stats: DashboardStats
  trends: TrendData[] // 최근 30일
  popularPosts: PopularPost[] // 상위 10개
  popularTags: PopularTag[] // 상위 10개
  recentComments: RecentComment[] // 최근 20개
}
