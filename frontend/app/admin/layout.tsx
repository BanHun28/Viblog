export default function AdminLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div>
      <nav>Admin Navigation</nav>
      <main>{children}</main>
    </div>
  )
}
