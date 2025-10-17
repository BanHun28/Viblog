/**
 * SEO utilities for Viblog frontend
 * Meta tags and SEO optimization helpers
 */

export interface SeoMetadata {
  title: string;
  description: string;
  keywords?: string[];
  image?: string;
  url?: string;
  type?: 'website' | 'article' | 'profile';
  author?: string;
  publishedTime?: string;
  modifiedTime?: string;
}

/**
 * Default SEO metadata
 */
export const DEFAULT_SEO: SeoMetadata = {
  title: 'Viblog',
  description: '개인 블로그 플랫폼',
  type: 'website'
};

/**
 * Build page title
 */
export function buildPageTitle(title?: string, siteName: string = 'Viblog'): string {
  if (!title) {
    return siteName;
  }

  return `${title} | ${siteName}`;
}

/**
 * Generate meta tags for Next.js metadata
 */
export function generateMetadata(seo: Partial<SeoMetadata>): Record<string, any> {
  const metadata: Record<string, any> = {
    title: buildPageTitle(seo.title),
    description: seo.description || DEFAULT_SEO.description
  };

  // Open Graph
  metadata.openGraph = {
    title: seo.title || DEFAULT_SEO.title,
    description: seo.description || DEFAULT_SEO.description,
    type: seo.type || DEFAULT_SEO.type,
    url: seo.url,
    images: seo.image ? [{ url: seo.image }] : undefined
  };

  // Twitter Card
  metadata.twitter = {
    card: 'summary_large_image',
    title: seo.title || DEFAULT_SEO.title,
    description: seo.description || DEFAULT_SEO.description,
    images: seo.image ? [seo.image] : undefined
  };

  // Article metadata
  if (seo.type === 'article') {
    metadata.openGraph.article = {
      author: seo.author,
      publishedTime: seo.publishedTime,
      modifiedTime: seo.modifiedTime
    };
  }

  // Keywords
  if (seo.keywords && seo.keywords.length > 0) {
    metadata.keywords = seo.keywords.join(', ');
  }

  return metadata;
}

/**
 * Generate structured data (JSON-LD) for article
 */
export function generateArticleJsonLd(article: {
  title: string;
  description: string;
  author: string;
  publishedTime: string;
  modifiedTime?: string;
  image?: string;
  url: string;
}): string {
  const jsonLd = {
    '@context': 'https://schema.org',
    '@type': 'BlogPosting',
    headline: article.title,
    description: article.description,
    author: {
      '@type': 'Person',
      name: article.author
    },
    datePublished: article.publishedTime,
    dateModified: article.modifiedTime || article.publishedTime,
    image: article.image,
    url: article.url
  };

  return JSON.stringify(jsonLd);
}

/**
 * Generate structured data (JSON-LD) for breadcrumbs
 */
export function generateBreadcrumbJsonLd(breadcrumbs: Array<{ name: string; url: string }>): string {
  const jsonLd = {
    '@context': 'https://schema.org',
    '@type': 'BreadcrumbList',
    itemListElement: breadcrumbs.map((item, index) => ({
      '@type': 'ListItem',
      position: index + 1,
      name: item.name,
      item: item.url
    }))
  };

  return JSON.stringify(jsonLd);
}

/**
 * Sanitize text for SEO
 */
export function sanitizeForSeo(text: string): string {
  return text
    .replace(/<[^>]*>/g, '') // Remove HTML tags
    .replace(/\s+/g, ' ') // Normalize whitespace
    .trim()
    .slice(0, 160); // Limit length for meta description
}

/**
 * Generate canonical URL
 */
export function generateCanonicalUrl(path: string, baseUrl?: string): string {
  const base = baseUrl || process.env.NEXT_PUBLIC_SITE_URL || 'http://localhost:30001';
  const cleanPath = path.startsWith('/') ? path : `/${path}`;
  return `${base}${cleanPath}`;
}
