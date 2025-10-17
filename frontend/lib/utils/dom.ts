/**
 * DOM utilities for Viblog frontend
 * DOM manipulation and event handling helpers
 */

/**
 * Check if element is in viewport
 */
export function isInViewport(element: HTMLElement): boolean {
  const rect = element.getBoundingClientRect();
  return (
    rect.top >= 0 &&
    rect.left >= 0 &&
    rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) &&
    rect.right <= (window.innerWidth || document.documentElement.clientWidth)
  );
}

/**
 * Scroll to element smoothly
 */
export function scrollToElement(
  element: HTMLElement | string,
  options?: ScrollIntoViewOptions
): void {
  const el = typeof element === 'string'
    ? document.querySelector(element) as HTMLElement
    : element;

  if (el) {
    el.scrollIntoView({
      behavior: 'smooth',
      block: 'start',
      ...options
    });
  }
}

/**
 * Scroll to top of page
 */
export function scrollToTop(smooth: boolean = true): void {
  window.scrollTo({
    top: 0,
    behavior: smooth ? 'smooth' : 'auto'
  });
}

/**
 * Get scroll position
 */
export function getScrollPosition(): { x: number; y: number } {
  return {
    x: window.pageXOffset || document.documentElement.scrollLeft,
    y: window.pageYOffset || document.documentElement.scrollTop
  };
}

/**
 * Check if scrolled to bottom
 */
export function isScrolledToBottom(threshold: number = 0): boolean {
  const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
  const scrollHeight = document.documentElement.scrollHeight;
  const clientHeight = window.innerHeight || document.documentElement.clientHeight;

  return scrollTop + clientHeight >= scrollHeight - threshold;
}

/**
 * Get element offset from top
 */
export function getElementOffset(element: HTMLElement): { top: number; left: number } {
  const rect = element.getBoundingClientRect();
  const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
  const scrollLeft = window.pageXOffset || document.documentElement.scrollLeft;

  return {
    top: rect.top + scrollTop,
    left: rect.left + scrollLeft
  };
}

/**
 * Add class to element
 */
export function addClass(element: HTMLElement, className: string): void {
  element.classList.add(className);
}

/**
 * Remove class from element
 */
export function removeClass(element: HTMLElement, className: string): void {
  element.classList.remove(className);
}

/**
 * Toggle class on element
 */
export function toggleClass(element: HTMLElement, className: string): void {
  element.classList.toggle(className);
}

/**
 * Check if element has class
 */
export function hasClass(element: HTMLElement, className: string): boolean {
  return element.classList.contains(className);
}

/**
 * Get element by id
 */
export function getById(id: string): HTMLElement | null {
  return document.getElementById(id);
}

/**
 * Query selector
 */
export function query<T extends HTMLElement = HTMLElement>(selector: string): T | null {
  return document.querySelector<T>(selector);
}

/**
 * Query selector all
 */
export function queryAll<T extends HTMLElement = HTMLElement>(selector: string): T[] {
  return Array.from(document.querySelectorAll<T>(selector));
}

/**
 * Add event listener with cleanup
 */
export function addEventListener<K extends keyof WindowEventMap>(
  element: Window | Document | HTMLElement,
  event: K,
  handler: (event: any) => void,
  options?: AddEventListenerOptions
): () => void {
  element.addEventListener(event as string, handler, options);

  return () => {
    element.removeEventListener(event as string, handler, options);
  };
}

/**
 * Lock body scroll (for modals)
 */
export function lockBodyScroll(): void {
  document.body.style.overflow = 'hidden';
}

/**
 * Unlock body scroll
 */
export function unlockBodyScroll(): void {
  document.body.style.overflow = '';
}

/**
 * Get window dimensions
 */
export function getWindowDimensions(): { width: number; height: number } {
  return {
    width: window.innerWidth || document.documentElement.clientWidth,
    height: window.innerHeight || document.documentElement.clientHeight
  };
}

/**
 * Check if mobile device
 */
export function isMobileDevice(): boolean {
  return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
    navigator.userAgent
  );
}

/**
 * Check if touch device
 */
export function isTouchDevice(): boolean {
  return 'ontouchstart' in window || navigator.maxTouchPoints > 0;
}

/**
 * Get element dimensions
 */
export function getElementDimensions(element: HTMLElement): {
  width: number;
  height: number;
} {
  const rect = element.getBoundingClientRect();
  return {
    width: rect.width,
    height: rect.height
  };
}

/**
 * Focus element
 */
export function focusElement(element: HTMLElement | string): void {
  const el = typeof element === 'string'
    ? document.querySelector(element) as HTMLElement
    : element;

  if (el) {
    el.focus();
  }
}

/**
 * Blur element
 */
export function blurElement(element: HTMLElement | string): void {
  const el = typeof element === 'string'
    ? document.querySelector(element) as HTMLElement
    : element;

  if (el) {
    el.blur();
  }
}

/**
 * Create element with attributes
 */
export function createElement<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  attributes?: Record<string, string>,
  children?: (HTMLElement | string)[]
): HTMLElementTagNameMap[K] {
  const element = document.createElement(tag);

  if (attributes) {
    Object.entries(attributes).forEach(([key, value]) => {
      element.setAttribute(key, value);
    });
  }

  if (children) {
    children.forEach(child => {
      if (typeof child === 'string') {
        element.appendChild(document.createTextNode(child));
      } else {
        element.appendChild(child);
      }
    });
  }

  return element;
}

/**
 * Remove element from DOM
 */
export function removeElement(element: HTMLElement | string): void {
  const el = typeof element === 'string'
    ? document.querySelector(element) as HTMLElement
    : element;

  if (el && el.parentNode) {
    el.parentNode.removeChild(el);
  }
}

/**
 * Insert HTML at position
 */
export function insertHTML(
  element: HTMLElement,
  position: InsertPosition,
  html: string
): void {
  element.insertAdjacentHTML(position, html);
}

/**
 * Wait for element to exist
 */
export function waitForElement(
  selector: string,
  timeout: number = 5000
): Promise<HTMLElement> {
  return new Promise((resolve, reject) => {
    const element = document.querySelector(selector) as HTMLElement;
    if (element) {
      resolve(element);
      return;
    }

    const observer = new MutationObserver(() => {
      const element = document.querySelector(selector) as HTMLElement;
      if (element) {
        observer.disconnect();
        resolve(element);
      }
    });

    observer.observe(document.body, {
      childList: true,
      subtree: true
    });

    setTimeout(() => {
      observer.disconnect();
      reject(new Error(`Element ${selector} not found within ${timeout}ms`));
    }, timeout);
  });
}

/**
 * Observe element visibility
 */
export function observeVisibility(
  element: HTMLElement,
  callback: (isVisible: boolean) => void,
  options?: IntersectionObserverInit
): () => void {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach(entry => {
        callback(entry.isIntersecting);
      });
    },
    options
  );

  observer.observe(element);

  return () => {
    observer.disconnect();
  };
}
