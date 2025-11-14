import { DomainError } from "./DomainError.js";

/**
 * Исключение, выбрасываемое при попытке создать Hash с невалидным значением
 */
export class InvalidHashError extends DomainError {
  constructor(hash: string, reason?: string) {
    const message = reason
      ? `Invalid hash "${hash}": ${reason}`
      : `Invalid hash format: "${hash}"`;
    super(message);
  }
}

