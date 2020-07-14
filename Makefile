.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/kyler-bot functions/kyler-bot/*.go

# Install relevant packages for deployments via Serverless.
sls-add:
	yarn add serverless serverless-domain-manager

# Remove relevant packages for deployments via Serverless.
sls-remove:
	rm -rf node_modules package.json package.lock yarn.lock

clean:
	rm -rf ./bin ./vendor *.lock

# Sets the Telegram bot post webhook using the variables in local .env file.
webhook:
	sh ./scripts/set-webhook.sh

# Deploy a clean build to AWS Lambda via Serverless.
deploy: clean build
	sls deploy --verbose
