import { NodeInputProps } from './helpers'

export function NodeInputHidden<T>({ attributes }: NodeInputProps) {
  // Render a hidden input field
  return (
    <input
      type={attributes.type}
      name={attributes.name}
      value={attributes.value || 'true'}
    />
  )
}
