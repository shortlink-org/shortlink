import { IUseCase } from "../use-cases/IUseCase.js";
import { Result } from "neverthrow";
import {
  IUseCaseInterceptor,
  UseCaseExecutionContext,
} from "./IUseCaseInterceptor.js";

/**
 * Pipeline для выполнения Use Case с интерцепторами
 * Реализует паттерн Chain of Responsibility для cross-cutting concerns
 */
export class UseCasePipeline {
  /**
   * Выполняет Use Case с цепочкой интерцепторов
   *
   * @param useCase - Use Case для выполнения
   * @param request - Входные данные запроса
   * @param interceptors - Массив интерцепторов для выполнения
   * @returns Promise с результатом выполнения Use Case
   */
  async execute<TRequest, TResponse>(
    useCase: IUseCase<TRequest, TResponse>,
    request: TRequest,
    interceptors: IUseCaseInterceptor<TRequest, TResponse>[] = []
  ): Promise<TResponse> {
    const useCaseName = useCase.constructor.name;
    const startTime = Date.now();
    const metadata = new Map<string, unknown>();

    // Создаем контекст выполнения
    const context: UseCaseExecutionContext<TRequest, TResponse> = {
      useCaseName,
      request,
      startTime,
      metadata,
    };

    let modifiedRequest = request;

    try {
      // Выполняем before интерцепторы
      for (const interceptor of interceptors) {
        modifiedRequest = await interceptor.before(context);
        context.request = modifiedRequest;
      }

      // Выполняем Use Case
      const result = await useCase.execute(modifiedRequest);

      // Заполняем контекст результатом
      context.response = result as any;
      context.endTime = Date.now();
      context.duration = context.endTime - context.startTime;

      // Проверяем, является ли результат Result типом (neverthrow)
      const isResultType = result && typeof result === "object" && "isErr" in result;

      if (isResultType && (result as any).isErr()) {
        const errorResult = result as any;
        context.error = errorResult.error as Error;

        // Выполняем onError интерцепторы
        for (const interceptor of interceptors) {
          await interceptor.onError(context);
        }

        // Выполняем finally интерцепторы
        for (const interceptor of interceptors) {
          if (interceptor.finally) {
            await interceptor.finally(context);
          }
        }

        // Пробрасываем ошибку дальше
        throw errorResult.error;
      }

      // Выполняем after интерцепторы для успешного результата
      for (const interceptor of interceptors) {
        await interceptor.after(context);
      }

      // Выполняем finally интерцепторы
      for (const interceptor of interceptors) {
        if (interceptor.finally) {
          await interceptor.finally(context);
        }
      }

      // Возвращаем успешный результат
      // Если это Result тип, возвращаем value, иначе сам результат
      if (isResultType) {
        return (result as any).value;
      }
      return result;
    } catch (error) {
      // Обработка неожиданных ошибок
      context.error = error instanceof Error ? error : new Error(String(error));
      context.endTime = Date.now();
      context.duration = context.endTime - context.startTime;

      // Выполняем onError интерцепторы
      for (const interceptor of interceptors) {
        await interceptor.onError(context);
      }

      // Выполняем finally интерцепторы
      for (const interceptor of interceptors) {
        if (interceptor.finally) {
          await interceptor.finally(context);
        }
      }

      throw error;
    }
  }
}

