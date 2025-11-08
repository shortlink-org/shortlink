import { LinkGateway, CreateLinkCommand, CreateLinkResult } from '@/domain/link/link.types'

export class CreateLinkUseCase {
  constructor(private readonly gateway: LinkGateway) {}

  async execute(command: CreateLinkCommand): Promise<CreateLinkResult> {
    return this.gateway.createLink(command)
  }
}
