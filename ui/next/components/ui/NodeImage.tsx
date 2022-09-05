import { UiNode, UiNodeImageAttributes } from '@ory/client'

interface Props {
  node: UiNode
  attributes: UiNodeImageAttributes
}

export const NodeImage = (
  { node, attributes }: Props, // eslint-disable-line
) => (
  <img
    data-testid={`node/image/${attributes.id}`}
    src={attributes.src}
    alt={node.meta.label?.text}
  />
)
