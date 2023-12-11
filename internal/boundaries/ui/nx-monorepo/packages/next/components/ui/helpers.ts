import { UiNode, UiNodeInputAttributes } from '@ory/client'
import { FormEvent, MouseEvent } from 'react'

export type ValueSetter = (
  value: string | number | boolean | undefined,
) => Promise<void>

export type FormDispatcher = (e: MouseEvent | React.FormEvent) => Promise<void>

export interface NodeInputProps {
  node: UiNode
  attributes: UiNodeInputAttributes
  value: any
  disabled: boolean
  dispatchSubmit: FormDispatcher
  setValue: ValueSetter
}
