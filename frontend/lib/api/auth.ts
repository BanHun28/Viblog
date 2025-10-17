import { apiClient } from './client'
import {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
  UpdateProfileRequest,
  UpdateProfileResponse,
  RefreshTokenRequest,
  RefreshTokenResponse,
  User,
} from '@/types/user'

export const authApi = {
  /**
   * Login user with email and password
   */
  async login(data: LoginRequest): Promise<LoginResponse> {
    return apiClient.post<LoginResponse>('/auth/login', data)
  },

  /**
   * Register new user
   */
  async register(data: RegisterRequest): Promise<RegisterResponse> {
    return apiClient.post<RegisterResponse>('/auth/register', data)
  },

  /**
   * Logout current user
   */
  async logout(): Promise<{ message: string }> {
    return apiClient.post<{ message: string }>('/auth/logout')
  },

  /**
   * Get current user profile
   */
  async getProfile(): Promise<{ user: User }> {
    return apiClient.get<{ user: User }>('/auth/me')
  },

  /**
   * Update current user profile
   */
  async updateProfile(data: UpdateProfileRequest): Promise<UpdateProfileResponse> {
    return apiClient.put<UpdateProfileResponse>('/auth/me', data)
  },

  /**
   * Refresh access token
   */
  async refreshToken(data: RefreshTokenRequest): Promise<RefreshTokenResponse> {
    return apiClient.post<RefreshTokenResponse>('/auth/refresh', data)
  },
}
