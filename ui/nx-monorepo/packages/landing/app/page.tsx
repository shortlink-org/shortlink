'use client'

import { NextPage } from 'next'
import React, { useState } from 'react'
import { useTheme, AppBar, Grid, Tabs, Tab, Box } from '@mui/material'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import TabPanel from '../components/TabPanel'
import TabContent from '../components/TabContent'

function a11yProps(index: number) {
  return {
    id: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  }
}

const Home: NextPage = () => {
  const theme = useTheme()
  const [value, setValue] = useState(0)

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue)
  }

  return (
    <React.Fragment>
      <ToggleDarkMode id="ToggleDarkMode" />

      <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
      >
        <Box sx={{ bgcolor: 'background.paper', width: 800 }}>
          <AppBar position="static" id="menu">
            <Tabs
              value={value}
              onChange={handleChange}
              indicatorColor="secondary"
              textColor="inherit"
              variant="fullWidth"
              aria-label="full width tabs example"
              selectionFollowsFocus
            >
              <Tab label="Application" {...a11yProps(0)} />
              <Tab label="Infrastructure" {...a11yProps(1)} />
              <Tab label="Observability" {...a11yProps(2)} />
              <Tab label="Docs" {...a11yProps(3)} />
            </Tabs>
          </AppBar>

          <TabPanel value={value} index={0}>
            <TabContent
              title="Shortlink service (Microservice example)"
              key="shortlink"
              cards={[
                { name: 'Next', url: '/next' },
                { name: 'ui-kit', url: '/storybook/' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={1} dir={theme.direction}>
            <TabContent
              title="Infrastructure services"
              key="infrastructure"
              cards={[
                { name: 'RabbitMQ', url: '/rabbitmq/' },
                { name: 'Kafka', url: '/kafka-ui' },
                { name: 'Argo CD', url: '/argo/cd' },
                { name: 'Argo Workflows', url: '/argo/workflows' },
                { name: 'Keycloak', url: 'https://keycloak.shortlink.best' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={2} dir={theme.direction}>
            <TabContent
              title="Observability services"
              key="observability"
              cards={[
                { name: 'Prometheus', url: '/prometheus' },
                { name: 'AlertManager', url: '/alertmanager' },
                { name: 'Grafana', url: 'https://grafana.shortlink.best' },
                { name: 'Pyroscope', url: 'https://pyroscope.shortlink.best' },
                { name: 'Kyverno', url: '/kyverno/' },
                { name: 'Testkube', url: 'https://testkube.shortlink.best' },
                { name: 'TraceTest', url: 'https://tracetest.shortlink.best' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={3} dir={theme.direction}>
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
              ]}
            />
          </TabPanel>
        </Box>
      </Grid>
    </React.Fragment>
  )
}

// @ts-ignore
export default Home
