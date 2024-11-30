const Pyroscope = require('@pyroscope/nodejs')

const SERVICE_NAME: string = process.env.SERVICE_NAME || 'proxy'
const PYROSCOPE_SERVER_ADDRESS: string = process.env.PYROSCOPE_SERVER_ADDRESS || 'http://pyroscope:4040'

Pyroscope.init({
  appName: SERVICE_NAME,
  serverAddress: PYROSCOPE_SERVER_ADDRESS,
  sourceMapPath: ['.'],
})

Pyroscope.start()
