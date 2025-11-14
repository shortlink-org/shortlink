import { z } from "zod";
import { ValidationError } from "../../../application/exceptions/index.js";

/**
 * HTTP DTO для запроса редиректа
 * Используется для валидации HTTP параметров запроса
 */
export const RedirectRequestDtoSchema = z.object({
  hash: z
    .string()
    .min(1, "Hash cannot be empty")
    .regex(/^[a-zA-Z0-9]+$/, "Hash must contain only alphanumeric characters"),
});

export type RedirectRequestDto = z.infer<typeof RedirectRequestDtoSchema>;

/**
 * Валидирует HTTP параметры запроса редиректа
 * @param params - Express request params
 * @returns Валидированный DTO или выбрасывает ValidationError
 */
export function validateRedirectRequest(params: {
  hash?: string;
}): RedirectRequestDto {
  try {
    return RedirectRequestDtoSchema.parse({ hash: params.hash });
  } catch (error) {
    if (error instanceof z.ZodError) {
      const firstError = error.issues[0];
      const field = firstError?.path.join(".") || "hash";
      throw new ValidationError(
        firstError?.message || "Validation failed",
        field,
        { issues: error.issues }
      );
    }
    throw new ValidationError("Invalid request parameters");
  }
}

