module.exports = {
  "customSyntax": "postcss-scss",
  "extends": "stylelint-config-standard",
  "plugins": ["stylelint-scss"],
  "rules": {
    "at-rule-no-unknown": null,
    "scss/at-rule-no-unknown": true,
    "no-descending-specificity": null,
    "max-line-length": 122,
    "color-function-notation": "legacy",
    "value-no-vendor-prefix": false,
    "alpha-value-notation": "number"
  },
}
