const http = require('http');
const server = http.createServer((req,res)=>{
    res.end('helloworld!');
});
server.listen(8087);
