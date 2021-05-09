import * as express from 'express'
import { inject } from 'inversify'
import { interfaces, controller, httpGet, httpPost, request, response } from "inversify-express-utils"
import { Logger } from "tslog"

import TYPES from '../../../types'

const log: Logger = new Logger({ type: "json" })

@controller(`/s/:hash`)
class ProxyController implements interfaces.Controller {
  // constructor( @inject(TYPES.SERVICE.StatsService) private statsService: ProxyController) { }

  @httpGet(`/`)
  public async redirect (@request() req: express.Request, @response() res: express.Response) {
    const { hash } = req.params

    try {
      // TODO:
      // 1. request by gRPC to LinkService: getLinkByHash
      const link = `https://batazor.ru`
      res.redirect(301, link)
    } catch (error) {
      res.status(400).json(error)
    }
  }
}

export default ProxyController
