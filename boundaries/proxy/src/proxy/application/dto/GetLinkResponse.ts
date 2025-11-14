import { Link } from "../../domain/entities/Link.js";

/**
 * Response DTO для получения ссылки
 * Application DTO - чистый, без зависимостей от Express, Prisma, HTTP, gRPC
 */
export class GetLinkResponse {
  constructor(public readonly link: Link) {}
}

