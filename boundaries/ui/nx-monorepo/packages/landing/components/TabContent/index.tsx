import { Button, Divider, Stack, Typography, useTheme } from '@mui/material'
import Link from 'next/link'
import * as React from 'react'

interface Card {
  name: string
  url: string
}

interface TabContentProps {
  title: string
  cards: Card[]
}

const TabContent: React.FC<TabContentProps> = ({ title, cards }) => {
  const theme = useTheme()

  return (
    <div className="my-5">
      <Typography
        variant="h5"
        align="center"
        color={theme.palette.mode === 'dark' ? 'primary' : 'inherit'}
      >
        {title}
      </Typography>
      <Stack
        spacing={{ xs: 1, sm: 1, md: 2 }}
        direction={{ xs: 'column', sm: 'row' }}
        divider={<Divider orientation="vertical" flexItem />}
        mt={2}
        justifyContent="center"
        alignItems="center"
        useFlexGap
        flexWrap="wrap"
      >
        {cards.map((card) => getCard(card.name, card.url, theme))}
      </Stack>
    </div>
  )
}

// @ts-ignore
function getCard(name: string, url: string, theme) {
  return (
    <Link href={url} key={url} passHref>
      <Button
        variant="outlined"
        color={theme.palette.mode === 'dark' ? 'primary' : 'inherit'}
        size="large"
      >
        {name}
      </Button>
    </Link>
  )
}

export default TabContent
