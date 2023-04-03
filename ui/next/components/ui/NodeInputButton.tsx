import { getNodeLabel } from '@ory/integrations/ui'
import Button from '@mui/material/Button'
import { FormEvent } from 'react'

import { NodeInputProps } from './helpers'

// @ts-ignore
export function NodeInputButton<T>({
  // eslint-disable-line
  node,
  attributes,
  setValue,
  disabled,
  dispatchSubmit,
}: NodeInputProps) {
  // Some attributes have dynamic JavaScript - this is for example required for WebAuthn.
  const onClick = (e: React.MouseEvent | React.FormEvent<HTMLFormElement>) => {
    // This section is only used for WebAuthn. The script is loaded via a <script> node
    // and the functions are available on the global window level. Unfortunately, there
    // is currently no better way than executing eval / function here at this moment.
    //
    // Please note that we also need to prevent the default action from happening.
    if (attributes.onclick) {
      e.stopPropagation()
      e.preventDefault()

      const run = new Function(attributes.onclick)
      run()
      return
    }

    setValue(attributes.value).then(() => dispatchSubmit(e))
  }

  return (
    <Button
      name={attributes.name}
      onClick={onClick}
      className="bg-sky-600 hover:bg-sky-700"
      value={attributes.value || ''}
      disabled={attributes.disabled || disabled}
    >
      {
        // @ts-ignore
        getNodeLabel(node)
      }
    </Button>
  )
}
