/**
 * File and image upload utilities for Viblog frontend
 * File validation, image processing, upload helpers
 */

import { IMAGE_CONFIG } from '../constants';

export interface FileValidationResult {
  isValid: boolean;
  error?: string;
}

/**
 * Validate file size
 */
export function validateFileSize(file: File, maxSize: number = IMAGE_CONFIG.MAX_SIZE): FileValidationResult {
  if (file.size > maxSize) {
    const maxSizeMB = (maxSize / (1024 * 1024)).toFixed(1);
    return {
      isValid: false,
      error: `파일 크기는 ${maxSizeMB}MB를 초과할 수 없습니다`
    };
  }

  return { isValid: true };
}

/**
 * Validate file type
 */
export function validateFileType(
  file: File,
  allowedTypes: readonly string[] = IMAGE_CONFIG.ALLOWED_TYPES
): FileValidationResult {
  if (!allowedTypes.includes(file.type)) {
    return {
      isValid: false,
      error: '지원하지 않는 파일 형식입니다'
    };
  }

  return { isValid: true };
}

/**
 * Validate image file
 */
export function validateImageFile(file: File): FileValidationResult {
  const sizeValidation = validateFileSize(file);
  if (!sizeValidation.isValid) {
    return sizeValidation;
  }

  const typeValidation = validateFileType(file);
  if (!typeValidation.isValid) {
    return typeValidation;
  }

  return { isValid: true };
}

/**
 * Read file as data URL
 */
export function readFileAsDataURL(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = () => {
      resolve(reader.result as string);
    };

    reader.onerror = () => {
      reject(new Error('파일을 읽을 수 없습니다'));
    };

    reader.readAsDataURL(file);
  });
}

/**
 * Read file as text
 */
export function readFileAsText(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = () => {
      resolve(reader.result as string);
    };

    reader.onerror = () => {
      reject(new Error('파일을 읽을 수 없습니다'));
    };

    reader.readAsText(file);
  });
}

/**
 * Compress image file
 */
export function compressImage(
  file: File,
  maxWidth: number = 1920,
  maxHeight: number = 1080,
  quality: number = 0.8
): Promise<Blob> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = (e) => {
      const img = new Image();

      img.onload = () => {
        const canvas = document.createElement('canvas');
        let width = img.width;
        let height = img.height;

        // Calculate new dimensions
        if (width > height) {
          if (width > maxWidth) {
            height = (height * maxWidth) / width;
            width = maxWidth;
          }
        } else {
          if (height > maxHeight) {
            width = (width * maxHeight) / height;
            height = maxHeight;
          }
        }

        canvas.width = width;
        canvas.height = height;

        const ctx = canvas.getContext('2d');
        if (!ctx) {
          reject(new Error('Canvas context를 가져올 수 없습니다'));
          return;
        }

        ctx.drawImage(img, 0, 0, width, height);

        canvas.toBlob(
          (blob) => {
            if (blob) {
              resolve(blob);
            } else {
              reject(new Error('이미지 압축에 실패했습니다'));
            }
          },
          file.type,
          quality
        );
      };

      img.onerror = () => {
        reject(new Error('이미지를 로드할 수 없습니다'));
      };

      img.src = e.target?.result as string;
    };

    reader.onerror = () => {
      reject(new Error('파일을 읽을 수 없습니다'));
    };

    reader.readAsDataURL(file);
  });
}

/**
 * Get image dimensions
 */
export function getImageDimensions(file: File): Promise<{ width: number; height: number }> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = (e) => {
      const img = new Image();

      img.onload = () => {
        resolve({
          width: img.width,
          height: img.height
        });
      };

      img.onerror = () => {
        reject(new Error('이미지를 로드할 수 없습니다'));
      };

      img.src = e.target?.result as string;
    };

    reader.onerror = () => {
      reject(new Error('파일을 읽을 수 없습니다'));
    };

    reader.readAsDataURL(file);
  });
}

/**
 * Convert blob to file
 */
export function blobToFile(blob: Blob, fileName: string): File {
  return new File([blob], fileName, { type: blob.type });
}

/**
 * Create file from data URL
 */
export function dataURLtoFile(dataURL: string, filename: string): File {
  const arr = dataURL.split(',');
  const mime = arr[0].match(/:(.*?);/)?.[1] || 'image/png';
  const bstr = atob(arr[1]);
  let n = bstr.length;
  const u8arr = new Uint8Array(n);

  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }

  return new File([u8arr], filename, { type: mime });
}

/**
 * Generate thumbnail from image
 */
export function generateThumbnail(
  file: File,
  width: number = 200,
  height: number = 200
): Promise<Blob> {
  return compressImage(file, width, height, 0.7);
}

/**
 * Check if file is image
 */
export function isImageFile(file: File): boolean {
  return file.type.startsWith('image/');
}

/**
 * Get file extension
 */
export function getFileExtension(filename: string): string {
  const parts = filename.split('.');
  return parts.length > 1 ? parts.pop()!.toLowerCase() : '';
}

/**
 * Generate unique filename
 */
export function generateUniqueFilename(originalFilename: string): string {
  const timestamp = Date.now();
  const random = Math.random().toString(36).substring(2, 8);
  const extension = getFileExtension(originalFilename);
  const nameWithoutExt = originalFilename.replace(`.${extension}`, '');

  return `${nameWithoutExt}_${timestamp}_${random}.${extension}`;
}

/**
 * Upload file with progress tracking
 */
export async function uploadFileWithProgress(
  file: File,
  uploadUrl: string,
  onProgress?: (progress: number) => void
): Promise<any> {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();

    xhr.upload.addEventListener('progress', (e) => {
      if (e.lengthComputable && onProgress) {
        const progress = (e.loaded / e.total) * 100;
        onProgress(progress);
      }
    });

    xhr.addEventListener('load', () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        try {
          const response = JSON.parse(xhr.responseText);
          resolve(response);
        } catch {
          resolve(xhr.responseText);
        }
      } else {
        reject(new Error(`Upload failed: ${xhr.statusText}`));
      }
    });

    xhr.addEventListener('error', () => {
      reject(new Error('Upload failed'));
    });

    xhr.addEventListener('abort', () => {
      reject(new Error('Upload aborted'));
    });

    const formData = new FormData();
    formData.append('file', file);

    xhr.open('POST', uploadUrl);
    xhr.send(formData);
  });
}

/**
 * Handle file drop event
 */
export function handleFileDrop(
  event: DragEvent,
  onFiles: (files: File[]) => void
): void {
  event.preventDefault();
  event.stopPropagation();

  const files = Array.from(event.dataTransfer?.files || []);
  if (files.length > 0) {
    onFiles(files);
  }
}

/**
 * Handle file input change
 */
export function handleFileInputChange(
  event: Event,
  onFiles: (files: File[]) => void
): void {
  const input = event.target as HTMLInputElement;
  const files = Array.from(input.files || []);

  if (files.length > 0) {
    onFiles(files);
  }

  // Reset input value to allow selecting the same file again
  input.value = '';
}
