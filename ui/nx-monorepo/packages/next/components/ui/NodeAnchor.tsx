import { UiNode, UiNodeAnchorAttributes } from '@ory/client'
import Button from '@mui/material/Button'

interface Props {
  node: UiNode
  attributes: UiNodeAnchorAttributes
}

// @ts-ignore
export const NodeAnchor = (
  { node, attributes }: Props, // eslint-disable-line
) => (
  <Button
    onClick={(e) => {
      e.stopPropagation()
      e.preventDefault()
      window.location.href = attributes.href
    }}
  >
    {attributes.title.text}
  </Button>
)

export default NodeAnchor
