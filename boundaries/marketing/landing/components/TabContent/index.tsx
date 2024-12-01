import { Divider, Stack, Typography, useTheme } from '@mui/material'
import Button from '@mui/material/Button'
import Link from 'next/link'
import React from 'react'

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
    <div className="my-2 mx-5 max-w-4xl mx-auto">
      <h2 className="text-2xl prose text-center my-5 text-gray-800 dark:text-white">{title}</h2>

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
        {cards.map((card) => (
          <Link href={card.url} key={card.url} passHref>
            <Button
              variant="outlined"
              size="large"
              sx={{
                minWidth: 160,
                color: theme.palette.text.primary,
                borderColor: theme.palette.divider,
                '&:hover': {
                  borderColor: theme.palette.primary.main,
                  backgroundColor: theme.palette.action.hover,
                },
              }}
            >
              {card.name}
            </Button>
          </Link>
        ))}
      </Stack>
    </div>
  )
}

export default TabContent
