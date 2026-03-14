// @ts-ignore
import { Header as UiKitPageHeader } from '@shortlink-org/ui-kit'

import PageSection from './Section'

type AppProps = {
  title: string
  description?: string
  eyebrow?: string
}

function Header({ title, description, eyebrow }: AppProps) {
  return (
    <PageSection className="py-6 lg:py-10">
      <UiKitPageHeader title={title} description={description} eyebrow={eyebrow} />
    </PageSection>
  )
}

export default Header
