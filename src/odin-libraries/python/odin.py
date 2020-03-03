from sys import exit
from ruamel.yaml import YAML
from odin_logger import OdinLogger as logger

class Odin:
    def __init__(self, config="job.yml", mongodb="mongodb://localhost:27017"):
        self.config = config
        data = YAML().load(open(self.config,"r").read())
        self.id = data["job"]["id"]
        self.mongodb = mongodb

    def condition(self, desc, expr):
        logger.log("condition", desc, expr, self.id, self.mongodb)
        return expr
    
    def watch(self, desc, value):
        logger.log("watch", desc, value, self.id, self.mongodb)

    def result(self, desc, status):
        logger.log("result", desc, status, self.id, self.mongodb)
        exit(0)
