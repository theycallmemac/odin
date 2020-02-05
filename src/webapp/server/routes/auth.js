var express = require('express');
var User = require('../models/User');
var router = express.Router();

const jwt = require('jsonwebtoken');
const {OAuth2Client} = require('google-auth-library');
const secret = process.env.SECRET;
const client = new OAuth2Client(process.env.CLIENTID);

/* GET home page. */
router.get('/auth', function(req, res, next) {
    res.render('index', { title: 'Auth' });
});

// Verify JWT
function verifyToken(token){
    console.log(token, secret)
    return jwt.verify(token, secret);
}

// Return a users account information
router.post('/user', function(req, res) {
    let token = req.body.token;
    // Verify Google OAuth Token
    let verifiedToken = verifyToken(token);

    User.findOne({userId:verifiedToken.userId}, function(err, user) {
    if (err) { return res.status(400).json({message:"Error finding user"}) }
    else { return res.status(200).json({email: user.email, name: user.fullname, photo: user.photo}) }
    });
})


// Log the user in or create an account then log in
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
            // set token expire time to 24h
            expiresIn: 86400,
            userid: user.userid
            });
        }

        User.findOne({email: email},function(err,doc){
            if(err) { return res.status(500).json({message:'error occured'}) }
            else {
            // user found, return token
            if(doc) { returnAuth(doc) }
            // user not found, create new user
            else {
                var record = new User({
                userid: userid,
                name: name,
                token: token,
                email: email,
                photo: image
                });

                record.save( (err,user) => {
                if(err){ return res.status(500).json({message: 'db error while saving'}) } 
                else { returnAuth(user) }
                });
            }
            }
            }
        );
    }
})


module.exports = router;
