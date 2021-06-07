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

const kratos = new PublicApi(new Configuration({ basePath: process.env.KRATOS_API || 'http://127.0.0.1:4433/.ory/kratos/public' }))

export default (req: Request, res: Response, next: NextFunction) => {
  // @ts-ignore
  kratos.toSession(
    req.header('Cookie'), req.header('Authorization')
  // @ts-ignore
  ).then(({ data: session }) => {
    // `whoami` returns the session or an error. We're changing the type here
    // because express-session is not detected by TypeScript automatically.
    (req as Request & { user: any }).user = { session };
    next();
  }).catch(() => {
    // If no session is found, redirect to login.
    res.redirect('/next/auth/login');
  });
};
