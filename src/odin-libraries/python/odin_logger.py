import motor.motor_asyncio
import asyncio
from os import environ

# start mongodb driver client
client = motor.motor_asyncio.AsyncIOMotorClient(environ['ODIN_MONGODB_URL'])
# get odin db
db = client['odin']

class Odin_logger:

    @classmethod
    def log(cls, type, desc, value_status, job_id, run_number):
        collection = db['odin-observability']
        collection.insert_one({
            'type' : str(type),
            'desc' : str(desc),
            'value' : str(value_status),
            'job_id' : str(job_id),
            'run_number' : str(run_number)
        })