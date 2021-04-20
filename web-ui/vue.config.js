module.exports = {
  lintOnSave: false,
  assetsDir: "static",
  devServer: {
    proxy: 'http://localhost:5000'
  }
}
