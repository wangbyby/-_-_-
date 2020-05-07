from flask import Flask, render_template, send_file,  json, jsonify, make_response, request,Response
import json
import  yagmail
import location

app = Flask(__name__, static_url_path='')

#send emails 的配置
email_args = {
    "user":"xxx@qq.com",
    "password":"xxx",
    "host":"smtp.qq.com",
    "port":"465"
}

@app.route('/')
def hello_flask(): #根目录
    return render_template('datastruct_inuse.htm')


@app.route('/query', methods=['GET', 'POST'])
def query_json():#返回经纬度
    with open('location.json') as f:
        json_file = json.load(f)
        return jsonify(json_file)
#返回图片
@app.route('/img/<img_name>', methods=['GET', 'POST'])
def get_img(img_name):
    img= 0
    img_path = "./img/"+img_name+".png"
    with open(img_path, 'rb') as f:
        img = f.read()
    return Response(img, mimetype='image/png')
#返回服务器资源, 脚本文件...
@app.route('/get_file/<path:file_name>', methods=['GET', 'POST'])
def get_file(file_name): 
    print(file_name)
    response = make_response(send_file(file_name))

    response.headers["Content-Disposition"] = "attachment; filename={};".format(file_name)
    return response

@app.route('/email',methods=['GET','POST'])
def send_email(): #发送邮件
    if request.method == 'GET':
        return render_template('tips.htm')
    elif request.method == 'POST':
        data = None
        for i in request.form:
            data = json.loads(i)
            print(data)

        # receive = data['emails']
        receive = data.get('emails', '')
        if receive == '':
            return  "No email"
        email = yagmail.SMTP(**email_args)
        subject = data['subject']
        contents = data['contents']
        with open('emails.htm','w') as f:
            f.write(contents)
        email.send(to=receive, subject=subject,contents=[contents,'emails.htm'])

        return "OK"
@app.route('/data', methods=['GET',"POST"])
def data(): #返回要查询的路径
    res = None
    g = location.Graph()
    g.read_from_file()
    for i in request.form:
        search = json.loads(i)
        res = g.search(search)
    print(res)
    return jsonify(res)
if __name__ == '__main__':
    app.run(debug=True)

