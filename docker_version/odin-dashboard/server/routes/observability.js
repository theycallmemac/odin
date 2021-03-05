var express = require('express');
var router = express.Router();
var Observability = require('../models/Observability');
var Job = require('../models/Job');
var auth = require('./auth');

router.post('/jobrunstatus', function(req , res) {
  let token = req.body.token;
  let jobId = req.body.jobId;

  try {
      auth.verifyToken(token);
  } catch(err) {
      // invalid token, secret does not match
      res.status(401).json({status: false, message: 'Unauthorized client'})
  }
  Observability.findOne({id: String(jobId), type: 'result'}, function(err, result) { 
      if (err) { 
        res.status(400).json({status: false, message: "Error finding status"});
        console.log(err)
       }
      else {
        if (result != null) {
          result = result.toObject()
          res.status(200).json({status: true, message:  { status : result.value, time: result.timestamp }})
        } else {
          res.status(200).json({status: false, message: "Status not found"})
        }
      }
  });
})

router.post('/jobmetrics', function(req , res) {
  let token = req.body.token;
  let jobId = req.body.jobId;

  try {
      auth.verifyToken(token);
  } catch(err) {
      // invalid token, secret does not match
      res.status(401).json({status: false, message: 'Unauthorized client'})
  }
  Job.findOne({id : jobId}, function(err, job) {
    if (err) { 
      return res.status(500).json({success: false , message: "Failed to query DB"})
    } else if (job == null) {
      return res.status(404).json({success: false , message: "Job doesn't exist"})
    } else {
      // finds metrics for the provided job, groups by metric description
      // sort applied to have consistent order since it's not predictive and can vary given same inputs
      Observability.aggregate([
        {
          $match : { id: jobId }
        },
        { 
          $group : { 
            _id : "$desc", 
            observability: { $push: "$$ROOT" } 
          }
        },
        {
          $sort : { _id : 1}
        }
        ], function(err, result) {
          if (err) {
            return res.status(500).json({success: false , message: "Failed to query DB"})
          } else if (result == null) {
            return res.status(404).json({success: false , message: "No job Metrics exist"})
          } else {
            return res.status(200).json({success: true , message: result})
          }
        }
      )
    }
  })
})
        
// io.on("connection", socket => {
//     let previousId;
//     const safeJoin = currentId => {
//       socket.leave(previousId);
//       socket.join(currentId);
//       previousId = currentId;
//     };
  
//     socket.on("getDoc", docId => {
//       safeJoin(docId);
//       socket.emit("document", documents[docId]);
//     });
  
//     socket.on("addDoc", doc => {
//       documents[doc.id] = doc;
//       safeJoin(doc.id);
//       io.emit("documents", Object.keys(documents));
//       socket.emit("document", doc);
//     });
  
//     socket.on("editDoc", doc => {
//       documents[doc.id] = doc;
//       socket.to(doc.id).emit("document", doc);
//     });
  
//     io.emit("documents", Object.keys(documents));
//   });


module.exports = router;