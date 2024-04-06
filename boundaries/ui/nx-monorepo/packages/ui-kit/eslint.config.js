const shortlink = require('eslint-config-shortlink')

module.exports = [
  ...shortlink,
  {
    ignores: ['.*', 'node_modules', 'dist', 'storybook-static'],
  },
]
