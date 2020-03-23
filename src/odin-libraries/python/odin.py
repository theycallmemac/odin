from os import environ
from sys import exit
from ruamel.yaml import YAML
import json
from odin_logger import OdinLogger as logger


class Odin:
    def __init__(self, config="job.yml", test=False):
        self.config = config
        try: 
            with open(self.config,"r") as config:
                configR = config.read()
            data = YAML().load(configR)
        except Exception as e:
            print(e)
        self.id = data["job"]["id"]

        if 'ODIN_EXEC_ENV' in environ or test != False:
            self.ENV_CONFIG = True
        else:
            self.ENV_CONFIG = False

    def condition(self, desc, expr):
        if self.ENV_CONFIG:
            logger.log("condition", desc, expr, self.id)
        return expr
    
    def watch(self, desc, value):
        if self.ENV_CONFIG:
            logger.log("watch", desc, value, self.id)

    def result(self, desc, status):
        if self.ENV_CONFIG:
            logger.log("result", desc, status, self.id)
        exit(0)