export interface User {
  id: string
  email: string
  nickname: string
  profileImage?: string
  role: 'admin' | 'user'
  createdAt: string
  updatedAt: string
}

export interface AuthTokens {
  accessToken: string
  refreshToken: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  nickname: string
  profileImage?: string
}

export interface UpdateProfileRequest {
  nickname?: string
  profileImage?: string
}
