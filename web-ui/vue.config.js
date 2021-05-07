const fs = require('fs')

module.exports = {
  lintOnSave: false,
  assetsDir: 'static',
  devServer: {
    proxy: 'http://localhost:5000',
    https: {
      key: fs.readFileSync('./localhost-key.pem'),
      cert: fs.readFileSync('./localhost.pem'),
    },
  }
}
