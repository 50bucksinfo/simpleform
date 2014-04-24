Simple Form
===========

##Setup
To set up simpleform you need go and postgresql. Run the following commands to
set it up.

~~~sh
go get github.com/minhajuddin/simpleform
cd $GOPATH/src/github.com/minhajuddin/simpleform
./setup.sh
go build
#tweak config.json as needed and then start
#simpleform by running the following
./simpleform
~~~

#This code is not ready for primetime :)

Code which drives http://getsimpleform.com/

Go to http://getsimpleform.com/downloads for more information

#TODO move this readme to a getsimpleform

## Done
  - Embedding of a form
  - Listing of messages

## Todo
  - Email notifications
  - Redirect to
  - Name of a form
  - Ability to upload files
  - Spam prevention
  - HTTP trigger?
  - Page to host forms for each form
  - Ajax form submissions
  - Access messages using JSON
  - Dont store messages and just email?
  - Validation


## To blog about
  - Form name
  - Ability to setup an email handler e.g. /messages?email=foo@bar.com
  - Send emails via httppost
  - Send requests without referrer to spam


## LICENSE
  - Check the [LICENSE] file
