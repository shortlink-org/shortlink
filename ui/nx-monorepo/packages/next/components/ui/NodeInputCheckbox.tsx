import { getNodeLabel } from '@ory/integrations/ui'
import { Checkbox } from '@ory/themes'

import { NodeInputProps } from './helpers'

export function NodeInputCheckbox<T>({
  // eslint-disable-line
  node,
  attributes,
  setValue,
  disabled,
}: NodeInputProps) {
  // Render a checkbox.s
  return (
    <Checkbox
      name={attributes.name}
      defaultChecked={attributes.value}
      onChange={(e) => setValue(e.target.checked)}
      disabled={attributes.disabled || disabled}
      label={
        // @ts-ignore
        getNodeLabel(node)
      }
      state={
        node.messages.find(({ type }) => type === 'error') ? 'error' : undefined
      }
      subtitle={node.messages.map(({ text }) => text).join('\n')}
    />
  )
}
