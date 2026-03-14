import { Temporal } from '@js-temporal/polyfill'

type ProtoTimestamp = {
  seconds?: number | string | null
  nanos?: number | null
}

export function protoTimestampToIsoString(timestamp?: ProtoTimestamp | null): string {
  if (!timestamp) return ''

  const epochNanoseconds =
    BigInt(String(timestamp.seconds ?? 0)) * 1_000_000_000n + BigInt(timestamp.nanos ?? 0)

  return Temporal.Instant.fromEpochNanoseconds(epochNanoseconds).toString()
}
