import { NextPage } from 'next'
import Script from 'next/script'
import Head from 'next/head'
import React, { useState } from 'react'
import { ArticleJsonLd, BreadcrumbJsonLd, NextSeo } from 'next-seo'
import {
  useTheme,
  AppBar,
  Grid,
  Tabs,
  Tab,
  Box,
} from '@mui/material'
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

  const GA_ID = process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS_ID

  return (
    <div>
      <Script
        src={`https://www.googletagmanager.com/gtag/js?id=${GA_ID}`}
        strategy="afterInteractive"
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', '${GA_ID}');
        `}
      </Script>

      <NextSeo
        title="Routing by project"
        description="Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you."
        openGraph={{
          title: 'Routing by project',
          description:
            "Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you.",
          type: 'article',
          article: {
            publishedTime: '2021-08-01T05:00:00.000Z',
            modifiedTime: '2021-08-01T05:00:00.000Z',
            section: 'FAQ',
            authors: ['https://batazor.ru'],
            tags: ['shortlink'],
          },
        }}
      />
      <ArticleJsonLd
        url="https://shortlink.best/"
        title="Main"
        images={['https://shortlink.best/images/logo.png']}
        datePublished="2021-08-01T05:00:00.000Z"
        dateModified="2021-08-01T05:00:00.000Z"
        authorName={[
          {
            name: 'Login Viktor',
            url: 'https://batazor.ru',
          },
        ]}
        publisherName="Login Viktor"
        publisherLogo="https://shortlink.best/images/logo.png"
        description="Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you."
      />
      <BreadcrumbJsonLd
        itemListElements={[
          {
            position: 1,
            name: 'Next UI',
            item: 'https://shortlink.best/next',
          },
          {
            position: 2,
            name: 'Prometheus',
            item: 'https://shortlink.best/prometheus/',
          },
          {
            position: 3,
            name: 'Grafana',
            item: 'https://grafana.shortlink.best',
          },
          {
            position: 4,
            name: 'Argo CD',
            item: 'https://shortlink.best/argo/cd/',
          },
          {
            position: 5,
            name: 'GitHab',
            item: 'https://github.com/shortlink-org/shortlink',
          },
        ]}
      />

      <Head>
        <link rel="icon" href="/favicon.ico" />
      </Head>

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
              cards={[
                { name: 'Next', url: '/next' },
                { name: 'ui-kit', url: '/storybook/' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={1} dir={theme.direction}>
            <TabContent
              title="Infrastructure services"
              cards={[
                { name: 'RabbitMQ', url: '/rabbitmq/' },
                { name: 'Kafka', url: '/kafka-ui' },
                { name: 'Argo CD', url: '/argo/cd' },
                { name: 'Argo Workflows', url: '/argo/workflows' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={2} dir={theme.direction}>
            <TabContent
              title="Observability services"
              cards={[
                { name: 'Prometheus', url: '/prometheus' },
                { name: 'AlertManager', url: '/alertmanager' },
                { name: 'Grafana', url: 'https://grafana.shortlink.best' },
                { name: 'Pyroscope', url: 'https://pyroscope.shortlink.best' },
                { name: 'Kyverno', url: '/kyverno/' },
              ]}
            />
          </TabPanel>

          <TabPanel value={value} index={3} dir={theme.direction}>
            <TabContent
              title="Documentation and etc..."
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
    </div>
  )
}

// @ts-ignore
export default Home
