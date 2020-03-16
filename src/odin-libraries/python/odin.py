from os import environ
from sys import exit
from ruamel.yaml import YAML
from odin_logger import OdinLogger as logger

ENV_CONFIG = 'ODIN_EXEC_ENV' in environ
MONGODB = environ.get('ODIN_MONGODB')

class Odin:
    def __init__(self, config="job.yml", mongodb=string(MONGODB)):
        self.config = config
        data = YAML().load(open(self.config,"r").read())
        self.id = data["job"]["id"]
        self.mongodb = mongodb

    def condition(self, desc, expr):
        if ENV_CONFIG:
            logger.log("condition", desc, expr, self.id, self.mongodb)
        return expr
    
    def watch(self, desc, value):
        if ENV_CONFIG:
            logger.log("watch", desc, value, self.id, self.mongodb)

    def result(self, desc, status):
        if ENV_CONFIG:
            logger.log("result", desc, status, self.id, self.mongodb)
        exit(0)
