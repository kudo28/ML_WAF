f = open("payload_train.csv","r")
baseurl = "http://localhost:8080/tienda1/index.jsp?query="
file= open("raw_payload_train.txt","w")
for line in f:
	one = line.split("\",\"")
	# print line
	if one[2]=="sqli":
		s= baseurl+one[0]
		file.write(s+"\n")