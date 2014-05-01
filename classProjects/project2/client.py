import xmlrpclib

# connect to your XML-RPC Server
proxy = xmlrpclib.ServerProxy('http://classwork-25688.usw1.actionbox.io:8002')

#print list of available methods
print proxy.system.listMethods()

# call methods on the server
print proxy.execute("CREATE TABLE test (lastname, age)")
print proxy.execute("INSERT INTO test VALUES ('Thompson', 21)")

# print the results of a query
print proxy.query("SELECT * FROM test")
