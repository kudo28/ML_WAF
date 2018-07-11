#!/usr/bin/python

from flask import Flask, request, jsonify
from flask_restful import Resource, Api
from json import dumps

app = Flask(__name__)
api = Api(app)

features=["SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "RENAME", "WHERE", "FROM", "UNION", "NOT", "AND", "OR", "XOR","EXEC","!", "&&", "||", "--", "#", "<", ">", "<=>", ">=", "<=", "==", "=", "!=", "<<", ">>", "<>", "%", "*", "?", "|", "&", "-", "+","/**/"]

# process raw data into features to classify
def processRawData(raw_url):
	array=[]
	array.append(len(raw_url))
	for x in features:
	    print(x +": ")
    	if raw_url.lower().__contains__(x.lower()):
    	   print('yes')     #    	array.append(1)
   	 	else:
		   print('no')

	return array

# using ML Model to classify
def classify(record):
	return "test"


class Raws(Resource):
	def get(self,raw_url):
		data = processRawData(raw_url)
		result={'code':200,'isMalicious':'false','data':data}
		return jsonify(result)

api.add_resource(Raws,'/classify/<raw_url>')

if __name__ == '__main__':
	app.run(port='5002')