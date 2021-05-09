import { Container } from 'inversify'
import { interfaces, TYPE } from 'inversify-express-utils'
import { makeLoggerMiddleware } from 'inversify-logger-middleware'

import { LinkService } from './proxy/service/links'
import { StatsService } from './proxy/service/stats'

import TYPES from './types'

// load everything needed to the Container
const container = new Container();

if (process.env.NODE_ENV === 'development') {
    let logger = makeLoggerMiddleware();
    container.applyMiddleware(logger);
}

container.bind<LinkService>(TYPES.SERVICE.LinkService).to(LinkService).inSingletonScope()
container.bind<StatsService>(TYPES.SERVICE.StatsService).to(StatsService).inSingletonScope()
export default container
