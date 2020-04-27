from pymongo import errors, MongoClient
from os import environ

try:
    client = MongoClient(environ.get('ODIN_MONGODB'), serverSelectionTimeoutMS=2000)
    db = client['odin']
    collection = db['observability']
    clientSuccess = True
except:
    clientSuccess = False

class OdinLogger:
    @classmethod
    def log(self, type, desc, value, id, timestamp):
        try:
            if (clientSuccess): 
                self.find_and_insert(collection, type, desc, value,  id, timestamp)
        except:
            # if logging fails, check connection 
            clientSuccess = self.check_connection(client)

    
    @staticmethod
    def find_and_insert(collection, type, desc, value, id, timestamp):
        collection.update_one(
            {'id': str(id), 'desc':str(desc), 'type': str(type), 'timestamp': str(timestamp)},
            {'$set' : {
                'type' : str(type),
                'desc' : str(desc),
                'value' : str(value),
                'id' : str(id),
                'timestamp': str(timestamp)
            }}, upsert=True
        )
    
    @staticmethod
    def check_connection(client):
        try:
            client.admin.command('ping')
            return True
        except errors.ServerSelectionTimeoutError:
            return False

