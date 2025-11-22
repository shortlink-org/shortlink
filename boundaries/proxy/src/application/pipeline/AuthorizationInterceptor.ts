import { IUseCaseInterceptor, UseCaseExecutionContext } from "./IUseCaseInterceptor.js";
import { ILogger } from "../../../infrastructure/logging/ILogger.js";
import { ApplicationError } from "../exceptions/ApplicationError.js";

/**
 * Ошибка авторизации
 */
export class AuthorizationError extends ApplicationError {
  constructor(message: string = "Unauthorized") {
    super(message, "AUTHORIZATION_ERROR", 403);
    this.name = "AuthorizationError";
    Object.setPrototypeOf(this, AuthorizationError.prototype);
  }
}

/**
 * Интерфейс для проверки авторизации
 */
export interface IAuthorizationChecker {
  /**
   * Проверяет, имеет ли пользователь право выполнить Use Case
   *
   * @param useCaseName - Название Use Case
   * @param request - Запрос Use Case
   * @param context - Контекст выполнения (может содержать информацию о пользователе)
   * @returns true, если авторизован, false иначе
   */
  isAuthorized(
    useCaseName: string,
    request: any,
    context: UseCaseExecutionContext<any, any>
  ): Promise<boolean> | boolean;
}

/**
 * Интерцептор для авторизации Use Cases
 * Проверяет права доступа перед выполнением Use Case
 */
export class AuthorizationInterceptor<TRequest = any, TResponse = any>
  implements IUseCaseInterceptor<TRequest, TResponse>
{
  constructor(
    private readonly logger: ILogger,
    private readonly authorizationChecker: IAuthorizationChecker
  ) {}

  async before(
    context: UseCaseExecutionContext<TRequest, TResponse>
  ): Promise<TRequest> {
    // Проверяем авторизацию
    const isAuthorized = await this.authorizationChecker.isAuthorized(
      context.useCaseName,
      context.request,
      context
    );

    if (!isAuthorized) {
      this.logger.warn(`[Authorization] Access denied for: ${context.useCaseName}`, {
        useCase: context.useCaseName,
      });
      throw new AuthorizationError(
        `Access denied for Use Case: ${context.useCaseName}`
      );
    }

    return context.request;
  }

  after(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    // Авторизация проверена, дополнительных действий не требуется
  }

  onError(context: UseCaseExecutionContext<TRequest, TResponse>): void {
    // Логируем ошибки авторизации
    if (context.error instanceof AuthorizationError) {
      this.logger.warn(`[Authorization] ${context.error.message}`, {
        useCase: context.useCaseName,
      });
    }
  }
}

/**
 * Простая реализация IAuthorizationChecker для демонстрации
 * В реальном приложении должна проверять права пользователя
 */
export class DefaultAuthorizationChecker implements IAuthorizationChecker {
  /**
   * Use Cases, которые не требуют авторизации (публичные)
   */
  private readonly publicUseCases = ["GetLinkByHashUseCase"];

  isAuthorized(
    useCaseName: string,
    request: any,
    context: UseCaseExecutionContext<any, any>
  ): boolean {
    // Публичные Use Cases доступны всем
    if (this.publicUseCases.includes(useCaseName)) {
      return true;
    }

    // Для остальных Use Cases проверяем наличие пользователя в контексте
    // В реальном приложении здесь должна быть проверка токена, ролей и т.д.
    const user = context.metadata.get("user");
    return user !== undefined && user !== null;
  }
}

