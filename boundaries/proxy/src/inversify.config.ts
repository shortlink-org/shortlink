import { Container } from 'inversify'

import TYPES from './types'

// link service
import { LinkService } from './proxy/service/links'

// stats service
// import { StatsService } from './proxy/service/stats'
// import StatsRepository from "./proxy/infrastructure/store";

import AMQPController from "./proxy/infrastructure/amqp/amqp";

// load everything needed to the Container
const container = new Container();

container.bind<LinkService>(TYPES.SERVICE.LinkService).to(LinkService).inSingletonScope()
// container.bind<StatsService>(TYPES.SERVICE.StatsService).to(StatsService).inSingletonScope()
container.bind<AMQPController>(TYPES.TAGS.AMQPController).to(AMQPController).inSingletonScope()
// container.bind<StatsRepository>(TYPES.REPOSITORY.StatsRepository).to(StatsRepository).inSingletonScope()
export default container
