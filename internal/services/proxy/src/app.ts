import "reflect-metadata"
import http from "http"
import * as express from "express"
import * as bodyParser from 'body-parser'
import helmet from 'helmet'
import { Logger } from "tslog"
import { InversifyExpressServer } from "inversify-express-utils"
import { createTerminus } from "@godaddy/terminus"

import container from "./inversify.config"

import './proxy/infrastructure/http/proxy'

const APP = express.default()
const PORT = process.env.PORT || 3020
const log: Logger = new Logger()
const SERVER_HTTP = http.createServer(APP)

// configuration application
APP.use(bodyParser.urlencoded({
  extended: true
}))
APP.use(bodyParser.json())
APP.use(helmet())

// start the server
const SERVER = new InversifyExpressServer(container, null, { rootPath: "/" }, APP);
const APP_CONFIGURED = SERVER.build();

async function onHealthCheck () {
  // checks if the system is healthy, like the db connection is live
  // resolves, if health, rejects if not
}

// graceful shutdown and Kubernetes readiness/liveness checks
createTerminus(SERVER_HTTP, {
  signal: 'SIGINT',
  healthChecks: { '/ready': onHealthCheck },
})

SERVER_HTTP.listen(PORT, () => log.info(`App running on ${PORT}`));

exports = module.exports = SERVER_HTTP
