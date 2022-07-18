import { UiNode, UiNodeInputAttributes } from '@ory/client'
import { getNodeLabel } from '@ory/integrations/ui'
import Button from '@mui/material/Button'

import { NodeInputProps } from './helpers'

// @ts-ignore
export function NodeInputButton<T>({
  node,
  attributes,
  setValue,
  disabled,
  dispatchSubmit,
}: NodeInputProps) {
  // Some attributes have dynamic JavaScript - this is for example required for WebAuthn.
  const onClick = () => {
    // This section is only used for WebAuthn. The script is loaded via a <script> node
    // and the functions are available on the global window level. Unfortunately, there
    // is currently no better way than executing eval / function here at this moment.
    if (attributes.onclick) {
      const run = new Function(attributes.onclick)
      run()
    }
  }

  return (
    <Button
      name={attributes.name}
      onClick={(e) => {
        onClick()
        setValue(attributes.value).then(() => dispatchSubmit(e))
      }}
      className="bg-sky-600 hover:bg-sky-700"
      value={attributes.value || ''}
      disabled={attributes.disabled || disabled}
    >
      {getNodeLabel(node)}
    </Button>
  )
}
