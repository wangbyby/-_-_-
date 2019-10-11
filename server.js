var exp = require('express')
var app = exp()
const nodemailer = require('nodemailer')

var bodyParser = require('body-parser');

// 创建 application/x-www-form-urlencoded 编码解析
var urlencodedParser = bodyParser.urlencoded({ extended: false })
const port = 7999
var server = app.listen(port, function () {
    var host = server.address().address
    var port = server.address().port
    console.log("实例访问地址 http://127.0.0.1:" + port + "/")
})
// 自定义跨域中间件
var allowCors = function (req, res, next) {
    res.header('Access-Control-Allow-Origin', req.headers.origin);
    res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS');
    res.header('Access-Control-Allow-Headers', 'Content-Type');
    res.header('Access-Control-Allow-Credentials', 'true');
    next();
};
app.use(allowCors);//使用跨域中间件

app.get('/', function (req, res) {
    res.sendFile(__dirname + "/" + "./view.htm");



})
app.post('/email', urlencodedParser, function (req, res) {
    var emailList = req.body['email[]']
    console.log(emailList)
})


let transporter = nodemailer.createTransport({
    service: 'qq',
    port: 465,
    secureConnection: true,
    auth: {
        user: '1329859972@qq.com',
        pass: 'entaros1c2MDZZ'
    }
})

let mainOp = {
    form: ' "bot" <1329859972@qq.com>',
    to: 'wang115byby@163.com',
    subject: 'Hello the path',
    html: '<b>Hello 我是火星黑洞</b>'
}

transporter.sendMail(mainOp, (error, info) => {
    if (error) {
        return console.log(error);
    }
    console.log('Message sent: %s', info.messageId);
})