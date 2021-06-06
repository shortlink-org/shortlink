import express, { Request, Response } from "express"
const next = require('next')

const cookieParser = require('cookie-parser')
const { v4: uuidv4 } = require('uuid')
// @ts-ignore
const port = parseInt(process.env.PORT, 10) || 3000
const dev = process.env.NODE_ENV !== 'production'
const app = next({ dev })
const handler = app.getRequestHandler()

// @ts-ignore
function sessionCookie(req, res, next) {
  const htmlPage =
    !req.path.match(/^\/(_next|static)/) &&
    !req.path.match(/\.(js|map)$/) &&
    req.accepts('text/html', 'text/css', 'image/png') === 'text/html'

  if (!htmlPage) {
    next()
    return
  }

  if (!req.cookies.sid || req.cookies.sid.length === 0) {
    req.cookies.sid = uuidv4()
    res.cookie('sid', req.cookies.sid)
  }

  next()
}

app.prepare().then(() => {
  // app.buildId is only available after app.prepare(), hence why we setup here

  express()
    .use(cookieParser())
    .use(sessionCookie)
    // Regular next.js request handler
    .use(handler)
    .listen(port, (err?: any) => {
      if (err) {
        throw err
      }
      // eslint-disable-next-line no-console
      console.log(`> Ready on http://localhost:${port}`)
    })
})
