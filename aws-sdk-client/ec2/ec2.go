package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Sdk struct {
	ec2 *ec2.EC2
}

type EC2Tags = map[string]string

func New(region string) (Sdk, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})

	if err != nil {
		return Sdk{}, err
	}

	return Sdk{ec2: ec2.New(sess)}, nil
}
