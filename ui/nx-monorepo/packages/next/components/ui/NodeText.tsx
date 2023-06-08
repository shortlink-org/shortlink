import { UiNode, UiNodeTextAttributes, UiText } from '@ory/client'
import * as React from 'react'

interface Props {
  node: UiNode
  attributes: UiNodeTextAttributes
}

function Content({ attributes }: Props) {
  switch (attributes.text.id) {
    case 1050015: {
      // This text node contains lookup secrets. Let's make them a bit more beautiful!
      const secrets = (attributes.text.context as any).secrets.map(
        (text: UiText, k: number) => (
          <div
            key={k}
            data-testid={`node/text/${attributes.id}/lookup_secret`}
            className="col-xs-3"
          >
            {/* Used lookup_secret has ID 1050014 */}
            <code>{text.id === 1050014 ? 'Used' : text.text}</code>
          </div>
        ),
      )
      return (
        <div
          className="container-fluid"
          data-testid={`node/text/${attributes.id}/text`}
        >
          <div className="row">{secrets}</div>
        </div>
      )
    }
    default:
    // Otherwise, we nothitng - the error will be handled by the Flow component
  }

  return (
    <div
      className="overflow-x-auto"
      data-testid={`node/text/${attributes.id}/text`}
    >
      <pre className="whitespace-pre overflow-x-scroll">
        {attributes.text.text}
      </pre>
    </div>
  )
}

export function NodeText({ node, attributes }: Props) {
  return (
    <>
      <p
        className="font-normal"
        data-testid={`node/text/${attributes.id}/label`}
      >
        {node.meta?.label?.text}
      </p>
      <Content node={node} attributes={attributes} />
    </>
  )
}
