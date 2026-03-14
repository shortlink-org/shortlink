'use client'

import { BeakerIcon } from '@heroicons/react/24/solid'

import DirectoryTable, { DirectoryTableItem } from '@/components/Page/admin/directoryTable'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

const groups: DirectoryTableItem[] = [
  {
    id: 'growth-team',
    name: 'Growth Team',
    title: 'Regional Paradigm Technicians',
    department: 'Optimization',
    role: 'Admin',
    email: 'growth-team@shortlink.dev',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'product-council',
    name: 'Product Council',
    title: 'Product Directives Officers',
    department: 'Intranet',
    role: 'Owner',
    email: 'product-council@shortlink.dev',
    status: 'Invited',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'directives-squad',
    name: 'Directives Squad',
    title: 'Forward Response Developers',
    department: 'Directives',
    role: 'Member',
    email: 'directives-squad@shortlink.dev',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'security-watch',
    name: 'Security Watch',
    title: 'Central Security Managers',
    department: 'Program',
    role: 'Member',
    email: 'security-watch@shortlink.dev',
    status: 'Suspended',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'mobility-labs',
    name: 'Mobility Labs',
    title: 'Lead Implementation Liaisons',
    department: 'Mobility',
    role: 'Admin',
    email: 'mobility-labs@shortlink.dev',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'platform-ops',
    name: 'Platform Ops',
    title: 'Internal Applications Engineers',
    department: 'Security',
    role: 'Member',
    email: 'platform-ops@shortlink.dev',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
]

function GroupContent() {
  return (
    <>
      {/*<NextSeo title="Groups" description="Admin groups" />*/}

      <Header title="Admin groups" description="Review team groups, member roles and administrative status." />

      <PageSection className="space-y-6 pb-10">
        <section className="body-font text-gray-600">
          <button
            type="button"
            className="group block max-w-xs rounded-2xl bg-white p-6 text-left ring-1 ring-slate-900/5 shadow-lg transition-all hover:bg-sky-500 hover:ring-sky-500 dark:bg-gray-800"
          >
            <div className="flex items-center space-x-3">
              <BeakerIcon className="h-6 w-6 stroke-sky-500 group-hover:stroke-white" />
              <h3 className="text-sm font-semibold text-slate-900 group-hover:text-white">Create group</h3>
            </div>
            <p className="mt-3 text-sm text-slate-500 group-hover:text-white">
              Create a workspace group to assign roles, permissions and ownership boundaries.
            </p>
          </button>
        </section>

        <DirectoryTable data={groups} />
      </PageSection>
    </>
  )
}

export default withAuthSync(GroupContent)
