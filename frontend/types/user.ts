export interface User {
  id: number
  email: string
  nickname: string
  avatar_url?: string
  bio?: string
  is_admin: boolean
  created_at: string
}

export interface AuthTokens {
  access_token: string
  refresh_token: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  user: User
  access_token: string
  refresh_token: string
}

export interface RegisterRequest {
  email: string
  password: string
  nickname: string
}

export interface RegisterResponse {
  message: string
  user: {
    id: number
    email: string
    nickname: string
    created_at: string
  }
}

export interface UpdateProfileRequest {
  nickname?: string
  avatar_url?: string
  bio?: string
}

export interface UpdateProfileResponse {
  message: string
  user: User
}

export interface RefreshTokenRequest {
  refresh_token: string
}

export interface RefreshTokenResponse {
  access_token: string
}
