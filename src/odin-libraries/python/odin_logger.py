from pymongo import MongoClient
from os import environ

client = MongoClient(environ.get('ODIN_MONGODB'))
# get odin db
db = client['odin']
collection = db['observability']

class OdinLogger:
    @classmethod
    def log(self, type, desc, value, id, collection=collection):
        self.find_and_insert(collection, type, desc, value,  id)
    
    @staticmethod
    def find_and_insert(collection, type, desc, value, id):
        collection.update_one(
            {'id': str(id), 'desc':str(desc), 'type': str(type)},
            {'$set' : {
                'type' : str(type),
                'desc' : str(desc),
                'value' : str(value),
                'id' : str(id),
            }}, upsert=True
        )