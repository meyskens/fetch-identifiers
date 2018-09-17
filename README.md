Fetch Identifiers
==============

Just like [fetch-comments](https://github.com/meyskens/fetch-comments) but somebody who is very friendly wanted it for identifiers

## How to use
To use bblfsh you will need a server set up using Docker.  
`docker run -d --name bblfshd --privileged -p 9432:9432 -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd`  
You will also need to have drivers for the languages you need to use. Since we want diverse data I suggest just installing all using:  
`docker exec -it bblfshd bblfshctl driver install --all`  

Now that bblfshd is running you can start running this repo. 

Currently the repos to analyse are listed in `main.go`
```
var repos = map[string]string{
	// "file name": "git url"
	"freeCodeCamp.coment": "https://github.com/freeCodeCamp/freeCodeCamp",
	"vue.coment":          "https://github.com/vuejs/vue",
	"springboot.coment":   "https://github.com/spring-projects/spring-boot",
	"moby.coment":         "https://github.com/moby/moby",
}
```
(This should be improved)

## Provided comment lists
This projects analyses a few popular open source projects and uploads the artifacts to S3 using Travis CI (thank you for the free build time!). You can find these at https://s3.eu-west-3.amazonaws.com/fetch-identifiers/