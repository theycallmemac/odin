const MongoClient = require('mongodb').MongoClient;
const assert = require('assert')

// Connection URL
const url = (process.env.ODIN_MONGODB)

class OdinLogger {
    constructor() {
        MongoClient.connect(url, function(err, client){
            assert.equal(null, err);
            this.db = client.db('odin')
            this.collection = this.db.collection('observability')
        })
    }

    async log(type, desc, value, id, timestamp, collection=this.collection){
        await OdinLogger.insert(collection, type, desc, value, id, timestamp)
    }

    static async insert(collection, type, desc, value, id, timestamp){
        await collection.insertOne({
            'type' : String(type),
            'desc' : String(desc),
            'value' : String(value),
            'id' : String(id),
            'timestamp' : String(timestamp)
        })
    }
}

module.exports.OdinLogger = OdinLogger;