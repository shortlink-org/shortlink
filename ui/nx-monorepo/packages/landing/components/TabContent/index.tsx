import Link from 'next/link'
import React from 'react'
import { Button, Divider, Stack, Typography } from '@mui/material'

interface Card {
  name: string
  url: string
}

interface TabContentProps {
  title: string
  cards: Card[]
}

const TabContent: React.FC<TabContentProps> = ({ title, cards }) => (
  <React.Fragment>
    <Typography variant="h5" align="center">
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
      {cards.map((card) => getCard(card.name, card.url))}
    </Stack>
  </React.Fragment>
)

function getCard(name: string, url: string) {
  return (
    <Link href={url} key={url} legacyBehavior>
      <Button variant="outlined">{name}</Button>
    </Link>
  )
}

export default TabContent
