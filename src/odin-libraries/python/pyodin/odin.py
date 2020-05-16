from os import environ, listdir, path
from sys import exit
from ruamel.yaml import YAML
import json
from pyodin.odin_logger import OdinLogger as logger
from time import time

class Odin:
    def __init__(self, config="job.yml", test=False, pathType="absolute"):
        if pathType == "absolute":
            for file in listdir("/etc/odin/jobs"):
                if path.exists("/etc/odin/jobs/" + file + "/" + config):
                    self.config = "/etc/odin/jobs/" + file + "/" + config
                    break
        else:
            self.config = config
        try: 
            with open(self.config,"r") as config:
                configR = config.read()
            data = YAML().load(configR)
            self.id = data["job"]["id"]
            self.timestamp = time()
        except Exception as e:
            print(e)
        if 'ODIN_EXEC_ENV' in environ or test != False:
            self.ENV_CONFIG = True
        else:
            self.ENV_CONFIG = False

    def condition(self, desc, expr):
        if self.ENV_CONFIG:
            logger.log("condition", desc, expr, self.id, self.timestamp)
        return expr
    
    def watch(self, desc, value):
        if self.ENV_CONFIG:
            logger.log("watch", desc, value, self.id, self.timestamp)

    def result(self, desc, status):
        if self.ENV_CONFIG:
            logger.log("result", desc, status, self.id, self.timestamp)
        exit(0)
