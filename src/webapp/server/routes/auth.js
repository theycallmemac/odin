var express = require('express');
var User = require('../models/User');
var router = express.Router();

const jwt = require('jsonwebtoken');
const {OAuth2Client} = require('google-auth-library');
const secret = "GEBQi33zH5sZ1ENNXbKZQlnN";
const client = new OAuth2Client(process.env.CLIENTID);

/* GET home page. */
router.get('/auth', function(req, res, next) {
    res.render('index', { title: 'Auth' });
});

// Verify JWT
function verifyToken(token){
    return jwt.verify(token, secret);
}

// Return users account information
router.post('/user', function(req, res) {
    let token = req.body.token;
    // Verify Google OAuth Token 
    let verifiedToken = verifyToken(token);

    User.findOne({userid:verifiedToken.userid}, function(err, user) {
    if (err) { return res.status(400).json({message: "User not found"}) }
    else {return res.status(200).json({email: user.email, name: user.fullname, photo: user.photo})}
    });
})


// Log the user in, if user account does not exist in DB, create account and authenticate
router.post('/login', function(req, res) {
    // Verify Google OAuth Token
    async function verify(token) {
        const ticket = await client.verifyIdToken({idToken: token, audience: process.env.CLIENT_ID});
        const payload = ticket.getPayload();
        const audience = payload['aud'];

        if(audience != process.env.CLIENTID){
            return false;
        } else {
            return true;
        }
    }
   
    result = verify(req.body.id_token);
    userid = req.body.userid;
    name = req.body.name;
    token = req.body.id_token;
    email = req.body.email;
    image = req.body.image;

    if (result != false) {
        let user;

        function returnAuth(user) {
            const retToken = jwt.sign({email: user.email, userid: user.userid}, secret);
            return res.status(200).json({
            token: retToken,
            // set token expiry time to 24h
            expiresIn: 86400,
            userid: user.userid
            });
        }

        User.findOne({email: email},function(err,doc){
            if(err) { return res.status(500).json({message:'error occured'}) }
            else {
            // user found in DB, return token
            if(doc) { returnAuth(doc) }
            // user not found, add new user record to DB
            else {
                var record = new User({
                userid: userid,
                name: name,
                token: token,
                email: email,
                photo: image
                });

                record.save( (err,user) => {
                if(err){ return res.status(500).json({message: 'DB error while saving'}) } 
                else { returnAuth(user) }
                });
            }
            }
            }
        );
    }
})


module.exports = router;
