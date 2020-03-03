import motor.motor_asyncio
import asyncio
from os import environ

# start mongodb driver client
client = motor.motor_asyncio.AsyncIOMotorClient("mongodb://localhost:27017")

# get odin db
db = client['odin']

class OdinLogger:
    @classmethod
    def log(self, type, desc, value, id):
        collection = db['observability']
        collection.insert_one({
            'type' : str(type),
            'desc' : str(desc),
            'value' : str(value),
            'id' : str(id),
        })
