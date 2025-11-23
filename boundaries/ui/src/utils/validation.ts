/**
 * Common validation utilities for forms
 */

export interface ValidationResult {
  isValid: boolean
  error?: string
}

/**
 * Validates a URL
 */
export function validateUrl(url: string, required = false): ValidationResult {
  if (!url) {
    return required 
      ? { isValid: false, error: 'URL is required' }
      : { isValid: true }
  }

  try {
    // Try with http:// if no protocol
    const urlToValidate = url.startsWith('http://') || url.startsWith('https://') 
      ? url 
      : `https://${url}`
    
    new URL(urlToValidate)
    return { isValid: true }
  } catch {
    return { isValid: false, error: 'Please enter a valid URL' }
  }
}

/**
 * Validates an email address
 */
export function validateEmail(email: string, required = false): ValidationResult {
  if (!email) {
    return required
      ? { isValid: false, error: 'Email is required' }
      : { isValid: true }
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(email)) {
    return { isValid: false, error: 'Please enter a valid email address' }
  }

  return { isValid: true }
}

/**
 * Validates a password with strength requirements
 */
export function validatePassword(
  password: string,
  options: {
    minLength?: number
    requireUppercase?: boolean
    requireLowercase?: boolean
    requireNumber?: boolean
    requireSpecial?: boolean
  } = {}
): ValidationResult {
  const {
    minLength = 8,
    requireUppercase = true,
    requireLowercase = true,
    requireNumber = true,
    requireSpecial = false,
  } = options

  if (!password) {
    return { isValid: false, error: 'Password is required' }
  }

  if (password.length < minLength) {
    return { isValid: false, error: `Password must be at least ${minLength} characters long` }
  }

  if (requireUppercase && !/(?=.*[A-Z])/.test(password)) {
    return { isValid: false, error: 'Password must contain at least one uppercase letter' }
  }

  if (requireLowercase && !/(?=.*[a-z])/.test(password)) {
    return { isValid: false, error: 'Password must contain at least one lowercase letter' }
  }

  if (requireNumber && !/(?=.*\d)/.test(password)) {
    return { isValid: false, error: 'Password must contain at least one number' }
  }

  if (requireSpecial && !/(?=.*[!@#$%^&*])/.test(password)) {
    return { isValid: false, error: 'Password must contain at least one special character' }
  }

  return { isValid: true }
}

/**
 * Validates that two passwords match
 */
export function validatePasswordMatch(
  password: string,
  confirmPassword: string
): ValidationResult {
  if (!confirmPassword) {
    return { isValid: false, error: 'Please confirm your password' }
  }

  if (password !== confirmPassword) {
    return { isValid: false, error: 'Passwords do not match' }
  }

  return { isValid: true }
}

/**
 * Validates required field
 */
export function validateRequired(value: string | null | undefined, fieldName = 'This field'): ValidationResult {
  if (!value || value.trim() === '') {
    return { isValid: false, error: `${fieldName} is required` }
  }
  return { isValid: true }
}

/**
 * Validates minimum length
 */
export function validateMinLength(
  value: string,
  minLength: number,
  fieldName = 'This field'
): ValidationResult {
  if (value.length < minLength) {
    return { isValid: false, error: `${fieldName} must be at least ${minLength} characters` }
  }
  return { isValid: true }
}

/**
 * Validates maximum length
 */
export function validateMaxLength(
  value: string,
  maxLength: number,
  fieldName = 'This field'
): ValidationResult {
  if (value.length > maxLength) {
    return { isValid: false, error: `${fieldName} must be no more than ${maxLength} characters` }
  }
  return { isValid: true }
}

