const { addDecoratorsLegacy, addLessLoader } = require('customize-cra')
// const path = require('path')

function webpack(config, env) {
  config = addDecoratorsLegacy()(config)
  config = addLessLoader({
    javascriptEnabled: true,
    modifyVars: {}
  })(config)
  // config = overrideEsLint(config)
  config.module.rules.push({
    test: /\.mjs$/,
    include: /node_modules/,
    type: 'javascript/auto'
  })
  config.resolve = {
    alias: {
      // '@components': path.resolve(__dirname, './src/components'),
    },
    extensions: ['.js', 'jsx', '.json', '.ts', '.tsx']
  }
  return config
}

module.exports = { webpack }
