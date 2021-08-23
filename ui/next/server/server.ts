import express from 'express'
import next from 'next'
// @ts-ignore
import cookieParser from 'cookie-parser'

import protect from './middleware/auth'

// @ts-ignore
const port = parseInt(process.env.PORT, 10) || 3000
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handler = app.getRequestHandler()

app.prepare().then(() => {
  // app.buildId is only available after app.prepare(), hence why we setup here
  const server = express()

  // add middleware
  server.use(cookieParser())

  // Routing
  // @ts-ignore
  server.all('/next/user/*', protect, handler)
  // @ts-ignore
  server.all('/next/admin/*', protect, handler)
  // @ts-ignore
  server.all('*', handler)

  // Run server
  server.listen(port, (err?: any) => {
    if (err) {
      throw err
    }
    // eslint-disable-next-line no-console
    console.log(`> Ready on http://localhost:${port}`)
  })
})
