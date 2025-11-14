import * as express from "express";
import { ErrorMapper } from "./ErrorMapper.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { WinstonLogger } from "../../../infrastructure/logging/WinstonLogger.js";

/**
 * Express middleware для централизованной обработки ошибок
 * Должен быть последним middleware в цепочке
 */
export function errorHandler(
  logger: ILogger = new WinstonLogger()
): express.ErrorRequestHandler {
  const errorMapper = new ErrorMapper(logger);

  return (
    err: unknown,
    req: express.Request,
    res: express.Response,
    next: express.NextFunction
  ) => {
    // Если ответ уже отправлен, передаем ошибку дальше
    if (res.headersSent) {
      return next(err);
    }

    // Маппим ошибку в HTTP ответ
    errorMapper.mapToHttpResponse(err, req, res);
  };
}

