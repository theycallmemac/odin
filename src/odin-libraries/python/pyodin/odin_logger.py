from pymongo import errors, MongoClient
from os import environ

client = MongoClient(environ.get('ODIN_MONGODB'), serverSelectionTimeoutMS=2000)

class OdinLogger:
    @classmethod
    def log(self, type, desc, value, id, timestamp):
        if self.check_connection(client): 
            db = client['odin']
            collection = db['observability']
            self.find_and_insert(collection, type, desc, value,  id, timestamp)
    
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
        except errors.ServerSelectionTimeoutError as e:
            return False

