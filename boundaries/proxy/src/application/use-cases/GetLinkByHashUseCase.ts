import { Hash } from "../../domain/entities/Hash.js";
import { LinkNotFoundError } from "../../domain/exceptions/index.js";
import { ILinkRepository } from "../../domain/repositories/ILinkRepository.js";
import { GetLinkRequest } from "../dto/GetLinkRequest.js";
import { GetLinkResponse } from "../dto/GetLinkResponse.js";

/**
 * Use Case для получения ссылки по хешу
 */
export class GetLinkByHashUseCase {
  constructor(
    private readonly linkRepository: ILinkRepository
  ) {}

  async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
    const hash = new Hash(request.hash);
    const link = await this.linkRepository.findByHash(hash);

    if (!link) {
      throw new LinkNotFoundError(hash);
    }

    return { link };
  }
}

