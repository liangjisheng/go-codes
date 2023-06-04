const https = require("https");
const fs = require("fs");

const options = {
  key: fs.readFileSync("127.0.0.1+1-key.pem"),
  cert: fs.readFileSync("127.0.0.1+1.pem")
};

https
  .createServer(options, (req, res) => {
    res.writeHead(200);
    res.end("hello world\n");
  })
  .listen(8000);
