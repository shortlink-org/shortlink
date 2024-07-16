'use client'

// @ts-nocheck

import FileCopyIcon from '@mui/icons-material/FileCopy'
import Alert from '@mui/material/Alert'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Grid from '@mui/material/Grid'
import IconButton from '@mui/material/IconButton'
import Link from '@mui/material/Link'
import Snackbar from '@mui/material/Snackbar'
import TextField from '@mui/material/TextField'
import Typography from '@mui/material/Typography'
import React, { useState } from 'react'
import { CopyToClipboard } from 'react-copy-to-clipboard'

import withAuthSync from 'components/Private'
import Header from 'components/Page/Header'

function Page() {
  const [open, setOpen] = useState(false)

  const [url, setURL] = useState({
    url: '',
  })

  const [response, setResponse] = useState({
    type: '',
    message: '',
    hash: '',
  })

  const handleChange = (e: { target: { name: any; value: any } }) => setURL({ ...url, [e.target.name]: e.target.value })

  const handleClose = (event: any, reason: string) => {
    if (reason === 'clickaway') {
      return
    }

    setOpen(false)
  }

  const handleSubmit = async (e: { preventDefault: () => void }) => {
    e.preventDefault()
    try {
      // TODO: use store.actions
      const res = await fetch(`/api/links`, {
        method: 'POST',
        body: JSON.stringify(url),
        headers: { 'Content-Type': 'application/json' },
      })
      const json = await res.json()

      if (res.status === 201) {
        setResponse({
          type: 'success',
          message: 'Success add your link.',
          hash: json.hash,
        })
      } else {
        setResponse({
          type: 'error',
          message: json.error,
          hash: '',
        })
      }

      setOpen(true)
    } catch (error) {
      console.error('An error occurred', error)
      setResponse({
        hash: '',
        type: 'error',
        message: 'An error occured while submitting the form',
      })
      setOpen(true)
    }
  }

  return (
    <>
      {/*<NextSeo title="Add link" description="Add a new link" />*/}
      <Grid container direction="column" justifyContent="space-around" alignItems="center">
        <div className="container mx-auto w-5/6 sm:w-2/3 h-full">
          <Header title="Add link" />

          {/* Code block starts */}
          <div className="w-full bg-white dark:bg-gray-800 py-5 flex flex-col xl:flex-row items-start xl:items-center justify-between px-5 xl:px-10 shadow rounded-t">
            <div className="mb-4 sm:mb-0 md:mb-0 lg:mb-0 xl:mb-0 lg:w-1/2">
              <h2 className="text-gray-800 dark:text-gray-100 text-lg font-bold">Add link</h2>
              <p className="font-normal text-sm text-gray-600 dark:text-gray-100 mt-1">
                This is a great place to add a short, friendly link that goes somewhere interesting.
              </p>
            </div>

            <Box
              component="form"
              sx={{
                '& > :not(style)': { m: 1, width: '25ch' },
                maxWidth: '100%',
              }}
              noValidate
              autoComplete="off"
              onSubmit={handleSubmit}
            >
              <TextField variant="outlined" label="Your URL" name="url" required fullWidth onChange={handleChange} />

              <TextField variant="outlined" label="Describe" name="describe" fullWidth onChange={handleChange} />

              <Button variant="contained" className="bg-sky-600 hover:bg-sky-700" type="submit">
                Add
              </Button>
            </Box>

            {response.type !== '' && response.type !== 'error' && (
              <div className="mb-4 sm:mb-0 md:mb-0 lg:mb-0 xl:mb-0 lg:w-1/2">
                <Typography component="p" className="text-gray-800 dark:text-gray-100 text-lg font-bold">
                  Your link: &nbsp;
                  <Link href={`/s/${response.hash}`} target="_blank" rel="noopener" variant="body2" underline="hover">
                    {window.location.host}/s/{response.hash}
                  </Link>
                  <CopyToClipboard
                    text={`${window.location.host}/s/${response.hash}`}
                    onCopy={() => {
                      setResponse({
                        type: 'success',
                        message: 'Success copy your link.',
                        hash: response.hash,
                      })
                    }}
                  >
                    <IconButton aria-label="copy" color="secondary" size="large">
                      <FileCopyIcon />
                    </IconButton>
                  </CopyToClipboard>
                </Typography>
              </div>
            )}
          </div>
        </div>

        <Snackbar
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left',
          }}
          open={open}
          autoHideDuration={6000}
          onClose={handleClose}
        >
          {/* @ts-ignore */}
          <Alert onClose={() => handleClose} severity={response.type}>
            {response.message}
          </Alert>
        </Snackbar>
      </Grid>
    </>
  )
}

export default withAuthSync(Page)
