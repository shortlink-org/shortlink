import { inject, injectable } from "inversify";
import { Hash } from "../../domain/entities/Hash.js";
import { LinkNotFoundError } from "../../domain/exceptions/index.js";
import { ILinkRepository } from "../../domain/repositories/ILinkRepository.js";
import { GetLinkRequest } from "../dto/GetLinkRequest.js";
import { GetLinkResponse } from "../dto/GetLinkResponse.js";
import TYPES from "../../../types.js";

/**
 * Use Case для получения ссылки по хешу
 */
@injectable()
export class GetLinkByHashUseCase {
  constructor(
    @inject(TYPES.REPOSITORY.LinkRepository)
    private readonly linkRepo: ILinkRepository
  ) {}

  async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
    const hash = new Hash(request.hash);
    const link = await this.linkRepo.findByHash(hash);

    if (!link) {
      throw new LinkNotFoundError(hash);
    }

    return new GetLinkResponse(link);
  }
}

