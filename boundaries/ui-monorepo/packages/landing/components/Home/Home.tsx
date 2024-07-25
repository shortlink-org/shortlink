'use client'

import Link from 'next/link'
import { useTheme, AppBar, Tabs, Tab } from '@mui/material'
import useMediaQuery from '@mui/material/useMediaQuery'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import React, { useState } from 'react'

import TabContent from '../TabContent'
import TabPanel from '../TabPanel'

function a11yProps(index: number) {
  return {
    id: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  }
}

const Home = () => {
  const theme = useTheme()
  const [value, setValue] = useState(0)

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue)
  }

  const appBarColor = theme.palette.mode === 'dark' ? 'inherit' : 'primary'
  const textColor = theme.palette.mode === 'dark' ? 'secondary' : 'inherit'
  // @ts-ignore
  const isMobile = useMediaQuery((props) => props.breakpoints.down('sm'))

  return (
    <>
      <ToggleDarkMode id="ToggleDarkMode" />

      <div className="relative flex flex-col text-gray-700 bg-white shadow-md bg-clip-border rounded-xl min-h-60">
        <AppBar position="static" id="menu" color={appBarColor} className="mt-[10em] md:mt-0">
          <Tabs
            value={value}
            onChange={handleChange}
            indicatorColor="secondary"
            textColor={textColor}
            variant={isMobile ? 'scrollable' : 'fullWidth'}
            aria-label="full width tabs example"
            selectionFollowsFocus
            scrollButtons="auto"
            allowScrollButtonsMobile
          >
            <Tab label="ShortLink" {...a11yProps(0)} />
            <Tab label="Infrastructure" {...a11yProps(1)} />
            <Tab label="Security" {...a11yProps(2)} />
            <Tab label="Observability" {...a11yProps(3)} />
            <Tab label="Docs" {...a11yProps(4)} />
          </Tabs>
        </AppBar>

        <TabPanel value={value} index={0}>
          <TabContent
            title="UI"
            key="shortlink-ui"
            cards={[
              { name: 'Next', url: '/next' },
              { name: 'ui-kit', url: '/storybook/' },
              { name: 'shop', url: 'https://shop.shortlink.best' },
            ]}
          />

          <TabContent
            title="Shortlink API"
            key="shortlink-api"
            cards={[
              { name: 'HTTP', url: '/api' },
              { name: 'gRPC-web', url: '/grpc/' },
              { name: 'CloudEvents', url: '/cloudevents/' },
              { name: 'GraphQL', url: '/graphql/' },
              { name: 'WebSocket', url: '/ws/' },
            ]}
          />
        </TabPanel>

        <TabPanel value={value} index={1} dir={theme.direction}>
          <TabContent
            title="Infrastructure services"
            key="infrastructure"
            cards={[
              { name: 'RabbitMQ', url: '/rabbitmq/' },
              { name: 'Kafka', url: '/kafka-ui/' },
              { name: 'Keycloak', url: 'https://keycloak.shortlink.best' },
            ]}
          />

          <TabContent
            title="Argo"
            key="infrastructure"
            cards={[
              { name: 'Argo CD', url: 'https://argo.shortlink.best' },
              {
                name: 'Argo Rollout',
                url: 'https://argo.shortlink.best/rollout',
              },
              {
                name: 'Argo Workflows',
                url: 'https://workflows.shortlink.best',
              },
            ]}
          />
        </TabPanel>

        <TabPanel value={value} index={2} dir={theme.direction}>
          <TabContent
            title="Security"
            key="observability"
            cards={[
              {
                name: 'Armosec',
                url: 'https://cloud.armosec.io/compliance/shortlink',
              },
              { name: 'KubeShark', url: 'https://kubeshark.shortlink.best' },
              { name: 'Kyverno', url: '/kyverno/#/' },
            ]}
          />
        </TabPanel>

        <TabPanel value={value} index={3} dir={theme.direction}>
          <TabContent
            title="Observability services"
            key="observability"
            cards={[
              { name: 'Prometheus', url: '/prometheus' },
              { name: 'AlertManager', url: '/alertmanager' },
              { name: 'Grafana', url: 'https://grafana.shortlink.best' },
              { name: 'Pyroscope', url: 'https://pyroscope.shortlink.best' },
              { name: 'Testkube', url: 'https://testkube.shortlink.best' },
              { name: 'TraceTest', url: 'https://tracetest.shortlink.best' },
              { name: 'Status Page', url: 'https://status.shortlink.best' },
            ]}
          />
        </TabPanel>

        <TabPanel value={value} index={4} dir={theme.direction}>
          <TabContent
            title="Documentation and etc..."
            key="docs"
            cards={[
              {
                name: 'GitHub',
                url: 'https://github.com/shortlink-org/shortlink',
              },
              {
                name: 'GitLab',
                url: 'https://gitlab.com/shortlink-org/shortlink/',
              },
              {
                name: 'Swagger API',
                url: 'https://shortlink-org.gitlab.io/shortlink/',
              },
              { name: 'Backstage', url: 'https://backstage.shortlink.best/' },
              { name: 'Landscape', url: 'https://landscape.shortlink.best/' },
            ]}
          />
        </TabPanel>
      </div>

      <p className="text-sm text-gray-600 text-center my-6">
        Version: <b>{process.env.NEXT_PUBLIC_GIT_TAG}</b>
        {' && '}
        <Link href={process.env.NEXT_PUBLIC_CI_PIPELINE_URL}>
          Pipeline: <b>{process.env.NEXT_PUBLIC_PIPELINE_ID}</b>
        </Link>
      </p>
    </>
  )
}

// @ts-ignore
export default Home
