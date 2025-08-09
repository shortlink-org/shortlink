const SERVICE = {
  StatsService: Symbol.for('StatsService'),
  LinkService: Symbol.for('LinkService'),
}

const REPOSITORY = {
  StatsRepository: Symbol.for('StatsRepository'),
}

const TAGS = {
  StatsController: Symbol.for('StatsController'),
  AMQPController: Symbol.for('AMQPController'),
}

export default { SERVICE, REPOSITORY, TAGS }
