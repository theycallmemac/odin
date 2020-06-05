const odin = require('../odinlib/odin.js');
const assert = require('assert');

const MongoClient = require('mongodb').MongoClient;
const url = process.env.ODIN_MONGODB;


class MongoDB {
    async checkMongo(testDesc) {
	const client = await MongoClient.connect(url, {useUnifiedTopology: true})
        const db = client.db("odin");
        const results = db.collection("observability").find({}, {"desc": testDesc}).toArray();
        results.then((results) => {
            results.forEach(element => {
                if (element.descripton == testDesc) {
                    assert.equal(element.description, testDesc);
                }
            });
            client.close()
        }).catch((error) => {
            assert.equal(0,1);
            client.close()
        })
    }
}

module.exports.MongoDB = MongoDB;

