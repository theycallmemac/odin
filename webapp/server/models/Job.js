var mongoose = require('mongoose');
var schema = mongoose.Schema;

var jobSchema = new schema({
    id:{
        type:String,
        required:true
    },
    uid:{
        type:String,
        required:true,
    },
    gid:{
        type:String,
        required:true,
    },
    name:{
        type:String,
        required:true,
    },
    description:{
        type:String,
        required:true,
    },
    language:{
        type:String,
        required:true,
    },
    file:{
        type:String,
        required:true,
    },
    stats:{
        type:String,
        required:false,
    },
    schedule:{
        type:String,
        required:true,
    },
    runs:{
        type:Number,
        required:true,
    },
    link:{
        type:String,
        required:false,
    }
})

module.exports = mongoose.model('jobs',jobSchema,'jobs');