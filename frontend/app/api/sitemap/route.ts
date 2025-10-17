export async function GET() {
  return new Response('Sitemap', {
    headers: {
      'Content-Type': 'application/xml',
    },
  })
}
