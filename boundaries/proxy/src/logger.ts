import { createLogger, format, transports } from 'winston'

const log = createLogger({
  level: process.env.LOG_LEVEL || 'info',
  transports: [new transports.Console({ handleExceptions: true })],
  format: format.combine(
    format.errors({ stack: true }),
    format.splat(),
    format.timestamp(),
    format.json(),
  ),
})

export default log
