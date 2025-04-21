# Using Go in AWS Lambda

A simple guide for packaging and running Go applications in AWS Lambda using Docker and the al2023 runtime.

## Getting Started

#### 📂 Clone the repository

```sh
git clone https://github.com/georgelopez7/aws-lambda-go-tutorial.git
cd aws-lambda-go-tutorial
```

#### 🔨 Build the project

**_NOTE:_** Ensure you have `Docker` running

```sh
sam build
```

#### 🚀 Start the project

```sh
sam local start-api
```

This will start the project at `localhost:3000`

#### 🛰 Send HTTP Request

Send `POST` request to `http://localhost:3000/submit` with the following **request body:**

```json
{ "name": "George" }
```

**Example resonse:**

```json
{
  "message": "Hello, George"
}
```

## How It Works

#### 🐳 Running Go in AWS Lambda

AWS now uses the `al2023` runtime for Go, replacing `go1.x`. This project demonstrates how to package your Go Lambda function using a Dockerfile for smooth deployment.

## Dockerfile Overview

The Dockerfile uses a multi-stage build:

```Docker
# ---- BUILD IMAGE FOR LAMBDA -----
FROM golang:1.22 as build

WORKDIR /lambda

COPY go.mod go.sum ./

COPY . .

RUN go build -tags lambda.norpc -o main main.go

# ---- RUNTIME IMAGE FOR LAMBDA ---
FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /lambda/main ./main

ENTRYPOINT [ "./main" ]
```

**Build Stage:** Uses `golang:1.22` to compile the project.

**Final Stage:** Uses `al2023` to create a _Lambda-compatible_ image.

## Cloudformation (template.yml)

To deploy using a Docker image,we define the Lambda function like the following inside the `template.yml`:

```yaml
LambdaGoFunction:
  Type: AWS::Serverless::Function
  Properties:
    PackageType: Image
    Events:
      PostRequest:
        Type: Api
        Properties:
          Path: /submit
          Method: post
          RestApiId: !Ref APIGateway
          Auth:
            Authorizer: LambdaAuthorizer
  Metadata:
    Dockerfile: Dockerfile
    DockerContext: ./lambda
```

- `PackageType: Image` tells AWS to use a Dockerfile.

- `Metadata` specifies the Dockerfile's path.

## Makefile

Inside this repo, there is also a `Makefile` to help build, test and deploy our Lambda function.

- `lambda-build` - builds the Lambda function (creates the `.aws-sam` folder)

- `lambda-start-api` - runs the Lambda function locally using Docker (`localhost:3000`)

- `lambda-deploy-guided` - deploys the Lambda to AWS

## Deploying to AWS

Once you're ready to **deploy**, simply run:

```sh
sam build
sam deploy --guided
```

This will package and deploy the function to AWS.

**_NOTE:_** This will also create an ECR (Elastic Container Registry) repository to hold all the Docker images for this Lambda function.

After deployment, you'll have an **API Gateway endpoint** exposed which will run the Lambda function!
