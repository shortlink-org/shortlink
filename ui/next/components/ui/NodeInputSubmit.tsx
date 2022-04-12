import { getNodeLabel } from '@ory/integrations/ui'
import Button from '@mui/material/Button'

/* eslint-disable */
// @ts-ignore
export function NodeInputSubmit<T>({ node, attributes, setValue, disabled, dispatchSubmit}) {
  return (
    <Button
      name={attributes.name}
      variant="contained"
      color="primary"
      type="submit"
      onClick={e => {
        // On click, we set this value, and once set, dispatch the submission!
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
