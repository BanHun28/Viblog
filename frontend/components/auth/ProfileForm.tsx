'use client'

import { useState, useEffect } from 'react'
import { useAuth } from '@/lib/hooks/useAuth'
import { Button } from '@/components/ui/Button'
import { Input } from '@/components/ui/Input'
import { Textarea } from '@/components/ui/Textarea'
import { Alert } from '@/components/ui/Alert'
import { Spinner } from '@/components/ui/Spinner'

export function ProfileForm() {
  const { user, updateProfile, isLoading } = useAuth()
  const [formData, setFormData] = useState({
    nickname: '',
    avatar_url: '',
    bio: '',
  })
  const [errors, setErrors] = useState<{
    nickname?: string
    avatar_url?: string
    bio?: string
  }>({})
  const [apiError, setApiError] = useState<string>('')
  const [success, setSuccess] = useState(false)

  useEffect(() => {
    if (user) {
      setFormData({
        nickname: user.nickname,
        avatar_url: user.avatar_url || '',
        bio: user.bio || '',
      })
    }
  }, [user])

  const validateForm = () => {
    const newErrors: typeof errors = {}

    // Nickname validation
    if (!formData.nickname) {
      newErrors.nickname = 'Nickname is required'
    } else if (formData.nickname.length < 2) {
      newErrors.nickname = 'Nickname must be at least 2 characters'
    } else if (formData.nickname.length > 30) {
      newErrors.nickname = 'Nickname must be less than 30 characters'
    }

    // Avatar URL validation (optional)
    if (formData.avatar_url && !/^https?:\/\/.+/.test(formData.avatar_url)) {
      newErrors.avatar_url = 'Must be a valid URL'
    }

    // Bio validation (optional)
    if (formData.bio && formData.bio.length > 500) {
      newErrors.bio = 'Bio must be less than 500 characters'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setApiError('')
    setSuccess(false)

    if (!validateForm()) return

    // Only send non-empty fields
    const updates: any = { nickname: formData.nickname }
    if (formData.avatar_url) updates.avatar_url = formData.avatar_url
    if (formData.bio) updates.bio = formData.bio

    const result = await updateProfile(updates)

    if (result.success) {
      setSuccess(true)
      setTimeout(() => setSuccess(false), 3000)
    } else {
      setApiError(result.error || 'Profile update failed')
    }
  }

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target
    setFormData((prev) => ({ ...prev, [name]: value }))
    // Clear error for this field
    if (errors[name as keyof typeof errors]) {
      setErrors((prev) => ({ ...prev, [name]: undefined }))
    }
  }

  if (!user) {
    return (
      <Alert variant="info">
        Please log in to view your profile.
      </Alert>
    )
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      {success && (
        <Alert variant="success" className="mb-4">
          Profile updated successfully!
        </Alert>
      )}

      {apiError && (
        <Alert variant="error" className="mb-4">
          {apiError}
        </Alert>
      )}

      <div>
        <label htmlFor="email" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Email
        </label>
        <Input
          id="email"
          name="email"
          type="email"
          value={user.email}
          disabled
          className="bg-gray-100 dark:bg-gray-800 cursor-not-allowed"
        />
        <p className="mt-1 text-xs text-gray-500 dark:text-gray-400">
          Email cannot be changed
        </p>
      </div>

      <div>
        <label htmlFor="nickname" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Nickname
        </label>
        <Input
          id="nickname"
          name="nickname"
          type="text"
          value={formData.nickname}
          onChange={handleChange}
          error={errors.nickname}
          disabled={isLoading}
          placeholder="Your display name"
        />
      </div>

      <div>
        <label htmlFor="avatar_url" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Avatar URL (Optional)
        </label>
        <Input
          id="avatar_url"
          name="avatar_url"
          type="url"
          value={formData.avatar_url}
          onChange={handleChange}
          error={errors.avatar_url}
          disabled={isLoading}
          placeholder="https://example.com/avatar.jpg"
        />
      </div>

      <div>
        <label htmlFor="bio" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Bio (Optional)
        </label>
        <Textarea
          id="bio"
          name="bio"
          value={formData.bio}
          onChange={handleChange}
          error={errors.bio}
          disabled={isLoading}
          placeholder="Tell us about yourself..."
          rows={4}
        />
        <p className="mt-1 text-xs text-gray-500 dark:text-gray-400">
          {formData.bio.length}/500 characters
        </p>
      </div>

      <div className="flex gap-4">
        <Button type="submit" disabled={isLoading}>
          {isLoading ? (
            <>
              <Spinner size="sm" className="mr-2" />
              Updating...
            </>
          ) : (
            'Update Profile'
          )}
        </Button>
      </div>
    </form>
  )
}
