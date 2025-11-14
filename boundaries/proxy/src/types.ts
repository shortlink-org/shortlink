const REPOSITORY = {
  LinkRepository: Symbol.for("LinkRepository"),
};

const DOMAIN = {
  LinkDomainService: Symbol.for("LinkDomainService"),
};

const APPLICATION = {
  GetLinkByHashUseCase: Symbol.for("GetLinkByHashUseCase"),
  PublishEventUseCase: Symbol.for("PublishEventUseCase"),
  LinkApplicationService: Symbol.for("LinkApplicationService"),
  UseCasePipeline: Symbol.for("UseCasePipeline"),
  LoggingInterceptor: Symbol.for("LoggingInterceptor"),
  MetricsInterceptor: Symbol.for("MetricsInterceptor"),
  AuthorizationInterceptor: Symbol.for("AuthorizationInterceptor"),
};

const INFRASTRUCTURE = {
  EventPublisher: Symbol.for("EventPublisher"),
  MessageBus: Symbol.for("MessageBus"),
  Logger: Symbol.for("Logger"),
  LinkServiceAdapter: Symbol.for("LinkServiceAdapter"),
  GrpcMetrics: Symbol.for("GrpcMetrics"),
  GrpcTracing: Symbol.for("GrpcTracing"),
  AuthorizationChecker: Symbol.for("AuthorizationChecker"),
  LinkCache: Symbol.for("LinkCache"),
};

const TAGS = {
  StatsController: Symbol.for("StatsController"),
};

export default {
  REPOSITORY,
  DOMAIN,
  APPLICATION,
  INFRASTRUCTURE,
  TAGS,
};
