import { Divider, Stack, Typography, useTheme } from '@mui/material'
import Button from '@mui/material/Button'
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
    <div className="my-5 max-w-4xl">
      <h2 className="prose text-center my-5 dark:text-white">{title}</h2>

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
        size="large"
      >
        {name}
      </Button>
    </Link>
  )
}

export default TabContent
