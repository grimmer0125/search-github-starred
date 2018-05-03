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
                    "CallbackURL": "http://YOUR_SITE_URL/auth/callback/",
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

### Set up AWS' Elasticserach

The keypoint is to create a AMI role to have a permission to write to AWS' Elasticsearch, then only give anonymous users the read permission. 

#### Set up an AMI to have the permission, CloudSearch Full access, at least 

#### Set up Elasticsearch 

Create a new Domain, e.g. searchgithub. Then modify the access policy like this, 

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

### Deployment on Heroku

Heroku' Redis add-on will automatically create the REDIS_URL as the environment config variable, shown in the dashboard setting page. The other variables needed to be added in the Heroku setting page. https://devcenter.heroku.com/articles/heroku-redis#configuring-your-instance indicates that its REDIS_URL may change at any time.

### Referenced repository 
The parameters and the flow about github api calls are from https://github.com/mjmsmith/starredsearch, which is a excellent project and uses Swift on server side to implement the function seraching the information on starred repositories. This repo is based on that repository and add the feature, fulll-text (elasticserach).  
