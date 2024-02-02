import { Container } from 'inversify'

import { LinkService } from './proxy/service/links'
import { StatsService } from './proxy/service/stats'

import TYPES from './types'
import AMQPController from "./proxy/infrastructure/amqp/amqp";
import StatsRepository from "./proxy/infrastructure/store";

// load everything needed to the Container
const container = new Container();

container.bind<LinkService>(TYPES.SERVICE.LinkService).to(LinkService).inSingletonScope()
container.bind<StatsService>(TYPES.SERVICE.StatsService).to(StatsService).inSingletonScope()
container.bind<AMQPController>(TYPES.TAGS.AMQPController).to(AMQPController).inSingletonScope()
container.bind<StatsRepository>(TYPES.REPOSITORY.StatsRepository).to(StatsRepository).inSingletonScope()
export default container
