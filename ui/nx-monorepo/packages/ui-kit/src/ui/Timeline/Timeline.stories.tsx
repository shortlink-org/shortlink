// @ts-nocheck
import { BeakerIcon } from '@heroicons/react/20/solid'
import { Meta, StoryFn } from '@storybook/react'

import Timeline, { TimelineProps } from './Timeline'

interface Event {
  date: string
  description: string
}

const meta: Meta<TimelineProps> = {
  title: 'UI/Timeline',
  component: Timeline,
}

export default meta

const Template: StoryFn<TimelineProps> = (args) => <Timeline {...args} />

export const Default = Template.bind({})
Default.args = {
  items: [
    {
      date: 'Apr 7, 2024',
      name: 'Mark Mikrol',
      action: 'opened the request',
      content:
        'Various versions have evolved over the years, sometimes by accident, sometimes on purpose injected humour and the like.',
      icon: <BeakerIcon width="16" height="16" className="fill-emerald-500" />,
    },
    {
      date: 'Apr 7, 2024',
      name: 'Mark Mikrol',
      action: 'opened the request',
      content: 'Various versions have evolved over the years.',
      icon: <BeakerIcon width="16" height="16" className="fill-gray-500" />,
    },
    {
      date: 'Apr 7, 2024',
      name: 'Mark Mikrol',
      action: 'opened the request',
      content: 'Various versions have evolved over the years.',
      icon: <BeakerIcon width="16" height="16" className="fill-gray-500" />,
    },
    {
      date: 'Apr 7, 2024',
      name: 'Mark Mikrol',
      action: 'opened the request',
      content: 'Various versions have evolved over the years.',
      icon: <BeakerIcon width="16" height="16" className="fill-red-500" />,
    },
  ],
}
