'use client'

import { BeakerIcon } from '@heroicons/react/24/solid'

import DirectoryTable, { DirectoryTableItem } from '@/components/Page/admin/directoryTable'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

const people: DirectoryTableItem[] = [
  {
    id: 'jane-cooper',
    name: 'Jane Cooper',
    title: 'Regional Paradigm Technician',
    department: 'Optimization',
    role: 'Admin',
    email: 'jane.cooper@example.com',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'cody-fisher',
    name: 'Cody Fisher',
    title: 'Product Directives Officer',
    department: 'Intranet',
    role: 'Owner',
    email: 'cody.fisher@example.com',
    status: 'Invited',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'esther-howard',
    name: 'Esther Howard',
    title: 'Forward Response Developer',
    department: 'Directives',
    role: 'Member',
    email: 'esther.howard@example.com',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'jenny-wilson',
    name: 'Jenny Wilson',
    title: 'Central Security Manager',
    department: 'Program',
    role: 'Member',
    email: 'jenny.wilson@example.com',
    status: 'Suspended',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'kristin-watson',
    name: 'Kristin Watson',
    title: 'Lead Implementation Liaison',
    department: 'Mobility',
    role: 'Admin',
    email: 'kristin.watson@example.com',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
  {
    id: 'cameron-williamson',
    name: 'Cameron Williamson',
    title: 'Internal Applications Engineer',
    department: 'Security',
    role: 'Member',
    email: 'cameron.williamson@example.com',
    status: 'Active',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
  },
]

function Page() {
  return (
    <>
      {/*<NextSeo title="Users" description="Admin users page" />*/}

      <Header title="Admin users" description="Browse and manage user accounts, roles and membership status." />

      <PageSection className="space-y-6 pb-10">
        <section className="text-gray-600 body-font">
          <button
            type="button"
            className="group block max-w-xs rounded-2xl bg-white p-6 text-left ring-1 ring-slate-900/5 shadow-lg transition-all hover:bg-sky-500 hover:ring-sky-500 dark:bg-gray-800"
          >
            <div className="flex items-center space-x-3">
              <BeakerIcon className="h-6 w-6 stroke-sky-500 group-hover:stroke-white" />
              <h3 className="text-sm font-semibold text-slate-900 group-hover:text-white">Invite user</h3>
            </div>
            <p className="mt-3 text-sm text-slate-500 group-hover:text-white">
              Start a new invitation flow for teammates, admins and workspace owners.
            </p>
          </button>
        </section>

        <DirectoryTable data={people} />
      </PageSection>
    </>
  )
}

export default withAuthSync(Page)
