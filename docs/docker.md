# DOCKER DEPLOY

`docker build -t smart-homev2 .`

`heroku container:login`

`heroku container:push web -a smart-homev2`

`heroku container:release web -a smart-homev2`
