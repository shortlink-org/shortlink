import * as express from "express";
import { inject } from "inversify";
import {
  controller,
  httpGet,
  request,
  response,
} from "inversify-express-utils";
import { LinkApplicationService } from "../../../application/services/LinkApplicationService.js";
import { Hash } from "../../../domain/entities/Hash.js";
import { validateRedirectRequest } from "../dto/RedirectRequestDto.js";
import { ILogger } from "../../../../infrastructure/logging/ILogger.js";
import TYPES from "../../../../types.js";

/**
 * HTTP Controller для обработки редиректов
 * Использует Application Service для оркестрации Use Cases
 * Преобразует HTTP запросы в Application DTO и обратно
 */
@controller("/s/:hash")
export class ProxyController {
  constructor(
    @inject(TYPES.APPLICATION.LinkApplicationService)
    private readonly linkApplicationService: LinkApplicationService,
    @inject(TYPES.INFRASTRUCTURE.Logger)
    private readonly logger: ILogger
  ) {}

  /**
   * @swagger
   * /s/{hash}:
   *   get:
   *     summary: Редирект на оригинальный URL по короткой ссылке
   *     description: |
   *       Выполняет редирект (301) на оригинальный URL по хешу короткой ссылки.
   *       Статистика редиректов собирается автоматически через eBPF.
   *     tags:
   *       - Redirect
   *     parameters:
   *       - in: path
   *         name: hash
   *         required: true
   *         schema:
   *           type: string
   *           pattern: '^[a-zA-Z0-9]+$'
   *           minLength: 1
   *           maxLength: 100
   *         description: Хеш короткой ссылки
   *         example: abc123
   *     responses:
   *       '301':
   *         description: Успешный редирект на оригинальный URL
   *         headers:
   *           Location:
   *             description: URL для редиректа
   *             schema:
   *               type: string
   *               example: https://example.com
   *       '400':
   *         $ref: '#/components/responses/BadRequest'
   *         description: Невалидный формат хеша
   *       '404':
   *         $ref: '#/components/responses/NotFound'
   *         description: Ссылка с указанным хешем не найдена
   *       '429':
   *         $ref: '#/components/responses/TooManyRequests'
   *         description: Превышен лимит запросов (100 запросов за 15 минут)
   *       '500':
   *         $ref: '#/components/responses/InternalServerError'
   *         description: Внутренняя ошибка сервера
   */
  @httpGet("/")
  public async redirect(
    @request() req: express.Request,
    @response() res: express.Response,
    next: express.NextFunction
  ): Promise<void> {
    try {
      // Валидация HTTP параметров через DTO
      const dto = validateRedirectRequest(req.params);

      // Преобразуем валидированный DTO в доменный Value Object
      const hash = new Hash(dto.hash);

      // Используем Application Service для обработки редиректа
      const result = await this.linkApplicationService.handleRedirect({ hash });

      if (result.isErr()) {
        // Пробрасываем ошибку в error handler middleware
        return next(result.error);
      }

      const { link } = result.value;

      // Выполняем редирект
      res.redirect(301, link.url);
    } catch (error) {
      // Пробрасываем все ошибки в error handler middleware
      next(error);
    }
  }
}

