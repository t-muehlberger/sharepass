const fs = require('fs')

function getHttpsConfig() {
  try {
    return {
      key: fs.readFileSync('./localhost-key.pem'),
      cert: fs.readFileSync('./localhost.pem'),
    }
  } catch {
    console.log("failed to load dev keys")
    return undefined;
  }
}

module.exports = {
  lintOnSave: false,
  assetsDir: 'static',
  devServer: {
    proxy: 'http://localhost:5000',
    https: getHttpsConfig(),
  }
}
