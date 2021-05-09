import "reflect-metadata"
import * as express from "express"
import * as bodyParser from 'body-parser'
import helmet from 'helmet'
import { Logger } from "tslog"
import { InversifyExpressServer } from "inversify-express-utils"

import container from "./inversify.config"

import './proxy/infrastructure/http/proxy'

const APP = express.default()
const PORT = process.env.PORT || 3000
const log: Logger = new Logger({ type: "json" })

// configuration application
APP.use(bodyParser.urlencoded({
  extended: true
}))
APP.use(bodyParser.json())
APP.use(helmet())

// start the server
const SERVER = new InversifyExpressServer(container, null, { rootPath: "/" }, APP);
const APP_CONFIGURED = SERVER.build();
APP_CONFIGURED.listen(PORT, () => log.info(`App running on ${PORT}`));

exports = module.exports = APP_CONFIGURED
