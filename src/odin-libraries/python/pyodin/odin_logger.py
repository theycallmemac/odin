from pymongo import errors, MongoClient
from os import environ

try:
    CLIENT = MongoClient(environ.get('ODIN_MONGODB'), serverSelectionTimeoutMS=2000)
    DB = CLIENT['odin']
    COLLECTION = DB['observability']
    CLIENT_SUCCESS = True
except:
    CLIENT_SUCCESS = False


class OdinLogger:
    @classmethod
    def log(cls, type, desc, value, id, timestamp):
        global CLIENT_SUCCESS
        try:
            if CLIENT_SUCCESS:
                cls.find_and_insert(COLLECTION, type, desc, value,  id, timestamp)
        except:
            # if logging fails, check connection
            global CLIENT
            CLIENT_SUCCESS = cls.check_connection(CLIENT)


    @staticmethod
    def find_and_insert(COLLECTION, type, desc, value, id, timestamp):
        COLLECTION.update_one(
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

