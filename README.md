# search-github-starred
Full-Text Search the readme, description, homepage and URL of your GitHub starred repository. 

This is the missing function on GitHub. GitHub site only supplies the function to search repo-descripton and also only exact phrase match. This site supports both types. Type "A B" for phrase A B and type A B for sequence not important case.

Please try: https://searchgithub.herokuapp.com.

It uses OAuth 2, React, Redux, Golang (server side), Elasticsearch, Redis and so on. Will open source later.

## Development locally 
1. npm install
2. install go extension of Visual Studio Code. 
3. change the necessary fields in .vscode/launch.json (YOUR_*** fields)
example: 
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
                "GITHUB_CLIENT_ID" : "YOUR_GITHUB_CLIENT_ID",
                "GITHUB_CLIENT_SECRET": "YOUR_GITHUB_CLIENT_SECRET",
                "CallbackURL": "http://localhost:5000/auth/callback/",
                "AWS_ACCESS_KEY_ID": "YOUR_AWS_ACCESS_KEY_ID",
                "AWS_SECRET_ACCESS_KEY": "YOUR_AWS_SECRET_ACCESS_KEY",
                "REDIS_URL": "YOUR_REDIS_URL"
            },
            "args": []
        }
    ]
}
