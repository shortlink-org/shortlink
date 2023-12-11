import type { StoryObj, Meta } from '@storybook/react'

import DiscountedComponent from '../../components/Billing/Discounted'

export default {
  /* 👇 The title prop is optional.
   * See https://storybook.js.org/docs/7.0/react/configure/overview#configure-story-loading
   * to learn how to generate automatic titles
   */
  title: 'Billing',
  component: DiscountedComponent,
  parameters: {},
} as Meta<any>

/*
 *👇 Render functions are a framework-specific feature to allow you to control on how the component renders.
 * See https://storybook.js.org/docs/7.0/react/api/csf
 * to learn how to use render functions.
 */
export const Discounted: StoryObj<typeof DiscountedComponent> = {
  render: () => <DiscountedComponent />,
}
