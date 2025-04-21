# Used to build the AWS Lambda function (creates a .aws-sam folder)
lambda-build:
	sam build

# Used to run the AWS Lambda function locally, creating an API that can be called
lambda-start-api:
	sam local start-api

# Used to deploy the AWS Lambda function to AWS
lambda-deploy-guided:
	sam deploy --guided