/**
 * Простое кеширование промисов для использования с React 19 use()
 * 
 * Важно: промисы должны быть закешированы, иначе они будут пересоздаваться
 * при каждом рендере и компонент будет бесконечно suspending
 */

type CacheEntry<T> = {
  promise: Promise<T>
  timestamp: number
  ttl: number // Time to live in milliseconds
}

class PromiseCache {
  private cache: Map<string, CacheEntry<any>> = new Map()

  /**
   * Получить промис из кеша или создать новый
   * 
   * @param key - уникальный ключ для кеша
   * @param fetcher - функция для получения данных
   * @param ttl - время жизни кеша в миллисекундах (по умолчанию 5 минут)
   */
  get<T>(key: string, fetcher: () => Promise<T>, ttl: number = 5 * 60 * 1000): Promise<T> {
    const cached = this.cache.get(key)
    const now = Date.now()

    // Проверяем есть ли валидный кеш
    if (cached && (now - cached.timestamp) < cached.ttl) {
      return cached.promise
    }

    // Создаём новый промис
    const promise = fetcher().catch((error) => {
      // Удаляем из кеша в случае ошибки
      this.cache.delete(key)
      throw error
    })

    // Кешируем
    this.cache.set(key, {
      promise,
      timestamp: now,
      ttl
    })

    return promise
  }

  /**
   * Инвалидировать кеш по ключу
   */
  invalidate(key: string) {
    this.cache.delete(key)
  }

  /**
   * Инвалидировать все ключи по паттерну
   */
  invalidatePattern(pattern: RegExp) {
    for (const key of this.cache.keys()) {
      if (pattern.test(key)) {
        this.cache.delete(key)
      }
    }
  }

  /**
   * Очистить весь кеш
   */
  clear() {
    this.cache.clear()
  }

  /**
   * Предзагрузить данные в кеш (prefetch)
   */
  prefetch<T>(key: string, fetcher: () => Promise<T>, ttl?: number) {
    this.get(key, fetcher, ttl)
  }
}

// Глобальный инстанс кеша
export const promiseCache = new PromiseCache()

/**
 * Вспомогательная функция для создания кешируемых fetcher'ов
 */
export function createCachedFetcher<TArgs extends any[], TResult>(
  keyGenerator: (...args: TArgs) => string,
  fetcher: (...args: TArgs) => Promise<TResult>,
  ttl?: number
) {
  return (...args: TArgs): Promise<TResult> => {
    const key = keyGenerator(...args)
    return promiseCache.get(key, () => fetcher(...args), ttl)
  }
}

/**
 * React Hook для инвалидации кеша
 */
export function useInvalidateCache() {
  return {
    invalidate: (key: string) => promiseCache.invalidate(key),
    invalidatePattern: (pattern: RegExp) => promiseCache.invalidatePattern(pattern),
    clear: () => promiseCache.clear(),
  }
}

export default promiseCache

