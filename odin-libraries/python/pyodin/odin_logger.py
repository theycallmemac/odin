from os import environ
from requests import post

class OdinLogger:
    @classmethod
    def log(cls, type, desc, value, id, timestamp):
        response = post(url="http://localhost:3939/stats/add", data = type + "," + desc + "," + str(value) + "," + id + "," + str(timestamp))
        return response.status_code
