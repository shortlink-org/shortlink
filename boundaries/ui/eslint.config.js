const shortlink = require('@shortlink-org/eslint-plugin-shortlink')

module.exports = [
  ...(Array.isArray(shortlink) ? shortlink : [shortlink.default || shortlink]),
  {
    ignores: ['node_modules', 'out', '.*', 'public', 'test'],
  },
]
