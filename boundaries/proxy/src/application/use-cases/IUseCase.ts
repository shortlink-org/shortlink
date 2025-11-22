/**
 * Базовый интерфейс для Use Cases
 * Use Case представляет одну бизнес-операцию в Application Layer
 */
export interface IUseCase<TRequest, TResponse> {
  /**
   * Выполняет Use Case с переданным запросом
   *
   * @param request - входные данные для Use Case
   * @returns результат выполнения Use Case
   */
  execute(request: TRequest): Promise<TResponse>;
}

