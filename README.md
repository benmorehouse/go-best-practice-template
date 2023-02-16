# 💸 Budget Service

This is a template for a golang application that can be used and deployed to AWS Lambda from the Monorepo. 

To get started, you can simply copy and paste this directory into a new project folder within the Monorepo.

## Development 

This uses Go, the Gin HTTP Package, makefiles, and github actions to be able to deploy to AWS Lambda directly.

For local development, there's a dockerfile ready to go for the application that will build and run the app.

## Infrastructure

`./infra` is a directory that houses Terraform files, which can then be used to create aws resources using terraform IaC. 
This is deployed using Terraform Cloud.

### Datastore


## Build locally with
```text
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

## Test locally with
```text
_LAMBDA_SERVER_PORT=2000 AWS_LAMBDA_RUNTIME_API=go1.x go run main.go
```


## API  

`/budget` GET request that takes user ID and returns the budget of said user
