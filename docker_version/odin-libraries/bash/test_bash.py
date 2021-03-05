import unittest
from pymongo import MongoClient
import json
from os import system, environ
import random

class odinSdkTest(unittest.TestCase):
    def setUp(self):
        client = MongoClient(environ.get('ODIN_MONGODB'))
        mongodb = client['odin']
        self.collection = mongodb['observability']

    def testConditionNotOdinEnv(self):
        r = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(r)
        
        environ["ODIN_EXEC_ENV"]  = "False"
        system("/bin/odinbash  -p relative -i job.yml")
        system("/bin/odinbash -p relative -d " + test_desc + " -c True")

        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(None, result)

    def testWatchNotOdinEnv(self):
        r = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(r)
        
        environ["ODIN_EXEC_ENV"]  = "False"
        system("/bin/odinbash  -p relative -i job.yml")
        system("/bin/odinbash -p relative -d " + test_desc + " -w True")

        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(None, result)

    def testCondition(self):
        r = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(r)
        
        environ["ODIN_EXEC_ENV"]  = "True"
        system("/bin/odinbash  -p relative -i job.yml -t")
        system("/bin/odinbash -p relative -d " + test_desc + " -c True -t")

        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(test_desc, result['description'])


    def testWatch(self):
        r = random.randint(100000, 999999)
        test_desc = 'test_desc' + str(r)

        environ["ODIN_EXEC_ENV"]  = "True"
        system("/bin/odinbash  -p relative -i job.yml -t")
        system("/bin/odinbash -p relative -d " + test_desc + " -c True -t")

        result = self.collection.find_one({"description" : test_desc})

        self.assertEqual(test_desc, result['description'])

if __name__ == "__main__":
    unittest.main() # run all tests
