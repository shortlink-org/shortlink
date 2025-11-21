import { Link } from "../../domain/entities/Link.js";

/**
 * Response DTO для получения ссылки
 * Application Boundary Contract (ABC): Single output DTO для GetLinkByHashUseCase
 * Application DTO - чистый, без зависимостей от Express, Prisma, HTTP, gRPC
 */
export interface GetLinkResponse {
  link: Link;
}

