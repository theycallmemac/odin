const odinlib = require('../odinlib/odin.js');
const mongo = require('./mongo.js');
const assert = require('assert');

mongodb = new mongo.MongoDB();

describe('Watch Operation - Not Odin Env ', () => {
    it('should check a failed watch operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=false);
        odin.watch(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});

describe('Watch Operation', () => {
    it('should check a successful watch operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=true);
        odin.watch(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});

describe('Condition Operation - Not Odin Env ', () => {
    it('should check a failed condition operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=false);
        odin.condition(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});

describe('Condition Operation', () => {
    it('should check a successful condition operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=true);
        odin.condition(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});

describe('Result Operation - Not Odin Env ', () => {
    it('should check a failed result operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=false);
        odin.result(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});

describe('Result Operation', () => {
    it('should check a successful result operation', () => {
        testDesc = 'test_desc' + Math.floor(Math.random() * (999999 - 100000) + 100000)
        odin = new odinlib.Odin(config="job.yml", path="relative", test=true);
        odin.result(testDesc, true)
        desc = mongodb.checkMongo(testDesc);
    });
});
