export interface ApiResponse<T> {
  data: T
  message?: string
}

export interface ApiError {
  code: string
  message: string
  details?: Record<string, string[]>
}

export interface PaginationMeta {
  total: number
  page: number
  limit: number
  totalPages: number
  hasNext: boolean
  hasPrev: boolean
}

export interface CursorPaginationMeta {
  hasNext: boolean
  nextCursor?: string
}

export interface PaginatedResponse<T> {
  data: T[]
  meta: PaginationMeta
}

export interface CursorPaginatedResponse<T> {
  data: T[]
  meta: CursorPaginationMeta
}
