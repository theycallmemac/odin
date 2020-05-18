var mongoose = require('mongoose');
var schema = mongoose.Schema;

var observabilitySchema = new schema({
    desc:{
        type:String,
        required:true
    },
    id:{
        type:String,
        required:true,
    },
    timestamp:{
        type:String,
        required:true,
    },
    type:{
        type:String,
        required:true,
    },
    valie:{
        type:String,
        required:true,
    }
})

module.exports = mongoose.model('observability',observabilitySchema,'observability');