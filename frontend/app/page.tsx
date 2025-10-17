import { MainLayout } from '@/components/layout/MainLayout'

export default function Home() {
  return (
    <MainLayout>
      <div className="text-center py-12">
        <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-4">Welcome to Viblog</h1>
        <p className="text-lg text-gray-600 dark:text-gray-400 mb-8">Share your thoughts and experiences</p>
        <div className="flex justify-center gap-4">
          <a
            href="/posts"
            className="px-6 py-3 bg-blue-600 dark:bg-blue-500 text-white rounded-lg hover:bg-blue-700 dark:hover:bg-blue-600 transition-colors"
          >
            Browse Posts
          </a>
          <a
            href="/register"
            className="px-6 py-3 border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors"
          >
            Sign Up
          </a>
        </div>
      </div>
    </MainLayout>
  )
}
