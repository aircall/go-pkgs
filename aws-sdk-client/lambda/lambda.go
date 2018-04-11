package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Sdk struct {
	lambda *lambda.Lambda
}

func New(region string) (Sdk, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})

	if err != nil {
		return Sdk{}, err
	}

	return Sdk{lambda: lambda.New(sess)}, nil
}

// Invoke awaits for the invoked lambda response and returns the result.
func (cli Sdk) Invoke(functionName string, payload []byte) ([]byte, error) {
	result, err := cli.lambda.Invoke(
		&lambda.InvokeInput{
			FunctionName:   aws.String(functionName),
			InvocationType: aws.String("RequestResponse"),
			Payload:        payload,
		},
	)

	return result.Payload, err
}

// InvokeAsync fires the invocation and forgets (asynchronous)
func (cli Sdk) InvokeAsync(functionName string, payload []byte) error {
	_, err := cli.lambda.Invoke(
		&lambda.InvokeInput{
			FunctionName:   aws.String(functionName),
			InvocationType: aws.String("Event"),
			Payload:        payload,
		},
	)
	return err
}
