/**
 * @js-temporal/polyfill does not attach `Temporal` to globalThis (see
 * https://github.com/js-temporal/temporal-polyfill ). The ui-kit Footer uses
 * `Temporal.Now.plainDateISO()` as a bare global, so we install it once.
 */
import { Temporal } from '@js-temporal/polyfill'

/* Global typing may not match the polyfill’s export; ui-kit only needs `Temporal.Now`. */
if ((globalThis as Record<string, unknown>).Temporal === undefined) {
  ;(globalThis as Record<string, unknown>).Temporal = Temporal
}
