if (process.env.DD_PROFILING_ENABLED === 'true') {
  try {
    const Pyroscope = require('@pyroscope/nodejs')

    const SERVICE_NAME = process.env.SERVICE_NAME || 'proxy'
    const PYROSCOPE_SERVER_ADDRESS = process.env.PYROSCOPE_SERVER_ADDRESS || 'http://pyroscope:4040'

    Pyroscope.init({
      appName: SERVICE_NAME,
      serverAddress: PYROSCOPE_SERVER_ADDRESS,
      sourceMapPath: ['.'],
    })

    Pyroscope.start()
    console.log(`[Pyroscope] enabled for "${SERVICE_NAME}"`)
  } catch (err) {
    console.error('[Pyroscope] failed to initialize:', err)
  }
}
