import { UiNode, UiNodeTextAttributes, UiText } from '@ory/client'
import React from 'react'

interface Props {
  node: UiNode
  attributes: UiNodeTextAttributes
}

const Content = ({ attributes }: Props) => {
  // eslint-disable-line
  switch (attributes.text.id) {
    case 1050015: {
      // This text node contains lookup secrets. Let's make them a bit more beautiful!
      const secrets = (attributes.text.context as any).secrets.map(
        // eslint-disable-line
        (text: UiText, k: number) => (
          <div
            key={k} // eslint-disable-line
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

export const NodeText = ({ node, attributes }: Props) => (
  <React.Fragment>
    <p className="font-normal" data-testid={`node/text/${attributes.id}/label`}>
      {node.meta?.label?.text}
    </p>
    <Content node={node} attributes={attributes} />
  </React.Fragment>
)
