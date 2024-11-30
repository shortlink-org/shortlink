import "reflect-metadata"
import http from "http"
import * as express from "express"
import * as bodyParser from 'body-parser'
import helmet from 'helmet'
import morgan from 'morgan'
import {InversifyExpressServer} from "inversify-express-utils"
import {createTerminus} from "@godaddy/terminus"
import * as dotenv from 'dotenv'

dotenv.config()

import './opentelemtery'
import './pprof'
import log from './logger'
// @ts-ignore
import container from "./inversify.config"

import './proxy/infrastructure/http/proxy'

const APP = express.default()
const PORT = process.env.PORT || 3020
const SERVER_HTTP = http.createServer(APP)

const morganMiddleware = morgan(
  ':method :url :status :res[content-length] - :response-time ms',
  {
    stream: {
      // Configure Morgan to use our custom logger with the http severity
      write: (message: string) => log.http(message.trim()),
    },
  }
);

// configuration application
APP.use(bodyParser.urlencoded({
  extended: true
}))
APP.use(bodyParser.json())
APP.use(helmet())
APP.use(morganMiddleware)

// start the server
const SERVER = new InversifyExpressServer(container, null, { rootPath: "/" }, APP);
const APP_CONFIGURED = SERVER.build();

async function onHealthCheck () {
  // checks if the system is healthy, like the db connection is live
  // resolves, if health, rejects if not
  return true
}

// graceful shutdown and Kubernetes readiness/liveness checks
createTerminus(SERVER_HTTP, {
  signal: 'SIGINT',
  healthChecks: { '/ready': onHealthCheck },
})

SERVER_HTTP.listen(PORT, () => log.info(`App running on ${PORT}`));

exports = module.exports = SERVER_HTTP
