const fs = require('fs');

module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    // public: "0.0.0.0",
    host: "127.0.0.1",
    port: 8080,
    hot: true,
    open: true,
    headers: { "Access-Control-Allow-Origin": "*" },
    disableHostCheck: true,
    proxy: {
      "/api/v1/": {
        target: "http://localhost:8081/",
        pathRewrite: {'^/api/v1' : ''},
        changeOrigin: true
      }
    },
    https: {
      key: fs.readFileSync('./certs/server.key'),
      cert: fs.readFileSync('./certs/server.cert'),
    }
  }
}