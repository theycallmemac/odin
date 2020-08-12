""" Docstring """
from os import environ, listdir, path
import sys
from time import time
from ruamel.yaml import YAML
from pyodin.odin_logger import OdinLogger as logger


class Odin(object):
    """ Odin class used for Configuration File to persist in a job """
    def __init__(self, config="job.yml", test=False, path_type="absolute"):
        if path_type == "absolute":
            for job_id in listdir("/etc/odin/jobs"):
                if path.exists("/etc/odin/jobs/" + job_id + "/" + config):
                    self.config = "/etc/odin/jobs/" + job_id + "/" + config
                    break
        else:
            self.config = config
        try:
            with open(self.config, "r") as yaml_config:
                config_r = yaml_config.read()
            data = YAML().load(config_r)
            self.job_id = data["job"]["id"]
            self.timestamp = time()
        except FileNotFoundError as fnf_error:
            print(fnf_error)
        if test:
            self.env_config = True
        elif 'ODIN_EXEC_ENV' in environ:
            self.env_config = True
        else:
            self.env_config = False

    def condition(self, desc, expr):
        """ Function condition takes a description and an expression """
        if self.env_config:
            logger.log("condition", desc, expr, (self.job_id, self.timestamp))
        return expr

    def watch(self, desc, value):
        """ Function watch takes a description and a value """
        if self.env_config:
            logger.log("watch", desc, value, (self.job_id, self.timestamp))

    def result(self, desc, status):
        """ Function result takes a description and a status"""
        if self.env_config:
            logger.log("result", desc, status, (self.job_id, self.timestamp))
        sys.exit(0)
