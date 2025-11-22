/**
 * Request DTO для получения ссылки по хешу
 * Application Boundary Contract (ABC): Single input DTO для GetLinkByHashUseCase
 * Application DTO - чистый, без зависимостей от Express, Prisma, HTTP, gRPC
 */
export interface GetLinkRequest {
  hash: string;
}

