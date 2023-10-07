import Table from './Table'
import { StoryObj, Meta, Preview } from '@storybook/react'

const meta: Meta<any> = {
  title: 'UI/Table',
  component: Table,
}

export default meta

function Template(args: any) {
  args = {
    list: [
      {"url":"https://batazor.ru","hash":"myHash1","describe":"My personal website","createdAt":"1970-01-01T00:00:12.500908Z","updatedAt":"1970-01-01T00:00:12.500908Z"},
      {"url":"https://github.com/batazor","hash":"myHash2","describe":"My accout of github","createdAt":"1970-01-01T00:00:12.500908Z","updatedAt":"1970-01-01T00:00:12.500908Z"},
      {"url":"https://vk.com/batazor","hash":"myHash3","describe":"My page on vk.com","createdAt":"1970-01-01T00:00:12.500908Z","updatedAt":"1970-01-01T00:00:12.500908Z"},
    ]
  }

  return <Table {...args}>Text</Table>
}

export const Default = Template.bind({})
// @ts-ignore
Default.args = {}
