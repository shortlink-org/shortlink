// This middleware uses ORY Krato's `/sessions/whoami` endpoint to check if the user is signed in or not:
//
//   import express from 'express'
//   import protect from './middleware/auth.ts'
//
//   const app = express()
//
//   app.get("/dashboard", protect, (req, res) => { /* ... */ })

import { Configuration, PublicApi } from '@ory/kratos-client'
import { NextFunction, Request, Response } from 'express'
import http from 'http'

const kratos = new PublicApi( // eslint-disable-line
  new Configuration({
    basePath: process.env.KRATOS_API || 'http://127.0.0.1:4433',
  }),
)

export default (req: Request, res: Response, next: NextFunction) => {
  const options = {
    host: '127.0.0.1',
    port: '4433',
    path: '/sessions/whoami',
    headers: {
      Cookie: req.header('Cookie'),
    },
  }

  http
    .request(options, (response) => {
      let session = ''

      // another chunk of data has been received, so append it to `str`
      response.on('data', (chunk) => {
        session += chunk
      })

      // the whole response has been received, so we just print it out here
      response.on('end', () => {
        ;(req as Request & { user: any }).user = JSON.parse(session)

        // @ts-ignore
        if (req.user.active) {
          next()
        } else {
          // If no session is found, redirect to login.
          res.redirect('/next/auth/login')
        }
      })
    })
    .end()

  // TODO: use official method
  // // @ts-ignore
  // kratos
  //   .toSession(
  //     req.header('Cookie'),
  //     req.header('Authorization'),
  //     // @ts-ignore
  //   )
  //   .then(({ data: session }) => {
  //     // `whoami` returns the session or an error. We're changing the type here
  //     // because express-session is not detected by TypeScript automatically.
  //     (req as Request & { user: any }).user = { session }
  //     next()
  //   })
  //   .catch((error) => {
  //     console.warn("TEST", error.response)
  //     // If no session is found, redirect to login.
  //     res.redirect('/next/auth/login')
  //   })
}
