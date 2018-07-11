from flask import Flask, request, jsonify
from flask_restful import Resource, Api
from json import dumps

app = Flask(__name__)
api = Api(app)


def processRawData(raw_url):

	return "checked: "+raw_url

class Raws(Resource):
	def get(self,raw_url):
		raw_url = processRawData(raw_url)
		result={'code':200,'isMalicious':'false','data':raw_url}
		return jsonify(result)

api.add_resource(Raws,'/classify/<raw_url>')

if __name__ == '__main__':
	app.run(port='5002')