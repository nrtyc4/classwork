from SimpleXMLRPCServer import SimpleXMLRPCServer
from SimpleXMLRPCServer import SimpleXMLRPCRequestHandler
import ConfigParser
import sqlite3

#read in from .conf file to get the path of the database and port to connect to
config = ConfigParser.RawConfigParser()
config.read('httpd.conf')
path = config.get('Section1', 'Path')
port = config.getint('Section1', 'Port')
#print port

# instantiate a server object
server = SimpleXMLRPCServer (("0.0.0.0", port), SimpleXMLRPCRequestHandler)
#server.register_introspection_functions()


#connect/create database and create cursor
conn = sqlite3.connect(path)
c = conn.cursor()

# define your methods in a class
class RPCMethods :
    def query (self, q) :
        c.execute(q)
        return c.fetchall(c)
      
    def execute (self, q) :
        c.execute(q)
        conn.commit()
        return "The query ",q," was successful"

# start the server
server.register_instance (RPCMethods ())
server.serve_forever ()
