# 🚀 Golang Lambda Application Template

This is template project showcasing an easy-to-develop and easy-to-deploy application using Go & AWS Lambda.

To use this example, you can directly copy and paste the code here into your own repo and deploy.

## Development

This project utilizes Go, the Gin HTTP Package, makefiles, and GitHub Actions for seamless deployment to AWS Lambda.

For local development, a Dockerfile is provided, enabling you to build and run the application effortlessly.

## Infrastructure

The `./infra` directory contains Terraform files that can be utilized to create AWS resources using Infrastructure 
as Code (IaC) with Terraform. The deployment process leverages Terraform Cloud.

### Deploying with Terraform and Github Actions

To deploy this application, you first need an AWS Account. Then, you need to create an IAM user with an AWS Access Key and Secret Access Key 
which has the appropriate IAM policies attached to it (I would personally start with Admin privs and work my way down).

Then, simply take your access and secret access key and export them as variables: 

```sh
export AWS_ACCESS_KEY_ID="your_access_key"
export AWS_SECRET_ACCESS_KEY="your_secret_key"
```

Next, run `terraform plan && terraform apply` and it should create an example lambda function in your aws environment that corresponds to this.

Finally, if you want to have deploys when you merge code using Github Actions, go into ./.github/workflows/deploy.yml
and replace the AWS Access and Secret access key values with your credentials. I strongly suggest doing so with Github Secrets.

### Datastore

(Note: The information about the datastore is not provided in the original text.)

## Build Locally

To build the application locally, run the following command:

```text
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

## Test Locally

To test the application locally, use the following command:

```text
_LAMBDA_SERVER_PORT=2000 AWS_LAMBDA_RUNTIME_API=go1.x go run main.go
```

## API

The application provides a endpoint that accepts a simple GET request.

## Project Structure

Sure! Here's a revised explanation of each component in the project structure:

- `.github/`: This directory contains configuration files related to GitHub, such as workflows, actions, and issue templates. It helps automate various tasks and streamline the development workflow.

- `config/`: This directory typically holds configuration files for the application, including environment-specific settings, database configurations, or any other parameters that can be customized.

- `handlers/`: This directory contains the handler functions responsible for processing incoming requests and generating appropriate responses. Each handler function represents a specific endpoint or API route and implements the business logic associated with it.

- `infra/`: The `infra` directory houses Terraform files that define the infrastructure required by the application. These files can be used with Terraform to provision and manage AWS resources, such as Lambda functions, API Gateways, databases, or any other cloud infrastructure components.

- `models/`: This directory typically includes data models and structures used within the application. It may contain Go struct definitions that represent entities, database schemas, or data transfer objects (DTOs) used for input/output operations.

- `CODEOWNERS`: This file specifies the individuals or teams who are responsible for maintaining and reviewing the code in this repository. It helps ensure that changes are reviewed by the designated owners before merging them into the main branch.

- `Dockerfile`: The Dockerfile is used to build a Docker image of the application. It defines the environment and dependencies required for running the application within a containerized environment.

- `go.mod` and `go.sum`: These files are part of Go's module system and manage the project's dependencies. They specify the required external packages and their versions, allowing for reproducible builds.

- `main.go`: This is the entry point of the application. It typically initializes the necessary components, sets up the server, registers the API routes, and starts the application's execution.

- `Makefile`: The Makefile contains a set of predefined commands and rules that make it easier to build, test, and manage the application. It simplifies repetitive tasks and provides a standardized way of interacting with the project.

- `README.md`: The README file provides essential information and instructions about the project. It typically includes an overview, setup instructions, usage examples, and any other relevant details to help developers understand and contribute to the project.
