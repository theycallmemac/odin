""" Runs tests for Ptyhon Odin SDK """

import unittest
from os import environ
import random
from pymongo import MongoClient
import pyodin as odin

class OdinSdkTest(unittest.TestCase):
    """ Establish OdinSdkTest object """

    def setUp(self):
        client = MongoClient(environ.get('ODIN_MONGODB'))
        mongodb = client['odin']
        self.collection = mongodb['observability']

    def tearDown(self):
        self.collection.delete_many({"id" : "test_id"})

    def test_condition_not_odin_env(self):
        """ Run condition operation outside of Odin Env """
        random_int = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(random_int)

        odin_test = odin.Odin(config="job.yml", path_type="relative")

        cond = odin_test.condition(test_desc, True)
        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(cond, True)
        self.assertEqual(None, result)

    def test_watch_not_odin_env(self):
        """ Run watch operation outside of Odin Env """
        random_int = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(random_int)

        odin_test = odin.Odin(config="job.yml", path_type="relative")

        odin_test.watch(test_desc, True)
        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(None, result)

    def test_condition(self):
        """ Run condition operation inside Odin Env """
        random_int = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(random_int)

        # test True sets odin exc env to true and in turn enables logging everything to the DB
        odin_test = odin.Odin(test=True, config="job.yml", path_type="relative")

        cond = odin_test.condition(test_desc, True)
        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(cond, True)
        self.assertEqual(test_desc, result['description'])

    def test_watch(self):
        """ Run watch operation inside Odin Env """
        random_int = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(random_int)

        # test True sets odin exc env to true and in turn enables logging everything to the DB
        odin_test = odin.Odin(test=True, config="job.yml", path_type="relative")

        odin_test.watch(test_desc, True)
        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(test_desc, result['description'])

if __name__ == "__main__":
    unittest.main() # run all tests
