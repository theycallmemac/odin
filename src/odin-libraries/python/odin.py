from os import environ
from ruamel.yaml import YAML
from odin_logger import OdinLogger as logger

class Odin:
    def __init__(self, config="job.yml"):
        data = YAML().load(open(config,"r").read())
        self.id = data["job"]["id"]

    def condition(self, desc, expr):
        logger.log("condition", desc, expr, self.id)
        return expr
    def watch(self, desc, value):
        logger.log("watch", desc, value, self.id)

    def result(self, desc, status):
        logger.log("result", desc, status, self.id)
