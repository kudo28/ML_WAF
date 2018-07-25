#!/usr/bin/python

from flask import Flask, request, jsonify
from flask_restful import Resource, Api
from sklearn.externals import joblib
from json import dumps

app = Flask(__name__)
api = Api(app)
# load model from file
filename = "finalized_model.sav"
model = joblib.load(filename)

features = ["SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "RENAME", "WHERE", "FROM", "UNION",
            "NOT", "AND", "OR", "XOR", "EXEC", "!", "&&", "||", "--", "#", "<", ">", "<=>", ">=", "<=", "==", "=", "!=",
            "<<", ">>", "<>", "%", "*", "?", "|", "&", "-", "+", "/**/"]


# process raw data into features to classify
def processRawData(raw_url):
    array = []
    array.append(len(raw_url))
    for x in features:
        if raw_url.lower().__contains__(x.lower()):
            # print('yes')
            array.append(1)
        else:
            # print('no')
            array.append(0)
    print(array.__len__())
    return array


# using ML Model to classify
def classify(record):
    return model.predict([record])[0]


class Raws(Resource):
    def post(self):
        raw_url = request.json['raw']
        data = processRawData(raw_url)
        checkresult = classify(data)
        print (checkresult)
        result = {'code': 200, 'isMalicious': checkresult, 'data': checkresult}
        return jsonify(result)


api.add_resource(Raws, '/classify')

if __name__ == '__main__':
    app.run(port='5002')
