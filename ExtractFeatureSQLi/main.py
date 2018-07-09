import numpy as np
import pandas as pd
from sklearn.naive_bayes import GaussianNB
from sklearn.preprocessing import LabelEncoder
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score

headers=["length","SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "RENAME", "WHERE", "FROM", "UNION", "NOT", "AND", "OR", "XOR","EXEC",
	"!", "&&", "||", "--", "#", "<", ">", "<=>", ">=", "<=", "==", "=", "!=", "<<", ">>", "<>", "%", "*", "?", "|", "&", "-", "+","/**/"]
SQLi_test = pd.read_csv("dataset.csv")
SQLi_test.head()
number = LabelEncoder()
SQLi_test['URL'] = number.fit_transform(SQLi_test['URL'])
SQLi_test['label'] = number.fit_transform(SQLi_test['label'])

for header in headers:
	SQLi_test[header] = number.fit_transform(SQLi_test[header])
features=headers
target = "label"
# print SQLi_test
features_train, features_test, target_train, target_test = train_test_split(SQLi_test[features],SQLi_test[target],test_size = 0.39,random_state = 54)
model = GaussianNB()
model.fit(features_train, target_train)
pred = model.predict(features_test)
accuracy = accuracy_score(target_test, pred)
print accuracy
# print model.predict([[1,2,0,1]])
