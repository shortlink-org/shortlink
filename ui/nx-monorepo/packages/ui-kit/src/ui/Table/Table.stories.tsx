import { Meta } from '@storybook/react'
import React from 'react'

import Table from './Table'

const meta: Meta<any> = {
  title: 'UI/Table',
  component: Table,
}

export default meta

function Template(args: any) {
  return <Table {...args} />
}

export const Default = Template.bind({})

// @ts-ignore
Default.args = {
  data: [
    {
      url: 'https://batazor.ru',
      hash: 'myHash1',
      describe: 'My personal website',
      createdAt: '1970-01-01T00:00:12.500908Z',
      updatedAt: '1970-01-01T00:00:12.500908Z',
    },
    {
      url: 'https://github.com/batazor',
      hash: 'myHash2',
      describe: 'My accout of github',
      createdAt: '1970-01-01T00:00:12.500908Z',
      updatedAt: '1970-01-01T00:00:12.500908Z',
    },
    {
      url: 'https://vk.com/batazor',
      hash: 'myHash3',
      describe: 'My page on vk.com',
      createdAt: '1970-01-01T00:00:12.500908Z',
      updatedAt: '1970-01-01T00:00:12.500908Z',
    },
  ],
  columns: [
    {
      accessorKey: 'url',
      header: 'URL',
      size: 150,
    },
    {
      accessorKey: 'hash',
      header: 'Hash',
      size: 150,
    },
    {
      accessorKey: 'describe',
      header: 'Describe',
      size: 150,
    },
    {
      accessorKey: 'createdAt',
      header: 'Created at',
      size: 150,
    },
    {
      accessorKey: 'updatedAt',
      header: 'Updated at',
      size: 150,
    },
  ],
}
