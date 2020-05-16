var express = require('express');
var router = express.Router();
var Job = require('../models/Job');
var auth = require('./auth');


router.post('/getJobs', function(req , res) {
    let token = req.body.token;

    try {
        var decoded = auth.verifyToken(token);
    } catch(err) {
        // invalid token, secret does not match
        return res.status(401).json({message: 'Unauthorized client'})
    }
    Job.find({}, function(err, jobs) { 
        if (err) { return res.status(400).json({message: "No jobs found"}) }
        else {
            var jobMap = {}; 
            jobs.forEach(function(job) { 
                jobMap[job._id] = job; 
            }); 
            
            res.status(200).json(jobs); 
        }
    });
})

router.post('/getJob', function(req , res) {
    let token = req.body.token;
    let jobid = req.body.jobid

    try {
        var decoded = auth.verifyToken(token);
    } catch(err) {
        // invalid token, secret does not match
        return res.status(401).json({message: 'Unauthorized client'})
    }
    Job.findOne({id : jobid}, function(err, job) { 
        if (err) { return res.status(404).json({message: "Job not found"}) }
        else {
            res.status(200).json(job); 
        }
    });
})


module.exports = router;
