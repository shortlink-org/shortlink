import TextField from '@mui/material/TextField'

import { NodeInputProps } from './helpers'

export function NodeInputDefault<T>(props: NodeInputProps) {
  const { node, attributes, value = '', setValue, disabled } = props

  // Some attributes have dynamic JavaScript - this is for example required for WebAuthn.
  const onClick = () => {
    // This section is only used for WebAuthn.
    // The script is loaded via a <script> node,
    // and the functions are available on the global window level.
    // Unfortunately, there
    // is currently no better way than executing eval / function here at this moment.
    if (attributes.onclick) {
      // eslint-disable-next-line no-new-func
      const run = new Function(attributes.onclick)
      run()
    }
  }

  // Render a generic text input field.
  //       error={node.messages.find(({ type }) => type === 'error') ? 'error' : undefined}
  return (
    <TextField
      name={attributes.name}
      id={node.meta.label?.text}
      type={attributes.type}
      required
      fullWidth
      // variant={value}
      label={node.meta.label?.text}
      value={value}
      disabled={attributes.disabled || disabled}
      error={!!node.messages.find(({ type }) => type === 'error')}
      onClick={onClick}
      onChange={(e) => {
        setValue(e.target.value)
      }}
    />
  )
}
