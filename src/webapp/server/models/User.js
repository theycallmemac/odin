var mongoose = require('mongoose');
var schema = mongoose.Schema;

var userSchema = new schema({
    userid:{
        type:String,
        required:true
    },
    name:{
        type:String,
        required:true,
    },
    token:{
        type:String,
        required:true,
    },
    email:{
        type:String,
        required:true,
    },
    photo:{
        type:String,
        required:true,
    }
})

module.exports = mongoose.model('User',userSchema,'User');
