import urllib
array=[]
url = raw_input('Enter url: ')
feature=["SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "RENAME", "WHERE", "FROM", "UNION", "NOT", "AND", "OR", "XOR","EXEC","!", "&&", "||", "--", "#", "<", ">", "<=>", ">=", "<=", "==", "=", "!=", "<<", ">>", "<>", "%", "*", "?", "|", "&", "-", "+","/**/"]
print "length: ",len(url)
array.append(len(url))


for x in feature:
    if url.__contains__(x):
        array.append(1)
    else:
        array.append(0)

print array
