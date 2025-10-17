import { User } from '@/types/user'
import { cn } from '@/lib/utils/cn'

interface UserAvatarProps {
  user: User
  size?: 'sm' | 'md' | 'lg' | 'xl'
  className?: string
  showBorder?: boolean
}

const sizeClasses = {
  sm: 'w-8 h-8 text-sm',
  md: 'w-12 h-12 text-base',
  lg: 'w-16 h-16 text-xl',
  xl: 'w-24 h-24 text-3xl',
}

export function UserAvatar({
  user,
  size = 'md',
  className,
  showBorder = false,
}: UserAvatarProps) {
  const sizeClass = sizeClasses[size]
  const borderClass = showBorder ? 'border-2 border-blue-500' : ''

  if (user.avatar_url) {
    return (
      <img
        src={user.avatar_url}
        alt={user.nickname}
        className={cn(
          'rounded-full object-cover',
          sizeClass,
          borderClass,
          className
        )}
      />
    )
  }

  return (
    <div
      className={cn(
        'rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white font-bold',
        sizeClass,
        borderClass,
        className
      )}
    >
      {user.nickname.charAt(0).toUpperCase()}
    </div>
  )
}
