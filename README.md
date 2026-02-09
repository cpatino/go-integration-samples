# go-integration-samples
Sample Golang Projects to use as playbooks for integrations with other services

The AWS SDK for Go v2 uses Go modules. Initialize your local project by running the following Go command.

            go mod init consumer

Next, execute  the go get command that will retrieve the core SDK module, and the config module.

            go get github.com/aws/aws-sdk-go-v2

            go get github.com/aws/aws-sdk-go-v2/config

And finally, execute the following command to retrieve the module for the SQS service.

            go get github.com/aws/aws-sdk-go-v2/service/sqs

