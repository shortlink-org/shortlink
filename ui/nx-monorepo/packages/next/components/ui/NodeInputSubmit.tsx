import Button from '@mui/material/Button'
import { getNodeLabel } from '@ory/integrations/ui'

import { NodeInputProps } from './helpers'

// @ts-ignore
export function NodeInputSubmit<T>({
  node,
  attributes,
  setValue,
  disabled,
  dispatchSubmit,
}: NodeInputProps) {
  return (
    <Button
      name={attributes.name}
      variant="contained"
      color="primary"
      type="submit"
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
