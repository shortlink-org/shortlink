import { CreateLinkUseCase } from './createLinkUseCase'
import { LinkHttpGateway } from '@/infrastructure/http/link/linkHttpGateway'

const linkGateway = new LinkHttpGateway()

export const createLinkUseCase = new CreateLinkUseCase(linkGateway)
