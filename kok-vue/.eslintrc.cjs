/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  'extends': [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier/skip-formatting',
    '@vue/airbnb'
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    parser: 'babel-eslint'
  },
  rules: {
    'indent' : ['warn',2],
    'no-console' : process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger' : process.env.NODE_ENV === 'production' ? 'error' : 'off'
  }
}