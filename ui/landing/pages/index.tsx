import type { NextPage } from 'next'
import Script from 'next/script'
import Head from 'next/head'
import * as React from 'react'
import { ArticleJsonLd, BreadcrumbJsonLd, NextSeo } from 'next-seo'
// @ts-ignore
import Divider from '@mui/material/Divider'
import { useTheme } from '@mui/material/styles'
import AppBar from '@mui/material/AppBar'
import Grid from '@mui/material/Grid'
import Stack from '@mui/material/Stack'
import Tabs from '@mui/material/Tabs'
import Button from '@mui/material/Button'
import Tab from '@mui/material/Tab'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Link from 'next/link'
// @ts-ignore
import { ToggleDarkMode } from '@shortlink-org/ui-kit'
import TabPanel from '../components/TabPanel'

function a11yProps(index: number) {
  return {
    id: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  }
}

function getCard(name: string, url: string) {
  return (
    <Link href={url} legacyBehavior>
      <Button variant="outlined">{name}</Button>
    </Link>
  )
}

const Home: NextPage = () => {
  const theme = useTheme()
  const [value, setValue] = React.useState(0)

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue)
  }

  return (
    <div>
      <Script
        src="https://www.googletagmanager.com/gtag/js?id=G-DBZDFPJCJ9"
        strategy="afterInteractive"
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', 'G-DBZDFPJCJ9');
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
            item: 'https://shortlink.best/grafana/',
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
        <Box sx={{ bgcolor: 'background.paper', width: 700 }}>
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
            <Typography variant="h5" align="center">
              Shortlink service (Microservice example)
            </Typography>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('Next', '/next')}
              {getCard('ui-kit', '/storybook/')}
            </Stack>
          </TabPanel>

          <TabPanel value={value} index={1} dir={theme.direction}>
            <Typography variant="h5" align="center">
              Infrastructure services
            </Typography>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('RabbitMQ', '/rabbitmq')}

              {getCard('Kafka', '/kafka-ui')}
            </Stack>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('Argo CD', '/argo/cd')}
              {getCard('Argo Workflows', '/argo/workflows')}
              {/*{getCard("Argo Dashboard", "/argo/dashboard")}*/}
            </Stack>
          </TabPanel>

          <TabPanel value={value} index={2} dir={theme.direction}>
            <Typography variant="h5" align="center">
              Observability
            </Typography>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('Prometheus', '/prometheus')}
              {getCard('AlertManager', '/alertmanager')}
            </Stack>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('Grafana', '/grafana')}

              {getCard('Pyroscope', 'https://pyroscope.shortlink.best')}

              {getCard('Kyverno', '/kyverno/#/')}
            </Stack>
          </TabPanel>

          <TabPanel value={value} index={3} dir={theme.direction}>
            <Typography variant="h5" align="center">
              Documentation and etc...
            </Typography>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard('GitHub', 'https://github.com/shortlink-org/shortlink')}

              {getCard('GitLab', 'https://gitlab.com/shortlink-org/shortlink/')}

              {getCard(
                'Swagger API',
                'https://shortlink-org.gitlab.io/shortlink/',
              )}

              {getCard('Backstage', 'https://shortlink.best/backstage')}
            </Stack>
          </TabPanel>
        </Box>
      </Grid>
    </div>
  )
}

export default Home
