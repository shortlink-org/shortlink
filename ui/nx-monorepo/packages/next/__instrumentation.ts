import { registerOTel } from '@vercel/otel'

export function register() {
  registerOTel('ui-next')
}
