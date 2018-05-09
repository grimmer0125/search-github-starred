# search-github-starred
Full-Text Search the readme, description, homepage and URL of your GitHub starred repository.

This is the missing function on GitHub. GitHub site only supplies the function to search repo-descripton and also only exact phrase match. This site supports both types. Type "A B" for phrase A B and type A B for sequence not important case.

Please try: https://searchgithub.herokuapp.com.

It uses OAuth 2, React, Redux, Golang (server side), Elasticsearch, Redis and so on. Will open source later.

### Local Development
1. npm install
2. install go extension of Visual Studio Code.
3. change the necessary fields in .vscode/launch.json (YOUR_ fields), example:
    ~~~ javascript
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch",
                "type": "go",
                "request": "launch",
                "mode": "debug",
                "remotePath": "",
                "preLaunchTask": "buildclient",
                "port": 2345,
                "host": "127.0.0.1",
                "program": "${workspaceRoot}",
                "env": {
                    "GITHUB_CLIENT_ID" : "YOUR_GITHUB_CLIENT_ID", //ouath of your github app
                    "GITHUB_CLIENT_SECRET": "YOUR_GITHUB_CLIENT_SECRET",  //ouath of your github app
                    "CallbackURL": "http://YOUR_SITE_ADDRESS/auth/callback/",
                    "AWS_ACCESS_KEY_ID": "YOUR_AWS_ACCESS_KEY_ID",  // elasticserach of aws
                    "AWS_SECRET_ACCESS_KEY": "YOUR_AWS_SECRET_ACCESS_KEY", // elasticserach of aws
                    "REDIS_URL": "YOUR_REDIS_URL" //setup your heroku redis or other service's redis
                },
                "args": []
            }
        ]
    }
    ~~~

4. use Visual Studio Code to launch the server.
5. open localhost:5000.

`YOUR_REDIS_URL` could be `redis://localhost:6379` or ` redis://h:YOUR_REDIS_PWD@REDIS_ADDRESS:PORT`. You can use `docker run -p 6379:6379 --name some-redis -d redis` to run a local dockerized Redis.

### Set up AWS' Elasticserach

The keypoint is to create a AMI role to have a permission to write to AWS' Elasticsearch, then only give anonymous users the read permission.

#### Set up an AWS AMI

This project uses [Resource-based Policies](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-ac.html#es-ac-types-resource), so it does not need to add any permission to this AMI user.

#### Set up Elasticsearch

Launch a Elasticsearch service and choose version 2.3. The current go API is not updated to the latest Elasticsearch version yet.

Then create a new Domain, e.g. `searchgithub`. Then modify the access policy like this,

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "es:*",
      "Resource": “Domain ARN/githubrepos/_search"
    },
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "User ARN"
      },
      "Action": "es:*",
      "Resource": "Domain ARN/*"
    }
  ]
}
```

`githubrepos` is the fixed Elasticsearch `index` in this project. The first statement is let browser have read permission, and the the second is to let the server have the write permission if it has AWS access key of the AMI user.

The account name of each user will be used as the `type` of Elasticsearch.

#### Modify the Elasticsearch setting in indexAPI.go and repos.js

indexAPI.go:
```
awsURL = "AWS_ELASTICSEARCH_DOMAIN_ENDPOINT"
```

repos.js:
```
const client = new elasticsearch.Client({
  host: 'AWS_ELASTICSEARCH_DOMAIN_ENDPOINT/githubrepos',
});
```

`AWS_ELASTICSEARCH_DOMAIN_ENDPOINT` could be found out in AWS dashboard. E.g. `https://search-searchgithub-XXXXXXXXXXXXXXXXXXXXXXXXXX.us-west-2.es.amazonaws.com`

**use locally dockerized Elasticsearch**

```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" grimmer0125/elasticsearch:2.3
```

Then use `http://localhost:9200` as the above `AWS_ELASTICSEARCH_DOMAIN_ENDPOINT` in the codes. 

### Deployment on Heroku

Heroku' Redis add-on will automatically create the REDIS_URL as the environment config variable, shown in the dashboard setting page. The other variables needed to be added in the Heroku setting page. https://devcenter.heroku.com/articles/heroku-redis#configuring-your-instance indicates that its REDIS_URL may change at any time.

### Referenced repository
The parameters and the flow about github api calls are from https://github.com/mjmsmith/starredsearch, which is a excellent project and uses Swift on server side to implement the function seraching the information on starred repositories. This repo is based on that repository and add the feature, fulll-text (elasticserach).  
