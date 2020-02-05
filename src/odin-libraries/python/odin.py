from os import environ
from ruamel.yaml import YAML
import odin_logger as logger

ENV_CONFIG = 'ODIN_EXEC_ENV' in environ

class Odin:

    def __init__(self, config_file="odin_job.yml"):
        data = YAML().load(open(config_file,"r").read())
        self.id = data['id']
        self.run = data['run']

    def watch(self, desc, value):
        if ENV_CONFIG:
            logger.log('watch', desc, value, self.id, self.run)

    def result(self, desc, status):
        if ENV_CONFIG:
            logger.log('result', desc, status, self.id, self.run)