const MongoClient = require('mongodb').MongoClient;
const assert = require('assert')

// Connection URL
const url = (process.env.ODIN_MONGODB)

class OdinLogger {
    constructor() {
        MongoClient.connect(url, function(err, client){
            assert.equal(null, err);
            if (err) {
                this.success = false;
            }
            this.db = client.db('odin');
            this.collection = this.db.collection('observability');
            this.success = true;
        })
    }

    async log(type, desc, value, id, timestamp, collection=this.collection){
        if (this.success == true){
            await OdinLogger.insert(collection, type, desc, value, id, timestamp)
        }
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