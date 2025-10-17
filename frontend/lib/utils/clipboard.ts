/**
 * Clipboard utilities for Viblog frontend
 * Copy text, code, URLs to clipboard
 */

/**
 * Copy text to clipboard
 */
export async function copyToClipboard(text: string): Promise<boolean> {
  // Modern Clipboard API
  if (navigator.clipboard && window.isSecureContext) {
    try {
      await navigator.clipboard.writeText(text);
      return true;
    } catch (error) {
      console.error('Failed to copy to clipboard:', error);
      return false;
    }
  }

  // Fallback for older browsers
  return copyToClipboardFallback(text);
}

/**
 * Fallback method using execCommand
 */
function copyToClipboardFallback(text: string): boolean {
  const textArea = document.createElement('textarea');
  textArea.value = text;
  textArea.style.position = 'fixed';
  textArea.style.left = '-999999px';
  textArea.style.top = '-999999px';
  document.body.appendChild(textArea);

  try {
    textArea.focus();
    textArea.select();
    const successful = document.execCommand('copy');
    document.body.removeChild(textArea);
    return successful;
  } catch (error) {
    console.error('Fallback: Failed to copy to clipboard:', error);
    document.body.removeChild(textArea);
    return false;
  }
}

/**
 * Copy code block to clipboard
 * Preserves formatting and indentation
 */
export async function copyCodeToClipboard(code: string): Promise<boolean> {
  return copyToClipboard(code);
}

/**
 * Copy current URL to clipboard
 */
export async function copyCurrentUrlToClipboard(): Promise<boolean> {
  return copyToClipboard(window.location.href);
}

/**
 * Copy post link to clipboard (for sharing)
 */
export async function copyPostLink(postId: string): Promise<boolean> {
  const baseUrl = window.location.origin;
  const postUrl = `${baseUrl}/posts/${postId}`;
  return copyToClipboard(postUrl);
}

/**
 * Read text from clipboard
 */
export async function readFromClipboard(): Promise<string | null> {
  if (navigator.clipboard && window.isSecureContext) {
    try {
      return await navigator.clipboard.readText();
    } catch (error) {
      console.error('Failed to read from clipboard:', error);
      return null;
    }
  }

  console.warn('Clipboard API not available');
  return null;
}

/**
 * Check if clipboard API is supported
 */
export function isClipboardSupported(): boolean {
  return !!(navigator.clipboard && window.isSecureContext);
}

/**
 * Copy with success/error feedback
 */
export async function copyWithFeedback(
  text: string,
  onSuccess?: () => void,
  onError?: () => void
): Promise<boolean> {
  const success = await copyToClipboard(text);

  if (success && onSuccess) {
    onSuccess();
  } else if (!success && onError) {
    onError();
  }

  return success;
}
