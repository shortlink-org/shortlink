import { UiNode } from '@ory/client'
import {
  isUiNodeAnchorAttributes,
  isUiNodeImageAttributes,
  isUiNodeInputAttributes,
  isUiNodeScriptAttributes,
  isUiNodeTextAttributes
} from '@ory/integrations/ui'

import { NodeAnchor } from './NodeAnchor'
import { NodeImage } from './NodeImage'
import { NodeInput } from './NodeInput'
import { NodeScript } from './NodeScript'
import { NodeText } from './NodeText'
import { FormDispatcher, ValueSetter } from './helpers'

interface Props {
  node: UiNode
  disabled: boolean
  value: any
  setValue: ValueSetter
  dispatchSubmit: FormDispatcher
}

export const Node = ({
  node,
  value,
  setValue,
  disabled,
  dispatchSubmit
}: Props) => {
  if (isUiNodeImageAttributes(node.attributes)) {
    return <NodeImage node={node} attributes={node.attributes} />
  }

  if (isUiNodeScriptAttributes(node.attributes)) {
    return <NodeScript node={node} attributes={node.attributes} />
  }

  if (isUiNodeTextAttributes(node.attributes)) {
    return <NodeText node={node} attributes={node.attributes} />
  }

  if (isUiNodeAnchorAttributes(node.attributes)) {
    return <NodeAnchor node={node} attributes={node.attributes} />
  }

  if (isUiNodeInputAttributes(node.attributes)) {
    return (
      <NodeInput
        dispatchSubmit={dispatchSubmit}
        value={value}
        setValue={setValue}
        node={node}
        disabled={disabled}
        attributes={node.attributes}
      />
    )
  }

  return null
}
