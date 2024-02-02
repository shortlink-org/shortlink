import { getNodeLabel } from '@ory/integrations/ui'
import React, { ChangeEvent } from 'react'

import { NodeInputProps } from './helpers'

export function NodeInputCheckbox<T>({
  node,
  attributes,
  setValue,
  disabled,
}: NodeInputProps) {
  const errorState = node.messages.find(({ type }) => type === 'error')
    ? 'border-red-500'
    : ''

  const handleCheckboxChange = (e: ChangeEvent<HTMLInputElement>) => {
    setValue(e.target.checked)
  }

  return (
    <label className={`block ${errorState}`}>
      <input
        type="checkbox"
        name={attributes.name}
        defaultChecked={attributes.value}
        onChange={handleCheckboxChange}
        disabled={attributes.disabled || disabled}
        className="form-checkbox h-5 w-5 text-blue-600"
      />
      <span className="ml-2 text-gray-700">
        {/* @ts-ignore */}
        {getNodeLabel(node)}
      </span>
      <p className="text-red-500">
        {node.messages.map(({ text }) => text).join('\n')}
      </p>
    </label>
  )
}
