// dependencies
var async = require('async');
var AWS = require('aws-sdk');
var util = require('util');

exports.handler = function(event, context) {
  // Read options from the event.
  console.log('Received event:', JSON.stringify(event, null, 2));
  console.log('Context:', JSON.stringify(context, null, 2));

  var q = (event.q === undefined ? 'No-Query' : event.q);


  context.done(null, {"Hello":"World", "Q":q});


  // var srcBucket = event.Records[0].s3.bucket.name;
  // // Object key may have spaces or unicode non-ASCII characters.
  //   var srcKey    =
  //   decodeURIComponent(event.Records[0].s3.object.key.replace(/\+/g, " "));
  // var dstBucket = srcBucket + "poop";
  // var dstKey    = "resized-" + srcKey;

  // // Sanity check: validate that source and destination are different buckets.
  // if (srcBucket == dstBucket) {
  //   console.error("Destination bucket must not match source bucket.");
  //   return;
  // }

  // // Infer the image type.
  // var typeMatch = srcKey.match(/\.([^.]*)$/);
  // if (!typeMatch) {
  //   console.error('unable to infer image type for key ' + srcKey);
  //   return;
  // }
  // var imageType = typeMatch[1];
  // if (imageType != "jpg" && imageType != "png") {
  //   console.log('skipping non-image ' + srcKey);
  //   return;
  // }

  // function check(number, next) {
  //   console.log(number)
  //   console.log(next)
  //   return true
  // }

  // // Download the image from S3, transform, and upload to a different S3 bucket.
  // async.waterfall([
  //   function checks(next) {
  //       console.log('****************')
  //       console.log('checks')
  //       console.log(dstBucket)
  //       next(null, '973477067')
  //     },
  //   function transform(number, next) {
  //     console.log('****************')
  //     console.log("transform")
  //     console.log(number)

  //     next(null)

  //   }], function (err) {
  //        if (err) {
  //           msg = 'Fail'
  //        } else {
  //           msg = 'Success'
  //        }
  //        context.done(err, msg);
  //   }
  // );



};
