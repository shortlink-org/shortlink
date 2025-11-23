import { useState, useCallback } from 'react'
import { ValidationResult } from '@/utils/validation'

export interface FormFieldState<T = string> {
  value: T
  error: string | null
  touched: boolean
}

export interface UseFormStateOptions<T extends Record<string, any>> {
  initialValues: T
  validate?: (values: T) => Partial<Record<keyof T, string | null>>
  onSubmit: (values: T) => Promise<void> | void
}

export function useFormState<T extends Record<string, any>>({
  initialValues,
  validate,
  onSubmit,
}: UseFormStateOptions<T>) {
  const [values, setValues] = useState<T>(initialValues)
  const [errors, setErrors] = useState<Partial<Record<keyof T, string | null>>>({})
  const [touched, setTouched] = useState<Partial<Record<keyof T, boolean>>>({})
  const [loading, setLoading] = useState(false)
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [submitSuccess, setSubmitSuccess] = useState<string | null>(null)

  const setValue = useCallback(<K extends keyof T>(field: K, value: T[K]) => {
    setValues(prev => ({ ...prev, [field]: value }))
    // Clear error when user starts typing
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: null }))
    }
  }, [errors])

  const setFieldTouched = useCallback(<K extends keyof T>(field: K) => {
    setTouched(prev => ({ ...prev, [field]: true }))
  }, [])

  const setFieldError = useCallback(<K extends keyof T>(field: K, error: string | null) => {
    setErrors(prev => ({ ...prev, [field]: error }))
  }, [])

  const validateForm = useCallback((): boolean => {
    if (!validate) return true

    const validationErrors = validate(values)
    setErrors(validationErrors)
    setTouched(
      Object.keys(values).reduce((acc, key) => {
        acc[key as keyof T] = true
        return acc
      }, {} as Partial<Record<keyof T, boolean>>)
    )

    return Object.keys(validationErrors).length === 0
  }, [values, validate])

  const handleSubmit = useCallback(async (e?: React.FormEvent) => {
    e?.preventDefault()
    
    setSubmitError(null)
    setSubmitSuccess(null)

    if (!validateForm()) {
      return
    }

    setLoading(true)
    try {
      await onSubmit(values)
      setSubmitSuccess('Saved successfully!')
    } catch (error: any) {
      setSubmitError(error.message || 'An error occurred')
    } finally {
      setLoading(false)
    }
  }, [values, validateForm, onSubmit])

  const reset = useCallback(() => {
    setValues(initialValues)
    setErrors({})
    setTouched({})
    setSubmitError(null)
    setSubmitSuccess(null)
  }, [initialValues])

  return {
    values,
    errors,
    touched,
    loading,
    submitError,
    submitSuccess,
    setValue,
    setFieldTouched,
    setFieldError,
    validateForm,
    handleSubmit,
    reset,
  }
}

