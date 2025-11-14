/**
 * Request DTO для получения ссылки по хешу
 * Application DTO - чистый, без зависимостей от Express, Prisma, HTTP, gRPC
 */
export class GetLinkRequest {
  constructor(public readonly hash: string) {}
}

