export default function AdminEditPostPage({ params }: { params: { id: string } }) {
  return (
    <div>
      <h1>Edit Post {params.id}</h1>
    </div>
  )
}
