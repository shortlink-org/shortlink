import express from 'express'

import protect from './middleware/auth'
import sessionCookie from './middleware/sessionCookie'

const next = require('next')
const cookieParser = require('cookie-parser')

// @ts-ignore
const port = parseInt(process.env.PORT, 10) || 3000
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handler = app.getRequestHandler()

app.prepare().then(() => {
  // app.buildId is only available after app.prepare(), hence why we setup here
  const app = express()

  // add middleware
  app.use(cookieParser())
  app.use(sessionCookie)

  // Routing
  app.all('/next/user/*', protect, handler)
  app.all('/next/admin/*', protect, handler)
  app.all('*', handler)

  // Run server
  app.listen(port, (err?: any) => {
    if (err) {
      throw err
    }
    // eslint-disable-next-line no-console
    console.log(`> Ready on http://localhost:${port}`)
  })
})
