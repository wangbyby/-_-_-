from flask import Flask, render_template, send_file,  json, jsonify, make_response, request
import json
import location
app = Flask(__name__, static_url_path='')


@app.route('/')
def hello_flask():
    return render_template('datastruct_inuse.htm')


@app.route('/query', methods=['GET', 'POST'])
def query_json():
    with open('location.json') as f:
        json_file = json.load(f)
        return jsonify(json_file)


@app.route('/get_file/<file_name>', methods=['GET', 'POST'])
def get_file(file_name):
    response = make_response(send_file(file_name))
    response.headers["Content-Disposition"] = "attachment; filename={};".format(file_name)
    return response

@app.route('/data', methods=['GET',"POST"])
def data():
    res = None
    for i in request.form:
        search = json.loads(i)

        g = location.Graph()
        g.read_from_file()
        res = g.search(search)
    print(res)
    return jsonify(res)
if __name__ == '__main__':
    app.run(debug=True)
    print("none")
