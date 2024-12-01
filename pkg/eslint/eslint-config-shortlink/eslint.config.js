const shortlink = require('eslint-config-shortlink')

module.exports = [
  ...shortlink,
  {
    ignores: ['dist/'],
  },
]
