import { DomainError } from "./DomainError.js";
import { Hash } from "../entities/Hash.js";

/**
 * Исключение, выбрасываемое когда ссылка не найдена
 */
export class LinkNotFoundError extends DomainError {
  constructor(hash: Hash) {
    super(`Link with hash "${hash.value}" not found`);
  }
}

