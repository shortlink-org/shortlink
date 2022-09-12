import type { NextPage } from 'next'
import Script from 'next/script'
import Head from 'next/head'
import * as React from 'react'
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

interface TabPanelProps {
  children?: React.ReactNode;
  dir?: string;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `full-width-tab-${index}`,
    'aria-controls': `full-width-tabpanel-${index}`,
  };
}

function getCard(name: string, url: string) {
  return (
    <Link href={url}>
      <Button variant="outlined">
        {name}
      </Button>
    </Link>
  )
}

const Home: NextPage = () => {
  const theme = useTheme();
  const [value, setValue] = React.useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

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
      
      <Head>
        <title>Shortlink | Landing</title>
        <meta name="description" content="Routing by project" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <ToggleDarkMode />

      <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
      >
        <Box sx={{ bgcolor: 'background.paper', width: 700 }}>
          <AppBar position="static">
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
              <Tab label="Docs" {...a11yProps(2)} />
            </Tabs>
          </AppBar>

          <TabPanel value={value} index={0}>
            <Typography variant="h5" align={"center"}>
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
              {getCard("Next", "/next")}
            </Stack>
          </TabPanel>

          <TabPanel value={value} index={1} dir={theme.direction}>
            <Typography variant="h5" align={"center"}>
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
              {getCard("Prometheus", "/prometheus")}
              {getCard("AlertManager", "/alertmanager")}
            </Stack>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard("Grafana", "/grafana")}

              {getCard("RabbitMQ", "/rabbitmq")}

              {getCard("Kyverno", "/kyverno/#/")}
            </Stack>

            <Stack
              spacing={{ xs: 1, sm: 2, md: 4 }}
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
              mt={2}
              justifyContent="center"
              alignItems="center"
            >
              {getCard("Argo CD", "/argo/cd")}
              {getCard("Argo Workflows", "/argo/workflows")}
              {/*{getCard("Argo Dashboard", "/argo/dashboard")}*/}
            </Stack>
          </TabPanel>

          <TabPanel value={value} index={2} dir={theme.direction}>
            <Typography variant="h5" align={"center"}>
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
              {getCard("GitHub", "https://github.com/batazor/shortlink")}

              {getCard("GitLab", "https://gitlab.com/shortlink-org/shortlink/")}

              {getCard("Swagger API", "https://shortlink-org.gitlab.io/shortlink/")}
            </Stack>
          </TabPanel>
        </Box>
      </Grid>
    </div>
  )
}

export default Home
