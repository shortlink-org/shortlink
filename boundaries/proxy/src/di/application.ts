import { DefaultAuthorizationChecker } from "../../application/pipeline/index.js";
import { LoggingInterceptor } from "../../application/pipeline/index.js";
import { MetricsInterceptor } from "../../application/pipeline/index.js";
import { AuthorizationInterceptor } from "../../application/pipeline/index.js";
import { UseCasePipeline } from "../../application/pipeline/index.js";
import { GetLinkByHashUseCase } from "../../application/use-cases/GetLinkByHashUseCase.js";
import { PublishEventUseCase } from "../../application/use-cases/PublishEventUseCase.js";
import { LinkApplicationService } from "../../application/services/LinkApplicationService.js";

export const APP = {
  authorizationChecker: DefaultAuthorizationChecker,
  loggingInterceptor: LoggingInterceptor,
  metricsInterceptor: MetricsInterceptor,
  authorizationInterceptor: AuthorizationInterceptor,
  useCasePipeline: UseCasePipeline,
  getLinkByHashUseCase: GetLinkByHashUseCase,
  publishEventUseCase: PublishEventUseCase,
  linkApplicationService: LinkApplicationService,
};

