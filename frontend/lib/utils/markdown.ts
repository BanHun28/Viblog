/**
 * Markdown utilities for Viblog frontend
 * Basic markdown parsing and manipulation (without external dependencies)
 */

/**
 * Extract headings from markdown
 */
export function extractHeadings(markdown: string): Array<{
  level: number;
  text: string;
  id: string;
}> {
  const headingRegex = /^(#{1,6})\s+(.+)$/gm;
  const headings: Array<{ level: number; text: string; id: string }> = [];
  let match;

  while ((match = headingRegex.exec(markdown)) !== null) {
    const level = match[1].length;
    const text = match[2].trim();
    const id = generateHeadingId(text);

    headings.push({ level, text, id });
  }

  return headings;
}

/**
 * Generate heading ID from text
 */
export function generateHeadingId(text: string): string {
  return text
    .toLowerCase()
    .replace(/[^\w\s가-힣-]/g, '')
    .replace(/\s+/g, '-')
    .replace(/-+/g, '-')
    .replace(/^-|-$/g, '');
}

/**
 * Extract code blocks from markdown
 */
export function extractCodeBlocks(markdown: string): Array<{
  language: string;
  code: string;
}> {
  const codeBlockRegex = /```(\w*)\n([\s\S]*?)```/g;
  const codeBlocks: Array<{ language: string; code: string }> = [];
  let match;

  while ((match = codeBlockRegex.exec(markdown)) !== null) {
    const language = match[1] || 'plaintext';
    const code = match[2].trim();

    codeBlocks.push({ language, code });
  }

  return codeBlocks;
}

/**
 * Extract images from markdown
 */
export function extractImages(markdown: string): Array<{
  alt: string;
  url: string;
}> {
  const imageRegex = /!\[([^\]]*)\]\(([^)]+)\)/g;
  const images: Array<{ alt: string; url: string }> = [];
  let match;

  while ((match = imageRegex.exec(markdown)) !== null) {
    const alt = match[1];
    const url = match[2];

    images.push({ alt, url });
  }

  return images;
}

/**
 * Extract links from markdown
 */
export function extractLinks(markdown: string): Array<{
  text: string;
  url: string;
}> {
  const linkRegex = /\[([^\]]+)\]\(([^)]+)\)/g;
  const links: Array<{ text: string; url: string }> = [];
  let match;

  while ((match = linkRegex.exec(markdown)) !== null) {
    const text = match[1];
    const url = match[2];

    links.push({ text, url });
  }

  return links;
}

/**
 * Strip markdown formatting
 */
export function stripMarkdown(markdown: string): string {
  return markdown
    // Remove code blocks
    .replace(/```[\s\S]*?```/g, '')
    // Remove inline code
    .replace(/`([^`]+)`/g, '$1')
    // Remove images
    .replace(/!\[([^\]]*)\]\(([^)]+)\)/g, '')
    // Remove links (keep text)
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '$1')
    // Remove headings
    .replace(/^#{1,6}\s+/gm, '')
    // Remove bold
    .replace(/\*\*([^*]+)\*\*/g, '$1')
    .replace(/__([^_]+)__/g, '$1')
    // Remove italic
    .replace(/\*([^*]+)\*/g, '$1')
    .replace(/_([^_]+)_/g, '$1')
    // Remove strikethrough
    .replace(/~~([^~]+)~~/g, '$1')
    // Remove blockquotes
    .replace(/^>\s+/gm, '')
    // Remove lists
    .replace(/^[-*+]\s+/gm, '')
    .replace(/^\d+\.\s+/gm, '')
    // Remove horizontal rules
    .replace(/^[-*_]{3,}$/gm, '')
    // Normalize whitespace
    .replace(/\n{3,}/g, '\n\n')
    .trim();
}

/**
 * Count words in markdown (excluding code)
 */
export function countWords(markdown: string): number {
  const plainText = stripMarkdown(markdown);
  const words = plainText.match(/[\w가-힣]+/g);
  return words ? words.length : 0;
}

/**
 * Estimate reading time (200 words per minute)
 */
export function estimateReadingTime(markdown: string): number {
  const wordCount = countWords(markdown);
  const minutes = Math.ceil(wordCount / 200);
  return Math.max(1, minutes);
}

/**
 * Generate table of contents from markdown
 */
export function generateTableOfContents(markdown: string): string {
  const headings = extractHeadings(markdown);

  if (headings.length === 0) {
    return '';
  }

  let toc = '## 목차\n\n';

  headings.forEach(heading => {
    const indent = '  '.repeat(heading.level - 1);
    toc += `${indent}- [${heading.text}](#${heading.id})\n`;
  });

  return toc;
}

/**
 * Add IDs to headings in markdown
 */
export function addHeadingIds(markdown: string): string {
  return markdown.replace(/^(#{1,6})\s+(.+)$/gm, (match, hashes, text) => {
    const id = generateHeadingId(text);
    return `${hashes} ${text} {#${id}}`;
  });
}

/**
 * Replace image URLs in markdown
 */
export function replaceImageUrls(
  markdown: string,
  replacer: (url: string) => string
): string {
  return markdown.replace(
    /!\[([^\]]*)\]\(([^)]+)\)/g,
    (match, alt, url) => {
      const newUrl = replacer(url);
      return `![${alt}](${newUrl})`;
    }
  );
}

/**
 * Escape markdown special characters
 */
export function escapeMarkdown(text: string): string {
  const specialChars = /([\\`*_{}[\]()#+\-.!])/g;
  return text.replace(specialChars, '\\$1');
}

/**
 * Unescape markdown special characters
 */
export function unescapeMarkdown(text: string): string {
  return text.replace(/\\([\\`*_{}[\]()#+\-.!])/g, '$1');
}

/**
 * Validate markdown structure
 */
export function validateMarkdown(markdown: string): {
  isValid: boolean;
  errors: string[];
} {
  const errors: string[] = [];

  // Check for unclosed code blocks
  const codeBlockCount = (markdown.match(/```/g) || []).length;
  if (codeBlockCount % 2 !== 0) {
    errors.push('코드 블록이 올바르게 닫히지 않았습니다');
  }

  // Check for unmatched brackets
  const openBrackets = (markdown.match(/\[/g) || []).length;
  const closeBrackets = (markdown.match(/\]/g) || []).length;
  if (openBrackets !== closeBrackets) {
    errors.push('대괄호가 올바르게 닫히지 않았습니다');
  }

  // Check for unmatched parentheses in links
  const openParens = (markdown.match(/\(/g) || []).length;
  const closeParens = (markdown.match(/\)/g) || []).length;
  if (openParens !== closeParens) {
    errors.push('괄호가 올바르게 닫히지 않았습니다');
  }

  return {
    isValid: errors.length === 0,
    errors
  };
}

/**
 * Get first image URL from markdown
 */
export function getFirstImage(markdown: string): string | null {
  const images = extractImages(markdown);
  return images.length > 0 ? images[0].url : null;
}

/**
 * Get first paragraph from markdown
 */
export function getFirstParagraph(markdown: string): string {
  const plainText = stripMarkdown(markdown);
  const paragraphs = plainText.split('\n\n').filter(p => p.trim().length > 0);
  return paragraphs.length > 0 ? paragraphs[0] : '';
}

/**
 * Truncate markdown to max length
 */
export function truncateMarkdown(markdown: string, maxLength: number): string {
  const plainText = stripMarkdown(markdown);

  if (plainText.length <= maxLength) {
    return plainText;
  }

  return plainText.slice(0, maxLength).trim() + '...';
}
