import { LoginForm } from '@/components/auth/LoginForm'
import { Card } from '@/components/ui/Card'
import { Container } from '@/components/ui/Container'

export default function LoginPage() {
  return (
    <Container className="py-16">
      <div className="max-w-md mx-auto">
        <Card className="p-8">
          <div className="text-center mb-8">
            <h1 className="text-3xl font-bold text-gray-900 dark:text-white mb-2">
              Welcome back
            </h1>
            <p className="text-gray-600 dark:text-gray-400">
              Sign in to your account to continue
            </p>
          </div>
          <LoginForm />
        </Card>
      </div>
    </Container>
  )
}
