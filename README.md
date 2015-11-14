# doorman
Answer twilio, open the door.

```bash
 → aws lambda update-function-code --function-name "doorman" --zip-file "fileb://dist/doorman_1-0-0_2015-10-9-22-5-55.zip" --debug
```


https://github.com/Tim-B/grunt-aws-lambda

Use arn, don't use function name. Function name is deprecated

```bash
$ npm install
$ grunt lambda_invoke
$ grunt lambda_package
$ grunt deploy --verbose
```


https://github.com/caolan/async


STREAM_NAME=$(echo `aws logs describe-log-streams --log-group-name "/aws/lambda/doorman"` \
| jq '.logStreams | .[0].logStreamName' | sed 's/"//g') \
&& aws logs get-log-events --start-from-head \
--log-group-name "/aws/lambda/doorman" \
--log-stream-name $STREAM_NAME | jq '. | .events | .[] | .message'
