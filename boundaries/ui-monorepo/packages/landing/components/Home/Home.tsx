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

      <div className="relative flex flex-col text-gray-700 bg-white shadow-lg bg-clip-border rounded-2xl max-w-4xl mx-auto mt-12">
        <AppBar position="static" id="menu" color={appBarColor} className="mt-[10em] md:mt-0">
          <Tabs
            value={value}
            onChange={handleChange}
            indicatorColor="secondary"
            textColor={textColor}
            variant={isMobile ? 'scrollable' : 'fullWidth'}
            aria-label="scrollable full width tabs example"
            selectionFollowsFocus
            scrollButtons="auto"
            allowScrollButtonsMobile
            className="md:max-w-3xl mx-auto"
          >
            <Tab label="ShortLink" {...a11yProps(0)} />
            <Tab label="Shop" {...a11yProps(1)} />
            <Tab label="Infrastructure" {...a11yProps(2)} />
            <Tab label="Security" {...a11yProps(3)} />
            <Tab label="Observability" {...a11yProps(4)} />
            <Tab label="Docs" {...a11yProps(5)} />
          </Tabs>
        </AppBar>

        <TabPanel value={value} index={0}>
          <TabContent
            title="UI"
            key="shortlink-ui"
            cards={[
              { name: 'Next', url: '/next' },
              { name: 'ui-kit', url: '/storybook/' },
            ]}
          />

          <TabContent title="Shortlink API" key="shortlink-api" cards={[{ name: 'HTTP', url: '/api' }]} />
        </TabPanel>

        <TabPanel value={value} index={1}>
          <TabContent
            title="Shop"
            key="shortlink-shop"
            cards={[
              { name: 'Shop', url: 'https://shop.shortlink.best' },
              { name: 'Admin', url: 'https://shop.shortlink.best/admin' },
              { name: 'Temporal', url: 'https://temporal.shortlink.best' },
              { name: 'Storybook', url: 'https://shop.shortlink.best/storybook' },
            ]}
          />
        </TabPanel>

        <TabPanel value={value} index={2} dir={theme.direction}>
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
            key="argo"
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

        <TabPanel value={value} index={3} dir={theme.direction}>
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

        <TabPanel value={value} index={4} dir={theme.direction}>
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

        <TabPanel value={value} index={5} dir={theme.direction}>
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

      <p className="text-sm text-gray-600 text-center my-4">
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
