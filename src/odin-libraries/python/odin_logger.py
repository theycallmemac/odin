from pymongo import MongoClient

class OdinLogger:
    @classmethod
    def log(self, type, desc, value, id, mongo):
        # start mongodb driver client
        client = MongoClient(mongo)
        # get odin db
        db = client['odin']
        collection = db['observability']
        self.find_and_insert(collection, type, desc, value,  id)

    def find_and_insert(collection, type, desc, value, id):
        collection.update(
            {'id': str(id), 'desc':str(desc), 'type': str(type)},
            {'$set' : {
                'type' : str(type),
                'desc' : str(desc),
                'value' : str(value),
                'id' : str(id),
            }}, upsert=True
        )

