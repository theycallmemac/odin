var assert = require('assert');
const request = require('supertest');
const jwt = require('jsonwebtoken');
const { testUser, HOST } = require('./tests-helper')


describe('Auth unit tests', function () {

    const secret = "GEBQi33zH5sZ1ENNXbKZQlnN";
    var testToken;

    // generate a signed test user JWT token before auth tests
    before(function() {
        testToken = jwt.sign({email: testUser.email, userid: testUser.userid}, secret);
    });

    // test disabled for now as google login tokens have a short expiry time, will look for solution
    // it('should return http status code 200 when provided with valid login details', function () {
    //     const payload = testUser;
    //     request(HOST)
    //     .post('/auth/login')
    //     .send(payload)
    //     .set('Accept', 'application/json')
    //     .expect(200)
    //     .end(function(err, res) {
    //         if (err) throw err;
    //         assert.equal(res.body.userid, "118141658548859160434");
    //         done();
    //     });
    // });

    it('should return http status code 200 if requested user exists in DB and is returned', (done) => {
        payload = {
            token : testToken
        }
        request(HOST)
            .post('/auth/user')
            .send(payload)
            .set('Accept', 'application/json')
            .expect(200)
            .end(function(err, res) {
                if (err) throw err;
                assert.equal(res.body.email, "odintestsuser@gmail.com");
                done();
            })
    });

    it('should return http status code 500 if requested user does not exist', (done) => {
        payload = {
            token : "totallyNotARandomString"
        }
        request(HOST)
            .post('/auth/user')
            .send(payload)
            .set('Accept', 'application/json')
            .expect(500)
            .end(function(err, res) {
                if (err) throw err;
                done();
            })
    });
});