const SERVICE = {
  StatsService: Symbol.for('StatsService'),
  LinkService: Symbol.for('LinkService'),
}

const REPOSITORY = {
  StatsRepositoryImpl: Symbol('StatsRepositoryImpl'),
}

const TAGS = {
  StatsController: Symbol('StatsController'),
}

export default { SERVICE, REPOSITORY, TAGS }
