from flask import Flask, render_template, send_file, send_from_directory, json, jsonify, make_response, config
import json

app = Flask(__name__, static_url_path='')


@app.route('/')
def hello_flask():
    return render_template('datastruct_inuse.htm')


@app.route('/query', methods=['GET', 'POST'])
def query_json():
    json_file = None
    with open('location.json') as f:
        json_file = json.load(f)

        return jsonify(json_file)


@app.route('/get_file/<file_name>', methods=['GET', 'POST'])
def get_file(file_name):
    response = make_response(send_file(file_name))
    response.headers["Content-Disposition"] = "attachment; filename={};".format(file_name)
    return response


if __name__ == '__main__':
    app.run(debug=True)
    print("none")
